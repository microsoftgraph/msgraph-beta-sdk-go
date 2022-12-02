package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalsItemOwnedObjectsRequestBuilder provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
type ServicePrincipalsItemOwnedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalsItemOwnedObjectsRequestBuilderGetQueryParameters retrieve a list of objects owned by the servicePrincipal.  This could include applications or groups.
type ServicePrincipalsItemOwnedObjectsRequestBuilderGetQueryParameters struct {
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
// ServicePrincipalsItemOwnedObjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalsItemOwnedObjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalsItemOwnedObjectsRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) Application()(*ServicePrincipalsItemOwnedObjectsApplicationRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalsItemOwnedObjectsRequestBuilderInternal instantiates a new OwnedObjectsRequestBuilder and sets the default values.
func NewServicePrincipalsItemOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemOwnedObjectsRequestBuilder) {
    m := &ServicePrincipalsItemOwnedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemOwnedObjectsRequestBuilder instantiates a new OwnedObjectsRequestBuilder and sets the default values.
func NewServicePrincipalsItemOwnedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemOwnedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemOwnedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) Count()(*ServicePrincipalsItemOwnedObjectsCountRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of objects owned by the servicePrincipal.  This could include applications or groups.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalsItemOwnedObjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Endpoint casts the previous resource to endpoint.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) Endpoint()(*ServicePrincipalsItemOwnedObjectsEndpointRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve a list of objects owned by the servicePrincipal.  This could include applications or groups.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalsItemOwnedObjectsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) Group()(*ServicePrincipalsItemOwnedObjectsGroupRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ServicePrincipalsItemOwnedObjectsRequestBuilder) ServicePrincipal()(*ServicePrincipalsItemOwnedObjectsServicePrincipalRequestBuilder) {
    return NewServicePrincipalsItemOwnedObjectsServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
