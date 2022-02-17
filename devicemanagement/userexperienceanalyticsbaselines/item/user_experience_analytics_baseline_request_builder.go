package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/apphealthmetrics"
    i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/workfromanywheremetrics"
    i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/batteryhealthmetrics"
    i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/devicebootperformancemetrics"
    icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/rebootanalyticsmetrics"
    iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/bestpracticesmetrics"
    iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/resourceperformancemetrics"
)

// UserExperienceAnalyticsBaselineRequestBuilder builds and executes requests for operations under \deviceManagement\userExperienceAnalyticsBaselines\{userExperienceAnalyticsBaseline-id}
type UserExperienceAnalyticsBaselineRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsBaselineRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsBaselineRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsBaselineRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsBaselineRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters user experience analytics baselines
type UserExperienceAnalyticsBaselineRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsBaselineRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsBaselineRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) AppHealthMetrics()(*i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.AppHealthMetricsRequestBuilder) {
    return i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.NewAppHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) BatteryHealthMetrics()(*i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530.BatteryHealthMetricsRequestBuilder) {
    return i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530.NewBatteryHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) BestPracticesMetrics()(*iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.BestPracticesMetricsRequestBuilder) {
    return iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.NewBestPracticesMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserExperienceAnalyticsBaselineRequestBuilderInternal instantiates a new UserExperienceAnalyticsBaselineRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselineRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsBaselineRequestBuilder) {
    m := &UserExperienceAnalyticsBaselineRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsBaselineRequestBuilder instantiates a new UserExperienceAnalyticsBaselineRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselineRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsBaselineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBaselineRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsBaselineRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsBaselineRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsBaselineRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Delete(options *UserExperienceAnalyticsBaselineRequestBuilderDeleteOptions)(error) {
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
func (m *UserExperienceAnalyticsBaselineRequestBuilder) DeviceBootPerformanceMetrics()(*i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.DeviceBootPerformanceMetricsRequestBuilder) {
    return i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.NewDeviceBootPerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Get(options *UserExperienceAnalyticsBaselineRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsBaseline() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsBaseline), nil
}
// Patch user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineRequestBuilder) Patch(options *UserExperienceAnalyticsBaselineRequestBuilderPatchOptions)(error) {
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
func (m *UserExperienceAnalyticsBaselineRequestBuilder) RebootAnalyticsMetrics()(*icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.RebootAnalyticsMetricsRequestBuilder) {
    return icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.NewRebootAnalyticsMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) ResourcePerformanceMetrics()(*iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.ResourcePerformanceMetricsRequestBuilder) {
    return iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.NewResourcePerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserExperienceAnalyticsBaselineRequestBuilder) WorkFromAnywhereMetrics()(*i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.WorkFromAnywhereMetricsRequestBuilder) {
    return i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.NewWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
