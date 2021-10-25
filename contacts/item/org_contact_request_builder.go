package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports"
    i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmembergroups"
    i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmembergroups"
    iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmemberobjects"
    ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/manager"
    id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports"
    ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmemberobjects"
    iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/restore"
    if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof"
)

type OrgContactRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type OrgContactRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *OrgContactRequestBuilder) CheckMemberGroups()(*i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.CheckMemberGroupsRequestBuilder) {
    return i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) CheckMemberObjects()(*iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.CheckMemberObjectsRequestBuilder) {
    return iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactRequestBuilder) {
    m := &OrgContactRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/contacts/{orgContact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewOrgContactRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OrgContactRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrgContactRequestBuilder) CreateGetRequestInformation(q func (value *OrgContactRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OrgContactRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrgContactRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContact, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *OrgContactRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OrgContactRequestBuilder) DirectReports()(*id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a.DirectReportsRequestBuilder) {
    return id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Get(q func (value *OrgContactRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContact, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOrgContact() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContact), nil
}
func (m *OrgContactRequestBuilder) GetMemberGroups()(*i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.GetMemberGroupsRequestBuilder) {
    return i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) GetMemberObjects()(*ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.GetMemberObjectsRequestBuilder) {
    return ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Manager()(*ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.ManagerRequestBuilder) {
    return ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) MemberOf()(*i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5.MemberOfRequestBuilder) {
    return i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContact, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *OrgContactRequestBuilder) Restore()(*iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.RestoreRequestBuilder) {
    return iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) TransitiveMemberOf()(*if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d.TransitiveMemberOfRequestBuilder) {
    return if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactRequestBuilder) TransitiveReports()(*i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724.TransitiveReportsRequestBuilder) {
    return i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724.NewTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
