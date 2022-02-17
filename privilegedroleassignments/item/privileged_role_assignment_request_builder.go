package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i360fd14cb1918888dde8889252044616ef9e8bae14b63f8ac32974d2ab3ab93f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/makepermanent"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iaea252828cfd0bc8915348077699bb4a6d57de37db1f39571a36aa9d242ade3c "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/makeeligible"
    ie1054a1b9157187681bd433f8846c01e56b7924bf3c589f1ade883924304dc7d "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo"
)

// PrivilegedRoleAssignmentRequestBuilder builds and executes requests for operations under \privilegedRoleAssignments\{privilegedRoleAssignment-id}
type PrivilegedRoleAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedRoleAssignmentRequestBuilderDeleteOptions options for Delete
type PrivilegedRoleAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleAssignmentRequestBuilderGetOptions options for Get
type PrivilegedRoleAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedRoleAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedRoleAssignmentRequestBuilderGetQueryParameters get entity from privilegedRoleAssignments by key
type PrivilegedRoleAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedRoleAssignmentRequestBuilderPatchOptions options for Patch
type PrivilegedRoleAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPrivilegedRoleAssignmentRequestBuilderInternal instantiates a new PrivilegedRoleAssignmentRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleAssignmentRequestBuilder) {
    m := &PrivilegedRoleAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedRoleAssignmentRequestBuilder instantiates a new PrivilegedRoleAssignmentRequestBuilder and sets the default values.
func NewPrivilegedRoleAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedRoleAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedRoleAssignments by key
func (m *PrivilegedRoleAssignmentRequestBuilder) CreateGetRequestInformation(options *PrivilegedRoleAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentRequestBuilder) CreatePatchRequestInformation(options *PrivilegedRoleAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentRequestBuilder) Delete(options *PrivilegedRoleAssignmentRequestBuilderDeleteOptions)(error) {
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
// Get get entity from privilegedRoleAssignments by key
func (m *PrivilegedRoleAssignmentRequestBuilder) Get(options *PrivilegedRoleAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedRoleAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedRoleAssignment), nil
}
func (m *PrivilegedRoleAssignmentRequestBuilder) MakeEligible()(*iaea252828cfd0bc8915348077699bb4a6d57de37db1f39571a36aa9d242ade3c.MakeEligibleRequestBuilder) {
    return iaea252828cfd0bc8915348077699bb4a6d57de37db1f39571a36aa9d242ade3c.NewMakeEligibleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrivilegedRoleAssignmentRequestBuilder) MakePermanent()(*i360fd14cb1918888dde8889252044616ef9e8bae14b63f8ac32974d2ab3ab93f.MakePermanentRequestBuilder) {
    return i360fd14cb1918888dde8889252044616ef9e8bae14b63f8ac32974d2ab3ab93f.NewMakePermanentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in privilegedRoleAssignments
func (m *PrivilegedRoleAssignmentRequestBuilder) Patch(options *PrivilegedRoleAssignmentRequestBuilderPatchOptions)(error) {
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
func (m *PrivilegedRoleAssignmentRequestBuilder) RoleInfo()(*ie1054a1b9157187681bd433f8846c01e56b7924bf3c589f1ade883924304dc7d.RoleInfoRequestBuilder) {
    return ie1054a1b9157187681bd433f8846c01e56b7924bf3c589f1ade883924304dc7d.NewRoleInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
