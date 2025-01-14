<template>
    <Header />
    <div class="settings-page">
        <h2>设置页面</h2>
        <p>在这里你可以更改你的个人信息和设置。</p>

        <form @submit.prevent="saveSettings">
            <div class="form-row">
                <div class="form-group">
                    <label for="avatar">头像链接</label>
                    <input type="text" v-model="avatar" id="avatar" placeholder="请输入头像URL" />
                </div>

                <div class="form-group">
                    <label for="nickname">昵称</label>
                    <input type="text" v-model="nickname" id="nickname" placeholder="请输入昵称" />
                </div>
            </div>

            <div class="form-row">
                <div class="form-group">
                    <label for="username">账户名</label>
                    <input type="text" v-model="username" id="username" placeholder="请输入账户名" />
                </div>

                <div class="form-group">
                    <label for="age">年龄</label>
                    <input type="number" v-model="age" id="age" placeholder="请输入年龄" />
                </div>
            </div>

            <div class="form-row">
                <div class="form-group">
                    <label for="gender">性别</label>
                    <select v-model="gender" id="gender">
                        <option value="男">男</option>
                        <option value="女">女</option>
                    </select>
                </div>

                <div class="form-group">
                    <label for="location">地区</label>
                    <input type="text" v-model="location" id="location" placeholder="请输入地区" />
                </div>
            </div>

            <div class="form-row">
                <div class="form-group">
                    <label for="school">学校</label>
                    <input type="text" v-model="school" id="school" placeholder="请输入学校" />
                </div>

                <div class="form-group">
                    <label for="hobbies">兴趣爱好</label>
                    <input type="text" v-model="hobbies" id="hobbies" placeholder="请输入兴趣爱好" />
                </div>
            </div>

            <div class="form-row">
                <div class="form-group">
                    <label for="signature">个性签名</label>
                    <input type="text" v-model="signature" id="signature" placeholder="请输入个性签名" />
                </div>

                <div class="form-group">
                    <label for="phone">手机号</label>
                    <input type="text" v-model="phone" id="phone" placeholder="请输入手机号" />
                </div>
            </div>

            <div class="form-row">
                <div class="form-group">
                    <label for="email">邮箱</label>
                    <input type="email" v-model="email" id="email" placeholder="请输入邮箱" />
                </div>
            </div>

            <button type="submit">保存设置</button>
        </form>

        <button @click="backToProfile">返回个人主页</button>
    </div>
</template>

<script>
import axios from 'axios';
import Header from "@/components/Header.vue";
axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
export default {
    components:{
        Header
    },
    data() {
        return {
            avatar: 'https://via.placeholder.com/150',
            nickname: 'JohnDoe123',
            username: 'john_doe_01',
            age: 30,
            gender: '男',
            location: '北京',
            school: '中山大学',
            hobbies: '旅行、摄影、音乐、编程',
            signature: '探索未知，享受生活。',
            phone: '+86 138 0000 0000',
            email: 'john.doe@email.com'
        };
    },
    methods: {
        async saveSettings() {
            try {
                // 发送用户更新信息到后端
                const response = await axios.post('http://127.0.0.1:8080/api/v1/updateuserDetail', {
                    profile_picture: this.avatar,
                    nickname: this.nickname,
                    username: this.username,
                    age: this.age,
                    gender: this.gender,
                    location: this.location,
                    school: this.school,
                    hobbies: this.hobbies,
                    signature: this.signature,
                    phone_number: this.phone,
                    email: this.email
                });
                console.log(response.data);
            } catch (error) {
                console.error('保存设置失败', error);
            }
        },
        backToProfile() {
            this.$router.push('/profile');
        }
    },
    created() {
        // 从后端获取用户信息
        axios.get('http://127.0.0.1:8080/api/v1/userdetail')
           .then(response => {
                this.avatar = response.data.profile_picture;
                this.nickname = response.data.nickname;
                this.username = response.data.username;
                this.age = response.data.age;
                this.gender = response.data.gender;
                this.location = response.data.location;
                this.school = response.data.school;
                this.hobbies = response.data.hobbies;
                this.signature = response.data.signature;
                this.phone = response.data.phone_number;
                this.email = response.data.email;
            })
           .catch(error => {
                console.error('获取用户信息失败', error);
            });
    }
};
</script>

<style scoped>
    .settings-page {
        max-width: 1000px; /* 调整为适应电脑屏幕的最大宽度 */
        margin: 0 auto;
        font-family: Arial, sans-serif;
        padding: 40px;
        background-color: #f9f9f9;
        border-radius: 20px;
        color: #333;
        box-sizing: border-box; /* 确保padding不影响布局 */
    }

    h2 {
        font-size: 3em; /* 增大标题字体 */
        color: #f0a7c2;
        margin-bottom: 40px;
    }

    .form-row {
        display: flex;
        flex-wrap: wrap;
        gap: 30px; /* 增加字段之间的间隔 */
    }

    .form-group {
        flex: 1 1 calc(50% - 30px); /* 每行显示两个表单项，宽度为50%，同时保持30px的间隔 */
        min-width: 300px;
        margin-bottom: 30px;
    }

        .form-group label {
            font-size: 1.4em; /* 增大标签字体 */
            margin-bottom: 10px;
            color: #f0a7c2;
        }

        .form-group input,
        .form-group select {
            width: 100%;
            padding: 15px;
            font-size: 1.2em;
            border: 1px solid #f0a7c2;
            border-radius: 10px;
            background-color: #fff0f5;
        }

    button {
        background-color: #f0a7c2;
        color: white;
        border: none;
        padding: 15px 25px;
        border-radius: 10px;
        cursor: pointer;
        margin-top: 40px;
        width: 100%;
        font-size: 1.5em; /* 增大按钮字体 */
    }

        button:hover {
            background-color: #f8c8d1;
        }

    /* 在更大的屏幕上确保按钮和表单项的间距合适 */
    @media (min-width: 1200px) {
        .form-row {
            gap: 40px; /* 增加大屏幕上的间隔 */
        }

        .form-group {
            flex: 1 1 calc(33% - 40px); /* 每行显示三个表单项，宽度为33% */
        }
    }
</style>