package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i111108be478a5c65efb26bd98fba3812504bcf09fab98bec275c9bfb036c6e76 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/close"
    i4d72be7dec3a86e988334df48fc2a5a60b7080974462a5c3b6f245eb595546e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians"
    i6a31f80a06495660e3184b68a5c151df035e02a18ac99cfd059d1236423c6758 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/operations"
    i6a6026b47c385592c13309b70526fe02801e107051c3ddd78444baca8e547380 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches"
    i729fc921d5b80a60d8f1623b0a50e0190155d43bb1437138a638905bcd343eb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/tags"
    i85c8e42d130da558580b9a46e52a3077a2dcdfb85df42f3dc3771e5e54494092 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/legalholds"
    i9c4817aecd1c7b6fe339c7ee61c758e77ffb79c46bd4b2c82da6fc313c8559f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reopen"
    iadbc9b9b01faf8819d968fb41fc7969583e10daeca4d6ca26b40b4ce48622a4f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets"
    idd68d7a96e6d4b3e6a6836e7a6e2e4ddfb4db1c4e2cbd40a89f985851fe47c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/settings"
    if55b3465b5d38c970f2ab12abe1f658e1c8eaaba08efeed80e2df798315efd0d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/noncustodialdatasources"
    i07e34a8004c952f926f66ce45ec267fb3e980eebe9e100ac4d9d587e26f4f10e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/noncustodialdatasources/item"
    i165a308de50d35f4c393d5346bd98bf82a032c98d62cf3853110ba65de984b2f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/operations/item"
    i50da0783ccd2492dd0029e547a26126c7a1fa0a518a1a7604e4700345a64d387 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item"
    i7f2c94cccb75af6bdb0124a47a951ae1e97dcddd956b0615ac4e2d6d698ec11b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item"
    i8cab881a3513cc272eec2a3472e709a6bdc7a3974398acf13e4a2281a871cc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/legalholds/item"
    i98eb84e02e294af8e8f8b700991ad14c1cb72bf5681fa872dd1d942a0bca0877 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/tags/item"
    if1715e408bc3525e0632493a6e5ae7e31a098303838372aef02b23c1a27e00f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item"
)

// EdiscoveryCaseItemRequestBuilder provides operations to manage the ediscoveryCases property of the microsoft.graph.security.casesRoot entity.
type EdiscoveryCaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCaseItemRequestBuilderGetQueryParameters get ediscoveryCases from security
type EdiscoveryCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCaseItemRequestBuilderGetQueryParameters
}
// EdiscoveryCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Close the close property
func (m *EdiscoveryCaseItemRequestBuilder) Close()(*i111108be478a5c65efb26bd98fba3812504bcf09fab98bec275c9bfb036c6e76.CloseRequestBuilder) {
    return i111108be478a5c65efb26bd98fba3812504bcf09fab98bec275c9bfb036c6e76.NewCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoveryCaseItemRequestBuilderInternal instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewEdiscoveryCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCaseItemRequestBuilder) {
    m := &EdiscoveryCaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryCaseItemRequestBuilder instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewEdiscoveryCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property ediscoveryCases for security
func (m *EdiscoveryCaseItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get ediscoveryCases from security
func (m *EdiscoveryCaseItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property ediscoveryCases in security
func (m *EdiscoveryCaseItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, requestConfiguration *EdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Custodians the custodians property
func (m *EdiscoveryCaseItemRequestBuilder) Custodians()(*i4d72be7dec3a86e988334df48fc2a5a60b7080974462a5c3b6f245eb595546e4.CustodiansRequestBuilder) {
    return i4d72be7dec3a86e988334df48fc2a5a60b7080974462a5c3b6f245eb595546e4.NewCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodiansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.custodians.item collection
func (m *EdiscoveryCaseItemRequestBuilder) CustodiansById(id string)(*i50da0783ccd2492dd0029e547a26126c7a1fa0a518a1a7604e4700345a64d387.EdiscoveryCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryCustodian%2Did"] = id
    }
    return i50da0783ccd2492dd0029e547a26126c7a1fa0a518a1a7604e4700345a64d387.NewEdiscoveryCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property ediscoveryCases for security
func (m *EdiscoveryCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get ediscoveryCases from security
func (m *EdiscoveryCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable), nil
}
// LegalHolds the legalHolds property
func (m *EdiscoveryCaseItemRequestBuilder) LegalHolds()(*i85c8e42d130da558580b9a46e52a3077a2dcdfb85df42f3dc3771e5e54494092.LegalHoldsRequestBuilder) {
    return i85c8e42d130da558580b9a46e52a3077a2dcdfb85df42f3dc3771e5e54494092.NewLegalHoldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LegalHoldsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.legalHolds.item collection
func (m *EdiscoveryCaseItemRequestBuilder) LegalHoldsById(id string)(*i8cab881a3513cc272eec2a3472e709a6bdc7a3974398acf13e4a2281a871cc9f.EdiscoveryHoldPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryHoldPolicy%2Did"] = id
    }
    return i8cab881a3513cc272eec2a3472e709a6bdc7a3974398acf13e4a2281a871cc9f.NewEdiscoveryHoldPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NoncustodialDataSources the noncustodialDataSources property
func (m *EdiscoveryCaseItemRequestBuilder) NoncustodialDataSources()(*if55b3465b5d38c970f2ab12abe1f658e1c8eaaba08efeed80e2df798315efd0d.NoncustodialDataSourcesRequestBuilder) {
    return if55b3465b5d38c970f2ab12abe1f658e1c8eaaba08efeed80e2df798315efd0d.NewNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialDataSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.noncustodialDataSources.item collection
func (m *EdiscoveryCaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*i07e34a8004c952f926f66ce45ec267fb3e980eebe9e100ac4d9d587e26f4f10e.EdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = id
    }
    return i07e34a8004c952f926f66ce45ec267fb3e980eebe9e100ac4d9d587e26f4f10e.NewEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *EdiscoveryCaseItemRequestBuilder) Operations()(*i6a31f80a06495660e3184b68a5c151df035e02a18ac99cfd059d1236423c6758.OperationsRequestBuilder) {
    return i6a31f80a06495660e3184b68a5c151df035e02a18ac99cfd059d1236423c6758.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.operations.item collection
func (m *EdiscoveryCaseItemRequestBuilder) OperationsById(id string)(*i165a308de50d35f4c393d5346bd98bf82a032c98d62cf3853110ba65de984b2f.CaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation%2Did"] = id
    }
    return i165a308de50d35f4c393d5346bd98bf82a032c98d62cf3853110ba65de984b2f.NewCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property ediscoveryCases in security
func (m *EdiscoveryCaseItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, requestConfiguration *EdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCaseable), nil
}
// Reopen the reopen property
func (m *EdiscoveryCaseItemRequestBuilder) Reopen()(*i9c4817aecd1c7b6fe339c7ee61c758e77ffb79c46bd4b2c82da6fc313c8559f9.ReopenRequestBuilder) {
    return i9c4817aecd1c7b6fe339c7ee61c758e77ffb79c46bd4b2c82da6fc313c8559f9.NewReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSets the reviewSets property
func (m *EdiscoveryCaseItemRequestBuilder) ReviewSets()(*iadbc9b9b01faf8819d968fb41fc7969583e10daeca4d6ca26b40b4ce48622a4f.ReviewSetsRequestBuilder) {
    return iadbc9b9b01faf8819d968fb41fc7969583e10daeca4d6ca26b40b4ce48622a4f.NewReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.reviewSets.item collection
func (m *EdiscoveryCaseItemRequestBuilder) ReviewSetsById(id string)(*if1715e408bc3525e0632493a6e5ae7e31a098303838372aef02b23c1a27e00f5.EdiscoveryReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewSet%2Did"] = id
    }
    return if1715e408bc3525e0632493a6e5ae7e31a098303838372aef02b23c1a27e00f5.NewEdiscoveryReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Searches the searches property
func (m *EdiscoveryCaseItemRequestBuilder) Searches()(*i6a6026b47c385592c13309b70526fe02801e107051c3ddd78444baca8e547380.SearchesRequestBuilder) {
    return i6a6026b47c385592c13309b70526fe02801e107051c3ddd78444baca8e547380.NewSearchesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.searches.item collection
func (m *EdiscoveryCaseItemRequestBuilder) SearchesById(id string)(*i7f2c94cccb75af6bdb0124a47a951ae1e97dcddd956b0615ac4e2d6d698ec11b.EdiscoverySearchItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoverySearch%2Did"] = id
    }
    return i7f2c94cccb75af6bdb0124a47a951ae1e97dcddd956b0615ac4e2d6d698ec11b.NewEdiscoverySearchItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings the settings property
func (m *EdiscoveryCaseItemRequestBuilder) Settings()(*idd68d7a96e6d4b3e6a6836e7a6e2e4ddfb4db1c4e2cbd40a89f985851fe47c7a.SettingsRequestBuilder) {
    return idd68d7a96e6d4b3e6a6836e7a6e2e4ddfb4db1c4e2cbd40a89f985851fe47c7a.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *EdiscoveryCaseItemRequestBuilder) Tags()(*i729fc921d5b80a60d8f1623b0a50e0190155d43bb1437138a638905bcd343eb4.TagsRequestBuilder) {
    return i729fc921d5b80a60d8f1623b0a50e0190155d43bb1437138a638905bcd343eb4.NewTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.tags.item collection
func (m *EdiscoveryCaseItemRequestBuilder) TagsById(id string)(*i98eb84e02e294af8e8f8b700991ad14c1cb72bf5681fa872dd1d942a0bca0877.EdiscoveryReviewTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewTag%2Did"] = id
    }
    return i98eb84e02e294af8e8f8b700991ad14c1cb72bf5681fa872dd1d942a0bca0877.NewEdiscoveryReviewTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
