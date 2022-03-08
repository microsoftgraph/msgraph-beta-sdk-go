package getmanagementconditionsforplatformwithplatform

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetManagementConditionsForPlatformWithPlatformRequestBuilder provides operations to call the getManagementConditionsForPlatform method.
type GetManagementConditionsForPlatformWithPlatformRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions options for Get
type GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, platform *string)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    m := &GetManagementConditionsForPlatformWithPlatformRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditions/microsoft.graph.getManagementConditionsForPlatform(platform='{platform}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if platform != nil {
        urlTplParams["platform"] = *platform
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilder instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformation(options *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) Get(options *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions)(GetManagementConditionsForPlatformWithPlatformResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagementConditionsForPlatformWithPlatformResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagementConditionsForPlatformWithPlatformResponseable), nil
}
