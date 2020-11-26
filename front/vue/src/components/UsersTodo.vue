<template>
  <div class="mypage">
    <div class="top-copy">
      <h1>{{ todos.User.UserName }}の日課</h1>
    </div>
    <ul>
      <li class="todo-set" v-for="todo in todos.TodoObj" :key="todo.TodoIndex">
        <div class="my-todoset">
          <h4 class="content">
            <b-icon icon="bookmark-star-fill" class="check"></b-icon
            >{{ todo.Content }}
          </h4>
        </div>
        <div class="achieved-info">
          <p class="count">実行日数：{{ todo.Count }}日</p>
          <p class="last-achieved">最終実行日：{{ todo.LastAchieved }}</p>
          <p class="created-at">{{ todo.CreatedAt }}</p>
        </div>
      </li>
    </ul>
    <br />
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",
  message: "Not yet",
  data() {
    return {
      todos: {
        User: {
          UserID: "0",
          UserName: "",
          UserHN: "",
          UserImg: "",
        },
        TodoArray: [],
      },
      content: "",
    };
  },
  computed: {
    id: function () {
      return this.$route.params.id;
    },
  },
  mounted: function () {
    axios.get("/todo/" + this.id).then((res) => {
      this.todos = res.data.Todo;
      console.log(res);
    });
  },
  methods: {},
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h4 {
  font-weight: normal;
  padding-top: 0.5em;
  margin-right: 1em;
}
.content {
  margin-bottom: 1.2em;
}
.achieved-info {
  padding-left: 1em;
}
a {
  color: #42b983;
}
.my-todoset {
  display: flex;
  justify-content: space-between;
}
.yatta {
  margin-right: 1em;
  padding: 0.5em;
  height: 2.5em;
  letter-spacing: 0.05em;
}
.check {
  color: #e46c86;
  margin-right: 0.5em;
}
.todo-add {
  margin-top: 4em;
  text-align: left;
}
.add-form {
  margin: 0.2em 2em 0.2em 2em;
}
h2 {
  font-size: 1.5em;
  margin-bottom: 0.8em;
}
</style>
