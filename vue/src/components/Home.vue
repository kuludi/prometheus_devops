<template>
  <el-container>
    <el-header>
      <div>
        <span>管理系统</span>
      </div>
      <el-button type="info" @click="logOut">退出</el-button>
    </el-header>
    <el-container>
      <el-aside width="200px">Aside</el-aside>
      <el-main>Main</el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: "Home"
}
</script>

<style lang="less" scoped>
.el-container {
  height: 100%;
}
.el-header {
  display: flex;
  justify-content: space-between;
  background-color: #383f4c;
  align-items: center;
  font-size: 14px;

  > div {
    color: #ffffff;
  }
}
.el-aside {
  height: 100%;
  background-color: red;
}
</style>