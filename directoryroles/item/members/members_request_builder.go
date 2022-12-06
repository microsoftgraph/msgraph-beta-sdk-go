package members

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1fc7a2a4ee4ef30eb8406d0471e417a8ef8c07d66f5e4a27ce8feca817684b57 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/user"
    i49e7796ef4796c80789cc964d33590055f77dda19544d468d69d69e93ad1dba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/application"
    i4cde313a997aacf46308327b831d6bb1b76c69f86e478d28fe47e0f1e39df596 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/count"
    i5649e937d3563442cbebc5cc78363bfeaf2cd68e060ccbd9723925e5bfb2034a "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/ref"
    i56c4df3f343dac1e4d5bf6ece09438661c3677baede6d32e66dc5fb12383c7b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/group"
    i74499d368d5be350cc0073c6e355c31e0b4c9a723e0bad327437826cf6565c3d "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/device"
    iabc5f531ded85ae394d072f775ee07cf5b0c76123cbb4a080d4b81f70a354ff0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/serviceprincipal"
    ic7a68d4746e9200eb13c01a7aee47225ee7d42c450b643f3c0237f19dbc36e0a "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/item/members/orgcontact"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.directoryRole entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersRequestBuilderGetQueryParameters users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
type MembersRequestBuilderGetQueryParameters struct {
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
// MembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MembersRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *MembersRequestBuilder) Application()(*i49e7796ef4796c80789cc964d33590055f77dda19544d468d69d69e93ad1dba5.ApplicationRequestBuilder) {
    return i49e7796ef4796c80789cc964d33590055f77dda19544d468d69d69e93ad1dba5.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMembersRequestBuilder instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MembersRequestBuilder) Count()(*i4cde313a997aacf46308327b831d6bb1b76c69f86e478d28fe47e0f1e39df596.CountRequestBuilder) {
    return i4cde313a997aacf46308327b831d6bb1b76c69f86e478d28fe47e0f1e39df596.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
func (m *MembersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MembersRequestBuilder) Device()(*i74499d368d5be350cc0073c6e355c31e0b4c9a723e0bad327437826cf6565c3d.DeviceRequestBuilder) {
    return i74499d368d5be350cc0073c6e355c31e0b4c9a723e0bad327437826cf6565c3d.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable. Supports $expand.
func (m *MembersRequestBuilder) Get(ctx context.Context, requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *MembersRequestBuilder) Group()(*i56c4df3f343dac1e4d5bf6ece09438661c3677baede6d32e66dc5fb12383c7b8.GroupRequestBuilder) {
    return i56c4df3f343dac1e4d5bf6ece09438661c3677baede6d32e66dc5fb12383c7b8.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *MembersRequestBuilder) OrgContact()(*ic7a68d4746e9200eb13c01a7aee47225ee7d42c450b643f3c0237f19dbc36e0a.OrgContactRequestBuilder) {
    return ic7a68d4746e9200eb13c01a7aee47225ee7d42c450b643f3c0237f19dbc36e0a.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directoryRole entities.
func (m *MembersRequestBuilder) Ref()(*i5649e937d3563442cbebc5cc78363bfeaf2cd68e060ccbd9723925e5bfb2034a.RefRequestBuilder) {
    return i5649e937d3563442cbebc5cc78363bfeaf2cd68e060ccbd9723925e5bfb2034a.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MembersRequestBuilder) ServicePrincipal()(*iabc5f531ded85ae394d072f775ee07cf5b0c76123cbb4a080d4b81f70a354ff0.ServicePrincipalRequestBuilder) {
    return iabc5f531ded85ae394d072f775ee07cf5b0c76123cbb4a080d4b81f70a354ff0.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MembersRequestBuilder) User()(*i1fc7a2a4ee4ef30eb8406d0471e417a8ef8c07d66f5e4a27ce8feca817684b57.UserRequestBuilder) {
    return i1fc7a2a4ee4ef30eb8406d0471e417a8ef8c07d66f5e4a27ce8feca817684b57.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
