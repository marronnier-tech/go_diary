<template>
  <div class="mypage">
    <div class="top-copy">
      <h1>Hello, {{ todos.User.UserName }}!</h1>
      <p>あなたの今の日課はこちらです。</p>
    </div>
    <ul>
      <li class="todo-set" v-for="todo in todos.TodoObj" :key="todo.TodoIndex">
        <div class="my-todoset">
          <h4 class="content">
            <b-icon icon="bookmark-star-fill" class="check"></b-icon
            >{{ todo.Content }}
          </h4>
          <b-button :pressed="false" variant="info" class="yatta"
            >Yatta!</b-button
          >
        </div>
        <div class="achieved-info">
          <p class="count">実行日数：{{ todo.Count }}日</p>
          <p class="last-achieved">最終実行日：{{ todo.LastAchieved }}</p>
          <p class="created-at">{{ todo.CreatedAt }}</p>
        </div>
      </li>
    </ul>
    <br />
    今日も一日頑張りましょう！<br />
    <br />
    <div class="todo-add">
      <h2>日課を追加する</h2>
      <div class="add-form">
        <b-form @submit.prevent="AddToDo">
          <b-form-input
            id="text-addtodo"
            placeholder="新しい日課"
          ></b-form-input>
          <b-button variant="info">追加</b-button>
        </b-form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",
  message: "Not yet",
  data() {
    return {
      todos: [],
    };
  },
  mounted: function () {
    axios.get("/mypage").then((res) => {
      this.todos = res.data.Todo;
      console.log(res);
    });
  },
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
