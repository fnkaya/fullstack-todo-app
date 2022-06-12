export const SET_TODOS = 'setTodos';

export const mutations = {
// eslint-disable-next-line no-shadow
  [SET_TODOS]: (state, todoList) => {
    state.todoList = todoList;
  },
};
