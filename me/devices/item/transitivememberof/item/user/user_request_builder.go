package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i039a0f2bebc212f71477ea7ed1a3d5dae97e7f98b5f0c372f0561e56ae049348 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i086ec21accb5a6f27ed4a7f5db328b16879fe952f8436acbfe5d1500d1523809 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    i0c698df8046ea58c1ae570f33bfd91a9c36af447d147893f8a21253fe5badeda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/findrooms"
    i15abd5c2c5875ae5858c4c771356d7c1d67b632ef199bb2a1ab7f15fad3ebb44 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/changepassword"
    i1b6c1b87730c0c1ea71bf9ec1264d8ebfcf3ef4bdfea6ebbf4d6ee1fb212b10d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getloggedonmanageddevices"
    i28a261cb3c2c9d5a7575851786766c63036d410829ea2539e12d567197b18cf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/assignlicense"
    i33d8e4f36675ee8bb165aade27f30e060307feb69823c892fb1467a438dc8a5c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    i37c135caabbbb489d2d5e176a47f8ae345b281f6c15467a8c3a15c45fd108d11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/findmeetingtimes"
    i39e4d9c37b86b57ce79f81762ca5d8844fb87ca984f92da9939b875b97ec58df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/ismanagedappuserblocked"
    i3a5a7ddc1b9ff865f10d5fa387b83f20f5e1100c68f3389a43e6c8fb0e6f19e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmanagedapppolicies"
    i400b632b7a1189fb79edaa26af15bc6e7a43f1c68d0069f6541b769a0af2e3e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmemberobjects"
    i52fe9e317f6b2628969d09d1afe6e21670ab520fddfc117e51967ffdee3faa58 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    i53e348f2a6861e33164186927fa71f1a60e7b34c794032131a9c94054e5dce41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/unblockmanagedapps"
    i58790cb97fa78f76e42b7d29d9235dac6b6792b181f92c80cc68158f6e3d2de5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/exportdeviceandappmanagementdata"
    i5fdc140054168c7c773728d1dad3f7bc67b2901952c1923aa07a30669ff8eb8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/checkmembergroups"
    i635bab25148834b637da8d62b902919184d33307a0130eb5bedb8527b71574a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/translateexchangeids"
    i6f3d352d3abc4ff20880471eab6730641bd53ea4ff20034eff42552f1af23e27 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    i7a8d0ac3a4800158ca80ebf1ee1ffb5e8b7125e3c1bf21fe83da879854242409 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/exportpersonaldata"
    i7de5bbe2ffcaae0e891731819bdf35d3a0e1479d5e1289f7bbe911e373e4d6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/restore"
    i7f89594fac9fd475f162c28e2fd7b7ac229f27780bed14d5824054b3d3eb69b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/findroomlists"
    i83cc467cad6a5a948fceff8c239c2e91f75198a674816b3a7771de73d65ba2af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/reprocesslicenseassignment"
    i8474dd4e6a603e80d09f2d24e35b69d5241bc6e66158cb7eacb7de42060c550d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/revokesigninsessions"
    i8738e18acade0e8779e3864d60964c1259017b28e4886302b699a7495f014ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/findroomswithroomlist"
    iad286d69a9571149e32fdb22e39861299ebc17a255e1532e491834bfa37d2ac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmailtips"
    ic1695858fc61b230523542611bdc63ad0e09233269d2592e2ff79eb54c774d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/checkmemberobjects"
    icba60e3388a9c63b34a17beab57cfcccf35ce9b271ef6207ce9808572cdbbfe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmanageddeviceswithappfailures"
    ice83f83793585b1f78d8dabc1a51dcdb2b82082bd74d8ad625aa010eae5daede "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    id66ae24beddbd0ff8578fa4397e544933b4257b87177eef1457a7d425ec153df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/invalidateallrefreshtokens"
    idcf2e53b48d3e4bbc84a9b2f421a07028629625d5465bdd9611b2f09a1df423f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/wipeandblockmanagedapps"
    ie25e3ac09acfed8953b581b7942b522e2d0c4c45a100e4c537500de0fb07f293 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/removealldevicesfrommanagement"
    ie7b4b5dc14d1211cb02f842c733e43c8224f7d6fcffb7c01f35da761be39bf47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmembergroups"
    iebfb5b5ee114544432fdeb603145d5f75cccb290c586dd10e91abf76b3ecd4db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    iee225a4b0cda07e0bd62ae1f52489a508f360fbe556cc0a5e2249387b03991db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    if09ccbe5c8df51b4bd9ac774fefd4aea667d584528202a23b40f5a4d78f32ee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/sendmail"
    if4394a6f26f056f06be35f8537fad791a4b3323a77a3227e6d0064891bd2f062 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/transitivememberof/item/user/activateserviceplan"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*if4394a6f26f056f06be35f8537fad791a4b3323a77a3227e6d0064891bd2f062.ActivateServicePlanRequestBuilder) {
    return if4394a6f26f056f06be35f8537fad791a4b3323a77a3227e6d0064891bd2f062.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i28a261cb3c2c9d5a7575851786766c63036d410829ea2539e12d567197b18cf3.AssignLicenseRequestBuilder) {
    return i28a261cb3c2c9d5a7575851786766c63036d410829ea2539e12d567197b18cf3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i15abd5c2c5875ae5858c4c771356d7c1d67b632ef199bb2a1ab7f15fad3ebb44.ChangePasswordRequestBuilder) {
    return i15abd5c2c5875ae5858c4c771356d7c1d67b632ef199bb2a1ab7f15fad3ebb44.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*i5fdc140054168c7c773728d1dad3f7bc67b2901952c1923aa07a30669ff8eb8e.CheckMemberGroupsRequestBuilder) {
    return i5fdc140054168c7c773728d1dad3f7bc67b2901952c1923aa07a30669ff8eb8e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*ic1695858fc61b230523542611bdc63ad0e09233269d2592e2ff79eb54c774d04.CheckMemberObjectsRequestBuilder) {
    return ic1695858fc61b230523542611bdc63ad0e09233269d2592e2ff79eb54c774d04.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i58790cb97fa78f76e42b7d29d9235dac6b6792b181f92c80cc68158f6e3d2de5.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i58790cb97fa78f76e42b7d29d9235dac6b6792b181f92c80cc68158f6e3d2de5.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ice83f83793585b1f78d8dabc1a51dcdb2b82082bd74d8ad625aa010eae5daede.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ice83f83793585b1f78d8dabc1a51dcdb2b82082bd74d8ad625aa010eae5daede.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i7a8d0ac3a4800158ca80ebf1ee1ffb5e8b7125e3c1bf21fe83da879854242409.ExportPersonalDataRequestBuilder) {
    return i7a8d0ac3a4800158ca80ebf1ee1ffb5e8b7125e3c1bf21fe83da879854242409.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i37c135caabbbb489d2d5e176a47f8ae345b281f6c15467a8c3a15c45fd108d11.FindMeetingTimesRequestBuilder) {
    return i37c135caabbbb489d2d5e176a47f8ae345b281f6c15467a8c3a15c45fd108d11.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i7f89594fac9fd475f162c28e2fd7b7ac229f27780bed14d5824054b3d3eb69b5.FindRoomListsRequestBuilder) {
    return i7f89594fac9fd475f162c28e2fd7b7ac229f27780bed14d5824054b3d3eb69b5.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i0c698df8046ea58c1ae570f33bfd91a9c36af447d147893f8a21253fe5badeda.FindRoomsRequestBuilder) {
    return i0c698df8046ea58c1ae570f33bfd91a9c36af447d147893f8a21253fe5badeda.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i8738e18acade0e8779e3864d60964c1259017b28e4886302b699a7495f014ddd.FindRoomsWithRoomListRequestBuilder) {
    return i8738e18acade0e8779e3864d60964c1259017b28e4886302b699a7495f014ddd.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i039a0f2bebc212f71477ea7ed1a3d5dae97e7f98b5f0c372f0561e56ae049348.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i039a0f2bebc212f71477ea7ed1a3d5dae97e7f98b5f0c372f0561e56ae049348.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i1b6c1b87730c0c1ea71bf9ec1264d8ebfcf3ef4bdfea6ebbf4d6ee1fb212b10d.GetLoggedOnManagedDevicesRequestBuilder) {
    return i1b6c1b87730c0c1ea71bf9ec1264d8ebfcf3ef4bdfea6ebbf4d6ee1fb212b10d.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*iad286d69a9571149e32fdb22e39861299ebc17a255e1532e491834bfa37d2ac6.GetMailTipsRequestBuilder) {
    return iad286d69a9571149e32fdb22e39861299ebc17a255e1532e491834bfa37d2ac6.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*iee225a4b0cda07e0bd62ae1f52489a508f360fbe556cc0a5e2249387b03991db.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return iee225a4b0cda07e0bd62ae1f52489a508f360fbe556cc0a5e2249387b03991db.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i3a5a7ddc1b9ff865f10d5fa387b83f20f5e1100c68f3389a43e6c8fb0e6f19e2.GetManagedAppPoliciesRequestBuilder) {
    return i3a5a7ddc1b9ff865f10d5fa387b83f20f5e1100c68f3389a43e6c8fb0e6f19e2.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*icba60e3388a9c63b34a17beab57cfcccf35ce9b271ef6207ce9808572cdbbfe5.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return icba60e3388a9c63b34a17beab57cfcccf35ce9b271ef6207ce9808572cdbbfe5.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*i52fe9e317f6b2628969d09d1afe6e21670ab520fddfc117e51967ffdee3faa58.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return i52fe9e317f6b2628969d09d1afe6e21670ab520fddfc117e51967ffdee3faa58.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*ie7b4b5dc14d1211cb02f842c733e43c8224f7d6fcffb7c01f35da761be39bf47.GetMemberGroupsRequestBuilder) {
    return ie7b4b5dc14d1211cb02f842c733e43c8224f7d6fcffb7c01f35da761be39bf47.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i400b632b7a1189fb79edaa26af15bc6e7a43f1c68d0069f6541b769a0af2e3e9.GetMemberObjectsRequestBuilder) {
    return i400b632b7a1189fb79edaa26af15bc6e7a43f1c68d0069f6541b769a0af2e3e9.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*id66ae24beddbd0ff8578fa4397e544933b4257b87177eef1457a7d425ec153df.InvalidateAllRefreshTokensRequestBuilder) {
    return id66ae24beddbd0ff8578fa4397e544933b4257b87177eef1457a7d425ec153df.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i39e4d9c37b86b57ce79f81762ca5d8844fb87ca984f92da9939b875b97ec58df.IsManagedAppUserBlockedRequestBuilder) {
    return i39e4d9c37b86b57ce79f81762ca5d8844fb87ca984f92da9939b875b97ec58df.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i6f3d352d3abc4ff20880471eab6730641bd53ea4ff20034eff42552f1af23e27.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i6f3d352d3abc4ff20880471eab6730641bd53ea4ff20034eff42552f1af23e27.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ie25e3ac09acfed8953b581b7942b522e2d0c4c45a100e4c537500de0fb07f293.RemoveAllDevicesFromManagementRequestBuilder) {
    return ie25e3ac09acfed8953b581b7942b522e2d0c4c45a100e4c537500de0fb07f293.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i83cc467cad6a5a948fceff8c239c2e91f75198a674816b3a7771de73d65ba2af.ReprocessLicenseAssignmentRequestBuilder) {
    return i83cc467cad6a5a948fceff8c239c2e91f75198a674816b3a7771de73d65ba2af.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i7de5bbe2ffcaae0e891731819bdf35d3a0e1479d5e1289f7bbe911e373e4d6bf.RestoreRequestBuilder) {
    return i7de5bbe2ffcaae0e891731819bdf35d3a0e1479d5e1289f7bbe911e373e4d6bf.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i8474dd4e6a603e80d09f2d24e35b69d5241bc6e66158cb7eacb7de42060c550d.RevokeSignInSessionsRequestBuilder) {
    return i8474dd4e6a603e80d09f2d24e35b69d5241bc6e66158cb7eacb7de42060c550d.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*if09ccbe5c8df51b4bd9ac774fefd4aea667d584528202a23b40f5a4d78f32ee2.SendMailRequestBuilder) {
    return if09ccbe5c8df51b4bd9ac774fefd4aea667d584528202a23b40f5a4d78f32ee2.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i635bab25148834b637da8d62b902919184d33307a0130eb5bedb8527b71574a3.TranslateExchangeIdsRequestBuilder) {
    return i635bab25148834b637da8d62b902919184d33307a0130eb5bedb8527b71574a3.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i53e348f2a6861e33164186927fa71f1a60e7b34c794032131a9c94054e5dce41.UnblockManagedAppsRequestBuilder) {
    return i53e348f2a6861e33164186927fa71f1a60e7b34c794032131a9c94054e5dce41.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*idcf2e53b48d3e4bbc84a9b2f421a07028629625d5465bdd9611b2f09a1df423f.WipeAndBlockManagedAppsRequestBuilder) {
    return idcf2e53b48d3e4bbc84a9b2f421a07028629625d5465bdd9611b2f09a1df423f.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i086ec21accb5a6f27ed4a7f5db328b16879fe952f8436acbfe5d1500d1523809.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i086ec21accb5a6f27ed4a7f5db328b16879fe952f8436acbfe5d1500d1523809.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*iebfb5b5ee114544432fdeb603145d5f75cccb290c586dd10e91abf76b3ecd4db.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return iebfb5b5ee114544432fdeb603145d5f75cccb290c586dd10e91abf76b3ecd4db.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*i33d8e4f36675ee8bb165aade27f30e060307feb69823c892fb1467a438dc8a5c.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return i33d8e4f36675ee8bb165aade27f30e060307feb69823c892fb1467a438dc8a5c.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
