package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"Reco-server/src/init/model"

	"gorm.io/gorm"
)

// seedデータ作成関数
func CreateSeed[T any](db *gorm.DB, filePath string, checkExistsFn func(db *gorm.DB, model T) bool) error {
	// seedデータが存在する場合は何も実行しない
	var count int64

	db.Model(&model.Users{}).Count(&count)
	if count > 0 {
		return nil
	}

	// JSONファイルを読み込む
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// JSONデコード
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var items []T
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// データの存在確認
	if len(items) == 0 {
		return fmt.Errorf("no data to seed")
	}
	
	firstItem := items[0]
	if checkExistsFn(db, firstItem) {
		fmt.Println("Data already exists. Skipping seeding.")
		return nil
	}

	// データベースに保存
	for _, item := range items {
		if err := db.Create(&item).Error; err != nil {
			return fmt.Errorf("failed to seed data: %w", err)
		}
	}

	fmt.Println("Seeding completed successfully.")
	return nil
}

// データの存在確認関数
func checkUserExists[T any](db *gorm.DB, model T) bool {
    var count int64
    db.Model(&model).Count(&count)
    return count > 0
}

// main実行関数
func Seeder(db *gorm.DB) {
	// Users
	if err := CreateSeed[model.Users](db, "users.json",checkUserExists); err != nil {
		fmt.Printf("Error seeding data: %v\n", err)
	}

	// Tweets
	if err := CreateSeed[model.Tweets](db, "tweets.json",checkUserExists); err != nil {
		fmt.Printf("Error seeding data: %v\n", err)
	}

	// RecommendItem
	if err := CreateSeed[model.RecommendItem](db, "recommenditem.json",checkUserExists); err != nil {
		fmt.Printf("Error seeding data: %v\n", err)
	}

	// retweets
	if err := CreateSeed[model.Retweets](db, "retweets.json",checkUserExists); err != nil {
		fmt.Printf("Error seeding data: %v\n", err)
	}

	// follow
	if err := CreateSeed[model.Follow](db, "follow.json",checkUserExists); err != nil {
		fmt.Printf("Error seeding data: %v\n", err)
	}
}