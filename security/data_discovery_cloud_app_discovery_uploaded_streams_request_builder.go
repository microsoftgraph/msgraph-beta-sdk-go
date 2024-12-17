package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder provides operations to manage the uploadedStreams property of the microsoft.graph.security.dataDiscoveryReport entity.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetQueryParameters get visibility into all the manually uploaded streams from your firewalls and proxies.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetQueryParameters
}
// DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudAppDiscoveryReportId provides operations to manage the uploadedStreams property of the microsoft.graph.security.dataDiscoveryReport entity.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) ByCloudAppDiscoveryReportId(cloudAppDiscoveryReportId string)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudAppDiscoveryReportId != "" {
        urlTplParams["cloudAppDiscoveryReport%2Did"] = cloudAppDiscoveryReportId
    }
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCloudAppDiscoveryReportItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderInternal instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) {
    m := &DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/dataDiscovery/cloudAppDiscovery/uploadedStreams{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder instantiates a new DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder and sets the default values.
func NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsCountRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) Count()(*DataDiscoveryCloudAppDiscoveryUploadedStreamsCountRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get visibility into all the manually uploaded streams from your firewalls and proxies.
// returns a CloudAppDiscoveryReportCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-datadiscoveryreport-list-uploadedstreams?view=graph-rest-beta
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) Get(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateCloudAppDiscoveryReportCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportCollectionResponseable), nil
}
// Post create new navigation property to uploadedStreams for security
// returns a CloudAppDiscoveryReportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get visibility into all the manually uploaded streams from your firewalls and proxies.
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to uploadedStreams for security
// returns a *RequestInformation when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CloudAppDiscoveryReportable, requestConfiguration *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder when successful
func (m *DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) WithUrl(rawUrl string)(*DataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder) {
    return NewDataDiscoveryCloudAppDiscoveryUploadedStreamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
