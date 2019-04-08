import Vue from 'vue'
import Router from 'vue-router'
import Experiment from "../components/Experiment"
import ListExperiments from "../components/ListExperiments"
import Header from "../components/Header"
import SideBar from "../components/SideBar"

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
    ]
})
