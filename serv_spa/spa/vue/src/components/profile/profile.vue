<template>
    <div class="profile">
      <!-- Заголовок профиля -->
      <header class="profile-header">
        <img :src="user.avatar" alt="Avatar" class="avatar" />
        <h1 class="name">{{ user.name }}</h1>
        <p class="email">{{ user.email }}</p>
      </header>
      
      <!-- Редактирование профиля -->
      <button @click="editUserInfo()" class="edit-profile-btn">Редактировать профиль</button>
  
      <!-- Форма поиска -->
      <form class="search-form">
        <input type="text" placeholder="Поиск..." v-model.trim="searchQuery" />
        <button @click.prevent="searchPlaces()">Найти</button>
      </form>
  
      <!-- Результат поиска -->
      <section v-if="filteredPlaces.length > 0 && searchQuery !== ''" class="search-results">
        <h2>Результаты поиска</h2>
        <ul>
          <li v-for="place in filteredPlaces" :key="place.id" class="place-item">
            <h3><a :href="place.editUrl">{{ place.title }}</a></h3>
            <p>{{ place.description }}</p>
          </li>
        </ul>
      </section>
  
      <!-- Список мест -->
      <section v-if="filteredPlaces.length > 0" class="places-list">
        <h2>Места</h2>
        <ul>
          <li v-for="place in filteredPlaces" :key="place.id" class="place-item">
            <h3><a :href="place.editUrl">{{ place.title }}</a></h3>
            <p>{{ place.description }}</p>
          </li>
        </ul>
      </section>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Profile',
    data() {
      return {
        user: {
          avatar: 'https://via.placeholder.com/150', // замените на реальный URL аватара
          name: 'Иван Иванов',
          email: 'ivan@example.com',
          places: [
            { id: 1, title: 'Место 1', description: 'Описание места 1', editUrl: '/edit-place/1' },
            { id: 2, title: 'Место 2', description: 'Описание места 2', editUrl: '/edit-place/2' }
          ]
        },
        searchQuery: ''
      };
    },
    computed: {
      filteredPlaces() {
        if (!this.searchQuery) {
          return this.user.places;
        }
        return this.user.places.filter(place => {
          return place.title.toLowerCase().includes(this.searchQuery.toLowerCase());
        });
      }
    },
    methods: {
      editUserInfo() {
        alert('Редактируем информацию о пользователе...');
      },
      searchPlaces() {
        console.log(`Поиск мест по запросу: ${this.searchQuery}`);
      }
    }
  };
  </script>
  
  <style scoped>
  .profile {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  
  .profile-header {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .avatar {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    object-fit: cover;
    margin-bottom: 10px;
  }
  
  .name {
    font-size: 24px;
    margin-top: 0;
  }
  
  .email {
    color: #666;
    font-size: 16px;
  }
  
  .edit-profile-btn {
    padding: 8px 12px;
    background-color: #4CAF50;
    color: white;
    border: none;
    cursor: pointer;
    margin-bottom: 20px;
  }
  
  .places-list h2,
  .search-results h2 {
    margin-top: 30px;
    margin-bottom: 10px;
  }
  
  .place-item {
    margin-bottom: 15px;
  }
  
  .search-form {
    max-width: 300px;
    width: 100%;
    margin-bottom: 20px;
  }
  
  .search-form input {
    width: calc(100% - 80px);
    padding: 8px;
    border: 1px solid #ccc;
    border-right: none;
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
  }
  
  .search-form button {
    width: 70px;
    padding: 8px;
    background-color: #2196F3;
    color: white;
    border: none;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
    cursor: pointer;
  }
  
  @media (max-width: 600px) {
    .profile {
      width: 90%;
    }
    
    .search-form {
      width: 100%;
    }
  }
  </style>