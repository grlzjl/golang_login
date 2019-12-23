<template>
    <div class="back">
      <!-- <img src="https://pic4.zhimg.com/80/v2-a47051e92cf74930bedd7469978e6c08_hd.png" class="logo"/>           -->
        <card class="form-container" width="400px" height="320px">

            <div class="login-form">
                <div class="title">用户登录</div>
                <div>
                    <el-input class="login-input" placeholder="用户名" v-model="username"></el-input>
                    <el-input type="password" class="login-input" placeholder="密码" v-model="password"></el-input>
                    <el-button class="login-button" type="primary" @click="login">登录</el-button>

                </div>
            </div>

        </card>
    </div>
</template>

<script>
  //import Card from "@/components/Card";
  import axios from "axios";

  export default {
    name: "Login",
    //components: {Card},
    data() {
      return {
        username: "",
        password: "",
      }
    },
    methods: {
      login() {
        axios.post("http://localhost:8083/login", {
          username: this.username,
          password: this.password,
        }).then(res => {
          if (res.data.data == "err") {
            // this.$message.error("登录失败")
            this.$message({
              type:'success',
              message: '登录失败'
            })
          } else {
            // this.$message.success("登录成功")
            this.$message({
              type:'success',
              message: '登录成功'
            })
          }
        })

      },
    },
  }
</script>

<style scoped>
    .back {
        margin: 0;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        /* background-image: url(https://static.zhihu.com/heifetz/assets/sign_bg.db29b0fb.png); */
        background-repeat: no-repeat;
        background-color: #b8e5f8;
        background-size: cover;
        width: 100%;
        height: 100vh;
        overflow: auto;
    }

    .logo {
        width: 128px;
        height: 81px;
        margin-bottom: 24px;
        margin-top: 47px;
    }

    .title {
        color: #1a1a1a;
        font-weight: 600;
        position: relative;
        margin-left: 0;
    }

    .form-container {
        margin: 0 auto;
    }

    .title {
        display: flex;
        color: #1a1a1a;
        font-weight: 600;
        font-synthesis: style;
        height: 60px;
        line-height: 60px;
    }

    .login-form {
        margin-left: 20px;
        margin-right: 20px;
    }

    .login-input {
        margin-top: 20px;
    }

    .login-button {
        margin-top: 20px;
        width: 100%;
        font-size: 14px;
    }

    .login-link {
        float: right;
        margin-top: 15px;
        color: #8590a6;
        font-size: 14px;
    }
</style>