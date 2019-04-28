<template>
    <div class="settings">
        <div class="title">个人设置</div>
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="编号" disabled="true">
                <el-input v-model="form.id" :disabled="true"></el-input>
            </el-form-item>
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

            <el-form-item label="密码" v-show="identity=='用户管理员'">
                <el-input v-model="form.password"></el-input>
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
                    id: '',
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
                this.form.id = getCookie("id");
                this.form.name = getCookie("name");
                this.form.email = getCookie("email");
                this.form.phone = getCookie("phone");
                this.form.teacherId = getCookie("teacherId");
            } else if (this.identity == "学生") {
                this.form.id = getCookie("id");
                this.form.name = getCookie("name");
                this.form.email = getCookie("email");
                this.form.phone = getCookie("phone");
                this.form.class = getCookie("class");
                this.form.studentId = getCookie("studentId");
            } else if (this.identity == "用户管理员") {
                this.form.id = getCookie("id");
                this.form.name = getCookie("name");
                this.form.email = getCookie("email");
                this.form.phone = getCookie("phone");
            } else if (this.identity == "系统管理员") {
                this.form.id = getCookie("id");
                this.form.name = getCookie("name");
                this.form.email = getCookie("email");
                this.form.phone = getCookie("phone");
            }
        },
        methods: {
            onSubmit() {
                var that = this
                if (this.identity == "教师") {
                    var { id, name, teacherId, email, phone } = this.form
                    this.$axios.get('/api/teacher/update?id=' + id + '&name=' + name + '&teacherId=' + teacherId + '&email=' + email + '&phone=' + phone)
                        .then(function (response) {
                            if (response.status == 200) {
                                if (name != "") {
                                    setCookie('name', name);
                                }
                                if (email != "") {
                                    setCookie('email', email);
                                }
                                if (phone != "") {
                                    setCookie('phone', phone);
                                }
                                if (teacherId != "") {
                                    setCookie('teacherId', teacherId);
                                }

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
                } else if (this.identity == "学生") {
                    var { id, name, studentId, email, phone } = this.form
                    var clas = this.form.class
                    this.$axios.get('/api/student/update?id=' + id + '&name=' + name + '&studentId=' + studentId + '&class=' + clas + '&email=' + email + '&phone=' + phone)
                        .then(function (response) {
                            if (response.status == 200) {
                                if (name != "") {
                                    setCookie('name', name);
                                }
                                if (email != "") {
                                    setCookie('email', email);
                                }
                                if (phone != "") {
                                    setCookie('phone', phone);
                                }
                                if (clas != "") {
                                    setCookie('class', clas);
                                }
                                if (studentId != "") {
                                    setCookie('studentId', studentId);
                                }

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
                } else if (this.identity == "用户管理员") {
                    var { id, name, email, phone, password} = this.form
                    this.$axios.get('/api/userAdmin/update?id=' + id + '&name=' + name + '&password=' + password + '&email=' + email + '&phone=' + phone)
                        .then(function (response) {
                            if (response.status == 200) {
                                if (name != "") {
                                    setCookie('name', name);
                                }
                                if (email != "") {
                                    setCookie('email', email);
                                }
                                if (phone != "") {
                                    setCookie('phone', phone);
                                }

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
                } else if (this.identity == "系统管理员") {
                    var { id, name, email, phone} = this.form
                    this.$axios.get('/api/sysAdmin/update?id=' + id + '&name=' + name + '&email=' + email + '&phone=' + phone)
                        .then(function (response) {
                            if (response.status == 200) {
                                if (name != "") {
                                    setCookie('name', name);
                                }
                                if (email != "") {
                                    setCookie('email', email);
                                }
                                if (phone != "") {
                                    setCookie('phone', phone);
                                }

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
