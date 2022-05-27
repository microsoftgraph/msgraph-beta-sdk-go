package memberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b47b7f619459696f6a035eedcd3bac8bd7a14f8aa6779da78b38764c65b9c55 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/device"
    i231078149d04ad3ba9c01158263b8cc7b91cb05562f88cfb6c8df354ed43cfdc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/application"
    i54741bc03e9cffff30b845e40f19242a2ccc1c3fd84338efeb85cc34a7d79795 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/group"
    i57a9b05210a746928010661e9c18014f8e235ff52afa313241d78c4d64f35435 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/user"
    i80effb16f1b9e1ee6056a140e9667c91b5fc622cf03bcfc6a53dd629d66d3907 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/serviceprincipal"
    i8fd97f95521da7b21f1e73fe02024d26e4274afe748e9effe5aaf77bd466e3f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/count"
    i9e5db44b01651144e5090d3413d67d004b1574f466555bf946d53c018a025725 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/memberof/orgcontact"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MemberOfRequestBuilderGetQueryParameters
}
// Application the application property
func (m *MemberOfRequestBuilder) Application()(*i231078149d04ad3ba9c01158263b8cc7b91cb05562f88cfb6c8df354ed43cfdc.ApplicationRequestBuilder) {
    return i231078149d04ad3ba9c01158263b8cc7b91cb05562f88cfb6c8df354ed43cfdc.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *MemberOfRequestBuilder) Count()(*i8fd97f95521da7b21f1e73fe02024d26e4274afe748e9effe5aaf77bd466e3f5.CountRequestBuilder) {
    return i8fd97f95521da7b21f1e73fe02024d26e4274afe748e9effe5aaf77bd466e3f5.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device the device property
func (m *MemberOfRequestBuilder) Device()(*i0b47b7f619459696f6a035eedcd3bac8bd7a14f8aa6779da78b38764c65b9c55.DeviceRequestBuilder) {
    return i0b47b7f619459696f6a035eedcd3bac8bd7a14f8aa6779da78b38764c65b9c55.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *MemberOfRequestBuilder) Group()(*i54741bc03e9cffff30b845e40f19242a2ccc1c3fd84338efeb85cc34a7d79795.GroupRequestBuilder) {
    return i54741bc03e9cffff30b845e40f19242a2ccc1c3fd84338efeb85cc34a7d79795.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MemberOfRequestBuilder) OrgContact()(*i9e5db44b01651144e5090d3413d67d004b1574f466555bf946d53c018a025725.OrgContactRequestBuilder) {
    return i9e5db44b01651144e5090d3413d67d004b1574f466555bf946d53c018a025725.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i80effb16f1b9e1ee6056a140e9667c91b5fc622cf03bcfc6a53dd629d66d3907.ServicePrincipalRequestBuilder) {
    return i80effb16f1b9e1ee6056a140e9667c91b5fc622cf03bcfc6a53dd629d66d3907.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MemberOfRequestBuilder) User()(*i57a9b05210a746928010661e9c18014f8e235ff52afa313241d78c4d64f35435.UserRequestBuilder) {
    return i57a9b05210a746928010661e9c18014f8e235ff52afa313241d78c4d64f35435.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
