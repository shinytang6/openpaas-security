<template>
<div class="sidebar">
    <el-row class="tac" v-show="identity!='未登录'">
      <el-col :span="12">
        <h5>自定义颜色</h5>
        <el-menu
          default-active="1"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b">

          <el-menu-item index="1" @click="goToExperiment" v-show="identity=='学生'">
            <i class="el-icon-menu"></i>
            <span slot="title">课程实验</span>
          </el-menu-item>

          <el-menu-item index="1" @click="addExperiment" v-show="identity=='教师'">
            <i class="el-icon-document"></i>
            <span slot="title">添加实验</span>
          </el-menu-item>
          <el-menu-item index="2" @click="manageExperiment" v-show="identity=='教师'">
            <i class="el-icon-setting"></i>
            <span slot="title">管理实验</span>
          </el-menu-item>

          <el-menu-item index="1" @click="goToDashboard" v-show="identity=='系统管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">控制面板</span>
          </el-menu-item>
          <el-menu-item index="2" @click="goToCadvisor" v-show="identity=='系统管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">流量监控</span>
          </el-menu-item>
          <el-menu-item index="3" @click="goToEFK" v-show="identity=='系统管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">日志监控</span>
          </el-menu-item>

          <el-menu-item index="1" @click="manageStudent" v-show="identity=='用户管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">学生管理</span>
          </el-menu-item>
          <el-menu-item index="3" @click="manageTeacher" v-show="identity=='用户管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">教师管理</span>
          </el-menu-item>
          <el-menu-item index="5" @click="manageSysAdmin" v-show="identity=='用户管理员'">
            <i class="el-icon-setting"></i>
            <span slot="title">系统管理员管理</span>
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
    var identity = this.isLogin
    if (identity == undefined || identity == "") {
      // this.$router.replace('/login');
      this.identity = "未登录"
    } else if (identity == "学生") {
      this.identity = "学生"
    } else if (identity == "教师") {
      this.identity = "教师"
    } else if (identity == "系统管理员") {
      this.identity = "系统管理员"
    } else if (identity == "用户管理员") {
      this.identity = "用户管理员"
    }
  },
  computed: {
    isLogin () {
      this.identity = getCookie("identity");
      return this.identity;
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
    goToEFK() {
      this.$router.push({
        name: "EFK",
      });
    },
    manageStudent() {
      this.$router.push({
        name: "ManageStudents",
      });
    },
    manageTeacher() {
      this.$router.push({
        name: "ManageTeachers",
      });
    },
    manageSysAdmin() {
      this.$router.push({
        name: "ManageSysAdmins",
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
