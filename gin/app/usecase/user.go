package usecase

import (
	"context"
	"fmt"
)

/*
User interface.
*/
type User interface {
	List(c context.Context, m *ListUsersRequest) (*ListUsersResponse, error)
	Get(c context.Context, m *GetUserRequest) (*UserResponse, error)
	Create(c context.Context, m *CreateUserRequest) (*UserResponse, error)
	Update(c context.Context, m *UpdateUserRequest) (*UserResponse, error)
	Delete(c context.Context, m *DeleteUserRequest) error
}

/*
user mode.
*/
type user struct{}

/*
user Request/Response
*/
type ListUsersRequest struct{}
type ListUsersResponse struct{}
type GetUserRequest struct{}
type UserResponse struct{}
type CreateUserRequest struct{}
type UpdateUserRequest struct{}
type DeleteUserRequest struct{}

func NewUser() *user {
	return &user{}
}

func (u *user) List(c context.Context, m *ListUsersRequest) (*ListUsersResponse, error) {
	fmt.Println("user.List executed.")
	return &ListUsersResponse{}, nil
}

func (u *user) Get(c context.Context, m *GetUserRequest) (*UserResponse, error) {
	fmt.Println("user.Get executed.")
	return &UserResponse{}, nil
}

func (u *user) Create(c context.Context, m *CreateUserRequest) (*UserResponse, error) {
	fmt.Println("user.Create executed.")
	return &UserResponse{}, nil
}

func (u *user) Update(c context.Context, m *UpdateUserRequest) (*UserResponse, error) {
	fmt.Println("user.Update executed.")
	return &UserResponse{}, nil
}

func (u *user) Delete(c context.Context, m *DeleteUserRequest) error {
	fmt.Println("user.Delete executed.")
	return nil
}
