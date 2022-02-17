package userexperienceanalyticsregressionsummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/operatingsystemregression"
    i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/modelregression"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/summarizedeviceregressionperformancewithsummarizeby"
    ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/manufacturerregression"
    i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/modelregression/item"
    i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/manufacturerregression/item"
    ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/operatingsystemregression/item"
)

// UserExperienceAnalyticsRegressionSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsRegressionSummary
type UserExperienceAnalyticsRegressionSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsRegressionSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters user experience analytics regression summary
type UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal instantiates a new UserExperienceAnalyticsRegressionSummaryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsRegressionSummaryRequestBuilder) {
    m := &UserExperienceAnalyticsRegressionSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsRegressionSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsRegressionSummaryRequestBuilder instantiates a new UserExperienceAnalyticsRegressionSummaryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRegressionSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsRegressionSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Delete(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Get(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsRegressionSummary() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary), nil
}
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ManufacturerRegression()(*ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef.ManufacturerRegressionRequestBuilder) {
    return ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef.NewManufacturerRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManufacturerRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.manufacturerRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ManufacturerRegressionById(id string)(*i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090.UserExperienceAnalyticsMetricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090.NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ModelRegression()(*i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026.ModelRegressionRequestBuilder) {
    return i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026.NewModelRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ModelRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.modelRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ModelRegressionById(id string)(*i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af.UserExperienceAnalyticsMetricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af.NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) OperatingSystemRegression()(*i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf.OperatingSystemRegressionRequestBuilder) {
    return i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf.NewOperatingSystemRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperatingSystemRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.operatingSystemRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) OperatingSystemRegressionById(id string)(*ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123.UserExperienceAnalyticsMetricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123.NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Patch(options *UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// SummarizeDeviceRegressionPerformanceWithSummarizeBy builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsRegressionSummary\microsoft.graph.summarizeDeviceRegressionPerformance(summarizeBy={summarizeBy})
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) SummarizeDeviceRegressionPerformanceWithSummarizeBy(summarizeBy *string)(*i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001.SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    return i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001.NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(m.pathParameters, m.requestAdapter, summarizeBy);
}
