/* eslint-disable */
<template>
    <div class="experiment-update">
        <div class="title">更新实验</div>
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
            <el-form-item label="实验人数" prop="age">
                <el-input type="age" v-model.number="ruleForm.num" autocomplete="off"></el-input>
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
        name: 'updateExperiments',
        data() {
            return {
                ruleForm: {
                    name: '',
                    region: '',
                    date1: '',
                    date2: '',
                    delivery: false,
                    type: [],
                    resource: '',
                    desc: '',
                    num: '',
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
                    desc: [
                        { required: true, message: '请填写活动形式', trigger: 'blur' }
                    ],
                    num: [
                        { required: true, message: '年龄不能为空'},
                        { type: 'number', message: '年龄必须为数字值'}
                    ]
                }
            };
        },
        beforeMount: function() {
            var index = this.$route.params.index;
            var data = this.$route.params.data;
            this.ruleForm.name = data.name
            this.ruleForm.date1 = data.date1
            this.ruleForm.date2 = data.date2
            this.ruleForm.num = data.num
            console.log(this.ruleForm, "fuck!!")
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        var { name, date1, date2, age } = this.$refs[formName].model
                        this.$axios.get('/api/experiment/add?name='+name+'&experimentNum='+age)
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
    .experiment-update {
        position: fixed;
        height: 500px;
        width: 500px;
        top: 140px;
        margin-left: 700px;
    }
    .experiment-update .title {
        text-align: center;
        margin-bottom: 30px;
        margin-left: 30px;
        font-size: 20px;
    }
</style>
