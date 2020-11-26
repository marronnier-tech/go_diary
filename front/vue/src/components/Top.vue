<template>
  <div class="top">
    <p><router-link to="/todo">みんなのToDoを覗いてみる</router-link></p>
    <p><router-link to="/login">login</router-link></p>
    <p><router-link to="/register">register</router-link></p>
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
a {
  color: #42b983;
}
</style>
