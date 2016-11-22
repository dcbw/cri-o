package server

import (
	"golang.org/x/net/context"
	pb "k8s.io/kubernetes/pkg/kubelet/api/v1alpha1/runtime"
)

// Status returns the status of the runtime
func (s *Server) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {

	// Deal with Runtime conditions
	runtimeReady, err := s.runtime.RuntimeReady()
	if err != nil {
		return nil, err
	}
	networkReady, err := s.runtime.NetworkReady()
	if err != nil {
		return nil, err
	}

	// Use vendored strings
	runtimeReadyConditionString := pb.RuntimeReady
	networkReadyConditionString := pb.NetworkReady

	resp := &pb.StatusResponse{
		Status: &pb.RuntimeStatus{
			Conditions: []*pb.RuntimeCondition{
				&pb.RuntimeCondition{
					Type:   &runtimeReadyConditionString,
					Status: &runtimeReady,
				},
				&pb.RuntimeCondition{
					Type:   &networkReadyConditionString,
					Status: &networkReady,
				},
			},
		},
	}

	return resp, nil
}
