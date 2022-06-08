package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1d7a7554141836bdcee003fd4c049465dea6d1078d88e5616ab101bd3389afad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportpersonaldata"
    i20f91bb829989f2811477987d8a82f03085a31146adb18b5f1c28c63951aac28 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/assignlicense"
    i24b823f4db40431c6b2d0be028705b8dd75407e3a37cb502fd750096503c6da9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    i25c214c1cc7ebca02f8cc2abe580d2a183544ec233de3d13de921e0625d3f8c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i26e79d91f924fa2f6a894f7ba357f6c74e6864615b5e1549152f8d7980ccc13c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findroomswithroomlist"
    i2f846c17ade4154b345c54832408d618420a98ef50ddfe0ce0c74375c51691f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i3433a392fdf2ce096ff6f6f77e8829db0504a87532e53a765d22d656acea4d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    i49804a0831abe635646cff5756ae0ddb74944e26d246e7d32feb84179e1cad6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i4fecba04dead397a484c3915c0b16620e19774d94cf233d10cc269e2cf07d369 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmailtips"
    i5235673dbd19c30073ba8189b70240dcdd8f9a3937415a118a81b3f9bbc87bf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i57f6df23484f509627e67b3fbf6f7af26ed4bbeeca827499218fba443e4e81bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/revokesigninsessions"
    i6adb86feb1d4e974abd702032c3cb5b2249c2d83894663e471fee162617931ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    i6baa215710c6ef7edff4f879d62d3fe3f2e9a301d7f8ff190af8801188e389f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmembergroups"
    i6cc7260eb248d6217451be698e549b971e06fc6f1f2018d5e37496697a583860 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    i7a51931434bb2a78b982a5705f7a1fafc813012fb72f1118ef41291cf5a63583 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    i822fe509a4e09852b20a5a484276459f736b2390055b426ad2d6dd3bcdc3e703 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmembergroups"
    i9e1a1c05f6c7beb7a85e149e927a75da844dd507f3e2c54da71364e25f324c8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/translateexchangeids"
    i9e35c9481b221f077fa641dd0729157c61610d40625119cf4213a618d3d97754 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    ia47e28d788b456ebd59ab82f6f928bdff7657f2ddac5e59005fbf85fb2fa9885 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    iaac569684e6c06d7393a21b3a5b22f31f821130748153edcf0e25afc2f2b2253 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/restore"
    iad90ad56f876c444d8b2c8fc9e4a0eb1433f223cf56e55ed225918a5fc837818 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findroomlists"
    iaf818d33ec3625755af7ff7aca778eb6de9b220d665a264b16daf2448121f8e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
    ib7ac1efb5faf2f5c48d5b3a9978a85cd71c87dea80d606c771b479988a6096c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmemberobjects"
    ib8d564ac577a41a7786a09abeb1bd9d9ee9177710042dcdf611f347dba502847 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findrooms"
    ib9999b194276cd51503512299f4dd8da7d28e798bb856e281cd3473e790fb577 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/unblockmanagedapps"
    ic5be7822fb8e592d9fd81c28766cb29ea1df937d0fed0e4521eaaa900a3c0206 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/activateserviceplan"
    id27e078b36fc3779e927b614107657e0e83ad4534e257075b721f0a81aa1d3a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findmeetingtimes"
    id45980f16647edd03561a14eab8536189e7a2c6c90dd22373d80ebc2f57af9cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
    id6443c699df84b1f87db15ff852b293c9dc5a22a8315bbf84e27c1fd87a5a8dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    ie1c8f4864c72f21c63b324793d52769bc7e77a79d89f28fcdeee7bef0eddcee7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    ie496d3f454690a9344aad47522df68f57042cd1f780b5317daade1b34a23c52a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmemberobjects"
    ie4d6a8e31fec0b137cf00cd351105f76da0c44e2a84661bae29f13067f7af58a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    ie6e7fd14ccdc8f08233f27fc319529a881080cc229046362adaa3386c0c1b7d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/sendmail"
    ie7df216d2417c127dc0e9e47133473f9655238456a6a5c47184006ee80fc918f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    if44d50d843c9f8a74c6daa49837276c1bbfb74f021722dad78e2335582546253 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/changepassword"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ic5be7822fb8e592d9fd81c28766cb29ea1df937d0fed0e4521eaaa900a3c0206.ActivateServicePlanRequestBuilder) {
    return ic5be7822fb8e592d9fd81c28766cb29ea1df937d0fed0e4521eaaa900a3c0206.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i20f91bb829989f2811477987d8a82f03085a31146adb18b5f1c28c63951aac28.AssignLicenseRequestBuilder) {
    return i20f91bb829989f2811477987d8a82f03085a31146adb18b5f1c28c63951aac28.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*if44d50d843c9f8a74c6daa49837276c1bbfb74f021722dad78e2335582546253.ChangePasswordRequestBuilder) {
    return if44d50d843c9f8a74c6daa49837276c1bbfb74f021722dad78e2335582546253.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i822fe509a4e09852b20a5a484276459f736b2390055b426ad2d6dd3bcdc3e703.CheckMemberGroupsRequestBuilder) {
    return i822fe509a4e09852b20a5a484276459f736b2390055b426ad2d6dd3bcdc3e703.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ib7ac1efb5faf2f5c48d5b3a9978a85cd71c87dea80d606c771b479988a6096c8.CheckMemberObjectsRequestBuilder) {
    return ib7ac1efb5faf2f5c48d5b3a9978a85cd71c87dea80d606c771b479988a6096c8.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*iaf818d33ec3625755af7ff7aca778eb6de9b220d665a264b16daf2448121f8e0.ExportDeviceAndAppManagementDataRequestBuilder) {
    return iaf818d33ec3625755af7ff7aca778eb6de9b220d665a264b16daf2448121f8e0.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ia47e28d788b456ebd59ab82f6f928bdff7657f2ddac5e59005fbf85fb2fa9885.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ia47e28d788b456ebd59ab82f6f928bdff7657f2ddac5e59005fbf85fb2fa9885.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i1d7a7554141836bdcee003fd4c049465dea6d1078d88e5616ab101bd3389afad.ExportPersonalDataRequestBuilder) {
    return i1d7a7554141836bdcee003fd4c049465dea6d1078d88e5616ab101bd3389afad.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*id27e078b36fc3779e927b614107657e0e83ad4534e257075b721f0a81aa1d3a1.FindMeetingTimesRequestBuilder) {
    return id27e078b36fc3779e927b614107657e0e83ad4534e257075b721f0a81aa1d3a1.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*iad90ad56f876c444d8b2c8fc9e4a0eb1433f223cf56e55ed225918a5fc837818.FindRoomListsRequestBuilder) {
    return iad90ad56f876c444d8b2c8fc9e4a0eb1433f223cf56e55ed225918a5fc837818.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ib8d564ac577a41a7786a09abeb1bd9d9ee9177710042dcdf611f347dba502847.FindRoomsRequestBuilder) {
    return ib8d564ac577a41a7786a09abeb1bd9d9ee9177710042dcdf611f347dba502847.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i26e79d91f924fa2f6a894f7ba357f6c74e6864615b5e1549152f8d7980ccc13c.FindRoomsWithRoomListRequestBuilder) {
    return i26e79d91f924fa2f6a894f7ba357f6c74e6864615b5e1549152f8d7980ccc13c.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i24b823f4db40431c6b2d0be028705b8dd75407e3a37cb502fd750096503c6da9.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i24b823f4db40431c6b2d0be028705b8dd75407e3a37cb502fd750096503c6da9.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*id45980f16647edd03561a14eab8536189e7a2c6c90dd22373d80ebc2f57af9cc.GetLoggedOnManagedDevicesRequestBuilder) {
    return id45980f16647edd03561a14eab8536189e7a2c6c90dd22373d80ebc2f57af9cc.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i4fecba04dead397a484c3915c0b16620e19774d94cf233d10cc269e2cf07d369.GetMailTipsRequestBuilder) {
    return i4fecba04dead397a484c3915c0b16620e19774d94cf233d10cc269e2cf07d369.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i5235673dbd19c30073ba8189b70240dcdd8f9a3937415a118a81b3f9bbc87bf1.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i5235673dbd19c30073ba8189b70240dcdd8f9a3937415a118a81b3f9bbc87bf1.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i6adb86feb1d4e974abd702032c3cb5b2249c2d83894663e471fee162617931ee.GetManagedAppPoliciesRequestBuilder) {
    return i6adb86feb1d4e974abd702032c3cb5b2249c2d83894663e471fee162617931ee.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i6cc7260eb248d6217451be698e549b971e06fc6f1f2018d5e37496697a583860.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i6cc7260eb248d6217451be698e549b971e06fc6f1f2018d5e37496697a583860.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i3433a392fdf2ce096ff6f6f77e8829db0504a87532e53a765d22d656acea4d58.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i3433a392fdf2ce096ff6f6f77e8829db0504a87532e53a765d22d656acea4d58.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i6baa215710c6ef7edff4f879d62d3fe3f2e9a301d7f8ff190af8801188e389f2.GetMemberGroupsRequestBuilder) {
    return i6baa215710c6ef7edff4f879d62d3fe3f2e9a301d7f8ff190af8801188e389f2.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ie496d3f454690a9344aad47522df68f57042cd1f780b5317daade1b34a23c52a.GetMemberObjectsRequestBuilder) {
    return ie496d3f454690a9344aad47522df68f57042cd1f780b5317daade1b34a23c52a.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i7a51931434bb2a78b982a5705f7a1fafc813012fb72f1118ef41291cf5a63583.InvalidateAllRefreshTokensRequestBuilder) {
    return i7a51931434bb2a78b982a5705f7a1fafc813012fb72f1118ef41291cf5a63583.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ie1c8f4864c72f21c63b324793d52769bc7e77a79d89f28fcdeee7bef0eddcee7.IsManagedAppUserBlockedRequestBuilder) {
    return ie1c8f4864c72f21c63b324793d52769bc7e77a79d89f28fcdeee7bef0eddcee7.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i9e35c9481b221f077fa641dd0729157c61610d40625119cf4213a618d3d97754.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i9e35c9481b221f077fa641dd0729157c61610d40625119cf4213a618d3d97754.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id6443c699df84b1f87db15ff852b293c9dc5a22a8315bbf84e27c1fd87a5a8dc.RemoveAllDevicesFromManagementRequestBuilder) {
    return id6443c699df84b1f87db15ff852b293c9dc5a22a8315bbf84e27c1fd87a5a8dc.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i25c214c1cc7ebca02f8cc2abe580d2a183544ec233de3d13de921e0625d3f8c5.ReprocessLicenseAssignmentRequestBuilder) {
    return i25c214c1cc7ebca02f8cc2abe580d2a183544ec233de3d13de921e0625d3f8c5.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*iaac569684e6c06d7393a21b3a5b22f31f821130748153edcf0e25afc2f2b2253.RestoreRequestBuilder) {
    return iaac569684e6c06d7393a21b3a5b22f31f821130748153edcf0e25afc2f2b2253.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i57f6df23484f509627e67b3fbf6f7af26ed4bbeeca827499218fba443e4e81bc.RevokeSignInSessionsRequestBuilder) {
    return i57f6df23484f509627e67b3fbf6f7af26ed4bbeeca827499218fba443e4e81bc.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie6e7fd14ccdc8f08233f27fc319529a881080cc229046362adaa3386c0c1b7d8.SendMailRequestBuilder) {
    return ie6e7fd14ccdc8f08233f27fc319529a881080cc229046362adaa3386c0c1b7d8.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9e1a1c05f6c7beb7a85e149e927a75da844dd507f3e2c54da71364e25f324c8e.TranslateExchangeIdsRequestBuilder) {
    return i9e1a1c05f6c7beb7a85e149e927a75da844dd507f3e2c54da71364e25f324c8e.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ib9999b194276cd51503512299f4dd8da7d28e798bb856e281cd3473e790fb577.UnblockManagedAppsRequestBuilder) {
    return ib9999b194276cd51503512299f4dd8da7d28e798bb856e281cd3473e790fb577.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ie4d6a8e31fec0b137cf00cd351105f76da0c44e2a84661bae29f13067f7af58a.WipeAndBlockManagedAppsRequestBuilder) {
    return ie4d6a8e31fec0b137cf00cd351105f76da0c44e2a84661bae29f13067f7af58a.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ie7df216d2417c127dc0e9e47133473f9655238456a6a5c47184006ee80fc918f.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ie7df216d2417c127dc0e9e47133473f9655238456a6a5c47184006ee80fc918f.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i2f846c17ade4154b345c54832408d618420a98ef50ddfe0ce0c74375c51691f6.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i2f846c17ade4154b345c54832408d618420a98ef50ddfe0ce0c74375c51691f6.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i49804a0831abe635646cff5756ae0ddb74944e26d246e7d32feb84179e1cad6d.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i49804a0831abe635646cff5756ae0ddb74944e26d246e7d32feb84179e1cad6d.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
