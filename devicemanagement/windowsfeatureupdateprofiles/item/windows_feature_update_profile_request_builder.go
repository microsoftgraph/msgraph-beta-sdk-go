package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assign"
    ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments"
    i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsfeatureupdateprofiles/item/assignments/item"
)

// Builds and executes requests for operations under \deviceManagement\windowsFeatureUpdateProfiles\{windowsFeatureUpdateProfile-id}
type WindowsFeatureUpdateProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type WindowsFeatureUpdateProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type WindowsFeatureUpdateProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// A collection of windows feature update profiles
type WindowsFeatureUpdateProfileRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type WindowsFeatureUpdateProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Assign()(*i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.AssignRequestBuilder) {
    return i17ae4c396051acc3d8969fe0d510a15f08f74b0af97608557d2125c16de1bb7d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsFeatureUpdateProfileRequestBuilder) Assignments()(*ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.AssignmentsRequestBuilder) {
    return ib9f19a8a1ed79d6335dbd9524123a1a111e94345212d4f98f256d4d77fb64865.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.windowsFeatureUpdateProfiles.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *WindowsFeatureUpdateProfileRequestBuilder) AssignmentsById(id string)(*i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.WindowsFeatureUpdateProfileAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsFeatureUpdateProfileAssignment_id"] = id
    }
    return i9fc8cb618d1d2cb2bd52c185cd35fcf810ea303a5692d3f159788228e6d3514f.NewWindowsFeatureUpdateProfileAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new WindowsFeatureUpdateProfileRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWindowsFeatureUpdateProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileRequestBuilder) {
    m := &WindowsFeatureUpdateProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsFeatureUpdateProfiles/{windowsFeatureUpdateProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WindowsFeatureUpdateProfileRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWindowsFeatureUpdateProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsFeatureUpdateProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsFeatureUpdateProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreateDeleteRequestInformation(options *WindowsFeatureUpdateProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreateGetRequestInformation(options *WindowsFeatureUpdateProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) CreatePatchRequestInformation(options *WindowsFeatureUpdateProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) Delete(options *WindowsFeatureUpdateProfileRequestBuilderDeleteOptions)(error) {
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
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) Get(options *WindowsFeatureUpdateProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsFeatureUpdateProfile() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsFeatureUpdateProfile), nil
}
// A collection of windows feature update profiles
// Parameters:
//  - options : Options for the request
func (m *WindowsFeatureUpdateProfileRequestBuilder) Patch(options *WindowsFeatureUpdateProfileRequestBuilderPatchOptions)(error) {
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
