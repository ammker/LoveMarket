<template>
    <div class="home">
        <Header />

        <!-- 帖子内容 -->
        <div class="post-detail">
            <div v-if="isLoading">{{ mes }}</div>

            <div v-else-if="post">
                <div class="post-header">
                    <div class="author_name">{{ post.author_name }}</div>
                    <h2 class="title">{{ post.title }}</h2>
                </div>

                <p class="post-body">{{ post.content }}</p>
                <div class="post-footer">
                    <div class="create_time">{{ formatTime(post.create_time) }}</div>
                    <div class="vote_count">
                        <img class="like-icon" 
                             :src="isVoted ? require('@/assets/voted.png') : require('@/assets/unvote.png')" 
                             alt="点赞按钮" 
                             @click="toggleLike" />
                        点赞数：{{ post.vote_count }}
                    </div>
                </div>
            </div>

            <div v-else>
                <p>帖子找不到了喔</p>
            </div>
        </div>

        <!-- 评论标题和回复按钮 -->
        <div class="comments-header">
            <span class="comments-title">评论</span>
            <div class="buttonstyle">
                <button class="reply-button" @click="showReplyInput">回复</button>
            </div>
        </div>

        <!-- 评论输入框 -->
        <div v-if="showInputBox" class="reply-section">
            <textarea v-model="newComment"
                      class="reply-textarea"
                      placeholder="请输入你的评论..."></textarea>
            <div class="button-container">
                <button class="reply-button" @click="submitComment">提交</button>
                <button class="cancel-button" @click="cancelReply">取消</button>
            </div>
        </div>

        <!-- 评论部分 -->
        <div>
            <div v-if="post.comments && post.comments.length > 0">
                <div v-for="(comment, index) in post.comments" :key="index" class="post-detail">
                    <div class="post-header">
                        <div class="author_name">{{ comment.user_name }}</div>
                    </div>
                    <p class="post-body">{{ comment.content }}</p>
                    <div class="post-footer">
                        <div class="create_time">{{ formatTime(comment.create_time) }}</div>
                    </div>
                </div>
            </div>
            <div v-else>
                <p>暂无评论</p>
            </div>
        </div>
    </div>
</template>


<script>
    import Header from '@/components/Header.vue';
    import axios from 'axios';

    export default {
        components: {
            Header,
        },
        name: 'PostDetail',
        data() {
            return {
                mes: 'loading...',
                showInputBox: false, // 控制输入框是否显示
                newComment: '', // 保存用户输入的评论
                post: {},
                isLoading: true, // 是否正在加载数据
                isVoted: false, // 是否已点赞
            };
        },
        mounted() {
            this.fetchPostDetail();
            this.fetchVoteStatus();
        },
        methods: {
            fetchPostDetail() {
                const postId = this.$route.params.id;
                axios
                    .get(`http://127.0.0.1:8080/api/v1/posts/${postId}`)
                    .then((response) => {
                        if (response.data.code === 1000) {
                            this.post = response.data.data;
                            this.isLoading = false;
                        } else {
                            console.error('请求失败');
                            this.isLoading = false;
                            this.mes = '加载失败: ' + response.data.msg;
                        }
                    })
                    .catch((error) => {
                        console.error('请求失败:', error);
                        this.isLoading = false;
                        this.mes = '请求失败，请稍后重试';
                    });
            },

            // 获取点赞状态
            fetchVoteStatus() {
                const postId = this.$route.params.id;
                axios
                    .get(`http://127.0.0.1:8080/api/v1/votestatus?postid=${postId}`)
                    .then((response) => {
                        if (response.data.code === 1000) {
                            this.isVoted = response.data.data.status === 1;
                        } else {
                            console.error('获取点赞状态失败');
                        }
                    })
                    .catch((error) => {
                        console.error('请求失败:', error);
                    });
            },

            // 切换点赞状态
            toggleLike() {
                const postId = this.$route.params.id;
                const direction = this.isVoted ? 0 : 1; // 如果已点赞，设置为取消点赞，否则点赞
                axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem('token')}`;
                axios
                    .post('http://127.0.0.1:8080/api/v1/vote', {
                        post_id: String(postId),
                        direction: direction,
                    })
                    .then((response) => {
                        if (response.data.code === 1000) {
                            this.isVoted = !this.isVoted; // 切换点赞状态
                            this.post.vote_count += this.isVoted ? 1 : -1; // 更新点赞数
                        } else {
                            alert('点赞失败: ' + response.data.msg);
                        }
                    })
                    .catch((error) => {
                        console.error('点赞请求失败:', error);
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

            showReplyInput() {
                this.showInputBox = true;
            },

            cancelReply() {
                this.showInputBox = false;
                this.newComment = '';
            },

            submitComment() {
                if (!this.newComment.trim()) {
                    alert('评论内容不能为空！');
                    return;
                }

                const commentData = {
                    postid: parseInt(this.$route.params.id, 10),
                    content: this.newComment,
                };

                axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem('token')}`;

                axios
                    .post('http://127.0.0.1:8080/api/v1/comment', commentData)
                    .then((response) => {
                        if (response.data.code === 1000) {
                            alert('评论成功！');
                            this.newComment = ''; // 清空输入框
                            this.showInputBox = false; // 隐藏输入框
                            this.fetchPostDetail();
                        } else {
                            alert('评论失败：' + response.data.msg);
                        }
                    })
                    .catch((error) => {
                        console.error('评论提交失败:', error);
                        alert('评论提交失败，请稍后再试');
                    });
            },
        },
    };
</script>

<style scoped>
    .home {
        font-family: Arial, sans-serif;
    }

    .post-detail {
        padding: 30px;
        background-color: #ffffff;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.25);
        border-radius: 10px;
        max-width: 80%;
        margin: 60px auto;
        box-sizing: border-box;
    }

    /* 帖子头部：用户名和标题 */
    .post-header {
        margin-bottom: 20px;
    }

    .author_name {
        font-size: 20px;
        font-weight: bold;
        text-align: left;
    }

    .title {
        font-size: 22px;
        font-weight: bold;
        margin-top: 10px;
        text-align: left;
    }

    .post-body {
        font-size: 16px;
        margin-top: 20px;
        text-align: left;
    }

    .post-footer {
    display: flex;
    justify-content: space-between;
    margin-top: 50px;
}

    .vote_count {
        font-size: 14px;
        color: #333;
        display: flex; /* 使用 Flexbox 布局 */
        align-items: center; /* 垂直居中对齐 */
    }

    .like-icon {
        width: 30px; /* 缩小图片宽度 */
        height: 30px; /* 缩小图片高度 */
        margin-right: 8px; /* 图片与点赞文字之间的间距 */
        cursor: pointer;
    }

    .create_time {
        font-size: 14px;
        color: #999;
    }

    /* 评论标题和回复按钮部分 */
    .comments-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 20px 0; /* 增加间距 */
    padding: 0 100px; /* 左右边距与帖子内容对齐 */
    }

    .comments-title {
    font-size: 25px;
    font-weight: bold;
    color: #333;
    text-align: left;
    }

    .buttonstyle{
    padding: 0 30px; /* 左右边距与帖子内容对齐 */
    }

    .button-container{
    display: flex; /* 启用 Flexbox */
    justify-content: center; /* 居中对齐 */
    align-items: center; /* 垂直居中对齐 */
    gap: 19px; /* 设置按钮之间的间距 */
    margin-top: 20px; /* 可选：为按钮组添加上边距 */
    }

    .reply-button {
    padding: 15px 24px;
    font-size: 14px;
    background-color: #f0a7c2;
    color: #ffffff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    }

    .reply-button:hover {
    background-color: #c298a8;
    }

    .reply-section {
    max-width: 80%;
    margin: 0 auto;
    padding: 20px 0;
    }

    .reply-textarea {
    width: 100%;
    height: 270px;
    border: 1px solid #ccc;
    border-radius: 4px;
    padding: 10px;
    font-size: 15px;
    resize: none;
    margin-bottom: 10px;
    box-sizing: border-box;
    }

    .cancel-button {
    padding: 15px 24px;
    font-size: 14px;
    background-color: #ccc;
    color: #ffffff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-left: 10px;
    }

    .cancel-button:hover {
    background-color: #999;
    }
    
</style>
