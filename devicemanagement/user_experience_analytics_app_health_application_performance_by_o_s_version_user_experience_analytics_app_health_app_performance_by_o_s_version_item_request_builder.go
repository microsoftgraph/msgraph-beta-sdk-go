package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by OS Version
type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion/{userExperienceAnalyticsAppHealthAppPerformanceByOSVersion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by OS Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance by OS Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
