package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationFido2MethodsRequestBuilder provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
type ItemAuthenticationFido2MethodsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationFido2MethodsRequestBuilderGetQueryParameters retrieve a list of a user's FIDO2 Security Key Authentication Method objects and their properties.
type ItemAuthenticationFido2MethodsRequestBuilderGetQueryParameters struct {
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
// ItemAuthenticationFido2MethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationFido2MethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationFido2MethodsRequestBuilderGetQueryParameters
}
// ByFido2AuthenticationMethodId provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationFido2MethodsRequestBuilder) ByFido2AuthenticationMethodId(fido2AuthenticationMethodId string)(*ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fido2AuthenticationMethodId != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = fido2AuthenticationMethodId
    }
    return NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationFido2MethodsRequestBuilderInternal instantiates a new ItemAuthenticationFido2MethodsRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2MethodsRequestBuilder) {
    m := &ItemAuthenticationFido2MethodsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/fido2Methods{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAuthenticationFido2MethodsRequestBuilder instantiates a new ItemAuthenticationFido2MethodsRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2MethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationFido2MethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAuthenticationFido2MethodsCountRequestBuilder when successful
func (m *ItemAuthenticationFido2MethodsRequestBuilder) Count()(*ItemAuthenticationFido2MethodsCountRequestBuilder) {
    return NewItemAuthenticationFido2MethodsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of a user's FIDO2 Security Key Authentication Method objects and their properties.
// returns a Fido2AuthenticationMethodCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/fido2authenticationmethod-list?view=graph-rest-1.0
func (m *ItemAuthenticationFido2MethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Fido2AuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateFido2AuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Fido2AuthenticationMethodCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of a user's FIDO2 Security Key Authentication Method objects and their properties.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationFido2MethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationFido2MethodsRequestBuilder when successful
func (m *ItemAuthenticationFido2MethodsRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationFido2MethodsRequestBuilder) {
    return NewItemAuthenticationFido2MethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
