package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageassignments"
    i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageresourcescope"
    i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackagesubject"
    if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageresourcerole"
    ia37a638d7f3d60c9112b0aec719d0a99d0cd2cd187495734627185fc987e9d75 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentresourceroles/item/accesspackageassignments/item"
)

// AccessPackageAssignmentResourceRoleItemRequestBuilder provides operations to manage the accessPackageAssignmentResourceRoles property of the microsoft.graph.accessPackageAssignment entity.
type AccessPackageAssignmentResourceRoleItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderGetOptions options for Get
type AccessPackageAssignmentResourceRoleItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters the resource roles delivered to the target user for this assignment. Read-only. Nullable.
type AccessPackageAssignmentResourceRoleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentResourceRoleItemRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentResourceRoleItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRoleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignments()(*i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc.AccessPackageAssignmentsRequestBuilder) {
    return i0cabe3741c90882c02a7845c85fc54aabd27cb743b143945d2182594f0d1e0bc.NewAccessPackageAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignments.item.accessPackageAssignmentResourceRoles.item.accessPackageAssignments.item collection
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageAssignmentsById(id string)(*ia37a638d7f3d60c9112b0aec719d0a99d0cd2cd187495734627185fc987e9d75.AccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment_id1"] = id
    }
    return ia37a638d7f3d60c9112b0aec719d0a99d0cd2cd187495734627185fc987e9d75.NewAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceRole()(*if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882.AccessPackageResourceRoleRequestBuilder) {
    return if1b7f7cae92d17e491826c7eb7dbfe64885f5e3a436dcba3a8a3e613f5121882.NewAccessPackageResourceRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageResourceScope()(*i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9.AccessPackageResourceScopeRequestBuilder) {
    return i2c32e45c5d65e51fced4cf39c1c3f67f8d4ed0608aab6b75b7d642a007b2f9e9.NewAccessPackageResourceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) AccessPackageSubject()(*i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4.AccessPackageSubjectRequestBuilder) {
    return i6eca23cd75a999dbb5c8e9d0117641b4ff3cecb3979b3e547793fe238df221c4.NewAccessPackageSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    m := &AccessPackageAssignmentResourceRoleItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentResourceRoleItemRequestBuilder instantiates a new AccessPackageAssignmentResourceRoleItemRequestBuilder and sets the default values.
func NewAccessPackageAssignmentResourceRoleItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentResourceRoleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentResourceRoleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentResourceRoleItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentResourceRoleItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property accessPackageAssignmentResourceRoles for identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Delete(options *AccessPackageAssignmentResourceRoleItemRequestBuilderDeleteOptions)(error) {
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
// Get the resource roles delivered to the target user for this assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Get(options *AccessPackageAssignmentResourceRoleItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentResourceRoleable), nil
}
// Patch update the navigation property accessPackageAssignmentResourceRoles in identityGovernance
func (m *AccessPackageAssignmentResourceRoleItemRequestBuilder) Patch(options *AccessPackageAssignmentResourceRoleItemRequestBuilderPatchOptions)(error) {
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
