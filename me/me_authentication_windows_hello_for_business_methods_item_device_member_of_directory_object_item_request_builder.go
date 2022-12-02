package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Application()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemApplicationRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    m := &MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Device()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemDeviceRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
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
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) Group()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemGroupRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemOrgContactRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfDirectoryObjectItemRequestBuilder) User()(*MeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemUserRequestBuilder) {
    return NewMeAuthenticationWindowsHelloForBusinessMethodsItemDeviceMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
