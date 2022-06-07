package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00ab3a8d7c4b9908756f88aed0f81770f541c09ab42fbfec97b55c3a75bb1ea3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmanageddeviceswithappfailures"
    i096e80204a624f6a101d220c26580107edb32eb339603f766e2165fa8bab3a54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/removealldevicesfrommanagement"
    i23fa842e27573c95f0f79a4eda5329c8bfc3f8573874b83381bbfa65eacd365b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/activateserviceplan"
    i2d5946f81310ef0055e4b25c720301ec2eaa96cb62306f40e64bc19c7eed2fd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i2db5d40aefa75a69c7227872924f37d10a7ad02b5310edb4791a2b4f158ba1ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmanagedapppolicies"
    i30706567b77ee6728efd4605be4d681697b9ca468593d7b692d34a98587e98c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmailtips"
    i327b93bb69e1cbbc26c04ccb408fb6d921ac3b1a855750bc5fa7e031517347bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/sendmail"
    i3d01246e6c39a267832983b91e9ae5beff7e1411e5f5424e1fa082765330496c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i3fe3c85c8b4c94715548485daddd5df74f4360a67567843c35c16d86ba38772c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/findroomswithroomlist"
    i45f0b3d43920eb7ea90eae5e0c83c1bf58aa484f355e741e500db32fffd450dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i4861fff6bde495d5c074afebbf72098f5bf54b17ac84de8b19db622696d9aaad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/assignlicense"
    i52af2147a60c10a98b4bf6cc0a04231693a0312ad4df77b062445a80a2b5aca5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/revokesigninsessions"
    i588a6cbe3624bf0e18ae8e0175df607330e63e7be2b7ab29369e109abebc4f5c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/exportdeviceandappmanagementdata"
    i59dd461f6765c956be823a85605f4761de614ed8fb15dd7f8d3e6c751ca41de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/checkmembergroups"
    i5efb7b97b3514848b21a941df9666dfba986716eabe2ada38ec7eb154ca6c54e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/unblockmanagedapps"
    i62bb60a0a298c34d1bf7a03de46eccba83eb914853383ec4cdb54d8bc6908c68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmemberobjects"
    i6deb51ee31aa92f6fdd306553f7f003547ab45fb2c02636ad72b06417b4ce20d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/findmeetingtimes"
    i744c64f811eb4d57f8f9437320b6ff3887868704a73d4a5034077fbdb4b7f503 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    i7f4962c75618e230fad6d0bb006ef3b1451e836e6b548cff96623274d7baf20f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i9f6759a7069306441745fe7ed2fe387319b01049a68cd0ddda2ed3b631b0825f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/findrooms"
    ia0a7b024e014796743258446c38fb80be469e4a8ffed3cd8b2fcc71a3cc8adb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/restore"
    ia86b4c5269e55b61c93f8a8f1c89b0dd27225283d9015c8a9200def53e47587e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmanagedappdiagnosticstatuses"
    ib4fa0ae3a353e47a3ca66a1b96ea03b7ba9f872e944e838ee9096a95c77952c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/exportpersonaldata"
    ib675e9fa89b52663295f21f3b3c69e83c94d47f4494c746b2b7d2959e5e6c5b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/translateexchangeids"
    ic1531357dea6a8e9beed4b860dc4d3482e9204e9255aaf2ec63c016ac8936c8a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/findroomlists"
    ic3939eb0457f452703e66af1ba4df8f3de7cb60442cb4eb65176e066fc413564 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getloggedonmanageddevices"
    ic4ba95f7845713525c440706539e33e62bee57cdccc11a61308ef683f80e11f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/changepassword"
    icf246c859cc6feeadbe947feed26380a0932b5ea7f50f5ca65e2db03cef4fd79 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/wipeandblockmanagedapps"
    id5d0c017d20b8840d04f71b947e9552132b0164abc60a0cae22184ba38130244 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/checkmemberobjects"
    id7f09b8b7d8fe1af2461546e78327d75ce5d23a05eafdf03f444fe484b20da02 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/reprocesslicenseassignment"
    ie0f734552ee9eab0f4bd9b6ddef460a1b1b36ea2a467051113855dd27977e14f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ie863ef18606e3227f6479353afd467f9c7f8b52a70f50b6fc3e656111802fe74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/ismanagedappuserblocked"
    ie955f0bfe7e92c84562e35d2145d2cbb8e55a7cce4a44f4f95c73f8e46255785 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/invalidateallrefreshtokens"
    if41d9bde7367ee75e6e1b740360459f16314d14079e0995ce08903bae199f008 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/wipemanagedappregistrationbydevicetag"
    ifaa0cda89693714142cf313b7216d2749be692e7dba9676a6740250a7fa4098e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/item/user/getmembergroups"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i23fa842e27573c95f0f79a4eda5329c8bfc3f8573874b83381bbfa65eacd365b.ActivateServicePlanRequestBuilder) {
    return i23fa842e27573c95f0f79a4eda5329c8bfc3f8573874b83381bbfa65eacd365b.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i4861fff6bde495d5c074afebbf72098f5bf54b17ac84de8b19db622696d9aaad.AssignLicenseRequestBuilder) {
    return i4861fff6bde495d5c074afebbf72098f5bf54b17ac84de8b19db622696d9aaad.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ic4ba95f7845713525c440706539e33e62bee57cdccc11a61308ef683f80e11f4.ChangePasswordRequestBuilder) {
    return ic4ba95f7845713525c440706539e33e62bee57cdccc11a61308ef683f80e11f4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i59dd461f6765c956be823a85605f4761de614ed8fb15dd7f8d3e6c751ca41de6.CheckMemberGroupsRequestBuilder) {
    return i59dd461f6765c956be823a85605f4761de614ed8fb15dd7f8d3e6c751ca41de6.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*id5d0c017d20b8840d04f71b947e9552132b0164abc60a0cae22184ba38130244.CheckMemberObjectsRequestBuilder) {
    return id5d0c017d20b8840d04f71b947e9552132b0164abc60a0cae22184ba38130244.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i588a6cbe3624bf0e18ae8e0175df607330e63e7be2b7ab29369e109abebc4f5c.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i588a6cbe3624bf0e18ae8e0175df607330e63e7be2b7ab29369e109abebc4f5c.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i3d01246e6c39a267832983b91e9ae5beff7e1411e5f5424e1fa082765330496c.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i3d01246e6c39a267832983b91e9ae5beff7e1411e5f5424e1fa082765330496c.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ib4fa0ae3a353e47a3ca66a1b96ea03b7ba9f872e944e838ee9096a95c77952c3.ExportPersonalDataRequestBuilder) {
    return ib4fa0ae3a353e47a3ca66a1b96ea03b7ba9f872e944e838ee9096a95c77952c3.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i6deb51ee31aa92f6fdd306553f7f003547ab45fb2c02636ad72b06417b4ce20d.FindMeetingTimesRequestBuilder) {
    return i6deb51ee31aa92f6fdd306553f7f003547ab45fb2c02636ad72b06417b4ce20d.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ic1531357dea6a8e9beed4b860dc4d3482e9204e9255aaf2ec63c016ac8936c8a.FindRoomListsRequestBuilder) {
    return ic1531357dea6a8e9beed4b860dc4d3482e9204e9255aaf2ec63c016ac8936c8a.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i9f6759a7069306441745fe7ed2fe387319b01049a68cd0ddda2ed3b631b0825f.FindRoomsRequestBuilder) {
    return i9f6759a7069306441745fe7ed2fe387319b01049a68cd0ddda2ed3b631b0825f.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i3fe3c85c8b4c94715548485daddd5df74f4360a67567843c35c16d86ba38772c.FindRoomsWithRoomListRequestBuilder) {
    return i3fe3c85c8b4c94715548485daddd5df74f4360a67567843c35c16d86ba38772c.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i45f0b3d43920eb7ea90eae5e0c83c1bf58aa484f355e741e500db32fffd450dc.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i45f0b3d43920eb7ea90eae5e0c83c1bf58aa484f355e741e500db32fffd450dc.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ic3939eb0457f452703e66af1ba4df8f3de7cb60442cb4eb65176e066fc413564.GetLoggedOnManagedDevicesRequestBuilder) {
    return ic3939eb0457f452703e66af1ba4df8f3de7cb60442cb4eb65176e066fc413564.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i30706567b77ee6728efd4605be4d681697b9ca468593d7b692d34a98587e98c3.GetMailTipsRequestBuilder) {
    return i30706567b77ee6728efd4605be4d681697b9ca468593d7b692d34a98587e98c3.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*ia86b4c5269e55b61c93f8a8f1c89b0dd27225283d9015c8a9200def53e47587e.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return ia86b4c5269e55b61c93f8a8f1c89b0dd27225283d9015c8a9200def53e47587e.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i2db5d40aefa75a69c7227872924f37d10a7ad02b5310edb4791a2b4f158ba1ac.GetManagedAppPoliciesRequestBuilder) {
    return i2db5d40aefa75a69c7227872924f37d10a7ad02b5310edb4791a2b4f158ba1ac.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i00ab3a8d7c4b9908756f88aed0f81770f541c09ab42fbfec97b55c3a75bb1ea3.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i00ab3a8d7c4b9908756f88aed0f81770f541c09ab42fbfec97b55c3a75bb1ea3.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ie0f734552ee9eab0f4bd9b6ddef460a1b1b36ea2a467051113855dd27977e14f.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ie0f734552ee9eab0f4bd9b6ddef460a1b1b36ea2a467051113855dd27977e14f.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ifaa0cda89693714142cf313b7216d2749be692e7dba9676a6740250a7fa4098e.GetMemberGroupsRequestBuilder) {
    return ifaa0cda89693714142cf313b7216d2749be692e7dba9676a6740250a7fa4098e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i62bb60a0a298c34d1bf7a03de46eccba83eb914853383ec4cdb54d8bc6908c68.GetMemberObjectsRequestBuilder) {
    return i62bb60a0a298c34d1bf7a03de46eccba83eb914853383ec4cdb54d8bc6908c68.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ie955f0bfe7e92c84562e35d2145d2cbb8e55a7cce4a44f4f95c73f8e46255785.InvalidateAllRefreshTokensRequestBuilder) {
    return ie955f0bfe7e92c84562e35d2145d2cbb8e55a7cce4a44f4f95c73f8e46255785.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ie863ef18606e3227f6479353afd467f9c7f8b52a70f50b6fc3e656111802fe74.IsManagedAppUserBlockedRequestBuilder) {
    return ie863ef18606e3227f6479353afd467f9c7f8b52a70f50b6fc3e656111802fe74.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i7f4962c75618e230fad6d0bb006ef3b1451e836e6b548cff96623274d7baf20f.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i7f4962c75618e230fad6d0bb006ef3b1451e836e6b548cff96623274d7baf20f.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i096e80204a624f6a101d220c26580107edb32eb339603f766e2165fa8bab3a54.RemoveAllDevicesFromManagementRequestBuilder) {
    return i096e80204a624f6a101d220c26580107edb32eb339603f766e2165fa8bab3a54.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*id7f09b8b7d8fe1af2461546e78327d75ce5d23a05eafdf03f444fe484b20da02.ReprocessLicenseAssignmentRequestBuilder) {
    return id7f09b8b7d8fe1af2461546e78327d75ce5d23a05eafdf03f444fe484b20da02.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*ia0a7b024e014796743258446c38fb80be469e4a8ffed3cd8b2fcc71a3cc8adb0.RestoreRequestBuilder) {
    return ia0a7b024e014796743258446c38fb80be469e4a8ffed3cd8b2fcc71a3cc8adb0.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i52af2147a60c10a98b4bf6cc0a04231693a0312ad4df77b062445a80a2b5aca5.RevokeSignInSessionsRequestBuilder) {
    return i52af2147a60c10a98b4bf6cc0a04231693a0312ad4df77b062445a80a2b5aca5.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i327b93bb69e1cbbc26c04ccb408fb6d921ac3b1a855750bc5fa7e031517347bd.SendMailRequestBuilder) {
    return i327b93bb69e1cbbc26c04ccb408fb6d921ac3b1a855750bc5fa7e031517347bd.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*ib675e9fa89b52663295f21f3b3c69e83c94d47f4494c746b2b7d2959e5e6c5b5.TranslateExchangeIdsRequestBuilder) {
    return ib675e9fa89b52663295f21f3b3c69e83c94d47f4494c746b2b7d2959e5e6c5b5.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i5efb7b97b3514848b21a941df9666dfba986716eabe2ada38ec7eb154ca6c54e.UnblockManagedAppsRequestBuilder) {
    return i5efb7b97b3514848b21a941df9666dfba986716eabe2ada38ec7eb154ca6c54e.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*icf246c859cc6feeadbe947feed26380a0932b5ea7f50f5ca65e2db03cef4fd79.WipeAndBlockManagedAppsRequestBuilder) {
    return icf246c859cc6feeadbe947feed26380a0932b5ea7f50f5ca65e2db03cef4fd79.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*if41d9bde7367ee75e6e1b740360459f16314d14079e0995ce08903bae199f008.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return if41d9bde7367ee75e6e1b740360459f16314d14079e0995ce08903bae199f008.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i2d5946f81310ef0055e4b25c720301ec2eaa96cb62306f40e64bc19c7eed2fd5.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i2d5946f81310ef0055e4b25c720301ec2eaa96cb62306f40e64bc19c7eed2fd5.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i744c64f811eb4d57f8f9437320b6ff3887868704a73d4a5034077fbdb4b7f503.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i744c64f811eb4d57f8f9437320b6ff3887868704a73d4a5034077fbdb4b7f503.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
