package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web/middleware"
	"web/models"
)

var users = make(map[string]models.User)

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} gin.H "注册成功"
// @Failure 400 {object} gin.H "请求错误"
// @Failure 409 {object} gin.H "用户已存在"
// @Router /register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	if _, exists := users[user.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	users[user.Username] = user

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"data": models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
		},
	})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "登录凭证"
// @Success 200 {object} gin.H "登录成功"
// @Failure 400 {object} gin.H "请求错误"
// @Failure 401 {object} gin.H "用户名或密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	user, exists := users[loginReq.Username]
	if !exists || loginReq.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	middleware.LoginUser(user.Username, c)
	c.JSON(http.StatusOK, gin.H{
		"message":  "登录成功",
		"username": user.Username,
		"data": models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
		},
	})
}

// Logout 用户退出
// @Summary 用户退出
// @Description 用户退出登录
// @Tags users
// @Produce json
// @Success 200 {object} gin.H "退出成功"
// @Router /logout [post]
func Logout(c *gin.Context) {
	username := middleware.GetCurrentUser(c)
	middleware.LogoutUser(username, c)
	c.JSON(http.StatusOK, gin.H{"message": "退出成功"})
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前用户密码
// @Tags users
// @Accept json
// @Produce json
// @Param password body models.ChangePasswordRequest true "密码信息"
// @Success 200 {object} gin.H "密码修改成功"
// @Failure 400 {object} gin.H "请求错误"
// @Failure 401 {object} gin.H "原密码错误"
// @Router /user/password [put]
func ChangePassword(c *gin.Context) {
	username := middleware.GetCurrentUser(c)
	user := users[username]

	var changeReq models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&changeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	if changeReq.OldPassword != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "原密码错误"})
		return
	}

	user.Password = changeReq.NewPassword
	user.UpdatedAt = time.Now()
	users[username] = user
	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 获取当前登录用户信息
// @Tags users
// @Produce json
// @Success 200 {object} models.UserResponse "用户信息"
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {
	username := middleware.GetCurrentUser(c)
	user := users[username]
	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data": models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// UpdateUserInfo 更新用户信息
// @Summary 更新用户信息
// @Description 更新当前用户昵称
// @Tags users
// @Accept json
// @Produce json
// @Param userInfo body models.UpdateUserRequest true "用户信息"
// @Success 200 {object} models.UserResponse "更新后的用户信息"
// @Failure 400 {object} gin.H "请求错误"
// @Router /user/info [put]
func UpdateUserInfo(c *gin.Context) {
	username := middleware.GetCurrentUser(c)
	user := users[username]

	var updateReq models.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	if updateReq.Nickname != nil {
		user.Nickname = *updateReq.Nickname
	}
	user.UpdatedAt = time.Now()
	users[username] = user

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"data": models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// GetUsers 获取所有用户列表
// @Summary 获取所有用户
// @Description 获取所有用户列表（需要登录）
// @Tags users
// @Produce json
// @Success 200 {array} models.UserResponse "用户列表"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var userList []models.UserResponse
	for _, user := range users {
		userList = append(userList, models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    userList,
		"total":   len(userList),
	})
}

// CheckAuth 检查认证状态
// @Summary 检查认证状态
// @Description 检查当前用户认证状态
// @Tags users
// @Produce json
// @Success 200 {object} gin.H "认证状态"
// @Router /auth/check [get]
func CheckAuth(c *gin.Context) {
	username := middleware.GetCurrentUser(c)
	user := users[username]
	c.JSON(http.StatusOK, gin.H{
		"message":       "已认证",
		"authenticated": true,
		"username":      username,
		"data": models.UserResponse{
			Username:  user.Username,
			Nickname:  user.Nickname,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
