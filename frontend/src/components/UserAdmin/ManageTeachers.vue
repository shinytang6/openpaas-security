/* eslint-disable */
<template>
    <div class="teachers-management">
        <div class="title">教师管理</div>
        <el-button type="primary" class="addBtn" @click="addTeacher">添加</el-button>
        <el-table
                :data="tableData.filter(data => !search || data.name.toLowerCase().includes(search.toLowerCase()))"
                style="width: 100%">
            <el-table-column
                    type="index"
                    label="用户id"
                    width="100">
            </el-table-column>
            <el-table-column
                    label="姓名"
                    prop="name">
            </el-table-column>
            <el-table-column
                    label="密码"
                    prop="password">
            </el-table-column>
            <el-table-column
                    label="工号"
                    prop="teacherId">
            </el-table-column>
            <el-table-column
                    label="邮箱"
                    prop="email"
                    width="200">
            </el-table-column>
            <el-table-column
                    label="联系电话"
                    prop="phone">
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
        name: 'manageTeachers',
        data() {
            return {
                tableData: [],
                search: ''
            }
        },
        inject: ['reload'],
        mounted: function(){
            var that = this
            this.$axios.get('/api/teacher/getall')
                .then(function (response) {
                    if(response.status == 200) {
                        // that.experimentArr = response.data.data

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
                    name: "UpdateTeachers",
                    params: {
                        index: index,
                        data: row
                    }
                });
            },
            handleDelete(index, row) {
                console.log(index, row);
                var that = this
                this.$axios.get('/api/teacher/delete?name='+row.name)
                    .then(function (response) {
                        if(response.status == 200) {
                            that.tableData = response.data.data
                            that.reload()
                        }
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            addTeacher() {
                this.$router.push({
                    name: "AddTeachers",
                });
            }
        },
    }
</script>

<style>
    .teachers-management {
        position: fixed;
        left: 220px;
        height: 100%;
        width: 1200px;
        top: 200px;
        margin-left: 200px;
    }

    .teachers-management .title {
        text-align: center;
        margin-bottom: 50px;
        margin-left: -80px;
        font-size: 20px;
    }

    .teachers-management .addBtn {
        position: absolute;
        top: 0px;
        right: 20px;
    }
</style>
