<template>
    <div>
        <h1>Home Page</h1>
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
import { mapActions, mapState } from "vuex";
import moment from "moment";
export default {
  data() {
    return {
      feed: []
    };
  },
  created() {
    this.fetchFeed({ type: "home" })
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
    ...mapActions({
      fetchFeed: "feed/list"
    })
  }
};
</script>

<style lang="scss" scoped>
</style>