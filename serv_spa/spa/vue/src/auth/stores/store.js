// store.js
import { createStore } from 'vuex';
import axios from 'axios';

export default createStore({
  state: {
    user: null,
    isAuthenticated: false,
  },
  mutations: {
    SET_USER(state, user) {
      state.user = user;
      state.isAuthenticated = !!user;
    },
  },
  actions: {
    async login({ commit }, credentials) {
      try {
        const response = await axios.post('/api/login', credentials);
        commit('SET_USER', response.data.user);
      } catch (error) {
        throw error;
      }
    },
    async logout({ commit }) {
      try {
        await axios.post('/api/logout');
        commit('SET_USER', null);
      } catch (error) {
        throw error;
      }
    },
    async checkAuth({ commit }) {
      try {
        const response = await axios.get('/api/user');
        commit('SET_USER', response.data.user);
      } catch (error) {
        commit('SET_USER', null);
      }
    },
  },
  getters: {
    isAuthenticated: (state) => state.isAuthenticated,
    user: (state) => state.user,
  },
});