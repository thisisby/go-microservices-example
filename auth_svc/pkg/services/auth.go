package services

import (
	"auth_svc/pkg/database"
	"auth_svc/pkg/models"
	"auth_svc/pkg/proto"
	"auth_svc/pkg/utils"
	"context"
	"net/http"
)

type Server struct {
	H   database.DBConnection
	Jwt utils.JwtWrapper
	proto.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var user models.User

	if result := s.H.Conn.Where("email = ?", req.Email).First(&user); result.Error == nil {
		return &proto.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	s.H.Conn.Create(&user)

	return &proto.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	var user models.User

	if result := s.H.Conn.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return &proto.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.ComparePassword(user.Password, req.Password)

	if !match {
		return &proto.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid password",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return &proto.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &proto.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &proto.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	if result := s.H.Conn.Where("id = ?", claims.Id).First(&user); result.Error != nil {
		return &proto.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &proto.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
