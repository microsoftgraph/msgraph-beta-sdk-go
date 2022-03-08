package userexperienceanalyticsapphealthapplicationperformancebyappversion

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i84373d61bfcce691515bbcc8278b8f9cd7d2914c714057d40e3eaf0967a9beb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsapphealthapplicationperformancebyappversion/count"
)

// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) Count()(*i84373d61bfcce691515bbcc8278b8f9cd7d2914c714057d40e3eaf0967a9beb9.CountRequestBuilder) {
    return i84373d61bfcce691515bbcc8278b8f9cd7d2914c714057d40e3eaf0967a9beb9.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation user experience analytics appHealth Application Performance by App Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics appHealth Application Performance by App Version
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) Get(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion for deviceManagement
func (m *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) Post(options *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable), nil
}
