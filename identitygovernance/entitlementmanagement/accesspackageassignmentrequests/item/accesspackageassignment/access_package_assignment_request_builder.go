package accesspackageassignment

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/reprocess"
    i13f41d0331cc43f6c13d7c89468e8cb6cf6311138579b9a58c63193c5f4a6eb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentrequests"
    i285c876d5a8cc7e7023c1b2fc27c8fb12db6522a9b8e69ebf43628adcad931b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentresourceroles"
    i5db2fe656380830458e049ff8fdf8c92a87b50cd763afece28481062b04f6f87 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentpolicy"
    icf5a215bf3787be000e186a6d751b896aaac3ce17b742665321dbc2b091d4954 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/target"
    ifec30412b3c200b86de4be263e81197960c12823efc7785fbd360bbb6be98c84 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackage"
    i2e7e1a5d943fabb7337bc89fb080c5087bc61ffc498b45630a0ff7889e031424 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentrequests/item"
    if55ecae688a10dbb886807245aafc677f67cea3c48410b03a702f3b53ea391ee "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/accesspackageassignmentresourceroles/item"
)

// AccessPackageAssignmentRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}\accessPackageAssignment
type AccessPackageAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestBuilderGetOptions options for Get
type AccessPackageAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestBuilderGetQueryParameters for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
type AccessPackageAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// AccessPackageAssignmentRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackage()(*ifec30412b3c200b86de4be263e81197960c12823efc7785fbd360bbb6be98c84.AccessPackageRequestBuilder) {
    return ifec30412b3c200b86de4be263e81197960c12823efc7785fbd360bbb6be98c84.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentPolicy()(*i5db2fe656380830458e049ff8fdf8c92a87b50cd763afece28481062b04f6f87.AccessPackageAssignmentPolicyRequestBuilder) {
    return i5db2fe656380830458e049ff8fdf8c92a87b50cd763afece28481062b04f6f87.NewAccessPackageAssignmentPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentRequests()(*i13f41d0331cc43f6c13d7c89468e8cb6cf6311138579b9a58c63193c5f4a6eb8.AccessPackageAssignmentRequestsRequestBuilder) {
    return i13f41d0331cc43f6c13d7c89468e8cb6cf6311138579b9a58c63193c5f4a6eb8.NewAccessPackageAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentRequests.item collection
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentRequestsById(id string)(*i2e7e1a5d943fabb7337bc89fb080c5087bc61ffc498b45630a0ff7889e031424.AccessPackageAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest_id1"] = id
    }
    return i2e7e1a5d943fabb7337bc89fb080c5087bc61ffc498b45630a0ff7889e031424.NewAccessPackageAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentResourceRoles()(*i285c876d5a8cc7e7023c1b2fc27c8fb12db6522a9b8e69ebf43628adcad931b7.AccessPackageAssignmentResourceRolesRequestBuilder) {
    return i285c876d5a8cc7e7023c1b2fc27c8fb12db6522a9b8e69ebf43628adcad931b7.NewAccessPackageAssignmentResourceRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentResourceRolesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.entitlementManagement.accessPackageAssignmentRequests.item.accessPackageAssignment.accessPackageAssignmentResourceRoles.item collection
func (m *AccessPackageAssignmentRequestBuilder) AccessPackageAssignmentResourceRolesById(id string)(*if55ecae688a10dbb886807245aafc677f67cea3c48410b03a702f3b53ea391ee.AccessPackageAssignmentResourceRoleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentResourceRole_id"] = id
    }
    return if55ecae688a10dbb886807245aafc677f67cea3c48410b03a702f3b53ea391ee.NewAccessPackageAssignmentResourceRoleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAccessPackageAssignmentRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestBuilder) {
    m := &AccessPackageAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}/accessPackageAssignment{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentRequestBuilder instantiates a new AccessPackageAssignmentRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) Delete(options *AccessPackageAssignmentRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) Get(options *AccessPackageAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignment() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment), nil
}
// Patch for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) Patch(options *AccessPackageAssignmentRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessPackageAssignmentRequestBuilder) Reprocess()(*i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b.ReprocessRequestBuilder) {
    return i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) Target()(*icf5a215bf3787be000e186a6d751b896aaac3ce17b742665321dbc2b091d4954.TargetRequestBuilder) {
    return icf5a215bf3787be000e186a6d751b896aaac3ce17b742665321dbc2b091d4954.NewTargetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
