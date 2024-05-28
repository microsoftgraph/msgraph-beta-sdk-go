package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
type ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetQueryParameters get appConsentRequestsForApproval from users
type ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetQueryParameters struct {
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
// ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetQueryParameters
}
// ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAppConsentRequestId provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
// returns a *ItemAppconsentrequestsforapprovalAppConsentRequestItemRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) ByAppConsentRequestId(appConsentRequestId string)(*ItemAppconsentrequestsforapprovalAppConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appConsentRequestId != "" {
        urlTplParams["appConsentRequest%2Did"] = appConsentRequestId
    }
    return NewItemAppconsentrequestsforapprovalAppConsentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderInternal instantiates a new ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) {
    m := &ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder instantiates a new ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAppconsentrequestsforapprovalCountRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) Count()(*ItemAppconsentrequestsforapprovalCountRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *ItemAppconsentrequestsforapprovalFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ItemAppconsentrequestsforapprovalFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get get appConsentRequestsForApproval from users
// returns a AppConsentRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppConsentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestCollectionResponseable), nil
}
// Post create new navigation property to appConsentRequestsForApproval for users
// returns a AppConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestable, requestConfiguration *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestable), nil
}
// ToGetRequestInformation get appConsentRequestsForApproval from users
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to appConsentRequestsForApproval for users
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppConsentRequestable, requestConfiguration *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) WithUrl(rawUrl string)(*ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
