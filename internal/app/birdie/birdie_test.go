package birdie

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

var service = Service{}

func TestService_Ping(t *testing.T) {
	res, err := service.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Error(err)
	}

	if res.Status != ResponseStatus_STATUS_OK {
		t.Error("wrong response status")
	}
}

func TestService_Store(t *testing.T) {
	res, err := service.Store(context.Background(), &Entry{Key: "key1", Value: []byte("value"), Ttl: 5})
	if err != nil {
		t.Error(err)
	}

	if res.Status != ResponseStatus_STATUS_OK {
		t.Error("wrong response status")
	}
}

func TestService_Load(t *testing.T) {
	res, err := service.Load(context.Background(), &LoadQuery{Key: "key1"})

	if err != nil {
		t.Error(err)
	}

	if res.Key != "key1" {
		t.Error("wrong key")
	}

	res, err = service.Load(context.Background(), &LoadQuery{Key: "unexpected"})
	st, ok := status.FromError(err)
	if !ok {
		t.Error("wrong err type")
	}
	if st.Code() != codes.NotFound {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestService_BulkStore(t *testing.T) {
	var entries = &Entries{Entry: []*Entry{
		&Entry{Key: "key2", Value: []byte("value2"), Ttl: 5},
		&Entry{Key: "key3", Value: []byte("value3"), Ttl: 5},
		&Entry{Key: "key4", Value: []byte("value4"), Ttl: 5},
		&Entry{Key: "key5", Value: []byte("value6"), Ttl: 5},
	}}

	res, err := service.BulkStore(context.Background(), entries)

	if err != nil {
		t.Error(err)
	}

	if res.Status != ResponseStatus_STATUS_OK {
		t.Error("wrong response status")
	}
}

func TestService_Search(t *testing.T) {
	_, err := service.Search(context.Background(), &SearchQuery{Pattern: "key*"})

	st, ok := status.FromError(err)
	if !ok {
		t.Error("wrong err type")
	}
	if st.Code() != codes.Unimplemented {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
