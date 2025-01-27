### 参照記事 & Memo

https://dev.classmethod.jp/articles/eventbridge-booting-with-s3-file-uploading/

https://zenn.dev/y16ra/articles/31e32e61db8a44
- 環境変数にEventBridgeから渡された値をセットしてコンテナを起動する

### dir

```
├── extract-audio
│   ├── Readme.md
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── .env
│   ├── build
│   │   └── Dockerfile
│   └── infra
│       └── s3repository.go
├── tmp
├── docker-compose.yaml
```

- https://github.com/golang-standards/project-layout/blob/master/README_ja.md
- https://zenn.dev/foxtail88/articles/824c5e8e0c6d82
- https://www.slideshare.net/slideshow/go-80591000/80591000

### docker

$ docker build -f build/Dockerfile -t extract-audio .

$ docker run --rm --name extract-audio -it extract-audio /bin/sh

$ docker run -v ./tmp:/tmp --name extract-audio -it --rm extract-audio /bin/sh
/app # ffmpeg -i /tmp/kokumin_gaito_enzetu.mp4 -q:a 0 -vn /tmp/output.mp3

