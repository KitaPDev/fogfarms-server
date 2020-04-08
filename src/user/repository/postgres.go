package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/lib/pq"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]models.User, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT userid, username,isadministrator,hash,salt,createdat FROM Users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.IsAdministrator,
			&user.Hash,
			&user.Salt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	db := database.GetDB()

	sqlStatement := `SELECT UserID, Username, IsAdministrator, Hash, Salt, CreatedAt FROM Users WHERE Username = $1;`

	rows, err := db.Query(sqlStatement, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.IsAdministrator,
			&user.Hash,
			&user.Salt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

	}

	return &user, nil
}

func GetUserByID(userID int) (*models.User, error) {
	db := database.GetDB()

	sqlStatement := `SELECT UserId, Username, IsAdministrator, Hash, Salt, CreatedAt FROM Users WHERE UserID = $1;`

	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.IsAdministrator,
			&user.Hash,
			&user.Salt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

	}

	return &user, nil
}

func GetUsersByID(userIDs []int) ([]models.User, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT UserID, Username, IsAdministrator, Hash, Salt, CreatedAt 
		FROM Users 
		WHERE UserID = ANY($1);`

	rows, err := db.Query(sqlStatement, pq.Array(userIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.UserID,
			&user.Username,
			&user.IsAdministrator,
			&user.Hash,
			&user.Salt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

func CreateUser(username string, password string, isAdministrator bool) error {
	db := database.GetDB()

	salt := generateSalt()
	hash, err := hash(password, salt)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO Users (Username, IsAdministrator, Hash, Salt, CreatedAt) 
		VALUES ($1, $2, $3, $4, Now());`

	db.QueryRow(sqlStatement, username, isAdministrator, hash, salt)

	return nil
}

func ValidateUserByUsername(username string, inputPassword string) (bool, error) {
	db := database.GetDB()

	sqlStatement := `SELECT UserID, Username, Hash, Salt FROM Users WHERE Username = $1;`

	user := models.User{}

	row := db.QueryRow(sqlStatement, username)

	switch err := row.Scan(
		&user.UserID,
		&user.Username,
		&user.Hash,
		&user.Salt,
	); err {

	case sql.ErrNoRows:
		fmt.Println("No Rows Returned!")

	case nil:
		password := inputPassword + user.Salt
		fmt.Println(username, user.Salt)
		fmt.Printf("this works \n")
		fmt.Printf("%+v , %+v \n", username, password)

		if username == username &&
			bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password)) == nil {
			return true, nil
		}
		return false, nil

	default:
		fmt.Printf("%+v", err)
		return false, nil
	}

	return false, nil
}

func generateSalt() string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	salt := make([]byte, 32)
	for i := range salt {
		salt[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(salt)
}

func hash(password string, salt string) (string, error) {
	s := password + salt
	h, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

func PopulateUserManagementPage(u *models.User) (map[string]map[string]int, error) {
	db := database.GetDB()
	users, err := GetAllUsers()
	if err != nil {
		return nil, err
	}
	var sqlStatement string
	var rows *sql.Rows
	if u.IsAdministrator {
		sqlStatement = "SELECT DISTINCT modulegrouplabel FROM modulegroup"
		rows, err = db.Query(sqlStatement)
	} else {
		sqlStatement = "SELECT DISTINCT modulegrouplabel FROM modulegroup,permission where modulegroup.modulegroupid=permission.modulegroupid AND USERID= $1 AND permissionlevel=3;"
		rows, err = db.Query(sqlStatement, u.UserID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modulegrouplabels []string
	for rows.Next() {
		var modulegrouplabel string

		err := rows.Scan(
			&modulegrouplabel,
		)
		if err != nil {
			return nil, err
		}

		modulegrouplabels = append(modulegrouplabels, modulegrouplabel)

	}

	var usernameMAP = make(map[string]map[string]int)
	for i := range users {
		var modulegrouplabelsMAP = make(map[string]int)
		for i := range modulegrouplabels {
			modulegrouplabelsMAP[modulegrouplabels[i]] = 0
		}
		if users[i].Username != u.Username {
			usernameMAP[users[i].Username] = modulegrouplabelsMAP
		}
	}
	fmt.Printf("%+v", usernameMAP)
	if u.IsAdministrator {
		sqlStatement = "Select username, permissionlevel,modulegrouplabel from permission,modulegroup,users where users.userid=permission.userid AND users.userid!= $1 AND modulegroup.modulegroupid=permission.modulegroupid;"

	} else {
		sqlStatement = "Select username, permissionlevel,modulegrouplabel from permission,modulegroup,users where users.userid=permission.userid AND users.userid!= $1 AND modulegroup.modulegroupid=permission.modulegroupid AND modulegroup.modulegroupid IN ( Select modulegroupid from permission where userid = $1 AND permissionlevel =3 );"

	}
	rows, err = db.Query(sqlStatement, u.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var permissionlevel int
		var modulegrouplabel string
		err := rows.Scan(
			&username,
			&permissionlevel,
			&modulegrouplabel,
		)
		if err != nil {
			return nil, err
		}

		usernameMAP[username][modulegrouplabel] = permissionlevel

	}
	return usernameMAP, err
}
