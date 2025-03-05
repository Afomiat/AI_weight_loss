package repository

import (
	"context"
	"time"

	"github.com/Afomiat/AI_weight_loss/backend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfoRepository interface {
	SaveUserInfo(ctx context.Context, userInfo *domain.UserInfo) error
	GetUserInfoByUserID(ctx context.Context, userID string) ([]domain.UserInfo, error)
}

type userInfoRepository struct {
	collection *mongo.Collection
}

func NewUserInfoRepository(db *mongo.Database, collectionName string) UserInfoRepository {
	return &userInfoRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *userInfoRepository) SaveUserInfo(ctx context.Context, userInfo *domain.UserInfo) error {
	userInfo.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, userInfo)
	return err
}

func (r *userInfoRepository) GetUserInfoByUserID(ctx context.Context, userID string) ([]domain.UserInfo, error) {
	var userInfos []domain.UserInfo
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var userInfo domain.UserInfo
		if err := cursor.Decode(&userInfo); err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}
	return userInfos, nil
}
