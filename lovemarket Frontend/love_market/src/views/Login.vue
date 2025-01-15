<template>
    <Header />
    <!--缩进决定层级-->
    <div id="bgimg" :class="{'start-animation': isRegisterClicked}">
        <div id="login_box" v-if="!isRegisterClicked">
            <h2>
                <span class="normal">登</span><span class="highlight">录</span>
            </h2>
            <div id="app">
                <router-view />
            </div>

            <div class="input_box" v-if="!isRegisterClicked">
                <input type="text" v-model="username" placeholder="请输入用户名" />
            </div>

            <div class="input_box" v-if="!isRegisterClicked">
                <input type="password" v-model="password" placeholder="请输入密码" />
            </div>

            <button @click="handleLogin" :disabled="isSubmitting">登录</button>

            <div class="msg" v-if="!isRegisterClicked">
                尚未注册？ <span class="register-link" @click="goToRegister">点击注册</span>
            </div>
        </div>       
        
    </div>
</template>

<script>

    import axios from 'axios'
    import Header from '../components/Header.vue'
    import { useChatStore } from '@/store/chatStore';

    export default {
        components:{
            Header,
            useChatStore
        },

        data() {
            return {
                username: '',  // 存储用户名
                password: '',      // 存储密码
                isSubmitting: false , // 控制按钮禁用，点击登录时防止重复点击
                isRegisterClicked: false, // 控制是否点击了注册
                chatStore:useChatStore()
            };
        },
        name: 'App',

        methods: {
            mouted(){
                chatStore.initData()
            },
            handleLogin() {
                if (!this.username || !this.password) {
                    alert('请输入完整的登录信息');
                    return;
                }

                try {
                  // 发送登录请求
                  axios.post('http://127.0.0.1:8080/api/v1/login', {
                    username: this.username,
                    password: this.password
                  })
                  .then(response => {
                    // 检查是否登录成功
                    if (response.data.code === 1000) {
                      // 登录成功，获取 token
                      const token = response.data.data.token;

                      // 将 token 存储到 localStorage（或者 Vuex 中）
                      localStorage.setItem('token', token);

                      // 可以存储用户信息（如用户名、userid）方便后续使用
                      localStorage.setItem('user', JSON.stringify(response.data.data));

                      alert('登录成功');
                      this.chatStore.setCurrentUser(response.data.data.userid)
                      this.chatStore.fetchMe()
                      this.chatStore.fetchUsers()
                      this.chatStore.fetchMessages() 
                      this.$router.push('/home'); // 跳转到首页或其他页面
                    } 
                    else {
                      alert(response.data.msg);
                    }

                  })
                  .catch(error => {
                    console.error('登录请求失败', error);
                    alert('登录失败，请稍后再试');
                  })

                  .finally(() => {
                    // 恢复按钮状态
                    this.isSubmitting = false;
                  });
                } 
                
                catch (error) {
                  // 这里可以捕获其他同步错误
                  console.error('错误', error);
                }
            },

            // 跳转到注册页面
            goToRegister() {
                this.isRegisterClicked = true;
                 // 延迟2秒后跳转到注册页面
                setTimeout(() => {
                    this.$router.push('/register'); // 跳转到注册页面
                }, 500);  // 延迟0.5秒
            },
        }
    };
</script>

<style scoped>

    #bgimg {
        margin: 0;
        padding: 0;
        border: 0;
        height: 90vh; /* 设置body的高度为100vh，确保背景图片填充整个页面 */
        display: flex;
        justify-content: center;
        align-items: center; /* 使用Flexbox居中内容 */
        position: relative; /* 为了让绝对定位的伪元素相对父元素定位 */
    }

        #bgimg::before, #bgimg::after {
            content: '';
            position: absolute;
            top: 0;
            bottom: 0;
            width: 50%;
            z-index: -1; /* 确保背景色块不会覆盖内容 */
        }

        #bgimg::before {
            left: 0;
            background-color: #f0a7c2; /* 左边为粉色 */
        }

        #bgimg::after {
            right: 0;
            background-color: white; /* 右边为白色 */
        } 

    /*登录框*/
    #login_box {
        width: 30%; /* 可以适当调整宽度 */
        text-align: center;
        border-radius: 10px;
        padding: 50px;
        position: absolute;
        left: 32%;
        top: 15%;
        z-index: 1; /* 确保登录框位于背景色块的上面 */
    }

    /* 点击注册后，启动背景动画 */
    .start-animation {
        background: linear-gradient(to right, #f0a7c2 50%, white 50%,);
        animation: moveBackground 0.5s forwards; /* 背景动画 */
    }

    /*平移动画 */
    @keyframes moveBackground {
        0% {
            transform: translateX(0);
        }

        10% {
            transform: translateX(5%);
        }

        20% {
            transform: translateX(10%);
        }

        30% {
            transform: translateX(15%);
        }

        40% {
            transform: translateX(20%);
        }

        50% {
            transform: translateX(25%);
        }

        60% {
            transform: translateX(30%);
        }

        70% {
            transform: translateX(35%);
        }

        80% {
            transform: translateX(40%);
        }

        90% {
            transform: translateX(45%);
        }

        100% {
            transform: translateX(50%);
        }
    }


    /* 使用组合方式修正 `::before` 和 `::after`的移动效果 */
    #bgimg::before, #bgimg::after {
        transition: transform 0.5s;
    }


    /*登录标题*/
    h2 {
        color: #ffffff90; /* 设置标题默认颜色为淡白色 */
        margin-top: 0;
        font-size: 56px;
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    .highlight {
        color: #f7a1c4; /* "录"字的颜色 */
        letter-spacing: 20px;
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    .normal {
        color: white; /* "登"字的颜色 */
        letter-spacing: 80px;
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    /*输入框*/
    #input_box {
        margin-top: 20px;
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    input {
        border: 0;
        width: 80%; /* 调整输入框宽度 */
        font-size: 15px;
        color: #000000;
        background: transparent;
        border-bottom: 2px solid #808080;
        padding: 5px 10px;
        outline: none;
        margin-top: 30px;
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    /*登录按钮*/
    button {
        margin-top: 30px;
        width: 80%;
        height: 40px;
        border-radius: 10px;
        border: 0;
        color: #fff;
        text-align: center;
        line-height: 40px;
        font-size: 15px;
        background-image: linear-gradient(to right, #e74c97, #f7a1c4); /* 更柔和的粉色渐变 */
        cursor: pointer; /* 鼠标悬停时显示手型光标 */
        z-index: 1; /* 确保位于背景色块的上面 */
    }

    /*点击跳转注册的文字链接*/
    .register-link {
        color: #f29ac1; /* 默认字体颜色 */
        cursor: pointer; /* 鼠标悬停时显示手型光标 */
        transition: color 0.3s; /* 添加过渡效果 */
        margin-left: 35px; /* 调整文字向左移动 */
        z-index: 1; /* 确保位于背景色块的上面 */
    }

        .register-link:hover {
            color: #1e90ff; /* 鼠标悬停时改变为深蓝色 */
        }

    .msg {
        margin-top: 20px;
        color: #fff;
        z-index: 1; /* 确保位于背景色块的上面 */
    }
    
</style>
