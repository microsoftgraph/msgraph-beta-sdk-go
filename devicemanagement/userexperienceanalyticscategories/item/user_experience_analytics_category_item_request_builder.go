package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticscategories/item/metricvalues"
    i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticscategories/item/metricvalues/item"
)

// UserExperienceAnalyticsCategoryItemRequestBuilder provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsCategoryItemRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsCategoryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsCategoryItemRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsCategoryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsCategoryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsCategoryItemRequestBuilderGetQueryParameters user experience analytics categories
type UserExperienceAnalyticsCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsCategoryItemRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsCategoryItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategoryable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsCategoryItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsCategoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsCategoryItemRequestBuilder) {
    m := &UserExperienceAnalyticsCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsCategories/{userExperienceAnalyticsCategory_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsCategoryItemRequestBuilder instantiates a new UserExperienceAnalyticsCategoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsCategoryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsCategories for deviceManagement
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsCategoryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsCategoryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsCategories in deviceManagement
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsCategoryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsCategories for deviceManagement
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) Delete(options *UserExperienceAnalyticsCategoryItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
// Get user experience analytics categories
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) Get(options *UserExperienceAnalyticsCategoryItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsCategoryable), nil
}
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) MetricValues()(*ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1.MetricValuesRequestBuilder) {
    return ia7cafd794b20eca3ff11b37b8e1eefd547ec1e595639eda0f03245a3f4253fb1.NewMetricValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MetricValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsCategories.item.metricValues.item collection
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) MetricValuesById(id string)(*i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42.UserExperienceAnalyticsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric_id"] = id
    }
    return i723c6e2ab05e2e91dea99207a63fe4f6ad55551bd1e2ce27c2f6a08a95d00d42.NewUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property userExperienceAnalyticsCategories in deviceManagement
func (m *UserExperienceAnalyticsCategoryItemRequestBuilder) Patch(options *UserExperienceAnalyticsCategoryItemRequestBuilderPatchOptions)(error) {
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
