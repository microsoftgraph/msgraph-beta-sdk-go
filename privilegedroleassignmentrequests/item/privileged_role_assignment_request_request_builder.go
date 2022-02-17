package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib239ddd157212e9b2de338617aed29750c7fa7cb6355343e0685a26014096aa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/cancel"
    ieff61d093ce682e939abbaecbc1b098cece4e7f58170d229e88f2f7f6b9511ee "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests/item/roleinfo"
)

// PrivilegedRoleAssignmentRequestRequestBuilder builds and executes requests for operations under \privilegedRoleAssignmentRequests\{privilegedRoleAssignmentRequest-id}
type PrivilegedRoleAssignmentRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedRoleAssignmentRequestRequestBuilderDeleteOptions options for Delete
type PrivilegedRoleAssignmentRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleAssignmentRequestRequestBuilderGetOptions options for Get
type PrivilegedRoleAssignmentRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedRoleAssignmentRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleAssignmentRequestRequestBuilderGetQueryParameters get entity from privilegedRoleAssignmentRequests by key
type PrivilegedRoleAssignmentRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedRoleAssignmentRequestRequestBuilderPatchOptions options for Patch
type PrivilegedRoleAssignmentRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignmentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) Cancel()(*ib239ddd157212e9b2de338617aed29750c7fa7cb6355343e0685a26014096aa8.CancelRequestBuilder) {
    return ib239ddd157212e9b2de338617aed29750c7fa7cb6355343e0685a26014096aa8.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrivilegedRoleAssignmentRequestRequestBuilderInternal instantiates a new PrivilegedRoleAssignmentRequestRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleAssignmentRequestRequestBuilder) {
    m := &PrivilegedRoleAssignmentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignmentRequests/{privilegedRoleAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleAssignmentRequestRequestBuilder instantiates a new PrivilegedRoleAssignmentRequestRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleAssignmentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleAssignmentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedRoleAssignmentRequests
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedRoleAssignmentRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedRoleAssignmentRequests by key
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) CreateGetRequestInformation(options *PrivilegedRoleAssignmentRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedRoleAssignmentRequests
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) CreatePatchRequestInformation(options *PrivilegedRoleAssignmentRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from privilegedRoleAssignmentRequests
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) Delete(options *PrivilegedRoleAssignmentRequestRequestBuilderDeleteOptions)(error) {
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
// Get get entity from privilegedRoleAssignmentRequests by key
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) Get(options *PrivilegedRoleAssignmentRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedRoleAssignmentRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignmentRequest), nil
}
// Patch update entity in privilegedRoleAssignmentRequests
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) Patch(options *PrivilegedRoleAssignmentRequestRequestBuilderPatchOptions)(error) {
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
func (m *PrivilegedRoleAssignmentRequestRequestBuilder) RoleInfo()(*ieff61d093ce682e939abbaecbc1b098cece4e7f58170d229e88f2f7f6b9511ee.RoleInfoRequestBuilder) {
    return ieff61d093ce682e939abbaecbc1b098cece4e7f58170d229e88f2f7f6b9511ee.NewRoleInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
