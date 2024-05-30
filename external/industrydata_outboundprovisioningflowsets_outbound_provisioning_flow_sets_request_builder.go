package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters get a list of the outboundProvisioningFlowSet objects and their properties.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters struct {
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
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetQueryParameters
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOutboundProvisioningFlowSetId provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) ByOutboundProvisioningFlowSetId(outboundProvisioningFlowSetId string)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if outboundProvisioningFlowSetId != "" {
        urlTplParams["outboundProvisioningFlowSet%2Did"] = outboundProvisioningFlowSetId
    }
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderInternal instantiates a new IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) {
    m := &IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder instantiates a new IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustrydataOutboundprovisioningflowsetsCountRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) Count()(*IndustrydataOutboundprovisioningflowsetsCountRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the outboundProvisioningFlowSet objects and their properties.
// returns a OutboundProvisioningFlowSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataroot-list-outboundprovisioningflowsets?view=graph-rest-beta
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetCollectionResponseable, error) {
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
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) Post(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
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
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
