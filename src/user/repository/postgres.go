package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"log"
	"time"
)

func GetAllUsers() []models.User {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM Users;")
	if err != nil {
		panic(err)
	}
	defer log.Fatal(rows.Close())

	var users []models.User
	for rows.Next() {
		var userID string
		var username string
		var hash string
		var salt string
		var isAdmin bool
		var createdAt time.Time
		err := rows.Scan(&userID, &username, &hash, &salt, &isAdmin, &createdAt)
		if err != nil {
			panic(err)
		}

		user := models.User{
			UserID:          userID,
			Username:        username,
			Salt:            salt,
			Hash:            hash,
			IsAdministrator: isAdmin,
			CreatedAt:       createdAt,
		}

		users = append(users, user)
	}

	return users
}
