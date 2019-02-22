package server

import (
	"golang.org/x/net/context"

	pb "grpc-swagger/protos/utlcore"
)

type smsService struct{}

func NewSmsService() *smsService {
	return &smsService{}
}

func (h smsService) SendSMS(ctx context.Context, r *pb.SMSRequest) (*pb.SMSResponse, error) {
	return &pb.SMSResponse{}, nil
}

func (h smsService) VerifySMS(ctx context.Context, r *pb.SMSVerifyRequest) (*pb.ServerStatus, error) {
	return &pb.ServerStatus{}, nil
}

func (h smsService) CheckPhone(ctx context.Context, r *pb.CheckPhoneRequest) (*pb.ServerStatus, error) {
	return &pb.ServerStatus{}, nil
}
