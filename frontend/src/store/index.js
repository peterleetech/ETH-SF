import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    apiToken: "",
    walletAddress: "",
    web3Client: null,
    erc20ContractAddress: {
      busd: "0xc4860463c59d59a9afac9fde35dff9da363e8425",
      link: "0x2c6e26b2fad89bc52d043e78e3d980a08af0ce88",
    },
    vaultContractAddress: "0xf9cd945f27ea67b5871b97446eb3c243ff866a89",
  },
  getters: {
  },
  mutations: {
    commitApiToken(state, val){
      state.apiToken = val;
    },
    commitWalletAddress(state, val){
      state.walletAddress = val;
    },
    commitWeb3Client(state, val){
      state.web3Client = val;
    },
  },
  actions: {

  },
  modules: {
  }
})
