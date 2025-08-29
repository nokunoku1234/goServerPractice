package transport

import "time"

type GetUserProfileResponse struct {
	User UserProfileDTO `json:"user"`
}

type UserProfileDTO struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Bio        string    `json:"bio"`
	Status     string    `json:"status"`
	Gender     string    `json:"gender"`
	Prefecture string    `json:"prefecture"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
