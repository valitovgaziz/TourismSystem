<template>
    <div class="place-details">
      <!-- Заглавная часть с фото и основной информацией -->
      <header class="place-header">
        <img src="../images/restObject.jpg" alt="Фото места" class="place-image" />
        <div class="place-info">
          <h1 class="place-name">{{ place.name }}</h1>
          <span class="place-type">Тип: {{ place.type }}</span>
          <span class="place-coordinates">Координаты: {{ place.coordinates }}</span>
        </div>
      </header>
  
      <!-- Контакты -->
      <section class="place-contacts">
        <h2>Контакты</h2>
        <ul>
          <li v-for="contact in place.contacts" :key="contact.id">
            <strong>{{ contact.label }}:</strong> {{ contact.value }}
          </li>
        </ul>
      </section>
  
      <!-- Отзывы -->
      <section class="place-reviews">
        <h2>Отзывы</h2>
        <ul>
          <li v-for="review in place.reviews" :key="review.id">
            <p><strong>Автор:</strong> {{ review.author }}</p>
            <p><strong>Отзыв:</strong> {{ review.text }}</p>
          </li>
        </ul>
      </section>
  
      <!-- Кнопка для редактирования -->
      <footer class="place-actions">
        <router-link :to="`/edit/${place.id}`" class="btn-edit">Редактировать</router-link>
      </footer>
    </div>
  </template>
  
  <script>
  export default {
    name: 'PlaceDetails',
    props: {
      place: Object
    }
  };
  </script>
  
  <style scoped>
  .place-details {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
  
  .place-header {
    display: flex;
    gap: 20px;
    align-items: center;
  }
  
  .place-image {
    width: 200px;
    height: 200px;
    object-fit: cover;
    border-radius: 10px;
  }
  
  .place-info {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  
  .place-name {
    font-size: 24px;
    margin: 0;
  }
  
  .place-type,
  .place-coordinates {
    font-size: 18px;
    color: #555;
  }
  
  .place-contacts,
  .place-reviews {
    border: 1px solid #ddd;
    padding: 20px;
    border-radius: 10px;
  }
  
  .place-contacts li,
  .place-reviews li {
    margin-bottom: 10px;
  }
  
  .btn-edit {
    padding: 10px 20px;
    background-color: #4CAF50;
    color: white;
    text-decoration: none;
    border-radius: 5px;
    display: inline-block;
    text-align: center;
  }
  
  @media (min-width: 768px) {
    .place-details {
      gap: 40px;
      padding: 40px;
    }
  
    .place-header {
      gap: 40px;
    }
  
    .place-image {
      width: 300px;
      height: 300px;
    }
  
    .place-name {
      font-size: 32px;
    }
  
    .place-type,
    .place-coordinates {
      font-size: 22px;
    }
  
    .place-contacts,
    .place-reviews {
      padding: 30px;
    }
  
    .place-contacts li,
    .place-reviews li {
      margin-bottom: 15px;
    }
  
    .btn-edit {
      padding: 15px 30px;
    }
  }
  </style>