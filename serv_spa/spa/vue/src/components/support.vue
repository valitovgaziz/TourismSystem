<template>
    <div class="support-form">
      <h2>Техническая поддержка</h2>
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="email">Ваш email:</label>
          <input
            type="email"
            id="email"
            v-model="formData.email"
            required
            placeholder="Введите ваш email"
          />
        </div>
        <div class="form-group">
          <label for="message">Сообщение:</label>
          <textarea
            id="message"
            v-model="formData.text"
            required
            placeholder="Опишите вашу проблему"
          ></textarea>
        </div>
        <button type="submit" :disabled="isLoading">
          {{ isLoading ? 'Отправка...' : 'Отправить' }}
        </button>
        <p v-if="successMessage" class="success-message">{{ successMessage }}</p>
        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        formData: {
          email: '',
          text: '',
        },
        isLoading: false,
        successMessage: '',
        errorMessage: '',
      };
    },
    methods: {
      async submitForm() {
        this.isLoading = true;
        this.successMessage = '';
        this.errorMessage = '';
  
        try {
          // Отправка данных на API
          console.log(this.formData);
          const response = await axios.post('https://yalarba.ru/api/support/support', this.formData);
  
          if (response.status === 201) {
            this.successMessage = 'Сообщение успешно отправлено!';
            this.formData.email = '';
            this.formData.text = '';
          } else {
            this.errorMessage = 'Произошла ошибка при отправке сообщения.';
          }
        } catch (error) {
          this.errorMessage = 'Произошла ошибка при отправке сообщения.';
        } finally {
          this.isLoading = false;
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .support-form {
    max-width: 700px;
    margin: 0 auto;
    padding: 1rem;
    border-radius: 1rem;
  }
  
  h2 {
    text-align: center;
    margin-bottom: 2rem;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  input,
  textarea {
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 1rem;
    font-size: 16px;
    background-color: var(--light-dark-background-color);
    box-shadow: 1px 2px 3px rgb(0, 255, 162);
  }

  input::placeholder,
  textarea::placeholder {
    color: var(--light-dark-text-color);
  }
  
  textarea {
    resize: vertical;
    min-height: 150px;
  }
  
  button {
    width: 50%;
    padding: 10px;
    color: var(--light-dark-text-color);
    background-color: var(--light-dark-background-color);
    border: none;
    border-radius: 0.5rem;
    font-size: 16px;
    cursor: pointer;
    box-shadow: 1px 2px 3px rgb(100, 199, 153);
  }
  
  button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  
  .success-message {
    color: green;
    text-align: center;
    margin-top: 10px;
  }
  
  .error-message {
    color: red;
    text-align: center;
    margin-top: 10px;
  }
  
  @media (max-width: 600px) {
    .support-form {
      padding: 15px;
    }
  
    input,
    textarea {
      font-size: 14px;
    }
  
    button {
      font-size: 14px;
    }
  }
  </style>