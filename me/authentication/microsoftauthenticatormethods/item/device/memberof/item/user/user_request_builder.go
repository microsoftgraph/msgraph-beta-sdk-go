package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1ab8de1eb467c9403fdd26c1ca1ac89fe97a64cc353b8b63932a8287f0368d99 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i20503999785990d8fd6dcde962ebee93a6d4e134909700f9eacf78bb1330642b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findroomlists"
    i23d2a1d548d7a967cb478877c739f670a72bb906ca8fa228af4ffccd67e9db71 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    i2acbed64f89e0252bf623d0deabb4f8d44c9015e3cc2208251d69d8adfc82bc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/changepassword"
    i2d67144b9d6f7ea68942126d92757f92448b4e913006fbdcac85fa39c1827fb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    i2fff97f8c001173976eb96b250db72ff435bbd4fe465dea1830e732f0273a37d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i3722432354593c87de216db7ece49ea0106696706078993478b55babaea0ddfc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i58109e07a3235070aa94e22918d002ddba39f94931f3a2ebb057a399412fb88e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i5d4070cac4abcaf6cff13f48a6790032ffde4a6cb282435fd168d4036a38f079 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    i5ea7ee79c09a3658520b3b077b1bb299fd543797ff903bbdaf1a4b2e6a0319b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/activateserviceplan"
    i68f5253a99ac01465814f6706134c2b8548bc4f8556515da5afec44a0f5019be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    i6b594e0fad42663232b37e30a840c718eb1940fcf218b7f771682ddefdf6d7d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    i74d25f367d6d1752e91847be7f66c9ecd5c0aea477469d33493abc333b6d5850 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i7c93a370a7a7063dfb191301b0e0e3124824a8c5af08147f5bd175856fac501a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/unblockmanagedapps"
    i874731698078af3a7a9cd34a0eda949de558e0b3c27869e1d7a19266b1abb277 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i87c1668d36db530ad4cf250da671c3595fe0441ede725717a7fb865caa15b469 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getloggedonmanageddevices"
    i8ec6deabe37e30ed4985fb02c03936d5a508821a36836558453dbbcc1b027dbf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    i932b5202b02b670a4d4bfd3d17d0f1d7f02538af91f58c4257231c3045ff4e87 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i94db6b4b4e839d5243418517b823d5acac768b6b9f5bbf95ac11581e1f191f9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findroomswithroomlist"
    i99f2b88a431974d518dfdeec3189ad14d34de69a0a686dd6f5760d6686add00d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i9c6d847407b096db7c3957ae1d18aa4ceed3c44842010f14bfc2d9fa4cd0b9e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i9edb5e48e94d2484497ac27456800802194e074fc82da7f69259c3dc99f4202a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/sendmail"
    i9ffbf53e386788c806074078104a7dc56b295edd8205db185d0cedb111460e20 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    ia6e3d5db0d6e8fa566623986e572c892366d53ec200c9ca33774e6df9abe8829 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    ia9b7ca215cde00bec5b9a69680bcc6de19a2fe1bd1e6c55f8e552f0a4fb4cf5b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
    iaf4fc91e9e7e70011af1d4f65d75d0b9fe88cbba96d3dce20e01cbe58f94b276 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    ib457b9eec7a2454940cd66374496db2d86db2faaab2ba3f538c625cdc9798093 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/restore"
    ic4fc160623b3634617f070b7499501c29224d31a6935a21b8438f653fd69cb5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    ic7f9b9351156c12a0bda4585c0b91d5e71d7679a3123e0f326a91857e819925d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    icb5c129789243567663f102b673c0d243b76d3cb6d43cd296b5a700d329782bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/ismanagedappuserblocked"
    id636c1148f0cdc67b790012c27fc369e036475cacb51666239be85da868c52be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    id94158fb60d932a42654513f7072698596325c6b8a14a5da7b92ea1593d86c1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    iecbf7f5653e9d10695889c732233810d846a1e48b42b78ecf710e99451bad89c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    iedf6fd8c3457275752cb1b7cc2dd0274ba44c7d6c0c3e2f04c4543389123c338 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findrooms"
    if92b87880e35f591c98112ffce4dd8bbeedf3c84f783d05baa52c532253b78ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
)

// UserRequestBuilder casts the previous resource to user.
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.user
type UserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserRequestBuilderGetQueryParameters
}
// ActivateServicePlan the activateServicePlan property
func (m *UserRequestBuilder) ActivateServicePlan()(*i5ea7ee79c09a3658520b3b077b1bb299fd543797ff903bbdaf1a4b2e6a0319b9.ActivateServicePlanRequestBuilder) {
    return i5ea7ee79c09a3658520b3b077b1bb299fd543797ff903bbdaf1a4b2e6a0319b9.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ic4fc160623b3634617f070b7499501c29224d31a6935a21b8438f653fd69cb5d.AssignLicenseRequestBuilder) {
    return ic4fc160623b3634617f070b7499501c29224d31a6935a21b8438f653fd69cb5d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i2acbed64f89e0252bf623d0deabb4f8d44c9015e3cc2208251d69d8adfc82bc6.ChangePasswordRequestBuilder) {
    return i2acbed64f89e0252bf623d0deabb4f8d44c9015e3cc2208251d69d8adfc82bc6.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i58109e07a3235070aa94e22918d002ddba39f94931f3a2ebb057a399412fb88e.CheckMemberGroupsRequestBuilder) {
    return i58109e07a3235070aa94e22918d002ddba39f94931f3a2ebb057a399412fb88e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i3722432354593c87de216db7ece49ea0106696706078993478b55babaea0ddfc.CheckMemberObjectsRequestBuilder) {
    return i3722432354593c87de216db7ece49ea0106696706078993478b55babaea0ddfc.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i5d4070cac4abcaf6cff13f48a6790032ffde4a6cb282435fd168d4036a38f079.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i5d4070cac4abcaf6cff13f48a6790032ffde4a6cb282435fd168d4036a38f079.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i874731698078af3a7a9cd34a0eda949de558e0b3c27869e1d7a19266b1abb277.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i874731698078af3a7a9cd34a0eda949de558e0b3c27869e1d7a19266b1abb277.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i6b594e0fad42663232b37e30a840c718eb1940fcf218b7f771682ddefdf6d7d2.ExportPersonalDataRequestBuilder) {
    return i6b594e0fad42663232b37e30a840c718eb1940fcf218b7f771682ddefdf6d7d2.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*iecbf7f5653e9d10695889c732233810d846a1e48b42b78ecf710e99451bad89c.FindMeetingTimesRequestBuilder) {
    return iecbf7f5653e9d10695889c732233810d846a1e48b42b78ecf710e99451bad89c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i20503999785990d8fd6dcde962ebee93a6d4e134909700f9eacf78bb1330642b.FindRoomListsRequestBuilder) {
    return i20503999785990d8fd6dcde962ebee93a6d4e134909700f9eacf78bb1330642b.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*iedf6fd8c3457275752cb1b7cc2dd0274ba44c7d6c0c3e2f04c4543389123c338.FindRoomsRequestBuilder) {
    return iedf6fd8c3457275752cb1b7cc2dd0274ba44c7d6c0c3e2f04c4543389123c338.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i94db6b4b4e839d5243418517b823d5acac768b6b9f5bbf95ac11581e1f191f9d.FindRoomsWithRoomListRequestBuilder) {
    return i94db6b4b4e839d5243418517b823d5acac768b6b9f5bbf95ac11581e1f191f9d.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i932b5202b02b670a4d4bfd3d17d0f1d7f02538af91f58c4257231c3045ff4e87.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i932b5202b02b670a4d4bfd3d17d0f1d7f02538af91f58c4257231c3045ff4e87.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i87c1668d36db530ad4cf250da671c3595fe0441ede725717a7fb865caa15b469.GetLoggedOnManagedDevicesRequestBuilder) {
    return i87c1668d36db530ad4cf250da671c3595fe0441ede725717a7fb865caa15b469.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ia6e3d5db0d6e8fa566623986e572c892366d53ec200c9ca33774e6df9abe8829.GetMailTipsRequestBuilder) {
    return ia6e3d5db0d6e8fa566623986e572c892366d53ec200c9ca33774e6df9abe8829.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i2d67144b9d6f7ea68942126d92757f92448b4e913006fbdcac85fa39c1827fb7.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i2d67144b9d6f7ea68942126d92757f92448b4e913006fbdcac85fa39c1827fb7.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*iaf4fc91e9e7e70011af1d4f65d75d0b9fe88cbba96d3dce20e01cbe58f94b276.GetManagedAppPoliciesRequestBuilder) {
    return iaf4fc91e9e7e70011af1d4f65d75d0b9fe88cbba96d3dce20e01cbe58f94b276.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i23d2a1d548d7a967cb478877c739f670a72bb906ca8fa228af4ffccd67e9db71.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i23d2a1d548d7a967cb478877c739f670a72bb906ca8fa228af4ffccd67e9db71.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i2fff97f8c001173976eb96b250db72ff435bbd4fe465dea1830e732f0273a37d.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i2fff97f8c001173976eb96b250db72ff435bbd4fe465dea1830e732f0273a37d.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i1ab8de1eb467c9403fdd26c1ca1ac89fe97a64cc353b8b63932a8287f0368d99.GetMemberGroupsRequestBuilder) {
    return i1ab8de1eb467c9403fdd26c1ca1ac89fe97a64cc353b8b63932a8287f0368d99.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i68f5253a99ac01465814f6706134c2b8548bc4f8556515da5afec44a0f5019be.GetMemberObjectsRequestBuilder) {
    return i68f5253a99ac01465814f6706134c2b8548bc4f8556515da5afec44a0f5019be.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *UserRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// InvalidateAllRefreshTokens the invalidateAllRefreshTokens property
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i9ffbf53e386788c806074078104a7dc56b295edd8205db185d0cedb111460e20.InvalidateAllRefreshTokensRequestBuilder) {
    return i9ffbf53e386788c806074078104a7dc56b295edd8205db185d0cedb111460e20.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*icb5c129789243567663f102b673c0d243b76d3cb6d43cd296b5a700d329782bc.IsManagedAppUserBlockedRequestBuilder) {
    return icb5c129789243567663f102b673c0d243b76d3cb6d43cd296b5a700d329782bc.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i99f2b88a431974d518dfdeec3189ad14d34de69a0a686dd6f5760d6686add00d.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i99f2b88a431974d518dfdeec3189ad14d34de69a0a686dd6f5760d6686add00d.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*if92b87880e35f591c98112ffce4dd8bbeedf3c84f783d05baa52c532253b78ea.RemoveAllDevicesFromManagementRequestBuilder) {
    return if92b87880e35f591c98112ffce4dd8bbeedf3c84f783d05baa52c532253b78ea.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i74d25f367d6d1752e91847be7f66c9ecd5c0aea477469d33493abc333b6d5850.ReprocessLicenseAssignmentRequestBuilder) {
    return i74d25f367d6d1752e91847be7f66c9ecd5c0aea477469d33493abc333b6d5850.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ib457b9eec7a2454940cd66374496db2d86db2faaab2ba3f538c625cdc9798093.RestoreRequestBuilder) {
    return ib457b9eec7a2454940cd66374496db2d86db2faaab2ba3f538c625cdc9798093.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i8ec6deabe37e30ed4985fb02c03936d5a508821a36836558453dbbcc1b027dbf.RevokeSignInSessionsRequestBuilder) {
    return i8ec6deabe37e30ed4985fb02c03936d5a508821a36836558453dbbcc1b027dbf.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i9edb5e48e94d2484497ac27456800802194e074fc82da7f69259c3dc99f4202a.SendMailRequestBuilder) {
    return i9edb5e48e94d2484497ac27456800802194e074fc82da7f69259c3dc99f4202a.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ia9b7ca215cde00bec5b9a69680bcc6de19a2fe1bd1e6c55f8e552f0a4fb4cf5b.TranslateExchangeIdsRequestBuilder) {
    return ia9b7ca215cde00bec5b9a69680bcc6de19a2fe1bd1e6c55f8e552f0a4fb4cf5b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i7c93a370a7a7063dfb191301b0e0e3124824a8c5af08147f5bd175856fac501a.UnblockManagedAppsRequestBuilder) {
    return i7c93a370a7a7063dfb191301b0e0e3124824a8c5af08147f5bd175856fac501a.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ic7f9b9351156c12a0bda4585c0b91d5e71d7679a3123e0f326a91857e819925d.WipeAndBlockManagedAppsRequestBuilder) {
    return ic7f9b9351156c12a0bda4585c0b91d5e71d7679a3123e0f326a91857e819925d.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*id94158fb60d932a42654513f7072698596325c6b8a14a5da7b92ea1593d86c1f.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return id94158fb60d932a42654513f7072698596325c6b8a14a5da7b92ea1593d86c1f.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*id636c1148f0cdc67b790012c27fc369e036475cacb51666239be85da868c52be.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return id636c1148f0cdc67b790012c27fc369e036475cacb51666239be85da868c52be.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i9c6d847407b096db7c3957ae1d18aa4ceed3c44842010f14bfc2d9fa4cd0b9e6.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i9c6d847407b096db7c3957ae1d18aa4ceed3c44842010f14bfc2d9fa4cd0b9e6.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
