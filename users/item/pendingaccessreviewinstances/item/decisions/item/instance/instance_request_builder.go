package instance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1959da9b23611a66111bba1103551538eaf9d8af14a2efceee1d0b8668dafa84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/stages"
    i22d1a30d9891fbc88c7630a69af15ccf11ce00473aa9a4206d9e31c8467e4105 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/applydecisions"
    i316785d27a0a31011d7966b83fd798f8077fa7d13aa5ca2e68097d831311c04d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/definition"
    i54a536afaf234612b42685996cf5d529ec881cbd332daec0af1f3028dfe49150 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/contactedreviewers"
    i586bd7c3dc27d304fc22c053703c5efe70c85fb02c776dd70a1cb912db05d294 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/stop"
    i97b9f71efab69fc1759cd914522f0dbc62af593a5edfe352e618e082fcdefa94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/sendreminder"
    ia9fce0c5c72025a51923562060dad8be0fda0f3f5af172754bfb2c7b752fdf5b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/batchrecorddecisions"
    ibe8dd3fcfec86d1d540afc1fe71ed45c15c81b82a4be534e59b34ca0f753aaaf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/resetdecisions"
    ic94443bd6d55294f6041d27ac88f5e3cf9781bb5b3040fe63a636a8415f6de0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/acceptrecommendations"
    i6ed9d28624c9122c38ba0dd2b1c13ce5c4f810e43240e2c63deeae20fdeeb56e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/stages/item"
    i9d55cc900f8ed2ffb05a846f05759d58df50ab72bc742ff92a55307c88592116 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/pendingaccessreviewinstances/item/decisions/item/instance/contactedreviewers/item"
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
// InstanceRequestBuilderDeleteOptions options for Delete
type InstanceRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// InstanceRequestBuilderGetOptions options for Get
type InstanceRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstanceRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// InstanceRequestBuilderGetQueryParameters there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
type InstanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InstanceRequestBuilderPatchOptions options for Patch
type InstanceRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AcceptRecommendations the acceptRecommendations property
func (m *InstanceRequestBuilder) AcceptRecommendations()(*ic94443bd6d55294f6041d27ac88f5e3cf9781bb5b3040fe63a636a8415f6de0a.AcceptRecommendationsRequestBuilder) {
    return ic94443bd6d55294f6041d27ac88f5e3cf9781bb5b3040fe63a636a8415f6de0a.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *InstanceRequestBuilder) ApplyDecisions()(*i22d1a30d9891fbc88c7630a69af15ccf11ce00473aa9a4206d9e31c8467e4105.ApplyDecisionsRequestBuilder) {
    return i22d1a30d9891fbc88c7630a69af15ccf11ce00473aa9a4206d9e31c8467e4105.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *InstanceRequestBuilder) BatchRecordDecisions()(*ia9fce0c5c72025a51923562060dad8be0fda0f3f5af172754bfb2c7b752fdf5b.BatchRecordDecisionsRequestBuilder) {
    return ia9fce0c5c72025a51923562060dad8be0fda0f3f5af172754bfb2c7b752fdf5b.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInstanceRequestBuilderInternal instantiates a new InstanceRequestBuilder and sets the default values.
func NewInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstanceRequestBuilder) {
    m := &InstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance{?%24select,%24expand}";
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
func (m *InstanceRequestBuilder) ContactedReviewers()(*i54a536afaf234612b42685996cf5d529ec881cbd332daec0af1f3028dfe49150.ContactedReviewersRequestBuilder) {
    return i54a536afaf234612b42685996cf5d529ec881cbd332daec0af1f3028dfe49150.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.decisions.item.instance.contactedReviewers.item collection
func (m *InstanceRequestBuilder) ContactedReviewersById(id string)(*i9d55cc900f8ed2ffb05a846f05759d58df50ab72bc742ff92a55307c88592116.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return i9d55cc900f8ed2ffb05a846f05759d58df50ab72bc742ff92a55307c88592116.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property instance for users
func (m *InstanceRequestBuilder) CreateDeleteRequestInformation(options *InstanceRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) CreateGetRequestInformation(options *InstanceRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instance in users
func (m *InstanceRequestBuilder) CreatePatchRequestInformation(options *InstanceRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Definition the definition property
func (m *InstanceRequestBuilder) Definition()(*i316785d27a0a31011d7966b83fd798f8077fa7d13aa5ca2e68097d831311c04d.DefinitionRequestBuilder) {
    return i316785d27a0a31011d7966b83fd798f8077fa7d13aa5ca2e68097d831311c04d.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instance for users
func (m *InstanceRequestBuilder) Delete(options *InstanceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get there is exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the recurrence of the access review the decision is made on.
func (m *InstanceRequestBuilder) Get(options *InstanceRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable), nil
}
// Patch update the navigation property instance in users
func (m *InstanceRequestBuilder) Patch(options *InstanceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ResetDecisions the resetDecisions property
func (m *InstanceRequestBuilder) ResetDecisions()(*ibe8dd3fcfec86d1d540afc1fe71ed45c15c81b82a4be534e59b34ca0f753aaaf.ResetDecisionsRequestBuilder) {
    return ibe8dd3fcfec86d1d540afc1fe71ed45c15c81b82a4be534e59b34ca0f753aaaf.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *InstanceRequestBuilder) SendReminder()(*i97b9f71efab69fc1759cd914522f0dbc62af593a5edfe352e618e082fcdefa94.SendReminderRequestBuilder) {
    return i97b9f71efab69fc1759cd914522f0dbc62af593a5edfe352e618e082fcdefa94.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages the stages property
func (m *InstanceRequestBuilder) Stages()(*i1959da9b23611a66111bba1103551538eaf9d8af14a2efceee1d0b8668dafa84.StagesRequestBuilder) {
    return i1959da9b23611a66111bba1103551538eaf9d8af14a2efceee1d0b8668dafa84.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.pendingAccessReviewInstances.item.decisions.item.instance.stages.item collection
func (m *InstanceRequestBuilder) StagesById(id string)(*i6ed9d28624c9122c38ba0dd2b1c13ce5c4f810e43240e2c63deeae20fdeeb56e.AccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage%2Did"] = id
    }
    return i6ed9d28624c9122c38ba0dd2b1c13ce5c4f810e43240e2c63deeae20fdeeb56e.NewAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop the stop property
func (m *InstanceRequestBuilder) Stop()(*i586bd7c3dc27d304fc22c053703c5efe70c85fb02c776dd70a1cb912db05d294.StopRequestBuilder) {
    return i586bd7c3dc27d304fc22c053703c5efe70c85fb02c776dd70a1cb912db05d294.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
