import { state } from '@/store';
import { mutations, SET_TODOS } from '@/store/mutations';
import { actions, ADD_TODO, FETCH_TODOS } from '@/store/actions';
import flushPromises from 'flush-promises';
// eslint-disable-next-line import/no-named-as-default
import API from '@/api';

let mockState;
const mockTodoList = [
  { id: 1, text: 'todo-1', done: false },
  { id: 2, text: 'todo-2', done: false },
];
jest.mock('@/api');

describe('mutation tests', () => {
  beforeEach(() => {
    mockState = { ...state };
  });

  it('should set todos successfully', () => {
    mutations[SET_TODOS](mockState, mockTodoList);
    expect(mockState.todoList).toStrictEqual(mockTodoList);
  });
});

describe('actions tests', () => {
  beforeEach(() => {
    mockState = { ...state };
  });

  it('should delegate SET_TODOS mutation', async () => {
    API.getTodos.mockResolvedValue(mockTodoList);
    const commit = jest.fn();
    actions[FETCH_TODOS]({ commit });
    await flushPromises();
    expect(commit).toHaveBeenCalledWith(SET_TODOS, mockTodoList);
  });

  it('should delegate FETCH_TODOS mutation', async () => {
    API.addTodo.mockResolvedValue(201);
    const dispatch = jest.fn();
    const todoObj = { text: 'new todo' };
    actions[ADD_TODO]({ dispatch }, todoObj);
    await flushPromises();
    expect(dispatch).toHaveBeenCalledWith(FETCH_TODOS);
  });
});
