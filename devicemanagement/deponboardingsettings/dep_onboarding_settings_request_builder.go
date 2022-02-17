package deponboardingsettings

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia653be664a0dd75c52af6e78b0cf052819bdf93d3b5de4be59dc76970e065611 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deponboardingsettings/getexpiringvpptokencountwithexpiringbeforedatetime"
)

// DepOnboardingSettingsRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings
type DepOnboardingSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DepOnboardingSettingsRequestBuilderGetOptions options for Get
type DepOnboardingSettingsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DepOnboardingSettingsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DepOnboardingSettingsRequestBuilderGetQueryParameters this collections of multiple DEP tokens per-tenant.
type DepOnboardingSettingsRequestBuilderGetQueryParameters struct {
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
// DepOnboardingSettingsRequestBuilderPostOptions options for Post
type DepOnboardingSettingsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDepOnboardingSettingsRequestBuilderInternal instantiates a new DepOnboardingSettingsRequestBuilder and sets the default values.
func NewDepOnboardingSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingsRequestBuilder) {
    m := &DepOnboardingSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDepOnboardingSettingsRequestBuilder instantiates a new DepOnboardingSettingsRequestBuilder and sets the default values.
func NewDepOnboardingSettingsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DepOnboardingSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsRequestBuilder) CreateGetRequestInformation(options *DepOnboardingSettingsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsRequestBuilder) CreatePostRequestInformation(options *DepOnboardingSettingsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsRequestBuilder) Get(options *DepOnboardingSettingsRequestBuilderGetOptions)(*DepOnboardingSettingsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepOnboardingSettingsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DepOnboardingSettingsResponse), nil
}
// GetExpiringVppTokenCountWithExpiringBeforeDateTime builds and executes requests for operations under \deviceManagement\depOnboardingSettings\microsoft.graph.getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')
func (m *DepOnboardingSettingsRequestBuilder) GetExpiringVppTokenCountWithExpiringBeforeDateTime(expiringBeforeDateTime *string)(*ia653be664a0dd75c52af6e78b0cf052819bdf93d3b5de4be59dc76970e065611.GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    return ia653be664a0dd75c52af6e78b0cf052819bdf93d3b5de4be59dc76970e065611.NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, expiringBeforeDateTime);
}
// Post this collections of multiple DEP tokens per-tenant.
func (m *DepOnboardingSettingsRequestBuilder) Post(options *DepOnboardingSettingsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDepOnboardingSetting() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DepOnboardingSetting), nil
}
