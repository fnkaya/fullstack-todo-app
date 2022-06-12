import { shallowMount } from '@vue/test-utils';
import TodoInput from '@/components/TodoInput.vue';
import { ADD_TODO } from '@/store/actions';

const factory = (options) => shallowMount(TodoInput, {
  data() {
    return {
      todoText: '',
    };
  },
  mocks: {
    $store: {
      dispatch: jest.fn(),
    },
  },
  ...options,
});

describe('TodoInput.vue tests', () => {
  it('should elements render', () => {
    const wrapper = factory();
    const inputEl = wrapper.find('#todo-input');
    const buttonEl = wrapper.find('#todo-add-button');
    expect(inputEl.exists()).toBeTruthy();
    expect(buttonEl.exists()).toBeTruthy();
    expect(buttonEl.text()).toEqual('ADD');
  });

  it('should data update when input changed', () => {
    const wrapper = factory();
    const inputEl = wrapper.find('#todo-input');
    const todoText = 'do something not stupid';
    inputEl.setValue(todoText);
    expect(wrapper.vm.todoText).toMatch(todoText);
  });

  it('if text not empty should delegate dispatch when button clicked', async () => {
    const mockDispatch = jest.fn();
    const wrapper = factory({
      mocks: {
        $store: {
          dispatch: mockDispatch,
        },
      },
    });
    const inputEl = wrapper.find('#todo-input');
    const todoText = 'do something not stupid';
    inputEl.setValue(todoText);
    const buttonEl = wrapper.find('#todo-add-button');
    await buttonEl.trigger('click');
    const todoObj = { text: todoText };
    expect(mockDispatch).toHaveBeenCalledWith(ADD_TODO, todoObj);
  });

  it('if text is empty should not delegate dispatch when button clicked', async () => {
    const mockDispatch = jest.fn();
    const wrapper = factory({
      mocks: {
        $store: {
          dispatch: mockDispatch,
        },
      },
    });
    const buttonEl = wrapper.find('#todo-add-button');
    await buttonEl.trigger('click');
    expect(mockDispatch).toBeCalledTimes(0);
  });
});
