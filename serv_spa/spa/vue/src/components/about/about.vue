<template>
    <div class="about">
        <FTA />
        <p>
            {{ t('messages.aboutL.donat-request') }}
        </p>
        <button v-on:click="donat" class="donat-coffee">
            {{ t('messages.aboutL.donat-coffee') }}
        </button>
        <img src="../images/photo_2025-01-25_05-57-24.jpg" alt="Донат QRcod" class="QRcod">
    </div>
</template>


<script>
import { useI18n } from 'vue-i18n';
import FTA from './filosofy.vue';

export default {
    setup() {
        const { t } = useI18n();
        return { t };
    },
    methods: {
        donat() {
            window.open("https://www.tinkoff.ru/rm/r_fFXNVwPxmH.VMEQXkfWxO/Ho4MN83486", "_blank");
        }
    },
    components: {
        FTA,
    }
}
</script>

<style scoped>
.about {
    height: fit-content;
    display: flex;
    flex-direction: column;
    align-items: center;
    filter: none;
    padding: 1rem;
    border-radius: 1rem;
    margin: 1rem;
}


.donat-coffee {
    margin: 1rem;
    background-color: rgb(89, 219, 109);
    box-shadow: 1px 4px 8px 0 rgba(27, 45, 37, 0.2);
    border-radius: 1rem;
}

.donat-coffee:hover {
    cursor: pointer;
    background-color: rgb(10, 150, 50);
    color: white;
    box-shadow: 1px 4px 8px rgb(10, 150, 50);
}

.QRcod {
    height: 200px;
    display: flexbox;
    border-radius: 0.5rem;
}
</style>
