package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder provides operations to manage the provisioningFlows property of the microsoft.graph.industryData.outboundProvisioningFlowSet entity.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetQueryParameters get a list of the provisioningFlow objects and their properties.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetQueryParameters struct {
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
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetQueryParameters
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByProvisioningFlowId provides operations to manage the provisioningFlows property of the microsoft.graph.industryData.outboundProvisioningFlowSet entity.
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) ByProvisioningFlowId(provisioningFlowId string)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if provisioningFlowId != "" {
        urlTplParams["provisioningFlow%2Did"] = provisioningFlowId
    }
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderInternal instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) {
    m := &IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}/provisioningFlows{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsCountRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) Count()(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsCountRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the provisioningFlow objects and their properties.
// returns a ProvisioningFlowCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-outboundprovisioningflowset-list-provisioningflows?view=graph-rest-beta
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateProvisioningFlowCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowCollectionResponseable), nil
}
// Post create new navigation property to provisioningFlows for external
// returns a ProvisioningFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) Post(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderPostRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateProvisioningFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable), nil
}
// ToGetRequestInformation get a list of the provisioningFlow objects and their properties.
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to provisioningFlows for external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) ToPostRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
