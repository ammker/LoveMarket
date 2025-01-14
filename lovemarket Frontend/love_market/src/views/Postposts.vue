<template>
    <div class="write-post">
      <Header />
  
      <!-- 写帖子表单 -->
      <div class="post-form" ref="dropdownContainer">
        <h2 class="form-title">写帖子</h2>
  
        <!-- 标题输入框 -->
        <div class="form-group">
          <label for="title">标题</label>
          <input
            type="text"
            id="title"
            v-model="postTitle"
            class="form-control"
            placeholder="请输入标题..."
          />
        </div>
  
        <!-- 分区标签选择框 -->
        <div class="form-group">
            <label for="community">选择分区标签</label>
            <div class="dropdown">
                <div class="dropdown-select" @click="toggleDropdown">
                    <span>{{ selectedCommunity ? selectedCommunity.name : '请选择分区标签' }}</span>
                    <i class="arrow"></i>
                </div>
                <ul v-if="dropdownOpen" class="dropdown-options">
                    <li
                        v-for="community in communities"
                        :key="community.id"
                        @click.stop="selectCommunity(community)"
                        :class="{ selected: selectedCommunity && selectedCommunity.id === community.id }"
                    >
                        {{ community.name }}
                    </li>
                </ul>
            </div>
        </div>

  
        <!-- 正文输入框 -->
        <div class="form-group">
          <label for="content">内容</label>
          <textarea
            id="content"
            v-model="postContent"
            class="form-control textarea"
            placeholder="请输入正文内容..."
          ></textarea>
        </div>
  
        <!-- 发布按钮 -->
        <div class="button-container">
          <button class="submit-button" @click="submitPost">发布</button>
          <button class="cancel-button" @click="cancelPost">取消</button>
        </div>
      </div>
    </div>
</template>
  

<script>
    import Header from "@/components/Header.vue";
    import axios from "axios";
    
    export default {
        name: "WritePost",
        components: {
        Header,
        },
        data() {
            return {
                postTitle: "", // 标题内容
                postContent: "", // 正文内容
                // 固定类别列表，每个类别有对应的 id 和名称
                communities: [
                { id: 1, name: "恋爱" },
                { id: 2, name: "校园" },
                { id: 3, name: "集市" },
                ],
                selectedCommunity: null, // 存储选中的类别对象
                dropdownOpen: false, // 控制下拉框是否显示
            };
        },
        mounted() {
        // 添加事件监听器，监听点击页面的任何地方
        document.addEventListener("click", this.closeDropdown);
        },
        beforeDestroy() {
        // 移除事件监听器，避免内存泄漏
        document.removeEventListener("click", this.closeDropdown);
        },
        methods: {
        // 切换下拉框状态
        toggleDropdown(event) {
            event.stopPropagation(); // 阻止冒泡，防止触发全局点击事件
            this.dropdownOpen = !this.dropdownOpen;
        },
        // 关闭下拉框
        closeDropdown(event) {
            const dropdownContainer = this.$refs.dropdownContainer;
            if (dropdownContainer && !dropdownContainer.contains(event.target)) {
                this.dropdownOpen = false;
            }
        },
        // 选择类别
        selectCommunity(community) {
            this.selectedCommunity = community; // 直接将选中的类别赋值
            this.dropdownOpen = false; // 收起下拉框
        },
        // 发布帖子
        submitPost() {
            if (!this.postTitle.trim() || !this.postContent.trim() || !this.selectedCommunity) {
                alert("标题、内容和分区标签不能为空！");
                return;
            }
            
            // 构建帖子数据
            const postData = {
                title: this.postTitle,
                content: this.postContent,
                community_id: this.selectedCommunity.id, // 传递类别 ID
            };

            // 发送请求
            axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
            axios
                .post("http://127.0.0.1:8080/api/v1/post", postData)
                .then((response) => {
                    if (response.data.code === 1000) {
                        const postId = parseInt(response.data.data.id, 10); // 将字符串转换为整数
                        alert("帖子发布成功！");
                        this.$router.push(`/posts/${postId}`); // 跳转到帖子详情页面
                    } else {
                        alert("帖子发布失败：" + response.data.meg);this.selectCommunity
                    }
                })
                .catch((error) => {
                    console.error("发布失败：", error);
                    alert("发布失败，请稍后再试");
                });
        },
        // 取消写帖子
        cancelPost() {
            this.postTitle = "";
            this.postContent = "";
            this.selectedCommunities = "";
            this.dropdownOpen = false;
            this.$router.push("/square"); // 返回主页或其他页面
        },
        },
    };
</script>
    
<style scoped>
    .form-control{
        width: 99.5%;
        height:25px;
    }
    .write-post {
        font-family: Arial, sans-serif;
        padding: 20px;
    }
    
    .post-form {
        max-width: 70%;
        margin: 50px auto;
        background-color: #fff;
        padding: 30px;
        border-radius: 8px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.15);
        box-sizing: border-box;
    }
    
    .form-title {
        font-size: 24px;
        font-weight: bold;
        margin-bottom: 20px;
        text-align: center;
    }
    
    .form-group {
        margin-bottom: 20px;
    }
    
    label {
        display: block;
        font-size: 16px;
        font-weight: bold;
        margin-bottom: 8px;
    }
    
    .dropdown {
        position: relative;
        width: 100%;
    }
    
    .dropdown-select {
        padding: 10px;
        font-size: 14px;
        border: 1px solid #ccc;
        border-radius: 4px;
        cursor: pointer;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    
    .dropdown-select:hover {
        background-color: #f9f9f9;
    }
    
    .dropdown-options {
        position: absolute;
        top: 100%;
        left: 0;
        right: 0;
        border: 1px solid #ccc;
        background-color: #fff;
        border-radius: 4px;
        max-height: 200px;
        overflow-y: auto;
        z-index: 10;
    }
    
    .dropdown-options li {
        padding: 10px;
        font-size: 14px;
        cursor: pointer;
    }
    
    .dropdown-options li:hover {
        background-color: #f0f0f0;
    }
    
    .dropdown-options li.selected {
        background-color: #e0e0e0;
    }
    
    .textarea {
        width: 100%;
        height: 200px;
        border: 1px solid #ccc;
        border-radius: 4px;
        padding: 10px;
        font-size: 14px;
        resize: none;
    }
    
    .button-container {
        display: flex; /* 启用 Flexbox */
        justify-content: center; /* 居中对齐 */
        align-items: center; /* 垂直居中对齐 */
        gap: 19px; /* 设置按钮之间的间距 */
        margin-top: 20px; /* 可选：为按钮组添加上边距 */
    }
    
    .submit-button {
        padding: 15px 24px;
        font-size: 14px;
        background-color: #f0a7c2;
        color: #ffffff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }
    
    .submit-button:hover {
        background-color: #c298a8;
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
  
  


  
