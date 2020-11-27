<template>
  <div id="service">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand to="/">ToDotCom</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav v-if="loginFlag">
          <b-nav-item to="/todo">みんなのToDo</b-nav-item>
          <b-nav-item to="/mypage">My ToDo</b-nav-item>
          <b-nav-item to="/mygoal">My Goaled</b-nav-item>
          <b-nav-item to="/profile">My Profile</b-nav-item>
          <b-nav-item to="/question">Q &amp; A</b-nav-item>
          <b-nav-item to="/logout">Logout</b-nav-item>
        </b-navbar-nav>

        <b-navbar-nav v-else>
          <b-nav-item to="/todo">みんなのToDo</b-nav-item>
          <b-nav-item to="/login">Login</b-nav-item>
          <b-nav-item to="/login">Register</b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input
              size="sm"
              class="mr-sm-2"
              placeholder="近日公開"
            ></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit"
              >検索</b-button
            >
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <div class="container">
      <br />

      <router-view></router-view>
    </div>
  </div>
</template>


<script>
import axios from "axios";
export default {
  name: "Service",
  data() {
    return {
      loginFlag: "",
    };
  },
  mounted: function () {
    this.getFlag();
  },
  methods: {
    getFlag() {
      axios.get("/adminflag").then((res) => {
        this.loginFlag = res.data.LoginFlag;
      });
    },
  },
};
</script>
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  color: #42b983;
}
.container {
  padding: 1em 10% 5em 10%;
}
</style>
