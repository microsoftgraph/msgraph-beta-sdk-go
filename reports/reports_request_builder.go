package reports

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetime"
    i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivityuserdetailwithperiod"
    i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appusercountswithperiod"
    i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmenttopfailureswithperiod"
    i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activationcounts"
    i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitycountswithperiod"
    i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagesitecountswithperiod"
    i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activeuserdetailwithperiod"
    i11672b5de3bc9d06bd8a600a8e1330f9218b2d50afdd893bb96da285d6b3eea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsteamactivitydetailwithdate"
    i14dceb42fdafcb299797921a57bd593b62efdccf10e177a8f72ba7b11efd7b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsteamactivitycountswithperiod"
    i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusageuserdetailwithdate"
    i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagepageswithperiod"
    i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessorganizeractivityminutecountswithperiod"
    i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivityusercountswithperiod"
    i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagedetailwithperiod"
    i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitystoragewithperiod"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i23ae3dadc25e31a369068a3ac3cd7ff57c8b4823dcea9650bd4b312d6f57e7e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammeractivityusercountswithperiod"
    i24783836c67f3cbffe50d47dcbcb2588e9de66ffd8f9a5c8da53fa898e5e9446 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmentabandonmentdetailswithskipwithtopwithfilterwithskiptoken"
    i250edbdb590c4bab056109a11a6cf515b7193da960e757e0325976fb00e6b7ef "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getrelyingpartydetailedsummarywithperiod"
    i26bbf530ddb18f9b5655ff8fa87d5f3b5ad1da8c08a1db360b8de025d0362210 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitydistributionusercountswithperiod"
    i27595bc84745a6124aef7bab74a2b609c8775cdb2e31edd5259685331293e617 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmentabandonmentsummarywithskipwithtopwithfilterwithskiptoken"
    i27a2b61ad92d96f585909f7842612603cbacfc26e12dc059d2f4878490c04101 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessorganizeractivityusercountswithperiod"
    i2993fd51754851864ba7c1a8dfc4c807b19f8200be6ce2e4964d2b687966d694 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyuser"
    i2bb1d1186591127bed84408c868b7d735b4f2f8bbd3d5294b967b9b6ae5594b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getgrouparchivedprintjobswithgroupidwithstartdatetimewithenddatetime"
    i2ca5100bc1b71c10f318988b462b80ebf432370aae38b64473c17dd121558bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusagetotalusercountswithperiod"
    i2dc572cfc357c8939c4ee58887229a9bc49bdbe95b636fb6d5460ee9894192d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageversionsusercountswithperiod"
    i3573cf7b32242b3add3fa97f751c5d986f3543d84edb73bd531d941dccc6d8c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getazureadapplicationsigninsummarywithperiod"
    i35fab8ead83b7a35b3fc3dff29cb309b459f147de1727e7671b1e3ea1085f7a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessactivityusercountswithperiod"
    i37eb861166f6d0555df75cb64e973d22d9bb9c2a8c03ca2f6c29444c37e3649f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyprinter"
    i38018b8ba711a07a4df4ba04370314b5a7eb52d04ded8775a2b76a35218402ef "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusagedistributionusercountswithperiod"
    i38d9dde997bfdced4eba6828af9a431514b84509f8ed82648d159166a512a10b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinesspeertopeeractivityminutecountswithperiod"
    i3c5eb35d4514961f4cbfafd47ac65dc5ed997312a3e247f4b53d8004051d0194 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyprinter"
    i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/applicationsignindetailedsummary"
    i434cde18544a55e4247ac040b6fccc58d1ed919655f86445cff7c94dcfda1437 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getattacksimulationsimulationusercoverage"
    i48a44087adc49827c59b599bc3e9c1ea8e284030a4416856821ea71b09310ed5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveusageaccountdetailwithdate"
    i492793d896d7d799d0ff9804f56092f81d113794a8916673f9c324f5a2053ffc "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmenttopfailures"
    i4a335c23d6834f12971f5a81936e763c8d804a2f22f65b093947ae6a7f4a1634 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessparticipantactivityminutecountswithperiod"
    i4a720b5f345e91b0bf967bf525d76d5febe85151037aa8327de70f7fd0aea25a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageuserdetailwithdate"
    i4a8fe9122551f2771b4a12b8839c70ffec863bf5960e9fcfc031a4e55d7adab6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessparticipantactivityusercountswithperiod"
    i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/credentialuserregistrationdetails"
    i4c2e0be102a6b6d64d75b2ce3aed7eee80548af7512aca09bc765525a936d621 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveusagestoragewithperiod"
    i4c439992e985ac97d20ac69debd6694765696eb8acf494c0dd9c1e843dde5943 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessactivitycountswithperiod"
    i4d0f15879165ef15ce895cf317adfa50fa15bf6ed244a753762d508f5d5ba5e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagefilecountswithperiod"
    i4f16df07e7ec9e6dd5cb52fe46110ce00f5f74418530157eaee04c72ffe1c4e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammeractivityuserdetailwithperiod"
    i50d831cff0d03367e9939eb5f21f0a71e577dce8e1e220d3877edf21418cf620 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getuserarchivedprintjobswithuseridwithstartdatetimewithenddatetime"
    i51d8f12ca8bef3b5a605ad76c9d75dc8e73de7c0bb0b203ff61c881d868d0391 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyprinter"
    i53891a3a3983e3f57fac19166c9af13eee82757751926f31b54f16d6fa23d29a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveactivityfilecountswithperiod"
    i5542a4e920cc301f9fd5b6d010166452a3ec1b96b595f0b1f8869da5b984d4ff "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitydetailwithdate"
    i557e886461aa53bc4de7819875552204d4fe40dab2a518db1a29f2061c23a669 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivitycountswithperiod"
    i59460f776b2ba7ec0487cd235a89ad5d26f27c3ec165f0790954cfc2a01b9637 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyprinter"
    i5abd3f1e1b1701310329336b2d99ca6c43f6556a82d34472ce1d9c428c9d9c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitygroupcountswithperiod"
    i5d36536b2d00cade7bb67be5034066eaa35f6bcc8eb10887fab5db989e0190b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getazureadlicenseusagewithperiod"
    i6419b1c71ac2474e588fa2df68b31ad5cf99896a1c878c0b2f23b8a51f4794c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessdeviceusagedistributionusercountswithperiod"
    i6735cb8d512199566e336179e484d86cb97201979137f1ef2ce9fdd8486c616e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammeractivitycountswithperiod"
    i6858d47e90b1334353133cd0114055d1772fbe2294c3c8b1f87448025955a7f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyuser"
    i6ad597f06da6b987f26da44b061b096836613ef8b752128cb12644670a0d63a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagequotastatusmailboxcountswithperiod"
    i6e57d0d859f4603fe4b6c70f2f421bec46c3044cb19b132e453e8c658e9e5b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/deviceconfigurationdeviceactivity"
    i700648709215c245e3f56a9bb592458abcaa4d198b5956df2836a08af3f74be4 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyuser"
    i7026ee7cc69ab0c0a76cfd363e6e41a9e764f78dfc9fc6b226788e43010b3585 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessdeviceusageuserdetailwithdate"
    i7143ac54c5e4852fa7c71022277b025f4e1140c3bb0b90e520b5505dfa09cc23 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagestoragewithperiod"
    i7217ecce10bb1170c5d1635128247d30624546692ebc20935f655e3ce4b68c47 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivityusercountswithperiod"
    i791f698d83fc125e8a1806cd3e8a916e376ace94be4518e908c757250c9c64e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmentfailuretrends"
    i7ad76340f4a0f8d49588345643b23f3a29ec1fd2161c87b7866114a35eda410a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveusageaccountcountswithperiod"
    i7c9189a010c6688a7fa9746221903390178e46db9ee2896c06cd2e3e781e3d3a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitydetailwithperiod"
    i8183070c6472c33737f601f9c6bf08439b4a91eed6e7b421036eecdd6e7c5cea "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getbrowseruserdetailwithperiod"
    i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityusercountswithperiod"
    i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitycountswithperiod"
    i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessactivityuserdetailwithperiod"
    i873f45675c7625cee9b526eb5a3af6e18edcd9f05511b0a172caeb817796d293 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsteamactivitydistributioncountswithperiod"
    i8a14754850e94320f257277f7b5983b2dfdea28ccf42833202bf6ecd26e1e9b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmentfailuredetails"
    i8deb95de65d724b352c9c68714d954885c65ccac1053a18e56072a07014b8d5b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusageuserdetailwithperiod"
    i8e6535187547a4b5d5c14160716a4c020758272be3e5ca22edba3b46b3087b52 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitytotalusercountswithperiod"
    i91a39045a7be5817859a873d2638b0c156baab984d6270f68a448faf3297b18c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveusagefilecountswithperiod"
    i920ca9bdad0e6d6c258a83d691a206e4dcbd817fcf55f61727be543b8f6a51e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageappsusercountswithperiod"
    i93a80e4540ddcb67a9776e160a250432a3c4dc682adbf3a39d6e2f9904bb9447 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagedetailwithperiod"
    i943fdce89c1d1ccca5b0f3b2db4cfc85447eafff319d6137111ed073b10c023e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusageusercountswithperiod"
    i944bfcd02a17a424e910488ba7b947a5fdad6f5a75edb524a1d32a729f796746 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivitypageswithperiod"
    i94b276da4083eda748c10c94dfe16c6d4742c5a65341271d9d71d7381dae0214 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getattacksimulationrepeatoffenders"
    i95b1164d20aad2d39c57c4b38fd44a5d2f2d05bf631cb13a075f37b71a7ac5b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appuserdetailwithdate"
    i973d4295c78aa9f58867f3b650c440eb7e1901bb527ccd20a8687328d4ecbbc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityfilecountswithperiod"
    i98b74c2f146a48ff844d3021737827fcc82215b40e873605b5b930c294e36b6b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusageusercountswithperiod"
    i9a872942082503636d059c4c1175930485a0f27ecbf3b0043154dc6658ef476f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusageuserdetailwithperiod"
    i9c1f3873c5c6f05ebb01275ba1e1cadbcac87226b44c1b890d5571cd53f331ef "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessorganizeractivitycountswithperiod"
    i9f1ae835c41e1a61632007ea0b4fd227f894397673ddb731c0d36bd49470d5c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365servicesusercountswithperiod"
    i9fe7970817241d27b4418d90bdbf078d909f78d7898d3a41062551253bdc2ef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getbrowserdistributionusercountswithperiod"
    ia0822ab92389697fe0392482a489e7576f172e2b499906e62bb3393cc7c11dbb "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitytotaldistributioncountswithperiod"
    ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getattacksimulationtrainingusercoverage"
    ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusagedistributionusercountswithperiod"
    ia905ae63e07e238ab152ce311e0525628c071873fd34232733350f04752d8678 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagedetailwithdate"
    ia983902699c45d63b5babc13320daacdb1564e3d45458d3d6153bc2112fd2eb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activationsusercounts"
    iab64d32b308e755b9724f53342c4489cd12c811324ffcfc1eebaeb1f6d63937f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinesspeertopeeractivitycountswithperiod"
    iab808a61daaca7b2e378fc8a0289994ae7ddf284a49112e4b8b742d6f3cba446 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinesspeertopeeractivityusercountswithperiod"
    iac59a09de72fb2fd2f0077d123253c2ef987a511fbd554c000b1e3714111f572 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appuserdetailwithperiod"
    iacb9bbe21799f4942d290c5e1ab7f9050f7193bb0196f318f8707bee63c0ac5d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagestoragewithperiod"
    ib217a9ab86db27e8b3dce55808a82185fd090bc6c4245f5330f5c1b0c08a3a61 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmentfailuredetailswithskipwithtopwithfilterwithskiptoken"
    ib278d65ffeb85ab37242d833ae021cc70a1c1347099b6536ec196e9a7db3128c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/authenticationmethods"
    ib546a3980e2bf4e36d4a57260b208c8311e952ceecd05de93cafd1097b8221d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveusageaccountdetailwithperiod"
    ibaf7d6c4cd3e8f19279c8ea66be35293ecddff285916af2302dad085654fb7a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitydetailwithdate"
    ic201bfab351b7c4943cb69827ffc49f717f8d030c2c09e022994a10813ccebb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitytotalcountswithperiod"
    ic2bba236674645b7dd61f3f93d935e166db2806a283ce92ca066bb178547b2b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessactivityuserdetailwithdate"
    ic3f682e7cc9725074cc23ecc64d8d813b5f6c39a93bcdd47f5b04e1e36e31a8c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveactivityusercountswithperiod"
    ic5da14e3ff43c5e6fd44a093a284bd0d470c3a9d4b65d5e2bce503b45cd374fd "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveactivityuserdetailwithdate"
    ic64ccf8d946071ab046b2c31875bb4f39c0a87b879466e64adf3aea9c73aa206 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitydistributiontotalusercountswithperiod"
    ic80bad1ddef41832846ff078a1a25173a3d2273a6871e39258db976a0150fc39 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsteamactivitydetailwithperiod"
    ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivityuserdetailwithdate"
    ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammeractivityuserdetailwithdate"
    id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityuserdetailwithperiod"
    id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivityfilecountswithperiod"
    id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageusercountswithperiod"
    id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveactivityuserdetailwithperiod"
    id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusageuserdetailwithdate"
    id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagemailboxcountswithperiod"
    ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityuserdetailwithdate"
    ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appplatformusercountswithperiod"
    ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/usercredentialusagedetails"
    ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitygroupcountswithperiod"
    ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusagedistributiontotalusercountswithperiod"
    ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getcredentialuserregistrationcount"
    ie9e4990bb33919415a8bb0cc6caa1b524b14e68c3ab1656a344070085517a253 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getbrowserusercountswithperiod"
    ieab560554737ae1c313ebdb1132ed35ae880a6f013ece3d8e5dc8786eeb39894 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activationsuserdetail"
    ieda57b6fe3b5918581e568356ba76f653b708c99154876060bace124d82a1bad "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/deviceconfigurationuseractivity"
    iedafbd1ed73588865f6b43d238dd6c81c00b86c2c48d2ba9ecb1e727bf24afe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivityuserdetailwithdate"
    iedb2c5fd4b1885cfde5af8d13d7955fcc21c8d6faeb064c46aaad884983733d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitydetailwithperiod"
    iee0f7f8db82d7bf47fa5eff92350ac6235e7af4d5a35bf22f0c190f09903fd62 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivityuserdetailwithperiod"
    if08b0f6462c0374cc04d6591fd00c6f903d80cb1aa4bcf0269bf16987dcf17a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activeusercountswithperiod"
    if220cda8a33cfb0dc18bc5f4b4ed4f3787a1bb1f4392a8c67322dae2cf9910ce "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessdeviceusageusercountswithperiod"
    if239d7585794526a1ec1b695c30f15b3ba8659d065ac77ec5d3622719b1c1fc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessdeviceusageuserdetailwithperiod"
    if261b2578ae42789e9c954041da7ecdc03603a05edf6ba0be09893dbe05b3142 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activeuserdetailwithdate"
    if34dee9f0158cce73ba132047204086837bb3168fd06c8f7e138d3d5d60b773e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getazureadfeatureusagewithperiod"
    if587a78dac698fac07fde7dbafd659ec926b72e5abce7d5081eac704ddb5fc12 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitycountswithperiod"
    if958ce2337186d28b6ff688907b51b0baf53d270d71152d85463560542cb332f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessparticipantactivitycountswithperiod"
    ifacf533bfae54389c178233012b220f0cc75de6629790a7585f840fac461be0b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getcredentialusagesummarywithperiod"
    ifbd8842a6664ec96c7d893455943698937b1af3fe05558ca8e628721e162064b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyuser"
    ifd317ad9016a7541a0ad611c0d46387984fd5c9e4b06f7fce9718fc0b5c0e52f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageuserdetailwithperiod"
    iff7cc4e7a992aa4c240079c5a2255ed095f68878fb58b3ad449dc4ad2229a76c "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getazureaduserfeatureusage"
    i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyuser/item"
    i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyuser/item"
    i6e7800f96429fcd3abddac192866febcab6b2f74bee1f02470e717ea53e849df "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/usercredentialusagedetails/item"
    i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyprinter/item"
    i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyprinter/item"
    i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyprinter/item"
    ib41269879aaaaafe023559bee72eced006e67c0981359cdf0e53448af198f4bc "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/applicationsignindetailedsummary/item"
    id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyprinter/item"
    ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyuser/item"
    ie8dfbe0f2a10d928f186e8a85be8a433fa4f6ef763d02998cd4f197b915e90ce "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/credentialuserregistrationdetails/item"
    if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyuser/item"
)

// ReportsRequestBuilder provides operations to manage the reportRoot singleton.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderGetQueryParameters get reports
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationSignInDetailedSummary the applicationSignInDetailedSummary property
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.ApplicationSignInDetailedSummaryRequestBuilder) {
    return i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.applicationSignInDetailedSummary.item collection
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*ib41269879aaaaafe023559bee72eced006e67c0981359cdf0e53448af198f4bc.ApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return ib41269879aaaaafe023559bee72eced006e67c0981359cdf0e53448af198f4bc.NewApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethods the authenticationMethods property
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ib278d65ffeb85ab37242d833ae021cc70a1c1347099b6536ec196e9a7db3128c.AuthenticationMethodsRequestBuilder) {
    return ib278d65ffeb85ab37242d833ae021cc70a1c1347099b6536ec196e9a7db3128c.NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get reports
func (m *ReportsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get reports
func (m *ReportsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update reports
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update reports
func (m *ReportsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CredentialUserRegistrationDetails the credentialUserRegistrationDetails property
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.CredentialUserRegistrationDetailsRequestBuilder) {
    return i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.credentialUserRegistrationDetails.item collection
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*ie8dfbe0f2a10d928f186e8a85be8a433fa4f6ef763d02998cd4f197b915e90ce.CredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return ie8dfbe0f2a10d928f186e8a85be8a433fa4f6ef763d02998cd4f197b915e90ce.NewCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByPrinter the dailyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*i3c5eb35d4514961f4cbfafd47ac65dc5ed997312a3e247f4b53d8004051d0194.DailyPrintUsageByPrinterRequestBuilder) {
    return i3c5eb35d4514961f4cbfafd47ac65dc5ed997312a3e247f4b53d8004051d0194.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByUser the dailyPrintUsageByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i2993fd51754851864ba7c1a8dfc4c807b19f8200be6ce2e4964d2b687966d694.DailyPrintUsageByUserRequestBuilder) {
    return i2993fd51754851864ba7c1a8dfc4c807b19f8200be6ce2e4964d2b687966d694.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinter the dailyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*i51d8f12ca8bef3b5a605ad76c9d75dc8e73de7c0bb0b203ff61c881d868d0391.DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return i51d8f12ca8bef3b5a605ad76c9d75dc8e73de7c0bb0b203ff61c881d868d0391.NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByUser the dailyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*i6858d47e90b1334353133cd0114055d1772fbe2294c3c8b1f87448025955a7f9.DailyPrintUsageSummariesByUserRequestBuilder) {
    return i6858d47e90b1334353133cd0114055d1772fbe2294c3c8b1f87448025955a7f9.NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*i6e57d0d859f4603fe4b6c70f2f421bec46c3044cb19b132e453e8c658e9e5b9c.DeviceConfigurationDeviceActivityRequestBuilder) {
    return i6e57d0d859f4603fe4b6c70f2f421bec46c3044cb19b132e453e8c658e9e5b9c.NewDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationUserActivity()(*ieda57b6fe3b5918581e568356ba76f653b708c99154876060bace124d82a1bad.DeviceConfigurationUserActivityRequestBuilder) {
    return ieda57b6fe3b5918581e568356ba76f653b708c99154876060bace124d82a1bad.NewDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get reports
func (m *ReportsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) GetAttackSimulationRepeatOffenders()(*i94b276da4083eda748c10c94dfe16c6d4742c5a65341271d9d71d7381dae0214.GetAttackSimulationRepeatOffendersRequestBuilder) {
    return i94b276da4083eda748c10c94dfe16c6d4742c5a65341271d9d71d7381dae0214.NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationSimulationUserCoverage()(*i434cde18544a55e4247ac040b6fccc58d1ed919655f86445cff7c94dcfda1437.GetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return i434cde18544a55e4247ac040b6fccc58d1ed919655f86445cff7c94dcfda1437.NewGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationTrainingUserCoverage()(*ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d.GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d.NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) GetAzureADApplicationSignInSummaryWithPeriod(period *string)(*i3573cf7b32242b3add3fa97f751c5d986f3543d84edb73bd531d941dccc6d8c8.GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return i3573cf7b32242b3add3fa97f751c5d986f3543d84edb73bd531d941dccc6d8c8.NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADFeatureUsageWithPeriod provides operations to call the getAzureADFeatureUsage method.
func (m *ReportsRequestBuilder) GetAzureADFeatureUsageWithPeriod(period *string)(*if34dee9f0158cce73ba132047204086837bb3168fd06c8f7e138d3d5d60b773e.GetAzureADFeatureUsageWithPeriodRequestBuilder) {
    return if34dee9f0158cce73ba132047204086837bb3168fd06c8f7e138d3d5d60b773e.NewGetAzureADFeatureUsageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADLicenseUsageWithPeriod provides operations to call the getAzureADLicenseUsage method.
func (m *ReportsRequestBuilder) GetAzureADLicenseUsageWithPeriod(period *string)(*i5d36536b2d00cade7bb67be5034066eaa35f6bcc8eb10887fab5db989e0190b7.GetAzureADLicenseUsageWithPeriodRequestBuilder) {
    return i5d36536b2d00cade7bb67be5034066eaa35f6bcc8eb10887fab5db989e0190b7.NewGetAzureADLicenseUsageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADUserFeatureUsage provides operations to call the getAzureADUserFeatureUsage method.
func (m *ReportsRequestBuilder) GetAzureADUserFeatureUsage()(*iff7cc4e7a992aa4c240079c5a2255ed095f68878fb58b3ad449dc4ad2229a76c.GetAzureADUserFeatureUsageRequestBuilder) {
    return iff7cc4e7a992aa4c240079c5a2255ed095f68878fb58b3ad449dc4ad2229a76c.NewGetAzureADUserFeatureUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserDistributionUserCountsWithPeriod(period *string)(*i9fe7970817241d27b4418d90bdbf078d909f78d7898d3a41062551253bdc2ef9.GetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return i9fe7970817241d27b4418d90bdbf078d909f78d7898d3a41062551253bdc2ef9.NewGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserUserCountsWithPeriod(period *string)(*ie9e4990bb33919415a8bb0cc6caa1b524b14e68c3ab1656a344070085517a253.GetBrowserUserCountsWithPeriodRequestBuilder) {
    return ie9e4990bb33919415a8bb0cc6caa1b524b14e68c3ab1656a344070085517a253.NewGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) GetBrowserUserDetailWithPeriod(period *string)(*i8183070c6472c33737f601f9c6bf08439b4a91eed6e7b421036eecdd6e7c5cea.GetBrowserUserDetailWithPeriodRequestBuilder) {
    return i8183070c6472c33737f601f9c6bf08439b4a91eed6e7b421036eecdd6e7c5cea.NewGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) GetCredentialUsageSummaryWithPeriod(period *string)(*ifacf533bfae54389c178233012b220f0cc75de6629790a7585f840fac461be0b.GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return ifacf533bfae54389c178233012b220f0cc75de6629790a7585f840fac461be0b.NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) GetCredentialUserRegistrationCount()(*ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0.GetCredentialUserRegistrationCountRequestBuilder) {
    return ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0.NewGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*i557e886461aa53bc4de7819875552204d4fe40dab2a518db1a29f2061c23a669.GetEmailActivityCountsWithPeriodRequestBuilder) {
    return i557e886461aa53bc4de7819875552204d4fe40dab2a518db1a29f2061c23a669.NewGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*i7217ecce10bb1170c5d1635128247d30624546692ebc20935f655e3ce4b68c47.GetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return i7217ecce10bb1170c5d1635128247d30624546692ebc20935f655e3ce4b68c47.NewGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd.GetEmailActivityUserDetailWithDateRequestBuilder) {
    return ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd.NewGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e.GetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e.NewGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*i920ca9bdad0e6d6c258a83d691a206e4dcbd817fcf55f61727be543b8f6a51e1.GetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return i920ca9bdad0e6d6c258a83d691a206e4dcbd817fcf55f61727be543b8f6a51e1.NewGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8.GetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8.NewGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i4a720b5f345e91b0bf967bf525d76d5febe85151037aa8327de70f7fd0aea25a.GetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return i4a720b5f345e91b0bf967bf525d76d5febe85151037aa8327de70f7fd0aea25a.NewGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*ifd317ad9016a7541a0ad611c0d46387984fd5c9e4b06f7fce9718fc0b5c0e52f.GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return ifd317ad9016a7541a0ad611c0d46387984fd5c9e4b06f7fce9718fc0b5c0e52f.NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*i2dc572cfc357c8939c4ee58887229a9bc49bdbe95b636fb6d5460ee9894192d6.GetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return i2dc572cfc357c8939c4ee58887229a9bc49bdbe95b636fb6d5460ee9894192d6.NewGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i2bb1d1186591127bed84408c868b7d735b4f2f8bbd3d5294b967b9b6ae5594b9.GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i2bb1d1186591127bed84408c868b7d735b4f2f8bbd3d5294b967b9b6ae5594b9.NewGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime);
}
// GetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppPlatformUserCountsWithPeriod(period *string)(*ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e.GetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e.NewGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppUserCountsWithPeriod(period *string)(*i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84.GetM365AppUserCountsWithPeriodRequestBuilder) {
    return i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84.NewGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i95b1164d20aad2d39c57c4b38fd44a5d2f2d05bf631cb13a075f37b71a7ac5b4.GetM365AppUserDetailWithDateRequestBuilder) {
    return i95b1164d20aad2d39c57c4b38fd44a5d2f2d05bf631cb13a075f37b71a7ac5b4.NewGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithPeriod(period *string)(*iac59a09de72fb2fd2f0077d123253c2ef987a511fbd554c000b1e3714111f572.GetM365AppUserDetailWithPeriodRequestBuilder) {
    return iac59a09de72fb2fd2f0077d123253c2ef987a511fbd554c000b1e3714111f572.NewGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34.GetMailboxUsageDetailWithPeriodRequestBuilder) {
    return i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34.NewGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230.GetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230.NewGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*i6ad597f06da6b987f26da44b061b096836613ef8b752128cb12644670a0d63a7.GetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return i6ad597f06da6b987f26da44b061b096836613ef8b752128cb12644670a0d63a7.NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*iacb9bbe21799f4942d290c5e1ab7f9050f7193bb0196f318f8707bee63c0ac5d.GetMailboxUsageStorageWithPeriodRequestBuilder) {
    return iacb9bbe21799f4942d290c5e1ab7f9050f7193bb0196f318f8707bee63c0ac5d.NewGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationCounts()(*i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a.GetOffice365ActivationCountsRequestBuilder) {
    return i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a.NewGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*ia983902699c45d63b5babc13320daacdb1564e3d45458d3d6153bc2112fd2eb6.GetOffice365ActivationsUserCountsRequestBuilder) {
    return ia983902699c45d63b5babc13320daacdb1564e3d45458d3d6153bc2112fd2eb6.NewGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*ieab560554737ae1c313ebdb1132ed35ae880a6f013ece3d8e5dc8786eeb39894.GetOffice365ActivationsUserDetailRequestBuilder) {
    return ieab560554737ae1c313ebdb1132ed35ae880a6f013ece3d8e5dc8786eeb39894.NewGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*if08b0f6462c0374cc04d6591fd00c6f903d80cb1aa4bcf0269bf16987dcf17a2.GetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return if08b0f6462c0374cc04d6591fd00c6f903d80cb1aa4bcf0269bf16987dcf17a2.NewGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*if261b2578ae42789e9c954041da7ecdc03603a05edf6ba0be09893dbe05b3142.GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return if261b2578ae42789e9c954041da7ecdc03603a05edf6ba0be09893dbe05b3142.NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d.GetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d.NewGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*if587a78dac698fac07fde7dbafd659ec926b72e5abce7d5081eac704ddb5fc12.GetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return if587a78dac698fac07fde7dbafd659ec926b72e5abce7d5081eac704ddb5fc12.NewGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i5542a4e920cc301f9fd5b6d010166452a3ec1b96b595f0b1f8869da5b984d4ff.GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return i5542a4e920cc301f9fd5b6d010166452a3ec1b96b595f0b1f8869da5b984d4ff.NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*iedb2c5fd4b1885cfde5af8d13d7955fcc21c8d6faeb064c46aaad884983733d9.GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return iedb2c5fd4b1885cfde5af8d13d7955fcc21c8d6faeb064c46aaad884983733d9.NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6.GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6.NewGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*i5abd3f1e1b1701310329336b2d99ca6c43f6556a82d34472ce1d9c428c9d9c9c.GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return i5abd3f1e1b1701310329336b2d99ca6c43f6556a82d34472ce1d9c428c9d9c9c.NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab.GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab.NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*i9f1ae835c41e1a61632007ea0b4fd227f894397673ddb731c0d36bd49470d5c3.GetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return i9f1ae835c41e1a61632007ea0b4fd227f894397673ddb731c0d36bd49470d5c3.NewGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*i53891a3a3983e3f57fac19166c9af13eee82757751926f31b54f16d6fa23d29a.GetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return i53891a3a3983e3f57fac19166c9af13eee82757751926f31b54f16d6fa23d29a.NewGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*ic3f682e7cc9725074cc23ecc64d8d813b5f6c39a93bcdd47f5b04e1e36e31a8c.GetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return ic3f682e7cc9725074cc23ecc64d8d813b5f6c39a93bcdd47f5b04e1e36e31a8c.NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ic5da14e3ff43c5e6fd44a093a284bd0d470c3a9d4b65d5e2bce503b45cd374fd.GetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return ic5da14e3ff43c5e6fd44a093a284bd0d470c3a9d4b65d5e2bce503b45cd374fd.NewGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9.GetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9.NewGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*i7ad76340f4a0f8d49588345643b23f3a29ec1fd2161c87b7866114a35eda410a.GetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return i7ad76340f4a0f8d49588345643b23f3a29ec1fd2161c87b7866114a35eda410a.NewGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i48a44087adc49827c59b599bc3e9c1ea8e284030a4416856821ea71b09310ed5.GetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return i48a44087adc49827c59b599bc3e9c1ea8e284030a4416856821ea71b09310ed5.NewGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*ib546a3980e2bf4e36d4a57260b208c8311e952ceecd05de93cafd1097b8221d0.GetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return ib546a3980e2bf4e36d4a57260b208c8311e952ceecd05de93cafd1097b8221d0.NewGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*i91a39045a7be5817859a873d2638b0c156baab984d6270f68a448faf3297b18c.GetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return i91a39045a7be5817859a873d2638b0c156baab984d6270f68a448faf3297b18c.NewGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*i4c2e0be102a6b6d64d75b2ce3aed7eee80548af7512aca09bc765525a936d621.GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return i4c2e0be102a6b6d64d75b2ce3aed7eee80548af7512aca09bc765525a936d621.NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e.GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e.NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime);
}
// GetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) GetRelyingPartyDetailedSummaryWithPeriod(period *string)(*i250edbdb590c4bab056109a11a6cf515b7193da960e757e0325976fb00e6b7ef.GetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return i250edbdb590c4bab056109a11a6cf515b7193da960e757e0325976fb00e6b7ef.NewGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*i973d4295c78aa9f58867f3b650c440eb7e1901bb527ccd20a8687328d4ecbbc5.GetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return i973d4295c78aa9f58867f3b650c440eb7e1901bb527ccd20a8687328d4ecbbc5.NewGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*i944bfcd02a17a424e910488ba7b947a5fdad6f5a75edb524a1d32a729f796746.GetSharePointActivityPagesWithPeriodRequestBuilder) {
    return i944bfcd02a17a424e910488ba7b947a5fdad6f5a75edb524a1d32a729f796746.NewGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927.GetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927.NewGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512.GetSharePointActivityUserDetailWithDateRequestBuilder) {
    return ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512.NewGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420.GetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420.NewGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ia905ae63e07e238ab152ce311e0525628c071873fd34232733350f04752d8678.GetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return ia905ae63e07e238ab152ce311e0525628c071873fd34232733350f04752d8678.NewGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*i93a80e4540ddcb67a9776e160a250432a3c4dc682adbf3a39d6e2f9904bb9447.GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return i93a80e4540ddcb67a9776e160a250432a3c4dc682adbf3a39d6e2f9904bb9447.NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*i4d0f15879165ef15ce895cf317adfa50fa15bf6ed244a753762d508f5d5ba5e3.GetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return i4d0f15879165ef15ce895cf317adfa50fa15bf6ed244a753762d508f5d5ba5e3.NewGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367.GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367.NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d.GetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d.NewGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*i7143ac54c5e4852fa7c71022277b025f4e1140c3bb0b90e520b5505dfa09cc23.GetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return i7143ac54c5e4852fa7c71022277b025f4e1140c3bb0b90e520b5505dfa09cc23.NewGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*i4c439992e985ac97d20ac69debd6694765696eb8acf494c0dd9c1e843dde5943.GetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return i4c439992e985ac97d20ac69debd6694765696eb8acf494c0dd9c1e843dde5943.NewGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*i35fab8ead83b7a35b3fc3dff29cb309b459f147de1727e7671b1e3ea1085f7a2.GetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return i35fab8ead83b7a35b3fc3dff29cb309b459f147de1727e7671b1e3ea1085f7a2.NewGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ic2bba236674645b7dd61f3f93d935e166db2806a283ce92ca066bb178547b2b8.GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return ic2bba236674645b7dd61f3f93d935e166db2806a283ce92ca066bb178547b2b8.NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5.GetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5.NewGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i6419b1c71ac2474e588fa2df68b31ad5cf99896a1c878c0b2f23b8a51f4794c0.GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i6419b1c71ac2474e588fa2df68b31ad5cf99896a1c878c0b2f23b8a51f4794c0.NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*if220cda8a33cfb0dc18bc5f4b4ed4f3787a1bb1f4392a8c67322dae2cf9910ce.GetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return if220cda8a33cfb0dc18bc5f4b4ed4f3787a1bb1f4392a8c67322dae2cf9910ce.NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i7026ee7cc69ab0c0a76cfd363e6e41a9e764f78dfc9fc6b226788e43010b3585.GetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return i7026ee7cc69ab0c0a76cfd363e6e41a9e764f78dfc9fc6b226788e43010b3585.NewGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*if239d7585794526a1ec1b695c30f15b3ba8659d065ac77ec5d3622719b1c1fc4.GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return if239d7585794526a1ec1b695c30f15b3ba8659d065ac77ec5d3622719b1c1fc4.NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*i9c1f3873c5c6f05ebb01275ba1e1cadbcac87226b44c1b890d5571cd53f331ef.GetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return i9c1f3873c5c6f05ebb01275ba1e1cadbcac87226b44c1b890d5571cd53f331ef.NewGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139.GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139.NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*i27a2b61ad92d96f585909f7842612603cbacfc26e12dc059d2f4878490c04101.GetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return i27a2b61ad92d96f585909f7842612603cbacfc26e12dc059d2f4878490c04101.NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*if958ce2337186d28b6ff688907b51b0baf53d270d71152d85463560542cb332f.GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return if958ce2337186d28b6ff688907b51b0baf53d270d71152d85463560542cb332f.NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*i4a335c23d6834f12971f5a81936e763c8d804a2f22f65b093947ae6a7f4a1634.GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return i4a335c23d6834f12971f5a81936e763c8d804a2f22f65b093947ae6a7f4a1634.NewGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*i4a8fe9122551f2771b4a12b8839c70ffec863bf5960e9fcfc031a4e55d7adab6.GetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return i4a8fe9122551f2771b4a12b8839c70ffec863bf5960e9fcfc031a4e55d7adab6.NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*iab64d32b308e755b9724f53342c4489cd12c811324ffcfc1eebaeb1f6d63937f.GetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return iab64d32b308e755b9724f53342c4489cd12c811324ffcfc1eebaeb1f6d63937f.NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*i38d9dde997bfdced4eba6828af9a431514b84509f8ed82648d159166a512a10b.GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i38d9dde997bfdced4eba6828af9a431514b84509f8ed82648d159166a512a10b.NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*iab808a61daaca7b2e378fc8a0289994ae7ddf284a49112e4b8b742d6f3cba446.GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return iab808a61daaca7b2e378fc8a0289994ae7ddf284a49112e4b8b742d6f3cba446.NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b.GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b.NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i38018b8ba711a07a4df4ba04370314b5a7eb52d04ded8775a2b76a35218402ef.GetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i38018b8ba711a07a4df4ba04370314b5a7eb52d04ded8775a2b76a35218402ef.NewGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*i2ca5100bc1b71c10f318988b462b80ebf432370aae38b64473c17dd121558bfd.GetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return i2ca5100bc1b71c10f318988b462b80ebf432370aae38b64473c17dd121558bfd.NewGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*i98b74c2f146a48ff844d3021737827fcc82215b40e873605b5b930c294e36b6b.GetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i98b74c2f146a48ff844d3021737827fcc82215b40e873605b5b930c294e36b6b.NewGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3.GetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3.NewGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*i8deb95de65d724b352c9c68714d954885c65ccac1053a18e56072a07014b8d5b.GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i8deb95de65d724b352c9c68714d954885c65ccac1053a18e56072a07014b8d5b.NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityCountsWithPeriod(period *string)(*i14dceb42fdafcb299797921a57bd593b62efdccf10e177a8f72ba7b11efd7b91.GetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return i14dceb42fdafcb299797921a57bd593b62efdccf10e177a8f72ba7b11efd7b91.NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i11672b5de3bc9d06bd8a600a8e1330f9218b2d50afdd893bb96da285d6b3eea5.GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return i11672b5de3bc9d06bd8a600a8e1330f9218b2d50afdd893bb96da285d6b3eea5.NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithPeriod(period *string)(*ic80bad1ddef41832846ff078a1a25173a3d2273a6871e39258db976a0150fc39.GetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return ic80bad1ddef41832846ff078a1a25173a3d2273a6871e39258db976a0150fc39.NewGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*i873f45675c7625cee9b526eb5a3af6e18edcd9f05511b0a172caeb817796d293.GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return i873f45675c7625cee9b526eb5a3af6e18edcd9f05511b0a172caeb817796d293.NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf.GetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf.NewGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*ic64ccf8d946071ab046b2c31875bb4f39c0a87b879466e64adf3aea9c73aa206.GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return ic64ccf8d946071ab046b2c31875bb4f39c0a87b879466e64adf3aea9c73aa206.NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*i26bbf530ddb18f9b5655ff8fa87d5f3b5ad1da8c08a1db360b8de025d0362210.GetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return i26bbf530ddb18f9b5655ff8fa87d5f3b5ad1da8c08a1db360b8de025d0362210.NewGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalCountsWithPeriod(period *string)(*ic201bfab351b7c4943cb69827ffc49f717f8d030c2c09e022994a10813ccebb6.GetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return ic201bfab351b7c4943cb69827ffc49f717f8d030c2c09e022994a10813ccebb6.NewGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*ia0822ab92389697fe0392482a489e7576f172e2b499906e62bb3393cc7c11dbb.GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return ia0822ab92389697fe0392482a489e7576f172e2b499906e62bb3393cc7c11dbb.NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*i8e6535187547a4b5d5c14160716a4c020758272be3e5ca22edba3b46b3087b52.GetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return i8e6535187547a4b5d5c14160716a4c020758272be3e5ca22edba3b46b3087b52.NewGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba.GetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba.NewGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*iedafbd1ed73588865f6b43d238dd6c81c00b86c2c48d2ba9ecb1e727bf24afe8.GetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return iedafbd1ed73588865f6b43d238dd6c81c00b86c2c48d2ba9ecb1e727bf24afe8.NewGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*iee0f7f8db82d7bf47fa5eff92350ac6235e7af4d5a35bf22f0c190f09903fd62.GetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return iee0f7f8db82d7bf47fa5eff92350ac6235e7af4d5a35bf22f0c190f09903fd62.NewGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*i50d831cff0d03367e9939eb5f21f0a71e577dce8e1e220d3877edf21418cf620.GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i50d831cff0d03367e9939eb5f21f0a71e577dce8e1e220d3877edf21418cf620.NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId);
}
// GetWithRequestConfigurationAndResponseHandler get reports
func (m *ReportsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ReportsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// GetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*i6735cb8d512199566e336179e484d86cb97201979137f1ef2ce9fdd8486c616e.GetYammerActivityCountsWithPeriodRequestBuilder) {
    return i6735cb8d512199566e336179e484d86cb97201979137f1ef2ce9fdd8486c616e.NewGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*i23ae3dadc25e31a369068a3ac3cd7ff57c8b4823dcea9650bd4b312d6f57e7e4.GetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return i23ae3dadc25e31a369068a3ac3cd7ff57c8b4823dcea9650bd4b312d6f57e7e4.NewGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b.GetYammerActivityUserDetailWithDateRequestBuilder) {
    return ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b.NewGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*i4f16df07e7ec9e6dd5cb52fe46110ce00f5f74418530157eaee04c72ffe1c4e1.GetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return i4f16df07e7ec9e6dd5cb52fe46110ce00f5f74418530157eaee04c72ffe1c4e1.NewGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42.GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42.NewGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*i943fdce89c1d1ccca5b0f3b2db4cfc85447eafff319d6137111ed073b10c023e.GetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i943fdce89c1d1ccca5b0f3b2db4cfc85447eafff319d6137111ed073b10c023e.NewGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45.GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45.NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*i9a872942082503636d059c4c1175930485a0f27ecbf3b0043154dc6658ef476f.GetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i9a872942082503636d059c4c1175930485a0f27ecbf3b0043154dc6658ef476f.NewGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa.GetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa.NewGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ibaf7d6c4cd3e8f19279c8ea66be35293ecddff285916af2302dad085654fb7a4.GetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return ibaf7d6c4cd3e8f19279c8ea66be35293ecddff285916af2302dad085654fb7a4.NewGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*i7c9189a010c6688a7fa9746221903390178e46db9ee2896c06cd2e3e781e3d3a.GetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return i7c9189a010c6688a7fa9746221903390178e46db9ee2896c06cd2e3e781e3d3a.NewGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2.GetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2.NewGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*i24783836c67f3cbffe50d47dcbcb2588e9de66ffd8f9a5c8da53fa898e5e9446.ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i24783836c67f3cbffe50d47dcbcb2588e9de66ffd8f9a5c8da53fa898e5e9446.NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*i27595bc84745a6124aef7bab74a2b609c8775cdb2e31edd5259685331293e617.ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i27595bc84745a6124aef7bab74a2b609c8775cdb2e31edd5259685331293e617.NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*i8a14754850e94320f257277f7b5983b2dfdea28ccf42833202bf6ecd26e1e9b5.ManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return i8a14754850e94320f257277f7b5983b2dfdea28ccf42833202bf6ecd26e1e9b5.NewManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ib217a9ab86db27e8b3dce55808a82185fd090bc6c4245f5330f5c1b0c08a3a61.ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return ib217a9ab86db27e8b3dce55808a82185fd090bc6c4245f5330f5c1b0c08a3a61.NewManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureTrends()(*i791f698d83fc125e8a1806cd3e8a916e376ace94be4518e908c757250c9c64e5.ManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return i791f698d83fc125e8a1806cd3e8a916e376ace94be4518e908c757250c9c64e5.NewManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*i492793d896d7d799d0ff9804f56092f81d113794a8916673f9c324f5a2053ffc.ManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return i492793d896d7d799d0ff9804f56092f81d113794a8916673f9c324f5a2053ffc.NewManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361.ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361.NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// MonthlyPrintUsageByPrinter the monthlyPrintUsageByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*i37eb861166f6d0555df75cb64e973d22d9bb9c2a8c03ca2f6c29444c37e3649f.MonthlyPrintUsageByPrinterRequestBuilder) {
    return i37eb861166f6d0555df75cb64e973d22d9bb9c2a8c03ca2f6c29444c37e3649f.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageByUser the monthlyPrintUsageByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*ifbd8842a6664ec96c7d893455943698937b1af3fe05558ca8e628721e162064b.MonthlyPrintUsageByUserRequestBuilder) {
    return ifbd8842a6664ec96c7d893455943698937b1af3fe05558ca8e628721e162064b.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinter the monthlyPrintUsageSummariesByPrinter property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*i59460f776b2ba7ec0487cd235a89ad5d26f27c3ec165f0790954cfc2a01b9637.MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return i59460f776b2ba7ec0487cd235a89ad5d26f27c3ec165f0790954cfc2a01b9637.NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUser the monthlyPrintUsageSummariesByUser property
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*i700648709215c245e3f56a9bb592458abcaa4d198b5956df2836a08af3f74be4.MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return i700648709215c245e3f56a9bb592458abcaa4d198b5956df2836a08af3f74be4.NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update reports
func (m *ReportsRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update reports
func (m *ReportsRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserCredentialUsageDetails the userCredentialUsageDetails property
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.UserCredentialUsageDetailsRequestBuilder) {
    return ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.userCredentialUsageDetails.item collection
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*i6e7800f96429fcd3abddac192866febcab6b2f74bee1f02470e717ea53e849df.UserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return i6e7800f96429fcd3abddac192866febcab6b2f74bee1f02470e717ea53e849df.NewUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
