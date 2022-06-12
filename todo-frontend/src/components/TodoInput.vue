<template>
  <div>
    <input id="todo-input" type="text" placeholder="Enter todo" v-model="todoText"/>
    <button id="todo-add-button" type="button" @click="addTodo">ADD</button>
  </div>
</template>

<script>
import { ADD_TODO } from '@/store/actions';

export default {
  name: 'TodoInput',
  data() {
    return {
      todoText: '',
    };
  },
  methods: {
    addTodo() {
      if (this.todoText !== '') {
        const todo = { text: this.todoText };
        this.$store.dispatch(ADD_TODO, todo);
      }
    },
  },
};
</script>

<style scoped>

</style>
