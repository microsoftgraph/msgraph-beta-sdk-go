package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsDeviceTimelineEventRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsDeviceTimelineEventRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters the user experience analytics device events entity contains NRT device timeline event details.
type UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsDeviceTimelineEventId provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceTimelineEventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ByUserExperienceAnalyticsDeviceTimelineEventId(userExperienceAnalyticsDeviceTimelineEventId string)(*UserExperienceAnalyticsDeviceTimelineEventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsDeviceTimelineEventId != "" {
        urlTplParams["userExperienceAnalyticsDeviceTimelineEvent%2Did"] = userExperienceAnalyticsDeviceTimelineEventId
    }
    return NewUserExperienceAnalyticsDeviceTimelineEventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal instantiates a new UserExperienceAnalyticsDeviceTimelineEventRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    m := &UserExperienceAnalyticsDeviceTimelineEventRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilder instantiates a new UserExperienceAnalyticsDeviceTimelineEventRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserExperienceAnalyticsDeviceTimelineEventCountRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Count()(*UserExperienceAnalyticsDeviceTimelineEventCountRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceTimelineEventCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the user experience analytics device events entity contains NRT device timeline event details.
// returns a UserExperienceAnalyticsDeviceTimelineEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventCollectionResponseable, error) {
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
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, requestConfiguration *UserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, error) {
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
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDeviceTimelineEventRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceTimelineEventable, requestConfiguration *UserExperienceAnalyticsDeviceTimelineEventRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceTimelineEvent", m.BaseRequestBuilder.PathParameters)
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
// returns a *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder when successful
func (m *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
