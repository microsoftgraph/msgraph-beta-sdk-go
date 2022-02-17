package accesspackageassignment

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/reprocess"
    i699cd609567d8d2b2ad367f5b49ef2229d2fbfd56ac49dd0c18eaf0745dc8848 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment/ref"
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
// AccessPackageAssignmentRequestBuilderGetQueryParameters for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
type AccessPackageAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
// CreateGetRequestInformation for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
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
// Get for a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequestBuilder) Get(options *AccessPackageAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignment), nil
}
func (m *AccessPackageAssignmentRequestBuilder) Ref()(*i699cd609567d8d2b2ad367f5b49ef2229d2fbfd56ac49dd0c18eaf0745dc8848.RefRequestBuilder) {
    return i699cd609567d8d2b2ad367f5b49ef2229d2fbfd56ac49dd0c18eaf0745dc8848.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestBuilder) Reprocess()(*i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b.ReprocessRequestBuilder) {
    return i050da783a45c88977b6b539e65ec872a9a528c347a261b3fb8decd462dcac33b.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
