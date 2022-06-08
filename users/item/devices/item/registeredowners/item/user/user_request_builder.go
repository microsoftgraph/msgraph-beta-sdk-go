package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a60f9ea2b6cce09831def3f509d4df59a4cab86059636d57af6afa9dceb712a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/invalidateallrefreshtokens"
    i117fdf5b2eac599662922eec9ad4be84b3227e500937ce8b54045f42f2f9b864 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmanagedapppolicies"
    i11afdfb304ea23402ce262feb0b1388ba316e975b68144a7e08af68b0d727246 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/translateexchangeids"
    i15cfd0907ef334819d192f0caae869887353a08aebd4f57a6d50609f90a38324 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmanageddeviceswithappfailures"
    i1e008e60a8eb48a26abdd8c95fe5c8356221274d610db379067882383cf00892 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i21b9f7afd263bff46da5eec64a5779ec05c7b8710c85a24989857b6d5fa907f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/exportdeviceandappmanagementdata"
    i253f8948a060f08ed80843c9f8ec8ffb9bc98a9dc85d1e77a4592bd3956e8fb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/wipeandblockmanagedapps"
    i37dda0334f2b355a2d6059b26d596a2995283976d7cc804d0949689e2b46ac3d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/revokesigninsessions"
    i3fd31d011299e4bd74e1cf12016fa1f926e6777b1d2d27d7fc28b9a3ee3b54ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/unblockmanagedapps"
    i410698a234b29aa9a151a24f43e41f20ed64b7619f28e2b57728e83398070822 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/checkmemberobjects"
    i51662d8d0f10f6679fedad6f510bf8ffca49639e415ab8ba60bbb2074818cfe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/changepassword"
    i51f2ccc7038aeea321321bd75f149854eb7892cd5c871c3e4c5354f018c73c48 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/findroomswithroomlist"
    i5c64c484fc87a05f75b11cb79b985d50f6a0ef5b796a0f8d6b38608d07c1dc80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/findmeetingtimes"
    i6f660b2006606b84de3153fd9d5b29547e9b86b0b7996ac3d5d9506c1d52e5b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/removealldevicesfrommanagement"
    i7201fc60f93506ab507aba55127067ad853520cd3b1bc02548cbb94c60ed14f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/exportpersonaldata"
    i77a8322a69419d2f12b42c265833cce281302b9c7cc7a22fdf7cb7674873fb95 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/ismanagedappuserblocked"
    i7d20cb7abcc98a52e8d49ef40215e6750c95a794dbd9e2e81be74631976880fa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    i7e6d1819cc25e13a93a3bbf72cc883fe05449ba38ff9aa1d59796a5b7a4e705e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/reprocesslicenseassignment"
    i859fa10100610ebcdde7ac39005732ad2c78c0956d868a2cdd060dbde71d0785 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/findroomlists"
    i98fd3840004bbe47a878dcf9635e0425c97446b349b35f5ab65013f5caa62d17 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmailtips"
    ia5c7e3e68cc560feb74cab0573eb2b92ee3caffee9bcd099ad9aaaaf1a8f2ece "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/findrooms"
    ia9c8860c2353c5adfde2cab02083d7ed0adb01098e53be4f21b20805dea4a77a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/activateserviceplan"
    iaaaf1941f01165dfcd733293a08b35738af3c0c57f52173888cdb0bc99e39602 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    ib623c31512b6e82ea374585c36ebe57ae562fcdb244e986b577bb24d9d76f291 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmemberobjects"
    ib878de120d70e19c2d02b15b53aab41a96f49794e2af6e5ce6a6ce7adc417b48 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/checkmembergroups"
    iba8d9c4db75c7d271f5066be861a9c614fa33ca32ed6fb0bc21cc3736723589f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    ibab6f315a80eee5cc5adc95b626ac41247d08ff54a8b33a80dcda11ab0f8d5dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/assignlicense"
    ic74dbb2c637e6585f461876e21663edbe94a58d6db8017de79cfa610d5fb69bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/restore"
    id375a73f8562f1afdea4c927701c0f4e21fb27a3727222c79c87a4f49a79c621 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    ie8b8fb27d7254e798aacb21b8417e8b0e30a7a923b637e03d1ce41c6e75a1448 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getloggedonmanageddevices"
    ief42293257452c74833dc67d9096853dd9d4f7d9d83d9a8a76da90904c35f220 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/getmembergroups"
    if0bdacbb8a78c8dbee95e09187dbd7133f8f3ff7b44b66075ce503671ee96ed1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/sendmail"
    ifbbbe5290364d20daa098ac5c2c4b3f2a0d67e1882746b298d0c4f90fb8dec53 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    iff75f05397b208c225848e8544853f552816e44d41e74edaa722798a39069705 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ifff60123811a96a13ebaf593c1fc70fa0d34d8408b9a1c03b6a1d8b16772eaf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ia9c8860c2353c5adfde2cab02083d7ed0adb01098e53be4f21b20805dea4a77a.ActivateServicePlanRequestBuilder) {
    return ia9c8860c2353c5adfde2cab02083d7ed0adb01098e53be4f21b20805dea4a77a.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ibab6f315a80eee5cc5adc95b626ac41247d08ff54a8b33a80dcda11ab0f8d5dd.AssignLicenseRequestBuilder) {
    return ibab6f315a80eee5cc5adc95b626ac41247d08ff54a8b33a80dcda11ab0f8d5dd.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i51662d8d0f10f6679fedad6f510bf8ffca49639e415ab8ba60bbb2074818cfe4.ChangePasswordRequestBuilder) {
    return i51662d8d0f10f6679fedad6f510bf8ffca49639e415ab8ba60bbb2074818cfe4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ib878de120d70e19c2d02b15b53aab41a96f49794e2af6e5ce6a6ce7adc417b48.CheckMemberGroupsRequestBuilder) {
    return ib878de120d70e19c2d02b15b53aab41a96f49794e2af6e5ce6a6ce7adc417b48.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i410698a234b29aa9a151a24f43e41f20ed64b7619f28e2b57728e83398070822.CheckMemberObjectsRequestBuilder) {
    return i410698a234b29aa9a151a24f43e41f20ed64b7619f28e2b57728e83398070822.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i21b9f7afd263bff46da5eec64a5779ec05c7b8710c85a24989857b6d5fa907f2.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i21b9f7afd263bff46da5eec64a5779ec05c7b8710c85a24989857b6d5fa907f2.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ifbbbe5290364d20daa098ac5c2c4b3f2a0d67e1882746b298d0c4f90fb8dec53.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ifbbbe5290364d20daa098ac5c2c4b3f2a0d67e1882746b298d0c4f90fb8dec53.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i7201fc60f93506ab507aba55127067ad853520cd3b1bc02548cbb94c60ed14f1.ExportPersonalDataRequestBuilder) {
    return i7201fc60f93506ab507aba55127067ad853520cd3b1bc02548cbb94c60ed14f1.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i5c64c484fc87a05f75b11cb79b985d50f6a0ef5b796a0f8d6b38608d07c1dc80.FindMeetingTimesRequestBuilder) {
    return i5c64c484fc87a05f75b11cb79b985d50f6a0ef5b796a0f8d6b38608d07c1dc80.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i859fa10100610ebcdde7ac39005732ad2c78c0956d868a2cdd060dbde71d0785.FindRoomListsRequestBuilder) {
    return i859fa10100610ebcdde7ac39005732ad2c78c0956d868a2cdd060dbde71d0785.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ia5c7e3e68cc560feb74cab0573eb2b92ee3caffee9bcd099ad9aaaaf1a8f2ece.FindRoomsRequestBuilder) {
    return ia5c7e3e68cc560feb74cab0573eb2b92ee3caffee9bcd099ad9aaaaf1a8f2ece.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i51f2ccc7038aeea321321bd75f149854eb7892cd5c871c3e4c5354f018c73c48.FindRoomsWithRoomListRequestBuilder) {
    return i51f2ccc7038aeea321321bd75f149854eb7892cd5c871c3e4c5354f018c73c48.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i7d20cb7abcc98a52e8d49ef40215e6750c95a794dbd9e2e81be74631976880fa.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i7d20cb7abcc98a52e8d49ef40215e6750c95a794dbd9e2e81be74631976880fa.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ie8b8fb27d7254e798aacb21b8417e8b0e30a7a923b637e03d1ce41c6e75a1448.GetLoggedOnManagedDevicesRequestBuilder) {
    return ie8b8fb27d7254e798aacb21b8417e8b0e30a7a923b637e03d1ce41c6e75a1448.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i98fd3840004bbe47a878dcf9635e0425c97446b349b35f5ab65013f5caa62d17.GetMailTipsRequestBuilder) {
    return i98fd3840004bbe47a878dcf9635e0425c97446b349b35f5ab65013f5caa62d17.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id375a73f8562f1afdea4c927701c0f4e21fb27a3727222c79c87a4f49a79c621.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id375a73f8562f1afdea4c927701c0f4e21fb27a3727222c79c87a4f49a79c621.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i117fdf5b2eac599662922eec9ad4be84b3227e500937ce8b54045f42f2f9b864.GetManagedAppPoliciesRequestBuilder) {
    return i117fdf5b2eac599662922eec9ad4be84b3227e500937ce8b54045f42f2f9b864.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i15cfd0907ef334819d192f0caae869887353a08aebd4f57a6d50609f90a38324.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i15cfd0907ef334819d192f0caae869887353a08aebd4f57a6d50609f90a38324.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*iba8d9c4db75c7d271f5066be861a9c614fa33ca32ed6fb0bc21cc3736723589f.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return iba8d9c4db75c7d271f5066be861a9c614fa33ca32ed6fb0bc21cc3736723589f.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ief42293257452c74833dc67d9096853dd9d4f7d9d83d9a8a76da90904c35f220.GetMemberGroupsRequestBuilder) {
    return ief42293257452c74833dc67d9096853dd9d4f7d9d83d9a8a76da90904c35f220.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ib623c31512b6e82ea374585c36ebe57ae562fcdb244e986b577bb24d9d76f291.GetMemberObjectsRequestBuilder) {
    return ib623c31512b6e82ea374585c36ebe57ae562fcdb244e986b577bb24d9d76f291.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i0a60f9ea2b6cce09831def3f509d4df59a4cab86059636d57af6afa9dceb712a.InvalidateAllRefreshTokensRequestBuilder) {
    return i0a60f9ea2b6cce09831def3f509d4df59a4cab86059636d57af6afa9dceb712a.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i77a8322a69419d2f12b42c265833cce281302b9c7cc7a22fdf7cb7674873fb95.IsManagedAppUserBlockedRequestBuilder) {
    return i77a8322a69419d2f12b42c265833cce281302b9c7cc7a22fdf7cb7674873fb95.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ifff60123811a96a13ebaf593c1fc70fa0d34d8408b9a1c03b6a1d8b16772eaf2.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ifff60123811a96a13ebaf593c1fc70fa0d34d8408b9a1c03b6a1d8b16772eaf2.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i6f660b2006606b84de3153fd9d5b29547e9b86b0b7996ac3d5d9506c1d52e5b6.RemoveAllDevicesFromManagementRequestBuilder) {
    return i6f660b2006606b84de3153fd9d5b29547e9b86b0b7996ac3d5d9506c1d52e5b6.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i7e6d1819cc25e13a93a3bbf72cc883fe05449ba38ff9aa1d59796a5b7a4e705e.ReprocessLicenseAssignmentRequestBuilder) {
    return i7e6d1819cc25e13a93a3bbf72cc883fe05449ba38ff9aa1d59796a5b7a4e705e.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ic74dbb2c637e6585f461876e21663edbe94a58d6db8017de79cfa610d5fb69bc.RestoreRequestBuilder) {
    return ic74dbb2c637e6585f461876e21663edbe94a58d6db8017de79cfa610d5fb69bc.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i37dda0334f2b355a2d6059b26d596a2995283976d7cc804d0949689e2b46ac3d.RevokeSignInSessionsRequestBuilder) {
    return i37dda0334f2b355a2d6059b26d596a2995283976d7cc804d0949689e2b46ac3d.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*if0bdacbb8a78c8dbee95e09187dbd7133f8f3ff7b44b66075ce503671ee96ed1.SendMailRequestBuilder) {
    return if0bdacbb8a78c8dbee95e09187dbd7133f8f3ff7b44b66075ce503671ee96ed1.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i11afdfb304ea23402ce262feb0b1388ba316e975b68144a7e08af68b0d727246.TranslateExchangeIdsRequestBuilder) {
    return i11afdfb304ea23402ce262feb0b1388ba316e975b68144a7e08af68b0d727246.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i3fd31d011299e4bd74e1cf12016fa1f926e6777b1d2d27d7fc28b9a3ee3b54ba.UnblockManagedAppsRequestBuilder) {
    return i3fd31d011299e4bd74e1cf12016fa1f926e6777b1d2d27d7fc28b9a3ee3b54ba.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i253f8948a060f08ed80843c9f8ec8ffb9bc98a9dc85d1e77a4592bd3956e8fb4.WipeAndBlockManagedAppsRequestBuilder) {
    return i253f8948a060f08ed80843c9f8ec8ffb9bc98a9dc85d1e77a4592bd3956e8fb4.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*iaaaf1941f01165dfcd733293a08b35738af3c0c57f52173888cdb0bc99e39602.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return iaaaf1941f01165dfcd733293a08b35738af3c0c57f52173888cdb0bc99e39602.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*iff75f05397b208c225848e8544853f552816e44d41e74edaa722798a39069705.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return iff75f05397b208c225848e8544853f552816e44d41e74edaa722798a39069705.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i1e008e60a8eb48a26abdd8c95fe5c8356221274d610db379067882383cf00892.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i1e008e60a8eb48a26abdd8c95fe5c8356221274d610db379067882383cf00892.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
