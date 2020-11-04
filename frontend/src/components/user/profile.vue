<template>
    <div>
        <h1>Profile Page</h1>
        <br><br><br>


        <div v-if="user" class="text item">
            <el-card class="box-card">
                <div slot="header" class="clearfix">
                    <span>User Info</span>
                    <el-button style="float: right; padding: 3px 0" type="primary" @click="toggle">Toggle Follow</el-button>
                    
                </div>
                <p>{{"Name: " + user.name}} </p>
                <p>{{"Username: " + user.username}} </p>
                <p>{{"Followers: " + follower}} </p>
                <p>{{"Followers: " + following}} </p>
                <p>{{"Created At: " + convertTime(user.createdAt)}} </p>
                

            </el-card>
        
            <br>
            <br>
        
        </div>


        <h2>Tweets</h2>
        <br><br><br>

        <div v-for="f in feed" :key="f.id" class="text item">
            <el-card class="box-card">
                <div slot="header" class="clearfix">
                    <span>{{ f.username }}</span>
                    <span style="float: right; padding: 3px 0" type="text">{{convertTime(f.createdAt)}}</span>
                </div>
                <p>{{ f.body }} </p>
            </el-card>
        
            <br>
            <br>
        
        </div>

    </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import moment from "moment";

export default {
  data() {
    return {
      user: null,
      following: 0,
      follower: 0,
      feed: []
    };
  },
  created() {
    const username = this.$route.params.username;
    this.getProfile({
      username
    })
      .then(result => {
        this.user = result.data.data.user;
        this.following = result.data.data.following;
        this.follower = result.data.data.follower;
      })
      .catch(() => {
        this.$noty.error("Failed to fetch profile, please try again.");
      });

    this.fetchFeed({ type: "profile", username })
      .then(result => {
        this.feed = result.data.data;
      })
      .catch(() => {
        this.$noty.error("Failed to fetch feed, please try again.");
      });
  },
  methods: {
    convertTime(t) {
      const utc = moment.utc(t);
      return moment(utc)
        .local()
        .format();
    },

    toggle() {
      const data = {
        profile: this.user.username
      };

      this.toggleFollow(data)
        .then(result => {
          this.$noty.success("Toggle Successful");
        })
        .catch(() => {
          this.$noty.error("Failed to fetch feed, please try again.");
        });
    },

    ...mapActions({
      getProfile: "user/getProfile",
      toggleFollow: "user/toggleFollow",
      fetchFeed: "feed/list"
    })
  }
};
</script>

<style lang="scss" scoped>
</style>
