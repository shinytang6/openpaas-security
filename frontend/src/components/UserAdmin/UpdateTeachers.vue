/* eslint-disable */
<template>
    <div class="teacher-update">
        <div class="title">教师信息更新</div>
        <el-form ref="ruleForm" :model="ruleForm" :rules="rules" label-width="80px">
            <el-form-item label="编号" disabled="true">
                <el-input v-model="ruleForm.id" :disabled="true"></el-input>
            </el-form-item>
            <el-form-item label="姓名">
                <el-input v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="ruleForm.password"></el-input>
            </el-form-item>
            <el-form-item label="工号">
                <el-input v-model="ruleForm.teacherId"></el-input>
            </el-form-item>
            <el-form-item label="电子邮箱">
                <el-input v-model="ruleForm.email"></el-input>
            </el-form-item>
            <el-form-item label="联系电话">
                <el-input v-model="ruleForm.phone"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">修改</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'updateStudents',
        data() {
            return {
                ruleForm: {
                    id: '',
                    name: '',
                    password: '',
                    teacherId: '',
                    email: '',
                    phone: '',
                },
                // rules: {
                //     name: [
                //         { required: true, message: '请输入活动名称', trigger: 'blur' },
                //         { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
                //     ],
                //     date1: [
                //         { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
                //     ],
                //     date2: [
                //         { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
                //     ],
                //     desc: [
                //         { required: true, message: '请填写活动形式', trigger: 'blur' }
                //     ],
                //     num: [
                //         { required: true, message: '年龄不能为空'},
                //         { type: 'number', message: '年龄必须为数字值'}
                //     ]
                // }
            };
        },
        beforeMount: function() {
            var index = this.$route.params.index;
            var data = this.$route.params.data;
            this.ruleForm.id = data.id
            this.ruleForm.name = data.name
            this.ruleForm.password = data.password
            this.ruleForm.teacherId = data.teacherId
            this.ruleForm.email = data.email
            this.ruleForm.phone = data.phone
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        var { name, password, teacherId, email, phone } = this.$refs[formName].model
                        var that = this
                        this.$axios.get('/api/teacher/update?id=' + id + '&name=' + name+'&password='+password+'&teacherId='+teacherId+'&email='+email+'&phone='+phone)
                            .then(function (response) {
                                if(response.status == 200) {
                                    console.log(response)
                                    that.$router.push({
                                        name: "ManageTeachers",
                                    });
                                }
                            })
                            .catch(function (error) {
                                console.log(error);
                            });
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        }
    }
</script>

<style>
    .teacher-update {
        position: fixed;
        height: 500px;
        width: 500px;
        top: 140px;
        margin-left: 700px;
    }
    .teacher-update .title {
        text-align: center;
        margin-bottom: 30px;
        margin-left: 30px;
        font-size: 20px;
    }
</style>
