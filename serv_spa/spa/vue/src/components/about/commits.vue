<template>
  <div v-if="loading">
    {{ t('messages.commits.download') }}
  </div>
  <ul v-else-if="error">
    {{ t('messages.commits.errorOc') }}: {{ error }}
  </ul>
  <ul v-else>
    <h2>Commits</h2>
    <li v-for="commit in commits" :key="commit.sha">
      <i>
        {{ commit.author.login }}
        {{ commit.commit.author.date.slice(0, 10) }}
        {{ commit.commit.message }}
      </i>
    </li>
  </ul>
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
      loading: true,
      commits: [],
      error: null,
    };
  },
  mounted() {
    this.fetchCommits();
  },
  methods: {
    async fetchCommits() {
      try {
        const response = await axios.get('https://api.github.com/repos/valitovgaziz/spa_/commits', {
          headers: {
            Accept: 'application/vnd.github.v3+json',
          },
        });
        this.commits = response.data;
      } catch (err) {
        this.error = err.message || 'Произошла неизвестная ошибка';
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
ul {
  max-width: 1024px;
  list-style-type: none;
  padding: 0;
}

li {
  margin: 0.5rem 0;
  white-space: pre-wrap;
  width: 100%;
}
</style>x``