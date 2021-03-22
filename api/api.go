package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/StukaNya/TgCrypter/model/session"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *ServerConfig
	logger *logrus.Logger
	users  session.UserRepository
}

func NewAPIServer(log *logrus.Logger, config *ServerConfig, users session.UserRepository) *APIServer {
	return &APIServer{
		config: config,
		logger: log,
		users:  users,
	}
}

func (s *APIServer) Serve(ctx context.Context) error {
	r := gin.Default()

	r.POST("/user", s.registerUser)
	r.GET("/user/:id", s.fetchUser)

	server := http.Server{
		Addr:    s.config.BindAddr,
		Handler: r,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Error during server listening: ", err)
		}
	}()
	s.logger.Info("Server is listening, URL: ", s.config.BindAddr)

	// Gracefully shutdown HTTP server...
	<-ctx.Done()
	s.logger.Info("Server stopped")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxShutDown); err != nil {
		s.logger.Fatal("Unable to shutdown server: ", err)
	}
	s.logger.Info("Gracefully shutdown server")

	return nil
}

func ServerError(err error) gin.H {
	return gin.H{"error": err}
}

func (s *APIServer) registerUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Register new user (API Mock!), id: ", uuid.NewV4())
}

func (s *APIServer) fetchUser(ctx *gin.Context) {
	userID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ServerError(err))
	}

	userInfo, err := s.users.FetchUser(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ServerError(err))
	}

	userView := gin.H{
		"id":           userID,
		"chatID":       userInfo.ChatID,
		"name":         userInfo.Name,
		"registeredAt": userInfo.RegisteredAt,
	}
	ctx.JSON(http.StatusOK, userView)

	s.logger.Info("[API] fetch user inso, ID: ", userID)
}
