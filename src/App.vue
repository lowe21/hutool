<script setup lang="ts">
</script>

<template>
  <el-container>
    <el-main>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
    </el-main>
    <el-aside></el-aside>
  </el-container>
</template>

<style scoped>
.el-container {
  height: 100vh;
  line-height: 1.5;
}

.el-main {
  background-color: green;
}

.el-aside {
  background-color: purple;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
