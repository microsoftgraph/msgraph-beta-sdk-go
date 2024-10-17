package msgraphbetasdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
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
    i164b4c6703708dd1b0670a17a07a4dd64e49cb0c1cb66e50a3146217db57a57a "github.com/microsoftgraph/msgraph-beta-sdk-go/workplace"
    i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontrols"
    i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/messageevents"
    i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a "github.com/microsoftgraph/msgraph-beta-sdk-go/schemaextensions"
    i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamtemplatedefinition"
    i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews"
    i1feaf5e7874bc1012d47d3fb9128ea97140f17a0526b7e8dd267bdb120026eae "github.com/microsoftgraph/msgraph-beta-sdk-go/applicationswithappid"
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
    i55ae935853de1792fdb46948985a94b1dba6c75d346e01b97f5bd49309b8959b "github.com/microsoftgraph/msgraph-beta-sdk-go/applicationswithuniquename"
    i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053 "github.com/microsoftgraph/msgraph-beta-sdk-go/messagerecipients"
    i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e "github.com/microsoftgraph/msgraph-beta-sdk-go/search"
    i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroletemplates"
    i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53 "github.com/microsoftgraph/msgraph-beta-sdk-go/domaindnsrecords"
    i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54 "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts"
    i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies"
    i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscriptions"
    i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval"
    i64fc24bb18135374ce200dde4281c7742414b6f0b008ed5c2eab2c2787c0619e "github.com/microsoftgraph/msgraph-beta-sdk-go/placeswithplaceid"
    i65ec2d26e2745c20b2909e217069fc9c66fc5d6b449f4f50e1f705a550f3e781 "github.com/microsoftgraph/msgraph-beta-sdk-go/filteringpolicies"
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
    i74f9164c3d2d3eadced923419244ab3c32717c28ae93424103196d4b781b759e "github.com/microsoftgraph/msgraph-beta-sdk-go/groupswithuniquename"
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
    i8f15afd67f0af97c484d7c65f1ee4d999eb4ea9fe08743b8099d031e04b49baf "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceswithdeviceid"
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
    ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals"
    ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingbusinesses"
    ibfdc6364c06d072aa2aefd5744c668e9033d90ef2a5b4075da3081bc413ea379 "github.com/microsoftgraph/msgraph-beta-sdk-go/networkaccess"
    ic095b493ce132d5d76dd2e98bfcff99e08f7de07490051885070d98614035110 "github.com/microsoftgraph/msgraph-beta-sdk-go/network"
    ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/places"
    icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork"
    icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
    id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821 "github.com/microsoftgraph/msgraph-beta-sdk-go/storage"
    id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/certificatebasedauthconfiguration"
    id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships"
    id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education"
    ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection"
    idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement"
    ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications"
    ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad "github.com/microsoftgraph/msgraph-beta-sdk-go/oauth2permissiongrants"
    ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd "github.com/microsoftgraph/msgraph-beta-sdk-go/bookingcurrencies"
    ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf "github.com/microsoftgraph/msgraph-beta-sdk-go/organization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/subscribedskus"
    if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignmentrequests"
    if298592eede46286a454444191c80a3a678c2136e14beb3f7c77b80b0b3ed812 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipalswithappid"
    if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification"
    if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e "github.com/microsoftgraph/msgraph-beta-sdk-go/approvalworkflowproviders"
    if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538 "github.com/microsoftgraph/msgraph-beta-sdk-go/programcontroltypes"
    if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications"
    ifc2b4dc2959942f88a5b666d62b682ecb910ab6189042234241247fd57d38877 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroleswithroletemplateid"
    ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219 "github.com/microsoftgraph/msgraph-beta-sdk-go/app"
    ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/governancesubjects"
    iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewDecisions provides operations to manage the collection of accessReviewDecision entities.
// returns a *AccessReviewDecisionsRequestBuilder when successful
func (m *GraphBaseServiceClient) AccessReviewDecisions()(*i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.AccessReviewDecisionsRequestBuilder) {
    return i424c8587488111febed6b9b4d2ad6d5226fb83e28676c38366dd47f76f319ffb.NewAccessReviewDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessReviews provides operations to manage the collection of accessReview entities.
// returns a *AccessReviewsRequestBuilder when successful
func (m *GraphBaseServiceClient) AccessReviews()(*i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.AccessReviewsRequestBuilder) {
    return i1dfcb6e17563ae78b6dbaf02d32cee89099a7795106760d7d401df42ce73b8fc.NewAccessReviewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Activitystatistics provides operations to manage the collection of activityStatistics entities.
// returns a *ActivitystatisticsRequestBuilder when successful
func (m *GraphBaseServiceClient) Activitystatistics()(*i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.ActivitystatisticsRequestBuilder) {
    return i8d3c03812535daaab5e9e28f499097b08f09a8a7ab62e664ebf24dd8af17e77c.NewActivitystatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Admin provides operations to manage the admin singleton.
// returns a *AdminRequestBuilder when successful
func (m *GraphBaseServiceClient) Admin()(*i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.AdminRequestBuilder) {
    return i65ed27543dee9887d3df7d7d883303dfead48cba6be4e357fa7d5c21332b4626.NewAdminRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AdministrativeUnits provides operations to manage the collection of administrativeUnit entities.
// returns a *AdministrativeUnitsRequestBuilder when successful
func (m *GraphBaseServiceClient) AdministrativeUnits()(*i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.AdministrativeUnitsRequestBuilder) {
    return i895a6f3a85eea8558280363fb928ce992d75c89f1c602b57f1d6352b0af78e5f.NewAdministrativeUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AgreementAcceptances provides operations to manage the collection of agreementAcceptance entities.
// returns a *AgreementAcceptancesRequestBuilder when successful
func (m *GraphBaseServiceClient) AgreementAcceptances()(*i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.AgreementAcceptancesRequestBuilder) {
    return i15e329825c659329448e12b30278e3a09efd68996edb65e6eb37bbba528b21d7.NewAgreementAcceptancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Agreements provides operations to manage the collection of agreement entities.
// returns a *AgreementsRequestBuilder when successful
func (m *GraphBaseServiceClient) Agreements()(*i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.AgreementsRequestBuilder) {
    return i493c694f665c6b8116f1d28cef9c35839e2b3810e4a8c9f326bfc1b2caa30afa.NewAgreementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AllowedDataLocations provides operations to manage the collection of allowedDataLocation entities.
// returns a *AllowedDataLocationsRequestBuilder when successful
func (m *GraphBaseServiceClient) AllowedDataLocations()(*i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.AllowedDataLocationsRequestBuilder) {
    return i8283dbd0b9136717b100e1ef6f4ac05d85e3412714dd0dea4b38d6de05ef4349.NewAllowedDataLocationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App provides operations to manage the commsApplication singleton.
// returns a *AppRequestBuilder when successful
func (m *GraphBaseServiceClient) App()(*ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.AppRequestBuilder) {
    return ifc59747dbaa83f8f51942823114f4abfa41e0c0a64d67957f17e6b60407ce219.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppCatalogs provides operations to manage the appCatalogs singleton.
// returns a *AppCatalogsRequestBuilder when successful
func (m *GraphBaseServiceClient) AppCatalogs()(*i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.AppCatalogsRequestBuilder) {
    return i2130b9a37453c245bc87d9a83666a92560714fc5bb3c0f5a77e999639d2f4e45.NewAppCatalogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications provides operations to manage the collection of application entities.
// returns a *ApplicationsRequestBuilder when successful
func (m *GraphBaseServiceClient) Applications()(*ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.ApplicationsRequestBuilder) {
    return ie1b2fd35e4b1f7cbc7bd808e462c966c4ec16a274923b50216bdd8a2ae0a3129.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplicationsWithAppId provides operations to manage the collection of application entities.
// returns a *ApplicationsWithAppIdRequestBuilder when successful
func (m *GraphBaseServiceClient) ApplicationsWithAppId(appId *string)(*i1feaf5e7874bc1012d47d3fb9128ea97140f17a0526b7e8dd267bdb120026eae.ApplicationsWithAppIdRequestBuilder) {
    return i1feaf5e7874bc1012d47d3fb9128ea97140f17a0526b7e8dd267bdb120026eae.NewApplicationsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// ApplicationsWithUniqueName provides operations to manage the collection of application entities.
// returns a *ApplicationsWithUniqueNameRequestBuilder when successful
func (m *GraphBaseServiceClient) ApplicationsWithUniqueName(uniqueName *string)(*i55ae935853de1792fdb46948985a94b1dba6c75d346e01b97f5bd49309b8959b.ApplicationsWithUniqueNameRequestBuilder) {
    return i55ae935853de1792fdb46948985a94b1dba6c75d346e01b97f5bd49309b8959b.NewApplicationsWithUniqueNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, uniqueName)
}
// ApplicationTemplates provides operations to manage the collection of applicationTemplate entities.
// returns a *ApplicationTemplatesRequestBuilder when successful
func (m *GraphBaseServiceClient) ApplicationTemplates()(*i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.ApplicationTemplatesRequestBuilder) {
    return i9fb9a4d9d99571d2cc1de51809c0dfccf1dae8bd81c7eb39e51d1382c2ec81ba.NewApplicationTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppRoleAssignments provides operations to manage the collection of appRoleAssignment entities.
// returns a *AppRoleAssignmentsRequestBuilder when successful
func (m *GraphBaseServiceClient) AppRoleAssignments()(*i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.AppRoleAssignmentsRequestBuilder) {
    return i35277464b5f866fcf2cb5324cd283216c9f6e9fc22956c71cb5b11c4ab649a29.NewAppRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApprovalWorkflowProviders provides operations to manage the collection of approvalWorkflowProvider entities.
// returns a *ApprovalWorkflowProvidersRequestBuilder when successful
func (m *GraphBaseServiceClient) ApprovalWorkflowProviders()(*if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.ApprovalWorkflowProvidersRequestBuilder) {
    return if4a9faac580b9d5510ead8eac155e0cecb2222152b913f0bedc9a44bbe2ee79e.NewApprovalWorkflowProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditLogs provides operations to manage the auditLogRoot singleton.
// returns a *AuditLogsRequestBuilder when successful
func (m *GraphBaseServiceClient) AuditLogs()(*i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.AuditLogsRequestBuilder) {
    return i6c3f8c4b4b571cf0fbb7c7c8791ae736e28cc3f4bb62262698b6291c13e127b9.NewAuditLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuthenticationMethodConfigurations provides operations to manage the collection of authenticationMethodConfiguration entities.
// returns a *AuthenticationMethodConfigurationsRequestBuilder when successful
func (m *GraphBaseServiceClient) AuthenticationMethodConfigurations()(*i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.AuthenticationMethodConfigurationsRequestBuilder) {
    return i8e667c6208be96da3103b8806ff97028502c18052414fe99a224c1565834ca0f.NewAuthenticationMethodConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy singleton.
// returns a *AuthenticationMethodsPolicyRequestBuilder when successful
func (m *GraphBaseServiceClient) AuthenticationMethodsPolicy()(*i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.AuthenticationMethodsPolicyRequestBuilder) {
    return i39dbae52481ac3c9530d9fae0a2292348b8f7327bab28ea21183045324adadbc.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BookingBusinesses provides operations to manage the collection of bookingBusiness entities.
// returns a *BookingBusinessesRequestBuilder when successful
func (m *GraphBaseServiceClient) BookingBusinesses()(*ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.BookingBusinessesRequestBuilder) {
    return ibf9394d7c54feda53ca523241dde659e8725041c25384ede68e72731d68d5abe.NewBookingBusinessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BookingCurrencies provides operations to manage the collection of bookingCurrency entities.
// returns a *BookingCurrenciesRequestBuilder when successful
func (m *GraphBaseServiceClient) BookingCurrencies()(*ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.BookingCurrenciesRequestBuilder) {
    return ie573fcc25112f624100d67f5c4380ef23bdf060e2876e90cec1bf404deffc3dd.NewBookingCurrenciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BusinessFlowTemplates provides operations to manage the collection of businessFlowTemplate entities.
// returns a *BusinessFlowTemplatesRequestBuilder when successful
func (m *GraphBaseServiceClient) BusinessFlowTemplates()(*i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.BusinessFlowTemplatesRequestBuilder) {
    return i97c9750160852aa25d52a4c6fa196b644ce728c6645ca520427ff4d85c76afa0.NewBusinessFlowTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CertificateBasedAuthConfiguration provides operations to manage the collection of certificateBasedAuthConfiguration entities.
// returns a *CertificateBasedAuthConfigurationRequestBuilder when successful
func (m *GraphBaseServiceClient) CertificateBasedAuthConfiguration()(*id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.CertificateBasedAuthConfigurationRequestBuilder) {
    return id53bdaa191b823f3e2f4009f4cc095b46d1c7a433bde3b6d09ef0bd8df3514c2.NewCertificateBasedAuthConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Chats provides operations to manage the collection of chat entities.
// returns a *ChatsRequestBuilder when successful
func (m *GraphBaseServiceClient) Chats()(*i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.ChatsRequestBuilder) {
    return i2bd7b88e5d2b5a20231449a09cd2703774075c53c19dc28ca3829e91d51dfd73.NewChatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commands provides operations to manage the collection of command entities.
// returns a *CommandsRequestBuilder when successful
func (m *GraphBaseServiceClient) Commands()(*i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.CommandsRequestBuilder) {
    return i81513db44a7796569414887f1110e1d158078cb199b3565d960425af9c2904ba.NewCommandsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Communications provides operations to manage the cloudCommunications singleton.
// returns a *CommunicationsRequestBuilder when successful
func (m *GraphBaseServiceClient) Communications()(*if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.CommunicationsRequestBuilder) {
    return if7bcb57951e8f2ae550fcf781dc209ed777854429fcfe2465a71b03112dfc346.NewCommunicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Compliance provides operations to manage the compliance singleton.
// returns a *ComplianceRequestBuilder when successful
func (m *GraphBaseServiceClient) Compliance()(*ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.ComplianceRequestBuilder) {
    return ia93a0d553904ea998c5c0e30cf6d6574ba3a42be402a2d1eb88d5bd221d1a0d6.NewComplianceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Connections provides operations to manage the collection of externalConnection entities.
// returns a *ConnectionsRequestBuilder when successful
func (m *GraphBaseServiceClient) Connections()(*i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.ConnectionsRequestBuilder) {
    return i29b3182886b8fc3d309db2628f3911e671f6745e9fdab71d07cdeb487c4453ad.NewConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGraphBaseServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewGraphBaseServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactory)(*GraphBaseServiceClient) {
    m := &GraphBaseServiceClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    m.BaseRequestBuilder.RequestAdapter.EnableBackingStore(backingStore);
    return m
}
// Contacts provides operations to manage the collection of orgContact entities.
// returns a *ContactsRequestBuilder when successful
func (m *GraphBaseServiceClient) Contacts()(*i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.ContactsRequestBuilder) {
    return i716e3204a4c47d24737c05f3b4c2ef2462fa5a1df29b57365f338e8f68ee16ef.NewContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Contracts provides operations to manage the collection of contract entities.
// returns a *ContractsRequestBuilder when successful
func (m *GraphBaseServiceClient) Contracts()(*i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.ContractsRequestBuilder) {
    return i5e88b3e8025e26d21777096f2c7525a182545bb8dc0137634b047a10fe14ab54.NewContractsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataClassification provides operations to manage the dataClassificationService singleton.
// returns a *DataClassificationRequestBuilder when successful
func (m *GraphBaseServiceClient) DataClassification()(*if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.DataClassificationRequestBuilder) {
    return if3c2dea1db099d8f9ce8b623a12f6291f276e2bbb50259658f584a3a85cf71b8.NewDataClassificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataPolicyOperations provides operations to manage the collection of dataPolicyOperation entities.
// returns a *DataPolicyOperationsRequestBuilder when successful
func (m *GraphBaseServiceClient) DataPolicyOperations()(*i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.DataPolicyOperationsRequestBuilder) {
    return i0f747ff1f24810ff51160697ed4229c9ca192f7b84644311b88fa3b475cc340d.NewDataPolicyOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceAppManagement provides operations to manage the deviceAppManagement singleton.
// returns a *DeviceAppManagementRequestBuilder when successful
func (m *GraphBaseServiceClient) DeviceAppManagement()(*idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.DeviceAppManagementRequestBuilder) {
    return idc4afe653def183ef95500aa004f556fdf3c3747771f17c8472ca3cad61cebf4.NewDeviceAppManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagement provides operations to manage the deviceManagement singleton.
// returns a *DeviceManagementRequestBuilder when successful
func (m *GraphBaseServiceClient) DeviceManagement()(*i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.DeviceManagementRequestBuilder) {
    return i09893664b20e7c846b2bc7aaaf1cd7f554ed3d2c00ac11336bea4c3c3d859e09.NewDeviceManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Devices provides operations to manage the collection of device entities.
// returns a *DevicesRequestBuilder when successful
func (m *GraphBaseServiceClient) Devices()(*i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.DevicesRequestBuilder) {
    return i0b4892b2f92a31e44541567b8065e8e7760cb336e17d7dacb9120a865d5b0a37.NewDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DevicesWithDeviceId provides operations to manage the collection of device entities.
// returns a *DevicesWithDeviceIdRequestBuilder when successful
func (m *GraphBaseServiceClient) DevicesWithDeviceId(deviceId *string)(*i8f15afd67f0af97c484d7c65f1ee4d999eb4ea9fe08743b8099d031e04b49baf.DevicesWithDeviceIdRequestBuilder) {
    return i8f15afd67f0af97c484d7c65f1ee4d999eb4ea9fe08743b8099d031e04b49baf.NewDevicesWithDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, deviceId)
}
// Directory provides operations to manage the directory singleton.
// returns a *DirectoryRequestBuilder when successful
func (m *GraphBaseServiceClient) Directory()(*i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.DirectoryRequestBuilder) {
    return i6e398703c86ec3814400d80161079e7253c4e25f4ba1adb0c8d31da236f7bcd7.NewDirectoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DirectoryObjects provides operations to manage the collection of directoryObject entities.
// returns a *DirectoryObjectsRequestBuilder when successful
func (m *GraphBaseServiceClient) DirectoryObjects()(*iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.DirectoryObjectsRequestBuilder) {
    return iaec68a3d2c3ba0a78ebb66cd93fd1c5d2a6e0450b97a0cf19d94cb58956bec1d.NewDirectoryObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DirectoryRoles provides operations to manage the collection of directoryRole entities.
// returns a *DirectoryRolesRequestBuilder when successful
func (m *GraphBaseServiceClient) DirectoryRoles()(*i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.DirectoryRolesRequestBuilder) {
    return i20621ebd49d2bb1ed6c592ae35dfa701db30564a91ff100d25b0dcdb142bd942.NewDirectoryRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DirectoryRolesWithRoleTemplateId provides operations to manage the collection of directoryRole entities.
// returns a *DirectoryRolesWithRoleTemplateIdRequestBuilder when successful
func (m *GraphBaseServiceClient) DirectoryRolesWithRoleTemplateId(roleTemplateId *string)(*ifc2b4dc2959942f88a5b666d62b682ecb910ab6189042234241247fd57d38877.DirectoryRolesWithRoleTemplateIdRequestBuilder) {
    return ifc2b4dc2959942f88a5b666d62b682ecb910ab6189042234241247fd57d38877.NewDirectoryRolesWithRoleTemplateIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, roleTemplateId)
}
// DirectoryRoleTemplates provides operations to manage the collection of directoryRoleTemplate entities.
// returns a *DirectoryRoleTemplatesRequestBuilder when successful
func (m *GraphBaseServiceClient) DirectoryRoleTemplates()(*i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.DirectoryRoleTemplatesRequestBuilder) {
    return i5b4eb770497618728398e41e6ed415ad2b92d20f7ad45ba75277a5800d9a2a12.NewDirectoryRoleTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DirectorySettingTemplates provides operations to manage the collection of directorySettingTemplate entities.
// returns a *DirectorySettingTemplatesRequestBuilder when successful
func (m *GraphBaseServiceClient) DirectorySettingTemplates()(*iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.DirectorySettingTemplatesRequestBuilder) {
    return iabd30d9ba6ae302fd2d3145e0da70036496a2167b9e0cbc049bab96d9d9a29b3.NewDirectorySettingTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DomainDnsRecords provides operations to manage the collection of domainDnsRecord entities.
// returns a *DomainDnsRecordsRequestBuilder when successful
func (m *GraphBaseServiceClient) DomainDnsRecords()(*i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.DomainDnsRecordsRequestBuilder) {
    return i5db103717157eeba829a6a8580002d8b9adfdfd549628f3c84152bf3164a7b53.NewDomainDnsRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Domains provides operations to manage the collection of domain entities.
// returns a *DomainsRequestBuilder when successful
func (m *GraphBaseServiceClient) Domains()(*i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.DomainsRequestBuilder) {
    return i9e22e53c888822daa9264a72be4d11f335e3170e9198aae4bba758214e319857.NewDomainsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the collection of drive entities.
// returns a *DrivesRequestBuilder when successful
func (m *GraphBaseServiceClient) Drives()(*i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.DrivesRequestBuilder) {
    return i2b9cd6123f9a7ca2d7c253973e81bb3965ee0b78a350919aafe66d7b73c70433.NewDrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Education provides operations to manage the educationRoot singleton.
// returns a *EducationRequestBuilder when successful
func (m *GraphBaseServiceClient) Education()(*id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.EducationRequestBuilder) {
    return id90d135edac1f1a3e952db4ad985001105d2e7c0133f8cc410765eb1af789cc0.NewEducationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmployeeExperience provides operations to manage the employeeExperience singleton.
// returns a *EmployeeExperienceRequestBuilder when successful
func (m *GraphBaseServiceClient) EmployeeExperience()(*i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.EmployeeExperienceRequestBuilder) {
    return i4b923a0ffb143d5def24980fce5d55c68c9634d5b55c33bb0b3029ac68415dd1.NewEmployeeExperienceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// External provides operations to manage the external singleton.
// returns a *ExternalRequestBuilder when successful
func (m *GraphBaseServiceClient) External()(*i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.ExternalRequestBuilder) {
    return i453454978d87486fd201df62ea4b775c5b4907e2a36395fb6fb1e9060fc3f1bb.NewExternalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilteringPolicies provides operations to manage the collection of filteringPolicy entities.
// returns a *FilteringPoliciesRequestBuilder when successful
func (m *GraphBaseServiceClient) FilteringPolicies()(*i65ec2d26e2745c20b2909e217069fc9c66fc5d6b449f4f50e1f705a550f3e781.FilteringPoliciesRequestBuilder) {
    return i65ec2d26e2745c20b2909e217069fc9c66fc5d6b449f4f50e1f705a550f3e781.NewFilteringPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterOperators provides operations to manage the collection of filterOperatorSchema entities.
// returns a *FilterOperatorsRequestBuilder when successful
func (m *GraphBaseServiceClient) FilterOperators()(*i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.FilterOperatorsRequestBuilder) {
    return i0b9d70018d3c267f9f34a818ce43cc889d06d87749a70e1ad1d45eead0c735e0.NewFilterOperatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Financials provides operations to manage the financials singleton.
// returns a *FinancialsRequestBuilder when successful
func (m *GraphBaseServiceClient) Financials()(*i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.FinancialsRequestBuilder) {
    return i0281e791d6eedd5b4b6d3c18b07336722b6152b5db6a2e9dc8385e98f565f677.NewFinancialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Functions provides operations to manage the collection of attributeMappingFunctionSchema entities.
// returns a *FunctionsRequestBuilder when successful
func (m *GraphBaseServiceClient) Functions()(*i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.FunctionsRequestBuilder) {
    return i054b68521cee54ec767d07cf7a8a6d50a0d24b6e6fc43b8296c34730fd8ca465.NewFunctionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceResources provides operations to manage the collection of governanceResource entities.
// returns a *GovernanceResourcesRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceResources()(*i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.GovernanceResourcesRequestBuilder) {
    return i3b7da1b693d5428b20b0bf3340acb4b879042a9393e45df9349b04a5b2830acb.NewGovernanceResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceRoleAssignmentRequests provides operations to manage the collection of governanceRoleAssignmentRequest entities.
// returns a *GovernanceRoleAssignmentRequestsRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceRoleAssignmentRequests()(*i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.GovernanceRoleAssignmentRequestsRequestBuilder) {
    return i3de07fd1833246d183a4d60c2329c7467381716c2cfdfc096ff627e4e9479cf8.NewGovernanceRoleAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceRoleAssignments provides operations to manage the collection of governanceRoleAssignment entities.
// returns a *GovernanceRoleAssignmentsRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceRoleAssignments()(*i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.GovernanceRoleAssignmentsRequestBuilder) {
    return i14752cfec59ab915e7c63922270765abf65744437d9135c191cef3986f08c3bb.NewGovernanceRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceRoleDefinitions provides operations to manage the collection of governanceRoleDefinition entities.
// returns a *GovernanceRoleDefinitionsRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceRoleDefinitions()(*i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.GovernanceRoleDefinitionsRequestBuilder) {
    return i8a18cb7418541221b2c3fd213a484d9e3029fab916358b16fb24015b078b8eba.NewGovernanceRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceRoleSettings provides operations to manage the collection of governanceRoleSetting entities.
// returns a *GovernanceRoleSettingsRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceRoleSettings()(*ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.GovernanceRoleSettingsRequestBuilder) {
    return ib71e32ed3a7f0f8a512aa55c1428492116ff2d1bae5015a9b89f910ecbc7c6bd.NewGovernanceRoleSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GovernanceSubjects provides operations to manage the collection of governanceSubject entities.
// returns a *GovernanceSubjectsRequestBuilder when successful
func (m *GraphBaseServiceClient) GovernanceSubjects()(*ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.GovernanceSubjectsRequestBuilder) {
    return ifcac309012d761a79a74e6d79fad6979f2117e7af36ff6e5ad131093412afcc7.NewGovernanceSubjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupLifecyclePolicies provides operations to manage the collection of groupLifecyclePolicy entities.
// returns a *GroupLifecyclePoliciesRequestBuilder when successful
func (m *GraphBaseServiceClient) GroupLifecyclePolicies()(*i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.GroupLifecyclePoliciesRequestBuilder) {
    return i71438b4a3f9d4a17f8c873a44b8ac76600403f5ce0cce2423bde35e0191f2c17.NewGroupLifecyclePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Groups provides operations to manage the collection of group entities.
// returns a *GroupsRequestBuilder when successful
func (m *GraphBaseServiceClient) Groups()(*i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupsRequestBuilder) {
    return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupsWithUniqueName provides operations to manage the collection of group entities.
// returns a *GroupsWithUniqueNameRequestBuilder when successful
func (m *GraphBaseServiceClient) GroupsWithUniqueName(uniqueName *string)(*i74f9164c3d2d3eadced923419244ab3c32717c28ae93424103196d4b781b759e.GroupsWithUniqueNameRequestBuilder) {
    return i74f9164c3d2d3eadced923419244ab3c32717c28ae93424103196d4b781b759e.NewGroupsWithUniqueNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, uniqueName)
}
// Identity provides operations to manage the identityContainer singleton.
// returns a *IdentityRequestBuilder when successful
func (m *GraphBaseServiceClient) Identity()(*i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.IdentityRequestBuilder) {
    return i20702653f98186060bd39b9fe8136743eafc0ddaa43435e527665ac75229a33a.NewIdentityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IdentityGovernance provides operations to manage the identityGovernance singleton.
// returns a *IdentityGovernanceRequestBuilder when successful
func (m *GraphBaseServiceClient) IdentityGovernance()(*i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.IdentityGovernanceRequestBuilder) {
    return i7ed0a0c7cc963cb1cf6c282e5e5c04cea2d4959e5f1bfd12698c0196858b1b52.NewIdentityGovernanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IdentityProtection provides operations to manage the identityProtectionRoot singleton.
// returns a *IdentityProtectionRequestBuilder when successful
func (m *GraphBaseServiceClient) IdentityProtection()(*i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.IdentityProtectionRequestBuilder) {
    return i444aa080f7aab1ac35ae55318eca861a558a57b4c5e3cbd071222f511baac5fe.NewIdentityProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IdentityProviders provides operations to manage the collection of identityProvider entities.
// returns a *IdentityProvidersRequestBuilder when successful
func (m *GraphBaseServiceClient) IdentityProviders()(*i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.IdentityProvidersRequestBuilder) {
    return i531b1efd1768fd272d51921ff5812bdeba5b46e0eeec0e4c818250cb7116aed5.NewIdentityProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InformationProtection provides operations to manage the informationProtection singleton.
// returns a *InformationProtectionRequestBuilder when successful
func (m *GraphBaseServiceClient) InformationProtection()(*ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.InformationProtectionRequestBuilder) {
    return ida1e35de3db89946d53da45357ba0e3da7b3e3a1d46b511191582f0d2a917caa.NewInformationProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Invitations provides operations to manage the collection of invitation entities.
// returns a *InvitationsRequestBuilder when successful
func (m *GraphBaseServiceClient) Invitations()(*i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.InvitationsRequestBuilder) {
    return i4cb6edb865a0e38bb1799dcb0c7881b92feed59596c1912cfe5e6142b61f9c91.NewInvitationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MessageEvents provides operations to manage the collection of messageEvent entities.
// returns a *MessageEventsRequestBuilder when successful
func (m *GraphBaseServiceClient) MessageEvents()(*i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.MessageEventsRequestBuilder) {
    return i19717748912ff29c95998304f534371c35864a4579ea92608b32aeecf7d18cc4.NewMessageEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MessageRecipients provides operations to manage the collection of messageRecipient entities.
// returns a *MessageRecipientsRequestBuilder when successful
func (m *GraphBaseServiceClient) MessageRecipients()(*i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.MessageRecipientsRequestBuilder) {
    return i56ad7deac03a612015589ab4fade2313d6df08086c7ee8d46177fc8ddc5b0053.NewMessageRecipientsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MessageTraces provides operations to manage the collection of messageTrace entities.
// returns a *MessageTracesRequestBuilder when successful
func (m *GraphBaseServiceClient) MessageTraces()(*i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.MessageTracesRequestBuilder) {
    return i43a1701c7f8902247fd60e3d6ff36be02d1a59f02884ce5734adb8deb69a47ef.NewMessageTracesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobilityManagementPolicies provides operations to manage the collection of mobilityManagementPolicy entities.
// returns a *MobilityManagementPoliciesRequestBuilder when successful
func (m *GraphBaseServiceClient) MobilityManagementPolicies()(*i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.MobilityManagementPoliciesRequestBuilder) {
    return i05bd1def68419ff4c77b9bdafb42d256c4c19b479003ef8adcd868ec38673e84.NewMobilityManagementPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Monitoring provides operations to manage the monitoring singleton.
// returns a *MonitoringRequestBuilder when successful
func (m *GraphBaseServiceClient) Monitoring()(*i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.MonitoringRequestBuilder) {
    return i3c2d5b7a05b0c58ca1d3b72343c5f03221f2bd894e75f89f62ccf8020a48250b.NewMonitoringRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Network provides operations to manage the network singleton.
// returns a *NetworkRequestBuilder when successful
func (m *GraphBaseServiceClient) Network()(*ic095b493ce132d5d76dd2e98bfcff99e08f7de07490051885070d98614035110.NetworkRequestBuilder) {
    return ic095b493ce132d5d76dd2e98bfcff99e08f7de07490051885070d98614035110.NewNetworkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NetworkAccess provides operations to manage the networkAccessRoot singleton.
// returns a *NetworkAccessRequestBuilder when successful
func (m *GraphBaseServiceClient) NetworkAccess()(*ibfdc6364c06d072aa2aefd5744c668e9033d90ef2a5b4075da3081bc413ea379.NetworkAccessRequestBuilder) {
    return ibfdc6364c06d072aa2aefd5744c668e9033d90ef2a5b4075da3081bc413ea379.NewNetworkAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Oauth2PermissionGrants provides operations to manage the collection of oAuth2PermissionGrant entities.
// returns a *Oauth2PermissionGrantsRequestBuilder when successful
func (m *GraphBaseServiceClient) Oauth2PermissionGrants()(*ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.Oauth2PermissionGrantsRequestBuilder) {
    return ie2e0818e93fcfbb33fde071a9354c3c22bedab0ec20855b7d5232d29bcc65bad.NewOauth2PermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnPremisesPublishingProfiles provides operations to manage the collection of onPremisesPublishingProfile entities.
// returns a *OnPremisesPublishingProfilesRequestBuilder when successful
func (m *GraphBaseServiceClient) OnPremisesPublishingProfiles()(*i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.OnPremisesPublishingProfilesRequestBuilder) {
    return i0d38f6e6ea6126fff7bb7a5c3c2d82fe471d00233209e2b6b2ce6ccb21ce50f5.NewOnPremisesPublishingProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organization provides operations to manage the collection of organization entities.
// returns a *OrganizationRequestBuilder when successful
func (m *GraphBaseServiceClient) Organization()(*ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.OrganizationRequestBuilder) {
    return ie58948149bb028757a64f16376df00cc5a99ad93e0d57affa0ac24ff6d096aaf.NewOrganizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PayloadResponse provides operations to manage the collection of payloadResponse entities.
// returns a *PayloadResponseRequestBuilder when successful
func (m *GraphBaseServiceClient) PayloadResponse()(*i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.PayloadResponseRequestBuilder) {
    return i22f037221f506c5b5751f13095cc17caaf248e93588f883123c452cb4f1ec6a9.NewPayloadResponseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermissionGrants provides operations to manage the collection of resourceSpecificPermissionGrant entities.
// returns a *PermissionGrantsRequestBuilder when successful
func (m *GraphBaseServiceClient) PermissionGrants()(*i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.PermissionGrantsRequestBuilder) {
    return i48a68a7c83dc874f9d9fdf942afed70a34b11f92d6b2ccb439359753116f65cc.NewPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Places the places property
// returns a *PlacesRequestBuilder when successful
func (m *GraphBaseServiceClient) Places()(*ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.PlacesRequestBuilder) {
    return ic21cf429efd6fc3199e67b5b4288a3193ff5e9cfb4e97a5e442e02ccd7748ec1.NewPlacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PlacesWithPlaceId provides operations to manage the collection of place entities.
// returns a *PlacesWithPlaceIdRequestBuilder when successful
func (m *GraphBaseServiceClient) PlacesWithPlaceId(placeId *string)(*i64fc24bb18135374ce200dde4281c7742414b6f0b008ed5c2eab2c2787c0619e.PlacesWithPlaceIdRequestBuilder) {
    return i64fc24bb18135374ce200dde4281c7742414b6f0b008ed5c2eab2c2787c0619e.NewPlacesWithPlaceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, placeId)
}
// Planner provides operations to manage the planner singleton.
// returns a *PlannerRequestBuilder when successful
func (m *GraphBaseServiceClient) Planner()(*i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.PlannerRequestBuilder) {
    return i2c16e7aacab2db540fbc4e11916f266c9c46936118144cc1383e798af27b1b6f.NewPlannerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Policies provides operations to manage the policyRoot singleton.
// returns a *PoliciesRequestBuilder when successful
func (m *GraphBaseServiceClient) Policies()(*i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.PoliciesRequestBuilder) {
    return i5f794bd004f1cc95da776bcb1947ffabf97b71aae1f5c9f15255b24451e2929b.NewPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Print provides operations to manage the print singleton.
// returns a *PrintRequestBuilder when successful
func (m *GraphBaseServiceClient) Print()(*i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.PrintRequestBuilder) {
    return i75b6dc07087cda1a9afc465878b0aa56ca3703a3ed530d5a22119b0960d159d3.NewPrintRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Privacy provides operations to manage the privacy singleton.
// returns a *PrivacyRequestBuilder when successful
func (m *GraphBaseServiceClient) Privacy()(*i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.PrivacyRequestBuilder) {
    return i7daddf90c53ab17a62cc6612f32636f5bea354e4e22c20c5656606d1b491bd03.NewPrivacyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedAccess provides operations to manage the collection of privilegedAccess entities.
// returns a *PrivilegedAccessRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedAccess()(*i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.PrivilegedAccessRequestBuilder) {
    return i2f87335be6a07866636e1f602f5beda6b47bc99969e216fac59efff432ff2345.NewPrivilegedAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedApproval provides operations to manage the collection of privilegedApproval entities.
// returns a *PrivilegedApprovalRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedApproval()(*i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.PrivilegedApprovalRequestBuilder) {
    return i613e9d4b5596faf0dba59087b1a65d06d17aab1ca9d69170b292a4e0d90063ab.NewPrivilegedApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedOperationEvents provides operations to manage the collection of privilegedOperationEvent entities.
// returns a *PrivilegedOperationEventsRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedOperationEvents()(*i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.PrivilegedOperationEventsRequestBuilder) {
    return i82672a497b8f66c510d59f2a80cb96da07ffad912f0c452f2d22f4a282e37720.NewPrivilegedOperationEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedRoleAssignmentRequests provides operations to manage the collection of privilegedRoleAssignmentRequest entities.
// returns a *PrivilegedRoleAssignmentRequestsRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedRoleAssignmentRequests()(*if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.PrivilegedRoleAssignmentRequestsRequestBuilder) {
    return if11203f2b5e6319285361e1998b4a25572cc1950d617d10ffd84e91e5f477349.NewPrivilegedRoleAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedRoleAssignments provides operations to manage the collection of privilegedRoleAssignment entities.
// returns a *PrivilegedRoleAssignmentsRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedRoleAssignments()(*i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.PrivilegedRoleAssignmentsRequestBuilder) {
    return i85eafc30b6e6233aefbca2b4e35ced0ac8836418de9401407873193666d812d9.NewPrivilegedRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedRoles provides operations to manage the collection of privilegedRole entities.
// returns a *PrivilegedRolesRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedRoles()(*i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.PrivilegedRolesRequestBuilder) {
    return i32f7b810493504f82bd6c97020faab5a8ff5f541a46dd3d6b9cc2aa77fc22fe4.NewPrivilegedRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedSignupStatus provides operations to manage the collection of privilegedSignupStatus entities.
// returns a *PrivilegedSignupStatusRequestBuilder when successful
func (m *GraphBaseServiceClient) PrivilegedSignupStatus()(*i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.PrivilegedSignupStatusRequestBuilder) {
    return i6b96a96c52bbdff1731b8a5490cd5f342e33866e0931912944d323bc79f663e4.NewPrivilegedSignupStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProgramControls provides operations to manage the collection of programControl entities.
// returns a *ProgramControlsRequestBuilder when successful
func (m *GraphBaseServiceClient) ProgramControls()(*i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.ProgramControlsRequestBuilder) {
    return i178f0aa4e5987fcbfe2e98cbb6dd777ebcdcdf124dd3478d2bf40f83912ca030.NewProgramControlsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProgramControlTypes provides operations to manage the collection of programControlType entities.
// returns a *ProgramControlTypesRequestBuilder when successful
func (m *GraphBaseServiceClient) ProgramControlTypes()(*if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.ProgramControlTypesRequestBuilder) {
    return if5cd0cf36bc86d9253920d73c41189ad8a30342e678d4f0138afa5095fd31538.NewProgramControlTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Programs provides operations to manage the collection of program entities.
// returns a *ProgramsRequestBuilder when successful
func (m *GraphBaseServiceClient) Programs()(*ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.ProgramsRequestBuilder) {
    return ib85b32f0384596c14f04b8d0f3dc8737da4b97428d7af145db2f1b06d7d9444e.NewProgramsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reportRoot singleton.
// returns a *ReportsRequestBuilder when successful
func (m *GraphBaseServiceClient) Reports()(*ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.ReportsRequestBuilder) {
    return ia5f30c11d37332bb7dfb48303b8e8880e60fc3560f60a813c66b9c1c3f2ff3f2.NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RiskDetections provides operations to manage the collection of riskDetection entities.
// returns a *RiskDetectionsRequestBuilder when successful
func (m *GraphBaseServiceClient) RiskDetections()(*i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.RiskDetectionsRequestBuilder) {
    return i24b998459b0fbcdc6fbd83b90048a098ff6bbdcdaff773a2886f5cf8b3d5545e.NewRiskDetectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RiskyUsers provides operations to manage the collection of riskyUser entities.
// returns a *RiskyUsersRequestBuilder when successful
func (m *GraphBaseServiceClient) RiskyUsers()(*iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.RiskyUsersRequestBuilder) {
    return iac77c9b5b86109e8ad626e30830db719efb3cc77c7babab332b409d84ae324a6.NewRiskyUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleManagement provides operations to manage the roleManagement singleton.
// returns a *RoleManagementRequestBuilder when successful
func (m *GraphBaseServiceClient) RoleManagement()(*i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.RoleManagementRequestBuilder) {
    return i310cda3e9f244aa61f9c9c78de433773f341b91c4a2310b8991671fe773be16e.NewRoleManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SchemaExtensions provides operations to manage the collection of schemaExtension entities.
// returns a *SchemaExtensionsRequestBuilder when successful
func (m *GraphBaseServiceClient) SchemaExtensions()(*i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.SchemaExtensionsRequestBuilder) {
    return i1b84a2c37ba0bbd175c6da40c8679db7d04dfcb044d8421d26d024db45218e4a.NewSchemaExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScopedRoleMemberships provides operations to manage the collection of scopedRoleMembership entities.
// returns a *ScopedRoleMembershipsRequestBuilder when successful
func (m *GraphBaseServiceClient) ScopedRoleMemberships()(*i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.ScopedRoleMembershipsRequestBuilder) {
    return i712907ad27a66d6ac32a26e01f88de1ad6484585eb7ed65f84b3a30571cec55b.NewScopedRoleMembershipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search provides operations to manage the searchEntity singleton.
// returns a *SearchRequestBuilder when successful
func (m *GraphBaseServiceClient) Search()(*i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.SearchRequestBuilder) {
    return i5840582f75a8eb78900edf3bd78566223ffee7aa1dc2f4cdca943ef635f6503e.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Security provides operations to manage the security singleton.
// returns a *SecurityRequestBuilder when successful
func (m *GraphBaseServiceClient) Security()(*i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.SecurityRequestBuilder) {
    return i761e9f0dec20dbf36c7fd626d107fb81ef94cafa7369422d2b2af143ffa16184.NewSecurityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePrincipals provides operations to manage the collection of servicePrincipal entities.
// returns a *ServicePrincipalsRequestBuilder when successful
func (m *GraphBaseServiceClient) ServicePrincipals()(*ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.ServicePrincipalsRequestBuilder) {
    return ibd3e65bb14e91a8a05d902c54fadec2c1b6931676c97f76da4969c975770aab2.NewServicePrincipalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePrincipalsWithAppId provides operations to manage the collection of servicePrincipal entities.
// returns a *ServicePrincipalsWithAppIdRequestBuilder when successful
func (m *GraphBaseServiceClient) ServicePrincipalsWithAppId(appId *string)(*if298592eede46286a454444191c80a3a678c2136e14beb3f7c77b80b0b3ed812.ServicePrincipalsWithAppIdRequestBuilder) {
    return if298592eede46286a454444191c80a3a678c2136e14beb3f7c77b80b0b3ed812.NewServicePrincipalsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// Settings provides operations to manage the collection of directorySetting entities.
// returns a *SettingsRequestBuilder when successful
func (m *GraphBaseServiceClient) Settings()(*i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.SettingsRequestBuilder) {
    return i714cbeb65962cb4d3e58007792fa4832d175c04614ba3aa7efb22871aea885bf.NewSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shares provides operations to manage the collection of sharedDriveItem entities.
// returns a *SharesRequestBuilder when successful
func (m *GraphBaseServiceClient) Shares()(*i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.SharesRequestBuilder) {
    return i73583652789c7aab226ac5bae66bc7b5fd924607d28350c4478c2a20524fd624.NewSharesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the collection of site entities.
// returns a *SitesRequestBuilder when successful
func (m *GraphBaseServiceClient) Sites()(*i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.SitesRequestBuilder) {
    return i2817c6849b286be20c56215e039110b08a4109a12669579f4597fbab6f720ed9.NewSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Solutions provides operations to manage the solutionsRoot singleton.
// returns a *SolutionsRequestBuilder when successful
func (m *GraphBaseServiceClient) Solutions()(*i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.SolutionsRequestBuilder) {
    return i7282246c788c7b44fcd028d8a1500ce1ffd32f655bee3b580b6a33d9ec10f610.NewSolutionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Storage provides operations to manage the storage singleton.
// returns a *StorageRequestBuilder when successful
func (m *GraphBaseServiceClient) Storage()(*id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.StorageRequestBuilder) {
    return id14bafb4ee71e1257662cacab67dd479e54eba65ce10c6816ee4676804caf821.NewStorageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubscribedSkus provides operations to manage the collection of subscribedSku entities.
// returns a *SubscribedSkusRequestBuilder when successful
func (m *GraphBaseServiceClient) SubscribedSkus()(*ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.SubscribedSkusRequestBuilder) {
    return ie934faa615fb56652e5964395b3dc205321ac84e8cf244796ebe59ba3713fbd9.NewSubscribedSkusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscriptions provides operations to manage the collection of subscription entities.
// returns a *SubscriptionsRequestBuilder when successful
func (m *GraphBaseServiceClient) Subscriptions()(*i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.SubscriptionsRequestBuilder) {
    return i60392bb6eb86abe7a3079e3b2b1e202f7dcf3452adf026db62ec93e2fafe3957.NewSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams provides operations to manage the collection of team entities.
// returns a *TeamsRequestBuilder when successful
func (m *GraphBaseServiceClient) Teams()(*iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.TeamsRequestBuilder) {
    return iff395ba1da21566390b02b5bed781aecf3bb849fc71f2359410792d1d1b67079.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TeamsTemplates provides operations to manage the collection of teamsTemplate entities.
// returns a *TeamsTemplatesRequestBuilder when successful
func (m *GraphBaseServiceClient) TeamsTemplates()(*i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.TeamsTemplatesRequestBuilder) {
    return i66f18ccab4e34309d26d1056f0e7dd8b563a5f8ee6f8d9c6e8e77c5fac50f8b5.NewTeamsTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TeamTemplateDefinition provides operations to manage the collection of teamTemplateDefinition entities.
// returns a *TeamTemplateDefinitionRequestBuilder when successful
func (m *GraphBaseServiceClient) TeamTemplateDefinition()(*i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.TeamTemplateDefinitionRequestBuilder) {
    return i1c7e7a5d0708841f8c98ec910d583f348cbffaad386ef9a24d3ee4eba285ea21.NewTeamTemplateDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teamwork provides operations to manage the teamwork singleton.
// returns a *TeamworkRequestBuilder when successful
func (m *GraphBaseServiceClient) Teamwork()(*icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.TeamworkRequestBuilder) {
    return icb4f253cb1cd35435f5752b611229032c618bbcfeb3be80ee4d6a06d404114fc.NewTeamworkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantRelationships provides operations to manage the tenantRelationship singleton.
// returns a *TenantRelationshipsRequestBuilder when successful
func (m *GraphBaseServiceClient) TenantRelationships()(*id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.TenantRelationshipsRequestBuilder) {
    return id5c2ef977a00dd1757d258dbbbfb4080031771e62e6c6b3b1339a0f03fc1c1f1.NewTenantRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the store singleton.
// returns a *TermStoreRequestBuilder when successful
func (m *GraphBaseServiceClient) TermStore()(*i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.TermStoreRequestBuilder) {
    return i98a1471d41b15330865bc87691830281af9ecf479bfc797e54f02448790b1e4e.NewTermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatSubmission provides operations to manage the threatSubmissionRoot singleton.
// returns a *ThreatSubmissionRequestBuilder when successful
func (m *GraphBaseServiceClient) ThreatSubmission()(*i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.ThreatSubmissionRequestBuilder) {
    return i231953b44a7f4aace0800ccf375537d423e6f60e82f2a6552d9613626e39aba5.NewThreatSubmissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TrustFramework provides operations to manage the trustFramework singleton.
// returns a *TrustFrameworkRequestBuilder when successful
func (m *GraphBaseServiceClient) TrustFramework()(*i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.TrustFrameworkRequestBuilder) {
    return i312c0a09d8ded5436957205a14adfc7e2facbcc6f26ef9872a5b5eb79228375f.NewTrustFrameworkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users provides operations to manage the collection of user entities.
// returns a *UsersRequestBuilder when successful
func (m *GraphBaseServiceClient) Users()(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UsersRequestBuilder) {
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Workplace provides operations to manage the workplace singleton.
// returns a *WorkplaceRequestBuilder when successful
func (m *GraphBaseServiceClient) Workplace()(*i164b4c6703708dd1b0670a17a07a4dd64e49cb0c1cb66e50a3146217db57a57a.WorkplaceRequestBuilder) {
    return i164b4c6703708dd1b0670a17a07a4dd64e49cb0c1cb66e50a3146217db57a57a.NewWorkplaceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
