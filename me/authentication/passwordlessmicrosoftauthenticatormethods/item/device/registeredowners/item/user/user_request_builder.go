package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0e69c298ca7bd0b018ab19e67fcd231ed9947ce73c90ab9818eeb513d181b4a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
    i0fae2ddb6e297dd1eac114f72da47da03250dd32d3c265f266fe9dd720ecb3a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    i10047cd3c09bd5ab2362fecd1dbc05f617180b081e505492c46dbe6ab26fb718 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    i100ae8ce36fa6a051be681efc0bfac95e04f65141e5b85596dc15b0a783740a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    i1d5143861ed57823eff96a768641863da2cbfa576c4516fc7bbf26951565fa14 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i2cd219f9d9343aeffcbbefac13ab8cbcec37703d7e4b3da327bb9c0999801e6c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    i2cf59c91756953b367f80e11b467f9f91a0358242f1f27a3b81ed3f44b172469 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    i39457aaef35346d3ac4501d615f573bdbe6df4389a7b2dcffeff1d52d5967495 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    i529f28be534ab0935c2bacea3a4fc0c028dc3fb861714a391e624377f8fc33b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    i53f13f12b075bb5e51cf810c5dea210dd22bb659c3001c67185404d33a11782e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i54147b6703e54ea25c258449e45c7ea4b1126c40a889e9852a070458f8786698 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    i54b0b1a775aa0c6b3b152c5b9cbe5b6ff4049e2ebb15ac02a20811c9aefc3ff8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    i5959356d3e9c31cacc74b0c90836bf0dca0e24a1f9412570c76f1f5573aca002 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    i5e27ab47f37f3cc5f667c9b03a6eb3e885f8ea2b8cb2bec2b37b5c35fdb289f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
    i6ee0b142d12dafeae8ec3a58e6d50cc3fcce5628ee6b3dffb813561c0b3f1da2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    i7e7d9d4d5039eaa6a201843d8abe0edc9551aa4ad430e0c5e38e5ddc3bb3aaf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    i94347efce6436c35225d8195ccb06f4adbb4226e64c56d8edf0b15239e577ed8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findroomswithroomlist"
    i951d8a45da93feee9929caaf5a896703c77020a3724969b32e8c66fb31504896 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    i9ef01eba45baa930b7fa3e4af1180e89f2f054fee8ac74baa3216542445a9488 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    ia0dcc458b4f4e9528bd92ea42b64ffb314f20a03eda866454d2551bfd7b5ec98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
    ib4ce79be99316ec9d59e19610485dd8b6104f3f5750b74438d340f6dc273cc6f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    ib998218a4ccb59a203ced83d9b0700d15975073dc8e073542c80a6f6ea1865a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ic736ac8b774ae934845e5ab5675129988409c195ac44d1e98cd09e7ea08acd25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    ic9a6ecb22fdf590029bbaacf1122791c1339db2257cc2b69ab28b5ae1553a236 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    icb266722556c478e27b24610d82f5b55dd557473af2fef9c0e158f1eb94675f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    icda6bd0bda35889f7e38b41311d51fe59c6dbb417ca5ebc95c65ce3a93f45a3f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
    icdd40cc6f1636fc9f3acc2a8165d364cb2a122e5b962acfd786fcf256fe7f941 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    id049fc524b7cb4c6b4a05a10ecbd960b88622f20bcd743acc8165c80596e1308 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findrooms"
    id375d4b0538670d7850ba94a542b43f1025f1ac566a634a29255be35fb477e44 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/unblockmanagedapps"
    id5a0c4977c7833e83151e903fe66472a1832c4f48a3441392b27f1e30c982ca4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
    id984371e971e545449f1530a7cd8ccfad784fd7633949b82ed93a2e97b9b38e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    idda60e7c4b3ead3f7c3dbc1a3f7b92dcc7b98130e2aa01e55df671bfc527efed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    ide4da2e590788e7a64831aa4bf51167338e55109eeb4dce03456e9feba551797 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/findroomlists"
    ie6f6de465b84c2f6f09642c492959fcf73b4f60bba7891890ddf4e724ba91130 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    if8897c942f49027a6921dfc209ae0a44b33ac1ed12f95e421d13090e8a474a4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user/activateserviceplan"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*if8897c942f49027a6921dfc209ae0a44b33ac1ed12f95e421d13090e8a474a4e.ActivateServicePlanRequestBuilder) {
    return if8897c942f49027a6921dfc209ae0a44b33ac1ed12f95e421d13090e8a474a4e.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i54147b6703e54ea25c258449e45c7ea4b1126c40a889e9852a070458f8786698.AssignLicenseRequestBuilder) {
    return i54147b6703e54ea25c258449e45c7ea4b1126c40a889e9852a070458f8786698.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*id984371e971e545449f1530a7cd8ccfad784fd7633949b82ed93a2e97b9b38e5.ChangePasswordRequestBuilder) {
    return id984371e971e545449f1530a7cd8ccfad784fd7633949b82ed93a2e97b9b38e5.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i6ee0b142d12dafeae8ec3a58e6d50cc3fcce5628ee6b3dffb813561c0b3f1da2.CheckMemberGroupsRequestBuilder) {
    return i6ee0b142d12dafeae8ec3a58e6d50cc3fcce5628ee6b3dffb813561c0b3f1da2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*icda6bd0bda35889f7e38b41311d51fe59c6dbb417ca5ebc95c65ce3a93f45a3f.CheckMemberObjectsRequestBuilder) {
    return icda6bd0bda35889f7e38b41311d51fe59c6dbb417ca5ebc95c65ce3a93f45a3f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*id5a0c4977c7833e83151e903fe66472a1832c4f48a3441392b27f1e30c982ca4.ExportDeviceAndAppManagementDataRequestBuilder) {
    return id5a0c4977c7833e83151e903fe66472a1832c4f48a3441392b27f1e30c982ca4.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ib998218a4ccb59a203ced83d9b0700d15975073dc8e073542c80a6f6ea1865a4.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ib998218a4ccb59a203ced83d9b0700d15975073dc8e073542c80a6f6ea1865a4.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i529f28be534ab0935c2bacea3a4fc0c028dc3fb861714a391e624377f8fc33b0.ExportPersonalDataRequestBuilder) {
    return i529f28be534ab0935c2bacea3a4fc0c028dc3fb861714a391e624377f8fc33b0.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i0e69c298ca7bd0b018ab19e67fcd231ed9947ce73c90ab9818eeb513d181b4a5.FindMeetingTimesRequestBuilder) {
    return i0e69c298ca7bd0b018ab19e67fcd231ed9947ce73c90ab9818eeb513d181b4a5.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ide4da2e590788e7a64831aa4bf51167338e55109eeb4dce03456e9feba551797.FindRoomListsRequestBuilder) {
    return ide4da2e590788e7a64831aa4bf51167338e55109eeb4dce03456e9feba551797.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*id049fc524b7cb4c6b4a05a10ecbd960b88622f20bcd743acc8165c80596e1308.FindRoomsRequestBuilder) {
    return id049fc524b7cb4c6b4a05a10ecbd960b88622f20bcd743acc8165c80596e1308.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i94347efce6436c35225d8195ccb06f4adbb4226e64c56d8edf0b15239e577ed8.FindRoomsWithRoomListRequestBuilder) {
    return i94347efce6436c35225d8195ccb06f4adbb4226e64c56d8edf0b15239e577ed8.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i5959356d3e9c31cacc74b0c90836bf0dca0e24a1f9412570c76f1f5573aca002.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i5959356d3e9c31cacc74b0c90836bf0dca0e24a1f9412570c76f1f5573aca002.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ia0dcc458b4f4e9528bd92ea42b64ffb314f20a03eda866454d2551bfd7b5ec98.GetLoggedOnManagedDevicesRequestBuilder) {
    return ia0dcc458b4f4e9528bd92ea42b64ffb314f20a03eda866454d2551bfd7b5ec98.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i10047cd3c09bd5ab2362fecd1dbc05f617180b081e505492c46dbe6ab26fb718.GetMailTipsRequestBuilder) {
    return i10047cd3c09bd5ab2362fecd1dbc05f617180b081e505492c46dbe6ab26fb718.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ie6f6de465b84c2f6f09642c492959fcf73b4f60bba7891890ddf4e724ba91130.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ie6f6de465b84c2f6f09642c492959fcf73b4f60bba7891890ddf4e724ba91130.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic736ac8b774ae934845e5ab5675129988409c195ac44d1e98cd09e7ea08acd25.GetManagedAppPoliciesRequestBuilder) {
    return ic736ac8b774ae934845e5ab5675129988409c195ac44d1e98cd09e7ea08acd25.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i2cd219f9d9343aeffcbbefac13ab8cbcec37703d7e4b3da327bb9c0999801e6c.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i2cd219f9d9343aeffcbbefac13ab8cbcec37703d7e4b3da327bb9c0999801e6c.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7e7d9d4d5039eaa6a201843d8abe0edc9551aa4ad430e0c5e38e5ddc3bb3aaf9.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7e7d9d4d5039eaa6a201843d8abe0edc9551aa4ad430e0c5e38e5ddc3bb3aaf9.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i5e27ab47f37f3cc5f667c9b03a6eb3e885f8ea2b8cb2bec2b37b5c35fdb289f0.GetMemberGroupsRequestBuilder) {
    return i5e27ab47f37f3cc5f667c9b03a6eb3e885f8ea2b8cb2bec2b37b5c35fdb289f0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i39457aaef35346d3ac4501d615f573bdbe6df4389a7b2dcffeff1d52d5967495.GetMemberObjectsRequestBuilder) {
    return i39457aaef35346d3ac4501d615f573bdbe6df4389a7b2dcffeff1d52d5967495.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i9ef01eba45baa930b7fa3e4af1180e89f2f054fee8ac74baa3216542445a9488.InvalidateAllRefreshTokensRequestBuilder) {
    return i9ef01eba45baa930b7fa3e4af1180e89f2f054fee8ac74baa3216542445a9488.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*idda60e7c4b3ead3f7c3dbc1a3f7b92dcc7b98130e2aa01e55df671bfc527efed.IsManagedAppUserBlockedRequestBuilder) {
    return idda60e7c4b3ead3f7c3dbc1a3f7b92dcc7b98130e2aa01e55df671bfc527efed.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i1d5143861ed57823eff96a768641863da2cbfa576c4516fc7bbf26951565fa14.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1d5143861ed57823eff96a768641863da2cbfa576c4516fc7bbf26951565fa14.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i54b0b1a775aa0c6b3b152c5b9cbe5b6ff4049e2ebb15ac02a20811c9aefc3ff8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i54b0b1a775aa0c6b3b152c5b9cbe5b6ff4049e2ebb15ac02a20811c9aefc3ff8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i2cf59c91756953b367f80e11b467f9f91a0358242f1f27a3b81ed3f44b172469.ReprocessLicenseAssignmentRequestBuilder) {
    return i2cf59c91756953b367f80e11b467f9f91a0358242f1f27a3b81ed3f44b172469.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ic9a6ecb22fdf590029bbaacf1122791c1339db2257cc2b69ab28b5ae1553a236.RestoreRequestBuilder) {
    return ic9a6ecb22fdf590029bbaacf1122791c1339db2257cc2b69ab28b5ae1553a236.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i100ae8ce36fa6a051be681efc0bfac95e04f65141e5b85596dc15b0a783740a0.RevokeSignInSessionsRequestBuilder) {
    return i100ae8ce36fa6a051be681efc0bfac95e04f65141e5b85596dc15b0a783740a0.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i951d8a45da93feee9929caaf5a896703c77020a3724969b32e8c66fb31504896.SendMailRequestBuilder) {
    return i951d8a45da93feee9929caaf5a896703c77020a3724969b32e8c66fb31504896.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i0fae2ddb6e297dd1eac114f72da47da03250dd32d3c265f266fe9dd720ecb3a2.TranslateExchangeIdsRequestBuilder) {
    return i0fae2ddb6e297dd1eac114f72da47da03250dd32d3c265f266fe9dd720ecb3a2.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*id375d4b0538670d7850ba94a542b43f1025f1ac566a634a29255be35fb477e44.UnblockManagedAppsRequestBuilder) {
    return id375d4b0538670d7850ba94a542b43f1025f1ac566a634a29255be35fb477e44.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ib4ce79be99316ec9d59e19610485dd8b6104f3f5750b74438d340f6dc273cc6f.WipeAndBlockManagedAppsRequestBuilder) {
    return ib4ce79be99316ec9d59e19610485dd8b6104f3f5750b74438d340f6dc273cc6f.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*icb266722556c478e27b24610d82f5b55dd557473af2fef9c0e158f1eb94675f1.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return icb266722556c478e27b24610d82f5b55dd557473af2fef9c0e158f1eb94675f1.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i53f13f12b075bb5e51cf810c5dea210dd22bb659c3001c67185404d33a11782e.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i53f13f12b075bb5e51cf810c5dea210dd22bb659c3001c67185404d33a11782e.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*icdd40cc6f1636fc9f3acc2a8165d364cb2a122e5b962acfd786fcf256fe7f941.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return icdd40cc6f1636fc9f3acc2a8165d364cb2a122e5b962acfd786fcf256fe7f941.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
