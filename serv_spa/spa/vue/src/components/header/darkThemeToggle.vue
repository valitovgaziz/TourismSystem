<template>
    <div class="theme-toggle">
      <label class="theme-toggle__label">
        <input
          type="checkbox"
          v-model="isDarkTheme"
          @change="toggleTheme"
          class="theme-toggle__input"
        />
        <span class="theme-toggle__slider"></span>
      </label>
    </div>
  </template>
  
  <script>
  import { ref, onMounted } from 'vue';
  
  export default {
    name: 'ThemeToggle',
    setup() {
      const isDarkTheme = ref(false);
  
      // Проверяем сохраненную тему в localStorage
      const savedTheme = localStorage.getItem('theme');
      if (savedTheme) {
        isDarkTheme.value = savedTheme === 'dark';
      } else {
        // Если тема не сохранена, используем системные настройки
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        isDarkTheme.value = prefersDark;
      }
  
      // Применяем тему при загрузке компонента
      const applyTheme = () => {
        const theme = isDarkTheme.value ? 'dark' : 'light';
        document.documentElement.setAttribute('data-theme', theme);
        localStorage.setItem('theme', theme);
      };
  
      // Переключение темы
      const toggleTheme = () => {
        applyTheme();
      };
  
      // Применяем тему при монтировании компонента
      onMounted(() => {
        applyTheme();
      });
  
      return {
        isDarkTheme,
        toggleTheme,
      };
    },
  };
  </script>
  
  <style scoped>
  .theme-toggle {
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 1px 2px 3px rgb(179, 210, 216);
    border-radius: 50%;
    height: 2rem;
    width: 2rem;
    padding: 0.5rem;
  }
  
  .theme-toggle:hover {
    box-shadow: 0px 0px 6px rgb(179, 210, 216);
  }

  .theme-toggle__label {
    display: flex;
    align-items: center;
    cursor: pointer;
  }
  
  .theme-toggle__input {
    display: none;
  }
  
  .theme-toggle__slider {
    width: 40px;
    height: 20px;
    background-color: #ccc;
    border-radius: 20px;
    position: relative;
    transition: background-color 0.3s;
  }
  
  .theme-toggle__slider::before {
    content: '';
    position: absolute;
    width: 16px;
    height: 16px;
    border-radius: 50%;
    background-color: white;
    top: 2px;
    left: 2px;
    transition: transform 0.3s;
  }
  
  .theme-toggle__input:checked + .theme-toggle__slider {
    background-color: #4caf50;
  }
  
  .theme-toggle__input:checked + .theme-toggle__slider::before {
    transform: translateX(20px);
  }
  </style>