<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>用户管理</el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>

      <el-row :gutter="20" class="search">

        <!--        查找用户-->
        <el-col :span="10">
          <el-input placeholder="请输入内容">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 添加用户-->
        <el-col :span="6">
          <el-button type="primary">添加用户</el-button>
        </el-col>

      </el-row>

      <!--用户展示列表-->
      <el-table :data="userList" style="width: 100%" stripe border>
        <el-table-column type="index" width="50" label="序号"></el-table-column>
        <el-table-column prop="username" label="姓名"></el-table-column>
        <el-table-column prop="email" label="邮箱" width="200px"></el-table-column>
        <el-table-column prop="mobile" label="电话"></el-table-column>
        <el-table-column prop="role_name" label="角色"></el-table-column>

        <el-table-column prop="" label="操作">
          <template v-slot="{row}">
            <el-row>
              <el-button size="mini" type="primary" icon="el-icon-edit"></el-button>
              <el-button size="mini" type="danger" icon="el-icon-delete"></el-button>
              <el-tooltip class="item" effect="dark" content="分配角色" placement="top" :enterable="false">
                <el-button size="mini" type="warning" icon="el-icon-setting" @click=""></el-button>
              </el-tooltip>
            </el-row>
          </template>
        </el-table-column>
      </el-table>

    </el-card>


  </div>
</template>

<script>
export default {
  name: "User",
  data() {
    return {

      userList: [],
    }
  },
  methods: {
    getUserList() {
      this.$http.get("getuser").then(r => {
        let res = r.data
        if (res.code !== 200) return this.$message.error(res.mes)
        this.userList = res.data
      })
    },
  },
}
</script>

<style lang="less" scoped>


</style>