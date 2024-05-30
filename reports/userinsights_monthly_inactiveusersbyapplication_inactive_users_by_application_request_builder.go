package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetQueryParameters get inactiveUsersByApplication from reports
type UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetQueryParameters struct {
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
// UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetQueryParameters
}
// ByMonthlyInactiveUsersByApplicationMetricId provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *UserinsightsMonthlyInactiveusersbyapplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) ByMonthlyInactiveUsersByApplicationMetricId(monthlyInactiveUsersByApplicationMetricId string)(*UserinsightsMonthlyInactiveusersbyapplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if monthlyInactiveUsersByApplicationMetricId != "" {
        urlTplParams["monthlyInactiveUsersByApplicationMetric%2Did"] = monthlyInactiveUsersByApplicationMetricId
    }
    return NewUserinsightsMonthlyInactiveusersbyapplicationMonthlyInactiveUsersByApplicationMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderInternal instantiates a new UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder and sets the default values.
func NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) {
    m := &UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/inactiveUsersByApplication{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder instantiates a new UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder and sets the default values.
func NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) Count()(*UserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilder) {
    return NewUserinsightsMonthlyInactiveusersbyapplicationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a MonthlyInactiveUsersByApplicationMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMonthlyInactiveUsersByApplicationMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyInactiveUsersByApplicationMetricCollectionResponseable), nil
}
// ToGetRequestInformation get inactiveUsersByApplication from reports
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Inactive Users By Application Metric is deprecated and will stop returning data on February 16, 2024. Please use the existing Inactive Users API. as of 2024-02/Remove_Breakdown_APIs
// returns a *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder when successful
func (m *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) WithUrl(rawUrl string)(*UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) {
    return NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
