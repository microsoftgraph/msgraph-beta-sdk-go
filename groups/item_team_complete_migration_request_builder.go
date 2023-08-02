package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type ItemTeamCompleteMigrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamCompleteMigrationRequestBuilderInternal instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewItemTeamCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamCompleteMigrationRequestBuilder) {
    m := &ItemTeamCompleteMigrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/completeMigration", pathParameters),
    }
    return m
}
// NewItemTeamCompleteMigrationRequestBuilder instantiates a new CompleteMigrationRequestBuilder and sets the default values.
func NewItemTeamCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete the message migration process by removing migration mode from a team. Migration mode is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a completeMigration request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-completemigration?view=graph-rest-1.0
func (m *ItemTeamCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation complete the message migration process by removing migration mode from a team. Migration mode is a special state where certain operations are barred, like message POST and membership operations during the data migration process. After a completeMigration request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *ItemTeamCompleteMigrationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
