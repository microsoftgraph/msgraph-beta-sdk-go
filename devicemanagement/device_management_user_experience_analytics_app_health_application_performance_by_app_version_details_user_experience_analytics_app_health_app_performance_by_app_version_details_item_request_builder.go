package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version details
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable), nil
}
