<template>
  <el-container>

    <el-header>
      <div>
        <span>管理系统</span>
      </div>
      <el-button type="info" @click="logOut">退出</el-button>
    </el-header>

    <el-container>

      <el-aside width="200px">

        <div class="zhetie"><span>|||</span></div>

        <el-menu
            class="el-menu-vertical-demo"
            background-color="#393f4b"
            text-color="#fff" unique-opened
            active-text-color="#579ff8" router>
          <el-submenu>
            <template slot="title">
              <i class="el-icon-user-solid"></i>
              <span>用户管理</span>
            </template>
            <el-menu-item index="users">
              <template slot="default">
                <i class="el-icon-menu"></i>
                <span>用户列表</span>
              </template>
            </el-menu-item>
          </el-submenu>

          <el-submenu index="2">
            <template slot="title">
              <i class="el-icon-location"></i>
              <span>监控管理</span>
            </template>
          </el-submenu>

        </el-menu>
      </el-aside>

      <el-main>
        <router-view></router-view>
      </el-main>


    </el-container>

  </el-container>
</template>

<script>
export default {
  name: "Home",
  methods: {
    logOut() {
      this.$router.push("/")
    },
  },
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
  background-color: #393f4b;
}

.zhetie {
  text-align: center;
  letter-spacing: 0.4ex;
  cursor: pointer;
  color: #ffffff;
}

.el-menu {
  border-right: 0;
}
</style>