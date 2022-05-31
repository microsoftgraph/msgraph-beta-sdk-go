package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i1bc6fd9fe0dcd9ce0dea094ac423b92762d3cfa374e3c7500989547fbce67a63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/ref"
    i251c4f7718e17dc5e1de956c1c02555dd1fa4c55bb86b19344cd58f990586d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/endpoint"
    i785bb04f33f683dcdeac7371e0b1c9bd9f9564869f1fe7a000332923b8d53fd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/user"
    id6fecc76cf24735fde7570c1278053ed345211a6c0c39ec2c12c2bbc9a82e40f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/registeredowners/item/serviceprincipal"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\devices\{device-id}\registeredOwners\{directoryObject-id}
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
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Endpoint()(*i251c4f7718e17dc5e1de956c1c02555dd1fa4c55bb86b19344cd58f990586d42.EndpointRequestBuilder) {
    return i251c4f7718e17dc5e1de956c1c02555dd1fa4c55bb86b19344cd58f990586d42.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i1bc6fd9fe0dcd9ce0dea094ac423b92762d3cfa374e3c7500989547fbce67a63.RefRequestBuilder) {
    return i1bc6fd9fe0dcd9ce0dea094ac423b92762d3cfa374e3c7500989547fbce67a63.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*id6fecc76cf24735fde7570c1278053ed345211a6c0c39ec2c12c2bbc9a82e40f.ServicePrincipalRequestBuilder) {
    return id6fecc76cf24735fde7570c1278053ed345211a6c0c39ec2c12c2bbc9a82e40f.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i785bb04f33f683dcdeac7371e0b1c9bd9f9564869f1fe7a000332923b8d53fd4.UserRequestBuilder) {
    return i785bb04f33f683dcdeac7371e0b1c9bd9f9564869f1fe7a000332923b8d53fd4.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
