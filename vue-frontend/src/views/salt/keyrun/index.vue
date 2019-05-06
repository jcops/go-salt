<template>
  <div class="edit-documents app-container">
    <el-form>
      <div class="el-col-11">
        <el-form-item label="主机" label-width="65px" prop="title">
          <el-input
            v-model="textarea.ip"
            type="textarea"
            :rows="10"
            placeholder="请输入内容"
          />
        </el-form-item>
        <el-form-item label="命令" label-width="65px">
          <el-input
            v-model="textarea.cmd"
            type="text"
            :rows="2"
            placeholder="请输入内容"
          />
        </el-form-item>
        <el-form-item label="操作" label-width="65px">
          <el-button type="default" size="small" @click="handlecancel">清空</el-button>
          <el-button type="primary" size="small" @click="handleSubmit">提交</el-button>
        </el-form-item>
      </div>
      <div class="el-col-11 ">
        <el-form-item label="状态" label-width="65px" prop="title">
          <pre>
          <el-input
          type="textarea"
          :rows="10"
          v-html="res.data"
          />
          </pre>

          <!--<textarea-->
          <!--:rows="13"-->
          <!--v-html="res.data">-->
          <!--</textarea>-->
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>
<script>
import { saltrun } from '@/api/salt'
export default {
  name: 'Index',
  data() {
    return {
      textarea: {
        ip: '*',
        cmd: 'ls'
      },
      res: []
    }
  },
  // created() {
  //   this.fetchData()
  // },
  methods: {
    handleSubmit() {
      saltrun({ match: this.textarea.ip, cmd: this.textarea.cmd }).then(res => {
        console.log('res', res.data)
        this.res = res
        // var str = res.replace(/\\n/gm,"<br/>");
      })
    },
    handlecancel() {
      this.textarea = '*'
    }
  }
}

</script>
