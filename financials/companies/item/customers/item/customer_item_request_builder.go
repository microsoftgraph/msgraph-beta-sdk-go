package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentterm"
    i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/picture"
    i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentmethod"
    i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/currency"
    i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/shipmentmethod"
    i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/picture/item"
)

// CustomerItemRequestBuilder provides operations to manage the customers property of the microsoft.graph.company entity.
type CustomerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CustomerItemRequestBuilderDeleteOptions options for Delete
type CustomerItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CustomerItemRequestBuilderGetOptions options for Get
type CustomerItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustomerItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CustomerItemRequestBuilderGetQueryParameters get customers from financials
type CustomerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CustomerItemRequestBuilderPatchOptions options for Patch
type CustomerItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewCustomerItemRequestBuilderInternal instantiates a new CustomerItemRequestBuilder and sets the default values.
func NewCustomerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomerItemRequestBuilder) {
    m := &CustomerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/customers/{customer%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomerItemRequestBuilder instantiates a new CustomerItemRequestBuilder and sets the default values.
func NewCustomerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustomerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customers for financials
func (m *CustomerItemRequestBuilder) CreateDeleteRequestInformation(options *CustomerItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get customers from financials
func (m *CustomerItemRequestBuilder) CreateGetRequestInformation(options *CustomerItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customers in financials
func (m *CustomerItemRequestBuilder) CreatePatchRequestInformation(options *CustomerItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CustomerItemRequestBuilder) Currency()(*i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.CurrencyRequestBuilder) {
    return i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customers for financials
func (m *CustomerItemRequestBuilder) Delete(options *CustomerItemRequestBuilderDeleteOptions)(error) {
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
// Get get customers from financials
func (m *CustomerItemRequestBuilder) Get(options *CustomerItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customers in financials
func (m *CustomerItemRequestBuilder) Patch(options *CustomerItemRequestBuilderPatchOptions)(error) {
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
// PaymentMethod the paymentMethod property
func (m *CustomerItemRequestBuilder) PaymentMethod()(*i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.PaymentMethodRequestBuilder) {
    return i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTerm the paymentTerm property
func (m *CustomerItemRequestBuilder) PaymentTerm()(*i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.PaymentTermRequestBuilder) {
    return i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Picture the picture property
func (m *CustomerItemRequestBuilder) Picture()(*i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.PictureRequestBuilder) {
    return i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customers.item.picture.item collection
func (m *CustomerItemRequestBuilder) PictureById(id string)(*i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethod the shipmentMethod property
func (m *CustomerItemRequestBuilder) ShipmentMethod()(*i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.ShipmentMethodRequestBuilder) {
    return i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
