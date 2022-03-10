package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/paymentmethods"
    i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotelines"
    i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes"
    i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorderlines"
    i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoicelines"
    i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/currencies"
    i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/companyinformation"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CompanyItemRequestBuilderDeleteOptions options for Delete
type CompanyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CompanyItemRequestBuilderGetOptions options for Get
type CompanyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CompanyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CompanyItemRequestBuilderGetQueryParameters get companies from financials
type CompanyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CompanyItemRequestBuilderPatchOptions options for Patch
type CompanyItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Companyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CompanyItemRequestBuilder) Accounts()(*ibe5eb65fe22e268725128bcf2845841605d9d1eef38f8068c89a3ccafdcb245d.AccountsRequestBuilder) {
    return ibe5eb65fe22e268725128bcf2845841605d9d1eef38f8068c89a3ccafdcb245d.NewAccountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccountsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.accounts.item collection
func (m *CompanyItemRequestBuilder) AccountsById(id string)(*i8bb41f9aaafe2872cf4a7f16919a0ea7a6265e4e8c2dde4370cbb7a491e99b17.AccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["account_id"] = id
    }
    return i8bb41f9aaafe2872cf4a7f16919a0ea7a6265e4e8c2dde4370cbb7a491e99b17.NewAccountItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) AgedAccountsPayable()(*i7d7d65e7b883cd40a4d185e170461c0df63bc2db571565c4b08605b6a6b39d56.AgedAccountsPayableRequestBuilder) {
    return i7d7d65e7b883cd40a4d185e170461c0df63bc2db571565c4b08605b6a6b39d56.NewAgedAccountsPayableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsPayableById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.agedAccountsPayable.item collection
func (m *CompanyItemRequestBuilder) AgedAccountsPayableById(id string)(*i01ad7e647f48e5295cfff5c82ccc53fd72b8ef5ab7d7e8bd7198d016f9c95a92.AgedAccountsPayableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsPayable_id"] = id
    }
    return i01ad7e647f48e5295cfff5c82ccc53fd72b8ef5ab7d7e8bd7198d016f9c95a92.NewAgedAccountsPayableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) AgedAccountsReceivable()(*icb4ac224f7dda176b38d0f2720a5c09dc25b1e4000589439ed85ea14acbf34b7.AgedAccountsReceivableRequestBuilder) {
    return icb4ac224f7dda176b38d0f2720a5c09dc25b1e4000589439ed85ea14acbf34b7.NewAgedAccountsReceivableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgedAccountsReceivableById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.agedAccountsReceivable.item collection
func (m *CompanyItemRequestBuilder) AgedAccountsReceivableById(id string)(*i438a48ca92e539d9b6fb008ddf3127a7f528133e671c42432929f28d397e72b3.AgedAccountsReceivableItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agedAccountsReceivable_id"] = id
    }
    return i438a48ca92e539d9b6fb008ddf3127a7f528133e671c42432929f28d397e72b3.NewAgedAccountsReceivableItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) CompanyInformation()(*i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18.CompanyInformationRequestBuilder) {
    return i4150f690ea3d8f1505214d237ce4d724022a2128b537709f0ee5446c49db5f18.NewCompanyInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompanyInformationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.companyInformation.item collection
func (m *CompanyItemRequestBuilder) CompanyInformationById(id string)(*i35979601231c435e27fd66cdf7e94aa3ff2c7ed6e60215138c38dddf7d8c728a.CompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["companyInformation_id"] = id
    }
    return i35979601231c435e27fd66cdf7e94aa3ff2c7ed6e60215138c38dddf7d8c728a.NewCompanyInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCompanyItemRequestBuilderInternal instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CompanyItemRequestBuilder) {
    m := &CompanyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCompanyItemRequestBuilder instantiates a new CompanyItemRequestBuilder and sets the default values.
func NewCompanyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CompanyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CompanyItemRequestBuilder) CountriesRegions()(*idbb5b19dc644044a2e4208047f31e9217b3185830f415d7189bb7a15e8dd4e7b.CountriesRegionsRequestBuilder) {
    return idbb5b19dc644044a2e4208047f31e9217b3185830f415d7189bb7a15e8dd4e7b.NewCountriesRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CountriesRegionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.countriesRegions.item collection
func (m *CompanyItemRequestBuilder) CountriesRegionsById(id string)(*ib4173258919a2470076856e048a74747e88c26282e025907c42dbae5ec6030ca.CountryRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["countryRegion_id"] = id
    }
    return ib4173258919a2470076856e048a74747e88c26282e025907c42dbae5ec6030ca.NewCountryRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property companies for financials
func (m *CompanyItemRequestBuilder) CreateDeleteRequestInformation(options *CompanyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get companies from financials
func (m *CompanyItemRequestBuilder) CreateGetRequestInformation(options *CompanyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property companies in financials
func (m *CompanyItemRequestBuilder) CreatePatchRequestInformation(options *CompanyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *CompanyItemRequestBuilder) Currencies()(*i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732.CurrenciesRequestBuilder) {
    return i39cfe575a2df965e199497b478ce23488c2f17f202a1ef77b208a12cad55f732.NewCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CurrenciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.currencies.item collection
func (m *CompanyItemRequestBuilder) CurrenciesById(id string)(*i10be44ffc518b9ffebfc696f5c34a3302d87645d22fc12694c9904aca1044d58.CurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["currency_id"] = id
    }
    return i10be44ffc518b9ffebfc696f5c34a3302d87645d22fc12694c9904aca1044d58.NewCurrencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) CustomerPaymentJournals()(*ia4bfc6e27a07d01219c32a3c969a7dd3bd15b29d5dd4888d96948af8c47938a4.CustomerPaymentJournalsRequestBuilder) {
    return ia4bfc6e27a07d01219c32a3c969a7dd3bd15b29d5dd4888d96948af8c47938a4.NewCustomerPaymentJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentJournalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customerPaymentJournals.item collection
func (m *CompanyItemRequestBuilder) CustomerPaymentJournalsById(id string)(*i2fa657acc59d6b7630ef46de77a50036226d4b074450ccaebde8cf00abc68dec.CustomerPaymentJournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPaymentJournal_id"] = id
    }
    return i2fa657acc59d6b7630ef46de77a50036226d4b074450ccaebde8cf00abc68dec.NewCustomerPaymentJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) CustomerPayments()(*ia59c5b754460daf1c032eae861b2a2a0cedc8a8b7ae80231abde5c98df3e9e78.CustomerPaymentsRequestBuilder) {
    return ia59c5b754460daf1c032eae861b2a2a0cedc8a8b7ae80231abde5c98df3e9e78.NewCustomerPaymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customerPayments.item collection
func (m *CompanyItemRequestBuilder) CustomerPaymentsById(id string)(*i156a2925868947a7298582f37c6bf9442997ed01292975d5f12ae6a54051c565.CustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment_id"] = id
    }
    return i156a2925868947a7298582f37c6bf9442997ed01292975d5f12ae6a54051c565.NewCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Customers()(*i57a164edcc262a1e7d0cc54eccb32cb5b3aba7f6bc58654ca5407f5ea88e4560.CustomersRequestBuilder) {
    return i57a164edcc262a1e7d0cc54eccb32cb5b3aba7f6bc58654ca5407f5ea88e4560.NewCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customers.item collection
func (m *CompanyItemRequestBuilder) CustomersById(id string)(*i0a22f9350aef914933c2bf7011f507fef010436ad65aed8403b06d3550b3f89d.CustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customer_id"] = id
    }
    return i0a22f9350aef914933c2bf7011f507fef010436ad65aed8403b06d3550b3f89d.NewCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property companies for financials
func (m *CompanyItemRequestBuilder) Delete(options *CompanyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CompanyItemRequestBuilder) Dimensions()(*if536310c607589fc4a894ef1c998601061419a56a998c732127ea6b48db63d8c.DimensionsRequestBuilder) {
    return if536310c607589fc4a894ef1c998601061419a56a998c732127ea6b48db63d8c.NewDimensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.dimensions.item collection
func (m *CompanyItemRequestBuilder) DimensionsById(id string)(*ib663fe9a724ce0e5ad43315594f7b7c19a9e8e35fe77f81dc1945c6e3936fe2f.DimensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimension_id"] = id
    }
    return ib663fe9a724ce0e5ad43315594f7b7c19a9e8e35fe77f81dc1945c6e3936fe2f.NewDimensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) DimensionValues()(*i5f4c33d5e43228d9388a0cc80ea3579373e6ec1fee4977e1a4307035fa641dc8.DimensionValuesRequestBuilder) {
    return i5f4c33d5e43228d9388a0cc80ea3579373e6ec1fee4977e1a4307035fa641dc8.NewDimensionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DimensionValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.dimensionValues.item collection
func (m *CompanyItemRequestBuilder) DimensionValuesById(id string)(*i11e6e00b839a93742244aa3549a6868991eeeed6ccefdbcf294e91fbbee4ff3a.DimensionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dimensionValue_id"] = id
    }
    return i11e6e00b839a93742244aa3549a6868991eeeed6ccefdbcf294e91fbbee4ff3a.NewDimensionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Employees()(*ib2dabce3215c46b19ea19a0fbff4f14d66ae9488d324c78fcb154a217eb7073c.EmployeesRequestBuilder) {
    return ib2dabce3215c46b19ea19a0fbff4f14d66ae9488d324c78fcb154a217eb7073c.NewEmployeesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmployeesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.employees.item collection
func (m *CompanyItemRequestBuilder) EmployeesById(id string)(*i1659c475e4793613cba0cab9ebd1d7cb4807826d1a70462923f4606aefaf2923.EmployeeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["employee_id"] = id
    }
    return i1659c475e4793613cba0cab9ebd1d7cb4807826d1a70462923f4606aefaf2923.NewEmployeeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) GeneralLedgerEntries()(*i6945f10b9fe3e8587a632f884b54c8c2c422de572bc738b4006fa3f708463503.GeneralLedgerEntriesRequestBuilder) {
    return i6945f10b9fe3e8587a632f884b54c8c2c422de572bc738b4006fa3f708463503.NewGeneralLedgerEntriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GeneralLedgerEntriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.generalLedgerEntries.item collection
func (m *CompanyItemRequestBuilder) GeneralLedgerEntriesById(id string)(*i3d03e8e7981c4382b0b7fec2f585dcb2af2df5976ea20ae1de1c5ef9996f16ae.GeneralLedgerEntryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["generalLedgerEntry_id"] = id
    }
    return i3d03e8e7981c4382b0b7fec2f585dcb2af2df5976ea20ae1de1c5ef9996f16ae.NewGeneralLedgerEntryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get companies from financials
func (m *CompanyItemRequestBuilder) Get(options *CompanyItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Companyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCompanyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Companyable), nil
}
func (m *CompanyItemRequestBuilder) ItemCategories()(*ia681f3c1a272d62548404c0148103d3f45c07c3d2d28c94b24c5faf621ca9eea.ItemCategoriesRequestBuilder) {
    return ia681f3c1a272d62548404c0148103d3f45c07c3d2d28c94b24c5faf621ca9eea.NewItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.itemCategories.item collection
func (m *CompanyItemRequestBuilder) ItemCategoriesById(id string)(*i86e272c5906623aaa71eb00713179597400af0b135903dd413a7d82af37200de.ItemCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemCategory_id"] = id
    }
    return i86e272c5906623aaa71eb00713179597400af0b135903dd413a7d82af37200de.NewItemCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Items()(*i9d6749cb386b81deff47b86d48c6a76416836b8b4fdee0aeec53545b4bc5c102.ItemsRequestBuilder) {
    return i9d6749cb386b81deff47b86d48c6a76416836b8b4fdee0aeec53545b4bc5c102.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.items.item collection
func (m *CompanyItemRequestBuilder) ItemsById(id string)(*i301463334e1dbc7c5c31947d397ff9c42f6aa8bb9cd25e2f928c02448534b4ce.ItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["item_id"] = id
    }
    return i301463334e1dbc7c5c31947d397ff9c42f6aa8bb9cd25e2f928c02448534b4ce.NewItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) JournalLines()(*iec2f8a5133e2722b7ee0b2f37cee34f1bbe214b8e5b888c82435c03e287828fa.JournalLinesRequestBuilder) {
    return iec2f8a5133e2722b7ee0b2f37cee34f1bbe214b8e5b888c82435c03e287828fa.NewJournalLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.journalLines.item collection
func (m *CompanyItemRequestBuilder) JournalLinesById(id string)(*i1e5b9e1d2db51439f7c83d74f5933dd50bab82a742de2085d38ded8ed8ba25db.JournalLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journalLine_id"] = id
    }
    return i1e5b9e1d2db51439f7c83d74f5933dd50bab82a742de2085d38ded8ed8ba25db.NewJournalLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Journals()(*i4daeb4853bc59f25aaa8ca4b3cb947d07e94af7f8f723a94381dab39c2e03c86.JournalsRequestBuilder) {
    return i4daeb4853bc59f25aaa8ca4b3cb947d07e94af7f8f723a94381dab39c2e03c86.NewJournalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JournalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.journals.item collection
func (m *CompanyItemRequestBuilder) JournalsById(id string)(*ie72c90b132e5b7adcd798955c629efd15447d563e9da1b63bdf14accfee955b3.JournalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["journal_id"] = id
    }
    return ie72c90b132e5b7adcd798955c629efd15447d563e9da1b63bdf14accfee955b3.NewJournalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property companies in financials
func (m *CompanyItemRequestBuilder) Patch(options *CompanyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CompanyItemRequestBuilder) PaymentMethods()(*i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c.PaymentMethodsRequestBuilder) {
    return i0b33ef160f8fead222ec024e8840780deb4ab2dbfa92a870ae77d1cc4968234c.NewPaymentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.paymentMethods.item collection
func (m *CompanyItemRequestBuilder) PaymentMethodsById(id string)(*i76069eb1d0c9bbb6e1ddd559e78b0833a04dd47f1a8d9e0b257bb372e8385bbe.PaymentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentMethod_id"] = id
    }
    return i76069eb1d0c9bbb6e1ddd559e78b0833a04dd47f1a8d9e0b257bb372e8385bbe.NewPaymentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) PaymentTerms()(*iae6b3d23c80fa7919b43dad2eea2141b6fe7958e6b757daeb35531a267b1e267.PaymentTermsRequestBuilder) {
    return iae6b3d23c80fa7919b43dad2eea2141b6fe7958e6b757daeb35531a267b1e267.NewPaymentTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTermsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.paymentTerms.item collection
func (m *CompanyItemRequestBuilder) PaymentTermsById(id string)(*iadc383c56ae6d955e657e8f3c19445f8290a6719a267f96792db35bd6b54360b.PaymentTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["paymentTerm_id"] = id
    }
    return iadc383c56ae6d955e657e8f3c19445f8290a6719a267f96792db35bd6b54360b.NewPaymentTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Picture()(*i5e18d51e07e898798b6ff5aee9efe5bacc623efdbcc4f25783d9b6d6c244fd86.PictureRequestBuilder) {
    return i5e18d51e07e898798b6ff5aee9efe5bacc623efdbcc4f25783d9b6d6c244fd86.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.picture.item collection
func (m *CompanyItemRequestBuilder) PictureById(id string)(*ie240be3e672744bf8362b6142f61189bc65992996aa5e71bf520255c5fa98e0a.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return ie240be3e672744bf8362b6142f61189bc65992996aa5e71bf520255c5fa98e0a.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) PurchaseInvoiceLines()(*i5cf0b22deb67122c3b9b804777f4654610ea33f2f2622da90c3c5f0692f2c11a.PurchaseInvoiceLinesRequestBuilder) {
    return i5cf0b22deb67122c3b9b804777f4654610ea33f2f2622da90c3c5f0692f2c11a.NewPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.purchaseInvoiceLines.item collection
func (m *CompanyItemRequestBuilder) PurchaseInvoiceLinesById(id string)(*ic1d69dea9cc76c8f4fc048660687abdf627bd9aa0c48a4b0f12a4f51e369b3e5.PurchaseInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine_id"] = id
    }
    return ic1d69dea9cc76c8f4fc048660687abdf627bd9aa0c48a4b0f12a4f51e369b3e5.NewPurchaseInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) PurchaseInvoices()(*ifa3c2f447ab84ba49123c2e10b5ab6f09f9a2fc5bfcf17b38e86f5b16ca24977.PurchaseInvoicesRequestBuilder) {
    return ifa3c2f447ab84ba49123c2e10b5ab6f09f9a2fc5bfcf17b38e86f5b16ca24977.NewPurchaseInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.purchaseInvoices.item collection
func (m *CompanyItemRequestBuilder) PurchaseInvoicesById(id string)(*i7cb486a3270e9b712ce5b39298bcbfc2b57bf6b95bb368048dafb6ec20554fdc.PurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoice_id"] = id
    }
    return i7cb486a3270e9b712ce5b39298bcbfc2b57bf6b95bb368048dafb6ec20554fdc.NewPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesCreditMemoLines()(*i85f9226d91c649fe1e2b3042eb0e0cb9f28e4ddf039a57a4cf3ca7ff0cbedda1.SalesCreditMemoLinesRequestBuilder) {
    return i85f9226d91c649fe1e2b3042eb0e0cb9f28e4ddf039a57a4cf3ca7ff0cbedda1.NewSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesCreditMemoLines.item collection
func (m *CompanyItemRequestBuilder) SalesCreditMemoLinesById(id string)(*ie940ae608019de62a242d937e70fdc6a489372a4d225a99b69483a2d8f07d28f.SalesCreditMemoLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine_id"] = id
    }
    return ie940ae608019de62a242d937e70fdc6a489372a4d225a99b69483a2d8f07d28f.NewSalesCreditMemoLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesCreditMemos()(*i5850b0e5e17f88c9d6f7fe7ed58df91a7a58ed68a07fe2e45383c1f95826ebb1.SalesCreditMemosRequestBuilder) {
    return i5850b0e5e17f88c9d6f7fe7ed58df91a7a58ed68a07fe2e45383c1f95826ebb1.NewSalesCreditMemosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesCreditMemos.item collection
func (m *CompanyItemRequestBuilder) SalesCreditMemosById(id string)(*ia5d480bc2a01a45ad601eaa1115f60ec63562dff910b91f3deaa3c8f23c43dba.SalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemo_id"] = id
    }
    return ia5d480bc2a01a45ad601eaa1115f60ec63562dff910b91f3deaa3c8f23c43dba.NewSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesInvoiceLines()(*i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6.SalesInvoiceLinesRequestBuilder) {
    return i332f46d9c7722cf45bd813ad625ff9365b88c83dacb95203737d7b2c420620d6.NewSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesInvoiceLines.item collection
func (m *CompanyItemRequestBuilder) SalesInvoiceLinesById(id string)(*i9c35c34fc619450ab198d055e4b3cf4a1cab7f96ff4114906a9df998e1d21369.SalesInvoiceLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine_id"] = id
    }
    return i9c35c34fc619450ab198d055e4b3cf4a1cab7f96ff4114906a9df998e1d21369.NewSalesInvoiceLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesInvoices()(*ib2010f51431a20b492efa7c1681142b3646e7c6454fce979a6dd8919a7c8f04b.SalesInvoicesRequestBuilder) {
    return ib2010f51431a20b492efa7c1681142b3646e7c6454fce979a6dd8919a7c8f04b.NewSalesInvoicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesInvoices.item collection
func (m *CompanyItemRequestBuilder) SalesInvoicesById(id string)(*i3978e2784e815c55b4d52990268524305c63abc98ce85a086e22764f5fcffda4.SalesInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoice_id"] = id
    }
    return i3978e2784e815c55b4d52990268524305c63abc98ce85a086e22764f5fcffda4.NewSalesInvoiceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesOrderLines()(*i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc.SalesOrderLinesRequestBuilder) {
    return i328f372295131e7ef46207b8702c52281f757fc7fe291d1259824e4f4fc924dc.NewSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrderLines.item collection
func (m *CompanyItemRequestBuilder) SalesOrderLinesById(id string)(*i9846434aac0b13311c51deaa6afdfc6e1fab10eda8f3d43f1bb2854df18aee02.SalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine_id"] = id
    }
    return i9846434aac0b13311c51deaa6afdfc6e1fab10eda8f3d43f1bb2854df18aee02.NewSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesOrders()(*if043e0c78816cb07629640b9329b95ecb1c729c74cef01c31dc93f710b016013.SalesOrdersRequestBuilder) {
    return if043e0c78816cb07629640b9329b95ecb1c729c74cef01c31dc93f710b016013.NewSalesOrdersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrdersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrders.item collection
func (m *CompanyItemRequestBuilder) SalesOrdersById(id string)(*i4c029305791a428f4d91149325952557c61cb631a48256abdb31d83808246516.SalesOrderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrder_id"] = id
    }
    return i4c029305791a428f4d91149325952557c61cb631a48256abdb31d83808246516.NewSalesOrderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesQuoteLines()(*i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24.SalesQuoteLinesRequestBuilder) {
    return i103871e8833aa81b6a751678ed468503e289f74fd4334356f22fbe48c0e66b24.NewSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesQuoteLines.item collection
func (m *CompanyItemRequestBuilder) SalesQuoteLinesById(id string)(*i5b30584ee167dc3eb4b88ff0d3279d8a5281517c79e75344ae87f0c6fe4c7c80.SalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine_id"] = id
    }
    return i5b30584ee167dc3eb4b88ff0d3279d8a5281517c79e75344ae87f0c6fe4c7c80.NewSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) SalesQuotes()(*i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d.SalesQuotesRequestBuilder) {
    return i264d8c8ae3f67c790cbd458ad458bb66295f8074cc2a150f8ed31ae9a8eb229d.NewSalesQuotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesQuotes.item collection
func (m *CompanyItemRequestBuilder) SalesQuotesById(id string)(*iac037dfe92e14a8772fe5c3815a269ac463e8c4bca84adc7dbd464941fbe1577.SalesQuoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuote_id"] = id
    }
    return iac037dfe92e14a8772fe5c3815a269ac463e8c4bca84adc7dbd464941fbe1577.NewSalesQuoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) ShipmentMethods()(*ie68804c62ab9b670efbb79712b214b095747ad5d0970b9af8547e7515614ccce.ShipmentMethodsRequestBuilder) {
    return ie68804c62ab9b670efbb79712b214b095747ad5d0970b9af8547e7515614ccce.NewShipmentMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShipmentMethodsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.shipmentMethods.item collection
func (m *CompanyItemRequestBuilder) ShipmentMethodsById(id string)(*i58aeaf7a36a4d971a858c890c12bf007187d01813c9ad9fd928e8a479673214a.ShipmentMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shipmentMethod_id"] = id
    }
    return i58aeaf7a36a4d971a858c890c12bf007187d01813c9ad9fd928e8a479673214a.NewShipmentMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) TaxAreas()(*i9c4b4f369d684a41d4d589022b498cbd7f033b11c2ff47ca608bd2038ec03f7a.TaxAreasRequestBuilder) {
    return i9c4b4f369d684a41d4d589022b498cbd7f033b11c2ff47ca608bd2038ec03f7a.NewTaxAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxAreasById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.taxAreas.item collection
func (m *CompanyItemRequestBuilder) TaxAreasById(id string)(*id553a6763a1078b0430a865a003eb6fc6d871225b59b6baa3a1e8c60e21486f7.TaxAreaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxArea_id"] = id
    }
    return id553a6763a1078b0430a865a003eb6fc6d871225b59b6baa3a1e8c60e21486f7.NewTaxAreaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) TaxGroups()(*i83151930b4fc33055bc41c29656d8bc1ccba155184559175e8523e6fa2611eec.TaxGroupsRequestBuilder) {
    return i83151930b4fc33055bc41c29656d8bc1ccba155184559175e8523e6fa2611eec.NewTaxGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaxGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.taxGroups.item collection
func (m *CompanyItemRequestBuilder) TaxGroupsById(id string)(*i390a707b0da26114fde002ae7c016ee59e47ebd9deaea17652eb6719aaa75ee4.TaxGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taxGroup_id"] = id
    }
    return i390a707b0da26114fde002ae7c016ee59e47ebd9deaea17652eb6719aaa75ee4.NewTaxGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) UnitsOfMeasure()(*ic2c6ec504e40a95618e6ce500464a091dc669eea42013d30fea10253e31a14bb.UnitsOfMeasureRequestBuilder) {
    return ic2c6ec504e40a95618e6ce500464a091dc669eea42013d30fea10253e31a14bb.NewUnitsOfMeasureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnitsOfMeasureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.unitsOfMeasure.item collection
func (m *CompanyItemRequestBuilder) UnitsOfMeasureById(id string)(*i43b91397b09aeedb6d1e1bcb250e49ac7a93101a0d0ba2122cdf554aa68abba8.UnitOfMeasureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unitOfMeasure_id"] = id
    }
    return i43b91397b09aeedb6d1e1bcb250e49ac7a93101a0d0ba2122cdf554aa68abba8.NewUnitOfMeasureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CompanyItemRequestBuilder) Vendors()(*ibc26fe33984b8d2f64068005ad1791207561cd451ba259344b2ffa77c27c8ba0.VendorsRequestBuilder) {
    return ibc26fe33984b8d2f64068005ad1791207561cd451ba259344b2ffa77c27c8ba0.NewVendorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VendorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.vendors.item collection
func (m *CompanyItemRequestBuilder) VendorsById(id string)(*i5e7a3640c332a9cd86f5c015ca58ce4c450090db118feb1bc32ead77f958303f.VendorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vendor_id"] = id
    }
    return i5e7a3640c332a9cd86f5c015ca58ce4c450090db118feb1bc32ead77f958303f.NewVendorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
