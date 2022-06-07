package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c05faef1dd4233910f6d8963b1c87d9d5bc2e52e2102631ffcb2f30f69dc5f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i115b3cbb1c133b2905ac3410179cb6e3d592bca8868a6a009565a9a1e575967e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/exportpersonaldata"
    i121095c0ae7af2b257786a00bf719c5146e767e547248d118badf50b17a82263 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/findrooms"
    i148ef37c2d0e876fb5dd6a2c748d83ca76a1a0ba4e70af4fd9fc20a740010b71 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/revokesigninsessions"
    i2fa8bfbe7d84162c83fe387d7c50e4b47db2b416b8d3400637083174d84f0ac8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/removealldevicesfrommanagement"
    i31bb1bb7152536e6aec8e96b2519602b80390e0bd6185c3acc986a5d8ec09161 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/sendmail"
    i34373e2c9a9bc36d0811d8f35baa3d01e192bbe889b035accdf4f4d073f769de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/checkmembergroups"
    i47b2a0df1b04e002585c5f4fab0e13940898a73f47604dc3f00f2843e0f74486 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i49e1483648112ab6cab2cce4228001101d4beab8738a1551dd5c799d53beeca8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    i4a67d35d9048bb17ae36a896a24699f8f3a9b4755585ac70edab5b0d71ee90e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmembergroups"
    i4ea840e2fb5df74cc75f05653fb71552f7617280a2b59a7f037086915aefec60 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/ismanagedappuserblocked"
    i66b43df4eee780a9a759f227e8fa5b61123207f11046d7a78b1ad3eafe054020 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/changepassword"
    i66c58b4c98a828eca6b37bd4c1822b1507cc924ecd3653079556c73d0d4f73b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmanageddeviceswithappfailures"
    i691aa52abff5b72b62be62c2b28c96beb8a2428371ddfb6218523fb9c68a56ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmailtips"
    i6dbf615c343680f11337beb42962fd8869bedee9381a86078ded772ea5021c50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/translateexchangeids"
    i7577cf8dd1775506683403908c0cc5083aef74ca2242671ab86e6ef9be8e7fe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getloggedonmanageddevices"
    i7f468f3291077a704820d55da1fee8cb2618d666e7efdf37538788b78548bb02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/activateserviceplan"
    i853626051f56d0bfd98b5ab371ca4ea6ac6a265bc17f8498f94985d78455baa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/restore"
    i8bb7e482a60cb3acca9d6835e9f6d4408d4b1ed52e431d0c22b39bef20bcb2b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/assignlicense"
    i90940345ddfdc15bcb24439ea6588afdd94cf361557f1c3d4cb19ffee3b7c88f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/findmeetingtimes"
    i9125904334965454f9202bd93c009bc9d5cf0756a2cc86bb7fbf24b41c92229d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    i956916fc39babd19cf75375ebdca02ef4e77ace8a4cf2ad0daa8e90c2aa832b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    ia13bf21dec81bfaf1866ad9645fc6aea3e394e9eeb396c6db811ecb88c6f081b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    ib007088026dd7b66b2dfa0564af7aae82372aefbbf4cee4d37df52fe22989d74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/unblockmanagedapps"
    ib197ea9ed0e96bf412a9a0195ba584cc168568883fab667053ea72d4f6cd3655 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    ib87aed14066b02c4844946d7f89f4b257194a4c1ce158fe6df3ca45b2fd9e401 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/invalidateallrefreshtokens"
    ic1570f22160326aa810acb784b77fecf27172482c5310abac2ad12317cc48bb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmemberobjects"
    ic9004ec5efc9719d96bc84d72ebdac450f83f3a3d6df0a04a255dbc1b4cfd029 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/getmanagedapppolicies"
    icf94a59958808a19c4c4e53bd6b1a59095904217604e5d8674c7232d5de5b73e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    id38ee4b1079fec932a30ac5f6cc3468c4e8a8adad22882f957bbe056306f8303 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/wipeandblockmanagedapps"
    ie0f2a5331399f74a43dc1151235a9082ba5d9f67ffd853f7bde771308f20bbab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/checkmemberobjects"
    ie47714071448ee4201f068e75a1f6eeaeb97bb1add4d6b945c50f0068deb0616 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/findroomswithroomlist"
    ief08280cad777555fb2d4fad3091eaddd715afe41a4a30d4a6186d9d9c140836 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/reprocesslicenseassignment"
    if028b3785e42fe83de1418b0d86c96376a5a3c6c2d9970745c1490c3c05cd512 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/exportdeviceandappmanagementdata"
    ifd121364c6094f3023609be59ea0ac82002b2455b3acd7ea3f7c896d5d58f691 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredusers/item/user/findroomlists"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i7f468f3291077a704820d55da1fee8cb2618d666e7efdf37538788b78548bb02.ActivateServicePlanRequestBuilder) {
    return i7f468f3291077a704820d55da1fee8cb2618d666e7efdf37538788b78548bb02.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i8bb7e482a60cb3acca9d6835e9f6d4408d4b1ed52e431d0c22b39bef20bcb2b2.AssignLicenseRequestBuilder) {
    return i8bb7e482a60cb3acca9d6835e9f6d4408d4b1ed52e431d0c22b39bef20bcb2b2.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i66b43df4eee780a9a759f227e8fa5b61123207f11046d7a78b1ad3eafe054020.ChangePasswordRequestBuilder) {
    return i66b43df4eee780a9a759f227e8fa5b61123207f11046d7a78b1ad3eafe054020.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i34373e2c9a9bc36d0811d8f35baa3d01e192bbe889b035accdf4f4d073f769de.CheckMemberGroupsRequestBuilder) {
    return i34373e2c9a9bc36d0811d8f35baa3d01e192bbe889b035accdf4f4d073f769de.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ie0f2a5331399f74a43dc1151235a9082ba5d9f67ffd853f7bde771308f20bbab.CheckMemberObjectsRequestBuilder) {
    return ie0f2a5331399f74a43dc1151235a9082ba5d9f67ffd853f7bde771308f20bbab.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*if028b3785e42fe83de1418b0d86c96376a5a3c6c2d9970745c1490c3c05cd512.ExportDeviceAndAppManagementDataRequestBuilder) {
    return if028b3785e42fe83de1418b0d86c96376a5a3c6c2d9970745c1490c3c05cd512.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i47b2a0df1b04e002585c5f4fab0e13940898a73f47604dc3f00f2843e0f74486.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i47b2a0df1b04e002585c5f4fab0e13940898a73f47604dc3f00f2843e0f74486.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i115b3cbb1c133b2905ac3410179cb6e3d592bca8868a6a009565a9a1e575967e.ExportPersonalDataRequestBuilder) {
    return i115b3cbb1c133b2905ac3410179cb6e3d592bca8868a6a009565a9a1e575967e.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i90940345ddfdc15bcb24439ea6588afdd94cf361557f1c3d4cb19ffee3b7c88f.FindMeetingTimesRequestBuilder) {
    return i90940345ddfdc15bcb24439ea6588afdd94cf361557f1c3d4cb19ffee3b7c88f.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ifd121364c6094f3023609be59ea0ac82002b2455b3acd7ea3f7c896d5d58f691.FindRoomListsRequestBuilder) {
    return ifd121364c6094f3023609be59ea0ac82002b2455b3acd7ea3f7c896d5d58f691.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i121095c0ae7af2b257786a00bf719c5146e767e547248d118badf50b17a82263.FindRoomsRequestBuilder) {
    return i121095c0ae7af2b257786a00bf719c5146e767e547248d118badf50b17a82263.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ie47714071448ee4201f068e75a1f6eeaeb97bb1add4d6b945c50f0068deb0616.FindRoomsWithRoomListRequestBuilder) {
    return ie47714071448ee4201f068e75a1f6eeaeb97bb1add4d6b945c50f0068deb0616.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i9125904334965454f9202bd93c009bc9d5cf0756a2cc86bb7fbf24b41c92229d.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i9125904334965454f9202bd93c009bc9d5cf0756a2cc86bb7fbf24b41c92229d.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i7577cf8dd1775506683403908c0cc5083aef74ca2242671ab86e6ef9be8e7fe7.GetLoggedOnManagedDevicesRequestBuilder) {
    return i7577cf8dd1775506683403908c0cc5083aef74ca2242671ab86e6ef9be8e7fe7.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i691aa52abff5b72b62be62c2b28c96beb8a2428371ddfb6218523fb9c68a56ad.GetMailTipsRequestBuilder) {
    return i691aa52abff5b72b62be62c2b28c96beb8a2428371ddfb6218523fb9c68a56ad.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i0c05faef1dd4233910f6d8963b1c87d9d5bc2e52e2102631ffcb2f30f69dc5f0.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i0c05faef1dd4233910f6d8963b1c87d9d5bc2e52e2102631ffcb2f30f69dc5f0.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic9004ec5efc9719d96bc84d72ebdac450f83f3a3d6df0a04a255dbc1b4cfd029.GetManagedAppPoliciesRequestBuilder) {
    return ic9004ec5efc9719d96bc84d72ebdac450f83f3a3d6df0a04a255dbc1b4cfd029.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i66c58b4c98a828eca6b37bd4c1822b1507cc924ecd3653079556c73d0d4f73b9.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i66c58b4c98a828eca6b37bd4c1822b1507cc924ecd3653079556c73d0d4f73b9.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ib197ea9ed0e96bf412a9a0195ba584cc168568883fab667053ea72d4f6cd3655.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ib197ea9ed0e96bf412a9a0195ba584cc168568883fab667053ea72d4f6cd3655.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i4a67d35d9048bb17ae36a896a24699f8f3a9b4755585ac70edab5b0d71ee90e0.GetMemberGroupsRequestBuilder) {
    return i4a67d35d9048bb17ae36a896a24699f8f3a9b4755585ac70edab5b0d71ee90e0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ic1570f22160326aa810acb784b77fecf27172482c5310abac2ad12317cc48bb7.GetMemberObjectsRequestBuilder) {
    return ic1570f22160326aa810acb784b77fecf27172482c5310abac2ad12317cc48bb7.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ib87aed14066b02c4844946d7f89f4b257194a4c1ce158fe6df3ca45b2fd9e401.InvalidateAllRefreshTokensRequestBuilder) {
    return ib87aed14066b02c4844946d7f89f4b257194a4c1ce158fe6df3ca45b2fd9e401.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i4ea840e2fb5df74cc75f05653fb71552f7617280a2b59a7f037086915aefec60.IsManagedAppUserBlockedRequestBuilder) {
    return i4ea840e2fb5df74cc75f05653fb71552f7617280a2b59a7f037086915aefec60.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i956916fc39babd19cf75375ebdca02ef4e77ace8a4cf2ad0daa8e90c2aa832b1.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i956916fc39babd19cf75375ebdca02ef4e77ace8a4cf2ad0daa8e90c2aa832b1.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i2fa8bfbe7d84162c83fe387d7c50e4b47db2b416b8d3400637083174d84f0ac8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i2fa8bfbe7d84162c83fe387d7c50e4b47db2b416b8d3400637083174d84f0ac8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ief08280cad777555fb2d4fad3091eaddd715afe41a4a30d4a6186d9d9c140836.ReprocessLicenseAssignmentRequestBuilder) {
    return ief08280cad777555fb2d4fad3091eaddd715afe41a4a30d4a6186d9d9c140836.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i853626051f56d0bfd98b5ab371ca4ea6ac6a265bc17f8498f94985d78455baa6.RestoreRequestBuilder) {
    return i853626051f56d0bfd98b5ab371ca4ea6ac6a265bc17f8498f94985d78455baa6.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i148ef37c2d0e876fb5dd6a2c748d83ca76a1a0ba4e70af4fd9fc20a740010b71.RevokeSignInSessionsRequestBuilder) {
    return i148ef37c2d0e876fb5dd6a2c748d83ca76a1a0ba4e70af4fd9fc20a740010b71.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i31bb1bb7152536e6aec8e96b2519602b80390e0bd6185c3acc986a5d8ec09161.SendMailRequestBuilder) {
    return i31bb1bb7152536e6aec8e96b2519602b80390e0bd6185c3acc986a5d8ec09161.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i6dbf615c343680f11337beb42962fd8869bedee9381a86078ded772ea5021c50.TranslateExchangeIdsRequestBuilder) {
    return i6dbf615c343680f11337beb42962fd8869bedee9381a86078ded772ea5021c50.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ib007088026dd7b66b2dfa0564af7aae82372aefbbf4cee4d37df52fe22989d74.UnblockManagedAppsRequestBuilder) {
    return ib007088026dd7b66b2dfa0564af7aae82372aefbbf4cee4d37df52fe22989d74.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*id38ee4b1079fec932a30ac5f6cc3468c4e8a8adad22882f957bbe056306f8303.WipeAndBlockManagedAppsRequestBuilder) {
    return id38ee4b1079fec932a30ac5f6cc3468c4e8a8adad22882f957bbe056306f8303.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i49e1483648112ab6cab2cce4228001101d4beab8738a1551dd5c799d53beeca8.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i49e1483648112ab6cab2cce4228001101d4beab8738a1551dd5c799d53beeca8.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*icf94a59958808a19c4c4e53bd6b1a59095904217604e5d8674c7232d5de5b73e.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return icf94a59958808a19c4c4e53bd6b1a59095904217604e5d8674c7232d5de5b73e.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ia13bf21dec81bfaf1866ad9645fc6aea3e394e9eeb396c6db811ecb88c6f081b.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ia13bf21dec81bfaf1866ad9645fc6aea3e394e9eeb396c6db811ecb88c6f081b.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
