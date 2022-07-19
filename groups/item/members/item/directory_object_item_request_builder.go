package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0554f116bfdffd94f642eff19d8111d848187556d2639aad20e8334f50ac410f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/ref"
    i302e76b03c5241f23781b72f91f61efdd300ecb011149443ac283284ca7762c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/serviceprincipal"
    i3196aa00a1ce299bcf796c3884e61126880c03f075bff5162d128f7c6d55c13e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/application"
    i3a337ddddf2d8793c348227a5f8c67e1edfce989ff9529d9a004942c45d0e5e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/orgcontact"
    i546e1ddc658614fd4779ae8e0fbbd2aa1d73c34462a7c7957c8a9d7544d0d046 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/group"
    i891dd04cc5613b4818d03f79193e6083c7b6213f8cd3bb83ad2a182b520cf403 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/device"
    i9b5247cc064ac6f25e02ad51c84a760a6cdd8864f16894568377f03c71d7cef2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/members/item/user"
)

// DirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\members\{directoryObject-id}
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*i3196aa00a1ce299bcf796c3884e61126880c03f075bff5162d128f7c6d55c13e.ApplicationRequestBuilder) {
    return i3196aa00a1ce299bcf796c3884e61126880c03f075bff5162d128f7c6d55c13e.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}";
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*i891dd04cc5613b4818d03f79193e6083c7b6213f8cd3bb83ad2a182b520cf403.DeviceRequestBuilder) {
    return i891dd04cc5613b4818d03f79193e6083c7b6213f8cd3bb83ad2a182b520cf403.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*i546e1ddc658614fd4779ae8e0fbbd2aa1d73c34462a7c7957c8a9d7544d0d046.GroupRequestBuilder) {
    return i546e1ddc658614fd4779ae8e0fbbd2aa1d73c34462a7c7957c8a9d7544d0d046.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i3a337ddddf2d8793c348227a5f8c67e1edfce989ff9529d9a004942c45d0e5e1.OrgContactRequestBuilder) {
    return i3a337ddddf2d8793c348227a5f8c67e1edfce989ff9529d9a004942c45d0e5e1.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *DirectoryObjectItemRequestBuilder) Ref()(*i0554f116bfdffd94f642eff19d8111d848187556d2639aad20e8334f50ac410f.RefRequestBuilder) {
    return i0554f116bfdffd94f642eff19d8111d848187556d2639aad20e8334f50ac410f.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i302e76b03c5241f23781b72f91f61efdd300ecb011149443ac283284ca7762c8.ServicePrincipalRequestBuilder) {
    return i302e76b03c5241f23781b72f91f61efdd300ecb011149443ac283284ca7762c8.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i9b5247cc064ac6f25e02ad51c84a760a6cdd8864f16894568377f03c71d7cef2.UserRequestBuilder) {
    return i9b5247cc064ac6f25e02ad51c84a760a6cdd8864f16894568377f03c71d7cef2.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
