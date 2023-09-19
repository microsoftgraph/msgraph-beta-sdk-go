package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersion%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by App Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
