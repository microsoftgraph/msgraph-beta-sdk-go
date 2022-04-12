package bookmarks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb "github.com/microsoftgraph/msgraph-beta-sdk-go/models/search"
    id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks/count"
)

// BookmarksRequestBuilder provides operations to manage the bookmarks property of the microsoft.graph.searchEntity entity.
type BookmarksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// BookmarksRequestBuilderGetOptions options for Get
type BookmarksRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BookmarksRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// BookmarksRequestBuilderGetQueryParameters administrative answer in Microsoft Search results for common search queries in an organization.
type BookmarksRequestBuilderGetQueryParameters struct {
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
// BookmarksRequestBuilderPostOptions options for Post
type BookmarksRequestBuilderPostOptions struct {
    // 
    Body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Bookmarkable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewBookmarksRequestBuilderInternal instantiates a new BookmarksRequestBuilder and sets the default values.
func NewBookmarksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookmarksRequestBuilder) {
    m := &BookmarksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search/bookmarks{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookmarksRequestBuilder instantiates a new BookmarksRequestBuilder and sets the default values.
func NewBookmarksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BookmarksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookmarksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *BookmarksRequestBuilder) Count()(*id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297.CountRequestBuilder) {
    return id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation administrative answer in Microsoft Search results for common search queries in an organization.
func (m *BookmarksRequestBuilder) CreateGetRequestInformation(options *BookmarksRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to bookmarks for search
func (m *BookmarksRequestBuilder) CreatePostRequestInformation(options *BookmarksRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get administrative answer in Microsoft Search results for common search queries in an organization.
func (m *BookmarksRequestBuilder) Get(options *BookmarksRequestBuilderGetOptions)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.BookmarkCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateBookmarkCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.BookmarkCollectionResponseable), nil
}
// Post create new navigation property to bookmarks for search
func (m *BookmarksRequestBuilder) Post(options *BookmarksRequestBuilderPostOptions)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Bookmarkable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateBookmarkFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Bookmarkable), nil
}
