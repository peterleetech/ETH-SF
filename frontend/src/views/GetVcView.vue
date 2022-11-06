<template>
  <div class="get-vc">
    <el-table
        :data="tableData"
        border
        id="vc-table"
        style="width: 100%"
        empty-text="vc not found"
        v-loading="isLoading"
    >
      <el-table-column
          prop="ID"
          label="ID"
          align="center"
          width="180">
      </el-table-column>
      <el-table-column
          prop="Asset"
          label="Asset"
          align="center">
      </el-table-column>
      <el-table-column
          prop="Amount"
          label="Amount"
          align="center">
      </el-table-column>
      <el-table-column
          prop="CreatedAt"
          label="Created At"
          align="center">
      </el-table-column>
      <el-table-column
          align="center"
          label="Action"
          width="200">
        <template slot-scope="scope">
          <el-button
              size="mini"
              type="primary"
              @click="handleVC(scope.$index, scope.row)">Get VC</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-card class="box-card" style="margin-top: 20px">
      <div slot="header" class="clearfix">
        <span><b>Download the client</b></span>
      </div>
      <div>
        <el-button type="primary" icon="el-icon-s-opportunity" @click="getClient('macos')">MacOS</el-button>
        <el-button type="primary" icon="el-icon-s-platform" @click="getClient('linux')">Linux</el-button>
      </div>
      <div style="margin-top: 20px">
        <p>download the client and vc file, run command below to create the vp.</p>
        <p>./mvp-client-{ amd64 | macos } vp create --vc_path="./vc.json" --key_path="./key" --out_path="./vp.json"</p>
      </div>
    </el-card>
    <el-dialog
        title="VC Detail"
        :visible.sync="dialogVisible"
        width="30%">
      <span id="span-vp">{{ nowVC }}</span>
      <p></p>
      <span>please download the VC file or copy it</span>
      <span slot="footer" class="dialog-footer">
        <el-button type="success" @click="downloadVC">Download</el-button>
        <el-button
            type="primary"
            id="btn-copy"
            data-clipboard-target="#span-vp"
        >Copy to clipboard</el-button>
        <el-button @click="dialogVisible = false">Close</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import ClipboardJS from 'clipboard';
import Download from 'downloadjs';

export default {
  data() {
    return {
      nowVC: '',
      dialogVisible: false,
      isLoading: false,
      tableData: [],
      clipboard: null,
    }
  },
  mounted() {
    // get vc
    this.getVc();

    this.clipboard = new ClipboardJS('#btn-copy');
    this.clipboard.on('success', e => {
      this.$message({
        message: 'copy success',
        type: 'success'
      });
      e.clearSelection();
    });
  },
  destroyed() {
    this.clipboard.destroy();
  },
  methods: {
    getVc() {
      this.isLoading = true;
      this.$axios.get('/user/vc',{
        headers: {
          "X-Token": this.$store.state.apiToken,
        }
      }).then(res => {
        const resObj = res.data;

        if (resObj.success) {
          for (const index in resObj.msg) {
            const item = resObj.msg[index];
            const vcObj = JSON.parse(item.VC);
            item.Asset = this.getAssetName(vcObj.credentialSubject.tokenAddress);
            item.Amount = vcObj.credentialSubject.amount / 1e18;
            this.tableData.push(item)
          }
        } else {
          this.$message.error("get VC failed. "+resObj.msg);
        }
        this.isLoading = false;
      }).catch(error => {
        if (error.response) {
          this.$message.error("get VC failed. "+error.response.data.err_msg);
        } else {
          this.$message.error("get VC failed. "+error.message);
        }
        this.isLoading = false;
      });
    },
    handleVC(index, row) {
      this.nowVC = row.VC;
      this.dialogVisible = true;
    },
    downloadVC() {
      Download(this.nowVC, "vc.json", "application/json");
    },
    getAssetName(address) {
      switch (address) {
        case "0xc4860463c59d59a9afac9fde35dff9da363e8425":
          return "BUSD";
        case "0x2c6e26b2fad89bc52d043e78e3d980a08af0ce88":
          return "LINK";
        case "0x6c4387c4f570aa8cadcaffc5e73ecb3d0f8fc593":
          return "WBTC";
        default:
          return "Unknown";
      }
    },
    getClient(sys) {
      let url = "";
      switch (sys) {
        case "linux":
          url = "https://veric-mvp.s3.us-west-2.amazonaws.com/public/client/mvp-client-amd64";
          break;
        case "macos":
          url = "https://veric-mvp.s3.us-west-2.amazonaws.com/public/client/mvp-client-macos";
          break;
      }

      window.open(url);
    },
  }
}
</script>
