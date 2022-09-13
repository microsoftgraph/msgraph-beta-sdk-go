package ownedobjects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i25be7a0abe3623e432f3d68d414148d114907c272263b51ade9c88e7396feede "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/serviceprincipal"
    i56c59dca0958c7e4544d9406366ddb1a4747561d1d8ebe793dff6b9c73fe3f96 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/count"
    i8018702d3846ca256f0d8c83b18ef29d5f76d470fbdee3bc0ae88763b4926944 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/group"
    i8bbbcf96afb909a4b9175c65e9efc752c3d426ed8dac3672190658251a65cc73 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/endpoint"
    ic9213a91056269d46044c3d272dddda936dac5e2969317d84a1ce5309de25111 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/application"
)

// OwnedObjectsRequestBuilder provides operations to manage the ownedObjects property of the microsoft.graph.servicePrincipal entity.
type OwnedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnedObjectsRequestBuilderGetQueryParameters directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
type OwnedObjectsRequestBuilderGetQueryParameters struct {
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
// OwnedObjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OwnedObjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OwnedObjectsRequestBuilderGetQueryParameters
}
// Application the application property
func (m *OwnedObjectsRequestBuilder) Application()(*ic9213a91056269d46044c3d272dddda936dac5e2969317d84a1ce5309de25111.ApplicationRequestBuilder) {
    return ic9213a91056269d46044c3d272dddda936dac5e2969317d84a1ce5309de25111.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOwnedObjectsRequestBuilderInternal instantiates a new OwnedObjectsRequestBuilder and sets the default values.
func NewOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedObjectsRequestBuilder) {
    m := &OwnedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects{?%24top*,%24skip*,%24search*,%24filter*,%24count*,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOwnedObjectsRequestBuilder instantiates a new OwnedObjectsRequestBuilder and sets the default values.
func NewOwnedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOwnedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *OwnedObjectsRequestBuilder) Count()(*i56c59dca0958c7e4544d9406366ddb1a4747561d1d8ebe793dff6b9c73fe3f96.CountRequestBuilder) {
    return i56c59dca0958c7e4544d9406366ddb1a4747561d1d8ebe793dff6b9c73fe3f96.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *OwnedObjectsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *OwnedObjectsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OwnedObjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Endpoint the endpoint property
func (m *OwnedObjectsRequestBuilder) Endpoint()(*i8bbbcf96afb909a4b9175c65e9efc752c3d426ed8dac3672190658251a65cc73.EndpointRequestBuilder) {
    return i8bbbcf96afb909a4b9175c65e9efc752c3d426ed8dac3672190658251a65cc73.NewEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand.
func (m *OwnedObjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *OwnedObjectsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *OwnedObjectsRequestBuilder) Group()(*i8018702d3846ca256f0d8c83b18ef29d5f76d470fbdee3bc0ae88763b4926944.GroupRequestBuilder) {
    return i8018702d3846ca256f0d8c83b18ef29d5f76d470fbdee3bc0ae88763b4926944.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnedObjectsRequestBuilder) ServicePrincipal()(*i25be7a0abe3623e432f3d68d414148d114907c272263b51ade9c88e7396feede.ServicePrincipalRequestBuilder) {
    return i25be7a0abe3623e432f3d68d414148d114907c272263b51ade9c88e7396feede.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
