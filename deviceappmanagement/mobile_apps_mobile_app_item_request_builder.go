package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsMobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppsMobileAppItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsMobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type MobileAppsMobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsMobileAppItemRequestBuilderGetQueryParameters
}
// MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) Assignments()(*MobileAppsItemAssignmentsRequestBuilder) {
    return NewMobileAppsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) AssignmentsById(id string)(*MobileAppsItemAssignmentsMobileAppAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppAssignment%2Did"] = id
    }
    return NewMobileAppsItemAssignmentsMobileAppAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) Categories()(*MobileAppsItemCategoriesRequestBuilder) {
    return NewMobileAppsItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) CategoriesById(id string)(*MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory%2Did"] = id
    }
    return NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewMobileAppsMobileAppItemRequestBuilderInternal instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    m := &MobileAppsMobileAppItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMobileAppsMobileAppItemRequestBuilder instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppsMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsMobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppsMobileAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) DeviceStatuses()(*MobileAppsItemDeviceStatusesRequestBuilder) {
    return NewMobileAppsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) DeviceStatusesById(id string)(*MobileAppsItemDeviceStatusesMobileAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus%2Did"] = id
    }
    return NewMobileAppsItemDeviceStatusesMobileAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the mobile apps.
func (m *MobileAppsMobileAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// InstallSummary provides operations to manage the installSummary property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) InstallSummary()(*MobileAppsItemInstallSummaryRequestBuilder) {
    return NewMobileAppsItemInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *MobileAppsMobileAppItemRequestBuilder) MicrosoftGraphAssign()(*MobileAppsItemMicrosoftGraphAssignAssignRequestBuilder) {
    return NewMobileAppsItemMicrosoftGraphAssignAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetRelatedAppStatesWithUserPrincipalNameWithDeviceId provides operations to call the getRelatedAppStates method.
func (m *MobileAppsMobileAppItemRequestBuilder) MicrosoftGraphGetRelatedAppStatesWithUserPrincipalNameWithDeviceId(deviceId *string, userPrincipalName *string)(*MobileAppsItemMicrosoftGraphGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return NewMobileAppsItemMicrosoftGraphGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, deviceId, userPrincipalName)
}
// MicrosoftGraphManagedMobileLobApp casts the previous resource to managedMobileLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) MicrosoftGraphManagedMobileLobApp()(*MobileAppsItemMicrosoftGraphManagedMobileLobAppManagedMobileLobAppRequestBuilder) {
    return NewMobileAppsItemMicrosoftGraphManagedMobileLobAppManagedMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphMobileLobApp casts the previous resource to mobileLobApp.
func (m *MobileAppsMobileAppItemRequestBuilder) MicrosoftGraphMobileLobApp()(*MobileAppsItemMicrosoftGraphMobileLobAppMobileLobAppRequestBuilder) {
    return NewMobileAppsItemMicrosoftGraphMobileLobAppMobileLobAppRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateRelationships provides operations to call the updateRelationships method.
func (m *MobileAppsMobileAppItemRequestBuilder) MicrosoftGraphUpdateRelationships()(*MobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsRequestBuilder) {
    return NewMobileAppsItemMicrosoftGraphUpdateRelationshipsUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppsMobileAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) Relationships()(*MobileAppsItemRelationshipsRequestBuilder) {
    return NewMobileAppsItemRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RelationshipsById provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) RelationshipsById(id string)(*MobileAppsItemRelationshipsMobileAppRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppRelationship%2Did"] = id
    }
    return NewMobileAppsItemRelationshipsMobileAppRelationshipItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppsMobileAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the mobile apps.
func (m *MobileAppsMobileAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsMobileAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppsMobileAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, requestConfiguration *MobileAppsMobileAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) UserStatuses()(*MobileAppsItemUserStatusesRequestBuilder) {
    return NewMobileAppsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsMobileAppItemRequestBuilder) UserStatusesById(id string)(*MobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAppInstallStatus%2Did"] = id
    }
    return NewMobileAppsItemUserStatusesUserAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
