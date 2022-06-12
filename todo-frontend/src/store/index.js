import Vue from 'vue';
import Vuex from 'vuex';
import { actions } from '@/store/actions';
import { mutations } from '@/store/mutations';

Vue.use(Vuex);

export const state = {
  todoList: [],
};

export const getters = {
  // eslint-disable-next-line no-shadow
  todos: (state) => state.todoList,
};

export const modules = {

};

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
  modules,
});
