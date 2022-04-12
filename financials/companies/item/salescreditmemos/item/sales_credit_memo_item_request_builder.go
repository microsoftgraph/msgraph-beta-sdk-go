package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer"
    i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines"
    i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/currency"
    i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/paymentterm"
    i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item"
)

// SalesCreditMemoItemRequestBuilder provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
type SalesCreditMemoItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SalesCreditMemoItemRequestBuilderDeleteOptions options for Delete
type SalesCreditMemoItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SalesCreditMemoItemRequestBuilderGetOptions options for Get
type SalesCreditMemoItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesCreditMemoItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SalesCreditMemoItemRequestBuilderGetQueryParameters get salesCreditMemos from financials
type SalesCreditMemoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesCreditMemoItemRequestBuilderPatchOptions options for Patch
type SalesCreditMemoItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewSalesCreditMemoItemRequestBuilderInternal instantiates a new SalesCreditMemoItemRequestBuilder and sets the default values.
func NewSalesCreditMemoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesCreditMemoItemRequestBuilder) {
    m := &SalesCreditMemoItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesCreditMemoItemRequestBuilder instantiates a new SalesCreditMemoItemRequestBuilder and sets the default values.
func NewSalesCreditMemoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesCreditMemoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesCreditMemoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesCreditMemos for financials
func (m *SalesCreditMemoItemRequestBuilder) CreateDeleteRequestInformation(options *SalesCreditMemoItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get salesCreditMemos from financials
func (m *SalesCreditMemoItemRequestBuilder) CreateGetRequestInformation(options *SalesCreditMemoItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property salesCreditMemos in financials
func (m *SalesCreditMemoItemRequestBuilder) CreatePatchRequestInformation(options *SalesCreditMemoItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Currency the currency property
func (m *SalesCreditMemoItemRequestBuilder) Currency()(*i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859.CurrencyRequestBuilder) {
    return i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer the customer property
func (m *SalesCreditMemoItemRequestBuilder) Customer()(*i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189.CustomerRequestBuilder) {
    return i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesCreditMemos for financials
func (m *SalesCreditMemoItemRequestBuilder) Delete(options *SalesCreditMemoItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get salesCreditMemos from financials
func (m *SalesCreditMemoItemRequestBuilder) Get(options *SalesCreditMemoItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable), nil
}
// Patch update the navigation property salesCreditMemos in financials
func (m *SalesCreditMemoItemRequestBuilder) Patch(options *SalesCreditMemoItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PaymentTerm the paymentTerm property
func (m *SalesCreditMemoItemRequestBuilder) PaymentTerm()(*i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412.PaymentTermRequestBuilder) {
    return i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLines the salesCreditMemoLines property
func (m *SalesCreditMemoItemRequestBuilder) SalesCreditMemoLines()(*i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51.SalesCreditMemoLinesRequestBuilder) {
    return i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51.NewSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesCreditMemos.item.salesCreditMemoLines.item collection
func (m *SalesCreditMemoItemRequestBuilder) SalesCreditMemoLinesById(id string)(*i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045.SalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine%2Did"] = id
    }
    return i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045.NewSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
