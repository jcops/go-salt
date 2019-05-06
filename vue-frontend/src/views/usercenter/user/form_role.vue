<template>
  <div class="user-role">
    <el-form ref="form" :model="form" label-width="100px" class="demo-form">

      <el-form-item label="用户名" prop="username">
        <el-input v-model="form.username" :disabled="true" />
      </el-form-item>

      <el-form-item label="角色" prop="groups">
        <el-select v-model="form.role" multiple filterable style="width:100%" placeholder="请选择">
          <el-option
            v-for="(item, index) in role_list"
            :key="index"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
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

import { getRoleList } from '@/api/role'

export default {
  name: 'UserRole',
  props: {
    form: { // 接受父组件传递过来的值渲染表单
      type: Object,
      default() {
        return {
          username: '',
          role: []
        }
      }
    },
    createloading: {
      type: Boolean,
      required: true
    }
  },

  data() {
    return {
      role_list: [],
      res: {}
    }
  },

  created() {
    // this.state = 1
    getRoleList().then(res => {
      this.role_list = res.data.lists
      console.log('组', this.role_list)
    })
  },

  methods: {
    submitForm() {
      this.$refs.form.validate(valid => {
        if (!valid) {
          return
        }
        this.res['id'] = this.form.id
        this.res['role_id'] = this.form.role
        console.log(this.res)
        this.$emit('submit', this.form)
      })
    },
    cancel() {
      this.$emit('cancel')
    }
  }
  // watch: {
  //   state() {
  //     getGroupList().then(res => {
  //       this.role_list = res.results
  //       console.log(this.role_list)
  //     })
  //   }
  // }

}
</script>

<style lang='scss' scoped>
.user-role {
  position: relative;
  display: block;
  .btn-wrapper{
    text-align: right;
  }
}
</style>

