package contenttypes

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i007bfa89748033f11654437f6e7ed4041238a7008f77fce348f7e30060cea1ca "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/addcopyfromcontenttypehub"
    ia7491f18503a7a56468628227e9ef82fd13d673f6565092396711e04a2c8c6bb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/addcopy"
    iaee576dc36a8127c08e2d0e9e9d875855ac261537dcbabae29813801021562fe "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/getcompatiblehubcontenttypes"
    ibc113b103dc1a0b976376c02ceda7690c4ee88bbc16325f0ce11ad848773300f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/contenttypes/count"
)

// ContentTypesRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentTypesRequestBuilderGetOptions options for Get
type ContentTypesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypesRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContentTypesRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypesRequestBuilderGetQueryParameters struct {
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
// ContentTypesRequestBuilderPostOptions options for Post
type ContentTypesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AddCopy the addCopy property
func (m *ContentTypesRequestBuilder) AddCopy()(*ia7491f18503a7a56468628227e9ef82fd13d673f6565092396711e04a2c8c6bb.AddCopyRequestBuilder) {
    return ia7491f18503a7a56468628227e9ef82fd13d673f6565092396711e04a2c8c6bb.NewAddCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddCopyFromContentTypeHub the addCopyFromContentTypeHub property
func (m *ContentTypesRequestBuilder) AddCopyFromContentTypeHub()(*i007bfa89748033f11654437f6e7ed4041238a7008f77fce348f7e30060cea1ca.AddCopyFromContentTypeHubRequestBuilder) {
    return i007bfa89748033f11654437f6e7ed4041238a7008f77fce348f7e30060cea1ca.NewAddCopyFromContentTypeHubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewContentTypesRequestBuilderInternal instantiates a new ContentTypesRequestBuilder and sets the default values.
func NewContentTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypesRequestBuilder) {
    m := &ContentTypesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypesRequestBuilder instantiates a new ContentTypesRequestBuilder and sets the default values.
func NewContentTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ContentTypesRequestBuilder) Count()(*ibc113b103dc1a0b976376c02ceda7690c4ee88bbc16325f0ce11ad848773300f.CountRequestBuilder) {
    return ibc113b103dc1a0b976376c02ceda7690c4ee88bbc16325f0ce11ad848773300f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypesRequestBuilder) CreateGetRequestInformation(options *ContentTypesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to contentTypes for sites
func (m *ContentTypesRequestBuilder) CreatePostRequestInformation(options *ContentTypesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the collection of content types present in this list.
func (m *ContentTypesRequestBuilder) Get(options *ContentTypesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeCollectionResponseable), nil
}
// GetCompatibleHubContentTypes provides operations to call the getCompatibleHubContentTypes method.
func (m *ContentTypesRequestBuilder) GetCompatibleHubContentTypes()(*iaee576dc36a8127c08e2d0e9e9d875855ac261537dcbabae29813801021562fe.GetCompatibleHubContentTypesRequestBuilder) {
    return iaee576dc36a8127c08e2d0e9e9d875855ac261537dcbabae29813801021562fe.NewGetCompatibleHubContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to contentTypes for sites
func (m *ContentTypesRequestBuilder) Post(options *ContentTypesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
