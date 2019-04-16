<template>
    <div class="cadvisor">
        <iframe class="content" src="http://192.168.99.100:4194/"></iframe>
    </div>
</template>

<script>
    export default {
        name: 'cadvisor',
    }
</script>

<style>
    .cadvisor {
        position: fixed;
        left: 220px;
        height: 500px;
        width: 100%;
        height: 100%;
        /*background-color: red;*/
    }
    .cadvisor .content {
        position: fixed;
        left: 220px;
        width: 100%;
        height: 100%;
        top: 0;
    }
</style>
