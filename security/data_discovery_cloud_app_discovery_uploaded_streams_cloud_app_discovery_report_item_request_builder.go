package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder provides operations to manage the uploadedStreams property of the microsoft.graph.security.dataDiscoveryReport entity.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetQueryParameters a collection of streams available for generating cloud discovery report.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetQueryParameters
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderInternal instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) {
    m := &DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/dataDiscovery/cloudAppDiscovery/uploadedStreams/{cloudAppDiscoveryReport%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property uploadedStreams for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of streams available for generating cloud discovery report.
// returns a CloudAppDiscoveryReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateCloudAppDiscoveryReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable), nil
}
// MicrosoftGraphSecurityAggregatedAppsDetailsWithPeriod provides operations to call the aggregatedAppsDetails method.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) MicrosoftGraphSecurityAggregatedAppsDetailsWithPeriod(period *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsItemMicrosoftGraphSecurityAggregatedAppsDetailsWithPeriodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, period)
}
// Patch update the navigation property uploadedStreams in security
// returns a CloudAppDiscoveryReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateCloudAppDiscoveryReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable), nil
}
// ToDeleteRequestInformation delete navigation property uploadedStreams for security
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of streams available for generating cloud discovery report.
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property uploadedStreams in security
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) WithUrl(rawUrl string)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
