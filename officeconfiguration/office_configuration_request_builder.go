package officeconfiguration

import (
    ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// OfficeConfigurationRequestBuilder provides operations to manage the officeConfiguration singleton.
type OfficeConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OfficeConfigurationRequestBuilderGetOptions options for Get
type OfficeConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OfficeConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OfficeConfigurationRequestBuilderGetQueryParameters get officeConfiguration
type OfficeConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OfficeConfigurationRequestBuilderPatchOptions options for Patch
type OfficeConfigurationRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OfficeConfigurationRequestBuilder) ClientConfigurations()(*ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f.ClientConfigurationsRequestBuilder) {
    return ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f.NewClientConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClientConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.officeConfiguration.clientConfigurations.item collection
func (m *OfficeConfigurationRequestBuilder) ClientConfigurationsById(id string)(*i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381.OfficeClientConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfiguration_id"] = id
    }
    return i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381.NewOfficeClientConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOfficeConfigurationRequestBuilderInternal instantiates a new OfficeConfigurationRequestBuilder and sets the default values.
func NewOfficeConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeConfigurationRequestBuilder) {
    m := &OfficeConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeConfigurationRequestBuilder instantiates a new OfficeConfigurationRequestBuilder and sets the default values.
func NewOfficeConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OfficeConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get officeConfiguration
func (m *OfficeConfigurationRequestBuilder) CreateGetRequestInformation(options *OfficeConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update officeConfiguration
func (m *OfficeConfigurationRequestBuilder) CreatePatchRequestInformation(options *OfficeConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Get get officeConfiguration
func (m *OfficeConfigurationRequestBuilder) Get(options *OfficeConfigurationRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOfficeConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OfficeConfigurationable), nil
}
// Patch update officeConfiguration
func (m *OfficeConfigurationRequestBuilder) Patch(options *OfficeConfigurationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
