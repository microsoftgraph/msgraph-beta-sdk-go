package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i31ccc83d457942971023031ced75254a984d1600fe1a6fe5c585cb6323ff2330 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments/item/subject"
    i3fca530ae915f2fd83ebf84d7cb0fc06dc230a339b01b80965c166a478eb27cd "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments/item/linkedeligibleroleassignment"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5f07357eefb5cbbd64d329aebf69d77092484b21ee353622d11dd8a898b8e64f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments/item/resource"
    ifd9f6ad135ab699731e92d205d00cb1267995251358d5e3a68252147a84b7456 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments/item/roledefinition"
)

// GovernanceRoleAssignmentRequestBuilder builds and executes requests for operations under \governanceRoleAssignments\{governanceRoleAssignment-id}
type GovernanceRoleAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceRoleAssignmentRequestBuilderDeleteOptions options for Delete
type GovernanceRoleAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleAssignmentRequestBuilderGetOptions options for Get
type GovernanceRoleAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceRoleAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleAssignmentRequestBuilderGetQueryParameters get entity from governanceRoleAssignments by key
type GovernanceRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceRoleAssignmentRequestBuilderPatchOptions options for Patch
type GovernanceRoleAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGovernanceRoleAssignmentRequestBuilderInternal instantiates a new GovernanceRoleAssignmentRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestBuilder) {
    m := &GovernanceRoleAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceRoleAssignments/{governanceRoleAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceRoleAssignmentRequestBuilder instantiates a new GovernanceRoleAssignmentRequestBuilder and sets the default values.
func NewGovernanceRoleAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from governanceRoleAssignments
func (m *GovernanceRoleAssignmentRequestBuilder) CreateDeleteRequestInformation(options *GovernanceRoleAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from governanceRoleAssignments by key
func (m *GovernanceRoleAssignmentRequestBuilder) CreateGetRequestInformation(options *GovernanceRoleAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in governanceRoleAssignments
func (m *GovernanceRoleAssignmentRequestBuilder) CreatePatchRequestInformation(options *GovernanceRoleAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from governanceRoleAssignments
func (m *GovernanceRoleAssignmentRequestBuilder) Delete(options *GovernanceRoleAssignmentRequestBuilderDeleteOptions)(error) {
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
// Get get entity from governanceRoleAssignments by key
func (m *GovernanceRoleAssignmentRequestBuilder) Get(options *GovernanceRoleAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceRoleAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignment), nil
}
func (m *GovernanceRoleAssignmentRequestBuilder) LinkedEligibleRoleAssignment()(*i3fca530ae915f2fd83ebf84d7cb0fc06dc230a339b01b80965c166a478eb27cd.LinkedEligibleRoleAssignmentRequestBuilder) {
    return i3fca530ae915f2fd83ebf84d7cb0fc06dc230a339b01b80965c166a478eb27cd.NewLinkedEligibleRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in governanceRoleAssignments
func (m *GovernanceRoleAssignmentRequestBuilder) Patch(options *GovernanceRoleAssignmentRequestBuilderPatchOptions)(error) {
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
func (m *GovernanceRoleAssignmentRequestBuilder) Resource()(*i5f07357eefb5cbbd64d329aebf69d77092484b21ee353622d11dd8a898b8e64f.ResourceRequestBuilder) {
    return i5f07357eefb5cbbd64d329aebf69d77092484b21ee353622d11dd8a898b8e64f.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestBuilder) RoleDefinition()(*ifd9f6ad135ab699731e92d205d00cb1267995251358d5e3a68252147a84b7456.RoleDefinitionRequestBuilder) {
    return ifd9f6ad135ab699731e92d205d00cb1267995251358d5e3a68252147a84b7456.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestBuilder) Subject()(*i31ccc83d457942971023031ced75254a984d1600fe1a6fe5c585cb6323ff2330.SubjectRequestBuilder) {
    return i31ccc83d457942971023031ced75254a984d1600fe1a6fe5c585cb6323ff2330.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
