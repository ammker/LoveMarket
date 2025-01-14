package controller

import (
	"lovemarket/global"
	"lovemarket/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	MID        int    `gorm:"column:MID" json:"mid"`                 // 消息ID，对应 MID
	SendUserID int    `gorm:"column:sendUserID" json:"send_user_id"` // 发送人ID，对应 sendUserID
	RevUserID  int    `gorm:"column:revUserID" json:"rev_user_id"`   // 接收人ID，对应 revUserID
	Content    string `gorm:"column:content" json:"content"`         // 消息内容，对应 content
	CreateTime string `gorm:"column:createTime" json:"create_time"`  // 创建时间，对应 createTime
}
type ChatRequest struct {
	UserID   int `json:"user_id" binding:"required"`   // 用户ID
	FriendID int `json:"friend_id" binding:"required"` // 好友ID
}

// FriendInfo 定义返回的好友信息结构体
type FriendInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// MessageCreateView 发布消息
func MessageCreateView(c *gin.Context) {
	// 从 Context 获取当前用户的 userid
	userid, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无法获取当前用户ID",
			"error":   err.Error(),
		})
		return
	}

	var cr Message

	// 绑定请求 JSON 参数到 MessageRequest（不接受 SendUserID，由后端确定）
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 将当前用户 ID 设置为发送人 ID
	cr.SendUserID = int(userid)

	var recvUser models.User

	// 查询接收人是否存在
	// err = global.Db.Take(&recvUser, cr.RevUserID).Error
	err = global.Db.Table("user").Find(&recvUser).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "接收人不存在",
		})
		return
	}

	// 创建消息记录并写入数据库
	message := Message{
		SendUserID: cr.SendUserID,
		RevUserID:  cr.RevUserID,
		Content:    cr.Content,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = global.Db.Create(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "消息发送失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "消息发送成功",
		"data": gin.H{
			"message_id":   message.MID,
			"send_user_id": message.SendUserID,
			"rev_user_id":  message.RevUserID,
			"content":      message.Content,
			"create_time":  message.CreateTime,
		},
	})
}

// 根据发送者id和接收者id获取消息
func GetChatMessages(c *gin.Context) {
	// 获取当前用户的ID
	userid, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无法获取当前用户ID",
			"error":   err.Error(),
		})
		return
	}

	// 创建请求对象并填充UserID
	var req ChatRequest
	req.UserID = int(userid)

	// 绑定好友ID参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 查询聊天记录
	var messages []Message
	err = global.Db.Where(
		"(sendUserID = ? AND revUserID = ?) OR (sendUserID = ? AND revUserID = ?)",
		req.UserID, req.FriendID, req.FriendID, req.UserID,
	).Order("createTime ASC").Find(&messages).Error

	// 检查查询错误
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取聊天记录失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回聊天记录
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"messages": messages,
	})
}

// 获取好友列表
func GetFriendsHandler(c *gin.Context) {
	// 获取当前用户的ID
	userid, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无法获取当前用户ID",
			"error":   err.Error(),
		})
		return
	}

	// 查询好友的 friend_id 列表
	var friendIDs []int
	err = global.Db.Table("friends").Select("friend_id").Where("user_id = ?", userid).Find(&friendIDs).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取好友ID列表失败",
			"error":   err.Error(),
		})
		return
	}

	// 如果没有好友，直接返回空列表
	if len(friendIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "好友列表为空",
			"data":    []FriendInfo{},
		})
		return
	}

	// 查询好友的详细信息
	var friends []FriendInfo
	err = global.Db.Table("user_details").Select("userID as id, username as name, profile_picture as avatar").
		Where("userID IN ?", friendIDs).Find(&friends).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取好友详情失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回好友列表
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取好友列表成功",
		"data":    friends,
	})
}

func GetMeHandler(c *gin.Context) {
	// 获取当前用户的ID
	userid, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无法获取当前用户ID",
			"error":   err.Error(),
		})
		return
	}

	var me FriendInfo

	err = global.Db.Table("user_details").Select("userID as id, username as name, profile_picture as avatar").
		Where("userID = ?", userid).Find(&me).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取好友详情失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回好友列表
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取好友列表成功",
		"data":    me,
	})
}
