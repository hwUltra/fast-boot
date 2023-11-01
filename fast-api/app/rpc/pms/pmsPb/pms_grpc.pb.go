// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pb/pms.proto

package pmsPb

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

// PmsClient is the client API for Pms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PmsClient interface {
	ShopGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Shop, error)
	ShopList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ShopListResp, error)
	ShopAdd(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
	ShopUpdate(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessResp, error)
	ShopDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
	ShopOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*OptionsResp, error)
	PmsCategoryGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsCategory, error)
	PmsCategoryList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*PmsCategoryListResp, error)
	PmsCategoryAdd(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
	PmsCategoryUpdate(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsCategoryDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsCategoryOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepOptionsResp, error)
	PmsCategoryAttributeList(ctx context.Context, in *PmsCategoryAttributeListReq, opts ...grpc.CallOption) (*PmsCategoryAttributeListResp, error)
	PmsCategoryAttributeAdd(ctx context.Context, in *PmsCategoryAttributeAddReq, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsBrandGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsBrand, error)
	PmsBrandList(ctx context.Context, in *PmsBrandListReq, opts ...grpc.CallOption) (*PmsBrandListResp, error)
	PmsBrandAdd(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
	PmsBrandUpdate(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsBrandDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsBrandOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OptionsResp, error)
	PmsGoodsList(ctx context.Context, in *PmsGoodsListReq, opts ...grpc.CallOption) (*PmsGoodsListResp, error)
	PmsGoodsGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsGoodsInfo, error)
	PmsGoodsDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
	PmsGoodsEdit(ctx context.Context, in *PmsGoodsForm, opts ...grpc.CallOption) (*SuccessIdResp, error)
}

type pmsClient struct {
	cc grpc.ClientConnInterface
}

func NewPmsClient(cc grpc.ClientConnInterface) PmsClient {
	return &pmsClient{cc}
}

func (c *pmsClient) ShopGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ShopList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ShopListResp, error) {
	out := new(ShopListResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ShopAdd(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	out := new(SuccessIdResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ShopUpdate(ctx context.Context, in *ShopForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ShopDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) ShopOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*OptionsResp, error) {
	out := new(OptionsResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/ShopOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsCategory, error) {
	out := new(PmsCategory)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*PmsCategoryListResp, error) {
	out := new(PmsCategoryListResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryAdd(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	out := new(SuccessIdResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryUpdate(ctx context.Context, in *PmsCategoryForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepOptionsResp, error) {
	out := new(RepOptionsResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryAttributeList(ctx context.Context, in *PmsCategoryAttributeListReq, opts ...grpc.CallOption) (*PmsCategoryAttributeListResp, error) {
	out := new(PmsCategoryAttributeListResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryAttributeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsCategoryAttributeAdd(ctx context.Context, in *PmsCategoryAttributeAddReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsCategoryAttributeAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsBrand, error) {
	out := new(PmsBrand)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandList(ctx context.Context, in *PmsBrandListReq, opts ...grpc.CallOption) (*PmsBrandListResp, error) {
	out := new(PmsBrandListResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandAdd(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	out := new(SuccessIdResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandUpdate(ctx context.Context, in *PmsBrandForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsBrandOptions(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*OptionsResp, error) {
	out := new(OptionsResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsBrandOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsGoodsList(ctx context.Context, in *PmsGoodsListReq, opts ...grpc.CallOption) (*PmsGoodsListResp, error) {
	out := new(PmsGoodsListResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsGoodsGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*PmsGoodsInfo, error) {
	out := new(PmsGoodsInfo)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsGoodsGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsGoodsDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	out := new(SuccessResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsGoodsDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pmsClient) PmsGoodsEdit(ctx context.Context, in *PmsGoodsForm, opts ...grpc.CallOption) (*SuccessIdResp, error) {
	out := new(SuccessIdResp)
	err := c.cc.Invoke(ctx, "/pmsPb.pms/PmsGoodsEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PmsServer is the server API for Pms service.
// All implementations must embed UnimplementedPmsServer
// for forward compatibility
type PmsServer interface {
	ShopGet(context.Context, *IdReq) (*Shop, error)
	ShopList(context.Context, *ListReq) (*ShopListResp, error)
	ShopAdd(context.Context, *ShopForm) (*SuccessIdResp, error)
	ShopUpdate(context.Context, *ShopForm) (*SuccessResp, error)
	ShopDel(context.Context, *IdsReq) (*SuccessResp, error)
	ShopOptions(context.Context, *AnyReq) (*OptionsResp, error)
	PmsCategoryGet(context.Context, *IdReq) (*PmsCategory, error)
	PmsCategoryList(context.Context, *ListReq) (*PmsCategoryListResp, error)
	PmsCategoryAdd(context.Context, *PmsCategoryForm) (*SuccessIdResp, error)
	PmsCategoryUpdate(context.Context, *PmsCategoryForm) (*SuccessResp, error)
	PmsCategoryDel(context.Context, *IdsReq) (*SuccessResp, error)
	PmsCategoryOptions(context.Context, *IdReq) (*RepOptionsResp, error)
	PmsCategoryAttributeList(context.Context, *PmsCategoryAttributeListReq) (*PmsCategoryAttributeListResp, error)
	PmsCategoryAttributeAdd(context.Context, *PmsCategoryAttributeAddReq) (*SuccessResp, error)
	PmsBrandGet(context.Context, *IdReq) (*PmsBrand, error)
	PmsBrandList(context.Context, *PmsBrandListReq) (*PmsBrandListResp, error)
	PmsBrandAdd(context.Context, *PmsBrandForm) (*SuccessIdResp, error)
	PmsBrandUpdate(context.Context, *PmsBrandForm) (*SuccessResp, error)
	PmsBrandDel(context.Context, *IdsReq) (*SuccessResp, error)
	PmsBrandOptions(context.Context, *IdReq) (*OptionsResp, error)
	PmsGoodsList(context.Context, *PmsGoodsListReq) (*PmsGoodsListResp, error)
	PmsGoodsGet(context.Context, *IdReq) (*PmsGoodsInfo, error)
	PmsGoodsDel(context.Context, *IdsReq) (*SuccessResp, error)
	PmsGoodsEdit(context.Context, *PmsGoodsForm) (*SuccessIdResp, error)
	mustEmbedUnimplementedPmsServer()
}

// UnimplementedPmsServer must be embedded to have forward compatible implementations.
type UnimplementedPmsServer struct {
}

func (UnimplementedPmsServer) ShopGet(context.Context, *IdReq) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopGet not implemented")
}
func (UnimplementedPmsServer) ShopList(context.Context, *ListReq) (*ShopListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopList not implemented")
}
func (UnimplementedPmsServer) ShopAdd(context.Context, *ShopForm) (*SuccessIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopAdd not implemented")
}
func (UnimplementedPmsServer) ShopUpdate(context.Context, *ShopForm) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopUpdate not implemented")
}
func (UnimplementedPmsServer) ShopDel(context.Context, *IdsReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopDel not implemented")
}
func (UnimplementedPmsServer) ShopOptions(context.Context, *AnyReq) (*OptionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShopOptions not implemented")
}
func (UnimplementedPmsServer) PmsCategoryGet(context.Context, *IdReq) (*PmsCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryGet not implemented")
}
func (UnimplementedPmsServer) PmsCategoryList(context.Context, *ListReq) (*PmsCategoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryList not implemented")
}
func (UnimplementedPmsServer) PmsCategoryAdd(context.Context, *PmsCategoryForm) (*SuccessIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryAdd not implemented")
}
func (UnimplementedPmsServer) PmsCategoryUpdate(context.Context, *PmsCategoryForm) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryUpdate not implemented")
}
func (UnimplementedPmsServer) PmsCategoryDel(context.Context, *IdsReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryDel not implemented")
}
func (UnimplementedPmsServer) PmsCategoryOptions(context.Context, *IdReq) (*RepOptionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryOptions not implemented")
}
func (UnimplementedPmsServer) PmsCategoryAttributeList(context.Context, *PmsCategoryAttributeListReq) (*PmsCategoryAttributeListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryAttributeList not implemented")
}
func (UnimplementedPmsServer) PmsCategoryAttributeAdd(context.Context, *PmsCategoryAttributeAddReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsCategoryAttributeAdd not implemented")
}
func (UnimplementedPmsServer) PmsBrandGet(context.Context, *IdReq) (*PmsBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandGet not implemented")
}
func (UnimplementedPmsServer) PmsBrandList(context.Context, *PmsBrandListReq) (*PmsBrandListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandList not implemented")
}
func (UnimplementedPmsServer) PmsBrandAdd(context.Context, *PmsBrandForm) (*SuccessIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandAdd not implemented")
}
func (UnimplementedPmsServer) PmsBrandUpdate(context.Context, *PmsBrandForm) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandUpdate not implemented")
}
func (UnimplementedPmsServer) PmsBrandDel(context.Context, *IdsReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandDel not implemented")
}
func (UnimplementedPmsServer) PmsBrandOptions(context.Context, *IdReq) (*OptionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsBrandOptions not implemented")
}
func (UnimplementedPmsServer) PmsGoodsList(context.Context, *PmsGoodsListReq) (*PmsGoodsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsGoodsList not implemented")
}
func (UnimplementedPmsServer) PmsGoodsGet(context.Context, *IdReq) (*PmsGoodsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsGoodsGet not implemented")
}
func (UnimplementedPmsServer) PmsGoodsDel(context.Context, *IdsReq) (*SuccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsGoodsDel not implemented")
}
func (UnimplementedPmsServer) PmsGoodsEdit(context.Context, *PmsGoodsForm) (*SuccessIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PmsGoodsEdit not implemented")
}
func (UnimplementedPmsServer) mustEmbedUnimplementedPmsServer() {}

// UnsafePmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PmsServer will
// result in compilation errors.
type UnsafePmsServer interface {
	mustEmbedUnimplementedPmsServer()
}

func RegisterPmsServer(s grpc.ServiceRegistrar, srv PmsServer) {
	s.RegisterService(&Pms_ServiceDesc, srv)
}

func _Pms_ShopGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopGet(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ShopList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ShopAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopAdd(ctx, req.(*ShopForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ShopUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopUpdate(ctx, req.(*ShopForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ShopDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopDel(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_ShopOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).ShopOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/ShopOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).ShopOptions(ctx, req.(*AnyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryGet(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsCategoryForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryAdd(ctx, req.(*PmsCategoryForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsCategoryForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryUpdate(ctx, req.(*PmsCategoryForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryDel(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryOptions(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryAttributeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsCategoryAttributeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryAttributeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryAttributeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryAttributeList(ctx, req.(*PmsCategoryAttributeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsCategoryAttributeAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsCategoryAttributeAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsCategoryAttributeAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsCategoryAttributeAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsCategoryAttributeAdd(ctx, req.(*PmsCategoryAttributeAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandGet(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsBrandListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandList(ctx, req.(*PmsBrandListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsBrandForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandAdd(ctx, req.(*PmsBrandForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsBrandForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandUpdate(ctx, req.(*PmsBrandForm))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandDel(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsBrandOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsBrandOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsBrandOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsBrandOptions(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsGoodsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsGoodsList(ctx, req.(*PmsGoodsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsGoodsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsGoodsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsGoodsGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsGoodsGet(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsGoodsDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsGoodsDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsGoodsDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsGoodsDel(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pms_PmsGoodsEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PmsGoodsForm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmsServer).PmsGoodsEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmsPb.pms/PmsGoodsEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmsServer).PmsGoodsEdit(ctx, req.(*PmsGoodsForm))
	}
	return interceptor(ctx, in, info, handler)
}

// Pms_ServiceDesc is the grpc.ServiceDesc for Pms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pmsPb.pms",
	HandlerType: (*PmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShopGet",
			Handler:    _Pms_ShopGet_Handler,
		},
		{
			MethodName: "ShopList",
			Handler:    _Pms_ShopList_Handler,
		},
		{
			MethodName: "ShopAdd",
			Handler:    _Pms_ShopAdd_Handler,
		},
		{
			MethodName: "ShopUpdate",
			Handler:    _Pms_ShopUpdate_Handler,
		},
		{
			MethodName: "ShopDel",
			Handler:    _Pms_ShopDel_Handler,
		},
		{
			MethodName: "ShopOptions",
			Handler:    _Pms_ShopOptions_Handler,
		},
		{
			MethodName: "PmsCategoryGet",
			Handler:    _Pms_PmsCategoryGet_Handler,
		},
		{
			MethodName: "PmsCategoryList",
			Handler:    _Pms_PmsCategoryList_Handler,
		},
		{
			MethodName: "PmsCategoryAdd",
			Handler:    _Pms_PmsCategoryAdd_Handler,
		},
		{
			MethodName: "PmsCategoryUpdate",
			Handler:    _Pms_PmsCategoryUpdate_Handler,
		},
		{
			MethodName: "PmsCategoryDel",
			Handler:    _Pms_PmsCategoryDel_Handler,
		},
		{
			MethodName: "PmsCategoryOptions",
			Handler:    _Pms_PmsCategoryOptions_Handler,
		},
		{
			MethodName: "PmsCategoryAttributeList",
			Handler:    _Pms_PmsCategoryAttributeList_Handler,
		},
		{
			MethodName: "PmsCategoryAttributeAdd",
			Handler:    _Pms_PmsCategoryAttributeAdd_Handler,
		},
		{
			MethodName: "PmsBrandGet",
			Handler:    _Pms_PmsBrandGet_Handler,
		},
		{
			MethodName: "PmsBrandList",
			Handler:    _Pms_PmsBrandList_Handler,
		},
		{
			MethodName: "PmsBrandAdd",
			Handler:    _Pms_PmsBrandAdd_Handler,
		},
		{
			MethodName: "PmsBrandUpdate",
			Handler:    _Pms_PmsBrandUpdate_Handler,
		},
		{
			MethodName: "PmsBrandDel",
			Handler:    _Pms_PmsBrandDel_Handler,
		},
		{
			MethodName: "PmsBrandOptions",
			Handler:    _Pms_PmsBrandOptions_Handler,
		},
		{
			MethodName: "PmsGoodsList",
			Handler:    _Pms_PmsGoodsList_Handler,
		},
		{
			MethodName: "PmsGoodsGet",
			Handler:    _Pms_PmsGoodsGet_Handler,
		},
		{
			MethodName: "PmsGoodsDel",
			Handler:    _Pms_PmsGoodsDel_Handler,
		},
		{
			MethodName: "PmsGoodsEdit",
			Handler:    _Pms_PmsGoodsEdit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/pms.proto",
}
