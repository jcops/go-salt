<template>
  <div class="user-form">
    <el-form ref="form" :model="form" :rules="rules" label-width="100px" class="demo-form" :status-icon="true">

      <!--<el-form-item label="用户名" prop="username">-->
      <!--<el-input v-model="form.username" placeholder="请输入用户名" :disabled="form.sidd >0 ? true :false" />-->
      <!--</el-form-item>-->
      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" placeholder="请输入用户名" />
      </el-form-item>

      <!--<el-form-item label="昵称" prop="nickname">-->
      <!--<el-input v-model="form.nickname" placeholder="请输入昵称"></el-input>-->
      <!--</el-form-item>-->

      <!--<el-form-item label="手机号" prop="mobile">-->
      <!--<el-input v-model="form.mobile" placeholder="请输入手机号"></el-input>-->
      <!--</el-form-item>-->
      <!--<el-form-item label="邮箱" prop="email">-->
      <!--<el-input v-model="form.email" placeholder="请输入邮箱"></el-input>-->
      <!--</el-form-item>-->
      <el-form-item label="密码" prop="password">
        <el-input v-model="form.password" type="password" placeholder="请输入密码" />
      </el-form-item>
      <el-form-item>
        <div class="btn-wrapper">
          <el-button size="small" @click="cancel">取消</el-button>
          <el-button size="small" type="primary" :loading="createloading" @click="submitForm">保存</el-button>
        </div>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>

export default {
  name: 'UserForm',
  props: {
    // sid: 1,
    form: { // 接受父组件传递过来的值渲染表单
      type: Object,
      default() {
        return {
          username: '',
          password: ''
          // name: '',
          // mobile: '',
          // email: ''
        }
      }
    },
    createloading: {
      type: Boolean,
      required: true
    }
  },

  data() {
    // const chechMobile = (role, value, callback) => {
    //   const REGEX_MOBILE = /^1[358]\d{9}$|^147\d{8}$|^176\d{8}$/
    //   //
    //   // if (value) {
    //   //   // if (/^\d+$/.test(value)) {
    //   //   if (!REGEX_MOBILE.test(value)) {
    //   //     // if (value.length === 11) {
    //   //     //   callback()
    //   //     // } else {
    //   //     //   callback(new Error('请输入11位的手机号'))
    //   //     // }
    //   //   } else {
    //   //     callback(new Error('请输入正确的手机号'))
    //   //   }
    //   // } else {
    //   //   callback(new Error('请输入手机号'))
    //   // }
    //
    //   if (value) {
    //     if (!REGEX_MOBILE.test(value)) {
    //       callback(new Error('请输入正确的手机号'))
    //     } else {
    //       callback()
    //     }
    //   } else {
    //     callback(new Error('请输入手机号'))
    //   }
    // }
    const validatePass = (rule, value, callback) => { // 校验密码长度
      // if (!value || value.length < 8) {
      //   callback(new Error('密码不能小于8位'))
      // } else {
      //   callback(new Error('请输入密码'))
      // }

      if (value) {
        if (value.length < 8) {
          callback(new Error('密码不能小于8位'))
        } else {
          callback()
        }
      } else {
        callback(new Error('请输入密码'))
      }
    }
    // const validateEmail = (rule, value, callback) => {
    //   const regEmail = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/
    //   if (value) {
    //     if (!regEmail.test(value)) {
    //       callback(new Error('请填写正确的邮箱，如: abc@126.com '))
    //     } else {
    //       callback()
    //     }
    //   } else {
    //     callback(new Error('请输入邮箱'))
    //   }
    // }
    return {
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        // nickname: [
        //   { required: true, message: '请输入昵称', trigger: 'blur' }
        // ],
        // mobile: [
        //   { required: true, trigger: 'blur', validator: chechMobile }
        // ],
        // email: [
        //   { required: true, trigger: 'blur', validator: validateEmail }
        // ]
        password: [
          { required: true, trigger: 'blur', validator: validatePass }
        ]
      },
      state: 0
    }
  },

  methods: {
    submitForm() {
      this.$refs.form.validate(valid => {
        if (!valid) {
          return
        }
        this.$emit('submit', this.form)
      })
    },
    cancel() {
      this.$emit('cancel')
    }
  }

}
</script>

<style lang='scss' scoped>
.user-form {
  position: relative;
  display: block;
  .btn-wrapper{
    text-align: right;
  }
}
</style>

