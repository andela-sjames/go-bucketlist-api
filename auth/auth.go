package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/andela-sjames/go-bucketlist-api/models"
	"github.com/andela-sjames/go-bucketlist-api/utils"
)

// RequestContextKey defined
type RequestContextKey string

const (
	// CtxKey context defined
	CtxKey RequestContextKey = "userObj"
)

// JWTAuthenticationMiddleware defined to intercept request before passing
// to the next middleware or handler
func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuthRequiredURL := []string{"/", "/api/auth/login", "/api/auth/signup"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                                 //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range noAuthRequiredURL {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("PASSPHRASE")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Println("User: ", tk.UserID, tk.Email) //Useful for monitoring

		userClaimContextValue := make(map[string]interface{})
		userClaimContextValue["userID"] = tk.UserID
		userClaimContextValue["userEmail"] = tk.Email

		reqCtxKey := context.WithValue(r.Context(), CtxKey, userClaimContextValue)
		next.ServeHTTP(w, r.WithContext(reqCtxKey)) //proceed in the middleware chain!
	})
}
