package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder provides operations to call the getRealTimeRemoteConnectionLatency method.
type VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal instantiates a new GetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    m := &VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRealTimeRemoteConnectionLatency(cloudPcId='{cloudPcId}')", pathParameters),
    }
    if cloudPcId != nil {
        m.BaseRequestBuilder.PathParameters["cloudPcId"] = *cloudPcId
    }
    return m
}
// NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder instantiates a new GetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getRealTimeRemoteConnectionLatency
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getRealTimeRemoteConnectionLatency
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionLatencyWithCloudPcIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
