package repository

import (
	"database/sql"
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

func AssignUserModuleGroupPermission(username string, moduleGroupLabel string, level int) error {
	db := database.GetDB()

	sqlStatement :=
		`CREATE OR REPLACE FUNCTION alterPermission(usernameI VARCHAR(256), moduleGroupLabelI VARCHAR(256), levelI INT)  RETURNS VOID
			AS $$
				BEGIN
				IF (SELECT COUNT(*) FROM Permission, Modulegroup, users WHERE users.userID = permission.userID AND Permission.ModuleGroupID = Modulegroup.moduleGroupID AND modulegrouplabel=moduleGroupLabelI AND username=usernameI) > 0 THEN
							UPDATE Permission SET PermissionLevel = levelI
							FROM Modulegroup, Users
							WHERE users.userID = permission.userID 
								AND Permission.ModuleGroupID = Modulegroup.moduleGroupID
								AND modulegrouplabel=moduleGroupLabelI 
								AND username=usernameI;
						ELSE
						INSERT INTO Permission (UserID,ModuleGroupID,Permissionlevel) 
						SELECT userid, modulegroupID, levelI
						FROM users, Modulegroup 
						WHERE modulegrouplabel=moduleGroupLabelI 
							AND username=usernameI
						;
						
					END IF;
				END;
			$$ LANGUAGE plpgsql;`

	_, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Query(`SELECT alterPermission($1, $2, $3)`, username, moduleGroupLabel, level)
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
