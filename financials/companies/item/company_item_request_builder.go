package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/paymentmethods"
    i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotelines"
    i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes"
    i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorderlines"
    i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoicelines"
    i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/currencies"
    i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/companyinformation"
    i4daeb4853bc59f25aaa8ca4b3cb947d07e94af7f8f723a94381dab39c2e03c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals"
    i57a164edcc262a1e7d0cc54eccb32cb5b3aba7f6bc58654ca5407f5ea88e4560 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers"
    i5850b0e5e17f88c9d6f7fe7ed58df91a7a58ed68a07fe2e45383c1f95826ebb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos"
    i5cf0b22deb67122c3b9b804777f4654610ea33f2f2622da90c3c5f0692f2c11a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoicelines"
    i5e18d51e07e898798b6ff5aee9efe5bacc623efdbcc4f25783d9b6d6c244fd86 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/picture"
    i5f4c33d5e43228d9388a0cc80ea3579373e6ec1fee4977e1a4307035fa641dc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensionvalues"
    i6945f10b9fe3e8587a632f884b54c8c2c422de572bc738b4006fa3f708463503 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/generalledgerentries"
    i7d7d65e7b883cd40a4d185e170461c0df63bc2db571565c4b08605b6a6b39d56 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/agedaccountspayable"
    i83151930b4fc33055bc41c29656d8bc1ccba155184559175e8523e6fa2611eec "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/taxgroups"
    i85f9226d91c649fe1e2b3042eb0e0cb9f28e4ddf039a57a4cf3ca7ff0cbedda1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemolines"
    i9c4b4f369d684a41d4d589022b498cbd7f033b11c2ff47ca608bd2038ec03f7a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/taxareas"
    i9d6749cb386b81deff47b86d48c6a76416836b8b4fdee0aeec53545b4bc5c102 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/items"
    ia4bfc6e27a07d01219c32a3c969a7dd3bd15b29d5dd4888d96948af8c47938a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpaymentjournals"
    ia59c5b754460daf1c032eae861b2a2a0cedc8a8b7ae80231abde5c98df3e9e78 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpayments"
    ia681f3c1a272d62548404c0148103d3f45c07c3d2d28c94b24c5faf621ca9eea "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/itemcategories"
    iae6b3d23c80fa7919b43dad2eea2141b6fe7958e6b757daeb35531a267b1e267 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/paymentterms"
    ib2010f51431a20b492efa7c1681142b3646e7c6454fce979a6dd8919a7c8f04b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices"
    ib2dabce3215c46b19ea19a0fbff4f14d66ae9488d324c78fcb154a217eb7073c "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/employees"
    ibc26fe33984b8d2f64068005ad1791207561cd451ba259344b2ffa77c27c8ba0 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors"
    ibe5eb65fe22e268725128bcf2845841605d9d1eef38f8068c89a3ccafdcb245d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/accounts"
    ic2c6ec504e40a95618e6ce500464a091dc669eea42013d30fea10253e31a14bb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/unitsofmeasure"
    icb4ac224f7dda176b38d0f2720a5c09dc25b1e4000589439ed85ea14acbf34b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/agedaccountsreceivable"
    idbb5b19dc644044a2e4208047f31e9217b3185830f415d7189bb7a15e8dd4e7b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/countriesregions"
    ie68804c62ab9b670efbb79712b214b095747ad5d0970b9af8547e7515614ccce "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/shipmentmethods"
    iec2f8a5133e2722b7ee0b2f37cee34f1bbe214b8e5b888c82435c03e287828fa "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journallines"
    if043e0c78816cb07629640b9329b95ecb1c729c74cef01c31dc93f710b016013 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders"
    if536310c607589fc4a894ef1c998601061419a56a998c732127ea6b48db63d8c "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensions"
    ifa3c2f447ab84ba49123c2e10b5ab6f09f9a2fc5bfcf17b38e86f5b16ca24977 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices"
    i01ad7e647f48e5295cfff5c82ccc53fd72b8ef5ab7d7e8bd7198d016f9c95a92 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/agedaccountspayable/item"
    i0a22f9350aef914933c2bf7011f507fef010436ad65aed8403b06d3550b3f89d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item"
    i10be44ffc518b9ffebfc696f5c34a3302d87645d22fc12694c9904aca1044d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/currencies/item"
    i11e6e00b839a93742244aa3549a6868991eeeed6ccefdbcf294e91fbbee4ff3a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensionvalues/item"
    i156a2925868947a7298582f37c6bf9442997ed01292975d5f12ae6a54051c565 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpayments/item"
    i1659c475e4793613cba0cab9ebd1d7cb4807826d1a70462923f4606aefaf2923 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/employees/item"
    i1e5b9e1d2db51439f7c83d74f5933dd50bab82a742de2085d38ded8ed8ba25db "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journallines/item"
    i2fa657acc59d6b7630ef46de77a50036226d4b074450ccaebde8cf00abc68dec "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpaymentjournals/item"
    i301463334e1dbc7c5c31947d397ff9c42f6aa8bb9cd25e2f928c02448534b4ce "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/items/item"
    i35979601231c435e27fd66cdf7e94aa3ff2c7ed6e60215138c38dddf7d8c728a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/companyinformation/item"
    i390a707b0da26114fde002ae7c016ee59e47ebd9deaea17652eb6719aaa75ee4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/taxgroups/item"
    i3978e2784e815c55b4d52990268524305c63abc98ce85a086e22764f5fcffda4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item"
    i3d03e8e7981c4382b0b7fec2f585dcb2af2df5976ea20ae1de1c5ef9996f16ae "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/generalledgerentries/item"
    i438a48ca92e539d9b6fb008ddf3127a7f528133e671c42432929f28d397e72b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/agedaccountsreceivable/item"
    i43b91397b09aeedb6d1e1bcb250e49ac7a93101a0d0ba2122cdf554aa68abba8 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/unitsofmeasure/item"
    i4c029305791a428f4d91149325952557c61cb631a48256abdb31d83808246516 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item"
    i58aeaf7a36a4d971a858c890c12bf007187d01813c9ad9fd928e8a479673214a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/shipmentmethods/item"
    i5b30584ee167dc3eb4b88ff0d3279d8a5281517c79e75344ae87f0c6fe4c7c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotelines/item"
    i5e7a3640c332a9cd86f5c015ca58ce4c450090db118feb1bc32ead77f958303f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors/item"
    i76069eb1d0c9bbb6e1ddd559e78b0833a04dd47f1a8d9e0b257bb372e8385bbe "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/paymentmethods/item"
    i7cb486a3270e9b712ce5b39298bcbfc2b57bf6b95bb368048dafb6ec20554fdc "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item"
    i86e272c5906623aaa71eb00713179597400af0b135903dd413a7d82af37200de "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/itemcategories/item"
    i8bb41f9aaafe2872cf4a7f16919a0ea7a6265e4e8c2dde4370cbb7a491e99b17 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/accounts/item"
    i9846434aac0b13311c51deaa6afdfc6e1fab10eda8f3d43f1bb2854df18aee02 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorderlines/item"
    i9c35c34fc619450ab198d055e4b3cf4a1cab7f96ff4114906a9df998e1d21369 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoicelines/item"
    ia5d480bc2a01a45ad601eaa1115f60ec63562dff910b91f3deaa3c8f23c43dba "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item"
    iac037dfe92e14a8772fe5c3815a269ac463e8c4bca84adc7dbd464941fbe1577 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item"
    iadc383c56ae6d955e657e8f3c19445f8290a6719a267f96792db35bd6b54360b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/paymentterms/item"
    ib4173258919a2470076856e048a74747e88c26282e025907c42dbae5ec6030ca "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/countriesregions/item"
    ib663fe9a724ce0e5ad43315594f7b7c19a9e8e35fe77f81dc1945c6e3936fe2f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/dimensions/item"
    ic1d69dea9cc76c8f4fc048660687abdf627bd9aa0c48a4b0f12a4f51e369b3e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoicelines/item"
    id553a6763a1078b0430a865a003eb6fc6d871225b59b6baa3a1e8c60e21486f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/taxareas/item"
    ie240be3e672744bf8362b6142f61189bc65992996aa5e71bf520255c5fa98e0a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/picture/item"
    ie72c90b132e5b7adcd798955c629efd15447d563e9da1b63bdf14accfee955b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/journals/item"
    ie940ae608019de62a242d937e70fdc6a489372a4d225a99b69483a2d8f07d28f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemolines/item"
)

// CompanyItemRequestBuilder provides operations to manage the companies property of the microsoft.graph.financials entity.
type CompanyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CompanyItemRequestBuilderGetQueryParameters get companies from financials
type CompanyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompanyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompanyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompanyItemRequestBuilderGetQueryParameters
}
// Accounts provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Accounts()(*ibe5eb65fe22e268725128bcf2845841605d9d1eef38f8068c89a3ccafdcb245d.AccountsRequestBuilder) {
    return ibe5eb65fe22e268725128bcf2845841605d9d1eef38f8068c89a3ccafdcb245d.NewAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountsById provides operations to manage the accounts property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) AccountsById(id string)(*i8bb41f9aaafe2872cf4a7f16919a0ea7a6265e4e8c2dde4370cbb7a491e99b17.AccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account%2Did"] = id
    }
    return i8bb41f9aaafe2872cf4a7f16919a0ea7a6265e4e8c2dde4370cbb7a491e99b17.NewAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgedAccountsPayable provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) AgedAccountsPayable()(*i7d7d65e7b883cd40a4d185e170461c0df63bc2db571565c4b08605b6a6b39d56.AgedAccountsPayableRequestBuilder) {
    return i7d7d65e7b883cd40a4d185e170461c0df63bc2db571565c4b08605b6a6b39d56.NewAgedAccountsPayableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsPayableById provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) AgedAccountsPayableById(id string)(*i01ad7e647f48e5295cfff5c82ccc53fd72b8ef5ab7d7e8bd7198d016f9c95a92.AgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsPayable%2Did"] = id
    }
    return i01ad7e647f48e5295cfff5c82ccc53fd72b8ef5ab7d7e8bd7198d016f9c95a92.NewAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgedAccountsReceivable provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) AgedAccountsReceivable()(*icb4ac224f7dda176b38d0f2720a5c09dc25b1e4000589439ed85ea14acbf34b7.AgedAccountsReceivableRequestBuilder) {
    return icb4ac224f7dda176b38d0f2720a5c09dc25b1e4000589439ed85ea14acbf34b7.NewAgedAccountsReceivableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsReceivableById provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) AgedAccountsReceivableById(id string)(*i438a48ca92e539d9b6fb008ddf3127a7f528133e671c42432929f28d397e72b3.AgedAccountsReceivableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsReceivable%2Did"] = id
    }
    return i438a48ca92e539d9b6fb008ddf3127a7f528133e671c42432929f28d397e72b3.NewAgedAccountsReceivableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompanyInformation provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CompanyInformation()(*i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18.CompanyInformationRequestBuilder) {
    return i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18.NewCompanyInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompanyInformationById provides operations to manage the companyInformation property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CompanyInformationById(id string)(*i35979601231c435e27fd66cdf7e94aa3ff2c7ed6e60215138c38dddf7d8c728a.CompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["companyInformation%2Did"] = id
    }
    return i35979601231c435e27fd66cdf7e94aa3ff2c7ed6e60215138c38dddf7d8c728a.NewCompanyInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCompanyItemRequestBuilderInternal instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompanyItemRequestBuilder) {
    m := &CompanyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCompanyItemRequestBuilder instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompanyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompanyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CountriesRegions provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CountriesRegions()(*idbb5b19dc644044a2e4208047f31e9217b3185830f415d7189bb7a15e8dd4e7b.CountriesRegionsRequestBuilder) {
    return idbb5b19dc644044a2e4208047f31e9217b3185830f415d7189bb7a15e8dd4e7b.NewCountriesRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CountriesRegionsById provides operations to manage the countriesRegions property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CountriesRegionsById(id string)(*ib4173258919a2470076856e048a74747e88c26282e025907c42dbae5ec6030ca.CountryRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["countryRegion%2Did"] = id
    }
    return ib4173258919a2470076856e048a74747e88c26282e025907c42dbae5ec6030ca.NewCountryRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation get companies from financials
func (m *CompanyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CompanyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currencies provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Currencies()(*i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732.CurrenciesRequestBuilder) {
    return i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732.NewCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CurrenciesById provides operations to manage the currencies property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CurrenciesById(id string)(*i10be44ffc518b9ffebfc696f5c34a3302d87645d22fc12694c9904aca1044d58.CurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["currency%2Did"] = id
    }
    return i10be44ffc518b9ffebfc696f5c34a3302d87645d22fc12694c9904aca1044d58.NewCurrencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomerPaymentJournals provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CustomerPaymentJournals()(*ia4bfc6e27a07d01219c32a3c969a7dd3bd15b29d5dd4888d96948af8c47938a4.CustomerPaymentJournalsRequestBuilder) {
    return ia4bfc6e27a07d01219c32a3c969a7dd3bd15b29d5dd4888d96948af8c47938a4.NewCustomerPaymentJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentJournalsById provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CustomerPaymentJournalsById(id string)(*i2fa657acc59d6b7630ef46de77a50036226d4b074450ccaebde8cf00abc68dec.CustomerPaymentJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPaymentJournal%2Did"] = id
    }
    return i2fa657acc59d6b7630ef46de77a50036226d4b074450ccaebde8cf00abc68dec.NewCustomerPaymentJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CustomerPayments()(*ia59c5b754460daf1c032eae861b2a2a0cedc8a8b7ae80231abde5c98df3e9e78.CustomerPaymentsRequestBuilder) {
    return ia59c5b754460daf1c032eae861b2a2a0cedc8a8b7ae80231abde5c98df3e9e78.NewCustomerPaymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentsById provides operations to manage the customerPayments property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CustomerPaymentsById(id string)(*i156a2925868947a7298582f37c6bf9442997ed01292975d5f12ae6a54051c565.CustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment%2Did"] = id
    }
    return i156a2925868947a7298582f37c6bf9442997ed01292975d5f12ae6a54051c565.NewCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Customers provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Customers()(*i57a164edcc262a1e7d0cc54eccb32cb5b3aba7f6bc58654ca5407f5ea88e4560.CustomersRequestBuilder) {
    return i57a164edcc262a1e7d0cc54eccb32cb5b3aba7f6bc58654ca5407f5ea88e4560.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById provides operations to manage the customers property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) CustomersById(id string)(*i0a22f9350aef914933c2bf7011f507fef010436ad65aed8403b06d3550b3f89d.CustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customer%2Did"] = id
    }
    return i0a22f9350aef914933c2bf7011f507fef010436ad65aed8403b06d3550b3f89d.NewCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Dimensions provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Dimensions()(*if536310c607589fc4a894ef1c998601061419a56a998c732127ea6b48db63d8c.DimensionsRequestBuilder) {
    return if536310c607589fc4a894ef1c998601061419a56a998c732127ea6b48db63d8c.NewDimensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionsById provides operations to manage the dimensions property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) DimensionsById(id string)(*ib663fe9a724ce0e5ad43315594f7b7c19a9e8e35fe77f81dc1945c6e3936fe2f.DimensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimension%2Did"] = id
    }
    return ib663fe9a724ce0e5ad43315594f7b7c19a9e8e35fe77f81dc1945c6e3936fe2f.NewDimensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DimensionValues provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) DimensionValues()(*i5f4c33d5e43228d9388a0cc80ea3579373e6ec1fee4977e1a4307035fa641dc8.DimensionValuesRequestBuilder) {
    return i5f4c33d5e43228d9388a0cc80ea3579373e6ec1fee4977e1a4307035fa641dc8.NewDimensionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionValuesById provides operations to manage the dimensionValues property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) DimensionValuesById(id string)(*i11e6e00b839a93742244aa3549a6868991eeeed6ccefdbcf294e91fbbee4ff3a.DimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue%2Did"] = id
    }
    return i11e6e00b839a93742244aa3549a6868991eeeed6ccefdbcf294e91fbbee4ff3a.NewDimensionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Employees provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Employees()(*ib2dabce3215c46b19ea19a0fbff4f14d66ae9488d324c78fcb154a217eb7073c.EmployeesRequestBuilder) {
    return ib2dabce3215c46b19ea19a0fbff4f14d66ae9488d324c78fcb154a217eb7073c.NewEmployeesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmployeesById provides operations to manage the employees property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) EmployeesById(id string)(*i1659c475e4793613cba0cab9ebd1d7cb4807826d1a70462923f4606aefaf2923.EmployeeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["employee%2Did"] = id
    }
    return i1659c475e4793613cba0cab9ebd1d7cb4807826d1a70462923f4606aefaf2923.NewEmployeeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GeneralLedgerEntries provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) GeneralLedgerEntries()(*i6945f10b9fe3e8587a632f884b54c8c2c422de572bc738b4006fa3f708463503.GeneralLedgerEntriesRequestBuilder) {
    return i6945f10b9fe3e8587a632f884b54c8c2c422de572bc738b4006fa3f708463503.NewGeneralLedgerEntriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GeneralLedgerEntriesById provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) GeneralLedgerEntriesById(id string)(*i3d03e8e7981c4382b0b7fec2f585dcb2af2df5976ea20ae1de1c5ef9996f16ae.GeneralLedgerEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["generalLedgerEntry%2Did"] = id
    }
    return i3d03e8e7981c4382b0b7fec2f585dcb2af2df5976ea20ae1de1c5ef9996f16ae.NewGeneralLedgerEntryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get companies from financials
func (m *CompanyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompanyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Companyable), nil
}
// ItemCategories provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) ItemCategories()(*ia681f3c1a272d62548404c0148103d3f45c07c3d2d28c94b24c5faf621ca9eea.ItemCategoriesRequestBuilder) {
    return ia681f3c1a272d62548404c0148103d3f45c07c3d2d28c94b24c5faf621ca9eea.NewItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemCategoriesById provides operations to manage the itemCategories property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) ItemCategoriesById(id string)(*i86e272c5906623aaa71eb00713179597400af0b135903dd413a7d82af37200de.ItemCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemCategory%2Did"] = id
    }
    return i86e272c5906623aaa71eb00713179597400af0b135903dd413a7d82af37200de.NewItemCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Items provides operations to manage the items property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Items()(*i9d6749cb386b81deff47b86d48c6a76416836b8b4fdee0aeec53545b4bc5c102.ItemsRequestBuilder) {
    return i9d6749cb386b81deff47b86d48c6a76416836b8b4fdee0aeec53545b4bc5c102.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) ItemsById(id string)(*i301463334e1dbc7c5c31947d397ff9c42f6aa8bb9cd25e2f928c02448534b4ce.ItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["item%2Did"] = id
    }
    return i301463334e1dbc7c5c31947d397ff9c42f6aa8bb9cd25e2f928c02448534b4ce.NewItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// JournalLines provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) JournalLines()(*iec2f8a5133e2722b7ee0b2f37cee34f1bbe214b8e5b888c82435c03e287828fa.JournalLinesRequestBuilder) {
    return iec2f8a5133e2722b7ee0b2f37cee34f1bbe214b8e5b888c82435c03e287828fa.NewJournalLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalLinesById provides operations to manage the journalLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) JournalLinesById(id string)(*i1e5b9e1d2db51439f7c83d74f5933dd50bab82a742de2085d38ded8ed8ba25db.JournalLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journalLine%2Did"] = id
    }
    return i1e5b9e1d2db51439f7c83d74f5933dd50bab82a742de2085d38ded8ed8ba25db.NewJournalLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Journals provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Journals()(*i4daeb4853bc59f25aaa8ca4b3cb947d07e94af7f8f723a94381dab39c2e03c86.JournalsRequestBuilder) {
    return i4daeb4853bc59f25aaa8ca4b3cb947d07e94af7f8f723a94381dab39c2e03c86.NewJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalsById provides operations to manage the journals property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) JournalsById(id string)(*ie72c90b132e5b7adcd798955c629efd15447d563e9da1b63bdf14accfee955b3.JournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journal%2Did"] = id
    }
    return ie72c90b132e5b7adcd798955c629efd15447d563e9da1b63bdf14accfee955b3.NewJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PaymentMethods provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PaymentMethods()(*i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c.PaymentMethodsRequestBuilder) {
    return i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c.NewPaymentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentMethodsById provides operations to manage the paymentMethods property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PaymentMethodsById(id string)(*i76069eb1d0c9bbb6e1ddd559e78b0833a04dd47f1a8d9e0b257bb372e8385bbe.PaymentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentMethod%2Did"] = id
    }
    return i76069eb1d0c9bbb6e1ddd559e78b0833a04dd47f1a8d9e0b257bb372e8385bbe.NewPaymentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PaymentTerms provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PaymentTerms()(*iae6b3d23c80fa7919b43dad2eea2141b6fe7958e6b757daeb35531a267b1e267.PaymentTermsRequestBuilder) {
    return iae6b3d23c80fa7919b43dad2eea2141b6fe7958e6b757daeb35531a267b1e267.NewPaymentTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTermsById provides operations to manage the paymentTerms property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PaymentTermsById(id string)(*iadc383c56ae6d955e657e8f3c19445f8290a6719a267f96792db35bd6b54360b.PaymentTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentTerm%2Did"] = id
    }
    return iadc383c56ae6d955e657e8f3c19445f8290a6719a267f96792db35bd6b54360b.NewPaymentTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Picture provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Picture()(*i5e18d51e07e898798b6ff5aee9efe5bacc623efdbcc4f25783d9b6d6c244fd86.PictureRequestBuilder) {
    return i5e18d51e07e898798b6ff5aee9efe5bacc623efdbcc4f25783d9b6d6c244fd86.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById provides operations to manage the picture property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PictureById(id string)(*ie240be3e672744bf8362b6142f61189bc65992996aa5e71bf520255c5fa98e0a.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return ie240be3e672744bf8362b6142f61189bc65992996aa5e71bf520255c5fa98e0a.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PurchaseInvoiceLines provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PurchaseInvoiceLines()(*i5cf0b22deb67122c3b9b804777f4654610ea33f2f2622da90c3c5f0692f2c11a.PurchaseInvoiceLinesRequestBuilder) {
    return i5cf0b22deb67122c3b9b804777f4654610ea33f2f2622da90c3c5f0692f2c11a.NewPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*ic1d69dea9cc76c8f4fc048660687abdf627bd9aa0c48a4b0f12a4f51e369b3e5.PurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine%2Did"] = id
    }
    return ic1d69dea9cc76c8f4fc048660687abdf627bd9aa0c48a4b0f12a4f51e369b3e5.NewPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PurchaseInvoices provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PurchaseInvoices()(*ifa3c2f447ab84ba49123c2e10b5ab6f09f9a2fc5bfcf17b38e86f5b16ca24977.PurchaseInvoicesRequestBuilder) {
    return ifa3c2f447ab84ba49123c2e10b5ab6f09f9a2fc5bfcf17b38e86f5b16ca24977.NewPurchaseInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoicesById provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) PurchaseInvoicesById(id string)(*i7cb486a3270e9b712ce5b39298bcbfc2b57bf6b95bb368048dafb6ec20554fdc.PurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoice%2Did"] = id
    }
    return i7cb486a3270e9b712ce5b39298bcbfc2b57bf6b95bb368048dafb6ec20554fdc.NewPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesCreditMemoLines()(*i85f9226d91c649fe1e2b3042eb0e0cb9f28e4ddf039a57a4cf3ca7ff0cbedda1.SalesCreditMemoLinesRequestBuilder) {
    return i85f9226d91c649fe1e2b3042eb0e0cb9f28e4ddf039a57a4cf3ca7ff0cbedda1.NewSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById provides operations to manage the salesCreditMemoLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesCreditMemoLinesById(id string)(*ie940ae608019de62a242d937e70fdc6a489372a4d225a99b69483a2d8f07d28f.SalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine%2Did"] = id
    }
    return ie940ae608019de62a242d937e70fdc6a489372a4d225a99b69483a2d8f07d28f.NewSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesCreditMemos provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesCreditMemos()(*i5850b0e5e17f88c9d6f7fe7ed58df91a7a58ed68a07fe2e45383c1f95826ebb1.SalesCreditMemosRequestBuilder) {
    return i5850b0e5e17f88c9d6f7fe7ed58df91a7a58ed68a07fe2e45383c1f95826ebb1.NewSalesCreditMemosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemosById provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesCreditMemosById(id string)(*ia5d480bc2a01a45ad601eaa1115f60ec63562dff910b91f3deaa3c8f23c43dba.SalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemo%2Did"] = id
    }
    return ia5d480bc2a01a45ad601eaa1115f60ec63562dff910b91f3deaa3c8f23c43dba.NewSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesInvoiceLines provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesInvoiceLines()(*i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6.SalesInvoiceLinesRequestBuilder) {
    return i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6.NewSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById provides operations to manage the salesInvoiceLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesInvoiceLinesById(id string)(*i9c35c34fc619450ab198d055e4b3cf4a1cab7f96ff4114906a9df998e1d21369.SalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine%2Did"] = id
    }
    return i9c35c34fc619450ab198d055e4b3cf4a1cab7f96ff4114906a9df998e1d21369.NewSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesInvoices provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesInvoices()(*ib2010f51431a20b492efa7c1681142b3646e7c6454fce979a6dd8919a7c8f04b.SalesInvoicesRequestBuilder) {
    return ib2010f51431a20b492efa7c1681142b3646e7c6454fce979a6dd8919a7c8f04b.NewSalesInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoicesById provides operations to manage the salesInvoices property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesInvoicesById(id string)(*i3978e2784e815c55b4d52990268524305c63abc98ce85a086e22764f5fcffda4.SalesInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoice%2Did"] = id
    }
    return i3978e2784e815c55b4d52990268524305c63abc98ce85a086e22764f5fcffda4.NewSalesInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesOrderLines provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesOrderLines()(*i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc.SalesOrderLinesRequestBuilder) {
    return i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc.NewSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById provides operations to manage the salesOrderLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesOrderLinesById(id string)(*i9846434aac0b13311c51deaa6afdfc6e1fab10eda8f3d43f1bb2854df18aee02.SalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine%2Did"] = id
    }
    return i9846434aac0b13311c51deaa6afdfc6e1fab10eda8f3d43f1bb2854df18aee02.NewSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesOrders provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesOrders()(*if043e0c78816cb07629640b9329b95ecb1c729c74cef01c31dc93f710b016013.SalesOrdersRequestBuilder) {
    return if043e0c78816cb07629640b9329b95ecb1c729c74cef01c31dc93f710b016013.NewSalesOrdersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrdersById provides operations to manage the salesOrders property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesOrdersById(id string)(*i4c029305791a428f4d91149325952557c61cb631a48256abdb31d83808246516.SalesOrderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrder%2Did"] = id
    }
    return i4c029305791a428f4d91149325952557c61cb631a48256abdb31d83808246516.NewSalesOrderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesQuoteLines provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesQuoteLines()(*i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24.SalesQuoteLinesRequestBuilder) {
    return i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24.NewSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById provides operations to manage the salesQuoteLines property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesQuoteLinesById(id string)(*i5b30584ee167dc3eb4b88ff0d3279d8a5281517c79e75344ae87f0c6fe4c7c80.SalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine%2Did"] = id
    }
    return i5b30584ee167dc3eb4b88ff0d3279d8a5281517c79e75344ae87f0c6fe4c7c80.NewSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SalesQuotes provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesQuotes()(*i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d.SalesQuotesRequestBuilder) {
    return i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d.NewSalesQuotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuotesById provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) SalesQuotesById(id string)(*iac037dfe92e14a8772fe5c3815a269ac463e8c4bca84adc7dbd464941fbe1577.SalesQuoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuote%2Did"] = id
    }
    return iac037dfe92e14a8772fe5c3815a269ac463e8c4bca84adc7dbd464941fbe1577.NewSalesQuoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethods provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) ShipmentMethods()(*ie68804c62ab9b670efbb79712b214b095747ad5d0970b9af8547e7515614ccce.ShipmentMethodsRequestBuilder) {
    return ie68804c62ab9b670efbb79712b214b095747ad5d0970b9af8547e7515614ccce.NewShipmentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethodsById provides operations to manage the shipmentMethods property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) ShipmentMethodsById(id string)(*i58aeaf7a36a4d971a858c890c12bf007187d01813c9ad9fd928e8a479673214a.ShipmentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shipmentMethod%2Did"] = id
    }
    return i58aeaf7a36a4d971a858c890c12bf007187d01813c9ad9fd928e8a479673214a.NewShipmentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaxAreas provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) TaxAreas()(*i9c4b4f369d684a41d4d589022b498cbd7f033b11c2ff47ca608bd2038ec03f7a.TaxAreasRequestBuilder) {
    return i9c4b4f369d684a41d4d589022b498cbd7f033b11c2ff47ca608bd2038ec03f7a.NewTaxAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxAreasById provides operations to manage the taxAreas property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) TaxAreasById(id string)(*id553a6763a1078b0430a865a003eb6fc6d871225b59b6baa3a1e8c60e21486f7.TaxAreaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxArea%2Did"] = id
    }
    return id553a6763a1078b0430a865a003eb6fc6d871225b59b6baa3a1e8c60e21486f7.NewTaxAreaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaxGroups provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) TaxGroups()(*i83151930b4fc33055bc41c29656d8bc1ccba155184559175e8523e6fa2611eec.TaxGroupsRequestBuilder) {
    return i83151930b4fc33055bc41c29656d8bc1ccba155184559175e8523e6fa2611eec.NewTaxGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxGroupsById provides operations to manage the taxGroups property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) TaxGroupsById(id string)(*i390a707b0da26114fde002ae7c016ee59e47ebd9deaea17652eb6719aaa75ee4.TaxGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxGroup%2Did"] = id
    }
    return i390a707b0da26114fde002ae7c016ee59e47ebd9deaea17652eb6719aaa75ee4.NewTaxGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnitsOfMeasure provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) UnitsOfMeasure()(*ic2c6ec504e40a95618e6ce500464a091dc669eea42013d30fea10253e31a14bb.UnitsOfMeasureRequestBuilder) {
    return ic2c6ec504e40a95618e6ce500464a091dc669eea42013d30fea10253e31a14bb.NewUnitsOfMeasureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnitsOfMeasureById provides operations to manage the unitsOfMeasure property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) UnitsOfMeasureById(id string)(*i43b91397b09aeedb6d1e1bcb250e49ac7a93101a0d0ba2122cdf554aa68abba8.UnitOfMeasureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unitOfMeasure%2Did"] = id
    }
    return i43b91397b09aeedb6d1e1bcb250e49ac7a93101a0d0ba2122cdf554aa68abba8.NewUnitOfMeasureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Vendors provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) Vendors()(*ibc26fe33984b8d2f64068005ad1791207561cd451ba259344b2ffa77c27c8ba0.VendorsRequestBuilder) {
    return ibc26fe33984b8d2f64068005ad1791207561cd451ba259344b2ffa77c27c8ba0.NewVendorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VendorsById provides operations to manage the vendors property of the microsoft.graph.company entity.
func (m *CompanyItemRequestBuilder) VendorsById(id string)(*i5e7a3640c332a9cd86f5c015ca58ce4c450090db118feb1bc32ead77f958303f.VendorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vendor%2Did"] = id
    }
    return i5e7a3640c332a9cd86f5c015ca58ce4c450090db118feb1bc32ead77f958303f.NewVendorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
