<template>
  <div>
    <BNavbar toggleable="lg" type="dark" variant="dark" fixed="top">
      <BNavbarBrand href="/">
        <img src="https://wow.zamimg.com/images/wow/icons/medium/inv_helmet_74.jpg" alt="Light">
      </BNavbarBrand>
      <BNavbarNav class="ml-auto">
        <BButton variant="dark" @click="showLoginDialog">Login</BButton>
        <BButton variant="dark">Contact Us</BButton>
      </BNavbarNav>
      <LoginFormDialog ref="LoginFormDialog"></LoginFormDialog>
    </BNavbar>
  </div>
</template>

<script>
import { BNavbar, BNavbarBrand, BNavbarNav, BButton } from 'bootstrap-vue'
import LoginFormDialog from '@/components/LoginFormDialog.vue'
export default {
  name: 'HeadNavBar',
  components: {
    BNavbar,
    BNavbarBrand,
    BNavbarNav,
    BButton,
    LoginFormDialog
  },
  methods: {
    showLoginDialog: function () {
      this.$refs.LoginFormDialog.dialogVisible = true
    }
  }
}
</script>
