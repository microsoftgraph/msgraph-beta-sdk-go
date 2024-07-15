package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder provides operations to manage the outboundProvisioningFlowSets property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters get outboundProvisioningFlowSets from external
type IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetQueryParameters
}
// IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderInternal instantiates a new IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) {
    m := &IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder instantiates a new IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property outboundProvisioningFlowSets for external
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
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
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, error) {
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
// returns a *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) ProvisioningFlows()(*IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsRequestBuilder) {
    return NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property outboundProvisioningFlowSets for external
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.OutboundProvisioningFlowSetable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) WithUrl(rawUrl string)(*IndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder) {
    return NewIndustryDataOutboundProvisioningFlowSetsOutboundProvisioningFlowSetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
