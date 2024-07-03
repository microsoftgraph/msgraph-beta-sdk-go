package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder provides operations to manage the keys_v2 property of the microsoft.graph.trustFrameworkKeySet entity.
type KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetQueryParameters read the properties and relationships of a trustFrameworkKeyv2 object.
type KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetQueryParameters
}
// NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderInternal instantiates a new KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder and sets the default values.
func NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) {
    m := &KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/keys_v2/{trustFrameworkKey_v2%2Dkid}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder instantiates a new KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder and sets the default values.
func NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a trustFrameworkKeyv2 object.
// returns a TrustFrameworkKey_v2able when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkey_v2-get?view=graph-rest-beta
func (m *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) Get(ctx context.Context, requestConfiguration *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKey_v2able, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKey_v2FromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKey_v2able), nil
}
// ToGetRequestInformation read the properties and relationships of a trustFrameworkKeyv2 object.
// returns a *RequestInformation when successful
func (m *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder when successful
func (m *KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) WithUrl(rawUrl string)(*KeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder) {
    return NewKeySetsItemKeys_v2TrustFrameworkKey_v2KItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
