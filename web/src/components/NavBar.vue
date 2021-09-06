<template>
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
</template>

<script>
export default {
  name: 'HeadNavBar'
}
</script>
