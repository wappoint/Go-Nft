// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: nft.proto

package nft

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Nft_GetMessageByHash_FullMethodName            = "/user.Nft/GetMessageByHash"
	Nft_CreateActivity_FullMethodName              = "/user.Nft/CreateActivity"
	Nft_GetDcFromActivity_FullMethodName           = "/user.Nft/GetDcFromActivity"
	Nft_GetAllActivity_FullMethodName              = "/user.Nft/GetAllActivity"
	Nft_GetOneActivity_FullMethodName              = "/user.Nft/GetOneActivity"
	Nft_SearchActivities_FullMethodName            = "/user.Nft/SearchActivities"
	Nft_GiveDc_FullMethodName                      = "/user.Nft/GiveDc"
	Nft_GetAllDc_FullMethodName                    = "/user.Nft/GetAllDc"
	Nft_SelectDc_FullMethodName                    = "/user.Nft/SelectDc"
	Nft_GetDcById_FullMethodName                   = "/user.Nft/GetDcById"
	Nft_GetMyDc_FullMethodName                     = "/user.Nft/GetMyDc"
	Nft_GetDigitalCollectionHistory_FullMethodName = "/user.Nft/GetDigitalCollectionHistory"
	Nft_CreatePool_FullMethodName                  = "/user.Nft/CreatePool"
	Nft_BuyFromPool_FullMethodName                 = "/user.Nft/BuyFromPool"
	Nft_SelectPool_FullMethodName                  = "/user.Nft/SelectPool"
	Nft_GetPoolById_FullMethodName                 = "/user.Nft/GetPoolById"
	Nft_GetAllPool_FullMethodName                  = "/user.Nft/GetAllPool"
	Nft_GetMyPool_FullMethodName                   = "/user.Nft/GetMyPool"
)

// NftClient is the client API for Nft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NftClient interface {
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

type nftClient struct {
	cc grpc.ClientConnInterface
}

func NewNftClient(cc grpc.ClientConnInterface) NftClient {
	return &nftClient{cc}
}

func (c *nftClient) GetMessageByHash(ctx context.Context, in *GetMessageByHashRequest, opts ...grpc.CallOption) (*GetMessageByHashDTO, error) {
	out := new(GetMessageByHashDTO)
	err := c.cc.Invoke(ctx, Nft_GetMessageByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Nft_CreateActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDcFromActivity(ctx context.Context, in *GetDcFromActivityRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Nft_GetDcFromActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetAllActivity(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	out := new(ActivityPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetAllActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetOneActivity(ctx context.Context, in *GetOneActivityRequest, opts ...grpc.CallOption) (*ActivityDetailsVO, error) {
	out := new(ActivityDetailsVO)
	err := c.cc.Invoke(ctx, Nft_GetOneActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SearchActivities(ctx context.Context, in *SearchActivitiesRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	out := new(ActivityPageVOList)
	err := c.cc.Invoke(ctx, Nft_SearchActivities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GiveDc(ctx context.Context, in *GiveDcRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Nft_GiveDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetAllDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetAllDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SelectDc(ctx context.Context, in *SelectDcRequest, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_SelectDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDcById(ctx context.Context, in *GetDcByIdRequest, opts ...grpc.CallOption) (*DcDetailVO, error) {
	out := new(DcDetailVO)
	err := c.cc.Invoke(ctx, Nft_GetDcById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetMyDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetMyDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDigitalCollectionHistory(ctx context.Context, in *GetDigitalCollectionHistoryRequest, opts ...grpc.CallOption) (*CollectionMessageOnChainVO, error) {
	out := new(CollectionMessageOnChainVO)
	err := c.cc.Invoke(ctx, Nft_GetDigitalCollectionHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Nft_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) BuyFromPool(ctx context.Context, in *BuyFromPoolRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Nft_BuyFromPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SelectPool(ctx context.Context, in *SelectPoolRequest, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_SelectPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetPoolById(ctx context.Context, in *GetPoolByIdRequest, opts ...grpc.CallOption) (*PoolDetailsVO, error) {
	out := new(PoolDetailsVO)
	err := c.cc.Invoke(ctx, Nft_GetPoolById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetAllPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetAllPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetMyPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetMyPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NftServer is the server API for Nft service.
// All implementations must embed UnimplementedNftServer
// for forward compatibility
type NftServer interface {
	GetMessageByHash(context.Context, *GetMessageByHashRequest) (*GetMessageByHashDTO, error)
	CreateActivity(context.Context, *CreateActivityRequest) (*Empty, error)
	GetDcFromActivity(context.Context, *GetDcFromActivityRequest) (*Empty, error)
	GetAllActivity(context.Context, *Empty) (*ActivityPageVOList, error)
	GetOneActivity(context.Context, *GetOneActivityRequest) (*ActivityDetailsVO, error)
	SearchActivities(context.Context, *SearchActivitiesRequest) (*ActivityPageVOList, error)
	GiveDc(context.Context, *GiveDcRequest) (*Empty, error)
	GetAllDc(context.Context, *Empty) (*DcPageVOList, error)
	SelectDc(context.Context, *SelectDcRequest) (*DcPageVOList, error)
	GetDcById(context.Context, *GetDcByIdRequest) (*DcDetailVO, error)
	GetMyDc(context.Context, *Empty) (*DcPageVOList, error)
	GetDigitalCollectionHistory(context.Context, *GetDigitalCollectionHistoryRequest) (*CollectionMessageOnChainVO, error)
	CreatePool(context.Context, *CreatePoolRequest) (*Empty, error)
	BuyFromPool(context.Context, *BuyFromPoolRequest) (*Empty, error)
	SelectPool(context.Context, *SelectPoolRequest) (*PoolPageVOList, error)
	GetPoolById(context.Context, *GetPoolByIdRequest) (*PoolDetailsVO, error)
	GetAllPool(context.Context, *Empty) (*PoolPageVOList, error)
	GetMyPool(context.Context, *Empty) (*PoolPageVOList, error)
	mustEmbedUnimplementedNftServer()
}

// UnimplementedNftServer must be embedded to have forward compatible implementations.
type UnimplementedNftServer struct {
}

func (UnimplementedNftServer) GetMessageByHash(context.Context, *GetMessageByHashRequest) (*GetMessageByHashDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageByHash not implemented")
}
func (UnimplementedNftServer) CreateActivity(context.Context, *CreateActivityRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedNftServer) GetDcFromActivity(context.Context, *GetDcFromActivityRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDcFromActivity not implemented")
}
func (UnimplementedNftServer) GetAllActivity(context.Context, *Empty) (*ActivityPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllActivity not implemented")
}
func (UnimplementedNftServer) GetOneActivity(context.Context, *GetOneActivityRequest) (*ActivityDetailsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneActivity not implemented")
}
func (UnimplementedNftServer) SearchActivities(context.Context, *SearchActivitiesRequest) (*ActivityPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchActivities not implemented")
}
func (UnimplementedNftServer) GiveDc(context.Context, *GiveDcRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveDc not implemented")
}
func (UnimplementedNftServer) GetAllDc(context.Context, *Empty) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDc not implemented")
}
func (UnimplementedNftServer) SelectDc(context.Context, *SelectDcRequest) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectDc not implemented")
}
func (UnimplementedNftServer) GetDcById(context.Context, *GetDcByIdRequest) (*DcDetailVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDcById not implemented")
}
func (UnimplementedNftServer) GetMyDc(context.Context, *Empty) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyDc not implemented")
}
func (UnimplementedNftServer) GetDigitalCollectionHistory(context.Context, *GetDigitalCollectionHistoryRequest) (*CollectionMessageOnChainVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDigitalCollectionHistory not implemented")
}
func (UnimplementedNftServer) CreatePool(context.Context, *CreatePoolRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedNftServer) BuyFromPool(context.Context, *BuyFromPoolRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyFromPool not implemented")
}
func (UnimplementedNftServer) SelectPool(context.Context, *SelectPoolRequest) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectPool not implemented")
}
func (UnimplementedNftServer) GetPoolById(context.Context, *GetPoolByIdRequest) (*PoolDetailsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolById not implemented")
}
func (UnimplementedNftServer) GetAllPool(context.Context, *Empty) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPool not implemented")
}
func (UnimplementedNftServer) GetMyPool(context.Context, *Empty) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyPool not implemented")
}
func (UnimplementedNftServer) mustEmbedUnimplementedNftServer() {}

// UnsafeNftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NftServer will
// result in compilation errors.
type UnsafeNftServer interface {
	mustEmbedUnimplementedNftServer()
}

func RegisterNftServer(s grpc.ServiceRegistrar, srv NftServer) {
	s.RegisterService(&Nft_ServiceDesc, srv)
}

func _Nft_GetMessageByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMessageByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMessageByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMessageByHash(ctx, req.(*GetMessageByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_CreateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDcFromActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDcFromActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDcFromActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDcFromActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDcFromActivity(ctx, req.(*GetDcFromActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetAllActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetAllActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetAllActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetAllActivity(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetOneActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetOneActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetOneActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetOneActivity(ctx, req.(*GetOneActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SearchActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SearchActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SearchActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SearchActivities(ctx, req.(*SearchActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GiveDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiveDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GiveDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GiveDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GiveDc(ctx, req.(*GiveDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetAllDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetAllDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetAllDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetAllDc(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SelectDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SelectDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SelectDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SelectDc(ctx, req.(*SelectDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDcById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDcByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDcById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDcById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDcById(ctx, req.(*GetDcByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetMyDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMyDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMyDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMyDc(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDigitalCollectionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDigitalCollectionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDigitalCollectionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDigitalCollectionHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDigitalCollectionHistory(ctx, req.(*GetDigitalCollectionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_BuyFromPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyFromPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).BuyFromPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_BuyFromPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).BuyFromPool(ctx, req.(*BuyFromPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SelectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SelectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SelectPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SelectPool(ctx, req.(*SelectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetPoolById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetPoolById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetPoolById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetPoolById(ctx, req.(*GetPoolByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetAllPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetAllPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetAllPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetAllPool(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetMyPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMyPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMyPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMyPool(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Nft_ServiceDesc is the grpc.ServiceDesc for Nft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Nft",
	HandlerType: (*NftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessageByHash",
			Handler:    _Nft_GetMessageByHash_Handler,
		},
		{
			MethodName: "CreateActivity",
			Handler:    _Nft_CreateActivity_Handler,
		},
		{
			MethodName: "GetDcFromActivity",
			Handler:    _Nft_GetDcFromActivity_Handler,
		},
		{
			MethodName: "GetAllActivity",
			Handler:    _Nft_GetAllActivity_Handler,
		},
		{
			MethodName: "GetOneActivity",
			Handler:    _Nft_GetOneActivity_Handler,
		},
		{
			MethodName: "SearchActivities",
			Handler:    _Nft_SearchActivities_Handler,
		},
		{
			MethodName: "GiveDc",
			Handler:    _Nft_GiveDc_Handler,
		},
		{
			MethodName: "GetAllDc",
			Handler:    _Nft_GetAllDc_Handler,
		},
		{
			MethodName: "SelectDc",
			Handler:    _Nft_SelectDc_Handler,
		},
		{
			MethodName: "GetDcById",
			Handler:    _Nft_GetDcById_Handler,
		},
		{
			MethodName: "GetMyDc",
			Handler:    _Nft_GetMyDc_Handler,
		},
		{
			MethodName: "GetDigitalCollectionHistory",
			Handler:    _Nft_GetDigitalCollectionHistory_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _Nft_CreatePool_Handler,
		},
		{
			MethodName: "BuyFromPool",
			Handler:    _Nft_BuyFromPool_Handler,
		},
		{
			MethodName: "SelectPool",
			Handler:    _Nft_SelectPool_Handler,
		},
		{
			MethodName: "GetPoolById",
			Handler:    _Nft_GetPoolById_Handler,
		},
		{
			MethodName: "GetAllPool",
			Handler:    _Nft_GetAllPool_Handler,
		},
		{
			MethodName: "GetMyPool",
			Handler:    _Nft_GetMyPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nft.proto",
}
