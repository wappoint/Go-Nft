// Code generated by goctl. DO NOT EDIT.
// Source: nft.proto

package nftclient

import (
	"context"

	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccountMessageVO                   = nft.AccountMessageVO
	ActivityDetailsVO                  = nft.ActivityDetailsVO
	ActivityPageVO                     = nft.ActivityPageVO
	ActivityPageVOList                 = nft.ActivityPageVOList
	BuyFromPoolBO                      = nft.BuyFromPoolBO
	BuyFromPoolRequest                 = nft.BuyFromPoolRequest
	CollectionMessageOnChainVO         = nft.CollectionMessageOnChainVO
	CommonResult                       = nft.CommonResult
	CreateActivityBO                   = nft.CreateActivityBO
	CreateActivityRequest              = nft.CreateActivityRequest
	CreatePoolBO                       = nft.CreatePoolBO
	CreatePoolRequest                  = nft.CreatePoolRequest
	DcDetailVO                         = nft.DcDetailVO
	DcOverviewVO                       = nft.DcOverviewVO
	DcPageVO                           = nft.DcPageVO
	DcPageVOList                       = nft.DcPageVOList
	Empty                              = nft.Empty
	GetDcByIdRequest                   = nft.GetDcByIdRequest
	GetDcFromActivityBO                = nft.GetDcFromActivityBO
	GetDcFromActivityRequest           = nft.GetDcFromActivityRequest
	GetDigitalCollectionHistoryRequest = nft.GetDigitalCollectionHistoryRequest
	GetMessageByHashDTO                = nft.GetMessageByHashDTO
	GetMessageByHashRequest            = nft.GetMessageByHashRequest
	GetOneActivityRequest              = nft.GetOneActivityRequest
	GetPoolByIdRequest                 = nft.GetPoolByIdRequest
	GiveDcBO                           = nft.GiveDcBO
	GiveDcRequest                      = nft.GiveDcRequest
	PoolDetailsVO                      = nft.PoolDetailsVO
	PoolPageVO                         = nft.PoolPageVO
	PoolPageVOList                     = nft.PoolPageVOList
	SearchActivitiesRequest            = nft.SearchActivitiesRequest
	SearchActivityBO                   = nft.SearchActivityBO
	SelectDcRequest                    = nft.SelectDcRequest
	SelectPoolBO                       = nft.SelectPoolBO
	SelectPoolRequest                  = nft.SelectPoolRequest
	TraceBackVO                        = nft.TraceBackVO

	Nft interface {
		GetMessageByHash(ctx context.Context, in *GetMessageByHashRequest, opts ...grpc.CallOption) (*GetMessageByHashDTO, error)
		CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*Empty, error)
		GetDcFromActivity(ctx context.Context, in *GetDcFromActivityRequest, opts ...grpc.CallOption) (*Empty, error)
		GetAllActivity(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ActivityPageVOList, error)
		GetOneActivity(ctx context.Context, in *GetOneActivityRequest, opts ...grpc.CallOption) (*ActivityDetailsVO, error)
		SearchActivities(ctx context.Context, in *SearchActivitiesRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error)
		GiveDc(ctx context.Context, in *GiveDcRequest, opts ...grpc.CallOption) (*Empty, error)
		GetAllDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error)
		SelectDc(ctx context.Context, in *SelectDcRequest, opts ...grpc.CallOption) (*DcPageVOList, error)
		GetDcById(ctx context.Context, in *GetDcByIdRequest, opts ...grpc.CallOption) (*DcDetailVO, error)
		GetMyDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error)
		GetDigitalCollectionHistory(ctx context.Context, in *GetDigitalCollectionHistoryRequest, opts ...grpc.CallOption) (*CollectionMessageOnChainVO, error)
		CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Empty, error)
		BuyFromPool(ctx context.Context, in *BuyFromPoolRequest, opts ...grpc.CallOption) (*Empty, error)
		SelectPool(ctx context.Context, in *SelectPoolRequest, opts ...grpc.CallOption) (*PoolPageVOList, error)
		GetPoolById(ctx context.Context, in *GetPoolByIdRequest, opts ...grpc.CallOption) (*PoolDetailsVO, error)
		GetAllPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error)
		GetMyPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error)
	}

	defaultNft struct {
		cli zrpc.Client
	}
)

func NewNft(cli zrpc.Client) Nft {
	return &defaultNft{
		cli: cli,
	}
}

func (m *defaultNft) GetMessageByHash(ctx context.Context, in *GetMessageByHashRequest, opts ...grpc.CallOption) (*GetMessageByHashDTO, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetMessageByHash(ctx, in, opts...)
}

func (m *defaultNft) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.CreateActivity(ctx, in, opts...)
}

func (m *defaultNft) GetDcFromActivity(ctx context.Context, in *GetDcFromActivityRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetDcFromActivity(ctx, in, opts...)
}

func (m *defaultNft) GetAllActivity(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetAllActivity(ctx, in, opts...)
}

func (m *defaultNft) GetOneActivity(ctx context.Context, in *GetOneActivityRequest, opts ...grpc.CallOption) (*ActivityDetailsVO, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetOneActivity(ctx, in, opts...)
}

func (m *defaultNft) SearchActivities(ctx context.Context, in *SearchActivitiesRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.SearchActivities(ctx, in, opts...)
}

func (m *defaultNft) GiveDc(ctx context.Context, in *GiveDcRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GiveDc(ctx, in, opts...)
}

func (m *defaultNft) GetAllDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetAllDc(ctx, in, opts...)
}

func (m *defaultNft) SelectDc(ctx context.Context, in *SelectDcRequest, opts ...grpc.CallOption) (*DcPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.SelectDc(ctx, in, opts...)
}

func (m *defaultNft) GetDcById(ctx context.Context, in *GetDcByIdRequest, opts ...grpc.CallOption) (*DcDetailVO, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetDcById(ctx, in, opts...)
}

func (m *defaultNft) GetMyDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetMyDc(ctx, in, opts...)
}

func (m *defaultNft) GetDigitalCollectionHistory(ctx context.Context, in *GetDigitalCollectionHistoryRequest, opts ...grpc.CallOption) (*CollectionMessageOnChainVO, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetDigitalCollectionHistory(ctx, in, opts...)
}

func (m *defaultNft) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.CreatePool(ctx, in, opts...)
}

func (m *defaultNft) BuyFromPool(ctx context.Context, in *BuyFromPoolRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.BuyFromPool(ctx, in, opts...)
}

func (m *defaultNft) SelectPool(ctx context.Context, in *SelectPoolRequest, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.SelectPool(ctx, in, opts...)
}

func (m *defaultNft) GetPoolById(ctx context.Context, in *GetPoolByIdRequest, opts ...grpc.CallOption) (*PoolDetailsVO, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetPoolById(ctx, in, opts...)
}

func (m *defaultNft) GetAllPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetAllPool(ctx, in, opts...)
}

func (m *defaultNft) GetMyPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	client := nft.NewNftClient(m.cli.Conn())
	return client.GetMyPool(ctx, in, opts...)
}
