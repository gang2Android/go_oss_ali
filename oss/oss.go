package oss

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	io "io/ioutil"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

// OssConfig ...
type OssConfig struct {
	BucketName      string `json:"bucketName"`
	EndPoint        string `json:"endPoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}

var fileLocker sync.Mutex

func loadConfig(path string) (OssConfig, bool) {
	var conf OssConfig
	fileLocker.Lock()
	data, err := io.ReadFile(path)
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, false
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	return conf, true
}

func Up(filePath string, isDelete bool) {
	config, ok := loadConfig("./oss_config.json")
	if !ok {
		return
	}
	//fmt.Println(config)
	bucketName := config.BucketName
	endPoint := config.EndPoint
	accessKeyId := config.AccessKeyId
	accessKeySecret := config.AccessKeySecret

	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		// HandleError(err)
		return
	}
	if client == nil {
		return
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		// HandleError(err)
		return
	}

	// 指定oss内的存储路径
	fileName := path.Base(filePath) + "-" + strconv.FormatInt(time.Now().Unix(), 10) + path.Ext(filePath)
	dirName := time.Now().Format("20060102")
	objectKey := "backup/db/" + dirName + "/" + fileName

	err = bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("success:" + objectKey)

	// 更改文件存储类型为归档存储类型
	_, err = bucket.CopyObject(objectKey, objectKey, oss.ObjectStorageClass(oss.StorageArchive))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	if isDelete {
		err2 := os.Remove(filePath)
		if err2 != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
	}

}
