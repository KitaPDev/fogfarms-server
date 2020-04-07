package repository

import (
	"github.com/KitaPDev/fogfarms-server/models"
	"github.com/KitaPDev/fogfarms-server/src/database"
	"github.com/jmoiron/sqlx"
)

func GetAllPermissions() []models.Permission {
	db := database.GetDB()

	sqlStatement := `SELECT * FROM Permission`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		permission := models.Permission{}

		err := rows.Scan(
			&permission.PermissionID,
			&permission.UserID,
			&permission.ModuleGroupID,
			&permission.PermissionLevel,
		)
		if err != nil {
			panic(err)
		}

		permissions = append(permissions, permission)
	}

	return permissions
}

func AssignUserModuleGroupPermission(userID int, moduleGroupID int, level int) error {
	db := database.GetDB()

	sqlStatement :=
		`DO $$
			BEGIN
				IF (SELECT COUNT(*) FROM Permission WHERE UserID = $2 AND ModuleGroupID = $3) > 0 THEN
					UPDATE Permission
				    SET PermissionLevel = $1
				    WHERE
				    	UserID = $2 AND ModuleGroupID = $3;
				ELSE
					INSERT INTO Permission (PermissionLevel, UserID, ModuleGroupID)
					VALUES ($1, $2, $3);
			END IF;
		END $$;`

	_, err := db.Query(sqlStatement, level, userID, moduleGroupID)
	if err != nil {
		return err
	}

	return nil
}

func GetSupervisorModuleGroups(userID int) []models.ModuleGroup {
	db := database.GetDB()

	rows, err := db.Query("SELECT ModuleGroupID, PermissionLevel FROM Permission WHERE UserID = ?", userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var supervisorModuleGroupIDs []int
	for rows.Next() {
		var moduleGroupID int
		var permissionLevel int

		err := rows.Scan(
			&moduleGroupID,
			&permissionLevel,
		)
		if err != nil {
			panic(err)
		}

		if permissionLevel == 3 {
			supervisorModuleGroupIDs = append(supervisorModuleGroupIDs, moduleGroupID)
		}
	}

	query, _, err := sqlx.In("SELECT * FROM ModuleGroup WHERE ModuleGroupID IN (?)",
		supervisorModuleGroupIDs)
	if err != nil {
		panic(err)
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	rows, err = db.Query(query)
	if err != nil {
		panic(err)
	}

	var moduleGroups []models.ModuleGroup
	for rows.Next() {
		moduleGroup := models.ModuleGroup{}

		err := rows.Scan(
			moduleGroup.ModuleGroupID,
			moduleGroup.ModuleGroupLabel,
			moduleGroup.PlantID,
			moduleGroup.OnAuto,
			moduleGroup.TDS,
			moduleGroup.PH,
			moduleGroup.Humidity,
			moduleGroup.LightsOnHour,
			moduleGroup.LightsOffHour,
		)
		if err != nil {
			panic(err)
		}

		moduleGroups = append(moduleGroups, moduleGroup)
	}

	return moduleGroups
}