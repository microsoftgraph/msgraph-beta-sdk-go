package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17bfa9a09af73dc5be89cb286359ba5ea12ceba4f7686e3ea0e2b622552397bc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/user"
    i2425d81dfbbe4494e094a557d850e84c599c20f280a5a6654e1f4ca4f522adef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/ref"
    i5bbdbad66e153e1cf11be4e2b0904709ec613c144843328cbae48a7d9ff08ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/serviceprincipal"
    ia5523f870a9ce09060dface3310d0548d54506e8c1edcbb0dea27595c991bfee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item/device/registeredowners/item/endpoint"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\windowsHelloForBusinessMethods\{windowsHelloForBusinessAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*ia5523f870a9ce09060dface3310d0548d54506e8c1edcbb0dea27595c991bfee.EndpointRequestBuilder) {
    return ia5523f870a9ce09060dface3310d0548d54506e8c1edcbb0dea27595c991bfee.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i2425d81dfbbe4494e094a557d850e84c599c20f280a5a6654e1f4ca4f522adef.RefRequestBuilder) {
    return i2425d81dfbbe4494e094a557d850e84c599c20f280a5a6654e1f4ca4f522adef.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i5bbdbad66e153e1cf11be4e2b0904709ec613c144843328cbae48a7d9ff08ee3.ServicePrincipalRequestBuilder) {
    return i5bbdbad66e153e1cf11be4e2b0904709ec613c144843328cbae48a7d9ff08ee3.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*i17bfa9a09af73dc5be89cb286359ba5ea12ceba4f7686e3ea0e2b622552397bc.UserRequestBuilder) {
    return i17bfa9a09af73dc5be89cb286359ba5ea12ceba4f7686e3ea0e2b622552397bc.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
