package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/juanignaciorc/microbloggin-pltf/internal/domain"
	mock_ports "github.com/juanignaciorc/microbloggin-pltf/mocks"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestTweetsService_CreateTweetOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCtx := context.Background()
	mockUUID := uuid.New()
	mockTw := domain.Tweet{
		ID:      mockUUID,
		UserID:  mockUUID,
		Message: "message mock",
	}

	mockRepo := mock_ports.NewMockTweetRepository(ctrl)
	s := NewTweetsService(mockRepo)

	mockRepo.
		EXPECT().
		CreateTweet(mockCtx, gomock.Any()).
		Return(mockTw, nil)

	got, err := s.CreateTweet(mockCtx, mockUUID, "message")
	if err != nil {
		t.Errorf("CreateTweet() error = %v, ", err)
		return
	}
	if !reflect.DeepEqual(got, mockTw) {
		t.Errorf("CreateTweet() got = %v, want %v", got, mockTw)
	}
}
