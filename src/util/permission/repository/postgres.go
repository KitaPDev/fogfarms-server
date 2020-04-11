package repository

import (
	"database/sql"
	"github.com/KitaPDev/fogfarms-server/src/util/user"
	"log"

	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
)

func GetAllPermissions() ([]models.Permission, error) {
	db := database.GetDB()

	sqlStatement := `SELECT * FROM Permission;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		permission := models.Permission{}

		err := rows.Scan(
			&permission.UserID,
			&permission.ModuleGroupID,
			&permission.PermissionLevel,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	return permissions, nil
}

func AssignUserModuleGroupPermission(userID int, moduleGroupID int, level int) error {
	db := database.GetDB()

	sqlStatement :=
		`CREATE OR REPLACE FUNCTION alterPermission(userIDI INT, moduleGroupIDI INT, levelI INT)  RETURNS VOID
			AS $$
				BEGIN
				IF (SELECT COUNT(*) FROM Permission WHERE UserID = userIDI AND ModuleGroupID = moduleGroupIDI) > 0 THEN
							UPDATE Permission SET PermissionLevel = levelI
							WHERE
								UserID = userIDI AND ModuleGroupID = moduleGroupIDI;
						ELSE
							INSERT INTO Permission (PermissionLevel, UserID, ModuleGroupID)
							VALUES (levelI, userIDI, moduleGroupIDI);
					END IF;
				END;
			$$ LANGUAGE plpgsql;`

	_, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Query(`SELECT alterPermission($1, $2, $3)`, userID, moduleGroupID, level)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetAssignedModuleGroupsWithPermissionLevel(userID int, permissionLevel int) (map[models.ModuleGroup]int, error) {
	db := database.GetDB()

	sqlStatement :=
		`SELECT m.ModuleGroupID, m.ModuleGroupID, m.PlantID, m.LocationID, m.Param_TDS, m.Param_PH, m.Param_Humidity,
       m.onAuto, m.LightsOffHour, m.LightsOnHour, p.PermissionLevel
		FROM ModuleGroup m, Permission p 
		WHERE p.UserID = $1 AND m.ModuleGroupID = p.ModuleGroupID`

	sqlStatementPermissionLevel := ` AND p.PermissionLevel = $2`

	var rows *sql.Rows
	var err error
	if permissionLevel != -1 {
		rows, err = db.Query(sqlStatement+sqlStatementPermissionLevel, userID, permissionLevel)
	} else {
		rows, err = db.Query(sqlStatement, userID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapModuleGroupPermissionLevel := make(map[models.ModuleGroup]int)
	for rows.Next() {
		moduleGroup := models.ModuleGroup{}
		var permissionLevel int

		err := rows.Scan(
			&moduleGroup.ModuleGroupID,
			&moduleGroup.ModuleGroupLabel,
			&moduleGroup.PlantID,
			&moduleGroup.LocationID,
			&moduleGroup.TDS,
			&moduleGroup.PH,
			&moduleGroup.Humidity,
			&moduleGroup.OnAuto,
			&moduleGroup.LightsOffHour,
			&moduleGroup.LightsOnHour,
			&permissionLevel,
		)
		if err != nil {
			return nil, err
		}

		mapModuleGroupPermissionLevel[moduleGroup] = permissionLevel
	}

	return mapModuleGroupPermissionLevel, nil
}

func PopulateUserManagementPage(u *models.User) (map[string]map[string]int, error) {
	db := database.GetDB()

	users, err := user.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var sqlStatement string
	var rows *sql.Rows

	if u.IsAdministrator {
		sqlStatement = `SELECT DISTINCT ModuleGroupID FROM ModuleGroup`
		rows, err = db.Query(sqlStatement)

	} else {
		sqlStatement =
			`SELECT DISTINCT ModuleGroupID 
			FROM ModuleGroup, Permission
			WHERE ModuleGroup.ModuleGroupID = Permission.ModuleGroupID
			AND UserID = $1 AND PermissionLevel = 3;`
		rows, err = db.Query(sqlStatement, u.UserID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moduleGroupLabels []string
	for rows.Next() {
		var moduleGroupLabel string

		err := rows.Scan(
			&moduleGroupLabel,
		)
		if err != nil {
			return nil, err
		}

		moduleGroupLabels = append(moduleGroupLabels, moduleGroupLabel)

	}

	var mapUsername = make(map[string]map[string]int)
	for i := range users {
		var mapModuleGroupLabels = make(map[string]int)

		for i := range moduleGroupLabels {
			mapModuleGroupLabels[moduleGroupLabels[i]] = 0
		}

		if users[i].Username != u.Username {
			mapUsername[users[i].Username] = mapModuleGroupLabels
		}
	}

	if u.IsAdministrator {
		sqlStatement =
			`SELECT Username, PermissionLevel, ModuleGroupLabel 
			FROM Permission, ModuleGroup, Users
			WHERE Users.UserID = Permission.UserID
			AND USERS.UserID != $1 AND ModuleGroup.ModuleGroupID = Permission.ModuleGroupID;`

	} else {
		sqlStatement =
			`SELECT Username, PermissionLevel, ModuleGroupLabel
			FROM Permission, ModuleGroup, Users
			WHERE Users.UserID = Permission.UserID
			AND Users.UserID != $1 AND ModuleGroup.ModuleGroupID = Permission.ModuleGroupID
			AND ModuleGroup.ModuleGroupID
			IN ( SELECT ModuleGroupID from Permission WHERE UserID = $1 AND PermissionLevel = 3 );`

	}

	rows, err = db.Query(sqlStatement, u.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		var permissionLevel int
		var moduleGroupLabel string

		err := rows.Scan(
			&username,
			&permissionLevel,
			&moduleGroupLabel,
		)
		if err != nil {
			return nil, err
		}

		mapUsername[username][moduleGroupLabel] = permissionLevel

	}
	return mapUsername, err
}