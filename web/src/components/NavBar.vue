<template>
  <!--
  <div>
    <b-navbar toggleable="lg" type="dark" variant="dark" fixed="top">
      <b-navbar-brand href="/">
        <img src="https://wow.zamimg.com/images/wow/icons/medium/inv_helmet_74.jpg" alt="Light">
      </b-navbar-brand>
      <b-navbar-nav class="ml-auto">
        <b-button variant="dark" @click="showLoginDialog">Login</b-button>
        <b-button variant="dark">Contact Us</b-button>
      </b-navbar-nav>
      <b-form-input v-model="text" placeholder="Enter your name"></b-form-input>
    </b-navbar>
  </div>
  -->
  <el-menu
    :default-active="active"
    mode="horizontal"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b"
    :uniqueOpened="true"
    router
  >
    <el-sub-menu v-for="(singleMenu, idx) in menuList" :index="idx.toString()" :key="singleMenu.id">
      <template #title>
        <i :class="singleMenu.headIcon"></i>
        <span>{{ singleMenu.name }}</span>
      </template>
      <el-menu-item-group>
        <template #title> {{ singleMenu.name }}</template>
        <el-menu-item v-for="item in singleMenu.subMenuList" :index="item.route" :key="item.id">
        {{ item.name }}
        </el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>
  </el-menu>
</template>

<script>
var menuList = [
  {
    name: 'Stock OverView',
    id: 'overview',
    headIcon: 'el-icon-s-data',
    subMenuList: [
      {
        name: 'Daily OverView',
        id: 'daily-overview',
        route: '/dailyRecommand'
      },
      {
        name: 'Query Stock',
        id: 'query-stock',
        route: '/queryStockView'
      },
      {
        name: 'Total Market',
        id: 'total-market',
        route: '/totalMarket'
      }
    ]
  },
  {
    name: 'Stock Position',
    id: 'position',
    headIcon: 'el-icon-cpu',
    subMenuList: [
      {
        name: 'Stock Hold',
        id: 'get-stock-hold',
        route: '/getStockHold'
      },
      {
        name: 'Statistic',
        id: 'statistic',
        route: '#'
      }
    ]
  },
  {
    name: 'DailyTask',
    id: 'daily-task',
    headIcon: 'el-icon-alarm-clock',
    subMenuList: [
      {
        name: 'Fetch Stock Data',
        id: 'fetch-stock-data',
        route: '#'
      },
      {
        name: 'Compute All Stock',
        id: 'compute-all-stock',
        route: '#'
      },
      {
        name: 'Run Statistic',
        id: 'run-statistic',
        route: '#'
      }
    ]
  },
  {
    name: 'News',
    id: 'news',
    headIcon: 'el-icon-s-comment',
    subMenuList: [
      {
        name: 'Investor Number',
        id: 'investor-number',
        route: '#'
      },
      {
        name: 'Annual Report',
        id: 'annual-report',
        route: '#'
      }
    ]
  }
//  {
//    name: 'Report',
//    id: 'report',
//    headIcon: 'el-icon-printer',
//    subMenuList: [
//      {
//        name: 'Index OverView',
//        id: 'index-overview',
//        route: '#'
//      },
//      {
//        name: 'Other Format',
//        id: 'other-format',
//        route: '#'
//      }
//    ]
//  },
//  {
//    name: 'Settings',
//    id: 'settings',
//    headIcon: 'el-icon-setting',
//    subMenuList: [
//      {
//        name: 'Server Settings',
//        id: 'server-settings',
//        route: '#'
//      },
//      {
//        name: 'Other Settings',
//        id: 'other-settings',
//        route: '#'
//      }
//    ]
//  },
//  {
//    name: 'Enjoy PornHub',
//    id: 'pornhub',
//    headIcon: 'el-icon-s-promotion',
//    subMenuList: [
//      {
//        name: 'Asia',
//        id: 'asia',
//        route: 'cn.pornhub.com'
//      },
//      {
//        name: 'Japan',
//        id: 'japan',
//        route: 'japan.pornhub.com'
//      }
//    ]
//  }
]

export default {
  name: 'HeadNavBar',
  data: function () {
    return {
      menuList: menuList
    }
  },
  computed: {
    collapse: {
      get() {
        return this.$store.state.collapse
      }
    }
  }
}
</script>
