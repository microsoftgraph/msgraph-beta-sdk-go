package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder provides operations to manage the dependentResources property of the microsoft.graph.educationSubmissionResource entity.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters get dependentResources from education
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderInternal instantiates a new MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) {
    m := &MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}/dependentResources/{educationSubmissionResource%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder instantiates a new MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property dependentResources for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get dependentResources from education
// returns a EducationSubmissionResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable), nil
}
// Patch update the navigation property dependentResources in education
// returns a EducationSubmissionResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable), nil
}
// ToDeleteRequestInformation delete navigation property dependentResources for education
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}/dependentResources/{educationSubmissionResource%2Did1}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get dependentResources from education
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property dependentResources in education
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}/dependentResources/{educationSubmissionResource%2Did1}", m.BaseRequestBuilder.PathParameters)
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
// returns a *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
