package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0dbfc2bad3cb7542ff06abf211d81cb42827d35f5de06f3e81a50365883f5af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i10fda1eb07bc21404c810e6b1fa9aa8edf51c7bbc99e4fa281a80661896fafa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    i15612adccb5cff0ec06c05224dd188eb05b8d8f08665a343592f2568d708eb4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    i2c1939f85c695210be9803438418f41446301b4e4655e897f5f6804335413a73 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/ismanagedappuserblocked"
    i2cdbac76c8915211c11b76e6407e5fe37a8722190461aeb3aeb794f8dd00bd31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findroomlists"
    i30359a61453f2e4b65a531f72793d51691f635abf684ed6e8bd45b3be739221f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i3569ef79b644106d2d160a1304b09fc869192b55a6a2514479f379c3b0e327f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/activateserviceplan"
    i4076b8a27bfd685eb52add0a1afa5a0416d297b43c25be217a66c77e9089f6d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i45e7fe20862623780ac1a743c0cb6a5be55e0fb1f211717c90cc8facef99671b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i51738ffc7f5fe838e0f617782651a02f9af9f398c3602a043c72f2862bbedc16 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i5675ca4b148817ce233b7e5496a898d1163395bab91cc543f9db295fc0d2b303 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    i5777ef9e1d25baa31b09d8eda9c9a78825386126b3641af1ee5aa85d8648c144 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getloggedonmanageddevices"
    i66788b8bc10f7bf47ad24244b543d092c76a6543ad9822e16f852d3fed89de05 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i6ad55677abca193f3ad3ffd23d3e700b8ab68e95c2c0b5c7db2626a930f47c10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    i6de0b6d8a90df9ada3bca0b2443a3dd5704e584425e33b29ff73a37138528b87 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    i7677b06fd442eac88cc422a91d28703ad2bfd7bb4bf5d7395d17937def105e0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    i7a4e882f0057d8fdcb24cec1859e5386cea28ec30f3e46e8f0491b21194606db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    i7dcb64003e21c05ad71a2bf0bad2c931270fffa77979425bc6a1aa46322d26a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findrooms"
    i7e2cf0d357948b8e2f1c30a44cae2d5f32076a864a1bf19b87062bcf88a5e02e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    i7f8ffc64e46ffb15ddde18e5b7b532efea3af814eef86dc8ac588456770d4e41 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i8e4eba9dd96c632c891891dbcf36c1fc2bc3a5ab3a12070e5d2fcbfa25d5f53a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/unblockmanagedapps"
    i972b9ab22869c06c5c2264ba49fc5395ddd4032d98f54f9ce3bfb341107fc41a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i9c14af240ae849321c683984ab11e29e6712d9bcad0c6d727607ea9be9f1b607 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/changepassword"
    ib14afc335bd53dfb5f2e37d3358a5f93300f7daaa35c189476e9ebb02884f73a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    ib420c77a75d9160666a6a85c8d56901c006b3fae980f576c17010e970cb674a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    ic5ca75cb9cc5d0aac3685dfd324765c82fa3d5b93104c09a5db82e89bf4088d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    icb4a2a4a349b49bf84a2e100104030f46b4c14529ca61da34ea302b592018604 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    id2152a747a8ffd67270c06171a5c4c1f5e5d72cec918361967c12e0063709a05 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    id6e2be4217fd3284a0652b8aabf2094213b2831c4aa74edda08bca1e775e2634 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    id806e3454f9eda2067009a5f967cf525bbb31a5f96383b2054f03b624afa5431 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    idb43164c5d18d9d44c40bda57533cf4748cc2c322781931255322222da909094 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findroomswithroomlist"
    ie5adccbf8ee196aa1a1e4e37f8e41a1794cfd96b4c2d1be278d14343a2aa54f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/sendmail"
    ie5da3e2410999f65eb694d6f2403c87565fbc2099f25c13633ee36fbdf9af514 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/restore"
    ie7ed9632a38877ce12029b4d20dff5e18a6d3516b599b25f0a251620196fcb03 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    if854877a0fbf4a04fbf3738da1bd18668b91476c74817ecdbcbe61cbec5ec7e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i3569ef79b644106d2d160a1304b09fc869192b55a6a2514479f379c3b0e327f8.ActivateServicePlanRequestBuilder) {
    return i3569ef79b644106d2d160a1304b09fc869192b55a6a2514479f379c3b0e327f8.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i7e2cf0d357948b8e2f1c30a44cae2d5f32076a864a1bf19b87062bcf88a5e02e.AssignLicenseRequestBuilder) {
    return i7e2cf0d357948b8e2f1c30a44cae2d5f32076a864a1bf19b87062bcf88a5e02e.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i9c14af240ae849321c683984ab11e29e6712d9bcad0c6d727607ea9be9f1b607.ChangePasswordRequestBuilder) {
    return i9c14af240ae849321c683984ab11e29e6712d9bcad0c6d727607ea9be9f1b607.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i0dbfc2bad3cb7542ff06abf211d81cb42827d35f5de06f3e81a50365883f5af2.CheckMemberGroupsRequestBuilder) {
    return i0dbfc2bad3cb7542ff06abf211d81cb42827d35f5de06f3e81a50365883f5af2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i45e7fe20862623780ac1a743c0cb6a5be55e0fb1f211717c90cc8facef99671b.CheckMemberObjectsRequestBuilder) {
    return i45e7fe20862623780ac1a743c0cb6a5be55e0fb1f211717c90cc8facef99671b.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i6ad55677abca193f3ad3ffd23d3e700b8ab68e95c2c0b5c7db2626a930f47c10.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i6ad55677abca193f3ad3ffd23d3e700b8ab68e95c2c0b5c7db2626a930f47c10.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*id6e2be4217fd3284a0652b8aabf2094213b2831c4aa74edda08bca1e775e2634.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return id6e2be4217fd3284a0652b8aabf2094213b2831c4aa74edda08bca1e775e2634.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ic5ca75cb9cc5d0aac3685dfd324765c82fa3d5b93104c09a5db82e89bf4088d0.ExportPersonalDataRequestBuilder) {
    return ic5ca75cb9cc5d0aac3685dfd324765c82fa3d5b93104c09a5db82e89bf4088d0.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*id2152a747a8ffd67270c06171a5c4c1f5e5d72cec918361967c12e0063709a05.FindMeetingTimesRequestBuilder) {
    return id2152a747a8ffd67270c06171a5c4c1f5e5d72cec918361967c12e0063709a05.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i2cdbac76c8915211c11b76e6407e5fe37a8722190461aeb3aeb794f8dd00bd31.FindRoomListsRequestBuilder) {
    return i2cdbac76c8915211c11b76e6407e5fe37a8722190461aeb3aeb794f8dd00bd31.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i7dcb64003e21c05ad71a2bf0bad2c931270fffa77979425bc6a1aa46322d26a3.FindRoomsRequestBuilder) {
    return i7dcb64003e21c05ad71a2bf0bad2c931270fffa77979425bc6a1aa46322d26a3.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*idb43164c5d18d9d44c40bda57533cf4748cc2c322781931255322222da909094.FindRoomsWithRoomListRequestBuilder) {
    return idb43164c5d18d9d44c40bda57533cf4748cc2c322781931255322222da909094.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i7f8ffc64e46ffb15ddde18e5b7b532efea3af814eef86dc8ac588456770d4e41.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i7f8ffc64e46ffb15ddde18e5b7b532efea3af814eef86dc8ac588456770d4e41.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i5777ef9e1d25baa31b09d8eda9c9a78825386126b3641af1ee5aa85d8648c144.GetLoggedOnManagedDevicesRequestBuilder) {
    return i5777ef9e1d25baa31b09d8eda9c9a78825386126b3641af1ee5aa85d8648c144.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i10fda1eb07bc21404c810e6b1fa9aa8edf51c7bbc99e4fa281a80661896fafa2.GetMailTipsRequestBuilder) {
    return i10fda1eb07bc21404c810e6b1fa9aa8edf51c7bbc99e4fa281a80661896fafa2.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id806e3454f9eda2067009a5f967cf525bbb31a5f96383b2054f03b624afa5431.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id806e3454f9eda2067009a5f967cf525bbb31a5f96383b2054f03b624afa5431.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i6de0b6d8a90df9ada3bca0b2443a3dd5704e584425e33b29ff73a37138528b87.GetManagedAppPoliciesRequestBuilder) {
    return i6de0b6d8a90df9ada3bca0b2443a3dd5704e584425e33b29ff73a37138528b87.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i5675ca4b148817ce233b7e5496a898d1163395bab91cc543f9db295fc0d2b303.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i5675ca4b148817ce233b7e5496a898d1163395bab91cc543f9db295fc0d2b303.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i972b9ab22869c06c5c2264ba49fc5395ddd4032d98f54f9ce3bfb341107fc41a.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i972b9ab22869c06c5c2264ba49fc5395ddd4032d98f54f9ce3bfb341107fc41a.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i51738ffc7f5fe838e0f617782651a02f9af9f398c3602a043c72f2862bbedc16.GetMemberGroupsRequestBuilder) {
    return i51738ffc7f5fe838e0f617782651a02f9af9f398c3602a043c72f2862bbedc16.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ie7ed9632a38877ce12029b4d20dff5e18a6d3516b599b25f0a251620196fcb03.GetMemberObjectsRequestBuilder) {
    return ie7ed9632a38877ce12029b4d20dff5e18a6d3516b599b25f0a251620196fcb03.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i7677b06fd442eac88cc422a91d28703ad2bfd7bb4bf5d7395d17937def105e0a.InvalidateAllRefreshTokensRequestBuilder) {
    return i7677b06fd442eac88cc422a91d28703ad2bfd7bb4bf5d7395d17937def105e0a.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i2c1939f85c695210be9803438418f41446301b4e4655e897f5f6804335413a73.IsManagedAppUserBlockedRequestBuilder) {
    return i2c1939f85c695210be9803438418f41446301b4e4655e897f5f6804335413a73.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i66788b8bc10f7bf47ad24244b543d092c76a6543ad9822e16f852d3fed89de05.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i66788b8bc10f7bf47ad24244b543d092c76a6543ad9822e16f852d3fed89de05.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i7a4e882f0057d8fdcb24cec1859e5386cea28ec30f3e46e8f0491b21194606db.RemoveAllDevicesFromManagementRequestBuilder) {
    return i7a4e882f0057d8fdcb24cec1859e5386cea28ec30f3e46e8f0491b21194606db.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i30359a61453f2e4b65a531f72793d51691f635abf684ed6e8bd45b3be739221f.ReprocessLicenseAssignmentRequestBuilder) {
    return i30359a61453f2e4b65a531f72793d51691f635abf684ed6e8bd45b3be739221f.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ie5da3e2410999f65eb694d6f2403c87565fbc2099f25c13633ee36fbdf9af514.RestoreRequestBuilder) {
    return ie5da3e2410999f65eb694d6f2403c87565fbc2099f25c13633ee36fbdf9af514.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ib420c77a75d9160666a6a85c8d56901c006b3fae980f576c17010e970cb674a8.RevokeSignInSessionsRequestBuilder) {
    return ib420c77a75d9160666a6a85c8d56901c006b3fae980f576c17010e970cb674a8.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie5adccbf8ee196aa1a1e4e37f8e41a1794cfd96b4c2d1be278d14343a2aa54f5.SendMailRequestBuilder) {
    return ie5adccbf8ee196aa1a1e4e37f8e41a1794cfd96b4c2d1be278d14343a2aa54f5.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*if854877a0fbf4a04fbf3738da1bd18668b91476c74817ecdbcbe61cbec5ec7e5.TranslateExchangeIdsRequestBuilder) {
    return if854877a0fbf4a04fbf3738da1bd18668b91476c74817ecdbcbe61cbec5ec7e5.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i8e4eba9dd96c632c891891dbcf36c1fc2bc3a5ab3a12070e5d2fcbfa25d5f53a.UnblockManagedAppsRequestBuilder) {
    return i8e4eba9dd96c632c891891dbcf36c1fc2bc3a5ab3a12070e5d2fcbfa25d5f53a.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*icb4a2a4a349b49bf84a2e100104030f46b4c14529ca61da34ea302b592018604.WipeAndBlockManagedAppsRequestBuilder) {
    return icb4a2a4a349b49bf84a2e100104030f46b4c14529ca61da34ea302b592018604.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i15612adccb5cff0ec06c05224dd188eb05b8d8f08665a343592f2568d708eb4b.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i15612adccb5cff0ec06c05224dd188eb05b8d8f08665a343592f2568d708eb4b.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i4076b8a27bfd685eb52add0a1afa5a0416d297b43c25be217a66c77e9089f6d5.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i4076b8a27bfd685eb52add0a1afa5a0416d297b43c25be217a66c77e9089f6d5.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ib14afc335bd53dfb5f2e37d3358a5f93300f7daaa35c189476e9ebb02884f73a.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ib14afc335bd53dfb5f2e37d3358a5f93300f7daaa35c189476e9ebb02884f73a.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
