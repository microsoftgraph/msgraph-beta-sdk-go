package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i066ca29975ef83000139d2166d4431efaff5d98bfc5ddbe86a58050408287fd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i09fdf9aeda4afadbaa27da5103be07127cd476efbc65033c2111a9b8d2fa18ce "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    i0a10437ec02487576aa6801abd7c976c078c8aa7066b739786bd4060f73d6f62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i0db6cec6d1f2c47a59a79705ef3cbf77b408b53c1febe793c35e63072e01d099 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    i1401c7603b31ae4133bc5969162f07c5117b2c26abc15af2f949cdd62d867c84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    i2020d58f955a64e3979c644ad39fbf8e9956a2ece55c2bc000a2b1c2a353dcbc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    i21a34652b6cfc766b2936946b5df4b703b33f8f5106e0524d15013e7763b440d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findrooms"
    i41377d1c3c4d03256b77c5aaca4eac0d6a8165147b8c38894e3f07baa5862483 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    i42ad0bd2bb82983052742600e9f559b97884018ed1f7ff65d7909e99622e77df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i4542081960f9c83e1cac51b8fa50fe7e84874566790684a9b0f0aca93aa3185b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i4c79fa957c34571e21c424f321ca73dd3dd276861d28af8630a5e6d0ca34e161 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findroomswithroomlist"
    i4f53af309d0d0f2a5bdfd7f1abf4674ae9b654813a06bb147fd76b1665186908 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i61c27b0793b9096b18aa844b43c4f5bf8420b7948a5c51219b68722a0ed9eb39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    i69d55a26041fab396ac17ec208b28b89b85944d5761366d7a0e2de9afb6c24ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    i754abaeef1dadb749c25b35eaf925b2ed2b3f08d5386b1ca54390e0c9e34fd8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findroomlists"
    i79d976b08b5f77c12beb0c974feaab69edd2cf105b50839c6beca4ce3cbb475b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    i7cdc98ec95872c0d6d9f18f1a35abce3f87acbabd2cebc9a86f5075c2e9c0de0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    i7e05382911dac6a3fe51ae0dfa8e4cd56c783e4f81e09e8e32fae18829ffe82e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    i8a2f85cfbdace84c9ed4a4b0e20b9ab8ee289d0d3cf6a186bd433d90aa907bed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    i8e50d49e99f464dcb76d7ca865c3d1199092f6935d93b1367f585dedc56ceb52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i933b12def25e0f4afca1ec858f6a3ea54714f2f4ea7ea57d1d71781784136f03 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    i943c56b9b4c338f2851d1720c065591aa7cbdcf9487e1fc49f3e66ae3eea7a9b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    ia9d59c599f171a94712371bd5ad0706e71847efc7e96e656412d967e1bd2a028 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    ib688deb34eef383286760fdda55142a4afb4e808a524108068614a134fea3524 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/unblockmanagedapps"
    ib77943d2c69ac31a25138fa1105171faab172ca7a0b5262aeac34cb1548a04e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ibdc205ef776ead4ef3670b5343aae4be09624147601939b08002f78f060debd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    ibdf647dcc57058bd6cf7a73e584ace56cc8be799c96c0bc13f45f9824ff7b2f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
    ic3f5480f2227e3094883b5ec91e015c2a21a43d6b27452dad00dd7e6696991e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    ic7689f29a02efe2902ad9cd1345f7e0cb46c515bd91e092ae6d5ed9462b83a77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    icceeae258b782d2a3c233fe6e190e50e99756cb9397cab2c3be0b0fc43312c1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    ie9c8a6ed09d3b6bf0aa09c99e29a8b978e0ccfc5a73cc970aff7834bffa868dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    iea37e9dade5909eaa383e87caa09b996a5c9ab844188a9319d42ea2ddbae7a22 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    ifa3fb5f36cdf83414573bc7106c18afc13ff24ec968f5f4438989a56045d9222 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    ifaac4fa51a3368590f4ba7e69751398994c47d58b6e4513a0568cba8b3f536a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    iff954656db20ca6fd13dec1799ded5d32bea6d7a52c0df1732188ee9ef5d8402 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/activateserviceplan"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*iff954656db20ca6fd13dec1799ded5d32bea6d7a52c0df1732188ee9ef5d8402.ActivateServicePlanRequestBuilder) {
    return iff954656db20ca6fd13dec1799ded5d32bea6d7a52c0df1732188ee9ef5d8402.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*icceeae258b782d2a3c233fe6e190e50e99756cb9397cab2c3be0b0fc43312c1c.AssignLicenseRequestBuilder) {
    return icceeae258b782d2a3c233fe6e190e50e99756cb9397cab2c3be0b0fc43312c1c.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i2020d58f955a64e3979c644ad39fbf8e9956a2ece55c2bc000a2b1c2a353dcbc.ChangePasswordRequestBuilder) {
    return i2020d58f955a64e3979c644ad39fbf8e9956a2ece55c2bc000a2b1c2a353dcbc.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ibdc205ef776ead4ef3670b5343aae4be09624147601939b08002f78f060debd2.CheckMemberGroupsRequestBuilder) {
    return ibdc205ef776ead4ef3670b5343aae4be09624147601939b08002f78f060debd2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i69d55a26041fab396ac17ec208b28b89b85944d5761366d7a0e2de9afb6c24ab.CheckMemberObjectsRequestBuilder) {
    return i69d55a26041fab396ac17ec208b28b89b85944d5761366d7a0e2de9afb6c24ab.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i7cdc98ec95872c0d6d9f18f1a35abce3f87acbabd2cebc9a86f5075c2e9c0de0.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i7cdc98ec95872c0d6d9f18f1a35abce3f87acbabd2cebc9a86f5075c2e9c0de0.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ib77943d2c69ac31a25138fa1105171faab172ca7a0b5262aeac34cb1548a04e9.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ib77943d2c69ac31a25138fa1105171faab172ca7a0b5262aeac34cb1548a04e9.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i4f53af309d0d0f2a5bdfd7f1abf4674ae9b654813a06bb147fd76b1665186908.ExportPersonalDataRequestBuilder) {
    return i4f53af309d0d0f2a5bdfd7f1abf4674ae9b654813a06bb147fd76b1665186908.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i933b12def25e0f4afca1ec858f6a3ea54714f2f4ea7ea57d1d71781784136f03.FindMeetingTimesRequestBuilder) {
    return i933b12def25e0f4afca1ec858f6a3ea54714f2f4ea7ea57d1d71781784136f03.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i754abaeef1dadb749c25b35eaf925b2ed2b3f08d5386b1ca54390e0c9e34fd8c.FindRoomListsRequestBuilder) {
    return i754abaeef1dadb749c25b35eaf925b2ed2b3f08d5386b1ca54390e0c9e34fd8c.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i21a34652b6cfc766b2936946b5df4b703b33f8f5106e0524d15013e7763b440d.FindRoomsRequestBuilder) {
    return i21a34652b6cfc766b2936946b5df4b703b33f8f5106e0524d15013e7763b440d.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i4c79fa957c34571e21c424f321ca73dd3dd276861d28af8630a5e6d0ca34e161.FindRoomsWithRoomListRequestBuilder) {
    return i4c79fa957c34571e21c424f321ca73dd3dd276861d28af8630a5e6d0ca34e161.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i943c56b9b4c338f2851d1720c065591aa7cbdcf9487e1fc49f3e66ae3eea7a9b.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i943c56b9b4c338f2851d1720c065591aa7cbdcf9487e1fc49f3e66ae3eea7a9b.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*iea37e9dade5909eaa383e87caa09b996a5c9ab844188a9319d42ea2ddbae7a22.GetLoggedOnManagedDevicesRequestBuilder) {
    return iea37e9dade5909eaa383e87caa09b996a5c9ab844188a9319d42ea2ddbae7a22.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i7e05382911dac6a3fe51ae0dfa8e4cd56c783e4f81e09e8e32fae18829ffe82e.GetMailTipsRequestBuilder) {
    return i7e05382911dac6a3fe51ae0dfa8e4cd56c783e4f81e09e8e32fae18829ffe82e.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i0a10437ec02487576aa6801abd7c976c078c8aa7066b739786bd4060f73d6f62.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i0a10437ec02487576aa6801abd7c976c078c8aa7066b739786bd4060f73d6f62.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic3f5480f2227e3094883b5ec91e015c2a21a43d6b27452dad00dd7e6696991e1.GetManagedAppPoliciesRequestBuilder) {
    return ic3f5480f2227e3094883b5ec91e015c2a21a43d6b27452dad00dd7e6696991e1.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i1401c7603b31ae4133bc5969162f07c5117b2c26abc15af2f949cdd62d867c84.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i1401c7603b31ae4133bc5969162f07c5117b2c26abc15af2f949cdd62d867c84.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i41377d1c3c4d03256b77c5aaca4eac0d6a8165147b8c38894e3f07baa5862483.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i41377d1c3c4d03256b77c5aaca4eac0d6a8165147b8c38894e3f07baa5862483.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i4542081960f9c83e1cac51b8fa50fe7e84874566790684a9b0f0aca93aa3185b.GetMemberGroupsRequestBuilder) {
    return i4542081960f9c83e1cac51b8fa50fe7e84874566790684a9b0f0aca93aa3185b.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i8a2f85cfbdace84c9ed4a4b0e20b9ab8ee289d0d3cf6a186bd433d90aa907bed.GetMemberObjectsRequestBuilder) {
    return i8a2f85cfbdace84c9ed4a4b0e20b9ab8ee289d0d3cf6a186bd433d90aa907bed.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i79d976b08b5f77c12beb0c974feaab69edd2cf105b50839c6beca4ce3cbb475b.InvalidateAllRefreshTokensRequestBuilder) {
    return i79d976b08b5f77c12beb0c974feaab69edd2cf105b50839c6beca4ce3cbb475b.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i61c27b0793b9096b18aa844b43c4f5bf8420b7948a5c51219b68722a0ed9eb39.IsManagedAppUserBlockedRequestBuilder) {
    return i61c27b0793b9096b18aa844b43c4f5bf8420b7948a5c51219b68722a0ed9eb39.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ie9c8a6ed09d3b6bf0aa09c99e29a8b978e0ccfc5a73cc970aff7834bffa868dd.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ie9c8a6ed09d3b6bf0aa09c99e29a8b978e0ccfc5a73cc970aff7834bffa868dd.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i066ca29975ef83000139d2166d4431efaff5d98bfc5ddbe86a58050408287fd8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i066ca29975ef83000139d2166d4431efaff5d98bfc5ddbe86a58050408287fd8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i8e50d49e99f464dcb76d7ca865c3d1199092f6935d93b1367f585dedc56ceb52.ReprocessLicenseAssignmentRequestBuilder) {
    return i8e50d49e99f464dcb76d7ca865c3d1199092f6935d93b1367f585dedc56ceb52.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ic7689f29a02efe2902ad9cd1345f7e0cb46c515bd91e092ae6d5ed9462b83a77.RestoreRequestBuilder) {
    return ic7689f29a02efe2902ad9cd1345f7e0cb46c515bd91e092ae6d5ed9462b83a77.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i09fdf9aeda4afadbaa27da5103be07127cd476efbc65033c2111a9b8d2fa18ce.RevokeSignInSessionsRequestBuilder) {
    return i09fdf9aeda4afadbaa27da5103be07127cd476efbc65033c2111a9b8d2fa18ce.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ibdf647dcc57058bd6cf7a73e584ace56cc8be799c96c0bc13f45f9824ff7b2f7.SendMailRequestBuilder) {
    return ibdf647dcc57058bd6cf7a73e584ace56cc8be799c96c0bc13f45f9824ff7b2f7.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ia9d59c599f171a94712371bd5ad0706e71847efc7e96e656412d967e1bd2a028.TranslateExchangeIdsRequestBuilder) {
    return ia9d59c599f171a94712371bd5ad0706e71847efc7e96e656412d967e1bd2a028.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ib688deb34eef383286760fdda55142a4afb4e808a524108068614a134fea3524.UnblockManagedAppsRequestBuilder) {
    return ib688deb34eef383286760fdda55142a4afb4e808a524108068614a134fea3524.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i42ad0bd2bb82983052742600e9f559b97884018ed1f7ff65d7909e99622e77df.WipeAndBlockManagedAppsRequestBuilder) {
    return i42ad0bd2bb82983052742600e9f559b97884018ed1f7ff65d7909e99622e77df.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ifa3fb5f36cdf83414573bc7106c18afc13ff24ec968f5f4438989a56045d9222.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ifa3fb5f36cdf83414573bc7106c18afc13ff24ec968f5f4438989a56045d9222.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ifaac4fa51a3368590f4ba7e69751398994c47d58b6e4513a0568cba8b3f536a2.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ifaac4fa51a3368590f4ba7e69751398994c47d58b6e4513a0568cba8b3f536a2.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i0db6cec6d1f2c47a59a79705ef3cbf77b408b53c1febe793c35e63072e01d099.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i0db6cec6d1f2c47a59a79705ef3cbf77b408b53c1febe793c35e63072e01d099.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
