package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
type ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetQueryParameters get servicePrincipalCreationPolicies from policies
type ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetQueryParameters struct {
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
// ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetQueryParameters
}
// ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByServicePrincipalCreationPolicyId provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
// returns a *ServiceprincipalcreationpoliciesServicePrincipalCreationPolicyItemRequestBuilder when successful
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) ByServicePrincipalCreationPolicyId(servicePrincipalCreationPolicyId string)(*ServiceprincipalcreationpoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if servicePrincipalCreationPolicyId != "" {
        urlTplParams["servicePrincipalCreationPolicy%2Did"] = servicePrincipalCreationPolicyId
    }
    return NewServiceprincipalcreationpoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderInternal instantiates a new ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder and sets the default values.
func NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) {
    m := &ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/servicePrincipalCreationPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder instantiates a new ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder and sets the default values.
func NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServiceprincipalcreationpoliciesCountRequestBuilder when successful
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) Count()(*ServiceprincipalcreationpoliciesCountRequestBuilder) {
    return NewServiceprincipalcreationpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get servicePrincipalCreationPolicies from policies
// returns a ServicePrincipalCreationPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyCollectionResponseable), nil
}
// Post create new navigation property to servicePrincipalCreationPolicies for policies
// returns a ServicePrincipalCreationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable), nil
}
// ToGetRequestInformation get servicePrincipalCreationPolicies from policies
// returns a *RequestInformation when successful
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to servicePrincipalCreationPolicies for policies
// returns a *RequestInformation when successful
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder when successful
func (m *ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) WithUrl(rawUrl string)(*ServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder) {
    return NewServiceprincipalcreationpoliciesServicePrincipalCreationPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
