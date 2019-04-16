/* eslint-disable */
<template>
    <div class="experiment-addition">
        <div class="title">新建实验</div>
        <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="实验名称" prop="name">
                <el-input v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="实验时间" required>
                <el-col :span="11">
                    <el-form-item prop="date1">
                        <el-date-picker type="date" placeholder="选择日期" v-model="ruleForm.date1" style="width: 100%;"></el-date-picker>
                    </el-form-item>
                </el-col>
                <el-col class="line" :span="2">-</el-col>
                <el-col :span="11">
                    <el-form-item prop="date2">
                        <el-time-picker placeholder="选择时间" v-model="ruleForm.date2" style="width: 100%;"></el-time-picker>
                    </el-form-item>
                </el-col>
            </el-form-item>
            <el-form-item label="实验人数" prop="people">
                <el-input type="people" v-model.number="ruleForm.people" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'addExperiments',
        data() {
            return {
                ruleForm: {
                    name: '',
                    // region: '',
                    date1: '',
                    date2: '',
                    // resource: '',
                    // desc: '',
                    people: '',
                },
                rules: {
                    name: [
                        { required: true, message: '请输入活动名称', trigger: 'blur' },
                        { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
                    ],
                    date1: [
                        { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
                    ],
                    date2: [
                        { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
                    ],
                    // desc: [
                    //     { required: true, message: '请填写活动形式', trigger: 'blur' }
                    // ],
                    people: [
                        { required: true, message: '年龄不能为空'},
                        { type: 'number', message: '年龄必须为数字值'}
                    ]
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        var { name, date1, date2, people } = this.$refs[formName].model
                        var date_value =date2.getFullYear() + '-' + (date2.getMonth() + 1) + '-' + date2.getDate() + ' ' + date2.getHours() + ':' + date2.getMinutes() + ':' + date2.getSeconds();
                        console.log("date", date_value)
                        var that= this
                        this.$axios.get('/api/experiment/add?name='+name+'&people='+people+'&date='+date_value)
                            .then(function (response) {
                                if(response.status == 200) {
                                    // that.$router.push({
                                    //     name: "Experiment",
                                    //     // params: {
                                    //     //     id: activity_id,
                                    //     //     data: response.data.data
                                    //     // }
                                    // });
                                    console.log(response)
                                    that.$router.push({
                                        name: "ListExperiments",
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
    .experiment-addition {
        position: fixed;
        height: 500px;
        width: 500px;
        top: 140px;
        margin-left: 700px;
    }
    .experiment-addition .title {
        text-align: center;
        margin-bottom: 30px;
        margin-left: 30px;
        font-size: 20px;
    }
</style>
