<template>
    <div class="menuH" v-on:click="toggleMenu">
        <HamburgerMenu :isMenuVisible="isMenuVisible" />
        <ul id="ul-m" :class="{ visible: isMenuVisible }">
            <li v-on:click="$refs.profile.$el.click()">
                <router-link to="/profile" ref="profile">
                    {{ t('messages.profile') }}
                </router-link>
            </li>
            <li v-on:click="$refs.settings.$el.click()">
                <router-link to="/settings" ref="settings">
                    {{ t('messages.settings') }}
                </router-link>
            </li>
            <li v-on:click="$refs.about.$el.click()">
                <router-link to="/about" ref="about">
                    {{ t('messages.about') }}
                </router-link>
            </li>
            <li v-on:click="$refs.support.$el.click()">
                <router-link to="/support" ref="support">
                    {{ t('messages.technicalSupport') }}
                </router-link>
            </li>
            <li v-on:click="$refs.feetback.$el.click()">
                <router-link to="/feetback" ref="feetback">
                    {{ t('messages.writeFeedback') }}
                </router-link>
            </li>
            <li v-on:click="$refs.registration.$el.click()">
                <router-link to="/registration" ref="registration">
                    {{ t('messages.registration') }}
                </router-link>
            </li>
            <li v-on:click="$refs.login.$el.click()">
                <router-link to="/login" ref="login">
                    {{ authStore.isAuthenticated ? t('messages.logout') : t('messages.login') }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
import { useI18n } from 'vue-i18n';
import HamburgerMenu from './toggleMenu.vue';
import { useAuthStore } from '@/auth/stores/auth.store';



export default {
    setup() {
        const { t } = useI18n();
        const authStore = useAuthStore();
        return { t, authStore };
    },
    name: 'menuH',
    components: {
        HamburgerMenu,
    },
    data() {
        return {
            isMenuVisible: false,
        }
    },
    mounted() {
        document.addEventListener('click', this.closeMenuIfClickedOutside)
    },
    beforeUnmount() {
        document.removeEventListener('click', this.closeMenuIfClickedOutside)
    },
    methods: {
        toggleMenu() {
            this.isMenuVisible = !this.isMenuVisible;
        },
        closeMenuIfClickedOutside(event) {
            const menuElement = this.$el;
            if (!menuElement.contains(event.target)) {
                this.isMenuVisible = false;
            }
        }
    }

}
</script>

<style scoped>
.menuH {
    height: 100%;
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    justify-content: center;
}

#ul-m {
    position: absolute;
    top: 3rem;
    right: 1rem;
    width: 25vw;
    min-width: 300px;
    visibility: hidden;
    border-radius: 1rem;
    box-shadow: 1px 2px 3px rgb(141, 225, 214);
    padding: 0.5rem;
    list-style-type: none;
    background-color: var(--background-color);  
}

li {
    width: 100%;
    padding: 0.3rem;
    text-decoration: none;
}

li * {
    text-decoration: none;
}

li:hover {
    cursor: pointer;
    background-color: rgb(92, 113, 113);
    border-radius: 0.3rem;
}

.visible {
    visibility: visible !important;
}
</style>
