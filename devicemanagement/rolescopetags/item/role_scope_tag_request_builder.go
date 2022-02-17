package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i75b21611d6b6b96c1c95e37193e653d7c1be7468679cf5785af829f0827b9b05 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/item/assign"
    iab305975e356746fe6505f00637e9813ac95de45da8b925112b990b92d6b0f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/item/assignments"
    ic12c5ef80a8671db45da105570fcb6fe91462ab2daa7529a50e81a4c171aa147 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/rolescopetags/item/assignments/item"
)

// RoleScopeTagRequestBuilder builds and executes requests for operations under \deviceManagement\roleScopeTags\{roleScopeTag-id}
type RoleScopeTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleScopeTagRequestBuilderDeleteOptions options for Delete
type RoleScopeTagRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleScopeTagRequestBuilderGetOptions options for Get
type RoleScopeTagRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleScopeTagRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleScopeTagRequestBuilderGetQueryParameters the Role Scope Tags.
type RoleScopeTagRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RoleScopeTagRequestBuilderPatchOptions options for Patch
type RoleScopeTagRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTag;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RoleScopeTagRequestBuilder) Assign()(*i75b21611d6b6b96c1c95e37193e653d7c1be7468679cf5785af829f0827b9b05.AssignRequestBuilder) {
    return i75b21611d6b6b96c1c95e37193e653d7c1be7468679cf5785af829f0827b9b05.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleScopeTagRequestBuilder) Assignments()(*iab305975e356746fe6505f00637e9813ac95de45da8b925112b990b92d6b0f57.AssignmentsRequestBuilder) {
    return iab305975e356746fe6505f00637e9813ac95de45da8b925112b990b92d6b0f57.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.roleScopeTags.item.assignments.item collection
func (m *RoleScopeTagRequestBuilder) AssignmentsById(id string)(*ic12c5ef80a8671db45da105570fcb6fe91462ab2daa7529a50e81a4c171aa147.RoleScopeTagAutoAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleScopeTagAutoAssignment_id"] = id
    }
    return ic12c5ef80a8671db45da105570fcb6fe91462ab2daa7529a50e81a4c171aa147.NewRoleScopeTagAutoAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleScopeTagRequestBuilderInternal instantiates a new RoleScopeTagRequestBuilder and sets the default values.
func NewRoleScopeTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagRequestBuilder) {
    m := &RoleScopeTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags/{roleScopeTag_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScopeTagRequestBuilder instantiates a new RoleScopeTagRequestBuilder and sets the default values.
func NewRoleScopeTagRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScopeTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) CreateDeleteRequestInformation(options *RoleScopeTagRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) CreateGetRequestInformation(options *RoleScopeTagRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) CreatePatchRequestInformation(options *RoleScopeTagRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) Delete(options *RoleScopeTagRequestBuilderDeleteOptions)(error) {
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
// Get the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) Get(options *RoleScopeTagRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTag, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRoleScopeTag() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleScopeTag), nil
}
// Patch the Role Scope Tags.
func (m *RoleScopeTagRequestBuilder) Patch(options *RoleScopeTagRequestBuilderPatchOptions)(error) {
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
