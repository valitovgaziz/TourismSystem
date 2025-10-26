import './assets/main.css';

import { createApp } from 'vue';
import SPA_VUE_App from './App.vue';
import appRouter from './router/index.js';
import i18n from './locales/i18n.js';
import './assets/fonts.css';
import { createPinia } from 'pinia';

const spaAppVue3 = createApp(SPA_VUE_App);
const pinia = createPinia();

spaAppVue3.use(pinia);
spaAppVue3.use(appRouter);
spaAppVue3.use(i18n);

spaAppVue3.mount('#app');
