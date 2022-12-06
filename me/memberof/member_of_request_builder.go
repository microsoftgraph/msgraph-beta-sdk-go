package memberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1cb40a66887bb35c5fd175cf151eaa4c1c7953df6ce1e1170700a5c0b5691ae5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/orgcontact"
    i1d5c132b29299359701073ab61dd22e264dfd64c4c32095b2073b159fcb4f2ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/device"
    i5d930866e68e92ea2c5bb29cd86ccee5e235ef0650c40b90af3e7a72578df986 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/serviceprincipal"
    i83784cd955b7d9802fa409865b08454105100b7fa8a9b5297196fb687eaf612f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/count"
    i9ebb526e7d7db68b75570f74fec13c8f2873bf44a2d6bddd8662c7f762d8b3ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/user"
    id47c84f55b7690b9d1d7c6b4856b98cae7283045760369958c28af41d8244f31 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/group"
    if1f4c2f4e530db035ebb1e4cbf2a2eff98abda4fdd5115da026d5b7de80f3470 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/application"
)

// MemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.user entity.
type MemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MemberOfRequestBuilderGetQueryParameters the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Application()(*if1f4c2f4e530db035ebb1e4cbf2a2eff98abda4fdd5115da026d5b7de80f3470.ApplicationRequestBuilder) {
    return if1f4c2f4e530db035ebb1e4cbf2a2eff98abda4fdd5115da026d5b7de80f3470.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MemberOfRequestBuilder) {
    m := &MemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
func (m *MemberOfRequestBuilder) Count()(*i83784cd955b7d9802fa409865b08454105100b7fa8a9b5297196fb687eaf612f.CountRequestBuilder) {
    return i83784cd955b7d9802fa409865b08454105100b7fa8a9b5297196fb687eaf612f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Device()(*i1d5c132b29299359701073ab61dd22e264dfd64c4c32095b2073b159fcb4f2ac.DeviceRequestBuilder) {
    return i1d5c132b29299359701073ab61dd22e264dfd64c4c32095b2073b159fcb4f2ac.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
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
func (m *MemberOfRequestBuilder) Group()(*id47c84f55b7690b9d1d7c6b4856b98cae7283045760369958c28af41d8244f31.GroupRequestBuilder) {
    return id47c84f55b7690b9d1d7c6b4856b98cae7283045760369958c28af41d8244f31.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *MemberOfRequestBuilder) OrgContact()(*i1cb40a66887bb35c5fd175cf151eaa4c1c7953df6ce1e1170700a5c0b5691ae5.OrgContactRequestBuilder) {
    return i1cb40a66887bb35c5fd175cf151eaa4c1c7953df6ce1e1170700a5c0b5691ae5.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MemberOfRequestBuilder) ServicePrincipal()(*i5d930866e68e92ea2c5bb29cd86ccee5e235ef0650c40b90af3e7a72578df986.ServicePrincipalRequestBuilder) {
    return i5d930866e68e92ea2c5bb29cd86ccee5e235ef0650c40b90af3e7a72578df986.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MemberOfRequestBuilder) User()(*i9ebb526e7d7db68b75570f74fec13c8f2873bf44a2d6bddd8662c7f762d8b3ad.UserRequestBuilder) {
    return i9ebb526e7d7db68b75570f74fec13c8f2873bf44a2d6bddd8662c7f762d8b3ad.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
