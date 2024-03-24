package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BuyFromPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBuyFromPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyFromPoolLogic {
	return &BuyFromPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BuyFromPoolLogic) BuyFromPool(in *nft.BuyFromPoolRequest) (*nft.CommonResult, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	dubbo := api.GetBlcService()
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		my := tx.PoolInfo
		//让PoolInfo指定id的数据中的left减一
		_, err = my.WithContext(l.ctx).Where(my.PoolId.Eq(in.BuyFromPoolBo.PoolId)).Update(my.Left, my.Left.Sub(1))
		if err != nil {
			return xerror.New("更新失败")
		}
		pool, err := my.WithContext(l.ctx).Where(my.PoolId.Eq(in.BuyFromPoolBo.PoolId)).First()
		if err != nil {
			return xerror.New("查询失败")
		}
		mint, err := dubbo.BeforeMint(l.ctx, &blc.BeforeMintRequest{
			Id: pool.PoolId,
		})
		if err != nil {
			return xerror.New("调用dubbo失败")
		}
		_, err = dubbo.Mint(l.ctx, &blc.MintRequest{
			UserKey: &blc.UserKey{UserKey: info.PrivateKey},
			PoolId:  pool.PoolId,
		})
		if err != nil {
			return xerror.New("调用dubbo失败")
		}
		dcInfo := model.DcInfo{
			Id:             int32(mint.DcId),
			Hash:           util.ByteArray2HexString(mint.UniqueId),
			Cid:            pool.Cid,
			Name:           pool.Name,
			Description:    pool.Description,
			Price:          pool.Price,
			OwnerName:      info.UserName,
			OwnerAddress:   info.Address,
			CreatorName:    pool.CreatorName,
			CreatorAddress: pool.CreatorAddress,
		}
		err = tx.DcInfo.WithContext(l.ctx).Create(&dcInfo)
		if err != nil {
			return xerror.New("创建失败")
		}
		redis := db.GetRedis()
		info.Balance = info.Balance - pool.Price
		_, err = redis.Set(l.ctx, string(info.UserId), info, 0).Result()
		if err != nil {
			return xerror.New("更新失败")
		}
		return nil
	})
	if err != nil {
		return nil, xerror.New("购买失败" + err.Error())
	}

	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
