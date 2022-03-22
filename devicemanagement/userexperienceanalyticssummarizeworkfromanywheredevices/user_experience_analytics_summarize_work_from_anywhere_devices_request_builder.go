package userexperienceanalyticssummarizeworkfromanywheredevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    m := &UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.userExperienceAnalyticsSummarizeWorkFromAnywhereDevices()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) Get(options *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsWorkFromAnywhereDevicesSummaryFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable), nil
}
