package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
type ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetQueryParameters retrieve a list of swapShiftsChangeRequest objects in a team.
type ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetQueryParameters struct {
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
// ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetQueryParameters
}
// ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySwapShiftsChangeRequestId provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) BySwapShiftsChangeRequestId(swapShiftsChangeRequestId string)(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if swapShiftsChangeRequestId != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = swapShiftsChangeRequestId
    }
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal instantiates a new ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    m := &ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/swapShiftsChangeRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder instantiates a new ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder and sets the default values.
func NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsCountRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) Count()(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsCountRequestBuilder) {
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of swapShiftsChangeRequest objects in a team.
// returns a SwapShiftsChangeRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/swapshiftschangerequest-list?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSwapShiftsChangeRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestCollectionResponseable), nil
}
// Post create an instance of a swapShiftsChangeRequest object.
// returns a SwapShiftsChangeRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/swapshiftschangerequest-post?view=graph-rest-1.0
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSwapShiftsChangeRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestable), nil
}
// ToGetRequestInformation retrieve a list of swapShiftsChangeRequest objects in a team.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an instance of a swapShiftsChangeRequest object.
// returns a *RequestInformation when successful
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SwapShiftsChangeRequestable, requestConfiguration *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/swapShiftsChangeRequests", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder when successful
func (m *ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemTeamDefinitionScheduleSwapShiftsChangeRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
