// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisvideosignalingchannelsiface provides an interface to enable mocking the Amazon Kinesis Video Signaling Channels service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisvideosignalingchannelsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels"
)

// KinesisVideoSignalingChannelsAPI provides an interface to enable mocking the
// kinesisvideosignalingchannels.KinesisVideoSignalingChannels service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Video Signaling Channels.
//    func myFunc(svc kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI) bool {
//        // Make svc.GetIceServerConfig request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kinesisvideosignalingchannels.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisVideoSignalingChannelsClient struct {
//        kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI
//    }
//    func (m *mockKinesisVideoSignalingChannelsClient) GetIceServerConfig(input *kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisVideoSignalingChannelsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KinesisVideoSignalingChannelsAPI interface {
	GetIceServerConfig(*kinesisvideosignalingchannels.GetIceServerConfigInput) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error)
	GetIceServerConfigWithContext(aws.Context, *kinesisvideosignalingchannels.GetIceServerConfigInput, ...request.Option) (*kinesisvideosignalingchannels.GetIceServerConfigOutput, error)
	GetIceServerConfigRequest(*kinesisvideosignalingchannels.GetIceServerConfigInput) (*request.Request, *kinesisvideosignalingchannels.GetIceServerConfigOutput)

	SendAlexaOfferToMaster(*kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error)
	SendAlexaOfferToMasterWithContext(aws.Context, *kinesisvideosignalingchannels.SendAlexaOfferToMasterInput, ...request.Option) (*kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput, error)
	SendAlexaOfferToMasterRequest(*kinesisvideosignalingchannels.SendAlexaOfferToMasterInput) (*request.Request, *kinesisvideosignalingchannels.SendAlexaOfferToMasterOutput)
}

var _ KinesisVideoSignalingChannelsAPI = (*kinesisvideosignalingchannels.KinesisVideoSignalingChannels)(nil)
