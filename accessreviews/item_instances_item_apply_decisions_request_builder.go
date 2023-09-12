package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesItemApplyDecisionsRequestBuilder provides operations to call the applyDecisions method.
type ItemInstancesItemApplyDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstancesItemApplyDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesItemApplyDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInstancesItemApplyDecisionsRequestBuilderInternal instantiates a new ApplyDecisionsRequestBuilder and sets the default values.
func NewItemInstancesItemApplyDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemApplyDecisionsRequestBuilder) {
    m := &ItemInstancesItemApplyDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}/applyDecisions", pathParameters),
    }
    return m
}
// NewItemInstancesItemApplyDecisionsRequestBuilder instantiates a new ApplyDecisionsRequestBuilder and sets the default values.
func NewItemInstancesItemApplyDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemApplyDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesItemApplyDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in the Azure AD access reviews feature, apply the decisions of a completed accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review.   After an access review is finished, either because it reached the end date or an administrator stopped it manually, and auto-apply wasn't configured for the review, you can call Apply to apply the changes. Until apply occurs, the decisions to remove access rights do not appear on the source resource, the users for instance retain their group memberships. By calling apply, the outcome of the review is implemented by updating the group or application. If a user's access was denied in the review, when an administrator calls this API, Azure AD removes their membership or application assignment.  After an access review is finished, and auto-apply was configured, then the status of the review will change from Completed through intermediate states and finally will change to state Applied. You should expect to see denied users, if any, being removed from the resource group membership or app assignment in a few minutes. A configured auto applying review, or selecting Apply doesn't have an effect on a group that originates in an on-premises directory or a dynamic group. If you want to change a group that originates on-premises, download the results and apply those changes to the representation of the group in that directory.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreview-apply?view=graph-rest-1.0
func (m *ItemInstancesItemApplyDecisionsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInstancesItemApplyDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation in the Azure AD access reviews feature, apply the decisions of a completed accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review.   After an access review is finished, either because it reached the end date or an administrator stopped it manually, and auto-apply wasn't configured for the review, you can call Apply to apply the changes. Until apply occurs, the decisions to remove access rights do not appear on the source resource, the users for instance retain their group memberships. By calling apply, the outcome of the review is implemented by updating the group or application. If a user's access was denied in the review, when an administrator calls this API, Azure AD removes their membership or application assignment.  After an access review is finished, and auto-apply was configured, then the status of the review will change from Completed through intermediate states and finally will change to state Applied. You should expect to see denied users, if any, being removed from the resource group membership or app assignment in a few minutes. A configured auto applying review, or selecting Apply doesn't have an effect on a group that originates in an on-premises directory or a dynamic group. If you want to change a group that originates on-premises, download the results and apply those changes to the representation of the group in that directory.
func (m *ItemInstancesItemApplyDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesItemApplyDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemInstancesItemApplyDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemInstancesItemApplyDecisionsRequestBuilder) {
    return NewItemInstancesItemApplyDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
