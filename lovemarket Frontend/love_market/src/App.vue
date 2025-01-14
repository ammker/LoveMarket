<template>
  <div id="app">
    <router-view :key="$route.fullPath" />
  </div>
</template>

<script>

  export default {
    components: {
    }
  }
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
