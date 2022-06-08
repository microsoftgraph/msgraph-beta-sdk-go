package owners

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1d617c3562064f6580fd55ae5e4013da163a0f961bf9b5f9fccc75f94b3f24e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/user"
    i30c6a904f09f7ba0c29704a7a72037781c52d7a49cd1e9f954d0e321ef034e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/serviceprincipal"
    i327eef57a0e308487ea882dc08c72a66bfd918bbaff92b23558cb16e66dd1ba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/ref"
    i5482c610da612e49b7b76e839398da0032b32da37a31641dd6c3fe2d5b66130c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/count"
    i5ac3da2b72c860c9d858ba8884f03ac263d08f62079ac3f62f3df69d20da6eb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/device"
    i9ffbfdaeeea3bb4efea4ad0e6db2b4ee4e6e0f52f5325b168da679a18859e18e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/orgcontact"
    ie46d71cd800842fcb2a698e0f4806ad7a00318a0cf8a6b36096801a9ab93016c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/group"
    iea108107f244d15625e030edb7372166349194daf3bc6cd117f48446d04e57f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/owners/application"
)

// OwnersRequestBuilder provides operations to manage the owners property of the microsoft.graph.group entity.
type OwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnersRequestBuilderGetQueryParameters the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
type OwnersRequestBuilderGetQueryParameters struct {
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
// OwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnersRequestBuilderGetQueryParameters
}
// Application the application property
func (m *OwnersRequestBuilder) Application()(*iea108107f244d15625e030edb7372166349194daf3bc6cd117f48446d04e57f2.ApplicationRequestBuilder) {
    return iea108107f244d15625e030edb7372166349194daf3bc6cd117f48446d04e57f2.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOwnersRequestBuilderInternal instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    m := &OwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOwnersRequestBuilder instantiates a new OwnersRequestBuilder and sets the default values.
func NewOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *OwnersRequestBuilder) Count()(*i5482c610da612e49b7b76e839398da0032b32da37a31641dd6c3fe2d5b66130c.CountRequestBuilder) {
    return i5482c610da612e49b7b76e839398da0032b32da37a31641dd6c3fe2d5b66130c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OwnersRequestBuilder) Device()(*i5ac3da2b72c860c9d858ba8884f03ac263d08f62079ac3f62f3df69d20da6eb5.DeviceRequestBuilder) {
    return i5ac3da2b72c860c9d858ba8884f03ac263d08f62079ac3f62f3df69d20da6eb5.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the owners of the group who can be users or service principals. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand including nested $select. For example, /groups?$filter=startsWith(displayName,'Role')&$select=id,displayName&$expand=owners($select=id,userPrincipalName,displayName).
func (m *OwnersRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OwnersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *OwnersRequestBuilder) Group()(*ie46d71cd800842fcb2a698e0f4806ad7a00318a0cf8a6b36096801a9ab93016c.GroupRequestBuilder) {
    return ie46d71cd800842fcb2a698e0f4806ad7a00318a0cf8a6b36096801a9ab93016c.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact the orgContact property
func (m *OwnersRequestBuilder) OrgContact()(*i9ffbfdaeeea3bb4efea4ad0e6db2b4ee4e6e0f52f5325b168da679a18859e18e.OrgContactRequestBuilder) {
    return i9ffbfdaeeea3bb4efea4ad0e6db2b4ee4e6e0f52f5325b168da679a18859e18e.NewOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref the ref property
func (m *OwnersRequestBuilder) Ref()(*i327eef57a0e308487ea882dc08c72a66bfd918bbaff92b23558cb16e66dd1ba4.RefRequestBuilder) {
    return i327eef57a0e308487ea882dc08c72a66bfd918bbaff92b23558cb16e66dd1ba4.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnersRequestBuilder) ServicePrincipal()(*i30c6a904f09f7ba0c29704a7a72037781c52d7a49cd1e9f954d0e321ef034e10.ServicePrincipalRequestBuilder) {
    return i30c6a904f09f7ba0c29704a7a72037781c52d7a49cd1e9f954d0e321ef034e10.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User the user property
func (m *OwnersRequestBuilder) User()(*i1d617c3562064f6580fd55ae5e4013da163a0f961bf9b5f9fccc75f94b3f24e8.UserRequestBuilder) {
    return i1d617c3562064f6580fd55ae5e4013da163a0f961bf9b5f9fccc75f94b3f24e8.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
