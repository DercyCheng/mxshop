// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gorm.io/gorm"
	"nd/user_srv/application/service"
	"nd/user_srv/infrastructure/persistence"
	"nd/user_srv/interfaces/grpc"
)

// ProvideUserHandler 提供用户处理器
func ProvideUserHandler(db *gorm.DB) *grpc.UserHandler {
	userRepositoryImpl := persistence.NewUserRepository(db)
	userServiceImpl := service.NewUserService(userRepositoryImpl)
	userHandler := grpc.NewUserHandler(userServiceImpl)
	return userHandler
}