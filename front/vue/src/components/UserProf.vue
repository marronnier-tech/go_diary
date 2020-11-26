<template>
  <div class="mypage">
    <div class="top-copy">
      <h1>＠{{ UserInfo.Name }}のプロフィール</h1>
      <p>
        <b-icon icon="award" class="award"></b-icon>ゴール数：{{
          UserInfo.GoaledCount
        }}
      </p>
    </div>
    <div class="whole-prof">
      <div class="prof-item">
        <h2>ユーザー名</h2>
        <p>{{ UserInfo.Name }}</p>
      </div>

      <div class="prof-item">
        <h2>ハンドルネーム</h2>
        <p>{{ UserInfo.HN }}</p>
      </div>
      <div class="prof-item">
        <h2>目標</h2>
        <p>{{ UserInfo.FinalGoal }}</p>
      </div>
      <div class="prof-item">
        <h2>自由プロフィール</h2>
        <p>{{ UserInfo.Profile }}</p>
      </div>
      <div class="prof-item">
        <h2>Twitter</h2>
        <p>
          <a
            v-bind:href="'https://twitter.com/' + UserInfo.Twitter"
            target="_blank"
            >@{{ UserInfo.Twitter }}</a
          >
        </p>
      </div>
      <div class="prof-item">
        <h2>Instagram</h2>
        <p>
          <a
            v-bind:href="'https://instagram.com/' + UserInfo.Instagram"
            target="_blank"
            >@{{ UserInfo.Instagram }}</a
          >
        </p>
      </div>
      <div class="prof-item">
        <h2>Facebook</h2>
        <p>
          <a
            v-bind:href="'https://facebook.com/' + UserInfo.Facebook"
            target="_blank"
            >@{{ UserInfo.Facebook }}</a
          >
        </p>
      </div>
      <div class="prof-item">
        <h2>Github</h2>
        <p>
          <a
            v-bind:href="'https://github.com/' + UserInfo.Github"
            target="_blank"
            >@{{ UserInfo.Github }}</a
          >
        </p>
      </div>
      <div class="prof-item">
        <h2>その他URL</h2>
        <p>
          <a v-bind:href="'' + UserInfo.URL" target="_blank">{{
            UserInfo.URL
          }}</a>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "Profile",
  data() {
    return {
      UserInfo: "",
    };
  },
  computed: {
    id: function () {
      return this.$route.params.id;
    },
  },
  mounted: function () {
    axios.get("/profile/" + this.id).then((res) => {
      this.UserInfo = res.data.UserInfo;
      console.log(res.data);
    });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.whole-prof {
  text-align: left;
  margin: 5% 0% 0% 0%;
}
.prof-item {
  margin-bottom: 3em;
}
.prof-item p {
  margin-left: 1.5em;
  margin-top: 0.8em;
}
h2 {
  font-weight: normal;
  font-size: 1.3em;
  border-bottom: dashed 1px #4fa4b3;
  padding: 0em 0em 0.2em 0.5em;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
</style>
