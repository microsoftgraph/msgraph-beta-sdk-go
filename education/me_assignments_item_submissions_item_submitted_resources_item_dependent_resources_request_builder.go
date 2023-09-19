package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder provides operations to manage the dependentResources property of the microsoft.graph.educationSubmissionResource entity.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetQueryParameters get dependentResources from education
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetQueryParameters
}
// MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEducationSubmissionResourceId1 provides operations to manage the dependentResources property of the microsoft.graph.educationSubmissionResource entity.
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) ByEducationSubmissionResourceId1(educationSubmissionResourceId1 string)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if educationSubmissionResourceId1 != "" {
        urlTplParams["educationSubmissionResource%2Did1"] = educationSubmissionResourceId1
    }
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderInternal instantiates a new DependentResourcesRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) {
    m := &MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}/dependentResources{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder instantiates a new DependentResourcesRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) Count()(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesCountRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dependentResources from education
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionResourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceCollectionResponseable), nil
}
// Post create new navigation property to dependentResources for education
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get dependentResources from education
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to dependentResources for education
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemSubmittedResourcesItemDependentResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
