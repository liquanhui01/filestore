// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"

	user "github.com/liquanhui01/filestore/internal/apiserver/controller/user"
	"github.com/liquanhui01/filestore/internal/apiserver/store/mysql"
)

func InitRouter() *gin.Engine {
	r := gin.New()

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
		}
	}

	return r
}
