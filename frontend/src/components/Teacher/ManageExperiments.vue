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
                    type="experimentId"
                    label="实验序号"
                    prop="experimentId"
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
                    label="实验状态"
                    prop="status">
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
                tableData: [],
                search: ''
            }
        },
        inject: ['reload'],
        mounted: function(){
            var that = this
            this.$axios.get('/api/experiment/getall')
                .then(function (response) {
                    if(response.status == 200) {
                        // that.experimentArr = response.data.data
                        console.log("tes!!")
                        console.log(response.data.data)
                        that.tableData = response.data.data
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        methods: {
            handleEdit(index, row) {
                console.log(index, row);
                this.$router.push({
                    name: "UpdateExperiments",
                    params: {
                        index: index,
                        data: row
                    }
                });
            },
            handleDelete(index, row) {
                console.log(index, row);
                var that = this
                this.$axios.get('/api/experiment/delete?name='+row.name)
                    .then(function (response) {
                        if(response.status == 200) {
                            // that.experimentArr = response.data.data
                            console.log("tes!!")
                            console.log(response.data.data)
                            that.tableData = response.data.data
                            that.reload()
                        }
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
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
