<template>
    <div>
        <h1>Post Tweet</h1>
        <br><br><br>
        <el-container width="200px">
            <el-form ref="form" :model="form" label-width="120px">
            
            <el-form-item label="TweetBody" prop="body">
                <el-input type="textarea" v-model="form.body"></el-input>
            </el-form-item>
            
            
            <el-form-item>
                <el-button type="primary" @click="onSubmit" round>Submit</el-button>
                <el-button native-type="reset" round>Cancel</el-button>
            </el-form-item>
            </el-form>
        </el-container>
    </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  data() {
    return {
      form: {
        body: ""
      }
    };
  },
  methods: {
    onSubmit() {
      const data = {
        body: this.form.body
      };
      this.post(data)
        .then(result => {
          this.$noty.success("Successfully posted tweet");
        })
        .catch(() => {
          this.$noty.error("Failed to log in, please try again.");
        });
    },
    ...mapActions({
      post: "tweet/postTweet"
    })
  }
};
</script>

<style lang="scss" scoped>
</style>