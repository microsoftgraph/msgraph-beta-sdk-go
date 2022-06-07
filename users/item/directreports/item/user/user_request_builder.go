package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03d991a20be82abccd631171c0d62677a35b3a32e62348076910e8f2ccf3194a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/findroomswithroomlist"
    i0a489f3c3481228010582a35eff8e34a03c848e49dc8d9c71f812db6e58c5fc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/findroomlists"
    i0c5f071eb9427c86e6fa56d82271fe14dc8786e7dbd8d78a85ed3440796a05b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/checkmemberobjects"
    i0fbe0430a13bf6f9cb9055828ea59fa43f5a8d588da8bd78b90e9f546c299701 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/exportdeviceandappmanagementdata"
    i1012aa8a0fdb5e4556374576959ed3144004689ce8df04f55f9a9866b99340a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/geteffectivedeviceenrollmentconfigurations"
    i16fd31fc5bcb01c49d1beaa40cfd0cebe3e3d7c3c8c8301408bfce3151699ca0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/ismanagedappuserblocked"
    i1a3054d15a8d559906fd929418feef79c3dda90334dfb9fbab4ee08a86cbd932 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/exportpersonaldata"
    i320173a3ccb5ca6414f55631768c3d5f161fe91325c0f95c14eb6afb9365a6a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/wipemanagedappregistrationsbydevicetag"
    i3c2af18b7ba7f406d9c5a07b84d5999316e75a6f7e9132290ffb232c68b9eb52 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmanageddeviceswithappfailures"
    i46518de386837dca1923bd20244d42b56d8ebf8d39113982f28eb86d06c3bb83 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/invalidateallrefreshtokens"
    i4bf764b73da74723564990877904e848f59fe65057d0f06885ebebcfe935dbbd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/findmeetingtimes"
    i512529e1b8c0edd8a4f75a589609a6560bbefa9619333832a16506bc73b5af9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/changepassword"
    i659a4671ee8000ebb9227e5bd4d51e56570b12793f671c067d9e24a3eeeb363a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/wipeandblockmanagedapps"
    i6b7079c755cee67bc673dfa08002a8de4c543e9cb166f99f261bc5eab033906f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmembergroups"
    i6c9f7b0cbb3d8f5c23a1b8d24ee33dd669ce488163446d0e691eb03ebcd0bc28 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmanageddeviceswithfailedorpendingapps"
    i732a69d46177bf3f0c2b08f7958164ae6baf8f62688d60acf66964c2428d489d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/restore"
    i7898dda5e030b677026e60bc8a5f31837d120cef7e7ca0c29d722cc51d8ad525 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i856a1366674c9f5c595824f5e39998da9163040a9248da2a602a7d79f84b4ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/reprocesslicenseassignment"
    i8d1ca96ba5d1dad51e3295547e1f6c910595ed561116c5797524614ea17c7814 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/assignlicense"
    i8dc52b9738cff0faad9aae5edb8590b355d823026263ed6860d39c9a0ef65099 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmailtips"
    i92253d4c5bb6f6dd7c27d0f4544e1117488a04d0457203b5a045528fe6f47ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/sendmail"
    i9373d16b6631dc33c761713ac38bbc2e5c2989f7b5d1f38b64c0d9dda1912519 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i94788952eed9836017a33560188d192e0244faf57aca852cd002e61ed4301d8a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/revokesigninsessions"
    i952a38b0dc8e99880b243706de0e75d226fa5fa066d88974ba9d6f8f69f31c3e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/unblockmanagedapps"
    i9a3abbb1c78a7a90ba9be5c79ccd2133c14feff77a0fc1bbbb0999c5050e6f86 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmanagedapppolicies"
    ia033d43a46309c2f1a2e613bccba1b17b6420991a18cc18708fdb359c34a65e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/findrooms"
    iadf56e1e4daee424a503c079faa4c3997420f5eabfebf525361c559769207909 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/checkmembergroups"
    ib2cafcd87dd5831bcd00001705be44bdbc951ce1a788824179a3b393419db10a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/activateserviceplan"
    ic4bd008b99e263e2854ba1d3455ade21bc4c52f8f4071054e4d38a2108f90cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/reminderviewwithstartdatetimewithenddatetime"
    iccbce095a530ae90b5f280f900d20b53da812ff66ad3d158dffab94b4183b0bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/removealldevicesfrommanagement"
    id2f948615a7f622a4fa861157fa27b23b0e93c4b558a8ff8b07bc9fc61e89053 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getloggedonmanageddevices"
    id65b2a36151100ce67c615e161daa7b1f037c4c78adf3538cc7f0ea90753be34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmemberobjects"
    id9296018bc035fe437bdb050178020ba610c43c0b3df5bbb68d29b28c3e56709 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/getmanagedappdiagnosticstatuses"
    idf441870b4c01ba0a040331da4c5712bbeed7fe04f2be121ae47c3a23fc70321 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/translateexchangeids"
    ifa2dcc5251c923e8f8bf51f493a3f1a327a07403fe6f68c07f8bf59b3a4fea47 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/directreports/item/user/wipemanagedappregistrationbydevicetag"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*ib2cafcd87dd5831bcd00001705be44bdbc951ce1a788824179a3b393419db10a.ActivateServicePlanRequestBuilder) {
    return ib2cafcd87dd5831bcd00001705be44bdbc951ce1a788824179a3b393419db10a.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i8d1ca96ba5d1dad51e3295547e1f6c910595ed561116c5797524614ea17c7814.AssignLicenseRequestBuilder) {
    return i8d1ca96ba5d1dad51e3295547e1f6c910595ed561116c5797524614ea17c7814.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i512529e1b8c0edd8a4f75a589609a6560bbefa9619333832a16506bc73b5af9e.ChangePasswordRequestBuilder) {
    return i512529e1b8c0edd8a4f75a589609a6560bbefa9619333832a16506bc73b5af9e.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*iadf56e1e4daee424a503c079faa4c3997420f5eabfebf525361c559769207909.CheckMemberGroupsRequestBuilder) {
    return iadf56e1e4daee424a503c079faa4c3997420f5eabfebf525361c559769207909.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i0c5f071eb9427c86e6fa56d82271fe14dc8786e7dbd8d78a85ed3440796a05b9.CheckMemberObjectsRequestBuilder) {
    return i0c5f071eb9427c86e6fa56d82271fe14dc8786e7dbd8d78a85ed3440796a05b9.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/directReports/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i0fbe0430a13bf6f9cb9055828ea59fa43f5a8d588da8bd78b90e9f546c299701.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i0fbe0430a13bf6f9cb9055828ea59fa43f5a8d588da8bd78b90e9f546c299701.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i7898dda5e030b677026e60bc8a5f31837d120cef7e7ca0c29d722cc51d8ad525.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i7898dda5e030b677026e60bc8a5f31837d120cef7e7ca0c29d722cc51d8ad525.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i1a3054d15a8d559906fd929418feef79c3dda90334dfb9fbab4ee08a86cbd932.ExportPersonalDataRequestBuilder) {
    return i1a3054d15a8d559906fd929418feef79c3dda90334dfb9fbab4ee08a86cbd932.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i4bf764b73da74723564990877904e848f59fe65057d0f06885ebebcfe935dbbd.FindMeetingTimesRequestBuilder) {
    return i4bf764b73da74723564990877904e848f59fe65057d0f06885ebebcfe935dbbd.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i0a489f3c3481228010582a35eff8e34a03c848e49dc8d9c71f812db6e58c5fc5.FindRoomListsRequestBuilder) {
    return i0a489f3c3481228010582a35eff8e34a03c848e49dc8d9c71f812db6e58c5fc5.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ia033d43a46309c2f1a2e613bccba1b17b6420991a18cc18708fdb359c34a65e1.FindRoomsRequestBuilder) {
    return ia033d43a46309c2f1a2e613bccba1b17b6420991a18cc18708fdb359c34a65e1.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i03d991a20be82abccd631171c0d62677a35b3a32e62348076910e8f2ccf3194a.FindRoomsWithRoomListRequestBuilder) {
    return i03d991a20be82abccd631171c0d62677a35b3a32e62348076910e8f2ccf3194a.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i1012aa8a0fdb5e4556374576959ed3144004689ce8df04f55f9a9866b99340a3.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i1012aa8a0fdb5e4556374576959ed3144004689ce8df04f55f9a9866b99340a3.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*id2f948615a7f622a4fa861157fa27b23b0e93c4b558a8ff8b07bc9fc61e89053.GetLoggedOnManagedDevicesRequestBuilder) {
    return id2f948615a7f622a4fa861157fa27b23b0e93c4b558a8ff8b07bc9fc61e89053.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i8dc52b9738cff0faad9aae5edb8590b355d823026263ed6860d39c9a0ef65099.GetMailTipsRequestBuilder) {
    return i8dc52b9738cff0faad9aae5edb8590b355d823026263ed6860d39c9a0ef65099.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id9296018bc035fe437bdb050178020ba610c43c0b3df5bbb68d29b28c3e56709.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id9296018bc035fe437bdb050178020ba610c43c0b3df5bbb68d29b28c3e56709.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i9a3abbb1c78a7a90ba9be5c79ccd2133c14feff77a0fc1bbbb0999c5050e6f86.GetManagedAppPoliciesRequestBuilder) {
    return i9a3abbb1c78a7a90ba9be5c79ccd2133c14feff77a0fc1bbbb0999c5050e6f86.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i3c2af18b7ba7f406d9c5a07b84d5999316e75a6f7e9132290ffb232c68b9eb52.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i3c2af18b7ba7f406d9c5a07b84d5999316e75a6f7e9132290ffb232c68b9eb52.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i6c9f7b0cbb3d8f5c23a1b8d24ee33dd669ce488163446d0e691eb03ebcd0bc28.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i6c9f7b0cbb3d8f5c23a1b8d24ee33dd669ce488163446d0e691eb03ebcd0bc28.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i6b7079c755cee67bc673dfa08002a8de4c543e9cb166f99f261bc5eab033906f.GetMemberGroupsRequestBuilder) {
    return i6b7079c755cee67bc673dfa08002a8de4c543e9cb166f99f261bc5eab033906f.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*id65b2a36151100ce67c615e161daa7b1f037c4c78adf3538cc7f0ea90753be34.GetMemberObjectsRequestBuilder) {
    return id65b2a36151100ce67c615e161daa7b1f037c4c78adf3538cc7f0ea90753be34.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i46518de386837dca1923bd20244d42b56d8ebf8d39113982f28eb86d06c3bb83.InvalidateAllRefreshTokensRequestBuilder) {
    return i46518de386837dca1923bd20244d42b56d8ebf8d39113982f28eb86d06c3bb83.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i16fd31fc5bcb01c49d1beaa40cfd0cebe3e3d7c3c8c8301408bfce3151699ca0.IsManagedAppUserBlockedRequestBuilder) {
    return i16fd31fc5bcb01c49d1beaa40cfd0cebe3e3d7c3c8c8301408bfce3151699ca0.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ic4bd008b99e263e2854ba1d3455ade21bc4c52f8f4071054e4d38a2108f90cdc.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ic4bd008b99e263e2854ba1d3455ade21bc4c52f8f4071054e4d38a2108f90cdc.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*iccbce095a530ae90b5f280f900d20b53da812ff66ad3d158dffab94b4183b0bc.RemoveAllDevicesFromManagementRequestBuilder) {
    return iccbce095a530ae90b5f280f900d20b53da812ff66ad3d158dffab94b4183b0bc.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i856a1366674c9f5c595824f5e39998da9163040a9248da2a602a7d79f84b4ddd.ReprocessLicenseAssignmentRequestBuilder) {
    return i856a1366674c9f5c595824f5e39998da9163040a9248da2a602a7d79f84b4ddd.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i732a69d46177bf3f0c2b08f7958164ae6baf8f62688d60acf66964c2428d489d.RestoreRequestBuilder) {
    return i732a69d46177bf3f0c2b08f7958164ae6baf8f62688d60acf66964c2428d489d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i94788952eed9836017a33560188d192e0244faf57aca852cd002e61ed4301d8a.RevokeSignInSessionsRequestBuilder) {
    return i94788952eed9836017a33560188d192e0244faf57aca852cd002e61ed4301d8a.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i92253d4c5bb6f6dd7c27d0f4544e1117488a04d0457203b5a045528fe6f47ea2.SendMailRequestBuilder) {
    return i92253d4c5bb6f6dd7c27d0f4544e1117488a04d0457203b5a045528fe6f47ea2.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*idf441870b4c01ba0a040331da4c5712bbeed7fe04f2be121ae47c3a23fc70321.TranslateExchangeIdsRequestBuilder) {
    return idf441870b4c01ba0a040331da4c5712bbeed7fe04f2be121ae47c3a23fc70321.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i952a38b0dc8e99880b243706de0e75d226fa5fa066d88974ba9d6f8f69f31c3e.UnblockManagedAppsRequestBuilder) {
    return i952a38b0dc8e99880b243706de0e75d226fa5fa066d88974ba9d6f8f69f31c3e.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i659a4671ee8000ebb9227e5bd4d51e56570b12793f671c067d9e24a3eeeb363a.WipeAndBlockManagedAppsRequestBuilder) {
    return i659a4671ee8000ebb9227e5bd4d51e56570b12793f671c067d9e24a3eeeb363a.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ifa2dcc5251c923e8f8bf51f493a3f1a327a07403fe6f68c07f8bf59b3a4fea47.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ifa2dcc5251c923e8f8bf51f493a3f1a327a07403fe6f68c07f8bf59b3a4fea47.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i9373d16b6631dc33c761713ac38bbc2e5c2989f7b5d1f38b64c0d9dda1912519.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i9373d16b6631dc33c761713ac38bbc2e5c2989f7b5d1f38b64c0d9dda1912519.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i320173a3ccb5ca6414f55631768c3d5f161fe91325c0f95c14eb6afb9365a6a4.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i320173a3ccb5ca6414f55631768c3d5f161fe91325c0f95c14eb6afb9365a6a4.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
