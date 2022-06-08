package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b8465d2922c734f211f40cb5c0781e486a89b7f1fa143622411905ed69b2cb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmembergroups"
    i1595e4d61c440bf4f64450db6c83cd1ebc18d90802208d995eebf1f810696ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/removealldevicesfrommanagement"
    i1f78bd68d9abf02706ff1f4370ce92877eeac4876fe588af9b1f19f117a6b38a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/ismanagedappuserblocked"
    i336a0c4381a170b8a802ee82a0f3406dfe929b2eb28c19f531a583b970438627 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmemberobjects"
    i384f692def41e5182e051c499e9b2125bfadfc49f5341c3050913c7c1afcc498 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findrooms"
    i3e41084d33d4be50e65fdeb02851cb33fe2759bdfda8877434d4cde8871b648e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reprocesslicenseassignment"
    i402e71cd240cac59258f96e9def743c3040c91a2f162801e81c3828b995dd56b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getloggedonmanageddevices"
    i470d128e2f6f03e9e1a235f8ebf764f09a5256139a84b2f51718069250f29dc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findroomlists"
    i4e3b0bdad408f6939ceddaa05bf361810a3a39eb6750835ae41465d4e0c8b4b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i4f37f62363ba427ca8fed67e79e08ecb1b2e797e8bcc400478e6c9c4a0447bb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/translateexchangeids"
    i554f0c4a270b2507768b9e0d5d3afe59c5e16efec805ee86f4bb413ac8382314 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/revokesigninsessions"
    i5567bd164403df0c74610f716a88d3973020dcb2f644080d80e0ce70affff7ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/invalidateallrefreshtokens"
    i565f79f1d5b7801d0f5f2753e44e29db33e5ff9294506a2295a95791cd2ccb1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findmeetingtimes"
    i5ec7ab1e4788242680a71ccc82d3be6dbaf6e2a39681a1e9425d914b36d2dea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i61a099640dadd1387c8fc0a012a3708ecc34df1a059eb225d081d9dd4e19b5af "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/restore"
    i67bf990a82a085a8b95a85547f695f71cb61dc223dd89321ca173cf5964f7ad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportdeviceandappmanagementdata"
    i68e58d492ce429f43999ae98319d480ce0ae5116a814839998cd260763a46cce "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/assignlicense"
    i69f79690fc6dc0d2e593f8f16160103512bf10299836d54fc2e142f23827fd26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmemberobjects"
    i72e8aa8fc9ae6d004e2615cd04c5251250b925b74eb7a42100a56596d05af00b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedapppolicies"
    i7ff0e9fed1c8f5854404f4566254e1b65fd452416c6722752b814694cbb2b626 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithappfailures"
    i84cd4a42e624a96d0d5e7e6a57820c7886391529eb773157ef7a9e2cf59ec4fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipeandblockmanagedapps"
    i85126e8fd5bc0cb2d04ae97bbfbecaf2d3a49ef8179575ae26880bb7ddc5e1c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmailtips"
    i9366b301c36d017fc802bea9557ca5f89e09e49897a6991ff106abb3f0d42ac4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i971ffaa03f8565df93af409b8b6cb97b818c7daf6ddfea2af838d81643bcc526 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/checkmembergroups"
    i9b05b8302a7e1b45eccd7fffbf759128c597710fe39be5171c2e4a9d55285297 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/activateserviceplan"
    ia12b95a67802529cd6c31d01b1091809d38c32e16d1f1104890e43763e42051b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationbydevicetag"
    ibbefb59740e2ec3433697946cc5a7194664a548b1f8ca19636e2d4cb31b98eb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/changepassword"
    id2efc94bd8ab6d28d89c648eac6ff7c307d5c43396a54122ca90950c6f29f341 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanagedappdiagnosticstatuses"
    id7dbaebedc15470a19dcd575e2366f04231eeb613ad2058f2764bb1b74ded287 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/exportpersonaldata"
    idbd03160806a995b6db754b8f6090c51c7b95adc958125efc08dd0f4217ca7c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ide667ffe93f0f84b82a4c884d8cf5edf87d42153e527d2b53c6ef9ca9c610336 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbydevicetag"
    ie3d5b404ea08386b45466b3a07186e7beb9ee032e8ad0f49985e9b2d7c7be7cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    ie78cffdb7ab69906caeff25c134cd2a9983f69977c17f79e1fc24d14add584ea "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/unblockmanagedapps"
    if1c8a237d7c79b932d2d94fd28851b1323a913dd66dd0f73a98af31385d35f72 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/findroomswithroomlist"
    if3b4b49d9e0afb9856388ddcb0830636820a50782c33a8fade4c1ad1bfac0d27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/microsoftauthenticatormethods/item/device/memberof/item/user/sendmail"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i9b05b8302a7e1b45eccd7fffbf759128c597710fe39be5171c2e4a9d55285297.ActivateServicePlanRequestBuilder) {
    return i9b05b8302a7e1b45eccd7fffbf759128c597710fe39be5171c2e4a9d55285297.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i68e58d492ce429f43999ae98319d480ce0ae5116a814839998cd260763a46cce.AssignLicenseRequestBuilder) {
    return i68e58d492ce429f43999ae98319d480ce0ae5116a814839998cd260763a46cce.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ibbefb59740e2ec3433697946cc5a7194664a548b1f8ca19636e2d4cb31b98eb0.ChangePasswordRequestBuilder) {
    return ibbefb59740e2ec3433697946cc5a7194664a548b1f8ca19636e2d4cb31b98eb0.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i971ffaa03f8565df93af409b8b6cb97b818c7daf6ddfea2af838d81643bcc526.CheckMemberGroupsRequestBuilder) {
    return i971ffaa03f8565df93af409b8b6cb97b818c7daf6ddfea2af838d81643bcc526.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i336a0c4381a170b8a802ee82a0f3406dfe929b2eb28c19f531a583b970438627.CheckMemberObjectsRequestBuilder) {
    return i336a0c4381a170b8a802ee82a0f3406dfe929b2eb28c19f531a583b970438627.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i67bf990a82a085a8b95a85547f695f71cb61dc223dd89321ca173cf5964f7ad7.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i67bf990a82a085a8b95a85547f695f71cb61dc223dd89321ca173cf5964f7ad7.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i4e3b0bdad408f6939ceddaa05bf361810a3a39eb6750835ae41465d4e0c8b4b2.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i4e3b0bdad408f6939ceddaa05bf361810a3a39eb6750835ae41465d4e0c8b4b2.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*id7dbaebedc15470a19dcd575e2366f04231eeb613ad2058f2764bb1b74ded287.ExportPersonalDataRequestBuilder) {
    return id7dbaebedc15470a19dcd575e2366f04231eeb613ad2058f2764bb1b74ded287.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i565f79f1d5b7801d0f5f2753e44e29db33e5ff9294506a2295a95791cd2ccb1c.FindMeetingTimesRequestBuilder) {
    return i565f79f1d5b7801d0f5f2753e44e29db33e5ff9294506a2295a95791cd2ccb1c.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i470d128e2f6f03e9e1a235f8ebf764f09a5256139a84b2f51718069250f29dc7.FindRoomListsRequestBuilder) {
    return i470d128e2f6f03e9e1a235f8ebf764f09a5256139a84b2f51718069250f29dc7.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i384f692def41e5182e051c499e9b2125bfadfc49f5341c3050913c7c1afcc498.FindRoomsRequestBuilder) {
    return i384f692def41e5182e051c499e9b2125bfadfc49f5341c3050913c7c1afcc498.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*if1c8a237d7c79b932d2d94fd28851b1323a913dd66dd0f73a98af31385d35f72.FindRoomsWithRoomListRequestBuilder) {
    return if1c8a237d7c79b932d2d94fd28851b1323a913dd66dd0f73a98af31385d35f72.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i9366b301c36d017fc802bea9557ca5f89e09e49897a6991ff106abb3f0d42ac4.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i9366b301c36d017fc802bea9557ca5f89e09e49897a6991ff106abb3f0d42ac4.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i402e71cd240cac59258f96e9def743c3040c91a2f162801e81c3828b995dd56b.GetLoggedOnManagedDevicesRequestBuilder) {
    return i402e71cd240cac59258f96e9def743c3040c91a2f162801e81c3828b995dd56b.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i85126e8fd5bc0cb2d04ae97bbfbecaf2d3a49ef8179575ae26880bb7ddc5e1c8.GetMailTipsRequestBuilder) {
    return i85126e8fd5bc0cb2d04ae97bbfbecaf2d3a49ef8179575ae26880bb7ddc5e1c8.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id2efc94bd8ab6d28d89c648eac6ff7c307d5c43396a54122ca90950c6f29f341.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id2efc94bd8ab6d28d89c648eac6ff7c307d5c43396a54122ca90950c6f29f341.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i72e8aa8fc9ae6d004e2615cd04c5251250b925b74eb7a42100a56596d05af00b.GetManagedAppPoliciesRequestBuilder) {
    return i72e8aa8fc9ae6d004e2615cd04c5251250b925b74eb7a42100a56596d05af00b.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i7ff0e9fed1c8f5854404f4566254e1b65fd452416c6722752b814694cbb2b626.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i7ff0e9fed1c8f5854404f4566254e1b65fd452416c6722752b814694cbb2b626.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*idbd03160806a995b6db754b8f6090c51c7b95adc958125efc08dd0f4217ca7c1.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return idbd03160806a995b6db754b8f6090c51c7b95adc958125efc08dd0f4217ca7c1.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i0b8465d2922c734f211f40cb5c0781e486a89b7f1fa143622411905ed69b2cb6.GetMemberGroupsRequestBuilder) {
    return i0b8465d2922c734f211f40cb5c0781e486a89b7f1fa143622411905ed69b2cb6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i69f79690fc6dc0d2e593f8f16160103512bf10299836d54fc2e142f23827fd26.GetMemberObjectsRequestBuilder) {
    return i69f79690fc6dc0d2e593f8f16160103512bf10299836d54fc2e142f23827fd26.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i5567bd164403df0c74610f716a88d3973020dcb2f644080d80e0ce70affff7ff.InvalidateAllRefreshTokensRequestBuilder) {
    return i5567bd164403df0c74610f716a88d3973020dcb2f644080d80e0ce70affff7ff.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1f78bd68d9abf02706ff1f4370ce92877eeac4876fe588af9b1f19f117a6b38a.IsManagedAppUserBlockedRequestBuilder) {
    return i1f78bd68d9abf02706ff1f4370ce92877eeac4876fe588af9b1f19f117a6b38a.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i5ec7ab1e4788242680a71ccc82d3be6dbaf6e2a39681a1e9425d914b36d2dea5.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i5ec7ab1e4788242680a71ccc82d3be6dbaf6e2a39681a1e9425d914b36d2dea5.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i1595e4d61c440bf4f64450db6c83cd1ebc18d90802208d995eebf1f810696ee8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i1595e4d61c440bf4f64450db6c83cd1ebc18d90802208d995eebf1f810696ee8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i3e41084d33d4be50e65fdeb02851cb33fe2759bdfda8877434d4cde8871b648e.ReprocessLicenseAssignmentRequestBuilder) {
    return i3e41084d33d4be50e65fdeb02851cb33fe2759bdfda8877434d4cde8871b648e.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i61a099640dadd1387c8fc0a012a3708ecc34df1a059eb225d081d9dd4e19b5af.RestoreRequestBuilder) {
    return i61a099640dadd1387c8fc0a012a3708ecc34df1a059eb225d081d9dd4e19b5af.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i554f0c4a270b2507768b9e0d5d3afe59c5e16efec805ee86f4bb413ac8382314.RevokeSignInSessionsRequestBuilder) {
    return i554f0c4a270b2507768b9e0d5d3afe59c5e16efec805ee86f4bb413ac8382314.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*if3b4b49d9e0afb9856388ddcb0830636820a50782c33a8fade4c1ad1bfac0d27.SendMailRequestBuilder) {
    return if3b4b49d9e0afb9856388ddcb0830636820a50782c33a8fade4c1ad1bfac0d27.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i4f37f62363ba427ca8fed67e79e08ecb1b2e797e8bcc400478e6c9c4a0447bb1.TranslateExchangeIdsRequestBuilder) {
    return i4f37f62363ba427ca8fed67e79e08ecb1b2e797e8bcc400478e6c9c4a0447bb1.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*ie78cffdb7ab69906caeff25c134cd2a9983f69977c17f79e1fc24d14add584ea.UnblockManagedAppsRequestBuilder) {
    return ie78cffdb7ab69906caeff25c134cd2a9983f69977c17f79e1fc24d14add584ea.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i84cd4a42e624a96d0d5e7e6a57820c7886391529eb773157ef7a9e2cf59ec4fb.WipeAndBlockManagedAppsRequestBuilder) {
    return i84cd4a42e624a96d0d5e7e6a57820c7886391529eb773157ef7a9e2cf59ec4fb.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ia12b95a67802529cd6c31d01b1091809d38c32e16d1f1104890e43763e42051b.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ia12b95a67802529cd6c31d01b1091809d38c32e16d1f1104890e43763e42051b.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ie3d5b404ea08386b45466b3a07186e7beb9ee032e8ad0f49985e9b2d7c7be7cf.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ie3d5b404ea08386b45466b3a07186e7beb9ee032e8ad0f49985e9b2d7c7be7cf.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ide667ffe93f0f84b82a4c884d8cf5edf87d42153e527d2b53c6ef9ca9c610336.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ide667ffe93f0f84b82a4c884d8cf5edf87d42153e527d2b53c6ef9ca9c610336.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
