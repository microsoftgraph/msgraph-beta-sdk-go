package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal instantiates a new UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    m := &UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview/{userExperienceAnalyticsAnomalyCorrelationGroupOverview%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder instantiates a new UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
// returns a UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
