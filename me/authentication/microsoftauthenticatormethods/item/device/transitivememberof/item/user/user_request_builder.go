package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2f3b138172fd3ca4da361be8e9e7c78ef830be67c272fb70df28ec4bbdd0bf00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i3ab967c7bde42b50e9e1bcb2ab46732a7330abad27d4b3772fa11fa91d98e51f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    i3f7ea5a0d80b3fcebe0b7f5fab33e100ae13300da716df7741f8ca6df99abe64 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i4372199fe7e0b4d3c1d8d7d7585c7b4f45475b7906067a9dee965cb1e84ca96e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i4c14a9df4e3e66b4c75256331ade0c9adc04783ec7e014c287a5c9ca91c40cac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/assignlicense"
    i4d9f40bcadb0e82344799d72bebd8e27372f89a743c069a6218e88add9ab7fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/restore"
    i4e52440a0aaefd83afc4c75ea6c3e928ddcc1f998d4197a177b011617cac5017 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/translateexchangeids"
    i5aa5efe557cfc7a5c6a30076e37cb69f201e354f1497685ce148f90544a8483c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmemberobjects"
    i6cc0121b43d91806ded2e46a719748a6785cd5b73bd547395830e57903a2b629 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmailtips"
    i6e6aee6b16e91b497e2a21fc1fef2dea4a3c69fabb8966c0e3dd09f78c2ded33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    i762cf4ce71ca5fd8dec0828d9832a04f42a2b1a341f0df97b28c7c43344d5a07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    i7718d76924bb7742c0fd2533e3ee6824b3e033bbc72d6ba5b4026db335e4fcc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i7a305b6f2674d4e0282beaa2dbd15f82a496d317339916212108fb35f541ff69 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmemberobjects"
    i808c197a689098175a0700109cd25fa1fd17c445ff75637ed317898b5502f902 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/activateserviceplan"
    i91f87993774b414e08827b71c67927ecd6f83da39bd08a409beeb4ad2db09ad4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    i95551ebcb223461d478bea0de8fe7aaaf84c221067d37fb50ab5a1e2c1d17379 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findmeetingtimes"
    i9993bc3646d20c12c524747a81ad1e830463732d8e1f7a9573e6d2ca858fb3f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i9c3d16c4cf239633698728da09c7dd7ede94a4144d1f69d32be04167b948fe10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    ia14531eef72874f755b3bdc3697699c0876fd0c32461b321aa04fd91d0d83ddc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    ia2a61537d977b5fe836ac15271b11df027f2a4f617d4f2b9bc9bcf86fd1fdf5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ia8502cce8a1be469bfd51150ef87399604ad447231b2228f9abb4e4051df0046 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    ia8ab2092df7e74872d5d6d2c662cda2cd8cc6e811a953994aefb675c013ead92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/checkmembergroups"
    ia92d33a6692e0217125485a86d817f166fabd528d285143454c80c4c6baea1aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    iacc65d7d5e101847dfb5a636952034d98a554b01cebc48cd873d60275025c1fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/changepassword"
    ib115aa00b53df99001077072eefdadc9093ce0f2d17e7c66caa1493b1ed14451 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/sendmail"
    ib77315011efe822ecb8d9edc81560de99b40a289907c4e622bc9496672ff6fdb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/exportpersonaldata"
    ibca8fc5014e0408c7fa59f6033791ae09bda81de25574e8f96d48a02d696469d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    id3b59f67f31075c8969cd5af88535d8632abdab35dc4eb656e49381e68e233df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    id4955fe28712eb6d9e140a90864a22c0a3441ec4f0ed3b72061b36c270b162fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findroomlists"
    id4a6b01a7878bc8ba9490a9f9c053f8c5f3c717d347ec818d5f43b005d526307 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/getmembergroups"
    ie3ea0d79ab353c49cbc55522c3b5dbaf805a697923c5e6957fa726520f16e61a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
    ie8f557bd36d6785d68647503dd00a12a0d449052e71b26f400b2c34def71faf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    if4c35389db304ceee05cab9f0d2888abbe1184d9e504332df1b780fdec55e81d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ifb8aa06a1325c2b3c959cbc8647d16eb1a51687a9f9e045b5fa52f6a4ca56bdf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findrooms"
    ifde5dfc1539b2accfd2eac950e1de44ee586db02f1c422285d57f70598175122 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/transitivememberof/item/user/findroomswithroomlist"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i808c197a689098175a0700109cd25fa1fd17c445ff75637ed317898b5502f902.ActivateServicePlanRequestBuilder) {
    return i808c197a689098175a0700109cd25fa1fd17c445ff75637ed317898b5502f902.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i4c14a9df4e3e66b4c75256331ade0c9adc04783ec7e014c287a5c9ca91c40cac.AssignLicenseRequestBuilder) {
    return i4c14a9df4e3e66b4c75256331ade0c9adc04783ec7e014c287a5c9ca91c40cac.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*iacc65d7d5e101847dfb5a636952034d98a554b01cebc48cd873d60275025c1fe.ChangePasswordRequestBuilder) {
    return iacc65d7d5e101847dfb5a636952034d98a554b01cebc48cd873d60275025c1fe.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ia8ab2092df7e74872d5d6d2c662cda2cd8cc6e811a953994aefb675c013ead92.CheckMemberGroupsRequestBuilder) {
    return ia8ab2092df7e74872d5d6d2c662cda2cd8cc6e811a953994aefb675c013ead92.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i7a305b6f2674d4e0282beaa2dbd15f82a496d317339916212108fb35f541ff69.CheckMemberObjectsRequestBuilder) {
    return i7a305b6f2674d4e0282beaa2dbd15f82a496d317339916212108fb35f541ff69.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i9c3d16c4cf239633698728da09c7dd7ede94a4144d1f69d32be04167b948fe10.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i9c3d16c4cf239633698728da09c7dd7ede94a4144d1f69d32be04167b948fe10.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i9993bc3646d20c12c524747a81ad1e830463732d8e1f7a9573e6d2ca858fb3f7.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i9993bc3646d20c12c524747a81ad1e830463732d8e1f7a9573e6d2ca858fb3f7.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ib77315011efe822ecb8d9edc81560de99b40a289907c4e622bc9496672ff6fdb.ExportPersonalDataRequestBuilder) {
    return ib77315011efe822ecb8d9edc81560de99b40a289907c4e622bc9496672ff6fdb.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i95551ebcb223461d478bea0de8fe7aaaf84c221067d37fb50ab5a1e2c1d17379.FindMeetingTimesRequestBuilder) {
    return i95551ebcb223461d478bea0de8fe7aaaf84c221067d37fb50ab5a1e2c1d17379.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*id4955fe28712eb6d9e140a90864a22c0a3441ec4f0ed3b72061b36c270b162fc.FindRoomListsRequestBuilder) {
    return id4955fe28712eb6d9e140a90864a22c0a3441ec4f0ed3b72061b36c270b162fc.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*ifb8aa06a1325c2b3c959cbc8647d16eb1a51687a9f9e045b5fa52f6a4ca56bdf.FindRoomsRequestBuilder) {
    return ifb8aa06a1325c2b3c959cbc8647d16eb1a51687a9f9e045b5fa52f6a4ca56bdf.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ifde5dfc1539b2accfd2eac950e1de44ee586db02f1c422285d57f70598175122.FindRoomsWithRoomListRequestBuilder) {
    return ifde5dfc1539b2accfd2eac950e1de44ee586db02f1c422285d57f70598175122.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ia2a61537d977b5fe836ac15271b11df027f2a4f617d4f2b9bc9bcf86fd1fdf5d.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ia2a61537d977b5fe836ac15271b11df027f2a4f617d4f2b9bc9bcf86fd1fdf5d.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i4372199fe7e0b4d3c1d8d7d7585c7b4f45475b7906067a9dee965cb1e84ca96e.GetLoggedOnManagedDevicesRequestBuilder) {
    return i4372199fe7e0b4d3c1d8d7d7585c7b4f45475b7906067a9dee965cb1e84ca96e.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i6cc0121b43d91806ded2e46a719748a6785cd5b73bd547395830e57903a2b629.GetMailTipsRequestBuilder) {
    return i6cc0121b43d91806ded2e46a719748a6785cd5b73bd547395830e57903a2b629.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i3f7ea5a0d80b3fcebe0b7f5fab33e100ae13300da716df7741f8ca6df99abe64.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i3f7ea5a0d80b3fcebe0b7f5fab33e100ae13300da716df7741f8ca6df99abe64.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*id3b59f67f31075c8969cd5af88535d8632abdab35dc4eb656e49381e68e233df.GetManagedAppPoliciesRequestBuilder) {
    return id3b59f67f31075c8969cd5af88535d8632abdab35dc4eb656e49381e68e233df.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i6e6aee6b16e91b497e2a21fc1fef2dea4a3c69fabb8966c0e3dd09f78c2ded33.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i6e6aee6b16e91b497e2a21fc1fef2dea4a3c69fabb8966c0e3dd09f78c2ded33.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7718d76924bb7742c0fd2533e3ee6824b3e033bbc72d6ba5b4026db335e4fcc0.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7718d76924bb7742c0fd2533e3ee6824b3e033bbc72d6ba5b4026db335e4fcc0.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*id4a6b01a7878bc8ba9490a9f9c053f8c5f3c717d347ec818d5f43b005d526307.GetMemberGroupsRequestBuilder) {
    return id4a6b01a7878bc8ba9490a9f9c053f8c5f3c717d347ec818d5f43b005d526307.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i5aa5efe557cfc7a5c6a30076e37cb69f201e354f1497685ce148f90544a8483c.GetMemberObjectsRequestBuilder) {
    return i5aa5efe557cfc7a5c6a30076e37cb69f201e354f1497685ce148f90544a8483c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ibca8fc5014e0408c7fa59f6033791ae09bda81de25574e8f96d48a02d696469d.InvalidateAllRefreshTokensRequestBuilder) {
    return ibca8fc5014e0408c7fa59f6033791ae09bda81de25574e8f96d48a02d696469d.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i762cf4ce71ca5fd8dec0828d9832a04f42a2b1a341f0df97b28c7c43344d5a07.IsManagedAppUserBlockedRequestBuilder) {
    return i762cf4ce71ca5fd8dec0828d9832a04f42a2b1a341f0df97b28c7c43344d5a07.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ia92d33a6692e0217125485a86d817f166fabd528d285143454c80c4c6baea1aa.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ia92d33a6692e0217125485a86d817f166fabd528d285143454c80c4c6baea1aa.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i3ab967c7bde42b50e9e1bcb2ab46732a7330abad27d4b3772fa11fa91d98e51f.RemoveAllDevicesFromManagementRequestBuilder) {
    return i3ab967c7bde42b50e9e1bcb2ab46732a7330abad27d4b3772fa11fa91d98e51f.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i91f87993774b414e08827b71c67927ecd6f83da39bd08a409beeb4ad2db09ad4.ReprocessLicenseAssignmentRequestBuilder) {
    return i91f87993774b414e08827b71c67927ecd6f83da39bd08a409beeb4ad2db09ad4.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i4d9f40bcadb0e82344799d72bebd8e27372f89a743c069a6218e88add9ab7fb2.RestoreRequestBuilder) {
    return i4d9f40bcadb0e82344799d72bebd8e27372f89a743c069a6218e88add9ab7fb2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i2f3b138172fd3ca4da361be8e9e7c78ef830be67c272fb70df28ec4bbdd0bf00.RevokeSignInSessionsRequestBuilder) {
    return i2f3b138172fd3ca4da361be8e9e7c78ef830be67c272fb70df28ec4bbdd0bf00.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ib115aa00b53df99001077072eefdadc9093ce0f2d17e7c66caa1493b1ed14451.SendMailRequestBuilder) {
    return ib115aa00b53df99001077072eefdadc9093ce0f2d17e7c66caa1493b1ed14451.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i4e52440a0aaefd83afc4c75ea6c3e928ddcc1f998d4197a177b011617cac5017.TranslateExchangeIdsRequestBuilder) {
    return i4e52440a0aaefd83afc4c75ea6c3e928ddcc1f998d4197a177b011617cac5017.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ia8502cce8a1be469bfd51150ef87399604ad447231b2228f9abb4e4051df0046.UnblockManagedAppsRequestBuilder) {
    return ia8502cce8a1be469bfd51150ef87399604ad447231b2228f9abb4e4051df0046.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ie3ea0d79ab353c49cbc55522c3b5dbaf805a697923c5e6957fa726520f16e61a.WipeAndBlockManagedAppsRequestBuilder) {
    return ie3ea0d79ab353c49cbc55522c3b5dbaf805a697923c5e6957fa726520f16e61a.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ia14531eef72874f755b3bdc3697699c0876fd0c32461b321aa04fd91d0d83ddc.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ia14531eef72874f755b3bdc3697699c0876fd0c32461b321aa04fd91d0d83ddc.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*if4c35389db304ceee05cab9f0d2888abbe1184d9e504332df1b780fdec55e81d.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return if4c35389db304ceee05cab9f0d2888abbe1184d9e504332df1b780fdec55e81d.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ie8f557bd36d6785d68647503dd00a12a0d449052e71b26f400b2c34def71faf3.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ie8f557bd36d6785d68647503dd00a12a0d449052e71b26f400b2c34def71faf3.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
