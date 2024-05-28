package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder provides operations to call the getCapabilities method.
type PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderInternal instantiates a new PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder and sets the default values.
func NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) {
    m := &PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/printers/{printer%2Did}/getCapabilities()", pathParameters),
    }
    return m
}
// NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder instantiates a new PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder and sets the default values.
func NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of capabilities for the printer.
// Deprecated: The getCapabilities API is deprecated and will stop returning data on July 31, 2023. Please use the capabilities property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrinterCapabilitiesable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/printer-getcapabilities?view=graph-rest-beta
func (m *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrinterCapabilitiesable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrinterCapabilitiesFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrinterCapabilitiesable), nil
}
// ToGetRequestInformation get a list of capabilities for the printer.
// Deprecated: The getCapabilities API is deprecated and will stop returning data on July 31, 2023. Please use the capabilities property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The getCapabilities API is deprecated and will stop returning data on July 31, 2023. Please use the capabilities property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder when successful
func (m *PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) WithUrl(rawUrl string)(*PrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder) {
    return NewPrintersItemGetcapabilitiesGetCapabilitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
