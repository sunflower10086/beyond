package code

import "beyond/pkg/codex"

var (
	GetBucketErr              = codex.New(30001, "获取bucket实例失败")
	PutBucketErr              = codex.New(30002, "上传bucket失败")
	ArticleTitleEmpty         = codex.New(30004, "文章标题为空")
	ArticleContentTooFewWords = codex.New(30005, "文章内容字数太少")
	ArticleCoverEmpty         = codex.New(30006, "文章封面为空")
)
