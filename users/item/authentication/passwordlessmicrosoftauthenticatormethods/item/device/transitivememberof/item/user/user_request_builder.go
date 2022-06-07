package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c3d4631adee6ee334df3c6e9f5a36de93d62f420e6eba8f7476c4f91bed511b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i0e14ba548bf134c852466a62a4ea1fdd068045e4b6a000c378d6bcc1036f4720 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    i179d26a094b7b1d92d392c7f52f00f352bf8d29d756a2907159fb38cff2bb341 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    i2145c20e77a05a5a987c631004f6919cc200a7e13ad43d2d463c720b2e8ed1e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i378ed81edc5f7659807ebcbbf6a15071b81eabea60b4bc9c3625c1cd920d25fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    i3e682571c730b8fbcaea5127e4fc5098edbdce95dc229bdcc6bc6a839fd13067 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i404437b92d38e2367e81306137b1805808cbdfd165d826dbc0d8d251fe2bf224 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i4f34b5b3b3f8fcf418f7bee5098aa39404110c83f3d6fe1000e68308a71dc984 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i57629ef934145c095878a194c2f009f07ae9d69635c8eb33e975294c332e5fdc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    i58f250010bcb3e430341c3372d66bb0543559166175f0b52f56e94f1d3626798 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i5d699944287eba5958f97b66df6d01b55fad4e28d25f8e3a9895a6ebc1f482c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i5e3457f7caf996c0212765e2462778b9f943d406bad548edf55068d3cb1a8b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findroomswithroomlist"
    i5f18baa5f3b33845bd1da4c0980d0acdc9ad32b0d274fe94ef241d8be4f14fc9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i64119174a682421d989bc15b022d48410683e668e982bebcfcfb570c2aa205d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    i715b042262dc9ed949747092d1a6d4c67f4e6a17af31f9c7d5009c8180d3a283 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    i77d820470b2b321323755cf8a59aa8a83c554a16214617b7463a42a6c6382205 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
    i7beca4b44ab53aa70feabd61a1f01b48ab66833b829c07de5894a814582516bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findrooms"
    i7dc39389e361fe6c068589d6472982f56cd01554a9eacdccf5546d58b61b22ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    i7e177c2b30d06eca5b8679486a190b47f61d08036e99cccd348987edfe106d19 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    i807f7ad94758834183333ac828bb2d2ba143442b51cc512e1796d0c1624f9383 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    i86831a89b833bd7a67b83a445c0954fe2810bae67c93b2d3de51ef0d91e57ec3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i875bacf36ec6afe008f45e95ce4ff1d674accafe1f7db527c9e4d059a43709d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    i89ebf0a17ea6af40cd498795effab2b40c1a728fbdc400a8dbb41122a24201bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
    i8e39ed91fefafe58a2e1e52e3c17f716582519ef0010019ce14c17b2fa6be426 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    ia5e25e2dd47557a9550abfd9f8b951c523317ba4d610c80f4b74304eb964e8f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    ia609a7fd58b7a403a148c3e7d4989441e7e2f4277f2900a7b02955a2cc5edb85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ia9574722694206daffb29cb76cfc618e6a99baa022c386324281e62b87cdfce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    ib1fb854c6f957236ec71f510bab72e4ae2b47ad7f8210f5fd9f829a9ec5ae510 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    ib2e90232be0b004b36af22bbbfd5f6c2764a30d271eb5a375b4237339bef40ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ib90ea9c2182aa8570683ef29aef2a9cdea47325a191baec33141a128eead4483 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    iba313d84dd503b98b839bd7c2c10b20446a8a6080cdb37c526cba11b243f7465 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findroomlists"
    ic69660e1409fcdd5572b77a61e5a0018c3b59a5b1a97637119bb9a9d012a3737 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    icd42f5a1297266f49fbd72565398db0b3a3c7e1ed877cd9e88e57e7c3097a1a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    id6870ea3aae3d9653c9ffdb8552acd8a7d40c2008785484d039b60cf31e2af06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/activateserviceplan"
    if2dd7619e101a0dd17d05cdbd37246f44c7f3812d254664e7b7893c13655489c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*id6870ea3aae3d9653c9ffdb8552acd8a7d40c2008785484d039b60cf31e2af06.ActivateServicePlanRequestBuilder) {
    return id6870ea3aae3d9653c9ffdb8552acd8a7d40c2008785484d039b60cf31e2af06.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ib90ea9c2182aa8570683ef29aef2a9cdea47325a191baec33141a128eead4483.AssignLicenseRequestBuilder) {
    return ib90ea9c2182aa8570683ef29aef2a9cdea47325a191baec33141a128eead4483.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i179d26a094b7b1d92d392c7f52f00f352bf8d29d756a2907159fb38cff2bb341.ChangePasswordRequestBuilder) {
    return i179d26a094b7b1d92d392c7f52f00f352bf8d29d756a2907159fb38cff2bb341.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ia9574722694206daffb29cb76cfc618e6a99baa022c386324281e62b87cdfce3.CheckMemberGroupsRequestBuilder) {
    return ia9574722694206daffb29cb76cfc618e6a99baa022c386324281e62b87cdfce3.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i0e14ba548bf134c852466a62a4ea1fdd068045e4b6a000c378d6bcc1036f4720.CheckMemberObjectsRequestBuilder) {
    return i0e14ba548bf134c852466a62a4ea1fdd068045e4b6a000c378d6bcc1036f4720.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i807f7ad94758834183333ac828bb2d2ba143442b51cc512e1796d0c1624f9383.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i807f7ad94758834183333ac828bb2d2ba143442b51cc512e1796d0c1624f9383.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ib2e90232be0b004b36af22bbbfd5f6c2764a30d271eb5a375b4237339bef40ed.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ib2e90232be0b004b36af22bbbfd5f6c2764a30d271eb5a375b4237339bef40ed.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i5f18baa5f3b33845bd1da4c0980d0acdc9ad32b0d274fe94ef241d8be4f14fc9.ExportPersonalDataRequestBuilder) {
    return i5f18baa5f3b33845bd1da4c0980d0acdc9ad32b0d274fe94ef241d8be4f14fc9.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i58f250010bcb3e430341c3372d66bb0543559166175f0b52f56e94f1d3626798.FindMeetingTimesRequestBuilder) {
    return i58f250010bcb3e430341c3372d66bb0543559166175f0b52f56e94f1d3626798.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*iba313d84dd503b98b839bd7c2c10b20446a8a6080cdb37c526cba11b243f7465.FindRoomListsRequestBuilder) {
    return iba313d84dd503b98b839bd7c2c10b20446a8a6080cdb37c526cba11b243f7465.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i7beca4b44ab53aa70feabd61a1f01b48ab66833b829c07de5894a814582516bc.FindRoomsRequestBuilder) {
    return i7beca4b44ab53aa70feabd61a1f01b48ab66833b829c07de5894a814582516bc.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i5e3457f7caf996c0212765e2462778b9f943d406bad548edf55068d3cb1a8b6a.FindRoomsWithRoomListRequestBuilder) {
    return i5e3457f7caf996c0212765e2462778b9f943d406bad548edf55068d3cb1a8b6a.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i0c3d4631adee6ee334df3c6e9f5a36de93d62f420e6eba8f7476c4f91bed511b.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i0c3d4631adee6ee334df3c6e9f5a36de93d62f420e6eba8f7476c4f91bed511b.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i2145c20e77a05a5a987c631004f6919cc200a7e13ad43d2d463c720b2e8ed1e5.GetLoggedOnManagedDevicesRequestBuilder) {
    return i2145c20e77a05a5a987c631004f6919cc200a7e13ad43d2d463c720b2e8ed1e5.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ib1fb854c6f957236ec71f510bab72e4ae2b47ad7f8210f5fd9f829a9ec5ae510.GetMailTipsRequestBuilder) {
    return ib1fb854c6f957236ec71f510bab72e4ae2b47ad7f8210f5fd9f829a9ec5ae510.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*if2dd7619e101a0dd17d05cdbd37246f44c7f3812d254664e7b7893c13655489c.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return if2dd7619e101a0dd17d05cdbd37246f44c7f3812d254664e7b7893c13655489c.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i5d699944287eba5958f97b66df6d01b55fad4e28d25f8e3a9895a6ebc1f482c8.GetManagedAppPoliciesRequestBuilder) {
    return i5d699944287eba5958f97b66df6d01b55fad4e28d25f8e3a9895a6ebc1f482c8.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ic69660e1409fcdd5572b77a61e5a0018c3b59a5b1a97637119bb9a9d012a3737.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ic69660e1409fcdd5572b77a61e5a0018c3b59a5b1a97637119bb9a9d012a3737.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i4f34b5b3b3f8fcf418f7bee5098aa39404110c83f3d6fe1000e68308a71dc984.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i4f34b5b3b3f8fcf418f7bee5098aa39404110c83f3d6fe1000e68308a71dc984.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i57629ef934145c095878a194c2f009f07ae9d69635c8eb33e975294c332e5fdc.GetMemberGroupsRequestBuilder) {
    return i57629ef934145c095878a194c2f009f07ae9d69635c8eb33e975294c332e5fdc.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i7e177c2b30d06eca5b8679486a190b47f61d08036e99cccd348987edfe106d19.GetMemberObjectsRequestBuilder) {
    return i7e177c2b30d06eca5b8679486a190b47f61d08036e99cccd348987edfe106d19.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i715b042262dc9ed949747092d1a6d4c67f4e6a17af31f9c7d5009c8180d3a283.InvalidateAllRefreshTokensRequestBuilder) {
    return i715b042262dc9ed949747092d1a6d4c67f4e6a17af31f9c7d5009c8180d3a283.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i64119174a682421d989bc15b022d48410683e668e982bebcfcfb570c2aa205d2.IsManagedAppUserBlockedRequestBuilder) {
    return i64119174a682421d989bc15b022d48410683e668e982bebcfcfb570c2aa205d2.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i86831a89b833bd7a67b83a445c0954fe2810bae67c93b2d3de51ef0d91e57ec3.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i86831a89b833bd7a67b83a445c0954fe2810bae67c93b2d3de51ef0d91e57ec3.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i875bacf36ec6afe008f45e95ce4ff1d674accafe1f7db527c9e4d059a43709d8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i875bacf36ec6afe008f45e95ce4ff1d674accafe1f7db527c9e4d059a43709d8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*icd42f5a1297266f49fbd72565398db0b3a3c7e1ed877cd9e88e57e7c3097a1a4.ReprocessLicenseAssignmentRequestBuilder) {
    return icd42f5a1297266f49fbd72565398db0b3a3c7e1ed877cd9e88e57e7c3097a1a4.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i89ebf0a17ea6af40cd498795effab2b40c1a728fbdc400a8dbb41122a24201bc.RestoreRequestBuilder) {
    return i89ebf0a17ea6af40cd498795effab2b40c1a728fbdc400a8dbb41122a24201bc.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i404437b92d38e2367e81306137b1805808cbdfd165d826dbc0d8d251fe2bf224.RevokeSignInSessionsRequestBuilder) {
    return i404437b92d38e2367e81306137b1805808cbdfd165d826dbc0d8d251fe2bf224.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i378ed81edc5f7659807ebcbbf6a15071b81eabea60b4bc9c3625c1cd920d25fc.SendMailRequestBuilder) {
    return i378ed81edc5f7659807ebcbbf6a15071b81eabea60b4bc9c3625c1cd920d25fc.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i7dc39389e361fe6c068589d6472982f56cd01554a9eacdccf5546d58b61b22ef.TranslateExchangeIdsRequestBuilder) {
    return i7dc39389e361fe6c068589d6472982f56cd01554a9eacdccf5546d58b61b22ef.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ia5e25e2dd47557a9550abfd9f8b951c523317ba4d610c80f4b74304eb964e8f3.UnblockManagedAppsRequestBuilder) {
    return ia5e25e2dd47557a9550abfd9f8b951c523317ba4d610c80f4b74304eb964e8f3.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i77d820470b2b321323755cf8a59aa8a83c554a16214617b7463a42a6c6382205.WipeAndBlockManagedAppsRequestBuilder) {
    return i77d820470b2b321323755cf8a59aa8a83c554a16214617b7463a42a6c6382205.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i3e682571c730b8fbcaea5127e4fc5098edbdce95dc229bdcc6bc6a839fd13067.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i3e682571c730b8fbcaea5127e4fc5098edbdce95dc229bdcc6bc6a839fd13067.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ia609a7fd58b7a403a148c3e7d4989441e7e2f4277f2900a7b02955a2cc5edb85.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ia609a7fd58b7a403a148c3e7d4989441e7e2f4277f2900a7b02955a2cc5edb85.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i8e39ed91fefafe58a2e1e52e3c17f716582519ef0010019ce14c17b2fa6be426.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i8e39ed91fefafe58a2e1e52e3c17f716582519ef0010019ce14c17b2fa6be426.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
