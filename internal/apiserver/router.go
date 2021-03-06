// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"

	"github.com/liquanhui01/filestore/internal/apiserver/controller/file"
	"github.com/liquanhui01/filestore/internal/apiserver/controller/folder"
	user "github.com/liquanhui01/filestore/internal/apiserver/controller/user"
	"github.com/liquanhui01/filestore/internal/apiserver/store/mysql"
	"github.com/liquanhui01/filestore/internal/pkg/auth"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(middleware.JWTAuthMiddleware())
	jwtMiddleware := auth.NewAuth()

	r.POST("/login", jwtMiddleware.LoginHandler)
	r.POST("/logout", jwtMiddleware.LogoutHandler)

	dbIns, _ := mysql.GetMySQLFactoryOr(nil)
	apiv1 := r.Group("/api/v1")
	{
		// user restful resource
		userv1 := apiv1.Group("/user")
		{
			userController := user.NewUserController(dbIns)

			userv1.POST("", userController.Create)
			userv1.DELETE("/:id", userController.Delete)
			userv1.GET("/:id", userController.Find)
			userv1.PUT("/:id", userController.Update)
			userv1.PUT("/change-password/:id", userController.ChangePassword)
		}

		// folder restful resource
		folderv1 := apiv1.Group("/folder")
		{
			folderController := folder.NewFolderController(dbIns)

			folderv1.POST("", folderController.Create)
			folderv1.GET("", folderController.Find)
			folderv1.GET("/:userid", folderController.FindByUserID)
			folderv1.PUT("", folderController.Update)
			folderv1.DELETE("", folderController.Delete)
		}

		// file restful resource
		filev1 := apiv1.Group("/file")
		{
			fileController := file.NewFileController(dbIns)

			filev1.POST("/:userid/:folderid", fileController.Upload)
			filev1.POST("/fastupload", fileController.FastUpload)
			filev1.GET("/:filesha1", fileController.Find)
			filev1.GET("", fileController.FindFilesByUserAndFolder)
			filev1.PUT("/folder", fileController.MoveFileToFolder)
			filev1.PUT("", fileController.Update)
			filev1.DELETE("/:filesha1", fileController.Delete)
		}
	}

	return r
}
