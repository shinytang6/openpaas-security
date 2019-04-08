/* eslint-disable */
<template>
<div class="experiments">

     <div class="experiments-wrapper">
        <!--eslint-disable-next-line-->
        <div class="experiment-item" v-for="(experiment,index) in experimentArr" :key="index">
            <!--<div class="poster" @click='showDetail(experiment.name)' :style="{ backgroundImage: 'url(' + experiment.poster+ ')' }"></div>-->
            <div class="content">
                <div class="title" @click="showDetail(experiment)" target="_blank">{{experiment.name}}</div>
                <div class="time">时间：{{experiment.name}}</div>
                <div class="place">城市：{{experiment.place}}</div>
            </div>
        </div>
     </div>
</div>
</template>

<script>
export default {
  name: 'listExperiments',
  data(){
      return {
          experimentArr: [],
      }
  },
  mounted: function(){
      var that = this
      this.$axios.get('/api/experiment/getall')
          .then(function (response) {
              if(response.status == 200) {
                  that.experimentArr = response.data.data
                  console.log(that.experimentArr)
              }
          })
          .catch(function (error) {
              console.log(error);
          });
  },
  methods: {
    showDetail: function(experiment){
        var that = this
        this.$axios.get('/api/experiment/get?id='+experiment.id+'&experimentId='+experiment.experimentId+'&name='+experiment.name)
            .then(function (response) {
                if(response.status == 200) {
                    that.$router.push({
                        name: "Experiment",
                        // params: {
                        //     id: activity_id,
                        //     data: response.data.data
                        // }
                    });
                }
            })
            .catch(function (error) {
                console.log(error);
            });
        // this.$route.push({
        //     name: "DetailActivity",
        //     params:{id: activity_id}
        // });
        // this.$router.push({path:'/activity/'+activity_id});
    }
  }
}
</script>

<style>
.experiments {
    position: fixed;
    left: 220px;
    height: 500px;
    width: 100%;
    height: 100%;

}

.experiments-wrapper {
    display: flex;
    flex-wrap: wrap;
    margin-top: 136px;
}

.experiments-wrapper .experiment-item {
    flex: 0 0 25%;
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
