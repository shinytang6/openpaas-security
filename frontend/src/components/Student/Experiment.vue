<template>
<div class="experiment-vnc">
     <!--<iframe class="content" src="http://192.168.99.100:32333/"></iframe>-->
     <iframe class="content" :src="address"></iframe>
</div>
</template>

<script>
export default {
  name: 'experiment',
  data() {
     return {
          address: ""
     };
  },
  beforeMount: function() {
     var id = this.$route.params.id;
     var addr = this.$route.params.address;

     this.address =  "http://" + addr
  },
}
</script>

<style>
.experiment-vnc {
     position: fixed;
     left: 220px;
     height: 500px;
     width: 100%;
     height: 100%;
     /*background-color: red;*/
}
.experiment-vnc .content {
     position: fixed;
     left: 220px;
     width: 100%;
     height: 100%;
     top: 0;
}
</style>
