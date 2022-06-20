package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item/item"
    i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item/account"
)

// SalesCreditMemoLineItemRequestBuilder provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
type SalesCreditMemoLineItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SalesCreditMemoLineItemRequestBuilderGetQueryParameters get salesCreditMemoLines from financials
type SalesCreditMemoLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesCreditMemoLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesCreditMemoLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesCreditMemoLineItemRequestBuilderGetQueryParameters
}
// SalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account the account property
func (m *SalesCreditMemoLineItemRequestBuilder) Account()(*i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65.AccountRequestBuilder) {
    return i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSalesCreditMemoLineItemRequestBuilderInternal instantiates a new SalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewSalesCreditMemoLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesCreditMemoLineItemRequestBuilder) {
    m := &SalesCreditMemoLineItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesCreditMemoLineItemRequestBuilder instantiates a new SalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewSalesCreditMemoLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesCreditMemoLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesCreditMemoLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SalesCreditMemoLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *SalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SalesCreditMemoLineItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoLineFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable), nil
}
// Item the item property
func (m *SalesCreditMemoLineItemRequestBuilder) Item()(*i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66.ItemRequestBuilder) {
    return i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66.NewItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *SalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
