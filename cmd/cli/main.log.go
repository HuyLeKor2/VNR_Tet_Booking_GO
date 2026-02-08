package main

import (
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//--- PHẦN 1: SugaredLogger vs Logger (Hình 2) ---

	//Sử dụng SugaredLogger: Tiện lợi, hỗ trợ định dạng chuỗi giống printf
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name:%s, age:%d ", "TipsGo", 40) // Giống fmt.Printf

	// Sử dụng Logger tiêu chuẩn: Hiệu năng cực cao, chặt chẽ về kiểu dữ liệu (Strongly typed)
	logger := zap.NewExample()
	logger.Info("Hello",
		zap.String("name", "TipsGo"),
		zap.Int("age", 40),
	)

	// --- PHẦN 2: Các chế độ khởi tạo Logger (Hình 3) ---

	// 1. NewExample: Dùng cho ví dụ hoặc testing (Output dạng JSON tối giản)
	loggerEx := zap.NewExample()
	loggerEx.Info("Hello NewExample")

	// 2. NewDevelopment: Dùng khi phát triển (Output dạng text dễ đọc, có màu, hiển thị file/dòng code)
	loggerDev, _ := zap.NewDevelopment()
	loggerDev.Info("Hello NewDevelopment")

	// 3. NewProduction: Dùng cho thực tế (Output dạng JSON đầy đủ, tối ưu cho các hệ thống thu thập log)
	loggerProd, _ := zap.NewProduction()
	loggerProd.Info("Hello NewProduction")

	// // sugar: Dễ sử dụng, hỗ trợ định dạng như fmt.Printf
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d ", "TipsGo", 40) // like fmt.Printf(format, a)

	// // logger: Hiệu năng cao, yêu cầu định nghĩa kiểu dữ liệu chặt chẽ
	// logger := zap.NewExample()
	// logger.Info("Hello",
	// 	zap.String("name", "TipsGo"),
	// 	zap.Int("age", 40),
	// )
}

func InitLogger() {
	encoder := getEncoderLog()
	sync := getWriterSync()

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// Hàm định dạng log (Encoder)
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Định dạng thời gian: 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	// Viết hoa level log: INFO, ERROR
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Hiển thị file và dòng code gọi log: cli/main.log.go:24
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// Hàm điều hướng nơi lưu log (WriteSyncer)
func getWriterSync() zapcore.WriteSyncer {
	// Mở file để lưu log
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)

	// Xuất log ra console
	syncConsole := zapcore.AddSync(os.Stderr)

	// Kết hợp cả hai: vừa ghi file vừa in ra màn hình
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
