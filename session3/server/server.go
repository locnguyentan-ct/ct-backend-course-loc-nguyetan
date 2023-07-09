package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(authMiddleware)
	e.Use(LogMiddleware)

	// Routes
	e.POST("/api/public/register", register)
	e.POST("/api/public/login", login)
	e.GET("/api/private/self", self, middleware.JWT([]byte("ct-secret-key")), authMiddleware)

	// Start server
	err := e.Start(":8090")
	if err != nil {
		return
	}
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		// Extract the JWT token from the Authorization header
		token, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Validate the JWT token
		username, err := ValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Store the username in the context for further use if needed
		c.Set("username", username)

		// Call the next handler
		return next(c)
	}
}

/*
	httpRequest -> BYTE -> Binding(JSON -> Struct) -> Validate -> Handle Logic -> (Struct -> JSON) -> BYTE -> httpResponse

	Framework: GO Practical (lib, wrapper)

	Echo (simple, easy to use) & Gin (interface more complex)
	-> Echo
*/

/*
	PATH : Handler
	Common prefix:
		api/public
		api/private (Middle authentication)
	Grouping => implement Middle

	* Middleware: Authentication, Logging, Recover (handler unexpected error)
*/

/*
		TODO #2:
		- implement the logic to register a new user (username, password, full_name, address)
	  	- Validate username (not empty and unique)
	  	- Validate password (length should at least 8)
*/
func register(c echo.Context) error {
	req := new(RegisterRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.JSON(http.StatusInternalServerError, "Failed to validate the request")
		}

		//var validationErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "Username":
				//validationErrors = append(validationErrors, "Username is required")
				return c.JSON(http.StatusBadRequest, "Username is required")
			case "Password":
				// If user is not existed, show error validate
				//validationErrors = append(validationErrors, "Password should be at least 8 characters long")
				return c.JSON(http.StatusBadRequest, "Password should be at least 8 characters long")
			default:
				// Validate unique Username
				_, err := userStore.Get(req.Username)
				if err == nil {
					return c.JSON(http.StatusBadRequest, "Username is existed")
				} else if err != ErrUserNotFound {
					return c.JSON(http.StatusInternalServerError, "Failed to check username availability")
				}
			}
		}

		//return c.JSON(http.StatusInternalServerError, validationErrors)
	}

	info := UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}

	if err := userStore.Save(info); err != nil {
		//http.Error(w, "Failed to save user", http.StatusInternalServerError)
		//return
		return c.JSON(http.StatusInternalServerError, "Failed to save user")
	}

	//w.WriteHeader(http.StatusOK)
	//return
	return c.JSON(http.StatusOK, "Registration successful")
}

/*
		TODO #3:
		- implement the logic to login
		- validate the user's credentials (username, password)
	  	- Return JWT token to client
*/
func login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return c.JSON(http.StatusInternalServerError, "Failed to validate the request")
		}

		for _, err := range err.(validator.ValidationErrors) {
			switch err.StructField() {
			case "Username":
				//validationErrors = append(validationErrors, "Username is required")
				return c.JSON(http.StatusBadRequest, "Username are required")
			case "Password":
				// If user is not existed, show error validate
				//validationErrors = append(validationErrors, "Password should be at least 8 characters long")
				return c.JSON(http.StatusBadRequest, "Password are required")
			default:
				// Validate unique Username
				user, err := userStore.Get(req.Username)
				if err != nil {
					return c.JSON(http.StatusUnauthorized, "Invalid username or password")
				}

				if user.Password != req.Password {
					return c.JSON(http.StatusUnauthorized, "Invalid username or password")
				}
			}
		}
	}

	token, err := GenerateToken(req.Username, 24*time.Hour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	resp := LoginResponse{Token: token}

	return c.JSON(http.StatusOK, resp)
}

/*
TODO #4:
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(c echo.Context) error {
	// Retrieve the username from the context
	username := c.Get("username").(string)

	// Retrieve user information from userStore
	user, err := userStore.Get(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
	}

	// Return the user information as JSON
	return c.JSON(http.StatusOK, user)
}

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		// Call the next handler
		err := next(c)

		// Log the request information
		duration := time.Since(startTime)
		log := "Path: " + c.Request().URL.Path + " | Status Code: " + http.StatusText(c.Response().Status) + " | Time Start: " + startTime.String() + " | Duration: " + duration.String()
		c.Logger().Info(log)

		return err
	}
}
