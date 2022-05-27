package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i008df3386410b86417773e46ba1e607b77b2336b4a65b8f2c1c0ac90504027eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmailtips"
    i0a6cdf7daef18067ec07a5bcce7f1781b26c4b9e8ca5ee266de87ec548286b67 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/findroomlists"
    i0ebf5023985a6c721deb30d6f5ee1ab147df64049bad80e9665a0cc55b24e652 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i157bca2f1939b96195fcb3f444d196b2fadce0448b6b7be996e1160d03000b97 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmembergroups"
    i183780b15802d6826e2e5d30d08709e61adce478ec603d1e9be278dbb9078a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmanagedapppolicies"
    i1d5a78ac87d3f55464d53ce3fc701b08e535dff334c3ec3dba63ac56bdc20843 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/findrooms"
    i31aecab1a2dbe9d7af59df7d0871356c71abba239aa25552ec79e6b50b87ba6e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/findmeetingtimes"
    i3cf4e00db271c0629a4f22b03b19d22a989284673326df091c25c0362acea50b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i44a1602b893f22bbdf43f7065416b981ea2d51b9c49bdc2ebc653a1bcfcefab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/checkmembergroups"
    i4e686d2f1119df3078438beebaf977a5213f075df380776bd7328e98b51d31a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/restore"
    i4ea9e0c06586127963c200cd0949ee86634e2a9785be397f6b16d18713dba929 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/assignlicense"
    i535e5789b5273a3e7c62ea793aba9d5babbc6b59b3b4ad2f4e4fea378b3da49f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getloggedonmanageddevices"
    i568bc0517188d10f43d0e30889a23d6bfbd211800629fb37defb0d51f26423d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/removealldevicesfrommanagement"
    i69d977b8e44b112a529f0ebd08708d5cb2efda279067963fdf4fd35dcace5581 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i6e34c3cfc7d589543c47744020f820b4061b80f3b7640888fa6939cd670072c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmemberobjects"
    i71e8e7bc37c9860836c92558264918ac93d4c1c852841e7f4841266fbdcf739a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/reprocesslicenseassignment"
    i72b1a294669b390a34f9438ad29018706d6db7a6c62781ff46b7d2b74f7fbb8a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/sendmail"
    i7b08a35513b304e5b64bf5e468c3f0c2c28f51990aa49adb5a05a6f8452e5872 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/checkmemberobjects"
    i869a214f8b079ad659d219f77a55e40f9ecacb3822f61004b913bf4841fae618 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/activateserviceplan"
    i90bcd575a1e89d52eb85a15a01eaf34e37a93fe5d0c4804e0649e49d40469997 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/findroomswithroomlist"
    i98f61688adbda43b0a211651d070ceb1df2b9c8a06b79fc11f42ddabd8e97d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i99e94a0a5e2b89b6b5b82d15ee41ac769d405e5e87621a9c56ddea42354209b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/getmanageddeviceswithappfailures"
    ia76374a60cb5575abf40354a0c9be495eebc629e090522f4a28072baeaef0c8c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    iab2db6e09cc48ae98571df82b8315ed5b9c6cab3898a47666762019d01449183 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/translateexchangeids"
    ib31ebac869aa7b521ddbf137a357323b623f06ca533966d0e0a96395d5bd4f97 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/wipeandblockmanagedapps"
    ib3dcb746380a375f993afc51f5a848e2f77cf082febfba67359bd88eb4e42ec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/invalidateallrefreshtokens"
    ibb8afb2a147beca7e71eb0f03057e986486543cca73aef2388ee9ecf665d28dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/exportpersonaldata"
    icc419ce3e510c97bd43268868fb28840d213120fef14de32a653d36fdc350c50 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/revokesigninsessions"
    ice7f634221f70fdd24bd5c418494fff2b0a8b38439df02a1e05b8d17448f873c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/changepassword"
    id53f20c20ad24a299ad92202ab5ac5740704e51eeb0814a4b6bd945fbeb34384 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/unblockmanagedapps"
    ie48ecebb14ab255f30dbf755cc29e7878eadc805baf9b3d60606e7dd5722e226 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    ie58a2ea46eafb2acda3d072e78481d222ecde20f55d7694e5fbb294ed56696ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    if0f58b97e10b3106282cc6b9e777d60f5c4df9c09b95308ec0bb9c0b5ca69403 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/exportdeviceandappmanagementdata"
    if3a21b0f5ef7abb01fdb63c49b59e1d424c6ce59435fc365f6207e52d11c11cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    if8db63faa154ae934393be9305501f8727743c2d7545e901308887b0e78033cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item/user/ismanagedappuserblocked"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i869a214f8b079ad659d219f77a55e40f9ecacb3822f61004b913bf4841fae618.ActivateServicePlanRequestBuilder) {
    return i869a214f8b079ad659d219f77a55e40f9ecacb3822f61004b913bf4841fae618.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i4ea9e0c06586127963c200cd0949ee86634e2a9785be397f6b16d18713dba929.AssignLicenseRequestBuilder) {
    return i4ea9e0c06586127963c200cd0949ee86634e2a9785be397f6b16d18713dba929.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ice7f634221f70fdd24bd5c418494fff2b0a8b38439df02a1e05b8d17448f873c.ChangePasswordRequestBuilder) {
    return ice7f634221f70fdd24bd5c418494fff2b0a8b38439df02a1e05b8d17448f873c.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i44a1602b893f22bbdf43f7065416b981ea2d51b9c49bdc2ebc653a1bcfcefab5.CheckMemberGroupsRequestBuilder) {
    return i44a1602b893f22bbdf43f7065416b981ea2d51b9c49bdc2ebc653a1bcfcefab5.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i7b08a35513b304e5b64bf5e468c3f0c2c28f51990aa49adb5a05a6f8452e5872.CheckMemberObjectsRequestBuilder) {
    return i7b08a35513b304e5b64bf5e468c3f0c2c28f51990aa49adb5a05a6f8452e5872.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*if0f58b97e10b3106282cc6b9e777d60f5c4df9c09b95308ec0bb9c0b5ca69403.ExportDeviceAndAppManagementDataRequestBuilder) {
    return if0f58b97e10b3106282cc6b9e777d60f5c4df9c09b95308ec0bb9c0b5ca69403.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*if3a21b0f5ef7abb01fdb63c49b59e1d424c6ce59435fc365f6207e52d11c11cc.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return if3a21b0f5ef7abb01fdb63c49b59e1d424c6ce59435fc365f6207e52d11c11cc.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ibb8afb2a147beca7e71eb0f03057e986486543cca73aef2388ee9ecf665d28dd.ExportPersonalDataRequestBuilder) {
    return ibb8afb2a147beca7e71eb0f03057e986486543cca73aef2388ee9ecf665d28dd.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i31aecab1a2dbe9d7af59df7d0871356c71abba239aa25552ec79e6b50b87ba6e.FindMeetingTimesRequestBuilder) {
    return i31aecab1a2dbe9d7af59df7d0871356c71abba239aa25552ec79e6b50b87ba6e.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i0a6cdf7daef18067ec07a5bcce7f1781b26c4b9e8ca5ee266de87ec548286b67.FindRoomListsRequestBuilder) {
    return i0a6cdf7daef18067ec07a5bcce7f1781b26c4b9e8ca5ee266de87ec548286b67.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i1d5a78ac87d3f55464d53ce3fc701b08e535dff334c3ec3dba63ac56bdc20843.FindRoomsRequestBuilder) {
    return i1d5a78ac87d3f55464d53ce3fc701b08e535dff334c3ec3dba63ac56bdc20843.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i90bcd575a1e89d52eb85a15a01eaf34e37a93fe5d0c4804e0649e49d40469997.FindRoomsWithRoomListRequestBuilder) {
    return i90bcd575a1e89d52eb85a15a01eaf34e37a93fe5d0c4804e0649e49d40469997.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i69d977b8e44b112a529f0ebd08708d5cb2efda279067963fdf4fd35dcace5581.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i69d977b8e44b112a529f0ebd08708d5cb2efda279067963fdf4fd35dcace5581.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i535e5789b5273a3e7c62ea793aba9d5babbc6b59b3b4ad2f4e4fea378b3da49f.GetLoggedOnManagedDevicesRequestBuilder) {
    return i535e5789b5273a3e7c62ea793aba9d5babbc6b59b3b4ad2f4e4fea378b3da49f.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i008df3386410b86417773e46ba1e607b77b2336b4a65b8f2c1c0ac90504027eb.GetMailTipsRequestBuilder) {
    return i008df3386410b86417773e46ba1e607b77b2336b4a65b8f2c1c0ac90504027eb.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i98f61688adbda43b0a211651d070ceb1df2b9c8a06b79fc11f42ddabd8e97d11.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i98f61688adbda43b0a211651d070ceb1df2b9c8a06b79fc11f42ddabd8e97d11.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i183780b15802d6826e2e5d30d08709e61adce478ec603d1e9be278dbb9078a89.GetManagedAppPoliciesRequestBuilder) {
    return i183780b15802d6826e2e5d30d08709e61adce478ec603d1e9be278dbb9078a89.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i99e94a0a5e2b89b6b5b82d15ee41ac769d405e5e87621a9c56ddea42354209b3.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i99e94a0a5e2b89b6b5b82d15ee41ac769d405e5e87621a9c56ddea42354209b3.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i3cf4e00db271c0629a4f22b03b19d22a989284673326df091c25c0362acea50b.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i3cf4e00db271c0629a4f22b03b19d22a989284673326df091c25c0362acea50b.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i157bca2f1939b96195fcb3f444d196b2fadce0448b6b7be996e1160d03000b97.GetMemberGroupsRequestBuilder) {
    return i157bca2f1939b96195fcb3f444d196b2fadce0448b6b7be996e1160d03000b97.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i6e34c3cfc7d589543c47744020f820b4061b80f3b7640888fa6939cd670072c4.GetMemberObjectsRequestBuilder) {
    return i6e34c3cfc7d589543c47744020f820b4061b80f3b7640888fa6939cd670072c4.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ib3dcb746380a375f993afc51f5a848e2f77cf082febfba67359bd88eb4e42ec6.InvalidateAllRefreshTokensRequestBuilder) {
    return ib3dcb746380a375f993afc51f5a848e2f77cf082febfba67359bd88eb4e42ec6.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*if8db63faa154ae934393be9305501f8727743c2d7545e901308887b0e78033cb.IsManagedAppUserBlockedRequestBuilder) {
    return if8db63faa154ae934393be9305501f8727743c2d7545e901308887b0e78033cb.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ie58a2ea46eafb2acda3d072e78481d222ecde20f55d7694e5fbb294ed56696ee.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ie58a2ea46eafb2acda3d072e78481d222ecde20f55d7694e5fbb294ed56696ee.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i568bc0517188d10f43d0e30889a23d6bfbd211800629fb37defb0d51f26423d5.RemoveAllDevicesFromManagementRequestBuilder) {
    return i568bc0517188d10f43d0e30889a23d6bfbd211800629fb37defb0d51f26423d5.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i71e8e7bc37c9860836c92558264918ac93d4c1c852841e7f4841266fbdcf739a.ReprocessLicenseAssignmentRequestBuilder) {
    return i71e8e7bc37c9860836c92558264918ac93d4c1c852841e7f4841266fbdcf739a.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i4e686d2f1119df3078438beebaf977a5213f075df380776bd7328e98b51d31a9.RestoreRequestBuilder) {
    return i4e686d2f1119df3078438beebaf977a5213f075df380776bd7328e98b51d31a9.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*icc419ce3e510c97bd43268868fb28840d213120fef14de32a653d36fdc350c50.RevokeSignInSessionsRequestBuilder) {
    return icc419ce3e510c97bd43268868fb28840d213120fef14de32a653d36fdc350c50.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i72b1a294669b390a34f9438ad29018706d6db7a6c62781ff46b7d2b74f7fbb8a.SendMailRequestBuilder) {
    return i72b1a294669b390a34f9438ad29018706d6db7a6c62781ff46b7d2b74f7fbb8a.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*iab2db6e09cc48ae98571df82b8315ed5b9c6cab3898a47666762019d01449183.TranslateExchangeIdsRequestBuilder) {
    return iab2db6e09cc48ae98571df82b8315ed5b9c6cab3898a47666762019d01449183.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*id53f20c20ad24a299ad92202ab5ac5740704e51eeb0814a4b6bd945fbeb34384.UnblockManagedAppsRequestBuilder) {
    return id53f20c20ad24a299ad92202ab5ac5740704e51eeb0814a4b6bd945fbeb34384.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ib31ebac869aa7b521ddbf137a357323b623f06ca533966d0e0a96395d5bd4f97.WipeAndBlockManagedAppsRequestBuilder) {
    return ib31ebac869aa7b521ddbf137a357323b623f06ca533966d0e0a96395d5bd4f97.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ie48ecebb14ab255f30dbf755cc29e7878eadc805baf9b3d60606e7dd5722e226.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ie48ecebb14ab255f30dbf755cc29e7878eadc805baf9b3d60606e7dd5722e226.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ia76374a60cb5575abf40354a0c9be495eebc629e090522f4a28072baeaef0c8c.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ia76374a60cb5575abf40354a0c9be495eebc629e090522f4a28072baeaef0c8c.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i0ebf5023985a6c721deb30d6f5ee1ab147df64049bad80e9665a0cc55b24e652.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i0ebf5023985a6c721deb30d6f5ee1ab147df64049bad80e9665a0cc55b24e652.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
