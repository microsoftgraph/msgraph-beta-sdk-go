package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataOutboundProvisioningFlowSetsRequestBuilder provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataOutboundProvisioningFlowSetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters get a list of the outboundProvisioningFlowSet objects and their properties.
type IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters struct {
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
// IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters
}
// IndustryDataOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOutboundProvisioningFlowSetId provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) ByOutboundProvisioningFlowSetId(outboundProvisioningFlowSetId string)(*IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if outboundProvisioningFlowSetId != "" {
        urlTplParams["outboundProvisioningFlowSet%2Did"] = outboundProvisioningFlowSetId
    }
    return NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustryDataOutboundProvisioningFlowSetsRequestBuilderInternal instantiates a new IndustryDataOutboundProvisioningFlowSetsRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsRequestBuilder) {
    m := &IndustryDataOutboundProvisioningFlowSetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustryDataOutboundProvisioningFlowSetsRequestBuilder instantiates a new IndustryDataOutboundProvisioningFlowSetsRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataOutboundProvisioningFlowSetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustryDataOutboundProvisioningFlowSetsCountRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) Count()(*IndustryDataOutboundProvisioningFlowSetsCountRequestBuilder) {
    return NewIndustryDataOutboundProvisioningFlowSetsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the outboundProvisioningFlowSet objects and their properties.
// returns a OutboundProvisioningFlowSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataroot-list-outboundprovisioningflowsets?view=graph-rest-beta
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateOutboundProvisioningFlowSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetCollectionResponseable), nil
}
// Post create new navigation property to outboundProvisioningFlowSets for external
// returns a OutboundProvisioningFlowSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) Post(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateOutboundProvisioningFlowSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable), nil
}
// ToGetRequestInformation get a list of the outboundProvisioningFlowSet objects and their properties.
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to outboundProvisioningFlowSets for external
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustryDataOutboundProvisioningFlowSetsRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsRequestBuilder) WithUrl(rawUrl string)(*IndustryDataOutboundProvisioningFlowSetsRequestBuilder) {
    return NewIndustryDataOutboundProvisioningFlowSetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
