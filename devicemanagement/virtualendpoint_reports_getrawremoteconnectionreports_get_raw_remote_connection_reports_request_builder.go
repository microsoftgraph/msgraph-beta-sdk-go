package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder provides operations to call the getRawRemoteConnectionReports method.
type VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderInternal instantiates a new VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) {
    m := &VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRawRemoteConnectionReports", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder instantiates a new VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the raw real-time remote connection report for a Cloud PC without any calculation, such as roundTripTime or available bandwidth, which are aggregated hourly from the raw event data.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getrawremoteconnectionreports?view=graph-rest-beta
func (m *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation get the raw real-time remote connection report for a Cloud PC without any calculation, such as roundTripTime or available bandwidth, which are aggregated hourly from the raw event data.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder) {
    return NewVirtualendpointReportsGetrawremoteconnectionreportsGetRawRemoteConnectionReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
