package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters the user experience analytics work from anywhere model performance
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereModelPerformance/{userExperienceAnalyticsWorkFromAnywhereModelPerformance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the user experience analytics work from anywhere model performance
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the user experience analytics work from anywhere model performance
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereModelPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable), nil
}
