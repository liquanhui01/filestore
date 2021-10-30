package user

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
}
