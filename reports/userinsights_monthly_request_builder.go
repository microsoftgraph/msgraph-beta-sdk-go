package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserinsightsMonthlyRequestBuilder provides operations to manage the monthly property of the microsoft.graph.userInsightsRoot entity.
type UserinsightsMonthlyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserinsightsMonthlyRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlyRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserinsightsMonthlyRequestBuilderGetQueryParameters summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
type UserinsightsMonthlyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserinsightsMonthlyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserinsightsMonthlyRequestBuilderGetQueryParameters
}
// UserinsightsMonthlyRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserinsightsMonthlyRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActiveUsers provides operations to manage the activeUsers property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyActiveusersActiveUsersRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) ActiveUsers()(*UserinsightsMonthlyActiveusersActiveUsersRequestBuilder) {
    return NewUserinsightsMonthlyActiveusersActiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Authentications provides operations to manage the authentications property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyAuthenticationsRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) Authentications()(*UserinsightsMonthlyAuthenticationsRequestBuilder) {
    return NewUserinsightsMonthlyAuthenticationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserinsightsMonthlyRequestBuilderInternal instantiates a new UserinsightsMonthlyRequestBuilder and sets the default values.
func NewUserinsightsMonthlyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyRequestBuilder) {
    m := &UserinsightsMonthlyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserinsightsMonthlyRequestBuilder instantiates a new UserinsightsMonthlyRequestBuilder and sets the default values.
func NewUserinsightsMonthlyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserinsightsMonthlyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserinsightsMonthlyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property monthly for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlyRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserinsightsMonthlyRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a MonthlyUserInsightMetricsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlyRequestBuilder) Get(ctx context.Context, requestConfiguration *UserinsightsMonthlyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMonthlyUserInsightMetricsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable), nil
}
// InactiveUsers provides operations to manage the inactiveUsers property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyInactiveusersInactiveUsersRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) InactiveUsers()(*UserinsightsMonthlyInactiveusersInactiveUsersRequestBuilder) {
    return NewUserinsightsMonthlyInactiveusersInactiveUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InactiveUsersByApplication provides operations to manage the inactiveUsersByApplication property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) InactiveUsersByApplication()(*UserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilder) {
    return NewUserinsightsMonthlyInactiveusersbyapplicationInactiveUsersByApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MfaCompletions provides operations to manage the mfaCompletions property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyMfacompletionsMfaCompletionsRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) MfaCompletions()(*UserinsightsMonthlyMfacompletionsMfaCompletionsRequestBuilder) {
    return NewUserinsightsMonthlyMfacompletionsMfaCompletionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property monthly in reports
// returns a MonthlyUserInsightMetricsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserinsightsMonthlyRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable, requestConfiguration *UserinsightsMonthlyRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMonthlyUserInsightMetricsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable), nil
}
// Requests provides operations to manage the requests property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlyRequestsRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) Requests()(*UserinsightsMonthlyRequestsRequestBuilder) {
    return NewUserinsightsMonthlyRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignUps provides operations to manage the signUps property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlySignupsSignUpsRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) SignUps()(*UserinsightsMonthlySignupsSignUpsRequestBuilder) {
    return NewUserinsightsMonthlySignupsSignUpsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Summary provides operations to manage the summary property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
// returns a *UserinsightsMonthlySummaryRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) Summary()(*UserinsightsMonthlySummaryRequestBuilder) {
    return NewUserinsightsMonthlySummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property monthly for reports
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlyRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserinsightsMonthlyRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra External ID for customers.
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserinsightsMonthlyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property monthly in reports
// returns a *RequestInformation when successful
func (m *UserinsightsMonthlyRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MonthlyUserInsightMetricsRootable, requestConfiguration *UserinsightsMonthlyRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserinsightsMonthlyRequestBuilder when successful
func (m *UserinsightsMonthlyRequestBuilder) WithUrl(rawUrl string)(*UserinsightsMonthlyRequestBuilder) {
    return NewUserinsightsMonthlyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
