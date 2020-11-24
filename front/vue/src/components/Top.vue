<template>
  <div class="top">
    <router-link to="/login">login</router-link>
    <router-link to="/register">register</router-link>
  </div>
</template>

<script>
export default {
  name: "Top",
  data() {
    return {
      msg: " ",
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.top {
  padding-top: 100px;
}
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
