package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetQueryParameters user experience analytics impacting process
type UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsImpactingProcessId provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessItemRequestBuilder when successful
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) ByUserExperienceAnalyticsImpactingProcessId(userExperienceAnalyticsImpactingProcessId string)(*UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsImpactingProcessId != "" {
        urlTplParams["userExperienceAnalyticsImpactingProcess%2Did"] = userExperienceAnalyticsImpactingProcessId
    }
    return NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderInternal instantiates a new UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder and sets the default values.
func NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) {
    m := &UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsImpactingProcess{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder instantiates a new UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder and sets the default values.
func NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsimpactingprocessCountRequestBuilder when successful
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) Count()(*UserexperienceanalyticsimpactingprocessCountRequestBuilder) {
    return NewUserexperienceanalyticsimpactingprocessCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics impacting process
// returns a UserExperienceAnalyticsImpactingProcessCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsImpactingProcessCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsImpactingProcess for deviceManagement
// returns a UserExperienceAnalyticsImpactingProcessable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessable, requestConfiguration *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessable), nil
}
// ToGetRequestInformation user experience analytics impacting process
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsImpactingProcess for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsImpactingProcessable, requestConfiguration *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder when successful
func (m *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
