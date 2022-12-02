package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformance/{userExperienceAnalyticsAppHealthApplicationPerformance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics appHealth Application Performance
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformance for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformance in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthApplicationPerformanceable), nil
}
