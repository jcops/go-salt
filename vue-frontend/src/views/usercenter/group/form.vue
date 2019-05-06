<template>
  <div class="group-form">
    <el-form ref="form" :model="form" :rules="rules" label-width="100px" class="demo-form">

      <el-form-item label="组名" prop="groupname">
        <el-input v-model="form.name" placeholder="请输入组名" />
      </el-form-item>

      <el-form-item>
        <div class="btn-wrapper">
          <el-button size="small" @click="cancel">取消</el-button>
          <el-button size="small" type="primary" :loading="createLoading" @click="submitForm">保存</el-button>
        </div>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>

export default {
  name: 'GroupForm',
  props: {
    // 接受父组件传递过来的值渲染表单
    form: {
      type: Object,
      default() {
        return {
          name: ''
        }
      }
    },
    createLoading: Boolean
  },

  data() {
    return {
      rules: {
        name: [
          { required: true, message: '请输入姓名', trigger: 'blur' }
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
.group-form {
  position: relative;
  display: block;
  .btn-wrapper{
    text-align: right;
  }
}
</style>

