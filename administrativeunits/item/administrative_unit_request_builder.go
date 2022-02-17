package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/scopedrolemembers"
    i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/members"
    i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/checkmembergroups"
    i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/restore"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/getmembergroups"
    ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/checkmemberobjects"
    ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/extensions"
    if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/getmemberobjects"
    ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/scopedrolemembers/item"
    ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits/item/extensions/item"
)

// AdministrativeUnitRequestBuilder builds and executes requests for operations under \administrativeUnits\{administrativeUnit-id}
type AdministrativeUnitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdministrativeUnitRequestBuilderDeleteOptions options for Delete
type AdministrativeUnitRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdministrativeUnitRequestBuilderGetOptions options for Get
type AdministrativeUnitRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdministrativeUnitRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdministrativeUnitRequestBuilderGetQueryParameters get entity from administrativeUnits by key
type AdministrativeUnitRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AdministrativeUnitRequestBuilderPatchOptions options for Patch
type AdministrativeUnitRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AdministrativeUnitRequestBuilder) CheckMemberGroups()(*i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956.CheckMemberGroupsRequestBuilder) {
    return i349829ba4fce29dd7987c6df11dbcaf800247af0fa3219b01dcdc6e1343f1956.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) CheckMemberObjects()(*ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c.CheckMemberObjectsRequestBuilder) {
    return ib099fb63054d0dd1784a47669d75c804d03b727cd11a2480311fa2d37e24ed6c.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAdministrativeUnitRequestBuilderInternal instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
func NewAdministrativeUnitRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    m := &AdministrativeUnitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/administrativeUnits/{administrativeUnit_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdministrativeUnitRequestBuilder instantiates a new AdministrativeUnitRequestBuilder and sets the default values.
func NewAdministrativeUnitRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdministrativeUnitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeUnitRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from administrativeUnits
func (m *AdministrativeUnitRequestBuilder) CreateDeleteRequestInformation(options *AdministrativeUnitRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from administrativeUnits by key
func (m *AdministrativeUnitRequestBuilder) CreateGetRequestInformation(options *AdministrativeUnitRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in administrativeUnits
func (m *AdministrativeUnitRequestBuilder) CreatePatchRequestInformation(options *AdministrativeUnitRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from administrativeUnits
func (m *AdministrativeUnitRequestBuilder) Delete(options *AdministrativeUnitRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AdministrativeUnitRequestBuilder) Extensions()(*ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c.ExtensionsRequestBuilder) {
    return ic10bcb37ee210a709c1b4b767679d61e4e281098278ef4698ae36b4863f0598c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.extensions.item collection
func (m *AdministrativeUnitRequestBuilder) ExtensionsById(id string)(*ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ieb753c78202e5499a95653d7cfff3918b3c7124418d5e65033612f2c7fbc09c2.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from administrativeUnits by key
func (m *AdministrativeUnitRequestBuilder) Get(options *AdministrativeUnitRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAdministrativeUnit() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministrativeUnit), nil
}
func (m *AdministrativeUnitRequestBuilder) GetMemberGroups()(*ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73.GetMemberGroupsRequestBuilder) {
    return ia216e040c36eeacd190ee60a98d2779e9f151c69020cf34b69a5a127c5a6db73.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) GetMemberObjects()(*if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74.GetMemberObjectsRequestBuilder) {
    return if3ef57b69f282bfad9e1e97bd91a5b30f4da893d66532d98d784f6d054b6fb74.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) Members()(*i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82.MembersRequestBuilder) {
    return i1fa710cc359019f3bf0614158e392cfc3471c6427ce07ada8e4627100cec5a82.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in administrativeUnits
func (m *AdministrativeUnitRequestBuilder) Patch(options *AdministrativeUnitRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AdministrativeUnitRequestBuilder) Restore()(*i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d.RestoreRequestBuilder) {
    return i4e5a3bd6a5396f5b0939bfce85a44230296cab8d3446ee6e3a5a29204173428d.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembers()(*i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c.ScopedRoleMembersRequestBuilder) {
    return i0f36168ca30e64e3bea01bb77d0f73154f5020719384d2f3ab434667a1277d0c.NewScopedRoleMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.administrativeUnits.item.scopedRoleMembers.item collection
func (m *AdministrativeUnitRequestBuilder) ScopedRoleMembersById(id string)(*ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780.ScopedRoleMembershipRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership_id"] = id
    }
    return ia0eae455c54ecade94b5dc033e0f792d508d4ae4d0e22ed0d19c83e392692780.NewScopedRoleMembershipRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
