<template>
  <div class="role_power">
    <el-form ref="form" :model="form" :rules="rules" label-width="100px" class="demo-form">

      <el-form-item label="组名" prop="name">
        <el-input v-model="form.name" readonly />
      </el-form-item>

      <el-form-item label="权限" prop="power">
        <el-transfer
          v-model="form.menu"
          filterable
          :titles="['权限列表', '角色权限']"
          :props="{
            key: 'id',
            label: 'name'
          }"
          :data="power_list"
        />
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

import { getPermissionList } from '@/api/role'

export default {
  name: 'RolePower',
  props: {
    form: { // 接受父组件传递过来的值渲染表单
      type: Object,
      default() {
        return {
          name: '',
          power: []
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
      power_list: [],
      state: 0
    }
  },
  watch: {
    state() {
      getPermissionList().then(res => {
        this.power_list = res.data.lists
        console.log(this.power_list)
      })
    }
  },

  created() {
    this.state = 1
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
.role-power {
  position: relative;
  display: block;
  .btn-wrapper{
    text-align: right;
  }
}
</style>

