package audience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members"
    i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/updateaudience"
    i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions"
    ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/updateaudiencebyid"
    icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/members/item"
    ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/deployments/item/audience/exclusions/item"
)

// AudienceRequestBuilder provides operations to manage the audience property of the microsoft.graph.windowsUpdates.deployment entity.
type AudienceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AudienceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AudienceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AudienceRequestBuilderGetQueryParameters specifies the audience to which content is deployed.
type AudienceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AudienceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AudienceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AudienceRequestBuilderGetQueryParameters
}
// AudienceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AudienceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAudienceRequestBuilderInternal instantiates a new AudienceRequestBuilder and sets the default values.
func NewAudienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AudienceRequestBuilder) {
    m := &AudienceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAudienceRequestBuilder instantiates a new AudienceRequestBuilder and sets the default values.
func NewAudienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AudienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAudienceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property audience for admin
func (m *AudienceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AudienceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AudienceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property audience in admin
func (m *AudienceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *AudienceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property audience for admin
func (m *AudienceRequestBuilder) Delete(ctx context.Context, requestConfiguration *AudienceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Exclusions provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AudienceRequestBuilder) Exclusions()(*i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde.ExclusionsRequestBuilder) {
    return i9d78b5d7e63632aad0dc14fc18eec733cc1a34e41c1bd899b904db2431e7bbde.NewExclusionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExclusionsById provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AudienceRequestBuilder) ExclusionsById(id string)(*ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac.UpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return ifa1c023a93cedbc885180b644196fd22db8e7768dfa64cd6a803d74884273eac.NewUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get specifies the audience to which content is deployed.
func (m *AudienceRequestBuilder) Get(ctx context.Context, requestConfiguration *AudienceRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AudienceRequestBuilder) Members()(*i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e.MembersRequestBuilder) {
    return i27c073677df2af7a178b750045cf7538776463d44b1fdb96205037112761934e.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
func (m *AudienceRequestBuilder) MembersById(id string)(*icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058.UpdatableAssetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["updatableAsset%2Did"] = id
    }
    return icf6ee88716e5da88403b56c76157382dd11930a0110b7137be42bf2fb2425058.NewUpdatableAssetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property audience in admin
func (m *AudienceRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, requestConfiguration *AudienceRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateDeploymentAudienceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.DeploymentAudienceable), nil
}
// UpdateAudience provides operations to call the updateAudience method.
func (m *AudienceRequestBuilder) UpdateAudience()(*i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426.UpdateAudienceRequestBuilder) {
    return i6ca7350a454d0eb7416a3ffe6d406c2f9a011b9818e8a5fc6714fbfab983c426.NewUpdateAudienceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateAudienceById provides operations to call the updateAudienceById method.
func (m *AudienceRequestBuilder) UpdateAudienceById()(*ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28.UpdateAudienceByIdRequestBuilder) {
    return ibb560c7d9d162fcf5a50b3a20535d73efbfb6eada7792589ef7f79232af0cd28.NewUpdateAudienceByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
