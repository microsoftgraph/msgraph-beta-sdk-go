package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0dd10804814bb6e23ddd71db12ca98df6e41800288823c4e0b53503139169358 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/assignlicense"
    i2bcc092da42e1bb694e34cb54c6e346c0a882863682c5382185ef6d757051af1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/restore"
    i2ec56c96a591f04a4ab60aee696e5e726bfc75a9f445c4ab808fc9dffbc611be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findrooms"
    i33c70d6acd0800a15089d8bcf3dcd6603b280f1c2aea548e57f86eef224c92b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmemberobjects"
    i35b02e67f96c23da06a3d979056142d9a931216ae8adc2988034f9af005da71d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    i3acce3bc46d3bca9c60e5be3a7fa20c71486725670c37026d120bab5527d54a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    i3c37afa52c0161a3a988caab9a66197973d94ac78ae4f8bbf88dd5c6464b4ced "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    i4cc28d51341cd985784a3b2c98fd9a2fdd8b328069a987690f4dbb42f3cb90df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    i4ecf30a6324319618e1c89d1f4182158c4e342f88100e38861cf51ac817ed8a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipeandblockmanagedapps"
    i518eb5883b8b92e390baca20aa337ae202ee0b052bf636a598b1f4354c87c8cc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i66074f40859f10d5d897bc923670baf015825274ef3bb0b245f50dbb9eada1e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/changepassword"
    i6cbd5e82e8d60540da36db05b454b1832090f41711f5ace3c54de0eb9250255e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/activateserviceplan"
    i6f756f3b3afa9a80198f976f45215b049aa255ba90e4339ee4fc732fa28089be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanageddeviceswithappfailures"
    i782491a55e93df4f6532da5a8328804b274a930b2c9d9aa885e57bf2b7109876 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/sendmail"
    i79be062ea894db489ba284033a8cc14c2ba3c3f34873a1b2b9b56b0f872839bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/translateexchangeids"
    i81c6498d6d75d343dff458293f2f5efbce7d98fa6d1eb58b9d927349a7962bdb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findroomswithroomlist"
    i8a40282557c5b74d0c10d04365aad23c2470d46e496d605a73c40c55e6dbcee1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/invalidateallrefreshtokens"
    i910eb89946ee6843ec916eaa5eef48a435e0af45675307cefa4839bcae288e4d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedappdiagnosticstatuses"
    i9526b5ff7df7b870c01d250ddf9aba4b6fa70c12bfd52e603d66f60f3d95b8f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findmeetingtimes"
    i9c03b45212f755cbe33560ff090e36e1ca69e094134a09ca38c1b424c1e19695 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmailtips"
    ia701d3ce273a24d5b88a18b2880d2d29a9b8be121c263885da30645fba4afa47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ia7a5c9e76e78c6eb9b853112bc7cdfa86fe935790b1b1e1906264df4e4247ce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    ic8e45b19260104b072392ac16d75b13de5b04f14650704d5b63ad98f15aa152b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/unblockmanagedapps"
    icab9a31e8baff054006e4be487b5f6525e4b9fcf7a3762b5063e6da2b6bb1612 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmanagedapppolicies"
    id5e6c4feb8ce80ab889b4def01090e5c6915eadc437209a0e032d41308bcf462 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getloggedonmanageddevices"
    id7109f2d04767ebf86c2bfb8d1b1d631b0c3a8d6ed2430053eb3880776a52005 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/findroomlists"
    id871ee7b6e36e745d256a2b48cfc2bdb32342e97888bba9fbbd9e5707db3a4c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/getmembergroups"
    idbe72874d8998c93bc6215e06d774bc63bea3b65d3726275e9c6aac412f076e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmembergroups"
    idd8cc3485b1b695d52c156d235e65ae43ef1c3c4c8b9632f2e1185db49e8f199 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/revokesigninsessions"
    ie038d3aa66c99421ee202b50f3ada75d87a9b9e049edd4638918e0a8b9c6b3b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportpersonaldata"
    iead9ea0294582d82a56e624810b65dab920e2e0388f7208ec4d962fa5dc3c788 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/exportdeviceandappmanagementdata"
    ied2c11a5e941a5bc0f4c15c0130d1e7eaaed318a9158d30fe9365172d084dc6d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/ismanagedappuserblocked"
    ief02c6e5cd345088f041db92513011db79fd3b8135a165c4682a9152ba59a8ea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/removealldevicesfrommanagement"
    if707d3b12cd8c8feec11504ad726ddb3a159e82387ded9a5c0d3a0b5e6ae9c5b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/reprocesslicenseassignment"
    ifbd98b9530cf0bf292f852924b2e24235288f4c071d29c30a40d149381cb8246 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user/checkmemberobjects"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i6cbd5e82e8d60540da36db05b454b1832090f41711f5ace3c54de0eb9250255e.ActivateServicePlanRequestBuilder) {
    return i6cbd5e82e8d60540da36db05b454b1832090f41711f5ace3c54de0eb9250255e.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i0dd10804814bb6e23ddd71db12ca98df6e41800288823c4e0b53503139169358.AssignLicenseRequestBuilder) {
    return i0dd10804814bb6e23ddd71db12ca98df6e41800288823c4e0b53503139169358.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i66074f40859f10d5d897bc923670baf015825274ef3bb0b245f50dbb9eada1e4.ChangePasswordRequestBuilder) {
    return i66074f40859f10d5d897bc923670baf015825274ef3bb0b245f50dbb9eada1e4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*idbe72874d8998c93bc6215e06d774bc63bea3b65d3726275e9c6aac412f076e8.CheckMemberGroupsRequestBuilder) {
    return idbe72874d8998c93bc6215e06d774bc63bea3b65d3726275e9c6aac412f076e8.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ifbd98b9530cf0bf292f852924b2e24235288f4c071d29c30a40d149381cb8246.CheckMemberObjectsRequestBuilder) {
    return ifbd98b9530cf0bf292f852924b2e24235288f4c071d29c30a40d149381cb8246.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*iead9ea0294582d82a56e624810b65dab920e2e0388f7208ec4d962fa5dc3c788.ExportDeviceAndAppManagementDataRequestBuilder) {
    return iead9ea0294582d82a56e624810b65dab920e2e0388f7208ec4d962fa5dc3c788.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ia701d3ce273a24d5b88a18b2880d2d29a9b8be121c263885da30645fba4afa47.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ia701d3ce273a24d5b88a18b2880d2d29a9b8be121c263885da30645fba4afa47.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*ie038d3aa66c99421ee202b50f3ada75d87a9b9e049edd4638918e0a8b9c6b3b3.ExportPersonalDataRequestBuilder) {
    return ie038d3aa66c99421ee202b50f3ada75d87a9b9e049edd4638918e0a8b9c6b3b3.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i9526b5ff7df7b870c01d250ddf9aba4b6fa70c12bfd52e603d66f60f3d95b8f4.FindMeetingTimesRequestBuilder) {
    return i9526b5ff7df7b870c01d250ddf9aba4b6fa70c12bfd52e603d66f60f3d95b8f4.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*id7109f2d04767ebf86c2bfb8d1b1d631b0c3a8d6ed2430053eb3880776a52005.FindRoomListsRequestBuilder) {
    return id7109f2d04767ebf86c2bfb8d1b1d631b0c3a8d6ed2430053eb3880776a52005.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i2ec56c96a591f04a4ab60aee696e5e726bfc75a9f445c4ab808fc9dffbc611be.FindRoomsRequestBuilder) {
    return i2ec56c96a591f04a4ab60aee696e5e726bfc75a9f445c4ab808fc9dffbc611be.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i81c6498d6d75d343dff458293f2f5efbce7d98fa6d1eb58b9d927349a7962bdb.FindRoomsWithRoomListRequestBuilder) {
    return i81c6498d6d75d343dff458293f2f5efbce7d98fa6d1eb58b9d927349a7962bdb.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i35b02e67f96c23da06a3d979056142d9a931216ae8adc2988034f9af005da71d.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i35b02e67f96c23da06a3d979056142d9a931216ae8adc2988034f9af005da71d.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*id5e6c4feb8ce80ab889b4def01090e5c6915eadc437209a0e032d41308bcf462.GetLoggedOnManagedDevicesRequestBuilder) {
    return id5e6c4feb8ce80ab889b4def01090e5c6915eadc437209a0e032d41308bcf462.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i9c03b45212f755cbe33560ff090e36e1ca69e094134a09ca38c1b424c1e19695.GetMailTipsRequestBuilder) {
    return i9c03b45212f755cbe33560ff090e36e1ca69e094134a09ca38c1b424c1e19695.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i910eb89946ee6843ec916eaa5eef48a435e0af45675307cefa4839bcae288e4d.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i910eb89946ee6843ec916eaa5eef48a435e0af45675307cefa4839bcae288e4d.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*icab9a31e8baff054006e4be487b5f6525e4b9fcf7a3762b5063e6da2b6bb1612.GetManagedAppPoliciesRequestBuilder) {
    return icab9a31e8baff054006e4be487b5f6525e4b9fcf7a3762b5063e6da2b6bb1612.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i6f756f3b3afa9a80198f976f45215b049aa255ba90e4339ee4fc732fa28089be.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i6f756f3b3afa9a80198f976f45215b049aa255ba90e4339ee4fc732fa28089be.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i3acce3bc46d3bca9c60e5be3a7fa20c71486725670c37026d120bab5527d54a9.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i3acce3bc46d3bca9c60e5be3a7fa20c71486725670c37026d120bab5527d54a9.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*id871ee7b6e36e745d256a2b48cfc2bdb32342e97888bba9fbbd9e5707db3a4c2.GetMemberGroupsRequestBuilder) {
    return id871ee7b6e36e745d256a2b48cfc2bdb32342e97888bba9fbbd9e5707db3a4c2.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i33c70d6acd0800a15089d8bcf3dcd6603b280f1c2aea548e57f86eef224c92b1.GetMemberObjectsRequestBuilder) {
    return i33c70d6acd0800a15089d8bcf3dcd6603b280f1c2aea548e57f86eef224c92b1.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i8a40282557c5b74d0c10d04365aad23c2470d46e496d605a73c40c55e6dbcee1.InvalidateAllRefreshTokensRequestBuilder) {
    return i8a40282557c5b74d0c10d04365aad23c2470d46e496d605a73c40c55e6dbcee1.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ied2c11a5e941a5bc0f4c15c0130d1e7eaaed318a9158d30fe9365172d084dc6d.IsManagedAppUserBlockedRequestBuilder) {
    return ied2c11a5e941a5bc0f4c15c0130d1e7eaaed318a9158d30fe9365172d084dc6d.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ia7a5c9e76e78c6eb9b853112bc7cdfa86fe935790b1b1e1906264df4e4247ce0.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ia7a5c9e76e78c6eb9b853112bc7cdfa86fe935790b1b1e1906264df4e4247ce0.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ief02c6e5cd345088f041db92513011db79fd3b8135a165c4682a9152ba59a8ea.RemoveAllDevicesFromManagementRequestBuilder) {
    return ief02c6e5cd345088f041db92513011db79fd3b8135a165c4682a9152ba59a8ea.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*if707d3b12cd8c8feec11504ad726ddb3a159e82387ded9a5c0d3a0b5e6ae9c5b.ReprocessLicenseAssignmentRequestBuilder) {
    return if707d3b12cd8c8feec11504ad726ddb3a159e82387ded9a5c0d3a0b5e6ae9c5b.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i2bcc092da42e1bb694e34cb54c6e346c0a882863682c5382185ef6d757051af1.RestoreRequestBuilder) {
    return i2bcc092da42e1bb694e34cb54c6e346c0a882863682c5382185ef6d757051af1.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*idd8cc3485b1b695d52c156d235e65ae43ef1c3c4c8b9632f2e1185db49e8f199.RevokeSignInSessionsRequestBuilder) {
    return idd8cc3485b1b695d52c156d235e65ae43ef1c3c4c8b9632f2e1185db49e8f199.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i782491a55e93df4f6532da5a8328804b274a930b2c9d9aa885e57bf2b7109876.SendMailRequestBuilder) {
    return i782491a55e93df4f6532da5a8328804b274a930b2c9d9aa885e57bf2b7109876.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i79be062ea894db489ba284033a8cc14c2ba3c3f34873a1b2b9b56b0f872839bd.TranslateExchangeIdsRequestBuilder) {
    return i79be062ea894db489ba284033a8cc14c2ba3c3f34873a1b2b9b56b0f872839bd.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ic8e45b19260104b072392ac16d75b13de5b04f14650704d5b63ad98f15aa152b.UnblockManagedAppsRequestBuilder) {
    return ic8e45b19260104b072392ac16d75b13de5b04f14650704d5b63ad98f15aa152b.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i4ecf30a6324319618e1c89d1f4182158c4e342f88100e38861cf51ac817ed8a2.WipeAndBlockManagedAppsRequestBuilder) {
    return i4ecf30a6324319618e1c89d1f4182158c4e342f88100e38861cf51ac817ed8a2.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i4cc28d51341cd985784a3b2c98fd9a2fdd8b328069a987690f4dbb42f3cb90df.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i4cc28d51341cd985784a3b2c98fd9a2fdd8b328069a987690f4dbb42f3cb90df.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i518eb5883b8b92e390baca20aa337ae202ee0b052bf636a598b1f4354c87c8cc.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i518eb5883b8b92e390baca20aa337ae202ee0b052bf636a598b1f4354c87c8cc.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i3c37afa52c0161a3a988caab9a66197973d94ac78ae4f8bbf88dd5c6464b4ced.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i3c37afa52c0161a3a988caab9a66197973d94ac78ae4f8bbf88dd5c6464b4ced.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
