package instance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/resetdecisions"
    i19220c6873e6f178a0065161bbaa90caa6620ef6dd66a5ec4758b123ac469187 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/definition"
    i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/acceptrecommendations"
    i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/stop"
    i8465856a57b39bfe6bbaee998f31cc1a6f95a416a097b32bdd5ae57760891e36 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/contactedreviewers"
    ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/sendreminder"
    ic5796aa32538952c642db31a0c8fa19775f1f8e0f83c910a8e3adc526597e81c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/stages"
    icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/batchrecorddecisions"
    iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/applydecisions"
    i42bb427c442cb4d4ddf1c7f5ca19862ac0bb20430fb80b4f937cdc05437f296a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/stages/item"
    ib94b9ab817e6471af61b50e251fd606a0d13640a636ee601a4261fc8d3edd511 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item/instance/contactedreviewers/item"
)

// InstanceRequestBuilder provides operations to manage the instance property of the microsoft.graph.accessReviewInstanceDecisionItem entity.
type InstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InstanceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type InstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InstanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstanceRequestBuilderGetQueryParameters
}
// InstanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptRecommendations the acceptRecommendations property
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d.AcceptRecommendationsRequestBuilder) {
    return i33c7c5706a7b5c0bf13b40b705217aed59b2d5208c669fd999e36faa62df7d8d.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *InstanceRequestBuilder) ApplyDecisions()(*iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba.ApplyDecisionsRequestBuilder) {
    return iebc762487b2890223499b50b54324ea41c534597ca8a925c094747c4cc9293ba.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153.BatchRecordDecisionsRequestBuilder) {
    return icd1d27bb795cd0bade6075661b23050cc531c7e953d93767d6867d925acd1153.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstanceRequestBuilder instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactedReviewers the contactedReviewers property
func (m *InstanceRequestBuilder) ContactedReviewers()(*i8465856a57b39bfe6bbaee998f31cc1a6f95a416a097b32bdd5ae57760891e36.ContactedReviewersRequestBuilder) {
    return i8465856a57b39bfe6bbaee998f31cc1a6f95a416a097b32bdd5ae57760891e36.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.decisions.item.instance.contactedReviewers.item collection
func (m *InstanceRequestBuilder) ContactedReviewersById(id string)(*ib94b9ab817e6471af61b50e251fd606a0d13640a636ee601a4261fc8d3edd511.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return ib94b9ab817e6471af61b50e251fd606a0d13640a636ee601a4261fc8d3edd511.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property instance for me
func (m *InstanceRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instance for me
func (m *InstanceRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instance in me
func (m *InstanceRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instance in me
func (m *InstanceRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Definition the definition property
func (m *InstanceRequestBuilder) Definition()(*i19220c6873e6f178a0065161bbaa90caa6620ef6dd66a5ec4758b123ac469187.DefinitionRequestBuilder) {
    return i19220c6873e6f178a0065161bbaa90caa6620ef6dd66a5ec4758b123ac469187.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instance for me
func (m *InstanceRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instance for me
func (m *InstanceRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *InstanceRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property instance in me
func (m *InstanceRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instance in me
func (m *InstanceRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *InstanceRequestBuilder) ResetDecisions()(*i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d.ResetDecisionsRequestBuilder) {
    return i03da4522f6ecfec1c06db2ac24e7752b8ccd635602f5fe82922c3a68e1d0dd7d.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *InstanceRequestBuilder) SendReminder()(*ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f.SendReminderRequestBuilder) {
    return ia755ae82953cc4dc7511248e4432b31518b4ceb68464488197f488ec3a27c58f.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages the stages property
func (m *InstanceRequestBuilder) Stages()(*ic5796aa32538952c642db31a0c8fa19775f1f8e0f83c910a8e3adc526597e81c.StagesRequestBuilder) {
    return ic5796aa32538952c642db31a0c8fa19775f1f8e0f83c910a8e3adc526597e81c.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.decisions.item.instance.stages.item collection
func (m *InstanceRequestBuilder) StagesById(id string)(*i42bb427c442cb4d4ddf1c7f5ca19862ac0bb20430fb80b4f937cdc05437f296a.AccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return i42bb427c442cb4d4ddf1c7f5ca19862ac0bb20430fb80b4f937cdc05437f296a.NewAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop the stop property
func (m *InstanceRequestBuilder) Stop()(*i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6.StopRequestBuilder) {
    return i74deae20af1cfa651b11f8e4117763fc91498e69766fedcf61e65d3873b58bd6.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
