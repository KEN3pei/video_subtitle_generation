package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	setEnvs()
	// ctx := context.Background()

	// // setup context(1.IAMロールの認証情報取得)
	// cfg, err := config.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
	// 	fmt.Println(err)
	// 	return
	// }
	// // download mov file
	// s3Client := infra.NewBucketBasics(cfg)
	// fmt.Println("download mov file")
	// err = s3Client.DownloadFile(ctx, os.Getenv("BUCKET_NAME"), os.Getenv("OBJECT_KEY"), os.Getenv("LOCAL_INPUT"))
	// if err != nil {
	// 	fmt.Println("Couldn't download movie file from s3client.")
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("complete")

	// // ffmpeg command
	// fmt.Println("exec ffmpeg command")
	// execFFmpegCommand(os.Getenv("LOCAL_INPUT"))
	// fmt.Println("complete")

	// // upload mp3 file
	// fmt.Println("upload mp3 file")
	// err = s3Client.UploadFile(ctx, os.Getenv("BUCKET_NAME"), os.Getenv("S3_UPLOAD"), os.Getenv("LOCAL_OUTPUT"))
	// if err != nil {
	// 	fmt.Println("Couldn't upload movie file from s3client.")
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("complete")
}

func execFFmpegCommand(filename string) {
	ls, err := exec.Command("ls", "-l", "/tmp").Output()
	fmt.Printf("ls:\n%s :Error:%v\n", ls, err)

	// 音声抽出コマンド
	// "ffmpeg -i {filename} -q:a 0 -vn /tmp/output.mp3"
	extract, err := exec.Command("ffmpeg", "-i", filename, "-q:a", "0", "-vn", os.Getenv("LOCAL_OUTPUT")).Output()
	if err != nil {
		fmt.Println("ffmpeg extract command execute error.")
		fmt.Println(err)
		return
	}
	fmt.Printf("ffmpeg extract:\n%s :Error:\n%v\n", extract, err)
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func setEnvs() {
	filepath := getFileNameWithoutExt(os.Getenv("OBJECT_KEY"))

	os.Setenv("LOCAL_INPUT", fmt.Sprintf("/tmp/%s.mp4", filepath))
	os.Setenv("LOCAL_OUTPUT", fmt.Sprintf("/tmp/%s.mp3", filepath))
	os.Setenv("S3_UPLOAD", fmt.Sprintf("mp3/%s.mp3", filepath))

	println("LOCAL_INPUT:", os.Getenv("LOCAL_INPUT"))
	println("LOCAL_OUTPUT:", os.Getenv("LOCAL_OUTPUT"))
	println("S3_UPLOAD:", os.Getenv("S3_UPLOAD"))
}
