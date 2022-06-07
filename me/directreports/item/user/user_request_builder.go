package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00c41c946f0c00951f00e1da8121095ebdf8b92b07bbeece8490090664f30842 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/ismanagedappuserblocked"
    i0e591b2025e03757c14e411fc4baa34a6d0a0bf2fff61ab7df845efe5d515222 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/revokesigninsessions"
    i16c44cb9b55da931c7bceed6eb95047a7f9e70212e55bc746ab7a4781c1960d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/checkmembergroups"
    i1b0269c63efaf01042f515c85a3819eeb8628f9ae79fa834f20d3b7bfd889db5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getloggedonmanageddevices"
    i1b4768675cc601632bb58caea86cf7a5b9640ff1f2d1d66b77d30a2b0ccace1f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/reminderviewwithstartdatetimewithenddatetime"
    i25c4e18a964472b923482d1bcc64be5955e3be22d221eea68fbd50faf9549152 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/unblockmanagedapps"
    i2bd69a9f4f874d431b6be9e7ac714fa877a1082f2aee92aa32494dfbc0408b34 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/findroomlists"
    i3447454b6822a9591bf982395438726636eff1ddd675c3dbae492a11c63cf97f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmanageddeviceswithfailedorpendingapps"
    i3c65c9ce35b7c3cf16d2f21fb38f9904e1e9761d708cdabe44ea7f0f60d66dfe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmanagedappdiagnosticstatuses"
    i4126f0fc7499872789a2c2fef29e3d8d38c036820cd87f11df112449cee13d1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/wipemanagedappregistrationsbydevicetag"
    i45613c159ee6c5070c44353e8fc0bc5818525a87d0f40a75ac89fbae3b16d80b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/geteffectivedeviceenrollmentconfigurations"
    i504c41245ac6ae9f13513059057db8aa8276947a25312c2ddb6d9119a3ad80b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i5991fc4cd941cfd6872c231378f0459c285f5cee42aa367585cd4a7b956a0239 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/exportpersonaldata"
    i6a2f15bd00d89f5df3f6d2bc511c0a2b76825fa890940e0e9ac2131cf8b7b660 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmanagedapppolicies"
    i6b2abd5e4365e72c22649d909de7e54b99a09c744a2ddd947b467855912b0fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/exportdeviceandappmanagementdata"
    i6d5a538ec82e9ca580bb228358f7fce6e7965308755b2d63209e44522bf4b3f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i707019f676fd3e0ce04699f813b2edfb1a486dc979c887231ae0d169bdd8cf5d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/activateserviceplan"
    i7442aff021da5ef0c9ec6c6acf443086645c4e455209637fe042c9c514562e24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/wipeandblockmanagedapps"
    i7d14b8783de938375520176e33cad78131ce68d3726c6db29df738dbd58e6e79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/translateexchangeids"
    i7e20d2be1fc1e7674dc94afcc9fdfff42354360828b8b3cfa35eed581e3d90df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmemberobjects"
    i8096e357aa3cffef62e324f7e5368f728517d2690389bad8847b921859115f3f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/sendmail"
    i8532e20f8d1b5ab7393cae2f670a789210719e638a64e7d33e5871f0c49ec676 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/restore"
    i96ea445a3e5f7c90de5bba3a0db29d83eb3ce0eb457d4052a1f87757e933eb4a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/invalidateallrefreshtokens"
    ia3e504ea4464264f63566a89fcee327934ca95a315fa9449a986492f1e0b3f88 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/wipemanagedappregistrationbydevicetag"
    ib67859a697dd537498f182d81811392bdadf90a76b6e62ffc2e8fc99144dd721 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/assignlicense"
    ic307cfb1f1a88dd2620a58e352a1beb254fb124bbc5c3890fae98cf38c49a0ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/changepassword"
    icc049e701dac9937b99b45e45bc22d2e186f3c8c964dfa8a0850181e77a7a399 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/removealldevicesfrommanagement"
    id31ab1ea06954aeddd2102e5adf70f8d43cf7f964e426373a6b99cf3e912b171 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/findmeetingtimes"
    id5580b5545b0276af3983f9e3bf2e79efa1c0578db80ae4681701ff0f9cf4b8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/reprocesslicenseassignment"
    idd6219cc165ac6e93dbdc7c0e1e4129f998f69b85694262f69986e473b315061 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/findrooms"
    ie03de28fcff41b2339e5fd96e80422af7704e91866c05e7fe0784609b415146f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/findroomswithroomlist"
    ie59e1b6209fc0e1299d0b028639ba7cb16dc1443a8e214a8c0a1457f67a746cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmailtips"
    if1c4830260398f102aa2494d96c53250e658a3792ad9777a9619de11d430ae91 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmanageddeviceswithappfailures"
    if796d43b5282ec23e6e91cb7d818078f9d61d8207df5c6b8504cc12a6a88c45d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/checkmemberobjects"
    ifcb308f3696c822cab7069c7164ad52cdc65a476be9f184f595bc2fc3a645306 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item/user/getmembergroups"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i707019f676fd3e0ce04699f813b2edfb1a486dc979c887231ae0d169bdd8cf5d.ActivateServicePlanRequestBuilder) {
    return i707019f676fd3e0ce04699f813b2edfb1a486dc979c887231ae0d169bdd8cf5d.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ib67859a697dd537498f182d81811392bdadf90a76b6e62ffc2e8fc99144dd721.AssignLicenseRequestBuilder) {
    return ib67859a697dd537498f182d81811392bdadf90a76b6e62ffc2e8fc99144dd721.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ic307cfb1f1a88dd2620a58e352a1beb254fb124bbc5c3890fae98cf38c49a0ce.ChangePasswordRequestBuilder) {
    return ic307cfb1f1a88dd2620a58e352a1beb254fb124bbc5c3890fae98cf38c49a0ce.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i16c44cb9b55da931c7bceed6eb95047a7f9e70212e55bc746ab7a4781c1960d5.CheckMemberGroupsRequestBuilder) {
    return i16c44cb9b55da931c7bceed6eb95047a7f9e70212e55bc746ab7a4781c1960d5.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*if796d43b5282ec23e6e91cb7d818078f9d61d8207df5c6b8504cc12a6a88c45d.CheckMemberObjectsRequestBuilder) {
    return if796d43b5282ec23e6e91cb7d818078f9d61d8207df5c6b8504cc12a6a88c45d.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/directReports/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i6b2abd5e4365e72c22649d909de7e54b99a09c744a2ddd947b467855912b0fc1.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i6b2abd5e4365e72c22649d909de7e54b99a09c744a2ddd947b467855912b0fc1.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i6d5a538ec82e9ca580bb228358f7fce6e7965308755b2d63209e44522bf4b3f6.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i6d5a538ec82e9ca580bb228358f7fce6e7965308755b2d63209e44522bf4b3f6.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i5991fc4cd941cfd6872c231378f0459c285f5cee42aa367585cd4a7b956a0239.ExportPersonalDataRequestBuilder) {
    return i5991fc4cd941cfd6872c231378f0459c285f5cee42aa367585cd4a7b956a0239.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*id31ab1ea06954aeddd2102e5adf70f8d43cf7f964e426373a6b99cf3e912b171.FindMeetingTimesRequestBuilder) {
    return id31ab1ea06954aeddd2102e5adf70f8d43cf7f964e426373a6b99cf3e912b171.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i2bd69a9f4f874d431b6be9e7ac714fa877a1082f2aee92aa32494dfbc0408b34.FindRoomListsRequestBuilder) {
    return i2bd69a9f4f874d431b6be9e7ac714fa877a1082f2aee92aa32494dfbc0408b34.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*idd6219cc165ac6e93dbdc7c0e1e4129f998f69b85694262f69986e473b315061.FindRoomsRequestBuilder) {
    return idd6219cc165ac6e93dbdc7c0e1e4129f998f69b85694262f69986e473b315061.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ie03de28fcff41b2339e5fd96e80422af7704e91866c05e7fe0784609b415146f.FindRoomsWithRoomListRequestBuilder) {
    return ie03de28fcff41b2339e5fd96e80422af7704e91866c05e7fe0784609b415146f.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i45613c159ee6c5070c44353e8fc0bc5818525a87d0f40a75ac89fbae3b16d80b.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i45613c159ee6c5070c44353e8fc0bc5818525a87d0f40a75ac89fbae3b16d80b.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i1b0269c63efaf01042f515c85a3819eeb8628f9ae79fa834f20d3b7bfd889db5.GetLoggedOnManagedDevicesRequestBuilder) {
    return i1b0269c63efaf01042f515c85a3819eeb8628f9ae79fa834f20d3b7bfd889db5.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ie59e1b6209fc0e1299d0b028639ba7cb16dc1443a8e214a8c0a1457f67a746cd.GetMailTipsRequestBuilder) {
    return ie59e1b6209fc0e1299d0b028639ba7cb16dc1443a8e214a8c0a1457f67a746cd.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i3c65c9ce35b7c3cf16d2f21fb38f9904e1e9761d708cdabe44ea7f0f60d66dfe.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i3c65c9ce35b7c3cf16d2f21fb38f9904e1e9761d708cdabe44ea7f0f60d66dfe.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i6a2f15bd00d89f5df3f6d2bc511c0a2b76825fa890940e0e9ac2131cf8b7b660.GetManagedAppPoliciesRequestBuilder) {
    return i6a2f15bd00d89f5df3f6d2bc511c0a2b76825fa890940e0e9ac2131cf8b7b660.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*if1c4830260398f102aa2494d96c53250e658a3792ad9777a9619de11d430ae91.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return if1c4830260398f102aa2494d96c53250e658a3792ad9777a9619de11d430ae91.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i3447454b6822a9591bf982395438726636eff1ddd675c3dbae492a11c63cf97f.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i3447454b6822a9591bf982395438726636eff1ddd675c3dbae492a11c63cf97f.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ifcb308f3696c822cab7069c7164ad52cdc65a476be9f184f595bc2fc3a645306.GetMemberGroupsRequestBuilder) {
    return ifcb308f3696c822cab7069c7164ad52cdc65a476be9f184f595bc2fc3a645306.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i7e20d2be1fc1e7674dc94afcc9fdfff42354360828b8b3cfa35eed581e3d90df.GetMemberObjectsRequestBuilder) {
    return i7e20d2be1fc1e7674dc94afcc9fdfff42354360828b8b3cfa35eed581e3d90df.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i96ea445a3e5f7c90de5bba3a0db29d83eb3ce0eb457d4052a1f87757e933eb4a.InvalidateAllRefreshTokensRequestBuilder) {
    return i96ea445a3e5f7c90de5bba3a0db29d83eb3ce0eb457d4052a1f87757e933eb4a.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i00c41c946f0c00951f00e1da8121095ebdf8b92b07bbeece8490090664f30842.IsManagedAppUserBlockedRequestBuilder) {
    return i00c41c946f0c00951f00e1da8121095ebdf8b92b07bbeece8490090664f30842.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i1b4768675cc601632bb58caea86cf7a5b9640ff1f2d1d66b77d30a2b0ccace1f.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1b4768675cc601632bb58caea86cf7a5b9640ff1f2d1d66b77d30a2b0ccace1f.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*icc049e701dac9937b99b45e45bc22d2e186f3c8c964dfa8a0850181e77a7a399.RemoveAllDevicesFromManagementRequestBuilder) {
    return icc049e701dac9937b99b45e45bc22d2e186f3c8c964dfa8a0850181e77a7a399.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*id5580b5545b0276af3983f9e3bf2e79efa1c0578db80ae4681701ff0f9cf4b8d.ReprocessLicenseAssignmentRequestBuilder) {
    return id5580b5545b0276af3983f9e3bf2e79efa1c0578db80ae4681701ff0f9cf4b8d.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i8532e20f8d1b5ab7393cae2f670a789210719e638a64e7d33e5871f0c49ec676.RestoreRequestBuilder) {
    return i8532e20f8d1b5ab7393cae2f670a789210719e638a64e7d33e5871f0c49ec676.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i0e591b2025e03757c14e411fc4baa34a6d0a0bf2fff61ab7df845efe5d515222.RevokeSignInSessionsRequestBuilder) {
    return i0e591b2025e03757c14e411fc4baa34a6d0a0bf2fff61ab7df845efe5d515222.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i8096e357aa3cffef62e324f7e5368f728517d2690389bad8847b921859115f3f.SendMailRequestBuilder) {
    return i8096e357aa3cffef62e324f7e5368f728517d2690389bad8847b921859115f3f.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i7d14b8783de938375520176e33cad78131ce68d3726c6db29df738dbd58e6e79.TranslateExchangeIdsRequestBuilder) {
    return i7d14b8783de938375520176e33cad78131ce68d3726c6db29df738dbd58e6e79.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i25c4e18a964472b923482d1bcc64be5955e3be22d221eea68fbd50faf9549152.UnblockManagedAppsRequestBuilder) {
    return i25c4e18a964472b923482d1bcc64be5955e3be22d221eea68fbd50faf9549152.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i7442aff021da5ef0c9ec6c6acf443086645c4e455209637fe042c9c514562e24.WipeAndBlockManagedAppsRequestBuilder) {
    return i7442aff021da5ef0c9ec6c6acf443086645c4e455209637fe042c9c514562e24.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ia3e504ea4464264f63566a89fcee327934ca95a315fa9449a986492f1e0b3f88.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ia3e504ea4464264f63566a89fcee327934ca95a315fa9449a986492f1e0b3f88.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i504c41245ac6ae9f13513059057db8aa8276947a25312c2ddb6d9119a3ad80b7.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i504c41245ac6ae9f13513059057db8aa8276947a25312c2ddb6d9119a3ad80b7.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i4126f0fc7499872789a2c2fef29e3d8d38c036820cd87f11df112449cee13d1e.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i4126f0fc7499872789a2c2fef29e3d8d38c036820cd87f11df112449cee13d1e.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
