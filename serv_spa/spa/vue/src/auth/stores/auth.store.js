// src/auth/stores/auth.store.js
import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';
import AuthService from '../services/auth.service.js';
import { jwtDecode } from 'jwt-decode';

export const useAuthStore = defineStore('auth', () => {
  const user = reactive({name: '', email: '', id: 0, token: ''});
  const isAuthenticated = ref(false);

  // Восстановление из localStorage при инициализации
  const initFromStorage = () => {
    const storedToken = localStorage.getItem('token');
    const storedUser = localStorage.getItem('user');
    
    if (storedToken && storedUser) {
      try {
        user.token = storedToken;
        const userData = JSON.parse(storedUser);
        user.name = userData.name;
        user.email = userData.email;
        user.id = userData.id;
        isAuthenticated.value = true;
      } catch (error) {
        console.error('Error restoring from storage:', error);
        logout();
      }
    }
  };

  // Вызываем при создании store
  initFromStorage();

  // ДОБАВЬТЕ ЭТОТ МЕТОД - регистрация
  const register = async (userData) => {
    try {
      const response = await AuthService.register(userData);
      
      // Если сервер возвращает токен при регистрации
      if (response.token) {
        const decodedToken = jwtDecode(response.token);
        alert(decodedToken.name)
        user.name = decodedToken.user?.name || userData.name;
        user.id = decodedToken.user?.id || 0;
        user.email = decodedToken.user?.email || userData.email;
        isAuthenticated.value = true;
        user.token = response.token;

        // Сохраняем в localStorage
        localStorage.setItem('token', response.token);
        localStorage.setItem('user', JSON.stringify({
          username: user.name,
          email: user.email,
          id: user.id
        }));
      }
      
      return response;
    } catch (error) {
      console.error('Registration failed', error);
      throw error;
    }
  };

  const login = async (credentials) => {
    try {
      const response = await AuthService.login(credentials);
      const decodedToken = jwtDecode(response.token);
      user.name = decodedToken.user?.name || '';
      user.id = decodedToken.user?.id || 0;
      user.email = decodedToken.user?.email || credentials.email;
      isAuthenticated.value = true;
      user.token = response.token;

      // Сохраняем в localStorage
      localStorage.setItem('token', response.token);
      localStorage.setItem('user', JSON.stringify({
        name: user.name,
        email: user.email,
        id: user.id
      }));

    } catch (error) {
      console.error('Login failed', error);
      throw error;
    }
  };

  const logout = () => {
    isAuthenticated.value = false;
    user.name = '';
    user.token = '';
    user.email = '';
    user.id = 0;

    // Удаляем из localStorage
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  };

  const checkAuth = async () => {
    try {
      const token = user.token || localStorage.getItem('token');
      if (token) {
        try {
          const response = await AuthService.checkAuth(token);
          // Обновляем данные пользователя
          user.name = response.user?.name || user.name;
          user.id = response.user?.id || user.id;
          user.email = response.user?.email || user.email;
          isAuthenticated.value = true;
        } catch (error) {
          console.error('Token validation failed:', error);
          logout();
        }
      }
    } catch (error) {
      console.error('Check auth failed', error);
      throw error;
    }
  };

  // ВАЖНО: добавьте register в return
  return { user, isAuthenticated, register, login, logout, checkAuth };
});