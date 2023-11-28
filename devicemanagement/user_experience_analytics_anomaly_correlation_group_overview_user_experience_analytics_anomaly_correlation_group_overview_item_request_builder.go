package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    m := &UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/{userExperienceAnalyticsAnomalyCorrelationGroupOverview%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder instantiates a new UserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable), nil
}
// Patch update the navigation property userExperienceAnalyticsAnomalyCorrelationGroupOverview in deviceManagement
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsAnomalyCorrelationGroupOverview in deviceManagement
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
