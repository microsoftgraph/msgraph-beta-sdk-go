package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/query"
    i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks"
    ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e "github.com/microsoftgraph/msgraph-beta-sdk-go/search/qnas"
    ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/acronyms"
    i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c "github.com/microsoftgraph/msgraph-beta-sdk-go/search/qnas/item"
    i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d "github.com/microsoftgraph/msgraph-beta-sdk-go/search/acronyms/item"
    i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks/item"
)

// SearchRequestBuilder provides operations to manage the searchEntity singleton.
type SearchRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SearchRequestBuilderGetQueryParameters get search
type SearchRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SearchRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SearchRequestBuilderGetQueryParameters
}
// SearchRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Acronyms the acronyms property
func (m *SearchRequestBuilder) Acronyms()(*ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072.AcronymsRequestBuilder) {
    return ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072.NewAcronymsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcronymsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.acronyms.item collection
func (m *SearchRequestBuilder) AcronymsById(id string)(*i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d.AcronymItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["acronym%2Did"] = id
    }
    return i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d.NewAcronymItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bookmarks the bookmarks property
func (m *SearchRequestBuilder) Bookmarks()(*i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8.BookmarksRequestBuilder) {
    return i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8.NewBookmarksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookmarksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.bookmarks.item collection
func (m *SearchRequestBuilder) BookmarksById(id string)(*i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192.BookmarkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookmark%2Did"] = id
    }
    return i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192.NewBookmarkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSearchRequestBuilderInternal instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    m := &SearchRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSearchRequestBuilder instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get search
func (m *SearchRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get search
func (m *SearchRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SearchRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update search
func (m *SearchRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SearchEntityable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update search
func (m *SearchRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SearchEntityable, requestConfiguration *SearchRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get search
func (m *SearchRequestBuilder) Get(ctx context.Context, requestConfiguration *SearchRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SearchEntityable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSearchEntityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SearchEntityable), nil
}
// Patch update search
func (m *SearchRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SearchEntityable, requestConfiguration *SearchRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Qnas the qnas property
func (m *SearchRequestBuilder) Qnas()(*ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e.QnasRequestBuilder) {
    return ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e.NewQnasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QnasById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.qnas.item collection
func (m *SearchRequestBuilder) QnasById(id string)(*i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c.QnaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["qna%2Did"] = id
    }
    return i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c.NewQnaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Query the query property
func (m *SearchRequestBuilder) Query()(*i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09.QueryRequestBuilder) {
    return i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09.NewQueryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
