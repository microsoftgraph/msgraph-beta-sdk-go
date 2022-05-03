package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages"
    i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/sendreminder"
    i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/resetdecisions"
    i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions"
    iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stop"
    ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers"
    ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/batchrecorddecisions"
    ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/applydecisions"
    id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/acceptrecommendations"
    ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/definition"
    i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/decisions/item"
    i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/contactedreviewers/item"
    i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/accessreviews/definitions/item/instances/item/stages/item"
)

// AccessReviewInstanceItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
type AccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewInstanceItemRequestBuilderGetQueryParameters set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
type AccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewInstanceItemRequestBuilderGetQueryParameters
}
// AccessReviewInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations the acceptRecommendations property
func (m *AccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5.AcceptRecommendationsRequestBuilder) {
    return id859fb99d92f5a45bd2ce6145abca710efa33e2a5b1fe88b4ff16f1f437a67c5.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729.ApplyDecisionsRequestBuilder) {
    return ic8e74603f5ee41faf457cafb716b509d060bc44d1f5078afd66ec5bf53e64729.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3.BatchRecordDecisionsRequestBuilder) {
    return ic519225da39bf9a83bff86a837186671f7908ceabe699cf8fab5579f59cb50c3.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    m := &AccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewInstanceItemRequestBuilder instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers the contactedReviewers property
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce.ContactedReviewersRequestBuilder) {
    return ic0e8dea6d9a1d97ce97cabb68e83dd9aa75ea89e6f734ec7f271d9ba27f052ce.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return i778254787cb626d9fe348e2d305e8501a888ae97b99b9154fb0f2bee3401ccea.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decisions the decisions property
func (m *AccessReviewInstanceItemRequestBuilder) Decisions()(*i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8.DecisionsRequestBuilder) {
    return i9327ed35d3b3a7e13ba1ba8f389450b462683ccdaccff95c747cb5f750f0d4d8.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.decisions.item collection
func (m *AccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did"] = id
    }
    return i425fb16e7ed420101998d1e876028364d5747a9897e3d5604f4a824ec3e04ffb.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition the definition property
func (m *AccessReviewInstanceItemRequestBuilder) Definition()(*ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3.DefinitionRequestBuilder) {
    return ie25b10e68fa81f5005eea5f527a61e7a311f8b60257495559d75d76510909fe3.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeleteWithResponseHandler delete navigation property instances for identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property instances for identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AccessReviewInstanceItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessReviewInstanceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewInstanceItemRequestBuilder) GetWithResponseHandler(requestConfiguration *AccessReviewInstanceItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// PatchWithResponseHandler update the navigation property instances in identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewInstanceItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property instances in identityGovernance
func (m *AccessReviewInstanceItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *AccessReviewInstanceItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// ResetDecisions the resetDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) ResetDecisions()(*i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534.ResetDecisionsRequestBuilder) {
    return i6a0a160cb4443ebb8d098e51ce55956f64d15462282145c7d7333c345b943534.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *AccessReviewInstanceItemRequestBuilder) SendReminder()(*i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe.SendReminderRequestBuilder) {
    return i4167933608f10c80b058ba8a45ffcaff5a3638d6f3458cded3059b347646eabe.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages the stages property
func (m *AccessReviewInstanceItemRequestBuilder) Stages()(*i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6.StagesRequestBuilder) {
    return i34029e4c2d0a245f6ca7822cbbd442b2a54ddd3d92e7f8f9797638a7756d56d6.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.accessReviews.definitions.item.instances.item.stages.item collection
func (m *AccessReviewInstanceItemRequestBuilder) StagesById(id string)(*i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e.AccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return i83c0dcbc9b5cafef3dad8f565d30a946fd80a79a1ba13afeef9c56a7f634d98e.NewAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop the stop property
func (m *AccessReviewInstanceItemRequestBuilder) Stop()(*iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4.StopRequestBuilder) {
    return iaee1412a0b517074ea75eb9acdb529e3a9e5a31a64757270c95df5946da814b4.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
