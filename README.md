# go_oss_ali
上传本地文件到阿里云oss

## 运行
 1. 先配置oss_config.json
 2. 修改./oss/oss.go文件中第70行的baseDir变量值(也可默认)
 3. 在根目录下执行`go run main.go`
 4. curl访问：`curl http://127.0.0.1:9996/up?filePath=/www/backup/db/20210326.sql`

## 打包为Linux执行包
1. 项目根目录下执行：
	```
	set GOARCH=amd64
	set GOOS=linux
	go build
	```
2. 上传到服务器上
3. 后台挂起,并记录pis到main.pid文件中：`nohup ./main & echo $! > main.pid`
4. curl访问：`curl http://127.0.0.1:9996/up?filePath=/www/backup/db/20210326.sql`
