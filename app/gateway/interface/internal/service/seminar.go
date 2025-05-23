package service

import (
	"context"

	v1 "github.com/Fl0rencess720/Ayana/api/gateway/seminar/v1"
	"github.com/Fl0rencess720/Ayana/app/gateway/interface/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type SeminarService struct {
	v1.UnimplementedSeminarServer

	uc *biz.SeminarUsecase
}

func NewSeminarService(uc *biz.SeminarUsecase) *SeminarService {
	return &SeminarService{uc: uc}
}

func (s *SeminarService) CreateTopic(ctx context.Context, req *v1.CreateTopicRequest) (*v1.CreateTopicReply, error) {
	reply, err := s.uc.CreateTopic(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (s *SeminarService) DeleteTopic(ctx context.Context, req *v1.DeleteTopicRequest) (*v1.DeleteTopicReply, error) {
	reply, err := s.uc.DeleteTopic(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (s *SeminarService) GetTopic(ctx context.Context, req *v1.GetTopicRequest) (*v1.GetTopicReply, error) {
	reply, err := s.uc.GetTopic(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func (s *SeminarService) GetTopicsMetadata(ctx context.Context, req *v1.GetTopicsMetadataRequest) (*v1.GetTopicsMetadataReply, error) {
	reply, err := s.uc.GetTopicsMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
func StartTopic(ctx http.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	h := ctx.Middleware(func(c context.Context, req interface{}) (interface{}, error) {
		return biz.StartTopic(ctx, c)
	})
	_, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SeminarService) StopTopic(ctx context.Context, req *v1.StopTopicRequest) (*v1.StopTopicReply, error) {
	reply, err := s.uc.StopTopic(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func ResumeTopic(ctx http.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	h := ctx.Middleware(func(c context.Context, req interface{}) (interface{}, error) {
		return biz.StartTopic(ctx, c)
	})
	_, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func GetTopicSream(ctx http.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	h := ctx.Middleware(func(c context.Context, req interface{}) (interface{}, error) {
		return biz.GetTopicStream(ctx)
	})
	_, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func UploadDocument(ctx http.Context) error {
	h := ctx.Middleware(func(c context.Context, req interface{}) (interface{}, error) {
		file, handler, err := ctx.Request().FormFile("file")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		return biz.UploadDocument(c, file, handler)
	})
	_, err := h(ctx, nil)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]interface{}{
		"message": "上传成功",
	})
}

func (s *SeminarService) GetDocuments(ctx context.Context, req *v1.GetDocumentsRequest) (*v1.GetDocumentsReply, error) {
	reply, err := s.uc.GetDocuments(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) AddMCPServer(ctx context.Context, req *v1.AddMCPServerReqeust) (*v1.AddMCPServerReply, error) {
	reply, err := s.uc.AddMCPServer(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) GetMCPServers(ctx context.Context, req *v1.GetMCPServersRequest) (*v1.GetMCPServersReply, error) {
	reply, err := s.uc.GetMCPServers(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) CheckMCPServerHealth(ctx context.Context, req *v1.CheckMCPServerHealthReqeust) (*v1.CheckMCPServerHealthReply, error) {
	reply, err := s.uc.CheckMCPServerHealth(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) DeleteMCPServer(ctx context.Context, req *v1.DeleteMCPServerRequest) (*v1.DeleteMCPServerReply, error) {
	reply, err := s.uc.DeleteMCPServer(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) EnableMCPServer(ctx context.Context, req *v1.EnableMCPServerRequest) (*v1.EnableMCPServerReply, error) {
	reply, err := s.uc.EnableMCPServer(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *SeminarService) DisableMCPServer(ctx context.Context, req *v1.DisableMCPServerRequest) (*v1.DisableMCPServerReply, error) {
	reply, err := s.uc.DisableMCPServer(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
