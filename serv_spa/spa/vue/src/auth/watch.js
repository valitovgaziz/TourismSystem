import { watch } from 'vue';
import { useAuthStore } from '../stores/auth.store';

export default {
  setup() {
    const authStore = useAuthStore();
    const { user, isAuthenticated } = storeToRefs(authStore);

    watch(isAuthenticated, (newVal) => {
      console.log('isAuthenticated changed:', newVal);
    });

    watch(user, (newVal) => {
      console.log('user changed:', newVal);
    });

    return {
      user,
      isAuthenticated,
    };
  },
};