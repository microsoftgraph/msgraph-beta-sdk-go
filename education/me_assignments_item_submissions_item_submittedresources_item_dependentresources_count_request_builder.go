package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder provides operations to count the resources in the collection.
type MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetQueryParameters get the number of the resource
type MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetQueryParameters
}
// NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderInternal instantiates a new MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) {
    m := &MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}/dependentResources/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder instantiates a new MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder when successful
func (m *MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) WithUrl(rawUrl string)(*MeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder) {
    return NewMeAssignmentsItemSubmissionsItemSubmittedresourcesItemDependentresourcesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
