package user

import (
	"context"
	"errors"
	"mage/pkg/configs"
	"mage/pkg/db"
	"mage/pkg/models"
	"mage/pkg/repository"
	"mage/pkg/utils"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UnimplementedUserServiceServer
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{userRepo: repository.NewUserRepository(db.DbHand.MongoDB)}
}

func (s *UserService) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	resp := &RegisterUserResponse{}

	if isOK, err := user.IsValidInfo(); !isOK {
		return resp, err
	}

	user.HashedPassword()
	user.Id = primitive.NewObjectID()
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return resp, errors.New("sorry, registration couldn't process")
	}

	resp.Status = "success"
	resp.Timestamp = utils.ConvertTimetoString(user.CreatedAt)
	resp.Result = &UserData{Id: &wrappers.Int64Value{Value: user.UserID}, Username: user.Username, Password: user.Password}

	return resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	resp := &LoginUserResponse{}

	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return resp, models.ErrINVALID_USERNAME_OR_PASSWORD
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return resp, models.ErrINVALID_USERNAME_OR_PASSWORD
	}

	token, err := utils.GenerateJWTToken(user.Id.Hex(), configs.Cfg.SecretKey)
	if err != nil {
		return resp, models.ErrLOGIN_ERROR
	}

	resp.Status = "success"
	resp.Timestamp = utils.ConvertTimetoString(time.Now())
	resp.Result = &ResultData{Username: user.Username, Password: user.Password, Jwt: token}
	return resp, nil
}
