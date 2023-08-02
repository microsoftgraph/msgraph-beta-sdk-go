package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteQueryParameters remove a homeRealmDiscoveryPolicy from a servicePrincipal.
type ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteQueryParameters
}
// NewItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder) {
    m := &ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a homeRealmDiscoveryPolicy from a servicePrincipal.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-delete-homerealmdiscoverypolicies?view=graph-rest-1.0
func (m *ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation remove a homeRealmDiscoveryPolicy from a servicePrincipal.
func (m *ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
