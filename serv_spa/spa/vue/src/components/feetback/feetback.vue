<template>
  <div class="reviews-container">
    <h2>{{ t('messages.reviews.h2review') }}</h2>
    <h4>{{ t('messages.reviews.h4review') }}</h4>

    <!-- Форма для добавления отзыва -->
    <div class="add-review">
      <textarea v-model="newReview" :placeholder="t('messages.reviews.textareaplaceholder')" rows="4"></textarea>
      <button @click="submitReview" :disabled="!newReview.trim()">
        {{ t('messages.reviews.button') }}
      </button>
      <button @click="toggleShowAll" :disabled="reviews.length === 0">
        {{ showAll ? t('messages.reviews.hide') : t('messages.reviews.viewAll') }}
      </button>
    </div>

    <!-- Список отзывов -->
    <div class="reviews-list">
      <div v-for="review in displayedReviews" :key="review.id" class="review-item">
        <p>{{ review.text }}</p>
        <small>{{ formatDate(review.created_at) }}</small>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useI18n } from 'vue-i18n';

export default {
  setup() {
    const { t } = useI18n();
    return { t };
  },
  data() {
    return {
      reviews: [], // Список всех отзывов
      newReview: '', // Текст нового отзыва
      showAll: false, // Флаг для отображения всех отзывов
    };
  },
  computed: {
    // Вычисляемое свойство для отображения отзывов
    displayedReviews() {
      return this.showAll ? this.reviews : this.reviews.slice(0, 5);
    }
  },
  mounted() {
    this.loadReviews(); // Загружаем отзывы при монтировании компонента
  },
  methods: {
    // Загрузка отзывов из базы данных
    async loadReviews() {
      try {
        const response = await axios.get('https://yalarba.ru/api/reviews/reviews');
        this.reviews = response.data.sort((a, b) => {
          return new Date(b.created_at) - new Date(a.created_at); // Сортируем по убыванию
        });
      } catch (error) {
        console.error('Ошибка при загрузке отзывов:', error);
      }
    },
    // Отправка нового отзыва
    async submitReview() {
      if (!this.newReview.trim()) return;

      try {
        const response = await axios.post('https://yalarba.ru/api/reviews/reviews', {
          text: this.newReview,
        });
        this.reviews.unshift(response.data); // Добавляем новый отзыв в начало списка
        this.newReview = ''; // Очищаем текстовое поле
      } catch (error) {
        console.error('Ошибка при отправке отзыва:', error);
      }
    },
    // Форматирование даты
    formatDate(dateString) {
      const date = new Date(dateString);
      return date.toLocaleString(); // Форматируем дату в удобный формат
    },
    // Переключение между "Показать все" и "Скрыть"
    toggleShowAll() {
      this.showAll = !this.showAll; // Инвертируем значение флага
    }
  },
};
</script>

<style scoped>
.reviews-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.add-review textarea {
  width: 90%;
  padding: 10px;
  margin-bottom: 10px;
  border-radius: 5px;
  border: 1px solid #ccc;
  background-color: var(--light-dark-background-color);
  box-shadow: 1px 2px 3px rgb(61, 154, 96);
}

.add-review textarea::placeholder {
  color: var(--light-dark-text-color);
}

.add-review button {
  padding: 10px 20px;
  background-color: var(--button-dark-background-color);
  color: var(--text-color);
  border: none;
  border-radius: 0.7rem;
  cursor: pointer;
  margin-right: 10px;
  box-shadow: 1px 2px 3px rgb(61, 154, 96);
}

.add-review button:disabled {
  background-color: var(--disabled-dark-background-color);
  cursor: not-allowed;
}

.reviews-list {
  margin-top: 2rem;
  width: 90%;
}

.review-item {
  background: var(--light-dark-background-color);
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 10px;
  box-shadow: 1px 2px 3px rgb(61, 154, 96);
}

.review-item p {
  margin: 0;
}

.review-item small {
  color: var(--light-dark-text-color);
  font-size: 0.9em;
}
</style>