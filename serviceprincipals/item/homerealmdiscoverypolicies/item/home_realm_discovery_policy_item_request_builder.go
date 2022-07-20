package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ibeac3ba6b888f46db18d69c8f17db9630a8f1ab21ebe30efa6e0f37fea42dae0 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/homerealmdiscoverypolicies/item/ref"
)

// HomeRealmDiscoveryPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\homeRealmDiscoveryPolicies\{homeRealmDiscoveryPolicy-id}
type HomeRealmDiscoveryPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPolicyItemRequestBuilder) {
    m := &HomeRealmDiscoveryPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewHomeRealmDiscoveryPolicyItemRequestBuilder instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the Ref property
func (m *HomeRealmDiscoveryPolicyItemRequestBuilder) Ref()(*ibeac3ba6b888f46db18d69c8f17db9630a8f1ab21ebe30efa6e0f37fea42dae0.RefRequestBuilder) {
    return ibeac3ba6b888f46db18d69c8f17db9630a8f1ab21ebe30efa6e0f37fea42dae0.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
