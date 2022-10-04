package acronyms

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb "github.com/microsoftgraph/msgraph-beta-sdk-go/models/search"
    i4676165c3a2c7675fa9df0975e39bd9ef9f4527be649374d7054fde3700e4256 "github.com/microsoftgraph/msgraph-beta-sdk-go/search/acronyms/count"
)

// AcronymsRequestBuilder provides operations to manage the acronyms property of the microsoft.graph.searchEntity entity.
type AcronymsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AcronymsRequestBuilderGetQueryParameters get a list of the acronym objects and their properties.
type AcronymsRequestBuilderGetQueryParameters struct {
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
// AcronymsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AcronymsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AcronymsRequestBuilderGetQueryParameters
}
// AcronymsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AcronymsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAcronymsRequestBuilderInternal instantiates a new AcronymsRequestBuilder and sets the default values.
func NewAcronymsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AcronymsRequestBuilder) {
    m := &AcronymsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/search/acronyms{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAcronymsRequestBuilder instantiates a new AcronymsRequestBuilder and sets the default values.
func NewAcronymsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AcronymsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAcronymsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *AcronymsRequestBuilder) Count()(*i4676165c3a2c7675fa9df0975e39bd9ef9f4527be649374d7054fde3700e4256.CountRequestBuilder) {
    return i4676165c3a2c7675fa9df0975e39bd9ef9f4527be649374d7054fde3700e4256.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the acronym objects and their properties.
func (m *AcronymsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AcronymsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create a new acronym object.
func (m *AcronymsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Acronymable, requestConfiguration *AcronymsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get a list of the acronym objects and their properties.
func (m *AcronymsRequestBuilder) Get(ctx context.Context, requestConfiguration *AcronymsRequestBuilderGetRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.AcronymCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateAcronymCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.AcronymCollectionResponseable), nil
}
// Post create a new acronym object.
func (m *AcronymsRequestBuilder) Post(ctx context.Context, body iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Acronymable, requestConfiguration *AcronymsRequestBuilderPostRequestConfiguration)(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Acronymable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.CreateAcronymFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iea48ada8ff44901e797bb459ff00d73b62bd6a3bff0f3314a5377938749128cb.Acronymable), nil
}
