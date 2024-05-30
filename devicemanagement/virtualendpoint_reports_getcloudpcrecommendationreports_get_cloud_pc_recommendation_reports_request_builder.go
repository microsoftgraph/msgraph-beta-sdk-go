package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder provides operations to call the getCloudPcRecommendationReports method.
type VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderInternal instantiates a new VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) {
    m := &VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getCloudPcRecommendationReports", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder instantiates a new VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the device recommendation reports for Cloud PCs, such as the usage category report. The usage category report categorizes a Cloud PC as Undersized, Oversized, Rightsized, or Underutilized, and also provides the recommended SKU when the Cloud PC isn't Rightsized.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getcloudpcrecommendationreports?view=graph-rest-beta
func (m *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation get the device recommendation reports for Cloud PCs, such as the usage category report. The usage category report categorizes a Cloud PC as Undersized, Oversized, Rightsized, or Underutilized, and also provides the recommended SKU when the Cloud PC isn't Rightsized.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder when successful
func (m *VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder) {
    return NewVirtualendpointReportsGetcloudpcrecommendationreportsGetCloudPcRecommendationReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
