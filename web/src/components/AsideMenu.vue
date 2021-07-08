<template>
  <div>
    <el-scrollbar class='page-component__scroll'>
      <!--
      <transition :duration="{ enter: 800, leave: 100}" mode="out-in">
        <el-menu
          :collapse="false"
          :collapse-transition="true"
          :default-active="active"
          :default-openeds="['0', '1', '2', '3', '4', '5']"
          class="el-menu-vertical"
          unique-opened
          router
          background-color="#545c64"
          active-text-color="#ffd04b"
          text-color="#fff">
          <el-submenu v-for="(singleMenu, idx) in menuList" :index="idx.toString()" :key="singleMenu.id">
            <template slot="title">
              <i :class="singleMenu.headIcon"></i>
              {{ singleMenu.name }}
            </template>
            <el-menu-item-group>
              <el-menu-item v-for="item in singleMenu.subMenuList" :index="item.route" :key="item.id">
              {{ item.name }}
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
        </el-menu>
      </transition>
      -->
      <li class="nav-item" v-for="singleMenu in menuList" :key="singleMenu.id">
        {{ singleMenu.name }}
        <ul class="pure-menu-list sub-nav">
          <li v-for="item in singleMenu.subMenuList" :key="item.id" class="nav-item">
            <a href="item.route">{{ item.name }}</a>
          </li>
        </ul>
      </li>
    </el-scrollbar>
  </div>
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
        route: '/dailyOverView'
      },
      {
        name: 'Query Stock',
        id: 'query-stock',
        route: '/queryStockView'
      }
    ]
  },
  {
    name: 'Algo Related',
    id: 'algo',
    headIcon: 'el-icon-cpu',
    subMenuList: [
      {
        name: 'Inference',
        id: 'inference',
        route: '#'
      },
      {
        name: 'Regression',
        id: 'regression',
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
  },
  {
    name: 'Report',
    id: 'report',
    headIcon: 'el-icon-printer',
    subMenuList: [
      {
        name: 'Index OverView',
        id: 'index-overview',
        route: '#'
      },
      {
        name: 'Other Format',
        id: 'other-format',
        route: '#'
      }
    ]
  },
  {
    name: 'Settings',
    id: 'settings',
    headIcon: 'el-icon-setting',
    subMenuList: [
      {
        name: 'Server Settings',
        id: 'server-settings',
        route: '#'
      },
      {
        name: 'Other Settings',
        id: 'other-settings',
        route: '#'
      }
    ]
  },
  {
    name: 'Enjoy PornHub',
    id: 'pornhub',
    headIcon: 'el-icon-s-promotion',
    subMenuList: [
      {
        name: 'Asia',
        id: 'asia',
        route: 'cn.pornhub.com'
      },
      {
        name: 'Japan',
        id: 'japan',
        route: 'japan.pornhub.com'
      }
    ]
  }
]

export default {
  name: 'AsideMenu',
  data: function () {
    return {
      menuList: menuList
    }
  }
}
</script>
