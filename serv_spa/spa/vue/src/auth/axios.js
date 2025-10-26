// axios.js
import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://yalarba.ru/api/v1/',
  withCredentials: true, // Это важно для работы с HTTP-only куки
});

export default instance;