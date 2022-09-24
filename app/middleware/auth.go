package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// OptionallyRequireAdminAuth wraps a handler requiring HTTP basic auth
// using "uploader" as the username.
// If a password isn't set then auth is skipped.
func OptionallyRequireAdminAuth(handler http.HandlerFunc, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Empty password means auth is not required.
		if password == "" {
			handler(w, r)
			return
		}

		username := "uploader"
		realm := "Tube uploader"

		// The following line is kind of a work around.
		// If you want HTTP Basic Auth + Cors it requires _explicit_ origins to be provided in the
		// Access-Control-Allow-Origin header.  So we just pull out the origin header and specify it.
		// If we want to lock down admin APIs to not be CORS accessible for anywhere, this is where we would do that.
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		// For request needing CORS, send a 204.
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		user, pass, ok := r.BasicAuth()

		// Failed
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Debugln("Failed uploader authentication")
			return
		}

		handler(w, r)
	}
}

func RequireSandstormPermission(handler http.HandlerFunc, permissionNeeded string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// Failed
		if !strings.Contains(r.Header.Get("X-Sandstorm-Permissions"), permissionNeeded) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Debugln("No upload capability granted")
			return
		}
		
		handler(w, r)
	}
}
