package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder provides operations to call the getFrontlineReport method.
type VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderInternal instantiates a new VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) {
    m := &VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getFrontlineReport", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder instantiates a new VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the Windows 365 Frontline reports, such as real-time or historical data reports.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getfrontlinereport?view=graph-rest-beta
func (m *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGetfrontlinereportGetFrontlineReportPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation get the Windows 365 Frontline reports, such as real-time or historical data reports.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGetfrontlinereportGetFrontlineReportPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder when successful
func (m *VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder) {
    return NewVirtualendpointReportsGetfrontlinereportGetFrontlineReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
