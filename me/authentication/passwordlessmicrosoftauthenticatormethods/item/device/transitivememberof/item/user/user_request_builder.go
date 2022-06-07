package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08706318cbe92e4334cf0a6d66175bd38c662f04a4a89399c38c48a277649e18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i0e75cbdf86077341bd485b785e3e8c0daa96413d0a5a69293a37e9031ff3027c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    i0fa792ab04d10cd768a34290880ff558c35d3b53720d98fb59c1ae7254e419e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i13f5ce949764f2327ba49cd034309e24175111b910e02e8a1499bfb648050386 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findrooms"
    i17ae3e891810c844ae3a1b01269846e63a8a0e0238e30ce68bf83cef59288581 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    i1b12ffe4c4bd6804cc1ec913663d894cbd7bb2c66ca70acd5ba95c3f7e8f3bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    i298e240dfce59463c67a6bf17de2f523b92c5e0abe786abd8fb3e2cb92d7b02c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findroomlists"
    i2be79ca1aa289dd5574ee2ba410a335a4f9e0d549bd69720bb5074b13597a53a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    i416b7853287cc360727d98accf94faedbb4a7de4f42f9804f4c2436c78f82e85 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i4d8266db43204c9eb0b842f1f808b8ea1a60a6f6238bd2ee381c660d6f75d22c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i51b2d54ea66eaa678c2c4523c7b6f77788c0921a48247ce07076eab8f83bc81b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    i574323293ead2f488c642d3f3cc4f1a832a5d0c5cdd70c84a5213ddfc14e0294 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i5b0d1868dda54419a50a2761c94971066a891c50343b3533fb9e3b402e04553f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i600d0689fc6eabe548eaba8132a64f2d018f05c2c30a624bc2db68f1d9ba59bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
    i607271fc3ac63047fa22a3e2382d5538e84abc1ead5e4d03267cd9684a1c6120 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    i6b42d18a8e1a345362e66c775dcbde5fd526a5e13ab9dc93a773bf286ae4aa6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i71357281e0d74d7dfc81eeef5b5dea1a5e014a97287019633f2d722d9d9fb0cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    i7dbd02fb66d2b43ba456123c175596bf87ca35ff001905f1e4fd682406b46e7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    i7e87afdfed535ae82394039a834f8a622175b962f3fd9f6c2a93e94145394e21 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    i84d4071abb115ff5801af1ba0642e103154fb98974d900bc71da1ce143d0c5b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/activateserviceplan"
    i90dad8efc2298e5b2874042602501569dfbdd228ccf2fbdb709736f8cf6d888a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i94755bc6fc1a3a7a3f3879f0b0f92f289b20eafb1ef7284f420f1b0ff33f9a63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    i99623f71561d2565b33366208de6227a25f51f9165626bc1a5271a184725cc6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    ia414fce22bb697794ad2468c2f01c129f34120cdf2f94fb1aacf685508605509 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/findroomswithroomlist"
    ia75fc0893c2819e9438d95ed7fe8bdee5677cf12fccd887c87fba3de2ee5a127 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    iaeeece6069a4d3739f534f1dbad71dfc34077e8cae93152b380791f059923168 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    ib0d1407470e13f989f9f2802d5b234ae96c3c4faa5d162b6cf310c4781e3d87a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    id78f887370da8f2d48536d4191d0a55a813d9533e3cd49bcde2146434543c9ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    idd9528c47ae9c8e2c13e8af6f9af4d275611e3e219756ed429d3911ee65adf58 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    iddabcc98aa0cb65ec8e2c121d3860319affc0761cbea3fad96ddfeb0a705636e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    ie0a07dfec2fb62d207c5934691bcc77376dc5d98a0c70c10644581f3c27d06e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    ie45e9e28847648ab17f3b828ebc97d14681190751b936f2a9815749e4ace2da2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    ie65476f3e1fbdb84f883a544e5c11d55104a9473f5558b520d1a75a0b95cf30e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    if2c82f8c95eee66de71108569af42541fb615d0622f90d94492eeefb7449dee1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
    ifd32c6828fd7861ee01b28df3d87155ed8eb05c30b0f96cf03d5cb43f74490d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i84d4071abb115ff5801af1ba0642e103154fb98974d900bc71da1ce143d0c5b8.ActivateServicePlanRequestBuilder) {
    return i84d4071abb115ff5801af1ba0642e103154fb98974d900bc71da1ce143d0c5b8.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*id78f887370da8f2d48536d4191d0a55a813d9533e3cd49bcde2146434543c9ec.AssignLicenseRequestBuilder) {
    return id78f887370da8f2d48536d4191d0a55a813d9533e3cd49bcde2146434543c9ec.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i607271fc3ac63047fa22a3e2382d5538e84abc1ead5e4d03267cd9684a1c6120.ChangePasswordRequestBuilder) {
    return i607271fc3ac63047fa22a3e2382d5538e84abc1ead5e4d03267cd9684a1c6120.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i71357281e0d74d7dfc81eeef5b5dea1a5e014a97287019633f2d722d9d9fb0cf.CheckMemberGroupsRequestBuilder) {
    return i71357281e0d74d7dfc81eeef5b5dea1a5e014a97287019633f2d722d9d9fb0cf.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ie65476f3e1fbdb84f883a544e5c11d55104a9473f5558b520d1a75a0b95cf30e.CheckMemberObjectsRequestBuilder) {
    return ie65476f3e1fbdb84f883a544e5c11d55104a9473f5558b520d1a75a0b95cf30e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*iaeeece6069a4d3739f534f1dbad71dfc34077e8cae93152b380791f059923168.ExportDeviceAndAppManagementDataRequestBuilder) {
    return iaeeece6069a4d3739f534f1dbad71dfc34077e8cae93152b380791f059923168.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ifd32c6828fd7861ee01b28df3d87155ed8eb05c30b0f96cf03d5cb43f74490d8.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ifd32c6828fd7861ee01b28df3d87155ed8eb05c30b0f96cf03d5cb43f74490d8.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*idd9528c47ae9c8e2c13e8af6f9af4d275611e3e219756ed429d3911ee65adf58.ExportPersonalDataRequestBuilder) {
    return idd9528c47ae9c8e2c13e8af6f9af4d275611e3e219756ed429d3911ee65adf58.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i99623f71561d2565b33366208de6227a25f51f9165626bc1a5271a184725cc6d.FindMeetingTimesRequestBuilder) {
    return i99623f71561d2565b33366208de6227a25f51f9165626bc1a5271a184725cc6d.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i298e240dfce59463c67a6bf17de2f523b92c5e0abe786abd8fb3e2cb92d7b02c.FindRoomListsRequestBuilder) {
    return i298e240dfce59463c67a6bf17de2f523b92c5e0abe786abd8fb3e2cb92d7b02c.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i13f5ce949764f2327ba49cd034309e24175111b910e02e8a1499bfb648050386.FindRoomsRequestBuilder) {
    return i13f5ce949764f2327ba49cd034309e24175111b910e02e8a1499bfb648050386.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ia414fce22bb697794ad2468c2f01c129f34120cdf2f94fb1aacf685508605509.FindRoomsWithRoomListRequestBuilder) {
    return ia414fce22bb697794ad2468c2f01c129f34120cdf2f94fb1aacf685508605509.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ib0d1407470e13f989f9f2802d5b234ae96c3c4faa5d162b6cf310c4781e3d87a.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ib0d1407470e13f989f9f2802d5b234ae96c3c4faa5d162b6cf310c4781e3d87a.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i574323293ead2f488c642d3f3cc4f1a832a5d0c5cdd70c84a5213ddfc14e0294.GetLoggedOnManagedDevicesRequestBuilder) {
    return i574323293ead2f488c642d3f3cc4f1a832a5d0c5cdd70c84a5213ddfc14e0294.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i1b12ffe4c4bd6804cc1ec913663d894cbd7bb2c66ca70acd5ba95c3f7e8f3bd6.GetMailTipsRequestBuilder) {
    return i1b12ffe4c4bd6804cc1ec913663d894cbd7bb2c66ca70acd5ba95c3f7e8f3bd6.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i416b7853287cc360727d98accf94faedbb4a7de4f42f9804f4c2436c78f82e85.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i416b7853287cc360727d98accf94faedbb4a7de4f42f9804f4c2436c78f82e85.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i5b0d1868dda54419a50a2761c94971066a891c50343b3533fb9e3b402e04553f.GetManagedAppPoliciesRequestBuilder) {
    return i5b0d1868dda54419a50a2761c94971066a891c50343b3533fb9e3b402e04553f.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i2be79ca1aa289dd5574ee2ba410a335a4f9e0d549bd69720bb5074b13597a53a.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i2be79ca1aa289dd5574ee2ba410a335a4f9e0d549bd69720bb5074b13597a53a.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i90dad8efc2298e5b2874042602501569dfbdd228ccf2fbdb709736f8cf6d888a.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i90dad8efc2298e5b2874042602501569dfbdd228ccf2fbdb709736f8cf6d888a.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i17ae3e891810c844ae3a1b01269846e63a8a0e0238e30ce68bf83cef59288581.GetMemberGroupsRequestBuilder) {
    return i17ae3e891810c844ae3a1b01269846e63a8a0e0238e30ce68bf83cef59288581.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i0e75cbdf86077341bd485b785e3e8c0daa96413d0a5a69293a37e9031ff3027c.GetMemberObjectsRequestBuilder) {
    return i0e75cbdf86077341bd485b785e3e8c0daa96413d0a5a69293a37e9031ff3027c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i7dbd02fb66d2b43ba456123c175596bf87ca35ff001905f1e4fd682406b46e7a.InvalidateAllRefreshTokensRequestBuilder) {
    return i7dbd02fb66d2b43ba456123c175596bf87ca35ff001905f1e4fd682406b46e7a.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i7e87afdfed535ae82394039a834f8a622175b962f3fd9f6c2a93e94145394e21.IsManagedAppUserBlockedRequestBuilder) {
    return i7e87afdfed535ae82394039a834f8a622175b962f3fd9f6c2a93e94145394e21.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i4d8266db43204c9eb0b842f1f808b8ea1a60a6f6238bd2ee381c660d6f75d22c.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i4d8266db43204c9eb0b842f1f808b8ea1a60a6f6238bd2ee381c660d6f75d22c.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ia75fc0893c2819e9438d95ed7fe8bdee5677cf12fccd887c87fba3de2ee5a127.RemoveAllDevicesFromManagementRequestBuilder) {
    return ia75fc0893c2819e9438d95ed7fe8bdee5677cf12fccd887c87fba3de2ee5a127.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i94755bc6fc1a3a7a3f3879f0b0f92f289b20eafb1ef7284f420f1b0ff33f9a63.ReprocessLicenseAssignmentRequestBuilder) {
    return i94755bc6fc1a3a7a3f3879f0b0f92f289b20eafb1ef7284f420f1b0ff33f9a63.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i600d0689fc6eabe548eaba8132a64f2d018f05c2c30a624bc2db68f1d9ba59bb.RestoreRequestBuilder) {
    return i600d0689fc6eabe548eaba8132a64f2d018f05c2c30a624bc2db68f1d9ba59bb.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ie0a07dfec2fb62d207c5934691bcc77376dc5d98a0c70c10644581f3c27d06e8.RevokeSignInSessionsRequestBuilder) {
    return ie0a07dfec2fb62d207c5934691bcc77376dc5d98a0c70c10644581f3c27d06e8.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*iddabcc98aa0cb65ec8e2c121d3860319affc0761cbea3fad96ddfeb0a705636e.SendMailRequestBuilder) {
    return iddabcc98aa0cb65ec8e2c121d3860319affc0761cbea3fad96ddfeb0a705636e.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i51b2d54ea66eaa678c2c4523c7b6f77788c0921a48247ce07076eab8f83bc81b.TranslateExchangeIdsRequestBuilder) {
    return i51b2d54ea66eaa678c2c4523c7b6f77788c0921a48247ce07076eab8f83bc81b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ie45e9e28847648ab17f3b828ebc97d14681190751b936f2a9815749e4ace2da2.UnblockManagedAppsRequestBuilder) {
    return ie45e9e28847648ab17f3b828ebc97d14681190751b936f2a9815749e4ace2da2.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*if2c82f8c95eee66de71108569af42541fb615d0622f90d94492eeefb7449dee1.WipeAndBlockManagedAppsRequestBuilder) {
    return if2c82f8c95eee66de71108569af42541fb615d0622f90d94492eeefb7449dee1.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i08706318cbe92e4334cf0a6d66175bd38c662f04a4a89399c38c48a277649e18.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i08706318cbe92e4334cf0a6d66175bd38c662f04a4a89399c38c48a277649e18.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i6b42d18a8e1a345362e66c775dcbde5fd526a5e13ab9dc93a773bf286ae4aa6d.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i6b42d18a8e1a345362e66c775dcbde5fd526a5e13ab9dc93a773bf286ae4aa6d.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i0fa792ab04d10cd768a34290880ff558c35d3b53720d98fb59c1ae7254e419e4.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i0fa792ab04d10cd768a34290880ff558c35d3b53720d98fb59c1ae7254e419e4.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
