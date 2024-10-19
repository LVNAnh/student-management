# Sử dụng hình ảnh Go chính thức
FROM golang:1.20

# Tạo thư mục làm việc
WORKDIR /app

# Sao chép file go.mod và go.sum trước để caching
COPY go.mod go.sum ./
RUN go mod download

# Sao chép toàn bộ mã nguồn vào thư mục làm việc
COPY . .

# Build ứng dụng
RUN go build -o main .

# Chạy ứng dụng
CMD ["/app/main"]
