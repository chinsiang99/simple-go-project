package handlers

import "github.com/chinsiang99/simple-go-project/internal/services"

type HandlerManager struct {
	AuthHandler *AuthHandler
	UserHandler IUserHandler
	// TicketHandler *TicketHandler
}

func NewHandlerManager(services *services.ServiceManager) *HandlerManager {
	return &HandlerManager{
		AuthHandler: NewAuthHandler(services),
		UserHandler: NewUserHandler(services.UserService),
		// TicketHandler: NewTicketHandler(service),
	}
}
