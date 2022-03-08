package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/setdefaultprofile"
    i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/updatedeviceprofileassignment"
    ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/enrollmentprofiles/item/exportmobileconfig"
)

// EnrollmentProfileItemRequestBuilder provides operations to manage the enrollmentProfiles property of the microsoft.graph.depOnboardingSetting entity.
type EnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EnrollmentProfileItemRequestBuilderDeleteOptions options for Delete
type EnrollmentProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EnrollmentProfileItemRequestBuilderGetOptions options for Get
type EnrollmentProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EnrollmentProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EnrollmentProfileItemRequestBuilderGetQueryParameters the enrollment profiles.
type EnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EnrollmentProfileItemRequestBuilderPatchOptions options for Patch
type EnrollmentProfileItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfileable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEnrollmentProfileItemRequestBuilderInternal instantiates a new EnrollmentProfileItemRequestBuilder and sets the default values.
func NewEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnrollmentProfileItemRequestBuilder) {
    m := &EnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/enrollmentProfiles/{enrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEnrollmentProfileItemRequestBuilder instantiates a new EnrollmentProfileItemRequestBuilder and sets the default values.
func NewEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property enrollmentProfiles for deviceManagement
func (m *EnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(options *EnrollmentProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(options *EnrollmentProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property enrollmentProfiles in deviceManagement
func (m *EnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(options *EnrollmentProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property enrollmentProfiles for deviceManagement
func (m *EnrollmentProfileItemRequestBuilder) Delete(options *EnrollmentProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportMobileConfig provides operations to call the exportMobileConfig method.
func (m *EnrollmentProfileItemRequestBuilder) ExportMobileConfig()(*ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294.ExportMobileConfigRequestBuilder) {
    return ie810f95f442b7e5b0679893cb8cade4ba1695eaa96df7ba757075302d6fd2294.NewExportMobileConfigRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the enrollment profiles.
func (m *EnrollmentProfileItemRequestBuilder) Get(options *EnrollmentProfileItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEnrollmentProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EnrollmentProfileable), nil
}
// Patch update the navigation property enrollmentProfiles in deviceManagement
func (m *EnrollmentProfileItemRequestBuilder) Patch(options *EnrollmentProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EnrollmentProfileItemRequestBuilder) SetDefaultProfile()(*i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f.SetDefaultProfileRequestBuilder) {
    return i2368cb001c0085f52ee4b80f8cbadbf0281840398627682c4832aae50f7e0d0f.NewSetDefaultProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EnrollmentProfileItemRequestBuilder) UpdateDeviceProfileAssignment()(*i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9.UpdateDeviceProfileAssignmentRequestBuilder) {
    return i43d975eb4ff9e5eee6d18830ab6b48dcbf7d0ba77a301b0a135586f21e707fa9.NewUpdateDeviceProfileAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
