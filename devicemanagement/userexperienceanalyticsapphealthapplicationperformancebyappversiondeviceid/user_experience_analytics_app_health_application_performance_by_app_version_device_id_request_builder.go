package userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ida22a59c71c1e40357ed507a50c96ee20e5f0497d5e85b947b96c9142ba9e63a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsapphealthapplicationperformancebyappversiondeviceid/count"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version Device Id
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Count()(*ida22a59c71c1e40357ed507a50c96ee20e5f0497d5e85b947b96c9142ba9e63a.CountRequestBuilder) {
    return ida22a59c71c1e40357ed507a50c96ee20e5f0497d5e85b947b96c9142ba9e63a.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation user experience analytics appHealth Application Performance by App Version Device Id
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics appHealth Application Performance by App Version Device Id
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Get(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) Post(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable), nil
}
