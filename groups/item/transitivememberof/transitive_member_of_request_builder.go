package transitivememberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1e6a6253c1d919a7804a310a9d5907ac3ab20e1b0f86b9349a61ab070fd65796 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/group"
    i332c6703bae87a933590c2e25cfca7b2af589aba0f159856e8c7262c31ed40c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/application"
    i392cebce218ea1c98676c9f0365f098dbe92f66150a6f8c0dfb7d2b8ec1150c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/serviceprincipal"
    i870a6ce3d01b361e259cf4ac3c7bdf5c5766b815800797cdf2a667aac697a6ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/device"
    i99da3d3b91f3dc375795ec5fa548dda2f955cf257ee9cfde928ee1418395445b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/user"
    if5eea3af17d8b2e389563487505984dcbcb17f59563628b1dfdb327440013f7e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/orgcontact"
    ifc0c576c85f182a8c24de4d78ba2a4aaa8336be5a9e7f5e0ec0889151419c3af "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/transitivememberof/count"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters the groups that a group is a member of, either directly and through nested membership. Nullable.
type TransitiveMemberOfRequestBuilderGetQueryParameters struct {
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
// TransitiveMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransitiveMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TransitiveMemberOfRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *TransitiveMemberOfRequestBuilder) Application()(*i332c6703bae87a933590c2e25cfca7b2af589aba0f159856e8c7262c31ed40c1.ApplicationRequestBuilder) {
    return i332c6703bae87a933590c2e25cfca7b2af589aba0f159856e8c7262c31ed40c1.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/transitiveMemberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTransitiveMemberOfRequestBuilder instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTransitiveMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *TransitiveMemberOfRequestBuilder) Count()(*ifc0c576c85f182a8c24de4d78ba2a4aaa8336be5a9e7f5e0ec0889151419c3af.CountRequestBuilder) {
    return ifc0c576c85f182a8c24de4d78ba2a4aaa8336be5a9e7f5e0ec0889151419c3af.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the groups that a group is a member of, either directly and through nested membership. Nullable.
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Device()(*i870a6ce3d01b361e259cf4ac3c7bdf5c5766b815800797cdf2a667aac697a6ef.DeviceRequestBuilder) {
    return i870a6ce3d01b361e259cf4ac3c7bdf5c5766b815800797cdf2a667aac697a6ef.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups that a group is a member of, either directly and through nested membership. Nullable.
func (m *TransitiveMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Group()(*i1e6a6253c1d919a7804a310a9d5907ac3ab20e1b0f86b9349a61ab070fd65796.GroupRequestBuilder) {
    return i1e6a6253c1d919a7804a310a9d5907ac3ab20e1b0f86b9349a61ab070fd65796.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*if5eea3af17d8b2e389563487505984dcbcb17f59563628b1dfdb327440013f7e.OrgContactRequestBuilder) {
    return if5eea3af17d8b2e389563487505984dcbcb17f59563628b1dfdb327440013f7e.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*i392cebce218ea1c98676c9f0365f098dbe92f66150a6f8c0dfb7d2b8ec1150c4.ServicePrincipalRequestBuilder) {
    return i392cebce218ea1c98676c9f0365f098dbe92f66150a6f8c0dfb7d2b8ec1150c4.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *TransitiveMemberOfRequestBuilder) User()(*i99da3d3b91f3dc375795ec5fa548dda2f955cf257ee9cfde928ee1418395445b.UserRequestBuilder) {
    return i99da3d3b91f3dc375795ec5fa548dda2f955cf257ee9cfde928ee1418395445b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
