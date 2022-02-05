package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/IBearSmile2319/go-rest-api/internal/comment"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pinter to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - return a pointer to a new Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// LoggingMiddleware - addss middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			},
		).Info("handling request")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handy middleware function that handles basic authentication
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic auth endpoint hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive!"}); err != nil {
			panic(err)
		}
	})
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("secret")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There has been an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

// JWTAuth - a decorator function that adds JWT authentication to our endpoints
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("JWT auth endpoint hit")
		token := r.Header["Authorization"]
		if token == nil {
			sendErrorResponse(w, "not authorized", errors.New("not authorized1"))
			return
		}
		// Bearer jwt-token
		authHeaderParts := strings.Split(token[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "not authorized", errors.New("not authorize2"))
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized3"))
			return
		}
	}
}
