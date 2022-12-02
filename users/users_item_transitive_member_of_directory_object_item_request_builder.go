package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
type UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters the groups, including nested groups, and directory roles that a user is a member of. Nullable.
type UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) Application()(*UsersItemTransitiveMemberOfItemApplicationRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    m := &UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) Device()(*UsersItemTransitiveMemberOfItemDeviceRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups, including nested groups, and directory roles that a user is a member of. Nullable.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) Group()(*UsersItemTransitiveMemberOfItemGroupRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*UsersItemTransitiveMemberOfItemOrgContactRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*UsersItemTransitiveMemberOfItemServicePrincipalRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *UsersItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) User()(*UsersItemTransitiveMemberOfItemUserRequestBuilder) {
    return NewUsersItemTransitiveMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
