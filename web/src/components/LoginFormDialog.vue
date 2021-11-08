<template>
  <div>
    <el-dialog
      text="Login"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'LoginFormDialog',
  components: {
  },
  data: function () {
    return {
      loginInfo: {
        account: '',
        password: ''
      },
      show: true,
      dialogVisible: false
    }
  },
  methods: {
    onSubmit: function () {
      alert(JSON.stringify(this.loginInfo))
    },
    onReset: function () {
      this.loginInfo.account = ''
      this.loginInfo.password = ''
      this.show = false
      this.$nextTick(() => {
        this.show = true
      })
    },
    handleClose: function (done) {
      this.$confirm('Sure to Close ?')
        .then(_ => {
          done()
        })
    }
  }
}
</script>
