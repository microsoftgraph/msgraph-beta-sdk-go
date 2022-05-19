package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06d0005d3e08114f086e40ae9aaf9014c1e2941b3d34d1e259f2a4721272cfb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/findroomlists"
    i08f093784082c7584eeed80410dfc43665e66bcc92b806fd50a0a49bbab53186 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/checkmemberobjects"
    i1f5c9f54c8e95106b91029db78cda487b3dec35fb5239e2f564c6b9606ec63d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i200c7a1309dda0069d2b78fb19e9c6e3da98aee255a2bc2358e633f761900587 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i27f07a50ed11fba9e9e80d411440818a5915243bcccb36d6ef6571170ba41797 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmanagedapppolicies"
    i3386910e8e3466365c88adc1df0449021faae977b5810dcfb2d59bc7b522c51d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/findroomswithroomlist"
    i38425d216794182cc2595d32e83b3210dd8a1ed5be3801bdcaa576426ec59862 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i3986dad4e0bec57f28b25bea54c87765eb5e38fd4756fec1b7554332f2d340d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/activateserviceplan"
    i3cefcfbeebdeb5b01558f53bfcf5423bd8198c799addae7b1c358b74a1eb4b07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmailtips"
    i40dcc81862a7014d06aa01e7791db4a4d2ddb9b651c7d2e5752d05c37dfa8c55 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/checkmembergroups"
    i53f1d35eb6c18d449e7731e30f543801df23913d9aba47de3bb4cb45e3cf6322 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i55769d9ada2dc3b28e70e204ffcb814c7e3e9d60538577117ecec1d3413f4e0c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmemberobjects"
    i5c641ed6f058403b9745f7af82c8f579fc9ddda6279ac87844e9d484f19379ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i5e5177f323c1ffe326c78295fe87257ed907912be0698cac966060b1b783a5d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/restore"
    i5eb63b72fe4ce7f0bef8cbddc22543544e7ca81edf1d2c394c322c476d45bfbb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/sendmail"
    i6de15caeb19193255bb6ba00799f713d6f487bcfa01ed44801e993c125abda6e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/exportdeviceandappmanagementdata"
    i80fbd8bbbcd3902c4c6d5082c7fa00ef8e367fa1bf5f340d648c9f3693ce9028 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/reprocesslicenseassignment"
    i840a5df5b41b629b2b2e5e29f1ece1065bba6e78c7014aac42c9d98d5975d01c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/findrooms"
    i87013f21fe46569b7386414cf5e93cd03add1fd237420df1e7c0d179b4a845f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/invalidateallrefreshtokens"
    i870ee559c28e4572280efa997e81663ff0b4fc26ccfb4a53c5851d16bcebc702 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i895b3abd81f960bf723a0f694ee3e19c780bc68dfcabf3fce2e7c9026bb1942c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmembergroups"
    i9ce5877f66d2d1ce704ec30ff4102e6f1f1ccc2678e6aedc8e026d5c1815caba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/ismanagedappuserblocked"
    ia6ed073c6b7138ce66dec87804768a570c6ced7a521acb5f77e0ef61aefb598b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/findmeetingtimes"
    ia77378d28cf70192f029ce5fe7633f8d0b6708a4c8ee1cee42366173ee23dc4c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ia9e64eb2592cd9b16347e2de58aac807cb77d1c0ddc01008b92b579f9845c467 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/unblockmanagedapps"
    ib072ba211a1fdcba5cf8ea7750f2410f13357828847ece6371159df1045d8ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/translateexchangeids"
    ib3c8cd942d46320c2fd34fb454791265ab3b207a8dc76dfa1f7094b3645a887f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getmanageddeviceswithappfailures"
    ib4a698ac8bb61bfc7884ad0f895c09a65997b9f50bd85dca996044e47860b769 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/wipeandblockmanagedapps"
    ibbd4d98e35c0cee0aa44fc7ec04e49a88fe59c1f2125f2679a868e21bc2e65f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/assignlicense"
    ic7da50c897bd1372fd8d351b29491e44d9ffe2965d602aa59f592072885b9fa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/changepassword"
    icaff82e6d54bc0182d42a26cb2b2d91be27cea04ab937cdd725e2f6d9959c09f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    ida58a91c692dc299466bf0bb1c96c0f9c3d4d5620c9542577426712cc3038232 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/exportpersonaldata"
    ie162ceb05fe3c55b816533d7ed757a034d956b762a08778c8199cb6f6459e1a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/getloggedonmanageddevices"
    ie7c52783bf847613ebaf73275f705022a0fc262866c43f495f384354ef2f5c32 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/removealldevicesfrommanagement"
    ie8e00017f331494e3c3b4b3ea69df62019ab36f641198bb8cb530a216f245e1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/transitivememberof/item/user/revokesigninsessions"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i3986dad4e0bec57f28b25bea54c87765eb5e38fd4756fec1b7554332f2d340d5.ActivateServicePlanRequestBuilder) {
    return i3986dad4e0bec57f28b25bea54c87765eb5e38fd4756fec1b7554332f2d340d5.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ibbd4d98e35c0cee0aa44fc7ec04e49a88fe59c1f2125f2679a868e21bc2e65f5.AssignLicenseRequestBuilder) {
    return ibbd4d98e35c0cee0aa44fc7ec04e49a88fe59c1f2125f2679a868e21bc2e65f5.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ic7da50c897bd1372fd8d351b29491e44d9ffe2965d602aa59f592072885b9fa2.ChangePasswordRequestBuilder) {
    return ic7da50c897bd1372fd8d351b29491e44d9ffe2965d602aa59f592072885b9fa2.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i40dcc81862a7014d06aa01e7791db4a4d2ddb9b651c7d2e5752d05c37dfa8c55.CheckMemberGroupsRequestBuilder) {
    return i40dcc81862a7014d06aa01e7791db4a4d2ddb9b651c7d2e5752d05c37dfa8c55.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i08f093784082c7584eeed80410dfc43665e66bcc92b806fd50a0a49bbab53186.CheckMemberObjectsRequestBuilder) {
    return i08f093784082c7584eeed80410dfc43665e66bcc92b806fd50a0a49bbab53186.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i6de15caeb19193255bb6ba00799f713d6f487bcfa01ed44801e993c125abda6e.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i6de15caeb19193255bb6ba00799f713d6f487bcfa01ed44801e993c125abda6e.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i5c641ed6f058403b9745f7af82c8f579fc9ddda6279ac87844e9d484f19379ff.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i5c641ed6f058403b9745f7af82c8f579fc9ddda6279ac87844e9d484f19379ff.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ida58a91c692dc299466bf0bb1c96c0f9c3d4d5620c9542577426712cc3038232.ExportPersonalDataRequestBuilder) {
    return ida58a91c692dc299466bf0bb1c96c0f9c3d4d5620c9542577426712cc3038232.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ia6ed073c6b7138ce66dec87804768a570c6ced7a521acb5f77e0ef61aefb598b.FindMeetingTimesRequestBuilder) {
    return ia6ed073c6b7138ce66dec87804768a570c6ced7a521acb5f77e0ef61aefb598b.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i06d0005d3e08114f086e40ae9aaf9014c1e2941b3d34d1e259f2a4721272cfb5.FindRoomListsRequestBuilder) {
    return i06d0005d3e08114f086e40ae9aaf9014c1e2941b3d34d1e259f2a4721272cfb5.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i840a5df5b41b629b2b2e5e29f1ece1065bba6e78c7014aac42c9d98d5975d01c.FindRoomsRequestBuilder) {
    return i840a5df5b41b629b2b2e5e29f1ece1065bba6e78c7014aac42c9d98d5975d01c.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i3386910e8e3466365c88adc1df0449021faae977b5810dcfb2d59bc7b522c51d.FindRoomsWithRoomListRequestBuilder) {
    return i3386910e8e3466365c88adc1df0449021faae977b5810dcfb2d59bc7b522c51d.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*icaff82e6d54bc0182d42a26cb2b2d91be27cea04ab937cdd725e2f6d9959c09f.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return icaff82e6d54bc0182d42a26cb2b2d91be27cea04ab937cdd725e2f6d9959c09f.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ie162ceb05fe3c55b816533d7ed757a034d956b762a08778c8199cb6f6459e1a5.GetLoggedOnManagedDevicesRequestBuilder) {
    return ie162ceb05fe3c55b816533d7ed757a034d956b762a08778c8199cb6f6459e1a5.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i3cefcfbeebdeb5b01558f53bfcf5423bd8198c799addae7b1c358b74a1eb4b07.GetMailTipsRequestBuilder) {
    return i3cefcfbeebdeb5b01558f53bfcf5423bd8198c799addae7b1c358b74a1eb4b07.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i200c7a1309dda0069d2b78fb19e9c6e3da98aee255a2bc2358e633f761900587.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i200c7a1309dda0069d2b78fb19e9c6e3da98aee255a2bc2358e633f761900587.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i27f07a50ed11fba9e9e80d411440818a5915243bcccb36d6ef6571170ba41797.GetManagedAppPoliciesRequestBuilder) {
    return i27f07a50ed11fba9e9e80d411440818a5915243bcccb36d6ef6571170ba41797.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ib3c8cd942d46320c2fd34fb454791265ab3b207a8dc76dfa1f7094b3645a887f.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ib3c8cd942d46320c2fd34fb454791265ab3b207a8dc76dfa1f7094b3645a887f.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ia77378d28cf70192f029ce5fe7633f8d0b6708a4c8ee1cee42366173ee23dc4c.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ia77378d28cf70192f029ce5fe7633f8d0b6708a4c8ee1cee42366173ee23dc4c.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i895b3abd81f960bf723a0f694ee3e19c780bc68dfcabf3fce2e7c9026bb1942c.GetMemberGroupsRequestBuilder) {
    return i895b3abd81f960bf723a0f694ee3e19c780bc68dfcabf3fce2e7c9026bb1942c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i55769d9ada2dc3b28e70e204ffcb814c7e3e9d60538577117ecec1d3413f4e0c.GetMemberObjectsRequestBuilder) {
    return i55769d9ada2dc3b28e70e204ffcb814c7e3e9d60538577117ecec1d3413f4e0c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i87013f21fe46569b7386414cf5e93cd03add1fd237420df1e7c0d179b4a845f4.InvalidateAllRefreshTokensRequestBuilder) {
    return i87013f21fe46569b7386414cf5e93cd03add1fd237420df1e7c0d179b4a845f4.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i9ce5877f66d2d1ce704ec30ff4102e6f1f1ccc2678e6aedc8e026d5c1815caba.IsManagedAppUserBlockedRequestBuilder) {
    return i9ce5877f66d2d1ce704ec30ff4102e6f1f1ccc2678e6aedc8e026d5c1815caba.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i1f5c9f54c8e95106b91029db78cda487b3dec35fb5239e2f564c6b9606ec63d1.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i1f5c9f54c8e95106b91029db78cda487b3dec35fb5239e2f564c6b9606ec63d1.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ie7c52783bf847613ebaf73275f705022a0fc262866c43f495f384354ef2f5c32.RemoveAllDevicesFromManagementRequestBuilder) {
    return ie7c52783bf847613ebaf73275f705022a0fc262866c43f495f384354ef2f5c32.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i80fbd8bbbcd3902c4c6d5082c7fa00ef8e367fa1bf5f340d648c9f3693ce9028.ReprocessLicenseAssignmentRequestBuilder) {
    return i80fbd8bbbcd3902c4c6d5082c7fa00ef8e367fa1bf5f340d648c9f3693ce9028.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i5e5177f323c1ffe326c78295fe87257ed907912be0698cac966060b1b783a5d5.RestoreRequestBuilder) {
    return i5e5177f323c1ffe326c78295fe87257ed907912be0698cac966060b1b783a5d5.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ie8e00017f331494e3c3b4b3ea69df62019ab36f641198bb8cb530a216f245e1c.RevokeSignInSessionsRequestBuilder) {
    return ie8e00017f331494e3c3b4b3ea69df62019ab36f641198bb8cb530a216f245e1c.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i5eb63b72fe4ce7f0bef8cbddc22543544e7ca81edf1d2c394c322c476d45bfbb.SendMailRequestBuilder) {
    return i5eb63b72fe4ce7f0bef8cbddc22543544e7ca81edf1d2c394c322c476d45bfbb.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ib072ba211a1fdcba5cf8ea7750f2410f13357828847ece6371159df1045d8ce5.TranslateExchangeIdsRequestBuilder) {
    return ib072ba211a1fdcba5cf8ea7750f2410f13357828847ece6371159df1045d8ce5.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ia9e64eb2592cd9b16347e2de58aac807cb77d1c0ddc01008b92b579f9845c467.UnblockManagedAppsRequestBuilder) {
    return ia9e64eb2592cd9b16347e2de58aac807cb77d1c0ddc01008b92b579f9845c467.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ib4a698ac8bb61bfc7884ad0f895c09a65997b9f50bd85dca996044e47860b769.WipeAndBlockManagedAppsRequestBuilder) {
    return ib4a698ac8bb61bfc7884ad0f895c09a65997b9f50bd85dca996044e47860b769.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i53f1d35eb6c18d449e7731e30f543801df23913d9aba47de3bb4cb45e3cf6322.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i53f1d35eb6c18d449e7731e30f543801df23913d9aba47de3bb4cb45e3cf6322.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i38425d216794182cc2595d32e83b3210dd8a1ed5be3801bdcaa576426ec59862.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i38425d216794182cc2595d32e83b3210dd8a1ed5be3801bdcaa576426ec59862.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i870ee559c28e4572280efa997e81663ff0b4fc26ccfb4a53c5851d16bcebc702.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i870ee559c28e4572280efa997e81663ff0b4fc26ccfb4a53c5851d16bcebc702.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
