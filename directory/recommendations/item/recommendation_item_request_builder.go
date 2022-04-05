package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/complete"
    i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/reactivate"
    i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources"
    ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/postpone"
    ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/dismiss"
    i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/recommendations/item/impactedresources/item"
)

// RecommendationItemRequestBuilder provides operations to manage the recommendations property of the microsoft.graph.directory entity.
type RecommendationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RecommendationItemRequestBuilderDeleteOptions options for Delete
type RecommendationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// RecommendationItemRequestBuilderGetOptions options for Get
type RecommendationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *RecommendationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// RecommendationItemRequestBuilderGetQueryParameters get recommendations from directory
type RecommendationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RecommendationItemRequestBuilderPatchOptions options for Patch
type RecommendationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Complete the complete property
func (m *RecommendationItemRequestBuilder) Complete()(*i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432.CompleteRequestBuilder) {
    return i4018c02f3fbf5218282cea136e679104019a4c34f3c1902f574244076c8fb432.NewCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRecommendationItemRequestBuilderInternal instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewRecommendationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationItemRequestBuilder) {
    m := &RecommendationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/recommendations/{recommendation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRecommendationItemRequestBuilder instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewRecommendationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property recommendations for directory
func (m *RecommendationItemRequestBuilder) CreateDeleteRequestInformation(options *RecommendationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get recommendations from directory
func (m *RecommendationItemRequestBuilder) CreateGetRequestInformation(options *RecommendationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property recommendations in directory
func (m *RecommendationItemRequestBuilder) CreatePatchRequestInformation(options *RecommendationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property recommendations for directory
func (m *RecommendationItemRequestBuilder) Delete(options *RecommendationItemRequestBuilderDeleteOptions)(error) {
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
func (m *RecommendationItemRequestBuilder) Dismiss()(*ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a.DismissRequestBuilder) {
    return ib872c308280e73f616388d32a30c0860fa7b0ee0f18b5fafe036548766c8a60a.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get recommendations from directory
func (m *RecommendationItemRequestBuilder) Get(options *RecommendationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRecommendationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable), nil
}
// ImpactedResources the impactedResources property
func (m *RecommendationItemRequestBuilder) ImpactedResources()(*i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6.ImpactedResourcesRequestBuilder) {
    return i9159a081c2a0d0abecb435cb8c52d12d3cb4c5cc3be944e4400678fa5eaf97f6.NewImpactedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImpactedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.directory.recommendations.item.impactedResources.item collection
func (m *RecommendationItemRequestBuilder) ImpactedResourcesById(id string)(*i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b.RecommendationResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendationResource_id"] = id
    }
    return i15d57ca0f29006194bb50f1f3633607a5d2e31816ad92ef8c6d5faf215b3732b.NewRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property recommendations in directory
func (m *RecommendationItemRequestBuilder) Patch(options *RecommendationItemRequestBuilderPatchOptions)(error) {
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
func (m *RecommendationItemRequestBuilder) Postpone()(*ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017.PostponeRequestBuilder) {
    return ib1bc9d8bdb20f875e5c28e8e4d42895f9148cf2ba25fb010132a667a01115017.NewPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reactivate the reactivate property
func (m *RecommendationItemRequestBuilder) Reactivate()(*i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da.ReactivateRequestBuilder) {
    return i85047cb55fafd88055bcdeb243251139d17240509426cf420901b907aa3f73da.NewReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
