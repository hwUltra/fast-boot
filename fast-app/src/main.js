import { createSSRApp } from 'vue';
import App from './App.vue';

import store from '@/store';
import { createPinia } from 'pinia';
import { createUnistorage } from 'pinia-plugin-unistorage'

import api from '@/api/index.js';
// 全局组件引入
import globalComponent from '@/components/index.js';
// 抽取公用的实例 - 操作成功与失败消息提醒内容等
import mixin from '@/utils/mixin.js';
// 全局过滤器
import { filters } from '@/utils/filters.js';

// 地图工具类
import qqmap from '@/utils/qqmap.js';

export function createApp() {
  const app = createSSRApp(App);

  const pinia = createPinia();
  pinia.use(createUnistorage());
  app.use(pinia);

  app.config.globalProperties.$store = store;
  app.config.globalProperties.$api = api;
  app.mixin(mixin);
  app.config.globalProperties.$filters = filters;
  app.config.globalProperties.$qqmap = qqmap;

  return {
    app,
  };
}
