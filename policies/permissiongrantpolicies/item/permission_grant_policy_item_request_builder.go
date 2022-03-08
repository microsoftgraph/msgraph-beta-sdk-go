package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i251a0310ac8ccd28324dd89dabf36e69621defd9d81620901f2d26bb5273ecec "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes"
    i81b08bf532ee9f74501f837d03893fe0b7ffe4476c1dcacfe6302490da5d3694 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes"
    i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/excludes/item"
    i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item/includes/item"
)

// PermissionGrantPolicyItemRequestBuilder provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
type PermissionGrantPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PermissionGrantPolicyItemRequestBuilderDeleteOptions options for Delete
type PermissionGrantPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyItemRequestBuilderGetOptions options for Get
type PermissionGrantPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PermissionGrantPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PermissionGrantPolicyItemRequestBuilderGetQueryParameters the policy that specifies the conditions under which consent can be granted.
type PermissionGrantPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PermissionGrantPolicyItemRequestBuilderPatchOptions options for Patch
type PermissionGrantPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPermissionGrantPolicyItemRequestBuilderInternal instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    m := &PermissionGrantPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/permissionGrantPolicies/{permissionGrantPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantPolicyItemRequestBuilder instantiates a new PermissionGrantPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PermissionGrantPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *PermissionGrantPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PermissionGrantPolicyItemRequestBuilder) CreateGetRequestInformation(options *PermissionGrantPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) CreatePatchRequestInformation(options *PermissionGrantPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property permissionGrantPolicies for policies
func (m *PermissionGrantPolicyItemRequestBuilder) Delete(options *PermissionGrantPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i307e2d11be62be54ec59a357c2601ceddc5616fa089cdfe3e037cd77aafd1e1e.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the policy that specifies the conditions under which consent can be granted.
func (m *PermissionGrantPolicyItemRequestBuilder) Get(options *PermissionGrantPolicyItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePermissionGrantPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PermissionGrantPolicyable), nil
}
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
        urlTplParams["permissionGrantConditionSet_id"] = id
    }
    return i5735f7ded85e85c94dfa2afa819e51282b06351a5a8baed9d71d2b45147a5aef.NewPermissionGrantConditionSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property permissionGrantPolicies in policies
func (m *PermissionGrantPolicyItemRequestBuilder) Patch(options *PermissionGrantPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
