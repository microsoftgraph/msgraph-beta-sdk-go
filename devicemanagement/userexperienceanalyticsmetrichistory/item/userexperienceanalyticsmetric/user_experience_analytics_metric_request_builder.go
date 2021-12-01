package userexperienceanalyticsmetric

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i27a1469231d4b52006b1c8056d75e7cb189874ecec7d87bb8beb239732216e81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsmetrichistory/item/userexperienceanalyticsmetric/ref"
)

// UserExperienceAnalyticsMetricRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsMetricHistory\{userExperienceAnalyticsMetricHistory-id}\userExperienceAnalyticsMetric
type UserExperienceAnalyticsMetricRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsMetricRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsMetricRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsMetricRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsMetricRequestBuilderGetQueryParameters user experience analytics metric.
type UserExperienceAnalyticsMetricRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewUserExperienceAnalyticsMetricRequestBuilderInternal instantiates a new UserExperienceAnalyticsMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsMetricRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsMetricRequestBuilder) {
    m := &UserExperienceAnalyticsMetricRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsMetricHistory/{userExperienceAnalyticsMetricHistory_id}/userExperienceAnalyticsMetric{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsMetricRequestBuilder instantiates a new UserExperienceAnalyticsMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsMetricRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsMetricRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation user experience analytics metric.
func (m *UserExperienceAnalyticsMetricRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsMetricRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics metric.
func (m *UserExperienceAnalyticsMetricRequestBuilder) Get(options *UserExperienceAnalyticsMetricRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsMetric, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsMetric() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsMetric), nil
}
func (m *UserExperienceAnalyticsMetricRequestBuilder) Ref()(*i27a1469231d4b52006b1c8056d75e7cb189874ecec7d87bb8beb239732216e81.RefRequestBuilder) {
    return i27a1469231d4b52006b1c8056d75e7cb189874ecec7d87bb8beb239732216e81.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
