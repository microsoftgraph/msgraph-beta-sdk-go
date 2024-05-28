package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder provides operations to manage the managedAppLogCollectionRequests property of the microsoft.graph.user entity.
type ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters zero or more log collection requests triggered for the user.
type ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters struct {
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
// ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters
}
// ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedAppLogCollectionRequestId provides operations to manage the managedAppLogCollectionRequests property of the microsoft.graph.user entity.
// returns a *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilder when successful
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ByManagedAppLogCollectionRequestId(managedAppLogCollectionRequestId string)(*ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedAppLogCollectionRequestId != "" {
        urlTplParams["managedAppLogCollectionRequest%2Did"] = managedAppLogCollectionRequestId
    }
    return NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal instantiates a new ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder and sets the default values.
func NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    m := &ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedAppLogCollectionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder instantiates a new ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder and sets the default values.
func NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemManagedapplogcollectionrequestsCountRequestBuilder when successful
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Count()(*ItemManagedapplogcollectionrequestsCountRequestBuilder) {
    return NewItemManagedapplogcollectionrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more log collection requests triggered for the user.
// returns a ManagedAppLogCollectionRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppLogCollectionRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestCollectionResponseable), nil
}
// Post create new navigation property to managedAppLogCollectionRequests for users
// returns a ManagedAppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, requestConfiguration *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable), nil
}
// ToGetRequestInformation zero or more log collection requests triggered for the user.
// returns a *RequestInformation when successful
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedAppLogCollectionRequests for users
// returns a *RequestInformation when successful
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, requestConfiguration *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder when successful
func (m *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    return NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
