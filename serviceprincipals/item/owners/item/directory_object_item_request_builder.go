package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i01cef5f2f4b78439c7071b511721966c9de680d40b4a4c9bbe65ce1f27f3c2c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/user"
    i0f1ec117611c54586be42e3afa734eac16f6d478bdfd25f16d2f2e88fd48ae28 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal"
    i8231db2f5f118a4e6a3bf8329ebc608ebb433814568298ec5d259f82ee22dec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/endpoint"
    i91f7323cc85f18c0b8a3f86072de0279123cda3d8a9dcaa4732ca1bd4c673e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/ref"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\owners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/owners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i8231db2f5f118a4e6a3bf8329ebc608ebb433814568298ec5d259f82ee22dec7.EndpointRequestBuilder) {
    return i8231db2f5f118a4e6a3bf8329ebc608ebb433814568298ec5d259f82ee22dec7.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i91f7323cc85f18c0b8a3f86072de0279123cda3d8a9dcaa4732ca1bd4c673e54.RefRequestBuilder) {
    return i91f7323cc85f18c0b8a3f86072de0279123cda3d8a9dcaa4732ca1bd4c673e54.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i0f1ec117611c54586be42e3afa734eac16f6d478bdfd25f16d2f2e88fd48ae28.ServicePrincipalRequestBuilder) {
    return i0f1ec117611c54586be42e3afa734eac16f6d478bdfd25f16d2f2e88fd48ae28.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i01cef5f2f4b78439c7071b511721966c9de680d40b4a4c9bbe65ce1f27f3c2c9.UserRequestBuilder) {
    return i01cef5f2f4b78439c7071b511721966c9de680d40b4a4c9bbe65ce1f27f3c2c9.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
