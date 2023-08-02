package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
type HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a homeRealmDiscoveryPolicy object.
type HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetQueryParameters
}
// HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppliesTo provides operations to manage the appliesTo property of the microsoft.graph.stsPolicy entity.
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) AppliesTo()(*HomeRealmDiscoveryPoliciesItemAppliesToRequestBuilder) {
    return NewHomeRealmDiscoveryPoliciesItemAppliesToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    m := &HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a homeRealmDiscoveryPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/homerealmdiscoverypolicy-delete?view=graph-rest-1.0
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a homeRealmDiscoveryPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/homerealmdiscoverypolicy-get?view=graph-rest-1.0
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable), nil
}
// Patch update the properties of a homeRealmDiscoveryPolicy object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/homerealmdiscoverypolicy-update?view=graph-rest-1.0
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable), nil
}
// ToDeleteRequestInformation delete a homeRealmDiscoveryPolicy object.
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a homeRealmDiscoveryPolicy object.
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the properties of a homeRealmDiscoveryPolicy object.
func (m *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HomeRealmDiscoveryPolicyable, requestConfiguration *HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
