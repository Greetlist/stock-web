<template>
  <el-row class="tac">
    <el-col :span="24">
      <el-menu
        :default-openeds="['1', '2', '3']"
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
    </el-col>
  </el-row>
</template>
<script>
var menuList = [
  {
    name: 'Stock OverView',
    id: 'stock',
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
