package workfromanywheremetrics

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0bdef080c03ea7bfb67fcc7efb5d00b12a7df0c8ff5ca61522a948881caa2191 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/workfromanywheremetrics/ref"
)

// WorkFromAnywhereMetricsRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsBaselines\{userExperienceAnalyticsBaseline-id}\workFromAnywhereMetrics
type WorkFromAnywhereMetricsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WorkFromAnywhereMetricsRequestBuilderGetOptions options for Get
type WorkFromAnywhereMetricsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WorkFromAnywhereMetricsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WorkFromAnywhereMetricsRequestBuilderGetQueryParameters the user experience analytics work from anywhere metrics.
type WorkFromAnywhereMetricsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewWorkFromAnywhereMetricsRequestBuilderInternal instantiates a new WorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewWorkFromAnywhereMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkFromAnywhereMetricsRequestBuilder) {
    m := &WorkFromAnywhereMetricsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline_id}/workFromAnywhereMetrics{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkFromAnywhereMetricsRequestBuilder instantiates a new WorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewWorkFromAnywhereMetricsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkFromAnywhereMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkFromAnywhereMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user experience analytics work from anywhere metrics.
func (m *WorkFromAnywhereMetricsRequestBuilder) CreateGetRequestInformation(options *WorkFromAnywhereMetricsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the user experience analytics work from anywhere metrics.
func (m *WorkFromAnywhereMetricsRequestBuilder) Get(options *WorkFromAnywhereMetricsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory, error) {
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
func (m *WorkFromAnywhereMetricsRequestBuilder) Ref()(*i0bdef080c03ea7bfb67fcc7efb5d00b12a7df0c8ff5ca61522a948881caa2191.RefRequestBuilder) {
    return i0bdef080c03ea7bfb67fcc7efb5d00b12a7df0c8ff5ca61522a948881caa2191.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
