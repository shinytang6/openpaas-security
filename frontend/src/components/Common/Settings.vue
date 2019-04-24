<template>
    <div class="settings">
        <div class="title">个人设置</div>
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="姓名">
                <el-input v-model="form.name"></el-input>
            </el-form-item>

            <el-form-item label="学号" v-show="identity=='学生'">
                <el-input v-model="form.studentId"></el-input>
            </el-form-item>
            <el-form-item label="工号" v-show="identity=='教师'">
                <el-input v-model="form.teacherId"></el-input>
            </el-form-item>
            <el-form-item label="所在班级" v-show="identity=='学生'">
                <el-input v-model="form.class"></el-input>
            </el-form-item>

            <el-form-item label="电子邮箱">
                <el-input v-model="form.email"></el-input>
            </el-form-item>
            <el-form-item label="联系电话">
                <el-input v-model="form.phone"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onSubmit">更新信息</el-button>
                <el-button @click="cancel">取消</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    import { setCookie, getCookie } from '../../js/cookieUtil'
    export default {
        name: 'settings',
        data() {
            return {
                form: {
                    name: '',
                    studentId: '',
                    teacherId: '',
                    class: '',
                    email: '',
                    phone: '',
                },
                identity: ""
            }
        },
        mounted () {
            this.identity = getCookie("identity");
            if (this.identity == "教师") {
                this.form.name = getCookie("name");
                this.form.email = getCookie("email");
                this.form.phone = getCookie("phone");
                this.form.teacherId = getCookie("teacherId");
            }

        },
        methods: {
            onSubmit() {
                var { name, teacherId, email, phone } = this.form
                var that = this
                if (this.identity == "教师") {
                    this.$axios.get('/api/teacher/update?name=' + name + '&teacherId=' + teacherId + '&email=' + email + '&phone=' + phone)
                        .then(function (response) {
                            if (response.status == 200) {
                                setCookie('name', name);
                                setCookie('email', email);
                                setCookie('phone', phone);
                                setCookie('teacherId', teacherId);

                                that.$message({
                                    showClose: true,
                                    message: '修改成功',
                                    type: 'success'
                                });
                            }
                        })
                        .catch(function (error) {
                            console.log(error);
                        });
                }
            },
            cancel() {
                this.$router.push({
                    name: "/",
                });
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
