<template>
  <el-container>
    <HeadNavBar></HeadNavBar>
    <el-container>
      <el-aside>
        <AsideMenu>
        </AsideMenu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import HeadNavBar from '@/components/NavBar.vue'
import AsideMenu from '@/components/AsideMenu.vue'
export default {
  name: 'HomeView',
  components: {
    HeadNavBar,
    AsideMenu
  }
}
</script>
