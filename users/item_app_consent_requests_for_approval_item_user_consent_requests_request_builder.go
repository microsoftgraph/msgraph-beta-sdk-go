package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
type ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetQueryParameters retrieve a collection of userConsentRequest objects and their properties.
type ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetQueryParameters struct {
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
// ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetQueryParameters
}
// ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserConsentRequestId provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
// returns a *ItemAppConsentRequestsForApprovalItemUserConsentRequestsUserConsentRequestItemRequestBuilder when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) ByUserConsentRequestId(userConsentRequestId string)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsUserConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userConsentRequestId != "" {
        urlTplParams["userConsentRequest%2Did"] = userConsentRequestId
    }
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsUserConsentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal instantiates a new ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder and sets the default values.
func NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) {
    m := &ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder instantiates a new ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder and sets the default values.
func NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAppConsentRequestsForApprovalItemUserConsentRequestsCountRequestBuilder when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Count()(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsCountRequestBuilder) {
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *ItemAppConsentRequestsForApprovalItemUserConsentRequestsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsFilterByCurrentUserWithOnRequestBuilder) {
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get retrieve a collection of userConsentRequest objects and their properties.
// returns a UserConsentRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appconsentrequest-list-userconsentrequests?view=graph-rest-1.0
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserConsentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestCollectionResponseable), nil
}
// Post create new navigation property to userConsentRequests for users
// returns a UserConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable), nil
}
// ToGetRequestInformation retrieve a collection of userConsentRequest objects and their properties.
// returns a *RequestInformation when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userConsentRequests for users
// returns a *RequestInformation when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder when successful
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) {
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
