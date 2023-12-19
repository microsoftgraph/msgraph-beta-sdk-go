package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataInboundFlowsInboundFlowItemRequestBuilder provides operations to manage the inboundFlows property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustryDataInboundFlowsInboundFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataInboundFlowsInboundFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataInboundFlowsInboundFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetQueryParameters read the properties and relationships of an inboundFileFlow object.
type IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetQueryParameters
}
// IndustryDataInboundFlowsInboundFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataInboundFlowsInboundFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataInboundFlowsInboundFlowItemRequestBuilderInternal instantiates a new InboundFlowItemRequestBuilder and sets the default values.
func NewIndustryDataInboundFlowsInboundFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataInboundFlowsInboundFlowItemRequestBuilder) {
    m := &IndustryDataInboundFlowsInboundFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/inboundFlows/{inboundFlow%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewIndustryDataInboundFlowsInboundFlowItemRequestBuilder instantiates a new InboundFlowItemRequestBuilder and sets the default values.
func NewIndustryDataInboundFlowsInboundFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataInboundFlowsInboundFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataInboundFlowsInboundFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataConnector provides operations to manage the dataConnector property of the microsoft.graph.industryData.inboundFlow entity.
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) DataConnector()(*IndustryDataInboundFlowsItemDataConnectorRequestBuilder) {
    return NewIndustryDataInboundFlowsItemDataConnectorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an inboundFileFlow object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundfileflow-delete?view=graph-rest-1.0
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an inboundFileFlow object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundfileflow-get?view=graph-rest-1.0
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the properties of an inboundFileFlow object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-inboundfileflow-update?view=graph-rest-1.0
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) Patch(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderPatchRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete an inboundFileFlow object.
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an inboundFileFlow object.
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an inboundFileFlow object.
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.InboundFlowable, requestConfiguration *IndustryDataInboundFlowsInboundFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) WithUrl(rawUrl string)(*IndustryDataInboundFlowsInboundFlowItemRequestBuilder) {
    return NewIndustryDataInboundFlowsInboundFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Year provides operations to manage the year property of the microsoft.graph.industryData.inboundFlow entity.
func (m *IndustryDataInboundFlowsInboundFlowItemRequestBuilder) Year()(*IndustryDataInboundFlowsItemYearRequestBuilder) {
    return NewIndustryDataInboundFlowsItemYearRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
