<template>
  <div class="post-list">
    <p v-if="posts.length == 0">{{ mes }}</p>
    <PostItem
        v-for="(post, index) in posts"
        :key="index"
        :vote_count="post.vote_count"
        :title="post.title"
        :content="post.content"
        :author_name="post.author_name"
        :create_time="post.create_time"
        :id="post.id"
        :summary="post.summary"
    />
    <!-- 提示加载 -->
    <div v-if="posts.length != 0 && isLoading" class="loading">{{ mes }}</div>
  </div>
</template>
  
  <script>
  import PostItem from "./PostItem.vue";
  import axios from "axios";
  
  export default {
    name: "PostList",
    components: {
      PostItem,
    },
    data() {
      return {
        mes: "loading...",
        posts: [],   // 存放帖子的数组
        page: 1,     // 当前页码，初始值为1
        size: 8,  // 每页请求的帖子数
        isLoading: false, // 是否正在加载数据
      };
    },
    mounted() {
      // 在组件加载时请求数据
      this.fetchPosts();
      // 添加滚动监听
      window.addEventListener("scroll", this.handleScroll.bind(this));
    },
    beforeDestroy() {
      // 移除滚动监听
      window.removeEventListener("scroll", this.handleScroll.bind(this));
    },
    methods: {
      // 请求数据的方法
      fetchPosts() {
        if (this.isLoading) return; // 防止重复加载
        this.isLoading = true;      // 标记加载中
        // 发起 GET 请求，获取帖子列表数据
        axios.get(`http://127.0.0.1:8080/api/v1/posts?page=${this.page}&size=${this.size}`)
          .then(response => {
            if (response.data.code === 1000) {
              const { data, hasMore } = response.data;
              // 请求成功，更新 posts 数据
              this.posts = [...this.posts, ...response.data.data];
              // 更新 page 变量，每次请求后加一
              this.page += 1;
              this.isLoading = false; // 加载完成
              if (!hasMore) {
                // 如果没有更多数据，移除滚动事件监听
                window.removeEventListener("scroll", this.handleScroll);
                this.mes = "没有更多数据了";
              }
            }
            else {
              console.error('请求失败:', response.data.meg);
              this.isLoading = false; // 加载完成
            }
          })
          .catch(error => {
            console.error('请求失败:', error);
            this.isLoading = false; // 加载完成
            this.mes = '请求失败，请稍后重试'; // 提示用户请求失败
          });
      },
      // 处理滚动事件
      handleScroll() {
        const scrollTop = document.documentElement.scrollTop; // 已滚动高度
        const clientHeight = document.documentElement.clientHeight; // 视口高度
        const scrollHeight = document.documentElement.scrollHeight; // 页面总高度

        // 判断是否到达底部
        if (scrollTop + clientHeight >= scrollHeight - 50) {
          this.fetchPosts(); // 加载更多数据
        }
      },
    },
  };
</script>
  
  <style scoped>
  .post-list {
    display: flex;
    flex-direction: column;
    gap: 20px; /* 每个帖子的间距 */
    width: 100%; /* 列表宽度最大化 */
    max-width: 1500px; /* 最大宽度 */
    margin: 0 auto;
    padding: 20px;
    height: 100%; /* 增加高度，使列表更长 */
    overflow-y: auto; /* 如果内容过长，允许滚动 */
  }
  </style>
  