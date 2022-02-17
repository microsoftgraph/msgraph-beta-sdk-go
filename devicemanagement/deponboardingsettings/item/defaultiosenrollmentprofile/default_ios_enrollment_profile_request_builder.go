package defaultiosenrollmentprofile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    idd246b53dd8b816e2149588d387d4e8dfb17d3f56a921204859f055f83125f9b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/item/defaultiosenrollmentprofile/ref"
)

// DefaultIosEnrollmentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\defaultIosEnrollmentProfile
type DefaultIosEnrollmentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DefaultIosEnrollmentProfileRequestBuilderGetOptions options for Get
type DefaultIosEnrollmentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DefaultIosEnrollmentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DefaultIosEnrollmentProfileRequestBuilderGetQueryParameters default iOS Enrollment Profile
type DefaultIosEnrollmentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDefaultIosEnrollmentProfileRequestBuilderInternal instantiates a new DefaultIosEnrollmentProfileRequestBuilder and sets the default values.
func NewDefaultIosEnrollmentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefaultIosEnrollmentProfileRequestBuilder) {
    m := &DefaultIosEnrollmentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/defaultIosEnrollmentProfile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDefaultIosEnrollmentProfileRequestBuilder instantiates a new DefaultIosEnrollmentProfileRequestBuilder and sets the default values.
func NewDefaultIosEnrollmentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DefaultIosEnrollmentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefaultIosEnrollmentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation default iOS Enrollment Profile
func (m *DefaultIosEnrollmentProfileRequestBuilder) CreateGetRequestInformation(options *DefaultIosEnrollmentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get default iOS Enrollment Profile
func (m *DefaultIosEnrollmentProfileRequestBuilder) Get(options *DefaultIosEnrollmentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepIOSEnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDepIOSEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepIOSEnrollmentProfile), nil
}
func (m *DefaultIosEnrollmentProfileRequestBuilder) Ref()(*idd246b53dd8b816e2149588d387d4e8dfb17d3f56a921204859f055f83125f9b.RefRequestBuilder) {
    return idd246b53dd8b816e2149588d387d4e8dfb17d3f56a921204859f055f83125f9b.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
