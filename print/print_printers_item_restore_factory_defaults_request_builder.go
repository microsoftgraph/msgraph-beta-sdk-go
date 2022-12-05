package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintPrintersItemRestoreFactoryDefaultsRequestBuilder provides operations to call the restoreFactoryDefaults method.
type PrintPrintersItemRestoreFactoryDefaultsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintPrintersItemRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintPrintersItemRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilderInternal instantiates a new RestoreFactoryDefaultsRequestBuilder and sets the default values.
func NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrintersItemRestoreFactoryDefaultsRequestBuilder) {
    m := &PrintPrintersItemRestoreFactoryDefaultsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printers/{printer%2Did}/microsoft.graph.restoreFactoryDefaults";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilder instantiates a new RestoreFactoryDefaultsRequestBuilder and sets the default values.
func NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintPrintersItemRestoreFactoryDefaultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintPrintersItemRestoreFactoryDefaultsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation restore a printer's default settings to the values specified by the manufacturer.
func (m *PrintPrintersItemRestoreFactoryDefaultsRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *PrintPrintersItemRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post restore a printer's default settings to the values specified by the manufacturer.
func (m *PrintPrintersItemRestoreFactoryDefaultsRequestBuilder) Post(ctx context.Context, requestConfiguration *PrintPrintersItemRestoreFactoryDefaultsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
