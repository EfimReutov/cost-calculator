package grpc

import (
	"context"
	"cost-calculator/gen/proto/base_service"
	"cost-calculator/models"
	"github.com/shopspring/decimal"
	_ "github.com/shopspring/decimal"
	_ "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	_ "time"
)

func (s *server) InsertCategory(ctx context.Context, req *base_service.Request_InsertCategory) (*base_service.Response_InsertCategory, error) {
	err := s.store.InsertCategory(&models.Category{
		Name: req.Category.Name,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_InsertCategory{Message: "successful inserted mutepacka"}, nil
}

func (s *server) GetCategory(ctx context.Context, req *base_service.Request_GetCategory) (*base_service.Response_GetCategory, error) {
	category, err := s.store.GetCategory(req.GetId())
	if err != nil {
		return nil, err
	}
	return &base_service.Response_GetCategory{
		Message: "this is work motherfucka",
		Category: &base_service.Category{
			Id:   category.ID,
			Name: category.Name,
		},
	}, nil
}

func (s *server) UpdateCategory(ctx context.Context, req *base_service.Request_UpdateCategory) (*base_service.Response_UpdateCategory, error) {
	err := s.store.UpdateCategory(&models.Category{
		ID:   req.Category.Id,
		Name: req.Category.Name,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_UpdateCategory{
		Message: "this is work motheCategory",
		Id:      req.Category.Id,
	}, nil
}

func (s *server) DeleteCategory(ctx context.Context, req *base_service.Request_DeleteCategory) (*base_service.Response_DeleteCategory, error) {
	err := s.store.DeleteCategory(req.Id)
	if err != nil {
		return nil, err
	}
	return &base_service.Response_DeleteCategory{
		Message: "udaleno normas",
		Id:      req.Id,
	}, nil
}

func (s *server) InsertSource(ctx context.Context, req *base_service.Request_InsertSource) (*base_service.Response_InsertSource, error) {
	err := s.store.InsertSource(&models.Source{
		Name: req.Source.Name,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_InsertSource{Message: "successful inserted mutepacka"}, nil
}

func (s *server) GetSource(ctx context.Context, req *base_service.Request_GetSource) (*base_service.Response_GetSource, error) {
	source, err := s.store.GetSource(req.GetId())
	if err != nil {
		return nil, err
	}
	return &base_service.Response_GetSource{
		Message: "this is work motherfucka",
		Source: &base_service.Source{
			Id:   int64(source.ID),
			Name: source.Name,
		},
	}, nil
}

func (s *server) UpdateSource(ctx context.Context, req *base_service.Request_UpdateSource) (*base_service.Response_UpdateSource, error) {
	err := s.store.UpdateSource(&models.Source{
		ID:   req.Source.Id,
		Name: req.Source.Name,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_UpdateSource{
		Message: "this is work motheCategory",
		Id:      req.Source.Id,
	}, nil
}

func (s *server) DeleteSource(ctx context.Context, req *base_service.Request_DeleteSource) (*base_service.Response_DeleteSource, error) {
	err := s.store.DeleteSource(req.Id)
	if err != nil {
		return nil, err
	}
	return &base_service.Response_DeleteSource{
		Message: "udaleno normas",
		Id:      req.Id,
	}, nil
}

func (s *server) InsertIncome(ctx context.Context, req *base_service.Request_InsertIncome) (*base_service.Response_InsertIncome, error) {
	amount, err := decimal.NewFromString(req.Incoming.Amount)
	if err != nil {
		return nil, err
	}
	err = s.store.InsertIncome(&models.Income{
		SourceID: req.Incoming.SourceId,
		Amount:   amount,
		Date:     req.Incoming.Date.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_InsertIncome{Message: "successful inserted incoming mutepacka"}, nil
}

func (s *server) GetIncome(ctx context.Context, req *base_service.Request_GetIncome) (*base_service.Response_GetIncome, error) {
	incoming, err := s.store.GetIncome(req.GetId())
	if err != nil {
		return nil, err
	}
	return &base_service.Response_GetIncome{
		Message: "this is work motherfucka",
		Incoming: &base_service.Incoming{
			Id:       incoming.ID,
			SourceId: incoming.SourceID,
			//Amount:   incoming.Amount,
			Date: timestamppb.New(incoming.Date),
		},
	}, nil
}

func (s *server) UpdateIncome(ctx context.Context, req *base_service.Request_UpdateIncome) (*base_service.Response_UpdateIncome, error) {
	err := s.store.UpdateIncome(&models.Income{
		ID:       req.Incoming.Id,
		SourceID: req.Incoming.SourceId,
		//Amount:   req.Amount,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_UpdateIncome{
		Message: "this is work motherfucka",
		Id:      req.Incoming.Id,
	}, nil
}

func (s *server) DeleteIncome(ctx context.Context, req *base_service.Request_DeleteIncome) (*base_service.Response_DeleteIncome, error) {
	err := s.store.DeleteIncome(req.Id)
	if err != nil {
		return nil, err
	}
	return &base_service.Response_DeleteIncome{
		Message: "udaleno normas",
		Id:      req.Id,
	}, nil
}

func (s *server) InsertSpend(ctx context.Context, req *base_service.Request_InsertSpend) (*base_service.Response_InsertSpend, error) {
	err := s.store.InsertSpend(&models.Spend{
		CategoryID: req.Spend.CategoryId,
		//Amount:      decimal.Decimal{},
		Description: req.Spend.Description,
		Date:        req.Spend.Date.AsTime(),
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_InsertSpend{Message: "successful inserted mutepacka"}, nil
}

func (s *server) GetSpend(ctx context.Context, req *base_service.Request_GetSpend) (*base_service.Response_GetSpend, error) {
	spend, err := s.store.GetSpend(req.GetId())
	if err != nil {
		return nil, err
	}
	return &base_service.Response_GetSpend{
		Message: "this is work motherfucka",
		Spend: &base_service.Spend{
			Id:         spend.ID,
			CategoryId: spend.CategoryID,
			//Amount:      spend.Amount,
			Description: spend.Description,
			Date:        timestamppb.New(spend.Date),
		},
	}, nil
}

func (s *server) UpdateSpend(ctx context.Context, req *base_service.Request_UpdateSpend) (*base_service.Response_UpdateSpend, error) {
	err := s.store.UpdateSpend(&models.Spend{
		ID:         req.Spend.Id,
		CategoryID: req.Spend.CategoryId,
		//Amount:   req.Spend.Amount,
	})
	if err != nil {
		return nil, err
	}
	return &base_service.Response_UpdateSpend{
		Message: "this is work motherfucka",
		Id:      req.Spend.Id,
	}, nil
}

func (s *server) DeleteSpend(ctx context.Context, req *base_service.Request_DeleteSpend) (*base_service.Response_DeleteSpend, error) {
	err := s.store.DeleteSpend(req.Id)
	if err != nil {
		return nil, err
	}
	return &base_service.Response_DeleteSpend{
		Message: "udaleno normas",
		Id:      req.Id,
	}, nil
}
