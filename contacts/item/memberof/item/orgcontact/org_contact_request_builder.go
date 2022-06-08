package orgcontact

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i546e5bf579c17dc5b1a8a8b5fba378ff01db4e6e9d9ad9884d5262d83ce0b4a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item/orgcontact/getmembergroups"
    i66686c67de546c2729e54a3d33f04e4364390fa107d2bc14e57fe9af395546b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item/orgcontact/getmemberobjects"
    i6d1ae0653c95a30cf39193ab9e9bbb6f543a147fb58491560035f4f7015bd41d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item/orgcontact/checkmembergroups"
    i93040ef40b8d951b56efa2e33fc779c75164f7ec745ce688bc6a70ba0944eaeb "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item/orgcontact/checkmemberobjects"
    id1de146679d835526d1cc33f5b6eafdc47409a1661b6149bb811441da1aa0321 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item/orgcontact/restore"
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
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*i6d1ae0653c95a30cf39193ab9e9bbb6f543a147fb58491560035f4f7015bd41d.CheckMemberGroupsRequestBuilder) {
    return i6d1ae0653c95a30cf39193ab9e9bbb6f543a147fb58491560035f4f7015bd41d.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*i93040ef40b8d951b56efa2e33fc779c75164f7ec745ce688bc6a70ba0944eaeb.CheckMemberObjectsRequestBuilder) {
    return i93040ef40b8d951b56efa2e33fc779c75164f7ec745ce688bc6a70ba0944eaeb.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.orgContact{?%24select,%24expand}";
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
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.orgContact
func (m *OrgContactRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i546e5bf579c17dc5b1a8a8b5fba378ff01db4e6e9d9ad9884d5262d83ce0b4a9.GetMemberGroupsRequestBuilder) {
    return i546e5bf579c17dc5b1a8a8b5fba378ff01db4e6e9d9ad9884d5262d83ce0b4a9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactRequestBuilder) GetMemberObjects()(*i66686c67de546c2729e54a3d33f04e4364390fa107d2bc14e57fe9af395546b1.GetMemberObjectsRequestBuilder) {
    return i66686c67de546c2729e54a3d33f04e4364390fa107d2bc14e57fe9af395546b1.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *OrgContactRequestBuilder) Restore()(*id1de146679d835526d1cc33f5b6eafdc47409a1661b6149bb811441da1aa0321.RestoreRequestBuilder) {
    return id1de146679d835526d1cc33f5b6eafdc47409a1661b6149bb811441da1aa0321.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
