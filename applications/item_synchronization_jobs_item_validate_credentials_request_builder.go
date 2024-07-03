package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemValidateCredentialsRequestBuilder provides operations to call the validateCredentials method.
type ItemSynchronizationJobsItemValidateCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemValidateCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemValidateCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemValidateCredentialsRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemValidateCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) {
    m := &ItemSynchronizationJobsItemValidateCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/validateCredentials", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemValidateCredentialsRequestBuilder instantiates a new ItemSynchronizationJobsItemValidateCredentialsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemValidateCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemValidateCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that the credentials are valid in the tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationjob-validatecredentials?view=graph-rest-beta
func (m *ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemValidateCredentialsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation validate that the credentials are valid in the tenant.
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemValidateCredentialsPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemValidateCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSynchronizationJobsItemValidateCredentialsRequestBuilder when successful
func (m *ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemValidateCredentialsRequestBuilder) {
    return NewItemSynchronizationJobsItemValidateCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
