package birdie

import (
	"context"
	"errors"
	"github.com/shamanis/birdie/internal/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct{}

func (s *Service) Ping(ctx context.Context, in *emptypb.Empty) (*BaseResponse, error) {
	logger.Debug("receive ping request")
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) Store(ctx context.Context, in *Entry) (*BaseResponse, error) {
	entry := storage.NewEntry(in.GetKey(), in.GetValue(), in.GetTtl())
	storage.Store(entry)
	logger.WithField("key", in.GetKey()).Debug("store entry")
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) BulkStore(ctx context.Context, in *Entries) (*BaseResponse, error) {
	var entries []*storage.Entry
	for _, e := range in.Entry {
		entries = append(entries, storage.NewEntry(e.GetKey(), e.GetValue(), e.GetTtl()))
	}
	storage.BulkStore(entries)
	logger.WithField("entries_count", len(in.Entry)).Debug("bulk store entries")
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) Load(ctx context.Context, in *LoadQuery) (*Entry, error) {
	value, err := storage.Load(in.GetKey())

	if err != nil {
		if errors.Is(err, storage.NotFoundError) {
			logger.WithField("key", in.GetKey()).Debug("not found entry")
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			logger.WithField("err", err).Error("unknown error")
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}
	logger.WithField("key", in.GetKey()).Debug("load entry")
	return &Entry{Key: in.GetKey(), Value: value.Value}, nil
}

func (s *Service) Search(ctx context.Context, in *SearchQuery) (*Entries, error) {
	logger.WithField("pattern", in.GetPattern()).Debug("search request receive")
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (s *Service) mustEmbedUnimplementedBirdieServer() {}
