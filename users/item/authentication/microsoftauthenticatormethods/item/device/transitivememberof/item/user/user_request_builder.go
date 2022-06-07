package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i19ed68126d709a82540812bba60d7d51d19aa6e4f11317d33cae1f051fa17976 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i1b3a6a82269384d6020b7c5a6325cde206132c9dee3b41d39f3ccabf4543d78e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    i1cb8dd35f65e9ea6d86f2bcc4bdde4a4ead0d78e24bb90152c10f4ed96856392 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    i204e5c913a92bb6ea8e49f417d63cfac139988b7c827d54ee4894227a198d548 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i24c5b9814f2073b9e7f22b3cf36f7daf51fefafb4eab1dcb3ef9cb884dc7c05f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    i295376a9e87d529493ee50dc6d93198a6b847124ae1c563a6267e2af07295b56 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
    i2d5cc840b7c0d226ff89b3d01a70db1b6bf8db09ae28f59e402a739b5335e084 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
    i35e13707db8c4d3199ea15f14015264f6c3062a1ed76c2371213ae3ad7bbe1c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i4171b656c38fef1c57019bffe1cd480bd65d61be89eeea98dca8a54794e50916 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    i485cad271b75388bf55bfca46b2e131f9db586cdef14cc9fbccd4c3a3b9269d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    i486324f589cf06d5592183853cedfb0cee61d7dda900d85ace70b3e3224e7d0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findroomlists"
    i56429eaf96d76c244fb3df17b1e4aee9764c53e2a2925774530a0411dddb9c3d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    i576a0a21c18956a80aefd2114d581672fb443866f81a26a6026f233767daddf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i5f1036ac0480dcf25fe3d7f1aaeef36ffeb055c64c1ccd3064f8d110e011da14 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findroomswithroomlist"
    i67c75a7976f24b55607f504127e865678bd4b5f9871cbd9be36d78ed5dffd2bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i696e0ee47219eb2ca67a75d716f4392392a61dcc9bfad41d9bca6246e382f0ca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    i69ff141cf7e8600fbbab56da2e82fee43766316164b8cfc8cf20b4d503edfaf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i6a7dae5caea32e17a8029984ac4c2f4ddacb4146a8ab183bab087a5da9be12b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    i6c4b761199124e4038eab42acaa34ea51786277b066ac95de0dda809d0b04670 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    i799512ab2e2f840da65bec587cca61c8b31af3ac501b0d89f89669c6be7a65a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i8306a5528aba6dae56eba498ac4b7b6c079c2e6509ee97e62074cb84ebc7f90c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i8e4578a11153aaee68d9eca6b1d93eb4a3fb1f1d57548097c3f4166543c12eb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    i99cf4d22d13544c6fc31bd9ac0864216f9396a9b1f05c3e346c528b4ebb56dff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    i9be732ff513a3fd2786b2be69abedf24b564f0fc38631c7cae0e3878ccbd21e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    ib0187e5f182839837ce4afdaef16eb467cf5ac2c3464cf94936e5e112b2a7e4f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findrooms"
    ib31d5ea274f4b4ea0c1ed4271cc3c5ac5df472bf9184bb756cd9703c2a795219 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ibbbd3bb572dc63bbf390e13b9b30364de3115be110a603fc4e3e57165aa3f874 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ic6f503187de5eefff58d23d03d0039e87106c7cc896b2809237f78fb9680321e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    idf6baf10effbe604b3095bf580189ccf757a31545da79646983eef2e22f9996c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    ie72d542a9047dbc1d7ffb48ce4a67aa918cad8042da2e57fdc297f43bc3769c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    ie982b90d8af737d7d1f146569e62141c8cc9f946e39b4d08b956a605fc5512ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/activateserviceplan"
    ie9caa3c501e0e31b83a05c63ec77635292c0ba486cd3777a24ddb0f6509211f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    ie9ebae9ee745ee1eb3f0d95c0d607beb875ef04ca992fb6f600de28a49e6efd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    if4618f906f6cbc777a99c3d5fb1299408f6e50ad8474f105c30a50b21e02ef9c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    ifbc53661a2b5de91bf09ee29f9ccb1ccf4f3326ea44c382276d002588effe1c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ie982b90d8af737d7d1f146569e62141c8cc9f946e39b4d08b956a605fc5512ed.ActivateServicePlanRequestBuilder) {
    return ie982b90d8af737d7d1f146569e62141c8cc9f946e39b4d08b956a605fc5512ed.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i6a7dae5caea32e17a8029984ac4c2f4ddacb4146a8ab183bab087a5da9be12b3.AssignLicenseRequestBuilder) {
    return i6a7dae5caea32e17a8029984ac4c2f4ddacb4146a8ab183bab087a5da9be12b3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*if4618f906f6cbc777a99c3d5fb1299408f6e50ad8474f105c30a50b21e02ef9c.ChangePasswordRequestBuilder) {
    return if4618f906f6cbc777a99c3d5fb1299408f6e50ad8474f105c30a50b21e02ef9c.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i1cb8dd35f65e9ea6d86f2bcc4bdde4a4ead0d78e24bb90152c10f4ed96856392.CheckMemberGroupsRequestBuilder) {
    return i1cb8dd35f65e9ea6d86f2bcc4bdde4a4ead0d78e24bb90152c10f4ed96856392.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ic6f503187de5eefff58d23d03d0039e87106c7cc896b2809237f78fb9680321e.CheckMemberObjectsRequestBuilder) {
    return ic6f503187de5eefff58d23d03d0039e87106c7cc896b2809237f78fb9680321e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*idf6baf10effbe604b3095bf580189ccf757a31545da79646983eef2e22f9996c.ExportDeviceAndAppManagementDataRequestBuilder) {
    return idf6baf10effbe604b3095bf580189ccf757a31545da79646983eef2e22f9996c.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ie9ebae9ee745ee1eb3f0d95c0d607beb875ef04ca992fb6f600de28a49e6efd7.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ie9ebae9ee745ee1eb3f0d95c0d607beb875ef04ca992fb6f600de28a49e6efd7.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ie9caa3c501e0e31b83a05c63ec77635292c0ba486cd3777a24ddb0f6509211f0.ExportPersonalDataRequestBuilder) {
    return ie9caa3c501e0e31b83a05c63ec77635292c0ba486cd3777a24ddb0f6509211f0.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i67c75a7976f24b55607f504127e865678bd4b5f9871cbd9be36d78ed5dffd2bd.FindMeetingTimesRequestBuilder) {
    return i67c75a7976f24b55607f504127e865678bd4b5f9871cbd9be36d78ed5dffd2bd.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i486324f589cf06d5592183853cedfb0cee61d7dda900d85ace70b3e3224e7d0e.FindRoomListsRequestBuilder) {
    return i486324f589cf06d5592183853cedfb0cee61d7dda900d85ace70b3e3224e7d0e.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ib0187e5f182839837ce4afdaef16eb467cf5ac2c3464cf94936e5e112b2a7e4f.FindRoomsRequestBuilder) {
    return ib0187e5f182839837ce4afdaef16eb467cf5ac2c3464cf94936e5e112b2a7e4f.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i5f1036ac0480dcf25fe3d7f1aaeef36ffeb055c64c1ccd3064f8d110e011da14.FindRoomsWithRoomListRequestBuilder) {
    return i5f1036ac0480dcf25fe3d7f1aaeef36ffeb055c64c1ccd3064f8d110e011da14.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ib31d5ea274f4b4ea0c1ed4271cc3c5ac5df472bf9184bb756cd9703c2a795219.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ib31d5ea274f4b4ea0c1ed4271cc3c5ac5df472bf9184bb756cd9703c2a795219.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i19ed68126d709a82540812bba60d7d51d19aa6e4f11317d33cae1f051fa17976.GetLoggedOnManagedDevicesRequestBuilder) {
    return i19ed68126d709a82540812bba60d7d51d19aa6e4f11317d33cae1f051fa17976.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i99cf4d22d13544c6fc31bd9ac0864216f9396a9b1f05c3e346c528b4ebb56dff.GetMailTipsRequestBuilder) {
    return i99cf4d22d13544c6fc31bd9ac0864216f9396a9b1f05c3e346c528b4ebb56dff.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i204e5c913a92bb6ea8e49f417d63cfac139988b7c827d54ee4894227a198d548.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i204e5c913a92bb6ea8e49f417d63cfac139988b7c827d54ee4894227a198d548.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i35e13707db8c4d3199ea15f14015264f6c3062a1ed76c2371213ae3ad7bbe1c1.GetManagedAppPoliciesRequestBuilder) {
    return i35e13707db8c4d3199ea15f14015264f6c3062a1ed76c2371213ae3ad7bbe1c1.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i56429eaf96d76c244fb3df17b1e4aee9764c53e2a2925774530a0411dddb9c3d.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i56429eaf96d76c244fb3df17b1e4aee9764c53e2a2925774530a0411dddb9c3d.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ibbbd3bb572dc63bbf390e13b9b30364de3115be110a603fc4e3e57165aa3f874.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ibbbd3bb572dc63bbf390e13b9b30364de3115be110a603fc4e3e57165aa3f874.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i696e0ee47219eb2ca67a75d716f4392392a61dcc9bfad41d9bca6246e382f0ca.GetMemberGroupsRequestBuilder) {
    return i696e0ee47219eb2ca67a75d716f4392392a61dcc9bfad41d9bca6246e382f0ca.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i485cad271b75388bf55bfca46b2e131f9db586cdef14cc9fbccd4c3a3b9269d0.GetMemberObjectsRequestBuilder) {
    return i485cad271b75388bf55bfca46b2e131f9db586cdef14cc9fbccd4c3a3b9269d0.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i6c4b761199124e4038eab42acaa34ea51786277b066ac95de0dda809d0b04670.InvalidateAllRefreshTokensRequestBuilder) {
    return i6c4b761199124e4038eab42acaa34ea51786277b066ac95de0dda809d0b04670.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1b3a6a82269384d6020b7c5a6325cde206132c9dee3b41d39f3ccabf4543d78e.IsManagedAppUserBlockedRequestBuilder) {
    return i1b3a6a82269384d6020b7c5a6325cde206132c9dee3b41d39f3ccabf4543d78e.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ie72d542a9047dbc1d7ffb48ce4a67aa918cad8042da2e57fdc297f43bc3769c7.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ie72d542a9047dbc1d7ffb48ce4a67aa918cad8042da2e57fdc297f43bc3769c7.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ifbc53661a2b5de91bf09ee29f9ccb1ccf4f3326ea44c382276d002588effe1c5.RemoveAllDevicesFromManagementRequestBuilder) {
    return ifbc53661a2b5de91bf09ee29f9ccb1ccf4f3326ea44c382276d002588effe1c5.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i24c5b9814f2073b9e7f22b3cf36f7daf51fefafb4eab1dcb3ef9cb884dc7c05f.ReprocessLicenseAssignmentRequestBuilder) {
    return i24c5b9814f2073b9e7f22b3cf36f7daf51fefafb4eab1dcb3ef9cb884dc7c05f.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i295376a9e87d529493ee50dc6d93198a6b847124ae1c563a6267e2af07295b56.RestoreRequestBuilder) {
    return i295376a9e87d529493ee50dc6d93198a6b847124ae1c563a6267e2af07295b56.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i69ff141cf7e8600fbbab56da2e82fee43766316164b8cfc8cf20b4d503edfaf7.RevokeSignInSessionsRequestBuilder) {
    return i69ff141cf7e8600fbbab56da2e82fee43766316164b8cfc8cf20b4d503edfaf7.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i8e4578a11153aaee68d9eca6b1d93eb4a3fb1f1d57548097c3f4166543c12eb6.SendMailRequestBuilder) {
    return i8e4578a11153aaee68d9eca6b1d93eb4a3fb1f1d57548097c3f4166543c12eb6.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9be732ff513a3fd2786b2be69abedf24b564f0fc38631c7cae0e3878ccbd21e1.TranslateExchangeIdsRequestBuilder) {
    return i9be732ff513a3fd2786b2be69abedf24b564f0fc38631c7cae0e3878ccbd21e1.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i4171b656c38fef1c57019bffe1cd480bd65d61be89eeea98dca8a54794e50916.UnblockManagedAppsRequestBuilder) {
    return i4171b656c38fef1c57019bffe1cd480bd65d61be89eeea98dca8a54794e50916.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i2d5cc840b7c0d226ff89b3d01a70db1b6bf8db09ae28f59e402a739b5335e084.WipeAndBlockManagedAppsRequestBuilder) {
    return i2d5cc840b7c0d226ff89b3d01a70db1b6bf8db09ae28f59e402a739b5335e084.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i576a0a21c18956a80aefd2114d581672fb443866f81a26a6026f233767daddf9.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i576a0a21c18956a80aefd2114d581672fb443866f81a26a6026f233767daddf9.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i8306a5528aba6dae56eba498ac4b7b6c079c2e6509ee97e62074cb84ebc7f90c.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i8306a5528aba6dae56eba498ac4b7b6c079c2e6509ee97e62074cb84ebc7f90c.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i799512ab2e2f840da65bec587cca61c8b31af3ac501b0d89f89669c6be7a65a6.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i799512ab2e2f840da65bec587cca61c8b31af3ac501b0d89f89669c6be7a65a6.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
