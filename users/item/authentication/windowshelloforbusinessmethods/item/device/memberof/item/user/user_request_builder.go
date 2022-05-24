package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0db2740294fa8089841c1072617ea729aef126d6623ad24c59c866656efaf7df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i10b735093bfb492238e268241c08e3bbc23abf518bd6ca7447397fbe86cf94f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/restore"
    i13e4801b242ae19dbfdf6e21da5bf17d0d04fea9b67f029217604cc69b30d6c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmailtips"
    i1b24166e9a5f19517be26bc02aafe609c4f4614e0e1fb3b4bf23b5e82650bf33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    i1ca3e90ce47b30c9b1a10faa1af83062940a12a76d75dff9199f22a815df060c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/revokesigninsessions"
    i3e03205aab7df19b11a5d372d028b0fb8b7cd650446063ba5b762015f01ddddc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmemberobjects"
    i44d0bc774f5245867ccc64328bd33234f807fbb243d6ca9a4e10a38cc4a30453 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/sendmail"
    i5019e0a37cff62bcf150496b2225d8e91e2544a6b4e3d956be7ed39b00ab17c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findroomlists"
    i5b19d64bbfbd56f658c9af1b98d65de363692d6047946c5462ee847a1441280e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findrooms"
    i5f517f6bc8d73b3b81033c06a04afe4b7c0e72187f1a660f83317ba7aa2a4f17 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findroomswithroomlist"
    i64daf8facf298b392228c1f8fbc21378bc0c8dc62e11a0dac1948221187db7e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmembergroups"
    i694cd79009823d2f5a5d13c723bda27e5edb5aa4c4fa83d4d98da872b9edf7d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/findmeetingtimes"
    i6b89b729682111d71f55bb444fea0b59cc3414dcc46e97498bd425ae1f10918a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i6ff7fc0372730c122b4babe8bf28f60a2245e1ece20ea5ca693e16a218d9c6aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/activateserviceplan"
    i7050ca76ec6a387cb08ec1134e681cb098cdf9d2ed3fec0b99a166fd7f2c8291 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmembergroups"
    i70e1b2be4e88dafb6c58959844b65d41c96f72061b2686508eae9ed17be0fc06 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/translateexchangeids"
    i7a198c110253c6692c34999de1fbc8b8f99b7d0aedd7263920e9fad91e31562e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/unblockmanagedapps"
    i7bf9f47041550eaca6a776837ea50de11dc1348704e9508f028484ca9c9c70f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i7e34b02d89b096a5000726e206dbb52a597c7b514a4ac30da1033bbb4b03fdb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedapppolicies"
    i8cbe1311cfb2623bf9606c48ffc155d951b5949948162e0b9340b626f04b8741 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i9463cd567421619a19e5341ef8afb88cfd09f7216694685d46892289d294f0ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/assignlicense"
    i9eaf941207ce5640ef72aa3704ce6223c384274eb30255326fd012e264a406ae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportpersonaldata"
    i9edacdd56843f19c386fe7376a907b5df2c93b768097be253e4b040c8d7f5a04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/ismanagedappuserblocked"
    i9f5d74313adbc1cee70c76cbd815d1ce170fc93f2f12ec13fb590ee91ef88f49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/changepassword"
    ia59cabc97279acd0b39780ca48eeb6dbfdfbaa906cacfe643643c4fca1acc3b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getloggedonmanageddevices"
    ia91e670d6c14be22de7af362c8a7fc0b8fb7d88d7b86fd7a1eeb3ade273221a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    iabb4798ecbab6e7885757a7a867b33c8998c38d62eadd9d68f39583f085f3c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    iae59d6fd2345d27fe78dfc078d36bf84351966f5b7b20dc6d6f8ac82a1cb5312 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ib5e8a06130dd4a8912ddd71ef5920792d70abe52d548ab45f56fd1ddad04b6bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    ic4b6b82e59f2c802f226bee712f014bce52dbda86d896b4159234aa4329fb084 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/checkmemberobjects"
    icf9637e82e732a717f507dfc543264ffd1de23a47b5bf0468fb39b7a883be476 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    id3a3d95622a2b03a996c9c6125c29aab8cff342a443b8b02ee8bb73ea8ee99f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    id5c529f39116c945503ac8bfa0b787bc6d032f884415e02e1315a39891a4edfe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    iddf84c133c3e60b644f39417d4d67ba4abe623d5f365246df5ace24375295581 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    ie4aa4f53c8accb61e9768af991587b52c7bd11a8aee60fc46e3d5d3defe9f79f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i6ff7fc0372730c122b4babe8bf28f60a2245e1ece20ea5ca693e16a218d9c6aa.ActivateServicePlanRequestBuilder) {
    return i6ff7fc0372730c122b4babe8bf28f60a2245e1ece20ea5ca693e16a218d9c6aa.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i9463cd567421619a19e5341ef8afb88cfd09f7216694685d46892289d294f0ed.AssignLicenseRequestBuilder) {
    return i9463cd567421619a19e5341ef8afb88cfd09f7216694685d46892289d294f0ed.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i9f5d74313adbc1cee70c76cbd815d1ce170fc93f2f12ec13fb590ee91ef88f49.ChangePasswordRequestBuilder) {
    return i9f5d74313adbc1cee70c76cbd815d1ce170fc93f2f12ec13fb590ee91ef88f49.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i7050ca76ec6a387cb08ec1134e681cb098cdf9d2ed3fec0b99a166fd7f2c8291.CheckMemberGroupsRequestBuilder) {
    return i7050ca76ec6a387cb08ec1134e681cb098cdf9d2ed3fec0b99a166fd7f2c8291.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ic4b6b82e59f2c802f226bee712f014bce52dbda86d896b4159234aa4329fb084.CheckMemberObjectsRequestBuilder) {
    return ic4b6b82e59f2c802f226bee712f014bce52dbda86d896b4159234aa4329fb084.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*icf9637e82e732a717f507dfc543264ffd1de23a47b5bf0468fb39b7a883be476.ExportDeviceAndAppManagementDataRequestBuilder) {
    return icf9637e82e732a717f507dfc543264ffd1de23a47b5bf0468fb39b7a883be476.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i0db2740294fa8089841c1072617ea729aef126d6623ad24c59c866656efaf7df.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i0db2740294fa8089841c1072617ea729aef126d6623ad24c59c866656efaf7df.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i9eaf941207ce5640ef72aa3704ce6223c384274eb30255326fd012e264a406ae.ExportPersonalDataRequestBuilder) {
    return i9eaf941207ce5640ef72aa3704ce6223c384274eb30255326fd012e264a406ae.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i694cd79009823d2f5a5d13c723bda27e5edb5aa4c4fa83d4d98da872b9edf7d5.FindMeetingTimesRequestBuilder) {
    return i694cd79009823d2f5a5d13c723bda27e5edb5aa4c4fa83d4d98da872b9edf7d5.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i5019e0a37cff62bcf150496b2225d8e91e2544a6b4e3d956be7ed39b00ab17c4.FindRoomListsRequestBuilder) {
    return i5019e0a37cff62bcf150496b2225d8e91e2544a6b4e3d956be7ed39b00ab17c4.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i5b19d64bbfbd56f658c9af1b98d65de363692d6047946c5462ee847a1441280e.FindRoomsRequestBuilder) {
    return i5b19d64bbfbd56f658c9af1b98d65de363692d6047946c5462ee847a1441280e.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i5f517f6bc8d73b3b81033c06a04afe4b7c0e72187f1a660f83317ba7aa2a4f17.FindRoomsWithRoomListRequestBuilder) {
    return i5f517f6bc8d73b3b81033c06a04afe4b7c0e72187f1a660f83317ba7aa2a4f17.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*iabb4798ecbab6e7885757a7a867b33c8998c38d62eadd9d68f39583f085f3c70.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return iabb4798ecbab6e7885757a7a867b33c8998c38d62eadd9d68f39583f085f3c70.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ia59cabc97279acd0b39780ca48eeb6dbfdfbaa906cacfe643643c4fca1acc3b3.GetLoggedOnManagedDevicesRequestBuilder) {
    return ia59cabc97279acd0b39780ca48eeb6dbfdfbaa906cacfe643643c4fca1acc3b3.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i13e4801b242ae19dbfdf6e21da5bf17d0d04fea9b67f029217604cc69b30d6c5.GetMailTipsRequestBuilder) {
    return i13e4801b242ae19dbfdf6e21da5bf17d0d04fea9b67f029217604cc69b30d6c5.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ib5e8a06130dd4a8912ddd71ef5920792d70abe52d548ab45f56fd1ddad04b6bd.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ib5e8a06130dd4a8912ddd71ef5920792d70abe52d548ab45f56fd1ddad04b6bd.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i7e34b02d89b096a5000726e206dbb52a597c7b514a4ac30da1033bbb4b03fdb5.GetManagedAppPoliciesRequestBuilder) {
    return i7e34b02d89b096a5000726e206dbb52a597c7b514a4ac30da1033bbb4b03fdb5.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ia91e670d6c14be22de7af362c8a7fc0b8fb7d88d7b86fd7a1eeb3ade273221a1.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ia91e670d6c14be22de7af362c8a7fc0b8fb7d88d7b86fd7a1eeb3ade273221a1.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ie4aa4f53c8accb61e9768af991587b52c7bd11a8aee60fc46e3d5d3defe9f79f.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ie4aa4f53c8accb61e9768af991587b52c7bd11a8aee60fc46e3d5d3defe9f79f.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i64daf8facf298b392228c1f8fbc21378bc0c8dc62e11a0dac1948221187db7e9.GetMemberGroupsRequestBuilder) {
    return i64daf8facf298b392228c1f8fbc21378bc0c8dc62e11a0dac1948221187db7e9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i3e03205aab7df19b11a5d372d028b0fb8b7cd650446063ba5b762015f01ddddc.GetMemberObjectsRequestBuilder) {
    return i3e03205aab7df19b11a5d372d028b0fb8b7cd650446063ba5b762015f01ddddc.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i1b24166e9a5f19517be26bc02aafe609c4f4614e0e1fb3b4bf23b5e82650bf33.InvalidateAllRefreshTokensRequestBuilder) {
    return i1b24166e9a5f19517be26bc02aafe609c4f4614e0e1fb3b4bf23b5e82650bf33.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i9edacdd56843f19c386fe7376a907b5df2c93b768097be253e4b040c8d7f5a04.IsManagedAppUserBlockedRequestBuilder) {
    return i9edacdd56843f19c386fe7376a907b5df2c93b768097be253e4b040c8d7f5a04.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i6b89b729682111d71f55bb444fea0b59cc3414dcc46e97498bd425ae1f10918a.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i6b89b729682111d71f55bb444fea0b59cc3414dcc46e97498bd425ae1f10918a.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*id3a3d95622a2b03a996c9c6125c29aab8cff342a443b8b02ee8bb73ea8ee99f0.RemoveAllDevicesFromManagementRequestBuilder) {
    return id3a3d95622a2b03a996c9c6125c29aab8cff342a443b8b02ee8bb73ea8ee99f0.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i7bf9f47041550eaca6a776837ea50de11dc1348704e9508f028484ca9c9c70f6.ReprocessLicenseAssignmentRequestBuilder) {
    return i7bf9f47041550eaca6a776837ea50de11dc1348704e9508f028484ca9c9c70f6.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i10b735093bfb492238e268241c08e3bbc23abf518bd6ca7447397fbe86cf94f4.RestoreRequestBuilder) {
    return i10b735093bfb492238e268241c08e3bbc23abf518bd6ca7447397fbe86cf94f4.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i1ca3e90ce47b30c9b1a10faa1af83062940a12a76d75dff9199f22a815df060c.RevokeSignInSessionsRequestBuilder) {
    return i1ca3e90ce47b30c9b1a10faa1af83062940a12a76d75dff9199f22a815df060c.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i44d0bc774f5245867ccc64328bd33234f807fbb243d6ca9a4e10a38cc4a30453.SendMailRequestBuilder) {
    return i44d0bc774f5245867ccc64328bd33234f807fbb243d6ca9a4e10a38cc4a30453.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i70e1b2be4e88dafb6c58959844b65d41c96f72061b2686508eae9ed17be0fc06.TranslateExchangeIdsRequestBuilder) {
    return i70e1b2be4e88dafb6c58959844b65d41c96f72061b2686508eae9ed17be0fc06.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i7a198c110253c6692c34999de1fbc8b8f99b7d0aedd7263920e9fad91e31562e.UnblockManagedAppsRequestBuilder) {
    return i7a198c110253c6692c34999de1fbc8b8f99b7d0aedd7263920e9fad91e31562e.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*iddf84c133c3e60b644f39417d4d67ba4abe623d5f365246df5ace24375295581.WipeAndBlockManagedAppsRequestBuilder) {
    return iddf84c133c3e60b644f39417d4d67ba4abe623d5f365246df5ace24375295581.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*id5c529f39116c945503ac8bfa0b787bc6d032f884415e02e1315a39891a4edfe.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return id5c529f39116c945503ac8bfa0b787bc6d032f884415e02e1315a39891a4edfe.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*iae59d6fd2345d27fe78dfc078d36bf84351966f5b7b20dc6d6f8ac82a1cb5312.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return iae59d6fd2345d27fe78dfc078d36bf84351966f5b7b20dc6d6f8ac82a1cb5312.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i8cbe1311cfb2623bf9606c48ffc155d951b5949948162e0b9340b626f04b8741.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i8cbe1311cfb2623bf9606c48ffc155d951b5949948162e0b9340b626f04b8741.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
