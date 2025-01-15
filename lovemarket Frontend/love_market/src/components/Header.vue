<template>
    <header class="header">
        <!-- Logo部分 -->
        <div class="logo">恋爱集市</div>
  
        <!-- 导航和搜索部分 -->
        <div class="nav-and-search">
            <!-- 导航栏 -->
            <nav class="navigation">
                <ul>
                    <li><a @click="goToHome" >首页</a></li>
                    <li><a @click="goToSqure" >广场</a></li>
                    <li><a @click="goToChat">消息</a></li>
                    <li><a @click="goToPostposts">发帖</a></li>
                    <li><a @click="goToMyposts">我的帖子</a></li>
                </ul>
            </nav>
  
            <!-- 搜索框 -->
            <div class="search-container">
                <input v-model="searchQuery" class="search-input" placeholder="找寻恋爱" />
                <button @click="goTosearch" class="search-btn">搜索</button>
            </div>
        </div>
  
        <!-- 个人主页入口 -->
        <div class="user-profile">
            <a @click="goToProfilePage">个人主页</a>
        </div>
    </header>
</template>
  
<script>
    import {useChatStore} from '@/store/chatStore'
    export default {
       
        data() {
            return {
                searchQuery: '',
                
            }
        },
        methods: {
            goTosearch() {
                if(this.searchQuery){
                    this.$router.push({ name: 'search',  params: { searchdetail: this.searchQuery }});
                }
                else{
                    alert("输入不可为空!")
                }
                
            },
            // 跳转到主页
            goToHome() {
                this.$router.push({ name: 'home'});
            },
            // 跳转到广场
            goToSqure() {
                this.$router.push({ name: 'square'});
            },
            goToPostposts(){
                this.$router.push({ name: 'postposts'});
            },
            goToProfilePage(){
                this.$router.push({ name: 'ProfilePage'});
            },
            goToChat(){
                
                this.$router.push({ name: 'chat', parmas:{friendId:0}});
            },
            goToMyposts(){
                this.$router.push({ name: 'myposts' });
            }

        }
    }
  </script>
  
<style scoped>
    /* 基本布局 */
    .header {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        background-color:  #f0a7c2;
        padding: 10px;
        color: white;
        display: flex;
        justify-content: space-between;
        align-items: center;
        z-index: 1000;
        box-sizing: border-box;
    }
    
    .logo {
        font-size: 24px;
        font-weight: bold;
    }
    
    .nav-and-search {
        display: flex;
        justify-content: center;
        flex-grow: 1;
        align-items: center;
        position: relative;
    }
    
    .navigation ul {
        list-style: none;
        display: flex;
        margin: -10px;
        padding: 0;
        position: absolute;
        left: 100px;  /* 导航栏靠 Logo 的右侧 */
    }
    
    .navigation li {
        margin-right: 30px;
    }
    
    .navigation a {
        color: white;
        text-decoration: none;
    }
    
    .navigation a:hover {
        color: #ccc;
        cursor: pointer; /* 鼠标样式为点击标识 */
    }
    
    .search-container {
        display: flex;
        align-items: center;
        margin-left: 20px;
    }
    
    .search-input {
        padding: 5px;
        margin-right: 10px;
        width: 200px;
    }
    
    .search-btn {
        padding: 5px 10px;
        background-color: #fff;
        color: #000000;
        border: none;
        cursor: pointer;
    }
    
    .search-btn:hover {
        background-color:#9a8d92;
    }
    
    .user-profile a {
        color: white;
        text-decoration: none;
        cursor: pointer;
    }
    
    .user-profile a:hover {
        color: #ccc;
    }
    
    /* 使主体内容不被 header 遮挡 */
    body {
        margin-top: 60px;
    }
    
    /* 响应式设计: 为更小的屏幕调整导航项的布局 */
    @media (max-width: 1200px) {
        .navigation ul {
            left: 50px;  /* 导航栏靠 Logo 的右侧 */
        }
        .navigation li {
        margin-right: 15px;
        }
    
        .search-input {
        width: 60%;
        }
    }
    
    @media (max-width: 992px) {
        .navigation li {
        margin-right: 10px;
        }
    
        .search-input {
        width: 50%;
        }
    
        .navigation ul {
        left: 23px;  /* 让导航稍微靠右一点 */
        }
    }
    
    @media (max-width: 768px) {
        .logo {
        font-size: 20px;  /* 缩小 Logo */
        }
    
        .navigation li {
        margin-right: 5px;  /* 缩小导航项间隔 */
        }
    
        .search-input {
        width: 40%;  /* 缩小搜索框 */
        }
    
        .search-container {
        margin-left: 10px; /* 调整搜索框位置 */
        }
    
        .user-profile a {
        font-size: 14px;  /* 调整个人主页入口字体 */
        }
    }
    
    @media (max-width: 480px) {
        .header {
        flex-direction: column;  /* 在小屏幕上将导航栏改为垂直布局 */
        align-items: flex-start;
        padding: 10px;
        }
    
        .nav-and-search {
        width: 100%;
        justify-content: space-between;
        }
    
        .navigation ul {
        position: relative;
        display: block;  /* 将导航项变为垂直排列 */
        }
    
        .navigation li {
        margin-right: 0;
        margin-bottom: 10px;
        }
    
        .search-input {
        width: 80%;  /* 增加搜索框的宽度 */
        }
    
        .search-btn {
        margin-top: 10px;
        width: 100%;
        }
    }
</style>
  