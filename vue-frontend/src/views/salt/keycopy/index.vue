<template>
  <div class="edit-documents app-container">
    <el-form>
      <div class="el-col-24">
        <el-form-item label="节点" label-width="65px" prop="title">
          <el-input
            v-model="textarea.ip"
            type="textarea"
            :rows="3"
            placeholder="输入IP"
          />
        </el-form-item>
        <div class="el-col-11">
          <el-form-item label="源" label-width="65px">

            <el-input
              v-model="textarea.src"
              type="text"
              :rows="3"
              placeholder="salt://"
            />

          </el-form-item>
        </div>
        <div class="el-col-11">
          <el-form-item label="目标" label-width="65px">

            <el-input
              v-model="textarea.dst"
              type="text"
              :rows="3"
              placeholder="如: /root/readme.txt"
            />

          </el-form-item>
        </div>
        <div class="el-col-24">
          <el-form-item label="操作" label-width="65px">
            <el-button
              type="primary"
              size="small"
              @click="handleCopySubmit"
            >执行复制
            </el-button>
          </el-form-item>
        </div>
        <div class="el-col-24">
          <el-form-item v-loading="CreateLoading" label="结果" label-width="65px" prop="title">
            <el-input
              type="textarea"
              :rows="3"
              v-html="textarea.res"
            />
          </el-form-item>
        </div>
      </div></el-form>
  </div>
</template>
<script>
import { saltcopy } from '@/api/salt'
export default {
  name: 'Index',
  data() {
    return {
      CreateLoading: false,
      textarea: {
        ip: '',
        src: '',
        dst: '',
        run: 'cp.get_url',
        res: []
      }
    }
  },
  methods: {
    handleCopySubmit() {
      console.log(this.match)
      if (this.textarea.radio === 1) {
        this.textarea.radio = 'redis'
      } else if (this.textarea.radio === 2) {
        this.textarea.radio = 'nginx'
      } else {
        this.textarea.radio = 'system_init'
      }
      console.log(this.textarea.radio)
      this.$confirm(`执行复制, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.CreateLoading = true
        saltcopy({ tgt: this.textarea.ip, arg1: this.textarea.src, arg2: this.textarea.dst, run: this.textarea.run }).then(res => {
          if (res.msg === 'ok') {
            this.CreateLoading = false
            console.log(res)
            this.textarea.res = res.data.data
            this.$notify({
              title: '成功',
              message: '成功',
              type: 'success'
            })
          }
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消'
        })
      })
    }
  }
}

</script>
