package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1258bff53ee3dfe3034024d8db6bb122ef83ef5ea72ee410710d9bdc6e0dec26 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/reprocesslicenseassignment"
    i12f3b3e700ecf5d7e8932ed3066219b78f5f1b7c39d64ecf883f73b6bd884077 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/findrooms"
    i208ab3fcbe3e068ccbd9d8556b3b258ee6fba1da5f234b47fb2d03d063df4df6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/exportdeviceandappmanagementdata"
    i2635347242f0b58806e9a29c0209cecde4c963551dc9d0cff4cfe05936d31c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/unblockmanagedapps"
    i28ae98ad9772ef86234e9f3d7c7ef1a8d057150817af4dca31d06638c8837aad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/findroomlists"
    i2b90d213d91a1340b4f1c06bc24b7662eb7f6b55a291a65a0dc79bed2861a56f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/ismanagedappuserblocked"
    i39060fdf532e71ac4bb3e79334a90b16b9e5b5832af56584feea297611dbbcea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i3fc88f18cef9e15d58ac0633413e5843446c388aa9945d3c05d0d0d01f52dcf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i429daee8b2df31247896347aac5bf2cbc3377c00cd2e379e995ede303399aa41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getloggedonmanageddevices"
    i4750fa3fc91f30391a65f8c63346fabc75748081226f11e20fdb29c28830b90f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmembergroups"
    i725ed0a57fe8f1f0efb8899b2329549d1c5e7ab922fdef8d6e2dc4d2f7151e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/invalidateallrefreshtokens"
    i76c89b91b5a4b730248b0c7f91915a18ff0571093fce1f267ffdc785569d3cf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmanagedappdiagnosticstatuses"
    i7fc0e1a7e8ec38efdd82faf9c1019cd206c649bf91510796e527c888ff2b3fa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmemberobjects"
    i827c3bd7b61c945d6f9cd6376bf9f05c14d5e0afef6869b666dbe27a99204376 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i8840773107d9093d04ae3d54a67fd1ab01751678b7a2a38d975cabf3b1c5b71b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/revokesigninsessions"
    i899264b6284d40a7357d9aa3233de4391ecfbbc923d09a4946c4f1fe7740eadf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/findmeetingtimes"
    i91c69c5d7708e73c966099b8f19d762750b05f9371773de93772bceb7be99521 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmailtips"
    i93731d4e550378da6953ba0737848017267c093aa10bca8157d4de43630f97f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i9fc260b11f5372ee47da752055b9777ee64458f1d4a0aaa8cf72ff68170bf364 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/findroomswithroomlist"
    ia399731e702081cae3690de8e22d3e8eefdb892175c18ecf818ed66b46b98d9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/removealldevicesfrommanagement"
    ib1e1896aaa7967668f9e691f0222c71b5a97d4bf02a064b3c17efbe98fd6d74a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/sendmail"
    ib6fd40c35ab05fa404f49bcf9bf487e888b689574abeb01ca58b256f65c46bea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/restore"
    ib99f4dfcca46caae3af244b66e19de57e185f9ffd4daf988865869aaa85a6f95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/checkmembergroups"
    ibdaaac51708da422da34c4a90d25832ca6532fac4ac1a48c312e3b374b402864 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/checkmemberobjects"
    ic1e3ae09bd9f0e6a697303b1b69c13f9f10d183df0fd719b3e899c2cfdf7c9cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/exportpersonaldata"
    ic279f6b967f06c6dc01dc868dceb5209c26521234317433b5928a0999474c0f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/changepassword"
    ic2eccbee5aaa43fc4f43d3e3cb872659a8633d2951ce56207e949027a0424d17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/activateserviceplan"
    ie417b018c7ae6bea2e4ea7101a187563743ee3f9bebbd1c94f45b246fe28b588 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/translateexchangeids"
    ie7b9eb0f22d55dfdb4345afd0907e4c33ac4f5347b74a25ea8f6a83588ee491c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ie812dd24f9476a11ff876469473aca1deddbd27cd83dad63a8c42991c68d5bc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/wipemanagedappregistrationbydevicetag"
    ie925c1730f074903a3e80f9c8bb6c5cd4e675924d01874ead821bb982d35f892 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmanagedapppolicies"
    if13175187724d41b86f44d31375823a94dc7ca267a2307b0786cb1a4fd75a6a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/getmanageddeviceswithappfailures"
    if4c038870c3b12ef0816636c9d3f772723383d86cae5ed0904010d1e10d5eb30 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    if985b193362059bf57d1be694cad52bff933890b4a3349a600ddb74c37eaf662 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/wipeandblockmanagedapps"
    ifd3ff375f5782eea0f4370a6387ff1a7bde950c701f958b4696a40e65a72fa5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user/assignlicense"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ic2eccbee5aaa43fc4f43d3e3cb872659a8633d2951ce56207e949027a0424d17.ActivateServicePlanRequestBuilder) {
    return ic2eccbee5aaa43fc4f43d3e3cb872659a8633d2951ce56207e949027a0424d17.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ifd3ff375f5782eea0f4370a6387ff1a7bde950c701f958b4696a40e65a72fa5d.AssignLicenseRequestBuilder) {
    return ifd3ff375f5782eea0f4370a6387ff1a7bde950c701f958b4696a40e65a72fa5d.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ic279f6b967f06c6dc01dc868dceb5209c26521234317433b5928a0999474c0f3.ChangePasswordRequestBuilder) {
    return ic279f6b967f06c6dc01dc868dceb5209c26521234317433b5928a0999474c0f3.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ib99f4dfcca46caae3af244b66e19de57e185f9ffd4daf988865869aaa85a6f95.CheckMemberGroupsRequestBuilder) {
    return ib99f4dfcca46caae3af244b66e19de57e185f9ffd4daf988865869aaa85a6f95.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ibdaaac51708da422da34c4a90d25832ca6532fac4ac1a48c312e3b374b402864.CheckMemberObjectsRequestBuilder) {
    return ibdaaac51708da422da34c4a90d25832ca6532fac4ac1a48c312e3b374b402864.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i208ab3fcbe3e068ccbd9d8556b3b258ee6fba1da5f234b47fb2d03d063df4df6.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i208ab3fcbe3e068ccbd9d8556b3b258ee6fba1da5f234b47fb2d03d063df4df6.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*if4c038870c3b12ef0816636c9d3f772723383d86cae5ed0904010d1e10d5eb30.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return if4c038870c3b12ef0816636c9d3f772723383d86cae5ed0904010d1e10d5eb30.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ic1e3ae09bd9f0e6a697303b1b69c13f9f10d183df0fd719b3e899c2cfdf7c9cf.ExportPersonalDataRequestBuilder) {
    return ic1e3ae09bd9f0e6a697303b1b69c13f9f10d183df0fd719b3e899c2cfdf7c9cf.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i899264b6284d40a7357d9aa3233de4391ecfbbc923d09a4946c4f1fe7740eadf.FindMeetingTimesRequestBuilder) {
    return i899264b6284d40a7357d9aa3233de4391ecfbbc923d09a4946c4f1fe7740eadf.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i28ae98ad9772ef86234e9f3d7c7ef1a8d057150817af4dca31d06638c8837aad.FindRoomListsRequestBuilder) {
    return i28ae98ad9772ef86234e9f3d7c7ef1a8d057150817af4dca31d06638c8837aad.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i12f3b3e700ecf5d7e8932ed3066219b78f5f1b7c39d64ecf883f73b6bd884077.FindRoomsRequestBuilder) {
    return i12f3b3e700ecf5d7e8932ed3066219b78f5f1b7c39d64ecf883f73b6bd884077.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i9fc260b11f5372ee47da752055b9777ee64458f1d4a0aaa8cf72ff68170bf364.FindRoomsWithRoomListRequestBuilder) {
    return i9fc260b11f5372ee47da752055b9777ee64458f1d4a0aaa8cf72ff68170bf364.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ie7b9eb0f22d55dfdb4345afd0907e4c33ac4f5347b74a25ea8f6a83588ee491c.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ie7b9eb0f22d55dfdb4345afd0907e4c33ac4f5347b74a25ea8f6a83588ee491c.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i429daee8b2df31247896347aac5bf2cbc3377c00cd2e379e995ede303399aa41.GetLoggedOnManagedDevicesRequestBuilder) {
    return i429daee8b2df31247896347aac5bf2cbc3377c00cd2e379e995ede303399aa41.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i91c69c5d7708e73c966099b8f19d762750b05f9371773de93772bceb7be99521.GetMailTipsRequestBuilder) {
    return i91c69c5d7708e73c966099b8f19d762750b05f9371773de93772bceb7be99521.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i76c89b91b5a4b730248b0c7f91915a18ff0571093fce1f267ffdc785569d3cf4.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i76c89b91b5a4b730248b0c7f91915a18ff0571093fce1f267ffdc785569d3cf4.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ie925c1730f074903a3e80f9c8bb6c5cd4e675924d01874ead821bb982d35f892.GetManagedAppPoliciesRequestBuilder) {
    return ie925c1730f074903a3e80f9c8bb6c5cd4e675924d01874ead821bb982d35f892.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*if13175187724d41b86f44d31375823a94dc7ca267a2307b0786cb1a4fd75a6a5.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return if13175187724d41b86f44d31375823a94dc7ca267a2307b0786cb1a4fd75a6a5.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i93731d4e550378da6953ba0737848017267c093aa10bca8157d4de43630f97f3.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i93731d4e550378da6953ba0737848017267c093aa10bca8157d4de43630f97f3.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i4750fa3fc91f30391a65f8c63346fabc75748081226f11e20fdb29c28830b90f.GetMemberGroupsRequestBuilder) {
    return i4750fa3fc91f30391a65f8c63346fabc75748081226f11e20fdb29c28830b90f.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i7fc0e1a7e8ec38efdd82faf9c1019cd206c649bf91510796e527c888ff2b3fa2.GetMemberObjectsRequestBuilder) {
    return i7fc0e1a7e8ec38efdd82faf9c1019cd206c649bf91510796e527c888ff2b3fa2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i725ed0a57fe8f1f0efb8899b2329549d1c5e7ab922fdef8d6e2dc4d2f7151e4e.InvalidateAllRefreshTokensRequestBuilder) {
    return i725ed0a57fe8f1f0efb8899b2329549d1c5e7ab922fdef8d6e2dc4d2f7151e4e.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i2b90d213d91a1340b4f1c06bc24b7662eb7f6b55a291a65a0dc79bed2861a56f.IsManagedAppUserBlockedRequestBuilder) {
    return i2b90d213d91a1340b4f1c06bc24b7662eb7f6b55a291a65a0dc79bed2861a56f.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i3fc88f18cef9e15d58ac0633413e5843446c388aa9945d3c05d0d0d01f52dcf9.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i3fc88f18cef9e15d58ac0633413e5843446c388aa9945d3c05d0d0d01f52dcf9.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ia399731e702081cae3690de8e22d3e8eefdb892175c18ecf818ed66b46b98d9d.RemoveAllDevicesFromManagementRequestBuilder) {
    return ia399731e702081cae3690de8e22d3e8eefdb892175c18ecf818ed66b46b98d9d.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i1258bff53ee3dfe3034024d8db6bb122ef83ef5ea72ee410710d9bdc6e0dec26.ReprocessLicenseAssignmentRequestBuilder) {
    return i1258bff53ee3dfe3034024d8db6bb122ef83ef5ea72ee410710d9bdc6e0dec26.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ib6fd40c35ab05fa404f49bcf9bf487e888b689574abeb01ca58b256f65c46bea.RestoreRequestBuilder) {
    return ib6fd40c35ab05fa404f49bcf9bf487e888b689574abeb01ca58b256f65c46bea.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i8840773107d9093d04ae3d54a67fd1ab01751678b7a2a38d975cabf3b1c5b71b.RevokeSignInSessionsRequestBuilder) {
    return i8840773107d9093d04ae3d54a67fd1ab01751678b7a2a38d975cabf3b1c5b71b.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ib1e1896aaa7967668f9e691f0222c71b5a97d4bf02a064b3c17efbe98fd6d74a.SendMailRequestBuilder) {
    return ib1e1896aaa7967668f9e691f0222c71b5a97d4bf02a064b3c17efbe98fd6d74a.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ie417b018c7ae6bea2e4ea7101a187563743ee3f9bebbd1c94f45b246fe28b588.TranslateExchangeIdsRequestBuilder) {
    return ie417b018c7ae6bea2e4ea7101a187563743ee3f9bebbd1c94f45b246fe28b588.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i2635347242f0b58806e9a29c0209cecde4c963551dc9d0cff4cfe05936d31c3b.UnblockManagedAppsRequestBuilder) {
    return i2635347242f0b58806e9a29c0209cecde4c963551dc9d0cff4cfe05936d31c3b.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*if985b193362059bf57d1be694cad52bff933890b4a3349a600ddb74c37eaf662.WipeAndBlockManagedAppsRequestBuilder) {
    return if985b193362059bf57d1be694cad52bff933890b4a3349a600ddb74c37eaf662.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ie812dd24f9476a11ff876469473aca1deddbd27cd83dad63a8c42991c68d5bc3.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ie812dd24f9476a11ff876469473aca1deddbd27cd83dad63a8c42991c68d5bc3.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i827c3bd7b61c945d6f9cd6376bf9f05c14d5e0afef6869b666dbe27a99204376.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i827c3bd7b61c945d6f9cd6376bf9f05c14d5e0afef6869b666dbe27a99204376.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i39060fdf532e71ac4bb3e79334a90b16b9e5b5832af56584feea297611dbbcea.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i39060fdf532e71ac4bb3e79334a90b16b9e5b5832af56584feea297611dbbcea.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
