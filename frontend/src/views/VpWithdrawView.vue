<template>
  <div class="vp-withdraw">
    <el-form label-position="top" label-width="80px" :model="formData">
      <el-form-item label="please copy your VP content into the textarea below">
        <el-input
            type="textarea"
            :rows="6"
            placeholder="VP content"
            v-model="formData.vpData">
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitVp" :disabled="formData.vpData === ''" :loading="txPending">Submit</el-button>
        <el-button @click="resetVp">Reset</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      txPending: false,
      formData: {
        vpData: '',
      }
    }
  },
  methods: {
    submitVp() {
      this.txPending = true

      this.$axios.post('/user/vp/withdraw', {
        vp: this.formData.vpData,
      },{
        headers: {
          "X-Token": this.$store.state.apiToken,
        }
      }).then(res => {
        const resObj = res.data;
        if (resObj.success) {
          this.$message({
            message: 'submit success',
            type: 'success'
          });
        } else {
          this.$message.error("submit VP failed. "+resObj.msg);
        }
        this.txPending = false
      }).catch(error => {
        if (error.response) {
          this.$message.error("submit VP failed. "+error.response.data.err_msg);
        } else {
          this.$message.error("submit VP failed. "+error.message);
        }
        this.txPending = false
      });
    },
    resetVp() {
      this.formData.vpData = "";
    }
  }
}
</script>
