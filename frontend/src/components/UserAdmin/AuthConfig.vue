/* eslint-disable */
<template>
    <div class="auth-config">
        <div class="title">授权配置</div>
        <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
            <el-form-item label="角色">
                <el-select v-model="ruleForm.user" placeholder="请选择授权用户">
                    <el-option label="teacher1" value="teacher1"></el-option>
                    <el-option label="teacher2" value="teacher2"></el-option>
                    <el-option label="teacher3" value="teacher3"></el-option>
                    <el-option label="sysAdmin" value="sysAdmin"></el-option>
                    <el-option label="other" value="other"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="命名空间">
                <el-select v-model="ruleForm.namespace" placeholder="请选择授权命名空间">
                    <el-option label="teacher1" value="teacher1"></el-option>
                    <el-option label="teacher2" value="teacher2"></el-option>
                    <el-option label="teacher3" value="teacher3"></el-option>
                    <el-option label="default" value="default"></el-option>
                    <el-option label="*" value="*"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="资源">
                <el-select v-model="ruleForm.resource" placeholder="请选择授权资源">
                    <el-option label="pod" value="pod"></el-option>
                    <el-option label="serverAcount" value="serverAcount"></el-option>
                    <el-option label="*" value="*"></el-option>
                </el-select>
            </el-form-item>

            <el-form-item label="API">
                <el-select v-model="ruleForm.api" placeholder="请选择授权API">
                    <el-option label="*" value="*"></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="是否只读">
                <el-switch v-model="ruleForm.readOnly"></el-switch>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">添加</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'authConfig',
        data() {
            return {
                ruleForm: {
                    name: '',
                    teacherId: '',
                    email: '',
                    phone: '',
                },
            }
        },
        inject: ['reload'],
        mounted: function(){
            var that = this
            that.tableData = [{user: '/', group: 'system:authenticated', namespace: '/', resource: '/', apiGroup: '*', readOnly: 'true'},
                {user: 'teacher1', group: '/', namespace: 'teacher1', resource: '*', apiGroup: '*', readOnly: 'false'},
                {user: 'teacher2', group: '/', namespace: 'teacher2', resource: '*', apiGroup: '*', readOnly: 'false'},
                {user: 'teacher3', group: '/', namespace: 'teacher3', resource: '*', apiGroup: '*', readOnly: 'false'},
                {user: 'sysAdmin', group: '/', namespace: '*', resource: '*', apiGroup: '*', readOnly: 'false'},
            ]
            // this.$axios.get('/api/student/getall')
            //     .then(function (response) {
            //         if(response.status == 200) {
            //             // that.experimentArr = response.data.data
            //
            //             that.tableData = response.data.data
            //         }
            //     })
            //     .catch(function (error) {
            //         console.log(error);
            //     });
        },
        methods: {
            handleEdit(index, row) {
                console.log(index, row);
                this.$router.push({
                    name: "UpdateStudents",
                    params: {
                        index: index,
                        data: row
                    }
                });
            },
            handleDelete(index, row) {
                console.log(index, row);
                var that = this
                this.$axios.get('/api/student/delete?name='+row.name)
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
            },
            addStudent() {
                this.$router.push({
                    name: "AddStudents",
                });
            }
        },
    }
</script>

<style>

    .auth-config {
        position: fixed;
        height: 500px;
        width: 500px;
        top: 140px;
        margin-left: 700px;
    }
    .auth-config .title {
        text-align: center;
        margin-bottom: 30px;
        margin-left: 30px;
        font-size: 20px;
    }
    .el-input--suffix .el-input__inner {
        width: 400px;
    }
</style>
