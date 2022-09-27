package ownedobjects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i38b6bcdb8583f7fc3c15b40a7b6804ec3035b7dd491bd08c078fe10d5e0a884b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects/serviceprincipal"
    i61bea30e212f931c96aa3ecc9e7f3e4d9094b11a682a42233bc692755285c5cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects/group"
    i6abdeca92b8bcaac9f60823fd72189369671d17efbcbb21d501d3073582f37fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects/count"
    ic041e40f6a1573129defd7e3a13bfaee595aac6da323d7f2a3dd69c1b291d000 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects/application"
)

// OwnedObjectsRequestBuilder provides operations to manage the ownedObjects property of the microsoft.graph.user entity.
type OwnedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OwnedObjectsRequestBuilderGetQueryParameters directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
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
func (m *OwnedObjectsRequestBuilder) Application()(*ic041e40f6a1573129defd7e3a13bfaee595aac6da323d7f2a3dd69c1b291d000.ApplicationRequestBuilder) {
    return ic041e40f6a1573129defd7e3a13bfaee595aac6da323d7f2a3dd69c1b291d000.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOwnedObjectsRequestBuilderInternal instantiates a new OwnedObjectsRequestBuilder and sets the default values.
func NewOwnedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OwnedObjectsRequestBuilder) {
    m := &OwnedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/ownedObjects{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
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
func (m *OwnedObjectsRequestBuilder) Count()(*i6abdeca92b8bcaac9f60823fd72189369671d17efbcbb21d501d3073582f37fc.CountRequestBuilder) {
    return i6abdeca92b8bcaac9f60823fd72189369671d17efbcbb21d501d3073582f37fc.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
func (m *OwnedObjectsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
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
// Get directory objects that are owned by the user. Read-only. Nullable. Supports $expand.
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
func (m *OwnedObjectsRequestBuilder) Group()(*i61bea30e212f931c96aa3ecc9e7f3e4d9094b11a682a42233bc692755285c5cf.GroupRequestBuilder) {
    return i61bea30e212f931c96aa3ecc9e7f3e4d9094b11a682a42233bc692755285c5cf.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal the servicePrincipal property
func (m *OwnedObjectsRequestBuilder) ServicePrincipal()(*i38b6bcdb8583f7fc3c15b40a7b6804ec3035b7dd491bd08c078fe10d5e0a884b.ServicePrincipalRequestBuilder) {
    return i38b6bcdb8583f7fc3c15b40a7b6804ec3035b7dd491bd08c078fe10d5e0a884b.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
