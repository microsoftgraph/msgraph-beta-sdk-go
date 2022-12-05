package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder provides operations to manage the collection of onPremisesPublishingProfile entities.
type OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property memberOf for onPremisesPublishingProfiles
type OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteQueryParameters
}
// NewOnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/connectors/{connector%2Did}/memberOf/{connectorGroup%2Did}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete ref of navigation property memberOf for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete ref of navigation property memberOf for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemConnectorsItemMemberOfItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
