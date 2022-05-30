import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import { createPinia } from 'pinia';
import { router } from "./router/index";
import axios from 'axios';
import App from './App.vue';

axios.defaults.baseURL = "/api";

const app = createApp(App);

app.use(ElementPlus, createPinia(), router)
    .provide("$axios", axios)
    .mount('#app')
