package api

import (
	"context"

	pb "github.com/Teerawat36167/PieFireDire/internal/pb"
)

func (s *Server) CountBeef(ctx context.Context, req *pb.BeefRequest) (*pb.BeefResponse, error) {
	text, err := FetchBaconIpsumText()
	if err != nil {
		return nil, err
	}

	counts := s.counter.GetBeefCounts(text)

	protoResponse := make(map[string]int32)
	for word, count := range counts {
		protoResponse[word] = int32(count)
	}

	return &pb.BeefResponse{
		BeefCounts: protoResponse,
	}, nil
}
