package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

// Initialize session store with a secret key.
// Replace "your-32-byte-secret-key-here" with a strong, randomly generated key.
var store = sessions.NewCookieStore([]byte("your-32-byte-secret-key-here"))

func init() {
	// Configure session options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
		Secure:   true, // Set to false for development without HTTPS
		SameSite: http.SameSiteLaxMode,
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// In a real application, validate credentials against a database
	if username == "user" && password == "password" {
		session, err := store.Get(r, "user-session")
		if err != nil {
			http.Error(w, "Session error", http.StatusInternalServerError)
			return
		}

		// Set session values
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Values["login_time"] = time.Now()

		// Save session
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Could not save session", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-session")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Check if the user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := session.Values["username"].(string)
	loginTime := session.Values["login_time"].(time.Time)

	fmt.Fprintf(w, "Welcome, %s! You logged in at %s.", username, loginTime.Format(time.RFC822))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "user-session")
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}

	// Invalidate the session
	session.Options.MaxAge = -1 // Deletes the cookie

	if err := session.Save(r, w); err != nil {
		http.Error(w, "Could not save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprint(w, `
				<form action="/login" method="post">
					Username: <input type="text" name="username"><br>
					Password: <input type="password" name="password"><br>
					<input type="submit" value="Login">
				</form>
			`)
		} else {
			loginHandler(w, r)
		}
	})
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/logout", logoutHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
