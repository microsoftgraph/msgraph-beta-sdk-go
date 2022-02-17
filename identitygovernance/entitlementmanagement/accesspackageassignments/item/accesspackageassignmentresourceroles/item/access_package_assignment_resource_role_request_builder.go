package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageassignments"
    i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageresourcescope"
    i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackagesubject"
    if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageresourcerole"
)

// AccessPackageAssignmentResourceRoleRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentResourceRoles\{accessPackageAssignmentResourceRole-id}
type AccessPackageAssignmentResourceRoleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleRequestBuilderGetOptions options for Get
type AccessPackageAssignmentResourceRoleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters the resource roles delivered to the target user for this assignment. Read-only. Nullable.
type AccessPackageAssignmentResourceRoleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageAssignments()(*i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc.AccessPackageAssignmentsRequestBuilder) {
    return i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceRole()(*if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882.AccessPackageResourceRoleRequestBuilder) {
    return if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882.NewAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceScope()(*i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9.AccessPackageResourceScopeRequestBuilder) {
    return i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9.NewAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageSubject()(*i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4.AccessPackageSubjectRequestBuilder) {
    return i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4.NewAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentResourceRoleRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleRequestBuilder) {
    m := &AccessPackageAssignmentResourceRoleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentResourceRoleRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Delete(options *AccessPackageAssignmentResourceRoleRequestBuilderDeleteOptions)(error) {
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
// Get the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Get(options *AccessPackageAssignmentResourceRoleRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignmentResourceRole() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRole), nil
}
// Patch the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleRequestBuilder) Patch(options *AccessPackageAssignmentResourceRoleRequestBuilderPatchOptions)(error) {
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
