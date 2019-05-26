<template>
<div class="experiment-vnc">
     <el-button plain @click="open" class="explanation">
          实验说明
     </el-button>
     <!--<iframe class="content" src="http://192.168.99.100:32333/"></iframe>-->
     <iframe class="content" :src="address"></iframe>
</div>
</template>

<script>
export default {
  name: 'experiment',
  data() {
     return {
          address: "",
          message: "",
     };
  },
  beforeMount: function() {
     var id = this.$route.params.id;
     var name = this.$route.params.name;
     var date = this.$route.params.date;
     var fileHash = this.$route.params.fileHash;
     var fileName = this.$route.params.fileName;
     var desc = this.$route.params.desc;
     var address = this.$route.params.address;

     this.address =  "http://" + address
       this.message = '<div>实验Id：' + id + '</div></br>' + '<div>实验名：' + name + '</div></br>' + '<div>实验日期：' + date + '</div></br>'
               + '<div>实验地址：' + this.address + '</div></br>' + '<div>实验简介：' + desc + '</div></br>' +
               '<div>实验指导书：<a href="' + 'http://localhost:8080/api/file/download?filehash=' + fileHash + '">' + fileName + '</a></div>'
  },
  methods: {
       open() {
            this.$notify({
                 title: '实验基本信息',
                 dangerouslyUseHTMLString: true,
                 message: this.message,
                 position: 'top-left',
                 duration: 0,
                 offset: 100,
            });
       }
  }
}
</script>

<style>
.experiment-vnc {
     position: fixed;
     left: 220px;
     height: 500px;
     width: 100%;
     height: 100%;
}
.experiment-vnc .content {
     position: fixed;
     left: 220px;
     width: 100%;
     height: 100%;
     top: 0;
}

.experiment-vnc .explanation{
     position: absolute;
     top: 130px;
     left: -154px;
     background-color: #545c64;
     color: #fff;
}

.el-notification {
     /*width: 220px;*/
     width: 300px;
}
.el-notification.left {
     left: 0px;
}
</style>
