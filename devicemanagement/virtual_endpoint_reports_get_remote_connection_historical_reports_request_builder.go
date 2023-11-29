package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder provides operations to call the getRemoteConnectionHistoricalReports method.
type VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    m := &VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getRemoteConnectionHistoricalReports", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the remote connection history records of a Cloud PC during a given period. This report contains data such as signInDateTime, signOutDateTime, usageInHour, remoteSignInTimeInSec and roundTripTimeInMsP50, and so on. This data is aggregated hourly for a specified time period, such as the last seven days.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getremoteconnectionhistoricalreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation get the remote connection history records of a Cloud PC during a given period. This report contains data such as signInDateTime, signOutDateTime, usageInHour, remoteSignInTimeInSec and roundTripTimeInMsP50, and so on. This data is aggregated hourly for a specified time period, such as the last seven days.
func (m *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
