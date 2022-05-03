package instance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/definition"
    i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/stop"
    i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/batchrecorddecisions"
    i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/acceptrecommendations"
    ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/applydecisions"
    ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/resetdecisions"
    id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/sendreminder"
    ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/contactedreviewers"
    if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/decisions"
    i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/decisions/item"
    iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/stages/item/decisions/item/instance/contactedreviewers/item"
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
func (m *InstanceRequestBuilder) AcceptRecommendations()(*i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f.AcceptRecommendationsRequestBuilder) {
    return i7f164e80748f8254d77062eb04b9fc70d24acae1f8ff55025019ff39b93f807f.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *InstanceRequestBuilder) ApplyDecisions()(*ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d.ApplyDecisionsRequestBuilder) {
    return ib32fea3d4d08f37811492a78820bd10fd5ed6a2b6ebe324905ab1721dcf9468d.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01.BatchRecordDecisionsRequestBuilder) {
    return i782c49834b7b944c9386d26b1447e1203c6c78b9faf1b852a88f039a1d2dcd01.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
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
func (m *InstanceRequestBuilder) ContactedReviewers()(*ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a.ContactedReviewersRequestBuilder) {
    return ie8f8ab4cd8e535acc957f1d84dfe870b17611bf063d1c7001c04bc5926b90f0a.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.stages.item.decisions.item.instance.contactedReviewers.item collection
func (m *InstanceRequestBuilder) ContactedReviewersById(id string)(*iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return iebc8a00efa2b50bd2e1e9d3b9cf13585be8d95411d38a44a5f752a1cdd32ea2b.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instance for users
func (m *InstanceRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instance for users
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
// CreateGetRequestInformationWithRequestConfiguration there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InstanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instance in users
func (m *InstanceRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instance in users
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
// Decisions the decisions property
func (m *InstanceRequestBuilder) Decisions()(*if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017.DecisionsRequestBuilder) {
    return if1bc878fe4076da3e59089857683f38dcd1d7348a735eaf6d2fa215d83f96017.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.stages.item.decisions.item.instance.decisions.item collection
func (m *InstanceRequestBuilder) DecisionsById(id string)(*i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = id
    }
    return i029ca4e7f214bf663eda08e436807f3035ef57d28cc79e27c6d721f043965f7c.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition the definition property
func (m *InstanceRequestBuilder) Definition()(*i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1.DefinitionRequestBuilder) {
    return i2b6e639b85012916127dc0d07246eb1ae49de6d165778c5f395e700affdd9ef1.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeleteWithResponseHandler delete navigation property instance for users
func (m *InstanceRequestBuilder) DeleteWithResponseHandler(requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property instance for users
func (m *InstanceRequestBuilder) DeleteWithResponseHandler(requestConfiguration *InstanceRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) GetWithResponseHandler(requestConfiguration *InstanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) GetWithResponseHandler(requestConfiguration *InstanceRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// PatchWithResponseHandler update the navigation property instance in users
func (m *InstanceRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property instance in users
func (m *InstanceRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, requestConfiguration *InstanceRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *InstanceRequestBuilder) ResetDecisions()(*ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d.ResetDecisionsRequestBuilder) {
    return ibc590043d3d2dc48e7d0a988e65b8d6ce44a64055edec92bbfb7f460a581677d.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *InstanceRequestBuilder) SendReminder()(*id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e.SendReminderRequestBuilder) {
    return id55ed28e16edbddd855986652f8486f24b2ce428790808ae55a7be780db9fb0e.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop the stop property
func (m *InstanceRequestBuilder) Stop()(*i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8.StopRequestBuilder) {
    return i4201497ed84875eef39ef0e9f79d2e857848f718dba808ab7cdb040dfdd886e8.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
