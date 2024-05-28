package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetQueryParameters user experience analytics device Startup Score History
type UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsScoreHistoryId provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilder when successful
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) ByUserExperienceAnalyticsScoreHistoryId(userExperienceAnalyticsScoreHistoryId string)(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsScoreHistoryId != "" {
        urlTplParams["userExperienceAnalyticsScoreHistory%2Did"] = userExperienceAnalyticsScoreHistoryId
    }
    return NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderInternal instantiates a new UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    m := &UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsScoreHistory{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder instantiates a new UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder and sets the default values.
func NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsscorehistoryCountRequestBuilder when successful
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) Count()(*UserexperienceanalyticsscorehistoryCountRequestBuilder) {
    return NewUserexperienceanalyticsscorehistoryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics device Startup Score History
// returns a UserExperienceAnalyticsScoreHistoryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsScoreHistoryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsScoreHistory for deviceManagement
// returns a UserExperienceAnalyticsScoreHistoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryable, requestConfiguration *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryable), nil
}
// ToGetRequestInformation user experience analytics device Startup Score History
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsScoreHistory for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsScoreHistoryable, requestConfiguration *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder when successful
func (m *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
