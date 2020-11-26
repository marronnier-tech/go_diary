<template>
  <div id="service">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-brand to="/">ToDotCom</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item to="/todo">みんなのToDo</b-nav-item>
          <b-nav-item to="/mypage">My ToDo</b-nav-item>
          <b-nav-item to="/mygoal">My Goaled</b-nav-item>
          <b-nav-item to="/profile">My Profile</b-nav-item>
          <b-nav-item to="/logout">Logout</b-nav-item>
          <b-nav-item href="#" disabled>Disabled</b-nav-item>
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
export default {
  name: "Service",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
    };
  },
};
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
