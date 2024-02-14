package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder provides operations to manage the documents property of the microsoft.graph.printJob entity.
type PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetQueryParameters get documents from print
type PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetQueryParameters
}
// PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderInternal instantiates a new PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder and sets the default values.
func NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) {
    m := &PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/documents/{printDocument%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder instantiates a new PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder and sets the default values.
func NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the print entity.
// returns a *PrintersItemJobsItemDocumentsItemValueContentRequestBuilder when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) Content()(*PrintersItemJobsItemDocumentsItemValueContentRequestBuilder) {
    return NewPrintersItemJobsItemDocumentsItemValueContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *PrintersItemJobsItemDocumentsItemCreateUploadSessionRequestBuilder when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) CreateUploadSession()(*PrintersItemJobsItemDocumentsItemCreateUploadSessionRequestBuilder) {
    return NewPrintersItemJobsItemDocumentsItemCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property documents for print
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get documents from print
// returns a PrintDocumentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintDocumentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable), nil
}
// Patch update the navigation property documents in print
// returns a PrintDocumentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintDocumentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable), nil
}
// ToDeleteRequestInformation delete navigation property documents for print
// returns a *RequestInformation when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/documents/{printDocument%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get documents from print
// returns a *RequestInformation when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property documents in print
// returns a *RequestInformation when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintDocumentable, requestConfiguration *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/print/printers/{printer%2Did}/jobs/{printJob%2Did}/documents/{printDocument%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder when successful
func (m *PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) WithUrl(rawUrl string)(*PrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder) {
    return NewPrintersItemJobsItemDocumentsPrintDocumentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
