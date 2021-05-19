<template>
  <div class="container">
    <div class="container-box">
      <div class="logo">
        <img src="../img/logo.png" alt="">
      </div>
      <el-form class="login_form" :model="loginForm" ref="FormRef" :rules="rules">
        <el-form-item prop="name">
          <el-input prefix-icon="el-icon-user-solid" v-model="loginForm.name"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input prefix-icon="el-icon-warning" v-model="loginForm.password" type="password"></el-input>
        </el-form-item>
        <el-form-item class="btn">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button @click="reset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>


<script>
export default {
  name: "Login",
  data() {
    return {
      loginForm: {
        name: '',
        password: '',
      },
      rules: {
        name: [
          {required: true, message: '请输入用户名', trigger: 'blur'},
          {min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '请输入密码', trigger: 'change'},
          {min: 6, max: 10, message: '长度在 6 到 10 个字符', trigger: 'blur'}
        ]
      },
    }
  },
  methods: {
    reset() {
      this.$refs.FormRef.resetFields()
    },
    login() {


      this.$http.post("/login", this.loginForm,).then(r => {
        let res = r.data
        if (res.code !== 200) return this.$message.error(res.msg)
        this.$message.success(res.msg)
        this.$router.push("/home")
      }).catch(r => {
        console.log(r.data)
      })

    }
  }
}

</script>

<style lang="less" scoped>
.container {
  position: relative;
  height: 100%;
  background-color: #2b2b2b;
}

.container-box {
  width: 450px;
  height: 300px;
  position: absolute;
  top: 50%;
  left: 50%;
  background-color: #fff;
  transform: translate(-50%, -50%);

  .logo {
    width: 150px;
    height: 150px;
    border-radius: 50%;
    padding: 10px;
    border: 1px solid #2b2b2b;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: #fff;


  }

  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    border: 1px solid #2b2b2b;
  }

  .login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    box-sizing: border-box;
    padding: 0 20px;


  }

  .btn {
    display: flex;
    justify-content: flex-end;
  }
}
</style>