package views

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/andela-sjames/go-bucketlist-api/auth"
	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// RefreshHandler functioned defined to handle renewal of close to expire token
func RefreshHandler(w http.ResponseWriter, r *http.Request) {

	userObj := r.Context().Value(auth.CtxKey).(map[string]interface{})
	userID := userObj["userID"].(uint)
	userEmail := userObj["userEmail"].(string)

	tokenHeader := r.Header.Get("Authorization")
	splitted := strings.Split(tokenHeader, " ")
	tokenPart := splitted[1]
	claims := &models.Token{}

	_, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PASSPHRASE")), nil
	})

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 60 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 60*time.Second {
		response := utils.Message(false, "Token refresh only applies to a valid token with less than 60s to expire")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		utils.Respond(w, response)
		return
	}

	// Create new JWT for the current use, with a renewed expiration time
	newClaims := models.GenerateUserClaims(userID, userEmail)

	newToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), newClaims)
	newTokenString, err := newToken.SignedString([]byte(os.Getenv("PASSPHRASE")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := utils.Message(true, "success")
	resp["token"] = newTokenString
	utils.Respond(w, resp)

}
