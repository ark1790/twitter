<template>
    <div class="error" v-if="value != null">
        <p>{{value}}</p>
    </div>
</template>

<script>
    export default {
        name: 'Validation',
        props: ['value']
    }
</script>

<style scoped>
    .error {
        color: #F00;
    }
</style>
