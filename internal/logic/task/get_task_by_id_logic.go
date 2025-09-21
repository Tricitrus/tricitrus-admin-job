package task

import (
	"context"

	"github.com/Tricitrus/tricitrus-admin-job/internal/svc"
	"github.com/Tricitrus/tricitrus-admin-job/internal/utils/dberrorhandler"
	"github.com/Tricitrus/tricitrus-admin-job/types/job"

	"github.com/Tricitrus/tricitrus-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskByIdLogic {
	return &GetTaskByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskByIdLogic) GetTaskById(in *job.IDReq) (*job.TaskInfo, error) {
	result, err := l.svcCtx.DB.Task.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.TaskInfo{
		Id:             &result.ID,
		CreatedAt:      pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:      pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:         pointy.GetPointer(uint32(result.Status)),
		Name:           &result.Name,
		TaskGroup:      &result.TaskGroup,
		CronExpression: &result.CronExpression,
		Pattern:        &result.Pattern,
		Payload:        &result.Payload,
	}, nil
}
