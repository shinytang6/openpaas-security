<template>
<div class="sidebar">
    <el-row class="tac">
      <el-col :span="12">
        <h5>自定义颜色</h5>
        <el-menu
          default-active="2"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b">
          <el-menu-item index="1">
            <i class="el-icon-menu"></i>
            <span slot="title" @click="goToExperiment">课程实验</span>
          </el-menu-item>
          <!--<el-menu-item index="2">-->
            <!--<i class="el-icon-document"></i>-->
            <!--<span slot="title">导航三</span>-->
          <!--</el-menu-item>-->
          <!--<el-menu-item index="3">-->
            <!--<i class="el-icon-setting"></i>-->
            <!--<span slot="title">导航四</span>-->
          <!--</el-menu-item>-->
        </el-menu>
      </el-col>
    </el-row>
</div>
</template>

<script>
export default {
  name: 'sidebar',
  methods: {
    goToExperiment() {
      this.$router.push({
        name: "ListExperiments",
      });
    },
  }
}
</script>

<style>
.sidebar {
    width: 220px;
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    height: 100%;
    background-color: #545c64;
}
.sidebar .tac .el-col{
    width: 220px;

}

</style>
