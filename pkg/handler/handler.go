package handler

import (
	"context"
	"fmt"
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_crawler/pkg/sdk/base"
	"github.com/ecnuvj/vhoj_crawler/pkg/sdk/crawlerpb"
	"github.com/ecnuvj/vhoj_crawler/pkg/service"
	"github.com/ecnuvj/vhoj_crawler/pkg/utils/reply_utils"
)

type CrawlerHandler struct {
	crawlerpb.UnimplementedCrawlerServiceServer
	CrawlerService *service.CrawlerService
}

func NewCrawlerHandler() (*CrawlerHandler, error) {
	return &CrawlerHandler{
		CrawlerService: &service.CrawlerService{},
	}, nil
}

func (c *CrawlerHandler) CrawlProblem(ctx context.Context, request *crawlerpb.CrawlProblemRequest) (*crawlerpb.CrawlProblemResponse, error) {
	if request == nil {
		return &crawlerpb.CrawlProblemResponse{
			BaseResponse: reply_utils.PbReplyf(base.REPLY_STATUS_FAILURE, "request is nil"),
		}, fmt.Errorf("request is nil")
	}
	err := c.CrawlerService.CrawlProblem(remote_oj.RemoteOJ(request.RemoteOj), request.RemoteProblemId)
	if err != nil {
		return &crawlerpb.CrawlProblemResponse{
			BaseResponse: reply_utils.PbReplyf(base.REPLY_STATUS_FAILURE, "service error: %v", err),
		}, fmt.Errorf("service error: %v", err)
	}
	return &crawlerpb.CrawlProblemResponse{
		BaseResponse: reply_utils.PbReplyf(base.REPLY_STATUS_SUCCESS, "success"),
	}, nil
}
