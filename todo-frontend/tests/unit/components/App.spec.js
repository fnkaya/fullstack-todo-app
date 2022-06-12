import { shallowMount } from '@vue/test-utils';
import App from '@/App.vue';
import TodoInput from '@/components/TodoInput.vue';
import TodoList from '@/components/TodoList.vue';

describe('App.vue tests', () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallowMount(App);
  });

  it('should components exists', () => {
    const todoInputComp = wrapper.findComponent(TodoInput);
    expect(todoInputComp.exists()).toBeTruthy();
    const todoListComp = wrapper.findComponent(TodoList);
    expect(todoListComp.exists()).toBeTruthy();
  });
});
