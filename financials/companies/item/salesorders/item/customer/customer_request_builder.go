package customer

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/paymentmethod"
    i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/paymentterm"
    i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/picture"
    i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/currency"
    ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/shipmentmethod"
    i9f22a7a8aaef43961b7c1fe0991070f3b76e85698895ef1b6b0eda32f2ef2caa "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/picture/item"
)

// CustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.salesOrder entity.
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
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}/customer{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property customer for financials
func (m *CustomerRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property customer for financials
func (m *CustomerRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration get customer from financials
func (m *CustomerRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get customer from financials
func (m *CustomerRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property customer in financials
func (m *CustomerRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property customer in financials
func (m *CustomerRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Currency the currency property
func (m *CustomerRequestBuilder) Currency()(*i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7.CurrencyRequestBuilder) {
    return i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeleteWithResponseHandler delete navigation property customer for financials
func (m *CustomerRequestBuilder) DeleteWithResponseHandler(requestConfiguration *CustomerRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property customer for financials
func (m *CustomerRequestBuilder) DeleteWithResponseHandler(requestConfiguration *CustomerRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler get customer from financials
func (m *CustomerRequestBuilder) GetWithResponseHandler(requestConfiguration *CustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get customer from financials
func (m *CustomerRequestBuilder) GetWithResponseHandler(requestConfiguration *CustomerRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PatchWithResponseHandler update the navigation property customer in financials
func (m *CustomerRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CustomerRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property customer in financials
func (m *CustomerRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CustomerRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PaymentMethod the paymentMethod property
func (m *CustomerRequestBuilder) PaymentMethod()(*i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e.PaymentMethodRequestBuilder) {
    return i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTerm the paymentTerm property
func (m *CustomerRequestBuilder) PaymentTerm()(*i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1.PaymentTermRequestBuilder) {
    return i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Picture the picture property
func (m *CustomerRequestBuilder) Picture()(*i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.PictureRequestBuilder) {
    return i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrders.item.customer.picture.item collection
func (m *CustomerRequestBuilder) PictureById(id string)(*i9f22a7a8aaef43961b7c1fe0991070f3b76e85698895ef1b6b0eda32f2ef2caa.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return i9f22a7a8aaef43961b7c1fe0991070f3b76e85698895ef1b6b0eda32f2ef2caa.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethod the shipmentMethod property
func (m *CustomerRequestBuilder) ShipmentMethod()(*ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632.ShipmentMethodRequestBuilder) {
    return ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
