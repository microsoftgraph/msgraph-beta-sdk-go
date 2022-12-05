package onpremisespublishingprofiles

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder provides operations to manage the collection of onPremisesPublishingProfile entities.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property agentGroups for onPremisesPublishingProfiles
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteQueryParameters
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder) {
    m := &OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile%2Did}/agentGroups/{onPremisesAgentGroup%2Did}/agents/{onPremisesAgent%2Did}/agentGroups/{onPremisesAgentGroup%2Did1}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete ref of navigation property agentGroups for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete ref of navigation property agentGroups for onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnPremisesPublishingProfilesItemAgentGroupsItemAgentsItemAgentGroupsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
