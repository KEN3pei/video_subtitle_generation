package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/config"

	"extract-audio/infra"
)

func main() {
	ctx := context.Background()

	// setup context(1.IAMロールの認証情報取得)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	// download mov file
	s3Client := infra.NewBucketBasics(cfg)
	fmt.Println("download mov file")
	err = s3Client.DownloadFile(ctx, os.Getenv("AWS_BUCKET"), os.Getenv("DL_FILE_PATH"), os.Getenv("LO_FILE_PATH"))
	if err != nil {
		fmt.Println("Couldn't download movie file from s3client.")
		fmt.Println(err)
		return
	}
	fmt.Println("complete")

	// ffmpeg command
	fmt.Println("exec ffmpeg command")
	execFFmpegCommand(os.Getenv("LO_FILE_PATH"))
	fmt.Println("complete")

	// upload mp3 file
	fmt.Println("upload mp3 file")
	err = s3Client.UploadFile(ctx, os.Getenv("AWS_BUCKET"), os.Getenv("UP_MP3_FILE_PATH"), os.Getenv("LO_MP3_FILE_PATH"))
	if err != nil {
		fmt.Println("Couldn't upload movie file from s3client.")
		fmt.Println(err)
		return
	}
	fmt.Println("complete")
}

func execFFmpegCommand(filename string) {
	ls, err := exec.Command("ls", "-l", "/tmp").Output()
	fmt.Printf("ls:\n%s :Error:%v\n", ls, err)

	// 音声抽出コマンド
	// "ffmpeg -i {filename} -q:a 0 -vn /tmp/output.mp3"
	extract, err := exec.Command("ffmpeg", "-i", filename, "-q:a", "0", "-vn", os.Getenv("LO_MP3_FILE_PATH")).Output()
	if err != nil {
		fmt.Println("ffmpeg extract command execute error.")
		fmt.Println(err)
		return
	}
	fmt.Printf("ffmpeg extract:\n%s :Error:\n%v\n", extract, err)
}
