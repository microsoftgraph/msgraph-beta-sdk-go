package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters user experience analytics appHealth Application Performance by App Version Device Id
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetQueryParameters
}
// DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    m := &DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId/{userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder instantiates a new UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder and sets the default values.
func NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId for deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable), nil
}
// Patch update the navigation property userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId in deviceManagement
func (m *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, requestConfiguration *DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable), nil
}
