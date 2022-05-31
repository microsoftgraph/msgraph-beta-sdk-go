package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b8e9547c3a9abd04ce5bf63f8620b7a3c19ba37118cbd4b37ac0ccdf2cf29b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/unblockmanagedapps"
    i1a3d372a5026fa47c22d8eb7b57e9306b040b20e664af0ef9d8c51022c7d85e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    i28be9e41913397ad9991329ef36d5c5aa371f5803da05367a4f2187dfc630abb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    i29610833768d9ec568effed0e024ae750bfbc15eab4b8e491fe2ecf459f2d88b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    i34a51d034fc4720f9506762ecb188d179071c75e104d223da34e0aa380ac2910 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    i384ea393cbdceb0b105eb270aa7e83545e4cc6aa92448555d762674d6b367d24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    i41ff77fa44ee8e22c3a74e39f08e8a5a15dfeced4d86b99a18cdfa4d2bc8f1cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    i42cd03aff0c02fc30b78a0fe1cf69191a0258e82022015c5f2aaec22f32d8e22 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    i487e96c9f6d903fdebc8d9ed062dbfb9924d39741046e66d30ac3c4a0708aae9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
    i51d10a08fe4111a953b7d67bef3209d035cb823fb6afd498b0af81e156749e2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    i58ba9c86e1c7dbaf786b58f37c75866772df8d61ad8c01e6c8312acaea2615fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i6286f18ec8f74600778adbc329b08c00a35985aa6cab31ce484687d10c39e711 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    i63cd32a90a1ded0c8898db808a91aae242db536e1ba76049fb3829bbd0967dcf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    i68dd2630351d65bf5b26b223181ec24c81a39d4385aebbb8fab40255a5459aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    i742d2c1aa082fd9bd428f0c84870f33cc9e198b5e9b764356530ca0f9263ee60 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    i81c0ae86db55508a5f2af6b4aa921b369654265c0046612c848125f8ebc6c402 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    i823184b7508d0e68b4f0a3db436e799d35c9082363343e07784b0cc9e67cd418 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    i85840e1783377434e86b7bfd3c7ce56325b9b0ae5e12d5e831d9189c19c0a08d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    i85d4f9986d7ca92d18cae380f307ecd3309e682f9dd67e3b690a81713257ad5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
    i8a7d4aea20ed9f1345cec47dc0430f68bb577ce11168ddca73fba388651eb78a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
    i918bc44194b135bb47f87164463d90ae40183f3caa5eb7331ec34b87a3ff9f85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    iabf18eadedf55157bcec506e65df64f673595ae2e4f265ae4ae0763f1301f2b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/activateserviceplan"
    ib67b25f3dd4778be57fd410613c11834ec2d38d1bb857305eb795d2de3566733 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    ib684ed05028ec819c90c10ae5afcf9d96478fdf9d6021884807183145ce2a7de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    ibbad446badf5303db94b1f9b939b7e28bb32f5c90cdb267b96711d54edd8833a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    ic71cc09cf83121fc84de637c9fe67ba9a44a48804b41f438a86c043a9ecc42d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    ica1bc56dc698a0c09990adef0bab770f5af4c80115bf2b88e2dd2739c57a15e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    id2f9dd316770ce0d3f0d2c8bdb28565a5ac0f8503b01bd35bd03dcef56dad5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findroomlists"
    id4beba52f2d44fd0aefb182b08c610445d9a90c8c1a9c96dbaa03cf79d8f5e8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ie1e48a7c6ae40be15e755a48675f3603e2de33be10841d8c1ee9cc4cd9bbbe37 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    ieaec50c0020679875db6e7ca25f870c1bab5a48a621fa0d6579c478de48e1b43 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    iebb8638d23bbb53d94c4797ec630f22ccf2cdc118b1e753a2af3d39aaa71c8b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findrooms"
    if47241a52f5dc32cf0a53d287e7e8a7b32f439ecdfb20ba97728b915755bf76e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
    if625c4384287ecf186540ddc73b70d915782534ea9b70f6bd10b54450b2c98de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findroomswithroomlist"
    ifc4ff5a8e8827e95828ccefdb8ff9ca91e2988da15f18a9e2bab682cad44ac14 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*iabf18eadedf55157bcec506e65df64f673595ae2e4f265ae4ae0763f1301f2b8.ActivateServicePlanRequestBuilder) {
    return iabf18eadedf55157bcec506e65df64f673595ae2e4f265ae4ae0763f1301f2b8.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i85840e1783377434e86b7bfd3c7ce56325b9b0ae5e12d5e831d9189c19c0a08d.AssignLicenseRequestBuilder) {
    return i85840e1783377434e86b7bfd3c7ce56325b9b0ae5e12d5e831d9189c19c0a08d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i51d10a08fe4111a953b7d67bef3209d035cb823fb6afd498b0af81e156749e2d.ChangePasswordRequestBuilder) {
    return i51d10a08fe4111a953b7d67bef3209d035cb823fb6afd498b0af81e156749e2d.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i742d2c1aa082fd9bd428f0c84870f33cc9e198b5e9b764356530ca0f9263ee60.CheckMemberGroupsRequestBuilder) {
    return i742d2c1aa082fd9bd428f0c84870f33cc9e198b5e9b764356530ca0f9263ee60.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i487e96c9f6d903fdebc8d9ed062dbfb9924d39741046e66d30ac3c4a0708aae9.CheckMemberObjectsRequestBuilder) {
    return i487e96c9f6d903fdebc8d9ed062dbfb9924d39741046e66d30ac3c4a0708aae9.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ifc4ff5a8e8827e95828ccefdb8ff9ca91e2988da15f18a9e2bab682cad44ac14.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ifc4ff5a8e8827e95828ccefdb8ff9ca91e2988da15f18a9e2bab682cad44ac14.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*id4beba52f2d44fd0aefb182b08c610445d9a90c8c1a9c96dbaa03cf79d8f5e8c.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return id4beba52f2d44fd0aefb182b08c610445d9a90c8c1a9c96dbaa03cf79d8f5e8c.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i29610833768d9ec568effed0e024ae750bfbc15eab4b8e491fe2ecf459f2d88b.ExportPersonalDataRequestBuilder) {
    return i29610833768d9ec568effed0e024ae750bfbc15eab4b8e491fe2ecf459f2d88b.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i8a7d4aea20ed9f1345cec47dc0430f68bb577ce11168ddca73fba388651eb78a.FindMeetingTimesRequestBuilder) {
    return i8a7d4aea20ed9f1345cec47dc0430f68bb577ce11168ddca73fba388651eb78a.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*id2f9dd316770ce0d3f0d2c8bdb28565a5ac0f8503b01bd35bd03dcef56dad5d2.FindRoomListsRequestBuilder) {
    return id2f9dd316770ce0d3f0d2c8bdb28565a5ac0f8503b01bd35bd03dcef56dad5d2.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*iebb8638d23bbb53d94c4797ec630f22ccf2cdc118b1e753a2af3d39aaa71c8b9.FindRoomsRequestBuilder) {
    return iebb8638d23bbb53d94c4797ec630f22ccf2cdc118b1e753a2af3d39aaa71c8b9.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*if625c4384287ecf186540ddc73b70d915782534ea9b70f6bd10b54450b2c98de.FindRoomsWithRoomListRequestBuilder) {
    return if625c4384287ecf186540ddc73b70d915782534ea9b70f6bd10b54450b2c98de.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i68dd2630351d65bf5b26b223181ec24c81a39d4385aebbb8fab40255a5459aaf.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i68dd2630351d65bf5b26b223181ec24c81a39d4385aebbb8fab40255a5459aaf.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*if47241a52f5dc32cf0a53d287e7e8a7b32f439ecdfb20ba97728b915755bf76e.GetLoggedOnManagedDevicesRequestBuilder) {
    return if47241a52f5dc32cf0a53d287e7e8a7b32f439ecdfb20ba97728b915755bf76e.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i384ea393cbdceb0b105eb270aa7e83545e4cc6aa92448555d762674d6b367d24.GetMailTipsRequestBuilder) {
    return i384ea393cbdceb0b105eb270aa7e83545e4cc6aa92448555d762674d6b367d24.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ie1e48a7c6ae40be15e755a48675f3603e2de33be10841d8c1ee9cc4cd9bbbe37.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ie1e48a7c6ae40be15e755a48675f3603e2de33be10841d8c1ee9cc4cd9bbbe37.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i1a3d372a5026fa47c22d8eb7b57e9306b040b20e664af0ef9d8c51022c7d85e2.GetManagedAppPoliciesRequestBuilder) {
    return i1a3d372a5026fa47c22d8eb7b57e9306b040b20e664af0ef9d8c51022c7d85e2.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i41ff77fa44ee8e22c3a74e39f08e8a5a15dfeced4d86b99a18cdfa4d2bc8f1cb.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i41ff77fa44ee8e22c3a74e39f08e8a5a15dfeced4d86b99a18cdfa4d2bc8f1cb.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i63cd32a90a1ded0c8898db808a91aae242db536e1ba76049fb3829bbd0967dcf.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i63cd32a90a1ded0c8898db808a91aae242db536e1ba76049fb3829bbd0967dcf.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i85d4f9986d7ca92d18cae380f307ecd3309e682f9dd67e3b690a81713257ad5a.GetMemberGroupsRequestBuilder) {
    return i85d4f9986d7ca92d18cae380f307ecd3309e682f9dd67e3b690a81713257ad5a.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i81c0ae86db55508a5f2af6b4aa921b369654265c0046612c848125f8ebc6c402.GetMemberObjectsRequestBuilder) {
    return i81c0ae86db55508a5f2af6b4aa921b369654265c0046612c848125f8ebc6c402.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i34a51d034fc4720f9506762ecb188d179071c75e104d223da34e0aa380ac2910.InvalidateAllRefreshTokensRequestBuilder) {
    return i34a51d034fc4720f9506762ecb188d179071c75e104d223da34e0aa380ac2910.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ieaec50c0020679875db6e7ca25f870c1bab5a48a621fa0d6579c478de48e1b43.IsManagedAppUserBlockedRequestBuilder) {
    return ieaec50c0020679875db6e7ca25f870c1bab5a48a621fa0d6579c478de48e1b43.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i58ba9c86e1c7dbaf786b58f37c75866772df8d61ad8c01e6c8312acaea2615fc.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i58ba9c86e1c7dbaf786b58f37c75866772df8d61ad8c01e6c8312acaea2615fc.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ic71cc09cf83121fc84de637c9fe67ba9a44a48804b41f438a86c043a9ecc42d5.RemoveAllDevicesFromManagementRequestBuilder) {
    return ic71cc09cf83121fc84de637c9fe67ba9a44a48804b41f438a86c043a9ecc42d5.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ib67b25f3dd4778be57fd410613c11834ec2d38d1bb857305eb795d2de3566733.ReprocessLicenseAssignmentRequestBuilder) {
    return ib67b25f3dd4778be57fd410613c11834ec2d38d1bb857305eb795d2de3566733.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i28be9e41913397ad9991329ef36d5c5aa371f5803da05367a4f2187dfc630abb.RestoreRequestBuilder) {
    return i28be9e41913397ad9991329ef36d5c5aa371f5803da05367a4f2187dfc630abb.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ib684ed05028ec819c90c10ae5afcf9d96478fdf9d6021884807183145ce2a7de.RevokeSignInSessionsRequestBuilder) {
    return ib684ed05028ec819c90c10ae5afcf9d96478fdf9d6021884807183145ce2a7de.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i42cd03aff0c02fc30b78a0fe1cf69191a0258e82022015c5f2aaec22f32d8e22.SendMailRequestBuilder) {
    return i42cd03aff0c02fc30b78a0fe1cf69191a0258e82022015c5f2aaec22f32d8e22.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i823184b7508d0e68b4f0a3db436e799d35c9082363343e07784b0cc9e67cd418.TranslateExchangeIdsRequestBuilder) {
    return i823184b7508d0e68b4f0a3db436e799d35c9082363343e07784b0cc9e67cd418.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i0b8e9547c3a9abd04ce5bf63f8620b7a3c19ba37118cbd4b37ac0ccdf2cf29b2.UnblockManagedAppsRequestBuilder) {
    return i0b8e9547c3a9abd04ce5bf63f8620b7a3c19ba37118cbd4b37ac0ccdf2cf29b2.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i6286f18ec8f74600778adbc329b08c00a35985aa6cab31ce484687d10c39e711.WipeAndBlockManagedAppsRequestBuilder) {
    return i6286f18ec8f74600778adbc329b08c00a35985aa6cab31ce484687d10c39e711.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ibbad446badf5303db94b1f9b939b7e28bb32f5c90cdb267b96711d54edd8833a.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ibbad446badf5303db94b1f9b939b7e28bb32f5c90cdb267b96711d54edd8833a.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ica1bc56dc698a0c09990adef0bab770f5af4c80115bf2b88e2dd2739c57a15e8.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ica1bc56dc698a0c09990adef0bab770f5af4c80115bf2b88e2dd2739c57a15e8.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i918bc44194b135bb47f87164463d90ae40183f3caa5eb7331ec34b87a3ff9f85.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i918bc44194b135bb47f87164463d90ae40183f3caa5eb7331ec34b87a3ff9f85.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
