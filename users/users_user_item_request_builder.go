package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersUserItemRequestBuilder provides operations to manage the collection of user entities.
type UsersUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersUserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
type UsersUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersUserItemRequestBuilderGetQueryParameters
}
// UsersUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateServicePlan provides operations to call the activateServicePlan method.
func (m *UsersUserItemRequestBuilder) ActivateServicePlan()(*UsersItemActivateServicePlanRequestBuilder) {
    return NewUsersItemActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Activities()(*UsersItemActivitiesRequestBuilder) {
    return NewUsersItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ActivitiesById(id string)(*UsersItemActivitiesUserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return NewUsersItemActivitiesUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AgreementAcceptances()(*UsersItemAgreementAcceptancesRequestBuilder) {
    return NewUsersItemAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AgreementAcceptancesById(id string)(*UsersItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewUsersItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Analytics()(*UsersItemAnalyticsRequestBuilder) {
    return NewUsersItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApproval provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppConsentRequestsForApproval()(*UsersItemAppConsentRequestsForApprovalRequestBuilder) {
    return NewUsersItemAppConsentRequestsForApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApprovalById provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppConsentRequestsForApprovalById(id string)(*UsersItemAppConsentRequestsForApprovalAppConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appConsentRequest%2Did"] = id
    }
    return NewUsersItemAppConsentRequestsForApprovalAppConsentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignedResources provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignedResources()(*UsersItemAppRoleAssignedResourcesRequestBuilder) {
    return NewUsersItemAppRoleAssignedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedResourcesById provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignedResourcesById(id string)(*UsersItemAppRoleAssignedResourcesServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return NewUsersItemAppRoleAssignedResourcesServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignments()(*UsersItemAppRoleAssignmentsRequestBuilder) {
    return NewUsersItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) AppRoleAssignmentsById(id string)(*UsersItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewUsersItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Approvals provides operations to manage the approvals property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Approvals()(*UsersItemApprovalsRequestBuilder) {
    return NewUsersItemApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalsById provides operations to manage the approvals property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ApprovalsById(id string)(*UsersItemApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewUsersItemApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense provides operations to call the assignLicense method.
func (m *UsersUserItemRequestBuilder) AssignLicense()(*UsersItemAssignLicenseRequestBuilder) {
    return NewUsersItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Authentication()(*UsersItemAuthenticationRequestBuilder) {
    return NewUsersItemAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Calendar()(*UsersItemCalendarRequestBuilder) {
    return NewUsersItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarGroups()(*UsersItemCalendarGroupsRequestBuilder) {
    return NewUsersItemCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarGroupsById(id string)(*UsersItemCalendarGroupsCalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return NewUsersItemCalendarGroupsCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Calendars()(*UsersItemCalendarsRequestBuilder) {
    return NewUsersItemCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarsById(id string)(*UsersItemCalendarsCalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return NewUsersItemCalendarsCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarView()(*UsersItemCalendarViewRequestBuilder) {
    return NewUsersItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CalendarViewById(id string)(*UsersItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewUsersItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChangePassword provides operations to call the changePassword method.
func (m *UsersUserItemRequestBuilder) ChangePassword()(*UsersItemChangePasswordRequestBuilder) {
    return NewUsersItemChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Chats()(*UsersItemChatsRequestBuilder) {
    return NewUsersItemChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ChatsById(id string)(*UsersItemChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return NewUsersItemChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *UsersUserItemRequestBuilder) CheckMemberGroups()(*UsersItemCheckMemberGroupsRequestBuilder) {
    return NewUsersItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *UsersUserItemRequestBuilder) CheckMemberObjects()(*UsersItemCheckMemberObjectsRequestBuilder) {
    return NewUsersItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CloudPCs()(*UsersItemCloudPCsRequestBuilder) {
    return NewUsersItemCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CloudPCsById(id string)(*UsersItemCloudPCsCloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return NewUsersItemCloudPCsCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUsersUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersUserItemRequestBuilder) {
    m := &UsersUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUsersUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactFolders()(*UsersItemContactFoldersRequestBuilder) {
    return NewUsersItemContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactFoldersById(id string)(*UsersItemContactFoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return NewUsersItemContactFoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Contacts()(*UsersItemContactsRequestBuilder) {
    return NewUsersItemContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ContactsById(id string)(*UsersItemContactsContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return NewUsersItemContactsContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UsersUserItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CreatedObjects()(*UsersItemCreatedObjectsRequestBuilder) {
    return NewUsersItemCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) CreatedObjectsById(id string)(*UsersItemCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UsersUserItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UsersUserItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UsersUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
func (m *UsersUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceEnrollmentConfigurations()(*UsersItemDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewUsersItemDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*UsersItemDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewUsersItemDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*UsersItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewUsersItemDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices provides operations to manage the devices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Devices()(*UsersItemDevicesRequestBuilder) {
    return NewUsersItemDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById provides operations to manage the devices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DevicesById(id string)(*UsersItemDevicesDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return NewUsersItemDevicesDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DirectReports()(*UsersItemDirectReportsRequestBuilder) {
    return NewUsersItemDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DirectReportsById(id string)(*UsersItemDirectReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemDirectReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Drive()(*UsersItemDriveRequestBuilder) {
    return NewUsersItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Drives()(*UsersItemDrivesRequestBuilder) {
    return NewUsersItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) DrivesById(id string)(*UsersItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewUsersItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Events()(*UsersItemEventsRequestBuilder) {
    return NewUsersItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) EventsById(id string)(*UsersItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewUsersItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UsersUserItemRequestBuilder) ExportDeviceAndAppManagementData()(*UsersItemExportDeviceAndAppManagementDataRequestBuilder) {
    return NewUsersItemExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UsersUserItemRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*UsersItemExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewUsersItemExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData provides operations to call the exportPersonalData method.
func (m *UsersUserItemRequestBuilder) ExportPersonalData()(*UsersItemExportPersonalDataRequestBuilder) {
    return NewUsersItemExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Extensions()(*UsersItemExtensionsRequestBuilder) {
    return NewUsersItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ExtensionsById(id string)(*UsersItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewUsersItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *UsersUserItemRequestBuilder) FindMeetingTimes()(*UsersItemFindMeetingTimesRequestBuilder) {
    return NewUsersItemFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UsersUserItemRequestBuilder) FindRoomLists()(*UsersItemFindRoomListsRequestBuilder) {
    return NewUsersItemFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UsersUserItemRequestBuilder) FindRooms()(*UsersItemFindRoomsRequestBuilder) {
    return NewUsersItemFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UsersUserItemRequestBuilder) FindRoomsWithRoomList(roomList *string)(*UsersItemFindRoomsWithRoomListRequestBuilder) {
    return NewUsersItemFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) FollowedSites()(*UsersItemFollowedSitesRequestBuilder) {
    return NewUsersItemFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSitesById provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) FollowedSitesById(id string)(*UsersItemFollowedSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewUsersItemFollowedSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *UsersUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UsersUserItemRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*UsersItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewUsersItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UsersUserItemRequestBuilder) GetLoggedOnManagedDevices()(*UsersItemGetLoggedOnManagedDevicesRequestBuilder) {
    return NewUsersItemGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips provides operations to call the getMailTips method.
func (m *UsersUserItemRequestBuilder) GetMailTips()(*UsersItemGetMailTipsRequestBuilder) {
    return NewUsersItemGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UsersUserItemRequestBuilder) GetManagedAppDiagnosticStatuses()(*UsersItemGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewUsersItemGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UsersUserItemRequestBuilder) GetManagedAppPolicies()(*UsersItemGetManagedAppPoliciesRequestBuilder) {
    return NewUsersItemGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UsersUserItemRequestBuilder) GetManagedDevicesWithAppFailures()(*UsersItemGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewUsersItemGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UsersUserItemRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*UsersItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return NewUsersItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *UsersUserItemRequestBuilder) GetMemberGroups()(*UsersItemGetMemberGroupsRequestBuilder) {
    return NewUsersItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *UsersUserItemRequestBuilder) GetMemberObjects()(*UsersItemGetMemberObjectsRequestBuilder) {
    return NewUsersItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) InferenceClassification()(*UsersItemInferenceClassificationRequestBuilder) {
    return NewUsersItemInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) InformationProtection()(*UsersItemInformationProtectionRequestBuilder) {
    return NewUsersItemInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Insights()(*UsersItemInsightsRequestBuilder) {
    return NewUsersItemInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvalidateAllRefreshTokens provides operations to call the invalidateAllRefreshTokens method.
func (m *UsersUserItemRequestBuilder) InvalidateAllRefreshTokens()(*UsersItemInvalidateAllRefreshTokensRequestBuilder) {
    return NewUsersItemInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UsersUserItemRequestBuilder) IsManagedAppUserBlocked()(*UsersItemIsManagedAppUserBlockedRequestBuilder) {
    return NewUsersItemIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedGroups provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) JoinedGroups()(*UsersItemJoinedGroupsRequestBuilder) {
    return NewUsersItemJoinedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) JoinedTeams()(*UsersItemJoinedTeamsRequestBuilder) {
    return NewUsersItemJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) JoinedTeamsById(id string)(*UsersItemJoinedTeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return NewUsersItemJoinedTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) LicenseDetails()(*UsersItemLicenseDetailsRequestBuilder) {
    return NewUsersItemLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) LicenseDetailsById(id string)(*UsersItemLicenseDetailsLicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return NewUsersItemLicenseDetailsLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MailFolders()(*UsersItemMailFoldersRequestBuilder) {
    return NewUsersItemMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MailFoldersById(id string)(*UsersItemMailFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return NewUsersItemMailFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedAppRegistrations()(*UsersItemManagedAppRegistrationsRequestBuilder) {
    return NewUsersItemManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedAppRegistrationsById(id string)(*UsersItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewUsersItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedDevices()(*UsersItemManagedDevicesRequestBuilder) {
    return NewUsersItemManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ManagedDevicesById(id string)(*UsersItemManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewUsersItemManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Manager()(*UsersItemManagerRequestBuilder) {
    return NewUsersItemManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MemberOf()(*UsersItemMemberOfRequestBuilder) {
    return NewUsersItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MemberOfById(id string)(*UsersItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Messages()(*UsersItemMessagesRequestBuilder) {
    return NewUsersItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MessagesById(id string)(*UsersItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewUsersItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppIntentAndStates provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MobileAppIntentAndStates()(*UsersItemMobileAppIntentAndStatesRequestBuilder) {
    return NewUsersItemMobileAppIntentAndStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppIntentAndStatesById provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MobileAppIntentAndStatesById(id string)(*UsersItemMobileAppIntentAndStatesMobileAppIntentAndStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppIntentAndState%2Did"] = id
    }
    return NewUsersItemMobileAppIntentAndStatesMobileAppIntentAndStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MobileAppTroubleshootingEvents()(*UsersItemMobileAppTroubleshootingEventsRequestBuilder) {
    return NewUsersItemMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*UsersItemMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return NewUsersItemMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notifications provides operations to manage the notifications property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Notifications()(*UsersItemNotificationsRequestBuilder) {
    return NewUsersItemNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationsById provides operations to manage the notifications property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) NotificationsById(id string)(*UsersItemNotificationsNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notification%2Did"] = id
    }
    return NewUsersItemNotificationsNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Oauth2PermissionGrants()(*UsersItemOauth2PermissionGrantsRequestBuilder) {
    return NewUsersItemOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Oauth2PermissionGrantsById(id string)(*UsersItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewUsersItemOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Onenote()(*UsersItemOnenoteRequestBuilder) {
    return NewUsersItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OnlineMeetings()(*UsersItemOnlineMeetingsRequestBuilder) {
    return NewUsersItemOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OnlineMeetingsById(id string)(*UsersItemOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewUsersItemOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Outlook()(*UsersItemOutlookRequestBuilder) {
    return NewUsersItemOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedDevices()(*UsersItemOwnedDevicesRequestBuilder) {
    return NewUsersItemOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevicesById provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedDevicesById(id string)(*UsersItemOwnedDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemOwnedDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedObjects()(*UsersItemOwnedObjectsRequestBuilder) {
    return NewUsersItemOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) OwnedObjectsById(id string)(*UsersItemOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *UsersUserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UsersUserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// PendingAccessReviewInstances provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PendingAccessReviewInstances()(*UsersItemPendingAccessReviewInstancesRequestBuilder) {
    return NewUsersItemPendingAccessReviewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PendingAccessReviewInstancesById provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PendingAccessReviewInstancesById(id string)(*UsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance%2Did"] = id
    }
    return NewUsersItemPendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) People()(*UsersItemPeopleRequestBuilder) {
    return NewUsersItemPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById provides operations to manage the people property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PeopleById(id string)(*UsersItemPeoplePersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return NewUsersItemPeoplePersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Photo()(*UsersItemPhotoRequestBuilder) {
    return NewUsersItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Photos()(*UsersItemPhotosRequestBuilder) {
    return NewUsersItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) PhotosById(id string)(*UsersItemPhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewUsersItemPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Planner()(*UsersItemPlannerRequestBuilder) {
    return NewUsersItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Presence()(*UsersItemPresenceRequestBuilder) {
    return NewUsersItemPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Profile provides operations to manage the profile property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Profile()(*UsersItemProfileRequestBuilder) {
    return NewUsersItemProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) RegisteredDevices()(*UsersItemRegisteredDevicesRequestBuilder) {
    return NewUsersItemRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevicesById provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) RegisteredDevicesById(id string)(*UsersItemRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UsersUserItemRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*UsersItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewUsersItemReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *UsersUserItemRequestBuilder) RemoveAllDevicesFromManagement()(*UsersItemRemoveAllDevicesFromManagementRequestBuilder) {
    return NewUsersItemRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *UsersUserItemRequestBuilder) ReprocessLicenseAssignment()(*UsersItemReprocessLicenseAssignmentRequestBuilder) {
    return NewUsersItemReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *UsersUserItemRequestBuilder) Restore()(*UsersItemRestoreRequestBuilder) {
    return NewUsersItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *UsersUserItemRequestBuilder) RevokeSignInSessions()(*UsersItemRevokeSignInSessionsRequestBuilder) {
    return NewUsersItemRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ScopedRoleMemberOf()(*UsersItemScopedRoleMemberOfRequestBuilder) {
    return NewUsersItemScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) ScopedRoleMemberOfById(id string)(*UsersItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewUsersItemScopedRoleMemberOfScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Security provides operations to manage the security property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Security()(*UsersItemSecurityRequestBuilder) {
    return NewUsersItemSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail provides operations to call the sendMail method.
func (m *UsersUserItemRequestBuilder) SendMail()(*UsersItemSendMailRequestBuilder) {
    return NewUsersItemSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Settings()(*UsersItemSettingsRequestBuilder) {
    return NewUsersItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Teamwork()(*UsersItemTeamworkRequestBuilder) {
    return NewUsersItemTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) Todo()(*UsersItemTodoRequestBuilder) {
    return NewUsersItemTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveMemberOf()(*UsersItemTransitiveMemberOfRequestBuilder) {
    return NewUsersItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveMemberOfById(id string)(*UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveReports provides operations to manage the transitiveReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveReports()(*UsersItemTransitiveReportsRequestBuilder) {
    return NewUsersItemTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveReportsById provides operations to manage the transitiveReports property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) TransitiveReportsById(id string)(*UsersItemTransitiveReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewUsersItemTransitiveReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *UsersUserItemRequestBuilder) TranslateExchangeIds()(*UsersItemTranslateExchangeIdsRequestBuilder) {
    return NewUsersItemTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps provides operations to call the unblockManagedApps method.
func (m *UsersUserItemRequestBuilder) UnblockManagedApps()(*UsersItemUnblockManagedAppsRequestBuilder) {
    return NewUsersItemUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) UsageRights()(*UsersItemUsageRightsRequestBuilder) {
    return NewUsersItemUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) UsageRightsById(id string)(*UsersItemUsageRightsUsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return NewUsersItemUsageRightsUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*UsersItemWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewUsersItemWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrationsById provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
func (m *UsersUserItemRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*UsersItemWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = id
    }
    return NewUsersItemWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WipeAndBlockManagedApps provides operations to call the wipeAndBlockManagedApps method.
func (m *UsersUserItemRequestBuilder) WipeAndBlockManagedApps()(*UsersItemWipeAndBlockManagedAppsRequestBuilder) {
    return NewUsersItemWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
func (m *UsersUserItemRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*UsersItemWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return NewUsersItemWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId provides operations to call the wipeManagedAppRegistrationsByAzureAdDeviceId method.
func (m *UsersUserItemRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*UsersItemWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return NewUsersItemWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *UsersUserItemRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*UsersItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewUsersItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
