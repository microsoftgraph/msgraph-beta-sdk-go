package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsworkfromanywheremetrics/item/metricdevices"
    ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsworkfromanywheremetrics/item/metricdevices/item"
)

// Builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsWorkFromAnywhereMetrics\{userExperienceAnalyticsWorkFromAnywhereMetric-id}
type UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// User experience analytics work from anywhere metrics.
type UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderInternal(urlParams, requestAdapter)
}
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) Delete(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) Get(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsWorkFromAnywhereMetric() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereMetric), nil
}
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) MetricDevices()(*iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503.MetricDevicesRequestBuilder) {
    return iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503.NewMetricDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.userExperienceAnalyticsWorkFromAnywhereMetrics.item.metricDevices.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) MetricDevicesById(id string)(*ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6.UserExperienceAnalyticsWorkFromAnywhereDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereDevice_id"] = id
    }
    return ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6.NewUserExperienceAnalyticsWorkFromAnywhereDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// User experience analytics work from anywhere metrics.
// Parameters:
//  - options : Options for the request
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilder) Patch(options *UserExperienceAnalyticsWorkFromAnywhereMetricRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
