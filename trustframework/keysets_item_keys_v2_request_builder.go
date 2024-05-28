package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeysetsItemKeys_v2RequestBuilder provides operations to manage the keys_v2 property of the microsoft.graph.trustFrameworkKeySet entity.
type KeysetsItemKeys_v2RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeysetsItemKeys_v2RequestBuilderGetQueryParameters read the properties and relationships of a trustFrameworkKeyv2 object.
type KeysetsItemKeys_v2RequestBuilderGetQueryParameters struct {
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
// KeysetsItemKeys_v2RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeysetsItemKeys_v2RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *KeysetsItemKeys_v2RequestBuilderGetQueryParameters
}
// ByTrustFrameworkKey_v2Kid provides operations to manage the keys_v2 property of the microsoft.graph.trustFrameworkKeySet entity.
// returns a *KeysetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder when successful
func (m *KeysetsItemKeys_v2RequestBuilder) ByTrustFrameworkKey_v2Kid(trustFrameworkKey_v2Kid string)(*KeysetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if trustFrameworkKey_v2Kid != "" {
        urlTplParams["trustFrameworkKey_v2%2Dkid"] = trustFrameworkKey_v2Kid
    }
    return NewKeysetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewKeysetsItemKeys_v2RequestBuilderInternal instantiates a new KeysetsItemKeys_v2RequestBuilder and sets the default values.
func NewKeysetsItemKeys_v2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemKeys_v2RequestBuilder) {
    m := &KeysetsItemKeys_v2RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/keys_v2{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewKeysetsItemKeys_v2RequestBuilder instantiates a new KeysetsItemKeys_v2RequestBuilder and sets the default values.
func NewKeysetsItemKeys_v2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeysetsItemKeys_v2RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeysetsItemKeys_v2RequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *KeysetsItemKeys_v2CountRequestBuilder when successful
func (m *KeysetsItemKeys_v2RequestBuilder) Count()(*KeysetsItemKeys_v2CountRequestBuilder) {
    return NewKeysetsItemKeys_v2CountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a trustFrameworkKeyv2 object.
// returns a TrustFrameworkKey_v2CollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *KeysetsItemKeys_v2RequestBuilder) Get(ctx context.Context, requestConfiguration *KeysetsItemKeys_v2RequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKey_v2CollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKey_v2CollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKey_v2CollectionResponseable), nil
}
// ToGetRequestInformation read the properties and relationships of a trustFrameworkKeyv2 object.
// returns a *RequestInformation when successful
func (m *KeysetsItemKeys_v2RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *KeysetsItemKeys_v2RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *KeysetsItemKeys_v2RequestBuilder when successful
func (m *KeysetsItemKeys_v2RequestBuilder) WithUrl(rawUrl string)(*KeysetsItemKeys_v2RequestBuilder) {
    return NewKeysetsItemKeys_v2RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
