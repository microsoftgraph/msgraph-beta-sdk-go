package orgcontact

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c194d9c76ac00401743b4274b2e641bbc11424974a100398e5b2b3dd33ce132 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item/orgcontact/checkmembergroups"
    i1c6c5601f055f8fda994b99255ce39947e22fb5c20a4991634feb6002bcced00 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item/orgcontact/getmemberobjects"
    i7fcf801c0e180273507b12a992a34c2109dcf80d1038933c0fc2a2b08affeada "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item/orgcontact/getmembergroups"
    i92107bebedaa23b60d8098d09f17a5dc81b4ca28fa13623b6ba4df811ca907ea "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item/orgcontact/restore"
    ib9e7d59bd6e6a0e37d0319cfe726eceb68d04ce0590ad367da479e6ceed3646a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item/orgcontact/checkmemberobjects"
)

// OrgContactRequestBuilder casts the previous resource to orgContact.
type OrgContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrgContactRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
type OrgContactRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrgContactRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrgContactRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrgContactRequestBuilderGetQueryParameters
}
// CheckMemberGroups the checkMemberGroups property
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*i0c194d9c76ac00401743b4274b2e641bbc11424974a100398e5b2b3dd33ce132.CheckMemberGroupsRequestBuilder) {
    return i0c194d9c76ac00401743b4274b2e641bbc11424974a100398e5b2b3dd33ce132.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*ib9e7d59bd6e6a0e37d0319cfe726eceb68d04ce0590ad367da479e6ceed3646a.CheckMemberObjectsRequestBuilder) {
    return ib9e7d59bd6e6a0e37d0319cfe726eceb68d04ce0590ad367da479e6ceed3646a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.orgContact{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrgContactRequestBuilder instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OrgContactRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i7fcf801c0e180273507b12a992a34c2109dcf80d1038933c0fc2a2b08affeada.GetMemberGroupsRequestBuilder) {
    return i7fcf801c0e180273507b12a992a34c2109dcf80d1038933c0fc2a2b08affeada.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactRequestBuilder) GetMemberObjects()(*i1c6c5601f055f8fda994b99255ce39947e22fb5c20a4991634feb6002bcced00.GetMemberObjectsRequestBuilder) {
    return i1c6c5601f055f8fda994b99255ce39947e22fb5c20a4991634feb6002bcced00.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OrgContactRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrgContactFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable), nil
}
// Restore the restore property
func (m *OrgContactRequestBuilder) Restore()(*i92107bebedaa23b60d8098d09f17a5dc81b4ca28fa13623b6ba4df811ca907ea.RestoreRequestBuilder) {
    return i92107bebedaa23b60d8098d09f17a5dc81b4ca28fa13623b6ba4df811ca907ea.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
