package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
type PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters get servicePrincipalCreationPolicies from policies
type PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetQueryParameters
}
// PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal instantiates a new ServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    m := &PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder instantiates a new ServicePrincipalCreationPolicyItemRequestBuilder and sets the default values.
func NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property servicePrincipalCreationPolicies for policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get servicePrincipalCreationPolicies from policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property servicePrincipalCreationPolicies in policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property servicePrincipalCreationPolicies for policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Excludes provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Excludes()(*PoliciesServicePrincipalCreationPoliciesItemExcludesRequestBuilder) {
    return NewPoliciesServicePrincipalCreationPoliciesItemExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExcludesById provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) ExcludesById(id string)(*PoliciesServicePrincipalCreationPoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet%2Did"] = id
    }
    return NewPoliciesServicePrincipalCreationPoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get servicePrincipalCreationPolicies from policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable), nil
}
// Includes provides operations to manage the includes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Includes()(*PoliciesServicePrincipalCreationPoliciesItemIncludesRequestBuilder) {
    return NewPoliciesServicePrincipalCreationPoliciesItemIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncludesById provides operations to manage the includes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) IncludesById(id string)(*PoliciesServicePrincipalCreationPoliciesItemIncludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet%2Did"] = id
    }
    return NewPoliciesServicePrincipalCreationPoliciesItemIncludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property servicePrincipalCreationPolicies in policies
func (m *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, requestConfiguration *PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationPolicyable), nil
}
