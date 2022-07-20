package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i07da294fbf4713215f8826177f1ec62eea48f5c448f6eec0a08af160cac85cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/ref"
    i5eb970599e80bdff8722a653ae7be2698466a3fddfa10c27e566939ab3a8f568 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/endpoint"
    iba5aa965134efe758e18c13b781bb4b09b79320015e1e3b534796c2d51d2a51d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user"
    id9fe0fde601403f08ef44c0eef01e2da302467c7f60e87170098e94b10d84087 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/serviceprincipal"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\passwordlessMicrosoftAuthenticatorMethods\{passwordlessMicrosoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint the endpoint property
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i5eb970599e80bdff8722a653ae7be2698466a3fddfa10c27e566939ab3a8f568.EndpointRequestBuilder) {
    return i5eb970599e80bdff8722a653ae7be2698466a3fddfa10c27e566939ab3a8f568.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i07da294fbf4713215f8826177f1ec62eea48f5c448f6eec0a08af160cac85cf7.RefRequestBuilder) {
    return i07da294fbf4713215f8826177f1ec62eea48f5c448f6eec0a08af160cac85cf7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*id9fe0fde601403f08ef44c0eef01e2da302467c7f60e87170098e94b10d84087.ServicePrincipalRequestBuilder) {
    return id9fe0fde601403f08ef44c0eef01e2da302467c7f60e87170098e94b10d84087.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*iba5aa965134efe758e18c13b781bb4b09b79320015e1e3b534796c2d51d2a51d.UserRequestBuilder) {
    return iba5aa965134efe758e18c13b781bb4b09b79320015e1e3b534796c2d51d2a51d.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
