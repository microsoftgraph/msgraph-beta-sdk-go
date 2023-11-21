package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device App Impact
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceAppImpact/{userExperienceAnalyticsBatteryHealthDeviceAppImpact%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder instantiates a new UserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsBatteryHealthDeviceAppImpact for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user Experience Analytics Battery Health Device App Impact
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable), nil
}
// Patch update the navigation property userExperienceAnalyticsBatteryHealthDeviceAppImpact in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsBatteryHealthDeviceAppImpact for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user Experience Analytics Battery Health Device App Impact
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsBatteryHealthDeviceAppImpact in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
