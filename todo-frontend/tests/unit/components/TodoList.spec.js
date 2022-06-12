import { shallowMount } from '@vue/test-utils';
import TodoList from '@/components/TodoList.vue';
import TodoItem from '@/components/TodoItem.vue';
import { FETCH_TODOS } from '@/store/actions';

const factory = (options) => shallowMount(TodoList, {
  mocks: {
    $store: {
      getters: { todos: [] },
      dispatch: jest.fn(),
    },
  },
  ...options,
});

describe('TodoList.vue tests', () => {
  it('should elements render', () => {
    const wrapper = factory();
    const todoListContainer = wrapper.find('#todo-list-container');
    expect(todoListContainer.exists()).toBeTruthy();
  });

  it('should todo item components correctly', async () => {
    const mockResponse = [
      {
        id: 1,
        text: 'write test',
        done: true,
      },
      {
        id: 2,
        text: 'write code',
        done: false,
      },
      {
        id: 3,
        text: 'make a pipeline',
        done: false,
      },
    ];
    const mockDispatch = jest.fn();
    const wrapper = factory({
      mocks: {
        $store: {
          getters: { todos: mockResponse },
          dispatch: mockDispatch,
        },
      },
    });

    const todoItemComps = wrapper.findAllComponents(TodoItem);
    expect(mockDispatch).toHaveBeenCalledWith(FETCH_TODOS);
    expect(todoItemComps).toHaveLength(mockResponse.length);
  });
});
