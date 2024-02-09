package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder provides operations to call the getRealTimeRemoteConnectionStatus method.
type VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal instantiates a new VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPcId *string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    m := &VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRealTimeRemoteConnectionStatus(cloudPcId='{cloudPcId}')", pathParameters),
    }
    if cloudPcId != nil {
        m.BaseRequestBuilder.PathParameters["cloudPcId"] = *cloudPcId
    }
    return m
}
// NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder instantiates a new VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getRealTimeRemoteConnectionStatus
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation invoke function getRealTimeRemoteConnectionStatus
// returns a *RequestInformation when successful
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder when successful
func (m *VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return NewVirtualEndpointReportsGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
