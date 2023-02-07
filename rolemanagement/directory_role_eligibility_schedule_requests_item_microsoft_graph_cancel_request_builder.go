package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder provides operations to call the cancel method.
type DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    m := &DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/microsoft.graph.cancel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder instantiates a new MicrosoftGraphCancelRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post immediately cancel a unifiedRoleEligibilityScheduleRequest that is in a `Granted` status, and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled unifiedRoleEligibilityScheduleRequest changes to `Revoked`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/unifiedroleeligibilityschedulerequest-cancel?view=graph-rest-1.0
func (m *DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation immediately cancel a unifiedRoleEligibilityScheduleRequest that is in a `Granted` status, and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled unifiedRoleEligibilityScheduleRequest changes to `Revoked`.
func (m *DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsItemMicrosoftGraphCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
