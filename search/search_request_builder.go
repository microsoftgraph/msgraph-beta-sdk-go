package search

import (
    i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/query"
    i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks"
    ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e "github.com/microsoftgraph/msgraph-beta-sdk-go/search/qnas"
    ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/acronyms"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c "github.com/microsoftgraph/msgraph-beta-sdk-go/search/qnas/item"
    i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d "github.com/microsoftgraph/msgraph-beta-sdk-go/search/acronyms/item"
    i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/bookmarks/item"
)

// SearchRequestBuilder builds and executes requests for operations under \search
type SearchRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SearchRequestBuilderGetOptions options for Get
type SearchRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SearchRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SearchRequestBuilderGetQueryParameters get search
type SearchRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SearchRequestBuilderPatchOptions options for Patch
type SearchRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SearchEntity;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SearchRequestBuilder) Acronyms()(*ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072.AcronymsRequestBuilder) {
    return ib721af151194f36463bd473df1277a52673500924a4562283e137a3f7ed94072.NewAcronymsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcronymsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.acronyms.item collection
func (m *SearchRequestBuilder) AcronymsById(id string)(*i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d.AcronymRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["acronym_id"] = id
    }
    return i74358e4db7f904e8f49fdcb30f91642f6639093ec5390b7ca96d2a396083031d.NewAcronymRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SearchRequestBuilder) Bookmarks()(*i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8.BookmarksRequestBuilder) {
    return i98acee39b1010d46fe63cc02c3afafb76788d106b29f8cbe8718dc6507b1faa8.NewBookmarksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookmarksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.bookmarks.item collection
func (m *SearchRequestBuilder) BookmarksById(id string)(*i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192.BookmarkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookmark_id"] = id
    }
    return i8b7285689d9aced3e20dd7af58f5c67035f37b8fe34ead1766f01cc855dbb192.NewBookmarkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSearchRequestBuilderInternal instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SearchRequestBuilder) {
    m := &SearchRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSearchRequestBuilder instantiates a new SearchRequestBuilder and sets the default values.
func NewSearchRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get search
func (m *SearchRequestBuilder) CreateGetRequestInformation(options *SearchRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update search
func (m *SearchRequestBuilder) CreatePatchRequestInformation(options *SearchRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Get get search
func (m *SearchRequestBuilder) Get(options *SearchRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SearchEntity, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSearchEntity() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SearchEntity), nil
}
// Patch update search
func (m *SearchRequestBuilder) Patch(options *SearchRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *SearchRequestBuilder) Qnas()(*ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e.QnasRequestBuilder) {
    return ia2bf4f3fe5831ab1e67fee5e544f828bb45264fc9a4823e780f848a593eed31e.NewQnasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QnasById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.search.qnas.item collection
func (m *SearchRequestBuilder) QnasById(id string)(*i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c.QnaRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["qna_id"] = id
    }
    return i5a7b1d161e5927fae4776d21687bd98261ca983357fb199581c959d3642cfc0c.NewQnaRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SearchRequestBuilder) Query()(*i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09.QueryRequestBuilder) {
    return i8bbfec18807af5d9aab08ad55e8f88a28fc74be92766b92e28c073d7011b7a09.NewQueryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
