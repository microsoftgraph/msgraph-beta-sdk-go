package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f32681b700fef2377ae14ea41d61216060d1adac48311cd377c00d7ff2f2331 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    i1258eb2d82a8aead77a046db96b06c3821284a475f09fcec37b5fbfcc76693ef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/ismanagedappuserblocked"
    i1390368fb08c8a1e1cc067c0a14459401ae820ebb24c80c65bdfae9cec7dfd17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i1ac554526dc49c723f9e6e44a2608aa1b70469a9f134debfd53458a1a70a37c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i245e16de11b91736d9869b2a59775583577fb06b94aaf900e9b861e8c9c52f8b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i2bd6a5d0cc1816774d27e032ea36c2b2a0ec110636c4a1e28654c0b786df83ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/restore"
    i2d9e12440bf8d77cf886578763b50603cc356605e5f62ddb4c588ef62139edcb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/sendmail"
    i38e8d5047247e354121209db6319a6d80e3e4f49aa29dfad664dda603a5bd8a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    i3fb1864fa5d4a5ad6b5191b1d20a5281cb33acc98246428a62fc0ad81666ac27 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    i43c174d0bfa7813a011c6cf05edd8ed635149478b84ddab7c80ff3179764fdf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    i443f676ff8dcb1ab136958f6876b15987b9cd66ae7bb92e90ba658b293f40467 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i446876bc2744afeb0ad82e3de4967b278cf840bbe160945769f26c2bede56414 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    i4e624d7ff10efb516a9cb7e8f44b4935343033143c17690d833218a0a7b08485 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    i69dddea1d2627e67b872d1c88b3555fd78e364c7d5a03d362395e79235f07dab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findroomlists"
    i6beb4778f97630b96dcce0d56f2b6294fb5b59c4e2abf093928a72434e61102c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findroomswithroomlist"
    i6f3130b7db67044fca0d7a22b113dea87c4a5c3c102246a6fd534d5c0377bee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    i7315c2f4d1d7a77aed5f6c90579caec3b7a472a6707b5b4f3cee429a28824b2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/unblockmanagedapps"
    i7bc3b41a467188e9521fcd3844f63ac98f17b343ea5f3cda4ebb3ae6f4fbc171 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i8bac059900bae33d364d78c9002bcb682f795407f4b106ff363a589309f0d5c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/activateserviceplan"
    i932ed76964e66d09920c60d2d3cae0f1e6766322806e21941acaf7d7b0428296 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i989e3bf7af67def0417feb8af23849df722c69c39b9e6c893f579e1f0c03f939 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    i992e0ef3be27332bb5729b466f93a4634fa30d4d361761b1bacddf05d6093963 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/changepassword"
    ia9ef3231e428c551f2810ffda96245629ef45f5d2a93cd5f120e0b481039ad0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    ib9216d86792a06b83ee306c0d7d5848afeb6e174008995e28a649d18520dab39 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    ibbc455632f0a481e714ea3f933fc6ce5d9bf7ea7156f514e73b182ffc7151ccf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
    id643479b6040ce66878ce6927a4c7105ff14630140a78d727fb1aa1e57fdea09 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    ida7aa9d6d2fdde8e9fe75c23d091691ecafc103c8966268eda00a85821646396 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    idc87f22a5838ed784d8bccd5d1dd3f9a27a71e133c24b7c27d09f2116dc60865 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/findrooms"
    ie2a9fc4944d6c98c3cf7b9a7a69ef8c58db383556811c7f4d5668dde1448839f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    ie539c20be5898c4091f1b361722b8e42f4e8a1b72de0d6a92ca4a837a9932811 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    ie63a6db75760914eb033447368f3a873e410c7351fd5b11dd03b93c0f28b9ed0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    ied4088f36e9f1a79d1aad1fe8b6619e1b171405b9e5910cdaf354b0f6f48d5ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    if53571c0d5eb89d7b8d686b3bd508dfe78da0769c6e5e2b203041d2aeca67da4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/getloggedonmanageddevices"
    ifc3055b54870f547db20ee6e4fbb6d0a040e99b951ee690e5b74ed80d1e1419a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    ifc90f5625598ae54567ce648a81674e7f70b0301d19d7a49fb0c4ec70e4340cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i8bac059900bae33d364d78c9002bcb682f795407f4b106ff363a589309f0d5c6.ActivateServicePlanRequestBuilder) {
    return i8bac059900bae33d364d78c9002bcb682f795407f4b106ff363a589309f0d5c6.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ie539c20be5898c4091f1b361722b8e42f4e8a1b72de0d6a92ca4a837a9932811.AssignLicenseRequestBuilder) {
    return ie539c20be5898c4091f1b361722b8e42f4e8a1b72de0d6a92ca4a837a9932811.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i992e0ef3be27332bb5729b466f93a4634fa30d4d361761b1bacddf05d6093963.ChangePasswordRequestBuilder) {
    return i992e0ef3be27332bb5729b466f93a4634fa30d4d361761b1bacddf05d6093963.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i245e16de11b91736d9869b2a59775583577fb06b94aaf900e9b861e8c9c52f8b.CheckMemberGroupsRequestBuilder) {
    return i245e16de11b91736d9869b2a59775583577fb06b94aaf900e9b861e8c9c52f8b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i1390368fb08c8a1e1cc067c0a14459401ae820ebb24c80c65bdfae9cec7dfd17.CheckMemberObjectsRequestBuilder) {
    return i1390368fb08c8a1e1cc067c0a14459401ae820ebb24c80c65bdfae9cec7dfd17.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i38e8d5047247e354121209db6319a6d80e3e4f49aa29dfad664dda603a5bd8a2.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i38e8d5047247e354121209db6319a6d80e3e4f49aa29dfad664dda603a5bd8a2.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ifc90f5625598ae54567ce648a81674e7f70b0301d19d7a49fb0c4ec70e4340cc.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ifc90f5625598ae54567ce648a81674e7f70b0301d19d7a49fb0c4ec70e4340cc.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i989e3bf7af67def0417feb8af23849df722c69c39b9e6c893f579e1f0c03f939.ExportPersonalDataRequestBuilder) {
    return i989e3bf7af67def0417feb8af23849df722c69c39b9e6c893f579e1f0c03f939.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i6f3130b7db67044fca0d7a22b113dea87c4a5c3c102246a6fd534d5c0377bee8.FindMeetingTimesRequestBuilder) {
    return i6f3130b7db67044fca0d7a22b113dea87c4a5c3c102246a6fd534d5c0377bee8.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i69dddea1d2627e67b872d1c88b3555fd78e364c7d5a03d362395e79235f07dab.FindRoomListsRequestBuilder) {
    return i69dddea1d2627e67b872d1c88b3555fd78e364c7d5a03d362395e79235f07dab.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*idc87f22a5838ed784d8bccd5d1dd3f9a27a71e133c24b7c27d09f2116dc60865.FindRoomsRequestBuilder) {
    return idc87f22a5838ed784d8bccd5d1dd3f9a27a71e133c24b7c27d09f2116dc60865.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i6beb4778f97630b96dcce0d56f2b6294fb5b59c4e2abf093928a72434e61102c.FindRoomsWithRoomListRequestBuilder) {
    return i6beb4778f97630b96dcce0d56f2b6294fb5b59c4e2abf093928a72434e61102c.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ida7aa9d6d2fdde8e9fe75c23d091691ecafc103c8966268eda00a85821646396.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ida7aa9d6d2fdde8e9fe75c23d091691ecafc103c8966268eda00a85821646396.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*if53571c0d5eb89d7b8d686b3bd508dfe78da0769c6e5e2b203041d2aeca67da4.GetLoggedOnManagedDevicesRequestBuilder) {
    return if53571c0d5eb89d7b8d686b3bd508dfe78da0769c6e5e2b203041d2aeca67da4.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i446876bc2744afeb0ad82e3de4967b278cf840bbe160945769f26c2bede56414.GetMailTipsRequestBuilder) {
    return i446876bc2744afeb0ad82e3de4967b278cf840bbe160945769f26c2bede56414.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i3fb1864fa5d4a5ad6b5191b1d20a5281cb33acc98246428a62fc0ad81666ac27.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i3fb1864fa5d4a5ad6b5191b1d20a5281cb33acc98246428a62fc0ad81666ac27.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i0f32681b700fef2377ae14ea41d61216060d1adac48311cd377c00d7ff2f2331.GetManagedAppPoliciesRequestBuilder) {
    return i0f32681b700fef2377ae14ea41d61216060d1adac48311cd377c00d7ff2f2331.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ib9216d86792a06b83ee306c0d7d5848afeb6e174008995e28a649d18520dab39.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ib9216d86792a06b83ee306c0d7d5848afeb6e174008995e28a649d18520dab39.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7bc3b41a467188e9521fcd3844f63ac98f17b343ea5f3cda4ebb3ae6f4fbc171.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7bc3b41a467188e9521fcd3844f63ac98f17b343ea5f3cda4ebb3ae6f4fbc171.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i1ac554526dc49c723f9e6e44a2608aa1b70469a9f134debfd53458a1a70a37c8.GetMemberGroupsRequestBuilder) {
    return i1ac554526dc49c723f9e6e44a2608aa1b70469a9f134debfd53458a1a70a37c8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i4e624d7ff10efb516a9cb7e8f44b4935343033143c17690d833218a0a7b08485.GetMemberObjectsRequestBuilder) {
    return i4e624d7ff10efb516a9cb7e8f44b4935343033143c17690d833218a0a7b08485.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ie63a6db75760914eb033447368f3a873e410c7351fd5b11dd03b93c0f28b9ed0.InvalidateAllRefreshTokensRequestBuilder) {
    return ie63a6db75760914eb033447368f3a873e410c7351fd5b11dd03b93c0f28b9ed0.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1258eb2d82a8aead77a046db96b06c3821284a475f09fcec37b5fbfcc76693ef.IsManagedAppUserBlockedRequestBuilder) {
    return i1258eb2d82a8aead77a046db96b06c3821284a475f09fcec37b5fbfcc76693ef.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i932ed76964e66d09920c60d2d3cae0f1e6766322806e21941acaf7d7b0428296.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i932ed76964e66d09920c60d2d3cae0f1e6766322806e21941acaf7d7b0428296.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id643479b6040ce66878ce6927a4c7105ff14630140a78d727fb1aa1e57fdea09.RemoveAllDevicesFromManagementRequestBuilder) {
    return id643479b6040ce66878ce6927a4c7105ff14630140a78d727fb1aa1e57fdea09.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ied4088f36e9f1a79d1aad1fe8b6619e1b171405b9e5910cdaf354b0f6f48d5ff.ReprocessLicenseAssignmentRequestBuilder) {
    return ied4088f36e9f1a79d1aad1fe8b6619e1b171405b9e5910cdaf354b0f6f48d5ff.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i2bd6a5d0cc1816774d27e032ea36c2b2a0ec110636c4a1e28654c0b786df83ce.RestoreRequestBuilder) {
    return i2bd6a5d0cc1816774d27e032ea36c2b2a0ec110636c4a1e28654c0b786df83ce.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i43c174d0bfa7813a011c6cf05edd8ed635149478b84ddab7c80ff3179764fdf9.RevokeSignInSessionsRequestBuilder) {
    return i43c174d0bfa7813a011c6cf05edd8ed635149478b84ddab7c80ff3179764fdf9.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i2d9e12440bf8d77cf886578763b50603cc356605e5f62ddb4c588ef62139edcb.SendMailRequestBuilder) {
    return i2d9e12440bf8d77cf886578763b50603cc356605e5f62ddb4c588ef62139edcb.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ibbc455632f0a481e714ea3f933fc6ce5d9bf7ea7156f514e73b182ffc7151ccf.TranslateExchangeIdsRequestBuilder) {
    return ibbc455632f0a481e714ea3f933fc6ce5d9bf7ea7156f514e73b182ffc7151ccf.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i7315c2f4d1d7a77aed5f6c90579caec3b7a472a6707b5b4f3cee429a28824b2e.UnblockManagedAppsRequestBuilder) {
    return i7315c2f4d1d7a77aed5f6c90579caec3b7a472a6707b5b4f3cee429a28824b2e.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ie2a9fc4944d6c98c3cf7b9a7a69ef8c58db383556811c7f4d5668dde1448839f.WipeAndBlockManagedAppsRequestBuilder) {
    return ie2a9fc4944d6c98c3cf7b9a7a69ef8c58db383556811c7f4d5668dde1448839f.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ia9ef3231e428c551f2810ffda96245629ef45f5d2a93cd5f120e0b481039ad0c.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ia9ef3231e428c551f2810ffda96245629ef45f5d2a93cd5f120e0b481039ad0c.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i443f676ff8dcb1ab136958f6876b15987b9cd66ae7bb92e90ba658b293f40467.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i443f676ff8dcb1ab136958f6876b15987b9cd66ae7bb92e90ba658b293f40467.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ifc3055b54870f547db20ee6e4fbb6d0a040e99b951ee690e5b74ed80d1e1419a.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ifc3055b54870f547db20ee6e4fbb6d0a040e99b951ee690e5b74ed80d1e1419a.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
