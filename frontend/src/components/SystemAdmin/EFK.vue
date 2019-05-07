<template>
    <div class="efk">
        <iframe class="content" :src="addr"></iframe>
    </div>
</template>

<script>
    export default {
        name: 'efk',
        data() {
            return {
                addr: "http://" + process.env.VUE_APP_CLUSTER_IP + ":" + process.env.VUE_APP_CLUSTER_EFK_PORT
            }
        },
    }
</script>

<style>
    .efk {
        position: fixed;
        left: 220px;
        height: 500px;
        width: 90%;
        height: 100%;
        /*background-color: red;*/
    }
    .efk .content {
        position: fixed;
        left: 220px;
        width: 1600px;
        height: 100%;
        top: 60px;
    }
</style>
