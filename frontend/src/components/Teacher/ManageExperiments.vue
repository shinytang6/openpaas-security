/* eslint-disable */
<template>
    <div class="experiments-management">
        <div class="title">管理实验</div>
        <el-table
                :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
                style="width: 100%">
            <el-table-column
                    type="index"
                    label="序号"
                    width="50">
            </el-table-column>
            <el-table-column
                    label="日期"
                    prop="date">
            </el-table-column>
            <el-table-column
                    label="实验名称"
                    prop="name">
            </el-table-column>
            <el-table-column
                    label="实验总数"
                    prop="name">
            </el-table-column>
            <el-table-column
                    align="right">
                <template slot="header" slot-scope="scope">
                    <el-input
                            v-model="search"
                            size="mini"
                            placeholder="输入关键字搜索"/>
                </template>
                <template slot-scope="scope">
                    <el-button
                            size="mini"
                            @click="handleEdit(scope.$index, scope.row)">Edit</el-button>
                    <el-button
                            size="mini"
                            type="danger"
                            @click="handleDelete(scope.$index, scope.row)">Delete</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
    export default {
        name: 'manageExperiments',
        data() {
            return {
                tableData: [{
                    date: '2016-05-02',
                    name: '王小虎1',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-04',
                    name: '王小虎2',
                    address: '上海市普陀区金沙江路 1517 弄'
                }, {
                    date: '2016-05-01',
                    name: '王小虎3',
                    address: '上海市普陀区金沙江路 1519 弄'
                }, {
                    date: '2016-05-03',
                    name: '王小虎4',
                    address: '上海市普陀区金沙江路 1516 弄'
                }],
                search: ''
            }
        },
        mounted: function(){
            var that = this
            this.$axios.get('/api/experiment/getall')
                .then(function (response) {
                    if(response.status == 200) {
                        // that.experimentArr = response.data.data
                        console.log("tes!!")
                        console.log(response.data.data)
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        methods: {
            handleEdit(index, row) {
                console.log(index, row);
            },
            handleDelete(index, row) {
                console.log(index, row);
            }
        },
    }
</script>

<style>
.experiments-management {
    position: fixed;
    left: 220px;
    height: 100%;
    width: 1200px;
    top: 200px;
    margin-left: 200px;
}

.experiments-management .title {
    text-align: center;
    margin-bottom: 50px;
    margin-left: -80px;
    font-size: 20px;
}
</style>
