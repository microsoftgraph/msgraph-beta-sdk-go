package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
type ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetQueryParameters the logical grouping of users in the schedule (usually by role).
type ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetQueryParameters struct {
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
// ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetQueryParameters
}
// ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySchedulingGroupId provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
// returns a *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupItemRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) BySchedulingGroupId(schedulingGroupId string)(*ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if schedulingGroupId != "" {
        urlTplParams["schedulingGroup%2Did"] = schedulingGroupId
    }
    return NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderInternal instantiates a new ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) {
    m := &ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/schedule/schedulingGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder instantiates a new ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder and sets the default values.
func NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamdefinitionScheduleSchedulinggroupsCountRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) Count()(*ItemTeamdefinitionScheduleSchedulinggroupsCountRequestBuilder) {
    return NewItemTeamdefinitionScheduleSchedulinggroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the logical grouping of users in the schedule (usually by role).
// returns a SchedulingGroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSchedulingGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupCollectionResponseable), nil
}
// Post create new navigation property to schedulingGroups for teamTemplateDefinition
// returns a SchedulingGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupable, requestConfiguration *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSchedulingGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupable), nil
}
// ToGetRequestInformation the logical grouping of users in the schedule (usually by role).
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to schedulingGroups for teamTemplateDefinition
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SchedulingGroupable, requestConfiguration *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder when successful
func (m *ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) {
    return NewItemTeamdefinitionScheduleSchedulinggroupsSchedulingGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
