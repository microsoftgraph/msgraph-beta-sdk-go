package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb "github.com/microsoftgraph/msgraph-beta-sdk-go/models/search"
)

// SearchQnasRequestBuilder provides operations to manage the qnas property of the microsoft.graph.searchEntity entity.
type SearchQnasRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SearchQnasRequestBuilderGetQueryParameters get a list of the qna objects and their properties.
type SearchQnasRequestBuilderGetQueryParameters struct {
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
// SearchQnasRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchQnasRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SearchQnasRequestBuilderGetQueryParameters
}
// SearchQnasRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SearchQnasRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSearchQnasRequestBuilderInternal instantiates a new QnasRequestBuilder and sets the default values.
func NewSearchQnasRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchQnasRequestBuilder) {
    m := &SearchQnasRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search/qnas{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSearchQnasRequestBuilder instantiates a new QnasRequestBuilder and sets the default values.
func NewSearchQnasRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SearchQnasRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSearchQnasRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *SearchQnasRequestBuilder) Count()(*SearchQnasCountRequestBuilder) {
    return NewSearchQnasCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the qna objects and their properties.
func (m *SearchQnasRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SearchQnasRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create a new qna object.
func (m *SearchQnasRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, requestConfiguration *SearchQnasRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get a list of the qna objects and their properties.
func (m *SearchQnasRequestBuilder) Get(ctx context.Context, requestConfiguration *SearchQnasRequestBuilderGetRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.QnaCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateQnaCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.QnaCollectionResponseable), nil
}
// Post create a new qna object.
func (m *SearchQnasRequestBuilder) Post(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, requestConfiguration *SearchQnasRequestBuilderPostRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateQnaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable), nil
}
