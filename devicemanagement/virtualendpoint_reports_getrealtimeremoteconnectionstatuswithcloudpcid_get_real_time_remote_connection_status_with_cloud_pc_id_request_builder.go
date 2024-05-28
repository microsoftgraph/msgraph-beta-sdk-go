package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder provides operations to call the getRealTimeRemoteConnectionStatus method.
type VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal instantiates a new VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPcId *string)(*VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    m := &VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRealTimeRemoteConnectionStatus(cloudPcId='{cloudPcId}')", pathParameters),
    }
    if cloudPcId != nil {
        m.BaseRequestBuilder.PathParameters["cloudPcId"] = *cloudPcId
    }
    return m
}
// NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder instantiates a new VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the real-time connection status information, such as signInStatus or daysSinceLastUse, for a Cloud PC.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getrealtimeremoteconnectionstatus?view=graph-rest-beta
func (m *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the real-time connection status information, such as signInStatus or daysSinceLastUse, for a Cloud PC.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder when successful
func (m *VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder) {
    return NewVirtualendpointReportsGetrealtimeremoteconnectionstatuswithcloudpcidGetRealTimeRemoteConnectionStatusWithCloudPcIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
