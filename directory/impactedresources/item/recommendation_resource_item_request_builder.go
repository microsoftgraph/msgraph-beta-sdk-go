package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/dismiss"
    i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/complete"
    i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/reactivate"
    ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/impactedresources/item/postpone"
)

// RecommendationResourceItemRequestBuilder provides operations to manage the impactedResources property of the microsoft.graph.directory entity.
type RecommendationResourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RecommendationResourceItemRequestBuilderDeleteOptions options for Delete
type RecommendationResourceItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RecommendationResourceItemRequestBuilderGetOptions options for Get
type RecommendationResourceItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RecommendationResourceItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// RecommendationResourceItemRequestBuilderGetQueryParameters get impactedResources from directory
type RecommendationResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RecommendationResourceItemRequestBuilderPatchOptions options for Patch
type RecommendationResourceItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecommendationResourceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Complete the complete property
func (m *RecommendationResourceItemRequestBuilder) Complete()(*i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6.CompleteRequestBuilder) {
    return i6a908001799100452fb9272829db83866d7952f0456ff214df6a4849088023c6.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRecommendationResourceItemRequestBuilderInternal instantiates a new RecommendationResourceItemRequestBuilder and sets the default values.
func NewRecommendationResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationResourceItemRequestBuilder) {
    m := &RecommendationResourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/impactedResources/{recommendationResource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecommendationResourceItemRequestBuilder instantiates a new RecommendationResourceItemRequestBuilder and sets the default values.
func NewRecommendationResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property impactedResources for directory
func (m *RecommendationResourceItemRequestBuilder) CreateDeleteRequestInformation(options *RecommendationResourceItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get impactedResources from directory
func (m *RecommendationResourceItemRequestBuilder) CreateGetRequestInformation(options *RecommendationResourceItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property impactedResources in directory
func (m *RecommendationResourceItemRequestBuilder) CreatePatchRequestInformation(options *RecommendationResourceItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property impactedResources for directory
func (m *RecommendationResourceItemRequestBuilder) Delete(options *RecommendationResourceItemRequestBuilderDeleteOptions)(error) {
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
// Dismiss the dismiss property
func (m *RecommendationResourceItemRequestBuilder) Dismiss()(*i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e.DismissRequestBuilder) {
    return i02454f4537287b6ab338217acb9968bd694788945e44d84da9b93503fb407a4e.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get impactedResources from directory
func (m *RecommendationResourceItemRequestBuilder) Get(options *RecommendationResourceItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecommendationResourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRecommendationResourceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RecommendationResourceable), nil
}
// Patch update the navigation property impactedResources in directory
func (m *RecommendationResourceItemRequestBuilder) Patch(options *RecommendationResourceItemRequestBuilderPatchOptions)(error) {
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
// Postpone the postpone property
func (m *RecommendationResourceItemRequestBuilder) Postpone()(*ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443.PostponeRequestBuilder) {
    return ie32cb720402ff9dd0b35e1e591941d01f5c513d30bea76e934259b963be0b443.NewPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reactivate the reactivate property
func (m *RecommendationResourceItemRequestBuilder) Reactivate()(*i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd.ReactivateRequestBuilder) {
    return i9a2ce2cc07814668197ce7e2daba99556437bce33769c58c3fa88ff13ad02bdd.NewReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
