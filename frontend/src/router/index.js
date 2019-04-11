import Vue from 'vue'
import Router from 'vue-router'
import Experiment from "../components/Student/Experiment"
import ListExperiments from "../components/Student/ListExperiments"
import AddExperiments from "../components/Teacher/AddExperiments"
import ManageExperiments from "../components/Teacher/ManageExperiments"
import Header from "../components/Common/Header"
import SideBar from "../components/Common/SideBar"

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
    ]
})
