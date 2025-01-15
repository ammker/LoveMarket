<template>
  <Header/>
  <div class="chat-layout">
    <!-- 左侧好友列表 -->
    <div class="friend-list-container">
      <div
        v-for="friend in chatStore.users"
        :key="friend.id"
        :class="['friend-item', { active: friend.id === chatStore.activeFriendId }]"
      >
        <RouterLink :to="{ name: 'ChatView', params: { friendId: friend.id } }" class="friend-link">
          <img :src="friend.avatar" alt="avatar" class="avatar" />
          <span class="friend-name">{{ friend.name }}</span>
        </RouterLink>
      </div>
    </div>

    <!-- 右侧聊天界面 -->
    <div class="chat-content-container">
      <RouterView />
    </div>
  </div>
</template>

<script setup>
import { onMounted,onUpdated } from 'vue'
import { useChatStore } from '../store/chatStore'
import Header from '@/components/Header.vue';

const chatStore = useChatStore()

// 初始化时获取用户列表
onMounted(async () => {
  chatStore.fetchMe()
  chatStore.fetchUsers()
})

</script>
vh
<style scoped>
.chat-layout {
  display: flex;
  flex-direction: row; /* 恢复为水平排列 */
  height: 95vh;
  box-sizing: border-box;
  background-color: #fff5f8; /* 粉色背景 */
}

.friend-list-container {
  width: 200px;
  background-color: #fff5f8; /* 粉色背景 */
  border-right: 1px solid #ffc0cb; /* 浅粉色边框 */
  overflow-y: auto;
  margin-right: 20px;
  border-radius: 10px;
}

.friend-item {
  display: flex;
  align-items: center;
  padding: 10px;
  cursor: pointer;
  border-bottom: 1px solid #ffc0cb; /* 浅粉色边框 */
}

.friend-item.active {
  background-color: #f0a7c2; /* 粉色主题色 */
  color: white;
}

.friend-item:hover {
  background-color: #ffc0cb; /* 浅粉色背景 */
}

/* 将整个 friend-item 设置为可点击 */
.friend-link {
  display: flex;
  align-items: center;
  width: 100%;
  text-decoration: none; /* 移除链接的下划线 */
  color: inherit; /* 继承父元素的颜色 */
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
}

.friend-name {
  font-size: 14px;
  color: #ff1493; /* 深粉色文字 */
}

.chat-content-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #fff5f8; /* 粉色背景 */
  border-radius: 10px;
  border: 1px solid #ffc0cb; /* 浅粉色边框 */
}
</style>