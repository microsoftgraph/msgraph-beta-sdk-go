package userexperienceanalyticsworkfromanywherehardwarereadinessmetric

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters user experience analytics work from anywhere hardware readiness metrics.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Delete(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderDeleteOptions)(error) {
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
// Get user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Get(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric), nil
}
// Patch user experience analytics work from anywhere hardware readiness metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) Patch(options *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderPatchOptions)(error) {
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
