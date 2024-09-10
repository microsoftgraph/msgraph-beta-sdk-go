package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApprovalApprovalItemsItemResponsesRequestBuilder provides operations to manage the responses property of the microsoft.graph.approvalItem entity.
type ApprovalApprovalItemsItemResponsesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApprovalApprovalItemsItemResponsesRequestBuilderGetQueryParameters get a list of approvalItemResponse objects.
type ApprovalApprovalItemsItemResponsesRequestBuilderGetQueryParameters struct {
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
// ApprovalApprovalItemsItemResponsesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalApprovalItemsItemResponsesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApprovalApprovalItemsItemResponsesRequestBuilderGetQueryParameters
}
// ApprovalApprovalItemsItemResponsesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApprovalApprovalItemsItemResponsesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApprovalItemResponseId provides operations to manage the responses property of the microsoft.graph.approvalItem entity.
// returns a *ApprovalApprovalItemsItemResponsesApprovalItemResponseItemRequestBuilder when successful
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) ByApprovalItemResponseId(approvalItemResponseId string)(*ApprovalApprovalItemsItemResponsesApprovalItemResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if approvalItemResponseId != "" {
        urlTplParams["approvalItemResponse%2Did"] = approvalItemResponseId
    }
    return NewApprovalApprovalItemsItemResponsesApprovalItemResponseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewApprovalApprovalItemsItemResponsesRequestBuilderInternal instantiates a new ApprovalApprovalItemsItemResponsesRequestBuilder and sets the default values.
func NewApprovalApprovalItemsItemResponsesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalApprovalItemsItemResponsesRequestBuilder) {
    m := &ApprovalApprovalItemsItemResponsesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/approval/approvalItems/{approvalItem%2Did}/responses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewApprovalApprovalItemsItemResponsesRequestBuilder instantiates a new ApprovalApprovalItemsItemResponsesRequestBuilder and sets the default values.
func NewApprovalApprovalItemsItemResponsesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApprovalApprovalItemsItemResponsesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApprovalApprovalItemsItemResponsesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ApprovalApprovalItemsItemResponsesCountRequestBuilder when successful
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) Count()(*ApprovalApprovalItemsItemResponsesCountRequestBuilder) {
    return NewApprovalApprovalItemsItemResponsesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of approvalItemResponse objects.
// returns a ApprovalItemResponseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/approvalitem-list-responses?view=graph-rest-beta
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) Get(ctx context.Context, requestConfiguration *ApprovalApprovalItemsItemResponsesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalItemResponseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseCollectionResponseable), nil
}
// Post create a new approvalItemResponse object.
// returns a ApprovalItemResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/approvalitem-post-responses?view=graph-rest-beta
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseable, requestConfiguration *ApprovalApprovalItemsItemResponsesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalItemResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseable), nil
}
// ToGetRequestInformation get a list of approvalItemResponse objects.
// returns a *RequestInformation when successful
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApprovalApprovalItemsItemResponsesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new approvalItemResponse object.
// returns a *RequestInformation when successful
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalItemResponseable, requestConfiguration *ApprovalApprovalItemsItemResponsesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ApprovalApprovalItemsItemResponsesRequestBuilder when successful
func (m *ApprovalApprovalItemsItemResponsesRequestBuilder) WithUrl(rawUrl string)(*ApprovalApprovalItemsItemResponsesRequestBuilder) {
    return NewApprovalApprovalItemsItemResponsesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
