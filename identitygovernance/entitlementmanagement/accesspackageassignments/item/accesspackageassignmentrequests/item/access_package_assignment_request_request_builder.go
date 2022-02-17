package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3e4d75bf61b18c4d849b716b4f69529b6637dcfa936d2e90da06ab2c5787ef34 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item/cancel"
    i60772d73e13472a6c6065256f75864c02575a58d62fead4e43965ec83ba44204 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item/accesspackage"
    ia880c7359c65f903496f2b17381f067f8882bb48664d81ba24b281bbfad8a387 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item/requestor"
    id4bd0f31ba7d15cc442ab19369499fcae39d2a8a874baf57cca5434731b36718 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item/accesspackageassignment"
    ifa9e53bae28ee002c0cb3cb70a24c9efe35bc3f7093405aed7714471cef46e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/entitlementmanagement/accesspackageassignments/item/accesspackageassignmentrequests/item/reprocess"
)

// AccessPackageAssignmentRequestRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackageAssignments\{accessPackageAssignment-id}\accessPackageAssignmentRequests\{accessPackageAssignmentRequest-id}
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
// AccessPackageAssignmentRequestRequestBuilderGetQueryParameters get accessPackageAssignmentRequests from identityGovernance
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
func (m *AccessPackageAssignmentRequestRequestBuilder) AccessPackage()(*i60772d73e13472a6c6065256f75864c02575a58d62fead4e43965ec83ba44204.AccessPackageRequestBuilder) {
    return i60772d73e13472a6c6065256f75864c02575a58d62fead4e43965ec83ba44204.NewAccessPackageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) AccessPackageAssignment()(*id4bd0f31ba7d15cc442ab19369499fcae39d2a8a874baf57cca5434731b36718.AccessPackageAssignmentRequestBuilder) {
    return id4bd0f31ba7d15cc442ab19369499fcae39d2a8a874baf57cca5434731b36718.NewAccessPackageAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) Cancel()(*i3e4d75bf61b18c4d849b716b4f69529b6637dcfa936d2e90da06ab2c5787ef34.CancelRequestBuilder) {
    return i3e4d75bf61b18c4d849b716b4f69529b6637dcfa936d2e90da06ab2c5787ef34.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessPackageAssignmentRequestRequestBuilderInternal instantiates a new AccessPackageAssignmentRequestRequestBuilder and sets the default values.
func NewAccessPackageAssignmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageAssignmentRequestRequestBuilder) {
    m := &AccessPackageAssignmentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment_id}/accessPackageAssignmentRequests/{accessPackageAssignmentRequest_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property accessPackageAssignmentRequests for identityGovernance
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
// CreateGetRequestInformation get accessPackageAssignmentRequests from identityGovernance
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
// CreatePatchRequestInformation update the navigation property accessPackageAssignmentRequests in identityGovernance
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
// Delete delete navigation property accessPackageAssignmentRequests for identityGovernance
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
// Get get accessPackageAssignmentRequests from identityGovernance
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
// Patch update the navigation property accessPackageAssignmentRequests in identityGovernance
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
func (m *AccessPackageAssignmentRequestRequestBuilder) Reprocess()(*ifa9e53bae28ee002c0cb3cb70a24c9efe35bc3f7093405aed7714471cef46e09.ReprocessRequestBuilder) {
    return ifa9e53bae28ee002c0cb3cb70a24c9efe35bc3f7093405aed7714471cef46e09.NewReprocessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageAssignmentRequestRequestBuilder) Requestor()(*ia880c7359c65f903496f2b17381f067f8882bb48664d81ba24b281bbfad8a387.RequestorRequestBuilder) {
    return ia880c7359c65f903496f2b17381f067f8882bb48664d81ba24b281bbfad8a387.NewRequestorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
