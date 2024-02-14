package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsDailyRequestBuilder provides operations to manage the daily property of the microsoft.graph.userInsightsRoot entity.
type UserInsightsDailyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsDailyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserInsightsDailyRequestBuilderGetQueryParameters summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
type UserInsightsDailyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsDailyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsDailyRequestBuilderGetQueryParameters
}
// UserInsightsDailyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActiveUsers provides operations to manage the activeUsers property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyActiveUsersRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) ActiveUsers()(*UserInsightsDailyActiveUsersRequestBuilder) {
    return NewUserInsightsDailyActiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ActiveUsersBreakdown provides operations to manage the activeUsersBreakdown property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyActiveUsersBreakdownRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) ActiveUsersBreakdown()(*UserInsightsDailyActiveUsersBreakdownRequestBuilder) {
    return NewUserInsightsDailyActiveUsersBreakdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Authentications provides operations to manage the authentications property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyAuthenticationsRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) Authentications()(*UserInsightsDailyAuthenticationsRequestBuilder) {
    return NewUserInsightsDailyAuthenticationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserInsightsDailyRequestBuilderInternal instantiates a new UserInsightsDailyRequestBuilder and sets the default values.
func NewUserInsightsDailyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyRequestBuilder) {
    m := &UserInsightsDailyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsDailyRequestBuilder instantiates a new UserInsightsDailyRequestBuilder and sets the default values.
func NewUserInsightsDailyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsDailyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property daily for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserInsightsDailyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a DailyUserInsightMetricsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsDailyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDailyUserInsightMetricsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable), nil
}
// InactiveUsers provides operations to manage the inactiveUsers property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyInactiveUsersRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) InactiveUsers()(*UserInsightsDailyInactiveUsersRequestBuilder) {
    return NewUserInsightsDailyInactiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InactiveUsersByApplication provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyInactiveUsersByApplicationRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) InactiveUsersByApplication()(*UserInsightsDailyInactiveUsersByApplicationRequestBuilder) {
    return NewUserInsightsDailyInactiveUsersByApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MfaCompletions provides operations to manage the mfaCompletions property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyMfaCompletionsRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) MfaCompletions()(*UserInsightsDailyMfaCompletionsRequestBuilder) {
    return NewUserInsightsDailyMfaCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property daily in reports
// returns a DailyUserInsightMetricsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, requestConfiguration *UserInsightsDailyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDailyUserInsightMetricsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable), nil
}
// SignUps provides operations to manage the signUps property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailySignUpsRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) SignUps()(*UserInsightsDailySignUpsRequestBuilder) {
    return NewUserInsightsDailySignUpsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Summary provides operations to manage the summary property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailySummaryRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) Summary()(*UserInsightsDailySummaryRequestBuilder) {
    return NewUserInsightsDailySummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property daily for reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/reports/userInsights/daily", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserInsightsDailyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property daily in reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, requestConfiguration *UserInsightsDailyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/reports/userInsights/daily", m.BaseRequestBuilder.PathParameters)
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
// UserCount provides operations to manage the userCount property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserInsightsDailyUserCountRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) UserCount()(*UserInsightsDailyUserCountRequestBuilder) {
    return NewUserInsightsDailyUserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserInsightsDailyRequestBuilder when successful
func (m *UserInsightsDailyRequestBuilder) WithUrl(rawUrl string)(*UserInsightsDailyRequestBuilder) {
    return NewUserInsightsDailyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
