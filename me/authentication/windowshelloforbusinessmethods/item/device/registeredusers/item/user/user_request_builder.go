package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00e412e0d921615577de9ef45bf7d8487f8e2b146d665ec27929dcc40aa5d249 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmailtips"
    i0183a420e1adf7a0e5b4673d52e8c38584ecf6840a13407a1cb22aa1ab8467ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i0d438253598336480ea6e5321e519fd8f8a4d5e064fb9cdf9163c0ebc342d4b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    i0fe5d4e20194af0ed86dcea25a2ff910bf31a3cf289832b920e4142fba0f1178 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    i18f3bdc8608c01d48e43cc8bc346df5a4207becf5c108e05ed77f27ab67a250f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmemberobjects"
    i23281cbd45ce85cdcf686f18da752247499c036cac07360733ae51f1d1c9008e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    i241af187f2c9132f90f8637aa21ad7382f0af9c21c417abd86fe5a47dceb04ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/assignlicense"
    i31f5050a8f7b60df1c8ddeaf884fd4d0006f0477c94fcd5b99adf5df8afb1eb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i33b2e260b24259029d6e526fe10fbca0f951a8135cc61a900f1247cfc9b0033c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/exportpersonaldata"
    i4122d581e5e42f864e7cc396199ba26fed8ccc9f15971e3daf2f70ba1a6befac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/restore"
    i415a6ee4477e053c8f254df656011677bab23bf424f8c6dd036d853346da676b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i4b40126bf795b81b4ca482e42213506cdbe79d55495e6c75db5c65bb7e3f2571 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findroomswithroomlist"
    i5575f1a19094aadea7f5a104c20e3b8d60391e2306cf77461098ba445c16a404 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    i5c4dac192f4ad23f95bb6b7a2e5bd88f44ac9ecb7fb412f5b7dfdf57aea4a487 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    i5c5ab6c54a14b07be3b65bd09dbb3211b7e6a302a0081b1c35674b8299a458c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/unblockmanagedapps"
    i648e105a158b93eeb1f75b32e003852fa4d46bb07302c3a7000d3766938b831a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findroomlists"
    i664fd63d09162036eb43c7c5aa679bfa9496b90a61c8c34d05a2a4987c9805ab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/checkmembergroups"
    i6f8cd378835127ee4d947e21d745da837a5da4c50382812993d2f3c535d86faa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i7a0775e5e65a177503bff9b2741bd8904c74a9dd211fccd5baad69e5128048b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmembergroups"
    i817b1d13283a26bc2c0fa710788863cc0adde4371193357ddc6df9b0a993955d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/activateserviceplan"
    i8249d410e2eaf50e7df6fa1d23e25f1b6f81d9b4de4f1654706882e5b9efd862 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i878f1aaf7c19a4cafdb06363cf5f534d1e0a4cef10f6ff0fb5871b3a35c9ff05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/sendmail"
    i955ffc7b0625a4f2266d9d1a9a6bfa7afc228bf8c6f67beb231bbc7cc8c6787d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ia3b7a267bed5cc58f641ce43ffe0aa74f23557621a07fabd77bd28e9908bed9c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    ia8c58872e4e9fbad5c500292d61411649e61166d99a7d9b622aa1a3ad0961917 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmemberobjects"
    iaa1c6edacf5feba0b2f962a849221cf85d5b7c28d8aac5a97aa838fcf3219c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    iaa4d047bbe74709daea43f0f48f0ed5fbe7d37f4460117da03c8cb9960ea8d5e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findrooms"
    ib3e3e46018b8f12d67a7c180fab11f6d3a70d22b507cba020f9e3cf5a73884b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/changepassword"
    ibdaac48dbde1aec4a8c246f3df6d690eb9e4e8d9a5c9c1a46cc9a17a438fbc56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/findmeetingtimes"
    ic9064c196b1eafcb6a94734daf3b957b68ce46d0fef1407fb37459013ff2691a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/revokesigninsessions"
    id74275d816899a3c55d91cd42f88022614ae74103753a6782e1d4bc99c7bd5b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    id8eff0d6a58a6bbe710a3158f2b444d51e014e83002bd032a0c6d417973f6d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    idb4477d2f3ad38f8a51d64de95e385c0e7d5395eaeb7513d8171e0859ad0821a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    ie5192d934b9d560a424770a4b92b860673860f40a76e5b882b4b612b0bb9734f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    ife4357fa284e05bcfcb62dcfac0bcc3a4d2196500d46bffceadfc5fb53bad962 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredusers/item/user/translateexchangeids"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i817b1d13283a26bc2c0fa710788863cc0adde4371193357ddc6df9b0a993955d.ActivateServicePlanRequestBuilder) {
    return i817b1d13283a26bc2c0fa710788863cc0adde4371193357ddc6df9b0a993955d.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i241af187f2c9132f90f8637aa21ad7382f0af9c21c417abd86fe5a47dceb04ea.AssignLicenseRequestBuilder) {
    return i241af187f2c9132f90f8637aa21ad7382f0af9c21c417abd86fe5a47dceb04ea.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ib3e3e46018b8f12d67a7c180fab11f6d3a70d22b507cba020f9e3cf5a73884b6.ChangePasswordRequestBuilder) {
    return ib3e3e46018b8f12d67a7c180fab11f6d3a70d22b507cba020f9e3cf5a73884b6.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i664fd63d09162036eb43c7c5aa679bfa9496b90a61c8c34d05a2a4987c9805ab.CheckMemberGroupsRequestBuilder) {
    return i664fd63d09162036eb43c7c5aa679bfa9496b90a61c8c34d05a2a4987c9805ab.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i18f3bdc8608c01d48e43cc8bc346df5a4207becf5c108e05ed77f27ab67a250f.CheckMemberObjectsRequestBuilder) {
    return i18f3bdc8608c01d48e43cc8bc346df5a4207becf5c108e05ed77f27ab67a250f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i0d438253598336480ea6e5321e519fd8f8a4d5e064fb9cdf9163c0ebc342d4b1.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i0d438253598336480ea6e5321e519fd8f8a4d5e064fb9cdf9163c0ebc342d4b1.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i31f5050a8f7b60df1c8ddeaf884fd4d0006f0477c94fcd5b99adf5df8afb1eb1.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i31f5050a8f7b60df1c8ddeaf884fd4d0006f0477c94fcd5b99adf5df8afb1eb1.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i33b2e260b24259029d6e526fe10fbca0f951a8135cc61a900f1247cfc9b0033c.ExportPersonalDataRequestBuilder) {
    return i33b2e260b24259029d6e526fe10fbca0f951a8135cc61a900f1247cfc9b0033c.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ibdaac48dbde1aec4a8c246f3df6d690eb9e4e8d9a5c9c1a46cc9a17a438fbc56.FindMeetingTimesRequestBuilder) {
    return ibdaac48dbde1aec4a8c246f3df6d690eb9e4e8d9a5c9c1a46cc9a17a438fbc56.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i648e105a158b93eeb1f75b32e003852fa4d46bb07302c3a7000d3766938b831a.FindRoomListsRequestBuilder) {
    return i648e105a158b93eeb1f75b32e003852fa4d46bb07302c3a7000d3766938b831a.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*iaa4d047bbe74709daea43f0f48f0ed5fbe7d37f4460117da03c8cb9960ea8d5e.FindRoomsRequestBuilder) {
    return iaa4d047bbe74709daea43f0f48f0ed5fbe7d37f4460117da03c8cb9960ea8d5e.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i4b40126bf795b81b4ca482e42213506cdbe79d55495e6c75db5c65bb7e3f2571.FindRoomsWithRoomListRequestBuilder) {
    return i4b40126bf795b81b4ca482e42213506cdbe79d55495e6c75db5c65bb7e3f2571.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i5575f1a19094aadea7f5a104c20e3b8d60391e2306cf77461098ba445c16a404.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i5575f1a19094aadea7f5a104c20e3b8d60391e2306cf77461098ba445c16a404.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i5c4dac192f4ad23f95bb6b7a2e5bd88f44ac9ecb7fb412f5b7dfdf57aea4a487.GetLoggedOnManagedDevicesRequestBuilder) {
    return i5c4dac192f4ad23f95bb6b7a2e5bd88f44ac9ecb7fb412f5b7dfdf57aea4a487.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i00e412e0d921615577de9ef45bf7d8487f8e2b146d665ec27929dcc40aa5d249.GetMailTipsRequestBuilder) {
    return i00e412e0d921615577de9ef45bf7d8487f8e2b146d665ec27929dcc40aa5d249.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i6f8cd378835127ee4d947e21d745da837a5da4c50382812993d2f3c535d86faa.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i6f8cd378835127ee4d947e21d745da837a5da4c50382812993d2f3c535d86faa.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*id74275d816899a3c55d91cd42f88022614ae74103753a6782e1d4bc99c7bd5b1.GetManagedAppPoliciesRequestBuilder) {
    return id74275d816899a3c55d91cd42f88022614ae74103753a6782e1d4bc99c7bd5b1.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*idb4477d2f3ad38f8a51d64de95e385c0e7d5395eaeb7513d8171e0859ad0821a.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return idb4477d2f3ad38f8a51d64de95e385c0e7d5395eaeb7513d8171e0859ad0821a.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ie5192d934b9d560a424770a4b92b860673860f40a76e5b882b4b612b0bb9734f.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ie5192d934b9d560a424770a4b92b860673860f40a76e5b882b4b612b0bb9734f.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i7a0775e5e65a177503bff9b2741bd8904c74a9dd211fccd5baad69e5128048b7.GetMemberGroupsRequestBuilder) {
    return i7a0775e5e65a177503bff9b2741bd8904c74a9dd211fccd5baad69e5128048b7.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*ia8c58872e4e9fbad5c500292d61411649e61166d99a7d9b622aa1a3ad0961917.GetMemberObjectsRequestBuilder) {
    return ia8c58872e4e9fbad5c500292d61411649e61166d99a7d9b622aa1a3ad0961917.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i0fe5d4e20194af0ed86dcea25a2ff910bf31a3cf289832b920e4142fba0f1178.InvalidateAllRefreshTokensRequestBuilder) {
    return i0fe5d4e20194af0ed86dcea25a2ff910bf31a3cf289832b920e4142fba0f1178.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i23281cbd45ce85cdcf686f18da752247499c036cac07360733ae51f1d1c9008e.IsManagedAppUserBlockedRequestBuilder) {
    return i23281cbd45ce85cdcf686f18da752247499c036cac07360733ae51f1d1c9008e.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ia3b7a267bed5cc58f641ce43ffe0aa74f23557621a07fabd77bd28e9908bed9c.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ia3b7a267bed5cc58f641ce43ffe0aa74f23557621a07fabd77bd28e9908bed9c.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i0183a420e1adf7a0e5b4673d52e8c38584ecf6840a13407a1cb22aa1ab8467ff.RemoveAllDevicesFromManagementRequestBuilder) {
    return i0183a420e1adf7a0e5b4673d52e8c38584ecf6840a13407a1cb22aa1ab8467ff.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i8249d410e2eaf50e7df6fa1d23e25f1b6f81d9b4de4f1654706882e5b9efd862.ReprocessLicenseAssignmentRequestBuilder) {
    return i8249d410e2eaf50e7df6fa1d23e25f1b6f81d9b4de4f1654706882e5b9efd862.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i4122d581e5e42f864e7cc396199ba26fed8ccc9f15971e3daf2f70ba1a6befac.RestoreRequestBuilder) {
    return i4122d581e5e42f864e7cc396199ba26fed8ccc9f15971e3daf2f70ba1a6befac.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ic9064c196b1eafcb6a94734daf3b957b68ce46d0fef1407fb37459013ff2691a.RevokeSignInSessionsRequestBuilder) {
    return ic9064c196b1eafcb6a94734daf3b957b68ce46d0fef1407fb37459013ff2691a.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i878f1aaf7c19a4cafdb06363cf5f534d1e0a4cef10f6ff0fb5871b3a35c9ff05.SendMailRequestBuilder) {
    return i878f1aaf7c19a4cafdb06363cf5f534d1e0a4cef10f6ff0fb5871b3a35c9ff05.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ife4357fa284e05bcfcb62dcfac0bcc3a4d2196500d46bffceadfc5fb53bad962.TranslateExchangeIdsRequestBuilder) {
    return ife4357fa284e05bcfcb62dcfac0bcc3a4d2196500d46bffceadfc5fb53bad962.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i5c5ab6c54a14b07be3b65bd09dbb3211b7e6a302a0081b1c35674b8299a458c9.UnblockManagedAppsRequestBuilder) {
    return i5c5ab6c54a14b07be3b65bd09dbb3211b7e6a302a0081b1c35674b8299a458c9.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i415a6ee4477e053c8f254df656011677bab23bf424f8c6dd036d853346da676b.WipeAndBlockManagedAppsRequestBuilder) {
    return i415a6ee4477e053c8f254df656011677bab23bf424f8c6dd036d853346da676b.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*iaa1c6edacf5feba0b2f962a849221cf85d5b7c28d8aac5a97aa838fcf3219c92.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return iaa1c6edacf5feba0b2f962a849221cf85d5b7c28d8aac5a97aa838fcf3219c92.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i955ffc7b0625a4f2266d9d1a9a6bfa7afc228bf8c6f67beb231bbc7cc8c6787d.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i955ffc7b0625a4f2266d9d1a9a6bfa7afc228bf8c6f67beb231bbc7cc8c6787d.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*id8eff0d6a58a6bbe710a3158f2b444d51e014e83002bd032a0c6d417973f6d5d.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return id8eff0d6a58a6bbe710a3158f2b444d51e014e83002bd032a0c6d417973f6d5d.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
