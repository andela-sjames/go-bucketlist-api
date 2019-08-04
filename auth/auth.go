package auth

import (
	"go-bucketlist-api/utils"
	"net/http"
)

// JWTAuthenticationMiddleware defined to intercept request before passing
// to the next middleware or handler
func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		noAuthRequiredURL := []string{"/api/user/new", "/api/user/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                         //current request path

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

	})
}
