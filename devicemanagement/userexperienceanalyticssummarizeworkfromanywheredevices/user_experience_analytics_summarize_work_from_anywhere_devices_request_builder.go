package userexperienceanalyticssummarizeworkfromanywheredevices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    m := &UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.userExperienceAnalyticsSummarizeWorkFromAnywhereDevices()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereDevicesSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryable), nil
}
