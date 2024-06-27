package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
type ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters get servicePrincipalCreationPolicies from policies
type ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters
}
// ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal instantiates a new ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    m := &ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder instantiates a new ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property servicePrincipalCreationPolicies for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Excludes provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
// returns a *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Excludes()(*ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemExcludesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get servicePrincipalCreationPolicies from policies
// returns a ServicePrincipalCreationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Includes provides operations to manage the includes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
// returns a *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Includes()(*ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemIncludesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property servicePrincipalCreationPolicies in policies
// returns a ServicePrincipalCreationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property servicePrincipalCreationPolicies for policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get servicePrincipalCreationPolicies from policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property servicePrincipalCreationPolicies in policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    return NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
