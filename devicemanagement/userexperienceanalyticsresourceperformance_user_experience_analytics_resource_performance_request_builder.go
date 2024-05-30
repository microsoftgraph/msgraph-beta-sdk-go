package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters user experience analytics resource performance
type UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsResourcePerformanceId provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) ByUserExperienceAnalyticsResourcePerformanceId(userExperienceAnalyticsResourcePerformanceId string)(*UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsResourcePerformanceId != "" {
        urlTplParams["userExperienceAnalyticsResourcePerformance%2Did"] = userExperienceAnalyticsResourcePerformanceId
    }
    return NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal instantiates a new UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    m := &UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder instantiates a new UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsresourceperformanceCountRequestBuilder when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) Count()(*UserexperienceanalyticsresourceperformanceCountRequestBuilder) {
    return NewUserexperienceanalyticsresourceperformanceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics resource performance
// returns a UserExperienceAnalyticsResourcePerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsResourcePerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsResourcePerformance for deviceManagement
// returns a UserExperienceAnalyticsResourcePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, requestConfiguration *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable), nil
}
// SummarizeDeviceResourcePerformanceWithSummarizeBy provides operations to call the summarizeDeviceResourcePerformance method.
// returns a *UserexperienceanalyticsresourceperformanceSummarizedeviceresourceperformancewithsummarizebySummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) SummarizeDeviceResourcePerformanceWithSummarizeBy(summarizeBy *string)(*UserexperienceanalyticsresourceperformanceSummarizedeviceresourceperformancewithsummarizebySummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    return NewUserexperienceanalyticsresourceperformanceSummarizedeviceresourceperformancewithsummarizebySummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, summarizeBy)
}
// ToGetRequestInformation user experience analytics resource performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsResourcePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, requestConfiguration *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder when successful
func (m *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
