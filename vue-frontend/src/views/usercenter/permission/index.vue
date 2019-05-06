<template>
  <div class="publish">
   {{ message }}
   {{ publishs }}
   <publish-list :value="publishs"></publish-list>
  </div>
</template>

<script>
import { getPublishList } from '@/api/books/publish'
import PublishList from './list'

export default {
  name: 'publish',
  components: {
    PublishList
  },
  data() {
    return {
      message: 'hello world',
      publishs: []
    }
  },
  methods: {
    fetchData() {
      getPublishList().then(
        res => {
          this.publishs = res.data.results
        }
      )
      console.log('welcome')
    }
  },
  created() {
    this.fetchData()
  }
}
</script>

<style lang='scss' scoped>
.publish {
  padding: 10px;
}
</style>

