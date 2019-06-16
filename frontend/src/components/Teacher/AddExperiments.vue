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
            <el-form-item label="配置文件" prop="config">
                <el-input v-model="ruleForm.config"></el-input>
            </el-form-item>
            <el-form-item label="实验人数" prop="people">
                <el-input type="people" v-model.number="ruleForm.people" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="实验描述">
                <el-input type="textarea" v-model="ruleForm.desc" :rows="6"></el-input>
            </el-form-item>
            <el-form-item label="实验指导书" prop="guide">
                <el-upload
                        class="upload"
                        ref="upload"
                        action=""
                        accept=".pdf,.doc,.docs,.txt,.PDF,.DOC,.DOCS,.TXT,.xls,.xlsx"
                        :http-request="handleFile"
                        :on-change="guide_path_file"
                        :multiple="false"
                        :limit="1"
                        :file-list="guide_path">
                    <el-button size="small" type="primary" @click="uploadGuide('upload')">点击上传</el-button>
                    <!--<div slot="tip" class="el-upload__tip">上传实验指导书</div>-->
                </el-upload>
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
                    date1: '',
                    date2: '',
                    desc: '',
                    description: '',
                    config: '',
                    people: '',
                },
                guide_path: "",
                rules: {
                    name: [
                        { required: true, message: '请输入实验名称', trigger: 'blur' },
                    ],
                    date1: [
                        { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
                    ],
                    date2: [
                        { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
                    ],
                    desc: [
                        { required: true, message: '请输入实验描述', trigger: 'blur' },
                    ],
                    config: [
                        { required: true, message: '请输入配置文件地址', trigger: 'blur' },
                    ],
                    people: [
                        { required: true, message: '人数不能为空'},
                        { type: 'number', message: '人数必须为数字值'}
                    ],
                }
            };
        },
        methods: {
            handleFile() {

            },
            uploadGuide(type) {
                this.$refs.upload.clearFiles();
                this.guide_path = [];
            },
            guide_path_file(file, fileList) {
                this.guide_path = fileList
            },
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        var { name, date1, date2, desc, config, people } = this.$refs[formName].model
                        var date_value = date2.getFullYear() + '-' + (date2.getMonth() + 1) + '-' + date2.getDate() + ' ' + date2.getHours() + ':' + date2.getMinutes() + ':' + date2.getSeconds();

                        let formData = new FormData()
                        formData.append('name', name)
                        formData.append('date', date_value)
                        formData.append('desc', desc)
                        formData.append('config', config)
                        formData.append('people', people)
                        formData.append('guide_path', this.guide_path[0] ? this.guide_path[0].raw : '')

                        let conf = {
                            headers: {
                                'Content-Type': 'multipart/form-data'
                            }
                        }

                        var that= this
                        this.$axios.post('/api/experiment/add', formData, conf)
                            .then(function (response) {
                                if(response.status == 200) {
                                    that.$message({
                                        showClose: true,
                                        message: '实验添加成功',
                                        type: 'success'
                                    });
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
