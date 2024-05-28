package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetQueryParameters user experience analytics devices not Windows Autopilot ready.
type UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal instantiates a new UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    m := &UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsNotAutopilotReadyDevice/{userExperienceAnalyticsNotAutopilotReadyDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder instantiates a new UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsNotAutopilotReadyDevice for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics devices not Windows Autopilot ready.
// returns a UserExperienceAnalyticsNotAutopilotReadyDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property userExperienceAnalyticsNotAutopilotReadyDevice in deviceManagement
// returns a UserExperienceAnalyticsNotAutopilotReadyDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsNotAutopilotReadyDevice for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics devices not Windows Autopilot ready.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsNotAutopilotReadyDevice in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsNotAutopilotReadyDeviceable, requestConfiguration *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder when successful
func (m *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    return NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
