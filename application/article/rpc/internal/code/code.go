package code

import (
	"beyond/pkg/codex"
)

var (
	SortTypeInvalid         = codex.New(60001, "排序类型无效")   // 排序类型无效
	UserIdInvalid           = codex.New(60002, "用户ID无效")   // 用户ID无效
	ArticleTitleCantEmpty   = codex.New(60003, "文章标题不能为空") // 文章标题不能为空
	ArticleContentCantEmpty = codex.New(60004, "文章内容不能为空") // 文章内容不能为空
	ArticleIdInvalid        = codex.New(60005, "文章ID无效")   // 文章ID无效
)
