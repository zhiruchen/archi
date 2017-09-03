package interfaces

import (
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"

	"github.com/zhiruchen/archi/domain"
	pb "github.com/zhiruchen/archi/pb"
)

type QuestionInteractor interface {
	Add(q *domain.Question) (*domain.Question, error)
	Update(id, userID, title, content string) (*domain.Question, error)
	Get(id string) (*domain.Question, error)
	GetQuestionList(previous, count int32) ([]*domain.Question, bool, error)
}

type RPCHandler struct {
	QuestionInteractor QuestionInteractor
}

func (h *RPCHandler) CreateQuestion(ctx context.Context, req *pb.CreateQuestionReq) (*pb.CreateQuestionResp, error) {
	q := &domain.Question{
		UserID:  req.UserID,
		Title:   req.QuestionTitle,
		Content: req.QuestionContent,
	}
	var err error
	q, err = h.QuestionInteractor.Add(q)

	if err != nil {
		return nil, err
	}

	q1 := &pb.Question{}
	copier.Copy(q1, q)

	return &pb.CreateQuestionResp{Quest: q1}, nil
}

func (h *RPCHandler) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionReq) (*pb.UpdateQuestionResp, error) {
	q, err := h.QuestionInteractor.Update(req.ID, req.UserID, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	q1 := &pb.Question{}
	copier.Copy(q1, q)

	return &pb.UpdateQuestionResp{Quest: q1}, nil
}

func (h *RPCHandler) GetQuestion(ctx context.Context, req *pb.GetQuestionReq) (*pb.GetQuestionResp, error) {
	q, err := h.QuestionInteractor.Get(req.QuestionID)
	if err != nil {
		return nil, err
	}

	q1 := &pb.Question{}
	copier.Copy(q1, q)
	return &pb.GetQuestionResp{Quest: q1}, nil
}

func (h *RPCHandler) GetQuestionList(ctx context.Context, req *pb.GetQuestionListReq) (*pb.GetQuestionListResp, error) {
	ql, hasNext, err := h.QuestionInteractor.GetQuestionList(req.PreviousPageNum, req.PerPage)

	if err != nil {
		return nil, err
	}
	ql1 := []*pb.Question{}

	for _, q := range ql {
		q1 := &pb.Question{}

		copier.Copy(q1, q)
		ql1 = append(ql1, q1)
	}

	return &pb.GetQuestionListResp{QList: ql1, HasNext: hasNext, TotalCount: 0}, nil
}
