<template>
    <Header />
    <div id="app">
        <!-- 个人信息栏 -->
        <div class="profile-bar">
            <div class="profile-info">
                <div class="user-details">
                    <!-- 头像 -->
                    <img class="avatar" :src="user.avatarUrl" alt="用户头像" />

                    <div class="user-text-details">
                        <span class="nickname">{{ user.nickname }}</span>
                        <!-- 简介 -->
                        <p class="user-bio">{{ user.signature }}</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- 贴子列表 -->
        <div class="post-list">
            <p v-if="paginatedPosts.length === 0" class="no-posts-message">没有更多帖子了。</p>

            <div v-for="post in paginatedPosts" :key="post.id" class="post">
                <div class="post-header">
                    <h3 class="post-title" @click="handlePostClick(post.id)">
                        {{ post.title }}
                    </h3>

                    <button class="more-options" @click.stop="deletePost(post.id)">删除</button>
                </div>
                <div class="post-summary">
                    <p>{{ post.summary.slice(0, 50) }}...</p>
                    <div class="post-meta">
                        <p class="post-date">{{ formatTime(post.create_time) }}</p>
                        <p class="post-likes">点赞数: {{ post.vote_count }}</p>
                    </div>
                </div>

                <!-- 分区标签 -->
                <div class="post-category">
                    <span class="category-label">分区标签：{{ post.community.name }}</span>
                </div>
            </div>
        </div>

        <!-- 分页 -->
        <div class="pagination">
            <button @click="prevPage" :disabled="currentPage === 1" class="pagination-button">上一页</button>
            <span class="pagination-page">页码: {{ currentPage }}</span>
            <button @click="nextPage" :disabled="currentPage === totalPages" class="pagination-button next-page">下一页</button>
        </div>

        <!-- 添加新帖子按钮 -->
        <button class="add-post-button" @click="goToCreatePostPage">+</button>
    </div>
</template>

<script>
    import axios from 'axios'; // 导入axios
    import Header from '@/components/Header.vue';

    export default {
        components:{
            Header
        },
        data() {
            return {
                user: {
                    nickname: '昵称', // 初始化为占位符
                    avatarUrl: 'https://example.com/avatar.jpg',  // 初始头像URL
                    signature: '这是用户的个性签名。'  // 初始个性签名
                },
                posts: [],
                currentPage: 1,
                postsPerPage: 10
            };
        },

        computed: {
            totalPages() {
                return Math.ceil(this.posts.length / this.postsPerPage);
            },
            paginatedPosts() {
                const startIndex = (this.currentPage - 1) * this.postsPerPage;
                const endIndex = this.currentPage * this.postsPerPage;
                return this.posts.slice(startIndex, endIndex);
            }
        },

        methods: {
            nextPage() {
                if (this.currentPage < this.totalPages) {
                    this.currentPage++;
                    this.fetchPosts(); // 重新获取帖子数据
                }
            },
            prevPage() {
                if (this.currentPage > 1) {
                    this.currentPage--;
                    this.fetchPosts(); // 重新获取帖子数据
                }
            },
            deletePost(postId) {
                if (confirm('确定删除这个帖子吗？')) {
                    axios.put('http://127.0.0.1:8080/api/v1/postdelete',
                        {
                            postid: String(postId) // 发送的数据结构 { postid: postId }
                        },
                        {
                            headers: {
                                'Authorization': `Bearer ${localStorage.getItem('token')}`,
                                'Content-Type': 'application/json'
                            }
                        })
                        .then(response => {
                            if (response.data.code === 1000) {
                                alert('删除成功');
                                this.fetchPosts();
                            } else {
                                alert('删除失败: ' + response.data.msg);
                            }
                        })
                        .catch(error => {
                            console.error('删除请求错误:', error);
                        });
                }
            },

            goToCreatePostPage() {
                this.$router.push({ name: 'postposts' });
            },

            handlePostClick(postId) {
                this.$router.push({ name: 'posts', params: { id: postId} });
            },


            fetchUserInfo() {
                axios.get('http://127.0.0.1:8080/api/v1/userdetail', {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`,
                        'Content-Type': 'application/json'
                    }
                })
                    .then(response => {

                            const userData = response.data;
                            this.user.nickname = userData.nickname; // 更新昵称
                            this.user.avatarUrl = userData.profile_picture; // 更新头像
                            this.user.signature = userData.signature; // 更新个性签名
                        
                    })
                    .catch(error => {
                        console.error('请求错误:', error); // 错误处理
                    });
            },

            fetchPosts() {
                axios.get('http://127.0.0.1:8080/api/v1/postsbyself', {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('token')}`,
                        'Content-Type': 'application/json'
                    }
                })
                    .then(response => {
                        if (response.data.code === 1000) {
                            this.posts = response.data.data;
                        } else {
                            console.error('获取数据失败:', response.data.msg);
                        }
                    })
                    .catch(error => {
                        console.error('请求错误:', error);
                    });
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

        },

        mounted() {
            this.fetchUserInfo();  // 获取个人信息
            this.fetchPosts();  // 获取帖子列表
        }
    };
</script>

<style scoped>
    #app {
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 100%;
    }

    .profile-bar {
        background-color: #f0a7c2;
        padding: 20px;
        display: flex;
        justify-content: space-between;
    }

    .profile-info {
        display: flex;
        align-items: center;
    }

    .avatar {
        border-radius: 50%;
        width: 80px;
        height: 80px;
        margin-right: 20px; /* 头像和昵称之间的间距 */
    }

    .user-text-details {
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    .nickname {
        font-size: 1.5rem;
        font-weight: bold;
    }

    .user-bio {
        font-size: 1rem;
        color: #777;
        margin-top: 5px;
    }

    .post-list {
        flex: 1;
        padding: 20px;
        overflow-y: auto;
    }

    .post {
        background-color: #fff;
        padding: 15px;
        margin-bottom: 10px;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    .post-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 10px;
    }

    .post-title {
        font-size: 1.25rem;
        font-weight: bold;
    }

    .more-options {
        background: none;
        border: none;
        font-size: 1.5rem;
        cursor: pointer;
        color: gray;
    }

    .post-summary {
        font-size: 1rem;
        margin-bottom: 10px;
    }

    .post-meta {
        font-size: 0.875rem;
        color: #777;
    }

    .pagination {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 10px;
        margin-top: 20px;
    }

    .pagination-button {
        padding: 8px 16px;
        font-size: 1.25rem;
        background-color: #ddd;
        border: 1px solid #ccc;
        border-radius: 5px;
        cursor: pointer;
        margin: 0 10px;
    }

        .pagination-button:hover {
            background-color: #bbb;
        }

        .pagination-button:disabled {
            background-color: #f5f5f5;
            cursor: not-allowed;
        }

    .next-page {
        background-color: #007bff;
        color: white;
    }

        .next-page:hover {
            background-color: #0056b3;
        }

    .pagination-page {
        font-size: 1.25rem;
        font-weight: bold;
        margin: 0 10px;
    }

    .add-post-button {
    position: fixed;
    bottom: 20px; /* 距离屏幕底部 20px，可以根据需要调整 */
    left: 50%; /* 水平居中 */
    transform: translateX(-50%); /* 向左偏移自身宽度的 50% */
    width: 60px;
    height: 60px;
    background-color: #f0a7c2;
    color: white;
    font-size: 30px;
    border: none;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
}

.add-post-button:hover {
    background-color: #f25f72;
}
</style>
