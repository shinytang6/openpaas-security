<template>
    <div class="dashboard">
        <iframe class="content" src="http://192.168.99.100:30000/"></iframe>
    </div>
</template>

<script>
    export default {
        name: 'dashboard',
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
