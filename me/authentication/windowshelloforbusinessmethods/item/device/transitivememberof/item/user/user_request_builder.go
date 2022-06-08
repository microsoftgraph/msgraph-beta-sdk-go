package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0580dd95c982b356e4fb1a5f9d8e9b45a1144aae2a1b04964c8a4d178256bf41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/restore"
    i0b03e66923482b6bafce850b9faffe1ad626a7d337b96e54356341c4cfac6623 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/removealldevicesfrommanagement"
    i0d1cea47b5a00a04ff83adca55cec4670ac81d1ded631fe5da7f6b1523e5cd76 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/sendmail"
    i202a1b749f41f1fea470fc234005fe53b0ce43910e9fe65b70f52a2612e73250 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmailtips"
    i2567c95b0764672f56e44582b22fed51186a694f6b2c83275b804b16b38da24b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i2926d108e8a370c469e897c5fd637bdbf4e59c41b578c927a748934065a7c30e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/geteffectivedeviceenrollmentconfigurations"
    i36e3055616af25088bcf8ba5c72416a6860ee6d6807fb3676590c225f8be2bd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/activateserviceplan"
    i421a0a949e5e3697a51987047d6d893ff0deee857ac39ae51ff1fdb24ad959b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedapppolicies"
    i54eeb9f2609c157ed5f58d71e74c146ed06117900495736137cfe1ff983806b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmembergroups"
    i5ffe6993f05c4715a4d256de9bd0274640a46156296d0bd853c448f7f009bdf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanagedappdiagnosticstatuses"
    i623a4c39b5f86bc611e4c7c79485e84315e8f204addbf279d18216e6fa4a0748 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/translateexchangeids"
    i6302f791def46e7bc84de5cdb9aa89991f9b11a917396c1b13ef1f904e2077a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportpersonaldata"
    i7669ec7815422fab8d6ff39effe64662f0e24d1411d1eb409e01105e6767e630 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getloggedonmanageddevices"
    i78e74354ec09acd0b98b4b4600f804ec7887ce3b3341a45657c366d6956c2523 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipeandblockmanagedapps"
    i809d3af15a46862830bb0a7bbe65141538eb61b2d47bc146069a1b7bad14eedb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/revokesigninsessions"
    i8af688f61dc6212ca8ed456af66f481f6b1d42099b8f612c45afb698082f5776 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findroomswithroomlist"
    i931235fc72db29a8627f2bf2928ee9d102bda6e555552570c18d50d51e82b02e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findrooms"
    i98699fe28f31f6aad757d2704a8b4d4a616f976d16a53129b6dc141e7567e81b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/unblockmanagedapps"
    i9c1812aceca306b433b89022ad806c2d55c49d4d1078bec4fea6e2d425a145b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/assignlicense"
    ia08e746479382b3bc22cccc08ed34698744f70a527af0624f0d5742df55abac3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/ismanagedappuserblocked"
    iae5f51a30530619cb6ca1bf4cc0e83792d392de9017c11af442a7babd5675f90 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmemberobjects"
    ib382f2ce2d135054992a6f83d600cf7b7db0f5612acc12559d9791a5d46e0b56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbydevicetag"
    ib4a63511d54de38d4611349f200716dd80d232b496bea64c94ae3c4f452a78ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findmeetingtimes"
    ib79e33d64ee727d931363a7557bdcbad45675b9a52090505ff8c96f740174fd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/findroomlists"
    ibd538f065b0a60a4a14454443524e48284b038c0035f81a5449c7b63c6591375 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanageddeviceswithfailedorpendingapps"
    ibe9d16b18b45790efb8dd821f6f01fbf965060fab56914a5242336817600d4ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/invalidateallrefreshtokens"
    id25033bd39711c598630f8d0d05b5fc928ff1883f7bb54b75112313fc15e12a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/getmanageddeviceswithappfailures"
    ida81d9030db55ff7aafc84a5c314148c06e3c815b85338b274542c202ff12922 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    idbefd5c8a0b6805f0518f7cee1954bdcfdd9df50c94ccd11e1cc63806a4ec336 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmembergroups"
    ie6479be6c19df4aec957773f6306aa32f894ebf5663b61db7e6d734044438e91 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/wipemanagedappregistrationbydevicetag"
    ie91934d6c0df672c3a4efe01bf837a147dd288c27ea0cc81cd8d0179d4b01eb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reminderviewwithstartdatetimewithenddatetime"
    if09453085e738876e0f6bf5bf71edca2b54748353988cf043978affb1a2a109a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/checkmemberobjects"
    if3c4838af67ca41e08f6a15ff364545b3218d0733ae603b526c689b492577f10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/reprocesslicenseassignment"
    if5bd027d91ee4e4c70ea446ee72329590c872cc0b68b98a1ef05110701fe2c93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/exportdeviceandappmanagementdata"
    ifa6b1ec49d487c39889edd689e9e472b7fd42d2cc9a665f21ea63e58625c9497 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/transitivememberof/item/user/changepassword"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i36e3055616af25088bcf8ba5c72416a6860ee6d6807fb3676590c225f8be2bd9.ActivateServicePlanRequestBuilder) {
    return i36e3055616af25088bcf8ba5c72416a6860ee6d6807fb3676590c225f8be2bd9.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i9c1812aceca306b433b89022ad806c2d55c49d4d1078bec4fea6e2d425a145b3.AssignLicenseRequestBuilder) {
    return i9c1812aceca306b433b89022ad806c2d55c49d4d1078bec4fea6e2d425a145b3.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*ifa6b1ec49d487c39889edd689e9e472b7fd42d2cc9a665f21ea63e58625c9497.ChangePasswordRequestBuilder) {
    return ifa6b1ec49d487c39889edd689e9e472b7fd42d2cc9a665f21ea63e58625c9497.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*idbefd5c8a0b6805f0518f7cee1954bdcfdd9df50c94ccd11e1cc63806a4ec336.CheckMemberGroupsRequestBuilder) {
    return idbefd5c8a0b6805f0518f7cee1954bdcfdd9df50c94ccd11e1cc63806a4ec336.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*if09453085e738876e0f6bf5bf71edca2b54748353988cf043978affb1a2a109a.CheckMemberObjectsRequestBuilder) {
    return if09453085e738876e0f6bf5bf71edca2b54748353988cf043978affb1a2a109a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*if5bd027d91ee4e4c70ea446ee72329590c872cc0b68b98a1ef05110701fe2c93.ExportDeviceAndAppManagementDataRequestBuilder) {
    return if5bd027d91ee4e4c70ea446ee72329590c872cc0b68b98a1ef05110701fe2c93.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i2567c95b0764672f56e44582b22fed51186a694f6b2c83275b804b16b38da24b.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i2567c95b0764672f56e44582b22fed51186a694f6b2c83275b804b16b38da24b.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i6302f791def46e7bc84de5cdb9aa89991f9b11a917396c1b13ef1f904e2077a6.ExportPersonalDataRequestBuilder) {
    return i6302f791def46e7bc84de5cdb9aa89991f9b11a917396c1b13ef1f904e2077a6.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ib4a63511d54de38d4611349f200716dd80d232b496bea64c94ae3c4f452a78ba.FindMeetingTimesRequestBuilder) {
    return ib4a63511d54de38d4611349f200716dd80d232b496bea64c94ae3c4f452a78ba.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*ib79e33d64ee727d931363a7557bdcbad45675b9a52090505ff8c96f740174fd3.FindRoomListsRequestBuilder) {
    return ib79e33d64ee727d931363a7557bdcbad45675b9a52090505ff8c96f740174fd3.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i931235fc72db29a8627f2bf2928ee9d102bda6e555552570c18d50d51e82b02e.FindRoomsRequestBuilder) {
    return i931235fc72db29a8627f2bf2928ee9d102bda6e555552570c18d50d51e82b02e.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i8af688f61dc6212ca8ed456af66f481f6b1d42099b8f612c45afb698082f5776.FindRoomsWithRoomListRequestBuilder) {
    return i8af688f61dc6212ca8ed456af66f481f6b1d42099b8f612c45afb698082f5776.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i2926d108e8a370c469e897c5fd637bdbf4e59c41b578c927a748934065a7c30e.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i2926d108e8a370c469e897c5fd637bdbf4e59c41b578c927a748934065a7c30e.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i7669ec7815422fab8d6ff39effe64662f0e24d1411d1eb409e01105e6767e630.GetLoggedOnManagedDevicesRequestBuilder) {
    return i7669ec7815422fab8d6ff39effe64662f0e24d1411d1eb409e01105e6767e630.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i202a1b749f41f1fea470fc234005fe53b0ce43910e9fe65b70f52a2612e73250.GetMailTipsRequestBuilder) {
    return i202a1b749f41f1fea470fc234005fe53b0ce43910e9fe65b70f52a2612e73250.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i5ffe6993f05c4715a4d256de9bd0274640a46156296d0bd853c448f7f009bdf0.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i5ffe6993f05c4715a4d256de9bd0274640a46156296d0bd853c448f7f009bdf0.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i421a0a949e5e3697a51987047d6d893ff0deee857ac39ae51ff1fdb24ad959b6.GetManagedAppPoliciesRequestBuilder) {
    return i421a0a949e5e3697a51987047d6d893ff0deee857ac39ae51ff1fdb24ad959b6.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*id25033bd39711c598630f8d0d05b5fc928ff1883f7bb54b75112313fc15e12a9.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return id25033bd39711c598630f8d0d05b5fc928ff1883f7bb54b75112313fc15e12a9.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ibd538f065b0a60a4a14454443524e48284b038c0035f81a5449c7b63c6591375.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ibd538f065b0a60a4a14454443524e48284b038c0035f81a5449c7b63c6591375.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i54eeb9f2609c157ed5f58d71e74c146ed06117900495736137cfe1ff983806b6.GetMemberGroupsRequestBuilder) {
    return i54eeb9f2609c157ed5f58d71e74c146ed06117900495736137cfe1ff983806b6.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*iae5f51a30530619cb6ca1bf4cc0e83792d392de9017c11af442a7babd5675f90.GetMemberObjectsRequestBuilder) {
    return iae5f51a30530619cb6ca1bf4cc0e83792d392de9017c11af442a7babd5675f90.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*ibe9d16b18b45790efb8dd821f6f01fbf965060fab56914a5242336817600d4ff.InvalidateAllRefreshTokensRequestBuilder) {
    return ibe9d16b18b45790efb8dd821f6f01fbf965060fab56914a5242336817600d4ff.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ia08e746479382b3bc22cccc08ed34698744f70a527af0624f0d5742df55abac3.IsManagedAppUserBlockedRequestBuilder) {
    return ia08e746479382b3bc22cccc08ed34698744f70a527af0624f0d5742df55abac3.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*ie91934d6c0df672c3a4efe01bf837a147dd288c27ea0cc81cd8d0179d4b01eb0.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return ie91934d6c0df672c3a4efe01bf837a147dd288c27ea0cc81cd8d0179d4b01eb0.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i0b03e66923482b6bafce850b9faffe1ad626a7d337b96e54356341c4cfac6623.RemoveAllDevicesFromManagementRequestBuilder) {
    return i0b03e66923482b6bafce850b9faffe1ad626a7d337b96e54356341c4cfac6623.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*if3c4838af67ca41e08f6a15ff364545b3218d0733ae603b526c689b492577f10.ReprocessLicenseAssignmentRequestBuilder) {
    return if3c4838af67ca41e08f6a15ff364545b3218d0733ae603b526c689b492577f10.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i0580dd95c982b356e4fb1a5f9d8e9b45a1144aae2a1b04964c8a4d178256bf41.RestoreRequestBuilder) {
    return i0580dd95c982b356e4fb1a5f9d8e9b45a1144aae2a1b04964c8a4d178256bf41.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i809d3af15a46862830bb0a7bbe65141538eb61b2d47bc146069a1b7bad14eedb.RevokeSignInSessionsRequestBuilder) {
    return i809d3af15a46862830bb0a7bbe65141538eb61b2d47bc146069a1b7bad14eedb.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i0d1cea47b5a00a04ff83adca55cec4670ac81d1ded631fe5da7f6b1523e5cd76.SendMailRequestBuilder) {
    return i0d1cea47b5a00a04ff83adca55cec4670ac81d1ded631fe5da7f6b1523e5cd76.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i623a4c39b5f86bc611e4c7c79485e84315e8f204addbf279d18216e6fa4a0748.TranslateExchangeIdsRequestBuilder) {
    return i623a4c39b5f86bc611e4c7c79485e84315e8f204addbf279d18216e6fa4a0748.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i98699fe28f31f6aad757d2704a8b4d4a616f976d16a53129b6dc141e7567e81b.UnblockManagedAppsRequestBuilder) {
    return i98699fe28f31f6aad757d2704a8b4d4a616f976d16a53129b6dc141e7567e81b.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i78e74354ec09acd0b98b4b4600f804ec7887ce3b3341a45657c366d6956c2523.WipeAndBlockManagedAppsRequestBuilder) {
    return i78e74354ec09acd0b98b4b4600f804ec7887ce3b3341a45657c366d6956c2523.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*ie6479be6c19df4aec957773f6306aa32f894ebf5663b61db7e6d734044438e91.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return ie6479be6c19df4aec957773f6306aa32f894ebf5663b61db7e6d734044438e91.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ida81d9030db55ff7aafc84a5c314148c06e3c815b85338b274542c202ff12922.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ida81d9030db55ff7aafc84a5c314148c06e3c815b85338b274542c202ff12922.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ib382f2ce2d135054992a6f83d600cf7b7db0f5612acc12559d9791a5d46e0b56.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ib382f2ce2d135054992a6f83d600cf7b7db0f5612acc12559d9791a5d46e0b56.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
