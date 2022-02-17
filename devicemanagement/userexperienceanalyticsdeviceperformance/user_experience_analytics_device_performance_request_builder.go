package userexperienceanalyticsdeviceperformance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id1528e2471ae36a0ae7981b45584b17b30e7a3e550591a6cc40dc5d26c932f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsdeviceperformance/summarizedeviceperformancedeviceswithsummarizeby"
)

// UserExperienceAnalyticsDevicePerformanceRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsDevicePerformanceRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsDevicePerformanceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters user experience analytics device performance
type UserExperienceAnalyticsDevicePerformanceRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsDevicePerformanceRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsDevicePerformanceRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsDevicePerformance;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal instantiates a new UserExperienceAnalyticsDevicePerformanceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    m := &UserExperienceAnalyticsDevicePerformanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicePerformance{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsDevicePerformanceRequestBuilder instantiates a new UserExperienceAnalyticsDevicePerformanceRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDevicePerformanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation user experience analytics device performance
func (m *UserExperienceAnalyticsDevicePerformanceRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsDevicePerformanceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation user experience analytics device performance
func (m *UserExperienceAnalyticsDevicePerformanceRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsDevicePerformanceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics device performance
func (m *UserExperienceAnalyticsDevicePerformanceRequestBuilder) Get(options *UserExperienceAnalyticsDevicePerformanceRequestBuilderGetOptions)(*UserExperienceAnalyticsDevicePerformanceResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDevicePerformanceResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserExperienceAnalyticsDevicePerformanceResponse), nil
}
// Post user experience analytics device performance
func (m *UserExperienceAnalyticsDevicePerformanceRequestBuilder) Post(options *UserExperienceAnalyticsDevicePerformanceRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsDevicePerformance, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsDevicePerformance() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsDevicePerformance), nil
}
// SummarizeDevicePerformanceDevicesWithSummarizeBy builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsDevicePerformance\microsoft.graph.summarizeDevicePerformanceDevices(summarizeBy={summarizeBy})
func (m *UserExperienceAnalyticsDevicePerformanceRequestBuilder) SummarizeDevicePerformanceDevicesWithSummarizeBy(summarizeBy *string)(*id1528e2471ae36a0ae7981b45584b17b30e7a3e550591a6cc40dc5d26c932f61.SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    return id1528e2471ae36a0ae7981b45584b17b30e7a3e550591a6cc40dc5d26c932f61.NewSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(m.pathParameters, m.requestAdapter, summarizeBy);
}
