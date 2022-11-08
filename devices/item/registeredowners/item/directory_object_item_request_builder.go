package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i31fc884ee21d81a9d223ec52b81984d7188925ebbb9590234ce695b09ee13f61 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item/serviceprincipal"
    i79dc4a4ce843bc9198866ba27871b1c21e9ef5e6b587c1524733e139950286d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item/ref"
    ic26eb1a27e689bf52dd4aea829f9f226d4a8538244a8ea55b7b4b1e167122789 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item/user"
    ie291eedd432f291637a50529c7b824577ad8ab84ee8549e7dbd01b08a0f5776a "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item/endpoint"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*ie291eedd432f291637a50529c7b824577ad8ab84ee8549e7dbd01b08a0f5776a.EndpointRequestBuilder) {
    return ie291eedd432f291637a50529c7b824577ad8ab84ee8549e7dbd01b08a0f5776a.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of device entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i79dc4a4ce843bc9198866ba27871b1c21e9ef5e6b587c1524733e139950286d8.RefRequestBuilder) {
    return i79dc4a4ce843bc9198866ba27871b1c21e9ef5e6b587c1524733e139950286d8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i31fc884ee21d81a9d223ec52b81984d7188925ebbb9590234ce695b09ee13f61.ServicePrincipalRequestBuilder) {
    return i31fc884ee21d81a9d223ec52b81984d7188925ebbb9590234ce695b09ee13f61.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*ic26eb1a27e689bf52dd4aea829f9f226d4a8538244a8ea55b7b4b1e167122789.UserRequestBuilder) {
    return ic26eb1a27e689bf52dd4aea829f9f226d4a8538244a8ea55b7b4b1e167122789.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
