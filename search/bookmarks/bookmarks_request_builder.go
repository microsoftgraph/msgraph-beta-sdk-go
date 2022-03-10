package bookmarks

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks/count"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// BookmarksRequestBuilder provides operations to manage the bookmarks property of the microsoft.graph.searchEntity entity.
type BookmarksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// BookmarksRequestBuilderGetOptions options for Get
type BookmarksRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *BookmarksRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// BookmarksRequestBuilderGetQueryParameters administrative answer in Microsoft Search results for common search queries in an organization.
type BookmarksRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// BookmarksRequestBuilderPostOptions options for Post
type BookmarksRequestBuilderPostOptions struct {
    // 
    Body id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.Bookmarkable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewBookmarksRequestBuilderInternal instantiates a new BookmarksRequestBuilder and sets the default values.
func NewBookmarksRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookmarksRequestBuilder) {
    m := &BookmarksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search/bookmarks{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewBookmarksRequestBuilder instantiates a new BookmarksRequestBuilder and sets the default values.
func NewBookmarksRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*BookmarksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBookmarksRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *BookmarksRequestBuilder) Count()(*id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297.CountRequestBuilder) {
    return id90852543ffe0a1612445ebbf1b2ba7ecfe3df20a8786bace4a5dd415613d297.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation administrative answer in Microsoft Search results for common search queries in an organization.
func (m *BookmarksRequestBuilder) CreateGetRequestInformation(options *BookmarksRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to bookmarks for search
func (m *BookmarksRequestBuilder) CreatePostRequestInformation(options *BookmarksRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get administrative answer in Microsoft Search results for common search queries in an organization.
func (m *BookmarksRequestBuilder) Get(options *BookmarksRequestBuilderGetOptions)(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.BookmarkCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.CreateBookmarkCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.BookmarkCollectionResponseable), nil
}
// Post create new navigation property to bookmarks for search
func (m *BookmarksRequestBuilder) Post(options *BookmarksRequestBuilderPostOptions)(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.Bookmarkable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.CreateBookmarkFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.Bookmarkable), nil
}
