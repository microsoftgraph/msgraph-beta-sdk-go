package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i040e75d2eb3112423eff124ce0f5182742f83088260e7a351f2db45b32dfea40 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/ref"
    i068373a8a0f03f9e4ca9a83266f61e7efbdc5a925e98b47aa62b3f890d7527e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/application"
    i4e65a2a10a9449e109a2ea98cdb2339c0e201b6da2adef5782cc32452ab1b9ce "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/user"
    i550f6b65d4f6f2e9b676fcb8d3993c69a5a6ed8f19448a2b815b455e253330b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/serviceprincipal"
    i833c8c31764983f877e16da504b9481a228bd85f2e9c2d56842f2abc67424e73 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/device"
    ibe02a8266d720c1c6d895dd221fd5de46093b8981d35f339d15acf47f5ee1f33 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/group"
    ie9bef146436c0d53c4b8e8839c022155ffd1bca7ddc3d7cc2e288dc8ddaea6be "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/item/orgcontact"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \directoryRoles\{directoryRole-id}\members\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *DirectoryObjectItemRequestBuilder) Application()(*i068373a8a0f03f9e4ca9a83266f61e7efbdc5a925e98b47aa62b3f890d7527e9.ApplicationRequestBuilder) {
    return i068373a8a0f03f9e4ca9a83266f61e7efbdc5a925e98b47aa62b3f890d7527e9.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole%2Did}/members/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*i833c8c31764983f877e16da504b9481a228bd85f2e9c2d56842f2abc67424e73.DeviceRequestBuilder) {
    return i833c8c31764983f877e16da504b9481a228bd85f2e9c2d56842f2abc67424e73.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *DirectoryObjectItemRequestBuilder) Group()(*ibe02a8266d720c1c6d895dd221fd5de46093b8981d35f339d15acf47f5ee1f33.GroupRequestBuilder) {
    return ibe02a8266d720c1c6d895dd221fd5de46093b8981d35f339d15acf47f5ee1f33.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*ie9bef146436c0d53c4b8e8839c022155ffd1bca7ddc3d7cc2e288dc8ddaea6be.OrgContactRequestBuilder) {
    return ie9bef146436c0d53c4b8e8839c022155ffd1bca7ddc3d7cc2e288dc8ddaea6be.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directoryRole entities.
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i040e75d2eb3112423eff124ce0f5182742f83088260e7a351f2db45b32dfea40.RefRequestBuilder) {
    return i040e75d2eb3112423eff124ce0f5182742f83088260e7a351f2db45b32dfea40.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i550f6b65d4f6f2e9b676fcb8d3993c69a5a6ed8f19448a2b815b455e253330b4.ServicePrincipalRequestBuilder) {
    return i550f6b65d4f6f2e9b676fcb8d3993c69a5a6ed8f19448a2b815b455e253330b4.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryObjectItemRequestBuilder) User()(*i4e65a2a10a9449e109a2ea98cdb2339c0e201b6da2adef5782cc32452ab1b9ce.UserRequestBuilder) {
    return i4e65a2a10a9449e109a2ea98cdb2339c0e201b6da2adef5782cc32452ab1b9ce.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
