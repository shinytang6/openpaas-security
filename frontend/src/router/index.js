import Vue from 'vue'
import Router from 'vue-router'
import Experiment from "../components/Student/Experiment"
import ListExperiments from "../components/Student/ListExperiments"
import AddExperiments from "../components/Teacher/AddExperiments"
import UpdateExperiments from "../components/Teacher/UpdateExperiments"
import ManageExperiments from "../components/Teacher/ManageExperiments"
import Header from "../components/Common/Header"
import SideBar from "../components/Common/SideBar"
import DashBoard from "../components/SystemAdmin/DashBoard"
import CAdvisor from "../components/SystemAdmin/CAdvisor"
import Settings from "../components/Common/Settings"
import LoginPage from "../components/Common/LoginPage"

import ManageStudents from "../components/UserAdmin/ManageStudents"
import UpdateStudents from "../components/UserAdmin/UpdateStudents"
import AddStudents from "../components/UserAdmin/AddStudents"

import ManageTeachers from "../components/UserAdmin/ManageTeachers"
import UpdateTeachers from "../components/UserAdmin/UpdateTeachers"
import AddTeachers from "../components/UserAdmin/AddTeachers"

import ManageSysAdmins from "../components/UserAdmin/ManageSysAdmins"
import UpdateSysAdmins from "../components/UserAdmin/UpdateSysAdmins"
import AddSysAdmins from "../components/UserAdmin/AddSysAdmins"

Vue.use(Router)

export default new Router({
    // 设置链接激活时使用的 CSS 类名
    linkActiveClass:"active",
    routes: [
        {
            path: '/',
            name: 'ListExperiments',
            component: ListExperiments
        },
        {
            path: '/header',
            name: 'Header',
            component: Header
        },
        {
            path: '/sideBar',
            name: 'SideBar',
            component: SideBar
        },
        {
            path: '/experiment',
            name: 'Experiment',
            component: Experiment
        },
        {
            path: '/teacher/add',
            name: 'AddExperiments',
            component: AddExperiments
        },
        {
            path: '/teacher/manage',
            name: 'ManageExperiments',
            component: ManageExperiments
        },
        {
            path: '/teacher/update',
            name: 'UpdateExperiments',
            component: UpdateExperiments
        },
        {
            path: '/dashboard',
            name: 'DashBoard',
            component: DashBoard
        },
        {
            path: '/cadvisor',
            name: 'CAdvisor',
            component: CAdvisor
        },
        {
            path: '/settings',
            name: 'Settings',
            component: Settings
        },
        {
            path: '/login',
            name: 'LoginPage',
            component: LoginPage
        },
        {
            path: '/useradmin/manage/student',
            name: 'ManageStudents',
            component: ManageStudents
        },
        {
            path: '/useradmin/update/student',
            name: 'UpdateStudents',
            component: UpdateStudents
        },
        {
            path: '/useradmin/add/student',
            name: 'AddStudents',
            component: AddStudents
        },
        {
            path: '/useradmin/manage/teacher',
            name: 'ManageTeachers',
            component: ManageTeachers
        },
        {
            path: '/useradmin/update/teacher',
            name: 'UpdateTeachers',
            component: UpdateTeachers
        },
        {
            path: '/useradmin/add/teacher',
            name: 'AddTeachers',
            component: AddTeachers
        },
        {
            path: '/useradmin/manage/sysAdmin',
            name: 'ManageSysAdmins',
            component: ManageSysAdmins
        },
        {
            path: '/useradmin/update/sysAdmin',
            name: 'UpdateSysAdmins',
            component: UpdateSysAdmins
        },
        {
            path: '/useradmin/add/sysAdmin',
            name: 'AddSysAdmins',
            component: AddSysAdmins
        },
    ]
})
