package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetQueryParameters the user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
type UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    m := &UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceScopes/{userExperienceAnalyticsDeviceScope%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder instantiates a new UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsDeviceScopes for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
// returns a UserExperienceAnalyticsDeviceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable), nil
}
// Patch update the navigation property userExperienceAnalyticsDeviceScopes in deviceManagement
// returns a UserExperienceAnalyticsDeviceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsDeviceScopes for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsDeviceScopes in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceScopeable, requestConfiguration *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TriggerDeviceScopeAction provides operations to call the triggerDeviceScopeAction method.
// returns a *UserExperienceAnalyticsDeviceScopesItemTriggerDeviceScopeActionRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) TriggerDeviceScopeAction()(*UserExperienceAnalyticsDeviceScopesItemTriggerDeviceScopeActionRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopesItemTriggerDeviceScopeActionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
