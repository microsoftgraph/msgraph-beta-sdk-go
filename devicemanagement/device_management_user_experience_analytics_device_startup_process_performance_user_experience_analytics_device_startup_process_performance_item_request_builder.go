package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters user experience analytics device Startup Process Performance
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDeviceStartupProcessPerformance/{userExperienceAnalyticsDeviceStartupProcessPerformance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics device Startup Process Performance
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsDeviceStartupProcessPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics device Startup Process Performance
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsDeviceStartupProcessPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsDeviceStartupProcessPerformanceable), nil
}
