package deleteditems

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2589430a289fe8e95a7acdf2d80ee9813ad5587a9980306516875cdabb3a3cb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/group"
    i42ef6142cdf65b13c8c619b92627f1ec8b7b6fce8f4b66e55c2b57da753f54e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/user"
    i4389b9aaa346a5fdf53bb167cf00b6bdc4aa52db66b0adf49337c8a1c78f8a32 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/getuserownedobjects"
    i582b3dfc89fb8335c62d5c1b642f724197229635d5f263e6cc636b61190df708 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/count"
    i59e6d9a924e0e45d290fd3bf7915868207b145fc6d55661a4adb34ab83d49d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/validateproperties"
    ie1025c6f61e31a1860f913437eacc721e9877d2b918c5dd29552da8ecb803f5f "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/getbyids"
    ifd1aec9b29d58b8b58e8239f298c349cd212a89eb493ad01299f5fbb91de0bc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/deleteditems/application"
)

// DeletedItemsRequestBuilder provides operations to manage the deletedItems property of the microsoft.graph.directory entity.
type DeletedItemsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeletedItemsRequestBuilderGetQueryParameters recently deleted items. Read-only. Nullable.
type DeletedItemsRequestBuilderGetQueryParameters struct {
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
// DeletedItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedItemsRequestBuilderGetQueryParameters
}
// DeletedItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Application the application property
func (m *DeletedItemsRequestBuilder) Application()(*ifd1aec9b29d58b8b58e8239f298c349cd212a89eb493ad01299f5fbb91de0bc5.ApplicationRequestBuilder) {
    return ifd1aec9b29d58b8b58e8239f298c349cd212a89eb493ad01299f5fbb91de0bc5.NewApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeletedItemsRequestBuilderInternal instantiates a new DeletedItemsRequestBuilder and sets the default values.
func NewDeletedItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsRequestBuilder) {
    m := &DeletedItemsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/deletedItems{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeletedItemsRequestBuilder instantiates a new DeletedItemsRequestBuilder and sets the default values.
func NewDeletedItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeletedItemsRequestBuilder) Count()(*i582b3dfc89fb8335c62d5c1b642f724197229635d5f263e6cc636b61190df708.CountRequestBuilder) {
    return i582b3dfc89fb8335c62d5c1b642f724197229635d5f263e6cc636b61190df708.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation recently deleted items. Read-only. Nullable.
func (m *DeletedItemsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration recently deleted items. Read-only. Nullable.
func (m *DeletedItemsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeletedItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deletedItems for directory
func (m *DeletedItemsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to deletedItems for directory
func (m *DeletedItemsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *DeletedItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get recently deleted items. Read-only. Nullable.
func (m *DeletedItemsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetByIds the getByIds property
func (m *DeletedItemsRequestBuilder) GetByIds()(*ie1025c6f61e31a1860f913437eacc721e9877d2b918c5dd29552da8ecb803f5f.GetByIdsRequestBuilder) {
    return ie1025c6f61e31a1860f913437eacc721e9877d2b918c5dd29552da8ecb803f5f.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *DeletedItemsRequestBuilder) GetUserOwnedObjects()(*i4389b9aaa346a5fdf53bb167cf00b6bdc4aa52db66b0adf49337c8a1c78f8a32.GetUserOwnedObjectsRequestBuilder) {
    return i4389b9aaa346a5fdf53bb167cf00b6bdc4aa52db66b0adf49337c8a1c78f8a32.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler recently deleted items. Read-only. Nullable.
func (m *DeletedItemsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeletedItemsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *DeletedItemsRequestBuilder) Group()(*i2589430a289fe8e95a7acdf2d80ee9813ad5587a9980306516875cdabb3a3cb4.GroupRequestBuilder) {
    return i2589430a289fe8e95a7acdf2d80ee9813ad5587a9980306516875cdabb3a3cb4.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deletedItems for directory
func (m *DeletedItemsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to deletedItems for directory
func (m *DeletedItemsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *DeletedItemsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// User the user property
func (m *DeletedItemsRequestBuilder) User()(*i42ef6142cdf65b13c8c619b92627f1ec8b7b6fce8f4b66e55c2b57da753f54e6.UserRequestBuilder) {
    return i42ef6142cdf65b13c8c619b92627f1ec8b7b6fce8f4b66e55c2b57da753f54e6.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *DeletedItemsRequestBuilder) ValidateProperties()(*i59e6d9a924e0e45d290fd3bf7915868207b145fc6d55661a4adb34ab83d49d96.ValidatePropertiesRequestBuilder) {
    return i59e6d9a924e0e45d290fd3bf7915868207b145fc6d55661a4adb34ab83d49d96.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
