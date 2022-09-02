package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i20feedbdbbe1098ec21a5bdba011824d8a82b5ff464d2eae69fe1d03ffc58470 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/additionalsources"
    i809d681c352cb5981157e3c438d466a1b04de6850a1f8211ff8456fd93289888 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/noncustodialsources"
    i84bdda408caa738046f53eaffe4066b7e8cbbac60c83398b34aafb24b42d5e9c "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/estimatestatistics"
    i9bf0d32205555a744ff805d0ce39eea026941eb4896b1268ef8e921ca5dfdf1d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/custodiansources"
    ica43e921341b90a322b56ec6c3f08593d7956347237d8db4d27114c4d354714d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/lastestimatestatisticsoperation"
    iec28b22188aa05851e7c8aca2473330329a279744509cf39b2c5175a754d5376 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/addtoreviewsetoperation"
    ifdaddfb3022f2acbe9a32248a06780d00ed48f808d39f914d8f26934b9c47ba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/purgedata"
    i4ef9cedb8e95d7629ba74b96b547fadee4741e9a4653f30050e66f19aa77f5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/custodiansources/item"
    i786b552748f08c251421195ccd48e540403334e7feba508548cb1e3d4f4a047b "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/additionalsources/item"
    ib0f2a8ff30297fdbdfb28756c7e08fd35d93f2139a8947c69fc06dab14a53738 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/cases/ediscoverycases/item/searches/item/noncustodialsources/item"
)

// EdiscoverySearchItemRequestBuilder provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
type EdiscoverySearchItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoverySearchItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoverySearchItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdiscoverySearchItemRequestBuilderGetQueryParameters returns a list of eDiscoverySearch objects associated with this case.
type EdiscoverySearchItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoverySearchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoverySearchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoverySearchItemRequestBuilderGetQueryParameters
}
// EdiscoverySearchItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoverySearchItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources the additionalSources property
func (m *EdiscoverySearchItemRequestBuilder) AdditionalSources()(*i20feedbdbbe1098ec21a5bdba011824d8a82b5ff464d2eae69fe1d03ffc58470.AdditionalSourcesRequestBuilder) {
    return i20feedbdbbe1098ec21a5bdba011824d8a82b5ff464d2eae69fe1d03ffc58470.NewAdditionalSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdditionalSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.searches.item.additionalSources.item collection
func (m *EdiscoverySearchItemRequestBuilder) AdditionalSourcesById(id string)(*i786b552748f08c251421195ccd48e540403334e7feba508548cb1e3d4f4a047b.DataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return i786b552748f08c251421195ccd48e540403334e7feba508548cb1e3d4f4a047b.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddToReviewSetOperation the addToReviewSetOperation property
func (m *EdiscoverySearchItemRequestBuilder) AddToReviewSetOperation()(*iec28b22188aa05851e7c8aca2473330329a279744509cf39b2c5175a754d5376.AddToReviewSetOperationRequestBuilder) {
    return iec28b22188aa05851e7c8aca2473330329a279744509cf39b2c5175a754d5376.NewAddToReviewSetOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEdiscoverySearchItemRequestBuilderInternal instantiates a new EdiscoverySearchItemRequestBuilder and sets the default values.
func NewEdiscoverySearchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoverySearchItemRequestBuilder) {
    m := &EdiscoverySearchItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoverySearchItemRequestBuilder instantiates a new EdiscoverySearchItemRequestBuilder and sets the default values.
func NewEdiscoverySearchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoverySearchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoverySearchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property searches for security
func (m *EdiscoverySearchItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property searches for security
func (m *EdiscoverySearchItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation returns a list of eDiscoverySearch objects associated with this case.
func (m *EdiscoverySearchItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of eDiscoverySearch objects associated with this case.
func (m *EdiscoverySearchItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EdiscoverySearchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property searches in security
func (m *EdiscoverySearchItemRequestBuilder) CreatePatchRequestInformation(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property searches in security
func (m *EdiscoverySearchItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable, requestConfiguration *EdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CustodianSources the custodianSources property
func (m *EdiscoverySearchItemRequestBuilder) CustodianSources()(*i9bf0d32205555a744ff805d0ce39eea026941eb4896b1268ef8e921ca5dfdf1d.CustodianSourcesRequestBuilder) {
    return i9bf0d32205555a744ff805d0ce39eea026941eb4896b1268ef8e921ca5dfdf1d.NewCustodianSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodianSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.searches.item.custodianSources.item collection
func (m *EdiscoverySearchItemRequestBuilder) CustodianSourcesById(id string)(*i4ef9cedb8e95d7629ba74b96b547fadee4741e9a4653f30050e66f19aa77f5ac.DataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSource%2Did"] = id
    }
    return i4ef9cedb8e95d7629ba74b96b547fadee4741e9a4653f30050e66f19aa77f5ac.NewDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property searches for security
func (m *EdiscoverySearchItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EstimateStatistics the estimateStatistics property
func (m *EdiscoverySearchItemRequestBuilder) EstimateStatistics()(*i84bdda408caa738046f53eaffe4066b7e8cbbac60c83398b34aafb24b42d5e9c.EstimateStatisticsRequestBuilder) {
    return i84bdda408caa738046f53eaffe4066b7e8cbbac60c83398b34aafb24b42d5e9c.NewEstimateStatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get returns a list of eDiscoverySearch objects associated with this case.
func (m *EdiscoverySearchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoverySearchItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateEdiscoverySearchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable), nil
}
// LastEstimateStatisticsOperation the lastEstimateStatisticsOperation property
func (m *EdiscoverySearchItemRequestBuilder) LastEstimateStatisticsOperation()(*ica43e921341b90a322b56ec6c3f08593d7956347237d8db4d27114c4d354714d.LastEstimateStatisticsOperationRequestBuilder) {
    return ica43e921341b90a322b56ec6c3f08593d7956347237d8db4d27114c4d354714d.NewLastEstimateStatisticsOperationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSources the noncustodialSources property
func (m *EdiscoverySearchItemRequestBuilder) NoncustodialSources()(*i809d681c352cb5981157e3c438d466a1b04de6850a1f8211ff8456fd93289888.NoncustodialSourcesRequestBuilder) {
    return i809d681c352cb5981157e3c438d466a1b04de6850a1f8211ff8456fd93289888.NewNoncustodialSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialSourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.cases.ediscoveryCases.item.searches.item.noncustodialSources.item collection
func (m *EdiscoverySearchItemRequestBuilder) NoncustodialSourcesById(id string)(*ib0f2a8ff30297fdbdfb28756c7e08fd35d93f2139a8947c69fc06dab14a53738.EdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = id
    }
    return ib0f2a8ff30297fdbdfb28756c7e08fd35d93f2139a8947c69fc06dab14a53738.NewEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property searches in security
func (m *EdiscoverySearchItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.EdiscoverySearchable, requestConfiguration *EdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(error) {
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
// PurgeData the purgeData property
func (m *EdiscoverySearchItemRequestBuilder) PurgeData()(*ifdaddfb3022f2acbe9a32248a06780d00ed48f808d39f914d8f26934b9c47ba1.PurgeDataRequestBuilder) {
    return ifdaddfb3022f2acbe9a32248a06780d00ed48f808d39f914d8f26934b9c47ba1.NewPurgeDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
