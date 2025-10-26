<template>
  <div class="developers">
    <!-- Заголовок -->
    <h1 class="title">
      {{ t('messages.aboutL.developers') }}
    </h1>

    <!-- Список разработчиков -->
    <ul class="developer-list">
      <li v-for="dev in developers" :key="dev.name" class="developer-item">
        <a :href="dev.profileUrl" target="_blank" rel="noopener noreferrer">
          {{ dev.nickname }} ({{ dev.name }})
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import { useI18n } from 'vue-i18n';
export default {
  setup() {
    const { t } = useI18n();
    return { t };
  },
  name: 'DevelopersList',
  data() {
    return {
      developers: [
        { nickname: 'valitovgaziz', name: 'Валитов Газиз', profileUrl: 'https://github.com/valitovgaziz' }
        // Добавьте больше разработчиков по аналогии...
      ]
    }
  }
}
</script>

<style scoped>
.developers {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
}

.title {
  text-align: center;
  font-size: 24px;
  margin-bottom: 20px;
}

.developer-list {
  list-style-type: none;
  padding-left: 0;
}

.developer-item {
  margin-bottom: 10px;
}

.developer-item a {
  color: #007bff;
  text-decoration: none;
}

.developer-item a:hover {
  text-decoration: underline;
}
</style>