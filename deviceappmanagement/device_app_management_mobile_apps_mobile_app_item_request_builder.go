package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementMobileAppsMobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type DeviceAppManagementMobileAppsMobileAppItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementMobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementMobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementMobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Assign()(*DeviceAppManagementMobileAppsItemAssignRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Assignments()(*DeviceAppManagementMobileAppsItemAssignmentsRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) AssignmentsById(id string)(*DeviceAppManagementMobileAppsItemAssignmentsMobileAppAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppAssignment%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemAssignmentsMobileAppAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Categories()(*DeviceAppManagementMobileAppsItemCategoriesRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) CategoriesById(id string)(*DeviceAppManagementMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilderInternal instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) {
    m := &DeviceAppManagementMobileAppsMobileAppItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilder instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the mobile apps.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mobileApps in deviceAppManagement
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property mobileApps for deviceAppManagement
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) DeviceStatuses()(*DeviceAppManagementMobileAppsItemDeviceStatusesRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) DeviceStatusesById(id string)(*DeviceAppManagementMobileAppsItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the mobile apps.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceId provides operations to call the getRelatedAppStates method.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(deviceId *string, userPrincipalName *string)(*DeviceAppManagementMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, deviceId, userPrincipalName);
}
// InstallSummary provides operations to manage the installSummary property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) InstallSummary()(*DeviceAppManagementMobileAppsItemInstallSummaryRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedMobileLobApp casts the previous resource to managedMobileLobApp.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) ManagedMobileLobApp()(*DeviceAppManagementMobileAppsItemManagedMobileLobAppRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemManagedMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileLobApp casts the previous resource to mobileLobApp.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) MobileLobApp()(*DeviceAppManagementMobileAppsItemMobileLobAppRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property mobileApps in deviceAppManagement
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *DeviceAppManagementMobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) Relationships()(*DeviceAppManagementMobileAppsItemRelationshipsRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationshipsById provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) RelationshipsById(id string)(*DeviceAppManagementMobileAppsItemRelationshipsMobileAppRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppRelationship%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemRelationshipsMobileAppRelationshipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateRelationships provides operations to call the updateRelationships method.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) UpdateRelationships()(*DeviceAppManagementMobileAppsItemUpdateRelationshipsRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) UserStatuses()(*DeviceAppManagementMobileAppsItemUserStatusesRequestBuilder) {
    return NewDeviceAppManagementMobileAppsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.mobileApp entity.
func (m *DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) UserStatusesById(id string)(*DeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAppInstallStatus%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
