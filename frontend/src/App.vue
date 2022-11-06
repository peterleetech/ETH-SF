<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-menu
            :default-active="$route.path"
            mode="horizontal"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
            router
        >
          <el-menu-item>
            <el-image
                style="width: 100px;padding-top: 9px;"
                :src="require('./assets/logo.png')"
                fit="fill"
            ></el-image>
          </el-menu-item>
          <el-menu-item index="/">Home</el-menu-item>
          <template v-if="$store.state.apiToken !== ''">
            <el-menu-item index="/deposit">Deposit</el-menu-item>
            <el-menu-item index="/get-vc">Get VC</el-menu-item>
            <el-menu-item index="/vp-withdraw">Withdraw</el-menu-item>
          </template>
          <el-menu-item v-if="$store.state.walletAddress === ''">
            <el-button type="warning" @click="connectWallet">Connect Wallet</el-button>
          </el-menu-item>
          <template v-else>
            <el-menu-item>
              <span>connected：{{ $store.state.walletAddress }}</span>
            </el-menu-item>
            <el-menu-item>
              <el-button type="danger" @click="disConnectWallet">Disconnect</el-button>
            </el-menu-item>
          </template>
        </el-menu>
      </el-header>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Web3 from "web3";
import Web3Modal from "web3modal";
import WalletConnectProvider from "@walletconnect/web3-provider";

export default {
  data() {
    return {
      web3Modal: null,
      web3ModalListened: false,
      providerOptions: {
        walletconnect: {
          package: WalletConnectProvider,
          options: {
            infuraId: "f46ac8e5497b4580aa43939ff4db19d3"
          }
        }
      },
    };
  },
  mounted() {
    this.web3Modal = new Web3Modal({
      network: "mainnet",
      cacheProvider: false,
      providerOptions: this.providerOptions
    });
  },
  methods: {
    async connectWallet() {
      let provider;
      try {
        provider = await this.web3Modal.connect();
      } catch (e) {
        this.$message.error("user refuse connect the wallet");
        return
      }

      if (provider.chainId !== "0x6357d2e0") {
        this.$message.error("error network. please use testnet of Harmony");
        return
      }

      console.log(provider)

      let web3Client = new Web3(provider);
      this.$store.commit("commitWeb3Client", web3Client);

      // get wallet address
      await this.$store.state.web3Client.eth.getAccounts().then((accounts) => {
        this.$store.commit("commitWalletAddress", accounts[0]);
      });

      if (!this.web3ModalListened) {
        provider.on("accountsChanged", async () => {
          // get wallet address
          await this.$store.state.web3Client.eth.getAccounts().then((accounts) => {
            this.$store.commit("commitWalletAddress", accounts[0]);

            if (accounts.length > 0) {
              // change wallet
              this.$message({
                message: 'change wallet success',
                showClose: true,
                type: 'success'
              });
            } else {
              // disconnect wallet
              this.$store.commit("commitWalletAddress", "");
              this.$store.commit("commitApiToken", "");

              this.web3Modal.clearCachedProvider();
              this.$message({
                message: 'disconnect wallet success',
                showClose: true,
                type: 'success'
              });
            }
          });
        });

        this.web3ModalListened = true
      }
    },
    disConnectWallet() {
      this.$store.commit("commitWalletAddress", "");
      this.$store.commit("commitApiToken", "");

      this.web3Modal.clearCachedProvider();

      this.$message({
        message: 'disconnect wallet success',
        showClose: true,
        type: 'success'
      });

      this.$router.push('/')
    },
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
