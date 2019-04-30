<template>
    <div class="efk">
        <iframe class="content" src="http://192.168.99.100:30003/"></iframe>
    </div>
</template>

<script>
    export default {
        name: 'efk',
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
