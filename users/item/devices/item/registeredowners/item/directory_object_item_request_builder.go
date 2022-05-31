package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i81690b3dc9d7e24db4b12a0b776f36e1dd0c19ed2fcade8c703d7b8ce1aa5157 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/user"
    i8476bf152666ee2bc0dd978deeffebb5424a1a20187e257d4fb56aec7ad14e78 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/ref"
    i90569534be322202b1875d3db27c3b988c0fe0b58f4d215450ad10619bcc6e8f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/endpoint"
    id597a991f65cd720682d2245fa11d32a3ce5d73d009939c2118a2f1e763ba610 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/registeredowners/item/serviceprincipal"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\devices\{device-id}\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i90569534be322202b1875d3db27c3b988c0fe0b58f4d215450ad10619bcc6e8f.EndpointRequestBuilder) {
    return i90569534be322202b1875d3db27c3b988c0fe0b58f4d215450ad10619bcc6e8f.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i8476bf152666ee2bc0dd978deeffebb5424a1a20187e257d4fb56aec7ad14e78.RefRequestBuilder) {
    return i8476bf152666ee2bc0dd978deeffebb5424a1a20187e257d4fb56aec7ad14e78.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*id597a991f65cd720682d2245fa11d32a3ce5d73d009939c2118a2f1e763ba610.ServicePrincipalRequestBuilder) {
    return id597a991f65cd720682d2245fa11d32a3ce5d73d009939c2118a2f1e763ba610.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i81690b3dc9d7e24db4b12a0b776f36e1dd0c19ed2fcade8c703d7b8ce1aa5157.UserRequestBuilder) {
    return i81690b3dc9d7e24db4b12a0b776f36e1dd0c19ed2fcade8c703d7b8ce1aa5157.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
