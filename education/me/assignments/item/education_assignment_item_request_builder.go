package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/setupresourcesfolder"
    i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions"
    i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources"
    i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories"
    iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/rubric"
    iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/publish"
    icaf5ade0d32878478f4fc3b3161eff9ab10621464aed9f1562c1f92c10d70e8c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/setupfeedbackresourcesfolder"
    i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item"
    i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources/item"
    ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories/item"
)

// EducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type EducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type EducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationAssignmentItemRequestBuilderGetQueryParameters
}
// EducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories provides operations to manage the categories property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) Categories()(*i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.CategoriesRequestBuilder) {
    return i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.categories.item collection
func (m *EducationAssignmentItemRequestBuilder) CategoriesById(id string)(*ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    m := &EducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *EducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *EducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Patch update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, requestConfiguration *EducationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationAssignmentable), nil
}
// Publish provides operations to call the publish method.
func (m *EducationAssignmentItemRequestBuilder) Publish()(*iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.PublishRequestBuilder) {
    return iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) Resources()(*i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.ResourcesRequestBuilder) {
    return i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) ResourcesById(id string)(*i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.EducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.NewEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Rubric provides operations to manage the rubric property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) Rubric()(*iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.RubricRequestBuilder) {
    return iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpFeedbackResourcesFolder provides operations to call the setUpFeedbackResourcesFolder method.
func (m *EducationAssignmentItemRequestBuilder) SetUpFeedbackResourcesFolder()(*icaf5ade0d32878478f4fc3b3161eff9ab10621464aed9f1562c1f92c10d70e8c.SetUpFeedbackResourcesFolderRequestBuilder) {
    return icaf5ade0d32878478f4fc3b3161eff9ab10621464aed9f1562c1f92c10d70e8c.NewSetUpFeedbackResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder provides operations to call the setUpResourcesFolder method.
func (m *EducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.SetUpResourcesFolderRequestBuilder) {
    return i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submissions provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) Submissions()(*i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.SubmissionsRequestBuilder) {
    return i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
func (m *EducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.EducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.NewEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
