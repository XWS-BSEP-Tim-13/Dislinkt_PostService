// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: post_service.proto

package proto

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

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetByUser(ctx context.Context, in *GetByUserRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	DeletePost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllRequest, error)
	CreatePost(ctx context.Context, in *NewPostRequest, opts ...grpc.CallOption) (*NewPost, error)
	ReactToPost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	GetFeedPosts(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	GetFeedPostsAnonymous(ctx context.Context, in *FeedRequestAnonymous, opts ...grpc.CallOption) (*FeedResponse, error)
	CreateCommentOnPost(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	GetImage(ctx context.Context, in *ImageResponse, opts ...grpc.CallOption) (*ImageRequest, error)
	GetMessagesForUsers(ctx context.Context, in *GetByUserRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	GetMessagesForUser(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetByUserResponse, error)
	SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetByUser(ctx context.Context, in *GetByUserRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAllRequest, error) {
	out := new(GetAllRequest)
	err := c.cc.Invoke(ctx, "/post.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *NewPostRequest, opts ...grpc.CallOption) (*NewPost, error) {
	out := new(NewPost)
	err := c.cc.Invoke(ctx, "/post.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReactToPost(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/ReactToPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetFeedPosts(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetFeedPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetFeedPostsAnonymous(ctx context.Context, in *FeedRequestAnonymous, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetFeedPostsAnonymous", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateCommentOnPost(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreateCommentOnPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UploadImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetImage(ctx context.Context, in *ImageResponse, opts ...grpc.CallOption) (*ImageRequest, error) {
	out := new(ImageRequest)
	err := c.cc.Invoke(ctx, "/post.PostService/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetMessagesForUsers(ctx context.Context, in *GetByUserRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetMessagesForUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetMessagesForUser(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetByUserResponse, error) {
	out := new(GetByUserResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetMessagesForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) SaveMessage(ctx context.Context, in *SaveMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/post.PostService/SaveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetByUser(context.Context, *GetByUserRequest) (*GetAllResponse, error)
	DeletePost(context.Context, *GetRequest) (*GetAllRequest, error)
	CreatePost(context.Context, *NewPostRequest) (*NewPost, error)
	ReactToPost(context.Context, *ReactionRequest) (*ReactionResponse, error)
	GetFeedPosts(context.Context, *FeedRequest) (*FeedResponse, error)
	GetFeedPostsAnonymous(context.Context, *FeedRequestAnonymous) (*FeedResponse, error)
	CreateCommentOnPost(context.Context, *CommentRequest) (*CommentResponse, error)
	UploadImage(context.Context, *ImageRequest) (*ImageResponse, error)
	GetImage(context.Context, *ImageResponse) (*ImageRequest, error)
	GetMessagesForUsers(context.Context, *GetByUserRequest) (*MessageResponse, error)
	GetMessagesForUser(context.Context, *GetAllRequest) (*GetByUserResponse, error)
	SaveMessage(context.Context, *SaveMessageRequest) (*Message, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPostServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPostServiceServer) GetByUser(context.Context, *GetByUserRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUser not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *GetRequest) (*GetAllRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) CreatePost(context.Context, *NewPostRequest) (*NewPost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) ReactToPost(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReactToPost not implemented")
}
func (UnimplementedPostServiceServer) GetFeedPosts(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedPosts not implemented")
}
func (UnimplementedPostServiceServer) GetFeedPostsAnonymous(context.Context, *FeedRequestAnonymous) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedPostsAnonymous not implemented")
}
func (UnimplementedPostServiceServer) CreateCommentOnPost(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommentOnPost not implemented")
}
func (UnimplementedPostServiceServer) UploadImage(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedPostServiceServer) GetImage(context.Context, *ImageResponse) (*ImageRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedPostServiceServer) GetMessagesForUsers(context.Context, *GetByUserRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesForUsers not implemented")
}
func (UnimplementedPostServiceServer) GetMessagesForUser(context.Context, *GetAllRequest) (*GetByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesForUser not implemented")
}
func (UnimplementedPostServiceServer) SaveMessage(context.Context, *SaveMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMessage not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetByUser(ctx, req.(*GetByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*NewPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReactToPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReactToPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/ReactToPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReactToPost(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetFeedPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetFeedPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetFeedPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetFeedPosts(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetFeedPostsAnonymous_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequestAnonymous)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetFeedPostsAnonymous(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetFeedPostsAnonymous",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetFeedPostsAnonymous(ctx, req.(*FeedRequestAnonymous))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreateCommentOnPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreateCommentOnPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreateCommentOnPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreateCommentOnPost(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UploadImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetImage(ctx, req.(*ImageResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetMessagesForUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetMessagesForUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetMessagesForUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetMessagesForUsers(ctx, req.(*GetByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetMessagesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetMessagesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetMessagesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetMessagesForUser(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_SaveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).SaveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/SaveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).SaveMessage(ctx, req.(*SaveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PostService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PostService_GetAll_Handler,
		},
		{
			MethodName: "GetByUser",
			Handler:    _PostService_GetByUser_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "ReactToPost",
			Handler:    _PostService_ReactToPost_Handler,
		},
		{
			MethodName: "GetFeedPosts",
			Handler:    _PostService_GetFeedPosts_Handler,
		},
		{
			MethodName: "GetFeedPostsAnonymous",
			Handler:    _PostService_GetFeedPostsAnonymous_Handler,
		},
		{
			MethodName: "CreateCommentOnPost",
			Handler:    _PostService_CreateCommentOnPost_Handler,
		},
		{
			MethodName: "UploadImage",
			Handler:    _PostService_UploadImage_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _PostService_GetImage_Handler,
		},
		{
			MethodName: "GetMessagesForUsers",
			Handler:    _PostService_GetMessagesForUsers_Handler,
		},
		{
			MethodName: "GetMessagesForUser",
			Handler:    _PostService_GetMessagesForUser_Handler,
		},
		{
			MethodName: "SaveMessage",
			Handler:    _PostService_SaveMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_service.proto",
}
