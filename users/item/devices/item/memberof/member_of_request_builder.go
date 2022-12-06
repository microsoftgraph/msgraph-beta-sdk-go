package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b2cb37660bc8ad793763719a09c0bee7908fc3b0f5f967d62dd7c3d8cde15c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/device"
    i157bc45bacb8e3a5e70316fd832ddc0902e14c272beec5cdac2988bae7d9ac71 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/count"
    i39ee332407b41e1f400a826c8f6b6ff08646e8c16fb58f84d2837a3835935f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/group"
    i3ef1e6c2fb10cd36d7a76d676918b60e6681b179e31ccf43c6865225ecd53216 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/serviceprincipal"
    i6d5666bebf3222e9735a52e21d1f585ab008411d483406f0672f97b9b8272985 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/user"
    ib7f7e9c113af9bf91ec6a29ac1da649de8294869b0dbbdb6e5d2a51584ad3109 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/application"
    id0a38355692932b2468469bbaf1232d56bc130e23e19e459e7f25289624e17d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/devices/item/memberof/orgcontact"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type MemberOfRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MemberOfRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *MemberOfRequestBuilder) Application()(*ib7f7e9c113af9bf91ec6a29ac1da649de8294869b0dbbdb6e5d2a51584ad3109.ApplicationRequestBuilder) {
    return ib7f7e9c113af9bf91ec6a29ac1da649de8294869b0dbbdb6e5d2a51584ad3109.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MemberOfRequestBuilder) Count()(*i157bc45bacb8e3a5e70316fd832ddc0902e14c272beec5cdac2988bae7d9ac71.CountRequestBuilder) {
    return i157bc45bacb8e3a5e70316fd832ddc0902e14c272beec5cdac2988bae7d9ac71.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device casts the previous resource to device.
func (m *MemberOfRequestBuilder) Device()(*i0b2cb37660bc8ad793763719a09c0bee7908fc3b0f5f967d62dd7c3d8cde15c2.DeviceRequestBuilder) {
    return i0b2cb37660bc8ad793763719a09c0bee7908fc3b0f5f967d62dd7c3d8cde15c2.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *MemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *MemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group casts the previous resource to group.
func (m *MemberOfRequestBuilder) Group()(*i39ee332407b41e1f400a826c8f6b6ff08646e8c16fb58f84d2837a3835935f57.GroupRequestBuilder) {
    return i39ee332407b41e1f400a826c8f6b6ff08646e8c16fb58f84d2837a3835935f57.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *MemberOfRequestBuilder) OrgContact()(*id0a38355692932b2468469bbaf1232d56bc130e23e19e459e7f25289624e17d5.OrgContactRequestBuilder) {
    return id0a38355692932b2468469bbaf1232d56bc130e23e19e459e7f25289624e17d5.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i3ef1e6c2fb10cd36d7a76d676918b60e6681b179e31ccf43c6865225ecd53216.ServicePrincipalRequestBuilder) {
    return i3ef1e6c2fb10cd36d7a76d676918b60e6681b179e31ccf43c6865225ecd53216.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MemberOfRequestBuilder) User()(*i6d5666bebf3222e9735a52e21d1f585ab008411d483406f0672f97b9b8272985.UserRequestBuilder) {
    return i6d5666bebf3222e9735a52e21d1f585ab008411d483406f0672f97b9b8272985.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
