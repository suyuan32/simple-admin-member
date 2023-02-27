package memberrank

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankByIdLogic {
	return &GetMemberRankByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankByIdLogic) GetMemberRankById(in *mms.IDReq) (*mms.MemberRankInfo, error) {
	result, err := l.svcCtx.DB.MemberRank.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(err, in)
	}

	return &mms.MemberRankInfo{
		Id:          result.ID,
		CreatedAt:   result.CreatedAt.UnixMilli(),
		UpdatedAt:   result.UpdatedAt.UnixMilli(),
		Name:        result.Name,
		Code:        result.Code,
		Description: result.Description,
		Remark:      result.Remark,
	}, nil
}
