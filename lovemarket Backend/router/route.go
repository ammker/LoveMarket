package router

import (
	"lovemarket/controller"
	"lovemarket/logger"
	"lovemarket/middlewares"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173", "http://127.0.0.1:5175", "http://localhost:5173"}, // 允许的前端来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                 // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},                                 // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                                          // 暴露的响应头
		AllowCredentials: true,                                                                                // 是否允许携带 Cookie
		MaxAge:           12 * time.Hour,                                                                      // 预检请求的缓存时间
	}))

	r.Static("/css", "./dist/css")
	r.Static("/img", "./dist/img") // 提供静态资源文件，映射到 dist/assets
	r.Static("/js", "./dist/js")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico") // 提供 favicon 图标

	// 处理 Vue3 的入口文件 index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	// 根据时间或分数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/postswithsearch", controller.GetPostListWithSearchHandler)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/posts/:id", controller.GetPostDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		v1.POST("/post", controller.CreatePostHandler)
		// 投票
		v1.GET("/postsbyself", controller.GetPostListByselfHandler)
		v1.POST("/vote", controller.PostVoteController)
		v1.GET("/votestatus", controller.VoteStatusController)
		v1.POST("/comment", controller.CreateCommentHandler)
		v1.PUT("/postdelete", controller.DeletePostHandler)
		//yexuweidedaima

		v1.GET("/userdetail", controller.GetUserDetail)
		v1.POST("/updateuserDetail", controller.UpdateUserDetail)
		v1.POST("/message/send", controller.MessageCreateView)
		v1.POST("/message/getmessages", controller.GetChatMessages)
		v1.GET("/message/getfriends", controller.GetFriendsHandler)
		v1.GET("/message/getme", controller.GetMeHandler)
	}

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html") // 返回前端入口文件
	})

	return r
}
