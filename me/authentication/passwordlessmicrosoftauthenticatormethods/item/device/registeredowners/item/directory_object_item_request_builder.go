package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0c4c71d995a32eef1a53c5c655d3472c01e4d43c33a4a7ba4373a7a7a129502c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/serviceprincipal"
    i42a9eb262894e60fd65f4851fe191e798a2538b8a9d3c5f65a1b7d05d6e77dcf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/endpoint"
    i60764cfd629997cfc22d214ddb56f24a488d58589a4a8d72c3c24cbbf02b9da1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/user"
    i8ee24ea35768bfc28d3b1b638395bc4c5dfdf443324904f5ec8f2e046138f6bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/passwordlessmicrosoftauthenticatormethods/item/device/registeredowners/item/ref"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\passwordlessMicrosoftAuthenticatorMethods\{passwordlessMicrosoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/me/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i42a9eb262894e60fd65f4851fe191e798a2538b8a9d3c5f65a1b7d05d6e77dcf.EndpointRequestBuilder) {
    return i42a9eb262894e60fd65f4851fe191e798a2538b8a9d3c5f65a1b7d05d6e77dcf.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i8ee24ea35768bfc28d3b1b638395bc4c5dfdf443324904f5ec8f2e046138f6bd.RefRequestBuilder) {
    return i8ee24ea35768bfc28d3b1b638395bc4c5dfdf443324904f5ec8f2e046138f6bd.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i0c4c71d995a32eef1a53c5c655d3472c01e4d43c33a4a7ba4373a7a7a129502c.ServicePrincipalRequestBuilder) {
    return i0c4c71d995a32eef1a53c5c655d3472c01e4d43c33a4a7ba4373a7a7a129502c.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i60764cfd629997cfc22d214ddb56f24a488d58589a4a8d72c3c24cbbf02b9da1.UserRequestBuilder) {
    return i60764cfd629997cfc22d214ddb56f24a488d58589a4a8d72c3c24cbbf02b9da1.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
