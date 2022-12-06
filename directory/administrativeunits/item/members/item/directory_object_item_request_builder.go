package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0109eb7ea97ec8d3ed5ad6014939bfbb3455c61c6131b6e96d65c56eebfccd52 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/ref"
    i20bb45ecbfcc0f537cf2ae02099b81a23741a5d5ecba52f2091af04773c41e72 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/application"
    i7bd8b700bed455d5c202602ab32552a36eb29af752cc266b0db6d3cc71ec58ed "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/group"
    i8a0e851b77f59af084270c73a0a351ce85b59c8d4d2f2e49931b11388023a544 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/user"
    ia614ffeeb88131f8fb813e727bfcdbd8d1396554732e0292f21f8df825273c66 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/serviceprincipal"
    ie5c2f0b491b3d3309853e4edc6149e9b647cf0b97a99e10f58e5a542e9645537 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/device"
    ied61f15217f23898025b9a311327e1de22bb040b5d1fe41ab02fca9862787ee4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/item/orgcontact"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \directory\administrativeUnits\{administrativeUnit-id}\members\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *DirectoryObjectItemRequestBuilder) Application()(*i20bb45ecbfcc0f537cf2ae02099b81a23741a5d5ecba52f2091af04773c41e72.ApplicationRequestBuilder) {
    return i20bb45ecbfcc0f537cf2ae02099b81a23741a5d5ecba52f2091af04773c41e72.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*ie5c2f0b491b3d3309853e4edc6149e9b647cf0b97a99e10f58e5a542e9645537.DeviceRequestBuilder) {
    return ie5c2f0b491b3d3309853e4edc6149e9b647cf0b97a99e10f58e5a542e9645537.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *DirectoryObjectItemRequestBuilder) Group()(*i7bd8b700bed455d5c202602ab32552a36eb29af752cc266b0db6d3cc71ec58ed.GroupRequestBuilder) {
    return i7bd8b700bed455d5c202602ab32552a36eb29af752cc266b0db6d3cc71ec58ed.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*ied61f15217f23898025b9a311327e1de22bb040b5d1fe41ab02fca9862787ee4.OrgContactRequestBuilder) {
    return ied61f15217f23898025b9a311327e1de22bb040b5d1fe41ab02fca9862787ee4.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directory entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i0109eb7ea97ec8d3ed5ad6014939bfbb3455c61c6131b6e96d65c56eebfccd52.RefRequestBuilder) {
    return i0109eb7ea97ec8d3ed5ad6014939bfbb3455c61c6131b6e96d65c56eebfccd52.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*ia614ffeeb88131f8fb813e727bfcdbd8d1396554732e0292f21f8df825273c66.ServicePrincipalRequestBuilder) {
    return ia614ffeeb88131f8fb813e727bfcdbd8d1396554732e0292f21f8df825273c66.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*i8a0e851b77f59af084270c73a0a351ce85b59c8d4d2f2e49931b11388023a544.UserRequestBuilder) {
    return i8a0e851b77f59af084270c73a0a351ce85b59c8d4d2f2e49931b11388023a544.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
