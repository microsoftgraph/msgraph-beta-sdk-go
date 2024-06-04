package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataInboundflowsInboundFlowItemRequestBuilder provides operations to manage the inboundFlows property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataInboundflowsInboundFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataInboundflowsInboundFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataInboundflowsInboundFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustrydataInboundflowsInboundFlowItemRequestBuilderGetQueryParameters read the properties and relationships of an inboundFileFlow object.
type IndustrydataInboundflowsInboundFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataInboundflowsInboundFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataInboundflowsInboundFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataInboundflowsInboundFlowItemRequestBuilderGetQueryParameters
}
// IndustrydataInboundflowsInboundFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataInboundflowsInboundFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataInboundflowsInboundFlowItemRequestBuilderInternal instantiates a new IndustrydataInboundflowsInboundFlowItemRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsInboundFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsInboundFlowItemRequestBuilder) {
    m := &IndustrydataInboundflowsInboundFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/inboundFlows/{inboundFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataInboundflowsInboundFlowItemRequestBuilder instantiates a new IndustrydataInboundflowsInboundFlowItemRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsInboundFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsInboundFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataInboundflowsInboundFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataConnector provides operations to manage the dataConnector property of the microsoft.graph.industryData.inboundFlow entity.
// returns a *IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) DataConnector()(*IndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilder) {
    return NewIndustrydataInboundflowsItemDataconnectorDataConnectorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an inboundFlow object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundflow-delete?view=graph-rest-beta
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an inboundFileFlow object.
// returns a InboundFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundfileflow-get?view=graph-rest-beta
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateInboundFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable), nil
}
// Patch update the properties of an inboundFlow object.
// returns a InboundFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundflow-update?view=graph-rest-beta
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateInboundFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable), nil
}
// ToDeleteRequestInformation delete an inboundFlow object.
// returns a *RequestInformation when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an inboundFileFlow object.
// returns a *RequestInformation when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an inboundFlow object.
// returns a *RequestInformation when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, requestConfiguration *IndustrydataInboundflowsInboundFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataInboundflowsInboundFlowItemRequestBuilder when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) WithUrl(rawUrl string)(*IndustrydataInboundflowsInboundFlowItemRequestBuilder) {
    return NewIndustrydataInboundflowsInboundFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Year provides operations to manage the year property of the microsoft.graph.industryData.inboundFlow entity.
// returns a *IndustrydataInboundflowsItemYearRequestBuilder when successful
func (m *IndustrydataInboundflowsInboundFlowItemRequestBuilder) Year()(*IndustrydataInboundflowsItemYearRequestBuilder) {
    return NewIndustrydataInboundflowsItemYearRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
