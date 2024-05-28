package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters the user experience analytics device events entity contains NRT device timeline event details.
type UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDeviceTimelineEventId provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilder when successful
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ByUserExperienceAnalyticsDeviceTimelineEventId(userExperienceAnalyticsDeviceTimelineEventId string)(*UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDeviceTimelineEventId != "" {
        urlTplParams["userExperienceAnalyticsDeviceTimelineEvent%2Did"] = userExperienceAnalyticsDeviceTimelineEventId
    }
    return NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal instantiates a new UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    m := &UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder instantiates a new UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder and sets the default values.
func NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsdevicetimelineeventCountRequestBuilder when successful
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Count()(*UserexperienceanalyticsdevicetimelineeventCountRequestBuilder) {
    return NewUserexperienceanalyticsdevicetimelineeventCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the user experience analytics device events entity contains NRT device timeline event details.
// returns a UserExperienceAnalyticsDeviceTimelineEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceTimelineEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsDeviceTimelineEvent for deviceManagement
// returns a UserExperienceAnalyticsDeviceTimelineEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, requestConfiguration *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceTimelineEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable), nil
}
// ToGetRequestInformation the user experience analytics device events entity contains NRT device timeline event details.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsDeviceTimelineEvent for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, requestConfiguration *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder when successful
func (m *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    return NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
