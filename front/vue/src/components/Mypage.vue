<template>
  <div class="hello">
    Hello,{{ user }}<br />
    {{ todos }}
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
    axios.get("/mypage").then((res) => {
      this.todos = res.data;
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
