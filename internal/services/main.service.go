package services

import "github.com/chinsiang99/simple-go-project/internal/repositories"

type ServiceManager struct {
	AuthService *AuthService
	// UserService   *UserService
	// TicketService *TicketService
}

func NewServiceManager(repositories *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		AuthService: NewAuthService(repositories),
		// UserService:   NewUserService(repositories),
		// TicketService: NewTicketService(repositories),
	}
}
