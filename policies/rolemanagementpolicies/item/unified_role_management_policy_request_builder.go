package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules"
    ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules"
    i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/effectiverules/item"
    ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item/rules/item"
)

// UnifiedRoleManagementPolicyRequestBuilder builds and executes requests for operations under \policies\roleManagementPolicies\{unifiedRoleManagementPolicy-id}
type UnifiedRoleManagementPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleManagementPolicyRequestBuilderDeleteOptions options for Delete
type UnifiedRoleManagementPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyRequestBuilderGetOptions options for Get
type UnifiedRoleManagementPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleManagementPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyRequestBuilderGetQueryParameters represents the role management policies.
type UnifiedRoleManagementPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleManagementPolicyRequestBuilderPatchOptions options for Patch
type UnifiedRoleManagementPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnifiedRoleManagementPolicyRequestBuilderInternal instantiates a new UnifiedRoleManagementPolicyRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyRequestBuilder) {
    m := &UnifiedRoleManagementPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleManagementPolicyRequestBuilder instantiates a new UnifiedRoleManagementPolicyRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleManagementPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents the role management policies.
func (m *UnifiedRoleManagementPolicyRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleManagementPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UnifiedRoleManagementPolicyRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleManagementPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents the role management policies.
func (m *UnifiedRoleManagementPolicyRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleManagementPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents the role management policies.
func (m *UnifiedRoleManagementPolicyRequestBuilder) Delete(options *UnifiedRoleManagementPolicyRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleManagementPolicyRequestBuilder) EffectiveRules()(*if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb.EffectiveRulesRequestBuilder) {
    return if0d35103551c602f5b2331b0ffa6028b98d7b4171e7c4db47010b54171dc2dcb.NewEffectiveRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EffectiveRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicies.item.effectiveRules.item collection
func (m *UnifiedRoleManagementPolicyRequestBuilder) EffectiveRulesById(id string)(*i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e.UnifiedRoleManagementPolicyRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule_id"] = id
    }
    return i811ba195cae1e905ab5d23c338b5965c64cfa723b2a7cbaf79f558bcd492389e.NewUnifiedRoleManagementPolicyRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get represents the role management policies.
func (m *UnifiedRoleManagementPolicyRequestBuilder) Get(options *UnifiedRoleManagementPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleManagementPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicy), nil
}
// Patch represents the role management policies.
func (m *UnifiedRoleManagementPolicyRequestBuilder) Patch(options *UnifiedRoleManagementPolicyRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleManagementPolicyRequestBuilder) Rules()(*ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649.RulesRequestBuilder) {
    return ifc39611d357a9b76d0ccfc15869a511b0e0c8f820d08446dd30dd89890511649.NewRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.roleManagementPolicies.item.rules.item collection
func (m *UnifiedRoleManagementPolicyRequestBuilder) RulesById(id string)(*ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7.UnifiedRoleManagementPolicyRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyRule_id"] = id
    }
    return ie226ae78525b483a512018ed5f5bf2023846009012cd463b4e7b22c6a6585ac7.NewUnifiedRoleManagementPolicyRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
