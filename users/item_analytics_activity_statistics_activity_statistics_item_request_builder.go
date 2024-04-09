package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder provides operations to manage the activityStatistics property of the microsoft.graph.userAnalytics entity.
type ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetQueryParameters the collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
type ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetQueryParameters
}
// ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderInternal instantiates a new ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder and sets the default values.
func NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) {
    m := &ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/analytics/activityStatistics/{activityStatistics%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder instantiates a new ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder and sets the default values.
func NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property activityStatistics for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
// returns a ActivityStatisticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityStatisticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable), nil
}
// Patch update the navigation property activityStatistics in users
// returns a ActivityStatisticsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActivityStatisticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable), nil
}
// ToDeleteRequestInformation delete navigation property activityStatistics for users
// returns a *RequestInformation when successful
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of work activities that a user spent time on during and outside of working hours. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property activityStatistics in users
// returns a *RequestInformation when successful
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActivityStatisticsable, requestConfiguration *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder when successful
func (m *ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) WithUrl(rawUrl string)(*ItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder) {
    return NewItemAnalyticsActivityStatisticsActivityStatisticsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
