package userexperienceanalyticsresourceperformance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9dfdde130af06bcea2bfd024361a3ac4535811cdd4d93ebdede50b24a52080ee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsresourceperformance/summarizedeviceresourceperformancewithsummarizeby"
)

// UserExperienceAnalyticsResourcePerformanceRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsResourcePerformance
type UserExperienceAnalyticsResourcePerformanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsResourcePerformanceRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsResourcePerformanceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters user experience analytics resource performance
type UserExperienceAnalyticsResourcePerformanceRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsResourcePerformanceRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsResourcePerformanceRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsResourcePerformance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal instantiates a new UserExperienceAnalyticsResourcePerformanceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    m := &UserExperienceAnalyticsResourcePerformanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsResourcePerformance{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsResourcePerformanceRequestBuilder instantiates a new UserExperienceAnalyticsResourcePerformanceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsResourcePerformanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation user experience analytics resource performance
func (m *UserExperienceAnalyticsResourcePerformanceRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsResourcePerformanceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation user experience analytics resource performance
func (m *UserExperienceAnalyticsResourcePerformanceRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsResourcePerformanceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics resource performance
func (m *UserExperienceAnalyticsResourcePerformanceRequestBuilder) Get(options *UserExperienceAnalyticsResourcePerformanceRequestBuilderGetOptions)(*UserExperienceAnalyticsResourcePerformanceResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsResourcePerformanceResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserExperienceAnalyticsResourcePerformanceResponse), nil
}
// Post user experience analytics resource performance
func (m *UserExperienceAnalyticsResourcePerformanceRequestBuilder) Post(options *UserExperienceAnalyticsResourcePerformanceRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsResourcePerformance, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsResourcePerformance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsResourcePerformance), nil
}
// SummarizeDeviceResourcePerformanceWithSummarizeBy builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsResourcePerformance\microsoft.graph.summarizeDeviceResourcePerformance(summarizeBy={summarizeBy})
func (m *UserExperienceAnalyticsResourcePerformanceRequestBuilder) SummarizeDeviceResourcePerformanceWithSummarizeBy(summarizeBy *string)(*i9dfdde130af06bcea2bfd024361a3ac4535811cdd4d93ebdede50b24a52080ee.SummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilder) {
    return i9dfdde130af06bcea2bfd024361a3ac4535811cdd4d93ebdede50b24a52080ee.NewSummarizeDeviceResourcePerformanceWithSummarizeByRequestBuilderInternal(m.pathParameters, m.requestAdapter, summarizeBy);
}
