<template>
    <el-container>
        <el-form ref="form" :model="form" label-width="120px">
        <el-form-item label="Name">
            <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="Username">
            <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="Password">
            <el-input v-model="form.password" show-password></el-input>
        </el-form-item>
        <el-form-item label="Private">
          <el-switch
            v-model="form.private"
            active-color="#13ce66"
            inactive-color="#ff4949">
          </el-switch>
        </el-form-item>
        
        
        
        <el-form-item>
            <el-button type="primary" @click="onSubmit" round>Register</el-button>
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
        name: "Arman Kamal",
        username: "ark",
        password: "ark",
        private: false
      }
    };
  },
  methods: {
    onSubmit() {
      const data = {
        name: this.form.name,
        username: this.form.username,
        password: this.form.password,
        private: this.form.private
      };

      this.signUp(data)
        .then(result => {
          this.$router.replace({ name: "auth" });
        })
        .catch(() => {
          this.$noty.error("Failed to sign up, please try again.");
        });
    },
    ...mapActions({
      signUp: "user/signUp"
    })
  }
};
</script>

<style lang="scss" scoped>
</style>