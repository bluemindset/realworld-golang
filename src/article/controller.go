package article

import "database/sql"

type ArticleContoller struct {
	service ArticleService
	conn    sql.Conn
}
