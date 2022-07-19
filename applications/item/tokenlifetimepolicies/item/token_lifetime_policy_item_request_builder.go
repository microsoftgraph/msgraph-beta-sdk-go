package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i250a9348e72a99486f82fad818d8875d90b681431fe0a7c4e91f7303fdb5e212 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenlifetimepolicies/item/ref"
)

// TokenLifetimePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenLifetimePolicies\{tokenLifetimePolicy-id}
type TokenLifetimePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewTokenLifetimePolicyItemRequestBuilderInternal instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenLifetimePolicyItemRequestBuilder) {
    m := &TokenLifetimePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTokenLifetimePolicyItemRequestBuilder instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the Ref property
func (m *TokenLifetimePolicyItemRequestBuilder) Ref()(*i250a9348e72a99486f82fad818d8875d90b681431fe0a7c4e91f7303fdb5e212.RefRequestBuilder) {
    return i250a9348e72a99486f82fad818d8875d90b681431fe0a7c4e91f7303fdb5e212.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
