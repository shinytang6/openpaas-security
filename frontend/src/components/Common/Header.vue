/* eslint-disable */
<template>
<div class="header">
    <el-menu
      :default-active="activeIndex2"
      class="menu"
      mode="horizontal"
      @select="handleSelect"
      background-color="#2F4056"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item index="1" class="logo"><img alt="Vue logo" src="../../assets/logo.jpg" ></el-menu-item>
      <el-menu-item index="2" class="title">网络安全实验平台</el-menu-item>
      <el-submenu index="3" class="identity" v-show="identity!='未登录'">
        <template slot="title">{{identity}}</template>
        <el-menu-item index="3-1" @click="goToSettings">个人设置</el-menu-item>
        <el-menu-item index="3-2" @click="logOut">退出</el-menu-item>
        <!--<el-menu-item index="3-3">选项3</el-menu-item>-->
      </el-submenu>
      <el-menu-item index="4" class="identity" v-show="identity=='未登录'" @click="login">未登录</el-menu-item>
    </el-menu>
</div>
</template>

<script>
import { setCookie, getCookie } from '../../js/cookieUtil'
export default {
  name: 'header',
  data() {
      return {
          identity: "",
      }
  },
  mounted () {
      if (this.isLogin==undefined || this.isLogin=="") {
          // this.$router.replace('/login');
          this.identity = "未登录"
      } else {
          // 执行后续操作
          this.identity = "管理员"
          // this.phone = getCookie("phone");
      }
  },
  computed: {
      isLogin () {
          this.userId = getCookie("userId");
          console.log("fuck")
          console.log(this.userId)
          return this.userId;
      }
  },
  methods: {
      login() {
          this.$router.push({
              name: "LoginPage",
          });
      },
      logOut() {
          setCookie('name', "");
          setCookie('email', "");
          setCookie('phone', "");
          setCookie('class', "");
          setCookie('studentId', "");
          setCookie('userId', "")
          window.location.reload()
          this.$router.push({
              name: "/",
          });
      },
      goToSettings() {
          this.$router.push({
              name: "Settings",
          });
      }
  }
}
</script>

<style>
.header {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 1;
}
.menu .logo {

}

.menu .logo img {
    width: 60px;
    heitght: 60px;
    background-color: red !important;
}
.menu .title {
    margin-left: 20px;
}

.menu .identity {
    position: absolute;
    right: 10px;
}
</style>
