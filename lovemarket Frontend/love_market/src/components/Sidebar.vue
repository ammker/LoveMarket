<template>
    <aside class="sidebar">
      <!-- 热榜 -->
      <div class="hotlist">
        <h3>热榜</h3>
        <!-- 遍历显示热榜帖子，每个帖子是一个独立的容器 -->
        <div v-for="post in hotPosts" :key="post.id" class="post-container">
          <!-- 点击帖子标题跳转到对应帖子详情页 -->
          <a @click="goToPostDetail(post.id)" class="post-title">{{ post.title }}</a>
        </div>
      </div>
    </aside>
  </template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      hotPosts: [{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},{author_name: "admin",
        id: "1",
        user_id: 1,
        community_id: 1,
        title: "科比绝杀凯文杜兰特",
        summary: "1234345464234",
        content: "凯文杜兰特被科比绝杀了啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊",
        create_time: "2024-12-25T15:11:19+08:00",
        vote_count: 1},],  // 用来存储热榜数据的数组
      intervalId: null, // 用于存储定时器ID
    };
  },
  mounted() {
    // 初始化数据并启动定时器
    this.fetchHotPosts();
    this.startPolling();
  },
  beforeDestroy() {
    // 组件销毁前清除定时器
    this.stopPolling();
  },
  methods: {
    // 获取热榜数据
    fetchHotPosts() {
      axios.get('http://127.0.0.1:8080/api/v1/posts2?size=10')
        .then(response => {
          if (response.data.code === 1000)
            this.hotPosts = response.data.data;
        })
      .catch(error => {
        console.error('获取热榜数据失败:', error);
      });
    },
    // 启动定时器，每隔一段时间请求数据
    startPolling() {
      this.intervalId = setInterval(() => {
        this.fetchHotPosts();
      }, 600000); // 每隔10分钟请求一次
    },
    // 停止定时器
    stopPolling() {
      if (this.intervalId) {
        clearInterval(this.intervalId);
        this.intervalId = null;
      }
    },
    // 跳转到帖子的详情页
    goToPostDetail(postid) {
      this.$router.push({ name: 'posts', params: { id: postid } });
    }
  },
};
</script>

<style scoped>
/* 侧边栏容器样式 */
.sidebar {
    margin-top: -30px;
  width: 300px;  /* 设置侧边栏的宽度 */
  background-color: white;  /* 设置背景颜色 */
  padding: 15px;  /* 内边距，使内容与边缘有间隔 */
  border-right: 1px solid #dddddd;  /* 右边框 */
}

/* 侧边栏菜单样式 */
.sidebar ul {
  list-style-type: none;  /* 去掉默认的列表样式 */
  padding: 0;  /* 去掉默认的内边距 */
}

.sidebar li {
  margin: 10px 0;  /* 设置每个菜单项之间的间距 */
}

.sidebar a {
  text-decoration: none;  /* 去除链接下划线 */
  color: #333;  /* 设置链接颜色 */
}

/* 鼠标悬停时改变链接颜色 */
.sidebar a:hover {
  color: #ff9900;  /* 悬停时改变文字颜色 */
}

/* 热榜部分样式 */
.hotlist {
  margin-top: 20px;  /* 与侧边栏菜单有间距 */
}

/* 热榜标题样式 */
.hotlist h3 {
  font-size: 18px;  /* 设置标题字体大小 */
  font-weight: bold;  /* 设置标题字体为粗体 */
  margin-bottom: 10px;  /* 设置标题与下方内容之间的间距 */
}

/* 单个帖子容器样式 */
.hotlist .post-container {
  margin-bottom: 5px;  /* 设置每个帖子之间的间距 */
  padding: 10px;  /* 设置内边距，使内容不贴边 */
  background-color: #fff;  /* 设置容器背景颜色为白色 */
  border: 1px solid #ddd;  /* 为容器添加边框 */
  border-radius: 5px;  /* 设置圆角效果 */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);  /* 给容器添加轻微阴影效果 */
}

/* 帖子标题的样式 */
.hotlist .post-title {
  font-size: 16px;  /* 设置标题字体大小 */
  color: #333;  /* 设置标题字体颜色 */
  text-decoration: none;  /* 去除链接下划线 */
  cursor: pointer; /* 显示为点击的手指图标 */
}

/* 鼠标悬停时改变标题颜色 */
.hotlist .post-title:hover {
  color: #ff9900;  /* 悬停时改变标题颜色 */
}
</style>
