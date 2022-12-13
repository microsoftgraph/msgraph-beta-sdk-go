package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiIndicatorsSubmitTiIndicatorsRequestBuilder provides operations to call the submitTiIndicators method.
type TiIndicatorsSubmitTiIndicatorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal instantiates a new SubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    m := &TiIndicatorsSubmitTiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators/microsoft.graph.submitTiIndicators";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTiIndicatorsSubmitTiIndicatorsRequestBuilder instantiates a new SubmitTiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsSubmitTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body TiIndicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post upload multiple threat intelligence (TI) indicators in one request instead of multiple requests.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/tiindicator-submittiindicators?view=graph-rest-1.0
func (m *TiIndicatorsSubmitTiIndicatorsRequestBuilder) Post(ctx context.Context, body TiIndicatorsSubmitTiIndicatorsPostRequestBodyable, requestConfiguration *TiIndicatorsSubmitTiIndicatorsRequestBuilderPostRequestConfiguration)(TiIndicatorsSubmitTiIndicatorsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateTiIndicatorsSubmitTiIndicatorsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TiIndicatorsSubmitTiIndicatorsResponseable), nil
}
