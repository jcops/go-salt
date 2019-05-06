<template>
  <div class="edit-documents app-container">
    <el-form>
      <div class="el-col-11">
        <el-form-item label="已认证" label-width="65px" prop="title">
          <el-select v-model="match" placeholder="请选择">
            <el-option
              v-for="(item, key) in keylist.minions"
              :key="key"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="操作" label-width="65px">
          <el-button
            type="primary"
            size="small"
            @click="handledeleteSubmit"
          >删除节点
          </el-button>
        </el-form-item>
      </div>
      <div class="el-col-11">
        <el-form-item label="未认证" label-width="65px" prop="title">
          <el-select v-model="match_pre" placeholder="请选择">
            <el-option
              v-for="(item, key) in keylist.minions_pre"
              :key="key"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="操作" label-width="65px">
          <el-button
            type="primary"
            size="small"
            @click="handleacceptSubmit"
          >认证节点
          </el-button>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>
<script>
import { saltlist, saltdelete, saltaccept } from '@/api/salt'
export default {
  name: 'Index',
  data() {
    return {
      keylist: {
        minions_pre: '',
        minions: ''
      },
      match: '',
      match_pre: ''
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      saltlist().then(res => {
        console.log('keylist', res.data)
        console.log(this.keylist.minions_pre)
        this.keylist.minions_pre = res.data.minions_pre
        this.keylist.minions = res.data.minions
      })
    },
    handledeleteSubmit() {
      console.log(this.match)
      saltdelete({ match: this.match }).then(res => {
        this.$notify({
          title: '成功',
          message: '删除节点成功',
          type: 'success'
        })
        this.fetchData()
      })
    },
    handleacceptSubmit() {
      console.log(this.match_pre)
      saltaccept({ match: this.match_pre }).then(res => {
        this.$notify({
          title: '成功',
          message: '认证节点成功',
          type: 'success'
        })
        this.fetchData()
      })
    }
  }
}

</script>
