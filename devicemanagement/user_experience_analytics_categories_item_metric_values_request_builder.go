package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder provides operations to manage the metricValues property of the microsoft.graph.userExperienceAnalyticsCategory entity.
type UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetQueryParameters the metric values for the user experience analytics category. Read-only.
type UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsMetricId provides operations to manage the metricValues property of the microsoft.graph.userExperienceAnalyticsCategory entity.
// returns a *UserExperienceAnalyticsCategoriesItemMetricValuesUserExperienceAnalyticsMetricItemRequestBuilder when successful
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) ByUserExperienceAnalyticsMetricId(userExperienceAnalyticsMetricId string)(*UserExperienceAnalyticsCategoriesItemMetricValuesUserExperienceAnalyticsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsMetricId != "" {
        urlTplParams["userExperienceAnalyticsMetric%2Did"] = userExperienceAnalyticsMetricId
    }
    return NewUserExperienceAnalyticsCategoriesItemMetricValuesUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderInternal instantiates a new UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) {
    m := &UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsCategories/{userExperienceAnalyticsCategory%2Did}/metricValues{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder instantiates a new UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserExperienceAnalyticsCategoriesItemMetricValuesCountRequestBuilder when successful
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) Count()(*UserExperienceAnalyticsCategoriesItemMetricValuesCountRequestBuilder) {
    return NewUserExperienceAnalyticsCategoriesItemMetricValuesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the metric values for the user experience analytics category. Read-only.
// returns a UserExperienceAnalyticsMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricCollectionResponseable), nil
}
// Post create new navigation property to metricValues for deviceManagement
// returns a UserExperienceAnalyticsMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricable, requestConfiguration *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricable), nil
}
// ToGetRequestInformation the metric values for the user experience analytics category. Read-only.
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to metricValues for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsMetricable, requestConfiguration *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/userExperienceAnalyticsCategories/{userExperienceAnalyticsCategory%2Did}/metricValues", m.BaseRequestBuilder.PathParameters)
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
// returns a *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder when successful
func (m *UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder) {
    return NewUserExperienceAnalyticsCategoriesItemMetricValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
