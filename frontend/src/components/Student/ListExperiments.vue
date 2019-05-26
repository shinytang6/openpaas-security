/* eslint-disable */
<template>
<div class="experiments">

     <div class="experiments-wrapper">
        <!--eslint-disable-next-line-->
        <div class="experiment-item" v-for="(experiment,index) in experimentArr" :key="index">
            <div class="content" @click="showDetail(experiment)" target="_blank">
                <div class="title">实验：{{experiment.name}}</div>
                <div class="time">时间：{{experiment.date}}</div>
                <div class="time">点击进行实验</div>
            </div>
        </div>
     </div>
</div>
</template>

<script>
import { setCookie, getCookie } from '../../js/cookieUtil'
export default {
  name: 'listExperiments',
  data(){
      return {
          experimentArr: [],
          identity: "",
      }
  },
  mounted: function(){
      var that = this
      var identity = this.isLogin
      var id = this.id
      if (identity == "学生") {
          this.$axios.get('/api/experiment/get?id='+id)
              .then(function (response) {
                  if (response.status == 200) {
                      that.experimentArr = response.data.data
                      console.log(that.experimentArr)
                  }
              })
              .catch(function (error) {
                  console.log(error);
              });
      }
  },
  computed: {
      isLogin () {
          this.identity = getCookie("identity");
          this.id = getCookie("id");
          return this.identity;
      }
  },
  methods: {
    showDetail: function(experiment){
        console.log(experiment)
        this.$router.push({
            name: "Experiment",
            params: {
                id: experiment.id,
                name: experiment.name,
                date: experiment.date,
                fileHash: experiment.fileHash,
                fileName: experiment.fileName,
                desc: experiment.description,
                address: experiment.Address
            }
        });
    }
  }
}
</script>

<style>
.experiments {
    position: fixed;
    left: 220px;
    height: 100%;
    width: 80%;
}

.experiments-wrapper {
    display: flex;
    flex-wrap: wrap;
    margin-top: 136px;
}

.experiments-wrapper .experiment-item {
    flex: 0 0 33%;
    margin-bottom: 30px;

}

.experiments-wrapper .experiment-item .poster {
    background-size: 150px 150px;
    width: 150px;
    height: 150px;
    display: inline-block;
    cursor: pointer;
}

.experiments-wrapper .experiment-item .content {
    display: inline-block;
    vertical-align: top;
    margin-left: 10px;
    font-size: 14px;
    color: #999;
    cursor: pointer;
}

.experiments-wrapper .experiment-item .content .title {
    font-size: 20px;
    width: 200px;
    color: #333;
    margin-bottom: 15px;
    cursor: pointer;
}
.experiments-wrapper .experiment-item .content .time {
   margin-bottom: 5px;
}
</style>
