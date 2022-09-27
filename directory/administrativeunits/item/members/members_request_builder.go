package members

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i043de7b1d022e021c4c76f9d76cab9f6b52cf472eede1ead8fac91b02eb27252 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/count"
    i237c599941fcfc47c272393e487ce3da971cb97acb8e1edfd142c83a9362cdee "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/group"
    i5ef6e20bd71d7b7f0f0b05521f85df4bf1379d65c99116ccef8920b4dd82c24f "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/serviceprincipal"
    i9968000560f9a77e3fbf4fa2a9e867a20a34063ee1b379280cee392254dee755 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/device"
    ic9f1a581b661445df39a48cb03d3d1dfd0d21c385ebb2bcd14d0e94952ed62a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/ref"
    id6dcfdab2cc2d290e94894b6092459bd5b1c63af624633ac3c70b29c86ab2302 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/orgcontact"
    ie245bba85841115e68099b3d2668341a8068bc2e63d9bbc39aeb2659a7dd30cf "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/application"
    if200fdd990abb8a354e4c02da185a06eaa8e3b958e93837240929b00ef5c842a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/administrativeunits/item/members/user"
)

// MembersRequestBuilder provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
type MembersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MembersRequestBuilderGetQueryParameters users and groups that are members of this administrative unit. Supports $expand.
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
// Application the application property
func (m *MembersRequestBuilder) Application()(*ie245bba85841115e68099b3d2668341a8068bc2e63d9bbc39aeb2659a7dd30cf.ApplicationRequestBuilder) {
    return ie245bba85841115e68099b3d2668341a8068bc2e63d9bbc39aeb2659a7dd30cf.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMembersRequestBuilderInternal instantiates a new MembersRequestBuilder and sets the default values.
func NewMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MembersRequestBuilder) {
    m := &MembersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *MembersRequestBuilder) Count()(*i043de7b1d022e021c4c76f9d76cab9f6b52cf472eede1ead8fac91b02eb27252.CountRequestBuilder) {
    return i043de7b1d022e021c4c76f9d76cab9f6b52cf472eede1ead8fac91b02eb27252.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Device the device property
func (m *MembersRequestBuilder) Device()(*i9968000560f9a77e3fbf4fa2a9e867a20a34063ee1b379280cee392254dee755.DeviceRequestBuilder) {
    return i9968000560f9a77e3fbf4fa2a9e867a20a34063ee1b379280cee392254dee755.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get users and groups that are members of this administrative unit. Supports $expand.
func (m *MembersRequestBuilder) Get(ctx context.Context, requestConfiguration *MembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Group the group property
func (m *MembersRequestBuilder) Group()(*i237c599941fcfc47c272393e487ce3da971cb97acb8e1edfd142c83a9362cdee.GroupRequestBuilder) {
    return i237c599941fcfc47c272393e487ce3da971cb97acb8e1edfd142c83a9362cdee.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *MembersRequestBuilder) OrgContact()(*id6dcfdab2cc2d290e94894b6092459bd5b1c63af624633ac3c70b29c86ab2302.OrgContactRequestBuilder) {
    return id6dcfdab2cc2d290e94894b6092459bd5b1c63af624633ac3c70b29c86ab2302.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the Ref property
func (m *MembersRequestBuilder) Ref()(*ic9f1a581b661445df39a48cb03d3d1dfd0d21c385ebb2bcd14d0e94952ed62a8.RefRequestBuilder) {
    return ic9f1a581b661445df39a48cb03d3d1dfd0d21c385ebb2bcd14d0e94952ed62a8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *MembersRequestBuilder) ServicePrincipal()(*i5ef6e20bd71d7b7f0f0b05521f85df4bf1379d65c99116ccef8920b4dd82c24f.ServicePrincipalRequestBuilder) {
    return i5ef6e20bd71d7b7f0f0b05521f85df4bf1379d65c99116ccef8920b4dd82c24f.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *MembersRequestBuilder) User()(*if200fdd990abb8a354e4c02da185a06eaa8e3b958e93837240929b00ef5c842a.UserRequestBuilder) {
    return if200fdd990abb8a354e4c02da185a06eaa8e3b958e93837240929b00ef5c842a.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
