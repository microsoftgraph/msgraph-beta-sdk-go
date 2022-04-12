package notebooks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0830ab625b9359734e1d1f0b8661894992aec435a3c30d278b52809328b654c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/getnotebookfromweburl"
    i19e59018d356110684d7b1f87f39ab0f94e288118dd13f9d1b0ef94a27c69c0d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/count"
    iedaf1f3b30812c4a7ab4084a3fbee8b9774634abe05eca1985e30fb58db329cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/getrecentnotebookswithincludepersonalnotebooks"
)

// NotebooksRequestBuilder provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
type NotebooksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NotebooksRequestBuilderGetOptions options for Get
type NotebooksRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NotebooksRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NotebooksRequestBuilderGetQueryParameters the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
type NotebooksRequestBuilderGetQueryParameters struct {
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
// NotebooksRequestBuilderPostOptions options for Post
type NotebooksRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Notebookable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewNotebooksRequestBuilderInternal instantiates a new NotebooksRequestBuilder and sets the default values.
func NewNotebooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotebooksRequestBuilder) {
    m := &NotebooksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/onenote/notebooks{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotebooksRequestBuilder instantiates a new NotebooksRequestBuilder and sets the default values.
func NewNotebooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NotebooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotebooksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *NotebooksRequestBuilder) Count()(*i19e59018d356110684d7b1f87f39ab0f94e288118dd13f9d1b0ef94a27c69c0d.CountRequestBuilder) {
    return i19e59018d356110684d7b1f87f39ab0f94e288118dd13f9d1b0ef94a27c69c0d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebooksRequestBuilder) CreateGetRequestInformation(options *NotebooksRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to notebooks for groups
func (m *NotebooksRequestBuilder) CreatePostRequestInformation(options *NotebooksRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebooksRequestBuilder) Get(options *NotebooksRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotebookCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateNotebookCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NotebookCollectionResponseable), nil
}
// GetNotebookFromWebUrl the getNotebookFromWebUrl property
func (m *NotebooksRequestBuilder) GetNotebookFromWebUrl()(*i0830ab625b9359734e1d1f0b8661894992aec435a3c30d278b52809328b654c6.GetNotebookFromWebUrlRequestBuilder) {
    return i0830ab625b9359734e1d1f0b8661894992aec435a3c30d278b52809328b654c6.NewGetNotebookFromWebUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRecentNotebooksWithIncludePersonalNotebooks provides operations to call the getRecentNotebooks method.
func (m *NotebooksRequestBuilder) GetRecentNotebooksWithIncludePersonalNotebooks(includePersonalNotebooks *bool)(*iedaf1f3b30812c4a7ab4084a3fbee8b9774634abe05eca1985e30fb58db329cb.GetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) {
    return iedaf1f3b30812c4a7ab4084a3fbee8b9774634abe05eca1985e30fb58db329cb.NewGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter, includePersonalNotebooks);
}
// Post create new navigation property to notebooks for groups
func (m *NotebooksRequestBuilder) Post(options *NotebooksRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Notebookable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateNotebookFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Notebookable), nil
}
