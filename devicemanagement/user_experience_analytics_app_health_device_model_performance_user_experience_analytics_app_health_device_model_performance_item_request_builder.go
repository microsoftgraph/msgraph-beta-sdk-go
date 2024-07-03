package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Model Performance
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthDeviceModelPerformance/{userExperienceAnalyticsAppHealthDeviceModelPerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Model Performance
// returns a UserExperienceAnalyticsAppHealthDeviceModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
// returns a UserExperienceAnalyticsAppHealthDeviceModelPerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Model Performance
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthDeviceModelPerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, requestConfiguration *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder when successful
func (m *UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
