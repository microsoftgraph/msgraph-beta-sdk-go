package search

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb "github.com/microsoftgraph/msgraph-beta-sdk-go/models/search"
)

// QnasRequestBuilder provides operations to manage the qnas property of the microsoft.graph.searchEntity entity.
type QnasRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// QnasRequestBuilderGetQueryParameters administrative answer in Microsoft Search results that provide answers for specific search keywords in an organization.
type QnasRequestBuilderGetQueryParameters struct {
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
// QnasRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type QnasRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *QnasRequestBuilderGetQueryParameters
}
// QnasRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type QnasRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByQnaId provides operations to manage the qnas property of the microsoft.graph.searchEntity entity.
// returns a *QnasQnaItemRequestBuilder when successful
func (m *QnasRequestBuilder) ByQnaId(qnaId string)(*QnasQnaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if qnaId != "" {
        urlTplParams["qna%2Did"] = qnaId
    }
    return NewQnasQnaItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewQnasRequestBuilderInternal instantiates a new QnasRequestBuilder and sets the default values.
func NewQnasRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*QnasRequestBuilder) {
    m := &QnasRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/search/qnas{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewQnasRequestBuilder instantiates a new QnasRequestBuilder and sets the default values.
func NewQnasRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*QnasRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewQnasRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *QnasCountRequestBuilder when successful
func (m *QnasRequestBuilder) Count()(*QnasCountRequestBuilder) {
    return NewQnasCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get administrative answer in Microsoft Search results that provide answers for specific search keywords in an organization.
// returns a QnaCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *QnasRequestBuilder) Get(ctx context.Context, requestConfiguration *QnasRequestBuilderGetRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.QnaCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateQnaCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.QnaCollectionResponseable), nil
}
// Post create new navigation property to qnas for search
// returns a Qnaable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *QnasRequestBuilder) Post(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, requestConfiguration *QnasRequestBuilderPostRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateQnaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable), nil
}
// ToGetRequestInformation administrative answer in Microsoft Search results that provide answers for specific search keywords in an organization.
// returns a *RequestInformation when successful
func (m *QnasRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *QnasRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to qnas for search
// returns a *RequestInformation when successful
func (m *QnasRequestBuilder) ToPostRequestInformation(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Qnaable, requestConfiguration *QnasRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *QnasRequestBuilder when successful
func (m *QnasRequestBuilder) WithUrl(rawUrl string)(*QnasRequestBuilder) {
    return NewQnasRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
