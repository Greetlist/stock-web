<template>
  <!--
  <div class="header">
    <div class="collapse-btn" v-on:click="changeStatus">
      <i v-if="collapse" class="el-icon-s-unfold"></i>
      <i v-else class="el-icon-s-fold"></i>
    </div>
  </div>
  -->
  <el-container>
    <el-header>
      <HeadNavBar>
      </HeadNavBar>
    </el-header>
    <!--
    <el-aside>
      <AsideMenu></AsideMenu>
    </el-aside>
    -->
    <!--<transition :duration="{ enter: 800, leave: 100}" mode="out-in" name="el-fade-in-linear">-->
    <el-main>
      <router-view></router-view>
    </el-main>
    <!--</transition>-->
  </el-container>
</template>

<script>
// import AsideMenu from '@/components/AsideMenu.vue'
import HeadNavBar from '@/components/NavBar.vue'
export default {
  name: 'HomeView',
  components: {
    // AsideMenu
    HeadNavBar
  },
  methods: {
    changeStatus: function () {
      console.log(this.$store.state.collapse)
      this.$store.commit('changeCollapse', !this.$store.state.collapse)
    },
    calcAsideWidth: function () {
      return 
    }
  },
  created() {
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.$store.commit('changeCollapse', true)
    }
  },
  computed: {
    collapse: {
      get () {
        return this.$store.state.collapse
      }
    }
  }
}
</script>

<style>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  color: #fff;
  background-color: #0cb7ea;
}
.collapse-btn {
  float: left;
  padding: 0 21px;
  cursor: pointer;
  line-height: 70px;
}
.el-aside {
  height: 100vh;
  width: auto;
  background-color: #545c64;
}

.el-main {
  padding-top: 20vh;
}
</style>
