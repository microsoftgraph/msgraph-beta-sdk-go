package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder provides operations to manage the mfaRegisteredUsers property of the microsoft.graph.monthlyUserInsightMetricsRoot entity.
type UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetQueryParameters get mfaRegisteredUsers from reports
type UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetQueryParameters
}
// UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderInternal instantiates a new UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) {
    m := &UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/monthly/mfaRegisteredUsers/{mfaUserCountMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder instantiates a new UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder and sets the default values.
func NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mfaRegisteredUsers for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get mfaRegisteredUsers from reports
// returns a MfaUserCountMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property mfaRegisteredUsers in reports
// returns a MfaUserCountMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property mfaRegisteredUsers for reports
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get mfaRegisteredUsers from reports
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mfaRegisteredUsers in reports
// returns a *RequestInformation when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaUserCountMetricable, requestConfiguration *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder when successful
func (m *UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder) {
    return NewUserInsightsMonthlyMfaRegisteredUsersMfaUserCountMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
