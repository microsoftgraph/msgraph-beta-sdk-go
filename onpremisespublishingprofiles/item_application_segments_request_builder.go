package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemApplicationSegmentsRequestBuilder provides operations to manage the applicationSegments property of the microsoft.graph.onPremisesPublishingProfile entity.
type ItemApplicationSegmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemApplicationSegmentsRequestBuilderGetQueryParameters represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
type ItemApplicationSegmentsRequestBuilderGetQueryParameters struct {
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
// ItemApplicationSegmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemApplicationSegmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemApplicationSegmentsRequestBuilderGetQueryParameters
}
// ByIpApplicationSegmentId provides operations to manage the applicationSegments property of the microsoft.graph.onPremisesPublishingProfile entity.
// returns a *ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder when successful
func (m *ItemApplicationSegmentsRequestBuilder) ByIpApplicationSegmentId(ipApplicationSegmentId string)(*ItemApplicationSegmentsIpApplicationSegmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ipApplicationSegmentId != "" {
        urlTplParams["ipApplicationSegment%2Did"] = ipApplicationSegmentId
    }
    return NewItemApplicationSegmentsIpApplicationSegmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemApplicationSegmentsRequestBuilderInternal instantiates a new ItemApplicationSegmentsRequestBuilder and sets the default values.
func NewItemApplicationSegmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationSegmentsRequestBuilder) {
    m := &ItemApplicationSegmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/applicationSegments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemApplicationSegmentsRequestBuilder instantiates a new ItemApplicationSegmentsRequestBuilder and sets the default values.
func NewItemApplicationSegmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApplicationSegmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemApplicationSegmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemApplicationSegmentsCountRequestBuilder when successful
func (m *ItemApplicationSegmentsRequestBuilder) Count()(*ItemApplicationSegmentsCountRequestBuilder) {
    return NewItemApplicationSegmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
// returns a IpApplicationSegmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemApplicationSegmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemApplicationSegmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IpApplicationSegmentCollectionResponseable, error) {
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
// ToGetRequestInformation represents the segment configurations that are allowed for an on-premises non-web application published through Microsoft Entra application proxy.
// returns a *RequestInformation when successful
func (m *ItemApplicationSegmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemApplicationSegmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemApplicationSegmentsRequestBuilder when successful
func (m *ItemApplicationSegmentsRequestBuilder) WithUrl(rawUrl string)(*ItemApplicationSegmentsRequestBuilder) {
    return NewItemApplicationSegmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
