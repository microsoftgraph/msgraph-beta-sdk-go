package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02428501daa95dedb37d3fe8144daca4b8682edba223f537dfb08c6b9d158642 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/sendmail"
    i027bac9ded71997d8ad8f9a2bb7dff43b6711b5a2bb9cadd0eafa3d711d787af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    i123c68f8c1023db62c35179518867c7736df82e1fc80a35fa04f77045d8ec3e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/changepassword"
    i13f78948ada404a3d2f3756ec32cec8edac180a05b75867e70ad98cd150cbfe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmemberobjects"
    i149005adc46e5d74cc135713307f8f42a00c4f36497f2da300d507a06a94b683 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/unblockmanagedapps"
    i16d6f5fc88c6265ff809bb11da44adb6b82b088e599340b50d87b78295c2afe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findrooms"
    i1b260a9dd1729f144a83f0ca6380c24d67b652b13052f267709358d2a8aef123 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/ismanagedappuserblocked"
    i3240ea71be52c7bd5d599f55db2ad623f06e05efafb418c399fa6b3f73606113 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/activateserviceplan"
    i3ed0658b877020d2b12654434098d9b427b7f55a13e2bbce2028fd65ce85a1d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i406fa4dd8a87f75e7fca78a2d6bdcf7e41e2b594c3e43033ca5d16c8a37a5a95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    i50eacfa72ec8794c3ea03d1c7f07b14dc8147b49c0e54feb4a4d21fbbce94f68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findroomswithroomlist"
    i592ab87f19ee92aa4f82e5dbf43e73405c059b591099ed23d1e43f682836ead3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    i67d9bd984d61c18ca3e0f9f2d22dd3c972346f0a271b098b2dbd94818f0a285e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i69596e17b5b2e10916382948c027858cc27e74bf883649120d64bc0202d760b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/restore"
    i6ca1a8a01dcaeac152ce05703626a32e31f14e1a353abe57da60249bdbf3fd49 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    i78a9ace5c56cd2c57f12273ee70b7d2ec6a4bb4d7be7690931b285f569b180c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmembergroups"
    i7f2c99b4f5335bdf3db594f77396e822dc4c2036fc9a02211921a850714401b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i8353ab2089d534e90c18b17a0ed885384be5a5e6534e52d58296fb51a6112da7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findroomlists"
    i87e5a6313b2019ea57ecc213f67f548d0052fbfc98f501800b49d3a55ed49c4b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmemberobjects"
    i885247ab76f9e6c4de5b3c26a9168bbcb409019951f75eb63edca75b122f86e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportpersonaldata"
    i934d69c8f07367fcab77c4308da8f2b2d5c6266c861208dda42ec1cf39ffbc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmembergroups"
    i99cb2ed70f983b41f86271d1513678d02f85c05d4a58043821e5fad40c2aca92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i9be54ed4053c6724f6b60c0222f69e50558a61d970cb8f677c7e498005d4eca9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    i9d035cd5e719b80fe0f8e5a28ca7e15ce60c26ad6e50caa32517c3fa6bf8a495 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/assignlicense"
    ia78d3d1570b182cc2021ca5c3927a2e157cf544cb8a0c80cb3c9a820b0b0fa49 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getloggedonmanageddevices"
    ib7cfb8ecb8b3bb21a9fef51671d56e5222979acc49ab5715ce760c0954e306c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    ibd80c0dfdbc53207d7b5f86b6e05b6dc212915988a59a35d6d8d92893c47abb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ic729a9ce3ca20125aac3ee70a8c3d9723e0f1536179de8024208a2727102c052 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reprocesslicenseassignment"
    ic8c76c25e40f52f87a86b800dd872c9eae37ba24c490229e064dddf80d99dec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/revokesigninsessions"
    ica1e29097da62e8b4da230339dd2e3a1dd190904e0703c351ac1a2b934f6e03b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findmeetingtimes"
    icdcfdd51f9049dc49147aa391046e85cd1b1181e7e2d23523baf5caa579159c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    id41fd58ff2eb4fe7c91b42a5d47d051c5a15b4fd82cab6d9452dc871435e37bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedapppolicies"
    idf16d626c896914293ef10f6b61277264162c2ade02399a0b95d19c8a969dbc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ieb80d2fdbe2091d66242d047d0a6a7b777b16308e5f8d6bb09c5a0ad4232a478 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/translateexchangeids"
    if8217ba5b859c2e36c010c00cb7f6d28e8927cf43d2328a95995e22d93b6b273 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmailtips"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i3240ea71be52c7bd5d599f55db2ad623f06e05efafb418c399fa6b3f73606113.ActivateServicePlanRequestBuilder) {
    return i3240ea71be52c7bd5d599f55db2ad623f06e05efafb418c399fa6b3f73606113.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i9d035cd5e719b80fe0f8e5a28ca7e15ce60c26ad6e50caa32517c3fa6bf8a495.AssignLicenseRequestBuilder) {
    return i9d035cd5e719b80fe0f8e5a28ca7e15ce60c26ad6e50caa32517c3fa6bf8a495.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i123c68f8c1023db62c35179518867c7736df82e1fc80a35fa04f77045d8ec3e0.ChangePasswordRequestBuilder) {
    return i123c68f8c1023db62c35179518867c7736df82e1fc80a35fa04f77045d8ec3e0.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i78a9ace5c56cd2c57f12273ee70b7d2ec6a4bb4d7be7690931b285f569b180c8.CheckMemberGroupsRequestBuilder) {
    return i78a9ace5c56cd2c57f12273ee70b7d2ec6a4bb4d7be7690931b285f569b180c8.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i13f78948ada404a3d2f3756ec32cec8edac180a05b75867e70ad98cd150cbfe5.CheckMemberObjectsRequestBuilder) {
    return i13f78948ada404a3d2f3756ec32cec8edac180a05b75867e70ad98cd150cbfe5.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*icdcfdd51f9049dc49147aa391046e85cd1b1181e7e2d23523baf5caa579159c9.ExportDeviceAndAppManagementDataRequestBuilder) {
    return icdcfdd51f9049dc49147aa391046e85cd1b1181e7e2d23523baf5caa579159c9.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*idf16d626c896914293ef10f6b61277264162c2ade02399a0b95d19c8a969dbc3.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return idf16d626c896914293ef10f6b61277264162c2ade02399a0b95d19c8a969dbc3.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i885247ab76f9e6c4de5b3c26a9168bbcb409019951f75eb63edca75b122f86e5.ExportPersonalDataRequestBuilder) {
    return i885247ab76f9e6c4de5b3c26a9168bbcb409019951f75eb63edca75b122f86e5.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ica1e29097da62e8b4da230339dd2e3a1dd190904e0703c351ac1a2b934f6e03b.FindMeetingTimesRequestBuilder) {
    return ica1e29097da62e8b4da230339dd2e3a1dd190904e0703c351ac1a2b934f6e03b.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i8353ab2089d534e90c18b17a0ed885384be5a5e6534e52d58296fb51a6112da7.FindRoomListsRequestBuilder) {
    return i8353ab2089d534e90c18b17a0ed885384be5a5e6534e52d58296fb51a6112da7.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i16d6f5fc88c6265ff809bb11da44adb6b82b088e599340b50d87b78295c2afe4.FindRoomsRequestBuilder) {
    return i16d6f5fc88c6265ff809bb11da44adb6b82b088e599340b50d87b78295c2afe4.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i50eacfa72ec8794c3ea03d1c7f07b14dc8147b49c0e54feb4a4d21fbbce94f68.FindRoomsWithRoomListRequestBuilder) {
    return i50eacfa72ec8794c3ea03d1c7f07b14dc8147b49c0e54feb4a4d21fbbce94f68.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i99cb2ed70f983b41f86271d1513678d02f85c05d4a58043821e5fad40c2aca92.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i99cb2ed70f983b41f86271d1513678d02f85c05d4a58043821e5fad40c2aca92.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ia78d3d1570b182cc2021ca5c3927a2e157cf544cb8a0c80cb3c9a820b0b0fa49.GetLoggedOnManagedDevicesRequestBuilder) {
    return ia78d3d1570b182cc2021ca5c3927a2e157cf544cb8a0c80cb3c9a820b0b0fa49.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*if8217ba5b859c2e36c010c00cb7f6d28e8927cf43d2328a95995e22d93b6b273.GetMailTipsRequestBuilder) {
    return if8217ba5b859c2e36c010c00cb7f6d28e8927cf43d2328a95995e22d93b6b273.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ib7cfb8ecb8b3bb21a9fef51671d56e5222979acc49ab5715ce760c0954e306c3.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ib7cfb8ecb8b3bb21a9fef51671d56e5222979acc49ab5715ce760c0954e306c3.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*id41fd58ff2eb4fe7c91b42a5d47d051c5a15b4fd82cab6d9452dc871435e37bf.GetManagedAppPoliciesRequestBuilder) {
    return id41fd58ff2eb4fe7c91b42a5d47d051c5a15b4fd82cab6d9452dc871435e37bf.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i9be54ed4053c6724f6b60c0222f69e50558a61d970cb8f677c7e498005d4eca9.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i9be54ed4053c6724f6b60c0222f69e50558a61d970cb8f677c7e498005d4eca9.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ibd80c0dfdbc53207d7b5f86b6e05b6dc212915988a59a35d6d8d92893c47abb6.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ibd80c0dfdbc53207d7b5f86b6e05b6dc212915988a59a35d6d8d92893c47abb6.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i934d69c8f07367fcab77c4308da8f2b2d5c6266c861208dda42ec1cf39ffbc28.GetMemberGroupsRequestBuilder) {
    return i934d69c8f07367fcab77c4308da8f2b2d5c6266c861208dda42ec1cf39ffbc28.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i87e5a6313b2019ea57ecc213f67f548d0052fbfc98f501800b49d3a55ed49c4b.GetMemberObjectsRequestBuilder) {
    return i87e5a6313b2019ea57ecc213f67f548d0052fbfc98f501800b49d3a55ed49c4b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i6ca1a8a01dcaeac152ce05703626a32e31f14e1a353abe57da60249bdbf3fd49.InvalidateAllRefreshTokensRequestBuilder) {
    return i6ca1a8a01dcaeac152ce05703626a32e31f14e1a353abe57da60249bdbf3fd49.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1b260a9dd1729f144a83f0ca6380c24d67b652b13052f267709358d2a8aef123.IsManagedAppUserBlockedRequestBuilder) {
    return i1b260a9dd1729f144a83f0ca6380c24d67b652b13052f267709358d2a8aef123.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i7f2c99b4f5335bdf3db594f77396e822dc4c2036fc9a02211921a850714401b7.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i7f2c99b4f5335bdf3db594f77396e822dc4c2036fc9a02211921a850714401b7.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i027bac9ded71997d8ad8f9a2bb7dff43b6711b5a2bb9cadd0eafa3d711d787af.RemoveAllDevicesFromManagementRequestBuilder) {
    return i027bac9ded71997d8ad8f9a2bb7dff43b6711b5a2bb9cadd0eafa3d711d787af.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ic729a9ce3ca20125aac3ee70a8c3d9723e0f1536179de8024208a2727102c052.ReprocessLicenseAssignmentRequestBuilder) {
    return ic729a9ce3ca20125aac3ee70a8c3d9723e0f1536179de8024208a2727102c052.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i69596e17b5b2e10916382948c027858cc27e74bf883649120d64bc0202d760b2.RestoreRequestBuilder) {
    return i69596e17b5b2e10916382948c027858cc27e74bf883649120d64bc0202d760b2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ic8c76c25e40f52f87a86b800dd872c9eae37ba24c490229e064dddf80d99dec7.RevokeSignInSessionsRequestBuilder) {
    return ic8c76c25e40f52f87a86b800dd872c9eae37ba24c490229e064dddf80d99dec7.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i02428501daa95dedb37d3fe8144daca4b8682edba223f537dfb08c6b9d158642.SendMailRequestBuilder) {
    return i02428501daa95dedb37d3fe8144daca4b8682edba223f537dfb08c6b9d158642.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ieb80d2fdbe2091d66242d047d0a6a7b777b16308e5f8d6bb09c5a0ad4232a478.TranslateExchangeIdsRequestBuilder) {
    return ieb80d2fdbe2091d66242d047d0a6a7b777b16308e5f8d6bb09c5a0ad4232a478.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i149005adc46e5d74cc135713307f8f42a00c4f36497f2da300d507a06a94b683.UnblockManagedAppsRequestBuilder) {
    return i149005adc46e5d74cc135713307f8f42a00c4f36497f2da300d507a06a94b683.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i406fa4dd8a87f75e7fca78a2d6bdcf7e41e2b594c3e43033ca5d16c8a37a5a95.WipeAndBlockManagedAppsRequestBuilder) {
    return i406fa4dd8a87f75e7fca78a2d6bdcf7e41e2b594c3e43033ca5d16c8a37a5a95.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i592ab87f19ee92aa4f82e5dbf43e73405c059b591099ed23d1e43f682836ead3.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i592ab87f19ee92aa4f82e5dbf43e73405c059b591099ed23d1e43f682836ead3.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i3ed0658b877020d2b12654434098d9b427b7f55a13e2bbce2028fd65ce85a1d7.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i3ed0658b877020d2b12654434098d9b427b7f55a13e2bbce2028fd65ce85a1d7.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i67d9bd984d61c18ca3e0f9f2d22dd3c972346f0a271b098b2dbd94818f0a285e.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i67d9bd984d61c18ca3e0f9f2d22dd3c972346f0a271b098b2dbd94818f0a285e.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
