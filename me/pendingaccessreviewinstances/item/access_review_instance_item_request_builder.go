package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/batchrecorddecisions"
    i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/contactedreviewers"
    i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stop"
    i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages"
    i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/applydecisions"
    i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions"
    i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/resetdecisions"
    ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/acceptrecommendations"
    ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/definition"
    icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/sendreminder"
    i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/decisions/item"
    ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/contactedreviewers/item"
    ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item/stages/item"
)

// AccessReviewInstanceItemRequestBuilder provides operations to manage the pendingAccessReviewInstances property of the microsoft.graph.user entity.
type AccessReviewInstanceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessReviewInstanceItemRequestBuilderDeleteOptions options for Delete
type AccessReviewInstanceItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessReviewInstanceItemRequestBuilderGetOptions options for Get
type AccessReviewInstanceItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AccessReviewInstanceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AccessReviewInstanceItemRequestBuilderGetQueryParameters navigation property to get list of access reviews pending approval by reviewer.
type AccessReviewInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AccessReviewInstanceItemRequestBuilderPatchOptions options for Patch
type AccessReviewInstanceItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AcceptRecommendations the acceptRecommendations property
func (m *AccessReviewInstanceItemRequestBuilder) AcceptRecommendations()(*ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628.AcceptRecommendationsRequestBuilder) {
    return ia0d94a184a000bb845f57559a39262d0cde1e4cbd62ba5b70e7f517eff968628.NewAcceptRecommendationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplyDecisions the applyDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) ApplyDecisions()(*i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8.ApplyDecisionsRequestBuilder) {
    return i78cfe79793d77a5fede4967db4de1713d8be67261324f74669c2af685e4134f8.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BatchRecordDecisions the batchRecordDecisions property
func (m *AccessReviewInstanceItemRequestBuilder) BatchRecordDecisions()(*i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656.BatchRecordDecisionsRequestBuilder) {
    return i226aac2d9e753eec2396e7a04e06f10afaa0784792e00d5fe99fea4069422656.NewBatchRecordDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewInstanceItemRequestBuilderInternal instantiates a new AccessReviewInstanceItemRequestBuilder and sets the default values.
func NewAccessReviewInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewInstanceItemRequestBuilder) {
    m := &AccessReviewInstanceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance_id}{?select,expand}";
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
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewers()(*i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04.ContactedReviewersRequestBuilder) {
    return i28cbfe53788555c1e8fb2bcbb1f7b5ee3a25273656373f13f5ff7b954d036e04.NewContactedReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactedReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.contactedReviewers.item collection
func (m *AccessReviewInstanceItemRequestBuilder) ContactedReviewersById(id string)(*ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return ia273f0c967440d14bb94006b177f70e6b83152c6fcad8ccbf86dc8739482bc95.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pendingAccessReviewInstances for me
func (m *AccessReviewInstanceItemRequestBuilder) CreateDeleteRequestInformation(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceItemRequestBuilder) CreateGetRequestInformation(options *AccessReviewInstanceItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pendingAccessReviewInstances in me
func (m *AccessReviewInstanceItemRequestBuilder) CreatePatchRequestInformation(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decisions the decisions property
func (m *AccessReviewInstanceItemRequestBuilder) Decisions()(*i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300.DecisionsRequestBuilder) {
    return i857369f7b2494f5fc984108c41f95436f6b0349231c74000340f9f1547e30300.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.decisions.item collection
func (m *AccessReviewInstanceItemRequestBuilder) DecisionsById(id string)(*i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e.AccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstanceDecisionItem_id"] = id
    }
    return i92a83f24221d12aedaef668e1fbb993200a107c50dd873c65f56c93ab8abaa7e.NewAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Definition the definition property
func (m *AccessReviewInstanceItemRequestBuilder) Definition()(*ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277.DefinitionRequestBuilder) {
    return ia8e711f9bce6e1855eb0b9ef5889095dbc6f1dff81d7c5cb1530e08f1e8af277.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property pendingAccessReviewInstances for me
func (m *AccessReviewInstanceItemRequestBuilder) Delete(options *AccessReviewInstanceItemRequestBuilderDeleteOptions)(error) {
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
// Get navigation property to get list of access reviews pending approval by reviewer.
func (m *AccessReviewInstanceItemRequestBuilder) Get(options *AccessReviewInstanceItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceable, error) {
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
// Patch update the navigation property pendingAccessReviewInstances in me
func (m *AccessReviewInstanceItemRequestBuilder) Patch(options *AccessReviewInstanceItemRequestBuilderPatchOptions)(error) {
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
func (m *AccessReviewInstanceItemRequestBuilder) ResetDecisions()(*i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e.ResetDecisionsRequestBuilder) {
    return i978a7302921342acb736ff821abff5eb57189f5f20e0b9d44a2800bd91c9bc2e.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *AccessReviewInstanceItemRequestBuilder) SendReminder()(*icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6.SendReminderRequestBuilder) {
    return icbbe17c4811e8e4855cc7841cf8914b2f2a20b1d41e16779ac8b42dd4ad533a6.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stages the stages property
func (m *AccessReviewInstanceItemRequestBuilder) Stages()(*i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108.StagesRequestBuilder) {
    return i6be43a98526e3cd95557ccc422ad7cdaa8c43585327deaffa99f4367f522a108.NewStagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// StagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item.stages.item collection
func (m *AccessReviewInstanceItemRequestBuilder) StagesById(id string)(*ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f.AccessReviewStageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewStage_id"] = id
    }
    return ie9afc88f4d0ebd933223975c8172435f458acf5c8c922ff0f1d55f973e0ed66f.NewAccessReviewStageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Stop the stop property
func (m *AccessReviewInstanceItemRequestBuilder) Stop()(*i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f.StopRequestBuilder) {
    return i5a33caa17867d8a52a52d95e786e9714ce0e9a6b07f5e9aa14f00f581f95c45f.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
