package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/mydecisions"
    i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/reviewers"
    i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/sendreminder"
    i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/stop"
    i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/resetdecisions"
    i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/applydecisions"
    ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances"
    ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/decisions"
    i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/decisions/item"
    i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item"
    i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/reviewers/item"
    ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/mydecisions/item"
)

// AccessReviewItemRequestBuilder provides operations to manage the collection of accessReview entities.
type AccessReviewItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewItemRequestBuilderGetQueryParameters in the Azure AD access reviews feature, retrieve an accessReview object.   To retrieve the reviewers of the access review, use the list accessReview reviewers API. To retrieve the decisions of the access review, use the list accessReview decisions API, or the list my accessReview decisions API. If this is a recurring access review, no decisions will be associated with the recurring access review series. Instead, use the `instances` relationship of that series to retrieve an accessReview collection of the past, current, and future instances of the access review. Each past and current instance will have decisions.
type AccessReviewItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewItemRequestBuilderGetQueryParameters
}
// AccessReviewItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyDecisions the applyDecisions property
func (m *AccessReviewItemRequestBuilder) ApplyDecisions()(*i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6.ApplyDecisionsRequestBuilder) {
    return i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAccessReviewItemRequestBuilderInternal instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewAccessReviewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewItemRequestBuilder) {
    m := &AccessReviewItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/accessReviews/{accessReview%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewItemRequestBuilder instantiates a new AccessReviewItemRequestBuilder and sets the default values.
func NewAccessReviewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation in the Azure AD access reviews feature, delete an accessReview object.
func (m *AccessReviewItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration in the Azure AD access reviews feature, delete an accessReview object.
func (m *AccessReviewItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AccessReviewItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation in the Azure AD access reviews feature, retrieve an accessReview object.   To retrieve the reviewers of the access review, use the list accessReview reviewers API. To retrieve the decisions of the access review, use the list accessReview decisions API, or the list my accessReview decisions API. If this is a recurring access review, no decisions will be associated with the recurring access review series. Instead, use the `instances` relationship of that series to retrieve an accessReview collection of the past, current, and future instances of the access review. Each past and current instance will have decisions.
func (m *AccessReviewItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration in the Azure AD access reviews feature, retrieve an accessReview object.   To retrieve the reviewers of the access review, use the list accessReview reviewers API. To retrieve the decisions of the access review, use the list accessReview decisions API, or the list my accessReview decisions API. If this is a recurring access review, no decisions will be associated with the recurring access review series. Instead, use the `instances` relationship of that series to retrieve an accessReview collection of the past, current, and future instances of the access review. Each past and current instance will have decisions.
func (m *AccessReviewItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AccessReviewItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation in the Azure AD access reviews feature, update an existing accessReview object to change one or more of its properties. This API is not intended to change the reviewers or decisions of a review.  To change the reviewers, use the addReviewer or removeReviewer APIs.  To stop an already-started one-time review, or an already-started instance of a recurring review, early, use the stop API. To apply the decisions to the target group or app access rights, use the apply API. 
func (m *AccessReviewItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration in the Azure AD access reviews feature, update an existing accessReview object to change one or more of its properties. This API is not intended to change the reviewers or decisions of a review.  To change the reviewers, use the addReviewer or removeReviewer APIs.  To stop an already-started one-time review, or an already-started instance of a recurring review, early, use the stop API. To apply the decisions to the target group or app access rights, use the apply API. 
func (m *AccessReviewItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *AccessReviewItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decisions the decisions property
func (m *AccessReviewItemRequestBuilder) Decisions()(*ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365.DecisionsRequestBuilder) {
    return ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviews.item.decisions.item collection
func (m *AccessReviewItemRequestBuilder) DecisionsById(id string)(*i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c.AccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete in the Azure AD access reviews feature, delete an accessReview object.
func (m *AccessReviewItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get in the Azure AD access reviews feature, retrieve an accessReview object.   To retrieve the reviewers of the access review, use the list accessReview reviewers API. To retrieve the decisions of the access review, use the list accessReview decisions API, or the list my accessReview decisions API. If this is a recurring access review, no decisions will be associated with the recurring access review series. Instead, use the `instances` relationship of that series to retrieve an accessReview collection of the past, current, and future instances of the access review. Each past and current instance will have decisions.
func (m *AccessReviewItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// Instances the instances property
func (m *AccessReviewItemRequestBuilder) Instances()(*ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021.InstancesRequestBuilder) {
    return ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviews.item.instances.item collection
func (m *AccessReviewItemRequestBuilder) InstancesById(id string)(*i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c.AccessReviewItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReview%2Did1"] = id
    }
    return i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c.NewAccessReviewItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MyDecisions the myDecisions property
func (m *AccessReviewItemRequestBuilder) MyDecisions()(*i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70.MyDecisionsRequestBuilder) {
    return i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70.NewMyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyDecisionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviews.item.myDecisions.item collection
func (m *AccessReviewItemRequestBuilder) MyDecisionsById(id string)(*ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4.AccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch in the Azure AD access reviews feature, update an existing accessReview object to change one or more of its properties. This API is not intended to change the reviewers or decisions of a review.  To change the reviewers, use the addReviewer or removeReviewer APIs.  To stop an already-started one-time review, or an already-started instance of a recurring review, early, use the stop API. To apply the decisions to the target group or app access rights, use the apply API. 
func (m *AccessReviewItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, requestConfiguration *AccessReviewItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewable), nil
}
// ResetDecisions the resetDecisions property
func (m *AccessReviewItemRequestBuilder) ResetDecisions()(*i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f.ResetDecisionsRequestBuilder) {
    return i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reviewers the reviewers property
func (m *AccessReviewItemRequestBuilder) Reviewers()(*i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92.ReviewersRequestBuilder) {
    return i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92.NewReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.accessReviews.item.reviewers.item collection
func (m *AccessReviewItemRequestBuilder) ReviewersById(id string)(*i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b.AccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewReviewer%2Did"] = id
    }
    return i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b.NewAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendReminder the sendReminder property
func (m *AccessReviewItemRequestBuilder) SendReminder()(*i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf.SendReminderRequestBuilder) {
    return i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop the stop property
func (m *AccessReviewItemRequestBuilder) Stop()(*i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0.StopRequestBuilder) {
    return i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
