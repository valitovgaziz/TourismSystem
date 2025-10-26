<template>
    <div class="container">
        <!-- Заголовок -->
        <h1>{{ t('messages.tripFreedomAdventure.name') }}</h1>

         <p>{{ t('messages.tripFreedomAdventure.text') }}</p>
    </div>
</template>

<script>
import { useI18n } from 'vue-i18n';

export default {
  setup() {
    const { t } = useI18n();
    return { t };
  }
}
</script>

<style scoped>
.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
    font-size: 32px;
    margin-bottom: 30px;
}

p {
    line-height: 1.6;
    font-size: 18px;
}

.additional-content {
    margin-top: 40px;
}

@media screen and (max-width: 600px) {
    .container {
        padding: 10px;
    }

    h1 {
        font-size: 24px;
    }

    p {
        font-size: 16px;
    }
}
</style>