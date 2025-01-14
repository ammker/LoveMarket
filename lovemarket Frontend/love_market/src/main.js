import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import request from './request/request.js'
import { useChatStore } from './store/chatStore'
import { createPinia } from 'pinia'

const pinia = createPinia()
const app = createApp(App)
// 使用 app.config.globalProperties 注册全局属性,不能修改prototype
app.config.globalProperties.$request = request;
app.use(store)
app.use(VueAxios, axios)
app.use(router)
app.use(pinia)
app.mount('#app')
