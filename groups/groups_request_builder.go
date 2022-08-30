package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/validateproperties"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/evaluatedynamicmembership"
    i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/getuserownedobjects"
    ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/getbyids"
    ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/delta"
)

// GroupsRequestBuilder provides operations to manage the collection of group entities.
type GroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsRequestBuilderGetQueryParameters list all the groups available in an organization, excluding dynamic distribution groups. To retrieve dynamic distribution groups, use the Exchange admin center. This operation returns by default only a subset of the more commonly used properties for each group. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the group and specify the properties in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
type GroupsRequestBuilderGetQueryParameters struct {
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsRequestBuilderGetQueryParameters
}
// GroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsRequestBuilderInternal instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsRequestBuilder) {
    m := &GroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups{?%24top,%24search,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsRequestBuilder instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation list all the groups available in an organization, excluding dynamic distribution groups. To retrieve dynamic distribution groups, use the Exchange admin center. This operation returns by default only a subset of the more commonly used properties for each group. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the group and specify the properties in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration list all the groups available in an organization, excluding dynamic distribution groups. To retrieve dynamic distribution groups, use the Exchange admin center. This operation returns by default only a subset of the more commonly used properties for each group. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the group and specify the properties in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create a new group as specified in the request body. You can create one of the following groups: This operation returns by default only a subset of the properties for each group. These default properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation and specify the properties in a `$select` OData query option. **Note**: To create a team, first create a group then add a team to it, see create team.
func (m *GroupsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create a new group as specified in the request body. You can create one of the following groups: This operation returns by default only a subset of the properties for each group. These default properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation and specify the properties in a `$select` OData query option. **Note**: To create a team, first create a group then add a team to it, see create team.
func (m *GroupsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *GroupsRequestBuilder) Delta()(*ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738.DeltaRequestBuilder) {
    return ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDynamicMembership the evaluateDynamicMembership property
func (m *GroupsRequestBuilder) EvaluateDynamicMembership()(*i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573.EvaluateDynamicMembershipRequestBuilder) {
    return i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get list all the groups available in an organization, excluding dynamic distribution groups. To retrieve dynamic distribution groups, use the Exchange admin center. This operation returns by default only a subset of the more commonly used properties for each group. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the group and specify the properties in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetByIds the getByIds property
func (m *GroupsRequestBuilder) GetByIds()(*ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7.GetByIdsRequestBuilder) {
    return ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *GroupsRequestBuilder) GetUserOwnedObjects()(*i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827.GetUserOwnedObjectsRequestBuilder) {
    return i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler list all the groups available in an organization, excluding dynamic distribution groups. To retrieve dynamic distribution groups, use the Exchange admin center. This operation returns by default only a subset of the more commonly used properties for each group. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the group and specify the properties in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query.
func (m *GroupsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GroupsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Post create a new group as specified in the request body. You can create one of the following groups: This operation returns by default only a subset of the properties for each group. These default properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation and specify the properties in a `$select` OData query option. **Note**: To create a team, first create a group then add a team to it, see create team.
func (m *GroupsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create a new group as specified in the request body. You can create one of the following groups: This operation returns by default only a subset of the properties for each group. These default properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation and specify the properties in a `$select` OData query option. **Note**: To create a team, first create a group then add a team to it, see create team.
func (m *GroupsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ValidateProperties the validateProperties property
func (m *GroupsRequestBuilder) ValidateProperties()(*i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027.ValidatePropertiesRequestBuilder) {
    return i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
