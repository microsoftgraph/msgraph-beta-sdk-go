package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules"
    ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules"
    i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules/item"
    ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules/item"
)

// UnifiedRoleManagementPolicyItemRequestBuilder provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
type UnifiedRoleManagementPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetOptions options for Get
type UnifiedRoleManagementPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters represents the role management policies.
type UnifiedRoleManagementPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnifiedRoleManagementPolicyItemRequestBuilderInternal instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyItemRequestBuilder) {
    m := &UnifiedRoleManagementPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleManagementPolicyItemRequestBuilder instantiates a new UnifiedRoleManagementPolicyItemRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleManagementPolicies for policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["unifiedRoleManagementPolicyRule_id"] = id
    }
    return i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e.NewUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents the role management policies.
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Get(options *UnifiedRoleManagementPolicyItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleManagementPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyable), nil
}
// Patch update the navigation property roleManagementPolicies in policies
func (m *UnifiedRoleManagementPolicyItemRequestBuilder) Patch(options *UnifiedRoleManagementPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["unifiedRoleManagementPolicyRule_id"] = id
    }
    return ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7.NewUnifiedRoleManagementPolicyRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
