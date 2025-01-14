<template>
    <div class="post-item">
      <!-- 左上角：点赞数 -->
      <div class="vote_count">
        {{ vote_count }}
      </div>
  
      <!-- 中间：标题和内容 -->
      <div class="post-content">
        <a class="post-title" @click="goToPostDetail">{{ title }}</a>
        <p class="post-body">{{ summary }}</p>
      </div>
  
      <!-- 右侧：用户名和时间 -->
      <div class="post-info">
        <div class="author_name">{{ author_name }}</div>
        <div class="time">{{ formatTime(create_time) }}</div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "PostItem",
    props: {
      vote_count: {
        type: Number,
        required: true,
      },
      title: {
        type: String,
        required: true,
      },
      content: {
        type: String,
        required: true,
      },
      author_name: {
        type: String,
        required: true,
      },
      create_time: {
        type: String,
        required: true,
      },
      id: {  // 确保传递帖子 id
        type: String,
        required: true,
      },
      summary: {
        type: String,
        required: true,
      }
    },
    methods:{
      // 跳转到帖子的详情页
      goToPostDetail() {
        this.$router.push({ name: 'posts', params: { id: this.id } });
      },
      formatTime(isoTime) {
        const date = new Date(isoTime);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        return `${year}-${month}-${day} ${hours}:${minutes}`;
      },
    }
  };
  </script>
  
  <style scoped>
  .post-item {
    display: flex;
    justify-content: flex-start; /* 左对齐所有子项 */
    align-items: flex-start; /* 顶部对齐 */
    background-color: #f9f9f9;
    border: 1px solid #e3e3e3;
    border-radius: 5px;
    padding: 15px;
    position: relative;
    width: 100%;
    max-width: 1400px;
    box-sizing: border-box;
    min-height: 180px;
    gap: 10px; /* 子项之间的间距 */
  }
  
  /* 左上角：回顾数量 */
  .vote_count {
    position: absolute;
    top: 10px;
    left: 10px;
    font-size: 15px;
    font-weight: bold;
    color: #666;
  }
  
  /* 中间部分：标题和正文 */
  .post-content {
    flex: 1;
    margin-left: 50px; /* 增加左边距，避免回顾数量遮挡 */
    display: block;
    text-align: left; /* 强制让内容左对齐 */
  }
  
  .post-title {
    font-size: 18px;
    font-weight: bold;
    margin-top: 15px;
    color: #4e89ef;
    text-decoration: none;
    display: block; /* 让标题成为块级元素 */
    margin-bottom: 30px; /* 标题和正文之间的间距 */
    text-align: left; /* 强制让标题左对齐 */
    cursor: pointer; /* 显示为点击的手指图标 */
  }
  
  .post-title:hover {
    text-decoration: underline;
  }
  
  .post-body {
    font-size: 14px;
    color: #333;
    margin-left: 50px;
    margin-top: 5px;
    white-space: normal; /* 允许换行 */
    word-wrap: break-word; /* 处理长单词换行 */
    text-align: left; /* 强制正文文本左对齐 */
  }
  
  /* 右侧部分：用户名和时间 */
  .post-info {
    text-align: right;
    margin-top: 25px;
    font-size: 12px;
    color: #999;
  }
  
  .author_name {
    font-weight: bold;
  }
  
  .time {
    margin-top: 5px;
  }
  </style>
  