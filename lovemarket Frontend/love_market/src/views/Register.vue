<template>
    <Header />
    <!--缩进决定层级-->
    <div id="bgimg" :class="{'start-animation': goback}">
        <div id="register_box" v-if="!goback">
            <!-- 返回登录界面 -->
            <span class="back-link" @click="goBack">返回</span>

            <h2>
                <span class="highlight">注</span><span class="normal">册</span>
            </h2>

            <div id="app">
                <router-view />
            </div>

            <div class="input_box">
                <input type="text" v-model="email" placeholder="请输入邮箱" />
            </div>

            <div class="input_box">
                <input type="text" v-model="userName" placeholder="请输入用户名" />
            </div>

            <div class="input_box">
                <input type="password" v-model="password" placeholder="请输入密码" />
            </div>

            <div class="input_box">
                <input type="password" v-model="repeat_pwd" placeholder="请再次输入密码" />
            </div>

            <button :disabled="isSubmitting" @click="handleRegister">注册</button><br>

        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    import Header from '../components/Header.vue'
    export default {
        components: {
            Header,
        },
        data() {
            return {
                email: '',  // 存储邮箱
                userName: '',     //存储用户名
                password: '',      // 存储密码
                repeat_pwd: '',    // 存储重复密码
                isSubmitting: false, // 控制按钮禁用，点击注册时防止重复点击
                goback: false
            };
        },
        name: 'App',

        methods: {

            handleRegister() {

                //前端检验部分--------------------------------------------------------
                //输入信息不完整
                if (!this.email || !this.userName || !this.password || !this.repeat_pwd) {
                    alert('请输入完整的用户信息');
                    return;
                }

                // 校验邮箱格式:xxx@xxx.xxx
                const emailPattern = /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/;
                if (!emailPattern.test(this.email)) {
                        alert('请输入有效的邮箱');
                        return;
                }

                // 校验用户名长度
                if (this.userName.length < 3 || this.userName.length > 15) {
                    alert('用户名长度应在3到15个字符之间');
                    return;
                }

                 // 校验密码强度（至少8位，包含字母和数字）
                const passwordPattern = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/;
                if (!passwordPattern.test(this.password)) {
                    alert('密码强度不够，请确保密码至少包含8个字符，并包含字母和数字');
                    return;
                }
                
                // 校验密码是否一致
                if (this.password !== this.repeat_pwd) {
                    alert('前后输入密码不一致');
                    return;
                }

                //前端检验部分--------------------------------------------------------


                //前后端交接部分~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 
                //设置提交状态
                this.isSubmitting = true;

                // 后端请求：发送注册请求
                axios.post('http://127.0.0.1:8080/api/v1/signup', {
                    "username": this.userName,
                    "password":  this.password, 
                    "email": this.email
                })
                .then(response => {
                    if (response.data.code === 1000) {
                        // 注册成功，跳转到登录页面
                        alert('注册成功');
                        this.isSubmitting = false; //请求结束，恢复按钮状态

                        //跳转回登陆页面
                        this.goback = true;
                        // 延迟2秒后跳转到登录页面
                        setTimeout(() => {
                            this.$router.push('/login');  // 跳转到注册页面
                        }, 500);  // 延迟0.5秒

                    } else {
                        // 显示错误信息（根据后端返回的msg）
                        alert(response.data.msg);
                        this.isSubmitting = false; //请求结束，恢复按钮状态
                    }
                })
                .catch(error => {
                    // 捕获请求失败的错误
                    console.error('注册请求失败', error);
                    alert('注册失败，请稍后再试');
                    this.isSubmitting = false; //请求结束，恢复按钮状态
                });

                //前后端交接部分~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                       
            },

            // 返回文字点击事件处理
            goBack() {
                this.goback = true;
                // 延迟2秒后跳转到登录页面
                setTimeout(() => {
                    this.$router.push('/login');  // 跳转到注册页面
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
        /*
        background: url('..\assets\pictures\login_background.jpg') no-repeat;
        background-size: cover;
        background-position: center;
        position: absolute; 
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        */
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
            background-color: white; /* 左边为白色 */
        }

        #bgimg::after {
            right: 0;
            background-color: #f0a7c2; /* 右边为粉色 */
        }

    #register_box {
        width: 30%; /* 可以适当调整宽度 */
        text-align: center;
        border-radius: 10px;
        padding: 50px;
        position: absolute;
        left: 32%;
        top: 15%;
        z-index: 1;
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
            transform: translateX(-5%);
        }

        20% {
            transform: translateX(-10%);
        }

        30% {
            transform: translateX(-15%);
        }

        40% {
            transform: translateX(-20%);
        }

        50% {
            transform: translateX(-25%);
        }

        60% {
            transform: translateX(-30%);
        }

        70% {
            transform: translateX(-35%);
        }

        80% {
            transform: translateX(-40%);
        }

        90% {
            transform: translateX(-45%);
        }

        100% {
            transform: translateX(-50%);
        }
    }



    h2 {
        color: #ffffff90; /* 设置标题默认颜色为淡白色 */
        margin-top: 0;
        font-size: 56px;
        z-index: 1;
    }

    .highlight {
        color: #f7a1c4; /* "注"字的颜色 */
        letter-spacing: 50px;
        z-index: 1;
    }

    .normal {
        color: white; /* "册"字的颜色 */
        letter-spacing: 0px;
        z-index: 1;
    }


    #input_box {
        margin-top: 50px;
        z-index: 1;
    }

    input {
        border: 0;
        width: 80%; /* 调整输入框宽度 */
        font-size: 15px;
        color: #000;
        background: transparent;
        padding: 5px 10px;
        outline: none;
        margin-top: 30px;
        border-bottom: 2px solid #808080;
        z-index: 1;
    }



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
        background-image: linear-gradient(to right, #f7a1c4, #e74c97); /* 更柔和的粉色渐变 */
        cursor: pointer; /* 鼠标悬停时显示手型光标 */
        z-index: 1;
    }



    /* 返回文字样式 */
    .back-link {
        position: absolute;
        top: 20px;
        left: 20px;
        font-size: 16px;
        color: #000;
        cursor: pointer;
        z-index: 1;
    }
        .back-link:hover {
            text-decoration: underline; /*鼠标悬停画下横线*/
        }

    .msg {
        margin-top: 20px;
        color: #fff;
    }

</style>
