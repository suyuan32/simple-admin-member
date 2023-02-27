package memberrank

import (
	"context"

	"github.com/suyuan32/simple-admin-member-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-member-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberRankLogic {
	return &CreateMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMemberRankLogic) CreateMemberRank(in *mms.MemberRankInfo) (*mms.BaseIDResp, error) {
	result, err := l.svcCtx.DB.MemberRank.Create().
		SetName(in.Name).
		SetCode(in.Code).
		SetDescription(in.Description).
		SetRemark(in.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(err, in)
	}

	return &mms.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
