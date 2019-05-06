<template>
  <div class="group">
    <div>
      <!--搜索-->
      <el-col :span="8">
        <el-input v-model="params.search" placeholder="搜索" @keyup.enter.native="searchClick">
          <el-button slot="append" icon="el-icon-search" @click="searchClick" />
        </el-input>
      </el-col>

      <!--添加按钮-->
      <el-col :span="16" style="text-align: right">
        <el-button type="primary" @click="handleAddBtn">创建组</el-button>
      </el-col>
    </div>

    <!--表格-->
    <group-list
      :value="groups"
      :list-loading="listLoading"
      @edit="handleEdit"
      @groupmember="handleGroupMember"
      @power="handlePower"
      @delete="handleDelete"
    />

    <!--模态窗增加表单-->
    <el-dialog
      title="添加"
      :visible.sync="dialogVisibleForAdd"
      width="50%"
    >
      <group-form
        ref="groupForm"
        :create-loading="CreateLoading"
        @submit="handleSubmitAdd"
        @cancel="handleCancelAdd"
      />
    </el-dialog>

    <!--模态窗更新表单-->
    <el-dialog
      title="更新"
      :visible.sync="dialogVisibleForEdit"
      width="50%"
    >
      <group-form
        ref="groupForm"
        :create-loading="CreateLoading"
        :form="currentValue"
        @submit="handleSubmitEdit"
        @cancel="handleCancelEdit"
      />
    </el-dialog>

    <!--模态窗权限表单-->
    <el-dialog
      title="权限"
      :visible.sync="dialogVisibleForPower"
      width="50%"
    >
      <role-power
        ref="powerForm"
        :form="currentValue"
        :create-loading="CreateLoading"
        @submit="handleSubmitPower"
        @cancel="handleCancelPower"
      />
    </el-dialog>

    <!--模态窗成员管理表格-->
    <el-dialog
      title="成员管理"
      :visible.sync="dialogVisibleForMember"
      width="50%"
    >
      <group-member
        :value="members"
        @deletemember="handleDeleteMember"
      />
    </el-dialog>

    <!--分页-->
    <div style="margin-bottom: 40px">
      <center>
        <el-pagination
          background
          layout="total, prev, pager, next, jumper"
          :page-sizes="[5,10,15]"
          :total="totalNum"
          @size-change="CurrentSizesChange"
          @current-change="handleCurrentChange"
        />
      </center>
    </div>
  </div>
</template>

<script>
import { getRoleList, createGroup, updateGroup, deleteGroup, updateGroupPower, deleteGroupMember } from '@/api/role'
import GroupList from './table'
import GroupForm from './form'
import RolePower from './power_form'
import GroupMember from './groupmember'

export default {
  name: 'Group',
  components: {
    GroupList,
    GroupForm,
    RolePower,
    GroupMember
  },

  data() {
    return {
      CreateLoading: false,
      dialogVisibleForAdd: false,
      dialogVisibleForEdit: false,
      dialogVisibleForPower: false,
      dialogVisibleForMember: false,
      currentValue: {},
      listLoading: false,
      groups: [],
      members: [],
      totalNum: 0,
      pagesize: 2,
      params: {
        page: 1,
        search: '',
        page_size: 10
      }
    }
  },

  created() {
    this.fetchData()
  },

  methods: {
    fetchData() {
      this.listLoading = true
      getRoleList(this.params).then(
        res => {
          this.groups = res.data.lists
          this.totalNum = res.data.total
          this.listLoading = false
        })
    },
    handleCurrentChange(val) {
      this.params.page = val
      this.fetchData()
      // console.log(this.params.page)
    },
    CurrentSizesChange(val) {
      this.params.page_size = val
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
      this.CreateLoading = true
      createGroup(value).then(res => {
        this.$notify({
          title: '成功',
          message: `创建组: ${value.name}成功`,
          type: 'success'
        })
        this.CreateLoading = false
        this.handleCancelAdd()
        this.fetchData()
      })
    },
    handleCancelAdd() {
      this.dialogVisibleForAdd = false
      this.$refs.groupForm.$refs.form.resetFields() // 清除掉子组件表单的数据
    },

    /* 更新，弹出模态窗、提交数据、取消 */
    handleEdit(value) {
      this.currentValue = { ...value } // 将table子组件传来的值给父组件的变量，然后渲染更新表单
      console.log(this.currentValue)
      this.dialogVisibleForEdit = true
    },

    handleSubmitEdit(value) {
      console.log(value)
      const { id, ...params } = value // 很神奇，能把表单的值拆解成自己想要的样子
      // console.log(id)
      // console.log(params)
      this.CreateLoading = true
      updateGroup(id, params).then(res => {
        this.$notify({
          title: '成功',
          message: `更新组: ${params.name} 信息成功`,
          type: 'success'
        })
        this.CreateLoading = false
        this.handleCancelEdit()
        this.fetchData()
      })
    },
    handleCancelEdit() {
      this.dialogVisibleForEdit = false
      this.$refs.groupForm.$refs.form.resetFields()
    },

    /* 权限，弹出模态窗、提交数据、取消 */
    handlePower(value) {
      this.currentValue = { ...value } // 将table子组件传来的值给父组件的变量，然后渲染更新表单
      // this.currentValue['menu'] = this.currentValue['menu'].map(it => it.id)
      // console.log(this.currentValue)
      console.log('value', this.currentValue['menu'])
      if (this.currentValue['menu'] != null || this.currentValue['menu'] > 0) {
        console.log(this.currentValue)
        this.currentValue['menu'] = this.currentValue['menu'].map(it => it.id)
        console.log('abc', this.currentValue)
        console.log('vvvvv', this.currentValue['menu'])
      } else {
        this.currentValue['menu'] = []
      }
      this.dialogVisibleForPower = true
    },

    handleSubmitPower(value) {
      console.log(value)
      // const { id, ...params } = value // 很神奇，能把表单的值拆解成自己想要的样子
      this.CreateLoading = true
      updateGroupPower({ id: value.id, menu_id: value.menu }).then(res => {
        this.$notify({
          title: '成功',
          // message: `更新组: ${params.name} 权限信息成功`,
          type: 'success'
        })
        this.CreateLoading = false
        this.handleCancelPower()
        this.fetchData()
      })
    },
    handleCancelPower() {
      this.dialogVisibleForPower = false
      this.$refs.powerForm.$refs.form.resetFields()
    },

    /* 弹出成员管理模态窗，并将组对成员渲染成到表格 */
    handleGroupMember(value) {
      this.currentValue = { ...value } // 将table子组件传来的值给父组件的变量，然后渲染更新表单
      this.members = this.currentValue.members // 直接将当前组的成员对象返回即可
      this.dialogVisibleForMember = true
    },

    /* 将组成员从组中删除 */
    handleDeleteMember(uid) {
      deleteGroupMember(this.currentValue.id, { 'uid': uid.id }).then(res => {
        this.$notify({
          title: '成功',
          message: `删除组成员: ${uid.username} 成功`,
          type: 'success'
        })
        this.params.page = 1
        this.dialogVisibleForMember = false
        this.fetchData()
      },
      err => {
        console.log(err.message)
      })
    },

    /* 删除 */
    handleDelete(id) {
      deleteGroup(id.id).then(res => {
        this.$notify({
          title: '成功',
          message: `删除组: ${id.name} 成功`,
          type: 'success'
        })
        this.params.page = 1
        this.fetchData()
      },
      err => {
        console.log(err.message)
      })
    }

  }

}
</script>

<style lang='scss' scoped>
.group {
  padding: 10px;
}
</style>

