package leaderboard

import (
	"context"
	"mage/pkg/db"
	"mage/pkg/models"
	"mage/pkg/repository"
	"mage/pkg/utils"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var redisKey = "leaderboard"

type LeaderBoardService struct {
	UnimplementedLeaderBoardServiceServer
	userRepo    *repository.UserRepository
	redisClient *redis.Client
}

func NewLeaderBoardService() *LeaderBoardService {
	return &LeaderBoardService{
		userRepo:    repository.NewUserRepository(db.DbHand.MongoDB),
		redisClient: db.DbHand.RedisClient,
	}
}

func (s *LeaderBoardService) Endgame(ctx context.Context, req *EndgameRequest) (*EndgameResponse, error) {
	resp := &EndgameResponse{}

	userID, isValid := ctx.Value("userId").(string)
	if !isValid {
		return resp, models.ErrINVALID_TOKEN_ERROR
	}
	currentScore, err := s.redisClient.ZScore(redisKey, userID).Result()
	if err != nil {
		if err == redis.Nil {
			currentScore = 0
		} else {
			return resp, status.Errorf(codes.Internal, "could not get existing score: %v", err)
		}
	}

	currentScore += req.Score
	userDetail := redis.Z{Score: currentScore, Member: userID}
	if err := s.redisClient.ZAdd(redisKey, userDetail).Err(); err != nil {
		return resp, status.Errorf(codes.Internal, "could not save score: %v", err)
	}

	_, err = s.redisClient.ZRevRangeWithScores(redisKey, 0, -1).Result()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update ranking: %v", err)
	}

	resp.Status = "success"
	resp.Timestamp = utils.ConvertTimetoString(time.Now())

	return resp, nil
}

func (s *LeaderBoardService) LeaderBoard(ctx context.Context, req *LeaderBoardRequest) (*LeaderBoardResponse, error) {
	resp := &LeaderBoardResponse{}
	count, _ := s.userRepo.CountAllUsers(ctx)
	scores, err := s.redisClient.ZRevRangeWithScores(redisKey, 0, count).Result()
	if err != nil {
		return nil, err
	}

	var list []*Result
	for _, score := range scores {
		userIDHex, _ := primitive.ObjectIDFromHex(score.Member.(string))
		user, err := s.userRepo.GetUserById(ctx, userIDHex)
		if err != nil {
			return nil, err
		}

		list = append(list, &Result{
			Username: user.Username,
			Score:    score.Score,
			Id:       &wrapperspb.Int64Value{Value: user.UserID},
		})
	}

	resp.Status = "success"
	resp.Timestamp = utils.ConvertTimetoString(time.Now())
	resp.Result = list

	return resp, nil
}
