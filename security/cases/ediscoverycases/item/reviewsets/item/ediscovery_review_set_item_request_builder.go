package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0eca40fca99f521380db599179b32d183a6d2074f8c40eac1aba336327909dce "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries"
    i5ff73ef2fe7116c20ed58367d30484db3d963ecc309ea6f5001b424ee079a433 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/files"
    i6ee8072b7be515f2616f558e647d2f0b734c1d8fd4d866b1080a2d6ba5c6a720 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/addtoreviewset"
    ib064e11e062e18e98877ab3fdf949472ea03e136c1932231c0139d9ae22ce1df "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/export"
    i043c811701556dc7039d9829366dfbd791bad059dc33385fdc660c2cb4fe3c7e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/files/item"
    ie5eca18cf54c63e4a7756b0ff25ab10a6405f77676774867434b697867eb2738 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries/item"
)

// EdiscoveryReviewSetItemRequestBuilder provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
type EdiscoveryReviewSetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryReviewSetItemRequestBuilderGetQueryParameters returns a list of eDiscoveryReviewSet objects in the case.
type EdiscoveryReviewSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryReviewSetItemRequestBuilderGetQueryParameters
}
// EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddToReviewSet the addToReviewSet property
func (m *EdiscoveryReviewSetItemRequestBuilder) AddToReviewSet()(*i6ee8072b7be515f2616f558e647d2f0b734c1d8fd4d866b1080a2d6ba5c6a720.AddToReviewSetRequestBuilder) {
    return i6ee8072b7be515f2616f558e647d2f0b734c1d8fd4d866b1080a2d6ba5c6a720.NewAddToReviewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoveryReviewSetItemRequestBuilderInternal instantiates a new EdiscoveryReviewSetItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetItemRequestBuilder) {
    m := &EdiscoveryReviewSetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryReviewSetItemRequestBuilder instantiates a new EdiscoveryReviewSetItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryReviewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable, requestConfiguration *EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property reviewSets for security
func (m *EdiscoveryReviewSetItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Export the export property
func (m *EdiscoveryReviewSetItemRequestBuilder) Export()(*ib064e11e062e18e98877ab3fdf949472ea03e136c1932231c0139d9ae22ce1df.ExportRequestBuilder) {
    return ib064e11e062e18e98877ab3fdf949472ea03e136c1932231c0139d9ae22ce1df.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Files the files property
func (m *EdiscoveryReviewSetItemRequestBuilder) Files()(*i5ff73ef2fe7116c20ed58367d30484db3d963ecc309ea6f5001b424ee079a433.FilesRequestBuilder) {
    return i5ff73ef2fe7116c20ed58367d30484db3d963ecc309ea6f5001b424ee079a433.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.reviewSets.item.files.item collection
func (m *EdiscoveryReviewSetItemRequestBuilder) FilesById(id string)(*i043c811701556dc7039d9829366dfbd791bad059dc33385fdc660c2cb4fe3c7e.EdiscoveryFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryFile%2Did"] = id
    }
    return i043c811701556dc7039d9829366dfbd791bad059dc33385fdc660c2cb4fe3c7e.NewEdiscoveryFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) Get()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler returns a list of eDiscoveryReviewSet objects in the case.
func (m *EdiscoveryReviewSetItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable), nil
}
// Patch update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) Patch(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property reviewSets in security
func (m *EdiscoveryReviewSetItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetable, requestConfiguration *EdiscoveryReviewSetItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Queries the queries property
func (m *EdiscoveryReviewSetItemRequestBuilder) Queries()(*i0eca40fca99f521380db599179b32d183a6d2074f8c40eac1aba336327909dce.QueriesRequestBuilder) {
    return i0eca40fca99f521380db599179b32d183a6d2074f8c40eac1aba336327909dce.NewQueriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// QueriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.reviewSets.item.queries.item collection
func (m *EdiscoveryReviewSetItemRequestBuilder) QueriesById(id string)(*ie5eca18cf54c63e4a7756b0ff25ab10a6405f77676774867434b697867eb2738.EdiscoveryReviewSetQueryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewSetQuery%2Did"] = id
    }
    return ie5eca18cf54c63e4a7756b0ff25ab10a6405f77676774867434b697867eb2738.NewEdiscoveryReviewSetQueryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
