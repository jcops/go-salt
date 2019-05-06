<template>
  <div ref="multipleTable" class="user-list" style="margin-top: 50px">
    <el-table
      v-loading="listLoading"
      :data="value"
      border
      stripe
      class="table"
      element-loading-text="给我一点时间"
      fit
      highlight-current-row
      tooltip-effect="dark"
      style="width: 100%"
    >

      <el-table-column
        type="index"
        label="ID"
        width="50"
      />

      <el-table-column
        label="用户名"

        prop="username"
      />

      <!--<el-table-column-->
      <!--label="昵称"-->

      <!--prop="nickname"-->
      <!--/>-->

      <!--<el-table-column-->
      <!--label="手机号"-->

      <!--prop="mobile"-->
      <!--/>-->

      <!--<el-table-column-->
      <!--label="是否为管理员"-->

      <!--type="scope"-->
      <!--prop="is_superuser"-->
      <!--&gt;-->
      <!--<template slot-scope="scope">-->
      <!--<div v-if="scope.row.is_superuser === true">-->
      <!--<span style="float: left;margin-left: 5px">-->
      <!--<el-tag disable-transitions>是</el-tag>-->
      <!--</span>-->
      <!--</div>-->

      <!--<div v-else>-->
      <!--<span style="float: left;margin-left: 5px">-->
      <!--<el-tag disable-transitions>否</el-tag>-->
      <!--</span>-->
      <!--</div>-->
      <!--</template>-->
      <!--</el-table-column>-->
      <el-table-column
        label="用户组"
        prop="role"
        type="scope"
      >
        <template slot-scope="scope">
          <div v-for="(item, index) in scope.row.role " :key="index">
            <span style="float: left;margin-left: 5px">
              <el-tag :key="index" disable-transitions>{{ item.name }}</el-tag>
            </span>

          </div>
        </template>
      </el-table-column>

      <el-table-column
        width="250"
        label="操作"
      >

        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"

            @click="handleEdit(scope.row)"
          >更新
          </el-button>

          <!--&lt;!&ndash;<div v-if="scope.row.is_superuser == true">&ndash;&gt;-->
          <!--<el-button-->
          <!--type="text"-->
          <!--size="small"-->
          <!--&gt;-->
          <!--<div v-if="scope.row.is_superuser === true">-->
          <!--&lt;!&ndash;取消管理员&ndash;&gt;-->
          <!--<span style="float: left;margin-left: 5px" @click="handleSetSuper(scope.row,0)">-->
          <!--取管理员-->
          <!--</span>-->
          <!--</div>-->
          <!--<div v-else>-->
          <!--<span style="float: left;margin-left: 5px" @click="handleSetSuper(scope.row,1)">-->
          <!--设管理员-->
          <!--</span>-->
          <!--</div>-->
          <!--</el-button>-->
          <!--<el-button-->
          <!--type="text"-->
          <!--size="small"-->
          <!--@click="handleRole(scope.row)"-->
          <!--&gt;设置组-->
          <!--</el-button>-->

          <el-button
            type="text"
            size="small"
            @click="handleDelete(scope.row)"
          >删除
          </el-button>
          <!--<el-button type="danger" icon="el-icon-delete" circle></el-button>-->
          <!--<el-button type="primary" icon="el-icon-edit" circle></el-button>-->
          <!--&lt;!&ndash;<el-button type="danger" icon="el-icon-delete" circle></el-button>&ndash;&gt;-->
          <!--<el-button type="danger" icon="el-icon-delete" circle></el-button>-->
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
export default {
  name: 'UserList',
  props: {
    value: {
      type: Array,
      required: false,
      default() {
        return []
      }
    },
    listLoading: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      loading: false,
      multipleSelection: []
    }
  },
  methods: {
    /* 点击编辑按钮，将子组件的事件传递给父组件 */
    handleEdit(value) {
      console.log(value)
      this.$emit('edit', value)
    },

    /* 点击角色按钮，将子组件的事件传递给父组件 */
    handleRole(value) {
      this.$emit('role', value)
    },
    // 设置用户为管理员
    handleSetSuper(value, id) {
      const uid = value
      const name = value.username
      if (id === 1) {
        this.$confirm(`确定设置用户: ${name}, 为超级管理员?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$emit('setsuper', uid, id)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          })
        })
      } else {
        this.$confirm(`确定取消用户: ${name}, 的超级管理员权限?`, '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.$emit('setsuper', uid, id)
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消'
          })
        })
      }
    },

    /* 删除 */
    handleDelete(user) {
      const id = user.id
      const name = user.username
      this.$confirm(`此操作将删除: ${name}, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$emit('delete', id, name)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    toggleSelection(rows) {
      console.log(rows)
      this.$emit('toggle', rows)
      // if (rows) {
      //   rows.forEach(row => {
      //     this.$refs.form.toggleRowSelection(row)
      //   })
      // } else {
      //   this.$refs.form.clearSelection()
      // }
    }
    // handleSelectionChange(val) {
    //   this.multipleSelection = val
    // }

  }
}
</script>

<style lang='scss'>
  .user-list {}
</style>

