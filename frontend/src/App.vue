<template>
  <div id="app">
    <vHeader></vHeader>
    <SideBar></SideBar>
    <router-view/>
  </div>
</template>

<script>
import vHeader from './components/Common/Header.vue'
import SideBar from './components/Common/SideBar.vue'

export default {
  name: 'app',
  components: {

    vHeader,
    SideBar,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;

}
body {
 background-color: #E8E5E5;

}
</style>
