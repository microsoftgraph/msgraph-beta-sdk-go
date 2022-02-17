package userexperienceanalyticsworkfromanywheremetrics

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsWorkFromAnywhereMetrics
type UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters user experience analytics work from anywhere metrics.
type UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) Get(options *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetOptions)(*UserExperienceAnalyticsWorkFromAnywhereMetricsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWorkFromAnywhereMetricsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserExperienceAnalyticsWorkFromAnywhereMetricsResponse), nil
}
// Post user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) Post(options *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsWorkFromAnywhereMetric() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric), nil
}
