package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetQueryParameters a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByThreatAssessmentResultId provides operations to manage the results property of the microsoft.graph.threatAssessmentRequest entity.
// returns a *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) ByThreatAssessmentResultId(threatAssessmentResultId string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if threatAssessmentResultId != "" {
        urlTplParams["threatAssessmentResult%2Did"] = threatAssessmentResultId
    }
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsThreatAssessmentResultItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) {
    m := &ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest%2Did}/results{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder instantiates a new ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) Count()(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
// returns a ThreatAssessmentResultCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultCollectionResponseable), nil
}
// Post create new navigation property to results for groups
// returns a ThreatAssessmentResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateThreatAssessmentResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable), nil
}
// ToGetRequestInformation a collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it.
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to results for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ThreatAssessmentResultable, requestConfiguration *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder) {
    return NewItemSitesItemInformationProtectionThreatAssessmentRequestsItemResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
