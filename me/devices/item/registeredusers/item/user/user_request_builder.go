package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d15d4eeb8035a6fed1343278372417aa984f1b07eb2c18710ce414a539725ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/reprocesslicenseassignment"
    i102b73a33e0cae9ec23033261b45ae56db944d3584ef9486b14a6635456d8037 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/checkmemberobjects"
    i1289fb9460edf89011fae3ab5791f65508fdbcce2e332697d7560bfb7efd4c13 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/findroomlists"
    i200bfbc68f91c35a109165bae7cf48e68183f8926955fc533b6065f707fb21cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/wipeandblockmanagedapps"
    i24a5796cf519da421723a4049812fd0db3ec53b85161baf0cada728beb70a5eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/changepassword"
    i2c9507cde89b959cfbd4c2ff8a8698cc5bfce28c6878d18da8451c7d5bd05c68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    i31f5b134fbd358ff2d6bef1b31f4b823c5c545bed790c9502150977fc2b8e5db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/exportpersonaldata"
    i3c35a2f86978cc8f1cf1937790792e9ba238c3a78538343dd194f0d109de16d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i3d498aed597b17291eea0df31a25904828cfc8e6bf318bfbf3faf8b9eb63cd03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/removealldevicesfrommanagement"
    i463b6ff5cc91a9511db7af15b4f3f36dd294231a61a92877f5dc5b32d8644398 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/findroomswithroomlist"
    i475de12c8f9fe3bb7b05a8a85d96a20c3e7915b62bc16ea6e092b56ed92251b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    i51a59cadff85fe051c001544d2f12379189e7eb51b2d8a953df277537af933d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/findrooms"
    i529599dd7fb1ef3c47000b8768003764a225b3618c3c36e6e08271e29e3ea550 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/findmeetingtimes"
    i53112a6bd2863e6209b35dd52c152c579f9525078b43ab978e47d44220e23a64 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/ismanagedappuserblocked"
    i544bd2aa4e974a5614024a968d0aa76b45fe37a487bfcfdd40aa47a12929c19a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmanagedapppolicies"
    i5d9bd5cc4b2a4891bb13ede6d3e08bb55737070049247a6a047c1860eadd61bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmembergroups"
    i63b2d9c28a036e124b7a7982a8217ebda1a9926634456d4d93ae6769563750f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/sendmail"
    i6da17f78947260f1e4b1a544ebe54b4a543e90921701e813844e4cf22aecf7a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getloggedonmanageddevices"
    i792ef6aa52d830c732b09e6feeb2ce238e1131419ef5c6c3b4def78ec205a983 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/revokesigninsessions"
    i7d7418eb7babd4c4163973a9758eedd2ce188d90c4b3ad69d7bc9e097eb94e9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    i87378262fa564dbfb48fc862c0ae57eb43fae1f44a20a11fb63ae163d8be4c88 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/translateexchangeids"
    i91e4def354eccc4d7fc9b70ec3d658dd0aa2461172b9b69bd8208ed7cfe0b7f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/assignlicense"
    i9c5a14a0e630b5a0802a63d599e925c8bea12c732edd23713469779b09919fce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    ib8d201e6479f445736747109279104fce9476482d9313307330369c221a4b68e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/activateserviceplan"
    iba049b07d017bb83baa50281d6bd63ce7fd57a823d88df991ef93c57fc97443f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    ic1beffef85a29ef2dac1f0f81c4ba35d6f960e7ce67b90bfa73dc7b7fe642ca6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/checkmembergroups"
    icb6cd645fe58e2a307fd0591be35ffcba4738c9de74750330b5793d5af8c8cd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/unblockmanagedapps"
    id1e16de4070eb79b0f400bff640b6d153f6f09e48d0835001d5cc4d052b87a3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/invalidateallrefreshtokens"
    id49a71c7c06f856927d8cf96464dfe23eac1b57f746c3f79eb16f6b596ca6967 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    id997bc7718a69aed6f9f1aaee99a99225323a27a36c4191869f9ac9f012e3ee5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/exportdeviceandappmanagementdata"
    ieb889d889a7d5621107733fa69543ff421143dd7906a8716c97a12ec65380b23 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/restore"
    iedff0f87d035dc019410ee601b4c801b473da9477f0cd05177800bb30f9164d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    if93ec5caaa2962fd32a42c531278cf46ea09ccfa16a192b21718e66dcd017fff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmailtips"
    ifb4346d50cdefa6b9f80f097cb5043926441153f714cda050c234929af347db8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmanageddeviceswithappfailures"
    ifbb871ef37e594d2581a28169444c588bf42bfcbb8346f947add6c3f87246843 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredusers/item/user/getmemberobjects"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ib8d201e6479f445736747109279104fce9476482d9313307330369c221a4b68e.ActivateServicePlanRequestBuilder) {
    return ib8d201e6479f445736747109279104fce9476482d9313307330369c221a4b68e.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i91e4def354eccc4d7fc9b70ec3d658dd0aa2461172b9b69bd8208ed7cfe0b7f9.AssignLicenseRequestBuilder) {
    return i91e4def354eccc4d7fc9b70ec3d658dd0aa2461172b9b69bd8208ed7cfe0b7f9.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i24a5796cf519da421723a4049812fd0db3ec53b85161baf0cada728beb70a5eb.ChangePasswordRequestBuilder) {
    return i24a5796cf519da421723a4049812fd0db3ec53b85161baf0cada728beb70a5eb.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ic1beffef85a29ef2dac1f0f81c4ba35d6f960e7ce67b90bfa73dc7b7fe642ca6.CheckMemberGroupsRequestBuilder) {
    return ic1beffef85a29ef2dac1f0f81c4ba35d6f960e7ce67b90bfa73dc7b7fe642ca6.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i102b73a33e0cae9ec23033261b45ae56db944d3584ef9486b14a6635456d8037.CheckMemberObjectsRequestBuilder) {
    return i102b73a33e0cae9ec23033261b45ae56db944d3584ef9486b14a6635456d8037.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*id997bc7718a69aed6f9f1aaee99a99225323a27a36c4191869f9ac9f012e3ee5.ExportDeviceAndAppManagementDataRequestBuilder) {
    return id997bc7718a69aed6f9f1aaee99a99225323a27a36c4191869f9ac9f012e3ee5.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i3c35a2f86978cc8f1cf1937790792e9ba238c3a78538343dd194f0d109de16d3.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i3c35a2f86978cc8f1cf1937790792e9ba238c3a78538343dd194f0d109de16d3.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i31f5b134fbd358ff2d6bef1b31f4b823c5c545bed790c9502150977fc2b8e5db.ExportPersonalDataRequestBuilder) {
    return i31f5b134fbd358ff2d6bef1b31f4b823c5c545bed790c9502150977fc2b8e5db.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i529599dd7fb1ef3c47000b8768003764a225b3618c3c36e6e08271e29e3ea550.FindMeetingTimesRequestBuilder) {
    return i529599dd7fb1ef3c47000b8768003764a225b3618c3c36e6e08271e29e3ea550.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i1289fb9460edf89011fae3ab5791f65508fdbcce2e332697d7560bfb7efd4c13.FindRoomListsRequestBuilder) {
    return i1289fb9460edf89011fae3ab5791f65508fdbcce2e332697d7560bfb7efd4c13.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i51a59cadff85fe051c001544d2f12379189e7eb51b2d8a953df277537af933d6.FindRoomsRequestBuilder) {
    return i51a59cadff85fe051c001544d2f12379189e7eb51b2d8a953df277537af933d6.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i463b6ff5cc91a9511db7af15b4f3f36dd294231a61a92877f5dc5b32d8644398.FindRoomsWithRoomListRequestBuilder) {
    return i463b6ff5cc91a9511db7af15b4f3f36dd294231a61a92877f5dc5b32d8644398.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i475de12c8f9fe3bb7b05a8a85d96a20c3e7915b62bc16ea6e092b56ed92251b1.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i475de12c8f9fe3bb7b05a8a85d96a20c3e7915b62bc16ea6e092b56ed92251b1.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i6da17f78947260f1e4b1a544ebe54b4a543e90921701e813844e4cf22aecf7a1.GetLoggedOnManagedDevicesRequestBuilder) {
    return i6da17f78947260f1e4b1a544ebe54b4a543e90921701e813844e4cf22aecf7a1.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*if93ec5caaa2962fd32a42c531278cf46ea09ccfa16a192b21718e66dcd017fff.GetMailTipsRequestBuilder) {
    return if93ec5caaa2962fd32a42c531278cf46ea09ccfa16a192b21718e66dcd017fff.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id49a71c7c06f856927d8cf96464dfe23eac1b57f746c3f79eb16f6b596ca6967.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id49a71c7c06f856927d8cf96464dfe23eac1b57f746c3f79eb16f6b596ca6967.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i544bd2aa4e974a5614024a968d0aa76b45fe37a487bfcfdd40aa47a12929c19a.GetManagedAppPoliciesRequestBuilder) {
    return i544bd2aa4e974a5614024a968d0aa76b45fe37a487bfcfdd40aa47a12929c19a.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ifb4346d50cdefa6b9f80f097cb5043926441153f714cda050c234929af347db8.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ifb4346d50cdefa6b9f80f097cb5043926441153f714cda050c234929af347db8.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7d7418eb7babd4c4163973a9758eedd2ce188d90c4b3ad69d7bc9e097eb94e9a.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7d7418eb7babd4c4163973a9758eedd2ce188d90c4b3ad69d7bc9e097eb94e9a.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i5d9bd5cc4b2a4891bb13ede6d3e08bb55737070049247a6a047c1860eadd61bc.GetMemberGroupsRequestBuilder) {
    return i5d9bd5cc4b2a4891bb13ede6d3e08bb55737070049247a6a047c1860eadd61bc.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ifbb871ef37e594d2581a28169444c588bf42bfcbb8346f947add6c3f87246843.GetMemberObjectsRequestBuilder) {
    return ifbb871ef37e594d2581a28169444c588bf42bfcbb8346f947add6c3f87246843.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*id1e16de4070eb79b0f400bff640b6d153f6f09e48d0835001d5cc4d052b87a3a.InvalidateAllRefreshTokensRequestBuilder) {
    return id1e16de4070eb79b0f400bff640b6d153f6f09e48d0835001d5cc4d052b87a3a.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i53112a6bd2863e6209b35dd52c152c579f9525078b43ab978e47d44220e23a64.IsManagedAppUserBlockedRequestBuilder) {
    return i53112a6bd2863e6209b35dd52c152c579f9525078b43ab978e47d44220e23a64.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*iba049b07d017bb83baa50281d6bd63ce7fd57a823d88df991ef93c57fc97443f.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return iba049b07d017bb83baa50281d6bd63ce7fd57a823d88df991ef93c57fc97443f.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i3d498aed597b17291eea0df31a25904828cfc8e6bf318bfbf3faf8b9eb63cd03.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3d498aed597b17291eea0df31a25904828cfc8e6bf318bfbf3faf8b9eb63cd03.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0d15d4eeb8035a6fed1343278372417aa984f1b07eb2c18710ce414a539725ba.ReprocessLicenseAssignmentRequestBuilder) {
    return i0d15d4eeb8035a6fed1343278372417aa984f1b07eb2c18710ce414a539725ba.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ieb889d889a7d5621107733fa69543ff421143dd7906a8716c97a12ec65380b23.RestoreRequestBuilder) {
    return ieb889d889a7d5621107733fa69543ff421143dd7906a8716c97a12ec65380b23.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i792ef6aa52d830c732b09e6feeb2ce238e1131419ef5c6c3b4def78ec205a983.RevokeSignInSessionsRequestBuilder) {
    return i792ef6aa52d830c732b09e6feeb2ce238e1131419ef5c6c3b4def78ec205a983.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i63b2d9c28a036e124b7a7982a8217ebda1a9926634456d4d93ae6769563750f4.SendMailRequestBuilder) {
    return i63b2d9c28a036e124b7a7982a8217ebda1a9926634456d4d93ae6769563750f4.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i87378262fa564dbfb48fc862c0ae57eb43fae1f44a20a11fb63ae163d8be4c88.TranslateExchangeIdsRequestBuilder) {
    return i87378262fa564dbfb48fc862c0ae57eb43fae1f44a20a11fb63ae163d8be4c88.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*icb6cd645fe58e2a307fd0591be35ffcba4738c9de74750330b5793d5af8c8cd5.UnblockManagedAppsRequestBuilder) {
    return icb6cd645fe58e2a307fd0591be35ffcba4738c9de74750330b5793d5af8c8cd5.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i200bfbc68f91c35a109165bae7cf48e68183f8926955fc533b6065f707fb21cb.WipeAndBlockManagedAppsRequestBuilder) {
    return i200bfbc68f91c35a109165bae7cf48e68183f8926955fc533b6065f707fb21cb.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i9c5a14a0e630b5a0802a63d599e925c8bea12c732edd23713469779b09919fce.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i9c5a14a0e630b5a0802a63d599e925c8bea12c732edd23713469779b09919fce.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*iedff0f87d035dc019410ee601b4c801b473da9477f0cd05177800bb30f9164d5.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return iedff0f87d035dc019410ee601b4c801b473da9477f0cd05177800bb30f9164d5.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i2c9507cde89b959cfbd4c2ff8a8698cc5bfce28c6878d18da8451c7d5bd05c68.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i2c9507cde89b959cfbd4c2ff8a8698cc5bfce28c6878d18da8451c7d5bd05c68.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
