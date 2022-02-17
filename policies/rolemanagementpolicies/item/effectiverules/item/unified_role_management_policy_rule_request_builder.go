package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UnifiedRoleManagementPolicyRuleRequestBuilder builds and executes requests for operations under \policies\roleManagementPolicies\{unifiedRoleManagementPolicy-id}\effectiveRules\{unifiedRoleManagementPolicyRule-id}
type UnifiedRoleManagementPolicyRuleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleManagementPolicyRuleRequestBuilderDeleteOptions options for Delete
type UnifiedRoleManagementPolicyRuleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyRuleRequestBuilderGetOptions options for Get
type UnifiedRoleManagementPolicyRuleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleManagementPolicyRuleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleManagementPolicyRuleRequestBuilderGetQueryParameters the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
type UnifiedRoleManagementPolicyRuleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleManagementPolicyRuleRequestBuilderPatchOptions options for Patch
type UnifiedRoleManagementPolicyRuleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyRule;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUnifiedRoleManagementPolicyRuleRequestBuilderInternal instantiates a new UnifiedRoleManagementPolicyRuleRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyRuleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyRuleRequestBuilder) {
    m := &UnifiedRoleManagementPolicyRuleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/roleManagementPolicies/{unifiedRoleManagementPolicy_id}/effectiveRules/{unifiedRoleManagementPolicyRule_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleManagementPolicyRuleRequestBuilder instantiates a new UnifiedRoleManagementPolicyRuleRequestBuilder and sets the default values.
func NewUnifiedRoleManagementPolicyRuleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleManagementPolicyRuleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleManagementPolicyRuleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleManagementPolicyRuleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleManagementPolicyRuleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleManagementPolicyRuleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) Delete(options *UnifiedRoleManagementPolicyRuleRequestBuilderDeleteOptions)(error) {
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
// Get the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) Get(options *UnifiedRoleManagementPolicyRuleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyRule, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleManagementPolicyRule() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleManagementPolicyRule), nil
}
// Patch the list of effective rules like approval rule, expiration rule, etc. evaluated based on inherited referenced rules. E.g. If there is a tenant wide policy to enforce enabling approval rule, the effective rule will be to enable approval even if the polcy has a rule to disable approval.
func (m *UnifiedRoleManagementPolicyRuleRequestBuilder) Patch(options *UnifiedRoleManagementPolicyRuleRequestBuilderPatchOptions)(error) {
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
