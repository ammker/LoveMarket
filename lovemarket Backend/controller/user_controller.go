package controller

import (
	"errors"
	"lovemarket/global"

	"fmt"
	"lovemarket/models2"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserDetail 获取个人主页信息
func GetUserDetail(ctx *gin.Context) {
	// 从Context获取userid
	userid, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userid not found in context"})
		return
	}

	var userDetail models2.UserDetail

	// 根据userid获取用户详情
	if err := global.Db.Where("userID = ?", userid).First(&userDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User detail not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 如果没有设置头像，返回默认头像
	if userDetail.ProfilePicture == "" {
		userDetail.ProfilePicture = "/path/to/default/avatar.png" // 默认头像路径
	}

	// 返回用户详情信息
	ctx.JSON(http.StatusOK, userDetail)
}

// 编辑个人主页信息
func UpdateUserDetail(ctx *gin.Context) {
	// 从Context获取userid
	userid, err := getCurrentUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userid not found in context"})
		return
	}

	var userDetail models2.UserDetail
	userDetail.UserID = uint(userid)
	if err := ctx.ShouldBindJSON(&userDetail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 处理上传的头像文件（如果存在）
	avatarFile, err := ctx.FormFile("profile_picture") // 获取表单中的头像文件字段
	if err == nil {                                    // 只有当上传了头像时才更新
		// 设置文件存储路径，可以是本地服务器路径
		avatarPath := fmt.Sprintf("./uploads/%s", avatarFile.Filename)

		// 保存文件到本地
		if saveErr := ctx.SaveUploadedFile(avatarFile, avatarPath); saveErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar file"})
			return
		}

		// 更新头像路径
		userDetail.ProfilePicture = avatarPath
	}

	// 更新用户信息
	if err := global.Db.Save(&userDetail).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":         "Profile updated successfully",
		"profile_picture": userDetail.ProfilePicture, // 返回当前头像路径
	})
}
