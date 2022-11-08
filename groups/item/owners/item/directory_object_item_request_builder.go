package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i7601bd0b917d6608f8df406e158abed23664e046c1686148079f13847da01518 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/group"
    i7e3f0928a979ee0571080bde1c0fbd80dfe479d75cc886b0ffad5d86f51031f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/ref"
    i95807e670d10184c232c5efde692dd4c0d37b143b25f239a20a26291ebeb438f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/serviceprincipal"
    ibc2070288865aa3ade1ced7adbb3977c68c8eb5140e18c05ccc1ef85455dec18 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/user"
    ic950a95da1cad019cce911a7110a242c225e5436ad2bfd560909922144cef6cf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/device"
    icfb500888979eddac94b6b9003f4cab6842c827b4d5fc9fd3bc5c6ab07368e35 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/orgcontact"
    if393380e19404b65cadde02bc08f8404b07e3964123ddea94d949e42fd27deb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/item/application"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\owners\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *DirectoryObjectItemRequestBuilder) Application()(*if393380e19404b65cadde02bc08f8404b07e3964123ddea94d949e42fd27deb1.ApplicationRequestBuilder) {
    return if393380e19404b65cadde02bc08f8404b07e3964123ddea94d949e42fd27deb1.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}";
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
// Device casts the previous resource to device.
func (m *DirectoryObjectItemRequestBuilder) Device()(*ic950a95da1cad019cce911a7110a242c225e5436ad2bfd560909922144cef6cf.DeviceRequestBuilder) {
    return ic950a95da1cad019cce911a7110a242c225e5436ad2bfd560909922144cef6cf.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *DirectoryObjectItemRequestBuilder) Group()(*i7601bd0b917d6608f8df406e158abed23664e046c1686148079f13847da01518.GroupRequestBuilder) {
    return i7601bd0b917d6608f8df406e158abed23664e046c1686148079f13847da01518.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*icfb500888979eddac94b6b9003f4cab6842c827b4d5fc9fd3bc5c6ab07368e35.OrgContactRequestBuilder) {
    return icfb500888979eddac94b6b9003f4cab6842c827b4d5fc9fd3bc5c6ab07368e35.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of group entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i7e3f0928a979ee0571080bde1c0fbd80dfe479d75cc886b0ffad5d86f51031f4.RefRequestBuilder) {
    return i7e3f0928a979ee0571080bde1c0fbd80dfe479d75cc886b0ffad5d86f51031f4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i95807e670d10184c232c5efde692dd4c0d37b143b25f239a20a26291ebeb438f.ServicePrincipalRequestBuilder) {
    return i95807e670d10184c232c5efde692dd4c0d37b143b25f239a20a26291ebeb438f.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*ibc2070288865aa3ade1ced7adbb3977c68c8eb5140e18c05ccc1ef85455dec18.UserRequestBuilder) {
    return ibc2070288865aa3ade1ced7adbb3977c68c8eb5140e18c05ccc1ef85455dec18.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
