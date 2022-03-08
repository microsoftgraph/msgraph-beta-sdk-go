package notebooks

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4d79beb16dbb767abe6a6ec38dcfce27f372db814540df7e9b0ab024dde9c914 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/count"
    ia6a9cde8a714b3afb76fa303b6a6d8a84db1d0d7b0fe049fa3f45da347407f3d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/getrecentnotebookswithincludepersonalnotebooks"
    ib811612ec152214543ab35a36cafe81286d994d7b599eee468af8f4498a4e5fe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/getnotebookfromweburl"
)

// NotebooksRequestBuilder provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
type NotebooksRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotebooksRequestBuilderGetOptions options for Get
type NotebooksRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotebooksRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotebooksRequestBuilderGetQueryParameters the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
type NotebooksRequestBuilderGetQueryParameters struct {
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
// NotebooksRequestBuilderPostOptions options for Post
type NotebooksRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotebooksRequestBuilderInternal instantiates a new NotebooksRequestBuilder and sets the default values.
func NewNotebooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebooksRequestBuilder) {
    m := &NotebooksRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotebooksRequestBuilder instantiates a new NotebooksRequestBuilder and sets the default values.
func NewNotebooksRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebooksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotebooksRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *NotebooksRequestBuilder) Count()(*i4d79beb16dbb767abe6a6ec38dcfce27f372db814540df7e9b0ab024dde9c914.CountRequestBuilder) {
    return i4d79beb16dbb767abe6a6ec38dcfce27f372db814540df7e9b0ab024dde9c914.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebooksRequestBuilder) CreateGetRequestInformation(options *NotebooksRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to notebooks for me
func (m *NotebooksRequestBuilder) CreatePostRequestInformation(options *NotebooksRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebooksRequestBuilder) Get(options *NotebooksRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotebookCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateNotebookCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotebookCollectionResponseable), nil
}
func (m *NotebooksRequestBuilder) GetNotebookFromWebUrl()(*ib811612ec152214543ab35a36cafe81286d994d7b599eee468af8f4498a4e5fe.GetNotebookFromWebUrlRequestBuilder) {
    return ib811612ec152214543ab35a36cafe81286d994d7b599eee468af8f4498a4e5fe.NewGetNotebookFromWebUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRecentNotebooksWithIncludePersonalNotebooks provides operations to call the getRecentNotebooks method.
func (m *NotebooksRequestBuilder) GetRecentNotebooksWithIncludePersonalNotebooks(includePersonalNotebooks *bool)(*ia6a9cde8a714b3afb76fa303b6a6d8a84db1d0d7b0fe049fa3f45da347407f3d.GetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) {
    return ia6a9cde8a714b3afb76fa303b6a6d8a84db1d0d7b0fe049fa3f45da347407f3d.NewGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter, includePersonalNotebooks);
}
// Post create new navigation property to notebooks for me
func (m *NotebooksRequestBuilder) Post(options *NotebooksRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateNotebookFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable), nil
}
