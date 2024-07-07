package logic

import (
	"beyond/application/article/api/code"
	"context"
	"fmt"
	"net/http"
	"time"

	"beyond/application/article/api/internal/svc"
	"beyond/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10MB

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(r *http.Request) (resp *types.UploadCoverResponse, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("cover")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fmt.Println(handler.Filename)

	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.ALiYunOSS.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed, err: %v", err)
		return nil, code.GetBucketErr
	}
	objectKey := genFilename(handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, code.PutBucketErr
	}

	url, err := bucket.SignURL(objectKey, http.MethodGet, 100000000)
	if err != nil {
		logx.Errorf("sign url failed, err: %v", err)
		return nil, err
	}
	return &types.UploadCoverResponse{CoverUrl: url}, nil
}
func genFilename(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), filename)
}
