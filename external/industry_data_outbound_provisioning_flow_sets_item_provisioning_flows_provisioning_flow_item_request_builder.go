package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder provides operations to manage the provisioningFlows property of the microsoft.graph.industryData.outboundProvisioningFlowSet entity.
type IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetQueryParameters read the properties and relationships of an administrativeUnitProvisioningFlow object.
type IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetQueryParameters
}
// IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderInternal instantiates a new IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) {
    m := &IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}/provisioningFlows/{provisioningFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder instantiates a new IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder and sets the default values.
func NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a classGroupProvisioningFlow object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-classgroupprovisioningflow-delete?view=graph-rest-1.0
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an administrativeUnitProvisioningFlow object.
// returns a ProvisioningFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-administrativeunitprovisioningflow-get?view=graph-rest-1.0
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of an administrativeUnitProvisioningFlow object.
// returns a ProvisioningFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-administrativeunitprovisioningflow-update?view=graph-rest-1.0
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete a classGroupProvisioningFlow object.
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an administrativeUnitProvisioningFlow object.
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an administrativeUnitProvisioningFlow object.
// returns a *RequestInformation when successful
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ProvisioningFlowable, requestConfiguration *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder when successful
func (m *IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) WithUrl(rawUrl string)(*IndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder) {
    return NewIndustryDataOutboundProvisioningFlowSetsItemProvisioningFlowsProvisioningFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
