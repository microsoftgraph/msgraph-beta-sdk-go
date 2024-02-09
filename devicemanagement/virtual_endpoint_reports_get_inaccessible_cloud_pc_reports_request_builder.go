package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder provides operations to call the getInaccessibleCloudPcReports method.
type VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderInternal instantiates a new VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) {
    m := &VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getInaccessibleCloudPcReports", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder instantiates a new VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get inaccessible Cloud PCs with details, including the latest health state, failed connection count, failed health check count, and system status. An inaccessible Cloud PC represents a Cloud PC that is in an unavailable state (at least one of the health checks failed) or has consecutive user connections failure.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getinaccessiblecloudpcreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetInaccessibleCloudPcReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation get inaccessible Cloud PCs with details, including the latest health state, failed connection count, failed health check count, and system status. An inaccessible Cloud PC represents a Cloud PC that is in an unavailable state (at least one of the health checks failed) or has consecutive user connections failure.
// returns a *RequestInformation when successful
func (m *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetInaccessibleCloudPcReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder when successful
func (m *VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetInaccessibleCloudPcReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
