package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version details
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get user experience analytics appHealth Application Performance by App Version details
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation user experience analytics appHealth Application Performance by App Version details
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
