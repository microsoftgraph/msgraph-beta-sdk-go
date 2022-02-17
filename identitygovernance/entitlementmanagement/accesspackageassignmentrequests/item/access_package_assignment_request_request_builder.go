package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i20e40892b0ebc7d6727eb7367496253e1dd2795803dfbfa588e57fec226e6bec "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackageassignment"
    i8cc72335bb5264f40443a0982ff3a60667a85d4e64c284cb2e45a1ba6754ad86 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/accesspackage"
    ia1207ce55af86318e1e73216f3a0b4be0d3c28fc5312ab2ae337e585e8f0444f "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/cancel"
    ic5984a5d6cbabd6e1a999485b785661d63586b2603e22e8744d6624a79b6c47e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/reprocess"
    ifa69905a653164d85c24018d6751bf3e19322529bbd454fc76751f0bdc4612ce "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignmentrequests/item/requestor"
)

// AccessPackageAssignmentRequestRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}
type AccessPackageAssignmentRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageAssignmentRequestRequestBuilderDeleteOptions options for Delete
type AccessPackageAssignmentRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestRequestBuilderGetOptions options for Get
type AccessPackageAssignmentRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageAssignmentRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageAssignmentRequestRequestBuilderGetQueryParameters represents access package assignment requests created by or on behalf of a user.
type AccessPackageAssignmentRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessPackageAssignmentRequestRequestBuilderPatchOptions options for Patch
type AccessPackageAssignmentRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AccessPackageAssignmentRequestRequestBuilder) AccessPackage()(*i8cc72335bb5264f40443a0982ff3a60667a85d4e64c284cb2e45a1ba6754ad86.AccessPackageRequestBuilder) {
    return i8cc72335bb5264f40443a0982ff3a60667a85d4e64c284cb2e45a1ba6754ad86.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) AccessPackageAssignment()(*i20e40892b0ebc7d6727eb7367496253e1dd2795803dfbfa588e57fec226e6bec.AccessPackageAssignmentRequestBuilder) {
    return i20e40892b0ebc7d6727eb7367496253e1dd2795803dfbfa588e57fec226e6bec.NewAccessPackageAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) Cancel()(*ia1207ce55af86318e1e73216f3a0b4be0d3c28fc5312ab2ae337e585e8f0444f.CancelRequestBuilder) {
    return ia1207ce55af86318e1e73216f3a0b4be0d3c28fc5312ab2ae337e585e8f0444f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentRequestRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestRequestBuilder) {
    m := &AccessPackageAssignmentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageAssignmentRequestRequestBuilder instantiates a new AccessPackageAssignmentRequestRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageAssignmentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) CreateDeleteRequestInformation(options *AccessPackageAssignmentRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) CreateGetRequestInformation(options *AccessPackageAssignmentRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) CreatePatchRequestInformation(options *AccessPackageAssignmentRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) Delete(options *AccessPackageAssignmentRequestRequestBuilderDeleteOptions)(error) {
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
// Get represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) Get(options *AccessPackageAssignmentRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessPackageAssignmentRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessPackageAssignmentRequest), nil
}
// Patch represents access package assignment requests created by or on behalf of a user.
func (m *AccessPackageAssignmentRequestRequestBuilder) Patch(options *AccessPackageAssignmentRequestRequestBuilderPatchOptions)(error) {
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
func (m *AccessPackageAssignmentRequestRequestBuilder) Reprocess()(*ic5984a5d6cbabd6e1a999485b785661d63586b2603e22e8744d6624a79b6c47e.ReprocessRequestBuilder) {
    return ic5984a5d6cbabd6e1a999485b785661d63586b2603e22e8744d6624a79b6c47e.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) Requestor()(*ifa69905a653164d85c24018d6751bf3e19322529bbd454fc76751f0bdc4612ce.RequestorRequestBuilder) {
    return ifa69905a653164d85c24018d6751bf3e19322529bbd454fc76751f0bdc4612ce.NewRequestorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
