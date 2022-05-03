package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes"
    i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes"
    i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes/item"
    i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes/item"
)

// PermissionGrantPolicyItemRequestBuilder provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
type PermissionGrantPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionGrantPolicyItemRequestBuilderGetQueryParameters the policy that specifies the conditions under which consent can be granted.
type PermissionGrantPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionGrantPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPolicyItemRequestBuilderGetQueryParameters
}
// PermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPolicyItemRequestBuilderInternal instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    m := &PermissionGrantPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantPolicyItemRequestBuilder instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable, requestConfiguration *PermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *PermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *PermissionGrantPolicyItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Excludes the excludes property
func (m *PermissionGrantPolicyItemRequestBuilder) Excludes()(*i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694.ExcludesRequestBuilder) {
    return i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExcludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.permissionGrantPolicies.item.excludes.item collection
func (m *PermissionGrantPolicyItemRequestBuilder) ExcludesById(id string)(*i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e.PermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet%2Did"] = id
    }
    return i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) GetWithResponseHandler(requestConfiguration *PermissionGrantPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) GetWithResponseHandler(requestConfiguration *PermissionGrantPolicyItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionGrantPolicyFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable), nil
}
// Includes the includes property
func (m *PermissionGrantPolicyItemRequestBuilder) Includes()(*i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec.IncludesRequestBuilder) {
    return i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.permissionGrantPolicies.item.includes.item collection
func (m *PermissionGrantPolicyItemRequestBuilder) IncludesById(id string)(*i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef.PermissionGrantConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet%2Did"] = id
    }
    return i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable, requestConfiguration *PermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPolicyable, requestConfiguration *PermissionGrantPolicyItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
