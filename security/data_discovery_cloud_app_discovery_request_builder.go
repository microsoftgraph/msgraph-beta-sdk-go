package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataDiscoveryCloudAppDiscoveryRequestBuilder provides operations to manage the cloudAppDiscovery property of the microsoft.graph.security.dataDiscoveryRoot entity.
type DataDiscoveryCloudAppDiscoveryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataDiscoveryCloudAppDiscoveryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DataDiscoveryCloudAppDiscoveryRequestBuilderGetQueryParameters the available entities are IP addresses, devices, and users who access a cloud app.
type DataDiscoveryCloudAppDiscoveryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DataDiscoveryCloudAppDiscoveryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataDiscoveryCloudAppDiscoveryRequestBuilderGetQueryParameters
}
// DataDiscoveryCloudAppDiscoveryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDataDiscoveryCloudAppDiscoveryRequestBuilderInternal instantiates a new DataDiscoveryCloudAppDiscoveryRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryRequestBuilder) {
    m := &DataDiscoveryCloudAppDiscoveryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/dataDiscovery/cloudAppDiscovery{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDataDiscoveryCloudAppDiscoveryRequestBuilder instantiates a new DataDiscoveryCloudAppDiscoveryRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataDiscoveryCloudAppDiscoveryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudAppDiscovery for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) Delete(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available entities are IP addresses, devices, and users who access a cloud app.
// returns a DataDiscoveryReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) Get(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDataDiscoveryReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable), nil
}
// Patch update the navigation property cloudAppDiscovery in security
// returns a DataDiscoveryReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateDataDiscoveryReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable), nil
}
// ToDeleteRequestInformation delete navigation property cloudAppDiscovery for security
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the available entities are IP addresses, devices, and users who access a cloud app.
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudAppDiscovery in security
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.DataDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UploadedStreams provides operations to manage the uploadedStreams property of the microsoft.graph.security.dataDiscoveryReport entity.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) UploadedStreams()(*DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DataDiscoveryCloudAppDiscoveryRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryRequestBuilder) WithUrl(rawUrl string)(*DataDiscoveryCloudAppDiscoveryRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
