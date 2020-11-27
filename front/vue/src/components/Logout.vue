<template>
  <div class="logout">
    <b-button
      v-on:click="logout"
      :pressed="false"
      variant="info"
      class="edit-profile"
      >ログアウトする</b-button
    >
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Logout",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
  },
  methods: {
    logout() {
      axios
        .delete("/logout")
        .then((res) => {
          this.$parent.loginFlag = false;
          this.$router.push({ name: "Top" });
        })
        .catch((error) => {
          alert("エラーが発生しました");
        });
    },
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
