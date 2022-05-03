package androidmanagedstoreappconfigurationschemas

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i50046ad7f0c0dc69b9a80160b1a8e94717b6c54b629503f2ca0036aa3d95523e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreappconfigurationschemas/count"
)

// AndroidManagedStoreAppConfigurationSchemasRequestBuilder provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
type AndroidManagedStoreAppConfigurationSchemasRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetQueryParameters android Enterprise app configuration schema entities.
type AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetQueryParameters struct {
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
// AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetQueryParameters
}
// AndroidManagedStoreAppConfigurationSchemasRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAppConfigurationSchemasRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal instantiates a new AndroidManagedStoreAppConfigurationSchemasRequestBuilder and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    m := &AndroidManagedStoreAppConfigurationSchemasRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAppConfigurationSchemas{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidManagedStoreAppConfigurationSchemasRequestBuilder instantiates a new AndroidManagedStoreAppConfigurationSchemasRequestBuilder and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemasRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) Count()(*i50046ad7f0c0dc69b9a80160b1a8e94717b6c54b629503f2ca0036aa3d95523e.CountRequestBuilder) {
    return i50046ad7f0c0dc69b9a80160b1a8e94717b6c54b629503f2ca0036aa3d95523e.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformationWithRequestConfiguration android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) GetWithResponseHandler(requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaCollectionResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) GetWithResponseHandler(requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAppConfigurationSchemaCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaCollectionResponseable), nil
}
// PostWithResponseHandler create new navigation property to androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler create new navigation property to androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable), nil
}
