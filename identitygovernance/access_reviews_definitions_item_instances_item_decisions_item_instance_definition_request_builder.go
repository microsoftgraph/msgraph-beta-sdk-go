package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetQueryParameters
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/definition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// Get there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable), nil
}
