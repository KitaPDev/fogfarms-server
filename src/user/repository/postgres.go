package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"time"
)

func GetAllUsers() []models.User {
	db := database.GetDB()

	defer db.Close()
	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		panic(err)
	}

	var users []models.User
	for rows.Next() {
		var id string
		var username string
		var hash string
		var salt string
		var r string
		err := rows.Scan(&id, &username, &hash, &salt, &r)
		if err != nil {
			panic(err)
		}

		role := models.AuthorizedUser
		if r == "Administrator" {
			role = models.Administrator
		}

		user := models.User{
			UserID:    id,
			Username:  username,
			Salt:      salt,
			Hash:      hash,
			Role:      role,
			CreatedAt: time.Time{},
		}

		users = append(users, user)
	}

	return users
}