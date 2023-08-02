package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder provides operations to call the setUpFeedbackResourcesFolder method.
type ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderInternal instantiates a new SetUpFeedbackResourcesFolderRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) {
    m := &ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/setUpFeedbackResourcesFolder", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder instantiates a new SetUpFeedbackResourcesFolderRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a SharePoint folder to upload feedback files for a given educationSubmission. Only teachers can perform this operation. The teacher determines the resources to upload in the feedback resources folder of a submission.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignment-setupfeedbackresourcesfolder?view=graph-rest-1.0
func (m *ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) Post(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
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
// ToPostRequestInformation create a SharePoint folder to upload feedback files for a given educationSubmission. Only teachers can perform this operation. The teacher determines the resources to upload in the feedback resources folder of a submission.
func (m *ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSetUpFeedbackResourcesFolderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
