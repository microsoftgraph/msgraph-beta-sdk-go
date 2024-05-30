package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserItemRequestBuilder provides operations to manage the collection of user entities.
type UserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserItemRequestBuilderGetQueryParameters retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These default properties are noted in the Properties section. To get properties that are not returned by default, do a GET operation for the user and specify the properties in a $select OData query option. Because the user resource supports extensions, you can also use the GET operation to get custom properties and extension data in a user instance. Customers through Microsoft Entra ID for customers can also use this API operation to retrieve their details.
type UserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserItemRequestBuilderGetQueryParameters
}
// UserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.user entity.
// returns a *ItemActivitiesRequestBuilder when successful
func (m *UserItemRequestBuilder) Activities()(*ItemActivitiesRequestBuilder) {
    return NewItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgreementAcceptances provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
// returns a *ItemAgreementacceptancesAgreementAcceptancesRequestBuilder when successful
func (m *UserItemRequestBuilder) AgreementAcceptances()(*ItemAgreementacceptancesAgreementAcceptancesRequestBuilder) {
    return NewItemAgreementacceptancesAgreementAcceptancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.user entity.
// returns a *ItemAnalyticsRequestBuilder when successful
func (m *UserItemRequestBuilder) Analytics()(*ItemAnalyticsRequestBuilder) {
    return NewItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppConsentRequestsForApproval provides operations to manage the appConsentRequestsForApproval property of the microsoft.graph.user entity.
// returns a *ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder when successful
func (m *UserItemRequestBuilder) AppConsentRequestsForApproval()(*ItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalAppConsentRequestsForApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignedResources provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
// returns a *ItemApproleassignedresourcesAppRoleAssignedResourcesRequestBuilder when successful
func (m *UserItemRequestBuilder) AppRoleAssignedResources()(*ItemApproleassignedresourcesAppRoleAssignedResourcesRequestBuilder) {
    return NewItemApproleassignedresourcesAppRoleAssignedResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignedResourcesWithAppId provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
// returns a *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder when successful
func (m *UserItemRequestBuilder) AppRoleAssignedResourcesWithAppId(appId *string)(*ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) {
    return NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.user entity.
// returns a *ItemApproleassignmentsAppRoleAssignmentsRequestBuilder when successful
func (m *UserItemRequestBuilder) AppRoleAssignments()(*ItemApproleassignmentsAppRoleAssignmentsRequestBuilder) {
    return NewItemApproleassignmentsAppRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Approvals provides operations to manage the approvals property of the microsoft.graph.user entity.
// returns a *ItemApprovalsRequestBuilder when successful
func (m *UserItemRequestBuilder) Approvals()(*ItemApprovalsRequestBuilder) {
    return NewItemApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignLicense provides operations to call the assignLicense method.
// returns a *ItemAssignlicenseAssignLicenseRequestBuilder when successful
func (m *UserItemRequestBuilder) AssignLicense()(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    return NewItemAssignlicenseAssignLicenseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Authentication provides operations to manage the authentication property of the microsoft.graph.user entity.
// returns a *ItemAuthenticationRequestBuilder when successful
func (m *UserItemRequestBuilder) Authentication()(*ItemAuthenticationRequestBuilder) {
    return NewItemAuthenticationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.user entity.
// returns a *ItemCalendarRequestBuilder when successful
func (m *UserItemRequestBuilder) Calendar()(*ItemCalendarRequestBuilder) {
    return NewItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CalendarGroups provides operations to manage the calendarGroups property of the microsoft.graph.user entity.
// returns a *ItemCalendargroupsCalendarGroupsRequestBuilder when successful
func (m *UserItemRequestBuilder) CalendarGroups()(*ItemCalendargroupsCalendarGroupsRequestBuilder) {
    return NewItemCalendargroupsCalendarGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendars provides operations to manage the calendars property of the microsoft.graph.user entity.
// returns a *ItemCalendarsRequestBuilder when successful
func (m *UserItemRequestBuilder) Calendars()(*ItemCalendarsRequestBuilder) {
    return NewItemCalendarsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.user entity.
// returns a *ItemCalendarviewCalendarViewRequestBuilder when successful
func (m *UserItemRequestBuilder) CalendarView()(*ItemCalendarviewCalendarViewRequestBuilder) {
    return NewItemCalendarviewCalendarViewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChangePassword provides operations to call the changePassword method.
// returns a *ItemChangepasswordChangePasswordRequestBuilder when successful
func (m *UserItemRequestBuilder) ChangePassword()(*ItemChangepasswordChangePasswordRequestBuilder) {
    return NewItemChangepasswordChangePasswordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Chats provides operations to manage the chats property of the microsoft.graph.user entity.
// returns a *ItemChatsRequestBuilder when successful
func (m *UserItemRequestBuilder) Chats()(*ItemChatsRequestBuilder) {
    return NewItemChatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
// returns a *ItemCheckmembergroupsCheckMemberGroupsRequestBuilder when successful
func (m *UserItemRequestBuilder) CheckMemberGroups()(*ItemCheckmembergroupsCheckMemberGroupsRequestBuilder) {
    return NewItemCheckmembergroupsCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
// returns a *ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder when successful
func (m *UserItemRequestBuilder) CheckMemberObjects()(*ItemCheckmemberobjectsCheckMemberObjectsRequestBuilder) {
    return NewItemCheckmemberobjectsCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudClipboard provides operations to manage the cloudClipboard property of the microsoft.graph.user entity.
// returns a *ItemCloudclipboardCloudClipboardRequestBuilder when successful
func (m *UserItemRequestBuilder) CloudClipboard()(*ItemCloudclipboardCloudClipboardRequestBuilder) {
    return NewItemCloudclipboardCloudClipboardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
// returns a *ItemCloudpcsCloudPCsRequestBuilder when successful
func (m *UserItemRequestBuilder) CloudPCs()(*ItemCloudpcsCloudPCsRequestBuilder) {
    return NewItemCloudpcsCloudPCsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserItemRequestBuilderInternal instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    m := &UserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserItemRequestBuilder instantiates a new UserItemRequestBuilder and sets the default values.
func NewUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders provides operations to manage the contactFolders property of the microsoft.graph.user entity.
// returns a *ItemContactfoldersContactFoldersRequestBuilder when successful
func (m *UserItemRequestBuilder) ContactFolders()(*ItemContactfoldersContactFoldersRequestBuilder) {
    return NewItemContactfoldersContactFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contacts provides operations to manage the contacts property of the microsoft.graph.user entity.
// returns a *ItemContactsRequestBuilder when successful
func (m *UserItemRequestBuilder) Contacts()(*ItemContactsRequestBuilder) {
    return NewItemContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConvertExternalToInternalMemberUser provides operations to call the convertExternalToInternalMemberUser method.
// returns a *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder when successful
func (m *UserItemRequestBuilder) ConvertExternalToInternalMemberUser()(*ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) {
    return NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedObjects provides operations to manage the createdObjects property of the microsoft.graph.user entity.
// returns a *ItemCreatedobjectsCreatedObjectsRequestBuilder when successful
func (m *UserItemRequestBuilder) CreatedObjects()(*ItemCreatedobjectsCreatedObjectsRequestBuilder) {
    return NewItemCreatedobjectsCreatedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-delete?view=graph-rest-beta
func (m *UserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeletePasswordSingleSignOnCredentials provides operations to call the deletePasswordSingleSignOnCredentials method.
// returns a *ItemDeletepasswordsinglesignoncredentialsDeletePasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *UserItemRequestBuilder) DeletePasswordSingleSignOnCredentials()(*ItemDeletepasswordsinglesignoncredentialsDeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemDeletepasswordsinglesignoncredentialsDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
// returns a *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *UserItemRequestBuilder) DeviceEnrollmentConfigurations()(*ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementTroubleshootingEvents provides operations to manage the deviceManagementTroubleshootingEvents property of the microsoft.graph.user entity.
// returns a *ItemDevicemanagementtroubleshootingeventsDeviceManagementTroubleshootingEventsRequestBuilder when successful
func (m *UserItemRequestBuilder) DeviceManagementTroubleshootingEvents()(*ItemDevicemanagementtroubleshootingeventsDeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewItemDevicemanagementtroubleshootingeventsDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Devices provides operations to manage the devices property of the microsoft.graph.user entity.
// returns a *ItemDevicesRequestBuilder when successful
func (m *UserItemRequestBuilder) Devices()(*ItemDevicesRequestBuilder) {
    return NewItemDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DevicesWithDeviceId provides operations to manage the devices property of the microsoft.graph.user entity.
// returns a *ItemDeviceswithdeviceidDevicesWithDeviceIdRequestBuilder when successful
func (m *UserItemRequestBuilder) DevicesWithDeviceId(deviceId *string)(*ItemDeviceswithdeviceidDevicesWithDeviceIdRequestBuilder) {
    return NewItemDeviceswithdeviceidDevicesWithDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, deviceId)
}
// DirectReports provides operations to manage the directReports property of the microsoft.graph.user entity.
// returns a *ItemDirectreportsDirectReportsRequestBuilder when successful
func (m *UserItemRequestBuilder) DirectReports()(*ItemDirectreportsDirectReportsRequestBuilder) {
    return NewItemDirectreportsDirectReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.user entity.
// returns a *ItemDriveRequestBuilder when successful
func (m *UserItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.user entity.
// returns a *ItemDrivesRequestBuilder when successful
func (m *UserItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmployeeExperience provides operations to manage the employeeExperience property of the microsoft.graph.user entity.
// returns a *ItemEmployeeexperienceEmployeeExperienceRequestBuilder when successful
func (m *UserItemRequestBuilder) EmployeeExperience()(*ItemEmployeeexperienceEmployeeExperienceRequestBuilder) {
    return NewItemEmployeeexperienceEmployeeExperienceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events provides operations to manage the events property of the microsoft.graph.user entity.
// returns a *ItemEventsRequestBuilder when successful
func (m *UserItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
// returns a *ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder when successful
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementData()(*ItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilder) {
    return NewItemExportdeviceandappmanagementdataExportDeviceAndAppManagementDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
// returns a *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder when successful
func (m *UserItemRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, skip, top)
}
// ExportPersonalData provides operations to call the exportPersonalData method.
// returns a *ItemExportpersonaldataExportPersonalDataRequestBuilder when successful
func (m *UserItemRequestBuilder) ExportPersonalData()(*ItemExportpersonaldataExportPersonalDataRequestBuilder) {
    return NewItemExportpersonaldataExportPersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.user entity.
// returns a *ItemExtensionsRequestBuilder when successful
func (m *UserItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindMeetingTimes provides operations to call the findMeetingTimes method.
// returns a *ItemFindmeetingtimesFindMeetingTimesRequestBuilder when successful
func (m *UserItemRequestBuilder) FindMeetingTimes()(*ItemFindmeetingtimesFindMeetingTimesRequestBuilder) {
    return NewItemFindmeetingtimesFindMeetingTimesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindRoomLists provides operations to call the findRoomLists method.
// returns a *ItemFindroomlistsFindRoomListsRequestBuilder when successful
func (m *UserItemRequestBuilder) FindRoomLists()(*ItemFindroomlistsFindRoomListsRequestBuilder) {
    return NewItemFindroomlistsFindRoomListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindRooms provides operations to call the findRooms method.
// returns a *ItemFindroomsFindRoomsRequestBuilder when successful
func (m *UserItemRequestBuilder) FindRooms()(*ItemFindroomsFindRoomsRequestBuilder) {
    return NewItemFindroomsFindRoomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
// returns a *ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder when successful
func (m *UserItemRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilder) {
    return NewItemFindroomswithroomlistFindRoomsWithRoomListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, roomList)
}
// FollowedSites provides operations to manage the followedSites property of the microsoft.graph.user entity.
// returns a *ItemFollowedsitesFollowedSitesRequestBuilder when successful
func (m *UserItemRequestBuilder) FollowedSites()(*ItemFollowedsitesFollowedSitesRequestBuilder) {
    return NewItemFollowedsitesFollowedSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These default properties are noted in the Properties section. To get properties that are not returned by default, do a GET operation for the user and specify the properties in a $select OData query option. Because the user resource supports extensions, you can also use the GET operation to get custom properties and extension data in a user instance. Customers through Microsoft Entra ID for customers can also use this API operation to retrieve their details.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-get?view=graph-rest-beta
func (m *UserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
// returns a *ItemGeteffectivedeviceenrollmentconfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ItemGeteffectivedeviceenrollmentconfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewItemGeteffectivedeviceenrollmentconfigurationsGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
// returns a *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder when successful
func (m *UserItemRequestBuilder) GetLoggedOnManagedDevices()(*ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) {
    return NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMailTips provides operations to call the getMailTips method.
// returns a *ItemGetmailtipsGetMailTipsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetMailTips()(*ItemGetmailtipsGetMailTipsRequestBuilder) {
    return NewItemGetmailtipsGetMailTipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
// returns a *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder when successful
func (m *UserItemRequestBuilder) GetManagedAppDiagnosticStatuses()(*ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
// returns a *ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder when successful
func (m *UserItemRequestBuilder) GetManagedAppPolicies()(*ItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilder) {
    return NewItemGetmanagedapppoliciesGetManagedAppPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
// returns a *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder when successful
func (m *UserItemRequestBuilder) GetManagedDevicesWithAppFailures()(*ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
// returns a *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberGroups provides operations to call the getMemberGroups method.
// returns a *ItemGetmembergroupsGetMemberGroupsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetMemberGroups()(*ItemGetmembergroupsGetMemberGroupsRequestBuilder) {
    return NewItemGetmembergroupsGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
// returns a *ItemGetmemberobjectsGetMemberObjectsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetMemberObjects()(*ItemGetmemberobjectsGetMemberObjectsRequestBuilder) {
    return NewItemGetmemberobjectsGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetPasswordSingleSignOnCredentials provides operations to call the getPasswordSingleSignOnCredentials method.
// returns a *ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *UserItemRequestBuilder) GetPasswordSingleSignOnCredentials()(*ItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemGetpasswordsinglesignoncredentialsGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InferenceClassification provides operations to manage the inferenceClassification property of the microsoft.graph.user entity.
// returns a *ItemInferenceclassificationInferenceClassificationRequestBuilder when successful
func (m *UserItemRequestBuilder) InferenceClassification()(*ItemInferenceclassificationInferenceClassificationRequestBuilder) {
    return NewItemInferenceclassificationInferenceClassificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.user entity.
// returns a *ItemInformationprotectionInformationProtectionRequestBuilder when successful
func (m *UserItemRequestBuilder) InformationProtection()(*ItemInformationprotectionInformationProtectionRequestBuilder) {
    return NewItemInformationprotectionInformationProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Insights provides operations to manage the insights property of the microsoft.graph.user entity.
// returns a *ItemInsightsRequestBuilder when successful
func (m *UserItemRequestBuilder) Insights()(*ItemInsightsRequestBuilder) {
    return NewItemInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InvalidateAllRefreshTokens provides operations to call the invalidateAllRefreshTokens method.
// returns a *ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder when successful
func (m *UserItemRequestBuilder) InvalidateAllRefreshTokens()(*ItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilder) {
    return NewItemInvalidateallrefreshtokensInvalidateAllRefreshTokensRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InvitedBy provides operations to manage the invitedBy property of the microsoft.graph.user entity.
// returns a *ItemInvitedbyInvitedByRequestBuilder when successful
func (m *UserItemRequestBuilder) InvitedBy()(*ItemInvitedbyInvitedByRequestBuilder) {
    return NewItemInvitedbyInvitedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
// returns a *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder when successful
func (m *UserItemRequestBuilder) IsManagedAppUserBlocked()(*ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) {
    return NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JoinedGroups provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
// returns a *ItemJoinedgroupsJoinedGroupsRequestBuilder when successful
func (m *UserItemRequestBuilder) JoinedGroups()(*ItemJoinedgroupsJoinedGroupsRequestBuilder) {
    return NewItemJoinedgroupsJoinedGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// JoinedTeams provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
// returns a *ItemJoinedteamsJoinedTeamsRequestBuilder when successful
func (m *UserItemRequestBuilder) JoinedTeams()(*ItemJoinedteamsJoinedTeamsRequestBuilder) {
    return NewItemJoinedteamsJoinedTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LicenseDetails provides operations to manage the licenseDetails property of the microsoft.graph.user entity.
// returns a *ItemLicensedetailsLicenseDetailsRequestBuilder when successful
func (m *UserItemRequestBuilder) LicenseDetails()(*ItemLicensedetailsLicenseDetailsRequestBuilder) {
    return NewItemLicensedetailsLicenseDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MailboxSettings the mailboxSettings property
// returns a *ItemMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *UserItemRequestBuilder) MailboxSettings()(*ItemMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewItemMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MailFolders provides operations to manage the mailFolders property of the microsoft.graph.user entity.
// returns a *ItemMailfoldersMailFoldersRequestBuilder when successful
func (m *UserItemRequestBuilder) MailFolders()(*ItemMailfoldersMailFoldersRequestBuilder) {
    return NewItemMailfoldersMailFoldersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppLogCollectionRequests provides operations to manage the managedAppLogCollectionRequests property of the microsoft.graph.user entity.
// returns a *ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder when successful
func (m *UserItemRequestBuilder) ManagedAppLogCollectionRequests()(*ItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    return NewItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
// returns a *ItemManagedappregistrationsManagedAppRegistrationsRequestBuilder when successful
func (m *UserItemRequestBuilder) ManagedAppRegistrations()(*ItemManagedappregistrationsManagedAppRegistrationsRequestBuilder) {
    return NewItemManagedappregistrationsManagedAppRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.user entity.
// returns a *ItemManageddevicesManagedDevicesRequestBuilder when successful
func (m *UserItemRequestBuilder) ManagedDevices()(*ItemManageddevicesManagedDevicesRequestBuilder) {
    return NewItemManageddevicesManagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Manager provides operations to manage the manager property of the microsoft.graph.user entity.
// returns a *ItemManagerRequestBuilder when successful
func (m *UserItemRequestBuilder) Manager()(*ItemManagerRequestBuilder) {
    return NewItemManagerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.user entity.
// returns a *ItemMemberofMemberOfRequestBuilder when successful
func (m *UserItemRequestBuilder) MemberOf()(*ItemMemberofMemberOfRequestBuilder) {
    return NewItemMemberofMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.user entity.
// returns a *ItemMessagesRequestBuilder when successful
func (m *UserItemRequestBuilder) Messages()(*ItemMessagesRequestBuilder) {
    return NewItemMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppIntentAndStates provides operations to manage the mobileAppIntentAndStates property of the microsoft.graph.user entity.
// returns a *ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder when successful
func (m *UserItemRequestBuilder) MobileAppIntentAndStates()(*ItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilder) {
    return NewItemMobileappintentandstatesMobileAppIntentAndStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.user entity.
// returns a *ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder when successful
func (m *UserItemRequestBuilder) MobileAppTroubleshootingEvents()(*ItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder) {
    return NewItemMobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications provides operations to manage the notifications property of the microsoft.graph.user entity.
// returns a *ItemNotificationsRequestBuilder when successful
func (m *UserItemRequestBuilder) Notifications()(*ItemNotificationsRequestBuilder) {
    return NewItemNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
// returns a *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder when successful
func (m *UserItemRequestBuilder) Oauth2PermissionGrants()(*ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.user entity.
// returns a *ItemOnenoteRequestBuilder when successful
func (m *UserItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnlineMeetings provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
// returns a *ItemOnlinemeetingsOnlineMeetingsRequestBuilder when successful
func (m *UserItemRequestBuilder) OnlineMeetings()(*ItemOnlinemeetingsOnlineMeetingsRequestBuilder) {
    return NewItemOnlinemeetingsOnlineMeetingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnlineMeetingsWithJoinWebUrl provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
// returns a *ItemOnlinemeetingswithjoinweburlOnlineMeetingsWithJoinWebUrlRequestBuilder when successful
func (m *UserItemRequestBuilder) OnlineMeetingsWithJoinWebUrl(joinWebUrl *string)(*ItemOnlinemeetingswithjoinweburlOnlineMeetingsWithJoinWebUrlRequestBuilder) {
    return NewItemOnlinemeetingswithjoinweburlOnlineMeetingsWithJoinWebUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, joinWebUrl)
}
// Outlook provides operations to manage the outlook property of the microsoft.graph.user entity.
// returns a *ItemOutlookRequestBuilder when successful
func (m *UserItemRequestBuilder) Outlook()(*ItemOutlookRequestBuilder) {
    return NewItemOutlookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OwnedDevices provides operations to manage the ownedDevices property of the microsoft.graph.user entity.
// returns a *ItemOwneddevicesOwnedDevicesRequestBuilder when successful
func (m *UserItemRequestBuilder) OwnedDevices()(*ItemOwneddevicesOwnedDevicesRequestBuilder) {
    return NewItemOwneddevicesOwnedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OwnedObjects provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
// returns a *ItemOwnedobjectsOwnedObjectsRequestBuilder when successful
func (m *UserItemRequestBuilder) OwnedObjects()(*ItemOwnedobjectsOwnedObjectsRequestBuilder) {
    return NewItemOwnedobjectsOwnedObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage. Customers through Microsoft Entra ID for customers can also use this API operation to update their details. See Default user permissions in customer tenants for the list of properties they can update.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-update?view=graph-rest-beta
func (m *UserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// PendingAccessReviewInstances provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
// returns a *ItemPendingaccessreviewinstancesPendingAccessReviewInstancesRequestBuilder when successful
func (m *UserItemRequestBuilder) PendingAccessReviewInstances()(*ItemPendingaccessreviewinstancesPendingAccessReviewInstancesRequestBuilder) {
    return NewItemPendingaccessreviewinstancesPendingAccessReviewInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// People provides operations to manage the people property of the microsoft.graph.user entity.
// returns a *ItemPeopleRequestBuilder when successful
func (m *UserItemRequestBuilder) People()(*ItemPeopleRequestBuilder) {
    return NewItemPeopleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.user entity.
// returns a *ItemPermissiongrantsPermissionGrantsRequestBuilder when successful
func (m *UserItemRequestBuilder) PermissionGrants()(*ItemPermissiongrantsPermissionGrantsRequestBuilder) {
    return NewItemPermissiongrantsPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.user entity.
// returns a *ItemPhotoRequestBuilder when successful
func (m *UserItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photos provides operations to manage the photos property of the microsoft.graph.user entity.
// returns a *ItemPhotosRequestBuilder when successful
func (m *UserItemRequestBuilder) Photos()(*ItemPhotosRequestBuilder) {
    return NewItemPhotosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Planner provides operations to manage the planner property of the microsoft.graph.user entity.
// returns a *ItemPlannerRequestBuilder when successful
func (m *UserItemRequestBuilder) Planner()(*ItemPlannerRequestBuilder) {
    return NewItemPlannerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Presence provides operations to manage the presence property of the microsoft.graph.user entity.
// returns a *ItemPresenceRequestBuilder when successful
func (m *UserItemRequestBuilder) Presence()(*ItemPresenceRequestBuilder) {
    return NewItemPresenceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Profile provides operations to manage the profile property of the microsoft.graph.user entity.
// returns a *ItemProfileRequestBuilder when successful
func (m *UserItemRequestBuilder) Profile()(*ItemProfileRequestBuilder) {
    return NewItemProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegisteredDevices provides operations to manage the registeredDevices property of the microsoft.graph.user entity.
// returns a *ItemRegistereddevicesRegisteredDevicesRequestBuilder when successful
func (m *UserItemRequestBuilder) RegisteredDevices()(*ItemRegistereddevicesRegisteredDevicesRequestBuilder) {
    return NewItemRegistereddevicesRegisteredDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
// returns a *ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *UserItemRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewItemReminderviewwithstartdatetimewithenddatetimeReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// RemoveAllDevicesFromManagement provides operations to call the removeAllDevicesFromManagement method.
// returns a *ItemRemovealldevicesfrommanagementRemoveAllDevicesFromManagementRequestBuilder when successful
func (m *UserItemRequestBuilder) RemoveAllDevicesFromManagement()(*ItemRemovealldevicesfrommanagementRemoveAllDevicesFromManagementRequestBuilder) {
    return NewItemRemovealldevicesfrommanagementRemoveAllDevicesFromManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReprocessLicenseAssignment provides operations to call the reprocessLicenseAssignment method.
// returns a *ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder when successful
func (m *UserItemRequestBuilder) ReprocessLicenseAssignment()(*ItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilder) {
    return NewItemReprocesslicenseassignmentReprocessLicenseAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemRestoreRequestBuilder when successful
func (m *UserItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetryServiceProvisioning provides operations to call the retryServiceProvisioning method.
// returns a *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder when successful
func (m *UserItemRequestBuilder) RetryServiceProvisioning()(*ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) {
    return NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RevokeSignInSessions provides operations to call the revokeSignInSessions method.
// returns a *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder when successful
func (m *UserItemRequestBuilder) RevokeSignInSessions()(*ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) {
    return NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScopedRoleMemberOf provides operations to manage the scopedRoleMemberOf property of the microsoft.graph.user entity.
// returns a *ItemScopedrolememberofScopedRoleMemberOfRequestBuilder when successful
func (m *UserItemRequestBuilder) ScopedRoleMemberOf()(*ItemScopedrolememberofScopedRoleMemberOfRequestBuilder) {
    return NewItemScopedrolememberofScopedRoleMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Security provides operations to manage the security property of the microsoft.graph.user entity.
// returns a *ItemSecurityRequestBuilder when successful
func (m *UserItemRequestBuilder) Security()(*ItemSecurityRequestBuilder) {
    return NewItemSecurityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendMail provides operations to call the sendMail method.
// returns a *ItemSendmailSendMailRequestBuilder when successful
func (m *UserItemRequestBuilder) SendMail()(*ItemSendmailSendMailRequestBuilder) {
    return NewItemSendmailSendMailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *ItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *UserItemRequestBuilder) ServiceProvisioningErrors()(*ItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewItemServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.user entity.
// returns a *ItemSettingsRequestBuilder when successful
func (m *UserItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sponsors provides operations to manage the sponsors property of the microsoft.graph.user entity.
// returns a *ItemSponsorsRequestBuilder when successful
func (m *UserItemRequestBuilder) Sponsors()(*ItemSponsorsRequestBuilder) {
    return NewItemSponsorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teamwork provides operations to manage the teamwork property of the microsoft.graph.user entity.
// returns a *ItemTeamworkRequestBuilder when successful
func (m *UserItemRequestBuilder) Teamwork()(*ItemTeamworkRequestBuilder) {
    return NewItemTeamworkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete user.   When deleted, user resources are moved to a temporary container and can be restored within 30 days.  After that time, they are permanently deleted.  To learn more, see deletedItems.
// returns a *RequestInformation when successful
func (m *UserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// Todo provides operations to manage the todo property of the microsoft.graph.user entity.
// returns a *ItemTodoRequestBuilder when successful
func (m *UserItemRequestBuilder) Todo()(*ItemTodoRequestBuilder) {
    return NewItemTodoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve the properties and relationships of user object. This operation returns by default only a subset of the more commonly used properties for each user. These default properties are noted in the Properties section. To get properties that are not returned by default, do a GET operation for the user and specify the properties in a $select OData query option. Because the user resource supports extensions, you can also use the GET operation to get custom properties and extension data in a user instance. Customers through Microsoft Entra ID for customers can also use this API operation to retrieve their details.
// returns a *RequestInformation when successful
func (m *UserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage. Customers through Microsoft Entra ID for customers can also use this API operation to update their details. See Default user permissions in customer tenants for the list of properties they can update.
// returns a *RequestInformation when successful
func (m *UserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
// returns a *ItemTransitivememberofTransitiveMemberOfRequestBuilder when successful
func (m *UserItemRequestBuilder) TransitiveMemberOf()(*ItemTransitivememberofTransitiveMemberOfRequestBuilder) {
    return NewItemTransitivememberofTransitiveMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TransitiveReports provides operations to manage the transitiveReports property of the microsoft.graph.user entity.
// returns a *ItemTransitivereportsTransitiveReportsRequestBuilder when successful
func (m *UserItemRequestBuilder) TransitiveReports()(*ItemTransitivereportsTransitiveReportsRequestBuilder) {
    return NewItemTransitivereportsTransitiveReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TranslateExchangeIds provides operations to call the translateExchangeIds method.
// returns a *ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder when successful
func (m *UserItemRequestBuilder) TranslateExchangeIds()(*ItemTranslateexchangeidsTranslateExchangeIdsRequestBuilder) {
    return NewItemTranslateexchangeidsTranslateExchangeIdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnblockManagedApps provides operations to call the unblockManagedApps method.
// returns a *ItemUnblockmanagedappsUnblockManagedAppsRequestBuilder when successful
func (m *UserItemRequestBuilder) UnblockManagedApps()(*ItemUnblockmanagedappsUnblockManagedAppsRequestBuilder) {
    return NewItemUnblockmanagedappsUnblockManagedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.user entity.
// returns a *ItemUsagerightsUsageRightsRequestBuilder when successful
func (m *UserItemRequestBuilder) UsageRights()(*ItemUsagerightsUsageRightsRequestBuilder) {
    return NewItemUsagerightsUsageRightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VirtualEvents provides operations to manage the virtualEvents property of the microsoft.graph.user entity.
// returns a *ItemVirtualeventsVirtualEventsRequestBuilder when successful
func (m *UserItemRequestBuilder) VirtualEvents()(*ItemVirtualeventsVirtualEventsRequestBuilder) {
    return NewItemVirtualeventsVirtualEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionDeviceRegistrations provides operations to manage the windowsInformationProtectionDeviceRegistrations property of the microsoft.graph.user entity.
// returns a *ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder when successful
func (m *UserItemRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*ItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return NewItemWindowsinformationprotectiondeviceregistrationsWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WipeAndBlockManagedApps provides operations to call the wipeAndBlockManagedApps method.
// returns a *ItemWipeandblockmanagedappsWipeAndBlockManagedAppsRequestBuilder when successful
func (m *UserItemRequestBuilder) WipeAndBlockManagedApps()(*ItemWipeandblockmanagedappsWipeAndBlockManagedAppsRequestBuilder) {
    return NewItemWipeandblockmanagedappsWipeAndBlockManagedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WipeManagedAppRegistrationByDeviceTag provides operations to call the wipeManagedAppRegistrationByDeviceTag method.
// returns a *ItemWipemanagedappregistrationbydevicetagWipeManagedAppRegistrationByDeviceTagRequestBuilder when successful
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ItemWipemanagedappregistrationbydevicetagWipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return NewItemWipemanagedappregistrationbydevicetagWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WipeManagedAppRegistrationsByAzureAdDeviceId provides operations to call the wipeManagedAppRegistrationsByAzureAdDeviceId method.
// returns a *ItemWipemanagedappregistrationsbyazureaddeviceidWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder when successful
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ItemWipemanagedappregistrationsbyazureaddeviceidWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return NewItemWipemanagedappregistrationsbyazureaddeviceidWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WipeManagedAppRegistrationsByDeviceTag provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
// returns a *ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder when successful
func (m *UserItemRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return NewItemWipemanagedappregistrationsbydevicetagWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UserItemRequestBuilder when successful
func (m *UserItemRequestBuilder) WithUrl(rawUrl string)(*UserItemRequestBuilder) {
    return NewUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
