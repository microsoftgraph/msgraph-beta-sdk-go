package deviceconfigurations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/gettargetedusersanddevices"
    i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/getiosavailableupdateversions"
    ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/haspayloadlinks"
)

// Builds and executes requests for operations under \deviceManagement\deviceConfigurations
type DeviceConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type DeviceConfigurationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceConfigurationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The device configurations.
type DeviceConfigurationsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Options for Post
type DeviceConfigurationsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new DeviceConfigurationsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationsRequestBuilder) {
    m := &DeviceConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeviceConfigurationsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceConfigurationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// The device configurations.
// Parameters:
//  - options : Options for the request
func (m *DeviceConfigurationsRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The device configurations.
// Parameters:
//  - options : Options for the request
func (m *DeviceConfigurationsRequestBuilder) CreatePostRequestInformation(options *DeviceConfigurationsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The device configurations.
// Parameters:
//  - options : Options for the request
func (m *DeviceConfigurationsRequestBuilder) Get(options *DeviceConfigurationsRequestBuilderGetOptions)(*DeviceConfigurationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DeviceConfigurationsResponse), nil
}
// Builds and executes requests for operations under \deviceManagement\deviceConfigurations\microsoft.graph.getIosAvailableUpdateVersions()
func (m *DeviceConfigurationsRequestBuilder) GetIosAvailableUpdateVersions()(*i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a.GetIosAvailableUpdateVersionsRequestBuilder) {
    return i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a.NewGetIosAvailableUpdateVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationsRequestBuilder) GetTargetedUsersAndDevices()(*i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a.GetTargetedUsersAndDevicesRequestBuilder) {
    return i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a.NewGetTargetedUsersAndDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceConfigurationsRequestBuilder) HasPayloadLinks()(*ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f.HasPayloadLinksRequestBuilder) {
    return ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The device configurations.
// Parameters:
//  - options : Options for the request
func (m *DeviceConfigurationsRequestBuilder) Post(options *DeviceConfigurationsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceConfiguration), nil
}
