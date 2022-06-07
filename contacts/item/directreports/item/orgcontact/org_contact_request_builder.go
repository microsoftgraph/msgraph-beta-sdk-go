package orgcontact

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d2c8f2417a88f8b6152e58ef065046918d0c5643a1fc8fb21a197d8103f82c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item/orgcontact/checkmembergroups"
    i4e84ba3476957065a451cb2d9384c4aee179d3b6cb0cec49dd736b34b8de1dd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item/orgcontact/getmemberobjects"
    i5475e8370e5c5e49ecee6be74dcfbff5f85a2adf11284ac89a5ee1bdd9ed6f5e "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item/orgcontact/getmembergroups"
    i7fa09bb322b69b9641e904b7e4946a8a0b6d7db60723691d59db1acbb8cb72e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item/orgcontact/restore"
    ie61da601a2a0b0b58f8f5dbf27a73352f4e988204388a0a5ce1a6282fa85a422 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item/orgcontact/checkmemberobjects"
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
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*i4d2c8f2417a88f8b6152e58ef065046918d0c5643a1fc8fb21a197d8103f82c8.CheckMemberGroupsRequestBuilder) {
    return i4d2c8f2417a88f8b6152e58ef065046918d0c5643a1fc8fb21a197d8103f82c8.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*ie61da601a2a0b0b58f8f5dbf27a73352f4e988204388a0a5ce1a6282fa85a422.CheckMemberObjectsRequestBuilder) {
    return ie61da601a2a0b0b58f8f5dbf27a73352f4e988204388a0a5ce1a6282fa85a422.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/directReports/{directoryObject%2Did}/microsoft.graph.orgContact{?%24select,%24expand}";
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
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i5475e8370e5c5e49ecee6be74dcfbff5f85a2adf11284ac89a5ee1bdd9ed6f5e.GetMemberGroupsRequestBuilder) {
    return i5475e8370e5c5e49ecee6be74dcfbff5f85a2adf11284ac89a5ee1bdd9ed6f5e.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactRequestBuilder) GetMemberObjects()(*i4e84ba3476957065a451cb2d9384c4aee179d3b6cb0cec49dd736b34b8de1dd0.GetMemberObjectsRequestBuilder) {
    return i4e84ba3476957065a451cb2d9384c4aee179d3b6cb0cec49dd736b34b8de1dd0.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *OrgContactRequestBuilder) Restore()(*i7fa09bb322b69b9641e904b7e4946a8a0b6d7db60723691d59db1acbb8cb72e3.RestoreRequestBuilder) {
    return i7fa09bb322b69b9641e904b7e4946a8a0b6d7db60723691d59db1acbb8cb72e3.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
