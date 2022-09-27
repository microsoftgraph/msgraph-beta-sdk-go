package userexperienceanalyticsregressionsummary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/operatingsystemregression"
    i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/modelregression"
    i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/summarizedeviceregressionperformancewithsummarizeby"
    ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/manufacturerregression"
    i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/modelregression/item"
    i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/manufacturerregression/item"
    ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsregressionsummary/operatingsystemregression/item"
)

// UserExperienceAnalyticsRegressionSummaryRequestBuilder provides operations to manage the userExperienceAnalyticsRegressionSummary property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsRegressionSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters user experience analytics regression summary
type UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsRegressionSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal instantiates a new UserExperienceAnalyticsRegressionSummaryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsRegressionSummaryRequestBuilder) {
    m := &UserExperienceAnalyticsRegressionSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsRegressionSummary{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsRegressionSummaryRequestBuilder instantiates a new UserExperienceAnalyticsRegressionSummaryRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsRegressionSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsRegressionSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsRegressionSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsRegressionSummary for deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property userExperienceAnalyticsRegressionSummary for deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsRegressionSummary in deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property userExperienceAnalyticsRegressionSummary in deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userExperienceAnalyticsRegressionSummary for deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get user experience analytics regression summary
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable), nil
}
// ManufacturerRegression the manufacturerRegression property
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ManufacturerRegression()(*ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef.ManufacturerRegressionRequestBuilder) {
    return ic83bad237f322810848e3865826b1cbcd8315bd015232234086583ad511fbdef.NewManufacturerRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManufacturerRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.manufacturerRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ManufacturerRegressionById(id string)(*i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090.UserExperienceAnalyticsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric%2Did"] = id
    }
    return i54bb50d57456c6f802caf30d89d436dccdca16bf7d2f43fc122831205ef1c090.NewUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ModelRegression the modelRegression property
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ModelRegression()(*i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026.ModelRegressionRequestBuilder) {
    return i42d726374ab44f3d882f404444614d1c77a96a52e30de57f4bd2616537880026.NewModelRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ModelRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.modelRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) ModelRegressionById(id string)(*i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af.UserExperienceAnalyticsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric%2Did"] = id
    }
    return i10ba7a2a99f49bf517ec16f33e2e22787b4abcbfbd747010335af7d1ae7c19af.NewUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OperatingSystemRegression the operatingSystemRegression property
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) OperatingSystemRegression()(*i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf.OperatingSystemRegressionRequestBuilder) {
    return i25b934baebf68a26cb5dd00d6359e6dd4b954d2959e5581085e9ded3a14dbdcf.NewOperatingSystemRegressionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperatingSystemRegressionById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsRegressionSummary.operatingSystemRegression.item collection
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) OperatingSystemRegressionById(id string)(*ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123.UserExperienceAnalyticsMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetric%2Did"] = id
    }
    return ifb56eb83a662054d3af4bc800c85e7a65c5bc3df936f5213775ecdd960ce4123.NewUserExperienceAnalyticsMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property userExperienceAnalyticsRegressionSummary in deviceManagement
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, requestConfiguration *UserExperienceAnalyticsRegressionSummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsRegressionSummaryable), nil
}
// SummarizeDeviceRegressionPerformanceWithSummarizeBy provides operations to call the summarizeDeviceRegressionPerformance method.
func (m *UserExperienceAnalyticsRegressionSummaryRequestBuilder) SummarizeDeviceRegressionPerformanceWithSummarizeBy(summarizeBy *string)(*i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001.SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    return i5dc2ebb3d97e030370ec345b6acf14891716de436f8309f8003f0644ec997001.NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(m.pathParameters, m.requestAdapter, summarizeBy);
}
