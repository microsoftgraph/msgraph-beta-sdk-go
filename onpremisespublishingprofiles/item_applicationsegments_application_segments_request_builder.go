package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemApplicationsegmentsApplicationSegmentsRequestBuilder provides operations to manage the applicationSegments property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemApplicationsegmentsApplicationSegmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetQueryParameters get applicationSegments from onPremisesPublishingProfiles
type ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetQueryParameters struct {
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
// ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetQueryParameters
}
// ByIpApplicationSegmentId provides operations to manage the applicationSegments property of the microsoft.graph.onPremisesPublishingProfile entity.
// returns a *ItemApplicationsegmentsIpApplicationSegmentItemRequestBuilder when successful
func (m *ItemApplicationsegmentsApplicationSegmentsRequestBuilder) ByIpApplicationSegmentId(ipApplicationSegmentId string)(*ItemApplicationsegmentsIpApplicationSegmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ipApplicationSegmentId != "" {
        urlTplParams["ipApplicationSegment%2Did"] = ipApplicationSegmentId
    }
    return NewItemApplicationsegmentsIpApplicationSegmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemApplicationsegmentsApplicationSegmentsRequestBuilderInternal instantiates a new ItemApplicationsegmentsApplicationSegmentsRequestBuilder and sets the default values.
func NewItemApplicationsegmentsApplicationSegmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationsegmentsApplicationSegmentsRequestBuilder) {
    m := &ItemApplicationsegmentsApplicationSegmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/applicationSegments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemApplicationsegmentsApplicationSegmentsRequestBuilder instantiates a new ItemApplicationsegmentsApplicationSegmentsRequestBuilder and sets the default values.
func NewItemApplicationsegmentsApplicationSegmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationsegmentsApplicationSegmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemApplicationsegmentsApplicationSegmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemApplicationsegmentsCountRequestBuilder when successful
func (m *ItemApplicationsegmentsApplicationSegmentsRequestBuilder) Count()(*ItemApplicationsegmentsCountRequestBuilder) {
    return NewItemApplicationsegmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get applicationSegments from onPremisesPublishingProfiles
// returns a IpApplicationSegmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemApplicationsegmentsApplicationSegmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IpApplicationSegmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIpApplicationSegmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IpApplicationSegmentCollectionResponseable), nil
}
// ToGetRequestInformation get applicationSegments from onPremisesPublishingProfiles
// returns a *RequestInformation when successful
func (m *ItemApplicationsegmentsApplicationSegmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemApplicationsegmentsApplicationSegmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemApplicationsegmentsApplicationSegmentsRequestBuilder when successful
func (m *ItemApplicationsegmentsApplicationSegmentsRequestBuilder) WithUrl(rawUrl string)(*ItemApplicationsegmentsApplicationSegmentsRequestBuilder) {
    return NewItemApplicationsegmentsApplicationSegmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
