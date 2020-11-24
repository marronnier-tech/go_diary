<template>
  <div class="todo">
    Todo!<br />
    <!-- {{ todos.TodoArray }} --><br />
    <ul>
      <li v-for="todo in todos.TodoArray" :key="todo.TodoIbj">
        {{ todo.TodoObj.Content }}<br />
        {{ todo.TodoObj.CreatedAt }}<br />
        {{ todo.TodoObj.Count }}<br />
        {{ todo.TodoObj.LastAchieved }}<br />
        {{ todo.TodayAchieved }}<br />
        {{ todo.User.UserName }}
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",

  data() {
    return {
      todos: [],
    };
  },
  mounted: function () {
    axios.get("/todo").then((res) => {
      this.todos = res.data;
      console.log(res);
    });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
