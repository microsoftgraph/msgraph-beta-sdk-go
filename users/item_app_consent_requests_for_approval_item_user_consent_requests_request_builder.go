package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
type ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal instantiates a new UserConsentRequestsRequestBuilder and sets the default values.
func NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) {
    m := &ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder instantiates a new UserConsentRequestsRequestBuilder and sets the default values.
func NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Count()(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsCountRequestBuilder) {
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get retrieve a collection of userConsentRequest objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/appconsentrequest-list-userconsentrequests?view=graph-rest-1.0
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserConsentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestCollectionResponseable), nil
}
// MicrosoftGraphFilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) MicrosoftGraphFilterByCurrentUserWithOn(on *string)(*ItemAppConsentRequestsForApprovalItemUserConsentRequestsMicrosoftGraphFilterByCurrentUserWithOnRequestBuilder) {
    return NewItemAppConsentRequestsForApprovalItemUserConsentRequestsMicrosoftGraphFilterByCurrentUserWithOnRequestBuilderInternal(m.pathParameters, m.requestAdapter, on)
}
// Post create new navigation property to userConsentRequests for users
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable), nil
}
// ToGetRequestInformation retrieve a collection of userConsentRequest objects and their properties.
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPostRequestInformation create new navigation property to userConsentRequests for users
func (m *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserConsentRequestable, requestConfiguration *ItemAppConsentRequestsForApprovalItemUserConsentRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
