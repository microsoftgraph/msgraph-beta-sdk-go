package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters user experience analytics baselines
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppHealthMetrics provides operations to manage the appHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) AppHealthMetrics()(*UserExperienceAnalyticsBaselinesItemAppHealthMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemAppHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatteryHealthMetrics provides operations to manage the batteryHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BatteryHealthMetrics()(*UserExperienceAnalyticsBaselinesItemBatteryHealthMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemBatteryHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BestPracticesMetrics provides operations to manage the bestPracticesMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BestPracticesMetrics()(*UserExperienceAnalyticsBaselinesItemBestPracticesMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemBestPracticesMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    m := &UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder instantiates a new UserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsBaselines for deviceManagement
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation user experience analytics baselines
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsBaselines in deviceManagement
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userExperienceAnalyticsBaselines for deviceManagement
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceBootPerformanceMetrics provides operations to manage the deviceBootPerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) DeviceBootPerformanceMetrics()(*UserExperienceAnalyticsBaselinesItemDeviceBootPerformanceMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemDeviceBootPerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get user experience analytics baselines
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable), nil
}
// Patch update the navigation property userExperienceAnalyticsBaselines in deviceManagement
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable), nil
}
// RebootAnalyticsMetrics provides operations to manage the rebootAnalyticsMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) RebootAnalyticsMetrics()(*UserExperienceAnalyticsBaselinesItemRebootAnalyticsMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemRebootAnalyticsMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcePerformanceMetrics provides operations to manage the resourcePerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ResourcePerformanceMetrics()(*UserExperienceAnalyticsBaselinesItemResourcePerformanceMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemResourcePerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkFromAnywhereMetrics provides operations to manage the workFromAnywhereMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) WorkFromAnywhereMetrics()(*UserExperienceAnalyticsBaselinesItemWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
