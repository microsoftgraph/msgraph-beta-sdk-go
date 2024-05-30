package privilegedaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder provides operations to call the updateRequest method.
type ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderInternal instantiates a new ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) {
    m := &ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedAccess/{privilegedAccess%2Did}/resources/{governanceResource%2Did}/roleAssignmentRequests/{governanceRoleAssignmentRequest%2Did}/updateRequest", pathParameters),
    }
    return m
}
// NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder instantiates a new ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder and sets the default values.
func NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateRequest
// returns a GovernanceRoleAssignmentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) Post(ctx context.Context, body ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestPostRequestBodyable, requestConfiguration *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGovernanceRoleAssignmentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GovernanceRoleAssignmentRequestable), nil
}
// ToPostRequestInformation invoke action updateRequest
// returns a *RequestInformation when successful
func (m *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestPostRequestBodyable, requestConfiguration *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder when successful
func (m *ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) WithUrl(rawUrl string)(*ItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder) {
    return NewItemResourcesItemRoleassignmentrequestsItemUpdaterequestUpdateRequestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
