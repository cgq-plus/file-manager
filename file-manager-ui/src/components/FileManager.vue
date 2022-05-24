<template>
  <div class="app-container">
    <el-row>
      <el-col :span="2">
        <el-button type="info" :disabled="queryParams.path === '/'" plain icon="el-icon-arrow-left"
                   @click.native="handleBack">返回上一级
        </el-button>
      </el-col>
      <el-col :span="20">
        <el-input
            type="text"
            readonly
            v-model="queryParams.path">
        </el-input>
      </el-col>
      <el-col :span="2">
        <el-button type="success" icon="el-icon-refresh"
                   @click.native="handleQuery">刷新
        </el-button>
      </el-col>
    </el-row>
    <el-divider></el-divider>
    <el-table ref="table"
              stripe
              border
              v-loading="loading"
              :data="fileList"
              style="width: 100%">
      <el-table-column
          prop="name"
          sortable
          align="left"
          label="文件名称">
        <template slot-scope="scope">
          <el-link v-if="scope.row.isDir === true" icon="el-icon-folder" :underline="false"
                   @click.native="handleClickFile(scope.row)">{{ scope.row.name }}
          </el-link>
          <el-link v-if="scope.row.isDir !== true" icon="el-icon-tickets" :underline="false">{{ scope.row.name }}
          </el-link>
        </template>
      </el-table-column>
      <el-table-column
          prop="isDir"
          sortable
          align="center"
          width="120"
          label="文件类型">
        <template slot-scope="scope">
          <el-tag type="success" v-if="scope.row.isDir === true">目录</el-tag>
          <el-tag type="info" v-if="scope.row.isDir === false">文件</el-tag>
        </template>
      </el-table-column>
      <el-table-column
          prop="modTime"
          sortable
          align="center"
          label="修改时间">
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button v-if="scope.row.isDir === true"
                     size="small"
                     type="primary"
                     plain
                     icon="el-icon-folder"
                     @click="handleClickFile(scope.row)">打开
          </el-button>
          <el-button v-if="scope.row.isDir !== true"
                     size="small"
                     type="success"
                     plain
                     icon="el-icon-download"
                     @click="handleDown(scope.row)">下载
          </el-button>
          <el-button
              size="small"
              type="danger"
              plain
              icon="el-icon-delete"
              @click="handleDel(scope.row)">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {listFile, downFile, DeleteFile} from "@/api/file"
import {Message} from 'element-ui'

export default {
  name: "FileManager",
  data() {
    return {
      queryParams: {
        path: "/",
        name: undefined,
        prePath: undefined,
      },
      loading:false,
      fileList: []
    }
  },
  created() {
    this.handleQuery()
  },
  methods: {
    handleQuery() {
      this.loading=true
      listFile(this.queryParams).then((res) => {
        if (res.list) {
          this.fileList = res.list
          this.queryParams.path = res.path
          this.queryParams.prePath = res.prePath
        } else {
          this.fileList = []
          Message.error("目录内容为空")
        }
      }).finally(() => {
        this.$refs.table.sort("modTime", "descending")
        this.loading=false
      });
    },
    handleClickFile(row) {
      if(this.queryParams.path !== '/'){
        this.queryParams.path = this.queryParams.path + "/" + row.name
      }else{
        this.queryParams.path = this.queryParams.path + row.name
      }
      this.handleQuery()
    },
    handleDown(row) {
      let params = {
        path: this.queryParams.path,
        name: row.name,
      }
      this.$confirm('此操作将下载文件【' + row.name + '】, 是否继续?', '提示', {
        confirmButtonText: '下载',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        downFile(params)
      }).catch(() => {
      });
    },
    handleDel(row) {
      let params = {
        path: this.queryParams.path,
        name: row.name,
        isDir: row.isDir,
      }
      let msg = '此操作将删除文件【' + row.name + '】, 是否继续?'
      if (row.isDir === true) {
        msg = '此操作将删除目录【' + row.name + '】, 是否继续?'
      }
      this.$confirm(msg, '警告', {
        confirmButtonText: '删除',
        cancelButtonText: '取消',
        type: 'error'
      }).then(() => {
        DeleteFile(params).then((res) => {
          Message.info(res.msg)
        }).finally(()=>{
          if (this.fileList.length > 1) {
            this.handleQuery()
          } else {
            this.handleBack()
          }
        })


      }).catch(() => {
      });
    },
    handleBack() {
      if (this.queryParams.path !== "/") {
        this.queryParams.path = this.queryParams.prePath
      }
      this.handleQuery()
    },
  }
}
</script>

<style scoped>

</style>
