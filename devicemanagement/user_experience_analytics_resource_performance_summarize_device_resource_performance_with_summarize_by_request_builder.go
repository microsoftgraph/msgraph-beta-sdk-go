package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder provides operations to call the summarizeDeviceResourcePerformance method.
type UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetQueryParameters invoke function summarizeDeviceResourcePerformance
type UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetQueryParameters
}
// NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    m := &UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance/summarizeDeviceResourcePerformance(summarizeBy='{summarizeBy}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if summarizeBy != nil {
        m.BaseRequestBuilder.PathParameters["summarizeBy"] = *summarizeBy
    }
    return m
}
// NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder instantiates a new SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function summarizeDeviceResourcePerformance
// Deprecated: This method is obsolete. Use GetAsSummarizeDeviceResourcePerformanceWithSummarizeByGetResponse instead.
func (m *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByResponseable), nil
}
// GetAsSummarizeDeviceResourcePerformanceWithSummarizeByGetResponse invoke function summarizeDeviceResourcePerformance
func (m *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) GetAsSummarizeDeviceResourcePerformanceWithSummarizeByGetResponse(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByGetResponseable), nil
}
// ToGetRequestInformation invoke function summarizeDeviceResourcePerformance
func (m *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    return NewUserExperienceAnalyticsResourcePerformanceSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
