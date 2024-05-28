package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsDailyRequestBuilder provides operations to manage the daily property of the microsoft.graph.userInsightsRoot entity.
type UserinsightsDailyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsDailyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsDailyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserinsightsDailyRequestBuilderGetQueryParameters summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
type UserinsightsDailyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserinsightsDailyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsDailyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsDailyRequestBuilderGetQueryParameters
}
// UserinsightsDailyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsDailyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActiveUsers provides operations to manage the activeUsers property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyActiveusersActiveUsersRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) ActiveUsers()(*UserinsightsDailyActiveusersActiveUsersRequestBuilder) {
    return NewUserinsightsDailyActiveusersActiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Authentications provides operations to manage the authentications property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyAuthenticationsRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) Authentications()(*UserinsightsDailyAuthenticationsRequestBuilder) {
    return NewUserinsightsDailyAuthenticationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserinsightsDailyRequestBuilderInternal instantiates a new UserinsightsDailyRequestBuilder and sets the default values.
func NewUserinsightsDailyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsDailyRequestBuilder) {
    m := &UserinsightsDailyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserinsightsDailyRequestBuilder instantiates a new UserinsightsDailyRequestBuilder and sets the default values.
func NewUserinsightsDailyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsDailyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsDailyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property daily for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsDailyRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserinsightsDailyRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UserinsightsDailyRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsDailyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, error) {
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
// returns a *UserinsightsDailyInactiveusersInactiveUsersRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) InactiveUsers()(*UserinsightsDailyInactiveusersInactiveUsersRequestBuilder) {
    return NewUserinsightsDailyInactiveusersInactiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InactiveUsersByApplication provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) InactiveUsersByApplication()(*UserinsightsDailyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) {
    return NewUserinsightsDailyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MfaCompletions provides operations to manage the mfaCompletions property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) MfaCompletions()(*UserinsightsDailyMfacompletionsMfaCompletionsRequestBuilder) {
    return NewUserinsightsDailyMfacompletionsMfaCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property daily in reports
// returns a DailyUserInsightMetricsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsDailyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, requestConfiguration *UserinsightsDailyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, error) {
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
// returns a *UserinsightsDailySignupsSignUpsRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) SignUps()(*UserinsightsDailySignupsSignUpsRequestBuilder) {
    return NewUserinsightsDailySignupsSignUpsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Summary provides operations to manage the summary property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailySummaryRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) Summary()(*UserinsightsDailySummaryRequestBuilder) {
    return NewUserinsightsDailySummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property daily for reports
// returns a *RequestInformation when successful
func (m *UserinsightsDailyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserinsightsDailyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserinsightsDailyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsDailyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserinsightsDailyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DailyUserInsightMetricsRootable, requestConfiguration *UserinsightsDailyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserCount provides operations to manage the userCount property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
// returns a *UserinsightsDailyUsercountUserCountRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) UserCount()(*UserinsightsDailyUsercountUserCountRequestBuilder) {
    return NewUserinsightsDailyUsercountUserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserinsightsDailyRequestBuilder when successful
func (m *UserinsightsDailyRequestBuilder) WithUrl(rawUrl string)(*UserinsightsDailyRequestBuilder) {
    return NewUserinsightsDailyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
