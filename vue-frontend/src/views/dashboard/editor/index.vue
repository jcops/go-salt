<!--<template>-->
  <!--<div class="dashboard-editor-container">-->
    <!--&lt;!&ndash;<div class=" clearfix">&ndash;&gt;-->
      <!--&lt;!&ndash;&lt;!&ndash;<pan-thumb :image="avatar" style="float: left">&ndash;&gt;&ndash;&gt;-->
        <!--&lt;!&ndash;&lt;!&ndash;Your roles:&ndash;&gt;&ndash;&gt;-->
        <!--&lt;!&ndash;&lt;!&ndash;<span v-for="item in roles" :key="item" class="pan-info-roles">{{ item }}</span>&ndash;&gt;&ndash;&gt;-->
      <!--&lt;!&ndash;&lt;!&ndash;</pan-thumb>&ndash;&gt;&ndash;&gt;-->
      <!--&lt;!&ndash;<github-corner style="position: absolute; top: 0px; border: 0; right: 0;" />&ndash;&gt;-->
      <!--&lt;!&ndash;<div class="info-container">&ndash;&gt;-->
        <!--&lt;!&ndash;<span class="display_name">{{ name }}</span>&ndash;&gt;-->
        <!--&lt;!&ndash;<span style="font-size:20px;padding-top:20px;display:inline-block;">Editor's Dashboard</span>&ndash;&gt;-->
      <!--&lt;!&ndash;</div>&ndash;&gt;-->
    <!--&lt;!&ndash;</div>&ndash;&gt;-->
    <!--<div>-->
      <!--<img :src="emptyGif" class="emptyGif">-->
    <!--</div>-->
  <!--</div>-->
<!--</template>-->

<!--<script>-->
<!--import { mapGetters } from 'vuex'-->
<!--import PanThumb from '@/components/PanThumb'-->
<!--import GithubCorner from '@/components/GithubCorner'-->

<!--export default {-->
  <!--name: 'DashboardEditor',-->
  <!--components: { PanThumb, GithubCorner },-->
  <!--data() {-->
    <!--return {-->
      <!--emptyGif: 'https://wpimg.wallstcn.com/0e03b7da-db9e-4819-ba10-9016ddfdaed3'-->
    <!--}-->
  <!--},-->
  <!--computed: {-->
    <!--...mapGetters([-->
      <!--'name',-->
      <!--'avatar',-->
      <!--'roles'-->
    <!--])-->
  <!--}-->
<!--}-->
<!--</script>-->

<!--<style lang="scss" scoped>-->
  <!--.emptyGif {-->
    <!--display: block;-->
    <!--width: 45%;-->
    <!--margin: 0 auto;-->
  <!--}-->

  <!--.dashboard-editor-container {-->
    <!--background-color: #e3e3e3;-->
    <!--min-height: 100vh;-->
    <!--padding: 50px 60px 0px;-->
    <!--.pan-info-roles {-->
      <!--font-size: 12px;-->
      <!--font-weight: 700;-->
      <!--color: #333;-->
      <!--display: block;-->
    <!--}-->
    <!--.info-container {-->
      <!--position: relative;-->
      <!--margin-left: 190px;-->
      <!--height: 150px;-->
      <!--line-height: 200px;-->
      <!--.display_name {-->
        <!--font-size: 48px;-->
        <!--line-height: 48px;-->
        <!--color: #212121;-->
        <!--position: absolute;-->
        <!--top: 25px;-->
      <!--}-->
    <!--}-->
  <!--}-->
<!--</style>-->
