package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3598d127f59a59e72d7319edfd9a37c7c4a4b75eda903c14f30f782e0447f205 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/user"
    i4897a28bdd47243347a2462e6adbf9f8f0dba3efa8b2451b50719c8b24ff2738 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/orgcontact"
    i78afae362412664b40439484c29196b09255982e678c42651be7354dbb9a9ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/serviceprincipal"
    i8f4e0c675a212171840dd4fc15f1a5ab29ca3a6554aaa39aa61a0ad5ee75e6c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/application"
    ib399c02d7456f14f876d8aefe91910d2b0c269e4eb007880b14e0dadeb5f9a1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/device"
    iea602daa89c5a9af8d0cd3ad2f980db11d7ec1fb1b0de06f4a23bd7dababc172 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/memberof/item/group"
)

// DirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.user entity.
type DirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryObjectItemRequestBuilderGetQueryParameters the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
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
func (m *DirectoryObjectItemRequestBuilder) Application()(*i8f4e0c675a212171840dd4fc15f1a5ab29ca3a6554aaa39aa61a0ad5ee75e6c4.ApplicationRequestBuilder) {
    return i8f4e0c675a212171840dd4fc15f1a5ab29ca3a6554aaa39aa61a0ad5ee75e6c4.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryObjectItemRequestBuilder) {
    m := &DirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}{?%24select,%24expand}";
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
// CreateGetRequestInformation the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectoryObjectItemRequestBuilder) Device()(*ib399c02d7456f14f876d8aefe91910d2b0c269e4eb007880b14e0dadeb5f9a1e.DeviceRequestBuilder) {
    return ib399c02d7456f14f876d8aefe91910d2b0c269e4eb007880b14e0dadeb5f9a1e.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the groups, directory roles and administrative units that the user is a member of. Read-only. Nullable. Supports $expand.
func (m *DirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
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
// Group the group property
func (m *DirectoryObjectItemRequestBuilder) Group()(*iea602daa89c5a9af8d0cd3ad2f980db11d7ec1fb1b0de06f4a23bd7dababc172.GroupRequestBuilder) {
    return iea602daa89c5a9af8d0cd3ad2f980db11d7ec1fb1b0de06f4a23bd7dababc172.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *DirectoryObjectItemRequestBuilder) OrgContact()(*i4897a28bdd47243347a2462e6adbf9f8f0dba3efa8b2451b50719c8b24ff2738.OrgContactRequestBuilder) {
    return i4897a28bdd47243347a2462e6adbf9f8f0dba3efa8b2451b50719c8b24ff2738.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *DirectoryObjectItemRequestBuilder) ServicePrincipal()(*i78afae362412664b40439484c29196b09255982e678c42651be7354dbb9a9ce5.ServicePrincipalRequestBuilder) {
    return i78afae362412664b40439484c29196b09255982e678c42651be7354dbb9a9ce5.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *DirectoryObjectItemRequestBuilder) User()(*i3598d127f59a59e72d7319edfd9a37c7c4a4b75eda903c14f30f782e0447f205.UserRequestBuilder) {
    return i3598d127f59a59e72d7319edfd9a37c7c4a4b75eda903c14f30f782e0447f205.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
