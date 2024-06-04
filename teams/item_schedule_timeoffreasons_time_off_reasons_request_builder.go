package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
type ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetQueryParameters the set of reasons for a time off in the schedule.
type ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetQueryParameters struct {
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
// ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetQueryParameters
}
// ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTimeOffReasonId provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
// returns a *ItemScheduleTimeoffreasonsTimeOffReasonItemRequestBuilder when successful
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) ByTimeOffReasonId(timeOffReasonId string)(*ItemScheduleTimeoffreasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if timeOffReasonId != "" {
        urlTplParams["timeOffReason%2Did"] = timeOffReasonId
    }
    return NewItemScheduleTimeoffreasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderInternal instantiates a new ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder and sets the default values.
func NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) {
    m := &ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/schedule/timeOffReasons{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder instantiates a new ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder and sets the default values.
func NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemScheduleTimeoffreasonsCountRequestBuilder when successful
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) Count()(*ItemScheduleTimeoffreasonsCountRequestBuilder) {
    return NewItemScheduleTimeoffreasonsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the set of reasons for a time off in the schedule.
// returns a TimeOffReasonCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/schedule-list-timeoffreasons?view=graph-rest-beta
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeOffReasonCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonCollectionResponseable), nil
}
// Post create a new timeOffReason.
// returns a TimeOffReasonable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/schedule-post-timeoffreasons?view=graph-rest-beta
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, requestConfiguration *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeOffReasonFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable), nil
}
// ToGetRequestInformation the set of reasons for a time off in the schedule.
// returns a *RequestInformation when successful
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new timeOffReason.
// returns a *RequestInformation when successful
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeOffReasonable, requestConfiguration *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder when successful
func (m *ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) WithUrl(rawUrl string)(*ItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) {
    return NewItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
