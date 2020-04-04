package user

import "github.com/KitaPDev/fogfarms-server/models"

type Repository interface {
	GetAllUsers() []models.User
	CreateUser(username string, password string, isAdministrator bool)
}