package msgraphbetasdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials"
    i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465 "github.com/microsoftgraph/msgraph-beta-sdk-go/functions"
    i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84 "github.com/microsoftgraph/msgraph-beta-sdk-go/mobilitymanagementpolicies"
    i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement"
    i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices"
    i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/filteroperators"
    i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles"
    i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d "github.com/microsoftgraph/msgraph-beta-sdk-go/datapolicyoperations"
    i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignments"
    i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreementacceptances"
    i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontrols"
    i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/messageevents"
    i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a "github.com/microsoftgraph/msgraph-beta-sdk-go/schemaextensions"
    i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition"
    i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews"
    i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles"
    i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a "github.com/microsoftgraph/msgraph-beta-sdk-go/identity"
    i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/appcatalogs"
    i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/payloadresponse"
    i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5 "github.com/microsoftgraph/msgraph-beta-sdk-go/threatsubmission"
    i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e "github.com/microsoftgraph/msgraph-beta-sdk-go/riskdetections"
    i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites"
    i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad "github.com/microsoftgraph/msgraph-beta-sdk-go/connections"
    i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives"
    i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats"
    i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f "github.com/microsoftgraph/msgraph-beta-sdk-go/planner"
    i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess"
    i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement"
    i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f "github.com/microsoftgraph/msgraph-beta-sdk-go/trustframework"
    i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroles"
    i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29 "github.com/microsoftgraph/msgraph-beta-sdk-go/approleassignments"
    i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc "github.com/microsoftgraph/msgraph-beta-sdk-go/authenticationmethodspolicy"
    i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceresources"
    i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b "github.com/microsoftgraph/msgraph-beta-sdk-go/monitoring"
    i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests"
    i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviewdecisions"
    i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef "github.com/microsoftgraph/msgraph-beta-sdk-go/messagetraces"
    i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection"
    i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb "github.com/microsoftgraph/msgraph-beta-sdk-go/external"
    i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc "github.com/microsoftgraph/msgraph-beta-sdk-go/permissiongrants"
    i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements"
    i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/employeeexperience"
    i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/invitations"
    i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityproviders"
    i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053 "github.com/microsoftgraph/msgraph-beta-sdk-go/messagerecipients"
    i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e "github.com/microsoftgraph/msgraph-beta-sdk-go/search"
    i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me"
    i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates"
    i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53 "github.com/microsoftgraph/msgraph-beta-sdk-go/domaindnsrecords"
    i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54 "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts"
    i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies"
    i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscriptions"
    i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval"
    i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin"
    i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamstemplates"
    i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus"
    i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/auditlogs"
    i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory"
    i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b "github.com/microsoftgraph/msgraph-beta-sdk-go/scopedrolememberships"
    i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17 "github.com/microsoftgraph/msgraph-beta-sdk-go/grouplifecyclepolicies"
    i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf "github.com/microsoftgraph/msgraph-beta-sdk-go/settings"
    i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts"
    i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions"
    i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares"
    i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print"
    i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184 "github.com/microsoftgraph/msgraph-beta-sdk-go/security"
    i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups"
    i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy"
    i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance"
    i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba "github.com/microsoftgraph/msgraph-beta-sdk-go/commands"
    i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedoperationevents"
    i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349 "github.com/microsoftgraph/msgraph-beta-sdk-go/alloweddatalocations"
    i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f "github.com/microsoftgraph/msgraph-beta-sdk-go/administrativeunits"
    i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroledefinitions"
    i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c "github.com/microsoftgraph/msgraph-beta-sdk-go/activitystatistics"
    i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f "github.com/microsoftgraph/msgraph-beta-sdk-go/authenticationmethodconfigurations"
    i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/businessflowtemplates"
    i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore"
    i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains"
    i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba "github.com/microsoftgraph/msgraph-beta-sdk-go/applicationtemplates"
    ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports"
    ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance"
    iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates"
    iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/riskyusers"
    iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects"
    ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd "github.com/microsoftgraph/msgraph-beta-sdk-go/governancerolesettings"
    ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e "github.com/microsoftgraph/msgraph-beta-sdk-go/programs"
    ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration"
    ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals"
    ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses"
    ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/places"
    icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork"
    icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95 "github.com/microsoftgraph/msgraph-beta-sdk-go/branding"
    icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
    id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821 "github.com/microsoftgraph/msgraph-beta-sdk-go/storage"
    id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/certificatebasedauthconfiguration"
    id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships"
    id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education"
    ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection"
    idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive"
    idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement"
    ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications"
    ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad "github.com/microsoftgraph/msgraph-beta-sdk-go/oauth2permissiongrants"
    ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingcurrencies"
    ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization"
    ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscribedskus"
    if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests"
    if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification"
    if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders"
    if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontroltypes"
    if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications"
    ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219 "github.com/microsoftgraph/msgraph-beta-sdk-go/app"
    ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/governancesubjects"
    iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewDecisions provides operations to manage the collection of accessReviewDecision entities.
func (m *GraphBaseServiceClient) AccessReviewDecisions()(*i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.AccessReviewDecisionsRequestBuilder) {
    return i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.NewAccessReviewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessReviewDecisionsById provides operations to manage the collection of accessReviewDecision entities.
func (m *GraphBaseServiceClient) AccessReviewDecisionsById(id string)(*i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.AccessReviewDecisionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewDecision%2Did"] = id
    }
    return i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.NewAccessReviewDecisionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessReviews provides operations to manage the collection of accessReview entities.
func (m *GraphBaseServiceClient) AccessReviews()(*i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.AccessReviewsRequestBuilder) {
    return i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.NewAccessReviewsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessReviewsById provides operations to manage the collection of accessReview entities.
func (m *GraphBaseServiceClient) AccessReviewsById(id string)(*i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.AccessReviewItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReview%2Did"] = id
    }
    return i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.NewAccessReviewItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Activitystatistics provides operations to manage the collection of activityStatistics entities.
func (m *GraphBaseServiceClient) Activitystatistics()(*i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.ActivitystatisticsRequestBuilder) {
    return i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.NewActivitystatisticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitystatisticsById provides operations to manage the collection of activityStatistics entities.
func (m *GraphBaseServiceClient) ActivitystatisticsById(id string)(*i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.ActivityStatisticsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["activityStatistics%2Did"] = id
    }
    return i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.NewActivityStatisticsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Admin provides operations to manage the admin singleton.
func (m *GraphBaseServiceClient) Admin()(*i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.AdminRequestBuilder) {
    return i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.NewAdminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdministrativeUnits provides operations to manage the collection of administrativeUnit entities.
func (m *GraphBaseServiceClient) AdministrativeUnits()(*i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.AdministrativeUnitsRequestBuilder) {
    return i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.NewAdministrativeUnitsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdministrativeUnitsById provides operations to manage the collection of administrativeUnit entities.
func (m *GraphBaseServiceClient) AdministrativeUnitsById(id string)(*i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.AdministrativeUnitItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["administrativeUnit%2Did"] = id
    }
    return i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.NewAdministrativeUnitItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptances()(*i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.AgreementAcceptancesRequestBuilder) {
    return i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById provides operations to manage the collection of agreementAcceptance entities.
func (m *GraphBaseServiceClient) AgreementAcceptancesById(id string)(*i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Agreements provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) Agreements()(*i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.AgreementsRequestBuilder) {
    return i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementsById provides operations to manage the collection of agreement entities.
func (m *GraphBaseServiceClient) AgreementsById(id string)(*i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.AgreementItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreement%2Did"] = id
    }
    return i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.NewAgreementItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AllowedDataLocations provides operations to manage the collection of allowedDataLocation entities.
func (m *GraphBaseServiceClient) AllowedDataLocations()(*i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.AllowedDataLocationsRequestBuilder) {
    return i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.NewAllowedDataLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllowedDataLocationsById provides operations to manage the collection of allowedDataLocation entities.
func (m *GraphBaseServiceClient) AllowedDataLocationsById(id string)(*i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.AllowedDataLocationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["allowedDataLocation%2Did"] = id
    }
    return i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.NewAllowedDataLocationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// App provides operations to manage the commsApplication singleton.
func (m *GraphBaseServiceClient) App()(*ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.AppRequestBuilder) {
    return ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.NewAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppCatalogs provides operations to manage the appCatalogs singleton.
func (m *GraphBaseServiceClient) AppCatalogs()(*i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.AppCatalogsRequestBuilder) {
    return i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.NewAppCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Applications provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) Applications()(*ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.ApplicationsRequestBuilder) {
    return ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.NewApplicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationsById provides operations to manage the collection of application entities.
func (m *GraphBaseServiceClient) ApplicationsById(id string)(*ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.ApplicationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["application%2Did"] = id
    }
    return ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.NewApplicationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ApplicationTemplates provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplates()(*i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.ApplicationTemplatesRequestBuilder) {
    return i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.NewApplicationTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationTemplatesById provides operations to manage the collection of applicationTemplate entities.
func (m *GraphBaseServiceClient) ApplicationTemplatesById(id string)(*i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.ApplicationTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationTemplate%2Did"] = id
    }
    return i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.NewApplicationTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the collection of appRoleAssignment entities.
func (m *GraphBaseServiceClient) AppRoleAssignments()(*i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.AppRoleAssignmentsRequestBuilder) {
    return i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the collection of appRoleAssignment entities.
func (m *GraphBaseServiceClient) AppRoleAssignmentsById(id string)(*i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ApprovalWorkflowProviders provides operations to manage the collection of approvalWorkflowProvider entities.
func (m *GraphBaseServiceClient) ApprovalWorkflowProviders()(*if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.ApprovalWorkflowProvidersRequestBuilder) {
    return if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.NewApprovalWorkflowProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalWorkflowProvidersById provides operations to manage the collection of approvalWorkflowProvider entities.
func (m *GraphBaseServiceClient) ApprovalWorkflowProvidersById(id string)(*if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.ApprovalWorkflowProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approvalWorkflowProvider%2Did"] = id
    }
    return if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.NewApprovalWorkflowProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditLogs provides operations to manage the auditLogRoot singleton.
func (m *GraphBaseServiceClient) AuditLogs()(*i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.AuditLogsRequestBuilder) {
    return i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.NewAuditLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodConfigurations provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurations()(*i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.AuthenticationMethodConfigurationsRequestBuilder) {
    return i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.NewAuthenticationMethodConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodConfigurationsById provides operations to manage the collection of authenticationMethodConfiguration entities.
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurationsById(id string)(*i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.AuthenticationMethodConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethodConfiguration%2Did"] = id
    }
    return i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.NewAuthenticationMethodConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy singleton.
func (m *GraphBaseServiceClient) AuthenticationMethodsPolicy()(*i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.AuthenticationMethodsPolicyRequestBuilder) {
    return i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookingBusinesses provides operations to manage the collection of bookingBusiness entities.
func (m *GraphBaseServiceClient) BookingBusinesses()(*ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.BookingBusinessesRequestBuilder) {
    return ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.NewBookingBusinessesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookingBusinessesById provides operations to manage the collection of bookingBusiness entities.
func (m *GraphBaseServiceClient) BookingBusinessesById(id string)(*ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.BookingBusinessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingBusiness%2Did"] = id
    }
    return ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.NewBookingBusinessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// BookingCurrencies provides operations to manage the collection of bookingCurrency entities.
func (m *GraphBaseServiceClient) BookingCurrencies()(*ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.BookingCurrenciesRequestBuilder) {
    return ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.NewBookingCurrenciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BookingCurrenciesById provides operations to manage the collection of bookingCurrency entities.
func (m *GraphBaseServiceClient) BookingCurrenciesById(id string)(*ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.BookingCurrencyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["bookingCurrency%2Did"] = id
    }
    return ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.NewBookingCurrencyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Branding provides operations to manage the organizationalBranding singleton.
func (m *GraphBaseServiceClient) Branding()(*icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95.BrandingRequestBuilder) {
    return icbfa8075dc5e1ad04a3bc48d231d9b422f250ca1e2b74477a880e0db7b0e7f95.NewBrandingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowTemplates provides operations to manage the collection of businessFlowTemplate entities.
func (m *GraphBaseServiceClient) BusinessFlowTemplates()(*i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.BusinessFlowTemplatesRequestBuilder) {
    return i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.NewBusinessFlowTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BusinessFlowTemplatesById provides operations to manage the collection of businessFlowTemplate entities.
func (m *GraphBaseServiceClient) BusinessFlowTemplatesById(id string)(*i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.BusinessFlowTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessFlowTemplate%2Did"] = id
    }
    return i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.NewBusinessFlowTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CertificateBasedAuthConfiguration provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfiguration()(*id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.CertificateBasedAuthConfigurationRequestBuilder) {
    return id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateBasedAuthConfigurationById provides operations to manage the collection of certificateBasedAuthConfiguration entities.
func (m *GraphBaseServiceClient) CertificateBasedAuthConfigurationById(id string)(*id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.CertificateBasedAuthConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateBasedAuthConfiguration%2Did"] = id
    }
    return id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.NewCertificateBasedAuthConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Chats provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) Chats()(*i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.ChatsRequestBuilder) {
    return i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById provides operations to manage the collection of chat entities.
func (m *GraphBaseServiceClient) ChatsById(id string)(*i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.ChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.NewChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Commands provides operations to manage the collection of command entities.
func (m *GraphBaseServiceClient) Commands()(*i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.CommandsRequestBuilder) {
    return i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById provides operations to manage the collection of command entities.
func (m *GraphBaseServiceClient) CommandsById(id string)(*i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Communications provides operations to manage the cloudCommunications singleton.
func (m *GraphBaseServiceClient) Communications()(*if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.CommunicationsRequestBuilder) {
    return if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.NewCommunicationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Compliance provides operations to manage the compliance singleton.
func (m *GraphBaseServiceClient) Compliance()(*ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.ComplianceRequestBuilder) {
    return ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.NewComplianceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Connections provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) Connections()(*i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.ConnectionsRequestBuilder) {
    return i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.NewConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectionsById provides operations to manage the collection of externalConnection entities.
func (m *GraphBaseServiceClient) ConnectionsById(id string)(*i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.ExternalConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalConnection%2Did"] = id
    }
    return i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.NewExternalConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGraphBaseServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewGraphBaseServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphBaseServiceClient) {
    m := &GraphBaseServiceClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
    }
    m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
    return m
}
// Contacts provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) Contacts()(*i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.ContactsRequestBuilder) {
    return i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById provides operations to manage the collection of orgContact entities.
func (m *GraphBaseServiceClient) ContactsById(id string)(*i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.OrgContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["orgContact%2Did"] = id
    }
    return i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.NewOrgContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contracts provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) Contracts()(*i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.ContractsRequestBuilder) {
    return i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.NewContractsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContractsById provides operations to manage the collection of contract entities.
func (m *GraphBaseServiceClient) ContractsById(id string)(*i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.ContractItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contract%2Did"] = id
    }
    return i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.NewContractItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DataClassification provides operations to manage the dataClassificationService singleton.
func (m *GraphBaseServiceClient) DataClassification()(*if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.DataClassificationRequestBuilder) {
    return if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.NewDataClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataPolicyOperations provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperations()(*i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.DataPolicyOperationsRequestBuilder) {
    return i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.NewDataPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataPolicyOperationsById provides operations to manage the collection of dataPolicyOperation entities.
func (m *GraphBaseServiceClient) DataPolicyOperationsById(id string)(*i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.DataPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataPolicyOperation%2Did"] = id
    }
    return i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.NewDataPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceAppManagement provides operations to manage the deviceAppManagement singleton.
func (m *GraphBaseServiceClient) DeviceAppManagement()(*idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.DeviceAppManagementRequestBuilder) {
    return idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.NewDeviceAppManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagement provides operations to manage the deviceManagement singleton.
func (m *GraphBaseServiceClient) DeviceManagement()(*i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.DeviceManagementRequestBuilder) {
    return i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Devices provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) Devices()(*i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.DevicesRequestBuilder) {
    return i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById provides operations to manage the collection of device entities.
func (m *GraphBaseServiceClient) DevicesById(id string)(*i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.DeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.NewDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Directory provides operations to manage the directory singleton.
func (m *GraphBaseServiceClient) Directory()(*i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.DirectoryRequestBuilder) {
    return i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryObjects provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjects()(*iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.DirectoryObjectsRequestBuilder) {
    return iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.NewDirectoryObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryObjectsById provides operations to manage the collection of directoryObject entities.
func (m *GraphBaseServiceClient) DirectoryObjectsById(id string)(*iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectoryRoles provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRoles()(*i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.DirectoryRolesRequestBuilder) {
    return i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.NewDirectoryRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRolesById provides operations to manage the collection of directoryRole entities.
func (m *GraphBaseServiceClient) DirectoryRolesById(id string)(*i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.DirectoryRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRole%2Did"] = id
    }
    return i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.NewDirectoryRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectoryRoleTemplates provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplates()(*i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.DirectoryRoleTemplatesRequestBuilder) {
    return i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.NewDirectoryRoleTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleTemplatesById provides operations to manage the collection of directoryRoleTemplate entities.
func (m *GraphBaseServiceClient) DirectoryRoleTemplatesById(id string)(*i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.DirectoryRoleTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryRoleTemplate%2Did"] = id
    }
    return i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.NewDirectoryRoleTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectorySettingTemplates provides operations to manage the collection of directorySettingTemplate entities.
func (m *GraphBaseServiceClient) DirectorySettingTemplates()(*iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.DirectorySettingTemplatesRequestBuilder) {
    return iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.NewDirectorySettingTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectorySettingTemplatesById provides operations to manage the collection of directorySettingTemplate entities.
func (m *GraphBaseServiceClient) DirectorySettingTemplatesById(id string)(*iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.DirectorySettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySettingTemplate%2Did"] = id
    }
    return iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.NewDirectorySettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DomainDnsRecords provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecords()(*i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.DomainDnsRecordsRequestBuilder) {
    return i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.NewDomainDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainDnsRecordsById provides operations to manage the collection of domainDnsRecord entities.
func (m *GraphBaseServiceClient) DomainDnsRecordsById(id string)(*i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Domains provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) Domains()(*i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.DomainsRequestBuilder) {
    return i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.NewDomainsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainsById provides operations to manage the collection of domain entities.
func (m *GraphBaseServiceClient) DomainsById(id string)(*i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.DomainItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domain%2Did"] = id
    }
    return i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.NewDomainItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive provides operations to manage the drive singleton.
func (m *GraphBaseServiceClient) Drive()(*idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a.DriveRequestBuilder) {
    return idc403d989fa3b67787cd15ceedea7e45400694aa996d133f9bf61a54fe5a497a.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) Drives()(*i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.DrivesRequestBuilder) {
    return i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the collection of drive entities.
func (m *GraphBaseServiceClient) DrivesById(id string)(*i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Education provides operations to manage the educationRoot singleton.
func (m *GraphBaseServiceClient) Education()(*id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.EducationRequestBuilder) {
    return id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.NewEducationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmployeeExperience provides operations to manage the employeeExperience singleton.
func (m *GraphBaseServiceClient) EmployeeExperience()(*i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.EmployeeExperienceRequestBuilder) {
    return i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.NewEmployeeExperienceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// External provides operations to manage the external singleton.
func (m *GraphBaseServiceClient) External()(*i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.ExternalRequestBuilder) {
    return i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.NewExternalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilterOperators provides operations to manage the collection of filterOperatorSchema entities.
func (m *GraphBaseServiceClient) FilterOperators()(*i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.FilterOperatorsRequestBuilder) {
    return i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.NewFilterOperatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilterOperatorsById provides operations to manage the collection of filterOperatorSchema entities.
func (m *GraphBaseServiceClient) FilterOperatorsById(id string)(*i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.FilterOperatorSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["filterOperatorSchema%2Did"] = id
    }
    return i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.NewFilterOperatorSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Financials provides operations to manage the financials singleton.
func (m *GraphBaseServiceClient) Financials()(*i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.FinancialsRequestBuilder) {
    return i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.NewFinancialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Functions provides operations to manage the collection of attributeMappingFunctionSchema entities.
func (m *GraphBaseServiceClient) Functions()(*i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.FunctionsRequestBuilder) {
    return i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.NewFunctionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FunctionsById provides operations to manage the collection of attributeMappingFunctionSchema entities.
func (m *GraphBaseServiceClient) FunctionsById(id string)(*i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.AttributeMappingFunctionSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attributeMappingFunctionSchema%2Did"] = id
    }
    return i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.NewAttributeMappingFunctionSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceResources provides operations to manage the collection of governanceResource entities.
func (m *GraphBaseServiceClient) GovernanceResources()(*i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.GovernanceResourcesRequestBuilder) {
    return i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.NewGovernanceResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceResourcesById provides operations to manage the collection of governanceResource entities.
func (m *GraphBaseServiceClient) GovernanceResourcesById(id string)(*i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.GovernanceResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceResource%2Did"] = id
    }
    return i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.NewGovernanceResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceRoleAssignmentRequests provides operations to manage the collection of governanceRoleAssignmentRequest entities.
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentRequests()(*i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.GovernanceRoleAssignmentRequestsRequestBuilder) {
    return i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.NewGovernanceRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceRoleAssignmentRequestsById provides operations to manage the collection of governanceRoleAssignmentRequest entities.
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentRequestsById(id string)(*i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.GovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest%2Did"] = id
    }
    return i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceRoleAssignments provides operations to manage the collection of governanceRoleAssignment entities.
func (m *GraphBaseServiceClient) GovernanceRoleAssignments()(*i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.GovernanceRoleAssignmentsRequestBuilder) {
    return i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.NewGovernanceRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceRoleAssignmentsById provides operations to manage the collection of governanceRoleAssignment entities.
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentsById(id string)(*i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.GovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment%2Did"] = id
    }
    return i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceRoleDefinitions provides operations to manage the collection of governanceRoleDefinition entities.
func (m *GraphBaseServiceClient) GovernanceRoleDefinitions()(*i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.GovernanceRoleDefinitionsRequestBuilder) {
    return i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.NewGovernanceRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceRoleDefinitionsById provides operations to manage the collection of governanceRoleDefinition entities.
func (m *GraphBaseServiceClient) GovernanceRoleDefinitionsById(id string)(*i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.GovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition%2Did"] = id
    }
    return i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.NewGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceRoleSettings provides operations to manage the collection of governanceRoleSetting entities.
func (m *GraphBaseServiceClient) GovernanceRoleSettings()(*ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.GovernanceRoleSettingsRequestBuilder) {
    return ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.NewGovernanceRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceRoleSettingsById provides operations to manage the collection of governanceRoleSetting entities.
func (m *GraphBaseServiceClient) GovernanceRoleSettingsById(id string)(*ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.GovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting%2Did"] = id
    }
    return ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.NewGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GovernanceSubjects provides operations to manage the collection of governanceSubject entities.
func (m *GraphBaseServiceClient) GovernanceSubjects()(*ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.GovernanceSubjectsRequestBuilder) {
    return ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.NewGovernanceSubjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GovernanceSubjectsById provides operations to manage the collection of governanceSubject entities.
func (m *GraphBaseServiceClient) GovernanceSubjectsById(id string)(*ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.GovernanceSubjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceSubject%2Did"] = id
    }
    return ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.NewGovernanceSubjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupLifecyclePolicies provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePolicies()(*i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.GroupLifecyclePoliciesRequestBuilder) {
    return i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.NewGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById provides operations to manage the collection of groupLifecyclePolicy entities.
func (m *GraphBaseServiceClient) GroupLifecyclePoliciesById(id string)(*i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.GroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.NewGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Groups provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) Groups()(*i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupsRequestBuilder) {
    return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) GroupsById(id string)(*i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Identity provides operations to manage the identityContainer singleton.
func (m *GraphBaseServiceClient) Identity()(*i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.IdentityRequestBuilder) {
    return i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.NewIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityGovernance provides operations to manage the identityGovernance singleton.
func (m *GraphBaseServiceClient) IdentityGovernance()(*i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.IdentityGovernanceRequestBuilder) {
    return i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.NewIdentityGovernanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProtection provides operations to manage the identityProtectionRoot singleton.
func (m *GraphBaseServiceClient) IdentityProtection()(*i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.IdentityProtectionRequestBuilder) {
    return i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.NewIdentityProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProviders provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProviders()(*i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.IdentityProvidersRequestBuilder) {
    return i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById provides operations to manage the collection of identityProvider entities.
func (m *GraphBaseServiceClient) IdentityProvidersById(id string)(*i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InformationProtection provides operations to manage the informationProtection singleton.
func (m *GraphBaseServiceClient) InformationProtection()(*ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.InformationProtectionRequestBuilder) {
    return ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Invitations provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) Invitations()(*i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.InvitationsRequestBuilder) {
    return i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.NewInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvitationsById provides operations to manage the collection of invitation entities.
func (m *GraphBaseServiceClient) InvitationsById(id string)(*i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.InvitationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["invitation%2Did"] = id
    }
    return i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.NewInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Me provides operations to manage the user singleton.
func (m *GraphBaseServiceClient) Me()(*i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb.MeRequestBuilder) {
    return i5a50ec7cbb7e52feaacf8d45f5e247e765502c09622e2fa44cd68445eab876eb.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageEvents provides operations to manage the collection of messageEvent entities.
func (m *GraphBaseServiceClient) MessageEvents()(*i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.MessageEventsRequestBuilder) {
    return i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.NewMessageEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageEventsById provides operations to manage the collection of messageEvent entities.
func (m *GraphBaseServiceClient) MessageEventsById(id string)(*i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.MessageEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageEvent%2Did"] = id
    }
    return i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.NewMessageEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MessageRecipients provides operations to manage the collection of messageRecipient entities.
func (m *GraphBaseServiceClient) MessageRecipients()(*i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.MessageRecipientsRequestBuilder) {
    return i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.NewMessageRecipientsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRecipientsById provides operations to manage the collection of messageRecipient entities.
func (m *GraphBaseServiceClient) MessageRecipientsById(id string)(*i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.MessageRecipientItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRecipient%2Did"] = id
    }
    return i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.NewMessageRecipientItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MessageTraces provides operations to manage the collection of messageTrace entities.
func (m *GraphBaseServiceClient) MessageTraces()(*i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.MessageTracesRequestBuilder) {
    return i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.NewMessageTracesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageTracesById provides operations to manage the collection of messageTrace entities.
func (m *GraphBaseServiceClient) MessageTracesById(id string)(*i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.MessageTraceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageTrace%2Did"] = id
    }
    return i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.NewMessageTraceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobilityManagementPolicies provides operations to manage the collection of mobilityManagementPolicy entities.
func (m *GraphBaseServiceClient) MobilityManagementPolicies()(*i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.MobilityManagementPoliciesRequestBuilder) {
    return i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.NewMobilityManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobilityManagementPoliciesById provides operations to manage the collection of mobilityManagementPolicy entities.
func (m *GraphBaseServiceClient) MobilityManagementPoliciesById(id string)(*i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.MobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.NewMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Monitoring provides operations to manage the monitoring singleton.
func (m *GraphBaseServiceClient) Monitoring()(*i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.MonitoringRequestBuilder) {
    return i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.NewMonitoringRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrants provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrants()(*ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.Oauth2PermissionGrantsRequestBuilder) {
    return ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById provides operations to manage the collection of oAuth2PermissionGrant entities.
func (m *GraphBaseServiceClient) Oauth2PermissionGrantsById(id string)(*ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OfficeConfiguration provides operations to manage the officeConfiguration singleton.
func (m *GraphBaseServiceClient) OfficeConfiguration()(*ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f.OfficeConfigurationRequestBuilder) {
    return ibcc2ba1bd11f45f4381ae4007e619ef47b74dbf731b922816aff1ed4d5d47a0f.NewOfficeConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesPublishingProfiles provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *GraphBaseServiceClient) OnPremisesPublishingProfiles()(*i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.OnPremisesPublishingProfilesRequestBuilder) {
    return i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.NewOnPremisesPublishingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesPublishingProfilesById provides operations to manage the collection of onPremisesPublishingProfile entities.
func (m *GraphBaseServiceClient) OnPremisesPublishingProfilesById(id string)(*i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.OnPremisesPublishingProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesPublishingProfile%2Did"] = id
    }
    return i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.NewOnPremisesPublishingProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Organization provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) Organization()(*ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.OrganizationRequestBuilder) {
    return ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.NewOrganizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrganizationById provides operations to manage the collection of organization entities.
func (m *GraphBaseServiceClient) OrganizationById(id string)(*ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.OrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["organization%2Did"] = id
    }
    return ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.NewOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PayloadResponse provides operations to manage the collection of payloadResponse entities.
func (m *GraphBaseServiceClient) PayloadResponse()(*i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.PayloadResponseRequestBuilder) {
    return i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.NewPayloadResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PayloadResponseById provides operations to manage the collection of payloadResponse entities.
func (m *GraphBaseServiceClient) PayloadResponseById(id string)(*i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.PayloadResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["payloadResponse%2Did"] = id
    }
    return i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.NewPayloadResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PermissionGrants provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrants()(*i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.PermissionGrantsRequestBuilder) {
    return i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.NewPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the collection of resourceSpecificPermissionGrant entities.
func (m *GraphBaseServiceClient) PermissionGrantsById(id string)(*i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Places provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) Places()(*ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.PlacesRequestBuilder) {
    return ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.NewPlacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlacesById provides operations to manage the collection of place entities.
func (m *GraphBaseServiceClient) PlacesById(id string)(*ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.PlaceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["place%2Did"] = id
    }
    return ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.NewPlaceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner singleton.
func (m *GraphBaseServiceClient) Planner()(*i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.PlannerRequestBuilder) {
    return i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Policies provides operations to manage the policyRoot singleton.
func (m *GraphBaseServiceClient) Policies()(*i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.PoliciesRequestBuilder) {
    return i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.NewPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Print provides operations to manage the print singleton.
func (m *GraphBaseServiceClient) Print()(*i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.PrintRequestBuilder) {
    return i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.NewPrintRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Privacy provides operations to manage the privacy singleton.
func (m *GraphBaseServiceClient) Privacy()(*i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.PrivacyRequestBuilder) {
    return i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.NewPrivacyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedAccess provides operations to manage the collection of privilegedAccess entities.
func (m *GraphBaseServiceClient) PrivilegedAccess()(*i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.PrivilegedAccessRequestBuilder) {
    return i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.NewPrivilegedAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedAccessById provides operations to manage the collection of privilegedAccess entities.
func (m *GraphBaseServiceClient) PrivilegedAccessById(id string)(*i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.PrivilegedAccessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccess%2Did"] = id
    }
    return i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.NewPrivilegedAccessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedApproval provides operations to manage the collection of privilegedApproval entities.
func (m *GraphBaseServiceClient) PrivilegedApproval()(*i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.PrivilegedApprovalRequestBuilder) {
    return i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.NewPrivilegedApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedApprovalById provides operations to manage the collection of privilegedApproval entities.
func (m *GraphBaseServiceClient) PrivilegedApprovalById(id string)(*i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.PrivilegedApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedApproval%2Did"] = id
    }
    return i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.NewPrivilegedApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedOperationEvents provides operations to manage the collection of privilegedOperationEvent entities.
func (m *GraphBaseServiceClient) PrivilegedOperationEvents()(*i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.PrivilegedOperationEventsRequestBuilder) {
    return i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.NewPrivilegedOperationEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedOperationEventsById provides operations to manage the collection of privilegedOperationEvent entities.
func (m *GraphBaseServiceClient) PrivilegedOperationEventsById(id string)(*i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.PrivilegedOperationEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedOperationEvent%2Did"] = id
    }
    return i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.NewPrivilegedOperationEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedRoleAssignmentRequests provides operations to manage the collection of privilegedRoleAssignmentRequest entities.
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentRequests()(*if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.PrivilegedRoleAssignmentRequestsRequestBuilder) {
    return if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.NewPrivilegedRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedRoleAssignmentRequestsById provides operations to manage the collection of privilegedRoleAssignmentRequest entities.
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentRequestsById(id string)(*if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.PrivilegedRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignmentRequest%2Did"] = id
    }
    return if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.NewPrivilegedRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedRoleAssignments provides operations to manage the collection of privilegedRoleAssignment entities.
func (m *GraphBaseServiceClient) PrivilegedRoleAssignments()(*i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.PrivilegedRoleAssignmentsRequestBuilder) {
    return i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.NewPrivilegedRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedRoleAssignmentsById provides operations to manage the collection of privilegedRoleAssignment entities.
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentsById(id string)(*i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment%2Did"] = id
    }
    return i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedRoles provides operations to manage the collection of privilegedRole entities.
func (m *GraphBaseServiceClient) PrivilegedRoles()(*i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.PrivilegedRolesRequestBuilder) {
    return i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.NewPrivilegedRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedRolesById provides operations to manage the collection of privilegedRole entities.
func (m *GraphBaseServiceClient) PrivilegedRolesById(id string)(*i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.PrivilegedRoleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRole%2Did"] = id
    }
    return i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.NewPrivilegedRoleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrivilegedSignupStatus provides operations to manage the collection of privilegedSignupStatus entities.
func (m *GraphBaseServiceClient) PrivilegedSignupStatus()(*i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.PrivilegedSignupStatusRequestBuilder) {
    return i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.NewPrivilegedSignupStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrivilegedSignupStatusById provides operations to manage the collection of privilegedSignupStatus entities.
func (m *GraphBaseServiceClient) PrivilegedSignupStatusById(id string)(*i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.PrivilegedSignupStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedSignupStatus%2Did"] = id
    }
    return i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.NewPrivilegedSignupStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ProgramControls provides operations to manage the collection of programControl entities.
func (m *GraphBaseServiceClient) ProgramControls()(*i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.ProgramControlsRequestBuilder) {
    return i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.NewProgramControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProgramControlsById provides operations to manage the collection of programControl entities.
func (m *GraphBaseServiceClient) ProgramControlsById(id string)(*i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.ProgramControlItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["programControl%2Did"] = id
    }
    return i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.NewProgramControlItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ProgramControlTypes provides operations to manage the collection of programControlType entities.
func (m *GraphBaseServiceClient) ProgramControlTypes()(*if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.ProgramControlTypesRequestBuilder) {
    return if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.NewProgramControlTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProgramControlTypesById provides operations to manage the collection of programControlType entities.
func (m *GraphBaseServiceClient) ProgramControlTypesById(id string)(*if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.ProgramControlTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["programControlType%2Did"] = id
    }
    return if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.NewProgramControlTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Programs provides operations to manage the collection of program entities.
func (m *GraphBaseServiceClient) Programs()(*ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.ProgramsRequestBuilder) {
    return ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.NewProgramsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProgramsById provides operations to manage the collection of program entities.
func (m *GraphBaseServiceClient) ProgramsById(id string)(*ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.ProgramItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["program%2Did"] = id
    }
    return ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.NewProgramItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports provides operations to manage the reportRoot singleton.
func (m *GraphBaseServiceClient) Reports()(*ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.ReportsRequestBuilder) {
    return ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskDetections provides operations to manage the collection of riskDetection entities.
func (m *GraphBaseServiceClient) RiskDetections()(*i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.RiskDetectionsRequestBuilder) {
    return i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.NewRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskDetectionsById provides operations to manage the collection of riskDetection entities.
func (m *GraphBaseServiceClient) RiskDetectionsById(id string)(*i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.RiskDetectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskDetection%2Did"] = id
    }
    return i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.NewRiskDetectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RiskyUsers provides operations to manage the collection of riskyUser entities.
func (m *GraphBaseServiceClient) RiskyUsers()(*iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.RiskyUsersRequestBuilder) {
    return iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskyUsersById provides operations to manage the collection of riskyUser entities.
func (m *GraphBaseServiceClient) RiskyUsersById(id string)(*iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.RiskyUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUser%2Did"] = id
    }
    return iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.NewRiskyUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagement provides operations to manage the roleManagement singleton.
func (m *GraphBaseServiceClient) RoleManagement()(*i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.RoleManagementRequestBuilder) {
    return i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.NewRoleManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchemaExtensions provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensions()(*i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.SchemaExtensionsRequestBuilder) {
    return i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.NewSchemaExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchemaExtensionsById provides operations to manage the collection of schemaExtension entities.
func (m *GraphBaseServiceClient) SchemaExtensionsById(id string)(*i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.SchemaExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schemaExtension%2Did"] = id
    }
    return i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.NewSchemaExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ScopedRoleMemberships provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMemberships()(*i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.ScopedRoleMembershipsRequestBuilder) {
    return i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.NewScopedRoleMembershipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMembershipsById provides operations to manage the collection of scopedRoleMembership entities.
func (m *GraphBaseServiceClient) ScopedRoleMembershipsById(id string)(*i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Search provides operations to manage the searchEntity singleton.
func (m *GraphBaseServiceClient) Search()(*i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.SearchRequestBuilder) {
    return i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.NewSearchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Security provides operations to manage the security singleton.
func (m *GraphBaseServiceClient) Security()(*i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.SecurityRequestBuilder) {
    return i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipals provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipals()(*ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.ServicePrincipalsRequestBuilder) {
    return ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.NewServicePrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalsById provides operations to manage the collection of servicePrincipal entities.
func (m *GraphBaseServiceClient) ServicePrincipalsById(id string)(*ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.ServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.NewServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings provides operations to manage the collection of directorySetting entities.
func (m *GraphBaseServiceClient) Settings()(*i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.SettingsRequestBuilder) {
    return i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the collection of directorySetting entities.
func (m *GraphBaseServiceClient) SettingsById(id string)(*i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.DirectorySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting%2Did"] = id
    }
    return i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.NewDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Shares provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) Shares()(*i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.SharesRequestBuilder) {
    return i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharesById provides operations to manage the collection of sharedDriveItem entities.
func (m *GraphBaseServiceClient) SharesById(id string)(*i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.SharedDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedDriveItem%2Did"] = id
    }
    return i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.NewSharedDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) Sites()(*i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.SitesRequestBuilder) {
    return i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the collection of site entities.
func (m *GraphBaseServiceClient) SitesById(id string)(*i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Solutions provides operations to manage the solutionsRoot singleton.
func (m *GraphBaseServiceClient) Solutions()(*i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.SolutionsRequestBuilder) {
    return i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.NewSolutionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Storage provides operations to manage the storage singleton.
func (m *GraphBaseServiceClient) Storage()(*id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.StorageRequestBuilder) {
    return id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.NewStorageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribedSkus provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkus()(*ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.SubscribedSkusRequestBuilder) {
    return ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.NewSubscribedSkusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribedSkusById provides operations to manage the collection of subscribedSku entities.
func (m *GraphBaseServiceClient) SubscribedSkusById(id string)(*ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.SubscribedSkuItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscribedSku%2Did"] = id
    }
    return ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.NewSubscribedSkuItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) Subscriptions()(*i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.SubscriptionsRequestBuilder) {
    return i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the collection of subscription entities.
func (m *GraphBaseServiceClient) SubscriptionsById(id string)(*i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teams provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) Teams()(*iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.TeamsRequestBuilder) {
    return iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.NewTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsById provides operations to manage the collection of team entities.
func (m *GraphBaseServiceClient) TeamsById(id string)(*iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.TeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.NewTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TeamsTemplates provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplates()(*i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.TeamsTemplatesRequestBuilder) {
    return i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.NewTeamsTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamsTemplatesById provides operations to manage the collection of teamsTemplate entities.
func (m *GraphBaseServiceClient) TeamsTemplatesById(id string)(*i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.TeamsTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTemplate%2Did"] = id
    }
    return i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.NewTeamsTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TeamTemplateDefinition provides operations to manage the collection of teamTemplateDefinition entities.
func (m *GraphBaseServiceClient) TeamTemplateDefinition()(*i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.TeamTemplateDefinitionRequestBuilder) {
    return i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.NewTeamTemplateDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeamTemplateDefinitionById provides operations to manage the collection of teamTemplateDefinition entities.
func (m *GraphBaseServiceClient) TeamTemplateDefinitionById(id string)(*i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.TeamTemplateDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamTemplateDefinition%2Did"] = id
    }
    return i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.NewTeamTemplateDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teamwork provides operations to manage the teamwork singleton.
func (m *GraphBaseServiceClient) Teamwork()(*icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.TeamworkRequestBuilder) {
    return icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantRelationships provides operations to manage the tenantRelationship singleton.
func (m *GraphBaseServiceClient) TenantRelationships()(*id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.TenantRelationshipsRequestBuilder) {
    return id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.NewTenantRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermStore provides operations to manage the store singleton.
func (m *GraphBaseServiceClient) TermStore()(*i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.TermStoreRequestBuilder) {
    return i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatSubmission provides operations to manage the threatSubmissionRoot singleton.
func (m *GraphBaseServiceClient) ThreatSubmission()(*i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.ThreatSubmissionRequestBuilder) {
    return i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.NewThreatSubmissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TrustFramework provides operations to manage the trustFramework singleton.
func (m *GraphBaseServiceClient) TrustFramework()(*i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.TrustFrameworkRequestBuilder) {
    return i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.NewTrustFrameworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Users provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) Users()(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UsersRequestBuilder) {
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) UsersById(id string)(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
