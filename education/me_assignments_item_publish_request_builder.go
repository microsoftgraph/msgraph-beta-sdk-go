package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemPublishRequestBuilder provides operations to call the publish method.
type MeAssignmentsItemPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemPublishRequestBuilderInternal instantiates a new PublishRequestBuilder and sets the default values.
func NewMeAssignmentsItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemPublishRequestBuilder) {
    m := &MeAssignmentsItemPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/publish", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemPublishRequestBuilder instantiates a new PublishRequestBuilder and sets the default values.
func NewMeAssignmentsItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post change the state of an educationAssignment from its original draft status to the published status.  You can change the state from draft to scheduled if the assignment is scheduled for a future date.  Only a teacher in the class can make this call. When an assignment is in draft status, students will not see the assignment, nor will there be any submission objects. When you call this API, educationSubmission objects are created and the assignment appears in the student's list. The state of the assignment goes back to draft if there is any backend failure during publish process. To update the properties of a published assignment, see update an assignment. To update the properties of a published assignment, see update an assignment. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignment-publish?view=graph-rest-1.0
func (m *MeAssignmentsItemPublishRequestBuilder) Post(ctx context.Context, requestConfiguration *MeAssignmentsItemPublishRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// ToPostRequestInformation change the state of an educationAssignment from its original draft status to the published status.  You can change the state from draft to scheduled if the assignment is scheduled for a future date.  Only a teacher in the class can make this call. When an assignment is in draft status, students will not see the assignment, nor will there be any submission objects. When you call this API, educationSubmission objects are created and the assignment appears in the student's list. The state of the assignment goes back to draft if there is any backend failure during publish process. To update the properties of a published assignment, see update an assignment. To update the properties of a published assignment, see update an assignment. This API is available in the following national cloud deployments.
func (m *MeAssignmentsItemPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MeAssignmentsItemPublishRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemPublishRequestBuilder) {
    return NewMeAssignmentsItemPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
