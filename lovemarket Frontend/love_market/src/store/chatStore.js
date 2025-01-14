import { defineStore } from 'pinia'
import axios from 'axios'

const baseURL = 'http://127.0.0.1:8080'//待填URL


export const useChatStore = defineStore('chat', {
  state: () => ({
    me:[],
    users: [], // 从后端获取
    messages: [], // 从后端获取
    activeFriendId: null, // 当前选中的好友 ID
    currentUserId: null // 当前好友 ID
  }),
  getters: {
    // 获取当前选中的好友信息
    activeFriend: (state) => state.users.find(user => user.id === state.activeFriendId),
    // 获取当前好友的聊天记录
    activeMessages: (state) =>
      state.messages.filter(
        (msg) =>
          (msg.sender_id === state.currentUserId && msg.receiver_id === state.activeFriendId) ||
          (msg.sender_id === state.activeFriendId && msg.receiver_id === state.currentUserId)
      )
  },
  actions: {
    //从后端获取个人信息
    async fetchMe() {
      try {
        console.log("Getting...")
        axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
        const response = await axios.get(`http://127.0.0.1:8080/api/v1/message/getme`);//代填url
        
        // 根据后端返回的格式，提取 "data" 字段
        if (response.data.code === 200) {
          console.log(response.data.data)
          
          console.log('个人信息获取成功:', this.me);
          this.me=response.data.data
        } else {
          console.error('获取个人信息失败，后端返回错误:', response.data.message);
        }
      } catch (error) {
        console.error('获取个人信息失败:', error);
        console.error('请求 URL:', error.config?.url || '');
        console.error('错误详情:', error.response ? error.response.data : error.message);
      }
    },
    // 从后端获取好友列表
    async fetchUsers() {
      try {
        axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
        const response = await axios.get(`${baseURL}/api/v1/message/getfriends`);//这里还要根据当前用户获取消息
        
        // 根据后端返回的格式，提取 "data" 字段
        if (response.data.code === 200) {
          this.users = response.data.data.map(user => ({
            id: user.id,
            name: user.name,
            avatar: user.avatar
          }));
          console.log('好友列表获取成功:', this.users);
        } else {
          console.error('获取好友列表失败，后端返回错误:', response.data.message);
        }
      } catch (error) {
        console.error('获取好友列表失败:', error);
        console.error('请求 URL:', error.config?.url || '');
        console.error('错误详情:', error.response ? error.response.data : error.message);
      }
    },

    // 从后端获取消息列表
    async fetchMessages() {
      try {
        axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
        const response = await axios.post(
          `http://127.0.0.1:8080/api/v1/message/getmessages`, // 后端 API 地址
          {
            friend_id: this.activeFriendId,
          },
          {
            headers: {
              'Content-Type': 'application/json', // 设置 Content-Type 为 JSON
            }
          }
        );
    
        // 将后端返回的消息存储到状态中
        if (response.data && response.data.code === 200) {
          console.log("Response messages:", response.data.messages);
          this.messages = response.data.messages.map((msg) => ({
            id: msg.mid,                        // 对应后端的 ID 字段
            sender_id: msg.send_user_id,        // 映射到 SenderID
            receiver_id: msg.rev_user_id,       // 映射到 ReceiverID
            text: msg.content,                  // 映射到 Text
            timestamp: String(msg.create_time), 
          }));
          console.log("Mapped messages:", this.messages); // 检查是否提取成功
          console.log("currentUserID:",this.currentUserId);
          console.log("ActiveFriendID:",this.activeFriend);
        } else {
          console.error('后端返回错误:', response.data);
        }
      } catch (error) {
        console.error('获取消息列表失败:', error);
        console.error('请求 URL:', error.config.url);
        console.error('错误详情:', error.response ? error.response.data : error.message);
      }
    },

    // 设置当前选中的好友
    setActiveFriend(friendId) {
      this.activeFriendId = friendId
    },

    // 发送消息到后端
    async sendMessage(messageContent) {
      try {
        const requestData = {
          rev_user_id: this.activeFriendId,
          content: messageContent.text, 
        };
        axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
        const response = await axios.post(
          `http://127.0.0.1:8080/api/v1/message/send`,
          requestData, // 请求体
          {
            headers: {
              'Content-Type': 'application/json', 
            }
          }
        );
        if (response.data && response.data.code === 200) {
          // 将返回的新消息添加到本地状态
          this.messages.push({
            id: response.data.message.mid,
            sender_id: this.currentUserId,
            receiver_id: requestData.rev_user_id,
            text: requestData.content,
            timestamp: response.data.message.create_time,
          });
        } else {
          console.error('后端返回错误:', response.data);
        }
      } catch (error) {
        console.error('发送消息失败:', error);
        console.error('请求 URL:', error.config.url);
        console.error('错误详情:', error.response ? error.response.data : error.message);
      }
      axios.defaults.headers.common["Authorization"] = `Bearer ${localStorage.getItem("token")}`;
      await this.fetchMessages();
      console.log('Message sent and messages list updated.');
    },



    // 设置当前用户
    setCurrentUser(userId) {
      this.currentUserId = userId
    }
  }
})