package repositories

import (
	"database/sql"

	"github.com/yutak03/go-http-api-server/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		INSERT INTO comments (article_id, message, created_at) VALUES
		(?, ?, now());
	`

	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	
	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()

	newComment.CommentID = int(id)
	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT *
		FROM comments
		WHERE article_id = ?;
	`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.ArticleID, &comment.CommentID, &comment.Message, &comment.CreatedAt)

		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}