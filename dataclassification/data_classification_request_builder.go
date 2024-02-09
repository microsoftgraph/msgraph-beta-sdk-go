package dataclassification

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataClassificationRequestBuilder provides operations to manage the dataClassificationService singleton.
type DataClassificationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *ClassifyExactMatchesRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ClassifyExactMatches()(*ClassifyExactMatchesRequestBuilder) {
    return NewClassifyExactMatchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClassifyFile provides operations to call the classifyFile method.
// returns a *ClassifyFileRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ClassifyFile()(*ClassifyFileRequestBuilder) {
    return NewClassifyFileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClassifyFileJobs provides operations to manage the classifyFileJobs property of the microsoft.graph.dataClassificationService entity.
// returns a *ClassifyFileJobsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ClassifyFileJobs()(*ClassifyFileJobsRequestBuilder) {
    return NewClassifyFileJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ClassifyTextJobs provides operations to manage the classifyTextJobs property of the microsoft.graph.dataClassificationService entity.
// returns a *ClassifyTextJobsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ClassifyTextJobs()(*ClassifyTextJobsRequestBuilder) {
    return NewClassifyTextJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDataClassificationRequestBuilderInternal instantiates a new DataClassificationRequestBuilder and sets the default values.
func NewDataClassificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationRequestBuilder) {
    m := &DataClassificationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dataClassification{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDataClassificationRequestBuilder instantiates a new DataClassificationRequestBuilder and sets the default values.
func NewDataClassificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataClassificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataClassificationRequestBuilderInternal(urlParams, requestAdapter)
}
// EvaluateDlpPoliciesJobs provides operations to manage the evaluateDlpPoliciesJobs property of the microsoft.graph.dataClassificationService entity.
// returns a *EvaluateDlpPoliciesJobsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobs()(*EvaluateDlpPoliciesJobsRequestBuilder) {
    return NewEvaluateDlpPoliciesJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateLabelJobs provides operations to manage the evaluateLabelJobs property of the microsoft.graph.dataClassificationService entity.
// returns a *EvaluateLabelJobsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) EvaluateLabelJobs()(*EvaluateLabelJobsRequestBuilder) {
    return NewEvaluateLabelJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExactMatchDataStores provides operations to manage the exactMatchDataStores property of the microsoft.graph.dataClassificationService entity.
// returns a *ExactMatchDataStoresRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ExactMatchDataStores()(*ExactMatchDataStoresRequestBuilder) {
    return NewExactMatchDataStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExactMatchUploadAgents provides operations to manage the exactMatchUploadAgents property of the microsoft.graph.dataClassificationService entity.
// returns a *ExactMatchUploadAgentsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgents()(*ExactMatchUploadAgentsRequestBuilder) {
    return NewExactMatchUploadAgentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dataClassification
// returns a DataClassificationServiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataClassificationRequestBuilder) Get(ctx context.Context, requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataClassificationServiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable), nil
}
// Jobs provides operations to manage the jobs property of the microsoft.graph.dataClassificationService entity.
// returns a *JobsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) Jobs()(*JobsRequestBuilder) {
    return NewJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update dataClassification
// returns a DataClassificationServiceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DataClassificationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataClassificationServiceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable), nil
}
// SensitiveTypes provides operations to manage the sensitiveTypes property of the microsoft.graph.dataClassificationService entity.
// returns a *SensitiveTypesRequestBuilder when successful
func (m *DataClassificationRequestBuilder) SensitiveTypes()(*SensitiveTypesRequestBuilder) {
    return NewSensitiveTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SensitivityLabels provides operations to manage the sensitivityLabels property of the microsoft.graph.dataClassificationService entity.
// returns a *SensitivityLabelsRequestBuilder when successful
func (m *DataClassificationRequestBuilder) SensitivityLabels()(*SensitivityLabelsRequestBuilder) {
    return NewSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get dataClassification
// returns a *RequestInformation when successful
func (m *DataClassificationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update dataClassification
// returns a *RequestInformation when successful
func (m *DataClassificationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/dataClassification", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DataClassificationRequestBuilder when successful
func (m *DataClassificationRequestBuilder) WithUrl(rawUrl string)(*DataClassificationRequestBuilder) {
    return NewDataClassificationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
