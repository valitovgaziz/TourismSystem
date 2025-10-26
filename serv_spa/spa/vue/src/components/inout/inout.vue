<template>
    <div v-if="isLoad">{{ t('messages.load') }} . . .</div>
    <div v-else-if="authStore.isAuthenticated" class="logout">
        <div>
            <p>Good by {{ authStore.user.name }}!</p>
            <button @click="logout">Logout</button>
        </div>
    </div>
    <div v-else class="login-form">
        <h1>
            {{ t('messages.inout.login') }}
        </h1>
        <form @submit.prevent="handleSubmit">
            <div class="form-group">
                <label for="email">
                    {{ t('messages.inout.email') }}:
                </label>
                <input v-model.trim="email" type="email" id="email" required :placeholder="t('messages.inout.email')" />
            </div>

            <div class="form-group">
                <label for="password">
                    {{ t('messages.inout.password') }}:
                </label>
                <input v-model.trim="password" type="password" id="password" required
                    :placeholder="t('messages.inout.password')" />
            </div>

            <button type="submit">
                {{ t('messages.inout.login') }}
            </button>
        </form>
        <router-link to="registration" class="toRegistration">to Registration</router-link>
    </div>
</template>

<script setup>
import { useI18n } from 'vue-i18n';
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/auth/stores/auth.store';
import { useRouter } from 'vue-router';

const email = ref('');
const isValidE = ref(true);

const password = ref('');
const isValidP = ref(true);

const { t } = useI18n();
const router = useRouter();
const authStore = useAuthStore();
let isLoad = ref(false);


onMounted(async () => {
    isLoad = true;
    authStore.checkAuth().catch(error => console.error('Check auth failed', error));
    isLoad = false;
});

async function handleSubmit() {
    isLoad = true;
    if (!isValidEmail(email.value) || !isValidPassword(password.value)) {
        alert("Пожалуйста, заполните все поля корректно.");
        isLoad = false;
        return;
    }
    try {
        await authStore.login({ email: email.value, password: password.value });
        isLoad = false;
        router.push('/');
    } catch (error) {
        console.error(error);
        alert('Неверный email или пароль. Попробуйте снова.');
    }
}

function isValidPassword(assword) {
    if (password.length === 0) {
        return false;
    }

    if (password.length < 6) {
        return false;
    }
    return true;
}

function isValidEmail(email) {
    if (email.length === 0) {
        return false;
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        return false;
    }
    return true;
}

function logout() {
    isLoad = true;
    authStore.logout();
    isLoad = false;
    router.push('/');
}
</script>


<style scoped>
.login-form {
    padding: 0 2rem 2rem 2rem;
    border: 1px solid #ccc;
    border-radius: 1rem;
    box-shadow: 1px 2px 3px #609f7d;
}

.form-group {
    margin-bottom: 1rem;
}

.form-group label {
    display: block;
}

.form-group input {
    padding: 0.7rem;
    border: none;
    border-radius: 1rem;
    background-color: var(--light-dark-background-color);
    color: var(--text-color);
    box-shadow: 1px 2px 3px #609f7d;
}

button {
    margin-top: 0.7rem;
    padding: 0.7rem 1.7rem;
    background-color: var(--button-dark-background-color);
    color: var(--text-color);
    border: none;
    border-radius: 1rem;
    cursor: pointer;
    box-shadow: 1px 2px 3px #609f7d;
}

button:hover {
    box-shadow: 0 0 6px #609f7d;
}

.toRegistration {
    display: inline-block;
    margin-top: 1rem;
    padding: 0.5rem 1.5rem;
    background-color: transparent;
    color: var(--text-color);
    border: 1px solid rgba(96, 159, 125, 0.5);
    border-radius: 1rem;
    cursor: pointer;
    text-decoration: none;
    font-size: 0.9rem;
    transition: all 0.3s ease;
    opacity: 0.8;
}

.toRegistration:hover {
    opacity: 1;
    box-shadow: 0 0 4px rgba(96, 159, 125, 0.4);
    background-color: rgba(96, 159, 125, 0.05);
    text-decoration: none;
}
</style>