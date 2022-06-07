package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i022c0d452ec49edee6125df6e39d7caecd3c17569888777aa531d842c084418e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    i02af864404dff3c17861da3813320579426ecc44c58d3bb3b17b0fbd3e1d313b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    i0654d88d801605b01ed745ba95f410385a18694d8ab5ac0cd3ab7ecec9f01e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    i0b0f58ca3cef5d4b707da758aa89ebe5e642d33e7664d4baede05c2b036e78a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findrooms"
    i18087b2d769e30d062e4bfbf5a728f1d87fabe95b6372d4e6d2dc4cfd627500d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    i1bcabe712d1412e6e39a7c10f6664abb09138aa3dc107a61a3caf71d53777937 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i23d50b8bbeafe9973bc26ec5e501b5d548dfef00fbc6c57399afb4877030c0d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i25849f9396847f2f97be4cd5aa8ee6eeacd432e5535711e9278078b8555fe627 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findroomswithroomlist"
    i2bc9310c2b97241192d642c385dae8591c025685c515a6d59063ff784b93cd61 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    i36090775e1468d6883a75058df077dd175a86ab1ab2a8a6165bbb3e8c5d240ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    i425e208e261b983f5bf0d6c6e3fb02b90ed256ca873fea2c5a03bea8e49f3c23 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    i47b6ddefa636cc83e63bde7431dbcf3917b83b203289696656c52905b1ee4733 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    i593c3769b1101b805fbb4820b26c016e28cff6eb1b21e54e76a3141dc0b66f4c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findroomlists"
    i5e7565983403aeecc6ff764a7b2ce34e7b58dde6093cceb454ed57d938eb2cbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
    i66e7cab32aad3e5f4b2907d81fbe921681e5785a22e4ae5b80b7bfe9c7cbdfd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    i6f22fe029b80aa71638fc9eb8ffb4d4582d2ce38cfb4ae469b09167ca2eb4f6c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/unblockmanagedapps"
    i881beff645ec95b6f47f3635ed88261bb8b95733f832c12be314bc04334ef55b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i92af046883a7c1b9d86405d8e171d7479db93c63deb15f003fc8b4a3d7a0a920 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    i9b55e8cbe8f7eeb345d7ea8cceff088b5a4345018d5a7e2e7d01ee10c425c363 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ia19a70092b31ff4e914cce8070e1720dd1c703ad1a8395521580f223442bbb1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    ia86a5838f6c83b2f233193573334e8d2acc4bb708a176aa42fd65bf971c77546 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
    iaaadf5cd707d110b0008f88e2e8219d0535848d00160d509e1a3761b84217299 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    ib95cfe1d672d1ebfcf0ee6d44f47e38de726f7fcc97fe4eb5b26eefd42dfcd55 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    ibdf30d5075dfa05f23c312bc74328b5802a832d917b3d696fa540a479a464045 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    ibec79873079cb8ab570d07be2f4b61079ebefbb22b5994e95f9341a92e42513b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    ic1f88e2c239f56e75bc9e4817c6b36e5c42b41e6d02d715c8e6fa3a777554fae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    ic6a1e1bd54c2566da73d0619b06de6972a318b6fd92953bb3ee4a0860f1c7e26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    icacf20cfb2bcb7b63679631fa01caf9a278081f5e21ba29104b66078b7fd0e55 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
    icb795c12ce79422a47c4a78a7a39be2d482af35d5fdeea850d1faccb972726c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    id218665f84ae1bc339bb3dfed296a944e294618f1d86778fe93dedf7fcaf4f30 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    id684f49c6cb7759086b5d465af11f869dd0e41a9a7df5239e9854c617bbf8621 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    ie318f8d0f5394064bdd61ab40e7e9b441c803b787dd515d33ac1fe26c96c14f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    ie62172f764c9f27995a87968c7bd29ba9566f12969b2b067c23c9f5e25922db3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/activateserviceplan"
    iebf26ab2d325ed21127edaf11606bf365742c08ee34b46020ffe568700091a33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
    if0a3dbdfdc948835280cd2897785af10cfd920712d107133f8f523fe3ea7f202 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ie62172f764c9f27995a87968c7bd29ba9566f12969b2b067c23c9f5e25922db3.ActivateServicePlanRequestBuilder) {
    return ie62172f764c9f27995a87968c7bd29ba9566f12969b2b067c23c9f5e25922db3.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ibec79873079cb8ab570d07be2f4b61079ebefbb22b5994e95f9341a92e42513b.AssignLicenseRequestBuilder) {
    return ibec79873079cb8ab570d07be2f4b61079ebefbb22b5994e95f9341a92e42513b.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i2bc9310c2b97241192d642c385dae8591c025685c515a6d59063ff784b93cd61.ChangePasswordRequestBuilder) {
    return i2bc9310c2b97241192d642c385dae8591c025685c515a6d59063ff784b93cd61.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i425e208e261b983f5bf0d6c6e3fb02b90ed256ca873fea2c5a03bea8e49f3c23.CheckMemberGroupsRequestBuilder) {
    return i425e208e261b983f5bf0d6c6e3fb02b90ed256ca873fea2c5a03bea8e49f3c23.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i5e7565983403aeecc6ff764a7b2ce34e7b58dde6093cceb454ed57d938eb2cbc.CheckMemberObjectsRequestBuilder) {
    return i5e7565983403aeecc6ff764a7b2ce34e7b58dde6093cceb454ed57d938eb2cbc.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*icacf20cfb2bcb7b63679631fa01caf9a278081f5e21ba29104b66078b7fd0e55.ExportDeviceAndAppManagementDataRequestBuilder) {
    return icacf20cfb2bcb7b63679631fa01caf9a278081f5e21ba29104b66078b7fd0e55.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ic6a1e1bd54c2566da73d0619b06de6972a318b6fd92953bb3ee4a0860f1c7e26.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ic6a1e1bd54c2566da73d0619b06de6972a318b6fd92953bb3ee4a0860f1c7e26.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i0654d88d801605b01ed745ba95f410385a18694d8ab5ac0cd3ab7ecec9f01e07.ExportPersonalDataRequestBuilder) {
    return i0654d88d801605b01ed745ba95f410385a18694d8ab5ac0cd3ab7ecec9f01e07.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ia86a5838f6c83b2f233193573334e8d2acc4bb708a176aa42fd65bf971c77546.FindMeetingTimesRequestBuilder) {
    return ia86a5838f6c83b2f233193573334e8d2acc4bb708a176aa42fd65bf971c77546.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i593c3769b1101b805fbb4820b26c016e28cff6eb1b21e54e76a3141dc0b66f4c.FindRoomListsRequestBuilder) {
    return i593c3769b1101b805fbb4820b26c016e28cff6eb1b21e54e76a3141dc0b66f4c.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i0b0f58ca3cef5d4b707da758aa89ebe5e642d33e7664d4baede05c2b036e78a8.FindRoomsRequestBuilder) {
    return i0b0f58ca3cef5d4b707da758aa89ebe5e642d33e7664d4baede05c2b036e78a8.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i25849f9396847f2f97be4cd5aa8ee6eeacd432e5535711e9278078b8555fe627.FindRoomsWithRoomListRequestBuilder) {
    return i25849f9396847f2f97be4cd5aa8ee6eeacd432e5535711e9278078b8555fe627.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ia19a70092b31ff4e914cce8070e1720dd1c703ad1a8395521580f223442bbb1f.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ia19a70092b31ff4e914cce8070e1720dd1c703ad1a8395521580f223442bbb1f.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*if0a3dbdfdc948835280cd2897785af10cfd920712d107133f8f523fe3ea7f202.GetLoggedOnManagedDevicesRequestBuilder) {
    return if0a3dbdfdc948835280cd2897785af10cfd920712d107133f8f523fe3ea7f202.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i36090775e1468d6883a75058df077dd175a86ab1ab2a8a6165bbb3e8c5d240ba.GetMailTipsRequestBuilder) {
    return i36090775e1468d6883a75058df077dd175a86ab1ab2a8a6165bbb3e8c5d240ba.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i23d50b8bbeafe9973bc26ec5e501b5d548dfef00fbc6c57399afb4877030c0d7.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i23d50b8bbeafe9973bc26ec5e501b5d548dfef00fbc6c57399afb4877030c0d7.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ib95cfe1d672d1ebfcf0ee6d44f47e38de726f7fcc97fe4eb5b26eefd42dfcd55.GetManagedAppPoliciesRequestBuilder) {
    return ib95cfe1d672d1ebfcf0ee6d44f47e38de726f7fcc97fe4eb5b26eefd42dfcd55.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ie318f8d0f5394064bdd61ab40e7e9b441c803b787dd515d33ac1fe26c96c14f8.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ie318f8d0f5394064bdd61ab40e7e9b441c803b787dd515d33ac1fe26c96c14f8.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ibdf30d5075dfa05f23c312bc74328b5802a832d917b3d696fa540a479a464045.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ibdf30d5075dfa05f23c312bc74328b5802a832d917b3d696fa540a479a464045.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*iebf26ab2d325ed21127edaf11606bf365742c08ee34b46020ffe568700091a33.GetMemberGroupsRequestBuilder) {
    return iebf26ab2d325ed21127edaf11606bf365742c08ee34b46020ffe568700091a33.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ic1f88e2c239f56e75bc9e4817c6b36e5c42b41e6d02d715c8e6fa3a777554fae.GetMemberObjectsRequestBuilder) {
    return ic1f88e2c239f56e75bc9e4817c6b36e5c42b41e6d02d715c8e6fa3a777554fae.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i92af046883a7c1b9d86405d8e171d7479db93c63deb15f003fc8b4a3d7a0a920.InvalidateAllRefreshTokensRequestBuilder) {
    return i92af046883a7c1b9d86405d8e171d7479db93c63deb15f003fc8b4a3d7a0a920.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i18087b2d769e30d062e4bfbf5a728f1d87fabe95b6372d4e6d2dc4cfd627500d.IsManagedAppUserBlockedRequestBuilder) {
    return i18087b2d769e30d062e4bfbf5a728f1d87fabe95b6372d4e6d2dc4cfd627500d.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i881beff645ec95b6f47f3635ed88261bb8b95733f832c12be314bc04334ef55b.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i881beff645ec95b6f47f3635ed88261bb8b95733f832c12be314bc04334ef55b.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id684f49c6cb7759086b5d465af11f869dd0e41a9a7df5239e9854c617bbf8621.RemoveAllDevicesFromManagementRequestBuilder) {
    return id684f49c6cb7759086b5d465af11f869dd0e41a9a7df5239e9854c617bbf8621.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i1bcabe712d1412e6e39a7c10f6664abb09138aa3dc107a61a3caf71d53777937.ReprocessLicenseAssignmentRequestBuilder) {
    return i1bcabe712d1412e6e39a7c10f6664abb09138aa3dc107a61a3caf71d53777937.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*icb795c12ce79422a47c4a78a7a39be2d482af35d5fdeea850d1faccb972726c3.RestoreRequestBuilder) {
    return icb795c12ce79422a47c4a78a7a39be2d482af35d5fdeea850d1faccb972726c3.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i66e7cab32aad3e5f4b2907d81fbe921681e5785a22e4ae5b80b7bfe9c7cbdfd2.RevokeSignInSessionsRequestBuilder) {
    return i66e7cab32aad3e5f4b2907d81fbe921681e5785a22e4ae5b80b7bfe9c7cbdfd2.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*id218665f84ae1bc339bb3dfed296a944e294618f1d86778fe93dedf7fcaf4f30.SendMailRequestBuilder) {
    return id218665f84ae1bc339bb3dfed296a944e294618f1d86778fe93dedf7fcaf4f30.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i47b6ddefa636cc83e63bde7431dbcf3917b83b203289696656c52905b1ee4733.TranslateExchangeIdsRequestBuilder) {
    return i47b6ddefa636cc83e63bde7431dbcf3917b83b203289696656c52905b1ee4733.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i6f22fe029b80aa71638fc9eb8ffb4d4582d2ce38cfb4ae469b09167ca2eb4f6c.UnblockManagedAppsRequestBuilder) {
    return i6f22fe029b80aa71638fc9eb8ffb4d4582d2ce38cfb4ae469b09167ca2eb4f6c.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i022c0d452ec49edee6125df6e39d7caecd3c17569888777aa531d842c084418e.WipeAndBlockManagedAppsRequestBuilder) {
    return i022c0d452ec49edee6125df6e39d7caecd3c17569888777aa531d842c084418e.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i02af864404dff3c17861da3813320579426ecc44c58d3bb3b17b0fbd3e1d313b.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i02af864404dff3c17861da3813320579426ecc44c58d3bb3b17b0fbd3e1d313b.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i9b55e8cbe8f7eeb345d7ea8cceff088b5a4345018d5a7e2e7d01ee10c425c363.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i9b55e8cbe8f7eeb345d7ea8cceff088b5a4345018d5a7e2e7d01ee10c425c363.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*iaaadf5cd707d110b0008f88e2e8219d0535848d00160d509e1a3761b84217299.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return iaaadf5cd707d110b0008f88e2e8219d0535848d00160d509e1a3761b84217299.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
