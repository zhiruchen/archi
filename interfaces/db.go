package interfaces

import (
	"database/sql"
	"time"

	"github.com/zhiruchen/archi/domain"
	infra "github.com/zhiruchen/archi/infrastructure"
)

type DBer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DbRepo struct {
	dbHandler DBer
}

type DbQuestion DbRepo

func (repo *DbQuestion) Create(q *domain.Question) (*domain.Question, error) {
	_sql := "insert into question(id, user_id, title, content) values(?, ?, ?, ?)"

	q.ID = infra.GenID()
	_, err := repo.dbHandler.Exec(_sql, q.ID, q.UserID, q.Title, q.Content)
	if err != nil {
		return nil, err
	}

	return &domain.Question{
		ID:          q.ID,
		UserID:      q.UserID,
		Title:       q.Title,
		Content:     q.Content,
		CreateTime:  time.Now().Unix(),
		FollowCount: 0,
	}, nil
}

func (repo *DbQuestion) Update(q *domain.Question) error {
	_sql := "update question set title=?, content=? where id=? and user_id=?"

	_, err := repo.dbHandler.Exec(_sql, q.Title, q.Content, q.ID, q.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DbQuestion) GetByID(id string) (*domain.Question, error) {
	_sql := "select id, user_id, title, content, unix_timestamp(create_time), follow_count from question where id=?"

	q := &domain.Question{}
	err := repo.dbHandler.QueryRow(_sql, id).Scan(&q.ID, &q.UserID, &q.Title, &q.Content, &q.CreateTime, &q.FollowCount)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (repo *DbQuestion) GetQuestionList(start, offset int32) ([]*domain.Question, bool, error) {
	_sql := "select count(id) from question"
	var count int32

	err := repo.dbHandler.QueryRow(_sql).Scan(&count)

	if err != nil {
		return nil, false, err
	}

	if count == 0 {
		return []*domain.Question{}, false, nil
	}

	var hasNext = true
	if count < offset {
		hasNext = false
	}

	_sql = "select id, user_id, title, content, unix_timestamp(create_time), follow_count from question limit ?, ?"
	rows, err1 := repo.dbHandler.Query(_sql, start, offset)
	if err1 != nil {
		return nil, false, err1
	}
	defer rows.Close()

	var ql = []*domain.Question{}
	for rows.Next() {
		var q = &domain.Question{}
		err = rows.Scan(&q.ID, &q.UserID, &q.Title, &q.Content, &q.CreateTime, &q.FollowCount)
		if err != nil {
			return nil, false, err
		}

		ql = append(ql, q)
	}

	err = rows.Err()
	if err != nil {
		return nil, false, err
	}

	return ql, hasNext, nil
}
