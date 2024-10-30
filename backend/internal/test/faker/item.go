package faker

import (
	"fmt"
	"math/rand"
)

func GenerateTitle() string {
	categories := []string{"時計", "カメラ", "パソコン", "スマートフォン", "タブレット"}
	brands := []string{"Apple", "Sony", "Samsung", "LG", "Panasonic"}

	category := categories[rand.Intn(len(categories))]
	brand := brands[rand.Intn(len(brands))]
	year := rand.Intn(5) + 2020 // 2020-2024

	return fmt.Sprintf("%s %s %d", brand, category, year)
}

func GenerateDescription() string {
	// 商品の説明を4文生成
	conditions1 := []string{"新品", "未使用", "美品", "やや傷あり", "使用感あり"}
	conditions2 := []string{"全体的に状態が良いです", "動作確認済み", "付属品はありません", "動作未確認", "ジャンク品"}
	conditions3 := []string{"お探しの方にお譲りします", "お値引き可能です", "お取引の際はコメントをお願いします", "即購入OK", "お値段交渉可能です"}
	conditions4 := []string{"お気軽にコメントください", "お取引の際はプロフィールをご覧ください", "他にも出品しています"}

	condition1 := conditions1[rand.Intn(len(conditions1))]
	condition2 := conditions2[rand.Intn(len(conditions2))]
	condition3 := conditions3[rand.Intn(len(conditions3))]
	condition4 := conditions4[rand.Intn(len(conditions4))]

	return fmt.Sprintf(`商品の状態: %s
商品の説明:
%s
%s
%s`,
		condition1,
		condition2,
		condition3,
		condition4)
}
