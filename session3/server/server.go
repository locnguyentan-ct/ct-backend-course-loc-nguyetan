package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/public/register", register)
	http.HandleFunc("/api/public/login", login)
	http.HandleFunc("/api/private/self", self)

	http.HandleFunc("/api/public/log/register", LogWrapper(register))
	http.HandleFunc("/api/public/log/login", LogWrapper(login))
	http.HandleFunc("/api/private/log/self", LogWrapper(self))

	http.ListenAndServe(":8090", nil)
}

/*
		TODO #2:
		- implement the logic to register a new user (username, password, full_name, address)
	  	- Validate username (not empty and unique)
	  	- Validate password (length should at least 8)
*/
func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := ParseJSONBody(r, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Validate username uniqueness (assuming userStore is a global variable)
	_, err := userStore.Get(req.Username)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	} else if err != ErrUserNotFound {
		http.Error(w, "Failed to check username availability", http.StatusInternalServerError)
		return
	}

	if len(req.Password) < 8 {
		http.Error(w, "Password should be at least 8 characters long", http.StatusBadRequest)
		return
	}

	info := UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}

	if err := userStore.Save(info); err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (username, password)
	  	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := ParseJSONBody(r, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	user, err := userStore.Get(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if user.Password != req.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(req.Username, 24*time.Hour)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Token: token}

	WriteJSONResponse(w, resp, http.StatusOK)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

/*
TODO #4:
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	// Extract the JWT token from the Authorization header
	token, err := ExtractTokenFromHeader(authHeader)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Validate the JWT token
	username, err := ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Retrieve user information from userStore
	user, err := userStore.Get(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Return the user information as JSON
	WriteJSONResponse(w, user, http.StatusOK)
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the original handler
		handler(w, r)

		// Log the request information
		duration := time.Since(startTime)
		log := "Path: " + r.URL.Path + " | Status Code: " + http.StatusText(http.StatusOK) + " | Time Start: " + startTime.String() + " | Duration: " + duration.String()
		println(log)
	}
}

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/
