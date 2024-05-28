package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetQueryParameters user experience analytics device Startup Process Performance
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDeviceStartupProcessPerformanceId provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) ByUserExperienceAnalyticsDeviceStartupProcessPerformanceId(userExperienceAnalyticsDeviceStartupProcessPerformanceId string)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDeviceStartupProcessPerformanceId != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcessPerformance%2Did"] = userExperienceAnalyticsDeviceStartupProcessPerformanceId
    }
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal instantiates a new UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    m := &UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder instantiates a new UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceCountRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) Count()(*UserexperienceanalyticsdevicestartupprocessperformanceCountRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessperformanceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics device Startup Process Performance
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a UserExperienceAnalyticsDeviceStartupProcessPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// ToGetRequestInformation user experience analytics device Startup Process Performance
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
