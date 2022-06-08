package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0d0aeb189d27f573d552fe48f36e47ca98471b76adcd830efce26cc601191c08 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/reprocesslicenseassignment"
    i2698b9df8f6487ecfc0ca9663fc2c1b0e8c6f67b0f056f7d3f1ad7556a78cf13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i2e6fe69ef64555ed2933be804bf5fab3d925a50b88cde5f365fceffa21d8e9d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/findrooms"
    i372cd79e2eb49f7720a4b195febe977df0de0b40ec46cbd6bd05ffaaed131660 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i3eb1026081fb4ee2ce0d5a26233229051d0e2ceac71c47a59649db446bc07fe3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/wipeandblockmanagedapps"
    i451b0c05682eb63370298fe3b54aad7910e1af7de0b4a1d50d1f060b05a773fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/removealldevicesfrommanagement"
    i526a0e2c8efe259a55cdfe30c37f2a431b4ff74c58ece7c5fcf623cca6580021 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/findmeetingtimes"
    i544f7fb86ba8748423d2452df9d8af80b087fd603741086c9a51f0ba8c7f24cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmailtips"
    i55dda30b75d6aa167034c6ff542c510593fc2a020da68d2209856d1d9cac4671 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/assignlicense"
    i55e78f4c4056b6365a4a76e0fb076ca2082be44df161b479cefa8576a678a3e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/invalidateallrefreshtokens"
    i5a641c6758f14020cd55c79864fa6f8de134e5bcf90b835aca1f5b13c442eec8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/activateserviceplan"
    i62279822abe0774faf9cf9577d7c49a880395bae379a96f35a05607aca81356f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/findroomswithroomlist"
    i641bdbcbe502c2e0c603f2549a069b156f691166f19f01926f1f40fea0246eb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i65aa85c77726585fce129c9ac2245c69dacd863f4da04c611fb3de15c244af89 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/revokesigninsessions"
    i66aa27d738a5694df7bf6d11e0d6c73a8e6b63851736a1f0d0ee4aad9a25f002 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmanagedapppolicies"
    i6763c92105e307571ba1e614b93ad187608074012122f9a4f943dd287d316f64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/ismanagedappuserblocked"
    i70af77338df14a4c5479d4a41ffaf801e271aac13e8cc8b80508692e705e3c52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/findroomlists"
    i7135d29f3c7978d4b37955189181218099c6cadb49ffd56d2f5ac5528dc810a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/changepassword"
    i713de87d8839effab083ccf53c640fc1ec3abd5c72ab10b908d79510d0263216 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmanageddeviceswithappfailures"
    i72a9247bcc82f91ed80a6ee2bf2b651a099146b9ef3708c65232f1f34e6bef26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getloggedonmanageddevices"
    i776419ab5456f3a2ea6e0e69c98f20d5088cc90b5f01fa8836aacdd1d26d64fc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i78eb7a39f62afea4ff8856ead63a3cc3845652a202fb45584411264f079e3660 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/exportpersonaldata"
    i87b611d81f2ddc2c556273ad5b36e342a3a20937fbf94a2605f9d9ba99a52ff0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/checkmembergroups"
    i9c965b40b8bdb23ab0ae1c710df15e41a75af8464d2636841b94df7f1e982ec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/checkmemberobjects"
    i9f53ba1abccd283ccc5ef7398a642b84c707249f228410bc5f7d9c40bc0fb50b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/translateexchangeids"
    ibb368cb17ee4e0a6f44744793429f27a9ce22e49ecc187c2d3a875c8cef5a10f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmembergroups"
    ibe76d55f8fb8b9659724818082d9aea3740c0bc0270a81cb771ff7a8d2e63555 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/restore"
    ic24c393692ceba3c9465e084682977e2444df16851301e2ee9d9fc84c4db5108 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/unblockmanagedapps"
    ic305b29721c669ace16654a966e3e8a121e11d2453e6313f6a813791f520a766 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/exportdeviceandappmanagementdata"
    ic49ebf14bc09f3dba0c564d57ca39e4a093369aafb3f91a154b0609a371ce21a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/wipemanagedappregistrationbydevicetag"
    ic7b7a0901e1f27b361204c820032565a525cdfa4f82be38fdc11915ee702666a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmemberobjects"
    ic7d5e8bac00429489fbca9c5182894c22331285d1a08bbd51e0cc90ba70caf3e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    idc68143cda5508b4cb7c8d37344ea56ef0907f02a592098cf38c632aa74cff5a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ie5f67d97e8dbcaba0d482a5c0880d30dfe0cc135c371793220cc8965a79bf589 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/sendmail"
    ie9cff2e437e9f631614b589cfe47a6f299e288897ef4b38cbb6d136bda7322c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user/getmanagedappdiagnosticstatuses"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i5a641c6758f14020cd55c79864fa6f8de134e5bcf90b835aca1f5b13c442eec8.ActivateServicePlanRequestBuilder) {
    return i5a641c6758f14020cd55c79864fa6f8de134e5bcf90b835aca1f5b13c442eec8.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i55dda30b75d6aa167034c6ff542c510593fc2a020da68d2209856d1d9cac4671.AssignLicenseRequestBuilder) {
    return i55dda30b75d6aa167034c6ff542c510593fc2a020da68d2209856d1d9cac4671.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i7135d29f3c7978d4b37955189181218099c6cadb49ffd56d2f5ac5528dc810a8.ChangePasswordRequestBuilder) {
    return i7135d29f3c7978d4b37955189181218099c6cadb49ffd56d2f5ac5528dc810a8.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i87b611d81f2ddc2c556273ad5b36e342a3a20937fbf94a2605f9d9ba99a52ff0.CheckMemberGroupsRequestBuilder) {
    return i87b611d81f2ddc2c556273ad5b36e342a3a20937fbf94a2605f9d9ba99a52ff0.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i9c965b40b8bdb23ab0ae1c710df15e41a75af8464d2636841b94df7f1e982ec0.CheckMemberObjectsRequestBuilder) {
    return i9c965b40b8bdb23ab0ae1c710df15e41a75af8464d2636841b94df7f1e982ec0.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ic305b29721c669ace16654a966e3e8a121e11d2453e6313f6a813791f520a766.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ic305b29721c669ace16654a966e3e8a121e11d2453e6313f6a813791f520a766.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i641bdbcbe502c2e0c603f2549a069b156f691166f19f01926f1f40fea0246eb4.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i641bdbcbe502c2e0c603f2549a069b156f691166f19f01926f1f40fea0246eb4.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i78eb7a39f62afea4ff8856ead63a3cc3845652a202fb45584411264f079e3660.ExportPersonalDataRequestBuilder) {
    return i78eb7a39f62afea4ff8856ead63a3cc3845652a202fb45584411264f079e3660.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i526a0e2c8efe259a55cdfe30c37f2a431b4ff74c58ece7c5fcf623cca6580021.FindMeetingTimesRequestBuilder) {
    return i526a0e2c8efe259a55cdfe30c37f2a431b4ff74c58ece7c5fcf623cca6580021.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i70af77338df14a4c5479d4a41ffaf801e271aac13e8cc8b80508692e705e3c52.FindRoomListsRequestBuilder) {
    return i70af77338df14a4c5479d4a41ffaf801e271aac13e8cc8b80508692e705e3c52.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i2e6fe69ef64555ed2933be804bf5fab3d925a50b88cde5f365fceffa21d8e9d4.FindRoomsRequestBuilder) {
    return i2e6fe69ef64555ed2933be804bf5fab3d925a50b88cde5f365fceffa21d8e9d4.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i62279822abe0774faf9cf9577d7c49a880395bae379a96f35a05607aca81356f.FindRoomsWithRoomListRequestBuilder) {
    return i62279822abe0774faf9cf9577d7c49a880395bae379a96f35a05607aca81356f.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*idc68143cda5508b4cb7c8d37344ea56ef0907f02a592098cf38c632aa74cff5a.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return idc68143cda5508b4cb7c8d37344ea56ef0907f02a592098cf38c632aa74cff5a.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i72a9247bcc82f91ed80a6ee2bf2b651a099146b9ef3708c65232f1f34e6bef26.GetLoggedOnManagedDevicesRequestBuilder) {
    return i72a9247bcc82f91ed80a6ee2bf2b651a099146b9ef3708c65232f1f34e6bef26.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i544f7fb86ba8748423d2452df9d8af80b087fd603741086c9a51f0ba8c7f24cf.GetMailTipsRequestBuilder) {
    return i544f7fb86ba8748423d2452df9d8af80b087fd603741086c9a51f0ba8c7f24cf.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ie9cff2e437e9f631614b589cfe47a6f299e288897ef4b38cbb6d136bda7322c4.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ie9cff2e437e9f631614b589cfe47a6f299e288897ef4b38cbb6d136bda7322c4.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i66aa27d738a5694df7bf6d11e0d6c73a8e6b63851736a1f0d0ee4aad9a25f002.GetManagedAppPoliciesRequestBuilder) {
    return i66aa27d738a5694df7bf6d11e0d6c73a8e6b63851736a1f0d0ee4aad9a25f002.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i713de87d8839effab083ccf53c640fc1ec3abd5c72ab10b908d79510d0263216.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i713de87d8839effab083ccf53c640fc1ec3abd5c72ab10b908d79510d0263216.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ic7d5e8bac00429489fbca9c5182894c22331285d1a08bbd51e0cc90ba70caf3e.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ic7d5e8bac00429489fbca9c5182894c22331285d1a08bbd51e0cc90ba70caf3e.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ibb368cb17ee4e0a6f44744793429f27a9ce22e49ecc187c2d3a875c8cef5a10f.GetMemberGroupsRequestBuilder) {
    return ibb368cb17ee4e0a6f44744793429f27a9ce22e49ecc187c2d3a875c8cef5a10f.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ic7b7a0901e1f27b361204c820032565a525cdfa4f82be38fdc11915ee702666a.GetMemberObjectsRequestBuilder) {
    return ic7b7a0901e1f27b361204c820032565a525cdfa4f82be38fdc11915ee702666a.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i55e78f4c4056b6365a4a76e0fb076ca2082be44df161b479cefa8576a678a3e0.InvalidateAllRefreshTokensRequestBuilder) {
    return i55e78f4c4056b6365a4a76e0fb076ca2082be44df161b479cefa8576a678a3e0.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i6763c92105e307571ba1e614b93ad187608074012122f9a4f943dd287d316f64.IsManagedAppUserBlockedRequestBuilder) {
    return i6763c92105e307571ba1e614b93ad187608074012122f9a4f943dd287d316f64.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i372cd79e2eb49f7720a4b195febe977df0de0b40ec46cbd6bd05ffaaed131660.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i372cd79e2eb49f7720a4b195febe977df0de0b40ec46cbd6bd05ffaaed131660.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i451b0c05682eb63370298fe3b54aad7910e1af7de0b4a1d50d1f060b05a773fb.RemoveAllDevicesFromManagementRequestBuilder) {
    return i451b0c05682eb63370298fe3b54aad7910e1af7de0b4a1d50d1f060b05a773fb.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0d0aeb189d27f573d552fe48f36e47ca98471b76adcd830efce26cc601191c08.ReprocessLicenseAssignmentRequestBuilder) {
    return i0d0aeb189d27f573d552fe48f36e47ca98471b76adcd830efce26cc601191c08.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ibe76d55f8fb8b9659724818082d9aea3740c0bc0270a81cb771ff7a8d2e63555.RestoreRequestBuilder) {
    return ibe76d55f8fb8b9659724818082d9aea3740c0bc0270a81cb771ff7a8d2e63555.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i65aa85c77726585fce129c9ac2245c69dacd863f4da04c611fb3de15c244af89.RevokeSignInSessionsRequestBuilder) {
    return i65aa85c77726585fce129c9ac2245c69dacd863f4da04c611fb3de15c244af89.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ie5f67d97e8dbcaba0d482a5c0880d30dfe0cc135c371793220cc8965a79bf589.SendMailRequestBuilder) {
    return ie5f67d97e8dbcaba0d482a5c0880d30dfe0cc135c371793220cc8965a79bf589.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i9f53ba1abccd283ccc5ef7398a642b84c707249f228410bc5f7d9c40bc0fb50b.TranslateExchangeIdsRequestBuilder) {
    return i9f53ba1abccd283ccc5ef7398a642b84c707249f228410bc5f7d9c40bc0fb50b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ic24c393692ceba3c9465e084682977e2444df16851301e2ee9d9fc84c4db5108.UnblockManagedAppsRequestBuilder) {
    return ic24c393692ceba3c9465e084682977e2444df16851301e2ee9d9fc84c4db5108.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i3eb1026081fb4ee2ce0d5a26233229051d0e2ceac71c47a59649db446bc07fe3.WipeAndBlockManagedAppsRequestBuilder) {
    return i3eb1026081fb4ee2ce0d5a26233229051d0e2ceac71c47a59649db446bc07fe3.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ic49ebf14bc09f3dba0c564d57ca39e4a093369aafb3f91a154b0609a371ce21a.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ic49ebf14bc09f3dba0c564d57ca39e4a093369aafb3f91a154b0609a371ce21a.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i2698b9df8f6487ecfc0ca9663fc2c1b0e8c6f67b0f056f7d3f1ad7556a78cf13.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i2698b9df8f6487ecfc0ca9663fc2c1b0e8c6f67b0f056f7d3f1ad7556a78cf13.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i776419ab5456f3a2ea6e0e69c98f20d5088cc90b5f01fa8836aacdd1d26d64fc.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i776419ab5456f3a2ea6e0e69c98f20d5088cc90b5f01fa8836aacdd1d26d64fc.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
