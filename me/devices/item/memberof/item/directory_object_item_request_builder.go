package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11b1de57dfddc58520c5148018c8b4c6db2425cae457f1f0b9c19bfc2b4c679c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/serviceprincipal"
    i210f27e3624a30c2228138de310dd05f6dc614441d5e90f13248da0e2bc80197 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/device"
    i3e1911ca9c1ffa4d0986c07b31900c3db9007dcc69d476eb1cbf8795dac3e6bc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/application"
    i5c741b93926a9e2328e66892e0d278560c681e12a9ddebc215556c6fe051a02d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/orgcontact"
    ib464cb525754ae1ea44bae50a75a58c8122fc4da582f2ffa9abab18df85c0fb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/group"
    ie4c91c81a60b80ce80b5e2b7d4cdfb20772af48e24249e078902cff10b3c84a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item/memberof/item/user"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.device entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
type DirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application the application property
func (m *DirectoryObjectItemRequestBuilder) Application()(*i3e1911ca9c1ffa4d0986c07b31900c3db9007dcc69d476eb1cbf8795dac3e6bc.ApplicationRequestBuilder) {
    return i3e1911ca9c1ffa4d0986c07b31900c3db9007dcc69d476eb1cbf8795dac3e6bc.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*i210f27e3624a30c2228138de310dd05f6dc614441d5e90f13248da0e2bc80197.DeviceRequestBuilder) {
    return i210f27e3624a30c2228138de310dd05f6dc614441d5e90f13248da0e2bc80197.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler groups and administrative units that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*ib464cb525754ae1ea44bae50a75a58c8122fc4da582f2ffa9abab18df85c0fb5.GroupRequestBuilder) {
    return ib464cb525754ae1ea44bae50a75a58c8122fc4da582f2ffa9abab18df85c0fb5.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i5c741b93926a9e2328e66892e0d278560c681e12a9ddebc215556c6fe051a02d.OrgContactRequestBuilder) {
    return i5c741b93926a9e2328e66892e0d278560c681e12a9ddebc215556c6fe051a02d.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i11b1de57dfddc58520c5148018c8b4c6db2425cae457f1f0b9c19bfc2b4c679c.ServicePrincipalRequestBuilder) {
    return i11b1de57dfddc58520c5148018c8b4c6db2425cae457f1f0b9c19bfc2b4c679c.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*ie4c91c81a60b80ce80b5e2b7d4cdfb20772af48e24249e078902cff10b3c84a2.UserRequestBuilder) {
    return ie4c91c81a60b80ce80b5e2b7d4cdfb20772af48e24249e078902cff10b3c84a2.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
