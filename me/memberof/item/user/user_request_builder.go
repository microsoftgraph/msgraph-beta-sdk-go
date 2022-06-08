package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04f69096ea20c82b076d67e6c9ac3b2005800db89190abfca4159eca6ed708b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmailtips"
    i0eb638e890c2f4821c63501846aa9753fa419b4703325e04278e9c51fc6ac407 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/activateserviceplan"
    i137ac7bf7c237b7d0ebea5f2514a8d0f0fd91adb041fc20549cf04ec7d1f2444 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/findrooms"
    i15f9493405c92d8b621c4f12b942958a48f90fdeb58db5afde1a5924eb54a569 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/wipemanagedappregistrationbydevicetag"
    i192093a1c69a2c88448e55e7d5640a54d348249267d5094f4a57fe4364c675df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i237c56c2e3ccaf6106393da0c3056e647dc802d4cb44415b6846fa43374c7e81 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i29b2520d299bdb49c83f60457ce5852384295e7494537e41373638cc4606c2d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getloggedonmanageddevices"
    i2bd83ae8df1bd75b07f0f9ee94d31c824d5e7400dcfd39bad9848a640dda90df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/findroomswithroomlist"
    i2eb194b76ca8b2061eb6580927eb2edca5a8df1d067fb2b6337e76bc9e090c2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/wipeandblockmanagedapps"
    i312f3d2b33726588bd7a2c33ed665292abc3db86b983a0f274467fbd2eb860bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i3b2423eb3be093361fc253933a8f9f81d9d65da0d12db5e807e595c84ff44804 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/sendmail"
    i42e27637da18adac1525745eeb572541c21308e97968cdde4c436989db00be17 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i57feebd25825f2d2108096008b20cc3b5a84824a7de65b62ca992afa5797eea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmemberobjects"
    i60e0a2ceacc98e4bbc356e5cf8471edb99e5d7955a79fd639b85d49614c2b92e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/removealldevicesfrommanagement"
    i612060847aa6846933e0b4ab9de6c229715f9eb7f097988914f3950e98975bf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/changepassword"
    i68f0be02094234deec342e993fa8cb74ee63200e2f5ee650a4a843dbb525ee18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmanagedapppolicies"
    i6a78d90b156b5e06774b4ad0fb255baab45f164e37600a8713bcdca92b4b9b33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmanageddeviceswithappfailures"
    i747e3b11261d8a2fbbe047788fe6e69602cf594122d308a0e9031bff56a448db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmanagedappdiagnosticstatuses"
    i7a2d1e40dc75911d1aa7f023955998442ca8e83dc8c0cecdb4b03d8c84992850 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/revokesigninsessions"
    i7b2e88eb5605465ffc14badcd397e1801abc36852f86d4031c4787b7818c95af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i7ed667bded6a01e461e24989c803259039ed910ad9d631f64f910885efdbd3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/restore"
    i867e1878f084c403c832d327859a56643fe2d34a381a5e4c6fae9481a1eae38e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/exportpersonaldata"
    ia6a4261a9552750ac11f3bb1c063991f7939cbb0e35208216ba77b670ed78448 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    iad0350af18cb03d1bab250f89f456d08d44c635b19106b571580f0eeb675596c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/invalidateallrefreshtokens"
    ib2403534548506478260dace5c63b42995930243498011e20bee79cbfc9a2929 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/findmeetingtimes"
    ib2569d9c64bcd7b45a9587d3d35b6610f0e2eb51374a9102ae5a604327c3f255 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/findroomlists"
    ibf3d69bcaf7aa3b6ecd17dbd8fe2a0357675775d51a13ea86fff62bad01c0926 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/exportdeviceandappmanagementdata"
    ibf8553a5b99f5abe8cc3791bacbacac0fcb38dbdf427c6d580de581e0c897986 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/checkmemberobjects"
    ic1803b1e0d944d17ef4628975f6fa4b515e11d1524f2ab5925f31f966faf100e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/assignlicense"
    ic1ad7e97452f28d3e91ab3069a5bb17a8e10b5e3d72d2ec2d6ec9f2a65a1aadc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/reprocesslicenseassignment"
    idaa18ce63fdede09002fd548c8ba23936ee7e6dcc75633324b936727fd17863f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/translateexchangeids"
    ie135ca3d785af2221d31ce4697dbf3c59cb0d639ed3348098bc88cf4278a7e4d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/checkmembergroups"
    ied57a755a0401eedfe633ef23f3be688949c86809dc60498c025523b550e94d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/getmembergroups"
    iedd57176c0e31c1c24dad6ff4a2b71f59133d0b750b0d76f59e1982d5866c972 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/unblockmanagedapps"
    ifdb22f578ee7b417fcffceb95a541aa9cd78c9fdfe7ea9e78c243d6eaf9a1a35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item/user/ismanagedappuserblocked"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i0eb638e890c2f4821c63501846aa9753fa419b4703325e04278e9c51fc6ac407.ActivateServicePlanRequestBuilder) {
    return i0eb638e890c2f4821c63501846aa9753fa419b4703325e04278e9c51fc6ac407.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ic1803b1e0d944d17ef4628975f6fa4b515e11d1524f2ab5925f31f966faf100e.AssignLicenseRequestBuilder) {
    return ic1803b1e0d944d17ef4628975f6fa4b515e11d1524f2ab5925f31f966faf100e.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i612060847aa6846933e0b4ab9de6c229715f9eb7f097988914f3950e98975bf2.ChangePasswordRequestBuilder) {
    return i612060847aa6846933e0b4ab9de6c229715f9eb7f097988914f3950e98975bf2.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ie135ca3d785af2221d31ce4697dbf3c59cb0d639ed3348098bc88cf4278a7e4d.CheckMemberGroupsRequestBuilder) {
    return ie135ca3d785af2221d31ce4697dbf3c59cb0d639ed3348098bc88cf4278a7e4d.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ibf8553a5b99f5abe8cc3791bacbacac0fcb38dbdf427c6d580de581e0c897986.CheckMemberObjectsRequestBuilder) {
    return ibf8553a5b99f5abe8cc3791bacbacac0fcb38dbdf427c6d580de581e0c897986.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ibf3d69bcaf7aa3b6ecd17dbd8fe2a0357675775d51a13ea86fff62bad01c0926.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ibf3d69bcaf7aa3b6ecd17dbd8fe2a0357675775d51a13ea86fff62bad01c0926.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ia6a4261a9552750ac11f3bb1c063991f7939cbb0e35208216ba77b670ed78448.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ia6a4261a9552750ac11f3bb1c063991f7939cbb0e35208216ba77b670ed78448.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i867e1878f084c403c832d327859a56643fe2d34a381a5e4c6fae9481a1eae38e.ExportPersonalDataRequestBuilder) {
    return i867e1878f084c403c832d327859a56643fe2d34a381a5e4c6fae9481a1eae38e.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ib2403534548506478260dace5c63b42995930243498011e20bee79cbfc9a2929.FindMeetingTimesRequestBuilder) {
    return ib2403534548506478260dace5c63b42995930243498011e20bee79cbfc9a2929.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ib2569d9c64bcd7b45a9587d3d35b6610f0e2eb51374a9102ae5a604327c3f255.FindRoomListsRequestBuilder) {
    return ib2569d9c64bcd7b45a9587d3d35b6610f0e2eb51374a9102ae5a604327c3f255.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i137ac7bf7c237b7d0ebea5f2514a8d0f0fd91adb041fc20549cf04ec7d1f2444.FindRoomsRequestBuilder) {
    return i137ac7bf7c237b7d0ebea5f2514a8d0f0fd91adb041fc20549cf04ec7d1f2444.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i2bd83ae8df1bd75b07f0f9ee94d31c824d5e7400dcfd39bad9848a640dda90df.FindRoomsWithRoomListRequestBuilder) {
    return i2bd83ae8df1bd75b07f0f9ee94d31c824d5e7400dcfd39bad9848a640dda90df.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i42e27637da18adac1525745eeb572541c21308e97968cdde4c436989db00be17.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i42e27637da18adac1525745eeb572541c21308e97968cdde4c436989db00be17.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i29b2520d299bdb49c83f60457ce5852384295e7494537e41373638cc4606c2d4.GetLoggedOnManagedDevicesRequestBuilder) {
    return i29b2520d299bdb49c83f60457ce5852384295e7494537e41373638cc4606c2d4.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i04f69096ea20c82b076d67e6c9ac3b2005800db89190abfca4159eca6ed708b2.GetMailTipsRequestBuilder) {
    return i04f69096ea20c82b076d67e6c9ac3b2005800db89190abfca4159eca6ed708b2.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i747e3b11261d8a2fbbe047788fe6e69602cf594122d308a0e9031bff56a448db.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i747e3b11261d8a2fbbe047788fe6e69602cf594122d308a0e9031bff56a448db.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i68f0be02094234deec342e993fa8cb74ee63200e2f5ee650a4a843dbb525ee18.GetManagedAppPoliciesRequestBuilder) {
    return i68f0be02094234deec342e993fa8cb74ee63200e2f5ee650a4a843dbb525ee18.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i6a78d90b156b5e06774b4ad0fb255baab45f164e37600a8713bcdca92b4b9b33.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i6a78d90b156b5e06774b4ad0fb255baab45f164e37600a8713bcdca92b4b9b33.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i7b2e88eb5605465ffc14badcd397e1801abc36852f86d4031c4787b7818c95af.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i7b2e88eb5605465ffc14badcd397e1801abc36852f86d4031c4787b7818c95af.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ied57a755a0401eedfe633ef23f3be688949c86809dc60498c025523b550e94d5.GetMemberGroupsRequestBuilder) {
    return ied57a755a0401eedfe633ef23f3be688949c86809dc60498c025523b550e94d5.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i57feebd25825f2d2108096008b20cc3b5a84824a7de65b62ca992afa5797eea1.GetMemberObjectsRequestBuilder) {
    return i57feebd25825f2d2108096008b20cc3b5a84824a7de65b62ca992afa5797eea1.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*iad0350af18cb03d1bab250f89f456d08d44c635b19106b571580f0eeb675596c.InvalidateAllRefreshTokensRequestBuilder) {
    return iad0350af18cb03d1bab250f89f456d08d44c635b19106b571580f0eeb675596c.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ifdb22f578ee7b417fcffceb95a541aa9cd78c9fdfe7ea9e78c243d6eaf9a1a35.IsManagedAppUserBlockedRequestBuilder) {
    return ifdb22f578ee7b417fcffceb95a541aa9cd78c9fdfe7ea9e78c243d6eaf9a1a35.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i192093a1c69a2c88448e55e7d5640a54d348249267d5094f4a57fe4364c675df.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i192093a1c69a2c88448e55e7d5640a54d348249267d5094f4a57fe4364c675df.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i60e0a2ceacc98e4bbc356e5cf8471edb99e5d7955a79fd639b85d49614c2b92e.RemoveAllDevicesFromManagementRequestBuilder) {
    return i60e0a2ceacc98e4bbc356e5cf8471edb99e5d7955a79fd639b85d49614c2b92e.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ic1ad7e97452f28d3e91ab3069a5bb17a8e10b5e3d72d2ec2d6ec9f2a65a1aadc.ReprocessLicenseAssignmentRequestBuilder) {
    return ic1ad7e97452f28d3e91ab3069a5bb17a8e10b5e3d72d2ec2d6ec9f2a65a1aadc.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i7ed667bded6a01e461e24989c803259039ed910ad9d631f64f910885efdbd3d8.RestoreRequestBuilder) {
    return i7ed667bded6a01e461e24989c803259039ed910ad9d631f64f910885efdbd3d8.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i7a2d1e40dc75911d1aa7f023955998442ca8e83dc8c0cecdb4b03d8c84992850.RevokeSignInSessionsRequestBuilder) {
    return i7a2d1e40dc75911d1aa7f023955998442ca8e83dc8c0cecdb4b03d8c84992850.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i3b2423eb3be093361fc253933a8f9f81d9d65da0d12db5e807e595c84ff44804.SendMailRequestBuilder) {
    return i3b2423eb3be093361fc253933a8f9f81d9d65da0d12db5e807e595c84ff44804.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*idaa18ce63fdede09002fd548c8ba23936ee7e6dcc75633324b936727fd17863f.TranslateExchangeIdsRequestBuilder) {
    return idaa18ce63fdede09002fd548c8ba23936ee7e6dcc75633324b936727fd17863f.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*iedd57176c0e31c1c24dad6ff4a2b71f59133d0b750b0d76f59e1982d5866c972.UnblockManagedAppsRequestBuilder) {
    return iedd57176c0e31c1c24dad6ff4a2b71f59133d0b750b0d76f59e1982d5866c972.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i2eb194b76ca8b2061eb6580927eb2edca5a8df1d067fb2b6337e76bc9e090c2e.WipeAndBlockManagedAppsRequestBuilder) {
    return i2eb194b76ca8b2061eb6580927eb2edca5a8df1d067fb2b6337e76bc9e090c2e.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i15f9493405c92d8b621c4f12b942958a48f90fdeb58db5afde1a5924eb54a569.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i15f9493405c92d8b621c4f12b942958a48f90fdeb58db5afde1a5924eb54a569.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i312f3d2b33726588bd7a2c33ed665292abc3db86b983a0f274467fbd2eb860bc.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i312f3d2b33726588bd7a2c33ed665292abc3db86b983a0f274467fbd2eb860bc.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i237c56c2e3ccaf6106393da0c3056e647dc802d4cb44415b6846fa43374c7e81.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i237c56c2e3ccaf6106393da0c3056e647dc802d4cb44415b6846fa43374c7e81.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
