import { SET_TODOS } from '@/store/mutations';
// eslint-disable-next-line import/no-named-as-default
import API from '@/api';

export const FETCH_TODOS = 'fetchTodos';
export const ADD_TODO = 'addTodo';

export const actions = {
  [FETCH_TODOS]: ({ commit }) => {
    API.getTodos().then((todos) => commit(SET_TODOS, todos));
  },

  [ADD_TODO]: ({ dispatch }, payload) => {
    API.addTodo(payload).then((statusCode) => statusCode === 201 && dispatch(FETCH_TODOS));
  },
};
