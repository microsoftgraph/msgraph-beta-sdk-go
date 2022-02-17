package defaultmacosenrollmentprofile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0812f16609971fd2afda2ff60c9b9126c8eb7ba65101a4d80ab915ac50208db7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/defaultmacosenrollmentprofile/ref"
)

// DefaultMacOsEnrollmentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\defaultMacOsEnrollmentProfile
type DefaultMacOsEnrollmentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DefaultMacOsEnrollmentProfileRequestBuilderGetOptions options for Get
type DefaultMacOsEnrollmentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DefaultMacOsEnrollmentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DefaultMacOsEnrollmentProfileRequestBuilderGetQueryParameters default MacOs Enrollment Profile
type DefaultMacOsEnrollmentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDefaultMacOsEnrollmentProfileRequestBuilderInternal instantiates a new DefaultMacOsEnrollmentProfileRequestBuilder and sets the default values.
func NewDefaultMacOsEnrollmentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefaultMacOsEnrollmentProfileRequestBuilder) {
    m := &DefaultMacOsEnrollmentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/defaultMacOsEnrollmentProfile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDefaultMacOsEnrollmentProfileRequestBuilder instantiates a new DefaultMacOsEnrollmentProfileRequestBuilder and sets the default values.
func NewDefaultMacOsEnrollmentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefaultMacOsEnrollmentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefaultMacOsEnrollmentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation default MacOs Enrollment Profile
func (m *DefaultMacOsEnrollmentProfileRequestBuilder) CreateGetRequestInformation(options *DefaultMacOsEnrollmentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get default MacOs Enrollment Profile
func (m *DefaultMacOsEnrollmentProfileRequestBuilder) Get(options *DefaultMacOsEnrollmentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepMacOSEnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDepMacOSEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepMacOSEnrollmentProfile), nil
}
func (m *DefaultMacOsEnrollmentProfileRequestBuilder) Ref()(*i0812f16609971fd2afda2ff60c9b9126c8eb7ba65101a4d80ab915ac50208db7.RefRequestBuilder) {
    return i0812f16609971fd2afda2ff60c9b9126c8eb7ba65101a4d80ab915ac50208db7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
