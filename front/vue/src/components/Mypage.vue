<template>
  <div class="hello">
    Hello,{{ user }}<br />
    {{ todoobj }}
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Todo",

  data() {
    return {
      user: null,
      todoobj: [],
    };
  },
  mounted: function () {
    axios.get("/todo/mypage").then((res) => {
      this.todos = res.data.Todo.TodoObj;
      console.log(res);
    });
  },
};
</script>
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
