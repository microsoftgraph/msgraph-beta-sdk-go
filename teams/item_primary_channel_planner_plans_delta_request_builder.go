package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPrimaryChannelPlannerPlansDeltaRequestBuilder provides operations to call the delta method.
type ItemPrimaryChannelPlannerPlansDeltaRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetQueryParameters get newly created, updated, or deleted Planner plans in either a group or a Planner roster without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
type ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetQueryParameters struct {
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
// ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetQueryParameters
}
// NewItemPrimaryChannelPlannerPlansDeltaRequestBuilderInternal instantiates a new ItemPrimaryChannelPlannerPlansDeltaRequestBuilder and sets the default values.
func NewItemPrimaryChannelPlannerPlansDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) {
    m := &ItemPrimaryChannelPlannerPlansDeltaRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel/planner/plans/delta(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPrimaryChannelPlannerPlansDeltaRequestBuilder instantiates a new ItemPrimaryChannelPlannerPlansDeltaRequestBuilder and sets the default values.
func NewItemPrimaryChannelPlannerPlansDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimaryChannelPlannerPlansDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get newly created, updated, or deleted Planner plans in either a group or a Planner roster without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// Deprecated: This method is obsolete. Use GetAsDeltaGetResponse instead.
// returns a ItemPrimaryChannelPlannerPlansDeltaResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplan-delta?view=graph-rest-beta
func (m *ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetRequestConfiguration)(ItemPrimaryChannelPlannerPlansDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPrimaryChannelPlannerPlansDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPrimaryChannelPlannerPlansDeltaResponseable), nil
}
// GetAsDeltaGetResponse get newly created, updated, or deleted Planner plans in either a group or a Planner roster without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a ItemPrimaryChannelPlannerPlansDeltaGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplan-delta?view=graph-rest-beta
func (m *ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) GetAsDeltaGetResponse(ctx context.Context, requestConfiguration *ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetRequestConfiguration)(ItemPrimaryChannelPlannerPlansDeltaGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemPrimaryChannelPlannerPlansDeltaGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemPrimaryChannelPlannerPlansDeltaGetResponseable), nil
}
// ToGetRequestInformation get newly created, updated, or deleted Planner plans in either a group or a Planner roster without having to perform a full read of the entire resource collection. For details, see Use delta query to track changes in Microsoft Graph data.
// returns a *RequestInformation when successful
func (m *ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimaryChannelPlannerPlansDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPrimaryChannelPlannerPlansDeltaRequestBuilder when successful
func (m *ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) WithUrl(rawUrl string)(*ItemPrimaryChannelPlannerPlansDeltaRequestBuilder) {
    return NewItemPrimaryChannelPlannerPlansDeltaRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
