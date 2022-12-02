package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeRequestBuilder provides operations to manage the user singleton.
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeRequestBuilderGetQueryParameters retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeRequestBuilderGetQueryParameters
}
// MeRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateServicePlan provides operations to call the activateServicePlan method.
func (m *MeRequestBuilder) ActivateServicePlan()(*MeActivateServicePlanRequestBuilder) {
    return NewMeActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Activities()(*MeActivitiesRequestBuilder) {
    return NewMeActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ActivitiesById(id string)(*MeActivitiesUserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return NewMeActivitiesUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AgreementAcceptances()(*MeAgreementAcceptancesRequestBuilder) {
    return NewMeAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AgreementAcceptancesById(id string)(*MeAgreementAcceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return NewMeAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Analytics()(*MeAnalyticsRequestBuilder) {
    return NewMeAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApproval provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppConsentRequestsForApproval()(*MeAppConsentRequestsForApprovalRequestBuilder) {
    return NewMeAppConsentRequestsForApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApprovalById provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppConsentRequestsForApprovalById(id string)(*MeAppConsentRequestsForApprovalAppConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appConsentRequest%2Did"] = id
    }
    return NewMeAppConsentRequestsForApprovalAppConsentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignedResources provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignedResources()(*MeAppRoleAssignedResourcesRequestBuilder) {
    return NewMeAppRoleAssignedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedResourcesById provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignedResourcesById(id string)(*MeAppRoleAssignedResourcesServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return NewMeAppRoleAssignedResourcesServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignments()(*MeAppRoleAssignmentsRequestBuilder) {
    return NewMeAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) AppRoleAssignmentsById(id string)(*MeAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewMeAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Approvals provides operations to manage the approvals property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Approvals()(*MeApprovalsRequestBuilder) {
    return NewMeApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalsById provides operations to manage the approvals property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ApprovalsById(id string)(*MeApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewMeApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense provides operations to call the assignLicense method.
func (m *MeRequestBuilder) AssignLicense()(*MeAssignLicenseRequestBuilder) {
    return NewMeAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Authentication()(*MeAuthenticationRequestBuilder) {
    return NewMeAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Calendar()(*MeCalendarRequestBuilder) {
    return NewMeCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarGroups()(*MeCalendarGroupsRequestBuilder) {
    return NewMeCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarGroupsById(id string)(*MeCalendarGroupsCalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return NewMeCalendarGroupsCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Calendars()(*MeCalendarsRequestBuilder) {
    return NewMeCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById provides operations to manage the calendars property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarsById(id string)(*MeCalendarsCalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return NewMeCalendarsCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarView()(*MeCalendarViewRequestBuilder) {
    return NewMeCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CalendarViewById(id string)(*MeCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewMeCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChangePassword provides operations to call the changePassword method.
func (m *MeRequestBuilder) ChangePassword()(*MeChangePasswordRequestBuilder) {
    return NewMeChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Chats()(*MeChatsRequestBuilder) {
    return NewMeChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById provides operations to manage the chats property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ChatsById(id string)(*MeChatsChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return NewMeChatsChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *MeRequestBuilder) CheckMemberGroups()(*MeCheckMemberGroupsRequestBuilder) {
    return NewMeCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *MeRequestBuilder) CheckMemberObjects()(*MeCheckMemberObjectsRequestBuilder) {
    return NewMeCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CloudPCs()(*MeCloudPCsRequestBuilder) {
    return NewMeCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CloudPCsById(id string)(*MeCloudPCsCloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return NewMeCloudPCsCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactFolders()(*MeContactFoldersRequestBuilder) {
    return NewMeContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById provides operations to manage the contactFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactFoldersById(id string)(*MeContactFoldersContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return NewMeContactFoldersContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Contacts()(*MeContactsRequestBuilder) {
    return NewMeContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the contacts property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ContactsById(id string)(*MeContactsContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return NewMeContactsContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CreatedObjects()(*MeCreatedObjectsRequestBuilder) {
    return NewMeCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById provides operations to manage the createdObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) CreatedObjectsById(id string)(*MeCreatedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeCreatedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *MeRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceEnrollmentConfigurations()(*MeDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewMeDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*MeDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return NewMeDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEvents()(*MeDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewMeDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*MeDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewMeDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices provides operations to manage the devices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Devices()(*MeDevicesRequestBuilder) {
    return NewMeDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById provides operations to manage the devices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DevicesById(id string)(*MeDevicesDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return NewMeDevicesDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DirectReports()(*MeDirectReportsRequestBuilder) {
    return NewMeDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById provides operations to manage the directReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DirectReportsById(id string)(*MeDirectReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeDirectReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Drive()(*MeDriveRequestBuilder) {
    return NewMeDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Drives()(*MeDrivesRequestBuilder) {
    return NewMeDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) DrivesById(id string)(*MeDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewMeDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Events()(*MeEventsRequestBuilder) {
    return NewMeEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) EventsById(id string)(*MeEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewMeEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) ExportDeviceAndAppManagementData()(*MeExportDeviceAndAppManagementDataRequestBuilder) {
    return NewMeExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*MeExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewMeExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData provides operations to call the exportPersonalData method.
func (m *MeRequestBuilder) ExportPersonalData()(*MeExportPersonalDataRequestBuilder) {
    return NewMeExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Extensions()(*MeExtensionsRequestBuilder) {
    return NewMeExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ExtensionsById(id string)(*MeExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewMeExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindMeetingTimes provides operations to call the findMeetingTimes method.
func (m *MeRequestBuilder) FindMeetingTimes()(*MeFindMeetingTimesRequestBuilder) {
    return NewMeFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *MeRequestBuilder) FindRoomLists()(*MeFindRoomListsRequestBuilder) {
    return NewMeFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *MeRequestBuilder) FindRooms()(*MeFindRoomsRequestBuilder) {
    return NewMeFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *MeRequestBuilder) FindRoomsWithRoomList(roomList *string)(*MeFindRoomsWithRoomListRequestBuilder) {
    return NewMeFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) FollowedSites()(*MeFollowedSitesRequestBuilder) {
    return NewMeFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSitesById provides operations to manage the followedSites property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) FollowedSitesById(id string)(*MeFollowedSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewMeFollowedSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option. Because the **user** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **user** instance.
func (m *MeRequestBuilder) Get(ctx context.Context, requestConfiguration *MeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
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
func (m *MeRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*MeGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewMeGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *MeRequestBuilder) GetLoggedOnManagedDevices()(*MeGetLoggedOnManagedDevicesRequestBuilder) {
    return NewMeGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips provides operations to call the getMailTips method.
func (m *MeRequestBuilder) GetMailTips()(*MeGetMailTipsRequestBuilder) {
    return NewMeGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *MeRequestBuilder) GetManagedAppDiagnosticStatuses()(*MeGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewMeGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *MeRequestBuilder) GetManagedAppPolicies()(*MeGetManagedAppPoliciesRequestBuilder) {
    return NewMeGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *MeRequestBuilder) GetManagedDevicesWithAppFailures()(*MeGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewMeGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *MeRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*MeGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return NewMeGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *MeRequestBuilder) GetMemberGroups()(*MeGetMemberGroupsRequestBuilder) {
    return NewMeGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *MeRequestBuilder) GetMemberObjects()(*MeGetMemberObjectsRequestBuilder) {
    return NewMeGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) InferenceClassification()(*MeInferenceClassificationRequestBuilder) {
    return NewMeInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) InformationProtection()(*MeInformationProtectionRequestBuilder) {
    return NewMeInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Insights()(*MeInsightsRequestBuilder) {
    return NewMeInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvalidateAllRefreshTokens provides operations to call the invalidateAllRefreshTokens method.
func (m *MeRequestBuilder) InvalidateAllRefreshTokens()(*MeInvalidateAllRefreshTokensRequestBuilder) {
    return NewMeInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *MeRequestBuilder) IsManagedAppUserBlocked()(*MeIsManagedAppUserBlockedRequestBuilder) {
    return NewMeIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedGroups provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) JoinedGroups()(*MeJoinedGroupsRequestBuilder) {
    return NewMeJoinedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) JoinedTeams()(*MeJoinedTeamsRequestBuilder) {
    return NewMeJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) JoinedTeamsById(id string)(*MeJoinedTeamsTeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return NewMeJoinedTeamsTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) LicenseDetails()(*MeLicenseDetailsRequestBuilder) {
    return NewMeLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) LicenseDetailsById(id string)(*MeLicenseDetailsLicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return NewMeLicenseDetailsLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MailFolders()(*MeMailFoldersRequestBuilder) {
    return NewMeMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById provides operations to manage the mailFolders property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MailFoldersById(id string)(*MeMailFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return NewMeMailFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedAppRegistrations()(*MeManagedAppRegistrationsRequestBuilder) {
    return NewMeManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedAppRegistrationsById(id string)(*MeManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewMeManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedDevices()(*MeManagedDevicesRequestBuilder) {
    return NewMeManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ManagedDevicesById(id string)(*MeManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewMeManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Manager()(*MeManagerRequestBuilder) {
    return NewMeManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MemberOf()(*MeMemberOfRequestBuilder) {
    return NewMeMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MemberOfById(id string)(*MeMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Messages()(*MeMessagesRequestBuilder) {
    return NewMeMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MessagesById(id string)(*MeMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewMeMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppIntentAndStates provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MobileAppIntentAndStates()(*MeMobileAppIntentAndStatesRequestBuilder) {
    return NewMeMobileAppIntentAndStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppIntentAndStatesById provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MobileAppIntentAndStatesById(id string)(*MeMobileAppIntentAndStatesMobileAppIntentAndStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppIntentAndState%2Did"] = id
    }
    return NewMeMobileAppIntentAndStatesMobileAppIntentAndStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MobileAppTroubleshootingEvents()(*MeMobileAppTroubleshootingEventsRequestBuilder) {
    return NewMeMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*MeMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return NewMeMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notifications provides operations to manage the notifications property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Notifications()(*MeNotificationsRequestBuilder) {
    return NewMeNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationsById provides operations to manage the notifications property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) NotificationsById(id string)(*MeNotificationsNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notification%2Did"] = id
    }
    return NewMeNotificationsNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Oauth2PermissionGrants()(*MeOauth2PermissionGrantsRequestBuilder) {
    return NewMeOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Oauth2PermissionGrantsById(id string)(*MeOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return NewMeOauth2PermissionGrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Onenote()(*MeOnenoteRequestBuilder) {
    return NewMeOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OnlineMeetings()(*MeOnlineMeetingsRequestBuilder) {
    return NewMeOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OnlineMeetingsById(id string)(*MeOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return NewMeOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Outlook()(*MeOutlookRequestBuilder) {
    return NewMeOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedDevices()(*MeOwnedDevicesRequestBuilder) {
    return NewMeOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevicesById provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedDevicesById(id string)(*MeOwnedDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeOwnedDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedObjects()(*MeOwnedObjectsRequestBuilder) {
    return NewMeOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) OwnedObjectsById(id string)(*MeOwnedObjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeOwnedObjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
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
func (m *MeRequestBuilder) PendingAccessReviewInstances()(*MePendingAccessReviewInstancesRequestBuilder) {
    return NewMePendingAccessReviewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PendingAccessReviewInstancesById provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) PendingAccessReviewInstancesById(id string)(*MePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance%2Did"] = id
    }
    return NewMePendingAccessReviewInstancesAccessReviewInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) People()(*MePeopleRequestBuilder) {
    return NewMePeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById provides operations to manage the people property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) PeopleById(id string)(*MePeoplePersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return NewMePeoplePersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Photo()(*MePhotoRequestBuilder) {
    return NewMePhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Photos()(*MePhotosRequestBuilder) {
    return NewMePhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) PhotosById(id string)(*MePhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewMePhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Planner()(*MePlannerRequestBuilder) {
    return NewMePlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Presence()(*MePresenceRequestBuilder) {
    return NewMePresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Profile provides operations to manage the profile property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Profile()(*MeProfileRequestBuilder) {
    return NewMeProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) RegisteredDevices()(*MeRegisteredDevicesRequestBuilder) {
    return NewMeRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevicesById provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) RegisteredDevicesById(id string)(*MeRegisteredDevicesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeRegisteredDevicesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *MeRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*MeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
func (m *MeRequestBuilder) RemoveAllDevicesFromManagement()(*MeRemoveAllDevicesFromManagementRequestBuilder) {
    return NewMeRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
func (m *MeRequestBuilder) ReprocessLicenseAssignment()(*MeReprocessLicenseAssignmentRequestBuilder) {
    return NewMeReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *MeRequestBuilder) Restore()(*MeRestoreRequestBuilder) {
    return NewMeRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions provides operations to call the revokeSignInSessions method.
func (m *MeRequestBuilder) RevokeSignInSessions()(*MeRevokeSignInSessionsRequestBuilder) {
    return NewMeRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ScopedRoleMemberOf()(*MeScopedRoleMemberOfRequestBuilder) {
    return NewMeScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) ScopedRoleMemberOfById(id string)(*MeScopedRoleMemberOfScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return NewMeScopedRoleMemberOfScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Security provides operations to manage the security property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Security()(*MeSecurityRequestBuilder) {
    return NewMeSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail provides operations to call the sendMail method.
func (m *MeRequestBuilder) SendMail()(*MeSendMailRequestBuilder) {
    return NewMeSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Settings()(*MeSettingsRequestBuilder) {
    return NewMeSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Teamwork()(*MeTeamworkRequestBuilder) {
    return NewMeTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) Todo()(*MeTodoRequestBuilder) {
    return NewMeTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveMemberOf()(*MeTransitiveMemberOfRequestBuilder) {
    return NewMeTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveMemberOfById(id string)(*MeTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveReports provides operations to manage the transitiveReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveReports()(*MeTransitiveReportsRequestBuilder) {
    return NewMeTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveReportsById provides operations to manage the transitiveReports property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) TransitiveReportsById(id string)(*MeTransitiveReportsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewMeTransitiveReportsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TranslateExchangeIds provides operations to call the translateExchangeIds method.
func (m *MeRequestBuilder) TranslateExchangeIds()(*MeTranslateExchangeIdsRequestBuilder) {
    return NewMeTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps provides operations to call the unblockManagedApps method.
func (m *MeRequestBuilder) UnblockManagedApps()(*MeUnblockManagedAppsRequestBuilder) {
    return NewMeUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) UsageRights()(*MeUsageRightsRequestBuilder) {
    return NewMeUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) UsageRightsById(id string)(*MeUsageRightsUsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return NewMeUsageRightsUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*MeWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewMeWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrationsById provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
func (m *MeRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*MeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = id
    }
    return NewMeWindowsInformationProtectionDeviceRegistrationsWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WipeAndBlockManagedApps provides operations to call the wipeAndBlockManagedApps method.
func (m *MeRequestBuilder) WipeAndBlockManagedApps()(*MeWipeAndBlockManagedAppsRequestBuilder) {
    return NewMeWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
func (m *MeRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*MeWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return NewMeWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId provides operations to call the wipeManagedAppRegistrationsByAzureAdDeviceId method.
func (m *MeRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*MeWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return NewMeWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
func (m *MeRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*MeWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewMeWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
