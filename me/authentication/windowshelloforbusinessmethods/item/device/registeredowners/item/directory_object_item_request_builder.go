package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4183ae1d532e3e427be8be681e57e19e32843a273e06e93073fc6acdf48e65b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/serviceprincipal"
    i6b67245ce2f0f56ab16a0fd5bad2b90b74d1cf41b2a0783ed8f309f34d7c8933 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/endpoint"
    ic50e3695ae71e0dc30677cefd1d9b391fb8f0ec7c051a05ddd897a89b856d87c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/ref"
    idf3d502bfe1d1b6a1e1ae9dfcce9fdfe4edeeb8f18c4276e139a648b06705dbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\windowsHelloForBusinessMethods\{windowsHelloForBusinessAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i6b67245ce2f0f56ab16a0fd5bad2b90b74d1cf41b2a0783ed8f309f34d7c8933.EndpointRequestBuilder) {
    return i6b67245ce2f0f56ab16a0fd5bad2b90b74d1cf41b2a0783ed8f309f34d7c8933.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*ic50e3695ae71e0dc30677cefd1d9b391fb8f0ec7c051a05ddd897a89b856d87c.RefRequestBuilder) {
    return ic50e3695ae71e0dc30677cefd1d9b391fb8f0ec7c051a05ddd897a89b856d87c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i4183ae1d532e3e427be8be681e57e19e32843a273e06e93073fc6acdf48e65b7.ServicePrincipalRequestBuilder) {
    return i4183ae1d532e3e427be8be681e57e19e32843a273e06e93073fc6acdf48e65b7.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*idf3d502bfe1d1b6a1e1ae9dfcce9fdfe4edeeb8f18c4276e139a648b06705dbb.UserRequestBuilder) {
    return idf3d502bfe1d1b6a1e1ae9dfcce9fdfe4edeeb8f18c4276e139a648b06705dbb.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
