package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07ea9426dc558b5b0c329fa6629e6450a3c3b47857174957eff02fb39151708c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findmeetingtimes"
    i11711209a5361d9de16ddab859bc848f471555bd3552c91461827db9b8353a96 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/restore"
    i123681f8c9035cd24ed6f5f2d1e4f1b18ae558a56793f108bf21272ae7de6a67 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    i127a34dfc4885325e12539ec28988683627ab6551e3174ea60385d697480ef8f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findrooms"
    i24c29a16bcb7e3171833cc95d27c27ca4c8fd7b9a5afe903d4f013fa9673845a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i288931c61b36c771b828474ad39a22adc5836b949ddd9c3264b01fa5ad364db2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmemberobjects"
    i2bcac46ff4c78d3e34c8c3de342b2d838834c6a19f418b3c87d66e0335bb79bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    i2beff65ac9b45d51f9f7f2f3f34acd22e1422c978e0a390f62d948870392b752 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmembergroups"
    i3791c51efa7487a4306aa44b71e4f7fdd5a24ef50a9fb17d92fe317110a850ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i41648ee7c80329421dd4930558c7f9df5e34a8265e6096c751e5e988a08b6605 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/activateserviceplan"
    i478b6e110d6d82827efa0a2d4a565e0bd59c058c7f54e34a79f3c3119618a8d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findroomlists"
    i524dc10161e92f6fb93296e5ddb276aa9a27e1012e1f050aa11cc29c1df7dcd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportpersonaldata"
    i53ce8ae9f81a119c5ddb15e3f2a12706f25e583ac0bd671bfcdf2ca713e1999f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmailtips"
    i549ef635dd7c772302ef56be1c9a1dff675e50a8f3cbaeb1fd948e83be80de55 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i60f32ef79c6ac8e7a2290563bfdc54bd0f33b2c432b8a39739dc4509f9e550d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/unblockmanagedapps"
    i7bbee660773e44d6361fa42938cd816ca03d849e08611e952595cb557845ec51 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i80e128991ede9a72a049be1b26e9b336dbd7b32d337d5b42b642124838a930fd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/translateexchangeids"
    i932a2c169109b2ef391689843544710eb8c5ab6aac44890c9b8bcdd947024023 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
    i9a510cffaf61d79946d18e93a5340163d4c302868630b5bcaca452e204823a7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    i9b0fc455dc045f2a20bce0e285ea9e35be41373d646a6009360a9b8e08836097 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmemberobjects"
    i9cbd13a06f16c60c6f6cb27851ab0596aff2e44e685dd847c5c0d05caa3c672c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
    ia7803215e8e6bf7c70ac1065fa6b4930720e87cfbb71698a8b8f58c3393a9394 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/sendmail"
    iabc322e88a8378d7f60b5d24fd385ac046eeb1be237cb7b5e896e28655ce7f5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/revokesigninsessions"
    iaf0e9b35edee5327d92f7b32c2c29bb0aa3fba7695625ed49aafb73fcacce4df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/checkmembergroups"
    ibe81ce1db33f361dc46fac394ebdaa2f4c6f7c9dee8dd858bdec3f3613a2472e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/findroomswithroomlist"
    ic48eee32109b6bd7e1b4c953770801b11ed9fb5e1d14b560f62d9e745d8a0a9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    ic7fe700e19e0b0c5969e2cf2077d84da7dd2e4feeb0a43a41b319f876c9b7443 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    ic8f1827b8744b17a04b290e375613de1dd341584b9c07bcfe3c2942645bbc3ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    ic91734e7a6bfc733a1bea394480741463152d7b0fe8dc49c6733b4b168f79d8a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    id935992fb11eece1b5abc1b559ca4e755d3ae870a8c57d716c1546d379940949 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/assignlicense"
    idd50169c749407aef5148d8bfb51c0dddd3fd9b1fc0b3eccdaf503dcb12d2293 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    ie3f806cfe1d2d944bd7c38b5348c74f4ddfbd5c717c98b71273aafef6d949c9b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    ieacce176e0811fa45db58facf7da2766c7cd4c373d3db6eb7ebc160eaa2dda99 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    iebbc29b2556ddd7c784f7f6030c57fe026a9caa9fa2e8e9c90d0b55c45cbf235 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    if68633e38c05d54f74f6f4f201a08689383470535eef536920823714d346abc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user/changepassword"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i41648ee7c80329421dd4930558c7f9df5e34a8265e6096c751e5e988a08b6605.ActivateServicePlanRequestBuilder) {
    return i41648ee7c80329421dd4930558c7f9df5e34a8265e6096c751e5e988a08b6605.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*id935992fb11eece1b5abc1b559ca4e755d3ae870a8c57d716c1546d379940949.AssignLicenseRequestBuilder) {
    return id935992fb11eece1b5abc1b559ca4e755d3ae870a8c57d716c1546d379940949.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*if68633e38c05d54f74f6f4f201a08689383470535eef536920823714d346abc6.ChangePasswordRequestBuilder) {
    return if68633e38c05d54f74f6f4f201a08689383470535eef536920823714d346abc6.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*iaf0e9b35edee5327d92f7b32c2c29bb0aa3fba7695625ed49aafb73fcacce4df.CheckMemberGroupsRequestBuilder) {
    return iaf0e9b35edee5327d92f7b32c2c29bb0aa3fba7695625ed49aafb73fcacce4df.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i9b0fc455dc045f2a20bce0e285ea9e35be41373d646a6009360a9b8e08836097.CheckMemberObjectsRequestBuilder) {
    return i9b0fc455dc045f2a20bce0e285ea9e35be41373d646a6009360a9b8e08836097.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i932a2c169109b2ef391689843544710eb8c5ab6aac44890c9b8bcdd947024023.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i932a2c169109b2ef391689843544710eb8c5ab6aac44890c9b8bcdd947024023.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*iebbc29b2556ddd7c784f7f6030c57fe026a9caa9fa2e8e9c90d0b55c45cbf235.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return iebbc29b2556ddd7c784f7f6030c57fe026a9caa9fa2e8e9c90d0b55c45cbf235.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i524dc10161e92f6fb93296e5ddb276aa9a27e1012e1f050aa11cc29c1df7dcd2.ExportPersonalDataRequestBuilder) {
    return i524dc10161e92f6fb93296e5ddb276aa9a27e1012e1f050aa11cc29c1df7dcd2.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i07ea9426dc558b5b0c329fa6629e6450a3c3b47857174957eff02fb39151708c.FindMeetingTimesRequestBuilder) {
    return i07ea9426dc558b5b0c329fa6629e6450a3c3b47857174957eff02fb39151708c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i478b6e110d6d82827efa0a2d4a565e0bd59c058c7f54e34a79f3c3119618a8d1.FindRoomListsRequestBuilder) {
    return i478b6e110d6d82827efa0a2d4a565e0bd59c058c7f54e34a79f3c3119618a8d1.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i127a34dfc4885325e12539ec28988683627ab6551e3174ea60385d697480ef8f.FindRoomsRequestBuilder) {
    return i127a34dfc4885325e12539ec28988683627ab6551e3174ea60385d697480ef8f.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ibe81ce1db33f361dc46fac394ebdaa2f4c6f7c9dee8dd858bdec3f3613a2472e.FindRoomsWithRoomListRequestBuilder) {
    return ibe81ce1db33f361dc46fac394ebdaa2f4c6f7c9dee8dd858bdec3f3613a2472e.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ieacce176e0811fa45db58facf7da2766c7cd4c373d3db6eb7ebc160eaa2dda99.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ieacce176e0811fa45db58facf7da2766c7cd4c373d3db6eb7ebc160eaa2dda99.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i9cbd13a06f16c60c6f6cb27851ab0596aff2e44e685dd847c5c0d05caa3c672c.GetLoggedOnManagedDevicesRequestBuilder) {
    return i9cbd13a06f16c60c6f6cb27851ab0596aff2e44e685dd847c5c0d05caa3c672c.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i53ce8ae9f81a119c5ddb15e3f2a12706f25e583ac0bd671bfcdf2ca713e1999f.GetMailTipsRequestBuilder) {
    return i53ce8ae9f81a119c5ddb15e3f2a12706f25e583ac0bd671bfcdf2ca713e1999f.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i549ef635dd7c772302ef56be1c9a1dff675e50a8f3cbaeb1fd948e83be80de55.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i549ef635dd7c772302ef56be1c9a1dff675e50a8f3cbaeb1fd948e83be80de55.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i123681f8c9035cd24ed6f5f2d1e4f1b18ae558a56793f108bf21272ae7de6a67.GetManagedAppPoliciesRequestBuilder) {
    return i123681f8c9035cd24ed6f5f2d1e4f1b18ae558a56793f108bf21272ae7de6a67.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ie3f806cfe1d2d944bd7c38b5348c74f4ddfbd5c717c98b71273aafef6d949c9b.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ie3f806cfe1d2d944bd7c38b5348c74f4ddfbd5c717c98b71273aafef6d949c9b.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i2bcac46ff4c78d3e34c8c3de342b2d838834c6a19f418b3c87d66e0335bb79bd.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i2bcac46ff4c78d3e34c8c3de342b2d838834c6a19f418b3c87d66e0335bb79bd.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i2beff65ac9b45d51f9f7f2f3f34acd22e1422c978e0a390f62d948870392b752.GetMemberGroupsRequestBuilder) {
    return i2beff65ac9b45d51f9f7f2f3f34acd22e1422c978e0a390f62d948870392b752.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i288931c61b36c771b828474ad39a22adc5836b949ddd9c3264b01fa5ad364db2.GetMemberObjectsRequestBuilder) {
    return i288931c61b36c771b828474ad39a22adc5836b949ddd9c3264b01fa5ad364db2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ic48eee32109b6bd7e1b4c953770801b11ed9fb5e1d14b560f62d9e745d8a0a9e.InvalidateAllRefreshTokensRequestBuilder) {
    return ic48eee32109b6bd7e1b4c953770801b11ed9fb5e1d14b560f62d9e745d8a0a9e.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i9a510cffaf61d79946d18e93a5340163d4c302868630b5bcaca452e204823a7c.IsManagedAppUserBlockedRequestBuilder) {
    return i9a510cffaf61d79946d18e93a5340163d4c302868630b5bcaca452e204823a7c.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ic8f1827b8744b17a04b290e375613de1dd341584b9c07bcfe3c2942645bbc3ad.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ic8f1827b8744b17a04b290e375613de1dd341584b9c07bcfe3c2942645bbc3ad.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i3791c51efa7487a4306aa44b71e4f7fdd5a24ef50a9fb17d92fe317110a850ad.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3791c51efa7487a4306aa44b71e4f7fdd5a24ef50a9fb17d92fe317110a850ad.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ic91734e7a6bfc733a1bea394480741463152d7b0fe8dc49c6733b4b168f79d8a.ReprocessLicenseAssignmentRequestBuilder) {
    return ic91734e7a6bfc733a1bea394480741463152d7b0fe8dc49c6733b4b168f79d8a.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i11711209a5361d9de16ddab859bc848f471555bd3552c91461827db9b8353a96.RestoreRequestBuilder) {
    return i11711209a5361d9de16ddab859bc848f471555bd3552c91461827db9b8353a96.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*iabc322e88a8378d7f60b5d24fd385ac046eeb1be237cb7b5e896e28655ce7f5a.RevokeSignInSessionsRequestBuilder) {
    return iabc322e88a8378d7f60b5d24fd385ac046eeb1be237cb7b5e896e28655ce7f5a.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ia7803215e8e6bf7c70ac1065fa6b4930720e87cfbb71698a8b8f58c3393a9394.SendMailRequestBuilder) {
    return ia7803215e8e6bf7c70ac1065fa6b4930720e87cfbb71698a8b8f58c3393a9394.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i80e128991ede9a72a049be1b26e9b336dbd7b32d337d5b42b642124838a930fd.TranslateExchangeIdsRequestBuilder) {
    return i80e128991ede9a72a049be1b26e9b336dbd7b32d337d5b42b642124838a930fd.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i60f32ef79c6ac8e7a2290563bfdc54bd0f33b2c432b8a39739dc4509f9e550d1.UnblockManagedAppsRequestBuilder) {
    return i60f32ef79c6ac8e7a2290563bfdc54bd0f33b2c432b8a39739dc4509f9e550d1.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ic7fe700e19e0b0c5969e2cf2077d84da7dd2e4feeb0a43a41b319f876c9b7443.WipeAndBlockManagedAppsRequestBuilder) {
    return ic7fe700e19e0b0c5969e2cf2077d84da7dd2e4feeb0a43a41b319f876c9b7443.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*idd50169c749407aef5148d8bfb51c0dddd3fd9b1fc0b3eccdaf503dcb12d2293.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return idd50169c749407aef5148d8bfb51c0dddd3fd9b1fc0b3eccdaf503dcb12d2293.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i7bbee660773e44d6361fa42938cd816ca03d849e08611e952595cb557845ec51.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i7bbee660773e44d6361fa42938cd816ca03d849e08611e952595cb557845ec51.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i24c29a16bcb7e3171833cc95d27c27ca4c8fd7b9a5afe903d4f013fa9673845a.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i24c29a16bcb7e3171833cc95d27c27ca4c8fd7b9a5afe903d4f013fa9673845a.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
