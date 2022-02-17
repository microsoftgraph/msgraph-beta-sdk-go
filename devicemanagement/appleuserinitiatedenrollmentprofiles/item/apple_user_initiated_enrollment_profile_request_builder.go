package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/setpriority"
    i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/assignments"
    i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/appleuserinitiatedenrollmentprofiles/item/assignments/item"
)

// AppleUserInitiatedEnrollmentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\appleUserInitiatedEnrollmentProfiles\{appleUserInitiatedEnrollmentProfile-id}
type AppleUserInitiatedEnrollmentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppleUserInitiatedEnrollmentProfileRequestBuilderDeleteOptions options for Delete
type AppleUserInitiatedEnrollmentProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppleUserInitiatedEnrollmentProfileRequestBuilderGetOptions options for Get
type AppleUserInitiatedEnrollmentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AppleUserInitiatedEnrollmentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppleUserInitiatedEnrollmentProfileRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type AppleUserInitiatedEnrollmentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AppleUserInitiatedEnrollmentProfileRequestBuilderPatchOptions options for Patch
type AppleUserInitiatedEnrollmentProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppleUserInitiatedEnrollmentProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) Assignments()(*i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031.AssignmentsRequestBuilder) {
    return i8f74274c95f6e543171df682273240e35174a6d9718f8eb24972864f87ab1031.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.appleUserInitiatedEnrollmentProfiles.item.assignments.item collection
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) AssignmentsById(id string)(*i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56.AppleEnrollmentProfileAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleEnrollmentProfileAssignment_id"] = id
    }
    return i3f2b91af36c1c85a91b2090f64ecc402fe3be3e6030dc816276efb34d3538a56.NewAppleEnrollmentProfileAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAppleUserInitiatedEnrollmentProfileRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfileRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppleUserInitiatedEnrollmentProfileRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppleUserInitiatedEnrollmentProfileRequestBuilder instantiates a new AppleUserInitiatedEnrollmentProfileRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppleUserInitiatedEnrollmentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) CreateDeleteRequestInformation(options *AppleUserInitiatedEnrollmentProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) CreateGetRequestInformation(options *AppleUserInitiatedEnrollmentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) CreatePatchRequestInformation(options *AppleUserInitiatedEnrollmentProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) Delete(options *AppleUserInitiatedEnrollmentProfileRequestBuilderDeleteOptions)(error) {
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
// Get apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) Get(options *AppleUserInitiatedEnrollmentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppleUserInitiatedEnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAppleUserInitiatedEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppleUserInitiatedEnrollmentProfile), nil
}
// Patch apple user initiated enrollment profiles
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) Patch(options *AppleUserInitiatedEnrollmentProfileRequestBuilderPatchOptions)(error) {
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
func (m *AppleUserInitiatedEnrollmentProfileRequestBuilder) SetPriority()(*i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.SetPriorityRequestBuilder) {
    return i45dabbbe435eb6199c9451e73874575e112663a6cbcd76c6d70f2e7434ffb64d.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
