<template>
    <div class="settings">
        <div class="title">个人设置</div>
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="姓名">
                <el-input v-model="form.name"></el-input>
            </el-form-item>
            <el-form-item label="学号">
                <el-input v-model="form.studentId"></el-input>
            </el-form-item>
            <el-form-item label="所在班级">
                <el-input v-model="form.class"></el-input>
            </el-form-item>
            <el-form-item label="电子邮箱">
                <el-input v-model="form.email"></el-input>
            </el-form-item>
            <el-form-item label="联系电话">
                <el-input v-model="form.phone"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onSubmit">立即创建</el-button>
                <el-button>取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'settings',
        data() {
            return {
                form: {
                    name: '',
                    studentId: '',
                    class: '',
                    email: '',
                    phone: '',
                }
            }
        },
        methods: {
            onSubmit() {
                console.log('submit!');
            }
        }
    }
</script>

<style>
    .settings {
        position: fixed;
        height: 500px;
        width: 500px;
        top: 140px;
        margin-left: 700px;
    }

    .settings .title {
        text-align: center;
        margin-bottom: 30px;
        margin-left: 30px;
        font-size: 20px;
    }
</style>
