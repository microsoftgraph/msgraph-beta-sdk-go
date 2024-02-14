package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *UserExperienceAnalyticsBaselinesItemAppHealthMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) AppHealthMetrics()(*UserExperienceAnalyticsBaselinesItemAppHealthMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemAppHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BatteryHealthMetrics provides operations to manage the batteryHealthMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemBatteryHealthMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BatteryHealthMetrics()(*UserExperienceAnalyticsBaselinesItemBatteryHealthMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemBatteryHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BestPracticesMetrics provides operations to manage the bestPracticesMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemBestPracticesMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) BestPracticesMetrics()(*UserExperienceAnalyticsBaselinesItemBestPracticesMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemBestPracticesMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    m := &UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder instantiates a new UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsBaselines for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceBootPerformanceMetrics provides operations to manage the deviceBootPerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemDeviceBootPerformanceMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) DeviceBootPerformanceMetrics()(*UserExperienceAnalyticsBaselinesItemDeviceBootPerformanceMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemDeviceBootPerformanceMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics baselines
// returns a UserExperienceAnalyticsBaselineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable), nil
}
// Patch update the navigation property userExperienceAnalyticsBaselines in deviceManagement
// returns a UserExperienceAnalyticsBaselineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable), nil
}
// RebootAnalyticsMetrics provides operations to manage the rebootAnalyticsMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemRebootAnalyticsMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) RebootAnalyticsMetrics()(*UserExperienceAnalyticsBaselinesItemRebootAnalyticsMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemRebootAnalyticsMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourcePerformanceMetrics provides operations to manage the resourcePerformanceMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemResourcePerformanceMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ResourcePerformanceMetrics()(*UserExperienceAnalyticsBaselinesItemResourcePerformanceMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemResourcePerformanceMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsBaselines for deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics baselines
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsBaselines in deviceManagement
// returns a *RequestInformation when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/userExperienceAnalyticsBaselines/{userExperienceAnalyticsBaseline%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) WithUrl(rawUrl string)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// WorkFromAnywhereMetrics provides operations to manage the workFromAnywhereMetrics property of the microsoft.graph.userExperienceAnalyticsBaseline entity.
// returns a *UserExperienceAnalyticsBaselinesItemWorkFromAnywhereMetricsRequestBuilder when successful
func (m *UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) WorkFromAnywhereMetrics()(*UserExperienceAnalyticsBaselinesItemWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesItemWorkFromAnywhereMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
