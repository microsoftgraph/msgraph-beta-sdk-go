package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/apphealthmetrics"
    i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/workfromanywheremetrics"
    i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/batteryhealthmetrics"
    i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/devicebootperformancemetrics"
    icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/rebootanalyticsmetrics"
    iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/bestpracticesmetrics"
    iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/userexperienceanalyticsbaselines/item/resourceperformancemetrics"
)

// UserExperienceAnalyticsBaselineItemRequestBuilder provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBaselineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters user experience analytics baselines
type UserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsBaselineItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppHealthMetrics the appHealthMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) AppHealthMetrics()(*i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.AppHealthMetricsRequestBuilder) {
    return i058f5de0abd53df643c0b1d4535f14317cbfcc0a5ef23108987eaeb903f7eb21.NewAppHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatteryHealthMetrics the batteryHealthMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) BatteryHealthMetrics()(*i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530.BatteryHealthMetricsRequestBuilder) {
    return i56415ec7de06a6b4fe9e8a6bf6d8c7e2b0a89b604e16c86ce803db1711e2c530.NewBatteryHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BestPracticesMetrics the bestPracticesMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) BestPracticesMetrics()(*iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.BestPracticesMetricsRequestBuilder) {
    return iec2f756b9554ce46110f6b3f733d0617bf3ef06dc480e86cb09cbf2d74628de8.NewBestPracticesMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserExperienceAnalyticsBaselineItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselineItemRequestBuilder) {
    m := &UserExperienceAnalyticsBaselineItemRequestBuilder{
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
// NewUserExperienceAnalyticsBaselineItemRequestBuilder instantiates a new UserExperienceAnalyticsBaselineItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBaselineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userExperienceAnalyticsBaselines for deviceManagement
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property userExperienceAnalyticsBaselines in deviceManagement
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property userExperienceAnalyticsBaselines for deviceManagement
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceBootPerformanceMetrics the deviceBootPerformanceMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) DeviceBootPerformanceMetrics()(*i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.DeviceBootPerformanceMetricsRequestBuilder) {
    return i69e441d79d9ffb70d727d5891fb03cb7264fe1ac972103efbef3cdeff428347c.NewDeviceBootPerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get user experience analytics baselines
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
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
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, requestConfiguration *UserExperienceAnalyticsBaselineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBaselineable, error) {
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
// RebootAnalyticsMetrics the rebootAnalyticsMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) RebootAnalyticsMetrics()(*icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.RebootAnalyticsMetricsRequestBuilder) {
    return icd0719662a86b52cac5fcea4da793c6735d376b486c1450eed1b5ddc8ff8e6e2.NewRebootAnalyticsMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcePerformanceMetrics the resourcePerformanceMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) ResourcePerformanceMetrics()(*iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.ResourcePerformanceMetricsRequestBuilder) {
    return iecb3f73f1e42490bd6888019100a7693b3f982d354ec76f90bf067fc6fabe7de.NewResourcePerformanceMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WorkFromAnywhereMetrics the workFromAnywhereMetrics property
func (m *UserExperienceAnalyticsBaselineItemRequestBuilder) WorkFromAnywhereMetrics()(*i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.WorkFromAnywhereMetricsRequestBuilder) {
    return i4f5cd14f349081a4fb0b52a2ea5ec7b6c90bc3d3cc26fbbd0891609c4b1b0d07.NewWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
