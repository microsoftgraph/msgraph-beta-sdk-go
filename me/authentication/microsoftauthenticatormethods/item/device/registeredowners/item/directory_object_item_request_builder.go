package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7a1cd3b19f067736f521155eaf51ca634eee7ac440186c255aecded54f4bd449 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/endpoint"
    ia1e69e347bd8c58f25eee4d465974bc0ce98d0545be94ef49ff97bf59b9b51c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/user"
    ic49018ade1090b1e65002d894fe4b1f7ada03431e581b2402bbf79c9e768812c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/serviceprincipal"
    id128f6478918749829b25fa2d36e4ad7839c33b29b33d9cb6f60b47b9d4e3e92 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/microsoftauthenticatormethods/item/device/registeredowners/item/ref"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\microsoftAuthenticatorMethods\{microsoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
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
// Endpoint casts the previous resource to endpoint.
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i7a1cd3b19f067736f521155eaf51ca634eee7ac440186c255aecded54f4bd449.EndpointRequestBuilder) {
    return i7a1cd3b19f067736f521155eaf51ca634eee7ac440186c255aecded54f4bd449.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*id128f6478918749829b25fa2d36e4ad7839c33b29b33d9cb6f60b47b9d4e3e92.RefRequestBuilder) {
    return id128f6478918749829b25fa2d36e4ad7839c33b29b33d9cb6f60b47b9d4e3e92.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*ic49018ade1090b1e65002d894fe4b1f7ada03431e581b2402bbf79c9e768812c.ServicePrincipalRequestBuilder) {
    return ic49018ade1090b1e65002d894fe4b1f7ada03431e581b2402bbf79c9e768812c.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*ia1e69e347bd8c58f25eee4d465974bc0ce98d0545be94ef49ff97bf59b9b51c4.UserRequestBuilder) {
    return ia1e69e347bd8c58f25eee4d465974bc0ce98d0545be94ef49ff97bf59b9b51c4.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
