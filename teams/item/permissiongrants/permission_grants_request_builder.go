package permissiongrants

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1a625e10aabbf5c4a31b1d9cc528aa53d67ede60c1f8bd97fd007dbb6a4a2189 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/getuserownedobjects"
    i28f94d0e9eca0c7cb482c21b35811b617ce21a754b0da440004c62a5f2478b2b "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/count"
    i50d4a71dabea8b5b3fd086fa51a466f7cfdb57afd7d91e7bcb6bcc80cd7d99ed "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/validateproperties"
    i6818d719c02f3aa456643a7d27de984f3fda3ba75a287e843f200e4b04df3b0a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/permissiongrants/getbyids"
)

// PermissionGrantsRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
type PermissionGrantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PermissionGrantsRequestBuilderGetQueryParameters list all resource-specific permission grants on the team. This list specifies the Azure AD apps that have access to the **team**, along with the corresponding kind of resource-specific access that each app has.
type PermissionGrantsRequestBuilderGetQueryParameters struct {
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
// PermissionGrantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantsRequestBuilderGetQueryParameters
}
// PermissionGrantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantsRequestBuilderInternal instantiates a new PermissionGrantsRequestBuilder and sets the default values.
func NewPermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantsRequestBuilder) {
    m := &PermissionGrantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/permissionGrants{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPermissionGrantsRequestBuilder instantiates a new PermissionGrantsRequestBuilder and sets the default values.
func NewPermissionGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PermissionGrantsRequestBuilder) Count()(*i28f94d0e9eca0c7cb482c21b35811b617ce21a754b0da440004c62a5f2478b2b.CountRequestBuilder) {
    return i28f94d0e9eca0c7cb482c21b35811b617ce21a754b0da440004c62a5f2478b2b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation list all resource-specific permission grants on the team. This list specifies the Azure AD apps that have access to the **team**, along with the corresponding kind of resource-specific access that each app has.
func (m *PermissionGrantsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to permissionGrants for teams
func (m *PermissionGrantsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *PermissionGrantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get list all resource-specific permission grants on the team. This list specifies the Azure AD apps that have access to the **team**, along with the corresponding kind of resource-specific access that each app has.
func (m *PermissionGrantsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantCollectionResponseable), nil
}
// GetByIds provides operations to call the getByIds method.
func (m *PermissionGrantsRequestBuilder) GetByIds()(*i6818d719c02f3aa456643a7d27de984f3fda3ba75a287e843f200e4b04df3b0a.GetByIdsRequestBuilder) {
    return i6818d719c02f3aa456643a7d27de984f3fda3ba75a287e843f200e4b04df3b0a.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects provides operations to call the getUserOwnedObjects method.
func (m *PermissionGrantsRequestBuilder) GetUserOwnedObjects()(*i1a625e10aabbf5c4a31b1d9cc528aa53d67ede60c1f8bd97fd007dbb6a4a2189.GetUserOwnedObjectsRequestBuilder) {
    return i1a625e10aabbf5c4a31b1d9cc528aa53d67ede60c1f8bd97fd007dbb6a4a2189.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to permissionGrants for teams
func (m *PermissionGrantsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *PermissionGrantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *PermissionGrantsRequestBuilder) ValidateProperties()(*i50d4a71dabea8b5b3fd086fa51a466f7cfdb57afd7d91e7bcb6bcc80cd7d99ed.ValidatePropertiesRequestBuilder) {
    return i50d4a71dabea8b5b3fd086fa51a466f7cfdb57afd7d91e7bcb6bcc80cd7d99ed.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
