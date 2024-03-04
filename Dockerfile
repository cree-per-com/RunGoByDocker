# 1단계: Go 빌드 환경 사용
FROM golang:1.18 AS builder

# 작업 디렉토리 설정
WORKDIR /app

# go.mod 파일 복사 후 의존성 다운로드
COPY go.mod ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 2단계: 실행 환경 준비
FROM alpine:latest

WORKDIR /root/

# 1단계에서 빌드한 실행 파일 복사
COPY --from=builder /app/myapp .

# 컨테이너 실행 시 애플리케이션 실행
CMD ["./myapp"]
