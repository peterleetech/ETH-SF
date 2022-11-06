<template>
  <div class="deposit">
    <el-card class="box-card">
      <el-row>
        <el-col :span="4" v-if="formData.asset !== ''">
          Now Balance:
        </el-col>
        <el-col :span="4" v-else>
          Please select the asset.
        </el-col>
        <el-col :span="5" v-if="formData.asset === 'busd'" class="asset-balance">
          <el-image
              style="width: 22px; height: 22px"
              :src="require('../assets/busd.png')"
              fit="fill"
          ></el-image>
          <span>BUSD: {{balance.busd}}</span>
        </el-col>
        <el-col :span="5" v-if="formData.asset === 'link'" class="asset-balance">
          <el-image
              style="width: 22px; height: 22px"
              :src="require('../assets/link.png')"
              fit="fill"
          ></el-image>
          <span>LINK: {{balance.link}}</span>
        </el-col>
        <el-col :span="10"></el-col>
      </el-row>
    </el-card>
    <el-card class="box-card">
      <el-row>
        <el-col :span="12">
          <el-form status-icon label-position="left" label-width="150px" :model="formData" ref="depositForm">
            <el-form-item
                label="Deposit Asset"
                prop="asset"
                :rules="[
                  { required: true, message: 'asset must be selected'},
                ]"
            >
              <el-select v-model="formData.asset" placeholder="please select the asset" style="width: 100%;" @change="changeAsset">
                <el-option label="BUSD" value="busd"></el-option>
                <el-option label="LINK" value="link"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item
                label="Asset Amount"
                prop="amount"
                :rules="[
                  { required: true, message: 'amount must be set'},
                  { type: 'number', message: 'amount must be a number', trigger: 'blur'},
                  { type: 'number', min: 1, message: 'amount mush bigger than 0', trigger: 'blur' },
                ]"
            >
              <el-input placeholder="please fill the amount" v-model.number="formData.amount" @blur="checkAllowance" :disabled="formData.asset === ''">
                <template slot="append">{{ formData.asset.toUpperCase() }}</template>
              </el-input>
            </el-form-item>
            <el-form-item style="text-align: right">
              <el-button type="warning" @click="submitApprove" :loading="txPending" v-if="needApprove">Approve</el-button>
              <el-button type="primary" @click="submitDeposit" :loading="txPending" v-else>Deposit</el-button>
              <el-button @click="resetForm">Reset</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
import {Erc20Abi} from "@/contract/erc20";
import {VaultAbi} from "@/contract/vault";

export default {
  data() {
    return {
      contractClient: {},
      vaultContractClient: null,
      defaultCallOption: {
        from: this.$store.state.walletAddress,
      },
      needApprove: false,
      txPending: false,
      balance: {
        busd: 0,
        link: 0,
      },
      allowanceAmount: {
        busd: 0,
        link: 0,
      },
      formData: {
        asset: "",
        amount: 1,
      },
      utils: this.$store.state.web3Client.utils,
    }
  },
  async mounted() {
    this.contractClient["busd"] = new this.$store.state.web3Client.eth.Contract(Erc20Abi,this.$store.state.erc20ContractAddress.busd);
    this.contractClient["link"] = new this.$store.state.web3Client.eth.Contract(Erc20Abi,this.$store.state.erc20ContractAddress.link);
    this.vaultContractClient = new this.$store.state.web3Client.eth.Contract(VaultAbi,this.$store.state.vaultContractAddress);
  },
  methods: {
    getBalance(asset) {
      return new Promise((done) => {
        this.contractClient[asset].methods.balanceOf(this.$store.state.walletAddress).call(this.defaultCallOption, (error, result) => {
          if (error !== null) {
            this.$notify.error({
              title: 'Contract Call Error',
              message: error.message,
            });
            return
          }

          this.balance[asset] = parseInt(result) / 1e18;
          done()
        });
      })

    },
    getAllowanceAmount(asset) {
      return new Promise((done) => {
        this.contractClient[asset].methods.allowance(this.$store.state.walletAddress, this.$store.state.vaultContractAddress).call(this.defaultCallOption, (error, result) => {
          if (error !== null) {
            this.$notify.error({
              title: 'Contract Call Error',
              message: error.message,
            });
            return
          }

          this.allowanceAmount[asset] = parseInt(result) / 1e18;

          this.checkAllowance();
          done();
        });
      })
    },
    async changeAsset(newAsset) {
      await this.getBalance(newAsset);
      await this.getAllowanceAmount(newAsset);
    },
    checkAllowance() {
      if (this.formData.asset === "" || this.formData.amount === 0) {
        this.needApprove = false;
        return
      }

      if (this.allowanceAmount[this.formData.asset] < this.formData.amount) {
        this.needApprove = true;
      } else {
        this.needApprove = false;
      }
    },
    submitApprove(){
      this.contractClient[this.formData.asset].methods.approve(this.$store.state.vaultContractAddress, this.utils.toWei('100000000000000', 'ether')).send(this.defaultCallOption)
          .on('transactionHash', (hash) => {
            this.txPending = true;
            this.$notify.info({
              title: 'Transaction Submit',
              message: hash
            });
          })
          .on('receipt', (receipt) => {
            this.$notify({
              title: 'Transaction Confirmed',
              message: 'approve asset success',
              type: 'success'
            });

            this.checkAllowance();

            this.txPending = false;
          })
          .on('error', (error, receipt) => {
            this.$notify.error({
              title: 'Contract Call Error',
              message: error.message,
            });

            this.txPending = false;
          });
    },
    submitDeposit() {
      this.$refs["depositForm"].validate((valid) => {
        if (valid) {
          console.log(this.formData.amount, this.utils.toWei(this.formData.amount+'', 'ether'))
          this.vaultContractClient.methods.deposit(this.$store.state.erc20ContractAddress[this.formData.asset], this.utils.toWei(this.formData.amount+'', 'ether')).send(this.defaultCallOption)
              .on('transactionHash', (hash) => {
                this.txPending = true;
                this.$notify.info({
                  title: 'Transaction Submit',
                  message: hash
                });
              })
              .on('receipt', (receipt) => {
                this.$notify({
                  title: 'Transaction Confirmed',
                  message: 'approve asset success',
                  type: 'success'
                });

                this.checkAllowance();

                this.txPending = false;
              })
              .on('error', (error, receipt) => {
                this.$notify.error({
                  title: 'Contract Call Error',
                  message: error.message,
                });

                this.txPending = false;
              });
        } else {
          this.$message.error("submit failed");
          return false;
        }
      });
    },
    resetForm() {
      this.$refs["depositForm"].resetFields();
    }
  }
}
</script>

<style scoped>
  .box-card {
    margin-bottom: 20px;
  }
  .asset-balance {
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }
  .asset-balance span {
    margin-left: 5px;
  }
</style>