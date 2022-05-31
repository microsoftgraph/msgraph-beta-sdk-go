package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0264c4ca12039672ded7292bbe6669606e22cd546a5e72af0fef66ddbb8a673c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/assignlicense"
    i0555e6ba9f67c8435107a6e9bca0efa144d4ef82228cd5b0defd8208285eb096 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/reminderviewwithstartdatetimewithenddatetime"
    i0dcc377f915da35fe0e6ad8652c4785d01b45679390a2dccfa47761181cf45eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    i0f982e20a30cf0f8b2588a0fcd4808bbecf32e9dd88342ef3df7ff6c7a30c9f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/reprocesslicenseassignment"
    i1b3f4dd175aaf9849c96cd6a22901bfc5213e22648fc4e4d0c558222370e1be3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/ismanagedappuserblocked"
    i2b9f05f288fc864078b6249c0e9b27f2d606174d7cf4413220c32a48233f323f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/unblockmanagedapps"
    i37122fddb96170635985ad36f7d774bb261618536ad6e6ea6cbd524d5707e56f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/activateserviceplan"
    i46794471075c1c919a84e734db895c5e1d3ef62190bc129d65c8d92fb0717e12 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getloggedonmanageddevices"
    i56c3bea3b26bf15450282abcaab987cd8a935a673d06b8584463b422c2d6efc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/revokesigninsessions"
    i57bc1367d5c01c62f841a2a8d267ee5c656a5983d223792224760ff514d7aea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/findroomlists"
    i58c7d242c974d43457fbfb2a5ebbb49c4e7f2d298a6b26eff3336ccaf4e8a37f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/checkmemberobjects"
    i5985b00b55ae23511ef9a90c37de892f6a60e345e33f1ee2c2065abe83781bf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/changepassword"
    i6393b1a5ca1cf6deea09d7a3e4b13cec0b47164a9c5bc2455e8629f30e51dbf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/translateexchangeids"
    i6559bbcce79f7e2f2e15c83e2a5e6aea921c19d27148d592b6241154ea114a38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/findrooms"
    i6a141803e4d427ba311abe5d59ac524af2d89937c39f3662671b2a9170f90528 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmembergroups"
    i7dca302e0d44eec0a519a93c25d5b3434059bf8cea5ff32a4b75c76818ef1823 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmemberobjects"
    i84b24ac6351c87bb54d74c1349d2b5f28c8373d0390f603b6e8fef916a029e5f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/exportdeviceandappmanagementdata"
    i961ab51d6b6436b52b6ae6ab2e42f38138e8290cb84498d7c8ee43d3e108fc9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/findroomswithroomlist"
    i96eaac76da372a370763ca0ea6da4e02484cf81b61a7e4e1ca5ace28b111d44f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/exportpersonaldata"
    i9bf64c354c48f58e22d0f0a63c4ee9274b4f47095c0927d0e4df445353e86dfb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/wipemanagedappregistrationbydevicetag"
    i9c374ef3eeac7090cad80c5f6e561d4bcbdb2cf53511147b5c3b21c03ad5a9aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/invalidateallrefreshtokens"
    ia4bb5fd9711b2fad95bf6876a640210412d8bbff9d44367d9fb7492cb98eddbd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/sendmail"
    ib7a9c593c531d91cac5a7d682fc36414b9bff4fe0e04779e8773c1dd43a0b730 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/removealldevicesfrommanagement"
    ibb03e676e838897d18ecfcfb8e02de10d5d33cfde862cf9598d26c5cf50243e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    ibc9ac011791d494750689288b0fdcc95715ecda16041beb0625a1814b4c48da0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmailtips"
    ibfbbb2c16b070e09d77ee28c08ed31d89430fb822fd27d3cdbb1964b5a804880 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/findmeetingtimes"
    ic02b783e7f9033e1637d07ed7e6a3960cc4c0a4ed07f42436228ed911710621a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/wipeandblockmanagedapps"
    ic1950bcf312a255fc57d28694068efd35398b9f203f0fd2b2691f24e0f280cbe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmanagedapppolicies"
    ic2f57e69b4d186453a6030d8c52e3cb6876fef7310ade1cee8fc60f0489835ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/geteffectivedeviceenrollmentconfigurations"
    ic98161422c8514306ce482cc385f11902f163641ddba9655463253c4bf317907 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmanageddeviceswithappfailures"
    id61669c81ae88c672ed1513ee2716a552aa682e46c040207077976fdbe76ab5e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/checkmembergroups"
    ieffcef10ba2ea81c645f5277a9f855c365b5a3f711cd73dab68aa1c34dbd758b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/wipemanagedappregistrationsbydevicetag"
    if23ea5597dd6b1fde1fb4dcc1b5ebf010c9b6d21e25fc3df7541981f4fa6cc8a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmanageddeviceswithfailedorpendingapps"
    if684996f56737ba66d52956a060b2b8bd636fc377d766ce8f7e01aaf16ba365c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/restore"
    iff18cf615c72dd07f75b6db8fb8a09bd93c96a876556c0707c635064860a5a8c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user/getmanagedappdiagnosticstatuses"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i37122fddb96170635985ad36f7d774bb261618536ad6e6ea6cbd524d5707e56f.ActivateServicePlanRequestBuilder) {
    return i37122fddb96170635985ad36f7d774bb261618536ad6e6ea6cbd524d5707e56f.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*i0264c4ca12039672ded7292bbe6669606e22cd546a5e72af0fef66ddbb8a673c.AssignLicenseRequestBuilder) {
    return i0264c4ca12039672ded7292bbe6669606e22cd546a5e72af0fef66ddbb8a673c.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i5985b00b55ae23511ef9a90c37de892f6a60e345e33f1ee2c2065abe83781bf4.ChangePasswordRequestBuilder) {
    return i5985b00b55ae23511ef9a90c37de892f6a60e345e33f1ee2c2065abe83781bf4.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*id61669c81ae88c672ed1513ee2716a552aa682e46c040207077976fdbe76ab5e.CheckMemberGroupsRequestBuilder) {
    return id61669c81ae88c672ed1513ee2716a552aa682e46c040207077976fdbe76ab5e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i58c7d242c974d43457fbfb2a5ebbb49c4e7f2d298a6b26eff3336ccaf4e8a37f.CheckMemberObjectsRequestBuilder) {
    return i58c7d242c974d43457fbfb2a5ebbb49c4e7f2d298a6b26eff3336ccaf4e8a37f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*i84b24ac6351c87bb54d74c1349d2b5f28c8373d0390f603b6e8fef916a029e5f.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i84b24ac6351c87bb54d74c1349d2b5f28c8373d0390f603b6e8fef916a029e5f.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ibb03e676e838897d18ecfcfb8e02de10d5d33cfde862cf9598d26c5cf50243e8.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ibb03e676e838897d18ecfcfb8e02de10d5d33cfde862cf9598d26c5cf50243e8.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i96eaac76da372a370763ca0ea6da4e02484cf81b61a7e4e1ca5ace28b111d44f.ExportPersonalDataRequestBuilder) {
    return i96eaac76da372a370763ca0ea6da4e02484cf81b61a7e4e1ca5ace28b111d44f.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*ibfbbb2c16b070e09d77ee28c08ed31d89430fb822fd27d3cdbb1964b5a804880.FindMeetingTimesRequestBuilder) {
    return ibfbbb2c16b070e09d77ee28c08ed31d89430fb822fd27d3cdbb1964b5a804880.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i57bc1367d5c01c62f841a2a8d267ee5c656a5983d223792224760ff514d7aea2.FindRoomListsRequestBuilder) {
    return i57bc1367d5c01c62f841a2a8d267ee5c656a5983d223792224760ff514d7aea2.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i6559bbcce79f7e2f2e15c83e2a5e6aea921c19d27148d592b6241154ea114a38.FindRoomsRequestBuilder) {
    return i6559bbcce79f7e2f2e15c83e2a5e6aea921c19d27148d592b6241154ea114a38.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i961ab51d6b6436b52b6ae6ab2e42f38138e8290cb84498d7c8ee43d3e108fc9d.FindRoomsWithRoomListRequestBuilder) {
    return i961ab51d6b6436b52b6ae6ab2e42f38138e8290cb84498d7c8ee43d3e108fc9d.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ic2f57e69b4d186453a6030d8c52e3cb6876fef7310ade1cee8fc60f0489835ba.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ic2f57e69b4d186453a6030d8c52e3cb6876fef7310ade1cee8fc60f0489835ba.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i46794471075c1c919a84e734db895c5e1d3ef62190bc129d65c8d92fb0717e12.GetLoggedOnManagedDevicesRequestBuilder) {
    return i46794471075c1c919a84e734db895c5e1d3ef62190bc129d65c8d92fb0717e12.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*ibc9ac011791d494750689288b0fdcc95715ecda16041beb0625a1814b4c48da0.GetMailTipsRequestBuilder) {
    return ibc9ac011791d494750689288b0fdcc95715ecda16041beb0625a1814b4c48da0.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*iff18cf615c72dd07f75b6db8fb8a09bd93c96a876556c0707c635064860a5a8c.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return iff18cf615c72dd07f75b6db8fb8a09bd93c96a876556c0707c635064860a5a8c.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ic1950bcf312a255fc57d28694068efd35398b9f203f0fd2b2691f24e0f280cbe.GetManagedAppPoliciesRequestBuilder) {
    return ic1950bcf312a255fc57d28694068efd35398b9f203f0fd2b2691f24e0f280cbe.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*ic98161422c8514306ce482cc385f11902f163641ddba9655463253c4bf317907.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return ic98161422c8514306ce482cc385f11902f163641ddba9655463253c4bf317907.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*if23ea5597dd6b1fde1fb4dcc1b5ebf010c9b6d21e25fc3df7541981f4fa6cc8a.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return if23ea5597dd6b1fde1fb4dcc1b5ebf010c9b6d21e25fc3df7541981f4fa6cc8a.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i6a141803e4d427ba311abe5d59ac524af2d89937c39f3662671b2a9170f90528.GetMemberGroupsRequestBuilder) {
    return i6a141803e4d427ba311abe5d59ac524af2d89937c39f3662671b2a9170f90528.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i7dca302e0d44eec0a519a93c25d5b3434059bf8cea5ff32a4b75c76818ef1823.GetMemberObjectsRequestBuilder) {
    return i7dca302e0d44eec0a519a93c25d5b3434059bf8cea5ff32a4b75c76818ef1823.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i9c374ef3eeac7090cad80c5f6e561d4bcbdb2cf53511147b5c3b21c03ad5a9aa.InvalidateAllRefreshTokensRequestBuilder) {
    return i9c374ef3eeac7090cad80c5f6e561d4bcbdb2cf53511147b5c3b21c03ad5a9aa.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1b3f4dd175aaf9849c96cd6a22901bfc5213e22648fc4e4d0c558222370e1be3.IsManagedAppUserBlockedRequestBuilder) {
    return i1b3f4dd175aaf9849c96cd6a22901bfc5213e22648fc4e4d0c558222370e1be3.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i0555e6ba9f67c8435107a6e9bca0efa144d4ef82228cd5b0defd8208285eb096.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i0555e6ba9f67c8435107a6e9bca0efa144d4ef82228cd5b0defd8208285eb096.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*ib7a9c593c531d91cac5a7d682fc36414b9bff4fe0e04779e8773c1dd43a0b730.RemoveAllDevicesFromManagementRequestBuilder) {
    return ib7a9c593c531d91cac5a7d682fc36414b9bff4fe0e04779e8773c1dd43a0b730.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i0f982e20a30cf0f8b2588a0fcd4808bbecf32e9dd88342ef3df7ff6c7a30c9f9.ReprocessLicenseAssignmentRequestBuilder) {
    return i0f982e20a30cf0f8b2588a0fcd4808bbecf32e9dd88342ef3df7ff6c7a30c9f9.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*if684996f56737ba66d52956a060b2b8bd636fc377d766ce8f7e01aaf16ba365c.RestoreRequestBuilder) {
    return if684996f56737ba66d52956a060b2b8bd636fc377d766ce8f7e01aaf16ba365c.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*i56c3bea3b26bf15450282abcaab987cd8a935a673d06b8584463b422c2d6efc3.RevokeSignInSessionsRequestBuilder) {
    return i56c3bea3b26bf15450282abcaab987cd8a935a673d06b8584463b422c2d6efc3.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*ia4bb5fd9711b2fad95bf6876a640210412d8bbff9d44367d9fb7492cb98eddbd.SendMailRequestBuilder) {
    return ia4bb5fd9711b2fad95bf6876a640210412d8bbff9d44367d9fb7492cb98eddbd.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i6393b1a5ca1cf6deea09d7a3e4b13cec0b47164a9c5bc2455e8629f30e51dbf0.TranslateExchangeIdsRequestBuilder) {
    return i6393b1a5ca1cf6deea09d7a3e4b13cec0b47164a9c5bc2455e8629f30e51dbf0.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i2b9f05f288fc864078b6249c0e9b27f2d606174d7cf4413220c32a48233f323f.UnblockManagedAppsRequestBuilder) {
    return i2b9f05f288fc864078b6249c0e9b27f2d606174d7cf4413220c32a48233f323f.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*ic02b783e7f9033e1637d07ed7e6a3960cc4c0a4ed07f42436228ed911710621a.WipeAndBlockManagedAppsRequestBuilder) {
    return ic02b783e7f9033e1637d07ed7e6a3960cc4c0a4ed07f42436228ed911710621a.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i9bf64c354c48f58e22d0f0a63c4ee9274b4f47095c0927d0e4df445353e86dfb.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i9bf64c354c48f58e22d0f0a63c4ee9274b4f47095c0927d0e4df445353e86dfb.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i0dcc377f915da35fe0e6ad8652c4785d01b45679390a2dccfa47761181cf45eb.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i0dcc377f915da35fe0e6ad8652c4785d01b45679390a2dccfa47761181cf45eb.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ieffcef10ba2ea81c645f5277a9f855c365b5a3f711cd73dab68aa1c34dbd758b.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ieffcef10ba2ea81c645f5277a9f855c365b5a3f711cd73dab68aa1c34dbd758b.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
