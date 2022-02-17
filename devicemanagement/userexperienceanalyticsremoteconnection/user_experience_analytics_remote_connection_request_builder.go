package userexperienceanalyticsremoteconnection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01796bfadca188d075b478c2219753fc62eb7269791c8a4122d34a3bcf9a54d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsremoteconnection/summarizedeviceremoteconnectionwithsummarizeby"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UserExperienceAnalyticsRemoteConnectionRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsRemoteConnection
type UserExperienceAnalyticsRemoteConnectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsRemoteConnectionRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsRemoteConnectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters user experience analytics remote connection
type UserExperienceAnalyticsRemoteConnectionRequestBuilderGetQueryParameters struct {
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
// UserExperienceAnalyticsRemoteConnectionRequestBuilderPostOptions options for Post
type UserExperienceAnalyticsRemoteConnectionRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRemoteConnection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal instantiates a new UserExperienceAnalyticsRemoteConnectionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    m := &UserExperienceAnalyticsRemoteConnectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsRemoteConnection{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsRemoteConnectionRequestBuilder instantiates a new UserExperienceAnalyticsRemoteConnectionRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRemoteConnectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation user experience analytics remote connection
func (m *UserExperienceAnalyticsRemoteConnectionRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsRemoteConnectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation user experience analytics remote connection
func (m *UserExperienceAnalyticsRemoteConnectionRequestBuilder) CreatePostRequestInformation(options *UserExperienceAnalyticsRemoteConnectionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get user experience analytics remote connection
func (m *UserExperienceAnalyticsRemoteConnectionRequestBuilder) Get(options *UserExperienceAnalyticsRemoteConnectionRequestBuilderGetOptions)(*UserExperienceAnalyticsRemoteConnectionResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsRemoteConnectionResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserExperienceAnalyticsRemoteConnectionResponse), nil
}
// Post user experience analytics remote connection
func (m *UserExperienceAnalyticsRemoteConnectionRequestBuilder) Post(options *UserExperienceAnalyticsRemoteConnectionRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRemoteConnection, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsRemoteConnection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRemoteConnection), nil
}
// SummarizeDeviceRemoteConnectionWithSummarizeBy builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsRemoteConnection\microsoft.graph.summarizeDeviceRemoteConnection(summarizeBy={summarizeBy})
func (m *UserExperienceAnalyticsRemoteConnectionRequestBuilder) SummarizeDeviceRemoteConnectionWithSummarizeBy(summarizeBy *string)(*i01796bfadca188d075b478c2219753fc62eb7269791c8a4122d34a3bcf9a54d0.SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    return i01796bfadca188d075b478c2219753fc62eb7269791c8a4122d34a3bcf9a54d0.NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(m.pathParameters, m.requestAdapter, summarizeBy);
}
