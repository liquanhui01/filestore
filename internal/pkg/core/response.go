// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Err     error       `json:"err"`
}

func WriteResponse(c *gin.Context, err error, code int, msg string, data interface{}) {
	if err != nil {
		fmt.Println("err is: ", err.Error())
		c.JSON(http.StatusBadRequest, Response{
			Code:    code,
			Message: msg,
			Err:     err,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: msg,
		Data:    data,
	})
}
