FROM golang:1.23-alpine

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.50.0

RUN go install github.com/volatiletech/sqlboiler/v4@4.17.1
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@4.17.1

# コード全部をコピー
COPY . .
RUN go mod download

# ビルド
# CGO_ENABLED=0はCGO（C言語とのインターフェースを提供するGoの機能）を無効にしている。
# これによりGoプログラムは純粋なGoコードとしてビルドされ、Cライブラリに依存しないバイナリが生成される.軽量なコンテナイメージを作成できる.
# -oフラグは出力ファイルのパスを指定し、この場合、ビルドされたバイナリは./tmp/appというパスに保存される.
# ./cmd/appはビルド対象のGoプログラムのエントリーポイントを指定. つまり./cmd/app/main.goがビルド対象となる.
RUN CGO_ENABLED=0 go build -o ./tmp/app ./cmd/app

# wbではマルチステージビルドを利用してalpineイメージでビルドした後、distrolessイメージにコードをコピーすることで
# アプリケーションとそのランタイムの依存関係だけにしてサイズを小さくしている
# doc: https://github.com/GoogleContainerTools/distroless
# FROM gcr.io/distroless/static-debian12
# COPY --from=builder /go/src/app/ /

EXPOSE 8080

# 実行コマンド
CMD ["/tmp/app"]
