package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder provides operations to call the triggerDeviceScopeAction method.
type UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderInternal instantiates a new TriggerDeviceScopeActionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) {
    m := &UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceScope/triggerDeviceScopeAction", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder instantiates a new TriggerDeviceScopeActionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action triggerDeviceScopeAction
func (m *UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) Post(ctx context.Context, body UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBodyable, requestConfiguration *UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceScopeActionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceScopeActionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceScopeActionResultable), nil
}
// ToPostRequestInformation invoke action triggerDeviceScopeAction
func (m *UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) ToPostRequestInformation(ctx context.Context, body UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionPostRequestBodyable, requestConfiguration *UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopeTriggerDeviceScopeActionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
