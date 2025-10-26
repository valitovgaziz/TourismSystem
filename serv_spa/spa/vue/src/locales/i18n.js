// i18n.js
import { createI18n } from 'vue-i18n';

// Загрузите языковые файлы
import en from '../locales/en.json';
import ru from '../locales/ru.json';
import bak from '../locales/bak.json';
import tat from '../locales/tat.json';

const messages = {
  en,
  ru,
  bak,
  tat,
};

const i18n = createI18n({
  legacy: false,
  locale: 'en', // Установите язык по умолчанию
  fallbackLocale: 'ru', // Язык, который будет использоваться, если перевод для текущего языка отсутствует
  messages,
});

export default i18n;