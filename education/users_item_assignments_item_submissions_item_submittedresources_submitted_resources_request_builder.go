package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
type UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetQueryParameters get submittedResources from education
type UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetQueryParameters struct {
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
// UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetQueryParameters
}
// UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEducationSubmissionResourceId provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
// returns a *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder when successful
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) ByEducationSubmissionResourceId(educationSubmissionResourceId string)(*UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if educationSubmissionResourceId != "" {
        urlTplParams["educationSubmissionResource%2Did"] = educationSubmissionResourceId
    }
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderInternal instantiates a new UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) {
    m := &UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder instantiates a new UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder and sets the default values.
func NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesCountRequestBuilder when successful
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) Count()(*UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesCountRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get submittedResources from education
// returns a EducationSubmissionResourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Post create new navigation property to submittedResources for education
// returns a EducationSubmissionResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get submittedResources from education
// returns a *RequestInformation when successful
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to submittedResources for education
// returns a *RequestInformation when successful
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionResourceable, requestConfiguration *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder when successful
func (m *UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) WithUrl(rawUrl string)(*UsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder) {
    return NewUsersItemAssignmentsItemSubmissionsItemSubmittedresourcesSubmittedResourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
