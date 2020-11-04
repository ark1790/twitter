<template>
    <el-container>
        <el-form ref="form" :model="form" label-width="120px">
        <el-form-item label="Username">
            <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="Password">
            <el-input v-model="form.password" show-password></el-input>
        </el-form-item>
        
        
        <el-form-item>
            <el-button type="primary" @click="onSubmit" round>Login</el-button>
            <el-button type="primary" @click="$router.push({name:'register'})" round>Register</el-button> 
            <el-button native-type="reset" round>Cancel</el-button>
        </el-form-item>
        </el-form>
    </el-container>
</template>
<script>
import { mapActions, mapGetters } from "vuex";

export default {
  data() {
    return {
      form: {
        username: "ark",
        password: "ark"
      }
    };
  },
  methods: {
    onSubmit() {
      this.logIn({
        username: this.form.username,
        password: this.form.password
      })
        .then(result => {
          this.$router.replace({ name: "home" });
        })
        .catch(() => {
          this.$noty.error("Failed to log in, please try again.");
        });
    },

    ...mapActions({
      logIn: "auth/logIn"
    })
  }
};
</script>

<style lang="scss" scoped>
</style>