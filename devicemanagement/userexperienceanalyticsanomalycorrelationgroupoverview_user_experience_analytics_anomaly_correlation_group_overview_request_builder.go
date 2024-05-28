package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetQueryParameters the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) ByUserExperienceAnalyticsAnomalyCorrelationGroupOverviewId(userExperienceAnalyticsAnomalyCorrelationGroupOverviewId string)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsAnomalyCorrelationGroupOverviewId != "" {
        urlTplParams["userExperienceAnalyticsAnomalyCorrelationGroupOverview%2Did"] = userExperienceAnalyticsAnomalyCorrelationGroupOverviewId
    }
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderInternal instantiates a new UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) {
    m := &UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder instantiates a new UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder and sets the default values.
func NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsanomalycorrelationgroupoverviewCountRequestBuilder when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) Count()(*UserexperienceanalyticsanomalycorrelationgroupoverviewCountRequestBuilder) {
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
// returns a UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
// returns a UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsAnomalyCorrelationGroupOverview for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, requestConfiguration *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder when successful
func (m *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) {
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
