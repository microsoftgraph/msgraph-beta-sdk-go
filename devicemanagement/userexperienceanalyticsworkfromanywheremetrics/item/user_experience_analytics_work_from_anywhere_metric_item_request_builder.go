package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsworkfromanywheremetrics/item/metricdevices"
    ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsworkfromanywheremetrics/item/metricdevices/item"
)

// UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters user experience analytics work from anywhere metrics.
type UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable), nil
}
// MetricDevices the metricDevices property
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) MetricDevices()(*iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503.MetricDevicesRequestBuilder) {
    return iba99f863645e713540730fd9068eedb628780d4a813936eff3986ef86a299503.NewMetricDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MetricDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.userExperienceAnalyticsWorkFromAnywhereMetrics.item.metricDevices.item collection
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) MetricDevicesById(id string)(*ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6.UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereDevice%2Did"] = id
    }
    return ic2caf0b6e98dd567ee96828b8d8b16a38ae5bc311f5a57fb5cdf060263d6b2e6.NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
