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
      <el-menu-item index="1" class="logo"><img alt="Vue logo" src="../assets/logo.jpg" ></el-menu-item>
      <el-menu-item index="2" class="title">网络安全实验平台</el-menu-item>
      <el-submenu index="3" class="identity">
        <template slot="title">管理员</template>
        <el-menu-item index="3-1">选项1</el-menu-item>
        <el-menu-item index="3-2">选项2</el-menu-item>
        <el-menu-item index="3-3">选项3</el-menu-item>
      </el-submenu>
    </el-menu>
</div>
</template>

<script>
export default {
  name: 'header',

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
