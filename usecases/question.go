package usecases

import (
	"github.com/zhiruchen/archi/domain"
)

type QuestionInteractor struct {
	QuestionStore domain.QuestionStore
}

func (qi *QuestionInteractor) Add(q *domain.Question) (*domain.Question, error) {
	return qi.QuestionStore.Create(q)
}

func (qi *QuestionInteractor) Update(id, userID, title, content string) (*domain.Question, error) {
	err := qi.QuestionStore.Update(&domain.Question{
		ID:      id,
		UserID:  userID,
		Title:   title,
		Content: content,
	})

	if err != nil {
		return nil, err
	}

	return qi.QuestionStore.GetByID(id)
}

func (qi *QuestionInteractor) Get(id string) (*domain.Question, error) {
	return qi.QuestionStore.GetByID(id)
}

func (qi *QuestionInteractor) GetQuestionList(previous, count int32) ([]*domain.Question, bool, error) {
	return qi.QuestionStore.GetQuestionList(previous, count)
}
