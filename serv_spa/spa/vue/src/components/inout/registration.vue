<template>
  <div class="register-form">
    <h3 class="form-name-h3">
      {{ t('messages.inout.registration') }}
    </h3>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="username">
          {{ t('messages.inout.name') }}:
        </label>
        <input v-model.trim="name" type="text" id="username" required />
      </div>

      <div class="form-group">
        <label for="email">
          {{ t('messages.inout.email') }}:
        </label>
        <input v-model.trim="email" type="email" id="email" required />
      </div>

      <div class="form-group">
        <label for="password">
          {{ t('messages.inout.password') }}:
        </label>
        <input v-model.trim="password" type="password" id="password" required />
      </div>

      <button type="submit">
        {{ t('messages.inout.registrationB') }}
      </button>
    </form>
  </div>
</template>

<script>
import { useI18n } from 'vue-i18n';
import { useAuthStore } from '@/auth/stores/auth.store';
import { useRouter } from 'vue-router';

export default {
  name: 'RegisterForm',
  setup() {
    const { t } = useI18n();
    const authStore = useAuthStore();
    const router = useRouter();
    return { t, authStore, router };
  },
  data() {
    return {
      name: '',
      email: '',
      password: '',
      isLoading: false
    };
  },
  methods: {
    async handleSubmit() {
      if (!this.isValid(this.name, this.email, this.password)) {
        alert("Пожалуйста, заполните все поля корректно.");
        return;
      }

      this.isLoading = true;
      
      try {
        // ИСПОЛЬЗУЕМ authStore.register вместо прямой отправки
        await this.authStore.register({
          name: this.name,
          email: this.email,
          password: this.password
        });
        
        // После успешной регистрации переходим в профиль
        this.router.push('/profile');
        
      } catch (error) {
        console.error('Registration error:', error);
        alert(error.message || 'Что-то пошло не так. Попробуйте еще раз.');
      } finally {
        this.isLoading = false;
      }
    },

    isValid(name, email, password) {
      if (name.length === 0 || email.length === 0 || password.length === 0) {
        return false;
      }

      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailRegex.test(email)) {
        return false;
      }

      if (password.length < 6) {
        return false;
      }

      return true;
    }
  }
};
</script>

<style scoped>
.form-name-h3 {
  margin-top: 0;
  padding-top: 0;
  height: 1.5rem;
}
.register-form {
  max-width: fit-content;
  padding: 1rem 2rem 2rem 2rem;
  border-radius: 1rem;
  box-shadow: 1px 2px 5px #609f7d;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
}

.form-group input {
  padding: 0.7rem;
  border-radius: 1rem;
  border: 1px solid #439c5f;
  background-color: var(--light-dark-background-color);
  color: var(--texgt-color);
  box-shadow: 1px 2px 3px #439c5f;
}

button {  
  margin-top: 1rem;
  padding: 0.7rem 1.4rem;
  background-color: var(--button-dark-background-color);
  color: white;
  border: none;
  border-radius: 1rem;
  cursor: pointer;
  box-shadow: 1px 2px 3px #a1c3ab;
  border: 1px solid rgb(124, 171, 156);
}

button:hover {
  box-shadow: 0px 0px 6px rgb(75, 103, 94);
}
</style>