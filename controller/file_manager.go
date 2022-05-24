package controller

import (
	"file-manager/domain"
	"file-manager/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	filepath := fmt.Sprintf("%s/%s", realPath, req.Name)
	if req.Path == "/" {
		filepath = fmt.Sprintf("%s%s", realPath, req.Name)
	}
	if req.IsDir {
		fmt.Printf("删除目录:%s\n", filepath)
		err := os.RemoveAll(filepath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else {
		fmt.Printf("删除文件:%s\n", filepath)
		err := os.Remove(filepath)
		if err != nil {
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
