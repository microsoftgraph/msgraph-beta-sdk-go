package reports

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetime"
    i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivityuserdetailwithperiod"
    i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appusercountswithperiod"
    i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/manageddeviceenrollmenttopfailureswithperiod"
    i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activationcounts"
    i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivitycountswithperiod"
    i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagesitecountswithperiod"
    i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365activeuserdetailwithperiod"
    i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusageuserdetailwithdate"
    i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointsiteusagepageswithperiod"
    i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessorganizeractivityminutecountswithperiod"
    i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsuseractivityusercountswithperiod"
    i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagedetailwithperiod"
    i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivitystoragewithperiod"
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
    i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityusercountswithperiod"
    i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitycountswithperiod"
    i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getskypeforbusinessactivityuserdetailwithperiod"
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
    ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getattacksimulationtrainingusercoverage"
    ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammerdeviceusagedistributionusercountswithperiod"
    ia83b8d65117b3e838e53c360304a89995009d10e47e5dce012a0e8c0406b67e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/gettenantsecurescoreswithperiod"
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
    ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailactivityuserdetailwithdate"
    ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammeractivityuserdetailwithdate"
    id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityuserdetailwithperiod"
    id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getoffice365groupsactivityfilecountswithperiod"
    id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getemailappusageusercountswithperiod"
    id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getonedriveactivityuserdetailwithperiod"
    id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusageuserdetailwithdate"
    id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getmailboxusagemailboxcountswithperiod"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getsharepointactivityuserdetailwithdate"
    ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getm365appplatformusercountswithperiod"
    ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/usercredentialusagedetails"
    ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getyammergroupsactivitygroupcountswithperiod"
    ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getteamsdeviceusagedistributiontotalusercountswithperiod"
    ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/getcredentialuserregistrationcount"
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
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyuser/item"
    i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyuser/item"
    i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyprinter/item"
    i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagebyprinter/item"
    i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyprinter/item"
    id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagebyprinter/item"
    ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/monthlyprintusagesummariesbyuser/item"
    if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/reports/dailyprintusagesummariesbyuser/item"
)

// ReportsRequestBuilder builds and executes requests for operations under \reports
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ReportsRequestBuilderGetOptions options for Get
type ReportsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReportsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReportsRequestBuilderGetQueryParameters get reports
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ReportsRequestBuilderPatchOptions options for Patch
type ReportsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.ApplicationSignInDetailedSummaryRequestBuilder) {
    return i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.applicationSignInDetailedSummary.item collection
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.ApplicationSignInDetailedSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary_id"] = id
    }
    return i3e5807f553ec696a6e36bac9cfe39f134eed4f44a99a041d12ec819c388bd840.NewApplicationSignInDetailedSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ib278d65ffeb85ab37242d833ae021cc70a1c1347099b6536ec196e9a7db3128c.AuthenticationMethodsRequestBuilder) {
    return ib278d65ffeb85ab37242d833ae021cc70a1c1347099b6536ec196e9a7db3128c.NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get reports
func (m *ReportsRequestBuilder) CreateGetRequestInformation(options *ReportsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update reports
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(options *ReportsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.CredentialUserRegistrationDetailsRequestBuilder) {
    return i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.credentialUserRegistrationDetails.item collection
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.CredentialUserRegistrationDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails_id"] = id
    }
    return i4b201efc99cccf9ab4d4acd90fb2df1aee6c33be6c22339b8c5424063e857dc2.NewCredentialUserRegistrationDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*i3c5eb35d4514961f4cbfafd47ac65dc5ed997312a3e247f4b53d8004051d0194.DailyPrintUsageByPrinterRequestBuilder) {
    return i3c5eb35d4514961f4cbfafd47ac65dc5ed997312a3e247f4b53d8004051d0194.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return id64618d3ddc94c3dc7162d9da67abd3172a3d1bf99f56c9991906d6421e665be.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i2993fd51754851864ba7c1a8dfc4c807b19f8200be6ce2e4964d2b687966d694.DailyPrintUsageByUserRequestBuilder) {
    return i2993fd51754851864ba7c1a8dfc4c807b19f8200be6ce2e4964d2b687966d694.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i6325a2796f62004e900be3097636ac3cdb4b5f89fff67309d7747a1f3c033bf5.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*i51d8f12ca8bef3b5a605ad76c9d75dc8e73de7c0bb0b203ff61c881d868d0391.DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return i51d8f12ca8bef3b5a605ad76c9d75dc8e73de7c0bb0b203ff61c881d868d0391.NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i81f53b8b26ae53be90ad15f1e85f74bb88ada9cad7e6d8f0b3cb9ce23e2e0f46.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*i6858d47e90b1334353133cd0114055d1772fbe2294c3c8b1f87448025955a7f9.DailyPrintUsageSummariesByUserRequestBuilder) {
    return i6858d47e90b1334353133cd0114055d1772fbe2294c3c8b1f87448025955a7f9.NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.dailyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return if50f0dbd8a90dfcd7816cc16a43a257b56afed52876190fb217ae57a9d4679e0.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationDeviceActivity builds and executes requests for operations under \reports\microsoft.graph.deviceConfigurationDeviceActivity()
func (m *ReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*i6e57d0d859f4603fe4b6c70f2f421bec46c3044cb19b132e453e8c658e9e5b9c.DeviceConfigurationDeviceActivityRequestBuilder) {
    return i6e57d0d859f4603fe4b6c70f2f421bec46c3044cb19b132e453e8c658e9e5b9c.NewDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationUserActivity builds and executes requests for operations under \reports\microsoft.graph.deviceConfigurationUserActivity()
func (m *ReportsRequestBuilder) DeviceConfigurationUserActivity()(*ieda57b6fe3b5918581e568356ba76f653b708c99154876060bace124d82a1bad.DeviceConfigurationUserActivityRequestBuilder) {
    return ieda57b6fe3b5918581e568356ba76f653b708c99154876060bace124d82a1bad.NewDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get reports
func (m *ReportsRequestBuilder) Get(options *ReportsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewReportRoot() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRoot), nil
}
// GetAttackSimulationRepeatOffenders builds and executes requests for operations under \reports\microsoft.graph.getAttackSimulationRepeatOffenders()
func (m *ReportsRequestBuilder) GetAttackSimulationRepeatOffenders()(*i94b276da4083eda748c10c94dfe16c6d4742c5a65341271d9d71d7381dae0214.GetAttackSimulationRepeatOffendersRequestBuilder) {
    return i94b276da4083eda748c10c94dfe16c6d4742c5a65341271d9d71d7381dae0214.NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationSimulationUserCoverage builds and executes requests for operations under \reports\microsoft.graph.getAttackSimulationSimulationUserCoverage()
func (m *ReportsRequestBuilder) GetAttackSimulationSimulationUserCoverage()(*i434cde18544a55e4247ac040b6fccc58d1ed919655f86445cff7c94dcfda1437.GetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return i434cde18544a55e4247ac040b6fccc58d1ed919655f86445cff7c94dcfda1437.NewGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationTrainingUserCoverage builds and executes requests for operations under \reports\microsoft.graph.getAttackSimulationTrainingUserCoverage()
func (m *ReportsRequestBuilder) GetAttackSimulationTrainingUserCoverage()(*ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d.GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return ia4a439ad585549ca6a264d3c659884bde60cd44ba79a5d03f941ef02cc50b09d.NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAzureADApplicationSignInSummaryWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getAzureADApplicationSignInSummary(period='{period}')
func (m *ReportsRequestBuilder) GetAzureADApplicationSignInSummaryWithPeriod(period *string)(*i3573cf7b32242b3add3fa97f751c5d986f3543d84edb73bd531d941dccc6d8c8.GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return i3573cf7b32242b3add3fa97f751c5d986f3543d84edb73bd531d941dccc6d8c8.NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADFeatureUsageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getAzureADFeatureUsage(period='{period}')
func (m *ReportsRequestBuilder) GetAzureADFeatureUsageWithPeriod(period *string)(*if34dee9f0158cce73ba132047204086837bb3168fd06c8f7e138d3d5d60b773e.GetAzureADFeatureUsageWithPeriodRequestBuilder) {
    return if34dee9f0158cce73ba132047204086837bb3168fd06c8f7e138d3d5d60b773e.NewGetAzureADFeatureUsageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADLicenseUsageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getAzureADLicenseUsage(period='{period}')
func (m *ReportsRequestBuilder) GetAzureADLicenseUsageWithPeriod(period *string)(*i5d36536b2d00cade7bb67be5034066eaa35f6bcc8eb10887fab5db989e0190b7.GetAzureADLicenseUsageWithPeriodRequestBuilder) {
    return i5d36536b2d00cade7bb67be5034066eaa35f6bcc8eb10887fab5db989e0190b7.NewGetAzureADLicenseUsageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetAzureADUserFeatureUsage builds and executes requests for operations under \reports\microsoft.graph.getAzureADUserFeatureUsage()
func (m *ReportsRequestBuilder) GetAzureADUserFeatureUsage()(*iff7cc4e7a992aa4c240079c5a2255ed095f68878fb58b3ad449dc4ad2229a76c.GetAzureADUserFeatureUsageRequestBuilder) {
    return iff7cc4e7a992aa4c240079c5a2255ed095f68878fb58b3ad449dc4ad2229a76c.NewGetAzureADUserFeatureUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCredentialUsageSummaryWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getCredentialUsageSummary(period='{period}')
func (m *ReportsRequestBuilder) GetCredentialUsageSummaryWithPeriod(period *string)(*ifacf533bfae54389c178233012b220f0cc75de6629790a7585f840fac461be0b.GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return ifacf533bfae54389c178233012b220f0cc75de6629790a7585f840fac461be0b.NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUserRegistrationCount builds and executes requests for operations under \reports\microsoft.graph.getCredentialUserRegistrationCount()
func (m *ReportsRequestBuilder) GetCredentialUserRegistrationCount()(*ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0.GetCredentialUserRegistrationCountRequestBuilder) {
    return ie687487da5ebce56cf90e0244fdf194dfab098133081e652570dd1678b6015e0.NewGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEmailActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*i557e886461aa53bc4de7819875552204d4fe40dab2a518db1a29f2061c23a669.GetEmailActivityCountsWithPeriodRequestBuilder) {
    return i557e886461aa53bc4de7819875552204d4fe40dab2a518db1a29f2061c23a669.NewGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*i7217ecce10bb1170c5d1635128247d30624546692ebc20935f655e3ce4b68c47.GetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return i7217ecce10bb1170c5d1635128247d30624546692ebc20935f655e3ce4b68c47.NewGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd.GetEmailActivityUserDetailWithDateRequestBuilder) {
    return ic8f641d2b8a763b7b6f48c7c85bdce9e87996fdd746bbaa9d7af7769245c71bd.NewGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e.GetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return i02290aae39c6d9344aa34fcba0783a063f860b4fcf84f34c4a01b0e2b489dd4e.NewGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageAppsUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageAppsUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*i920ca9bdad0e6d6c258a83d691a206e4dcbd817fcf55f61727be543b8f6a51e1.GetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return i920ca9bdad0e6d6c258a83d691a206e4dcbd817fcf55f61727be543b8f6a51e1.NewGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8.GetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return id16f9ce1d99f7d1c50c8c3911b7b3317fd937252ea9673d1e38cedfec41b08e8.NewGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserDetail(date={date})
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i4a720b5f345e91b0bf967bf525d76d5febe85151037aa8327de70f7fd0aea25a.GetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return i4a720b5f345e91b0bf967bf525d76d5febe85151037aa8327de70f7fd0aea25a.NewGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailAppUsageUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*ifd317ad9016a7541a0ad611c0d46387984fd5c9e4b06f7fce9718fc0b5c0e52f.GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return ifd317ad9016a7541a0ad611c0d46387984fd5c9e4b06f7fce9718fc0b5c0e52f.NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageVersionsUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getEmailAppUsageVersionsUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*i2dc572cfc357c8939c4ee58887229a9bc49bdbe95b636fb6d5460ee9894192d6.GetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return i2dc572cfc357c8939c4ee58887229a9bc49bdbe95b636fb6d5460ee9894192d6.NewGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime builds and executes requests for operations under \reports\microsoft.graph.getGroupArchivedPrintJobs(groupId='{groupId}',startDateTime={startDateTime},endDateTime={endDateTime})
func (m *ReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i2bb1d1186591127bed84408c868b7d735b4f2f8bbd3d5294b967b9b6ae5594b9.GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i2bb1d1186591127bed84408c868b7d735b4f2f8bbd3d5294b967b9b6ae5594b9.NewGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, groupId, startDateTime, endDateTime);
}
// GetM365AppPlatformUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getM365AppPlatformUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetM365AppPlatformUserCountsWithPeriod(period *string)(*ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e.GetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return ie14903a7ed9a3dee9cbc3334e7ce7a1476f522c78051624951404accb2167f9e.NewGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getM365AppUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetM365AppUserCountsWithPeriod(period *string)(*i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84.GetM365AppUserCountsWithPeriodRequestBuilder) {
    return i02b366b70fc1ec373b8983375e316d896baf6faad4edf2a4121986b8230d5a84.NewGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getM365AppUserDetail(date={date})
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i95b1164d20aad2d39c57c4b38fd44a5d2f2d05bf631cb13a075f37b71a7ac5b4.GetM365AppUserDetailWithDateRequestBuilder) {
    return i95b1164d20aad2d39c57c4b38fd44a5d2f2d05bf631cb13a075f37b71a7ac5b4.NewGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetM365AppUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getM365AppUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithPeriod(period *string)(*iac59a09de72fb2fd2f0077d123253c2ef987a511fbd554c000b1e3714111f572.GetM365AppUserDetailWithPeriodRequestBuilder) {
    return iac59a09de72fb2fd2f0077d123253c2ef987a511fbd554c000b1e3714111f572.NewGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageDetail(period='{period}')
func (m *ReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34.GetMailboxUsageDetailWithPeriodRequestBuilder) {
    return i1db07755998954593b7fa5949e7cf86000be094a324a5936d567e29944523e34.NewGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageMailboxCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageMailboxCounts(period='{period}')
func (m *ReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230.GetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return id819e13d5d8b37a469ab265d456430d17b707e1443609d80dec1be87354e1230.NewGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageQuotaStatusMailboxCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageQuotaStatusMailboxCounts(period='{period}')
func (m *ReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*i6ad597f06da6b987f26da44b061b096836613ef8b752128cb12644670a0d63a7.GetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return i6ad597f06da6b987f26da44b061b096836613ef8b752128cb12644670a0d63a7.NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageStorageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getMailboxUsageStorage(period='{period}')
func (m *ReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*iacb9bbe21799f4942d290c5e1ab7f9050f7193bb0196f318f8707bee63c0ac5d.GetMailboxUsageStorageWithPeriodRequestBuilder) {
    return iacb9bbe21799f4942d290c5e1ab7f9050f7193bb0196f318f8707bee63c0ac5d.NewGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActivationCounts builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationCounts()
func (m *ReportsRequestBuilder) GetOffice365ActivationCounts()(*i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a.GetOffice365ActivationCountsRequestBuilder) {
    return i052f94a1678ad3a8ae670a83cf6aef6ddbcfd78a8b9342c2cdafd1d680943f1a.NewGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserCounts builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationsUserCounts()
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*ia983902699c45d63b5babc13320daacdb1564e3d45458d3d6153bc2112fd2eb6.GetOffice365ActivationsUserCountsRequestBuilder) {
    return ia983902699c45d63b5babc13320daacdb1564e3d45458d3d6153bc2112fd2eb6.NewGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserDetail builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActivationsUserDetail()
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*ieab560554737ae1c313ebdb1132ed35ae880a6f013ece3d8e5dc8786eeb39894.GetOffice365ActivationsUserDetailRequestBuilder) {
    return ieab560554737ae1c313ebdb1132ed35ae880a6f013ece3d8e5dc8786eeb39894.NewGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActiveUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*if08b0f6462c0374cc04d6591fd00c6f903d80cb1aa4bcf0269bf16987dcf17a2.GetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return if08b0f6462c0374cc04d6591fd00c6f903d80cb1aa4bcf0269bf16987dcf17a2.NewGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActiveUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserDetail(date={date})
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*if261b2578ae42789e9c954041da7ecdc03603a05edf6ba0be09893dbe05b3142.GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return if261b2578ae42789e9c954041da7ecdc03603a05edf6ba0be09893dbe05b3142.NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365ActiveUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365ActiveUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d.GetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return i0fd2673ecfb12b3ec83069cf88b196dc71f0f84a3d2c7338c7750a1a62cef69d.NewGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*if587a78dac698fac07fde7dbafd659ec926b72e5abce7d5081eac704ddb5fc12.GetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return if587a78dac698fac07fde7dbafd659ec926b72e5abce7d5081eac704ddb5fc12.NewGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityDetail(date={date})
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i5542a4e920cc301f9fd5b6d010166452a3ec1b96b595f0b1f8869da5b984d4ff.GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return i5542a4e920cc301f9fd5b6d010166452a3ec1b96b595f0b1f8869da5b984d4ff.NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365GroupsActivityDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityDetail(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*iedb2c5fd4b1885cfde5af8d13d7955fcc21c8d6faeb064c46aaad884983733d9.GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return iedb2c5fd4b1885cfde5af8d13d7955fcc21c8d6faeb064c46aaad884983733d9.NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityFileCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityFileCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6.GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return id15f1fe25c88100898caf929a8f7f6c572e5e85b62bb67478ea5f270f24cf6f6.NewGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityGroupCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityGroupCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*i5abd3f1e1b1701310329336b2d99ca6c43f6556a82d34472ce1d9c428c9d9c9c.GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return i5abd3f1e1b1701310329336b2d99ca6c43f6556a82d34472ce1d9c428c9d9c9c.NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityStorageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365GroupsActivityStorage(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab.GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return i1f6b0d46fc4b338b7b612e7abff7bb8e48704a3d9c51a3b2122ebc3e77b54fab.NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ServicesUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOffice365ServicesUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*i9f1ae835c41e1a61632007ea0b4fd227f894397673ddb731c0d36bd49470d5c3.GetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return i9f1ae835c41e1a61632007ea0b4fd227f894397673ddb731c0d36bd49470d5c3.NewGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityFileCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityFileCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*i53891a3a3983e3f57fac19166c9af13eee82757751926f31b54f16d6fa23d29a.GetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return i53891a3a3983e3f57fac19166c9af13eee82757751926f31b54f16d6fa23d29a.NewGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*ic3f682e7cc9725074cc23ecc64d8d813b5f6c39a93bcdd47f5b04e1e36e31a8c.GetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return ic3f682e7cc9725074cc23ecc64d8d813b5f6c39a93bcdd47f5b04e1e36e31a8c.NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ic5da14e3ff43c5e6fd44a093a284bd0d470c3a9d4b65d5e2bce503b45cd374fd.GetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return ic5da14e3ff43c5e6fd44a093a284bd0d470c3a9d4b65d5e2bce503b45cd374fd.NewGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9.GetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return id1b5f80bb654118e8428176f73bc7413eb087dc72b5acfe75aa4b6d09d7405b9.NewGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*i7ad76340f4a0f8d49588345643b23f3a29ec1fd2161c87b7866114a35eda410a.GetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return i7ad76340f4a0f8d49588345643b23f3a29ec1fd2161c87b7866114a35eda410a.NewGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountDetail(date={date})
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i48a44087adc49827c59b599bc3e9c1ea8e284030a4416856821ea71b09310ed5.GetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return i48a44087adc49827c59b599bc3e9c1ea8e284030a4416856821ea71b09310ed5.NewGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveUsageAccountDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageAccountDetail(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*ib546a3980e2bf4e36d4a57260b208c8311e952ceecd05de93cafd1097b8221d0.GetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return ib546a3980e2bf4e36d4a57260b208c8311e952ceecd05de93cafd1097b8221d0.NewGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageFileCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageFileCounts(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*i91a39045a7be5817859a873d2638b0c156baab984d6270f68a448faf3297b18c.GetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return i91a39045a7be5817859a873d2638b0c156baab984d6270f68a448faf3297b18c.NewGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageStorageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getOneDriveUsageStorage(period='{period}')
func (m *ReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*i4c2e0be102a6b6d64d75b2ce3aed7eee80548af7512aca09bc765525a936d621.GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return i4c2e0be102a6b6d64d75b2ce3aed7eee80548af7512aca09bc765525a936d621.NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime builds and executes requests for operations under \reports\microsoft.graph.getPrinterArchivedPrintJobs(printerId='{printerId}',startDateTime={startDateTime},endDateTime={endDateTime})
func (m *ReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e.GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i0075784d86aaab84840c6dfaee5d8126c33aad7617cd0b1b92b629aec8f8834e.NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, printerId, startDateTime, endDateTime);
}
// GetRelyingPartyDetailedSummaryWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getRelyingPartyDetailedSummary(period='{period}')
func (m *ReportsRequestBuilder) GetRelyingPartyDetailedSummaryWithPeriod(period *string)(*i250edbdb590c4bab056109a11a6cf515b7193da960e757e0325976fb00e6b7ef.GetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return i250edbdb590c4bab056109a11a6cf515b7193da960e757e0325976fb00e6b7ef.NewGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityFileCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityFileCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*i973d4295c78aa9f58867f3b650c440eb7e1901bb527ccd20a8687328d4ecbbc5.GetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return i973d4295c78aa9f58867f3b650c440eb7e1901bb527ccd20a8687328d4ecbbc5.NewGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityPagesWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityPages(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*i944bfcd02a17a424e910488ba7b947a5fdad6f5a75edb524a1d32a729f796746.GetSharePointActivityPagesWithPeriodRequestBuilder) {
    return i944bfcd02a17a424e910488ba7b947a5fdad6f5a75edb524a1d32a729f796746.NewGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927.GetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return i826cac8af5755c1771dcf21902dcbe7a8c861f3d2cc2659b97318fa6ca801927.NewGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512.GetSharePointActivityUserDetailWithDateRequestBuilder) {
    return ide13a00269d183accb3fb129eacee077596bbbd9e1af9886a8fa5b82f24f9512.NewGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420.GetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return id0a1a6564efba17eb292c6d47fa074ac2b98c19c12c1a1495d8b7813b7b74420.NewGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageDetail(date={date})
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ia905ae63e07e238ab152ce311e0525628c071873fd34232733350f04752d8678.GetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return ia905ae63e07e238ab152ce311e0525628c071873fd34232733350f04752d8678.NewGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointSiteUsageDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageDetail(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*i93a80e4540ddcb67a9776e160a250432a3c4dc682adbf3a39d6e2f9904bb9447.GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return i93a80e4540ddcb67a9776e160a250432a3c4dc682adbf3a39d6e2f9904bb9447.NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageFileCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageFileCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*i4d0f15879165ef15ce895cf317adfa50fa15bf6ed244a753762d508f5d5ba5e3.GetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return i4d0f15879165ef15ce895cf317adfa50fa15bf6ed244a753762d508f5d5ba5e3.NewGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsagePagesWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsagePages(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367.GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return i1bec555870a869ece7648150219ddbeab418f800e6fad9515d22cbf803333367.NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageSiteCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageSiteCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d.GetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return i0cd404348cbf2ffec8d63dfcc57d473fd6cc07e2de00393f339b1dddab3abb9d.NewGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageStorageWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSharePointSiteUsageStorage(period='{period}')
func (m *ReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*i7143ac54c5e4852fa7c71022277b025f4e1140c3bb0b90e520b5505dfa09cc23.GetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return i7143ac54c5e4852fa7c71022277b025f4e1140c3bb0b90e520b5505dfa09cc23.NewGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*i4c439992e985ac97d20ac69debd6694765696eb8acf494c0dd9c1e843dde5943.GetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return i4c439992e985ac97d20ac69debd6694765696eb8acf494c0dd9c1e843dde5943.NewGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*i35fab8ead83b7a35b3fc3dff29cb309b459f147de1727e7671b1e3ea1085f7a2.GetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return i35fab8ead83b7a35b3fc3dff29cb309b459f147de1727e7671b1e3ea1085f7a2.NewGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ic2bba236674645b7dd61f3f93d935e166db2806a283ce92ca066bb178547b2b8.GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return ic2bba236674645b7dd61f3f93d935e166db2806a283ce92ca066bb178547b2b8.NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5.GetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return i86328c0af7138fd234511a05c1e2c56a024dd34613e898482f7ded7f6e3bebb5.NewGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageDistributionUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i6419b1c71ac2474e588fa2df68b31ad5cf99896a1c878c0b2f23b8a51f4794c0.GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i6419b1c71ac2474e588fa2df68b31ad5cf99896a1c878c0b2f23b8a51f4794c0.NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*if220cda8a33cfb0dc18bc5f4b4ed4f3787a1bb1f4392a8c67322dae2cf9910ce.GetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return if220cda8a33cfb0dc18bc5f4b4ed4f3787a1bb1f4392a8c67322dae2cf9910ce.NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserDetail(date={date})
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i7026ee7cc69ab0c0a76cfd363e6e41a9e764f78dfc9fc6b226788e43010b3585.GetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return i7026ee7cc69ab0c0a76cfd363e6e41a9e764f78dfc9fc6b226788e43010b3585.NewGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessDeviceUsageUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*if239d7585794526a1ec1b695c30f15b3ba8659d065ac77ec5d3622719b1c1fc4.GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return if239d7585794526a1ec1b695c30f15b3ba8659d065ac77ec5d3622719b1c1fc4.NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*i9c1f3873c5c6f05ebb01275ba1e1cadbcac87226b44c1b890d5571cd53f331ef.GetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return i9c1f3873c5c6f05ebb01275ba1e1cadbcac87226b44c1b890d5571cd53f331ef.NewGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityMinuteCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139.GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i1c1aeb0a39d86448beafcab36d8f25bbd683572663c82a4c40c4117725cbb139.NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessOrganizerActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*i27a2b61ad92d96f585909f7842612603cbacfc26e12dc059d2f4878490c04101.GetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return i27a2b61ad92d96f585909f7842612603cbacfc26e12dc059d2f4878490c04101.NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*if958ce2337186d28b6ff688907b51b0baf53d270d71152d85463560542cb332f.GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return if958ce2337186d28b6ff688907b51b0baf53d270d71152d85463560542cb332f.NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityMinuteCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*i4a335c23d6834f12971f5a81936e763c8d804a2f22f65b093947ae6a7f4a1634.GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return i4a335c23d6834f12971f5a81936e763c8d804a2f22f65b093947ae6a7f4a1634.NewGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessParticipantActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*i4a8fe9122551f2771b4a12b8839c70ffec863bf5960e9fcfc031a4e55d7adab6.GetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return i4a8fe9122551f2771b4a12b8839c70ffec863bf5960e9fcfc031a4e55d7adab6.NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*iab64d32b308e755b9724f53342c4489cd12c811324ffcfc1eebaeb1f6d63937f.GetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return iab64d32b308e755b9724f53342c4489cd12c811324ffcfc1eebaeb1f6d63937f.NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityMinuteCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*i38d9dde997bfdced4eba6828af9a431514b84509f8ed82648d159166a512a10b.GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return i38d9dde997bfdced4eba6828af9a431514b84509f8ed82648d159166a512a10b.NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getSkypeForBusinessPeerToPeerActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*iab808a61daaca7b2e378fc8a0289994ae7ddf284a49112e4b8b742d6f3cba446.GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return iab808a61daaca7b2e378fc8a0289994ae7ddf284a49112e4b8b742d6f3cba446.NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageDistributionTotalUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b.GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return ie5262971fb14e5bb4e229a4a02e0be2ec0626336ca05f560fa1193d0dbd0ed2b.NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageDistributionUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*i38018b8ba711a07a4df4ba04370314b5a7eb52d04ded8775a2b76a35218402ef.GetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return i38018b8ba711a07a4df4ba04370314b5a7eb52d04ded8775a2b76a35218402ef.NewGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageTotalUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageTotalUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*i2ca5100bc1b71c10f318988b462b80ebf432370aae38b64473c17dd121558bfd.GetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return i2ca5100bc1b71c10f318988b462b80ebf432370aae38b64473c17dd121558bfd.NewGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*i98b74c2f146a48ff844d3021737827fcc82215b40e873605b5b930c294e36b6b.GetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i98b74c2f146a48ff844d3021737827fcc82215b40e873605b5b930c294e36b6b.NewGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserDetail(date={date})
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3.GetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return id78f0043579ad6440a94b391bf422870702c33dd35a5f50c4bb1e102ee06b3a3.NewGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsDeviceUsageUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*i8deb95de65d724b352c9c68714d954885c65ccac1053a18e56072a07014b8d5b.GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i8deb95de65d724b352c9c68714d954885c65ccac1053a18e56072a07014b8d5b.NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf.GetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return i0c4bdb7ddbfb1808024438760b4f61f6ed2b3cc8f01006a4aff2c07b79309ebf.NewGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityDistributionTotalUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*ic64ccf8d946071ab046b2c31875bb4f39c0a87b879466e64adf3aea9c73aa206.GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return ic64ccf8d946071ab046b2c31875bb4f39c0a87b879466e64adf3aea9c73aa206.NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityDistributionUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*i26bbf530ddb18f9b5655ff8fa87d5f3b5ad1da8c08a1db360b8de025d0362210.GetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return i26bbf530ddb18f9b5655ff8fa87d5f3b5ad1da8c08a1db360b8de025d0362210.NewGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityTotalCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalCountsWithPeriod(period *string)(*ic201bfab351b7c4943cb69827ffc49f717f8d030c2c09e022994a10813ccebb6.GetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return ic201bfab351b7c4943cb69827ffc49f717f8d030c2c09e022994a10813ccebb6.NewGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityTotalUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*i8e6535187547a4b5d5c14160716a4c020758272be3e5ca22edba3b46b3087b52.GetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return i8e6535187547a4b5d5c14160716a4c020758272be3e5ca22edba3b46b3087b52.NewGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba.GetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return i1c97679a2593f53f0d56c8ca04b6129e80fb9633559094ab0a4b44e77ab52bba.NewGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*iedafbd1ed73588865f6b43d238dd6c81c00b86c2c48d2ba9ecb1e727bf24afe8.GetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return iedafbd1ed73588865f6b43d238dd6c81c00b86c2c48d2ba9ecb1e727bf24afe8.NewGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsUserActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTeamsUserActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*iee0f7f8db82d7bf47fa5eff92350ac6235e7af4d5a35bf22f0c190f09903fd62.GetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return iee0f7f8db82d7bf47fa5eff92350ac6235e7af4d5a35bf22f0c190f09903fd62.NewGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTenantSecureScoresWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getTenantSecureScores(period={period})
func (m *ReportsRequestBuilder) GetTenantSecureScoresWithPeriod(period *int32)(*ia83b8d65117b3e838e53c360304a89995009d10e47e5dce012a0e8c0406b67e2.GetTenantSecureScoresWithPeriodRequestBuilder) {
    return ia83b8d65117b3e838e53c360304a89995009d10e47e5dce012a0e8c0406b67e2.NewGetTenantSecureScoresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime builds and executes requests for operations under \reports\microsoft.graph.getUserArchivedPrintJobs(userId='{userId}',startDateTime={startDateTime},endDateTime={endDateTime})
func (m *ReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(userId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*i50d831cff0d03367e9939eb5f21f0a71e577dce8e1e220d3877edf21418cf620.GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return i50d831cff0d03367e9939eb5f21f0a71e577dce8e1e220d3877edf21418cf620.NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, userId, startDateTime, endDateTime);
}
// GetYammerActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*i6735cb8d512199566e336179e484d86cb97201979137f1ef2ce9fdd8486c616e.GetYammerActivityCountsWithPeriodRequestBuilder) {
    return i6735cb8d512199566e336179e484d86cb97201979137f1ef2ce9fdd8486c616e.NewGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*i23ae3dadc25e31a369068a3ac3cd7ff57c8b4823dcea9650bd4b312d6f57e7e4.GetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return i23ae3dadc25e31a369068a3ac3cd7ff57c8b4823dcea9650bd4b312d6f57e7e4.NewGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserDetail(date={date})
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b.GetYammerActivityUserDetailWithDateRequestBuilder) {
    return ica598676eb8a952415f863c021b34052e8169be21d16ff5ea2e1e76d21584e0b.NewGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerActivityUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerActivityUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*i4f16df07e7ec9e6dd5cb52fe46110ce00f5f74418530157eaee04c72ffe1c4e1.GetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return i4f16df07e7ec9e6dd5cb52fe46110ce00f5f74418530157eaee04c72ffe1c4e1.NewGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageDistributionUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageDistributionUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42.GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return ia54f5885db60032a888c2668e50091452912a071e528a3da9eba97c6d431fa42.NewGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*i943fdce89c1d1ccca5b0f3b2db4cfc85447eafff319d6137111ed073b10c023e.GetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return i943fdce89c1d1ccca5b0f3b2db4cfc85447eafff319d6137111ed073b10c023e.NewGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserDetail(date={date})
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45.GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return i1b46d4e69ba1b93deaf785955d5b648a3c2c33ff773c7f59b4f4318c1b70fd45.NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerDeviceUsageUserDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserDetail(period='{period}')
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*i9a872942082503636d059c4c1175930485a0f27ecbf3b0043154dc6658ef476f.GetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return i9a872942082503636d059c4c1175930485a0f27ecbf3b0043154dc6658ef476f.NewGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa.GetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return i85af9b00ffe076a423eb5fd31460ced62f5ace48b1303f5ddde47a395bc696aa.NewGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityDetailWithDate builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityDetail(date={date})
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*ibaf7d6c4cd3e8f19279c8ea66be35293ecddff285916af2302dad085654fb7a4.GetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return ibaf7d6c4cd3e8f19279c8ea66be35293ecddff285916af2302dad085654fb7a4.NewGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerGroupsActivityDetailWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityDetail(period='{period}')
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*i7c9189a010c6688a7fa9746221903390178e46db9ee2896c06cd2e3e781e3d3a.GetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return i7c9189a010c6688a7fa9746221903390178e46db9ee2896c06cd2e3e781e3d3a.NewGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityGroupCountsWithPeriod builds and executes requests for operations under \reports\microsoft.graph.getYammerGroupsActivityGroupCounts(period='{period}')
func (m *ReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2.GetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return ie2fde0ac510d57a06bc25774f0c38ce72bf91bfbe2bd2abeb497f770a0b349c2.NewGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentAbandonmentDetails(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(skip *int32, top *int32, filter *string, skipToken *string)(*i24783836c67f3cbffe50d47dcbcb2588e9de66ffd8f9a5c8da53fa898e5e9446.ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i24783836c67f3cbffe50d47dcbcb2588e9de66ffd8f9a5c8da53fa898e5e9446.NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top, filter, skipToken);
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentAbandonmentSummary(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(skip *int32, top *int32, filter *string, skipToken *string)(*i27595bc84745a6124aef7bab74a2b609c8775cdb2e31edd5259685331293e617.ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return i27595bc84745a6124aef7bab74a2b609c8775cdb2e31edd5259685331293e617.NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top, filter, skipToken);
}
// ManagedDeviceEnrollmentFailureDetails builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentFailureDetails()
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*i8a14754850e94320f257277f7b5983b2dfdea28ccf42833202bf6ecd26e1e9b5.ManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return i8a14754850e94320f257277f7b5983b2dfdea28ccf42833202bf6ecd26e1e9b5.NewManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentFailureDetails(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(skip *int32, top *int32, filter *string, skipToken *string)(*ib217a9ab86db27e8b3dce55808a82185fd090bc6c4245f5330f5c1b0c08a3a61.ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return ib217a9ab86db27e8b3dce55808a82185fd090bc6c4245f5330f5c1b0c08a3a61.NewManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top, filter, skipToken);
}
// ManagedDeviceEnrollmentFailureTrends builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentFailureTrends()
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureTrends()(*i791f698d83fc125e8a1806cd3e8a916e376ace94be4518e908c757250c9c64e5.ManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return i791f698d83fc125e8a1806cd3e8a916e376ace94be4518e908c757250c9c64e5.NewManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailures builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentTopFailures()
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*i492793d896d7d799d0ff9804f56092f81d113794a8916673f9c324f5a2053ffc.ManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return i492793d896d7d799d0ff9804f56092f81d113794a8916673f9c324f5a2053ffc.NewManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailuresWithPeriod builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentTopFailures(period='{period}')
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361.ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return i03f4f94b2b9e3fb286bbe022526518bbd7fc8be03e1497a413610d279818c361.NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*i37eb861166f6d0555df75cb64e973d22d9bb9c2a8c03ca2f6c29444c37e3649f.MonthlyPrintUsageByPrinterRequestBuilder) {
    return i37eb861166f6d0555df75cb64e973d22d9bb9c2a8c03ca2f6c29444c37e3649f.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i9d3332b918775296ebe9eb5ed2e432af9b61c784aceacaca03a043c49a0ff6db.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*ifbd8842a6664ec96c7d893455943698937b1af3fe05558ca8e628721e162064b.MonthlyPrintUsageByUserRequestBuilder) {
    return ifbd8842a6664ec96c7d893455943698937b1af3fe05558ca8e628721e162064b.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i5983a969121cc458289b2e9bf5b347bab6120a3ecd4adca075e2189cc9460c6f.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*i59460f776b2ba7ec0487cd235a89ad5d26f27c3ec165f0790954cfc2a01b9637.MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return i59460f776b2ba7ec0487cd235a89ad5d26f27c3ec165f0790954cfc2a01b9637.NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3.PrintUsageByPrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i9e0e6696a72655a2583635eb551d9445661402ca5fa5c8f22f82920d3cbae9a3.NewPrintUsageByPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*i700648709215c245e3f56a9bb592458abcaa4d198b5956df2836a08af3f74be4.MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return i700648709215c245e3f56a9bb592458abcaa4d198b5956df2836a08af3f74be4.NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.monthlyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8.PrintUsageByUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return ie40751cac1559213f1bbd658ede6043abd8ffd3463aa7abc8847eea9e36a36f8.NewPrintUsageByUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update reports
func (m *ReportsRequestBuilder) Patch(options *ReportsRequestBuilderPatchOptions)(error) {
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
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.UserCredentialUsageDetailsRequestBuilder) {
    return ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.reports.userCredentialUsageDetails.item collection
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.UserCredentialUsageDetailsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails_id"] = id
    }
    return ie1a7217ebe434d381aad98ca8c1e72bdf39d5d523e00d2462f814224c339246b.NewUserCredentialUsageDetailsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
