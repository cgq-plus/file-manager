package controller

import (
	"file-manager/domain"
	"file-manager/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// List 读取文件列表
func List(ctx *gin.Context) {
	// 定义参数接受结构体
	var req *domain.ListRequest
	// 绑定json参数
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 本地真实目录
	realPath := fmt.Sprintf("%s%s", global.CONFIG.Application.RootPath, req.Path)
	// 读取目录文件
	files, _ := ioutil.ReadDir(realPath)
	var fileList []*domain.FileItem
	for _, file := range files {
		item := &domain.FileItem{
			Name:    file.Name(),
			Size:    file.Size(),
			Mode:    file.Mode(),
			ModTime: file.ModTime().Format("2006-01-02 15:04:05"),
			IsDir:   file.IsDir(),
		}
		fileList = append(fileList, item)
	}
	// 上级目录
	prePath := "/"
	if req.Name == "" && req.Path != "/" {
		prePathArr := strings.FieldsFunc(req.Path, func(r rune) bool {
			return r == '/'
		})
		length := len(prePathArr)
		if length > 0 {
			prePath = "/"
			for i := 0; i < len(prePathArr)-1; i++ {
				prePath += prePathArr[i] + "/"
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":    fileList,
			"path":    req.Path,
			"prePath": prePath,
		},
	})
}

// Download 下载文件
func Download(ctx *gin.Context) {
	// 定义参数接受结构体
	var req *domain.ListRequest
	// 绑定json参数
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 本地真实目录
	realPath := fmt.Sprintf("%s%s", global.CONFIG.Application.RootPath, req.Path)
	filepath := fmt.Sprintf("%s/%s", realPath, req.Name)
	if req.Path == "/" {
		filepath = fmt.Sprintf("%s%s", realPath, req.Name)
	}
	fmt.Println(filepath)
	ctx.File(filepath)
	return
}

// Delete 删除文件
func Delete(ctx *gin.Context) {
	// 定义参数接受结构体
	var req *domain.DeleteRequest
	// 绑定json参数
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 本地真实目录
	realPath := fmt.Sprintf("%s%s", global.CONFIG.Application.RootPath, req.Path)
	// 备份真实目录
	backupPath := fmt.Sprintf("%s%s", global.CONFIG.Application.BackUpPath, req.Path)
	filepath := fmt.Sprintf("%s/%s", realPath, req.Name)
	if req.Path == "/" {
		filepath = fmt.Sprintf("%s%s", realPath, req.Name)
	}
	//获取文件名称带后缀
	fileNameWithSuffix := path.Base(filepath)
	//获取文件的后缀(文件类型)
	fileType := path.Ext(fileNameWithSuffix)
	//获取文件名称(不带后缀)
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	backUpFilePath := fmt.Sprintf("%s%s", backupPath, fmt.Sprintf("%s_%s%s", fileNameOnly, time.Now().Format("2006_01_02_150405"), fileType))
	if req.IsDir {
		global.LOG.Sugar().Infof("删除目录:%s \n", filepath)
		if global.CONFIG.Application.LogicDelete {
			err = os.Rename(filepath, backUpFilePath)
		} else {
			err = os.RemoveAll(filepath)
		}
		if err != nil {
			global.LOG.Sugar().Infof("删除目录异常.%s \n", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else {
		global.LOG.Sugar().Infof("删除文件:%s \n", filepath)
		if global.CONFIG.Application.LogicDelete {
			err = os.Rename(filepath, backUpFilePath)
		} else {
			err = os.Remove(filepath)
		}
		if err != nil {
			global.LOG.Sugar().Infof("删除目录异常.%s \n", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
	})
	return
}
