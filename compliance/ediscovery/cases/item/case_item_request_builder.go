package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reopen"
    i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations"
    i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections"
    i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds"
    i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians"
    ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/settings"
    ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources"
    ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags"
    icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/close"
    ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets"
    i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/noncustodialdatasources/item"
    i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/reviewsets/item"
    i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/item"
    i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/sourcecollections/item"
    i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/tags/item"
    i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/operations/item"
    ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/legalholds/item"
)

// CaseItemRequestBuilder provides operations to manage the cases property of the microsoft.graph.ediscovery.ediscoveryroot entity.
type CaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CaseItemRequestBuilderDeleteOptions options for Delete
type CaseItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CaseItemRequestBuilderGetOptions options for Get
type CaseItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CaseItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CaseItemRequestBuilderGetQueryParameters get cases from compliance
type CaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CaseItemRequestBuilderPatchOptions options for Patch
type CaseItemRequestBuilderPatchOptions struct {
    // 
    Body i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.Case_escapedable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CaseItemRequestBuilder) Close()(*icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013.CloseRequestBuilder) {
    return icbe808dca9f215e46a201507bb9fd3d6682ffd22c4228bd2934fa6d092ec2013.NewCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCaseItemRequestBuilderInternal instantiates a new CaseItemRequestBuilder and sets the default values.
func NewCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseItemRequestBuilder) {
    m := &CaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCaseItemRequestBuilder instantiates a new CaseItemRequestBuilder and sets the default values.
func NewCaseItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cases for compliance
func (m *CaseItemRequestBuilder) CreateDeleteRequestInformation(options *CaseItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get cases from compliance
func (m *CaseItemRequestBuilder) CreateGetRequestInformation(options *CaseItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cases in compliance
func (m *CaseItemRequestBuilder) CreatePatchRequestInformation(options *CaseItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CaseItemRequestBuilder) Custodians()(*i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a.CustodiansRequestBuilder) {
    return i7ee9cd47ece1636db5f23d48461a1f54657a90bc6e0ee26ce62b3d5e84e35d9a.NewCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodiansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.custodians.item collection
func (m *CaseItemRequestBuilder) CustodiansById(id string)(*i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6.CustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["custodian_id"] = id
    }
    return i6f12c740d8ad4a07e1c35e99570b5c006a24cd6710fc2378825d45f2cebe46e6.NewCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property cases for compliance
func (m *CaseItemRequestBuilder) Delete(options *CaseItemRequestBuilderDeleteOptions)(error) {
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
// Get get cases from compliance
func (m *CaseItemRequestBuilder) Get(options *CaseItemRequestBuilderGetOptions)(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.Case_escapedable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.CreateCase_escapedFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.Case_escapedable), nil
}
func (m *CaseItemRequestBuilder) LegalHolds()(*i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f.LegalHoldsRequestBuilder) {
    return i7e8c62f91304400acff09ffefec4fe5fc66f1640b9a230f33487eb271fea144f.NewLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LegalHoldsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.legalHolds.item collection
func (m *CaseItemRequestBuilder) LegalHoldsById(id string)(*ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36.LegalHoldItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["legalHold_id"] = id
    }
    return ia8b23b6ab36c2b10209f024c3f8396df9a37f4b7c7dbf0cdcc02383064607d36.NewLegalHoldItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) NoncustodialDataSources()(*ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea.NoncustodialDataSourcesRequestBuilder) {
    return ibce968213312dcf70b6cebf357574ebaff57289bb8d2d5537d0fd2e4c8f604ea.NewNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialDataSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.noncustodialDataSources.item collection
func (m *CaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849.NoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["noncustodialDataSource_id"] = id
    }
    return i2d904c4ba0f64303f7ec694166a350f070305a6d3282664966c3c7ee03bee849.NewNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) Operations()(*i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955.OperationsRequestBuilder) {
    return i2ba9e3b5dde4370f9d4ee21032c5d1cf837c8172fdd1c29a522a55e73f42e955.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.operations.item collection
func (m *CaseItemRequestBuilder) OperationsById(id string)(*i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df.CaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation_id"] = id
    }
    return i85cba7b9ace1e93810a0e9c5210f57fed67ca0ac127d57c09a74eadcbfe393df.NewCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property cases in compliance
func (m *CaseItemRequestBuilder) Patch(options *CaseItemRequestBuilderPatchOptions)(error) {
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
func (m *CaseItemRequestBuilder) Reopen()(*i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da.ReopenRequestBuilder) {
    return i01f2133292f0b08e5c7ad62ad1a919908e250568ad80f9fa7990e01a1a4675da.NewReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) ReviewSets()(*ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64.ReviewSetsRequestBuilder) {
    return ie37f7b6846339f3be72955eebac7b1a5f9a9b0e6768d7276f5894199b0fc4e64.NewReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.reviewSets.item collection
func (m *CaseItemRequestBuilder) ReviewSetsById(id string)(*i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705.ReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["reviewSet_id"] = id
    }
    return i6c4f9c432182d1a8383e210ea7282204df1fd76b96287e8d3e0ac99ac798c705.NewReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) Settings()(*ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3.SettingsRequestBuilder) {
    return ib38135aba802e63a95783cb40679431dfd8b0b006b136ce2cf392d13cbcf39b3.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) SourceCollections()(*i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b.SourceCollectionsRequestBuilder) {
    return i5a4461258b51ae222704978e5dc6e8b0f90af958cbbfe0291aa23989e511468b.NewSourceCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SourceCollectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.sourceCollections.item collection
func (m *CaseItemRequestBuilder) SourceCollectionsById(id string)(*i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351.SourceCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sourceCollection_id"] = id
    }
    return i700c8523a708bdda9b2320cf511ec749704101f5da23584bcca53db2919ad351.NewSourceCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CaseItemRequestBuilder) Tags()(*ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b.TagsRequestBuilder) {
    return ibf27493a3efa3e33589ec8953bae196df0f5b31cd7f0010efe7b3550bdd8d64b.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.compliance.ediscovery.cases.item.tags.item collection
func (m *CaseItemRequestBuilder) TagsById(id string)(*i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7.TagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tag_id"] = id
    }
    return i71c6b856224a9a329639e731310314e9aaa3c89985c5bf77e3adab5540d605e7.NewTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
