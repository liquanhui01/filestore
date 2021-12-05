// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

// UserInfo define the infomation for login.
// type UserInfo struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func (u *UserController) Login(c *gin.Context) {
// 	var r UserInfo
// 	if err := c.ShouldBindJSON(&r); err != nil {
// 		core.WriteResponse(c, nil, http.StatusInternalServerError, "登陆失败", nil)
// 		return
// 	}

// 	user, err := u.srv.Users().FindUserByName(c, r.Username)
// 	if err != nil {
// 		core.WriteResponse(c, err, http.StatusBadRequest, "用户不存在", nil)
// 		return
// 	}

// 	fmt.Println("user is: ", user)

// 	if err = core.ValidatePassword(r.Password, user.Password); err != nil {
// 		core.WriteResponse(c, err, http.StatusBadRequest, "用户密码错误", nil)
// 		return
// 	}

// 	token, err := auth.GenerateToken(user.ID)
// 	if err != nil {
// 		core.WriteResponse(c, err, http.StatusInternalServerError, "登陆失败", nil)
// 		return
// 	}

// 	core.WriteResponse(c, nil, http.StatusOK, "登陆成功", token)
// }
