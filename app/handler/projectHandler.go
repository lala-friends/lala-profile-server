package handler

import (
	"goframework/app/cmm"
	"database/sql"
	"goframework/app/util"
	"strings"
	"goframework/app/domain"
	"github.com/go-sql-driver/mysql"
)

func HandleGetProjects(s *cmm.Server, db *sql.DB)  {
	s.HandleFunc("GET", "/developer/:username/projects", func(c *cmm.Context) {
		id := util.GetUserId(db, c.Params["username"].(string))
		rows, err := db.Query(util.SELECT_PROJECTS, id)
		util.HandleSqlErr(err)
		defer rows.Close()

		for rows.Next() {
			var projectName string
			var periodFrom, periodTo mysql.NullTime
			var projectIntroduce, projectDescription, projectTechs, personalRole, projectLink sql.NullString
			err := rows.Scan(&projectName, &periodFrom, &periodTo, &projectIntroduce, &projectDescription, &projectTechs, &personalRole, &projectLink)
			util.HandleSqlErr(err)
			projectTechsArr := strings.Split(projectTechs.String, "\n")
			pjt := domain.Project{projectName, periodFrom.Time, periodTo.Time, projectIntroduce.String, projectDescription.String, projectTechsArr, personalRole.String, projectLink.String}
			c.RenderJson(pjt)
		}
	})
}
