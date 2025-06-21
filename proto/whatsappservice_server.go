// versions:
// 	gofr-cli v0.6.0
// 	gofr.dev v1.37.0
// 	source: wa.proto

package proto

import "gofr.dev/pkg/gofr"

// Register the gRPC service in your app using the following code in your main.go:
//
// proto.RegisterWhatsappServiceServerWithGofr(app, &proto.NewWhatsappServiceGoFrServer())
//
// WhatsappServiceGoFrServer defines the gRPC server implementation.
// Customize the struct with required dependencies and fields as needed.

type WhatsappServiceGoFrServer struct {
 health *healthServer
}
func (s *WhatsappServiceGoFrServer) CheckUserState(ctx *gofr.Context) (any, error) {
	return &StatusQueryProto{}, nil
}
func (s *WhatsappServiceGoFrServer) Initialize(ctx *gofr.Context) (any, error) {
	return &InitializeQueryProto{}, nil
}
func (s *WhatsappServiceGoFrServer) SendMessage(ctx *gofr.Context) (any, error) {
	return &SendMessageResultProto{}, nil
}
func (s *WhatsappServiceGoFrServer) SayHello(ctx *gofr.Context) (any, error) {
	return &GreetingResponse{}, nil
}
func (s *WhatsappServiceGoFrServer) SignOut(ctx *gofr.Context) (any, error) {
	return &StatusQueryProto{}, nil
}
