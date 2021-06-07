<template>
  <div>
    <el-dialog
      text="Login"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="handleClose">
      <BForm @submit="onSubmit" @reset="onReset" v-if="show">
        <BFormGroup
          id="input-account-group"
          label="Account"
          label-for="inpout-account"
          description="Enter Your Account">
          <BFormInput
            id="input-account"
            v-model="loginInfo.account"
            required
            placeholder="Account">
          </BFormInput>
        </BFormGroup>
        <BFormGroup
          id="input-password-group"
          label="Account"
          label-for="inpout-password"
          description="Enter Your Password">
          <BFormInput
            id="input-password"
            v-model="loginInfo.password"
            required
            placeholder="Password">
          </BFormInput>
        </BFormGroup>
        <BButton type='submit' variant="primary">Login</BButton>
        <BButton type='reset' variant="danger">Reset</BButton>
      </BForm>
    </el-dialog>
  </div>
</template>

<script>
import { BForm, BFormGroup, BFormInput, BButton } from 'bootstrap-vue'
export default {
  name: 'LoginFormDialog',
  components: {
    BForm,
    BFormGroup,
    BFormInput,
    BButton
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
