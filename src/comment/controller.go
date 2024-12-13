package comment

import "database/sql"

type CommentContoller struct {
	service CommentService
	conn    sql.Conn
}
