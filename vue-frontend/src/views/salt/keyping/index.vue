<template>
  <div class="edit-documents app-container">
    <el-form>
      <div class="el-col-11">
        <el-form-item label="主机" label-width="65px" prop="title">
          <el-input
            v-model="textarea"
            type="textarea"
            :rows="10"
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
          v-html="keylist.data"
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
import { saltKeyping } from '@/api/salt'
export default {
  name: 'Index',
  data() {
    return {
      textarea: '*',
      keylist: []
    }
  },
  // created() {
  //   this.fetchData()
  // },
  methods: {
    handleSubmit() {
      saltKeyping().then(res => {
        console.log('keylist', res.data)
        this.keylist = res
      })
    },
    handlecancel() {
      this.textarea = '*'
    }
  }
}

</script>
