package birdie

import (
	"context"
	"errors"
	"github.com/shamanis/birdie/internal/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Service struct{}

func (s *Service) Ping(ctx context.Context, in *emptypb.Empty) (*BaseResponse, error) {
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) Store(ctx context.Context, in *Entry) (*BaseResponse, error) {
	entry := storage.NewEntry(in.GetKey(), in.GetValue(), time.Duration(in.GetTtl())*time.Second)
	storage.Store(entry)
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) BulkStore(ctx context.Context, in *Entries) (*BaseResponse, error) {
	return &BaseResponse{Status: ResponseStatus_STATUS_OK, Error: ""}, nil
}

func (s *Service) Load(ctx context.Context, in *LoadQuery) (*Entry, error) {
	value, err := storage.Load(in.GetKey())
	if err != nil {
		if errors.Is(err, storage.NotFoundError) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, storage.TypeCastError) {
			return nil, status.Error(codes.Internal, err.Error())
		} else {
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return &Entry{Key: in.GetKey(), Value: value.Value}, nil
}

func (s *Service) Search(ctx context.Context, in *SearchQuery) (*Entries, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented method")
}

func (s *Service) mustEmbedUnimplementedBirdieServer() {}
