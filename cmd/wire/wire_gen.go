// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/internal/data"
	"echo-mangosteen/internal/repo"
	"echo-mangosteen/internal/router"
	"echo-mangosteen/internal/service"
	"echo-mangosteen/pkg/cache"
	"echo-mangosteen/pkg/config"
	"echo-mangosteen/pkg/jwt"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// Injectors from wire.go:

func NewApp(configConfig *config.Config) (*echo.Echo, func(), error) {
	jwtJWT := jwt.New(configConfig)
	dataData, cleanup, err := data.NewData(configConfig)
	if err != nil {
		return nil, nil, err
	}
	userRepo := repo.NewUserRepo(dataData)
	cacheCache, err := cache.NewCahce()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userService := service.NewUserService(userRepo, cacheCache, jwtJWT, configConfig)
	userController := controller.NewUserController(userService)
	codeController := controller.NewCodeController()
	tagRepo := repo.NewTagRepo(dataData)
	tagService := service.NewTagService(tagRepo)
	tagController := controller.NewTagController(tagService)
	itemRepo := repo.NewItemRepo(dataData)
	itemService := service.NewItemService(itemRepo)
	itemController := controller.NewItemController(itemService)
	echoEcho := router.NewRouter(jwtJWT, userController, codeController, tagController, itemController)
	return echoEcho, func() {
		cleanup()
	}, nil
}

// wire.go:

var controllerProvider = wire.NewSet(controller.NewUserController, controller.NewCodeController, controller.NewTagController, controller.NewItemController)

var repoProvider = wire.NewSet(repo.NewUserRepo, repo.NewTagRepo, repo.NewItemRepo)

var serviceProvider = wire.NewSet(service.NewUserService, service.NewTagService, service.NewItemService)
