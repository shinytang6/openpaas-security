<template>
    <div class="dashboard">
        <iframe class="content" :src="addr"></iframe>
    </div>
</template>

<script>
    export default {
        name: 'dashboard',
        data() {
            return {
                addr: "http://" + process.env.VUE_APP_CLUSTER_IP + ":" + process.env.VUE_APP_CLUSTER_DASHBOARD_PORT
            }
        },
    }
</script>

<style>
    .dashboard {
        position: fixed;
        left: 220px;
        height: 500px;
        width: 90%;
        height: 100%;
        /*background-color: red;*/
    }
    .dashboard .content {
        position: fixed;
        left: 220px;
        width: 1600px;
        height: 100%;
        top: 0;
    }
</style>
