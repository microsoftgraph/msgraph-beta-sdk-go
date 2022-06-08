package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08d78bbb509dccfabaa4a1981c2b22d3a6613788995fa3d379c9ba717933a2a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/changepassword"
    i0f906ba1af6fec67a49fd01f35e4f5698b9f35e24cbe8038412c4c96c5c9e5c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i116b0a4bea9f2b005dd0fb976c53bfbbf4a87606cbe6f5b141e7101d7e31480f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/findmeetingtimes"
    i14a2aed56e0c80a2646707f11a39f88160dc01fa00d6270c2737a0f5dac5cefe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/ismanagedappuserblocked"
    i15b36031b62152e68d962692af422fac609fc305087bab875cd1c28a1c529b1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i21c08d22e424e5764463cdb928e9b4e6d3979c7c07f9f4da05a13ee1bd702c3f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/translateexchangeids"
    i2595e49cb5cf9890fb30eeabe168ed85b177a84e772539267abdafed2b41b548 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/removealldevicesfrommanagement"
    i2de631e1b07958afe1d1eda5bf9ed0f52b09a565d29f4d0c87fcf843d435a3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/revokesigninsessions"
    i346d1cb8fc17a0107bd5a3e6e9b1ed1565662b690bb3c54179a745cfdc07e9a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmembergroups"
    i3c6071cd8369e036fa070776d73e87ca39d4a87d40e307897b504a12c3e44121 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/findroomlists"
    i3f163cfb09590e5da53bf1255e9143b0912aacbca7fac02476ee78cf59dd4910 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/exportpersonaldata"
    i57ea816aaf36534e1eb1d164ad67d05ac020206a665672df5c0a02584546d217 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/unblockmanagedapps"
    i5cb2211e6a1d4a85e3dcec303172e1bea12138e8289bdee1ef0bbf072dd63056 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/findroomswithroomlist"
    i600847c6dcbae5f445c6c17415504237feb51f06f1e51caef140ffb86ae21fbe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmemberobjects"
    i61082cc74ba04790e9cd8687759343236d5d3070aa507c043405add022ebed49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/checkmemberobjects"
    i6249d0159b115dc0a30326b214498d865593d2b15290d49d5728b7b7cbd9b508 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i669a65dfba359e4236d9f8869558e5d7ca8e835e42be5aec67f62cbd6de57973 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getloggedonmanageddevices"
    i6e233063ab5075a57291ebc43bf426bbfb1cb1528c761db75ed1860d97fbf8c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/wipeandblockmanagedapps"
    i6e40fc4ff902aec1ce77b9fb98beb7c1e5cebbc140852ec2d8fb8e526d115238 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmanagedapppolicies"
    i70d40c34e17f3c6c142668f12b717672e119107f2ba2db2e20c857b9b5f34fa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i7e762a0efadef8950544f1cdf26e63382a2430eccb80ece1cd3421d0d5a1373b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i87134ad961b4273a1e750b793af91b1653c1813f5eda54974e69df7055a6041e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmailtips"
    i8f5b8e09dda85a36321442649169e860a4ebc76433a188e4ead7136016076264 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/sendmail"
    i97ade316e25985462a308c43b418d1b6c67f763afcf3824631232e8e6750926c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/exportdeviceandappmanagementdata"
    i99a90ef7387dde71a63a337643a1591c25c83fba3086bfa36ae1bf28c68e8b4a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/restore"
    i9b794a9b39b4286a18b8c99b67a7fd123ee3fe9796d0ac78dd2be21aab2853d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    ia14e35ff132ae85af06be1e30988abc6f8d45620813e4e8390cb3d17c6af9094 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ia3581132dfa4bf4dfc193542fa7e59c6040812843365984ac611b1fa17d4a05c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ia725c636b6102fd3eed6976b3f17e8a4713474d81563aa9ce02e0ad085adb88f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/checkmembergroups"
    iba524bcb80c5114e090948462fae294caee9617228c9833ffe2ed5f5991e246c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/findrooms"
    icecefb7e3b1d189eb0e9d5c0aff0d2dd87179f257e27af1881e2f2a1da489a80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/assignlicense"
    ie611a1a92b68ed8ee07e817471df24fa062e760f6469a04ae822ff4dc4cb6ef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/activateserviceplan"
    iebdf725787d634ab739590777fff48c7ac956f6a112c2bcafd19b5549bac9d80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/invalidateallrefreshtokens"
    if26c68fbe1b7416ee8f90624c493d7f77884284c9ef165ae7b121e0430c57afd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/reprocesslicenseassignment"
    iff55b4d570a2846cec9f514b5c1a7cd373e4d94420444bf3bbc97040b06e2205 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/item/user/getmanageddeviceswithappfailures"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ie611a1a92b68ed8ee07e817471df24fa062e760f6469a04ae822ff4dc4cb6ef9.ActivateServicePlanRequestBuilder) {
    return ie611a1a92b68ed8ee07e817471df24fa062e760f6469a04ae822ff4dc4cb6ef9.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*icecefb7e3b1d189eb0e9d5c0aff0d2dd87179f257e27af1881e2f2a1da489a80.AssignLicenseRequestBuilder) {
    return icecefb7e3b1d189eb0e9d5c0aff0d2dd87179f257e27af1881e2f2a1da489a80.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i08d78bbb509dccfabaa4a1981c2b22d3a6613788995fa3d379c9ba717933a2a9.ChangePasswordRequestBuilder) {
    return i08d78bbb509dccfabaa4a1981c2b22d3a6613788995fa3d379c9ba717933a2a9.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ia725c636b6102fd3eed6976b3f17e8a4713474d81563aa9ce02e0ad085adb88f.CheckMemberGroupsRequestBuilder) {
    return ia725c636b6102fd3eed6976b3f17e8a4713474d81563aa9ce02e0ad085adb88f.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i61082cc74ba04790e9cd8687759343236d5d3070aa507c043405add022ebed49.CheckMemberObjectsRequestBuilder) {
    return i61082cc74ba04790e9cd8687759343236d5d3070aa507c043405add022ebed49.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i97ade316e25985462a308c43b418d1b6c67f763afcf3824631232e8e6750926c.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i97ade316e25985462a308c43b418d1b6c67f763afcf3824631232e8e6750926c.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i0f906ba1af6fec67a49fd01f35e4f5698b9f35e24cbe8038412c4c96c5c9e5c4.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i0f906ba1af6fec67a49fd01f35e4f5698b9f35e24cbe8038412c4c96c5c9e5c4.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i3f163cfb09590e5da53bf1255e9143b0912aacbca7fac02476ee78cf59dd4910.ExportPersonalDataRequestBuilder) {
    return i3f163cfb09590e5da53bf1255e9143b0912aacbca7fac02476ee78cf59dd4910.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i116b0a4bea9f2b005dd0fb976c53bfbbf4a87606cbe6f5b141e7101d7e31480f.FindMeetingTimesRequestBuilder) {
    return i116b0a4bea9f2b005dd0fb976c53bfbbf4a87606cbe6f5b141e7101d7e31480f.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i3c6071cd8369e036fa070776d73e87ca39d4a87d40e307897b504a12c3e44121.FindRoomListsRequestBuilder) {
    return i3c6071cd8369e036fa070776d73e87ca39d4a87d40e307897b504a12c3e44121.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*iba524bcb80c5114e090948462fae294caee9617228c9833ffe2ed5f5991e246c.FindRoomsRequestBuilder) {
    return iba524bcb80c5114e090948462fae294caee9617228c9833ffe2ed5f5991e246c.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i5cb2211e6a1d4a85e3dcec303172e1bea12138e8289bdee1ef0bbf072dd63056.FindRoomsWithRoomListRequestBuilder) {
    return i5cb2211e6a1d4a85e3dcec303172e1bea12138e8289bdee1ef0bbf072dd63056.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ia3581132dfa4bf4dfc193542fa7e59c6040812843365984ac611b1fa17d4a05c.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ia3581132dfa4bf4dfc193542fa7e59c6040812843365984ac611b1fa17d4a05c.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i669a65dfba359e4236d9f8869558e5d7ca8e835e42be5aec67f62cbd6de57973.GetLoggedOnManagedDevicesRequestBuilder) {
    return i669a65dfba359e4236d9f8869558e5d7ca8e835e42be5aec67f62cbd6de57973.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i87134ad961b4273a1e750b793af91b1653c1813f5eda54974e69df7055a6041e.GetMailTipsRequestBuilder) {
    return i87134ad961b4273a1e750b793af91b1653c1813f5eda54974e69df7055a6041e.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i6249d0159b115dc0a30326b214498d865593d2b15290d49d5728b7b7cbd9b508.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i6249d0159b115dc0a30326b214498d865593d2b15290d49d5728b7b7cbd9b508.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i6e40fc4ff902aec1ce77b9fb98beb7c1e5cebbc140852ec2d8fb8e526d115238.GetManagedAppPoliciesRequestBuilder) {
    return i6e40fc4ff902aec1ce77b9fb98beb7c1e5cebbc140852ec2d8fb8e526d115238.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*iff55b4d570a2846cec9f514b5c1a7cd373e4d94420444bf3bbc97040b06e2205.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return iff55b4d570a2846cec9f514b5c1a7cd373e4d94420444bf3bbc97040b06e2205.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i70d40c34e17f3c6c142668f12b717672e119107f2ba2db2e20c857b9b5f34fa2.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i70d40c34e17f3c6c142668f12b717672e119107f2ba2db2e20c857b9b5f34fa2.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i346d1cb8fc17a0107bd5a3e6e9b1ed1565662b690bb3c54179a745cfdc07e9a8.GetMemberGroupsRequestBuilder) {
    return i346d1cb8fc17a0107bd5a3e6e9b1ed1565662b690bb3c54179a745cfdc07e9a8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i600847c6dcbae5f445c6c17415504237feb51f06f1e51caef140ffb86ae21fbe.GetMemberObjectsRequestBuilder) {
    return i600847c6dcbae5f445c6c17415504237feb51f06f1e51caef140ffb86ae21fbe.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*iebdf725787d634ab739590777fff48c7ac956f6a112c2bcafd19b5549bac9d80.InvalidateAllRefreshTokensRequestBuilder) {
    return iebdf725787d634ab739590777fff48c7ac956f6a112c2bcafd19b5549bac9d80.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i14a2aed56e0c80a2646707f11a39f88160dc01fa00d6270c2737a0f5dac5cefe.IsManagedAppUserBlockedRequestBuilder) {
    return i14a2aed56e0c80a2646707f11a39f88160dc01fa00d6270c2737a0f5dac5cefe.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i15b36031b62152e68d962692af422fac609fc305087bab875cd1c28a1c529b1c.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i15b36031b62152e68d962692af422fac609fc305087bab875cd1c28a1c529b1c.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i2595e49cb5cf9890fb30eeabe168ed85b177a84e772539267abdafed2b41b548.RemoveAllDevicesFromManagementRequestBuilder) {
    return i2595e49cb5cf9890fb30eeabe168ed85b177a84e772539267abdafed2b41b548.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*if26c68fbe1b7416ee8f90624c493d7f77884284c9ef165ae7b121e0430c57afd.ReprocessLicenseAssignmentRequestBuilder) {
    return if26c68fbe1b7416ee8f90624c493d7f77884284c9ef165ae7b121e0430c57afd.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i99a90ef7387dde71a63a337643a1591c25c83fba3086bfa36ae1bf28c68e8b4a.RestoreRequestBuilder) {
    return i99a90ef7387dde71a63a337643a1591c25c83fba3086bfa36ae1bf28c68e8b4a.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i2de631e1b07958afe1d1eda5bf9ed0f52b09a565d29f4d0c87fcf843d435a3d8.RevokeSignInSessionsRequestBuilder) {
    return i2de631e1b07958afe1d1eda5bf9ed0f52b09a565d29f4d0c87fcf843d435a3d8.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i8f5b8e09dda85a36321442649169e860a4ebc76433a188e4ead7136016076264.SendMailRequestBuilder) {
    return i8f5b8e09dda85a36321442649169e860a4ebc76433a188e4ead7136016076264.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i21c08d22e424e5764463cdb928e9b4e6d3979c7c07f9f4da05a13ee1bd702c3f.TranslateExchangeIdsRequestBuilder) {
    return i21c08d22e424e5764463cdb928e9b4e6d3979c7c07f9f4da05a13ee1bd702c3f.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i57ea816aaf36534e1eb1d164ad67d05ac020206a665672df5c0a02584546d217.UnblockManagedAppsRequestBuilder) {
    return i57ea816aaf36534e1eb1d164ad67d05ac020206a665672df5c0a02584546d217.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i6e233063ab5075a57291ebc43bf426bbfb1cb1528c761db75ed1860d97fbf8c7.WipeAndBlockManagedAppsRequestBuilder) {
    return i6e233063ab5075a57291ebc43bf426bbfb1cb1528c761db75ed1860d97fbf8c7.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i7e762a0efadef8950544f1cdf26e63382a2430eccb80ece1cd3421d0d5a1373b.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i7e762a0efadef8950544f1cdf26e63382a2430eccb80ece1cd3421d0d5a1373b.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ia14e35ff132ae85af06be1e30988abc6f8d45620813e4e8390cb3d17c6af9094.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ia14e35ff132ae85af06be1e30988abc6f8d45620813e4e8390cb3d17c6af9094.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i9b794a9b39b4286a18b8c99b67a7fd123ee3fe9796d0ac78dd2be21aab2853d7.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i9b794a9b39b4286a18b8c99b67a7fd123ee3fe9796d0ac78dd2be21aab2853d7.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
