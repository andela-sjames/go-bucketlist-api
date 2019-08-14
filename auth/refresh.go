package auth

import (
	"net/http"
	"time"
)

// Refresh functioned defined to handle renewal of expired token
func Refresh(w http.ResponseWriter, r *http.Request) {

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
