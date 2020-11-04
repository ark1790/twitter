<script>
import appConfig from "@src/app.config";
import "vuejs-noty/dist/vuejs-noty.css";

export default {
  page: {
    // All subcomponent titles will be injected into this template.
    titleTemplate(title) {
      title = typeof title === "function" ? title(this.$store) : title;
      return title ? `${title} | ${appConfig.title}` : appConfig.title;
    }
  }
};
</script>

<template>
    <div>
        <router-view></router-view>
    </div>
</template>


<style lang="scss">
</style>
