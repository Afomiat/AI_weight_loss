package usecase

import (
	"context"
	"time"

	"github.com/Afomiat/AI_weight_loss/domain"
	"github.com/Afomiat/AI_weight_loss/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfoUsecase interface {
	AddUserInfo(ctx context.Context, userInfo *domain.UserInfo) error
	GetUserProgress(ctx context.Context, userID string) ([]domain.UserInfo, error)
}

type userInfoUsecase struct {
	userInfoRepo repository.UserInfoRepository
	contextTime  time.Duration
}

func NewUserInfoUsecase(repo repository.UserInfoRepository, timeout time.Duration) UserInfoUsecase {
	return &userInfoUsecase{
		userInfoRepo: repo,
		contextTime:  timeout,
	}
}

func (u *userInfoUsecase) AddUserInfo(ctx context.Context, userInfo *domain.UserInfo) error {
	userInfo.UserID = primitive.NewObjectID()
	userInfo.CreatedAt = time.Now()

	timeoutCtx, cancel := context.WithTimeout(ctx, u.contextTime)
	defer cancel()

	return u.userInfoRepo.SaveUserInfo(timeoutCtx, userInfo)
}

func (u *userInfoUsecase) GetUserProgress(ctx context.Context, userID string) ([]domain.UserInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTime)
	defer cancel()
	return u.userInfoRepo.GetUserInfoByUserID(ctx, userID)
}
