package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder provides operations to manage the itemCategory property of the microsoft.graph.item entity.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetQueryParameters get itemCategory from financials
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderInternal instantiates a new ItemCategoryRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) {
    m := &FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesInvoices/{salesInvoice%2Did}/salesInvoiceLines/{salesInvoiceLine%2Did}/item/itemCategory{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder instantiates a new ItemCategoryRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property itemCategory for financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get itemCategory from financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property itemCategory in financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property itemCategory for financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) Delete(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get itemCategory from financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable), nil
}
// Patch update the navigation property itemCategory in financials
func (m *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, requestConfiguration *FinancialsCompaniesItemSalesInvoicesItemSalesInvoiceLinesItemItemItemCategoryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemCategoryable), nil
}
