package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b39097657fbe88b406fc1b6eb76d967394330fbd70d43c5b00584005d5245d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findroomswithroomlist"
    i0e53131874c90de60c7fc2ae1935b28d78cecf24d50a5f0ecd11f421e3e202e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i168d3500fb56817429a7bdc3eaab3449ba6be14733419d3e349e4dc298ffc78c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    i1f0a5fa96b97d9bca8d455835fa1c564151dc00f916c7d2f987b36a0a241304e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    i26b9051a79724e72b9a2512bbd3aabdce9071c3deaf2e48487a7a49e4c2d4f89 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i2ece9e43d66dc3566a1b48c72063e10ae0b31ad2ab5d7889f13949f71d79b471 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    i35063fd7f1ad0589fc94fb3b6709383ece8aa11f4dc0f9b8e14dc5a5746554d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    i482737223efe8dc02b6546ed734fc77004eff46fb9e1e8c65528db4fa3dbba52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    i4915fe5bb53e06c02bd8e3b8a79a11f606f6e3c9fef9f5922db7b5e30ce5f3f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    i4da4b388440a3c0081b0e338987c597f5acfd63087113a16eda8a9d6470208e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    i51225c312ec43bff2bc259c7b06388e0746b8bbbc86e0c2fc604dfcf4b0759f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i5166400e837b94070d39d0f014dda18b1079fa90623de5748871e244389f04a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    i52ab2e51018f1b3efc8f533ba6d360c21a1ef57a47fca077007dabf7a112fbbe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/unblockmanagedapps"
    i5f8c2564049e25a27802d0cfd962784015841b1f258fae1a62d0dc68349a04c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/activateserviceplan"
    i67da788c4d3c4b7fd02d651252e0de29864355b61b8c77e337fd9b0c66036bac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
    i682bc70d0442cfa8e2a3e10c18a413d44877bb859a973c8257c3a379982af4ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    i6c9b18facb8fad827fcbdbe230f3745e18ceab8d91b9a422aa6ab5c6d9ad8f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i6d222ff92d14ce4110b7e01d31b1f087405b9c8ace22bac5d8c0bc12bb40c9cc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findrooms"
    i6d64fea2a0b0e1895f236262a212af51360e9416e3308a614e50a52768510ba9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    i7421f569518b6d911f85f6ad0d9c80a510968d9b4dbc1dfaa6a3de58f7e01400 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    i7bb7af3e4ffdf11994b91bc232d949fe47a72edc5937c09c9d09dd8d74fc66a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    i8207065a4c07d73ecae8dfadff355be778d398b00fd5869960b00c65760d1997 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i8606a6d0bde4527a5685ef9ce8fae4475d25fb6171ef9169a1d9933f72ad47ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i8e87e6bd8b282d92061805d7fc1fb846933b00861571b2c7c52934a71e6a8cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    i914a5296dd31b5629c12f333c911dceb2b70d67f6b9c127b5ee1897ebfce9b21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i98dc5c151ff6f7d85030e85bc231e1bac7f07c0d3199c93283033d77882ca17f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i9d988c9f15ff515c4a60e2c9c1d63f08b95c532c810d48fde287c4671df73320 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    ia1ea56b38e735a7b1958cf6dda16705671af71f4dd1a53d3ff462e17b48b8402 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    ib7e8a57a233b8605c617878983f83774822a84e0787f7f7c45204228ecf7e8b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findroomlists"
    ib8e88fe7e3688aca9e2b0dca5ad5851ff2acd35c2c4d81999a643ac27d1390a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    id64f5a6d9f2c0b8f62bc375576c508d23e1a0f114c48846e32bbb0573031ccda "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    idcb80e15afddfb892c1a90f7152accaa4c561738893f238c3bf92169544794b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    idd3ace18bfe75607474e31976e474dd5fad5f83a5a77cd6040011c821bdad754 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ieebe41c97b427d47bd5685216485f1b8eda686c00c0fd340edad3ecb3ee1d9b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    if907b6d9d4480b39bd0ab57a7686b150c27544885e14f4729604a03ea4e727ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i5f8c2564049e25a27802d0cfd962784015841b1f258fae1a62d0dc68349a04c3.ActivateServicePlanRequestBuilder) {
    return i5f8c2564049e25a27802d0cfd962784015841b1f258fae1a62d0dc68349a04c3.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ieebe41c97b427d47bd5685216485f1b8eda686c00c0fd340edad3ecb3ee1d9b8.AssignLicenseRequestBuilder) {
    return ieebe41c97b427d47bd5685216485f1b8eda686c00c0fd340edad3ecb3ee1d9b8.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i168d3500fb56817429a7bdc3eaab3449ba6be14733419d3e349e4dc298ffc78c.ChangePasswordRequestBuilder) {
    return i168d3500fb56817429a7bdc3eaab3449ba6be14733419d3e349e4dc298ffc78c.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i8e87e6bd8b282d92061805d7fc1fb846933b00861571b2c7c52934a71e6a8cc2.CheckMemberGroupsRequestBuilder) {
    return i8e87e6bd8b282d92061805d7fc1fb846933b00861571b2c7c52934a71e6a8cc2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ia1ea56b38e735a7b1958cf6dda16705671af71f4dd1a53d3ff462e17b48b8402.CheckMemberObjectsRequestBuilder) {
    return ia1ea56b38e735a7b1958cf6dda16705671af71f4dd1a53d3ff462e17b48b8402.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i4915fe5bb53e06c02bd8e3b8a79a11f606f6e3c9fef9f5922db7b5e30ce5f3f9.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i4915fe5bb53e06c02bd8e3b8a79a11f606f6e3c9fef9f5922db7b5e30ce5f3f9.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i914a5296dd31b5629c12f333c911dceb2b70d67f6b9c127b5ee1897ebfce9b21.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i914a5296dd31b5629c12f333c911dceb2b70d67f6b9c127b5ee1897ebfce9b21.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i51225c312ec43bff2bc259c7b06388e0746b8bbbc86e0c2fc604dfcf4b0759f8.ExportPersonalDataRequestBuilder) {
    return i51225c312ec43bff2bc259c7b06388e0746b8bbbc86e0c2fc604dfcf4b0759f8.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i9d988c9f15ff515c4a60e2c9c1d63f08b95c532c810d48fde287c4671df73320.FindMeetingTimesRequestBuilder) {
    return i9d988c9f15ff515c4a60e2c9c1d63f08b95c532c810d48fde287c4671df73320.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ib7e8a57a233b8605c617878983f83774822a84e0787f7f7c45204228ecf7e8b3.FindRoomListsRequestBuilder) {
    return ib7e8a57a233b8605c617878983f83774822a84e0787f7f7c45204228ecf7e8b3.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i6d222ff92d14ce4110b7e01d31b1f087405b9c8ace22bac5d8c0bc12bb40c9cc.FindRoomsRequestBuilder) {
    return i6d222ff92d14ce4110b7e01d31b1f087405b9c8ace22bac5d8c0bc12bb40c9cc.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i0b39097657fbe88b406fc1b6eb76d967394330fbd70d43c5b00584005d5245d0.FindRoomsWithRoomListRequestBuilder) {
    return i0b39097657fbe88b406fc1b6eb76d967394330fbd70d43c5b00584005d5245d0.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i35063fd7f1ad0589fc94fb3b6709383ece8aa11f4dc0f9b8e14dc5a5746554d0.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i35063fd7f1ad0589fc94fb3b6709383ece8aa11f4dc0f9b8e14dc5a5746554d0.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i482737223efe8dc02b6546ed734fc77004eff46fb9e1e8c65528db4fa3dbba52.GetLoggedOnManagedDevicesRequestBuilder) {
    return i482737223efe8dc02b6546ed734fc77004eff46fb9e1e8c65528db4fa3dbba52.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i682bc70d0442cfa8e2a3e10c18a413d44877bb859a973c8257c3a379982af4ef.GetMailTipsRequestBuilder) {
    return i682bc70d0442cfa8e2a3e10c18a413d44877bb859a973c8257c3a379982af4ef.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i0e53131874c90de60c7fc2ae1935b28d78cecf24d50a5f0ecd11f421e3e202e5.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i0e53131874c90de60c7fc2ae1935b28d78cecf24d50a5f0ecd11f421e3e202e5.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*idcb80e15afddfb892c1a90f7152accaa4c561738893f238c3bf92169544794b5.GetManagedAppPoliciesRequestBuilder) {
    return idcb80e15afddfb892c1a90f7152accaa4c561738893f238c3bf92169544794b5.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i1f0a5fa96b97d9bca8d455835fa1c564151dc00f916c7d2f987b36a0a241304e.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i1f0a5fa96b97d9bca8d455835fa1c564151dc00f916c7d2f987b36a0a241304e.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7bb7af3e4ffdf11994b91bc232d949fe47a72edc5937c09c9d09dd8d74fc66a1.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7bb7af3e4ffdf11994b91bc232d949fe47a72edc5937c09c9d09dd8d74fc66a1.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i8207065a4c07d73ecae8dfadff355be778d398b00fd5869960b00c65760d1997.GetMemberGroupsRequestBuilder) {
    return i8207065a4c07d73ecae8dfadff355be778d398b00fd5869960b00c65760d1997.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i2ece9e43d66dc3566a1b48c72063e10ae0b31ad2ab5d7889f13949f71d79b471.GetMemberObjectsRequestBuilder) {
    return i2ece9e43d66dc3566a1b48c72063e10ae0b31ad2ab5d7889f13949f71d79b471.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*id64f5a6d9f2c0b8f62bc375576c508d23e1a0f114c48846e32bbb0573031ccda.InvalidateAllRefreshTokensRequestBuilder) {
    return id64f5a6d9f2c0b8f62bc375576c508d23e1a0f114c48846e32bbb0573031ccda.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ib8e88fe7e3688aca9e2b0dca5ad5851ff2acd35c2c4d81999a643ac27d1390a0.IsManagedAppUserBlockedRequestBuilder) {
    return ib8e88fe7e3688aca9e2b0dca5ad5851ff2acd35c2c4d81999a643ac27d1390a0.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i26b9051a79724e72b9a2512bbd3aabdce9071c3deaf2e48487a7a49e4c2d4f89.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i26b9051a79724e72b9a2512bbd3aabdce9071c3deaf2e48487a7a49e4c2d4f89.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i8606a6d0bde4527a5685ef9ce8fae4475d25fb6171ef9169a1d9933f72ad47ff.RemoveAllDevicesFromManagementRequestBuilder) {
    return i8606a6d0bde4527a5685ef9ce8fae4475d25fb6171ef9169a1d9933f72ad47ff.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i98dc5c151ff6f7d85030e85bc231e1bac7f07c0d3199c93283033d77882ca17f.ReprocessLicenseAssignmentRequestBuilder) {
    return i98dc5c151ff6f7d85030e85bc231e1bac7f07c0d3199c93283033d77882ca17f.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i5166400e837b94070d39d0f014dda18b1079fa90623de5748871e244389f04a2.RestoreRequestBuilder) {
    return i5166400e837b94070d39d0f014dda18b1079fa90623de5748871e244389f04a2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i6d64fea2a0b0e1895f236262a212af51360e9416e3308a614e50a52768510ba9.RevokeSignInSessionsRequestBuilder) {
    return i6d64fea2a0b0e1895f236262a212af51360e9416e3308a614e50a52768510ba9.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i67da788c4d3c4b7fd02d651252e0de29864355b61b8c77e337fd9b0c66036bac.SendMailRequestBuilder) {
    return i67da788c4d3c4b7fd02d651252e0de29864355b61b8c77e337fd9b0c66036bac.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i7421f569518b6d911f85f6ad0d9c80a510968d9b4dbc1dfaa6a3de58f7e01400.TranslateExchangeIdsRequestBuilder) {
    return i7421f569518b6d911f85f6ad0d9c80a510968d9b4dbc1dfaa6a3de58f7e01400.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i52ab2e51018f1b3efc8f533ba6d360c21a1ef57a47fca077007dabf7a112fbbe.UnblockManagedAppsRequestBuilder) {
    return i52ab2e51018f1b3efc8f533ba6d360c21a1ef57a47fca077007dabf7a112fbbe.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i6c9b18facb8fad827fcbdbe230f3745e18ceab8d91b9a422aa6ab5c6d9ad8f27.WipeAndBlockManagedAppsRequestBuilder) {
    return i6c9b18facb8fad827fcbdbe230f3745e18ceab8d91b9a422aa6ab5c6d9ad8f27.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i4da4b388440a3c0081b0e338987c597f5acfd63087113a16eda8a9d6470208e8.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i4da4b388440a3c0081b0e338987c597f5acfd63087113a16eda8a9d6470208e8.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*idd3ace18bfe75607474e31976e474dd5fad5f83a5a77cd6040011c821bdad754.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return idd3ace18bfe75607474e31976e474dd5fad5f83a5a77cd6040011c821bdad754.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*if907b6d9d4480b39bd0ab57a7686b150c27544885e14f4729604a03ea4e727ab.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return if907b6d9d4480b39bd0ab57a7686b150c27544885e14f4729604a03ea4e727ab.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
