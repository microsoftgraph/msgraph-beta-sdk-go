package createdobjects

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5447f788a23739b35cc59dcf432024071a2db85a939ba1e0ffc85d2bb0ab734f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/createdobjects/serviceprincipal"
    ibbdd6a42ebe179b71d926db2993ac628ac1b9bab0a0c0ce1e4b25a77b89ce770 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/createdobjects/count"
)

// CreatedObjectsRequestBuilder provides operations to manage the createdObjects property of the microsoft.graph.user entity.
type CreatedObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreatedObjectsRequestBuilderGetQueryParameters directory objects that were created by the user. Read-only. Nullable.
type CreatedObjectsRequestBuilderGetQueryParameters struct {
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
// CreatedObjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreatedObjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CreatedObjectsRequestBuilderGetQueryParameters
}
// NewCreatedObjectsRequestBuilderInternal instantiates a new CreatedObjectsRequestBuilder and sets the default values.
func NewCreatedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreatedObjectsRequestBuilder) {
    m := &CreatedObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/createdObjects{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreatedObjectsRequestBuilder instantiates a new CreatedObjectsRequestBuilder and sets the default values.
func NewCreatedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreatedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *CreatedObjectsRequestBuilder) Count()(*ibbdd6a42ebe179b71d926db2993ac628ac1b9bab0a0c0ce1e4b25a77b89ce770.CountRequestBuilder) {
    return ibbdd6a42ebe179b71d926db2993ac628ac1b9bab0a0c0ce1e4b25a77b89ce770.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation directory objects that were created by the user. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration directory objects that were created by the user. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CreatedObjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get directory objects that were created by the user. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler directory objects that were created by the user. Read-only. Nullable.
func (m *CreatedObjectsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CreatedObjectsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
// ServicePrincipal the servicePrincipal property
func (m *CreatedObjectsRequestBuilder) ServicePrincipal()(*i5447f788a23739b35cc59dcf432024071a2db85a939ba1e0ffc85d2bb0ab734f.ServicePrincipalRequestBuilder) {
    return i5447f788a23739b35cc59dcf432024071a2db85a939ba1e0ffc85d2bb0ab734f.NewServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
