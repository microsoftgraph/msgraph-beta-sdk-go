// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetQueryParameters user experience analytics resource performance
type UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    m := &UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance/{userExperienceAnalyticsResourcePerformance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsResourcePerformance for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics resource performance
// returns a UserExperienceAnalyticsResourcePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsResourcePerformance in deviceManagement
// returns a UserExperienceAnalyticsResourcePerformanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsResourcePerformance for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics resource performance
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsResourcePerformance in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsResourcePerformanceable, requestConfiguration *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder when successful
func (m *UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    return NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
