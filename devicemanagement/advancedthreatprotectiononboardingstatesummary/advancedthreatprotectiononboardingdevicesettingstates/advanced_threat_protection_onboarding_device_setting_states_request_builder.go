package advancedthreatprotectiononboardingdevicesettingstates

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder builds and executes requests for operations under \deviceManagement\advancedThreatProtectionOnboardingStateSummary\advancedThreatProtectionOnboardingDeviceSettingStates
type AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetOptions options for Get
type AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters not yet documented
type AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetQueryParameters struct {
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
// AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostOptions options for Post
type AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingDeviceSettingState;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    m := &AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation not yet documented
func (m *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) CreateGetRequestInformation(options *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation not yet documented
func (m *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) CreatePostRequestInformation(options *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get not yet documented
func (m *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Get(options *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderGetOptions)(*AdvancedThreatProtectionOnboardingDeviceSettingStatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdvancedThreatProtectionOnboardingDeviceSettingStatesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*AdvancedThreatProtectionOnboardingDeviceSettingStatesResponse), nil
}
// Post not yet documented
func (m *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) Post(options *AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingDeviceSettingState, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAdvancedThreatProtectionOnboardingDeviceSettingState() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingDeviceSettingState), nil
}
