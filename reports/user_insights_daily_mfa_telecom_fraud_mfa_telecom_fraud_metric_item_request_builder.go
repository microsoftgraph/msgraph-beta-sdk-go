package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder provides operations to manage the mfaTelecomFraud property of the microsoft.graph.dailyUserInsightMetricsRoot entity.
type UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetQueryParameters get mfaTelecomFraud from reports
type UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetQueryParameters
}
// UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderInternal instantiates a new UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder and sets the default values.
func NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) {
    m := &UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/userInsights/daily/mfaTelecomFraud/{mfaTelecomFraudMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder instantiates a new UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder and sets the default values.
func NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mfaTelecomFraud for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get mfaTelecomFraud from reports
// returns a MfaTelecomFraudMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaTelecomFraudMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable), nil
}
// Patch update the navigation property mfaTelecomFraud in reports
// returns a MfaTelecomFraudMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMfaTelecomFraudMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable), nil
}
// ToDeleteRequestInformation delete navigation property mfaTelecomFraud for reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get mfaTelecomFraud from reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mfaTelecomFraud in reports
// returns a *RequestInformation when successful
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MfaTelecomFraudMetricable, requestConfiguration *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder when successful
func (m *UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder) {
    return NewUserInsightsDailyMfaTelecomFraudMfaTelecomFraudMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
