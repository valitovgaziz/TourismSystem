import Vue from 'vue';
import router from '@/router/index.js'; // Подключаем роутер
import store from '@/auth/stores/store.js';   // Подключаем стор
import VueAxios from 'vue-axios';
import axios from 'axios';
import VueAuth from '@websanova/vue-auth';
import Cookie from '@websanova/vue-auth/drivers/auth/cookie';
import HttpAxios from '@websanova/vue-auth/drivers/http/axios.1.x';
import Router from '@websanova/vue-auth/drivers/router/vue-router.2.x';

Vue.use(VueAxios, axios);

const options = {
  auth: Cookie,
  http: HttpAxios,
  router: Router,
};

Vue.use(VueAuth, options);

// Настройки API
Vue.axios.defaults.baseURL = 'http://yalarba.ru/api';

// Добавляем авторизацию через куки в заголовки запросов
Vue.axios.interceptors.request.use(config => {
  return config;
}, error => Promise.reject(error));

// Обработчик ошибок при получении ответа от сервера
Vue.axios.interceptors.response.use(response => response, error => {
  const { statusCode, message } = error.response.data;
  
  switch (statusCode) {
    case 401:
      store.dispatch('clearAuthData'); // Очистка данных об аутентификации
      router.push('/login');           // Переход на страницу логина
      break;
    default:
      console.error(message);          // Логирование ошибки
  }

  return Promise.reject(error);
});

export default options;