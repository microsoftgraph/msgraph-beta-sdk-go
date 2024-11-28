package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyMfaRegisteredUsersRequestBuilder provides operations to manage the mfaRegisteredUsers property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyMfaRegisteredUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetQueryParameters get mfaRegisteredUsers from reports
type UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetQueryParameters struct {
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
// UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetQueryParameters
}
// UserInsightsMonthlyMfaRegisteredUsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaRegisteredUsersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMfaUserCountMetricId provides operations to manage the mfaRegisteredUsers property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) ByMfaUserCountMetricId(mfaUserCountMetricId string)(*UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mfaUserCountMetricId != "" {
        urlTplParams["mfaUserCountMetric%2Did"] = mfaUserCountMetricId
    }
    return NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilderInternal instantiates a new UserInsightsMonthlyMfaRegisteredUsersRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) {
    m := &UserInsightsMonthlyMfaRegisteredUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/mfaRegisteredUsers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilder instantiates a new UserInsightsMonthlyMfaRegisteredUsersRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserInsightsMonthlyMfaRegisteredUsersCountRequestBuilder when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) Count()(*UserInsightsMonthlyMfaRegisteredUsersCountRequestBuilder) {
    return NewUserInsightsMonthlyMfaRegisteredUsersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get mfaRegisteredUsers from reports
// returns a MfaUserCountMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaUserCountMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricCollectionResponseable), nil
}
// Post create new navigation property to mfaRegisteredUsers for reports
// returns a MfaUserCountMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaUserCountMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable), nil
}
// ToGetRequestInformation get mfaRegisteredUsers from reports
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mfaRegisteredUsers for reports
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyMfaRegisteredUsersRequestBuilder) {
    return NewUserInsightsMonthlyMfaRegisteredUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
