package user

import (
	"context"
	"goServerPractice/internal/repository"
	"goServerPractice/internal/transport"

	"github.com/uptrace/bun"
)

type UserListRequest struct {
	Status     *string `query:"status"`
	Gender     *string `query:"gender"`
	Prefecture *string `query:"pref"`
	Page       int     `query:"page"`
	Limit      int     `query:"limit"`
}

type UserListResponse struct {
	Users      []transport.UserProfileDTO `json:"users"`
	TotalCount int                        `json:"totalCount"`
	Page       int                        `json:"page"`
	Limit      int                        `json:"limit"`
}

func GetUserList(db *bun.DB, ctx context.Context, req UserListRequest) (*UserListResponse, error) {
	repo := repository.NewUserRepository(db)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 20
	}
	offset := (req.Page - 1) * req.Limit

	filter := repository.UserFilter{
		Status:     req.Status,
		Gender:     req.Gender,
		Prefecture: req.Prefecture,
		Limit:      req.Limit,
		Offset:     offset,
		OrderBy:    "updated_at_desc",
	}

	users, totalCount, err := repo.GetUsers(ctx, filter)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]transport.UserProfileDTO, len(users))
	for i, user := range users {
		userDTOs[i] = transport.UserProfileDTO{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Status:     user.Status,
			Bio:        user.Bio,
			Gender:     user.Gender,
			Prefecture: user.Prefecture,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}
	}

	return &UserListResponse{
		Users:      userDTOs,
		TotalCount: totalCount,
		Limit:      req.Limit,
		Page:       req.Page,
	}, nil
}
