package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0846bd4ebbc4a16b1efb085ceaac083d04901f26d9aaddca03547bd8d53e7bff "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/item/ref"
    i3bce1e328e6c573308e0c6d53f934fbd921b13dc4125c7df4fbe884dbcf6b1c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/item/endpoint"
    i3d1151b8bd439fb793d74884a019bf6e2dfc1b0db70392ef1dcf843f2403e0bc "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/item/serviceprincipal"
    i83ddad6113fb39d307b8caab7e9294d5f3cd96c6141632f4aaad54444d164b58 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/item/user"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\owners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/owners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i3bce1e328e6c573308e0c6d53f934fbd921b13dc4125c7df4fbe884dbcf6b1c2.EndpointRequestBuilder) {
    return i3bce1e328e6c573308e0c6d53f934fbd921b13dc4125c7df4fbe884dbcf6b1c2.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of application entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i0846bd4ebbc4a16b1efb085ceaac083d04901f26d9aaddca03547bd8d53e7bff.RefRequestBuilder) {
    return i0846bd4ebbc4a16b1efb085ceaac083d04901f26d9aaddca03547bd8d53e7bff.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i3d1151b8bd439fb793d74884a019bf6e2dfc1b0db70392ef1dcf843f2403e0bc.ServicePrincipalRequestBuilder) {
    return i3d1151b8bd439fb793d74884a019bf6e2dfc1b0db70392ef1dcf843f2403e0bc.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*i83ddad6113fb39d307b8caab7e9294d5f3cd96c6141632f4aaad54444d164b58.UserRequestBuilder) {
    return i83ddad6113fb39d307b8caab7e9294d5f3cd96c6141632f4aaad54444d164b58.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
