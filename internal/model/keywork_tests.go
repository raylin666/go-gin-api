package model

import (
	"github.com/raylin666/go-utils/database"
	"go-gin-api/internal/constant"
	"gorm.io/gorm"
)

type (
	KeywordTestsModel struct {
		Connection *gorm.DB
		Table      string
	}

	KeywordTests struct {
		database.Model
		Keyword    string `json:"keyword" gorm:"uniqueIndex:uni_keyword"`
		ResContent string `json:"res_content"`
	}
)

func NewKeywordTestsModel() *KeywordTestsModel {
	var connection = database.Get(constant.DefaultDatabaseConnection)
	return &KeywordTestsModel{
		Connection: connection.Conn,
		Table:      connection.Conn.Config.NamingStrategy.TableName("keyword_tests"),
	}
}

// 获取关键词数据
func (model *KeywordTestsModel) GetFirst(keyword string) *KeywordTests {
	var keyword_tests *KeywordTests
	model.Connection.Table(model.Table).Where("keyword = ?", keyword).First(&keyword_tests)
	return keyword_tests
}