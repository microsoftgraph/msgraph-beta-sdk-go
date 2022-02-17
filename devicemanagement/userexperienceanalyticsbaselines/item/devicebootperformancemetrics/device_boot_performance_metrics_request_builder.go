package devicebootperformancemetrics

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b5cd12f65bfb5c953148247de1040135ed1d9fabecb7dac731c36d12cfc7bd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/devicebootperformancemetrics/ref"
)

// DeviceBootPerformanceMetricsRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsBaselines\{userExperienceAnalyticsBaseline-id}\deviceBootPerformanceMetrics
type DeviceBootPerformanceMetricsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceBootPerformanceMetricsRequestBuilderGetOptions options for Get
type DeviceBootPerformanceMetricsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceBootPerformanceMetricsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceBootPerformanceMetricsRequestBuilderGetQueryParameters the user experience analytics device boot performance metrics.
type DeviceBootPerformanceMetricsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewDeviceBootPerformanceMetricsRequestBuilderInternal instantiates a new DeviceBootPerformanceMetricsRequestBuilder and sets the default values.
func NewDeviceBootPerformanceMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceBootPerformanceMetricsRequestBuilder) {
    m := &DeviceBootPerformanceMetricsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline_id}/deviceBootPerformanceMetrics{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceBootPerformanceMetricsRequestBuilder instantiates a new DeviceBootPerformanceMetricsRequestBuilder and sets the default values.
func NewDeviceBootPerformanceMetricsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceBootPerformanceMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceBootPerformanceMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user experience analytics device boot performance metrics.
func (m *DeviceBootPerformanceMetricsRequestBuilder) CreateGetRequestInformation(options *DeviceBootPerformanceMetricsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the user experience analytics device boot performance metrics.
func (m *DeviceBootPerformanceMetricsRequestBuilder) Get(options *DeviceBootPerformanceMetricsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsCategory() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory), nil
}
func (m *DeviceBootPerformanceMetricsRequestBuilder) Ref()(*i5b5cd12f65bfb5c953148247de1040135ed1d9fabecb7dac731c36d12cfc7bd1.RefRequestBuilder) {
    return i5b5cd12f65bfb5c953148247de1040135ed1d9fabecb7dac731c36d12cfc7bd1.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
