package domain

type QuestionStore interface {
	Create(q *Question) (*Question, error)
	Update(q *Question) error
	GetByID(id string) (*Question, error)
	GetQuestionList(start, offset int32) ([]*Question, bool, error)
}

type Question struct {
	ID          string
	UserID      string
	Title       string
	Content     string
	CreateTime  int64
	FollowCount int64
}
