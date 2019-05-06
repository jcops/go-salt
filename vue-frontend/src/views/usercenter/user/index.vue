<template>
  <div class="user">
    <div>
      <!--搜索-->
      <el-col :span="8">
        <el-input v-model="params.search" placeholder="搜索" @keyup.enter.native="searchClick">
          <el-button slot="append" icon="el-icon-search" @click="searchClick" />
        </el-input>
      </el-col>

      <!--添加按钮-->
      <el-col :span="16" style="text-align: right">
        <el-button type="primary" @click="handleAddBtn">创建用户</el-button>
      </el-col>
    </div>

    <!--表格-->
    <user-list
      ref="userTable"
      :value="users"
      :list-loading="listLoading"
      @toggle="toggleSelection"
      @delete="handleDelete"
      @edit="handleEdit"
    />
    <!--:list-loading="listLoading"-->
    <!--@edit="handleEdit"-->
    <!--@role="handleRole"-->

    <!--@setsuper="handleSetsSuper"-->
    <!--模态窗增加表单-->
    <el-dialog
      title="用户创建"
      :visible.sync="dialogVisibleForAdd"
      width="50%"
    >
      <user-form
        ref="userForm"
        :createloading="CreateLoading"
        @submit="handleSubmitAdd"
        @cancel="handleCancelAdd"
      />
      <!--</user-form>-->
    </el-dialog>

    <!--模态窗更新表单-->
    <el-dialog
      title="用户更新"
      :visible.sync="dialogVisibleForEdit"
      width="50%"
    >
      <user-form
        ref="userForm"
        :createloading="CreateLoading"
        :form="currentValue"
        @submit="handleSubmitEdit"
        @cancel="handleCancelEdit"
      />
    </el-dialog>

    <!--模态窗角色表单-->
    <el-dialog
      title="分配角色"
      :visible.sync="dialogVisibleForRole"
      width="50%"
    >
      <user-role
        ref="userRole"
        :form="currentValue"
        :createloading="CreateLoading"
        @submit="handleSubmitRole"
        @cancel="handleCancelRole"
      />
    </el-dialog>
    <div>
      <center style="margin-bottom: 40px">
        <el-pagination
          background
          layout="total, prev, pager, next, jumper"
          :page-sizes="[10,20,50]"
          :total="totalNum"
          @size-change="CurrentSizesChange"
          @current-change="handleCurrentChange"
        />
      </center>
    </div>
  </div>
</template>

<script>
import { getUserList, createUser, updateUser, deleteUser, setUserSuper } from '@/api/user'
import UserList from './table'
import UserForm from './form'
import UserRole from './form_role'

export default {
  name: 'User',
  components: {
    UserList,
    UserForm,
    UserRole
  },
  data() {
    return {
      dialogVisibleForAdd: false,
      dialogVisibleForEdit: false,
      dialogVisibleForRole: false,
      currentValue: {},
      CreateLoading: false,
      users: [],
      totalNum: 0,
      params: {
        page: 1
        // search: '',
        // page_size: 15
      },
      listLoading: false,
      updateUserLoading: false,
      sid: 1
      // multipleTable: []
    }
  },

  created() {
    this.fetchData()
  },

  methods: {
    fetchData() {
      this.listLoading = true
      getUserList(this.params).then(res => {
        console.log('lists', res.data.lists)
        this.users = res.data.lists
        this.totalNum = res.data.total
        this.listLoading = false
        // console.log('获取的用户信息', res)
      })
    },
    CurrentSizesChange(val) {
      this.params.page_size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.params.page = val
      this.fetchData()
    },
    searchClick() {
      this.fetchData()
    },
    /* 添加,弹出模态窗、提交数据、取消 */
    handleAddBtn() {
      this.dialogVisibleForAdd = true
    },
    handleSubmitAdd(value) {
      console.log('表单', value)
      this.CreateLoading = true
      createUser(value).then(res => {
        this.$notify({
          title: '成功',
          message: `创建用户: ${value.username}成功`,
          type: 'success'
        })
        this.CreateLoading = false
        this.handleCancelAdd()
        this.fetchData()
      })
    },
    handleCancelAdd() {
      this.dialogVisibleForAdd = false
      this.$refs.userForm.$refs.form.resetFields() // 清除掉子组件表单的数据
    },

    /* 更新，弹出模态窗、提交数据、取消 */
    // handleEdit(value) {
    //   this.currentValue = { ...value } // 将table子组件传来的值给父组件的变量，然后渲染更新表单
    //   // this.currentValue['sidd'] = 1
    //   // console.log(this.currentValue)
    //   this.dialogVisibleForEdit = true
    // },

    handleSubmitEdit(value) {
      const { id, ...params } = value // 很神奇，能把表单的值拆解成自己想要的样子
      // console.log('更改用户', id, params)
      this.CreateLoading = true
      updateUser(id, params).then(res => {
        this.$notify({
          title: '成功',
          message: `更新用户: ${params.username} 信息成功`,
          type: 'success'
        })
        this.CreateLoading = false
        this.handleCancelEdit()
        this.fetchData()
      })
    },
    handleCancelEdit() {
      this.dialogVisibleForEdit = false
      this.$refs.userForm.$refs.form.resetFields()
    },
    /* 分配角色，弹出模态窗、提交数据、取消 */
    handleEdit: function(value) {
      this.currentValue = { ...value } // 将table子组件传来的值给父组件的变量，然后渲染更新表单
      console.log('value', this.currentValue['role'])
      if (this.currentValue['role'] != null || this.currentValue['role'] > 0) {
        console.log(this.currentValue)
        this.currentValue['role'] = this.currentValue['role'].map(it => it.id)
        console.log('abc', this.currentValue)
        console.log('vvvvv', this.currentValue['role'].map(it => it.id))
      } else {
        this.currentValue['role'] = []
      }
      //   console.log('roleid', this.currentValue['role'])
      // this.currentValue['role'].map(val => {
      //   console.log(this.currentValue)
      //   console.log('val', val)
      //   if (val !== '') {
      //     this.currentValue['role'] = val
      //     console.log('abc', this.currentValue)
      //   } else {
      //     console.log('val11111', val)
      //     this.currentValue['role'] = []
      //   }
      // }
      // this.currentValue['role'] = this.currentValue['role'].map(it => it.id)
      // this.current['user_permissions'] = this.currentValue['user_permissions'].map(it => it.id)
      this.dialogVisibleForRole = true
      // // this.chooseRole = Object.assign({}, row)
    },
    handleSubmitRole(value) {
      console.log('update user value', value)
      console.log({ id: value.id, role_id: value.role })
      // const { params } = value // 很神奇，能把表单的值拆解成自己想要的样子
      // console.log(params)
      // console.log(id, params)
      // console.log({ id: params.id, role_id: params.role })
      updateUser({ id: value.id, role_id: value.role }).then(res => {
        this.$notify({
          title: '成功',
          message: `更新用户: ${value.username} 信息成功`,
          type: 'success'
        })
        this.handleCancelRole()
        this.fetchData()
      })
    },
    handleCancelRole() {
      this.dialogVisibleForRole = false
      this.$refs.userRole.$refs.form.resetFields()
    },
    // 设置管理员
    handleSetsSuper(value, sid) {
      setUserSuper(value.id, { 'superid': sid }).then(res => {
        if (res.status === 1) {
          this.$notify({
            title: '成功',
            message: `设置用户: ${value.username} 管理员成功`,
            type: 'success'
          })
          this.fetchData()
        } else {
          this.$notify({
            title: '成功',
            message: `取消用户: ${value.username} 管理员成功`,
            type: 'success'
          })
        }
      })
      this.fetchData()
    },
    /* 删除 */
    handleDelete(id, name) {
      deleteUser(id).then(res => {
        this.$notify({
          title: '成功',
          message: `删除用户: ${name} 成功`,
          type: 'success'
        })
        this.params.page = 1
        this.fetchData()
      },
      err => {
        console.log(err.message)
      })
    },
    toggleSelection(rows) {
      // this.multipleTable = rows
      console.log(rows)
      if (rows) {
        rows.forEach(row => {
          this.$refs.userTable.$refs.multipleTable.toggleRowSelection(row)
        })
      } else {
        this.$refs.userTable.$refs.multipleTable.clearSelection()
      }
    }
    // toggleSelectionCancel() {
    //   const multipleTable = this.multipleTable
    //   console.log(multipleTable)
    //   if (multipleTable) {
    //     multipleTable.forEach(res => {
    //       this.$refs.userTable.$refs.form.toggleRowSelection(res)
    //     })
    //   } else {
    //     this.$refs.userTable.$refs.form.clearSelection()
    //   }
    // }

  }

}
</script>

<style lang='scss' scoped>
.user {
  padding: 10px;
}
</style>

