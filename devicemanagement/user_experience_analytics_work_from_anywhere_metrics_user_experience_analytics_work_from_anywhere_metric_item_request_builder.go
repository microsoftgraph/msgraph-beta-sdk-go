package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters user experience analytics work from anywhere metrics.
type UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable), nil
}
// MetricDevices provides operations to manage the metricDevices property of the microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric entity.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) MetricDevices()(*UserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsItemMetricDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
