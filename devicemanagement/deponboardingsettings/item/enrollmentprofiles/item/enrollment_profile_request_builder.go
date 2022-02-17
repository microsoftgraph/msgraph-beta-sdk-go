package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/setdefaultprofile"
    i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/updatedeviceprofileassignment"
    ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/exportmobileconfig"
)

// EnrollmentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\enrollmentProfiles\{enrollmentProfile-id}
type EnrollmentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EnrollmentProfileRequestBuilderDeleteOptions options for Delete
type EnrollmentProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EnrollmentProfileRequestBuilderGetOptions options for Get
type EnrollmentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EnrollmentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EnrollmentProfileRequestBuilderGetQueryParameters the enrollment profiles.
type EnrollmentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EnrollmentProfileRequestBuilderPatchOptions options for Patch
type EnrollmentProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEnrollmentProfileRequestBuilderInternal instantiates a new EnrollmentProfileRequestBuilder and sets the default values.
func NewEnrollmentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnrollmentProfileRequestBuilder) {
    m := &EnrollmentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/enrollmentProfiles/{enrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnrollmentProfileRequestBuilder instantiates a new EnrollmentProfileRequestBuilder and sets the default values.
func NewEnrollmentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnrollmentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnrollmentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) CreateDeleteRequestInformation(options *EnrollmentProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) CreateGetRequestInformation(options *EnrollmentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) CreatePatchRequestInformation(options *EnrollmentProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) Delete(options *EnrollmentProfileRequestBuilderDeleteOptions)(error) {
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
// ExportMobileConfig builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\enrollmentProfiles\{enrollmentProfile-id}\microsoft.graph.exportMobileConfig()
func (m *EnrollmentProfileRequestBuilder) ExportMobileConfig()(*ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294.ExportMobileConfigRequestBuilder) {
    return ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294.NewExportMobileConfigRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) Get(options *EnrollmentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfile), nil
}
// Patch the enrollment profiles.
func (m *EnrollmentProfileRequestBuilder) Patch(options *EnrollmentProfileRequestBuilderPatchOptions)(error) {
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
func (m *EnrollmentProfileRequestBuilder) SetDefaultProfile()(*i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f.SetDefaultProfileRequestBuilder) {
    return i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f.NewSetDefaultProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EnrollmentProfileRequestBuilder) UpdateDeviceProfileAssignment()(*i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9.UpdateDeviceProfileAssignmentRequestBuilder) {
    return i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9.NewUpdateDeviceProfileAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
