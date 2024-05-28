package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetQueryParameters user experience analytics device Startup Processes
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDeviceStartupProcessId provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) ByUserExperienceAnalyticsDeviceStartupProcessId(userExperienceAnalyticsDeviceStartupProcessId string)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDeviceStartupProcessId != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcess%2Did"] = userExperienceAnalyticsDeviceStartupProcessId
    }
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal instantiates a new UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    m := &UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder instantiates a new UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsdevicestartupprocessesCountRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) Count()(*UserexperienceanalyticsdevicestartupprocessesCountRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics device Startup Processes
// returns a UserExperienceAnalyticsDeviceStartupProcessCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsDeviceStartupProcesses for deviceManagement
// returns a UserExperienceAnalyticsDeviceStartupProcessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessable), nil
}
// ToGetRequestInformation user experience analytics device Startup Processes
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsDeviceStartupProcesses for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessable, requestConfiguration *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder when successful
func (m *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
