package customer

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ia27fcb608bc27e619a6746e06fc2099e02761ad0c68657dab8478355dc20c1b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/paymentmethod"
    ib32dc3d0cca5f019b1f5cb3ebd8b343ddcd21728340a514fe503ec7e62407578 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/picture"
    ic78d20d5074f96f1a460b92d608484f5ba5fdbab3e5dd4a7daa053161effdcd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/currency"
    ie5209b6c64f31dc1e00d82a43d68749014c830ea0d76d4202f27060c5fb24d3b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/paymentterm"
    if405c3fece48b6f107df71d8ca78bf7a71025d7ada612a47debd18bcf8c417f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/shipmentmethod"
    ib2cf02bd35e7425c02697fcbb126f5deeca5bc9bb4563905d62f531e12d40011 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer/picture/item"
)

// CustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.salesCreditMemo entity.
type CustomerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CustomerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CustomerRequestBuilderGetQueryParameters get customer from financials
type CustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustomerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomerRequestBuilderGetQueryParameters
}
// CustomerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustomerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCustomerRequestBuilderInternal instantiates a new CustomerRequestBuilder and sets the default values.
func NewCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomerRequestBuilder) {
    m := &CustomerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/customer{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomerRequestBuilder instantiates a new CustomerRequestBuilder and sets the default values.
func NewCustomerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customer for financials
func (m *CustomerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get customer from financials
func (m *CustomerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customer in financials
func (m *CustomerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) Currency()(*ic78d20d5074f96f1a460b92d608484f5ba5fdbab3e5dd4a7daa053161effdcd4.CurrencyRequestBuilder) {
    return ic78d20d5074f96f1a460b92d608484f5ba5fdbab3e5dd4a7daa053161effdcd4.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customer for financials
func (m *CustomerRequestBuilder) Delete(ctx context.Context, requestConfiguration *CustomerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get customer from financials
func (m *CustomerRequestBuilder) Get(ctx context.Context, requestConfiguration *CustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customer in financials
func (m *CustomerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CustomerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PaymentMethod provides operations to manage the paymentMethod property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) PaymentMethod()(*ia27fcb608bc27e619a6746e06fc2099e02761ad0c68657dab8478355dc20c1b5.PaymentMethodRequestBuilder) {
    return ia27fcb608bc27e619a6746e06fc2099e02761ad0c68657dab8478355dc20c1b5.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) PaymentTerm()(*ie5209b6c64f31dc1e00d82a43d68749014c830ea0d76d4202f27060c5fb24d3b.PaymentTermRequestBuilder) {
    return ie5209b6c64f31dc1e00d82a43d68749014c830ea0d76d4202f27060c5fb24d3b.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Picture provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) Picture()(*ib32dc3d0cca5f019b1f5cb3ebd8b343ddcd21728340a514fe503ec7e62407578.PictureRequestBuilder) {
    return ib32dc3d0cca5f019b1f5cb3ebd8b343ddcd21728340a514fe503ec7e62407578.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) PictureById(id string)(*ib2cf02bd35e7425c02697fcbb126f5deeca5bc9bb4563905d62f531e12d40011.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return ib2cf02bd35e7425c02697fcbb126f5deeca5bc9bb4563905d62f531e12d40011.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.customer entity.
func (m *CustomerRequestBuilder) ShipmentMethod()(*if405c3fece48b6f107df71d8ca78bf7a71025d7ada612a47debd18bcf8c417f7.ShipmentMethodRequestBuilder) {
    return if405c3fece48b6f107df71d8ca78bf7a71025d7ada612a47debd18bcf8c417f7.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
