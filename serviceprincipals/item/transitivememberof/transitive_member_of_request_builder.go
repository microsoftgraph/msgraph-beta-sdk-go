package transitivememberof

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i57e43716024bf990f9ee9ecb520d568c37cee1e0673714998defffd729bd51ab "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/device"
    i682bd07732f36a27ac56266aa5d657c4bfb49cd5346ae7ca8188619c9beeb4b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/count"
    i819d83578f04968066c7aaca89dd7a8b1c166a9040a4a9a047ec27875ea5becd "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/serviceprincipal"
    i82cab2da32bd101420472df5a91a014d0cfabf5a10c027ed51c89229cb98f32c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/application"
    ia55ef1d25c26b74fdf2c93357c94cd35648d36e3755356daaf0d57d494fd9045 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/user"
    iaf435cf82173ab3191667a29f4ac658d23d4c775b9517619e238eb52fb1745a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/group"
    ib3af12524f1b4582d58ac4231ac0f402d73bc114f42f927e244b90f56acb92fd "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/orgcontact"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.servicePrincipal entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters get transitiveMemberOf from servicePrincipals
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
// Application the application property
func (m *TransitiveMemberOfRequestBuilder) Application()(*i82cab2da32bd101420472df5a91a014d0cfabf5a10c027ed51c89229cb98f32c.ApplicationRequestBuilder) {
    return i82cab2da32bd101420472df5a91a014d0cfabf5a10c027ed51c89229cb98f32c.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/transitiveMemberOf{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
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
// Count the Count property
func (m *TransitiveMemberOfRequestBuilder) Count()(*i682bd07732f36a27ac56266aa5d657c4bfb49cd5346ae7ca8188619c9beeb4b7.CountRequestBuilder) {
    return i682bd07732f36a27ac56266aa5d657c4bfb49cd5346ae7ca8188619c9beeb4b7.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get transitiveMemberOf from servicePrincipals
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from servicePrincipals
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Device()(*i57e43716024bf990f9ee9ecb520d568c37cee1e0673714998defffd729bd51ab.DeviceRequestBuilder) {
    return i57e43716024bf990f9ee9ecb520d568c37cee1e0673714998defffd729bd51ab.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from servicePrincipals
func (m *TransitiveMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *TransitiveMemberOfRequestBuilder) Group()(*iaf435cf82173ab3191667a29f4ac658d23d4c775b9517619e238eb52fb1745a6.GroupRequestBuilder) {
    return iaf435cf82173ab3191667a29f4ac658d23d4c775b9517619e238eb52fb1745a6.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*ib3af12524f1b4582d58ac4231ac0f402d73bc114f42f927e244b90f56acb92fd.OrgContactRequestBuilder) {
    return ib3af12524f1b4582d58ac4231ac0f402d73bc114f42f927e244b90f56acb92fd.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*i819d83578f04968066c7aaca89dd7a8b1c166a9040a4a9a047ec27875ea5becd.ServicePrincipalRequestBuilder) {
    return i819d83578f04968066c7aaca89dd7a8b1c166a9040a4a9a047ec27875ea5becd.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMemberOfRequestBuilder) User()(*ia55ef1d25c26b74fdf2c93357c94cd35648d36e3755356daaf0d57d494fd9045.UserRequestBuilder) {
    return ia55ef1d25c26b74fdf2c93357c94cd35648d36e3755356daaf0d57d494fd9045.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
