package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules"
    ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules"
    i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules/item"
    ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules/item"
)

// UnifiedRoleManagementPolicyItemRequestBuilder provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
type UnifiedRoleManagementPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters represents the role management policies.
type UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters
}
// UnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUnifiedRoleManagementPolicyItemRequestBuilderInternal instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleManagementPolicyItemRequestBuilder) {
    m := &UnifiedRoleManagementPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleManagementPolicyItemRequestBuilder instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRoleManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// EffectiveRules the effectiveRules property
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) EffectiveRules()(*if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb.EffectiveRulesRequestBuilder) {
    return if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb.NewEffectiveRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EffectiveRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicies.item.effectiveRules.item collection
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) EffectiveRulesById(id string)(*i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e.UnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = id
    }
    return i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e.NewUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable), nil
}
// PatchWithResponseHandler update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, requestConfiguration *UnifiedRoleManagementPolicyItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Rules the rules property
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Rules()(*ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649.RulesRequestBuilder) {
    return ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649.NewRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicies.item.rules.item collection
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) RulesById(id string)(*ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7.UnifiedRoleManagementPolicyRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule%2Did"] = id
    }
    return ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7.NewUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
