package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters get outboundProvisioningFlowSets from external
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters
}
// IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderInternal instantiates a new IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) {
    m := &IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder instantiates a new IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property outboundProvisioningFlowSets for external
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get outboundProvisioningFlowSets from external
// returns a OutboundProvisioningFlowSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property outboundProvisioningFlowSets in external
// returns a OutboundProvisioningFlowSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ProvisioningFlows provides operations to manage the provisioningFlows property of the microsoft.graph.industryData.outboundProvisioningFlowSet entity.
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) ProvisioningFlows()(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsProvisioningFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property outboundProvisioningFlowSets for external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get outboundProvisioningFlowSets from external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property outboundProvisioningFlowSets in external
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsOutboundProvisioningFlowSetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
