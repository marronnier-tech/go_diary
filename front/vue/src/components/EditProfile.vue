<template>
  <div class="profile">
    <div class="top-copy">
      <h1>プロフィール</h1>
      <p>公開したくない情報は空欄でかまいません。<br /></p>
    </div>
    <div class="whole-prof">
      <b-form @submit="SubmitProf">
        <b-form-group
          id="input-group-1"
          label="ハンドルネーム："
          label-for="input-1"
        >
          <b-form-input id="handle-name" v-model="UserInfo.HN"></b-form-input>
        </b-form-group>

        <b-form-group id="input-group-2" label="目標：" label-for="input-2">
          <b-form-input
            id="final-goal"
            v-model="UserInfo.FinalGoal"
          ></b-form-input>
        </b-form-group>

        <b-form-group
          id="input-group-3"
          label="自由プロフィール："
          label-for="input-3"
        >
          <b-form-textarea
            id="free-profile"
            v-model="UserInfo.Profile"
            rows="5"
          ></b-form-textarea>
        </b-form-group>

        <b-form-group id="input-group-4" label="Twitter：" label-for="input-2">
          <b-form-input
            id="handle-name"
            v-model="UserInfo.Twitter"
          ></b-form-input>
        </b-form-group>

        <b-form-group
          id="input-group-5"
          label="Instagram："
          label-for="input-2"
        >
          <b-form-input
            id="handle-name"
            v-model="UserInfo.Instagram"
          ></b-form-input>
        </b-form-group>

        <b-form-group id="input-group-6" label="Facebook：" label-for="input-2">
          <b-form-input
            id="handle-name"
            v-model="UserInfo.Facebook"
          ></b-form-input>
        </b-form-group>

        <b-form-group id="input-group-7" label="GitHub：" label-for="input-2">
          <b-form-input
            id="handle-name"
            v-model="UserInfo.Github"
          ></b-form-input>
        </b-form-group>

        <b-form-group
          id="input-group-8"
          label="その他URL："
          label-for="input-2"
        >
          <b-form-input id="handle-name" v-model="UserInfo.URL"></b-form-input>
        </b-form-group>

        <b-button type="submit" variant="info">Submit</b-button>
      </b-form>
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
  mounted: function () {
    axios.get("/profile").then((res) => {
      this.UserInfo = res.data.UserInfo;
      console.log(res.data);
    });
  },
  methods: {
    SubmitProf() {
      const params = new URLSearchParams();
      params.append("HN", this.UserInfo.HN);
      params.append("FinalGoal", this.UserInfo.FinalGoal);
      params.append("Profile", this.UserInfo.Profile);
      params.append("Twitter", this.UserInfo.Twitter);
      params.append("Instagram", this.UserInfo.Instagram);
      params.append("Facebook", this.UserInfo.Facebook);
      params.append("Github", this.UserInfo.Github);
      params.append("URL".this.UserInfo.URL);
      axios.patch("/profile", params).then((res) => {
        this.$router.push({ name: "Profile" });
      });
    },
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
a {
  color: #42b983;
}
</style>
