package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iabf54a6528230cc6fd8becb59f999d1ad986020b16420c27db539498922c2851 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intunebrandingprofiles/item/assign"
    id933fdbaba7880973435cc8cf621ae774d7ea0758e1ed6de778e48418e22e695 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intunebrandingprofiles/item/assignments"
    i43e2d15146d44703f2826df7d20b0717eda3da488c1e8099cae91884e38e54c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/intunebrandingprofiles/item/assignments/item"
)

// IntuneBrandingProfileRequestBuilder builds and executes requests for operations under \deviceManagement\intuneBrandingProfiles\{intuneBrandingProfile-id}
type IntuneBrandingProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IntuneBrandingProfileRequestBuilderDeleteOptions options for Delete
type IntuneBrandingProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IntuneBrandingProfileRequestBuilderGetOptions options for Get
type IntuneBrandingProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IntuneBrandingProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IntuneBrandingProfileRequestBuilderGetQueryParameters intune branding profiles targeted to AAD groups
type IntuneBrandingProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IntuneBrandingProfileRequestBuilderPatchOptions options for Patch
type IntuneBrandingProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IntuneBrandingProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IntuneBrandingProfileRequestBuilder) Assign()(*iabf54a6528230cc6fd8becb59f999d1ad986020b16420c27db539498922c2851.AssignRequestBuilder) {
    return iabf54a6528230cc6fd8becb59f999d1ad986020b16420c27db539498922c2851.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IntuneBrandingProfileRequestBuilder) Assignments()(*id933fdbaba7880973435cc8cf621ae774d7ea0758e1ed6de778e48418e22e695.AssignmentsRequestBuilder) {
    return id933fdbaba7880973435cc8cf621ae774d7ea0758e1ed6de778e48418e22e695.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.intuneBrandingProfiles.item.assignments.item collection
func (m *IntuneBrandingProfileRequestBuilder) AssignmentsById(id string)(*i43e2d15146d44703f2826df7d20b0717eda3da488c1e8099cae91884e38e54c4.IntuneBrandingProfileAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["intuneBrandingProfileAssignment_id"] = id
    }
    return i43e2d15146d44703f2826df7d20b0717eda3da488c1e8099cae91884e38e54c4.NewIntuneBrandingProfileAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIntuneBrandingProfileRequestBuilderInternal instantiates a new IntuneBrandingProfileRequestBuilder and sets the default values.
func NewIntuneBrandingProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IntuneBrandingProfileRequestBuilder) {
    m := &IntuneBrandingProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/intuneBrandingProfiles/{intuneBrandingProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIntuneBrandingProfileRequestBuilder instantiates a new IntuneBrandingProfileRequestBuilder and sets the default values.
func NewIntuneBrandingProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IntuneBrandingProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntuneBrandingProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) CreateDeleteRequestInformation(options *IntuneBrandingProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) CreateGetRequestInformation(options *IntuneBrandingProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) CreatePatchRequestInformation(options *IntuneBrandingProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) Delete(options *IntuneBrandingProfileRequestBuilderDeleteOptions)(error) {
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
// Get intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) Get(options *IntuneBrandingProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IntuneBrandingProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIntuneBrandingProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IntuneBrandingProfile), nil
}
// Patch intune branding profiles targeted to AAD groups
func (m *IntuneBrandingProfileRequestBuilder) Patch(options *IntuneBrandingProfileRequestBuilderPatchOptions)(error) {
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
