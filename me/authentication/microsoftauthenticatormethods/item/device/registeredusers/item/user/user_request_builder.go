package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i05401c38a2ff845e7506f3c966111a10c70cdfcb7009f296fa3529a09be08cd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/removealldevicesfrommanagement"
    i0cc6f6dafe24dcfd9aadc470c721dc4c7bbd45ef4b58b2bf96c3aebc3b3be9b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/changepassword"
    i0f3cee9d84cb6241ffc88176838b79e481bde0b67e41707add90338fcd1a9be1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/invalidateallrefreshtokens"
    i1578ce2996dbd32c17a6332f08b09baf3c7edefdff3d5657bc20925848bdc3fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedapppolicies"
    i1617f4e269fe86fab6463837507bb353ccfa29a1ad4d93b40d44155815b735e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/restore"
    i2d7fc29e3cda9661c99bbd18276774ba24eeae059eb61b16f0539cd05deab300 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdatawithskipwithtop"
    i3737ebd5572d69d2715ddda26bfe2ae4d729d4d1eaf18e982348cc76e30b079c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipeandblockmanagedapps"
    i3ffe64f35f4121932436d9775de667b1a624bb6f2b5e099d46b9514f8efa63d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reminderviewwithstartdatetimewithenddatetime"
    i401033a8300d72d35fd8d7ce60abfe901cf8331875838df53810cf339c55a47c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/unblockmanagedapps"
    i42e396ed3fe7369efd500eb96c7851522b42ddf5ef5fec0721a5442a9127bc82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportpersonaldata"
    i496021d32cebe28e2a991551d29ddf05d5908270bcf572b7186ffa2d6c8656f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/reprocesslicenseassignment"
    i4b7f7148766c788ed5288450e25e816a8ca40fbb452bb8f564d4d5af2601104f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmemberobjects"
    i4c3928560e486f0c0599ddb1170e82751da66005a202725808897d1b9c49a24f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findrooms"
    i71a068f09e78a3cf78c551533bdbf31f82a1307afea8f4e469e0543bc02a5f0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/activateserviceplan"
    i7396345b1b30b277e242c724ccad71ff21739f1e3cdeb5e1a63a3cae58e4a5fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmembergroups"
    i82b32fa5a26a1184799bace930a38f815db967d58671b2e69bdcf76571d22f22 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithappfailures"
    i8ad242f6f2ea7a531e9d40164dc277990904bd851b2e64c9a4def07ac6d7e491 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationbydevicetag"
    i97a94fb2da7a1c9f537f6dcd5c00a9b9f33d0d21cb5aec4adfecb763b2775d43 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findmeetingtimes"
    i9dae52c29acf44c21289c837a7b318ab1947b799f8f977714aa9078bb75c62ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmemberobjects"
    i9fc53566fb828cab1efa5c1e7f51121f170bbb2e0e6934d8ca896d6940f850dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getloggedonmanageddevices"
    ia500b81cbdf70ef3674787704da3e6caa87ea85a53fc0a7489e9f9dd8a254813 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/revokesigninsessions"
    ia91f4e56e45ba82d7e3d2cfc9fec7d8c0241ae61103f530de5fdee2aa1870e73 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findroomswithroomlist"
    iab25ce47a0333bc4708a8ebaf81429a20769a7f5475eafce2d6785c0cc782311 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmailtips"
    iadee832d838640fdd6dbd05bdf87333b3c29fa9f0b068542f510d6fd6924a03d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/translateexchangeids"
    iba3fc4df1143344f173464b2fd9eaad9e64aac3d1817e519837ddae0f61a677c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/sendmail"
    ibf348a07774be49675329719dcf3f60716d98887178b7d32801d3fee664bbd61 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/assignlicense"
    ibfbf49f7c77c702270c40c812f2b3501bae057b0875e57739f77428bebff189d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/exportdeviceandappmanagementdata"
    ic00a010a0b31b973a04252122b304bfa39d1348b0b6d94c254c9e78161f1506f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/geteffectivedeviceenrollmentconfigurations"
    ic1a807c2bdd5ca3e6c209e08e0113afb813088aadbcdb4aebefb13a212e86b50 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/checkmembergroups"
    ic57c7e91a898de3aa79ae16e46160111aecb107acbaf11355c62ad9952d1e114 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/ismanagedappuserblocked"
    iccb63e5153c47970deb14544b4b776957ee82ce917073ea56d350be8a60bf36c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbydevicetag"
    icce39d438772b33aba29d3e4a9b7d4106ce34a855f60504b37ae0708b1c49f5b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanageddeviceswithfailedorpendingapps"
    id01b7339c473a740c63e6f15b4cf1e4e626460f1bb809cf61b10fad56d6375a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/getmanagedappdiagnosticstatuses"
    id3e91120c2a8a2e923da7b9b6274bd5fdce6ba0f9e68c5c293f9c7714bb8a449 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/wipemanagedappregistrationsbyazureaddeviceid"
    if56270f5a88c39bdc6bccde7473d7462bcefeff6892e488c125405041cf9d483 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredusers/item/user/findroomlists"
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
func (m *UserRequestBuilder) ActivateServicePlan()(*i71a068f09e78a3cf78c551533bdbf31f82a1307afea8f4e469e0543bc02a5f0c.ActivateServicePlanRequestBuilder) {
    return i71a068f09e78a3cf78c551533bdbf31f82a1307afea8f4e469e0543bc02a5f0c.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *UserRequestBuilder) AssignLicense()(*ibf348a07774be49675329719dcf3f60716d98887178b7d32801d3fee664bbd61.AssignLicenseRequestBuilder) {
    return ibf348a07774be49675329719dcf3f60716d98887178b7d32801d3fee664bbd61.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *UserRequestBuilder) ChangePassword()(*i0cc6f6dafe24dcfd9aadc470c721dc4c7bbd45ef4b58b2bf96c3aebc3b3be9b2.ChangePasswordRequestBuilder) {
    return i0cc6f6dafe24dcfd9aadc470c721dc4c7bbd45ef4b58b2bf96c3aebc3b3be9b2.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *UserRequestBuilder) CheckMemberGroups()(*ic1a807c2bdd5ca3e6c209e08e0113afb813088aadbcdb4aebefb13a212e86b50.CheckMemberGroupsRequestBuilder) {
    return ic1a807c2bdd5ca3e6c209e08e0113afb813088aadbcdb4aebefb13a212e86b50.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *UserRequestBuilder) CheckMemberObjects()(*i4b7f7148766c788ed5288450e25e816a8ca40fbb452bb8f564d4d5af2601104f.CheckMemberObjectsRequestBuilder) {
    return i4b7f7148766c788ed5288450e25e816a8ca40fbb452bb8f564d4d5af2601104f.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}/microsoft.graph.user{?%24select,%24expand}";
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
func (m *UserRequestBuilder) ExportDeviceAndAppManagementData()(*ibfbf49f7c77c702270c40c812f2b3501bae057b0875e57739f77428bebff189d.ExportDeviceAndAppManagementDataRequestBuilder) {
    return ibfbf49f7c77c702270c40c812f2b3501bae057b0875e57739f77428bebff189d.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *UserRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i2d7fc29e3cda9661c99bbd18276774ba24eeae059eb61b16f0539cd05deab300.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i2d7fc29e3cda9661c99bbd18276774ba24eeae059eb61b16f0539cd05deab300.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *UserRequestBuilder) ExportPersonalData()(*i42e396ed3fe7369efd500eb96c7851522b42ddf5ef5fec0721a5442a9127bc82.ExportPersonalDataRequestBuilder) {
    return i42e396ed3fe7369efd500eb96c7851522b42ddf5ef5fec0721a5442a9127bc82.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *UserRequestBuilder) FindMeetingTimes()(*i97a94fb2da7a1c9f537f6dcd5c00a9b9f33d0d21cb5aec4adfecb763b2775d43.FindMeetingTimesRequestBuilder) {
    return i97a94fb2da7a1c9f537f6dcd5c00a9b9f33d0d21cb5aec4adfecb763b2775d43.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *UserRequestBuilder) FindRoomLists()(*if56270f5a88c39bdc6bccde7473d7462bcefeff6892e488c125405041cf9d483.FindRoomListsRequestBuilder) {
    return if56270f5a88c39bdc6bccde7473d7462bcefeff6892e488c125405041cf9d483.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRooms()(*i4c3928560e486f0c0599ddb1170e82751da66005a202725808897d1b9c49a24f.FindRoomsRequestBuilder) {
    return i4c3928560e486f0c0599ddb1170e82751da66005a202725808897d1b9c49a24f.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *UserRequestBuilder) FindRoomsWithRoomList(roomList *string)(*ia91f4e56e45ba82d7e3d2cfc9fec7d8c0241ae61103f530de5fdee2aa1870e73.FindRoomsWithRoomListRequestBuilder) {
    return ia91f4e56e45ba82d7e3d2cfc9fec7d8c0241ae61103f530de5fdee2aa1870e73.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.user
func (m *UserRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *UserRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*ic00a010a0b31b973a04252122b304bfa39d1348b0b6d94c254c9e78161f1506f.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return ic00a010a0b31b973a04252122b304bfa39d1348b0b6d94c254c9e78161f1506f.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *UserRequestBuilder) GetLoggedOnManagedDevices()(*i9fc53566fb828cab1efa5c1e7f51121f170bbb2e0e6934d8ca896d6940f850dc.GetLoggedOnManagedDevicesRequestBuilder) {
    return i9fc53566fb828cab1efa5c1e7f51121f170bbb2e0e6934d8ca896d6940f850dc.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *UserRequestBuilder) GetMailTips()(*iab25ce47a0333bc4708a8ebaf81429a20769a7f5475eafce2d6785c0cc782311.GetMailTipsRequestBuilder) {
    return iab25ce47a0333bc4708a8ebaf81429a20769a7f5475eafce2d6785c0cc782311.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *UserRequestBuilder) GetManagedAppDiagnosticStatuses()(*id01b7339c473a740c63e6f15b4cf1e4e626460f1bb809cf61b10fad56d6375a5.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return id01b7339c473a740c63e6f15b4cf1e4e626460f1bb809cf61b10fad56d6375a5.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *UserRequestBuilder) GetManagedAppPolicies()(*i1578ce2996dbd32c17a6332f08b09baf3c7edefdff3d5657bc20925848bdc3fd.GetManagedAppPoliciesRequestBuilder) {
    return i1578ce2996dbd32c17a6332f08b09baf3c7edefdff3d5657bc20925848bdc3fd.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *UserRequestBuilder) GetManagedDevicesWithAppFailures()(*i82b32fa5a26a1184799bace930a38f815db967d58671b2e69bdcf76571d22f22.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i82b32fa5a26a1184799bace930a38f815db967d58671b2e69bdcf76571d22f22.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *UserRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*icce39d438772b33aba29d3e4a9b7d4106ce34a855f60504b37ae0708b1c49f5b.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return icce39d438772b33aba29d3e4a9b7d4106ce34a855f60504b37ae0708b1c49f5b.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *UserRequestBuilder) GetMemberGroups()(*i7396345b1b30b277e242c724ccad71ff21739f1e3cdeb5e1a63a3cae58e4a5fc.GetMemberGroupsRequestBuilder) {
    return i7396345b1b30b277e242c724ccad71ff21739f1e3cdeb5e1a63a3cae58e4a5fc.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *UserRequestBuilder) GetMemberObjects()(*i9dae52c29acf44c21289c837a7b318ab1947b799f8f977714aa9078bb75c62ba.GetMemberObjectsRequestBuilder) {
    return i9dae52c29acf44c21289c837a7b318ab1947b799f8f977714aa9078bb75c62ba.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *UserRequestBuilder) InvalidateAllRefreshTokens()(*i0f3cee9d84cb6241ffc88176838b79e481bde0b67e41707add90338fcd1a9be1.InvalidateAllRefreshTokensRequestBuilder) {
    return i0f3cee9d84cb6241ffc88176838b79e481bde0b67e41707add90338fcd1a9be1.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *UserRequestBuilder) IsManagedAppUserBlocked()(*ic57c7e91a898de3aa79ae16e46160111aecb107acbaf11355c62ad9952d1e114.IsManagedAppUserBlockedRequestBuilder) {
    return ic57c7e91a898de3aa79ae16e46160111aecb107acbaf11355c62ad9952d1e114.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *UserRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*i3ffe64f35f4121932436d9775de667b1a624bb6f2b5e099d46b9514f8efa63d3.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i3ffe64f35f4121932436d9775de667b1a624bb6f2b5e099d46b9514f8efa63d3.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *UserRequestBuilder) RemoveAllDevicesFromManagement()(*i05401c38a2ff845e7506f3c966111a10c70cdfcb7009f296fa3529a09be08cd6.RemoveAllDevicesFromManagementRequestBuilder) {
    return i05401c38a2ff845e7506f3c966111a10c70cdfcb7009f296fa3529a09be08cd6.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *UserRequestBuilder) ReprocessLicenseAssignment()(*i496021d32cebe28e2a991551d29ddf05d5908270bcf572b7186ffa2d6c8656f2.ReprocessLicenseAssignmentRequestBuilder) {
    return i496021d32cebe28e2a991551d29ddf05d5908270bcf572b7186ffa2d6c8656f2.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *UserRequestBuilder) Restore()(*i1617f4e269fe86fab6463837507bb353ccfa29a1ad4d93b40d44155815b735e2.RestoreRequestBuilder) {
    return i1617f4e269fe86fab6463837507bb353ccfa29a1ad4d93b40d44155815b735e2.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *UserRequestBuilder) RevokeSignInSessions()(*ia500b81cbdf70ef3674787704da3e6caa87ea85a53fc0a7489e9f9dd8a254813.RevokeSignInSessionsRequestBuilder) {
    return ia500b81cbdf70ef3674787704da3e6caa87ea85a53fc0a7489e9f9dd8a254813.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *UserRequestBuilder) SendMail()(*iba3fc4df1143344f173464b2fd9eaad9e64aac3d1817e519837ddae0f61a677c.SendMailRequestBuilder) {
    return iba3fc4df1143344f173464b2fd9eaad9e64aac3d1817e519837ddae0f61a677c.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *UserRequestBuilder) TranslateExchangeIds()(*iadee832d838640fdd6dbd05bdf87333b3c29fa9f0b068542f510d6fd6924a03d.TranslateExchangeIdsRequestBuilder) {
    return iadee832d838640fdd6dbd05bdf87333b3c29fa9f0b068542f510d6fd6924a03d.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *UserRequestBuilder) UnblockManagedApps()(*i401033a8300d72d35fd8d7ce60abfe901cf8331875838df53810cf339c55a47c.UnblockManagedAppsRequestBuilder) {
    return i401033a8300d72d35fd8d7ce60abfe901cf8331875838df53810cf339c55a47c.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *UserRequestBuilder) WipeAndBlockManagedApps()(*i3737ebd5572d69d2715ddda26bfe2ae4d729d4d1eaf18e982348cc76e30b079c.WipeAndBlockManagedAppsRequestBuilder) {
    return i3737ebd5572d69d2715ddda26bfe2ae4d729d4d1eaf18e982348cc76e30b079c.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*i8ad242f6f2ea7a531e9d40164dc277990904bd851b2e64c9a4def07ac6d7e491.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return i8ad242f6f2ea7a531e9d40164dc277990904bd851b2e64c9a4def07ac6d7e491.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*id3e91120c2a8a2e923da7b9b6274bd5fdce6ba0f9e68c5c293f9c7714bb8a449.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return id3e91120c2a8a2e923da7b9b6274bd5fdce6ba0f9e68c5c293f9c7714bb8a449.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*iccb63e5153c47970deb14544b4b776957ee82ce917073ea56d350be8a60bf36c.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return iccb63e5153c47970deb14544b4b776957ee82ce917073ea56d350be8a60bf36c.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
