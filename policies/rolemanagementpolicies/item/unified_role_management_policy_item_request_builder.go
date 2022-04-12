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
// UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetOptions options for Get
type UnifiedRoleManagementPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters represents the role management policies.
type UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreateDeleteRequestInformation delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Delete(options *UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
// Get represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Get(options *UnifiedRoleManagementPolicyItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleManagementPolicyable), nil
}
// Patch update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Patch(options *UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
