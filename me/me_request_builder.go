package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i03229d0781e0f5cd7e37b6bfd0cc1ec32a80065860d46ca6653e2824b55e7a9b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/exportdeviceandappmanagementdata"
    i03a3fdcb3f1d65e925013bb34fbf99663faa80c904b6f351d0e6428219666962 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/wipeandblockmanagedapps"
    i0615d7ba982e06a459c77c0876aee46ad012da5184e784858f5b99991e2b0442 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/registereddevices"
    i07a4f67d8feaf18ceaf6ff3228c4e2151126ceafbbfc83b3ec78e7051ad29e3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/oauth2permissiongrants"
    i09926aa3bed79b432266994f222c004b6ae84f55e1a732e52c5b3bcfd12c0dbd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmanageddeviceswithappfailures"
    i0eb8ecd1000d66dda1934d83eac5577b8ea0cf4a8da7d8e15e8e571a3472f70f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/checkmembergroups"
    i0f2bfaba642e50d9558b61b34b334847148895996f43398d274c0ab5863e3691 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/invalidateallrefreshtokens"
    i14f21981f966d7340ab241a9387463ba0fa4f07fa6c8f1727484c38fd9f16445 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages"
    i1511fdcb4792e1cd0954759201108e358dce36a83cc13223eb05070615e77a63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication"
    i1b5b0e98fa954be818e2b28d677bf536697e31009fabfc6339e28a5c7f35b844 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmanagedappdiagnosticstatuses"
    i1e3b0d7a9333d29fe71e530e04f6c5df334b5fc86243d3de0a534cd1921b4342 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/licensedetails"
    i1fbcfb89922aa73b86452575a4cc9ab699fca4815a0e676f6034b57fd5d44b48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/exportpersonaldata"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i21ad1b6dc11aa1d930604e3ce5077412ee57e50680df809a175042e68a73614f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/appconsentrequestsforapproval"
    i28d4fe3704f9f62380a603c7ae29a13c1aa06a7213de7ee6c9fcdc652c40d463 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manager"
    i292db412d9ce5ec2a850c1c6b4baeedec3e63413f2b63ac27c792b9f269d30ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/unblockmanagedapps"
    i29d15cf2ebb01a69779caa08b86678cd70e064e98f4f681c20dbb2f7c5068dcc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/profile"
    i30c1e1ac829691448f09a03d6d733f5968c3f85b0bffdf5d7b090718320e20d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/managedappregistrations"
    i324d21921cd9355af9f2a6ec8c94175c55a026ea9a5fb294255e1dce8e50392b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ismanagedappuserblocked"
    i32b15a9294b27af382e25627185a91297460ea69c1f81010b12afb6a1002db10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups"
    i32d8dc857112f42b52a344de9badc89e743eef960325ebe8910ed2013ecc8d37 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner"
    i3340a668307034bde84b6a65c934778fd43edb7307d82098a167906d463eedf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/photos"
    i34ae02cf2ff34fd803e19d8c2f9af8318285107c66cfcf00845014f3af686197 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events"
    i352fab4a04ba762c64656855e65ab019a254014b3080f8077b4c95cb2a1bc817 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/notifications"
    i3aea592c6cb33f9e19fb4c62bdecfe7aa46cc52af88598190451b79c5e046e26 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approleassignments"
    i3d2fc4aaa671f5674c171e95891301e78ae4f8ba90acf335402d69b338dd13e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmembergroups"
    i43b9ac10e794b8d72997f0f596f132feef5df0a633cd05dfca8089dcb0fda9e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/geteffectivedeviceenrollmentconfigurations"
    i477445cf5f8c496c3730b4a0183177997d93184d877ac862e8af6b8963a3985d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/translateexchangeids"
    i50cc07d8bbb211c98359ee0edb3bbe464c3db02260cf3b52b5e024439787f960 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders"
    i54eac61560762ff611b6700b2de929f760e7e3284932b7576e7ba46ccca69cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/analytics"
    i56a77360b87654a0c9a4dc6869de45fd14b76ff5a92744f6b05b9ae81ad57de8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approvals"
    i58d75b863c46999036fda79f4f851a6091e0d01581ba2e143d130dfeeacb7900 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview"
    i59bc57755ba56f6eb5fb2534c9eb2c79f00639d8861695b6b49b5e54cb223606 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivereports"
    i59f8bb1b42aa564603d502b2fba98210a12453d9b9b17878081d967b73a86343 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats"
    i5db4fee05cecd2bed7cb42dc5c6046db8f61d4b37063c6e5b2e1cf1e39ed1ff6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices"
    i6094b414759c1e98e84cd17f8f0abd06625fbedd0396bf8e9f07ae4fb687ebd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/agreementacceptances"
    i6237fa040b5eaf455abf5f698c389202ad776ba022b1566206439de15b5da115 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/presence"
    i639da1f678f79126f23fc2a686a6d6b7a6fa40fa1f13db941966bde0b30a5d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmailtips"
    i65eadd82fcbfd207b16c3e7dbf0c8d1067289ab718ba0ce591cfeea315d5d9cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote"
    i68ae3b4d592d0057eb4a687d378300648953a03d22862f907fd2ff66d201eeac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook"
    i73355d68106cd7f92ed5b9ffb4a52e59ba3ec00b272cb67ff6e7d1e181314937 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings"
    i7724a6079225e5b4fb109ecbcdfd2d5f8d5f2e78a512dff130978a6bd7ee0aa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/wipemanagedappregistrationsbyazureaddeviceid"
    i772944f98e62be37df8bbf30068a248f0f1b2b5e7b1938cbf0e1c6fe680be3c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/removealldevicesfrommanagement"
    i7b613e1f229abaea577330e8430e25c30f3890e517f541d6c2fa0bc338c85ebe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/teamwork"
    i7c69216b282055d4e66d8afba1065734bd28d022f420d5bbcb72d7908ab5be9b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof"
    i803324c6cdf723fa7f5b938e982a7ffcf6f95545f0ea1bc56a5f1eb50564bd46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection"
    i817e3209fc02c599989a60b33b6d0f340e5c9e6b1b3bbbbd86ddc77d00c52c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devicemanagementtroubleshootingevents"
    i832752c7fc0bc2aa43f3e00911325be0e4d0bda70e3caea85e2f66813f2639de "github.com/microsoftgraph/msgraph-beta-sdk-go/me/scopedrolememberof"
    i8739463faa09b1ded16c7fdad85ed43b78a2933e4ee6d9d96c9b06b283a3512c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/followedsites"
    i8901d4e1325ace1429a8c17db365e494f21be9f8e446d11691208513cffb80b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/insights"
    i89a1dab18d1109c45e8c0c86a6d955a75e944cfe5e7e065442996fd88fa2c678 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmanagedapppolicies"
    i8a7e58760e5040d83674b84187ec6926c3f92735d35bae15c23754128925bd0a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/findrooms"
    i8cce4b6e74c8589123f3152b94d3da2fe384b14aaf490d31fcac3986bd90ea29 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations"
    i8e4aadd7ed94b8a9e55552f74c51e245a91ebb89ada3b45c4e66b5cf03cd9f79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/findroomswithroomlist"
    i90fa087b5c274048383161e7735aac09f360b7664473b3d487cd2b6704001d83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts"
    i92acb123b06e31af45c2264dee537df2ecd185721affa39b8e1030fafa5511ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/assignlicense"
    i93405a8f89e4184693c09c0e84c792b9567480596ed3662ca8f3958b1922f361 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups"
    i94ad8fd049d36ab89c2fcb747ee48bc3c2cc92a488cb7ccc4156b4f110bf708d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/exportdeviceandappmanagementdatawithskipwithtop"
    i96b86fb89d12c60762a9b35297bd48c5f5a1f0a29a853fcaee1bcb8408c8a5a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/people"
    i9707f15743f2bc7fd08620852e99c80974fc9e9ce7762c2b1f6e4183508d9946 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/windowsinformationprotectiondeviceregistrations"
    i9859a1065a4c497ab3647c006153457dba946e2848eb7a50b67818697c69417b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedteams"
    ia16187a55eb9c3b9d429798be330d5df578e28729956c97a8f63c887d6ae66ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices"
    ia675b038deca3e32dad85e01ba41f1c919b5d1c779f5258a2941327eb3bfc17e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/restore"
    ia8fd6d4235fbe5ae19767a14b931820ca5e311a4a03fc9dae71c1801c7d3092f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/extensions"
    ia96d1f92e78f5de8132d5ce2d35b245bc47a307071b141447e1003202987f9f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/wipemanagedappregistrationsbydevicetag"
    iaa6383ba91d1d220deef98e9ce97a49a107e15b05eb6b2b3ce3c9d6c48e4f326 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/settings"
    iad3415fdc487951976de6996b1757be3243496c6f390b601c1a613f27ed8532c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects"
    iad50c20589920c64f75c4cdd1f4c1f88512e5a1f950831214803a7371a3e08d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/createdobjects"
    iae69e180638c5aed1f619923028c2eaece62f16c4c270d564d5db9fe857c53dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof"
    ib1105a7971b2cea2de83264504bc098dcb6964362763345d03825f5dc4820205 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/findroomlists"
    ib30fb712d3d83e6954d3924903b95a5b7b6e5583943adf5f99bff52563c27b53 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmemberobjects"
    ib4b34bad27f650f5cc9e03461b6d4c452608904535bd12df112a16ed4ac7edd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives"
    ib4c685509c6fd2907ec160446f20f296a4dbe123da549c4e6b9e7ad4ade1b374 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/sendmail"
    ib5b4d50478937130cb8d53f4c75d274229e4c2a9061b1c9a485907eed04f6aef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/usagerights"
    ib772ed3906151e13ada7c53574b5c2f5d30c9e8ecb1ed9791b44b643dc36193a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getmanageddeviceswithfailedorpendingapps"
    ib7801ea348e04a14861bd095693d71eb64b7e24fb1b05827984cd33d80abade1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/revokesigninsessions"
    ib8352abd928cf31d46bbc0656ef7642b2860ab7b592e379055bb077b4e4cf0bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents"
    ibcfbd279423471d7ced24105aee3850ee4270273f5274f95f9fd6d84f7c38720 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar"
    ic0ab9313e76e3a6d55ea39ac104e3622b84d0a0d01c5dfeaa1269e4831d8e98c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/inferenceclassification"
    ic0dbc76bcf27b6d539c0c34397a288e535f2347b4d4495b35ea9d88800664c13 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/getloggedonmanageddevices"
    ic3e86b248df5ff542e4828dcd02d564185edf5341445413104e1a9484f0da2b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders"
    ic3fe384252788f8c3b5dfb7e9064552b08d0785f3f25b9abb64555be87652a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/reprocesslicenseassignment"
    ic4947066d01e412f641a12a3bc8459085b3cfcfae771df1108dc6319f3536eeb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports"
    ic675e3e79f073212aeb3e72b7a32c8529405e4fe1cd4bd69b885ace5326ca895 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/todo"
    ic93f01ea250d4839ef099dc4c41fc62d7f5aa62afb84e4edfd9fcdf0c69e48cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/owneddevices"
    icfe0fa0def53eb8ead53a9c7e14cfa7b99aa67d289bbd5786ae5577155574961 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/reminderviewwithstartdatetimewithenddatetime"
    id9569f521a060e524166adc08c72508f54ed83854d63694ef777241e7b5c7d54 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/wipemanagedappregistrationbydevicetag"
    idd48ab55f5d1da1d01a7519debf68afdf64ff43e4e90500cda63eeb76402e355 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/tasks"
    ie92980fd5edaafba8f00572025a84bd9779e54cea477e030692babbc26d2d460 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances"
    iedfcf53badc587bde8ee6b1219608face4b6f46e0c89c51a0078a7a22748abef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/findmeetingtimes"
    if0ad722a34fdbb042c20f761755bd2d4f9dddac3c4c1392721a1ccd0148b3751 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drive"
    if3a811497f1bf492f3bbe053e28833b1140d2c90725fd72ed4cd2e0b4068fbf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security"
    if53d06a6944a15601fc9e0ecdcd8e010fb0ac38ce7857531732470170377fa9c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/activities"
    if54b727c8f82e94874a8380d8e6c910ef5f2b561abf736f53bd43812b1e23237 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars"
    if6e3044a34bc88e3264e7687443bb1bb7614a36ae9afdf6f7b962a1086748067 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs"
    if7c226ac1fb697611e5cc28b6bee0ac1f1c6d51898fd4bf78542842ae06536db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approleassignedresources"
    if9a235a81139ac366dce38caf3d59b07903735e656934a8208ee758d24cf000e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/checkmemberobjects"
    ifa6354c03b0a297cab93d13f8bc78708546aaeb1be40e137a7d07a262fce0ea9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/photo"
    ifae3a8927e14c995c44faea104b1594d3360d9fcbfabb3aba01f5da5671d0ba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileappintentandstates"
    ife93540a444acfcfb4f3255ba514e3f0e0ccf9edd12b1c46210ac705e0bdcbb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/changepassword"
    iff3e752f4ecf4dec684b2068c0ed6eed3b3822d4c0d9c259e1f976d12f88e194 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/activateserviceplan"
    i03cc0125fff403f10c9217cf6465d79f8b49c6205fad977f867396abf1c69779 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contacts/item"
    i044d2d5a570570eb3a834c3174faaabdad4efba49df934fb07c40ac12ad4b23d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devicemanagementtroubleshootingevents/item"
    i05218ac15109586aa0bf9d6b77b99bb99266067c1544901cbf50ca419e3741de "github.com/microsoftgraph/msgraph-beta-sdk-go/me/manageddevices/item"
    i0653b5ffd95682aea4c98ecb6c539dcb517315be3e2dc79188b7541f859ea8f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/extensions/item"
    i08f6400faf92b934d06b0b6861e5bf74d25a77d5b3d67c02b4779a1c0c52e960 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/scopedrolememberof/item"
    i0c7e22c7b3c700585549fd382bfa2115e0602d341e57daae7fcadf078a118fe1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/followedsites/item"
    i0feb3876f59829427e859bbade1e51c0d003a6a64ce67ad441b6d60c3087afce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/memberof/item"
    i18a300ce18206c9aa35cc46abd9486be5d2faa2ceb82ea15418d59e63400891b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations/item"
    i2cc8c08facade07f4bb7f43d3d598d3e49ab04058fe12463d7fbae54f4163f48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivereports/item"
    i311c4a65493f2e31d79a3e80ede9ce3128c8b2d5daa9d17e0c2524819eac5475 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item"
    i32120ccc8322d0ae4c79827cab0e21691e8b625292351396f2ee7ac9681caaed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/windowsinformationprotectiondeviceregistrations/item"
    i35b6d595de33f1d461b0d50276e049227be20e321f99ec0e204fb7be2adcec52 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/photos/item"
    i36124d7f986d43728dd4e4ec2537be43e1bb06aafbfb36badf4c29a92469fb74 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/createdobjects/item"
    i3a460a91d75d2133059711f9a41626eb89eb964605a87101f4ce61ca4b32eccc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item"
    i3c8eaf1c6dffbb0f85330d394a60715c6b08b77fd1865fa06fb70c1395e95efd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item"
    i4013f2adca2e80594707aa753bbc43cfa3c88b1f94557550a9de9a87010a65c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/notifications/item"
    i44d43c5e67fd914a879c6699eebbfd2589387d74bdc6e95565084cb16342c066 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item"
    i47e5b20f37978ff329fe5a3a058924b96b429a635613f9be0625a219ebbbf612 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/registereddevices/item"
    i4839dd1e34130c46f192e50856e4104f29c91c197b815b494b1a69ba6342e13b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/activities/item"
    i4ae3afd01d0851c73223650832b296a7fc705ea82f59dc6662bdc0e4bd134cb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item"
    i4b4dce57580b9d6f6781b551395b4c975ce6835ca4860fc6b89a5b72ff02daca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/transitivememberof/item"
    i52e44ebeddeea9833d46e37baa56c121056336e4d9cb3a738d6c05114eec6d10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedteams/item"
    i61d4b18aabe612d406051c63cdd71a82846d2421be7afc134002abc5e139f3e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approleassignments/item"
    i634abce1c174c1a4b8f987a99848dda449764fbd35f786e01464e876d22d956f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/usagerights/item"
    i74f1402adb8d2c11ef5ac6f6525b503a0958d79d4e161c048f6e2a01614b07a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/agreementacceptances/item"
    i88cb1f65e88e2be976a13788c7251e4bd43b629bce1c058fad78b21dc16f7062 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item"
    i8ce653616e56cac134ff5b2659d920c1b77db0e8cb12f0164e596429c11be8d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/ownedobjects/item"
    i964afe288259202a4e1a7dec3dd29d7b0e636062c4231f642d5958afce7f732f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/appconsentrequestsforapproval/item"
    i9bd0541127f761aca94a07e02c775d2bb43f46e5a893d047575b16b9ca24532d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/managedappregistrations/item"
    ia7d1c875f6bea5636947a7bc93b746fa5ee6b4cd81be4f22e93d7ec1ebd1a694 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/directreports/item"
    ia9d1397713a75de885b5f1c5f94ee39bba9f659b5ea81f0dc7406a018b1a48c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approvals/item"
    iac4e70fbd3847a4ded822f798b430d1011dd4ece8cf666478eff963f7ae42837 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileapptroubleshootingevents/item"
    ib2bb171f9206e1ef2294f99eecb6860ae7f1a3bd79cd3ca6732a4e3373030c46 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/licensedetails/item"
    ibdf6f337e272c2c489b46bc7469a81713a5f1141291912d55abe2a182b61324d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/owneddevices/item"
    icc221144e4bd42176c4e6641427cf56400dd43be43d6c4b3c993c43b91a69c05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item"
    icccda8c1e029711a6c2c010bde446ea283c88f9167583edd4cb4b8891df462bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/contactfolders/item"
    id23f8c7e0c3929eeb36c197eb371fd7242f148f968513d3b42532b35904d2c7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/devices/item"
    id7edb04c2fe4b42c55656962bf3bb27391a93462cc5a89234df1ddf704bdf49d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mobileappintentandstates/item"
    ie03252d8bc0ab84afe8d33f54e4b7dcd6b92db4b4d16eb6d427177dc1271a9c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/approleassignedresources/item"
    ie23caf2134464ac7a439e71ae9e04c43d1e7edeb9f60834c6450c46b2a73ef3d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item"
    iedd9b42c61703c4f6442a904b683a6508ee4bf335825ff84ce57fb2616e2cce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/people/item"
    if68ffab8944bb223858e55142c815f29e64d6051ab2ef6f6ce0d495d4d87f4a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item"
    if7c78154b39b2da1e1c12f7bf63ddd85b353b2d31d33f3ba11e1dfd43b13ad7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/oauth2permissiongrants/item"
    if8b811d7fbecff009f6661aef902befaba11f32ccbaaa91de716f8c4b103c293 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item"
    ifd2db3d8aa0a0ec27c89cb2949e6936c2d3de139fb986c5773dfe17008496fc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/pendingaccessreviewinstances/item"
)

// MeRequestBuilder provides operations to manage the user singleton.
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeRequestBuilderGetQueryParameters returns the user or organizational contact assigned as the user's manager. Optionally, you can expand the manager's chain up to the root node.
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeRequestBuilderGetQueryParameters
}
// MeRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivateServicePlan the activateServicePlan property
func (m *MeRequestBuilder) ActivateServicePlan()(*iff3e752f4ecf4dec684b2068c0ed6eed3b3822d4c0d9c259e1f976d12f88e194.ActivateServicePlanRequestBuilder) {
    return iff3e752f4ecf4dec684b2068c0ed6eed3b3822d4c0d9c259e1f976d12f88e194.NewActivateServicePlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Activities the activities property
func (m *MeRequestBuilder) Activities()(*if53d06a6944a15601fc9e0ecdcd8e010fb0ac38ce7857531732470170377fa9c.ActivitiesRequestBuilder) {
    return if53d06a6944a15601fc9e0ecdcd8e010fb0ac38ce7857531732470170377fa9c.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.activities.item collection
func (m *MeRequestBuilder) ActivitiesById(id string)(*i4839dd1e34130c46f192e50856e4104f29c91c197b815b494b1a69ba6342e13b.UserActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userActivity%2Did"] = id
    }
    return i4839dd1e34130c46f192e50856e4104f29c91c197b815b494b1a69ba6342e13b.NewUserActivityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AgreementAcceptances the agreementAcceptances property
func (m *MeRequestBuilder) AgreementAcceptances()(*i6094b414759c1e98e84cd17f8f0abd06625fbedd0396bf8e9f07ae4fb687ebd2.AgreementAcceptancesRequestBuilder) {
    return i6094b414759c1e98e84cd17f8f0abd06625fbedd0396bf8e9f07ae4fb687ebd2.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.agreementAcceptances.item collection
func (m *MeRequestBuilder) AgreementAcceptancesById(id string)(*i74f1402adb8d2c11ef5ac6f6525b503a0958d79d4e161c048f6e2a01614b07a4.AgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance%2Did"] = id
    }
    return i74f1402adb8d2c11ef5ac6f6525b503a0958d79d4e161c048f6e2a01614b07a4.NewAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *MeRequestBuilder) Analytics()(*i54eac61560762ff611b6700b2de929f760e7e3284932b7576e7ba46ccca69cdc.AnalyticsRequestBuilder) {
    return i54eac61560762ff611b6700b2de929f760e7e3284932b7576e7ba46ccca69cdc.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApproval the appConsentRequestsForApproval property
func (m *MeRequestBuilder) AppConsentRequestsForApproval()(*i21ad1b6dc11aa1d930604e3ce5077412ee57e50680df809a175042e68a73614f.AppConsentRequestsForApprovalRequestBuilder) {
    return i21ad1b6dc11aa1d930604e3ce5077412ee57e50680df809a175042e68a73614f.NewAppConsentRequestsForApprovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppConsentRequestsForApprovalById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.appConsentRequestsForApproval.item collection
func (m *MeRequestBuilder) AppConsentRequestsForApprovalById(id string)(*i964afe288259202a4e1a7dec3dd29d7b0e636062c4231f642d5958afce7f732f.AppConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appConsentRequest%2Did"] = id
    }
    return i964afe288259202a4e1a7dec3dd29d7b0e636062c4231f642d5958afce7f732f.NewAppConsentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignedResources the appRoleAssignedResources property
func (m *MeRequestBuilder) AppRoleAssignedResources()(*if7c226ac1fb697611e5cc28b6bee0ac1f1c6d51898fd4bf78542842ae06536db.AppRoleAssignedResourcesRequestBuilder) {
    return if7c226ac1fb697611e5cc28b6bee0ac1f1c6d51898fd4bf78542842ae06536db.NewAppRoleAssignedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.appRoleAssignedResources.item collection
func (m *MeRequestBuilder) AppRoleAssignedResourcesById(id string)(*ie03252d8bc0ab84afe8d33f54e4b7dcd6b92db4b4d16eb6d427177dc1271a9c0.ServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipal%2Did"] = id
    }
    return ie03252d8bc0ab84afe8d33f54e4b7dcd6b92db4b4d16eb6d427177dc1271a9c0.NewServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AppRoleAssignments the appRoleAssignments property
func (m *MeRequestBuilder) AppRoleAssignments()(*i3aea592c6cb33f9e19fb4c62bdecfe7aa46cc52af88598190451b79c5e046e26.AppRoleAssignmentsRequestBuilder) {
    return i3aea592c6cb33f9e19fb4c62bdecfe7aa46cc52af88598190451b79c5e046e26.NewAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.appRoleAssignments.item collection
func (m *MeRequestBuilder) AppRoleAssignmentsById(id string)(*i61d4b18aabe612d406051c63cdd71a82846d2421be7afc134002abc5e139f3e0.AppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return i61d4b18aabe612d406051c63cdd71a82846d2421be7afc134002abc5e139f3e0.NewAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Approvals the approvals property
func (m *MeRequestBuilder) Approvals()(*i56a77360b87654a0c9a4dc6869de45fd14b76ff5a92744f6b05b9ae81ad57de8.ApprovalsRequestBuilder) {
    return i56a77360b87654a0c9a4dc6869de45fd14b76ff5a92744f6b05b9ae81ad57de8.NewApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApprovalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.approvals.item collection
func (m *MeRequestBuilder) ApprovalsById(id string)(*ia9d1397713a75de885b5f1c5f94ee39bba9f659b5ea81f0dc7406a018b1a48c3.ApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return ia9d1397713a75de885b5f1c5f94ee39bba9f659b5ea81f0dc7406a018b1a48c3.NewApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense the assignLicense property
func (m *MeRequestBuilder) AssignLicense()(*i92acb123b06e31af45c2264dee537df2ecd185721affa39b8e1030fafa5511ed.AssignLicenseRequestBuilder) {
    return i92acb123b06e31af45c2264dee537df2ecd185721affa39b8e1030fafa5511ed.NewAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Authentication the authentication property
func (m *MeRequestBuilder) Authentication()(*i1511fdcb4792e1cd0954759201108e358dce36a83cc13223eb05070615e77a63.AuthenticationRequestBuilder) {
    return i1511fdcb4792e1cd0954759201108e358dce36a83cc13223eb05070615e77a63.NewAuthenticationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar the calendar property
func (m *MeRequestBuilder) Calendar()(*ibcfbd279423471d7ced24105aee3850ee4270273f5274f95f9fd6d84f7c38720.CalendarRequestBuilder) {
    return ibcfbd279423471d7ced24105aee3850ee4270273f5274f95f9fd6d84f7c38720.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroups the calendarGroups property
func (m *MeRequestBuilder) CalendarGroups()(*i32b15a9294b27af382e25627185a91297460ea69c1f81010b12afb6a1002db10.CalendarGroupsRequestBuilder) {
    return i32b15a9294b27af382e25627185a91297460ea69c1f81010b12afb6a1002db10.NewCalendarGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item collection
func (m *MeRequestBuilder) CalendarGroupsById(id string)(*i88cb1f65e88e2be976a13788c7251e4bd43b629bce1c058fad78b21dc16f7062.CalendarGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarGroup%2Did"] = id
    }
    return i88cb1f65e88e2be976a13788c7251e4bd43b629bce1c058fad78b21dc16f7062.NewCalendarGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendars the calendars property
func (m *MeRequestBuilder) Calendars()(*if54b727c8f82e94874a8380d8e6c910ef5f2b561abf736f53bd43812b1e23237.CalendarsRequestBuilder) {
    return if54b727c8f82e94874a8380d8e6c910ef5f2b561abf736f53bd43812b1e23237.NewCalendarsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item collection
func (m *MeRequestBuilder) CalendarsById(id string)(*i44d43c5e67fd914a879c6699eebbfd2589387d74bdc6e95565084cb16342c066.CalendarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendar%2Did"] = id
    }
    return i44d43c5e67fd914a879c6699eebbfd2589387d74bdc6e95565084cb16342c066.NewCalendarItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *MeRequestBuilder) CalendarView()(*i58d75b863c46999036fda79f4f851a6091e0d01581ba2e143d130dfeeacb7900.CalendarViewRequestBuilder) {
    return i58d75b863c46999036fda79f4f851a6091e0d01581ba2e143d130dfeeacb7900.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item collection
func (m *MeRequestBuilder) CalendarViewById(id string)(*if8b811d7fbecff009f6661aef902befaba11f32ccbaaa91de716f8c4b103c293.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return if8b811d7fbecff009f6661aef902befaba11f32ccbaaa91de716f8c4b103c293.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChangePassword the changePassword property
func (m *MeRequestBuilder) ChangePassword()(*ife93540a444acfcfb4f3255ba514e3f0e0ccf9edd12b1c46210ac705e0bdcbb7.ChangePasswordRequestBuilder) {
    return ife93540a444acfcfb4f3255ba514e3f0e0ccf9edd12b1c46210ac705e0bdcbb7.NewChangePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Chats the chats property
func (m *MeRequestBuilder) Chats()(*i59f8bb1b42aa564603d502b2fba98210a12453d9b9b17878081d967b73a86343.ChatsRequestBuilder) {
    return i59f8bb1b42aa564603d502b2fba98210a12453d9b9b17878081d967b73a86343.NewChatsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChatsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item collection
func (m *MeRequestBuilder) ChatsById(id string)(*if68ffab8944bb223858e55142c815f29e64d6051ab2ef6f6ce0d495d4d87f4a9.ChatItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chat%2Did"] = id
    }
    return if68ffab8944bb223858e55142c815f29e64d6051ab2ef6f6ce0d495d4d87f4a9.NewChatItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *MeRequestBuilder) CheckMemberGroups()(*i0eb8ecd1000d66dda1934d83eac5577b8ea0cf4a8da7d8e15e8e571a3472f70f.CheckMemberGroupsRequestBuilder) {
    return i0eb8ecd1000d66dda1934d83eac5577b8ea0cf4a8da7d8e15e8e571a3472f70f.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *MeRequestBuilder) CheckMemberObjects()(*if9a235a81139ac366dce38caf3d59b07903735e656934a8208ee758d24cf000e.CheckMemberObjectsRequestBuilder) {
    return if9a235a81139ac366dce38caf3d59b07903735e656934a8208ee758d24cf000e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCs the cloudPCs property
func (m *MeRequestBuilder) CloudPCs()(*if6e3044a34bc88e3264e7687443bb1bb7614a36ae9afdf6f7b962a1086748067.CloudPCsRequestBuilder) {
    return if6e3044a34bc88e3264e7687443bb1bb7614a36ae9afdf6f7b962a1086748067.NewCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.cloudPCs.item collection
func (m *MeRequestBuilder) CloudPCsById(id string)(*i311c4a65493f2e31d79a3e80ede9ce3128c8b2d5daa9d17e0c2524819eac5475.CloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return i311c4a65493f2e31d79a3e80ede9ce3128c8b2d5daa9d17e0c2524819eac5475.NewCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
// ContactFolders the contactFolders property
func (m *MeRequestBuilder) ContactFolders()(*i50cc07d8bbb211c98359ee0edb3bbe464c3db02260cf3b52b5e024439787f960.ContactFoldersRequestBuilder) {
    return i50cc07d8bbb211c98359ee0edb3bbe464c3db02260cf3b52b5e024439787f960.NewContactFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contactFolders.item collection
func (m *MeRequestBuilder) ContactFoldersById(id string)(*icccda8c1e029711a6c2c010bde446ea283c88f9167583edd4cb4b8891df462bd.ContactFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contactFolder%2Did"] = id
    }
    return icccda8c1e029711a6c2c010bde446ea283c88f9167583edd4cb4b8891df462bd.NewContactFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Contacts the contacts property
func (m *MeRequestBuilder) Contacts()(*i90fa087b5c274048383161e7735aac09f360b7664473b3d487cd2b6704001d83.ContactsRequestBuilder) {
    return i90fa087b5c274048383161e7735aac09f360b7664473b3d487cd2b6704001d83.NewContactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContactsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.contacts.item collection
func (m *MeRequestBuilder) ContactsById(id string)(*i03cc0125fff403f10c9217cf6465d79f8b49c6205fad977f867396abf1c69779.ContactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contact%2Did"] = id
    }
    return i03cc0125fff403f10c9217cf6465d79f8b49c6205fad977f867396abf1c69779.NewContactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreatedObjects the createdObjects property
func (m *MeRequestBuilder) CreatedObjects()(*iad50c20589920c64f75c4cdd1f4c1f88512e5a1f950831214803a7371a3e08d3.CreatedObjectsRequestBuilder) {
    return iad50c20589920c64f75c4cdd1f4c1f88512e5a1f950831214803a7371a3e08d3.NewCreatedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatedObjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.createdObjects.item collection
func (m *MeRequestBuilder) CreatedObjectsById(id string)(*i36124d7f986d43728dd4e4ec2537be43e1bb06aafbfb36badf4c29a92469fb74.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i36124d7f986d43728dd4e4ec2537be43e1bb06aafbfb36badf4c29a92469fb74.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateGetRequestInformation returns the user or organizational contact assigned as the user's manager. Optionally, you can expand the manager's chain up to the root node.
func (m *MeRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns the user or organizational contact assigned as the user's manager. Optionally, you can expand the manager's chain up to the root node.
func (m *MeRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeviceEnrollmentConfigurations the deviceEnrollmentConfigurations property
func (m *MeRequestBuilder) DeviceEnrollmentConfigurations()(*i8cce4b6e74c8589123f3152b94d3da2fe384b14aaf490d31fcac3986bd90ea29.DeviceEnrollmentConfigurationsRequestBuilder) {
    return i8cce4b6e74c8589123f3152b94d3da2fe384b14aaf490d31fcac3986bd90ea29.NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.deviceEnrollmentConfigurations.item collection
func (m *MeRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*i18a300ce18206c9aa35cc46abd9486be5d2faa2ceb82ea15418d59e63400891b.DeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return i18a300ce18206c9aa35cc46abd9486be5d2faa2ceb82ea15418d59e63400891b.NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementTroubleshootingEvents the deviceManagementTroubleshootingEvents property
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEvents()(*i817e3209fc02c599989a60b33b6d0f340e5c9e6b1b3bbbbd86ddc77d00c52c38.DeviceManagementTroubleshootingEventsRequestBuilder) {
    return i817e3209fc02c599989a60b33b6d0f340e5c9e6b1b3bbbbd86ddc77d00c52c38.NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.deviceManagementTroubleshootingEvents.item collection
func (m *MeRequestBuilder) DeviceManagementTroubleshootingEventsById(id string)(*i044d2d5a570570eb3a834c3174faaabdad4efba49df934fb07c40ac12ad4b23d.DeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return i044d2d5a570570eb3a834c3174faaabdad4efba49df934fb07c40ac12ad4b23d.NewDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Devices the devices property
func (m *MeRequestBuilder) Devices()(*i5db4fee05cecd2bed7cb42dc5c6046db8f61d4b37063c6e5b2e1cf1e39ed1ff6.DevicesRequestBuilder) {
    return i5db4fee05cecd2bed7cb42dc5c6046db8f61d4b37063c6e5b2e1cf1e39ed1ff6.NewDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.devices.item collection
func (m *MeRequestBuilder) DevicesById(id string)(*id23f8c7e0c3929eeb36c197eb371fd7242f148f968513d3b42532b35904d2c7e.DeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["device%2Did"] = id
    }
    return id23f8c7e0c3929eeb36c197eb371fd7242f148f968513d3b42532b35904d2c7e.NewDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DirectReports the directReports property
func (m *MeRequestBuilder) DirectReports()(*ic4947066d01e412f641a12a3bc8459085b3cfcfae771df1108dc6319f3536eeb.DirectReportsRequestBuilder) {
    return ic4947066d01e412f641a12a3bc8459085b3cfcfae771df1108dc6319f3536eeb.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.directReports.item collection
func (m *MeRequestBuilder) DirectReportsById(id string)(*ia7d1c875f6bea5636947a7bc93b746fa5ee6b4cd81be4f22e93d7ec1ebd1a694.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ia7d1c875f6bea5636947a7bc93b746fa5ee6b4cd81be4f22e93d7ec1ebd1a694.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Drive the drive property
func (m *MeRequestBuilder) Drive()(*if0ad722a34fdbb042c20f761755bd2d4f9dddac3c4c1392721a1ccd0148b3751.DriveRequestBuilder) {
    return if0ad722a34fdbb042c20f761755bd2d4f9dddac3c4c1392721a1ccd0148b3751.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *MeRequestBuilder) Drives()(*ib4b34bad27f650f5cc9e03461b6d4c452608904535bd12df112a16ed4ac7edd9.DrivesRequestBuilder) {
    return ib4b34bad27f650f5cc9e03461b6d4c452608904535bd12df112a16ed4ac7edd9.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item collection
func (m *MeRequestBuilder) DrivesById(id string)(*i3c8eaf1c6dffbb0f85330d394a60715c6b08b77fd1865fa06fb70c1395e95efd.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i3c8eaf1c6dffbb0f85330d394a60715c6b08b77fd1865fa06fb70c1395e95efd.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Events the events property
func (m *MeRequestBuilder) Events()(*i34ae02cf2ff34fd803e19d8c2f9af8318285107c66cfcf00845014f3af686197.EventsRequestBuilder) {
    return i34ae02cf2ff34fd803e19d8c2f9af8318285107c66cfcf00845014f3af686197.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item collection
func (m *MeRequestBuilder) EventsById(id string)(*icc221144e4bd42176c4e6641427cf56400dd43be43d6c4b3c993c43b91a69c05.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return icc221144e4bd42176c4e6641427cf56400dd43be43d6c4b3c993c43b91a69c05.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExportDeviceAndAppManagementData provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) ExportDeviceAndAppManagementData()(*i03229d0781e0f5cd7e37b6bfd0cc1ec32a80065860d46ca6653e2824b55e7a9b.ExportDeviceAndAppManagementDataRequestBuilder) {
    return i03229d0781e0f5cd7e37b6bfd0cc1ec32a80065860d46ca6653e2824b55e7a9b.NewExportDeviceAndAppManagementDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportDeviceAndAppManagementDataWithSkipWithTop provides operations to call the exportDeviceAndAppManagementData method.
func (m *MeRequestBuilder) ExportDeviceAndAppManagementDataWithSkipWithTop(skip *int32, top *int32)(*i94ad8fd049d36ab89c2fcb747ee48bc3c2cc92a488cb7ccc4156b4f110bf708d.ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return i94ad8fd049d36ab89c2fcb747ee48bc3c2cc92a488cb7ccc4156b4f110bf708d.NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(m.pathParameters, m.requestAdapter, skip, top);
}
// ExportPersonalData the exportPersonalData property
func (m *MeRequestBuilder) ExportPersonalData()(*i1fbcfb89922aa73b86452575a4cc9ab699fca4815a0e676f6034b57fd5d44b48.ExportPersonalDataRequestBuilder) {
    return i1fbcfb89922aa73b86452575a4cc9ab699fca4815a0e676f6034b57fd5d44b48.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *MeRequestBuilder) Extensions()(*ia8fd6d4235fbe5ae19767a14b931820ca5e311a4a03fc9dae71c1801c7d3092f.ExtensionsRequestBuilder) {
    return ia8fd6d4235fbe5ae19767a14b931820ca5e311a4a03fc9dae71c1801c7d3092f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.extensions.item collection
func (m *MeRequestBuilder) ExtensionsById(id string)(*i0653b5ffd95682aea4c98ecb6c539dcb517315be3e2dc79188b7541f859ea8f7.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i0653b5ffd95682aea4c98ecb6c539dcb517315be3e2dc79188b7541f859ea8f7.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FindMeetingTimes the findMeetingTimes property
func (m *MeRequestBuilder) FindMeetingTimes()(*iedfcf53badc587bde8ee6b1219608face4b6f46e0c89c51a0078a7a22748abef.FindMeetingTimesRequestBuilder) {
    return iedfcf53badc587bde8ee6b1219608face4b6f46e0c89c51a0078a7a22748abef.NewFindMeetingTimesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomLists provides operations to call the findRoomLists method.
func (m *MeRequestBuilder) FindRoomLists()(*ib1105a7971b2cea2de83264504bc098dcb6964362763345d03825f5dc4820205.FindRoomListsRequestBuilder) {
    return ib1105a7971b2cea2de83264504bc098dcb6964362763345d03825f5dc4820205.NewFindRoomListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRooms provides operations to call the findRooms method.
func (m *MeRequestBuilder) FindRooms()(*i8a7e58760e5040d83674b84187ec6926c3f92735d35bae15c23754128925bd0a.FindRoomsRequestBuilder) {
    return i8a7e58760e5040d83674b84187ec6926c3f92735d35bae15c23754128925bd0a.NewFindRoomsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FindRoomsWithRoomList provides operations to call the findRooms method.
func (m *MeRequestBuilder) FindRoomsWithRoomList(roomList *string)(*i8e4aadd7ed94b8a9e55552f74c51e245a91ebb89ada3b45c4e66b5cf03cd9f79.FindRoomsWithRoomListRequestBuilder) {
    return i8e4aadd7ed94b8a9e55552f74c51e245a91ebb89ada3b45c4e66b5cf03cd9f79.NewFindRoomsWithRoomListRequestBuilderInternal(m.pathParameters, m.requestAdapter, roomList);
}
// FollowedSites the followedSites property
func (m *MeRequestBuilder) FollowedSites()(*i8739463faa09b1ded16c7fdad85ed43b78a2933e4ee6d9d96c9b06b283a3512c.FollowedSitesRequestBuilder) {
    return i8739463faa09b1ded16c7fdad85ed43b78a2933e4ee6d9d96c9b06b283a3512c.NewFollowedSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowedSitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.followedSites.item collection
func (m *MeRequestBuilder) FollowedSitesById(id string)(*i0c7e22c7b3c700585549fd382bfa2115e0602d341e57daae7fcadf078a118fe1.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return i0c7e22c7b3c700585549fd382bfa2115e0602d341e57daae7fcadf078a118fe1.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get returns the user or organizational contact assigned as the user's manager. Optionally, you can expand the manager's chain up to the root node.
func (m *MeRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetEffectiveDeviceEnrollmentConfigurations provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
func (m *MeRequestBuilder) GetEffectiveDeviceEnrollmentConfigurations()(*i43b9ac10e794b8d72997f0f596f132feef5df0a633cd05dfca8089dcb0fda9e1.GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    return i43b9ac10e794b8d72997f0f596f132feef5df0a633cd05dfca8089dcb0fda9e1.NewGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetLoggedOnManagedDevices provides operations to call the getLoggedOnManagedDevices method.
func (m *MeRequestBuilder) GetLoggedOnManagedDevices()(*ic0dbc76bcf27b6d539c0c34397a288e535f2347b4d4495b35ea9d88800664c13.GetLoggedOnManagedDevicesRequestBuilder) {
    return ic0dbc76bcf27b6d539c0c34397a288e535f2347b4d4495b35ea9d88800664c13.NewGetLoggedOnManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMailTips the getMailTips property
func (m *MeRequestBuilder) GetMailTips()(*i639da1f678f79126f23fc2a686a6d6b7a6fa40fa1f13db941966bde0b30a5d12.GetMailTipsRequestBuilder) {
    return i639da1f678f79126f23fc2a686a6d6b7a6fa40fa1f13db941966bde0b30a5d12.NewGetMailTipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppDiagnosticStatuses provides operations to call the getManagedAppDiagnosticStatuses method.
func (m *MeRequestBuilder) GetManagedAppDiagnosticStatuses()(*i1b5b0e98fa954be818e2b28d677bf536697e31009fabfc6339e28a5c7f35b844.GetManagedAppDiagnosticStatusesRequestBuilder) {
    return i1b5b0e98fa954be818e2b28d677bf536697e31009fabfc6339e28a5c7f35b844.NewGetManagedAppDiagnosticStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppPolicies provides operations to call the getManagedAppPolicies method.
func (m *MeRequestBuilder) GetManagedAppPolicies()(*i89a1dab18d1109c45e8c0c86a6d955a75e944cfe5e7e065442996fd88fa2c678.GetManagedAppPoliciesRequestBuilder) {
    return i89a1dab18d1109c45e8c0c86a6d955a75e944cfe5e7e065442996fd88fa2c678.NewGetManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithAppFailures provides operations to call the getManagedDevicesWithAppFailures method.
func (m *MeRequestBuilder) GetManagedDevicesWithAppFailures()(*i09926aa3bed79b432266994f222c004b6ae84f55e1a732e52c5b3bcfd12c0dbd.GetManagedDevicesWithAppFailuresRequestBuilder) {
    return i09926aa3bed79b432266994f222c004b6ae84f55e1a732e52c5b3bcfd12c0dbd.NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedDevicesWithFailedOrPendingApps provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
func (m *MeRequestBuilder) GetManagedDevicesWithFailedOrPendingApps()(*ib772ed3906151e13ada7c53574b5c2f5d30c9e8ecb1ed9791b44b643dc36193a.GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return ib772ed3906151e13ada7c53574b5c2f5d30c9e8ecb1ed9791b44b643dc36193a.NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberGroups the getMemberGroups property
func (m *MeRequestBuilder) GetMemberGroups()(*i3d2fc4aaa671f5674c171e95891301e78ae4f8ba90acf335402d69b338dd13e8.GetMemberGroupsRequestBuilder) {
    return i3d2fc4aaa671f5674c171e95891301e78ae4f8ba90acf335402d69b338dd13e8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *MeRequestBuilder) GetMemberObjects()(*ib30fb712d3d83e6954d3924903b95a5b7b6e5583943adf5f99bff52563c27b53.GetMemberObjectsRequestBuilder) {
    return ib30fb712d3d83e6954d3924903b95a5b7b6e5583943adf5f99bff52563c27b53.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler returns the user or organizational contact assigned as the user's manager. Optionally, you can expand the manager's chain up to the root node.
func (m *MeRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MeRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// InferenceClassification the inferenceClassification property
func (m *MeRequestBuilder) InferenceClassification()(*ic0ab9313e76e3a6d55ea39ac104e3622b84d0a0d01c5dfeaa1269e4831d8e98c.InferenceClassificationRequestBuilder) {
    return ic0ab9313e76e3a6d55ea39ac104e3622b84d0a0d01c5dfeaa1269e4831d8e98c.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InformationProtection the informationProtection property
func (m *MeRequestBuilder) InformationProtection()(*i803324c6cdf723fa7f5b938e982a7ffcf6f95545f0ea1bc56a5f1eb50564bd46.InformationProtectionRequestBuilder) {
    return i803324c6cdf723fa7f5b938e982a7ffcf6f95545f0ea1bc56a5f1eb50564bd46.NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Insights the insights property
func (m *MeRequestBuilder) Insights()(*i8901d4e1325ace1429a8c17db365e494f21be9f8e446d11691208513cffb80b2.InsightsRequestBuilder) {
    return i8901d4e1325ace1429a8c17db365e494f21be9f8e446d11691208513cffb80b2.NewInsightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InvalidateAllRefreshTokens the invalidateAllRefreshTokens property
func (m *MeRequestBuilder) InvalidateAllRefreshTokens()(*i0f2bfaba642e50d9558b61b34b334847148895996f43398d274c0ab5863e3691.InvalidateAllRefreshTokensRequestBuilder) {
    return i0f2bfaba642e50d9558b61b34b334847148895996f43398d274c0ab5863e3691.NewInvalidateAllRefreshTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IsManagedAppUserBlocked provides operations to call the isManagedAppUserBlocked method.
func (m *MeRequestBuilder) IsManagedAppUserBlocked()(*i324d21921cd9355af9f2a6ec8c94175c55a026ea9a5fb294255e1dce8e50392b.IsManagedAppUserBlockedRequestBuilder) {
    return i324d21921cd9355af9f2a6ec8c94175c55a026ea9a5fb294255e1dce8e50392b.NewIsManagedAppUserBlockedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedGroups the joinedGroups property
func (m *MeRequestBuilder) JoinedGroups()(*i93405a8f89e4184693c09c0e84c792b9567480596ed3662ca8f3958b1922f361.JoinedGroupsRequestBuilder) {
    return i93405a8f89e4184693c09c0e84c792b9567480596ed3662ca8f3958b1922f361.NewJoinedGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeams the joinedTeams property
func (m *MeRequestBuilder) JoinedTeams()(*i9859a1065a4c497ab3647c006153457dba946e2848eb7a50b67818697c69417b.JoinedTeamsRequestBuilder) {
    return i9859a1065a4c497ab3647c006153457dba946e2848eb7a50b67818697c69417b.NewJoinedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// JoinedTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedTeams.item collection
func (m *MeRequestBuilder) JoinedTeamsById(id string)(*i52e44ebeddeea9833d46e37baa56c121056336e4d9cb3a738d6c05114eec6d10.TeamItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["team%2Did"] = id
    }
    return i52e44ebeddeea9833d46e37baa56c121056336e4d9cb3a738d6c05114eec6d10.NewTeamItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LicenseDetails the licenseDetails property
func (m *MeRequestBuilder) LicenseDetails()(*i1e3b0d7a9333d29fe71e530e04f6c5df334b5fc86243d3de0a534cd1921b4342.LicenseDetailsRequestBuilder) {
    return i1e3b0d7a9333d29fe71e530e04f6c5df334b5fc86243d3de0a534cd1921b4342.NewLicenseDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LicenseDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.licenseDetails.item collection
func (m *MeRequestBuilder) LicenseDetailsById(id string)(*ib2bb171f9206e1ef2294f99eecb6860ae7f1a3bd79cd3ca6732a4e3373030c46.LicenseDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["licenseDetails%2Did"] = id
    }
    return ib2bb171f9206e1ef2294f99eecb6860ae7f1a3bd79cd3ca6732a4e3373030c46.NewLicenseDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MailFolders the mailFolders property
func (m *MeRequestBuilder) MailFolders()(*ic3e86b248df5ff542e4828dcd02d564185edf5341445413104e1a9484f0da2b5.MailFoldersRequestBuilder) {
    return ic3e86b248df5ff542e4828dcd02d564185edf5341445413104e1a9484f0da2b5.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item collection
func (m *MeRequestBuilder) MailFoldersById(id string)(*i4ae3afd01d0851c73223650832b296a7fc705ea82f59dc6662bdc0e4bd134cb8.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did"] = id
    }
    return i4ae3afd01d0851c73223650832b296a7fc705ea82f59dc6662bdc0e4bd134cb8.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations the managedAppRegistrations property
func (m *MeRequestBuilder) ManagedAppRegistrations()(*i30c1e1ac829691448f09a03d6d733f5968c3f85b0bffdf5d7b090718320e20d1.ManagedAppRegistrationsRequestBuilder) {
    return i30c1e1ac829691448f09a03d6d733f5968c3f85b0bffdf5d7b090718320e20d1.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedAppRegistrations.item collection
func (m *MeRequestBuilder) ManagedAppRegistrationsById(id string)(*i9bd0541127f761aca94a07e02c775d2bb43f46e5a893d047575b16b9ca24532d.ManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return i9bd0541127f761aca94a07e02c775d2bb43f46e5a893d047575b16b9ca24532d.NewManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDevices the managedDevices property
func (m *MeRequestBuilder) ManagedDevices()(*ia16187a55eb9c3b9d429798be330d5df578e28729956c97a8f63c887d6ae66ba.ManagedDevicesRequestBuilder) {
    return ia16187a55eb9c3b9d429798be330d5df578e28729956c97a8f63c887d6ae66ba.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.managedDevices.item collection
func (m *MeRequestBuilder) ManagedDevicesById(id string)(*i05218ac15109586aa0bf9d6b77b99bb99266067c1544901cbf50ca419e3741de.ManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return i05218ac15109586aa0bf9d6b77b99bb99266067c1544901cbf50ca419e3741de.NewManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Manager the manager property
func (m *MeRequestBuilder) Manager()(*i28d4fe3704f9f62380a603c7ae29a13c1aa06a7213de7ee6c9fcdc652c40d463.ManagerRequestBuilder) {
    return i28d4fe3704f9f62380a603c7ae29a13c1aa06a7213de7ee6c9fcdc652c40d463.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
func (m *MeRequestBuilder) MemberOf()(*iae69e180638c5aed1f619923028c2eaece62f16c4c270d564d5db9fe857c53dd.MemberOfRequestBuilder) {
    return iae69e180638c5aed1f619923028c2eaece62f16c4c270d564d5db9fe857c53dd.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.memberOf.item collection
func (m *MeRequestBuilder) MemberOfById(id string)(*i0feb3876f59829427e859bbade1e51c0d003a6a64ce67ad441b6d60c3087afce.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i0feb3876f59829427e859bbade1e51c0d003a6a64ce67ad441b6d60c3087afce.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *MeRequestBuilder) Messages()(*i14f21981f966d7340ab241a9387463ba0fa4f07fa6c8f1727484c38fd9f16445.MessagesRequestBuilder) {
    return i14f21981f966d7340ab241a9387463ba0fa4f07fa6c8f1727484c38fd9f16445.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.messages.item collection
func (m *MeRequestBuilder) MessagesById(id string)(*i3a460a91d75d2133059711f9a41626eb89eb964605a87101f4ce61ca4b32eccc.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return i3a460a91d75d2133059711f9a41626eb89eb964605a87101f4ce61ca4b32eccc.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppIntentAndStates the mobileAppIntentAndStates property
func (m *MeRequestBuilder) MobileAppIntentAndStates()(*ifae3a8927e14c995c44faea104b1594d3360d9fcbfabb3aba01f5da5671d0ba4.MobileAppIntentAndStatesRequestBuilder) {
    return ifae3a8927e14c995c44faea104b1594d3360d9fcbfabb3aba01f5da5671d0ba4.NewMobileAppIntentAndStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppIntentAndStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mobileAppIntentAndStates.item collection
func (m *MeRequestBuilder) MobileAppIntentAndStatesById(id string)(*id7edb04c2fe4b42c55656962bf3bb27391a93462cc5a89234df1ddf704bdf49d.MobileAppIntentAndStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppIntentAndState%2Did"] = id
    }
    return id7edb04c2fe4b42c55656962bf3bb27391a93462cc5a89234df1ddf704bdf49d.NewMobileAppIntentAndStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents the mobileAppTroubleshootingEvents property
func (m *MeRequestBuilder) MobileAppTroubleshootingEvents()(*ib8352abd928cf31d46bbc0656ef7642b2860ab7b592e379055bb077b4e4cf0bd.MobileAppTroubleshootingEventsRequestBuilder) {
    return ib8352abd928cf31d46bbc0656ef7642b2860ab7b592e379055bb077b4e4cf0bd.NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mobileAppTroubleshootingEvents.item collection
func (m *MeRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*iac4e70fbd3847a4ded822f798b430d1011dd4ece8cf666478eff963f7ae42837.MobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return iac4e70fbd3847a4ded822f798b430d1011dd4ece8cf666478eff963f7ae42837.NewMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Notifications the notifications property
func (m *MeRequestBuilder) Notifications()(*i352fab4a04ba762c64656855e65ab019a254014b3080f8077b4c95cb2a1bc817.NotificationsRequestBuilder) {
    return i352fab4a04ba762c64656855e65ab019a254014b3080f8077b4c95cb2a1bc817.NewNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.notifications.item collection
func (m *MeRequestBuilder) NotificationsById(id string)(*i4013f2adca2e80594707aa753bbc43cfa3c88b1f94557550a9de9a87010a65c9.NotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notification%2Did"] = id
    }
    return i4013f2adca2e80594707aa753bbc43cfa3c88b1f94557550a9de9a87010a65c9.NewNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Oauth2PermissionGrants the oauth2PermissionGrants property
func (m *MeRequestBuilder) Oauth2PermissionGrants()(*i07a4f67d8feaf18ceaf6ff3228c4e2151126ceafbbfc83b3ec78e7051ad29e3a.Oauth2PermissionGrantsRequestBuilder) {
    return i07a4f67d8feaf18ceaf6ff3228c4e2151126ceafbbfc83b3ec78e7051ad29e3a.NewOauth2PermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Oauth2PermissionGrantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.oauth2PermissionGrants.item collection
func (m *MeRequestBuilder) Oauth2PermissionGrantsById(id string)(*if7c78154b39b2da1e1c12f7bf63ddd85b353b2d31d33f3ba11e1dfd43b13ad7a.OAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = id
    }
    return if7c78154b39b2da1e1c12f7bf63ddd85b353b2d31d33f3ba11e1dfd43b13ad7a.NewOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *MeRequestBuilder) Onenote()(*i65eadd82fcbfd207b16c3e7dbf0c8d1067289ab718ba0ce591cfeea315d5d9cb.OnenoteRequestBuilder) {
    return i65eadd82fcbfd207b16c3e7dbf0c8d1067289ab718ba0ce591cfeea315d5d9cb.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetings the onlineMeetings property
func (m *MeRequestBuilder) OnlineMeetings()(*i73355d68106cd7f92ed5b9ffb4a52e59ba3ec00b272cb67ff6e7d1e181314937.OnlineMeetingsRequestBuilder) {
    return i73355d68106cd7f92ed5b9ffb4a52e59ba3ec00b272cb67ff6e7d1e181314937.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onlineMeetings.item collection
func (m *MeRequestBuilder) OnlineMeetingsById(id string)(*ie23caf2134464ac7a439e71ae9e04c43d1e7edeb9f60834c6450c46b2a73ef3d.OnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting%2Did"] = id
    }
    return ie23caf2134464ac7a439e71ae9e04c43d1e7edeb9f60834c6450c46b2a73ef3d.NewOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Outlook the outlook property
func (m *MeRequestBuilder) Outlook()(*i68ae3b4d592d0057eb4a687d378300648953a03d22862f907fd2ff66d201eeac.OutlookRequestBuilder) {
    return i68ae3b4d592d0057eb4a687d378300648953a03d22862f907fd2ff66d201eeac.NewOutlookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevices the ownedDevices property
func (m *MeRequestBuilder) OwnedDevices()(*ic93f01ea250d4839ef099dc4c41fc62d7f5aa62afb84e4edfd9fcdf0c69e48cd.OwnedDevicesRequestBuilder) {
    return ic93f01ea250d4839ef099dc4c41fc62d7f5aa62afb84e4edfd9fcdf0c69e48cd.NewOwnedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.ownedDevices.item collection
func (m *MeRequestBuilder) OwnedDevicesById(id string)(*ibdf6f337e272c2c489b46bc7469a81713a5f1141291912d55abe2a182b61324d.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return ibdf6f337e272c2c489b46bc7469a81713a5f1141291912d55abe2a182b61324d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OwnedObjects the ownedObjects property
func (m *MeRequestBuilder) OwnedObjects()(*iad3415fdc487951976de6996b1757be3243496c6f390b601c1a613f27ed8532c.OwnedObjectsRequestBuilder) {
    return iad3415fdc487951976de6996b1757be3243496c6f390b601c1a613f27ed8532c.NewOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnedObjectsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.ownedObjects.item collection
func (m *MeRequestBuilder) OwnedObjectsById(id string)(*i8ce653616e56cac134ff5b2659d920c1b77db0e8cb12f0164e596429c11be8d5.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i8ce653616e56cac134ff5b2659d920c1b77db0e8cb12f0164e596429c11be8d5.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the properties of a user object. Not all properties can be updated by Member or Guest users with their default permissions without Administrator roles. Compare member and guest default permissions to see properties they can manage.
func (m *MeRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PendingAccessReviewInstances the pendingAccessReviewInstances property
func (m *MeRequestBuilder) PendingAccessReviewInstances()(*ie92980fd5edaafba8f00572025a84bd9779e54cea477e030692babbc26d2d460.PendingAccessReviewInstancesRequestBuilder) {
    return ie92980fd5edaafba8f00572025a84bd9779e54cea477e030692babbc26d2d460.NewPendingAccessReviewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PendingAccessReviewInstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.pendingAccessReviewInstances.item collection
func (m *MeRequestBuilder) PendingAccessReviewInstancesById(id string)(*ifd2db3d8aa0a0ec27c89cb2949e6936c2d3de139fb986c5773dfe17008496fc0.AccessReviewInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessReviewInstance%2Did"] = id
    }
    return ifd2db3d8aa0a0ec27c89cb2949e6936c2d3de139fb986c5773dfe17008496fc0.NewAccessReviewInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// People the people property
func (m *MeRequestBuilder) People()(*i96b86fb89d12c60762a9b35297bd48c5f5a1f0a29a853fcaee1bcb8408c8a5a3.PeopleRequestBuilder) {
    return i96b86fb89d12c60762a9b35297bd48c5f5a1f0a29a853fcaee1bcb8408c8a5a3.NewPeopleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PeopleById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.people.item collection
func (m *MeRequestBuilder) PeopleById(id string)(*iedd9b42c61703c4f6442a904b683a6508ee4bf335825ff84ce57fb2616e2cce3.PersonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["person%2Did"] = id
    }
    return iedd9b42c61703c4f6442a904b683a6508ee4bf335825ff84ce57fb2616e2cce3.NewPersonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo the photo property
func (m *MeRequestBuilder) Photo()(*ifa6354c03b0a297cab93d13f8bc78708546aaeb1be40e137a7d07a262fce0ea9.PhotoRequestBuilder) {
    return ifa6354c03b0a297cab93d13f8bc78708546aaeb1be40e137a7d07a262fce0ea9.NewPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos the photos property
func (m *MeRequestBuilder) Photos()(*i3340a668307034bde84b6a65c934778fd43edb7307d82098a167906d463eedf3.PhotosRequestBuilder) {
    return i3340a668307034bde84b6a65c934778fd43edb7307d82098a167906d463eedf3.NewPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.photos.item collection
func (m *MeRequestBuilder) PhotosById(id string)(*i35b6d595de33f1d461b0d50276e049227be20e321f99ec0e204fb7be2adcec52.ProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return i35b6d595de33f1d461b0d50276e049227be20e321f99ec0e204fb7be2adcec52.NewProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner the planner property
func (m *MeRequestBuilder) Planner()(*i32d8dc857112f42b52a344de9badc89e743eef960325ebe8910ed2013ecc8d37.PlannerRequestBuilder) {
    return i32d8dc857112f42b52a344de9badc89e743eef960325ebe8910ed2013ecc8d37.NewPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Presence the presence property
func (m *MeRequestBuilder) Presence()(*i6237fa040b5eaf455abf5f698c389202ad776ba022b1566206439de15b5da115.PresenceRequestBuilder) {
    return i6237fa040b5eaf455abf5f698c389202ad776ba022b1566206439de15b5da115.NewPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Profile the profile property
func (m *MeRequestBuilder) Profile()(*i29d15cf2ebb01a69779caa08b86678cd70e064e98f4f681c20dbb2f7c5068dcc.ProfileRequestBuilder) {
    return i29d15cf2ebb01a69779caa08b86678cd70e064e98f4f681c20dbb2f7c5068dcc.NewProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevices the registeredDevices property
func (m *MeRequestBuilder) RegisteredDevices()(*i0615d7ba982e06a459c77c0876aee46ad012da5184e784858f5b99991e2b0442.RegisteredDevicesRequestBuilder) {
    return i0615d7ba982e06a459c77c0876aee46ad012da5184e784858f5b99991e2b0442.NewRegisteredDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.registeredDevices.item collection
func (m *MeRequestBuilder) RegisteredDevicesById(id string)(*i47e5b20f37978ff329fe5a3a058924b96b429a635613f9be0625a219ebbbf612.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i47e5b20f37978ff329fe5a3a058924b96b429a635613f9be0625a219ebbbf612.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReminderViewWithStartDateTimeWithEndDateTime provides operations to call the reminderView method.
func (m *MeRequestBuilder) ReminderViewWithStartDateTimeWithEndDateTime(endDateTime *string, startDateTime *string)(*icfe0fa0def53eb8ead53a9c7e14cfa7b99aa67d289bbd5786ae5577155574961.ReminderViewWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return icfe0fa0def53eb8ead53a9c7e14cfa7b99aa67d289bbd5786ae5577155574961.NewReminderViewWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime);
}
// RemoveAllDevicesFromManagement the removeAllDevicesFromManagement property
func (m *MeRequestBuilder) RemoveAllDevicesFromManagement()(*i772944f98e62be37df8bbf30068a248f0f1b2b5e7b1938cbf0e1c6fe680be3c8.RemoveAllDevicesFromManagementRequestBuilder) {
    return i772944f98e62be37df8bbf30068a248f0f1b2b5e7b1938cbf0e1c6fe680be3c8.NewRemoveAllDevicesFromManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReprocessLicenseAssignment the reprocessLicenseAssignment property
func (m *MeRequestBuilder) ReprocessLicenseAssignment()(*ic3fe384252788f8c3b5dfb7e9064552b08d0785f3f25b9abb64555be87652a0c.ReprocessLicenseAssignmentRequestBuilder) {
    return ic3fe384252788f8c3b5dfb7e9064552b08d0785f3f25b9abb64555be87652a0c.NewReprocessLicenseAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *MeRequestBuilder) Restore()(*ia675b038deca3e32dad85e01ba41f1c919b5d1c779f5258a2941327eb3bfc17e.RestoreRequestBuilder) {
    return ia675b038deca3e32dad85e01ba41f1c919b5d1c779f5258a2941327eb3bfc17e.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeSignInSessions the revokeSignInSessions property
func (m *MeRequestBuilder) RevokeSignInSessions()(*ib7801ea348e04a14861bd095693d71eb64b7e24fb1b05827984cd33d80abade1.RevokeSignInSessionsRequestBuilder) {
    return ib7801ea348e04a14861bd095693d71eb64b7e24fb1b05827984cd33d80abade1.NewRevokeSignInSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOf the scopedRoleMemberOf property
func (m *MeRequestBuilder) ScopedRoleMemberOf()(*i832752c7fc0bc2aa43f3e00911325be0e4d0bda70e3caea85e2f66813f2639de.ScopedRoleMemberOfRequestBuilder) {
    return i832752c7fc0bc2aa43f3e00911325be0e4d0bda70e3caea85e2f66813f2639de.NewScopedRoleMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScopedRoleMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.scopedRoleMemberOf.item collection
func (m *MeRequestBuilder) ScopedRoleMemberOfById(id string)(*i08f6400faf92b934d06b0b6861e5bf74d25a77d5b3d67c02b4779a1c0c52e960.ScopedRoleMembershipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["scopedRoleMembership%2Did"] = id
    }
    return i08f6400faf92b934d06b0b6861e5bf74d25a77d5b3d67c02b4779a1c0c52e960.NewScopedRoleMembershipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Security the security property
func (m *MeRequestBuilder) Security()(*if3a811497f1bf492f3bbe053e28833b1140d2c90725fd72ed4cd2e0b4068fbf0.SecurityRequestBuilder) {
    return if3a811497f1bf492f3bbe053e28833b1140d2c90725fd72ed4cd2e0b4068fbf0.NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendMail the sendMail property
func (m *MeRequestBuilder) SendMail()(*ib4c685509c6fd2907ec160446f20f296a4dbe123da549c4e6b9e7ad4ade1b374.SendMailRequestBuilder) {
    return ib4c685509c6fd2907ec160446f20f296a4dbe123da549c4e6b9e7ad4ade1b374.NewSendMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *MeRequestBuilder) Settings()(*iaa6383ba91d1d220deef98e9ce97a49a107e15b05eb6b2b3ce3c9d6c48e4f326.SettingsRequestBuilder) {
    return iaa6383ba91d1d220deef98e9ce97a49a107e15b05eb6b2b3ce3c9d6c48e4f326.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tasks the tasks property
func (m *MeRequestBuilder) Tasks()(*idd48ab55f5d1da1d01a7519debf68afdf64ff43e4e90500cda63eeb76402e355.TasksRequestBuilder) {
    return idd48ab55f5d1da1d01a7519debf68afdf64ff43e4e90500cda63eeb76402e355.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Teamwork the teamwork property
func (m *MeRequestBuilder) Teamwork()(*i7b613e1f229abaea577330e8430e25c30f3890e517f541d6c2fa0bc338c85ebe.TeamworkRequestBuilder) {
    return i7b613e1f229abaea577330e8430e25c30f3890e517f541d6c2fa0bc338c85ebe.NewTeamworkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Todo the todo property
func (m *MeRequestBuilder) Todo()(*ic675e3e79f073212aeb3e72b7a32c8529405e4fe1cd4bd69b885ace5326ca895.TodoRequestBuilder) {
    return ic675e3e79f073212aeb3e72b7a32c8529405e4fe1cd4bd69b885ace5326ca895.NewTodoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
func (m *MeRequestBuilder) TransitiveMemberOf()(*i7c69216b282055d4e66d8afba1065734bd28d022f420d5bbcb72d7908ab5be9b.TransitiveMemberOfRequestBuilder) {
    return i7c69216b282055d4e66d8afba1065734bd28d022f420d5bbcb72d7908ab5be9b.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.transitiveMemberOf.item collection
func (m *MeRequestBuilder) TransitiveMemberOfById(id string)(*i4b4dce57580b9d6f6781b551395b4c975ce6835ca4860fc6b89a5b72ff02daca.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i4b4dce57580b9d6f6781b551395b4c975ce6835ca4860fc6b89a5b72ff02daca.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveReports the transitiveReports property
func (m *MeRequestBuilder) TransitiveReports()(*i59bc57755ba56f6eb5fb2534c9eb2c79f00639d8861695b6b49b5e54cb223606.TransitiveReportsRequestBuilder) {
    return i59bc57755ba56f6eb5fb2534c9eb2c79f00639d8861695b6b49b5e54cb223606.NewTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.transitiveReports.item collection
func (m *MeRequestBuilder) TransitiveReportsById(id string)(*i2cc8c08facade07f4bb7f43d3d598d3e49ab04058fe12463d7fbae54f4163f48.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i2cc8c08facade07f4bb7f43d3d598d3e49ab04058fe12463d7fbae54f4163f48.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TranslateExchangeIds the translateExchangeIds property
func (m *MeRequestBuilder) TranslateExchangeIds()(*i477445cf5f8c496c3730b4a0183177997d93184d877ac862e8af6b8963a3985d.TranslateExchangeIdsRequestBuilder) {
    return i477445cf5f8c496c3730b4a0183177997d93184d877ac862e8af6b8963a3985d.NewTranslateExchangeIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnblockManagedApps the unblockManagedApps property
func (m *MeRequestBuilder) UnblockManagedApps()(*i292db412d9ce5ec2a850c1c6b4baeedec3e63413f2b63ac27c792b9f269d30ec.UnblockManagedAppsRequestBuilder) {
    return i292db412d9ce5ec2a850c1c6b4baeedec3e63413f2b63ac27c792b9f269d30ec.NewUnblockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRights the usageRights property
func (m *MeRequestBuilder) UsageRights()(*ib5b4d50478937130cb8d53f4c75d274229e4c2a9061b1c9a485907eed04f6aef.UsageRightsRequestBuilder) {
    return ib5b4d50478937130cb8d53f4c75d274229e4c2a9061b1c9a485907eed04f6aef.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.usageRights.item collection
func (m *MeRequestBuilder) UsageRightsById(id string)(*i634abce1c174c1a4b8f987a99848dda449764fbd35f786e01464e876d22d956f.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return i634abce1c174c1a4b8f987a99848dda449764fbd35f786e01464e876d22d956f.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrations the windowsInformationProtectionDeviceRegistrations property
func (m *MeRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*i9707f15743f2bc7fd08620852e99c80974fc9e9ce7762c2b1f6e4183508d9946.WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return i9707f15743f2bc7fd08620852e99c80974fc9e9ce7762c2b1f6e4183508d9946.NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.windowsInformationProtectionDeviceRegistrations.item collection
func (m *MeRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*i32120ccc8322d0ae4c79827cab0e21691e8b625292351396f2ee7ac9681caaed.WindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = id
    }
    return i32120ccc8322d0ae4c79827cab0e21691e8b625292351396f2ee7ac9681caaed.NewWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WipeAndBlockManagedApps the wipeAndBlockManagedApps property
func (m *MeRequestBuilder) WipeAndBlockManagedApps()(*i03a3fdcb3f1d65e925013bb34fbf99663faa80c904b6f351d0e6428219666962.WipeAndBlockManagedAppsRequestBuilder) {
    return i03a3fdcb3f1d65e925013bb34fbf99663faa80c904b6f351d0e6428219666962.NewWipeAndBlockManagedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationByDeviceTag the wipeManagedAppRegistrationByDeviceTag property
func (m *MeRequestBuilder) WipeManagedAppRegistrationByDeviceTag()(*id9569f521a060e524166adc08c72508f54ed83854d63694ef777241e7b5c7d54.WipeManagedAppRegistrationByDeviceTagRequestBuilder) {
    return id9569f521a060e524166adc08c72508f54ed83854d63694ef777241e7b5c7d54.NewWipeManagedAppRegistrationByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByAzureAdDeviceId the wipeManagedAppRegistrationsByAzureAdDeviceId property
func (m *MeRequestBuilder) WipeManagedAppRegistrationsByAzureAdDeviceId()(*i7724a6079225e5b4fb109ecbcdfd2d5f8d5f2e78a512dff130978a6bd7ee0aa5.WipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilder) {
    return i7724a6079225e5b4fb109ecbcdfd2d5f8d5f2e78a512dff130978a6bd7ee0aa5.NewWipeManagedAppRegistrationsByAzureAdDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WipeManagedAppRegistrationsByDeviceTag the wipeManagedAppRegistrationsByDeviceTag property
func (m *MeRequestBuilder) WipeManagedAppRegistrationsByDeviceTag()(*ia96d1f92e78f5de8132d5ce2d35b245bc47a307071b141447e1003202987f9f7.WipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    return ia96d1f92e78f5de8132d5ce2d35b245bc47a307071b141447e1003202987f9f7.NewWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
