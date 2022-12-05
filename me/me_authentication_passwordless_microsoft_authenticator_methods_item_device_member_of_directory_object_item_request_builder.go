package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Application()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    m := &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device casts the previous resource to device.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Device()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemDeviceRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Group()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemGroupRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemOrgContactRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) User()(*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemUserRequestBuilder) {
    return NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
