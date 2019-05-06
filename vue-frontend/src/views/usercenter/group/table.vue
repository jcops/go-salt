<template>
  <div class="group-list">
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
        label="组名"
        prop="name"
        align="center"
      />

      <!--<el-table-column-->
      <!--label="成员管理"-->
      <!--prop="number"-->
      <!--align="center"-->
      <!--&gt;-->
      <!--<template slot-scope="scope">-->
      <!--<el-button-->
      <!--type="text"-->
      <!--size="small"-->
      <!--@click="handleGroupMember(scope.row)"-->
      <!--&gt;-->
      <!--{{ '成员列表（' + scope.row.member_count + '）人' }}-->
      <!--</el-button>-->
      <!--</template>-->
      <!--</el-table-column>-->

      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          <!--<el-button-->
          <!--size="mini"-->
          <!--@click="handleEdit(scope.row)"-->
          <!--&gt;编辑</el-button>-->

          <el-button
            size="mini"
            type="primary"
            @click="handlePower(scope.row)"
          >编辑</el-button>

          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
export default {
  name: 'GroupList',
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
  methods: {
    /* 点击编辑按钮，将子组件的事件传递给父组件 */
    handleEdit(value) {
      this.$emit('edit', value)
    },

    handlePower(value) {
      console.log('power', value)
      this.$emit('power', value)
    },

    handleGroupMember(value) {
      this.$emit('groupmember', value)
    },

    /* 删除 */
    handleDelete(group) {
      const id = group
      const name = group.name
      this.$confirm(`此操作将删除: ${name}, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$emit('delete', id)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    }
  }
}
</script>

<style lang='scss'>
.group-list {}
</style>

