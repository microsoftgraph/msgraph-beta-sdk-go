package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i359350673f27e90b8f801922beab931ba54203022a3848484d68117d8d800b44 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/orgcontact"
    i38e8f0696c351e82d135d1edd7e8d853fa3c70bb00122f4509c117d56ea6c52a "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/ref"
    i8000cf82ba33728fdbed16e195e95de76b4717885ae31277e07f210f5372d3fe "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/group"
    ic23977302b8be3c90b4f3899a5edbd90c6b0c8f6f0c5e64148ae8420b28fa2dd "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/user"
    ic8cc83c5b3300ff0275700cd9528c22b6c31ecf03fef11d0530ee264cf20e818 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/device"
    ie6a6343540e9c0fb419a6a0ead1f913732ce1eea549cdc3a3f1bfccc01fe676c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/application"
    ifc226f8f8331c540ad5d6951d28ecdaeee717432e077ee5abd5bfda4e0efc1b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members/item/serviceprincipal"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \administrativeUnits\{administrativeUnit-id}\members\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*ie6a6343540e9c0fb419a6a0ead1f913732ce1eea549cdc3a3f1bfccc01fe676c.ApplicationRequestBuilder) {
    return ie6a6343540e9c0fb419a6a0ead1f913732ce1eea549cdc3a3f1bfccc01fe676c.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/administrativeUnits/{administrativeUnit%2Did}/members/{directoryObject%2Did}";
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
// Device the device property
func (m *DirectoryObjectItemRequestBuilder) Device()(*ic8cc83c5b3300ff0275700cd9528c22b6c31ecf03fef11d0530ee264cf20e818.DeviceRequestBuilder) {
    return ic8cc83c5b3300ff0275700cd9528c22b6c31ecf03fef11d0530ee264cf20e818.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*i8000cf82ba33728fdbed16e195e95de76b4717885ae31277e07f210f5372d3fe.GroupRequestBuilder) {
    return i8000cf82ba33728fdbed16e195e95de76b4717885ae31277e07f210f5372d3fe.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i359350673f27e90b8f801922beab931ba54203022a3848484d68117d8d800b44.OrgContactRequestBuilder) {
    return i359350673f27e90b8f801922beab931ba54203022a3848484d68117d8d800b44.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i38e8f0696c351e82d135d1edd7e8d853fa3c70bb00122f4509c117d56ea6c52a.RefRequestBuilder) {
    return i38e8f0696c351e82d135d1edd7e8d853fa3c70bb00122f4509c117d56ea6c52a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*ifc226f8f8331c540ad5d6951d28ecdaeee717432e077ee5abd5bfda4e0efc1b1.ServicePrincipalRequestBuilder) {
    return ifc226f8f8331c540ad5d6951d28ecdaeee717432e077ee5abd5bfda4e0efc1b1.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*ic23977302b8be3c90b4f3899a5edbd90c6b0c8f6f0c5e64148ae8420b28fa2dd.UserRequestBuilder) {
    return ic23977302b8be3c90b4f3899a5edbd90c6b0c8f6f0c5e64148ae8420b28fa2dd.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
