package transitivememberof

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1870c3dfd2fdbfbc7712af600ccc47fd24dd2efa681364e3fa8dd2ca286d2b9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/device"
    i1dac8c4fe0461c7469025e3f54442a35c26f21fa28f5099658ef8caa006e654c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/count"
    i841ccb6a6785b98c9cbab93bbeaaad250c28a17363b3886a4a107388dd10e3d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/user"
    ibc33f635d541cd472bf897b5cdd365326f3992f998b28f6de014a7dba294adb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/group"
    ic3abd355a87e72d89f12d456a1bd572723919d05ac42c20bca153dbd0b8edb22 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/application"
    id684ed14cc70822a890140c16460630a0b6ae956225e64b69d27ea59c7808b96 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/serviceprincipal"
    if817e641883e207c82efe1a7f9d195036fada4d905d96f55a0f34ec1c561375d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/transitivememberof/orgcontact"
)

// TransitiveMemberOfRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.user entity.
type TransitiveMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TransitiveMemberOfRequestBuilderGetQueryParameters get transitiveMemberOf from users
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
func (m *TransitiveMemberOfRequestBuilder) Application()(*ic3abd355a87e72d89f12d456a1bd572723919d05ac42c20bca153dbd0b8edb22.ApplicationRequestBuilder) {
    return ic3abd355a87e72d89f12d456a1bd572723919d05ac42c20bca153dbd0b8edb22.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTransitiveMemberOfRequestBuilderInternal instantiates a new TransitiveMemberOfRequestBuilder and sets the default values.
func NewTransitiveMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TransitiveMemberOfRequestBuilder) {
    m := &TransitiveMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/transitiveMemberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
// Count the count property
func (m *TransitiveMemberOfRequestBuilder) Count()(*i1dac8c4fe0461c7469025e3f54442a35c26f21fa28f5099658ef8caa006e654c.CountRequestBuilder) {
    return i1dac8c4fe0461c7469025e3f54442a35c26f21fa28f5099658ef8caa006e654c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get transitiveMemberOf from users
func (m *TransitiveMemberOfRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get transitiveMemberOf from users
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
func (m *TransitiveMemberOfRequestBuilder) Device()(*i1870c3dfd2fdbfbc7712af600ccc47fd24dd2efa681364e3fa8dd2ca286d2b9a.DeviceRequestBuilder) {
    return i1870c3dfd2fdbfbc7712af600ccc47fd24dd2efa681364e3fa8dd2ca286d2b9a.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get transitiveMemberOf from users
func (m *TransitiveMemberOfRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get transitiveMemberOf from users
func (m *TransitiveMemberOfRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TransitiveMemberOfRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group the group property
func (m *TransitiveMemberOfRequestBuilder) Group()(*ibc33f635d541cd472bf897b5cdd365326f3992f998b28f6de014a7dba294adb3.GroupRequestBuilder) {
    return ibc33f635d541cd472bf897b5cdd365326f3992f998b28f6de014a7dba294adb3.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *TransitiveMemberOfRequestBuilder) OrgContact()(*if817e641883e207c82efe1a7f9d195036fada4d905d96f55a0f34ec1c561375d.OrgContactRequestBuilder) {
    return if817e641883e207c82efe1a7f9d195036fada4d905d96f55a0f34ec1c561375d.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *TransitiveMemberOfRequestBuilder) ServicePrincipal()(*id684ed14cc70822a890140c16460630a0b6ae956225e64b69d27ea59c7808b96.ServicePrincipalRequestBuilder) {
    return id684ed14cc70822a890140c16460630a0b6ae956225e64b69d27ea59c7808b96.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *TransitiveMemberOfRequestBuilder) User()(*i841ccb6a6785b98c9cbab93bbeaaad250c28a17363b3886a4a107388dd10e3d5.UserRequestBuilder) {
    return i841ccb6a6785b98c9cbab93bbeaaad250c28a17363b3886a4a107388dd10e3d5.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
