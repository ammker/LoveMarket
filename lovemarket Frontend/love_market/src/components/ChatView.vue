<template>
  <div class="chat-content">
    <div class="chat-header">
      <h2>与 {{ chatStore.activeFriend?.name }} 聊天</h2>
    </div>
    <div ref="chatMessages" class="chat-messages">
      <div
        v-for="message in chatStore.activeMessages"
        :key="message.id"
        :class="['msg', message.sender_id === chatStore.currentUserId ? 'me' : 'other']"
        @contextmenu.prevent="handleRightClick($event, message)"
      >
        <img :src="getAvatar(message.sender_id)" alt="avatar" class="avatar" />
        <div :class="['message-bubble', message.sender_id === chatStore.currentUserId ? 'me' : 'other']">
          <!-- 文本消息 -->
          <div v-if="message.text" class="message-content">
            {{ message.text }}
            <div class="timestamp">{{ message.timestamp }}</div>
          </div>
          <!-- 图片消息 -->
          <div v-if="message.imageUrl" class="message-content">
            <img
              :src="message.imageUrl"
              alt="图片"
              class="chat-image"
              @dblclick="openImageModal(message.imageUrl)"
            />
            <div class="timestamp">{{ message.timestamp }}</div>
          </div>
        </div>
      </div>
    </div>



    <div class="chat-input">
      <!-- 表情按钮 -->
      <button class="emoji-button" @click="toggleEmojiPicker">😀</button>
      <!-- 表情选择器 -->
      <div v-if="showEmojiPicker" class="emoji-picker">
        <span
          v-for="emoji in emojis"
          :key="emoji"
          class="emoji"
          @click="insertEmoji(emoji)"
        >
          {{ emoji }}
        </span>
      </div>


      <!-- 隐藏的文件输入 -->
      <input
        type="file"
        ref="fileInput"
        style="display: none"
        accept="image/*"
        @change="handleImageUpload"
      />
      <!-- 文本输入框 -->
      <textarea
        v-model="newMessage"
        @keyup="handleKeyUp"
        placeholder="输入消息...CTRL+ENTER换行..."
        rows="2"
        ref="textarea"
      ></textarea>
      <!-- 发送按钮 -->
      <button @click="sendMessage">发送</button>
    </div>


  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue' // 移除 onUpdated
import { useChatStore } from '../store/chatStore'
import { useRoute } from 'vue-router'

const chatStore = useChatStore()
const route = useRoute()

const newMessage = ref('')
const chatMessages = ref(null)
const fileInput = ref(null)
const isImageModalOpen = ref(false)
const selectedImageUrl = ref('')
const showEmojiPicker = ref(false)
const textarea = ref(null)



// 初始化时获取用户列表和消息列表
onMounted(() => {

  chatStore.fetchMe()
  chatStore.fetchUsers()
  chatStore.fetchMessages()
  console.log("True Messages:"+chatStore.messages)
  console.log("True Messages:"+chatStore.activeMessages)
  scrollToBottom() // 发送消息后滚动到底部
})


// 发送消息
const sendMessage =  () => {
  if (newMessage.value.trim() !== '') {
    const message = {
      sender_id: chatStore.currentUserId,
      receiver_id: chatStore.activeFriendId,
      text: newMessage.value.trim(),
      image_url: null // 如果是图片消息，这里替换为图片 URL
    }
    chatStore.sendMessage(message)
    newMessage.value = ''
    scrollToBottom() // 发送消息后滚动到底部
  }
}


const emojis = ref([
  '😀', '😃', '😄', '😁', '😆', '😅', '😂', '🤣', '😊', '😇',
  '🙂', '🙃', '😉', '😌', '😍', '🥰', '😘', '😗', '😙', '😚',
  '😋', '😛', '😝', '😜', '🤪', '🤨', '🧐', '🤓', '😎', '🥸',
  '🤩', '🥳', '😏', '😒', '😞', '😔', '😟', '😕', '🙁', '☹️',
  '😣', '😖', '😫', '😩', '🥺', '😢', '😭', '😤', '😠', '😡',
  '🤬', '🤯', '😳', '🥵', '🥶', '😱', '😨', '😰', '😥', '😓',
  '🤗', '🤔', '🤭', '🤫', '🤥', '😶', '😐', '😑', '😬', '🙄',
  '😯', '😦', '😧', '😮', '😲', '🥱', '😴', '🤤', '😪', '😵',
  '🤐', '🥴', '🤢', '🤮', '🤧', '😷', '🤒', '🤕', '🤑', '🤠',
  '😈', '👿', '👹', '👺', '🤡', '💩', '👻', '💀', '☠️', '👽',
  '👾', '🤖', '🎃', '😺', '😸', '😹', '😻', '😼', '😽', '🙀',
  '😿', '😾'
])

// 获取用户头像
const getAvatar = (userId) => {
  console.log("Updating...")
  if(userId === chatStore.currentUserId){
    const user=chatStore.me
    return user ? user.avatar : ''
  }
  else {
    const user = chatStore.users.find((user) => user.id === userId)
    
    return user ? user.avatar : ''
  }
  
}

// 监听路由参数变化，更新当前选中的好友
watch(
  () => route.params.friendId,
  (friendId) => {
    if (friendId) {
      chatStore.setActiveFriend(Number(friendId))
      chatStore.fetchMessages()
    }
  },
  { immediate: true }
)

// 处理键盘事件
const handleKeyUp = (event) => {
  if (event.key === 'Enter' && !event.ctrlKey) {
    event.preventDefault()
    sendMessage()
  } else if (event.ctrlKey && event.key === 'Enter') {
    event.preventDefault()
    newMessage.value += '\n'
  }
}

// 滚动到底部
const scrollToBottom = () => {
  if (chatMessages.value) {
    chatMessages.value.scrollTop = chatMessages.value.scrollHeight
  }
}
const triggerImageUpload = () => {
  fileInput.value.click()
}

// 处理图片上传
const handleImageUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    
    const reader = new FileReader()
    reader.onload = async (e) => {
      const imageUrl = e.target.result
      const message = {
        sender_id: chatStore.currentUserId,
        receiver_id: chatStore.activeFriendId,
        text: null,
        image_url: imageUrl // 如果是图片消息，这里替换为图片 URL
      }
      chatStore.sendMessage(message)   
      
    }
    scrollToBottom() // 发送图片后滚动到底部
    reader.readAsDataURL(file)
  }
}
// 打开图片预览模态框
const openImageModal = (imageUrl) => {
  selectedImageUrl.value = imageUrl
  isImageModalOpen.value = true
}

// 关闭图片预览模态框
const closeImageModal = () => {
  isImageModalOpen.value = false
  selectedImageUrl.value = ''
}

// 切换表情选择器
const toggleEmojiPicker = () => {
  showEmojiPicker.value = !showEmojiPicker.value
}

// 插入表情到输入框
const insertEmoji = (emoji) => {
  const textareaElement = textarea.value
  const startPos = textareaElement.selectionStart
  const endPos = textareaElement.selectionEnd

  newMessage.value =
    newMessage.value.substring(0, startPos) +
    emoji +
    newMessage.value.substring(endPos)

  textareaElement.focus()
  textareaElement.setSelectionRange(startPos + emoji.length, startPos + emoji.length)

  showEmojiPicker.value = false
}
</script>

<style scoped>
.chat-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 90vh; /* 确保聊天界面占满整个视口高度 */
  box-sizing: border-box;
  background-color: #fff5f8; /* 粉色背景 */
}

.chat-header {
  background-color: #f0a7c2; /* 粉色主题色 */
  color: white;
  padding: 10px;
  text-align: center;
  font-size: 18px;
  font-weight: bold;
  border-radius: 10px;
  border-bottom: 1px solid #f0a7c2; /* 深粉色边框 */
}

.chat-messages {
  flex: 1;
  padding: 10px;
  overflow-y: auto; /* 允许消息容器滚动 */
  display: flex;
  flex-direction: column;
  background-color: #fff;
  max-height: calc(100vh - 200px); /* 限制消息容器的高度，避免溢出 */
  border: 1px solid #ffc0cb; /* 浅粉色边框 */
}

.msg {
  display: flex;
  align-items: flex-start;
  margin-bottom: 15px;
}

.msg.me {
  align-self: flex-end;
  flex-direction: row-reverse;
}

.msg.other {
  align-self: flex-start;
  flex-direction: row;
}

.message-bubble {
  max-width: 70%;
  padding: 10px;
  border-radius: 8px;
  position: relative;
  word-wrap: break-word;
  white-space: pre-wrap;
}

.message-bubble.me {
  background-color: #f0a7c2; /* 粉色主题色 */
  color: white;
  margin-left: 10px;
}

.message-bubble.other {
  background-color: #f1f1f1;
  color: #333;
  margin-right: 10px;
}

.message-content {
  display: flex;
  flex-direction: column;
}

.timestamp {
  font-size: 12px;
  color: #888;
  margin-top: 5px;
  align-self: flex-end;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 10px;
  margin-left: 10px;
  object-fit: cover;
}

.chat-image {
  max-width: 200px; /* 限制图片的最大宽度 */
  max-height: 200px; /* 限制图片的最大高度 */
  border-radius: 8px;
  margin-top: 5px;
  cursor: pointer;
  object-fit: cover; /* 保持图片比例 */
}

.chat-input {
  display: flex;
  padding: 10px;
  background-color: #fff;
  border-top: 1px solid #ffc0cb; /* 浅粉色边框 */
  align-items: center;
  position: relative; /* 用于定位表情选择器 */
}

.chat-input textarea {
  flex: 1;
  padding: 10px;
  border: 1px solid #ffc0cb; /* 浅粉色边框 */
  border-radius: 5px;
  resize: none;
  overflow: hidden;
  font-family: inherit;
  font-size: inherit;
  margin: 0 10px;
  background-color: #fff5f8; /* 粉色背景 */
}

.chat-input button {
  padding: 10px 20px;
  background-color: #f0a7c2; /* 粉色主题色 */
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
}

.chat-input button:hover {
  background-color: #ff1493; /* 深粉色 */
}

.upload-button {
  padding: 10px;
  background-color: #ff69b4; /* 粉色主题色 */
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
}

.upload-button:hover {
  background-color: #ff1493; /* 深粉色 */
}

.emoji-button {
  padding: 10px;
  background-color: #ff69b4; /* 粉色主题色 */
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 16px;
  margin-right: 10px;
}

.emoji-button:hover {
  background-color: #ff1493; /* 深粉色 */
}

.emoji-picker {
  position: absolute;
  bottom: 60px; /* 位于输入框上方 */
  left: 10px;
  background-color: #fff;
  border: 1px solid #ffc0cb; /* 浅粉色边框 */
  border-radius: 8px;
  padding: 10px;
  display: grid;
  grid-template-columns: repeat(8, 1fr); /* 每行显示 8 个表情 */
  gap: 5px;
  max-height: 200px;
  overflow-y: auto;
  z-index: 1000; /* 确保表情选择器在最上层 */
}

.emoji {
  cursor: pointer;
  font-size: 20px;
  transition: transform 0.2s ease;
}

.emoji:hover {
  transform: scale(1.2); /* 悬停时放大表情 */
}

/* 图片预览模态框 */
.image-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-image {
  max-width: 90%;
  max-height: 90%;
  border-radius: 8px;
}

/* 右键菜单样式 */
.context-menu {
  position: fixed;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.menu-item {
  padding: 8px 16px;
  cursor: pointer;
  font-size: 14px;
}

.menu-item:hover {
  background-color: #f5f5f5;
}
</style>