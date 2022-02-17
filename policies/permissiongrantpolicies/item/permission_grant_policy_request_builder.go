package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes"
    i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes"
    i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes/item"
    i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes/item"
)

// PermissionGrantPolicyRequestBuilder builds and executes requests for operations under \policies\permissionGrantPolicies\{permissionGrantPolicy-id}
type PermissionGrantPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PermissionGrantPolicyRequestBuilderDeleteOptions options for Delete
type PermissionGrantPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyRequestBuilderGetOptions options for Get
type PermissionGrantPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PermissionGrantPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyRequestBuilderGetQueryParameters the policy that specifies the conditions under which consent can be granted.
type PermissionGrantPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PermissionGrantPolicyRequestBuilderPatchOptions options for Patch
type PermissionGrantPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPermissionGrantPolicyRequestBuilderInternal instantiates a new PermissionGrantPolicyRequestBuilder and sets the default values.
func NewPermissionGrantPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyRequestBuilder) {
    m := &PermissionGrantPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantPolicyRequestBuilder instantiates a new PermissionGrantPolicyRequestBuilder and sets the default values.
func NewPermissionGrantPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) CreateDeleteRequestInformation(options *PermissionGrantPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) CreateGetRequestInformation(options *PermissionGrantPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) CreatePatchRequestInformation(options *PermissionGrantPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) Delete(options *PermissionGrantPolicyRequestBuilderDeleteOptions)(error) {
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
func (m *PermissionGrantPolicyRequestBuilder) Excludes()(*i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694.ExcludesRequestBuilder) {
    return i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExcludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.permissionGrantPolicies.item.excludes.item collection
func (m *PermissionGrantPolicyRequestBuilder) ExcludesById(id string)(*i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e.PermissionGrantConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e.NewPermissionGrantConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) Get(options *PermissionGrantPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPermissionGrantPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicy), nil
}
func (m *PermissionGrantPolicyRequestBuilder) Includes()(*i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec.IncludesRequestBuilder) {
    return i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncludesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.policies.permissionGrantPolicies.item.includes.item collection
func (m *PermissionGrantPolicyRequestBuilder) IncludesById(id string)(*i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef.PermissionGrantConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef.NewPermissionGrantConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyRequestBuilder) Patch(options *PermissionGrantPolicyRequestBuilderPatchOptions)(error) {
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
