package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticscategories/item/metricvalues"
    i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticscategories/item/metricvalues/item"
)

// UserExperienceAnalyticsCategoryRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsCategories\{userExperienceAnalyticsCategory-id}
type UserExperienceAnalyticsCategoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsCategoryRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsCategoryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsCategoryRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsCategoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsCategoryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsCategoryRequestBuilderGetQueryParameters user experience analytics categories
type UserExperienceAnalyticsCategoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsCategoryRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsCategoryRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsCategoryRequestBuilderInternal instantiates a new UserExperienceAnalyticsCategoryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsCategoryRequestBuilder) {
    m := &UserExperienceAnalyticsCategoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsCategories/{userExperienceAnalyticsCategory_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsCategoryRequestBuilder instantiates a new UserExperienceAnalyticsCategoryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsCategoryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsCategoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsCategoryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) Delete(options *UserExperienceAnalyticsCategoryRequestBuilderDeleteOptions)(error) {
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
// Get user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) Get(options *UserExperienceAnalyticsCategoryRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategory, error) {
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
func (m *UserExperienceAnalyticsCategoryRequestBuilder) MetricValues()(*ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1.MetricValuesRequestBuilder) {
    return ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1.NewMetricValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MetricValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsCategories.item.metricValues.item collection
func (m *UserExperienceAnalyticsCategoryRequestBuilder) MetricValuesById(id string)(*i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42.UserExperienceAnalyticsMetricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42.NewUserExperienceAnalyticsMetricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch user experience analytics categories
func (m *UserExperienceAnalyticsCategoryRequestBuilder) Patch(options *UserExperienceAnalyticsCategoryRequestBuilderPatchOptions)(error) {
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
