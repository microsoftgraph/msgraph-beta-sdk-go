package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataClassificationRequestBuilder provides operations to manage the dataClassificationService singleton.
type DataClassificationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DataClassificationRequestBuilderGetQueryParameters get dataClassification
type DataClassificationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DataClassificationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataClassificationRequestBuilderGetQueryParameters
}
// DataClassificationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassifyExactMatches provides operations to call the classifyExactMatches method.
func (m *DataClassificationRequestBuilder) ClassifyExactMatches()(*ClassifyExactMatchesRequestBuilder) {
    return NewClassifyExactMatchesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFile provides operations to call the classifyFile method.
func (m *DataClassificationRequestBuilder) ClassifyFile()(*ClassifyFileRequestBuilder) {
    return NewClassifyFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFileJobs provides operations to manage the classifyFileJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ClassifyFileJobs()(*ClassifyFileJobsRequestBuilder) {
    return NewClassifyFileJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFileJobsById provides operations to manage the classifyFileJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ClassifyFileJobsById(id string)(*ClassifyFileJobsJobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return NewClassifyFileJobsJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ClassifyTextJobs provides operations to manage the classifyTextJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ClassifyTextJobs()(*ClassifyTextJobsRequestBuilder) {
    return NewClassifyTextJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyTextJobsById provides operations to manage the classifyTextJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ClassifyTextJobsById(id string)(*ClassifyTextJobsJobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return NewClassifyTextJobsJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDataClassificationRequestBuilderInternal instantiates a new DataClassificationRequestBuilder and sets the default values.
func NewDataClassificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationRequestBuilder) {
    m := &DataClassificationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDataClassificationRequestBuilder instantiates a new DataClassificationRequestBuilder and sets the default values.
func NewDataClassificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataClassificationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get dataClassification
func (m *DataClassificationRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update dataClassification
func (m *DataClassificationRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EvaluateDlpPoliciesJobs provides operations to manage the evaluateDlpPoliciesJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobs()(*EvaluateDlpPoliciesJobsRequestBuilder) {
    return NewEvaluateDlpPoliciesJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDlpPoliciesJobsById provides operations to manage the evaluateDlpPoliciesJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobsById(id string)(*EvaluateDlpPoliciesJobsJobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return NewEvaluateDlpPoliciesJobsJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EvaluateLabelJobs provides operations to manage the evaluateLabelJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) EvaluateLabelJobs()(*EvaluateLabelJobsRequestBuilder) {
    return NewEvaluateLabelJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateLabelJobsById provides operations to manage the evaluateLabelJobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) EvaluateLabelJobsById(id string)(*EvaluateLabelJobsJobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return NewEvaluateLabelJobsJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExactMatchDataStores provides operations to manage the exactMatchDataStores property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ExactMatchDataStores()(*ExactMatchDataStoresRequestBuilder) {
    return NewExactMatchDataStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExactMatchDataStoresById provides operations to manage the exactMatchDataStores property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ExactMatchDataStoresById(id string)(*ExactMatchDataStoresExactMatchDataStoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchDataStore%2Did"] = id
    }
    return NewExactMatchDataStoresExactMatchDataStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExactMatchUploadAgents provides operations to manage the exactMatchUploadAgents property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgents()(*ExactMatchUploadAgentsRequestBuilder) {
    return NewExactMatchUploadAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExactMatchUploadAgentsById provides operations to manage the exactMatchUploadAgents property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgentsById(id string)(*ExactMatchUploadAgentsExactMatchUploadAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchUploadAgent%2Did"] = id
    }
    return NewExactMatchUploadAgentsExactMatchUploadAgentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get dataClassification
func (m *DataClassificationRequestBuilder) Get(ctx context.Context, requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataClassificationServiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable), nil
}
// Jobs provides operations to manage the jobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) Jobs()(*JobsRequestBuilder) {
    return NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById provides operations to manage the jobs property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) JobsById(id string)(*JobsJobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return NewJobsJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update dataClassification
func (m *DataClassificationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataClassificationServiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable), nil
}
// SensitiveTypes provides operations to manage the sensitiveTypes property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) SensitiveTypes()(*SensitiveTypesRequestBuilder) {
    return NewSensitiveTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitiveTypesById provides operations to manage the sensitiveTypes property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) SensitiveTypesById(id string)(*SensitiveTypesSensitiveTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitiveType%2Did"] = id
    }
    return NewSensitiveTypesSensitiveTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) SensitivityLabels()(*SensitivityLabelsRequestBuilder) {
    return NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById provides operations to manage the sensitivityLabels property of the microsoft.graph.dataClassificationService entity.
func (m *DataClassificationRequestBuilder) SensitivityLabelsById(id string)(*SensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return NewSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
