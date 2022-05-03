package dataclassification

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchuploadagents"
    i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytext"
    i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatelabeljobs"
    i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/jobs"
    i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivitylabels"
    i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfilejobs"
    i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyexactmatches"
    ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfile"
    ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatedlppoliciesjobs"
    if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivetypes"
    ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytextjobs"
    i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivitylabels/item"
    i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytextjobs/item"
    i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfilejobs/item"
    i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatedlppoliciesjobs/item"
    i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchuploadagents/item"
    i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/jobs/item"
    i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivetypes/item"
    i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item"
    id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatelabeljobs/item"
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DataClassificationRequestBuilderGetQueryParameters
}
// DataClassificationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataClassificationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassifyExactMatches the classifyExactMatches property
func (m *DataClassificationRequestBuilder) ClassifyExactMatches()(*i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860.ClassifyExactMatchesRequestBuilder) {
    return i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860.NewClassifyExactMatchesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFile the classifyFile property
func (m *DataClassificationRequestBuilder) ClassifyFile()(*ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e.ClassifyFileRequestBuilder) {
    return ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e.NewClassifyFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFileJobs the classifyFileJobs property
func (m *DataClassificationRequestBuilder) ClassifyFileJobs()(*i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9.ClassifyFileJobsRequestBuilder) {
    return i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9.NewClassifyFileJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyFileJobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.classifyFileJobs.item collection
func (m *DataClassificationRequestBuilder) ClassifyFileJobsById(id string)(*i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5.JobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5.NewJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ClassifyText the classifyText property
func (m *DataClassificationRequestBuilder) ClassifyText()(*i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb.ClassifyTextRequestBuilder) {
    return i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb.NewClassifyTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyTextJobs the classifyTextJobs property
func (m *DataClassificationRequestBuilder) ClassifyTextJobs()(*ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70.ClassifyTextJobsRequestBuilder) {
    return ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70.NewClassifyTextJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassifyTextJobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.classifyTextJobs.item collection
func (m *DataClassificationRequestBuilder) ClassifyTextJobsById(id string)(*i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f.JobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f.NewJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *DataClassificationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get dataClassification
func (m *DataClassificationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update dataClassification
func (m *DataClassificationRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update dataClassification
func (m *DataClassificationRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// EvaluateDlpPoliciesJobs the evaluateDlpPoliciesJobs property
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobs()(*ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d.EvaluateDlpPoliciesJobsRequestBuilder) {
    return ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d.NewEvaluateDlpPoliciesJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDlpPoliciesJobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.evaluateDlpPoliciesJobs.item collection
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobsById(id string)(*i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f.JobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f.NewJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EvaluateLabelJobs the evaluateLabelJobs property
func (m *DataClassificationRequestBuilder) EvaluateLabelJobs()(*i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e.EvaluateLabelJobsRequestBuilder) {
    return i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e.NewEvaluateLabelJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateLabelJobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.evaluateLabelJobs.item collection
func (m *DataClassificationRequestBuilder) EvaluateLabelJobsById(id string)(*id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4.JobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4.NewJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExactMatchDataStores the exactMatchDataStores property
func (m *DataClassificationRequestBuilder) ExactMatchDataStores()(*i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815.ExactMatchDataStoresRequestBuilder) {
    return i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815.NewExactMatchDataStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExactMatchDataStoresById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.exactMatchDataStores.item collection
func (m *DataClassificationRequestBuilder) ExactMatchDataStoresById(id string)(*i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230.ExactMatchDataStoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchDataStore%2Did"] = id
    }
    return i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230.NewExactMatchDataStoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExactMatchUploadAgents the exactMatchUploadAgents property
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgents()(*i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798.ExactMatchUploadAgentsRequestBuilder) {
    return i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798.NewExactMatchUploadAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExactMatchUploadAgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.exactMatchUploadAgents.item collection
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgentsById(id string)(*i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b.ExactMatchUploadAgentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchUploadAgent%2Did"] = id
    }
    return i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b.NewExactMatchUploadAgentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get dataClassification
func (m *DataClassificationRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get dataClassification
func (m *DataClassificationRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DataClassificationRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataClassificationServiceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable), nil
}
// Jobs the jobs property
func (m *DataClassificationRequestBuilder) Jobs()(*i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0.JobsRequestBuilder) {
    return i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JobsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.jobs.item collection
func (m *DataClassificationRequestBuilder) JobsById(id string)(*i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a.JobResponseBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase%2Did"] = id
    }
    return i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a.NewJobResponseBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update dataClassification
func (m *DataClassificationRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update dataClassification
func (m *DataClassificationRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataClassificationServiceable, requestConfiguration *DataClassificationRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// SensitiveTypes the sensitiveTypes property
func (m *DataClassificationRequestBuilder) SensitiveTypes()(*if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986.SensitiveTypesRequestBuilder) {
    return if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986.NewSensitiveTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitiveTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.sensitiveTypes.item collection
func (m *DataClassificationRequestBuilder) SensitiveTypesById(id string)(*i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c.SensitiveTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitiveType%2Did"] = id
    }
    return i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c.NewSensitiveTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityLabels the sensitivityLabels property
func (m *DataClassificationRequestBuilder) SensitivityLabels()(*i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf.SensitivityLabelsRequestBuilder) {
    return i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.dataClassification.sensitivityLabels.item collection
func (m *DataClassificationRequestBuilder) SensitivityLabelsById(id string)(*i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf.SensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf.NewSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
