package repository

import (
	"fmt"
	"log"

	"github.com/with-insomnia/Forum-Golang/internal/model"
)

func (p *postQuery) CommentPost(comment model.Comment) error {
	stmt, err := p.db.Prepare("INSERT INTO comments(post_id, user_id,username, message, like, dislike, born) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(comment.PostId, comment.UserId, comment.Username, comment.Message, 0, 0, comment.Born)
	if err != nil {
		return err
	}
	return nil
}

func (p *postQuery) GetAllCommentByPostID(postId int64) ([]model.Comment, error) {
	stmt, err := p.db.Prepare("SELECT comment_id, post_id, user_id, username, message, like, dislike, born FROM comments WHERE post_id = ?")
	if err != nil {
		return nil, err
	}
	row, err := stmt.Query(postId)
	if err != nil {
		log.Println(err)
	}
	var comments []model.Comment
	for row.Next() {
		var comment model.Comment
		if err := row.Scan(&comment.ID, &comment.PostId, &comment.UserId, &comment.Username, &comment.Message, &comment.Like, &comment.Dislike, &comment.Born); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (p *postQuery) GetCommentByCommentID(commentId int64) (model.Comment, error) {
	stmt, err := p.db.Prepare("SELECT comment_id, post_id, user_id, username, message, like, dislike, born FROM comments WHERE comment_id = ?")
	if err != nil {
		log.Println(err)
	}
	row := stmt.QueryRow(commentId)
	var comment model.Comment
	if err := row.Scan(&comment.ID, &comment.PostId, &comment.UserId, &comment.Username, &comment.Message, &comment.Like, &comment.Dislike, &comment.Born); err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}
