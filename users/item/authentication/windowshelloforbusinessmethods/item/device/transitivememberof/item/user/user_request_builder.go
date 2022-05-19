package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01317bd929bf5d8d48da537b70f192e6729f56f77cc14c1af8a10895e1a7fe42 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i02c169a0c2edbec14d14cc4b89c9311c5a4c82f78af52ba0c646e2454c1be706 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i0587a01a07e1edd9ba0f5d61ebbe7ada5d8b9536679e2651377005b741dfd30f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    i2023f96904e737c3ade5dfd10c575a934f1e85d2c435f8125f9a02b03244f45c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/restore"
    i355c763c6ca6cd855c026717f574587ff038cb36ace3ed0cd88a13b1dc6fe22e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i35611a0576d01005a284ff53e927f161bd807155c783ce0935713eb6c34329de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmailtips"
    i395fe3a7f1e4b8e1332eda0553797d7d7c631dc10b6c577571edee5f30b5df0d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i42b1b01eaaebafc162481de9a9751cbbe16de02987112fc3529a3b8b35bdaf21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i49e3c8b7296fbdd47c64ae4cd9150dd1325a0e30b5398b9fa2f2af4dc918c99a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i527863e77616f93b0f1f28287821e408926308d1f8100ab26df526dab7f82ef5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findroomswithroomlist"
    i5a1f9dadab9c056d0d68e02b9891074bac69411273a346a9c50cd164ce208b22 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i63ae9c7cd6393326c7b643f8bf295d693921354ac6774051e874eed0774afbf6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/assignlicense"
    i7da5ee362a179726f658fa0d527efbcc5b878138c9e30670f7338e67edd12058 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findroomlists"
    i7dbe311aa11d9fe4b612594c893d7f937887510900cef9080a09214754fc49af "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/activateserviceplan"
    i85fd1eb970078ed5fdcff2d03e7a5846636fcdb9b95098b33be372b5aa9daae8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    i875ed252268a7432d064bc2b5589f13c9894ef2e09e6c23dee4ad07a89975028 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/sendmail"
    i95974f84b8eff95c72e139983b632cfa981e6ef0990a768161caaba5f3dcb8be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    iaa591f9b4d79558add7183f6bad1f61b83fb162d66401ccd0bd50c25088d2544 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    ib1eaaa01a6b495cd91ef33b0f6ddcb5cfdef4cd6ffddfb6e10052caf385c7ea4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/translateexchangeids"
    ib21cda8a149368b69d83c7410507c2945987a069c92fc789bbbabdfb55c45d6e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    ibc45d1701a6bc6b179919d2429d05f2ec07e4864cbeeee411cc093e3286fd09d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    ibd27044063631a4d8ae028b172f34832b41a7ca947a4a5634e4e85b180e47eed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ic084b419aa393a4d169eae67d8dbbfe3c75d40e2dab87bcae89568db9be522ca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmembergroups"
    ic8db0164bb3f96027d122c99532b990938aae086b99da33f144f54391ada536c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    icfaf78820cfdbc57db2a12e969fa9198402d57eec7fc4573d0d074c80abd0163 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    id2b8757b41410bd338d3f55d856e2e2c3d8419649ca81ae973e486f9ca06cbc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    id31326dc40a93902d4063034e15824ff754ffd832e69df2da8232fce16d24b87 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    ie4338fefeafd6f3267c4c8c9776399910d0debe1309d9ca65aec1c4b0d59af62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmemberobjects"
    ie4c96b0fddf59dad7d402e5a819c5d7e7958bed3c97b4e9847b05bd30b38b676 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmemberobjects"
    ie8a482d69ec23b2455bcac57f01c2f823323262906f779c2ba0e965728b1c649 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportpersonaldata"
    ie906903106172793c0d928032e5249a3af1623ad9972e49d8d592c5a84f8c7a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findrooms"
    ied928352d9a45703cf0ea0ce262ed53b8d9f60731f0c3769ae6d1d0ccbd58262 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmembergroups"
    if0524bf2c1e3abcc28c9d067ec72ad49aa68aaf65a5250ec3003424991b74154 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/changepassword"
    if5fd4900140fb9b42285ef963204fed00acee3a4165e229a4a28c505876943f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ifd917ffd5b53e232f6f70ec41abd49b0a174b9b5fd8a9f65d2b3003b1fb8c715 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i7dbe311aa11d9fe4b612594c893d7f937887510900cef9080a09214754fc49af.ActivateServicePlanRequestBuilder) {
    return i7dbe311aa11d9fe4b612594c893d7f937887510900cef9080a09214754fc49af.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i63ae9c7cd6393326c7b643f8bf295d693921354ac6774051e874eed0774afbf6.AssignLicenseRequestBuilder) {
    return i63ae9c7cd6393326c7b643f8bf295d693921354ac6774051e874eed0774afbf6.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*if0524bf2c1e3abcc28c9d067ec72ad49aa68aaf65a5250ec3003424991b74154.ChangePasswordRequestBuilder) {
    return if0524bf2c1e3abcc28c9d067ec72ad49aa68aaf65a5250ec3003424991b74154.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ied928352d9a45703cf0ea0ce262ed53b8d9f60731f0c3769ae6d1d0ccbd58262.CheckMemberGroupsRequestBuilder) {
    return ied928352d9a45703cf0ea0ce262ed53b8d9f60731f0c3769ae6d1d0ccbd58262.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ie4338fefeafd6f3267c4c8c9776399910d0debe1309d9ca65aec1c4b0d59af62.CheckMemberObjectsRequestBuilder) {
    return ie4338fefeafd6f3267c4c8c9776399910d0debe1309d9ca65aec1c4b0d59af62.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ib21cda8a149368b69d83c7410507c2945987a069c92fc789bbbabdfb55c45d6e.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ib21cda8a149368b69d83c7410507c2945987a069c92fc789bbbabdfb55c45d6e.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*if5fd4900140fb9b42285ef963204fed00acee3a4165e229a4a28c505876943f1.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return if5fd4900140fb9b42285ef963204fed00acee3a4165e229a4a28c505876943f1.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ie8a482d69ec23b2455bcac57f01c2f823323262906f779c2ba0e965728b1c649.ExportPersonalDataRequestBuilder) {
    return ie8a482d69ec23b2455bcac57f01c2f823323262906f779c2ba0e965728b1c649.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i395fe3a7f1e4b8e1332eda0553797d7d7c631dc10b6c577571edee5f30b5df0d.FindMeetingTimesRequestBuilder) {
    return i395fe3a7f1e4b8e1332eda0553797d7d7c631dc10b6c577571edee5f30b5df0d.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i7da5ee362a179726f658fa0d527efbcc5b878138c9e30670f7338e67edd12058.FindRoomListsRequestBuilder) {
    return i7da5ee362a179726f658fa0d527efbcc5b878138c9e30670f7338e67edd12058.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ie906903106172793c0d928032e5249a3af1623ad9972e49d8d592c5a84f8c7a2.FindRoomsRequestBuilder) {
    return ie906903106172793c0d928032e5249a3af1623ad9972e49d8d592c5a84f8c7a2.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i527863e77616f93b0f1f28287821e408926308d1f8100ab26df526dab7f82ef5.FindRoomsWithRoomListRequestBuilder) {
    return i527863e77616f93b0f1f28287821e408926308d1f8100ab26df526dab7f82ef5.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i5a1f9dadab9c056d0d68e02b9891074bac69411273a346a9c50cd164ce208b22.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i5a1f9dadab9c056d0d68e02b9891074bac69411273a346a9c50cd164ce208b22.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i42b1b01eaaebafc162481de9a9751cbbe16de02987112fc3529a3b8b35bdaf21.GetLoggedOnManagedDevicesRequestBuilder) {
    return i42b1b01eaaebafc162481de9a9751cbbe16de02987112fc3529a3b8b35bdaf21.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i35611a0576d01005a284ff53e927f161bd807155c783ce0935713eb6c34329de.GetMailTipsRequestBuilder) {
    return i35611a0576d01005a284ff53e927f161bd807155c783ce0935713eb6c34329de.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i02c169a0c2edbec14d14cc4b89c9311c5a4c82f78af52ba0c646e2454c1be706.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i02c169a0c2edbec14d14cc4b89c9311c5a4c82f78af52ba0c646e2454c1be706.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i95974f84b8eff95c72e139983b632cfa981e6ef0990a768161caaba5f3dcb8be.GetManagedAppPoliciesRequestBuilder) {
    return i95974f84b8eff95c72e139983b632cfa981e6ef0990a768161caaba5f3dcb8be.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*id2b8757b41410bd338d3f55d856e2e2c3d8419649ca81ae973e486f9ca06cbc2.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return id2b8757b41410bd338d3f55d856e2e2c3d8419649ca81ae973e486f9ca06cbc2.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ibd27044063631a4d8ae028b172f34832b41a7ca947a4a5634e4e85b180e47eed.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ibd27044063631a4d8ae028b172f34832b41a7ca947a4a5634e4e85b180e47eed.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ic084b419aa393a4d169eae67d8dbbfe3c75d40e2dab87bcae89568db9be522ca.GetMemberGroupsRequestBuilder) {
    return ic084b419aa393a4d169eae67d8dbbfe3c75d40e2dab87bcae89568db9be522ca.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ie4c96b0fddf59dad7d402e5a819c5d7e7958bed3c97b4e9847b05bd30b38b676.GetMemberObjectsRequestBuilder) {
    return ie4c96b0fddf59dad7d402e5a819c5d7e7958bed3c97b4e9847b05bd30b38b676.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i85fd1eb970078ed5fdcff2d03e7a5846636fcdb9b95098b33be372b5aa9daae8.InvalidateAllRefreshTokensRequestBuilder) {
    return i85fd1eb970078ed5fdcff2d03e7a5846636fcdb9b95098b33be372b5aa9daae8.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i0587a01a07e1edd9ba0f5d61ebbe7ada5d8b9536679e2651377005b741dfd30f.IsManagedAppUserBlockedRequestBuilder) {
    return i0587a01a07e1edd9ba0f5d61ebbe7ada5d8b9536679e2651377005b741dfd30f.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*id31326dc40a93902d4063034e15824ff754ffd832e69df2da8232fce16d24b87.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return id31326dc40a93902d4063034e15824ff754ffd832e69df2da8232fce16d24b87.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ibc45d1701a6bc6b179919d2429d05f2ec07e4864cbeeee411cc093e3286fd09d.RemoveAllDevicesFromManagementRequestBuilder) {
    return ibc45d1701a6bc6b179919d2429d05f2ec07e4864cbeeee411cc093e3286fd09d.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ic8db0164bb3f96027d122c99532b990938aae086b99da33f144f54391ada536c.ReprocessLicenseAssignmentRequestBuilder) {
    return ic8db0164bb3f96027d122c99532b990938aae086b99da33f144f54391ada536c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i2023f96904e737c3ade5dfd10c575a934f1e85d2c435f8125f9a02b03244f45c.RestoreRequestBuilder) {
    return i2023f96904e737c3ade5dfd10c575a934f1e85d2c435f8125f9a02b03244f45c.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i01317bd929bf5d8d48da537b70f192e6729f56f77cc14c1af8a10895e1a7fe42.RevokeSignInSessionsRequestBuilder) {
    return i01317bd929bf5d8d48da537b70f192e6729f56f77cc14c1af8a10895e1a7fe42.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i875ed252268a7432d064bc2b5589f13c9894ef2e09e6c23dee4ad07a89975028.SendMailRequestBuilder) {
    return i875ed252268a7432d064bc2b5589f13c9894ef2e09e6c23dee4ad07a89975028.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ib1eaaa01a6b495cd91ef33b0f6ddcb5cfdef4cd6ffddfb6e10052caf385c7ea4.TranslateExchangeIdsRequestBuilder) {
    return ib1eaaa01a6b495cd91ef33b0f6ddcb5cfdef4cd6ffddfb6e10052caf385c7ea4.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*icfaf78820cfdbc57db2a12e969fa9198402d57eec7fc4573d0d074c80abd0163.UnblockManagedAppsRequestBuilder) {
    return icfaf78820cfdbc57db2a12e969fa9198402d57eec7fc4573d0d074c80abd0163.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ifd917ffd5b53e232f6f70ec41abd49b0a174b9b5fd8a9f65d2b3003b1fb8c715.WipeAndBlockManagedAppsRequestBuilder) {
    return ifd917ffd5b53e232f6f70ec41abd49b0a174b9b5fd8a9f65d2b3003b1fb8c715.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i355c763c6ca6cd855c026717f574587ff038cb36ace3ed0cd88a13b1dc6fe22e.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i355c763c6ca6cd855c026717f574587ff038cb36ace3ed0cd88a13b1dc6fe22e.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i49e3c8b7296fbdd47c64ae4cd9150dd1325a0e30b5398b9fa2f2af4dc918c99a.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i49e3c8b7296fbdd47c64ae4cd9150dd1325a0e30b5398b9fa2f2af4dc918c99a.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*iaa591f9b4d79558add7183f6bad1f61b83fb162d66401ccd0bd50c25088d2544.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return iaa591f9b4d79558add7183f6bad1f61b83fb162d66401ccd0bd50c25088d2544.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
