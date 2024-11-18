package repo

import (
	"context"
	"time"

	"github.com/kenmobility/grpc-project/models"
	"github.com/kenmobility/grpc-project/models/dto"
	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateUser(ctx context.Context, req dto.CreateUserParams) (*models.User, error)
	GetUser(ctx context.Context, userId string) (*models.User, error)
	CreateOrder(ctx context.Context, user models.User) (*models.Order, error)
	GetOrder(ctx context.Context, orderId string) (*models.Order, error)
}

type GormStoreRepository struct {
	db *gorm.DB
}

func NewGormStoreRepository(db *gorm.DB) StoreRepository {
	return &GormStoreRepository{db: db}
}

func (r *GormStoreRepository) CreateUser(ctx context.Context, req dto.CreateUserParams) (*models.User, error) {
	user := models.User{
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormStoreRepository) GetUser(ctx context.Context, userId string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("public_id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormStoreRepository) CreateOrder(ctx context.Context, user models.User) (*models.Order, error) {
	order := models.Order{
		Status:    "new",
		UserId:    uint(user.ID),
		User:      &user,
		CreatedAt: time.Now(),
	}
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *GormStoreRepository) GetOrder(ctx context.Context, orderId string) (*models.Order, error) {
	var order models.Order

	if err := r.db.Where("public_id = ?", orderId).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}
