package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03f3e0d3baf12c177e4672910cc93cd351204faa39b38e9125b94b0465177cbd "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/sitesources"
    i1d6b86465bb3bc751dc147d093bd1285a99a35daf39ad5608f2d968591cb3ee4 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/updateindex"
    i250ddd0aa330bb7d9ea55c916947840d121a66c3dde4b7545742b82e8f6e7ac1 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/removehold"
    i25f82dbb8cf7ca352ce0f843e8fc8cd4fcfbd99f773b8a5f96233e8b67e0134b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/release"
    i4cded81d6396494b55b9783880804a76233d8daa057ffd6582c7638ebbc5b0e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/activate"
    i8207885b0bf2ea2d63475bae7d26fa0498c6807c91130210ada883efffcdff1e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/lastindexoperation"
    idc366a2de9455855aad941d033dcd4e11623a2239a8994da3404c6e79873f8ed "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/applyhold"
    idf7d17e258ae3eb429fcc8b0a4fb77cbb89c88802013e167f54b495ebd387b15 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/usersources"
    ie416f7f615f3b8cc0d03e5759e3f2a5be926bc05b1c601c65a9cd332cd0752da "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/unifiedgroupsources"
    i089a870e693072fe784f18deae6aef7a21348d641c984511c981fa961f3e7024 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/unifiedgroupsources/item"
    i143d5b3bc52e8356276951d7cc0a9c058b11eba317e2d6b3e7ac0db12a0f480a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/sitesources/item"
    i5bd1512a36bd7cc5e41d7ce737db1fe666eea2a3c15417a8d9e1b607a6d57151 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/custodians/item/usersources/item"
)

// EdiscoveryCustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
type EdiscoveryCustodianItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryCustodianItemRequestBuilderGetQueryParameters returns a list of case ediscoveryCustodian objects for this case.
type EdiscoveryCustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCustodianItemRequestBuilderGetQueryParameters
}
// EdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate the activate property
func (m *EdiscoveryCustodianItemRequestBuilder) Activate()(*i4cded81d6396494b55b9783880804a76233d8daa057ffd6582c7638ebbc5b0e9.ActivateRequestBuilder) {
    return i4cded81d6396494b55b9783880804a76233d8daa057ffd6582c7638ebbc5b0e9.NewActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyHold the applyHold property
func (m *EdiscoveryCustodianItemRequestBuilder) ApplyHold()(*idc366a2de9455855aad941d033dcd4e11623a2239a8994da3404c6e79873f8ed.ApplyHoldRequestBuilder) {
    return idc366a2de9455855aad941d033dcd4e11623a2239a8994da3404c6e79873f8ed.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoveryCustodianItemRequestBuilderInternal instantiates a new EdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewEdiscoveryCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCustodianItemRequestBuilder) {
    m := &EdiscoveryCustodianItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryCustodianItemRequestBuilder instantiates a new EdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewEdiscoveryCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property custodians for security
func (m *EdiscoveryCustodianItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property custodians for security
func (m *EdiscoveryCustodianItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of case ediscoveryCustodian objects for this case.
func (m *EdiscoveryCustodianItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of case ediscoveryCustodian objects for this case.
func (m *EdiscoveryCustodianItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property custodians in security
func (m *EdiscoveryCustodianItemRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property custodians in security
func (m *EdiscoveryCustodianItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, requestConfiguration *EdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property custodians for security
func (m *EdiscoveryCustodianItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get returns a list of case ediscoveryCustodian objects for this case.
func (m *EdiscoveryCustodianItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable), nil
}
// LastIndexOperation the lastIndexOperation property
func (m *EdiscoveryCustodianItemRequestBuilder) LastIndexOperation()(*i8207885b0bf2ea2d63475bae7d26fa0498c6807c91130210ada883efffcdff1e.LastIndexOperationRequestBuilder) {
    return i8207885b0bf2ea2d63475bae7d26fa0498c6807c91130210ada883efffcdff1e.NewLastIndexOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property custodians in security
func (m *EdiscoveryCustodianItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryCustodianable, requestConfiguration *EdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Release the release property
func (m *EdiscoveryCustodianItemRequestBuilder) Release()(*i25f82dbb8cf7ca352ce0f843e8fc8cd4fcfbd99f773b8a5f96233e8b67e0134b.ReleaseRequestBuilder) {
    return i25f82dbb8cf7ca352ce0f843e8fc8cd4fcfbd99f773b8a5f96233e8b67e0134b.NewReleaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveHold the removeHold property
func (m *EdiscoveryCustodianItemRequestBuilder) RemoveHold()(*i250ddd0aa330bb7d9ea55c916947840d121a66c3dde4b7545742b82e8f6e7ac1.RemoveHoldRequestBuilder) {
    return i250ddd0aa330bb7d9ea55c916947840d121a66c3dde4b7545742b82e8f6e7ac1.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSources the siteSources property
func (m *EdiscoveryCustodianItemRequestBuilder) SiteSources()(*i03f3e0d3baf12c177e4672910cc93cd351204faa39b38e9125b94b0465177cbd.SiteSourcesRequestBuilder) {
    return i03f3e0d3baf12c177e4672910cc93cd351204faa39b38e9125b94b0465177cbd.NewSiteSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SiteSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.custodians.item.siteSources.item collection
func (m *EdiscoveryCustodianItemRequestBuilder) SiteSourcesById(id string)(*i143d5b3bc52e8356276951d7cc0a9c058b11eba317e2d6b3e7ac0db12a0f480a.SiteSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["siteSource%2Did"] = id
    }
    return i143d5b3bc52e8356276951d7cc0a9c058b11eba317e2d6b3e7ac0db12a0f480a.NewSiteSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnifiedGroupSources the unifiedGroupSources property
func (m *EdiscoveryCustodianItemRequestBuilder) UnifiedGroupSources()(*ie416f7f615f3b8cc0d03e5759e3f2a5be926bc05b1c601c65a9cd332cd0752da.UnifiedGroupSourcesRequestBuilder) {
    return ie416f7f615f3b8cc0d03e5759e3f2a5be926bc05b1c601c65a9cd332cd0752da.NewUnifiedGroupSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnifiedGroupSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.custodians.item.unifiedGroupSources.item collection
func (m *EdiscoveryCustodianItemRequestBuilder) UnifiedGroupSourcesById(id string)(*i089a870e693072fe784f18deae6aef7a21348d641c984511c981fa961f3e7024.UnifiedGroupSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedGroupSource%2Did"] = id
    }
    return i089a870e693072fe784f18deae6aef7a21348d641c984511c981fa961f3e7024.NewUnifiedGroupSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateIndex the updateIndex property
func (m *EdiscoveryCustodianItemRequestBuilder) UpdateIndex()(*i1d6b86465bb3bc751dc147d093bd1285a99a35daf39ad5608f2d968591cb3ee4.UpdateIndexRequestBuilder) {
    return i1d6b86465bb3bc751dc147d093bd1285a99a35daf39ad5608f2d968591cb3ee4.NewUpdateIndexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSources the userSources property
func (m *EdiscoveryCustodianItemRequestBuilder) UserSources()(*idf7d17e258ae3eb429fcc8b0a4fb77cbb89c88802013e167f54b495ebd387b15.UserSourcesRequestBuilder) {
    return idf7d17e258ae3eb429fcc8b0a4fb77cbb89c88802013e167f54b495ebd387b15.NewUserSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.custodians.item.userSources.item collection
func (m *EdiscoveryCustodianItemRequestBuilder) UserSourcesById(id string)(*i5bd1512a36bd7cc5e41d7ce737db1fe666eea2a3c15417a8d9e1b607a6d57151.UserSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSource%2Did"] = id
    }
    return i5bd1512a36bd7cc5e41d7ce737db1fe666eea2a3c15417a8d9e1b607a6d57151.NewUserSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
