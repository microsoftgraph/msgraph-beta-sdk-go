package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
type DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters a list of Group Policy migration reports.
type DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderInternal instantiates a new GroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder instantiates a new GroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyMigrationReports for deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a list of Group Policy migration reports.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyMigrationReports in deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property groupPolicyMigrationReports for deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a list of Group Policy migration reports.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyMigrationReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable), nil
}
// GroupPolicySettingMappings provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) GroupPolicySettingMappings()(*DeviceManagementGroupPolicyMigrationReportsItemGroupPolicySettingMappingsRequestBuilder) {
    return NewDeviceManagementGroupPolicyMigrationReportsItemGroupPolicySettingMappingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicySettingMappingsById provides operations to manage the groupPolicySettingMappings property of the microsoft.graph.groupPolicyMigrationReport entity.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) GroupPolicySettingMappingsById(id string)(*DeviceManagementGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicySettingMapping%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyMigrationReportsItemGroupPolicySettingMappingsGroupPolicySettingMappingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property groupPolicyMigrationReports in deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyMigrationReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyMigrationReportable), nil
}
// UnsupportedGroupPolicyExtensions provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) UnsupportedGroupPolicyExtensions()(*DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) {
    return NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsupportedGroupPolicyExtensionsById provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
func (m *DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) UnsupportedGroupPolicyExtensionsById(id string)(*DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unsupportedGroupPolicyExtension%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
