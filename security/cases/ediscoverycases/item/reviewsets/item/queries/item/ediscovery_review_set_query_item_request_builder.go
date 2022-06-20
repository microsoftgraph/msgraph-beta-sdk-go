package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5778ef3de4edb4bac772a49fde90f34aa49d80488a90906b0688ef00dd758cde "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries/item/run"
    i9e1e4e7e8e8cd680826d746191f2ddd4e1784912c6628647895647279baf2784 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries/item/export"
    if9834b50a79825154b32059bdb0d29f98aa338e9369f666af5f862b9e6d2360e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/reviewsets/item/queries/item/applytags"
)

// EdiscoveryReviewSetQueryItemRequestBuilder provides operations to manage the queries property of the microsoft.graph.security.ediscoveryReviewSet entity.
type EdiscoveryReviewSetQueryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters represents queries within the review set.
type EdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryReviewSetQueryItemRequestBuilderGetQueryParameters
}
// EdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyTags the applyTags property
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) ApplyTags()(*if9834b50a79825154b32059bdb0d29f98aa338e9369f666af5f862b9e6d2360e.ApplyTagsRequestBuilder) {
    return if9834b50a79825154b32059bdb0d29f98aa338e9369f666af5f862b9e6d2360e.NewApplyTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoveryReviewSetQueryItemRequestBuilderInternal instantiates a new EdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetQueryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetQueryItemRequestBuilder) {
    m := &EdiscoveryReviewSetQueryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets/{ediscoveryReviewSet%2Did}/queries/{ediscoveryReviewSetQuery%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryReviewSetQueryItemRequestBuilder instantiates a new EdiscoveryReviewSetQueryItemRequestBuilder and sets the default values.
func NewEdiscoveryReviewSetQueryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryReviewSetQueryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryReviewSetQueryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property queries for security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property queries for security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents queries within the review set.
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration represents queries within the review set.
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property queries in security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property queries in security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property queries for security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property queries for security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) Export()(*i9e1e4e7e8e8cd680826d746191f2ddd4e1784912c6628647895647279baf2784.ExportRequestBuilder) {
    return i9e1e4e7e8e8cd680826d746191f2ddd4e1784912c6628647895647279baf2784.NewExportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents queries within the review set.
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) Get()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler represents queries within the review set.
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoveryReviewSetQueryFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable), nil
}
// Patch update the navigation property queries in security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) Patch(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property queries in security
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoveryReviewSetQueryable, requestConfiguration *EdiscoveryReviewSetQueryItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Run provides operations to call the run method.
func (m *EdiscoveryReviewSetQueryItemRequestBuilder) Run()(*i5778ef3de4edb4bac772a49fde90f34aa49d80488a90906b0688ef00dd758cde.RunRequestBuilder) {
    return i5778ef3de4edb4bac772a49fde90f34aa49d80488a90906b0688ef00dd758cde.NewRunRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
