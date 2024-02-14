package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetQueryParameters user experience analytics devices not Windows Autopilot ready.
type UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsNotAutopilotReadyDeviceId provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder when successful
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) ByUserExperienceAnalyticsNotAutopilotReadyDeviceId(userExperienceAnalyticsNotAutopilotReadyDeviceId string)(*UserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsNotAutopilotReadyDeviceId != "" {
        urlTplParams["userExperienceAnalyticsNotAutopilotReadyDevice%2Did"] = userExperienceAnalyticsNotAutopilotReadyDeviceId
    }
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal instantiates a new UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    m := &UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder instantiates a new UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserExperienceAnalyticsNotAutopilotReadyDeviceCountRequestBuilder when successful
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) Count()(*UserExperienceAnalyticsNotAutopilotReadyDeviceCountRequestBuilder) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics devices not Windows Autopilot ready.
// returns a UserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsNotAutopilotReadyDevice for deviceManagement
// returns a UserExperienceAnalyticsNotAutopilotReadyDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, requestConfiguration *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable), nil
}
// ToGetRequestInformation user experience analytics devices not Windows Autopilot ready.
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsNotAutopilotReadyDevice for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, requestConfiguration *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice", m.BaseRequestBuilder.PathParameters)
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
// returns a *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder when successful
func (m *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
