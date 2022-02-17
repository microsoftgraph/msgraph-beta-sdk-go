package userexperienceanalyticsapphealthoverview

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i76d76df01534fbe41283b7237bc4b6635f2316d4abaf6113e056758ee16ef5cc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsapphealthoverview/metricvalues"
    i52615d5294b7c09cc7abf48113e76d304d971ab6e369a32e6e648d565d5d7767 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsapphealthoverview/metricvalues/item"
)

// UserExperienceAnalyticsAppHealthOverviewRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsAppHealthOverview
type UserExperienceAnalyticsAppHealthOverviewRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsAppHealthOverviewRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsAppHealthOverviewRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetQueryParameters user experience analytics appHealth overview
type UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsAppHealthOverviewRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsAppHealthOverviewRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthOverviewRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    m := &UserExperienceAnalyticsAppHealthOverviewRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthOverview{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsAppHealthOverviewRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthOverviewRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsAppHealthOverviewRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) Delete(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderDeleteOptions)(error) {
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
// Get user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) Get(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsCategory() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory), nil
}
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) MetricValues()(*i76d76df01534fbe41283b7237bc4b6635f2316d4abaf6113e056758ee16ef5cc.MetricValuesRequestBuilder) {
    return i76d76df01534fbe41283b7237bc4b6635f2316d4abaf6113e056758ee16ef5cc.NewMetricValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MetricValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsAppHealthOverview.metricValues.item collection
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) MetricValuesById(id string)(*i52615d5294b7c09cc7abf48113e76d304d971ab6e369a32e6e648d565d5d7767.UserExperienceAnalyticsMetricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return i52615d5294b7c09cc7abf48113e76d304d971ab6e369a32e6e648d565d5d7767.NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch user experience analytics appHealth overview
func (m *UserExperienceAnalyticsAppHealthOverviewRequestBuilder) Patch(options *UserExperienceAnalyticsAppHealthOverviewRequestBuilderPatchOptions)(error) {
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
