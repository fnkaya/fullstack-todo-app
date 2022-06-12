import { shallowMount } from '@vue/test-utils';
import TodoItem from '@/components/TodoItem.vue';

describe('TodoItem.vue tests', () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallowMount(TodoItem, {
      propsData: {
        todo: {},
      },
    });
  });

  it('should elements render', () => {
    const todoContainer = wrapper.find('#todo-item-container');
    expect(todoContainer.exists()).toBeTruthy();
    const todoTextEl = todoContainer.find('#todo-text');
    expect(todoTextEl.exists()).toBeTruthy();
    const todoCheckEl = todoContainer.find('#todo-done');
    expect(todoCheckEl.exists()).toBeTruthy();
  });
});
