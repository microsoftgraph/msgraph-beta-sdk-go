package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i007fec0b3a965e07c71f17f6427ef8604eeeb9b27794ca23df30733a59bcaa75 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    i0255c5470f73dc1a0a62537343e3bd3e15616d75969630784eb8ddb33826510e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i04c431170916fd940828f550194e7fbaab06f60a87ec5aa45f9e8f380574a99b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/unblockmanagedapps"
    i04c5b2c478c62aa52f44040fcb91c6b1d2b51ecee68d967687009c34d04ceed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    i0b209ab80eba0c3a0466e76aec7ab4ae07fdf8bdb6602cf43c70d55aad08920c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i13cfe5886134f42ccf7f29b2e60414ff623499ff56e3be5a218042547c31420c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i1a691f56c165483ed05af1ca09b2c3a73ae3084a9420a9c7810c2e18de3f2a4d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    i1aabccdbcbee9752be04e7f16686a77de3c1c3b85a6824c063e6062512d79aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findrooms"
    i1ab18d8d353548f39732422139a6a33da1ba12bf0e58174a8126a4b3a3f7ab1b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    i2ce5e7e3e49e206d51d5bb5fb8e1cfc8bee50d810e4d55af8c130431321c209e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    i2e9156ab0cd38b250cdc395e86db12bb4bcd573b3f9dff00ae5774276f51dbc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    i3bb7b5b74491ebdcb0fea95a62c7eea699d66d1fb6d08d2e00cecd4bfbdb6c7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
    i41df25f16a4348bf78e773719c5eea9678b264b2bb252999aec8a873b7d22c37 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    i7295b388107c0dae20946cd10b22ccac7608b3277add1ac19fd96c229c942536 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i7e4733aa9bbaa2311705b26b1acd96004b33fb54a9d9d36a7b982f879a4c0b06 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    i81dbe4d8daf3a1281f8beead36cb103ed343e48f3e66bb9273427ba7a3ddcbe7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findroomswithroomlist"
    i83f10dc8fc9801712d33a961856aeeda52889c5883fbd74ec91dce054ff3f0c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    i8c976a3a8cd372d1eef810e8b50f1dc5ed10fccd94bcbfaa4712dddfd2c6da47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/activateserviceplan"
    i8d04f20b206d9387bf0c98621c714799510d14ef8fe9631617b5b5d2531e2d3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findroomlists"
    i8e69ebadfb8b9907e7040ea6a81ab5774c81bad1ca6252fb755b7291f33c1fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    i9419c020a7815eb02a5ff67e8aa644c6f099fec901c326df18cddae52bf0db88 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    ia0e0d7d86e4f91b314d07619d266a90e8310bd4b4bc1f20e1303b0804af7861d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    ia96ae76aa544cbe9946f920165cd6f6356b77ff82b8d9b63d25ce096e558a365 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    iab283b8cfc68528edab6a01e85ce01b0f23dbac486547875a56d8ba8269bcf74 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    iaf73f50fe028e2395a2c6948134e47ad6be42fbeefdb9f968804133e00a317c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    ib8edc6864ed83193f9456050bfb5b3744ac8ce410cb275c9bbc79db1b91b7302 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    id1c41a700ce2d1558bb997d97218c0325bc42cf86847154b4ca21cbf553d5cad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    ie479b19a00979067ec8e0dcfb4073be45e2360a896ccd9db0db2db90a4481244 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    ie9fbc8b6f503bb27e82a47ff9523e2624b7127d941658c83dbc0dd98e2cc76b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    ieb7a8ff883ae51ec565c3101415016042c4004093ee6e79a57ebe67ada5d3dcd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    iec693967c264567f84902bc7b11e7c5a2a119e56453d4f6b4712d32d43005c82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    iecff993c4004ddd0febd375b6220a347dbe095b8be26153314494cc49ed18d24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    if02721c7d0ecc119eba123a0b8016423fbb3de130dcc416742f46a16aad705b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    ifb8185a928afd76461ff87d20f17acc0fbc717532555fe23a21815800986e9b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    ifec8446ad2ff0f59b50e42b64ce12913318ef4f01ec640be8e9ddf346e44bfec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i8c976a3a8cd372d1eef810e8b50f1dc5ed10fccd94bcbfaa4712dddfd2c6da47.ActivateServicePlanRequestBuilder) {
    return i8c976a3a8cd372d1eef810e8b50f1dc5ed10fccd94bcbfaa4712dddfd2c6da47.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*iab283b8cfc68528edab6a01e85ce01b0f23dbac486547875a56d8ba8269bcf74.AssignLicenseRequestBuilder) {
    return iab283b8cfc68528edab6a01e85ce01b0f23dbac486547875a56d8ba8269bcf74.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i7e4733aa9bbaa2311705b26b1acd96004b33fb54a9d9d36a7b982f879a4c0b06.ChangePasswordRequestBuilder) {
    return i7e4733aa9bbaa2311705b26b1acd96004b33fb54a9d9d36a7b982f879a4c0b06.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ie479b19a00979067ec8e0dcfb4073be45e2360a896ccd9db0db2db90a4481244.CheckMemberGroupsRequestBuilder) {
    return ie479b19a00979067ec8e0dcfb4073be45e2360a896ccd9db0db2db90a4481244.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i2ce5e7e3e49e206d51d5bb5fb8e1cfc8bee50d810e4d55af8c130431321c209e.CheckMemberObjectsRequestBuilder) {
    return i2ce5e7e3e49e206d51d5bb5fb8e1cfc8bee50d810e4d55af8c130431321c209e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*iec693967c264567f84902bc7b11e7c5a2a119e56453d4f6b4712d32d43005c82.ExportDeviceAndAppManagementDataRequestBuilder) {
    return iec693967c264567f84902bc7b11e7c5a2a119e56453d4f6b4712d32d43005c82.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*ia96ae76aa544cbe9946f920165cd6f6356b77ff82b8d9b63d25ce096e558a365.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return ia96ae76aa544cbe9946f920165cd6f6356b77ff82b8d9b63d25ce096e558a365.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i0b209ab80eba0c3a0466e76aec7ab4ae07fdf8bdb6602cf43c70d55aad08920c.ExportPersonalDataRequestBuilder) {
    return i0b209ab80eba0c3a0466e76aec7ab4ae07fdf8bdb6602cf43c70d55aad08920c.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i8e69ebadfb8b9907e7040ea6a81ab5774c81bad1ca6252fb755b7291f33c1fb6.FindMeetingTimesRequestBuilder) {
    return i8e69ebadfb8b9907e7040ea6a81ab5774c81bad1ca6252fb755b7291f33c1fb6.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*i8d04f20b206d9387bf0c98621c714799510d14ef8fe9631617b5b5d2531e2d3a.FindRoomListsRequestBuilder) {
    return i8d04f20b206d9387bf0c98621c714799510d14ef8fe9631617b5b5d2531e2d3a.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i1aabccdbcbee9752be04e7f16686a77de3c1c3b85a6824c063e6062512d79aaf.FindRoomsRequestBuilder) {
    return i1aabccdbcbee9752be04e7f16686a77de3c1c3b85a6824c063e6062512d79aaf.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i81dbe4d8daf3a1281f8beead36cb103ed343e48f3e66bb9273427ba7a3ddcbe7.FindRoomsWithRoomListRequestBuilder) {
    return i81dbe4d8daf3a1281f8beead36cb103ed343e48f3e66bb9273427ba7a3ddcbe7.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*iecff993c4004ddd0febd375b6220a347dbe095b8be26153314494cc49ed18d24.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return iecff993c4004ddd0febd375b6220a347dbe095b8be26153314494cc49ed18d24.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*ie9fbc8b6f503bb27e82a47ff9523e2624b7127d941658c83dbc0dd98e2cc76b9.GetLoggedOnManagedDevicesRequestBuilder) {
    return ie9fbc8b6f503bb27e82a47ff9523e2624b7127d941658c83dbc0dd98e2cc76b9.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*i83f10dc8fc9801712d33a961856aeeda52889c5883fbd74ec91dce054ff3f0c4.GetMailTipsRequestBuilder) {
    return i83f10dc8fc9801712d33a961856aeeda52889c5883fbd74ec91dce054ff3f0c4.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*i2e9156ab0cd38b250cdc395e86db12bb4bcd573b3f9dff00ae5774276f51dbc4.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i2e9156ab0cd38b250cdc395e86db12bb4bcd573b3f9dff00ae5774276f51dbc4.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*ieb7a8ff883ae51ec565c3101415016042c4004093ee6e79a57ebe67ada5d3dcd.GetManagedAppPoliciesRequestBuilder) {
    return ieb7a8ff883ae51ec565c3101415016042c4004093ee6e79a57ebe67ada5d3dcd.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i9419c020a7815eb02a5ff67e8aa644c6f099fec901c326df18cddae52bf0db88.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i9419c020a7815eb02a5ff67e8aa644c6f099fec901c326df18cddae52bf0db88.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*id1c41a700ce2d1558bb997d97218c0325bc42cf86847154b4ca21cbf553d5cad.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return id1c41a700ce2d1558bb997d97218c0325bc42cf86847154b4ca21cbf553d5cad.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i13cfe5886134f42ccf7f29b2e60414ff623499ff56e3be5a218042547c31420c.GetMemberGroupsRequestBuilder) {
    return i13cfe5886134f42ccf7f29b2e60414ff623499ff56e3be5a218042547c31420c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i04c5b2c478c62aa52f44040fcb91c6b1d2b51ecee68d967687009c34d04ceed2.GetMemberObjectsRequestBuilder) {
    return i04c5b2c478c62aa52f44040fcb91c6b1d2b51ecee68d967687009c34d04ceed2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i41df25f16a4348bf78e773719c5eea9678b264b2bb252999aec8a873b7d22c37.InvalidateAllRefreshTokensRequestBuilder) {
    return i41df25f16a4348bf78e773719c5eea9678b264b2bb252999aec8a873b7d22c37.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*i1a691f56c165483ed05af1ca09b2c3a73ae3084a9420a9c7810c2e18de3f2a4d.IsManagedAppUserBlockedRequestBuilder) {
    return i1a691f56c165483ed05af1ca09b2c3a73ae3084a9420a9c7810c2e18de3f2a4d.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*iaf73f50fe028e2395a2c6948134e47ad6be42fbeefdb9f968804133e00a317c1.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return iaf73f50fe028e2395a2c6948134e47ad6be42fbeefdb9f968804133e00a317c1.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i7295b388107c0dae20946cd10b22ccac7608b3277add1ac19fd96c229c942536.RemoveAllDevicesFromManagementRequestBuilder) {
    return i7295b388107c0dae20946cd10b22ccac7608b3277add1ac19fd96c229c942536.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*ifb8185a928afd76461ff87d20f17acc0fbc717532555fe23a21815800986e9b9.ReprocessLicenseAssignmentRequestBuilder) {
    return ifb8185a928afd76461ff87d20f17acc0fbc717532555fe23a21815800986e9b9.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*if02721c7d0ecc119eba123a0b8016423fbb3de130dcc416742f46a16aad705b4.RestoreRequestBuilder) {
    return if02721c7d0ecc119eba123a0b8016423fbb3de130dcc416742f46a16aad705b4.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ia0e0d7d86e4f91b314d07619d266a90e8310bd4b4bc1f20e1303b0804af7861d.RevokeSignInSessionsRequestBuilder) {
    return ia0e0d7d86e4f91b314d07619d266a90e8310bd4b4bc1f20e1303b0804af7861d.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*i3bb7b5b74491ebdcb0fea95a62c7eea699d66d1fb6d08d2e00cecd4bfbdb6c7e.SendMailRequestBuilder) {
    return i3bb7b5b74491ebdcb0fea95a62c7eea699d66d1fb6d08d2e00cecd4bfbdb6c7e.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*i1ab18d8d353548f39732422139a6a33da1ba12bf0e58174a8126a4b3a3f7ab1b.TranslateExchangeIdsRequestBuilder) {
    return i1ab18d8d353548f39732422139a6a33da1ba12bf0e58174a8126a4b3a3f7ab1b.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i04c431170916fd940828f550194e7fbaab06f60a87ec5aa45f9e8f380574a99b.UnblockManagedAppsRequestBuilder) {
    return i04c431170916fd940828f550194e7fbaab06f60a87ec5aa45f9e8f380574a99b.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i0255c5470f73dc1a0a62537343e3bd3e15616d75969630784eb8ddb33826510e.WipeAndBlockManagedAppsRequestBuilder) {
    return i0255c5470f73dc1a0a62537343e3bd3e15616d75969630784eb8ddb33826510e.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i007fec0b3a965e07c71f17f6427ef8604eeeb9b27794ca23df30733a59bcaa75.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i007fec0b3a965e07c71f17f6427ef8604eeeb9b27794ca23df30733a59bcaa75.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*ifec8446ad2ff0f59b50e42b64ce12913318ef4f01ec640be8e9ddf346e44bfec.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return ifec8446ad2ff0f59b50e42b64ce12913318ef4f01ec640be8e9ddf346e44bfec.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ib8edc6864ed83193f9456050bfb5b3744ac8ce410cb275c9bbc79db1b91b7302.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ib8edc6864ed83193f9456050bfb5b3744ac8ce410cb275c9bbc79db1b91b7302.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
