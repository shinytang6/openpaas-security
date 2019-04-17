<template>
<div class="sidebar">
    <el-row class="tac" v-show="identity!='未登录'">
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
          <el-menu-item index="1" @click="goToExperiment">
            <i class="el-icon-menu"></i>
            <span slot="title">课程实验</span>
          </el-menu-item>
          <el-menu-item index="2" @click="addExperiment">
            <i class="el-icon-document"></i>
            <span slot="title">添加实验</span>
          </el-menu-item>
          <el-menu-item index="3" @click="manageExperiment">
            <i class="el-icon-setting"></i>
            <span slot="title">管理实验</span>
          </el-menu-item>
          <el-menu-item index="4" @click="goToDashboard">
            <i class="el-icon-setting"></i>
            <span slot="title">Dashboard</span>
          </el-menu-item>
          <el-menu-item index="5" @click="goToCadvisor">
            <i class="el-icon-setting"></i>
            <span slot="title">Cadvisor</span>
          </el-menu-item>

          <el-menu-item index="5" @click="manageStudent">
            <i class="el-icon-setting"></i>
            <span slot="title">学生管理</span>
          </el-menu-item>
          <el-menu-item index="5" @click="addStudent">
            <i class="el-icon-setting"></i>
            <span slot="title">添加学生</span>
          </el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
</div>
</template>

<script>
import { setCookie, getCookie } from '../../js/cookieUtil'
export default {
  name: 'sidebar',
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
      return this.userId;
    }
  },
  methods: {
    goToExperiment() {
      this.$router.push({
        name: "ListExperiments",
      });
    },
    addExperiment() {
      this.$router.push({
        name: "AddExperiments",
      });
    },
    manageExperiment() {
      this.$router.push({
        name: "ManageExperiments",
      });
    },
    goToDashboard() {
      this.$router.push({
        name: "DashBoard",
      });
    },
    goToCadvisor() {
      this.$router.push({
        name: "CAdvisor",
      });
    },
    manageStudent() {
      this.$router.push({
        name: "ManageStudents",
      });
    },
    addStudent() {
      this.$router.push({
        name: "AddStudents",
      });
    }
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
