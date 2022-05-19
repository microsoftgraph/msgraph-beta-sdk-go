package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04be2cecc4c300da4ecaa0432e75f2d8eb31fee2e80c803c8bda38680f8cc502 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmembergroups"
    i0894843d665d7b2e46679c258eced7236a968226700ec2e9248782d2187c0865 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i1683b626ac3d1f4833b8f996f95a91d26aa753d8aff232f50d1c00a598f2fbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i18ef6cd1edecc2dfbb9980f7c44438f1cda41cf58474c09505f46d8da1f6552c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findmeetingtimes"
    i1dc4bc27efb6baba7ef8051199356d9bfc34ea42ee38317993847fce175a64c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/changepassword"
    i2026f0581185ea22125f00dae84fe949dd402c3701bd9849df13a1c7f5c3e31d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    i2087e981b923c0d96a6dcd7b6d55bb17104b96375cd842e0d776e6312458ea47 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    i23b545a15c0eafc6666fbf38ad10f672d94120fd61222f5780d3c2b34c3cdc0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findroomlists"
    i2a8ef37fcaa79b5c347856a7cbf4ea2314b43eaa29a35dc0079f782bb124e936 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i3fa50447af209da1d89cb7562fc8b41a17175c267dd64246e2a28dd816bef5d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/restore"
    i3fbce24b83be4284e7d3bdca2d0fd24a77a5d6b863236b34e0a42336641c9d9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/sendmail"
    i4293ea05ff19bd8a07a35a221ec116a339d641e2343a8bbe5c2803e59d000870 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/assignlicense"
    i4591b18115f05b37f145764bd5962f05233cfdf0c1dc1c887302ef22ad020bd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findrooms"
    i58bf1b40a24d87bd46ad5ef5caca38e2607ede048d07a202d25b0d5d9791f96d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    i6d9e27eb13d6ca22917bfebc861d2d82738d54ae094622bf5144490e5d2cce7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmembergroups"
    i7b892c073aa78e2c8f859ff77bbb8f6f7407013c746fc349e7faf4c5694fdb6f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    i7cee93e99bb961c457eed32be8055c914fa78601b783b8238bcdde11eb06317f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i84eeaeb5a7fcbb63b8903ce4a5ab01d5308e5d8ceaddaaa79965cf30d965c47a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i8d6fd017c5bf98ee4b805092ae88be70b182eab28d1b2eed99f07acdc0071449 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    ia35eab144665019e8794b813d235750bcb360e91670b02b3407ab664efab817c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    ia5965697734f93884ba2b5eb5ec84acb4d42268f1c9a430dc4947f18693cf009 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportpersonaldata"
    ib68f04ed89460b82663c165697a642fcb8814234ed7be338202c294dc6bab3a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/revokesigninsessions"
    ib9c1d1c825156f16328f273a352774cf4fbbd0f553e1b7e71065bd56968d44d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    ic4bab54d8a7be797272f3910b0707c5b34f217a6c05827d427b512b8368796f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/unblockmanagedapps"
    icbf88501d7ced1ac11a5b9388b964bc81a84ba36815e366f1fffc8eca1d02e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    icf5095f5514ebaddf37609f046bf76e6569ba175257314bbf4dfd3c7dd8c4c32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmailtips"
    id5a0c8f196b9209b627ce6cf8da126bf045c509b46493721bcfc27c84ae7866f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findroomswithroomlist"
    id5af24e0bd4e0218731413457fd7736309619a794daa5042bdf851de00ccf78c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    id82afc7cdbd31053ee3e3da2382e101c88149f2b19b95c41a63c4df5d900ac50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    id90ab6a9b46e27c96908fa5d8330e4b617c647f9267c356864beae80cd5b44c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    idb3b9d01d49b42616f2e9b3e7de22edabd033c30d201498eb5de53d570e9f740 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmemberobjects"
    idc4ff3b34bb0d1d9f58d8a64f94769b50c7b66641b858535e6f7d0812793a12e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmemberobjects"
    idf674db88f51845dbeb372c6d05e090a5add02f1e5cbb74ebdb2ab05fd0c4ad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/activateserviceplan"
    ied46e2bb6c9d6095adbb39a331806ad02b957b12a241b404998d21708cdf4a67 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    if784a48870b6a2f0eac8ac59027dbb554e2f63b8ebc696c4291efa024341b0ae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/translateexchangeids"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*idf674db88f51845dbeb372c6d05e090a5add02f1e5cbb74ebdb2ab05fd0c4ad7.ActivateServicePlanRequestBuilder) {
    return idf674db88f51845dbeb372c6d05e090a5add02f1e5cbb74ebdb2ab05fd0c4ad7.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i4293ea05ff19bd8a07a35a221ec116a339d641e2343a8bbe5c2803e59d000870.AssignLicenseRequestBuilder) {
    return i4293ea05ff19bd8a07a35a221ec116a339d641e2343a8bbe5c2803e59d000870.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i1dc4bc27efb6baba7ef8051199356d9bfc34ea42ee38317993847fce175a64c1.ChangePasswordRequestBuilder) {
    return i1dc4bc27efb6baba7ef8051199356d9bfc34ea42ee38317993847fce175a64c1.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i6d9e27eb13d6ca22917bfebc861d2d82738d54ae094622bf5144490e5d2cce7c.CheckMemberGroupsRequestBuilder) {
    return i6d9e27eb13d6ca22917bfebc861d2d82738d54ae094622bf5144490e5d2cce7c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*idb3b9d01d49b42616f2e9b3e7de22edabd033c30d201498eb5de53d570e9f740.CheckMemberObjectsRequestBuilder) {
    return idb3b9d01d49b42616f2e9b3e7de22edabd033c30d201498eb5de53d570e9f740.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ied46e2bb6c9d6095adbb39a331806ad02b957b12a241b404998d21708cdf4a67.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ied46e2bb6c9d6095adbb39a331806ad02b957b12a241b404998d21708cdf4a67.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i7cee93e99bb961c457eed32be8055c914fa78601b783b8238bcdde11eb06317f.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i7cee93e99bb961c457eed32be8055c914fa78601b783b8238bcdde11eb06317f.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ia5965697734f93884ba2b5eb5ec84acb4d42268f1c9a430dc4947f18693cf009.ExportPersonalDataRequestBuilder) {
    return ia5965697734f93884ba2b5eb5ec84acb4d42268f1c9a430dc4947f18693cf009.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i18ef6cd1edecc2dfbb9980f7c44438f1cda41cf58474c09505f46d8da1f6552c.FindMeetingTimesRequestBuilder) {
    return i18ef6cd1edecc2dfbb9980f7c44438f1cda41cf58474c09505f46d8da1f6552c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i23b545a15c0eafc6666fbf38ad10f672d94120fd61222f5780d3c2b34c3cdc0f.FindRoomListsRequestBuilder) {
    return i23b545a15c0eafc6666fbf38ad10f672d94120fd61222f5780d3c2b34c3cdc0f.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i4591b18115f05b37f145764bd5962f05233cfdf0c1dc1c887302ef22ad020bd7.FindRoomsRequestBuilder) {
    return i4591b18115f05b37f145764bd5962f05233cfdf0c1dc1c887302ef22ad020bd7.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*id5a0c8f196b9209b627ce6cf8da126bf045c509b46493721bcfc27c84ae7866f.FindRoomsWithRoomListRequestBuilder) {
    return id5a0c8f196b9209b627ce6cf8da126bf045c509b46493721bcfc27c84ae7866f.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i7b892c073aa78e2c8f859ff77bbb8f6f7407013c746fc349e7faf4c5694fdb6f.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i7b892c073aa78e2c8f859ff77bbb8f6f7407013c746fc349e7faf4c5694fdb6f.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*id90ab6a9b46e27c96908fa5d8330e4b617c647f9267c356864beae80cd5b44c4.GetLoggedOnManagedDevicesRequestBuilder) {
    return id90ab6a9b46e27c96908fa5d8330e4b617c647f9267c356864beae80cd5b44c4.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*icf5095f5514ebaddf37609f046bf76e6569ba175257314bbf4dfd3c7dd8c4c32.GetMailTipsRequestBuilder) {
    return icf5095f5514ebaddf37609f046bf76e6569ba175257314bbf4dfd3c7dd8c4c32.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ia35eab144665019e8794b813d235750bcb360e91670b02b3407ab664efab817c.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ia35eab144665019e8794b813d235750bcb360e91670b02b3407ab664efab817c.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i58bf1b40a24d87bd46ad5ef5caca38e2607ede048d07a202d25b0d5d9791f96d.GetManagedAppPoliciesRequestBuilder) {
    return i58bf1b40a24d87bd46ad5ef5caca38e2607ede048d07a202d25b0d5d9791f96d.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*icbf88501d7ced1ac11a5b9388b964bc81a84ba36815e366f1fffc8eca1d02e3d.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return icbf88501d7ced1ac11a5b9388b964bc81a84ba36815e366f1fffc8eca1d02e3d.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*id5af24e0bd4e0218731413457fd7736309619a794daa5042bdf851de00ccf78c.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return id5af24e0bd4e0218731413457fd7736309619a794daa5042bdf851de00ccf78c.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i04be2cecc4c300da4ecaa0432e75f2d8eb31fee2e80c803c8bda38680f8cc502.GetMemberGroupsRequestBuilder) {
    return i04be2cecc4c300da4ecaa0432e75f2d8eb31fee2e80c803c8bda38680f8cc502.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*idc4ff3b34bb0d1d9f58d8a64f94769b50c7b66641b858535e6f7d0812793a12e.GetMemberObjectsRequestBuilder) {
    return idc4ff3b34bb0d1d9f58d8a64f94769b50c7b66641b858535e6f7d0812793a12e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i2026f0581185ea22125f00dae84fe949dd402c3701bd9849df13a1c7f5c3e31d.InvalidateAllRefreshTokensRequestBuilder) {
    return i2026f0581185ea22125f00dae84fe949dd402c3701bd9849df13a1c7f5c3e31d.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i2087e981b923c0d96a6dcd7b6d55bb17104b96375cd842e0d776e6312458ea47.IsManagedAppUserBlockedRequestBuilder) {
    return i2087e981b923c0d96a6dcd7b6d55bb17104b96375cd842e0d776e6312458ea47.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i84eeaeb5a7fcbb63b8903ce4a5ab01d5308e5d8ceaddaaa79965cf30d965c47a.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i84eeaeb5a7fcbb63b8903ce4a5ab01d5308e5d8ceaddaaa79965cf30d965c47a.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ib9c1d1c825156f16328f273a352774cf4fbbd0f553e1b7e71065bd56968d44d9.RemoveAllDevicesFromManagementRequestBuilder) {
    return ib9c1d1c825156f16328f273a352774cf4fbbd0f553e1b7e71065bd56968d44d9.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0894843d665d7b2e46679c258eced7236a968226700ec2e9248782d2187c0865.ReprocessLicenseAssignmentRequestBuilder) {
    return i0894843d665d7b2e46679c258eced7236a968226700ec2e9248782d2187c0865.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i3fa50447af209da1d89cb7562fc8b41a17175c267dd64246e2a28dd816bef5d3.RestoreRequestBuilder) {
    return i3fa50447af209da1d89cb7562fc8b41a17175c267dd64246e2a28dd816bef5d3.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ib68f04ed89460b82663c165697a642fcb8814234ed7be338202c294dc6bab3a6.RevokeSignInSessionsRequestBuilder) {
    return ib68f04ed89460b82663c165697a642fcb8814234ed7be338202c294dc6bab3a6.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i3fbce24b83be4284e7d3bdca2d0fd24a77a5d6b863236b34e0a42336641c9d9a.SendMailRequestBuilder) {
    return i3fbce24b83be4284e7d3bdca2d0fd24a77a5d6b863236b34e0a42336641c9d9a.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*if784a48870b6a2f0eac8ac59027dbb554e2f63b8ebc696c4291efa024341b0ae.TranslateExchangeIdsRequestBuilder) {
    return if784a48870b6a2f0eac8ac59027dbb554e2f63b8ebc696c4291efa024341b0ae.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ic4bab54d8a7be797272f3910b0707c5b34f217a6c05827d427b512b8368796f0.UnblockManagedAppsRequestBuilder) {
    return ic4bab54d8a7be797272f3910b0707c5b34f217a6c05827d427b512b8368796f0.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i2a8ef37fcaa79b5c347856a7cbf4ea2314b43eaa29a35dc0079f782bb124e936.WipeAndBlockManagedAppsRequestBuilder) {
    return i2a8ef37fcaa79b5c347856a7cbf4ea2314b43eaa29a35dc0079f782bb124e936.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i8d6fd017c5bf98ee4b805092ae88be70b182eab28d1b2eed99f07acdc0071449.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i8d6fd017c5bf98ee4b805092ae88be70b182eab28d1b2eed99f07acdc0071449.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i1683b626ac3d1f4833b8f996f95a91d26aa753d8aff232f50d1c00a598f2fbcd.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i1683b626ac3d1f4833b8f996f95a91d26aa753d8aff232f50d1c00a598f2fbcd.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*id82afc7cdbd31053ee3e3da2382e101c88149f2b19b95c41a63c4df5d900ac50.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return id82afc7cdbd31053ee3e3da2382e101c88149f2b19b95c41a63c4df5d900ac50.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
