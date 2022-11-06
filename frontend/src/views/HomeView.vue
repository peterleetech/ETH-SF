<template>
  <div class="home">
    <el-row v-if="$store.state.apiToken === ''">
      <el-col :span="24">
        <el-empty description="no content">
          <el-button type="primary" @click="login">Login</el-button>
        </el-empty>
      </el-col>
    </el-row>
    <el-row v-else>
      <el-col :span="24">
        <el-result icon="success" title="Login Success ✌️" subTitle="you can do actions as follows">
          <template slot="extra">
            <el-button type="primary" size="medium" @click="jumpNavTo('deposit')">Deposit</el-button>
            <el-button type="primary" size="medium" @click="jumpNavTo('get-vc')">Get VC</el-button>
            <el-button type="primary" size="medium" @click="jumpNavTo('vp-withdraw')">Withdraw</el-button>
          </template>
        </el-result>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'HomeView',
  data() {
    return {
    }
  },
  methods: {
    jumpNavTo(path) {
      this.$router.push('/'+path);
    },
    async login() {
      if (this.$store.state.walletAddress === "") {
        this.$message.error("please connect wallet first");
        return
      }

      let sign = ""
      await this.$store.state.web3Client.eth.personal.sign("sign in", this.$store.state.walletAddress, null).then((res) => {
        sign = res
      })

      // login
      this.$axios.post('/user/login', {
        eth_address: this.$store.state.walletAddress,
        sign_data: sign,
      }).then(res => {
        const resObj = res.data;
        if (resObj.success) {
          this.$message({
            message: 'login success',
            type: 'success'
          });

          this.$store.commit("commitApiToken", sign)
        } else {
          this.$message.error("login failed. "+resObj.msg);
        }
      }).catch(error => {
        if (error.response) {
          this.$message.error("login failed. "+error.response.data.err_msg);
        } else {
          this.$message.error("login failed. "+error.message);
        }
      });
    },
  },
}
</script>
