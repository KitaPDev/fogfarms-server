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
		var id string
		var username string
		var hash string
		var salt string
		var isAdmin bool
		err := rows.Scan(&id, &username, &hash, &salt, &isAdmin)
		if err != nil {
			panic(err)
		}

		user := models.User{
			UserID:          id,
			Username:        username,
			Salt:            salt,
			Hash:            hash,
			IsAdministrator: isAdmin,
			CreatedAt:       time.Time{},
		}

		users = append(users, user)
	}

	return users
}
