<template>
  <button v-on:click="onClick">Test</button>
</template>

<script>
const axios = require('axios').default
export default {
  name: 'TestButton',
  data: function () {
    return {
      all_data: 'Test'
    }
  },
  methods: {
    onClick () {
      const header = {
        'Access-Control-Allow-Origin': '*'
      }
      axios.get('http://121.5.100.186:8081/test', header)
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
        .then(function () {
          console.log('Finish Post Request')
        })
    }
  }
}
</script>
