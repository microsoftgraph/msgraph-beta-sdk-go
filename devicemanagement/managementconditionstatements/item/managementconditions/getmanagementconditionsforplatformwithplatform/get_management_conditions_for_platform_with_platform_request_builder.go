package getmanagementconditionsforplatformwithplatform

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetManagementConditionsForPlatformWithPlatformRequestBuilder builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}\managementConditions\microsoft.graph.getManagementConditionsForPlatform(platform={platform})
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
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditionStatements/{managementConditionStatement_id}/managementConditions/microsoft.graph.getManagementConditionsForPlatform(platform={platform})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if platform != nil {
        urlTplParams["platform"] = *platform
    }
    m.pathParameters = pathParameters;
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
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) Get(options *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions)([]GetManagementConditionsForPlatformWithPlatform, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetManagementConditionsForPlatformWithPlatform() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetManagementConditionsForPlatformWithPlatform, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetManagementConditionsForPlatformWithPlatform))
    }
    return val, nil
}
