package functions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0050d32d11128323eba61e1e75eb83c5547c444e23b016815365cda7028c0797 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/tbilleq"
    i00bd7d5caf4f643e4d966eb8a89a829df9dd5b0dbee54b766d2097ce7036aa2d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rows"
    i01096f2b34ec8a2f8a1148bdb674f5bae59ae8e931d384bab99065956d16489d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/factdouble"
    i014f5fef465cec5d69bfe6228aa49769a28be347b59b869cc6d201c421c5fc55 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/concatenate"
    i01ade93e966b7ef3ad3978d2f1d90704ba31d11a2d416ea326f25b442162d957 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/min"
    i01d278483174ac301bcbeac39516b56e97a8a5f1352e96fe34181f0ff94b1ee0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sec"
    i02e07df104c71232444b95fc5b1e62f0293d740e0b69b73e13add9e338bb0d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/nominal"
    i0453e21d45f8ed2aab3a312355a9cdffae8562a8747da1d10b0964833762a270 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/if_escaped"
    i0494ccdd128ce2f5c8821e6a849105650c4b4c0bee4f2530ceb6fbdeb4e8c7c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/stdevpa"
    i04b1ad703d7e9b5de1347bbe2e166dc427117fc2d5694354e996a14506873246 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pmt"
    i04bc21ac86d9392376c6b121e63d25df0567e9a3fc8da2e91db63efdd803c37b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/vara"
    i051f24a077b7189aa242fd3fe1798aebf0b3261717546eba2dc8be57c7d50c0d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oct2bin"
    i05bb4ed647eed6938ebc562b6011f32153484483264b9f46dd53577d8bb071f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sinh"
    i084289164bd03459fb7e41d139af885ebcd76e0cd01fb224a5df9263ede20802 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/year"
    i09a091e227ad1cadaf693bfcf5cabc15f2a578b85013c5809d603aed05bc797f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsqrt"
    i09c5c179d03bde65cb63adabd7bd4fbba135050439becd052c9e02d98ccc445d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/second"
    i0b6d1a4ef16cfd3065c4c6c4aa740a8f4ff77019773eb7e384205e813ce1468b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fvschedule"
    i0bbe7782eee3501d57bb814a1a1c58ce8cf07efaa239ed9577f240d94d413bc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oddfyield"
    i0d4bf14b866d57e7379422ab460cef0187cd905710c4fb759a2015829d01f655 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/delta"
    i0e88fa40bea4437ac83388050ce7e5ce9de47dc2f2ce412db5125c807029050d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/erf"
    i114f33a34a95c8aeeb6f01a515b9f584c007b96d3e8469f93ccaf4f74d1473ac "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/vdb"
    i11635e1f1864a4c8e9365eca479c56e1761cc4c1cb596a971d3e34a740a3c085 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/datevalue"
    i11a43b6060df9542091805a286115ebb955c1aa0d49662e197fdf6a4b2e50e4c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imaginary"
    i11c7d19b25404b514e263be9547aef03f560afcfcde0489c7c332825f429122b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bin2dec"
    i11d6df0ff757addbcc18e3f5e5696d91da8693a9560cbb682fa52b8f4539be2c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/besseli"
    i1240f098751c9c3ede3953da47ed5f1592e3bcdea4a678ff30bac2fa804160d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/intrate"
    i1325a7cc878f353726715d2b70625d31aa9f68c17022b85fd116b3ec2e2815ca "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/tan"
    i146daa43f721553e08c4e14d4fe76761df9364bd7bd6b5239d7f796913654307 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/small"
    i159670de4e02e4a74091f861f99a99ca141f71ec0b87a11f48af8d33b4829654 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/exact"
    i17c08c29b663ddc5369e52b456d6e68dba92c794e5a5b6ddbc0e076bfbd015f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imdiv"
    i18f1cdd18e81e750ed450d7b6aab0e82e96eff3d5624dc029aafd78d132463b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ecma_ceiling"
    i19464df7609672eb389ea5c0b7823da82b500d02a65e21c56836a21fdda12a5e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/norm_s_dist"
    i19b00c3ca966260792b57fd09e773e9ce9e699138711859ca47c20f1a9e34b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/acoth"
    i19b74ab0a3dd49a1976bab9450d4ef387991450a9dd1f439a0f11e82b1287321 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t_dist_2t"
    i1a280cdde157745d3f87a26d747c812b1665bc8b66c0fa4890d28140984b06e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/cot"
    i1a5c221c73ce54ab21f5cd625811db70165d0058939149337b4464d4fb657cfa "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ispmt"
    i1b1e5f6c5fe8a56a140323e96bcda3645d409917a438e57ec2f42590bfbd0bf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/days360"
    i1baed36ad116a1d420e1a7d605ebfacb8a1686fd4c1db3e10ed48bf62c9e39c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fisherinv"
    i1d9c84764ea0a52054b0a1ff763c3e328e31f451bbb49c623958cb42d31d56df "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rounddown"
    i1dde59714c4439bf7a5432f226d87b56892b824dae30192790cec46f5fa75ed6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/amordegrc"
    i1de8619dd7518db3d99462880cfbf87a747b0bde4d139b30efb92a2897d26416 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rank_avg"
    i1ed5a81b94a902275caf64d5646aa6a5d64df5ee721b688f4945f5422aec8c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/accrint"
    i2193ffa9da564442c3818c80c3a3934f5f1ab209ed23eaa6af49b0fca52698ff "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rri"
    i225dd267832dae6eab68b4f07ab1629f61b9ae136b37950bc41a4eb2121048ff "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rept"
    i231393df28bed923560304a9df79e689e7d63cdf353ea10aa4965010bbb27b59 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/harmean"
    i239cf5ab85067bafc56c3ea278efdda656968240fda3d73eb82eda937e180bdb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/cumipmt"
    i23d1cf8a00fdb2931404bd1ff8b0d447a926d8d3874700e2553064cf8c0d13e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/phi"
    i25279acfc780c5710e2ac0bd05598cb053419123f87b0e9b3f2cc83f197024c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/binom_dist_range"
    i2544fd0935f5ac2f12f79b4a02a43b030e17165c43e65638f812efe9c7d29892 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/yielddisc"
    i271789ad8387c2c7897caaa1d23fdb3f98f1e3cf2f792b2a03f62a7a586d9df3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/acos"
    i27a16833edf596f6be428ee862a5d9faf3ff3a73862d538865760670037b2004 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fv"
    i29da91cf8167cc7cfc9d8a229b680f1dfed88338c9faeb9401a09aacf7ad5ace "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/columns"
    i2b92e5915b7bdf7e95d7baea79b7488f604113d133d1d5a4afebacf7f0a4b470 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/time"
    i2b962df7680ae6985d6808308b7cb3aeecdbe628568137708567ef5ed78c52be "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/xnpv"
    i2be0b0806f84053729758b3570db185ad402da1d3ffd988392771290ec7d78d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/chisq_inv"
    i2db9198ed6278fd3b0fa28899357367386848879a0248c9bbb758b03bb05cdf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/norm_s_inv"
    i2e9731503ea02b7ef4747c359c541a15c2b4c70bce0cc4ef82d5c00c080a72cf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oddlprice"
    i2f1e41b93eba14f69f07e64f451b07a9d162c3b2b5ac0833f310b56e7dff93ce "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/power"
    i2f884c270bdc5f79a1b5c6d1fd764a4097f9f98cc08bcab6c13418d844f3a6f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/base"
    i2fe518f7ce30627483c22947ac5fae593cbb1137e090d90d29067f66da898e01 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/f_dist_rt"
    i3064bbf2a3c2deacd05d96d8e11b79f271f609d4a294f5cae347ed3c910c0696 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imlog10"
    i3122bd5b6a80eb07305ce98d54e54b51fb400d53e589e1c9177c2dca1b0e498c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hlookup"
    i317145d7cb1b1dd1c7a24b6b155a43af039f3cdaafcab3a3758411b12c50f560 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gcd"
    i31b6a6d9d88e1a77eb013a0a4f84b7c8f4f0e1487cd1befa282da7c3a5467b16 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coth"
    i33ef7455cd3354f695f10748ea17e47e4d0622a2a93d13f279966caa3f74e31c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/averageifs"
    i34dc66540bc7342bc45b426f0e1fedcbafa46564955a9848253415d4ac46b4d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isnontext"
    i34f2a4a24d744d3fc6016654b9341e19775b8dea258192b301b9d8cf4e72d763 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ceiling_precise"
    i3516812fdea5289d5c1239ff45469b8cd752ae305dd572fc2315b1a817738e3d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dec2hex"
    i3723d97baf5744837f171bd68f52e502d01d70a165e007898cd9489339572ae4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/erfc_precise"
    i3728955a08a79c9a0b07d504a99948f05e706d80b3740b69a04dbf94261e8dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/yieldmat"
    i37b034b37e59ca49f27a237e569deee9dd6f9020957c697c75489bbc106b3ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/roman"
    i37cb723b274316a2f034aa8b2a909f4d75dfe2fdcbcd35a4fdb5a661e02d602a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/value"
    i38da5a55b58b6bd8bc84fa109d38ad5ba55b676692f6ebb60a25def11141d862 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/atan2"
    i392b51c7d8691d97c55045a593139b9654ca512cf8ca6940e3716a0765231661 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/combina"
    i3a5b50be37d36553bf1119b198b458ef940eeda87d384cfb98fb6b76bcc3ea70 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/choose"
    i3a713ed3e471d758071d6320615cb630620a44eda7882e59205ee608b16ab8eb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dcounta"
    i3adc2f025c48fd250c29a04aa202729a054a29d0502ccf4ea72d4ec5a64635de "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fixed"
    i3b37f1d056cb7cf047c630f85e8d49d51384a42baa20b671b8f98d2d440feba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/avedev"
    i3b57e79beeef2f08027bdf9fd759cc0858a70f39d4a35a6c44414778563012f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/confidence_t"
    i3b6a93a263fef9b643219a486ba14d1a90efceba62eaddbec90cfa7ea241dcce "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/erf_precise"
    i3b889c48e56bbf6a183fe19f9853c22829238e3a7d488ba07cae69afed667f9f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fisher"
    i3baa4f1f70089ce805f6a5096e8b3a29c96246baf3605f030ef2951df8e61cdd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rank_eq"
    i3be8431b385db6978cd20fc820abcefaaa7a43b0bd9794564078d5be763a9937 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imconjugate"
    i3c257e281eaba9df97dc69b692bc1e7a702a67ef1f4d7673544a162df195a9c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sln"
    i3c61aac67e71e26228c623aa235d91a3be22189765dd40077b181ba6aa6ac074 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/error_type"
    i3c6e2bd2b92cb13f1e949c31c034ca7519729cbd902289447d8fb1f45d82ba7f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/asin"
    i3cbb023d286ddbc3ab8433a36661c912b2091cf8e12c0464bf56e7fddaa70c16 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gestep"
    i3e27b43c227d8f50f706ff6461aa17d93a7add1a68aaf7a5a7a31b13a5d05d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lower"
    i3e7338d395a2213f344bce5f8bd106c48c9613368c44472b4a68c636a47267e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mid"
    i3e83145b3125a8eb2f1a717e5d6d6f44f41b1443ad97bb094bc03f0a892b75db "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bitand"
    i3ed3d358b4dce2bb6a00a026fbf5a2343557dce788227b550f1f5ccbd12b1615 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coupncd"
    i3f458892db0406f566cc4ae0bfc1ec462432e8d9275bb6e6cd964aaf774a1e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rightb"
    i3f56cdb4a69ef320a75d1586206b7dfb437b5a604fce57a5019673319644e355 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/unichar"
    i405aebdde6e9c90bf532ba8258294f2f8162cedfeccf8c83ca2d1e4abea16a31 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/percentrank_inc"
    i41988a8388286f2ac1bbc4d6af302f9d1f081f6b63ee9e8fa3682d84c8fc86c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/int"
    i41d373dba62b6f70cc8ef9047fdafa598a1a6fd341eb627992cd1d45bbfce880 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t"
    i44c5b3ce8f085350158ab68112995ed6e115d0a6ee80c6b9d8d74e39fdd4f26c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gamma"
    i45864aa1e5af9a8d5fef6d7689757c11aa4f550e8509bd22356840584e8a4f74 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oddfprice"
    i45a317c4a16779aa0191decfdd8eafccaee07defa32b4f71d86074431bc29a0a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mina"
    i4748eaa62f242c480beca098f37730bd85581f0454d96962ab77c5bcc4077c53 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/atan"
    i486d88d3ffdf365f62b4928cda0b21c0c76426936a8b8ac8ccf19250c5debac9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/workday_intl"
    i48c062699c214b911f9435316ac4c100152635b2cee12fc5cb6ea644b079214f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sheet"
    i4ab58e6b08f32389c6302b37da6352ef4e50c7099f63e72b1eec5408402a3b25 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mduration"
    i4ac7e344d13c36c2e3c23d87816150a649beaf2012302f6610653444ecb9b234 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/weibull_dist"
    i4be11536eae45fd4ea7adf5d4299b71cd862146faa910c723899f6911f0552a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/daverage"
    i4c51dd2f6f6f5160da6053cc1291ae4a43ad7af4dbef329a7b3f7fa4d282750d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imcsch"
    i4d5acd27fec5eb7998fbc0ba0103be8eb685259f312784459dcdc8791fbd4ab9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/percentrank_exc"
    i4e3aa948ee53f2131c385f041992741d89d3bca88109722ca599726083d91928 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/islogical"
    i4eb6a2fb4e8f255bbffdc6e5b9a7f2ded94bce657445c460aadb52f5b34458ac "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/clean"
    i4f07a98ba1b3a8f551935097937b9a4cd8b34d18c75ac58ac182fd2835ed51b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/skew"
    i4f9aa0e6ed219c6f029aab669ac0680c2fceffd3c383d5478af4b9fddd950513 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imln"
    i4fb5a0beaa11f0d431ffe7b2776a63f1519b2b0730f7b34c1557c95dce210d0f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/chisq_dist"
    i5102dc4301948efe9759bb91a61ce696b9a7eae3cdfa559114a7b37dd8074776 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dollarde"
    i513653df7d3cb928e656cb4c60e76d7c5355f0856052209394f8da8b64e5e225 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/numbervalue"
    i51e821e10b850a9c97fc458dc827715fe679e8ba556de07fcb94e8afcf7924b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/poisson_dist"
    i52ae3aaa09a8851ecd5a25513977e69abde753c7b68342bc599ec5609538490e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imcosh"
    i5376958355630f6f6240607052fdce472946f9e94ea3e8560aa27167dae3ef3f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/randbetween"
    i539158791087b6399c0cb327f13cb7790c9745fa153c61e552c01b787c91efbf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/beta_dist"
    i544cc180c4727e4718f970d997bd828304c5c469cb353a04eb4136d9002fc7d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/received"
    i55563060001cd02d8ec59c905f1cbcdacb601fa67184d73a2799ca52783d5423 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/minute"
    i57abd1bbb332b1fb6f15dc1b9f5d038188388426d82da6b3ae052ab3713f7794 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pi"
    i5873f11c8e7d2e163103011598914ed1291b59404c3dffe95bdfea5a9837dc6b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/networkdays_intl"
    i58fe0cf7e33ef2cc42c1aa9ddecea577e48b03f4cdf37d375591026954f46c7c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/upper"
    i5929444bf113c70d20e2d4295b6983a9724d0227d6c4d438f913f22f07823801 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dollarfr"
    i5b2286301b4d4e486cb0f56c4227fcfef09387503478e065a21767ef0f527cfa "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dget"
    i5c04cac2d711234a9025c8e14e1c9fddeb0af9d850a46fc99376ee1f28840835 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/weeknum"
    i5c087781b52d9dcd7ec53ed914b813a5ad2a143bb9b4d06a86ae1d23917eb3f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/max"
    i5c3d49a1ba7296574afa46cd49edc3ad92da6e76655ea5a04f1913bcbf28aa69 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/irr"
    i5d1af52be2b3187ec0e03be085c84d7655295a9764e715c8352272c991a5b31a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/asc"
    i5d6b82b7c343af5d6b219c51afe859aee013f35c73c2e618180f18bfda38f7a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/or"
    i5e2ca6a19b1a68e98eeb5a22cdcb32b257cbd898afc03c740da42dbdaa146731 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/chisq_inv_rt"
    i5e90c8df3a37e03a7f9daa30813e1a9a6651dbf7a729a7284aafa183e84e909d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sqrtpi"
    i5f73d2d1a7d4781f4525b034525e9df3deb205e8e4ac758982e431dc8eb9e284 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gammaln"
    i60ef136bd90310658d302308c48fea77adcc95e0242c2f24f08b9ac1f74470bc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bin2oct"
    i6170b518e26aeac41ed27c586fcc4d593624a8dfc1898f24e26195a1b08adea9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/chisq_dist_rt"
    i61d4785a8acc4838fe5a5dc4bb2982e512218f586c0c26eb046d65eb08247df0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/false"
    i6421c408950fba8073283495764103519c7b83b532bec18bb40598c0dbb1dce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oct2hex"
    i64400f50c71ff8b2b271f1ceae5f9e839b64ae25073adf239505b644be0b5bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/devsq"
    i64f90fa888a48017c96edf4e2d4540dc46e262036416338183cd5fe46e76fa5f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/leftb"
    i6503d7704ccbe9c818734b31eb667f6181eacf1a7a022cd7f58f363dcf403849 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/quartile_exc"
    i6549418b40dc32ee3c2bda4721e2013b43682373bdc47a7e919d206cfd6bc36c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsub"
    i65fb0cdaf5c4b3936c8b0cfcc9c69c364c6f66cbf16cfac1a196908e8e0af0e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lognorm_inv"
    i67edd6c2573b42e032248aa937dbe9903d23f394d8ec053e7f97a11f2b38c3ea "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/percentile_inc"
    i68aeb1985b4e03322a4c461c36962d4d5d0c61538152110f09a942acaac1394a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/replaceb"
    i69c2649aba19ccdd2dd0f5bf255ddf8d0aa623a6c022229dbd77634b2c88bed8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/yield"
    i6a61af20e5713d0e4714cb59cf2e4d3c54a8491ed03cc4eef0303857d4c09194 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/cos"
    i6a81c49a0cbd281c42b4e3a543ce420c090faa981e28dac41ae63e33eddc10a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imreal"
    i6aa6e261dc646277372246f6ec53c74cf1d75e098b9e0f81a7f87dda1919e74e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isformula"
    i6aacbe20dc6e50258c9a6e8d391b05d4e7c7f96c9d960dcf3afd0ac5c111cda2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/csc"
    i6af5ea8a5eae6ce0d467e8d79bd8a793918c81b05535fac4d218577b2f1033c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/norm_dist"
    i6bd5f7f5f528d90cb5622517509fe271b78cb6337ca6bc94db79462be9bac8e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/product"
    i6c3e7e3f69674bcdedacded58afe5875ffcee90414e278f5b8841357e5d64df7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/percentile_exc"
    i6d8676c0302c91f80ee45dee8c2947ffc0870c9b7bfaa2dead2be59fa069aa05 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/subtotal"
    i6d9183c34f2b76db26c0d25b436c794deed367981ac978d37de88c2166b7aec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/left"
    i6f940147c420e91e85cb05f8f948d0d2e450e4a2783950aa886229ff61a8189c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/timevalue"
    i7063fea931c4bd333bf58cae4ae95006d880ae51dd69347fcd0112ce942d6e56 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/true"
    i71085b17f9c510fc838d37422083076af4290773fdac918f19c26e7aa8cfb541 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/asinh"
    i71bb5e68a9c6004f9feea3f5b87729cdae98a1ab46518322bb458595067046f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gamma_dist"
    i728699d88374e91f88187e481684864d04eab0cea4a217e8854f8b3a2723f232 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imargument"
    i7322c133f56d7c2d62c17303e5865ee7fbe2b4b417b832ab2cd4cc88a2be8ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hex2oct"
    i737c5d48df084950792e74f03a89bbffee5cbd3676f6378d4a47838452db791c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/z_test"
    i73debe587dffbb4f4db9d9bed9a0d4d4fecc4b1e819504a326cb68d7a674fd8c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sech"
    i74c8bcbbb2397de63f9e8cbdf6652c485da2c060b6c26fc01b4870ff2660609b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dmax"
    i750e48618ea35ce290e17ae4e72d7b69de05921c0c97083a399b25eb35ce6924 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/accrintm"
    i77fc9cd08813563fe05c08fb8c1d0c196d2acc8c15e091cbff3713a5326d51f8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/kurt"
    i782c59c275368a34292a362686f23a3045f8e4b667c49f94302584781b7ab08b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/abs"
    i7aee590138e19a1b27ad74be0e89a72f68c0974cb718a9e4486336810b4e6d85 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bitor"
    i7afa6354755ab604b2000c207aaa310d7f80ce2d8ac7ad23a797c12b57a4bf37 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/countif"
    i7b3b5d10f9dddfffc344d3ea44785438bf332088789e8512860cbec2800a0409 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mirr"
    i7b45dd95f5c1244495c5095b798515bb00c86ca2ea5024d3861f84b6d993c596 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ipmt"
    i7b57475004e70295b93ef2615abeb481c477997c718c386d064208924ccfc6a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/maxa"
    i7c36657c2cef60c6b31bd5fc09db6d2a130306d77e26c1d597d8361db50f0e59 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/round"
    i7c6ca0f53af8043bc128eda59012a093b72c668fd62d115a64b3d24d8478fcf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/eomonth"
    i7c89957eea7f5d1a9ead132e515642f958a7d519b693bc9aa4a5168ff3d6a7ee "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/trim"
    i7cb7ca6d53631bf58f2266e1292e6c6331539af3696b718238bad702aaa8480d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ceiling_math"
    i7cc44ae074d8c78d8b236dfd967fb7230ff47ed7d77c49554bd34d6b20435c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dvarp"
    i7f104846fcb71d8d34066552109bb9ebb1921483ac42b4c84d9622742e45c3a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coupdays"
    i7f8ec93114a050419ca8698a92c1ed22b189436b0d42e9c5451c7abaa3dd95eb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pduration"
    i805b9eb159e70babbb463fdade294b172fd8126ddea9b57a1899afa3be216b61 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/nper"
    i80a969686b2a16a658360a90a8280dface60a263cccd4ac61782a7ab3aa075c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imcos"
    i80cae408a9fb23c0b3d5973b3b36036f8f86da13604a2df9e4f5d82335af9b89 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/var_p"
    i81f182730936652d80d21614c48a4de559dd8e4a172aef9bd71b56f481f09e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/odd"
    i822d92a8a540ae6b23416a1ba6edd6eff35f4abed7953e4928fbd2b4ee76d341 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/type_escaped"
    i8263a2a05c5a495103660fe69cb0c9bd613de25e80e8b7679cf8eaf7fcfdb389 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/erfc"
    i82ad818b67403f8f447c71dc2a945eaf3946b28652d715954fd417474a3bb758 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/log10"
    i82ed0aa8211f2c0cb3c379fcec65435bc44073411288d8ff080fdcb136272548 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/proper"
    i834687913aeb2dfb3df0401881633d320023843d6c09e5d88ce93d598db4e98f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/binom_dist"
    i857d6144810367f6d7cd741fdc5a24ee633e1241ee716fc3c34ec46d66cd353a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mround"
    i861da83d39292b36ce13f5c2b964c9924cc966b749fc25d2b5b6a669d83dd8c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/couppcd"
    i86b37cd528c9b0fbbd43b89e4026ed456e78a6b499a1c8b89d4b4bb72448a3d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/floor_precise"
    i87768552c7fd50ad36dadf66d75df68add33f80dca108d619208d0955a2311ea "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imlog2"
    i8942bff90eda23d84d12cbc5b1bcd0f5ad61c132f149f9345ec5725a177ebf14 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/arabic"
    i8a12a2b5f264a7c74820e9b113e4a38aed7a746d39e9da6e8d180db215e3c28c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/disc"
    i8bcb057ec8007b84ba78097f278d3988968527d460ef14fb3a3150960e187052 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dstdev"
    i8dbc7a210e86514658b6b7b58e46f4977dd33fe20198e1a8cf092493e3ef7122 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t_inv_2t"
    i8dce8a2f51873024131884e899491f95d9d6c9304be7c61361d0f9d33abfb188 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imexp"
    i8f4c6eeb4303ba527f8f345e5726779c14970ba71a2280bc32aa7301236ad2e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bitlshift"
    i9000ffdfcb17d10e09832a5fda64717cb5f77f9ba0daa420e6680c0bf9b50128 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/count"
    i90d52a146df567b4dd6dfdabe26ec25d0096a30b2b054edbc0a038e81515513a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/quartile_inc"
    i91d54d1b34c20ed6b9e3d724abea350e57457e087e1724b981a173dcf2863cc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/date"
    i9205761ddefe6680928f97b76a8e0de932c39501da2b6760101861f096bbb281 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dmin"
    i924df40fd16c8928384201c28d4fb809abdfbba3aeec4eae55afc75ec94685d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/effect"
    i9272ce275af1b9564f11a7d1e165b408dc842cf4d755acf8576501d96eef63a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/exp"
    i92fb305ff14f82f0506ecb4f2acbf0653dc257c9cbb95fb04bce2b4ce1fd900e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/radians"
    i9400ff8a2db8c4ce9ec4300a6fba694671fe73b126c7864b06515a4f329c745b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/decimal"
    i9403886d8363871971936746538ecefa7791fc349af497df5bc1277c5783d44d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lenb"
    i9591d406740b7642ef993df1ac69b6e39bf409d7daee05f3abc58402bda7e28f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sumif"
    i96228adffa7f7c2c4fabef938340b0e347e15559d0e63f7b36c6874502933832 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/besselk"
    i978242c1185a2557dcea381fa50ed6e2710ddbf3163a20d5e785bc7363a70154 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsum"
    i97c850bd34ef5bb4796243f0cd6452958a6816910bcb04c5fef1ebbb05cf8ee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/day"
    i99d62393032ec4efdcaacf249e0301202697d7b6fba7914bf39098553f84803b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/norm_inv"
    i9a186f78fe48f3a957442b1ed17fccca4928be511af537fbf762cc585b3be753 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hour"
    i9ef2fd435a266fea3206c285815b97a69f2a30b0e9e2358d804969a866eb337c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/countblank"
    i9f86a18034de4b30e9787efe02f4e0a4284676cd1f4401006517a31aa55ffa6b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/acot"
    ia07dbb8476d456306a1137e9feae6f70a68fc248ed129f143d15460b2129f4e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/xirr"
    ia10e03feb00942388c2542ed5b8c978c7e5851279ff4dedaf8eef4dec3e83b3e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/iserr"
    ia1959188eafd4bcd478ecf33a6287ceea78edb264b17875aac2ffe130ca69250 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/degrees"
    ia1e08a8c5ea700baa7a7adb97501abab922a41f60100dd6d938863a562bc973f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/istext"
    ia41376203552bcfaba0bc85c9235047123337e891d1c22d7936e78d36d871ed9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oddlyield"
    ia489b4341b29bd08b05402f5a65f8ecd4424d329b3dbf6daec1cc94e0ba7569f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/db"
    ia591b7807d5cbcdaf4b604db80322f5bcf8895f8f6566a04ba9b89dfaf1caf3f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/permutationa"
    ia59debb70c35806b26ac204748774dcb3d74fce9e80e89b0347efe6d56c2bbc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/workday"
    ia6288640436d5fa51554e6bb7265db0ae494d329885ecd45cd9de18927e25160 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lognorm_dist"
    ia73960e7b4da95c25edd07615a5587c36efa222e28463006ebf38230600502aa "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/geomean"
    ia874a740cc8fc9d80c5d38f6d847606bc8805a7749b782542042db71a5c6a1eb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dec2bin"
    iaaf4f922056c865e9b5213ff8a651682b31518b0bd7b7475f5f133d6fe68e71b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bitxor"
    iab13452fe1cf2fbff7b3ef1f2840e01fd1ecd0a0acdd8ebea3c83698c958e723 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/convert"
    iab2d32928daa4a2afac27ad11c4719622a434897b5d5271728dc6194c74b327e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dvar"
    iabaec7f58e32d0ba15b2f608134b387ffc3dbf171ae18d2d78f0fdb2603db827 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/trimmean"
    iac41516568b082ded7ba5277f71216c2d4b48c3497906e5bed96d890ad7c4994 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t_dist_rt"
    iacd13e9f55fa4756e8bb88cac5e77f31e49dca8e18f11d86e62e585de0a4b76a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/xor"
    iae7b528f9d9de970748d7d77bcefdfd78cd0d575b555bb4ed9f9bdead7e1ada8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/month"
    iaf0f35ff858eeae62b0dc3e885a456c3cb25f0ad94b3d748b8e25f2832546481 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/iserror"
    iaf6aad8ff5f9f5a94542995fec8311088ffaf94e57ba86fefad541089ed2a5e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/replace"
    iafdeaffeb4387b6753161d084bb5c7369fbcbd19961da7f33df1c37fe2c57c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dbcs"
    ib147fed6b5db44580c7acf996a2806f0bb8515459362219dd8cade718c06172f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/right"
    ib21b0bfa06713e745699452b9806b6e49faafc31f2098f6348b6e710d4dcfadb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/duration"
    ib27c742fbbd687c6ead89d689434bf79c2ca3ff53b03369668d84c142ddd5359 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pricedisc"
    ib37ba68a5fab1c602fae7e488eea5313b78cd71fa7132a748d7b9402190fbe46 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/code"
    ib455b5a90ec038d50df092019cb07bc6983d528925397d7bc0fb1a3acebf820d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ln"
    ib4640205a5ee55bc3b8a20bcbe1259f213daf2b26928179da4ba1125ad3d3c95 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/beta_inv"
    ib479884246004b472c3710d8868d87ab00d93948463b7e5be5a8dc6329641309 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/mod"
    ib4f820ed510f8f1aa69d6b111c66898e64509ab1648adb13b7a2718624d00f7c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/averagea"
    ib5af4e9a26a97744aea932645181a72915add2b5f542a4712f8a88e0fb711a6a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/trunc"
    ib6796fab19a8e38b93a3039b8faeebbb4bb49035349deb590caa195682ab55b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hyperlink"
    ib8172c174b7898498cd0e4fac1ea4c3847990fe6f96023962f3ed4c8032776f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dsum"
    ib8ba910ad32637cc152e0085ee699ee1cc26b29b4001fd5e5f6a368de40c4228 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/counta"
    ib9d1f9ecf825ea1ba386243b2d81f7d6f2e5e0d078ecd6ab556ae1176298b85a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/amorlinc"
    ib9d47bc5a4b46376eb93154925b463120b0c1b160682670d4ddb93d6001ad874 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bahttext"
    ib9f89317bd5c3fd7fe775dbcdda3b3a0049814dbf7a1cf903cd14e5384478033 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coupdaysnc"
    iba9a57b0cf80608e77654b3a94673f296c392fd17d9c2605fad518d6f8c4d18c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/averageif"
    ibaa8c2b8d48ff4b0848a07a1597e4d8c0842561ec4c6e0608fe24e0754f36765 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/n"
    ibb5a055522b92a57289f795953bc91a9047dd934fe9577923e13cd420be37ca9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/weekday"
    ibbcf6f4e007d6d73f76fed9eba9506277256367d70978c6ffbeedde11713d6d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pricemat"
    ibbfe64cfb60ca74a478178ca79042cf60ba03437dc03a754280148b9b3176587 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rate"
    ibc4fcaf29a5c8243ec98cb8505e6ba9d34dfee284e3358f0f05b1f686a62307e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sign"
    ibdd44510acdea53f5b37d091d420e4d288afe3820de5bbb0acef871af1977257 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/oct2dec"
    ibe04576cf9d5d63b6da0268d3cc3055d39a57721c71cea7dbf5f42c48b66c27f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/iseven"
    ibf1f5480f81b3b3696ab2d68c81f82220ed397116e34e6fb6efbb2851afefb93 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bin2hex"
    ic050d8168c9176d6687a5793c1fadec3ddf9f24ac430f78bba9013749748fd44 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/average"
    ic0653bdbf21ea84cbb5d5019a15f895237d04c65caca6c28185d25442ba7f369 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/complex"
    ic17145b99fdaba106e94f63d1fc5b6b0611a0fc6c1c5528f3b1bf69810a32a01 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isref"
    ic1c776e8f5bbf1aad14e8f62953dd29ae80f14bf4e5b2c7caa6d1ada9a7ae721 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/usdollar"
    ic24f007baaa8f53b154215e49485e0157bb841692afa3867ccb2c9c4ae61e022 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/f_inv"
    ic304fed62fa3d5a1cc95fd13ab3f2a1fde03b29af58ae54b9e8a39e294df5989 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imcsc"
    ic3583f8f5afc77ad8b847c5324d2d64e7fe873bea81f387539aca6056a3857d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/seriessum"
    ic42d36d02b28aa18548e0d940d1dff43b455a6503f32bec64e0b0d15698d1757 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/match"
    ic437e437d0e71b1d3cc7b04d1ab93a5f450917005d23d28cde22e2a557125385 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/today"
    ic4d8e49ce2a6532fcbbf6f64f180cdde43079498aab024848e5f8284fb6d197f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gauss"
    ic4eac566db71e14385e8ec72cd1979f90b074926778f8d00063effaaee1753b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/unicode"
    ic6226a166835f996dbcfdcba053ecd941c3435f27f4f8f9118b06373abae987e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/even"
    ic62551399736ff711f08b65ebb2c6c45dd5104478737a483075cf85e6b50a312 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dcount"
    ic706605c4bbb256effc937096f4ffe327337e51464de84b2a9ac009a7c5d5da4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/cosh"
    ic72126fb1b037996cabc6adbf309508be5276a8796ba813d30b44410593639cd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/median"
    ic787c76c3a4d7eafff0f967dfc5e037834a6e91c0a559608a912d2aa9d7de413 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/skew_p"
    ic7d1ebf32cd5371fa01a5068d8262f1e6cd7ba173fab384d0eb5c5335a3998e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/tanh"
    ic9303038a56639aca96d0a80154f111d9ac533dbacbeb52f9b2b41a622969d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/multinomial"
    icb6302ea49bc055bcc2890ea5225e1281ba1899d3b568c7790ed9f7a6b5dffeb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/besselj"
    icbb06724890f1f9eee228ebe7fe6b69845b988c2801b9284456800b631008eaa "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bessely"
    icd9c390016ac3cd52b33ff2f290463ccc95fb31b19d45dc5aae9850d0dd11c4f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/yearfrac"
    icd9f65c4c8dadcdc84336ea37b49c2de9e0249e1c5999cd4e34a373bfd3190df "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isoweeknum"
    ice38cc067ca7908b46f938cf870e7a7526ae2628878f13edb8141e7652825aa4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/and"
    ice7e4f6925eb7d4acd3adce5b6e13f99707d1adf6c6737526bd71a868cd56e22 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/now"
    iceeedb2d9423f2e1adad3827b258347b4576b32867eefd181fbf9d1eb00770b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ddb"
    icf3adc3356eac42b32e3ac5d862933f3bfa9bbf0d63e5da7bfcf8586c675c1e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/days"
    icfa5346e1359d617cef0e40dc92cffb0050874c2e36c2eeb709d033339e6faa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/areas"
    id01758ffe2f0f268965d1810951affba607e2846a4d38de06c90fa826e0ca646 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/impower"
    id04facfe8d3bb3050bef6bf2c05e1933d3ac3f8a2bf4b8065a7b736b90e6bcd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/permut"
    id108db16dde59511c17ec892beae748163d0a234c67bf6260e480f03b728ca1d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/stdev_s"
    id1df436faa503aff9086193529309c67d045ec3b8755c7e00ade9240c0771a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dollar"
    id2ca82ac9bf9907a7f3c6159707ea8f884481039b2a9f0e473b19cf4c452c44e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/atanh"
    id2e785e1992b73d2862c5f6713800dca6a10b0279f7dfa2390cd3de87c98a982 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsin"
    id3bd2821ec8246f8cc97afc48556fce63ae66cd7db2cad74efec2a2fcdd8134b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/findb"
    id3e78ad2851361e3dea5140cea20ac6e2a699248ed96eae657254dbfd285c917 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/tbillprice"
    id3eb4904482ac31c14e02e367462325b3f88310776cbf05b20218c79ee5e8dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hex2bin"
    id42b4bcbb523a4bfd99220e19da63a21928fd9e1abc900d4fe74791d40a7f44d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gamma_inv"
    id4ffe48042bb0104433516252781cafba47bdd900942a4826d89736fe93ec6b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/char"
    id5ec5e3a45e7ee1f01952c4a4d9bc19dfe91a032fac85e7b0779e7813ded9f5b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/standardize"
    id60116db5c866cc5761b7029c3448e9e53c5724442d3964225071973f3046e1d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coupdaybs"
    id66d330b86ce1c22be92754c31257b7e7fb4c39eebe39e086b3ab2833fc30b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/find"
    id7014948bb05b38a010de7d8ee6c8910b1084bb4974c710ce61cbf7e75c3fe65 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/vlookup"
    id742f89db038135ff55db923244279c7da2917bf7156592a277a7660e0bf2292 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/len"
    id7a16762b42d506809d333c334ee18086ca130a944d2d04f8c3902f273dc8a55 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sumsq"
    id7fd94c4c6345dfca7b605745f22ab04c9acec6f4dc9d1b1e80e53f35f71ae50 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/networkdays"
    id85fe98d73d528eb11688fc89fdadcb47512cb9d2f22de02e2af1f3125e360f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t_dist"
    id8942bdf6d4be967967837cf31114bfc0c30976ed4db096b9a586023b010d8c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hypgeom_dist"
    id89cd427d046dec5165cda135df3c54e70068bde7f5fd6f602bfe3207c2d7a8d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/binom_inv"
    id905b9402e0b9e259051a16830ef10696c305595d77843f560ca8b949c4bbc45 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/confidence_norm"
    ida2deeed82524cc3dfb4db9745f9ac68bd6bd46ca0996c86cc6552be00df989d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imabs"
    ida7b63d49a94b8ee39c93be34b500df57dba6c5be3483beadcb70cf97a13aa49 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/csch"
    idad0bb41b082d0269aa802164dce05c2a807fbc97a7ebd6069b225b5f6b64a39 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/midb"
    idada3620833af990b46a5c879e30430cb1543a2abebbc00f5c07e1c14a873b62 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/edate"
    idb4c3873441488d04313aaacd1cd362d380887783837426b2b4d6f1259c8f5be "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dproduct"
    idc0b79ac63ba2c8b3402dc56406f039561bf704065726033d98f42081b94d7b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isnumber"
    idc0bb38ae510a3b852b847ac7d5fd2d2def09000af4b33633be4fe0fbf2c5d37 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lcm"
    idc55d61c4fe53f23e33f4d0c3e62391402534d7c47cb00c2d867a340f4ad680e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsinh"
    idcd01cb21b27b9b15225c6c93e8dc2d29a633a4bca76b7dea84c298644ef1375 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/lookup"
    idce3c3d2910ab657d92dfa59ea57a85b9a9134fbd6ad6a2976fdb9e96347b232 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/log"
    idd96fd773b6cbeff56d5db250cfef3d2fab48aab4d11ed9c706c56d6489aa890 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/fact"
    ide313f7c15c109ec4520431747e1c1ad59b98641175faed943944e8b90b4ff9e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/f_inv_rt"
    idf877ceae99180689e8e9a8bf9eee09332189d3fd7e31643a899b0538cc0dcef "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/ppmt"
    idff66a72240da9492b8b2fd994dc3c423b6620bd778050d7d19b21b3b50d24f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/iso_ceiling"
    ie004a57ca723f01f9805f68e4c7807c360c84b1420011ff08a32ec2e039ce9ed "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/gammaln_precise"
    ie0a1929468b6da507dbfb84dc9baf3f3753b6c5c0d9032028149798c714291b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/floor_math"
    ie0f6fae4b375577574ec3a3d10be3c1e7b7dff4f2a33c922ebb6b6011b3fc6c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/bitrshift"
    ie1788a287e583aaee3f2bb77e8e1a468fa3b2a0e293faef36a82104552219f03 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/large"
    ie347231f0f0a8ce2cadb635d237561360591f79cd4435bab6bf441298258887e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsec"
    ie496f58abf287e229fa39afb5ddb533e54ea59897e8c4cac11239aee71cb2ff5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/npv"
    ie4a5e796f7fa374cef652c75e70f5a1b1f87110404de05bc90276b31af680f01 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/hex2dec"
    ie53ed9255c4361ef56a74a112e81fc4784b28bfbf0cdacc882c262a5c28db6d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dstdevp"
    ie5577128b8e24362fb1c5a485d4911522b5ced09cde805cc1a05969da2097603 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sqrt"
    ie5678a893191a28298bfcb69e8f697efffc2d356b0f9ee2c39a9e5b3e5bd485a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/price"
    ie598c23cd821465563c1a517655127af97b775b456b274545e564c741d25d5cf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/tbillyield"
    ie7ed1bf0ad464b9c429a113f7c45516d7c13166cd0c581a3e34aa78ba6423b7b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/varpa"
    ie84ebf6c1f6d5fd5590422fb45bc8d7ee6a943ca1a287247547419e876d52501 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/improduct"
    ie95ea4c78471cf9e8da86904a28af4b814e78189e787241116f733848361a7b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sumifs"
    iea0f3b887bb91d47447172b7d3ea7fcec40b7c9dccc043a99fefa454a07cbec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/t_inv"
    ieac58a41e841549fbd8af3e54f87582f50250dca151693068af1909ecec35019 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/roundup"
    ieb15ac389058ef18a973e324edaff73830c8dddc6f6de9fd9d1ce8e074fdf8c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/negbinom_dist"
    ieb56fecb0739f2e710e721eb6508a3c2c25c0bbe74af47a66d0379daabee2fdc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isodd"
    iecec2b18f5b5d30ea473c416117fa75c41083c7e9471e526028e1b7dfbb1e37d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/isna"
    ied2f846e9c4ada5ae5bc068debda29a9a32acc68ba06aee53a90e47f13432261 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/cumprinc"
    ied40e56292d4efe1ae5f8b24055751ee82932520d99ae7cbbd34919617fc342b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sum"
    iedadcc399469fc51ce774f195b317d33cd4324ca9dd16a18f898ade9fdbf4a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/pv"
    iedfa6bb32f314fe3f8c5968834a5fb0f1b78436ba0dd57eb691229eb7ca9e939 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/combin"
    ief2b188c34ec07e72592e39a3d9a153a92765342857c1a9e782c05d12b30f62e "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/syd"
    if12a6595734390e605b40df53030c075e2bcd558a071685ee7442876b9c91217 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/coupnum"
    if1770654587a61f420eb20ffc8a5ea1f8b4c1742716cd5e46b4bff76c3a5264a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/stdev_p"
    if394010a3655e451e9502ee93545f5fe9f2d41f16ca348052e9a061dfb3bae6c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/var_s"
    if579faccaec7d615e7f4e3259a551dcb170851c1b18626ed6b952606d4f270c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imcot"
    if5c0edc21dffeb380c263c85352bea05578a6349dd520aeb5d85a35c845c2645 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/expon_dist"
    if73719a1dc2b6d6991df34ae8b7e7dfb30eb62bf72a584adf4ecd9394edee50d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/countifs"
    if7e94ba699066e6986bb11cb8c1fe185dfeac6238c220ba28a74982e8abf6364 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imtan"
    if8cd7bff2592738d8202fe4b64adb52b0349bd683eb7b2883788632cf859eb8f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/stdeva"
    if91020b076527317cc3d69f9150d268d684be9d42a938d778b2cf996fb23a4b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/acosh"
    if9bd3673059c5d91d6a5844a0702b4022748696d7caaebd69431763ab33549c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/f_dist"
    ifa3194c7763f634c1efa4e8ae657eaa8c118864c0f01ba521f2c5ae992f44bca "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sin"
    ifbe89311210db7e5aebd2c2a5941f6c6ee75028b80166be8dccd2d48242d4026 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/rand"
    ifc30dc22fa9fe32b7a51832bdd0cf2e35370072463fc48030297308894f031d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/quotient"
    ifc5534ed3042a1a7f49eba82e143c4978c0db242a2657d9bbe9576e636343d22 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/na"
    ifcab8491af5591341ba1b19b47b840ce4366d7e6e178032392893feb761d1f4d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/dec2oct"
    ifd1dbe5d2897f560ce7816e3d3dc4208b6ed57b5ae5446df8d93e5d261c102fb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/not"
    ifd522acfa264f2b97ee138f4abcf7d087cfe0cd47eb1c57d8121172e2f5db7ee "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/imsech"
    ifdfada0f682e13dd6856925da661b74617c1e052d16c650a5b1520225acf60dc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/text"
    iff2f96c56ce6b791c099573e926303c159821091509cb5a3cdffd2974d8cfb13 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/substitute"
    iff5d1b1deb3b8f2c2d3de6f2d3cbbc8961c3a069683ac3268635178bc77155cc "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/functions/sheets"
)

// FunctionsRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}\workbook\functions
type FunctionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FunctionsRequestBuilderDeleteOptions options for Delete
type FunctionsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FunctionsRequestBuilderGetOptions options for Get
type FunctionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FunctionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FunctionsRequestBuilderGetQueryParameters get functions from workbooks
type FunctionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// FunctionsRequestBuilderPatchOptions options for Patch
type FunctionsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFunctions;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *FunctionsRequestBuilder) Abs()(*i782c59c275368a34292a362686f23a3045f8e4b667c49f94302584781b7ab08b.AbsRequestBuilder) {
    return i782c59c275368a34292a362686f23a3045f8e4b667c49f94302584781b7ab08b.NewAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AccrInt()(*i1ed5a81b94a902275caf64d5646aa6a5d64df5ee721b688f4945f5422aec8c38.AccrIntRequestBuilder) {
    return i1ed5a81b94a902275caf64d5646aa6a5d64df5ee721b688f4945f5422aec8c38.NewAccrIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AccrIntM()(*i750e48618ea35ce290e17ae4e72d7b69de05921c0c97083a399b25eb35ce6924.AccrIntMRequestBuilder) {
    return i750e48618ea35ce290e17ae4e72d7b69de05921c0c97083a399b25eb35ce6924.NewAccrIntMRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acos()(*i271789ad8387c2c7897caaa1d23fdb3f98f1e3cf2f792b2a03f62a7a586d9df3.AcosRequestBuilder) {
    return i271789ad8387c2c7897caaa1d23fdb3f98f1e3cf2f792b2a03f62a7a586d9df3.NewAcosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acosh()(*if91020b076527317cc3d69f9150d268d684be9d42a938d778b2cf996fb23a4b7.AcoshRequestBuilder) {
    return if91020b076527317cc3d69f9150d268d684be9d42a938d778b2cf996fb23a4b7.NewAcoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acot()(*i9f86a18034de4b30e9787efe02f4e0a4284676cd1f4401006517a31aa55ffa6b.AcotRequestBuilder) {
    return i9f86a18034de4b30e9787efe02f4e0a4284676cd1f4401006517a31aa55ffa6b.NewAcotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Acoth()(*i19b00c3ca966260792b57fd09e773e9ce9e699138711859ca47c20f1a9e34b4e.AcothRequestBuilder) {
    return i19b00c3ca966260792b57fd09e773e9ce9e699138711859ca47c20f1a9e34b4e.NewAcothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AmorDegrc()(*i1dde59714c4439bf7a5432f226d87b56892b824dae30192790cec46f5fa75ed6.AmorDegrcRequestBuilder) {
    return i1dde59714c4439bf7a5432f226d87b56892b824dae30192790cec46f5fa75ed6.NewAmorDegrcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AmorLinc()(*ib9d1f9ecf825ea1ba386243b2d81f7d6f2e5e0d078ecd6ab556ae1176298b85a.AmorLincRequestBuilder) {
    return ib9d1f9ecf825ea1ba386243b2d81f7d6f2e5e0d078ecd6ab556ae1176298b85a.NewAmorLincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) And()(*ice38cc067ca7908b46f938cf870e7a7526ae2628878f13edb8141e7652825aa4.AndRequestBuilder) {
    return ice38cc067ca7908b46f938cf870e7a7526ae2628878f13edb8141e7652825aa4.NewAndRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Arabic()(*i8942bff90eda23d84d12cbc5b1bcd0f5ad61c132f149f9345ec5725a177ebf14.ArabicRequestBuilder) {
    return i8942bff90eda23d84d12cbc5b1bcd0f5ad61c132f149f9345ec5725a177ebf14.NewArabicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Areas()(*icfa5346e1359d617cef0e40dc92cffb0050874c2e36c2eeb709d033339e6faa6.AreasRequestBuilder) {
    return icfa5346e1359d617cef0e40dc92cffb0050874c2e36c2eeb709d033339e6faa6.NewAreasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asc()(*i5d1af52be2b3187ec0e03be085c84d7655295a9764e715c8352272c991a5b31a.AscRequestBuilder) {
    return i5d1af52be2b3187ec0e03be085c84d7655295a9764e715c8352272c991a5b31a.NewAscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asin()(*i3c6e2bd2b92cb13f1e949c31c034ca7519729cbd902289447d8fb1f45d82ba7f.AsinRequestBuilder) {
    return i3c6e2bd2b92cb13f1e949c31c034ca7519729cbd902289447d8fb1f45d82ba7f.NewAsinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Asinh()(*i71085b17f9c510fc838d37422083076af4290773fdac918f19c26e7aa8cfb541.AsinhRequestBuilder) {
    return i71085b17f9c510fc838d37422083076af4290773fdac918f19c26e7aa8cfb541.NewAsinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atan()(*i4748eaa62f242c480beca098f37730bd85581f0454d96962ab77c5bcc4077c53.AtanRequestBuilder) {
    return i4748eaa62f242c480beca098f37730bd85581f0454d96962ab77c5bcc4077c53.NewAtanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atan2()(*i38da5a55b58b6bd8bc84fa109d38ad5ba55b676692f6ebb60a25def11141d862.Atan2RequestBuilder) {
    return i38da5a55b58b6bd8bc84fa109d38ad5ba55b676692f6ebb60a25def11141d862.NewAtan2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Atanh()(*id2ca82ac9bf9907a7f3c6159707ea8f884481039b2a9f0e473b19cf4c452c44e.AtanhRequestBuilder) {
    return id2ca82ac9bf9907a7f3c6159707ea8f884481039b2a9f0e473b19cf4c452c44e.NewAtanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AveDev()(*i3b37f1d056cb7cf047c630f85e8d49d51384a42baa20b671b8f98d2d440feba1.AveDevRequestBuilder) {
    return i3b37f1d056cb7cf047c630f85e8d49d51384a42baa20b671b8f98d2d440feba1.NewAveDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Average()(*ic050d8168c9176d6687a5793c1fadec3ddf9f24ac430f78bba9013749748fd44.AverageRequestBuilder) {
    return ic050d8168c9176d6687a5793c1fadec3ddf9f24ac430f78bba9013749748fd44.NewAverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageA()(*ib4f820ed510f8f1aa69d6b111c66898e64509ab1648adb13b7a2718624d00f7c.AverageARequestBuilder) {
    return ib4f820ed510f8f1aa69d6b111c66898e64509ab1648adb13b7a2718624d00f7c.NewAverageARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageIf()(*iba9a57b0cf80608e77654b3a94673f296c392fd17d9c2605fad518d6f8c4d18c.AverageIfRequestBuilder) {
    return iba9a57b0cf80608e77654b3a94673f296c392fd17d9c2605fad518d6f8c4d18c.NewAverageIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) AverageIfs()(*i33ef7455cd3354f695f10748ea17e47e4d0622a2a93d13f279966caa3f74e31c.AverageIfsRequestBuilder) {
    return i33ef7455cd3354f695f10748ea17e47e4d0622a2a93d13f279966caa3f74e31c.NewAverageIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BahtText()(*ib9d47bc5a4b46376eb93154925b463120b0c1b160682670d4ddb93d6001ad874.BahtTextRequestBuilder) {
    return ib9d47bc5a4b46376eb93154925b463120b0c1b160682670d4ddb93d6001ad874.NewBahtTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Base()(*i2f884c270bdc5f79a1b5c6d1fd764a4097f9f98cc08bcab6c13418d844f3a6f2.BaseRequestBuilder) {
    return i2f884c270bdc5f79a1b5c6d1fd764a4097f9f98cc08bcab6c13418d844f3a6f2.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselI()(*i11d6df0ff757addbcc18e3f5e5696d91da8693a9560cbb682fa52b8f4539be2c.BesselIRequestBuilder) {
    return i11d6df0ff757addbcc18e3f5e5696d91da8693a9560cbb682fa52b8f4539be2c.NewBesselIRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselJ()(*icb6302ea49bc055bcc2890ea5225e1281ba1899d3b568c7790ed9f7a6b5dffeb.BesselJRequestBuilder) {
    return icb6302ea49bc055bcc2890ea5225e1281ba1899d3b568c7790ed9f7a6b5dffeb.NewBesselJRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselK()(*i96228adffa7f7c2c4fabef938340b0e347e15559d0e63f7b36c6874502933832.BesselKRequestBuilder) {
    return i96228adffa7f7c2c4fabef938340b0e347e15559d0e63f7b36c6874502933832.NewBesselKRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) BesselY()(*icbb06724890f1f9eee228ebe7fe6b69845b988c2801b9284456800b631008eaa.BesselYRequestBuilder) {
    return icbb06724890f1f9eee228ebe7fe6b69845b988c2801b9284456800b631008eaa.NewBesselYRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Beta_Dist()(*i539158791087b6399c0cb327f13cb7790c9745fa153c61e552c01b787c91efbf.Beta_DistRequestBuilder) {
    return i539158791087b6399c0cb327f13cb7790c9745fa153c61e552c01b787c91efbf.NewBeta_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Beta_Inv()(*ib4640205a5ee55bc3b8a20bcbe1259f213daf2b26928179da4ba1125ad3d3c95.Beta_InvRequestBuilder) {
    return ib4640205a5ee55bc3b8a20bcbe1259f213daf2b26928179da4ba1125ad3d3c95.NewBeta_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Dec()(*i11c7d19b25404b514e263be9547aef03f560afcfcde0489c7c332825f429122b.Bin2DecRequestBuilder) {
    return i11c7d19b25404b514e263be9547aef03f560afcfcde0489c7c332825f429122b.NewBin2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Hex()(*ibf1f5480f81b3b3696ab2d68c81f82220ed397116e34e6fb6efbb2851afefb93.Bin2HexRequestBuilder) {
    return ibf1f5480f81b3b3696ab2d68c81f82220ed397116e34e6fb6efbb2851afefb93.NewBin2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bin2Oct()(*i60ef136bd90310658d302308c48fea77adcc95e0242c2f24f08b9ac1f74470bc.Bin2OctRequestBuilder) {
    return i60ef136bd90310658d302308c48fea77adcc95e0242c2f24f08b9ac1f74470bc.NewBin2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Dist()(*i834687913aeb2dfb3df0401881633d320023843d6c09e5d88ce93d598db4e98f.Binom_DistRequestBuilder) {
    return i834687913aeb2dfb3df0401881633d320023843d6c09e5d88ce93d598db4e98f.NewBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Dist_Range()(*i25279acfc780c5710e2ac0bd05598cb053419123f87b0e9b3f2cc83f197024c0.Binom_Dist_RangeRequestBuilder) {
    return i25279acfc780c5710e2ac0bd05598cb053419123f87b0e9b3f2cc83f197024c0.NewBinom_Dist_RangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Binom_Inv()(*id89cd427d046dec5165cda135df3c54e70068bde7f5fd6f602bfe3207c2d7a8d.Binom_InvRequestBuilder) {
    return id89cd427d046dec5165cda135df3c54e70068bde7f5fd6f602bfe3207c2d7a8d.NewBinom_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitand()(*i3e83145b3125a8eb2f1a717e5d6d6f44f41b1443ad97bb094bc03f0a892b75db.BitandRequestBuilder) {
    return i3e83145b3125a8eb2f1a717e5d6d6f44f41b1443ad97bb094bc03f0a892b75db.NewBitandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitlshift()(*i8f4c6eeb4303ba527f8f345e5726779c14970ba71a2280bc32aa7301236ad2e7.BitlshiftRequestBuilder) {
    return i8f4c6eeb4303ba527f8f345e5726779c14970ba71a2280bc32aa7301236ad2e7.NewBitlshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitor()(*i7aee590138e19a1b27ad74be0e89a72f68c0974cb718a9e4486336810b4e6d85.BitorRequestBuilder) {
    return i7aee590138e19a1b27ad74be0e89a72f68c0974cb718a9e4486336810b4e6d85.NewBitorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitrshift()(*ie0f6fae4b375577574ec3a3d10be3c1e7b7dff4f2a33c922ebb6b6011b3fc6c7.BitrshiftRequestBuilder) {
    return ie0f6fae4b375577574ec3a3d10be3c1e7b7dff4f2a33c922ebb6b6011b3fc6c7.NewBitrshiftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Bitxor()(*iaaf4f922056c865e9b5213ff8a651682b31518b0bd7b7475f5f133d6fe68e71b.BitxorRequestBuilder) {
    return iaaf4f922056c865e9b5213ff8a651682b31518b0bd7b7475f5f133d6fe68e71b.NewBitxorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ceiling_Math()(*i7cb7ca6d53631bf58f2266e1292e6c6331539af3696b718238bad702aaa8480d.Ceiling_MathRequestBuilder) {
    return i7cb7ca6d53631bf58f2266e1292e6c6331539af3696b718238bad702aaa8480d.NewCeiling_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ceiling_Precise()(*i34f2a4a24d744d3fc6016654b9341e19775b8dea258192b301b9d8cf4e72d763.Ceiling_PreciseRequestBuilder) {
    return i34f2a4a24d744d3fc6016654b9341e19775b8dea258192b301b9d8cf4e72d763.NewCeiling_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Char()(*id4ffe48042bb0104433516252781cafba47bdd900942a4826d89736fe93ec6b3.CharRequestBuilder) {
    return id4ffe48042bb0104433516252781cafba47bdd900942a4826d89736fe93ec6b3.NewCharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Dist()(*i4fb5a0beaa11f0d431ffe7b2776a63f1519b2b0730f7b34c1557c95dce210d0f.ChiSq_DistRequestBuilder) {
    return i4fb5a0beaa11f0d431ffe7b2776a63f1519b2b0730f7b34c1557c95dce210d0f.NewChiSq_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Dist_RT()(*i6170b518e26aeac41ed27c586fcc4d593624a8dfc1898f24e26195a1b08adea9.ChiSq_Dist_RTRequestBuilder) {
    return i6170b518e26aeac41ed27c586fcc4d593624a8dfc1898f24e26195a1b08adea9.NewChiSq_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Inv()(*i2be0b0806f84053729758b3570db185ad402da1d3ffd988392771290ec7d78d8.ChiSq_InvRequestBuilder) {
    return i2be0b0806f84053729758b3570db185ad402da1d3ffd988392771290ec7d78d8.NewChiSq_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ChiSq_Inv_RT()(*i5e2ca6a19b1a68e98eeb5a22cdcb32b257cbd898afc03c740da42dbdaa146731.ChiSq_Inv_RTRequestBuilder) {
    return i5e2ca6a19b1a68e98eeb5a22cdcb32b257cbd898afc03c740da42dbdaa146731.NewChiSq_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Choose()(*i3a5b50be37d36553bf1119b198b458ef940eeda87d384cfb98fb6b76bcc3ea70.ChooseRequestBuilder) {
    return i3a5b50be37d36553bf1119b198b458ef940eeda87d384cfb98fb6b76bcc3ea70.NewChooseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Clean()(*i4eb6a2fb4e8f255bbffdc6e5b9a7f2ded94bce657445c460aadb52f5b34458ac.CleanRequestBuilder) {
    return i4eb6a2fb4e8f255bbffdc6e5b9a7f2ded94bce657445c460aadb52f5b34458ac.NewCleanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Code()(*ib37ba68a5fab1c602fae7e488eea5313b78cd71fa7132a748d7b9402190fbe46.CodeRequestBuilder) {
    return ib37ba68a5fab1c602fae7e488eea5313b78cd71fa7132a748d7b9402190fbe46.NewCodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Columns()(*i29da91cf8167cc7cfc9d8a229b680f1dfed88338c9faeb9401a09aacf7ad5ace.ColumnsRequestBuilder) {
    return i29da91cf8167cc7cfc9d8a229b680f1dfed88338c9faeb9401a09aacf7ad5ace.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Combin()(*iedfa6bb32f314fe3f8c5968834a5fb0f1b78436ba0dd57eb691229eb7ca9e939.CombinRequestBuilder) {
    return iedfa6bb32f314fe3f8c5968834a5fb0f1b78436ba0dd57eb691229eb7ca9e939.NewCombinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Combina()(*i392b51c7d8691d97c55045a593139b9654ca512cf8ca6940e3716a0765231661.CombinaRequestBuilder) {
    return i392b51c7d8691d97c55045a593139b9654ca512cf8ca6940e3716a0765231661.NewCombinaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Complex()(*ic0653bdbf21ea84cbb5d5019a15f895237d04c65caca6c28185d25442ba7f369.ComplexRequestBuilder) {
    return ic0653bdbf21ea84cbb5d5019a15f895237d04c65caca6c28185d25442ba7f369.NewComplexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Concatenate()(*i014f5fef465cec5d69bfe6228aa49769a28be347b59b869cc6d201c421c5fc55.ConcatenateRequestBuilder) {
    return i014f5fef465cec5d69bfe6228aa49769a28be347b59b869cc6d201c421c5fc55.NewConcatenateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Confidence_Norm()(*id905b9402e0b9e259051a16830ef10696c305595d77843f560ca8b949c4bbc45.Confidence_NormRequestBuilder) {
    return id905b9402e0b9e259051a16830ef10696c305595d77843f560ca8b949c4bbc45.NewConfidence_NormRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Confidence_T()(*i3b57e79beeef2f08027bdf9fd759cc0858a70f39d4a35a6c44414778563012f6.Confidence_TRequestBuilder) {
    return i3b57e79beeef2f08027bdf9fd759cc0858a70f39d4a35a6c44414778563012f6.NewConfidence_TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFunctionsRequestBuilderInternal instantiates a new FunctionsRequestBuilder and sets the default values.
func NewFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FunctionsRequestBuilder) {
    m := &FunctionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/workbook/functions{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFunctionsRequestBuilder instantiates a new FunctionsRequestBuilder and sets the default values.
func NewFunctionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *FunctionsRequestBuilder) Convert()(*iab13452fe1cf2fbff7b3ef1f2840e01fd1ecd0a0acdd8ebea3c83698c958e723.ConvertRequestBuilder) {
    return iab13452fe1cf2fbff7b3ef1f2840e01fd1ecd0a0acdd8ebea3c83698c958e723.NewConvertRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cos()(*i6a61af20e5713d0e4714cb59cf2e4d3c54a8491ed03cc4eef0303857d4c09194.CosRequestBuilder) {
    return i6a61af20e5713d0e4714cb59cf2e4d3c54a8491ed03cc4eef0303857d4c09194.NewCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cosh()(*ic706605c4bbb256effc937096f4ffe327337e51464de84b2a9ac009a7c5d5da4.CoshRequestBuilder) {
    return ic706605c4bbb256effc937096f4ffe327337e51464de84b2a9ac009a7c5d5da4.NewCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Cot()(*i1a280cdde157745d3f87a26d747c812b1665bc8b66c0fa4890d28140984b06e5.CotRequestBuilder) {
    return i1a280cdde157745d3f87a26d747c812b1665bc8b66c0fa4890d28140984b06e5.NewCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Coth()(*i31b6a6d9d88e1a77eb013a0a4f84b7c8f4f0e1487cd1befa282da7c3a5467b16.CothRequestBuilder) {
    return i31b6a6d9d88e1a77eb013a0a4f84b7c8f4f0e1487cd1befa282da7c3a5467b16.NewCothRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Count()(*i9000ffdfcb17d10e09832a5fda64717cb5f77f9ba0daa420e6680c0bf9b50128.CountRequestBuilder) {
    return i9000ffdfcb17d10e09832a5fda64717cb5f77f9ba0daa420e6680c0bf9b50128.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountA()(*ib8ba910ad32637cc152e0085ee699ee1cc26b29b4001fd5e5f6a368de40c4228.CountARequestBuilder) {
    return ib8ba910ad32637cc152e0085ee699ee1cc26b29b4001fd5e5f6a368de40c4228.NewCountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountBlank()(*i9ef2fd435a266fea3206c285815b97a69f2a30b0e9e2358d804969a866eb337c.CountBlankRequestBuilder) {
    return i9ef2fd435a266fea3206c285815b97a69f2a30b0e9e2358d804969a866eb337c.NewCountBlankRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountIf()(*i7afa6354755ab604b2000c207aaa310d7f80ce2d8ac7ad23a797c12b57a4bf37.CountIfRequestBuilder) {
    return i7afa6354755ab604b2000c207aaa310d7f80ce2d8ac7ad23a797c12b57a4bf37.NewCountIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CountIfs()(*if73719a1dc2b6d6991df34ae8b7e7dfb30eb62bf72a584adf4ecd9394edee50d.CountIfsRequestBuilder) {
    return if73719a1dc2b6d6991df34ae8b7e7dfb30eb62bf72a584adf4ecd9394edee50d.NewCountIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDayBs()(*id60116db5c866cc5761b7029c3448e9e53c5724442d3964225071973f3046e1d.CoupDayBsRequestBuilder) {
    return id60116db5c866cc5761b7029c3448e9e53c5724442d3964225071973f3046e1d.NewCoupDayBsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDays()(*i7f104846fcb71d8d34066552109bb9ebb1921483ac42b4c84d9622742e45c3a9.CoupDaysRequestBuilder) {
    return i7f104846fcb71d8d34066552109bb9ebb1921483ac42b4c84d9622742e45c3a9.NewCoupDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupDaysNc()(*ib9f89317bd5c3fd7fe775dbcdda3b3a0049814dbf7a1cf903cd14e5384478033.CoupDaysNcRequestBuilder) {
    return ib9f89317bd5c3fd7fe775dbcdda3b3a0049814dbf7a1cf903cd14e5384478033.NewCoupDaysNcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupNcd()(*i3ed3d358b4dce2bb6a00a026fbf5a2343557dce788227b550f1f5ccbd12b1615.CoupNcdRequestBuilder) {
    return i3ed3d358b4dce2bb6a00a026fbf5a2343557dce788227b550f1f5ccbd12b1615.NewCoupNcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupNum()(*if12a6595734390e605b40df53030c075e2bcd558a071685ee7442876b9c91217.CoupNumRequestBuilder) {
    return if12a6595734390e605b40df53030c075e2bcd558a071685ee7442876b9c91217.NewCoupNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CoupPcd()(*i861da83d39292b36ce13f5c2b964c9924cc966b749fc25d2b5b6a669d83dd8c6.CoupPcdRequestBuilder) {
    return i861da83d39292b36ce13f5c2b964c9924cc966b749fc25d2b5b6a669d83dd8c6.NewCoupPcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property functions for workbooks
func (m *FunctionsRequestBuilder) CreateDeleteRequestInformation(options *FunctionsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get functions from workbooks
func (m *FunctionsRequestBuilder) CreateGetRequestInformation(options *FunctionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property functions in workbooks
func (m *FunctionsRequestBuilder) CreatePatchRequestInformation(options *FunctionsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *FunctionsRequestBuilder) Csc()(*i6aacbe20dc6e50258c9a6e8d391b05d4e7c7f96c9d960dcf3afd0ac5c111cda2.CscRequestBuilder) {
    return i6aacbe20dc6e50258c9a6e8d391b05d4e7c7f96c9d960dcf3afd0ac5c111cda2.NewCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Csch()(*ida7b63d49a94b8ee39c93be34b500df57dba6c5be3483beadcb70cf97a13aa49.CschRequestBuilder) {
    return ida7b63d49a94b8ee39c93be34b500df57dba6c5be3483beadcb70cf97a13aa49.NewCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CumIPmt()(*i239cf5ab85067bafc56c3ea278efdda656968240fda3d73eb82eda937e180bdb.CumIPmtRequestBuilder) {
    return i239cf5ab85067bafc56c3ea278efdda656968240fda3d73eb82eda937e180bdb.NewCumIPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) CumPrinc()(*ied2f846e9c4ada5ae5bc068debda29a9a32acc68ba06aee53a90e47f13432261.CumPrincRequestBuilder) {
    return ied2f846e9c4ada5ae5bc068debda29a9a32acc68ba06aee53a90e47f13432261.NewCumPrincRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Date()(*i91d54d1b34c20ed6b9e3d724abea350e57457e087e1724b981a173dcf2863cc8.DateRequestBuilder) {
    return i91d54d1b34c20ed6b9e3d724abea350e57457e087e1724b981a173dcf2863cc8.NewDateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Datevalue()(*i11635e1f1864a4c8e9365eca479c56e1761cc4c1cb596a971d3e34a740a3c085.DatevalueRequestBuilder) {
    return i11635e1f1864a4c8e9365eca479c56e1761cc4c1cb596a971d3e34a740a3c085.NewDatevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Daverage()(*i4be11536eae45fd4ea7adf5d4299b71cd862146faa910c723899f6911f0552a7.DaverageRequestBuilder) {
    return i4be11536eae45fd4ea7adf5d4299b71cd862146faa910c723899f6911f0552a7.NewDaverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Day()(*i97c850bd34ef5bb4796243f0cd6452958a6816910bcb04c5fef1ebbb05cf8ee2.DayRequestBuilder) {
    return i97c850bd34ef5bb4796243f0cd6452958a6816910bcb04c5fef1ebbb05cf8ee2.NewDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Days()(*icf3adc3356eac42b32e3ac5d862933f3bfa9bbf0d63e5da7bfcf8586c675c1e2.DaysRequestBuilder) {
    return icf3adc3356eac42b32e3ac5d862933f3bfa9bbf0d63e5da7bfcf8586c675c1e2.NewDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Days360()(*i1b1e5f6c5fe8a56a140323e96bcda3645d409917a438e57ec2f42590bfbd0bf2.Days360RequestBuilder) {
    return i1b1e5f6c5fe8a56a140323e96bcda3645d409917a438e57ec2f42590bfbd0bf2.NewDays360RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Db()(*ia489b4341b29bd08b05402f5a65f8ecd4424d329b3dbf6daec1cc94e0ba7569f.DbRequestBuilder) {
    return ia489b4341b29bd08b05402f5a65f8ecd4424d329b3dbf6daec1cc94e0ba7569f.NewDbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dbcs()(*iafdeaffeb4387b6753161d084bb5c7369fbcbd19961da7f33df1c37fe2c57c49.DbcsRequestBuilder) {
    return iafdeaffeb4387b6753161d084bb5c7369fbcbd19961da7f33df1c37fe2c57c49.NewDbcsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dcount()(*ic62551399736ff711f08b65ebb2c6c45dd5104478737a483075cf85e6b50a312.DcountRequestBuilder) {
    return ic62551399736ff711f08b65ebb2c6c45dd5104478737a483075cf85e6b50a312.NewDcountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DcountA()(*i3a713ed3e471d758071d6320615cb630620a44eda7882e59205ee608b16ab8eb.DcountARequestBuilder) {
    return i3a713ed3e471d758071d6320615cb630620a44eda7882e59205ee608b16ab8eb.NewDcountARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ddb()(*iceeedb2d9423f2e1adad3827b258347b4576b32867eefd181fbf9d1eb00770b9.DdbRequestBuilder) {
    return iceeedb2d9423f2e1adad3827b258347b4576b32867eefd181fbf9d1eb00770b9.NewDdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Bin()(*ia874a740cc8fc9d80c5d38f6d847606bc8805a7749b782542042db71a5c6a1eb.Dec2BinRequestBuilder) {
    return ia874a740cc8fc9d80c5d38f6d847606bc8805a7749b782542042db71a5c6a1eb.NewDec2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Hex()(*i3516812fdea5289d5c1239ff45469b8cd752ae305dd572fc2315b1a817738e3d.Dec2HexRequestBuilder) {
    return i3516812fdea5289d5c1239ff45469b8cd752ae305dd572fc2315b1a817738e3d.NewDec2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dec2Oct()(*ifcab8491af5591341ba1b19b47b840ce4366d7e6e178032392893feb761d1f4d.Dec2OctRequestBuilder) {
    return ifcab8491af5591341ba1b19b47b840ce4366d7e6e178032392893feb761d1f4d.NewDec2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Decimal()(*i9400ff8a2db8c4ce9ec4300a6fba694671fe73b126c7864b06515a4f329c745b.DecimalRequestBuilder) {
    return i9400ff8a2db8c4ce9ec4300a6fba694671fe73b126c7864b06515a4f329c745b.NewDecimalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Degrees()(*ia1959188eafd4bcd478ecf33a6287ceea78edb264b17875aac2ffe130ca69250.DegreesRequestBuilder) {
    return ia1959188eafd4bcd478ecf33a6287ceea78edb264b17875aac2ffe130ca69250.NewDegreesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property functions for workbooks
func (m *FunctionsRequestBuilder) Delete(options *FunctionsRequestBuilderDeleteOptions)(error) {
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
func (m *FunctionsRequestBuilder) Delta()(*i0d4bf14b866d57e7379422ab460cef0187cd905710c4fb759a2015829d01f655.DeltaRequestBuilder) {
    return i0d4bf14b866d57e7379422ab460cef0187cd905710c4fb759a2015829d01f655.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DevSq()(*i64400f50c71ff8b2b271f1ceae5f9e839b64ae25073adf239505b644be0b5bd6.DevSqRequestBuilder) {
    return i64400f50c71ff8b2b271f1ceae5f9e839b64ae25073adf239505b644be0b5bd6.NewDevSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dget()(*i5b2286301b4d4e486cb0f56c4227fcfef09387503478e065a21767ef0f527cfa.DgetRequestBuilder) {
    return i5b2286301b4d4e486cb0f56c4227fcfef09387503478e065a21767ef0f527cfa.NewDgetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Disc()(*i8a12a2b5f264a7c74820e9b113e4a38aed7a746d39e9da6e8d180db215e3c28c.DiscRequestBuilder) {
    return i8a12a2b5f264a7c74820e9b113e4a38aed7a746d39e9da6e8d180db215e3c28c.NewDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dmax()(*i74c8bcbbb2397de63f9e8cbdf6652c485da2c060b6c26fc01b4870ff2660609b.DmaxRequestBuilder) {
    return i74c8bcbbb2397de63f9e8cbdf6652c485da2c060b6c26fc01b4870ff2660609b.NewDmaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dmin()(*i9205761ddefe6680928f97b76a8e0de932c39501da2b6760101861f096bbb281.DminRequestBuilder) {
    return i9205761ddefe6680928f97b76a8e0de932c39501da2b6760101861f096bbb281.NewDminRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dollar()(*id1df436faa503aff9086193529309c67d045ec3b8755c7e00ade9240c0771a89.DollarRequestBuilder) {
    return id1df436faa503aff9086193529309c67d045ec3b8755c7e00ade9240c0771a89.NewDollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DollarDe()(*i5102dc4301948efe9759bb91a61ce696b9a7eae3cdfa559114a7b37dd8074776.DollarDeRequestBuilder) {
    return i5102dc4301948efe9759bb91a61ce696b9a7eae3cdfa559114a7b37dd8074776.NewDollarDeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DollarFr()(*i5929444bf113c70d20e2d4295b6983a9724d0227d6c4d438f913f22f07823801.DollarFrRequestBuilder) {
    return i5929444bf113c70d20e2d4295b6983a9724d0227d6c4d438f913f22f07823801.NewDollarFrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dproduct()(*idb4c3873441488d04313aaacd1cd362d380887783837426b2b4d6f1259c8f5be.DproductRequestBuilder) {
    return idb4c3873441488d04313aaacd1cd362d380887783837426b2b4d6f1259c8f5be.NewDproductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DstDev()(*i8bcb057ec8007b84ba78097f278d3988968527d460ef14fb3a3150960e187052.DstDevRequestBuilder) {
    return i8bcb057ec8007b84ba78097f278d3988968527d460ef14fb3a3150960e187052.NewDstDevRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DstDevP()(*ie53ed9255c4361ef56a74a112e81fc4784b28bfbf0cdacc882c262a5c28db6d1.DstDevPRequestBuilder) {
    return ie53ed9255c4361ef56a74a112e81fc4784b28bfbf0cdacc882c262a5c28db6d1.NewDstDevPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dsum()(*ib8172c174b7898498cd0e4fac1ea4c3847990fe6f96023962f3ed4c8032776f2.DsumRequestBuilder) {
    return ib8172c174b7898498cd0e4fac1ea4c3847990fe6f96023962f3ed4c8032776f2.NewDsumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Duration()(*ib21b0bfa06713e745699452b9806b6e49faafc31f2098f6348b6e710d4dcfadb.DurationRequestBuilder) {
    return ib21b0bfa06713e745699452b9806b6e49faafc31f2098f6348b6e710d4dcfadb.NewDurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Dvar()(*iab2d32928daa4a2afac27ad11c4719622a434897b5d5271728dc6194c74b327e.DvarRequestBuilder) {
    return iab2d32928daa4a2afac27ad11c4719622a434897b5d5271728dc6194c74b327e.NewDvarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) DvarP()(*i7cc44ae074d8c78d8b236dfd967fb7230ff47ed7d77c49554bd34d6b20435c38.DvarPRequestBuilder) {
    return i7cc44ae074d8c78d8b236dfd967fb7230ff47ed7d77c49554bd34d6b20435c38.NewDvarPRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ecma_Ceiling()(*i18f1cdd18e81e750ed450d7b6aab0e82e96eff3d5624dc029aafd78d132463b1.Ecma_CeilingRequestBuilder) {
    return i18f1cdd18e81e750ed450d7b6aab0e82e96eff3d5624dc029aafd78d132463b1.NewEcma_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Edate()(*idada3620833af990b46a5c879e30430cb1543a2abebbc00f5c07e1c14a873b62.EdateRequestBuilder) {
    return idada3620833af990b46a5c879e30430cb1543a2abebbc00f5c07e1c14a873b62.NewEdateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Effect()(*i924df40fd16c8928384201c28d4fb809abdfbba3aeec4eae55afc75ec94685d7.EffectRequestBuilder) {
    return i924df40fd16c8928384201c28d4fb809abdfbba3aeec4eae55afc75ec94685d7.NewEffectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) EoMonth()(*i7c6ca0f53af8043bc128eda59012a093b72c668fd62d115a64b3d24d8478fcf4.EoMonthRequestBuilder) {
    return i7c6ca0f53af8043bc128eda59012a093b72c668fd62d115a64b3d24d8478fcf4.NewEoMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Erf()(*i0e88fa40bea4437ac83388050ce7e5ce9de47dc2f2ce412db5125c807029050d.ErfRequestBuilder) {
    return i0e88fa40bea4437ac83388050ce7e5ce9de47dc2f2ce412db5125c807029050d.NewErfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Erf_Precise()(*i3b6a93a263fef9b643219a486ba14d1a90efceba62eaddbec90cfa7ea241dcce.Erf_PreciseRequestBuilder) {
    return i3b6a93a263fef9b643219a486ba14d1a90efceba62eaddbec90cfa7ea241dcce.NewErf_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ErfC()(*i8263a2a05c5a495103660fe69cb0c9bd613de25e80e8b7679cf8eaf7fcfdb389.ErfCRequestBuilder) {
    return i8263a2a05c5a495103660fe69cb0c9bd613de25e80e8b7679cf8eaf7fcfdb389.NewErfCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ErfC_Precise()(*i3723d97baf5744837f171bd68f52e502d01d70a165e007898cd9489339572ae4.ErfC_PreciseRequestBuilder) {
    return i3723d97baf5744837f171bd68f52e502d01d70a165e007898cd9489339572ae4.NewErfC_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Error_Type()(*i3c61aac67e71e26228c623aa235d91a3be22189765dd40077b181ba6aa6ac074.Error_TypeRequestBuilder) {
    return i3c61aac67e71e26228c623aa235d91a3be22189765dd40077b181ba6aa6ac074.NewError_TypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Even()(*ic6226a166835f996dbcfdcba053ecd941c3435f27f4f8f9118b06373abae987e.EvenRequestBuilder) {
    return ic6226a166835f996dbcfdcba053ecd941c3435f27f4f8f9118b06373abae987e.NewEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Exact()(*i159670de4e02e4a74091f861f99a99ca141f71ec0b87a11f48af8d33b4829654.ExactRequestBuilder) {
    return i159670de4e02e4a74091f861f99a99ca141f71ec0b87a11f48af8d33b4829654.NewExactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Exp()(*i9272ce275af1b9564f11a7d1e165b408dc842cf4d755acf8576501d96eef63a0.ExpRequestBuilder) {
    return i9272ce275af1b9564f11a7d1e165b408dc842cf4d755acf8576501d96eef63a0.NewExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Expon_Dist()(*if5c0edc21dffeb380c263c85352bea05578a6349dd520aeb5d85a35c845c2645.Expon_DistRequestBuilder) {
    return if5c0edc21dffeb380c263c85352bea05578a6349dd520aeb5d85a35c845c2645.NewExpon_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Dist()(*if9bd3673059c5d91d6a5844a0702b4022748696d7caaebd69431763ab33549c3.F_DistRequestBuilder) {
    return if9bd3673059c5d91d6a5844a0702b4022748696d7caaebd69431763ab33549c3.NewF_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Dist_RT()(*i2fe518f7ce30627483c22947ac5fae593cbb1137e090d90d29067f66da898e01.F_Dist_RTRequestBuilder) {
    return i2fe518f7ce30627483c22947ac5fae593cbb1137e090d90d29067f66da898e01.NewF_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Inv()(*ic24f007baaa8f53b154215e49485e0157bb841692afa3867ccb2c9c4ae61e022.F_InvRequestBuilder) {
    return ic24f007baaa8f53b154215e49485e0157bb841692afa3867ccb2c9c4ae61e022.NewF_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) F_Inv_RT()(*ide313f7c15c109ec4520431747e1c1ad59b98641175faed943944e8b90b4ff9e.F_Inv_RTRequestBuilder) {
    return ide313f7c15c109ec4520431747e1c1ad59b98641175faed943944e8b90b4ff9e.NewF_Inv_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fact()(*idd96fd773b6cbeff56d5db250cfef3d2fab48aab4d11ed9c706c56d6489aa890.FactRequestBuilder) {
    return idd96fd773b6cbeff56d5db250cfef3d2fab48aab4d11ed9c706c56d6489aa890.NewFactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FactDouble()(*i01096f2b34ec8a2f8a1148bdb674f5bae59ae8e931d384bab99065956d16489d.FactDoubleRequestBuilder) {
    return i01096f2b34ec8a2f8a1148bdb674f5bae59ae8e931d384bab99065956d16489d.NewFactDoubleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) False()(*i61d4785a8acc4838fe5a5dc4bb2982e512218f586c0c26eb046d65eb08247df0.FalseRequestBuilder) {
    return i61d4785a8acc4838fe5a5dc4bb2982e512218f586c0c26eb046d65eb08247df0.NewFalseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Find()(*id66d330b86ce1c22be92754c31257b7e7fb4c39eebe39e086b3ab2833fc30b24.FindRequestBuilder) {
    return id66d330b86ce1c22be92754c31257b7e7fb4c39eebe39e086b3ab2833fc30b24.NewFindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FindB()(*id3bd2821ec8246f8cc97afc48556fce63ae66cd7db2cad74efec2a2fcdd8134b.FindBRequestBuilder) {
    return id3bd2821ec8246f8cc97afc48556fce63ae66cd7db2cad74efec2a2fcdd8134b.NewFindBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fisher()(*i3b889c48e56bbf6a183fe19f9853c22829238e3a7d488ba07cae69afed667f9f.FisherRequestBuilder) {
    return i3b889c48e56bbf6a183fe19f9853c22829238e3a7d488ba07cae69afed667f9f.NewFisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) FisherInv()(*i1baed36ad116a1d420e1a7d605ebfacb8a1686fd4c1db3e10ed48bf62c9e39c3.FisherInvRequestBuilder) {
    return i1baed36ad116a1d420e1a7d605ebfacb8a1686fd4c1db3e10ed48bf62c9e39c3.NewFisherInvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fixed()(*i3adc2f025c48fd250c29a04aa202729a054a29d0502ccf4ea72d4ec5a64635de.FixedRequestBuilder) {
    return i3adc2f025c48fd250c29a04aa202729a054a29d0502ccf4ea72d4ec5a64635de.NewFixedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Floor_Math()(*ie0a1929468b6da507dbfb84dc9baf3f3753b6c5c0d9032028149798c714291b8.Floor_MathRequestBuilder) {
    return ie0a1929468b6da507dbfb84dc9baf3f3753b6c5c0d9032028149798c714291b8.NewFloor_MathRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Floor_Precise()(*i86b37cd528c9b0fbbd43b89e4026ed456e78a6b499a1c8b89d4b4bb72448a3d6.Floor_PreciseRequestBuilder) {
    return i86b37cd528c9b0fbbd43b89e4026ed456e78a6b499a1c8b89d4b4bb72448a3d6.NewFloor_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fv()(*i27a16833edf596f6be428ee862a5d9faf3ff3a73862d538865760670037b2004.FvRequestBuilder) {
    return i27a16833edf596f6be428ee862a5d9faf3ff3a73862d538865760670037b2004.NewFvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Fvschedule()(*i0b6d1a4ef16cfd3065c4c6c4aa740a8f4ff77019773eb7e384205e813ce1468b.FvscheduleRequestBuilder) {
    return i0b6d1a4ef16cfd3065c4c6c4aa740a8f4ff77019773eb7e384205e813ce1468b.NewFvscheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma()(*i44c5b3ce8f085350158ab68112995ed6e115d0a6ee80c6b9d8d74e39fdd4f26c.GammaRequestBuilder) {
    return i44c5b3ce8f085350158ab68112995ed6e115d0a6ee80c6b9d8d74e39fdd4f26c.NewGammaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma_Dist()(*i71bb5e68a9c6004f9feea3f5b87729cdae98a1ab46518322bb458595067046f9.Gamma_DistRequestBuilder) {
    return i71bb5e68a9c6004f9feea3f5b87729cdae98a1ab46518322bb458595067046f9.NewGamma_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gamma_Inv()(*id42b4bcbb523a4bfd99220e19da63a21928fd9e1abc900d4fe74791d40a7f44d.Gamma_InvRequestBuilder) {
    return id42b4bcbb523a4bfd99220e19da63a21928fd9e1abc900d4fe74791d40a7f44d.NewGamma_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GammaLn()(*i5f73d2d1a7d4781f4525b034525e9df3deb205e8e4ac758982e431dc8eb9e284.GammaLnRequestBuilder) {
    return i5f73d2d1a7d4781f4525b034525e9df3deb205e8e4ac758982e431dc8eb9e284.NewGammaLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GammaLn_Precise()(*ie004a57ca723f01f9805f68e4c7807c360c84b1420011ff08a32ec2e039ce9ed.GammaLn_PreciseRequestBuilder) {
    return ie004a57ca723f01f9805f68e4c7807c360c84b1420011ff08a32ec2e039ce9ed.NewGammaLn_PreciseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gauss()(*ic4d8e49ce2a6532fcbbf6f64f180cdde43079498aab024848e5f8284fb6d197f.GaussRequestBuilder) {
    return ic4d8e49ce2a6532fcbbf6f64f180cdde43079498aab024848e5f8284fb6d197f.NewGaussRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Gcd()(*i317145d7cb1b1dd1c7a24b6b155a43af039f3cdaafcab3a3758411b12c50f560.GcdRequestBuilder) {
    return i317145d7cb1b1dd1c7a24b6b155a43af039f3cdaafcab3a3758411b12c50f560.NewGcdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GeoMean()(*ia73960e7b4da95c25edd07615a5587c36efa222e28463006ebf38230600502aa.GeoMeanRequestBuilder) {
    return ia73960e7b4da95c25edd07615a5587c36efa222e28463006ebf38230600502aa.NewGeoMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) GeStep()(*i3cbb023d286ddbc3ab8433a36661c912b2091cf8e12c0464bf56e7fddaa70c16.GeStepRequestBuilder) {
    return i3cbb023d286ddbc3ab8433a36661c912b2091cf8e12c0464bf56e7fddaa70c16.NewGeStepRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get functions from workbooks
func (m *FunctionsRequestBuilder) Get(options *FunctionsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFunctions, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookFunctions() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookFunctions), nil
}
func (m *FunctionsRequestBuilder) HarMean()(*i231393df28bed923560304a9df79e689e7d63cdf353ea10aa4965010bbb27b59.HarMeanRequestBuilder) {
    return i231393df28bed923560304a9df79e689e7d63cdf353ea10aa4965010bbb27b59.NewHarMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Bin()(*id3eb4904482ac31c14e02e367462325b3f88310776cbf05b20218c79ee5e8dcb.Hex2BinRequestBuilder) {
    return id3eb4904482ac31c14e02e367462325b3f88310776cbf05b20218c79ee5e8dcb.NewHex2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Dec()(*ie4a5e796f7fa374cef652c75e70f5a1b1f87110404de05bc90276b31af680f01.Hex2DecRequestBuilder) {
    return ie4a5e796f7fa374cef652c75e70f5a1b1f87110404de05bc90276b31af680f01.NewHex2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hex2Oct()(*i7322c133f56d7c2d62c17303e5865ee7fbe2b4b417b832ab2cd4cc88a2be8ccd.Hex2OctRequestBuilder) {
    return i7322c133f56d7c2d62c17303e5865ee7fbe2b4b417b832ab2cd4cc88a2be8ccd.NewHex2OctRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hlookup()(*i3122bd5b6a80eb07305ce98d54e54b51fb400d53e589e1c9177c2dca1b0e498c.HlookupRequestBuilder) {
    return i3122bd5b6a80eb07305ce98d54e54b51fb400d53e589e1c9177c2dca1b0e498c.NewHlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hour()(*i9a186f78fe48f3a957442b1ed17fccca4928be511af537fbf762cc585b3be753.HourRequestBuilder) {
    return i9a186f78fe48f3a957442b1ed17fccca4928be511af537fbf762cc585b3be753.NewHourRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Hyperlink()(*ib6796fab19a8e38b93a3039b8faeebbb4bb49035349deb590caa195682ab55b1.HyperlinkRequestBuilder) {
    return ib6796fab19a8e38b93a3039b8faeebbb4bb49035349deb590caa195682ab55b1.NewHyperlinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) HypGeom_Dist()(*id8942bdf6d4be967967837cf31114bfc0c30976ed4db096b9a586023b010d8c7.HypGeom_DistRequestBuilder) {
    return id8942bdf6d4be967967837cf31114bfc0c30976ed4db096b9a586023b010d8c7.NewHypGeom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) If_escaped()(*i0453e21d45f8ed2aab3a312355a9cdffae8562a8747da1d10b0964833762a270.IfRequestBuilder) {
    return i0453e21d45f8ed2aab3a312355a9cdffae8562a8747da1d10b0964833762a270.NewIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImAbs()(*ida2deeed82524cc3dfb4db9745f9ac68bd6bd46ca0996c86cc6552be00df989d.ImAbsRequestBuilder) {
    return ida2deeed82524cc3dfb4db9745f9ac68bd6bd46ca0996c86cc6552be00df989d.NewImAbsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Imaginary()(*i11a43b6060df9542091805a286115ebb955c1aa0d49662e197fdf6a4b2e50e4c.ImaginaryRequestBuilder) {
    return i11a43b6060df9542091805a286115ebb955c1aa0d49662e197fdf6a4b2e50e4c.NewImaginaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImArgument()(*i728699d88374e91f88187e481684864d04eab0cea4a217e8854f8b3a2723f232.ImArgumentRequestBuilder) {
    return i728699d88374e91f88187e481684864d04eab0cea4a217e8854f8b3a2723f232.NewImArgumentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImConjugate()(*i3be8431b385db6978cd20fc820abcefaaa7a43b0bd9794564078d5be763a9937.ImConjugateRequestBuilder) {
    return i3be8431b385db6978cd20fc820abcefaaa7a43b0bd9794564078d5be763a9937.NewImConjugateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCos()(*i80a969686b2a16a658360a90a8280dface60a263cccd4ac61782a7ab3aa075c6.ImCosRequestBuilder) {
    return i80a969686b2a16a658360a90a8280dface60a263cccd4ac61782a7ab3aa075c6.NewImCosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCosh()(*i52ae3aaa09a8851ecd5a25513977e69abde753c7b68342bc599ec5609538490e.ImCoshRequestBuilder) {
    return i52ae3aaa09a8851ecd5a25513977e69abde753c7b68342bc599ec5609538490e.NewImCoshRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCot()(*if579faccaec7d615e7f4e3259a551dcb170851c1b18626ed6b952606d4f270c3.ImCotRequestBuilder) {
    return if579faccaec7d615e7f4e3259a551dcb170851c1b18626ed6b952606d4f270c3.NewImCotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCsc()(*ic304fed62fa3d5a1cc95fd13ab3f2a1fde03b29af58ae54b9e8a39e294df5989.ImCscRequestBuilder) {
    return ic304fed62fa3d5a1cc95fd13ab3f2a1fde03b29af58ae54b9e8a39e294df5989.NewImCscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImCsch()(*i4c51dd2f6f6f5160da6053cc1291ae4a43ad7af4dbef329a7b3f7fa4d282750d.ImCschRequestBuilder) {
    return i4c51dd2f6f6f5160da6053cc1291ae4a43ad7af4dbef329a7b3f7fa4d282750d.NewImCschRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImDiv()(*i17c08c29b663ddc5369e52b456d6e68dba92c794e5a5b6ddbc0e076bfbd015f3.ImDivRequestBuilder) {
    return i17c08c29b663ddc5369e52b456d6e68dba92c794e5a5b6ddbc0e076bfbd015f3.NewImDivRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImExp()(*i8dce8a2f51873024131884e899491f95d9d6c9304be7c61361d0f9d33abfb188.ImExpRequestBuilder) {
    return i8dce8a2f51873024131884e899491f95d9d6c9304be7c61361d0f9d33abfb188.NewImExpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLn()(*i4f9aa0e6ed219c6f029aab669ac0680c2fceffd3c383d5478af4b9fddd950513.ImLnRequestBuilder) {
    return i4f9aa0e6ed219c6f029aab669ac0680c2fceffd3c383d5478af4b9fddd950513.NewImLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLog10()(*i3064bbf2a3c2deacd05d96d8e11b79f271f609d4a294f5cae347ed3c910c0696.ImLog10RequestBuilder) {
    return i3064bbf2a3c2deacd05d96d8e11b79f271f609d4a294f5cae347ed3c910c0696.NewImLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImLog2()(*i87768552c7fd50ad36dadf66d75df68add33f80dca108d619208d0955a2311ea.ImLog2RequestBuilder) {
    return i87768552c7fd50ad36dadf66d75df68add33f80dca108d619208d0955a2311ea.NewImLog2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImPower()(*id01758ffe2f0f268965d1810951affba607e2846a4d38de06c90fa826e0ca646.ImPowerRequestBuilder) {
    return id01758ffe2f0f268965d1810951affba607e2846a4d38de06c90fa826e0ca646.NewImPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImProduct()(*ie84ebf6c1f6d5fd5590422fb45bc8d7ee6a943ca1a287247547419e876d52501.ImProductRequestBuilder) {
    return ie84ebf6c1f6d5fd5590422fb45bc8d7ee6a943ca1a287247547419e876d52501.NewImProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImReal()(*i6a81c49a0cbd281c42b4e3a543ce420c090faa981e28dac41ae63e33eddc10a5.ImRealRequestBuilder) {
    return i6a81c49a0cbd281c42b4e3a543ce420c090faa981e28dac41ae63e33eddc10a5.NewImRealRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSec()(*ie347231f0f0a8ce2cadb635d237561360591f79cd4435bab6bf441298258887e.ImSecRequestBuilder) {
    return ie347231f0f0a8ce2cadb635d237561360591f79cd4435bab6bf441298258887e.NewImSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSech()(*ifd522acfa264f2b97ee138f4abcf7d087cfe0cd47eb1c57d8121172e2f5db7ee.ImSechRequestBuilder) {
    return ifd522acfa264f2b97ee138f4abcf7d087cfe0cd47eb1c57d8121172e2f5db7ee.NewImSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSin()(*id2e785e1992b73d2862c5f6713800dca6a10b0279f7dfa2390cd3de87c98a982.ImSinRequestBuilder) {
    return id2e785e1992b73d2862c5f6713800dca6a10b0279f7dfa2390cd3de87c98a982.NewImSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSinh()(*idc55d61c4fe53f23e33f4d0c3e62391402534d7c47cb00c2d867a340f4ad680e.ImSinhRequestBuilder) {
    return idc55d61c4fe53f23e33f4d0c3e62391402534d7c47cb00c2d867a340f4ad680e.NewImSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSqrt()(*i09a091e227ad1cadaf693bfcf5cabc15f2a578b85013c5809d603aed05bc797f.ImSqrtRequestBuilder) {
    return i09a091e227ad1cadaf693bfcf5cabc15f2a578b85013c5809d603aed05bc797f.NewImSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSub()(*i6549418b40dc32ee3c2bda4721e2013b43682373bdc47a7e919d206cfd6bc36c.ImSubRequestBuilder) {
    return i6549418b40dc32ee3c2bda4721e2013b43682373bdc47a7e919d206cfd6bc36c.NewImSubRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImSum()(*i978242c1185a2557dcea381fa50ed6e2710ddbf3163a20d5e785bc7363a70154.ImSumRequestBuilder) {
    return i978242c1185a2557dcea381fa50ed6e2710ddbf3163a20d5e785bc7363a70154.NewImSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ImTan()(*if7e94ba699066e6986bb11cb8c1fe185dfeac6238c220ba28a74982e8abf6364.ImTanRequestBuilder) {
    return if7e94ba699066e6986bb11cb8c1fe185dfeac6238c220ba28a74982e8abf6364.NewImTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Int()(*i41988a8388286f2ac1bbc4d6af302f9d1f081f6b63ee9e8fa3682d84c8fc86c5.IntRequestBuilder) {
    return i41988a8388286f2ac1bbc4d6af302f9d1f081f6b63ee9e8fa3682d84c8fc86c5.NewIntRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IntRate()(*i1240f098751c9c3ede3953da47ed5f1592e3bcdea4a678ff30bac2fa804160d2.IntRateRequestBuilder) {
    return i1240f098751c9c3ede3953da47ed5f1592e3bcdea4a678ff30bac2fa804160d2.NewIntRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ipmt()(*i7b45dd95f5c1244495c5095b798515bb00c86ca2ea5024d3861f84b6d993c596.IpmtRequestBuilder) {
    return i7b45dd95f5c1244495c5095b798515bb00c86ca2ea5024d3861f84b6d993c596.NewIpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Irr()(*i5c3d49a1ba7296574afa46cd49edc3ad92da6e76655ea5a04f1913bcbf28aa69.IrrRequestBuilder) {
    return i5c3d49a1ba7296574afa46cd49edc3ad92da6e76655ea5a04f1913bcbf28aa69.NewIrrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsErr()(*ia10e03feb00942388c2542ed5b8c978c7e5851279ff4dedaf8eef4dec3e83b3e.IsErrRequestBuilder) {
    return ia10e03feb00942388c2542ed5b8c978c7e5851279ff4dedaf8eef4dec3e83b3e.NewIsErrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsError()(*iaf0f35ff858eeae62b0dc3e885a456c3cb25f0ad94b3d748b8e25f2832546481.IsErrorRequestBuilder) {
    return iaf0f35ff858eeae62b0dc3e885a456c3cb25f0ad94b3d748b8e25f2832546481.NewIsErrorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsEven()(*ibe04576cf9d5d63b6da0268d3cc3055d39a57721c71cea7dbf5f42c48b66c27f.IsEvenRequestBuilder) {
    return ibe04576cf9d5d63b6da0268d3cc3055d39a57721c71cea7dbf5f42c48b66c27f.NewIsEvenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsFormula()(*i6aa6e261dc646277372246f6ec53c74cf1d75e098b9e0f81a7f87dda1919e74e.IsFormulaRequestBuilder) {
    return i6aa6e261dc646277372246f6ec53c74cf1d75e098b9e0f81a7f87dda1919e74e.NewIsFormulaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsLogical()(*i4e3aa948ee53f2131c385f041992741d89d3bca88109722ca599726083d91928.IsLogicalRequestBuilder) {
    return i4e3aa948ee53f2131c385f041992741d89d3bca88109722ca599726083d91928.NewIsLogicalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNA()(*iecec2b18f5b5d30ea473c416117fa75c41083c7e9471e526028e1b7dfbb1e37d.IsNARequestBuilder) {
    return iecec2b18f5b5d30ea473c416117fa75c41083c7e9471e526028e1b7dfbb1e37d.NewIsNARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNonText()(*i34dc66540bc7342bc45b426f0e1fedcbafa46564955a9848253415d4ac46b4d5.IsNonTextRequestBuilder) {
    return i34dc66540bc7342bc45b426f0e1fedcbafa46564955a9848253415d4ac46b4d5.NewIsNonTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsNumber()(*idc0b79ac63ba2c8b3402dc56406f039561bf704065726033d98f42081b94d7b2.IsNumberRequestBuilder) {
    return idc0b79ac63ba2c8b3402dc56406f039561bf704065726033d98f42081b94d7b2.NewIsNumberRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Iso_Ceiling()(*idff66a72240da9492b8b2fd994dc3c423b6620bd778050d7d19b21b3b50d24f2.Iso_CeilingRequestBuilder) {
    return idff66a72240da9492b8b2fd994dc3c423b6620bd778050d7d19b21b3b50d24f2.NewIso_CeilingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsOdd()(*ieb56fecb0739f2e710e721eb6508a3c2c25c0bbe74af47a66d0379daabee2fdc.IsOddRequestBuilder) {
    return ieb56fecb0739f2e710e721eb6508a3c2c25c0bbe74af47a66d0379daabee2fdc.NewIsOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsoWeekNum()(*icd9f65c4c8dadcdc84336ea37b49c2de9e0249e1c5999cd4e34a373bfd3190df.IsoWeekNumRequestBuilder) {
    return icd9f65c4c8dadcdc84336ea37b49c2de9e0249e1c5999cd4e34a373bfd3190df.NewIsoWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ispmt()(*i1a5c221c73ce54ab21f5cd625811db70165d0058939149337b4464d4fb657cfa.IspmtRequestBuilder) {
    return i1a5c221c73ce54ab21f5cd625811db70165d0058939149337b4464d4fb657cfa.NewIspmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Isref()(*ic17145b99fdaba106e94f63d1fc5b6b0611a0fc6c1c5528f3b1bf69810a32a01.IsrefRequestBuilder) {
    return ic17145b99fdaba106e94f63d1fc5b6b0611a0fc6c1c5528f3b1bf69810a32a01.NewIsrefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) IsText()(*ia1e08a8c5ea700baa7a7adb97501abab922a41f60100dd6d938863a562bc973f.IsTextRequestBuilder) {
    return ia1e08a8c5ea700baa7a7adb97501abab922a41f60100dd6d938863a562bc973f.NewIsTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Kurt()(*i77fc9cd08813563fe05c08fb8c1d0c196d2acc8c15e091cbff3713a5326d51f8.KurtRequestBuilder) {
    return i77fc9cd08813563fe05c08fb8c1d0c196d2acc8c15e091cbff3713a5326d51f8.NewKurtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Large()(*ie1788a287e583aaee3f2bb77e8e1a468fa3b2a0e293faef36a82104552219f03.LargeRequestBuilder) {
    return ie1788a287e583aaee3f2bb77e8e1a468fa3b2a0e293faef36a82104552219f03.NewLargeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lcm()(*idc0bb38ae510a3b852b847ac7d5fd2d2def09000af4b33633be4fe0fbf2c5d37.LcmRequestBuilder) {
    return idc0bb38ae510a3b852b847ac7d5fd2d2def09000af4b33633be4fe0fbf2c5d37.NewLcmRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Left()(*i6d9183c34f2b76db26c0d25b436c794deed367981ac978d37de88c2166b7aec6.LeftRequestBuilder) {
    return i6d9183c34f2b76db26c0d25b436c794deed367981ac978d37de88c2166b7aec6.NewLeftRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Leftb()(*i64f90fa888a48017c96edf4e2d4540dc46e262036416338183cd5fe46e76fa5f.LeftbRequestBuilder) {
    return i64f90fa888a48017c96edf4e2d4540dc46e262036416338183cd5fe46e76fa5f.NewLeftbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Len()(*id742f89db038135ff55db923244279c7da2917bf7156592a277a7660e0bf2292.LenRequestBuilder) {
    return id742f89db038135ff55db923244279c7da2917bf7156592a277a7660e0bf2292.NewLenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lenb()(*i9403886d8363871971936746538ecefa7791fc349af497df5bc1277c5783d44d.LenbRequestBuilder) {
    return i9403886d8363871971936746538ecefa7791fc349af497df5bc1277c5783d44d.NewLenbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ln()(*ib455b5a90ec038d50df092019cb07bc6983d528925397d7bc0fb1a3acebf820d.LnRequestBuilder) {
    return ib455b5a90ec038d50df092019cb07bc6983d528925397d7bc0fb1a3acebf820d.NewLnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Log()(*idce3c3d2910ab657d92dfa59ea57a85b9a9134fbd6ad6a2976fdb9e96347b232.LogRequestBuilder) {
    return idce3c3d2910ab657d92dfa59ea57a85b9a9134fbd6ad6a2976fdb9e96347b232.NewLogRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Log10()(*i82ad818b67403f8f447c71dc2a945eaf3946b28652d715954fd417474a3bb758.Log10RequestBuilder) {
    return i82ad818b67403f8f447c71dc2a945eaf3946b28652d715954fd417474a3bb758.NewLog10RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) LogNorm_Dist()(*ia6288640436d5fa51554e6bb7265db0ae494d329885ecd45cd9de18927e25160.LogNorm_DistRequestBuilder) {
    return ia6288640436d5fa51554e6bb7265db0ae494d329885ecd45cd9de18927e25160.NewLogNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) LogNorm_Inv()(*i65fb0cdaf5c4b3936c8b0cfcc9c69c364c6f66cbf16cfac1a196908e8e0af0e3.LogNorm_InvRequestBuilder) {
    return i65fb0cdaf5c4b3936c8b0cfcc9c69c364c6f66cbf16cfac1a196908e8e0af0e3.NewLogNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lookup()(*idcd01cb21b27b9b15225c6c93e8dc2d29a633a4bca76b7dea84c298644ef1375.LookupRequestBuilder) {
    return idcd01cb21b27b9b15225c6c93e8dc2d29a633a4bca76b7dea84c298644ef1375.NewLookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Lower()(*i3e27b43c227d8f50f706ff6461aa17d93a7add1a68aaf7a5a7a31b13a5d05d09.LowerRequestBuilder) {
    return i3e27b43c227d8f50f706ff6461aa17d93a7add1a68aaf7a5a7a31b13a5d05d09.NewLowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Match()(*ic42d36d02b28aa18548e0d940d1dff43b455a6503f32bec64e0b0d15698d1757.MatchRequestBuilder) {
    return ic42d36d02b28aa18548e0d940d1dff43b455a6503f32bec64e0b0d15698d1757.NewMatchRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Max()(*i5c087781b52d9dcd7ec53ed914b813a5ad2a143bb9b4d06a86ae1d23917eb3f3.MaxRequestBuilder) {
    return i5c087781b52d9dcd7ec53ed914b813a5ad2a143bb9b4d06a86ae1d23917eb3f3.NewMaxRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MaxA()(*i7b57475004e70295b93ef2615abeb481c477997c718c386d064208924ccfc6a3.MaxARequestBuilder) {
    return i7b57475004e70295b93ef2615abeb481c477997c718c386d064208924ccfc6a3.NewMaxARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mduration()(*i4ab58e6b08f32389c6302b37da6352ef4e50c7099f63e72b1eec5408402a3b25.MdurationRequestBuilder) {
    return i4ab58e6b08f32389c6302b37da6352ef4e50c7099f63e72b1eec5408402a3b25.NewMdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Median()(*ic72126fb1b037996cabc6adbf309508be5276a8796ba813d30b44410593639cd.MedianRequestBuilder) {
    return ic72126fb1b037996cabc6adbf309508be5276a8796ba813d30b44410593639cd.NewMedianRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mid()(*i3e7338d395a2213f344bce5f8bd106c48c9613368c44472b4a68c636a47267e9.MidRequestBuilder) {
    return i3e7338d395a2213f344bce5f8bd106c48c9613368c44472b4a68c636a47267e9.NewMidRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Midb()(*idad0bb41b082d0269aa802164dce05c2a807fbc97a7ebd6069b225b5f6b64a39.MidbRequestBuilder) {
    return idad0bb41b082d0269aa802164dce05c2a807fbc97a7ebd6069b225b5f6b64a39.NewMidbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Min()(*i01ade93e966b7ef3ad3978d2f1d90704ba31d11a2d416ea326f25b442162d957.MinRequestBuilder) {
    return i01ade93e966b7ef3ad3978d2f1d90704ba31d11a2d416ea326f25b442162d957.NewMinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MinA()(*i45a317c4a16779aa0191decfdd8eafccaee07defa32b4f71d86074431bc29a0a.MinARequestBuilder) {
    return i45a317c4a16779aa0191decfdd8eafccaee07defa32b4f71d86074431bc29a0a.NewMinARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Minute()(*i55563060001cd02d8ec59c905f1cbcdacb601fa67184d73a2799ca52783d5423.MinuteRequestBuilder) {
    return i55563060001cd02d8ec59c905f1cbcdacb601fa67184d73a2799ca52783d5423.NewMinuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mirr()(*i7b3b5d10f9dddfffc344d3ea44785438bf332088789e8512860cbec2800a0409.MirrRequestBuilder) {
    return i7b3b5d10f9dddfffc344d3ea44785438bf332088789e8512860cbec2800a0409.NewMirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mod()(*ib479884246004b472c3710d8868d87ab00d93948463b7e5be5a8dc6329641309.ModRequestBuilder) {
    return ib479884246004b472c3710d8868d87ab00d93948463b7e5be5a8dc6329641309.NewModRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Month()(*iae7b528f9d9de970748d7d77bcefdfd78cd0d575b555bb4ed9f9bdead7e1ada8.MonthRequestBuilder) {
    return iae7b528f9d9de970748d7d77bcefdfd78cd0d575b555bb4ed9f9bdead7e1ada8.NewMonthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Mround()(*i857d6144810367f6d7cd741fdc5a24ee633e1241ee716fc3c34ec46d66cd353a.MroundRequestBuilder) {
    return i857d6144810367f6d7cd741fdc5a24ee633e1241ee716fc3c34ec46d66cd353a.NewMroundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) MultiNomial()(*ic9303038a56639aca96d0a80154f111d9ac533dbacbeb52f9b2b41a622969d39.MultiNomialRequestBuilder) {
    return ic9303038a56639aca96d0a80154f111d9ac533dbacbeb52f9b2b41a622969d39.NewMultiNomialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) N()(*ibaa8c2b8d48ff4b0848a07a1597e4d8c0842561ec4c6e0608fe24e0754f36765.NRequestBuilder) {
    return ibaa8c2b8d48ff4b0848a07a1597e4d8c0842561ec4c6e0608fe24e0754f36765.NewNRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Na()(*ifc5534ed3042a1a7f49eba82e143c4978c0db242a2657d9bbe9576e636343d22.NaRequestBuilder) {
    return ifc5534ed3042a1a7f49eba82e143c4978c0db242a2657d9bbe9576e636343d22.NewNaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NegBinom_Dist()(*ieb15ac389058ef18a973e324edaff73830c8dddc6f6de9fd9d1ce8e074fdf8c9.NegBinom_DistRequestBuilder) {
    return ieb15ac389058ef18a973e324edaff73830c8dddc6f6de9fd9d1ce8e074fdf8c9.NewNegBinom_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NetworkDays()(*id7fd94c4c6345dfca7b605745f22ab04c9acec6f4dc9d1b1e80e53f35f71ae50.NetworkDaysRequestBuilder) {
    return id7fd94c4c6345dfca7b605745f22ab04c9acec6f4dc9d1b1e80e53f35f71ae50.NewNetworkDaysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NetworkDays_Intl()(*i5873f11c8e7d2e163103011598914ed1291b59404c3dffe95bdfea5a9837dc6b.NetworkDays_IntlRequestBuilder) {
    return i5873f11c8e7d2e163103011598914ed1291b59404c3dffe95bdfea5a9837dc6b.NewNetworkDays_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Nominal()(*i02e07df104c71232444b95fc5b1e62f0293d740e0b69b73e13add9e338bb0d96.NominalRequestBuilder) {
    return i02e07df104c71232444b95fc5b1e62f0293d740e0b69b73e13add9e338bb0d96.NewNominalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_Dist()(*i6af5ea8a5eae6ce0d467e8d79bd8a793918c81b05535fac4d218577b2f1033c0.Norm_DistRequestBuilder) {
    return i6af5ea8a5eae6ce0d467e8d79bd8a793918c81b05535fac4d218577b2f1033c0.NewNorm_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_Inv()(*i99d62393032ec4efdcaacf249e0301202697d7b6fba7914bf39098553f84803b.Norm_InvRequestBuilder) {
    return i99d62393032ec4efdcaacf249e0301202697d7b6fba7914bf39098553f84803b.NewNorm_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_S_Dist()(*i19464df7609672eb389ea5c0b7823da82b500d02a65e21c56836a21fdda12a5e.Norm_S_DistRequestBuilder) {
    return i19464df7609672eb389ea5c0b7823da82b500d02a65e21c56836a21fdda12a5e.NewNorm_S_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Norm_S_Inv()(*i2db9198ed6278fd3b0fa28899357367386848879a0248c9bbb758b03bb05cdf7.Norm_S_InvRequestBuilder) {
    return i2db9198ed6278fd3b0fa28899357367386848879a0248c9bbb758b03bb05cdf7.NewNorm_S_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Not()(*ifd1dbe5d2897f560ce7816e3d3dc4208b6ed57b5ae5446df8d93e5d261c102fb.NotRequestBuilder) {
    return ifd1dbe5d2897f560ce7816e3d3dc4208b6ed57b5ae5446df8d93e5d261c102fb.NewNotRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Now()(*ice7e4f6925eb7d4acd3adce5b6e13f99707d1adf6c6737526bd71a868cd56e22.NowRequestBuilder) {
    return ice7e4f6925eb7d4acd3adce5b6e13f99707d1adf6c6737526bd71a868cd56e22.NewNowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Nper()(*i805b9eb159e70babbb463fdade294b172fd8126ddea9b57a1899afa3be216b61.NperRequestBuilder) {
    return i805b9eb159e70babbb463fdade294b172fd8126ddea9b57a1899afa3be216b61.NewNperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Npv()(*ie496f58abf287e229fa39afb5ddb533e54ea59897e8c4cac11239aee71cb2ff5.NpvRequestBuilder) {
    return ie496f58abf287e229fa39afb5ddb533e54ea59897e8c4cac11239aee71cb2ff5.NewNpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) NumberValue()(*i513653df7d3cb928e656cb4c60e76d7c5355f0856052209394f8da8b64e5e225.NumberValueRequestBuilder) {
    return i513653df7d3cb928e656cb4c60e76d7c5355f0856052209394f8da8b64e5e225.NewNumberValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Bin()(*i051f24a077b7189aa242fd3fe1798aebf0b3261717546eba2dc8be57c7d50c0d.Oct2BinRequestBuilder) {
    return i051f24a077b7189aa242fd3fe1798aebf0b3261717546eba2dc8be57c7d50c0d.NewOct2BinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Dec()(*ibdd44510acdea53f5b37d091d420e4d288afe3820de5bbb0acef871af1977257.Oct2DecRequestBuilder) {
    return ibdd44510acdea53f5b37d091d420e4d288afe3820de5bbb0acef871af1977257.NewOct2DecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Oct2Hex()(*i6421c408950fba8073283495764103519c7b83b532bec18bb40598c0dbb1dce5.Oct2HexRequestBuilder) {
    return i6421c408950fba8073283495764103519c7b83b532bec18bb40598c0dbb1dce5.NewOct2HexRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Odd()(*i81f182730936652d80d21614c48a4de559dd8e4a172aef9bd71b56f481f09e64.OddRequestBuilder) {
    return i81f182730936652d80d21614c48a4de559dd8e4a172aef9bd71b56f481f09e64.NewOddRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddFPrice()(*i45864aa1e5af9a8d5fef6d7689757c11aa4f550e8509bd22356840584e8a4f74.OddFPriceRequestBuilder) {
    return i45864aa1e5af9a8d5fef6d7689757c11aa4f550e8509bd22356840584e8a4f74.NewOddFPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddFYield()(*i0bbe7782eee3501d57bb814a1a1c58ce8cf07efaa239ed9577f240d94d413bc2.OddFYieldRequestBuilder) {
    return i0bbe7782eee3501d57bb814a1a1c58ce8cf07efaa239ed9577f240d94d413bc2.NewOddFYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddLPrice()(*i2e9731503ea02b7ef4747c359c541a15c2b4c70bce0cc4ef82d5c00c080a72cf.OddLPriceRequestBuilder) {
    return i2e9731503ea02b7ef4747c359c541a15c2b4c70bce0cc4ef82d5c00c080a72cf.NewOddLPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) OddLYield()(*ia41376203552bcfaba0bc85c9235047123337e891d1c22d7936e78d36d871ed9.OddLYieldRequestBuilder) {
    return ia41376203552bcfaba0bc85c9235047123337e891d1c22d7936e78d36d871ed9.NewOddLYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Or()(*i5d6b82b7c343af5d6b219c51afe859aee013f35c73c2e618180f18bfda38f7a1.OrRequestBuilder) {
    return i5d6b82b7c343af5d6b219c51afe859aee013f35c73c2e618180f18bfda38f7a1.NewOrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property functions in workbooks
func (m *FunctionsRequestBuilder) Patch(options *FunctionsRequestBuilderPatchOptions)(error) {
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
func (m *FunctionsRequestBuilder) Pduration()(*i7f8ec93114a050419ca8698a92c1ed22b189436b0d42e9c5451c7abaa3dd95eb.PdurationRequestBuilder) {
    return i7f8ec93114a050419ca8698a92c1ed22b189436b0d42e9c5451c7abaa3dd95eb.NewPdurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Percentile_Exc()(*i6c3e7e3f69674bcdedacded58afe5875ffcee90414e278f5b8841357e5d64df7.Percentile_ExcRequestBuilder) {
    return i6c3e7e3f69674bcdedacded58afe5875ffcee90414e278f5b8841357e5d64df7.NewPercentile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Percentile_Inc()(*i67edd6c2573b42e032248aa937dbe9903d23f394d8ec053e7f97a11f2b38c3ea.Percentile_IncRequestBuilder) {
    return i67edd6c2573b42e032248aa937dbe9903d23f394d8ec053e7f97a11f2b38c3ea.NewPercentile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PercentRank_Exc()(*i4d5acd27fec5eb7998fbc0ba0103be8eb685259f312784459dcdc8791fbd4ab9.PercentRank_ExcRequestBuilder) {
    return i4d5acd27fec5eb7998fbc0ba0103be8eb685259f312784459dcdc8791fbd4ab9.NewPercentRank_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PercentRank_Inc()(*i405aebdde6e9c90bf532ba8258294f2f8162cedfeccf8c83ca2d1e4abea16a31.PercentRank_IncRequestBuilder) {
    return i405aebdde6e9c90bf532ba8258294f2f8162cedfeccf8c83ca2d1e4abea16a31.NewPercentRank_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Permut()(*id04facfe8d3bb3050bef6bf2c05e1933d3ac3f8a2bf4b8065a7b736b90e6bcd7.PermutRequestBuilder) {
    return id04facfe8d3bb3050bef6bf2c05e1933d3ac3f8a2bf4b8065a7b736b90e6bcd7.NewPermutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Permutationa()(*ia591b7807d5cbcdaf4b604db80322f5bcf8895f8f6566a04ba9b89dfaf1caf3f.PermutationaRequestBuilder) {
    return ia591b7807d5cbcdaf4b604db80322f5bcf8895f8f6566a04ba9b89dfaf1caf3f.NewPermutationaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Phi()(*i23d1cf8a00fdb2931404bd1ff8b0d447a926d8d3874700e2553064cf8c0d13e3.PhiRequestBuilder) {
    return i23d1cf8a00fdb2931404bd1ff8b0d447a926d8d3874700e2553064cf8c0d13e3.NewPhiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pi()(*i57abd1bbb332b1fb6f15dc1b9f5d038188388426d82da6b3ae052ab3713f7794.PiRequestBuilder) {
    return i57abd1bbb332b1fb6f15dc1b9f5d038188388426d82da6b3ae052ab3713f7794.NewPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pmt()(*i04b1ad703d7e9b5de1347bbe2e166dc427117fc2d5694354e996a14506873246.PmtRequestBuilder) {
    return i04b1ad703d7e9b5de1347bbe2e166dc427117fc2d5694354e996a14506873246.NewPmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Poisson_Dist()(*i51e821e10b850a9c97fc458dc827715fe679e8ba556de07fcb94e8afcf7924b7.Poisson_DistRequestBuilder) {
    return i51e821e10b850a9c97fc458dc827715fe679e8ba556de07fcb94e8afcf7924b7.NewPoisson_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Power()(*i2f1e41b93eba14f69f07e64f451b07a9d162c3b2b5ac0833f310b56e7dff93ce.PowerRequestBuilder) {
    return i2f1e41b93eba14f69f07e64f451b07a9d162c3b2b5ac0833f310b56e7dff93ce.NewPowerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Ppmt()(*idf877ceae99180689e8e9a8bf9eee09332189d3fd7e31643a899b0538cc0dcef.PpmtRequestBuilder) {
    return idf877ceae99180689e8e9a8bf9eee09332189d3fd7e31643a899b0538cc0dcef.NewPpmtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Price()(*ie5678a893191a28298bfcb69e8f697efffc2d356b0f9ee2c39a9e5b3e5bd485a.PriceRequestBuilder) {
    return ie5678a893191a28298bfcb69e8f697efffc2d356b0f9ee2c39a9e5b3e5bd485a.NewPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PriceDisc()(*ib27c742fbbd687c6ead89d689434bf79c2ca3ff53b03369668d84c142ddd5359.PriceDiscRequestBuilder) {
    return ib27c742fbbd687c6ead89d689434bf79c2ca3ff53b03369668d84c142ddd5359.NewPriceDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) PriceMat()(*ibbcf6f4e007d6d73f76fed9eba9506277256367d70978c6ffbeedde11713d6d0.PriceMatRequestBuilder) {
    return ibbcf6f4e007d6d73f76fed9eba9506277256367d70978c6ffbeedde11713d6d0.NewPriceMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Product()(*i6bd5f7f5f528d90cb5622517509fe271b78cb6337ca6bc94db79462be9bac8e9.ProductRequestBuilder) {
    return i6bd5f7f5f528d90cb5622517509fe271b78cb6337ca6bc94db79462be9bac8e9.NewProductRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Proper()(*i82ed0aa8211f2c0cb3c379fcec65435bc44073411288d8ff080fdcb136272548.ProperRequestBuilder) {
    return i82ed0aa8211f2c0cb3c379fcec65435bc44073411288d8ff080fdcb136272548.NewProperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Pv()(*iedadcc399469fc51ce774f195b317d33cd4324ca9dd16a18f898ade9fdbf4a02.PvRequestBuilder) {
    return iedadcc399469fc51ce774f195b317d33cd4324ca9dd16a18f898ade9fdbf4a02.NewPvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quartile_Exc()(*i6503d7704ccbe9c818734b31eb667f6181eacf1a7a022cd7f58f363dcf403849.Quartile_ExcRequestBuilder) {
    return i6503d7704ccbe9c818734b31eb667f6181eacf1a7a022cd7f58f363dcf403849.NewQuartile_ExcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quartile_Inc()(*i90d52a146df567b4dd6dfdabe26ec25d0096a30b2b054edbc0a038e81515513a.Quartile_IncRequestBuilder) {
    return i90d52a146df567b4dd6dfdabe26ec25d0096a30b2b054edbc0a038e81515513a.NewQuartile_IncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Quotient()(*ifc30dc22fa9fe32b7a51832bdd0cf2e35370072463fc48030297308894f031d7.QuotientRequestBuilder) {
    return ifc30dc22fa9fe32b7a51832bdd0cf2e35370072463fc48030297308894f031d7.NewQuotientRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Radians()(*i92fb305ff14f82f0506ecb4f2acbf0653dc257c9cbb95fb04bce2b4ce1fd900e.RadiansRequestBuilder) {
    return i92fb305ff14f82f0506ecb4f2acbf0653dc257c9cbb95fb04bce2b4ce1fd900e.NewRadiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rand()(*ifbe89311210db7e5aebd2c2a5941f6c6ee75028b80166be8dccd2d48242d4026.RandRequestBuilder) {
    return ifbe89311210db7e5aebd2c2a5941f6c6ee75028b80166be8dccd2d48242d4026.NewRandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RandBetween()(*i5376958355630f6f6240607052fdce472946f9e94ea3e8560aa27167dae3ef3f.RandBetweenRequestBuilder) {
    return i5376958355630f6f6240607052fdce472946f9e94ea3e8560aa27167dae3ef3f.NewRandBetweenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rank_Avg()(*i1de8619dd7518db3d99462880cfbf87a747b0bde4d139b30efb92a2897d26416.Rank_AvgRequestBuilder) {
    return i1de8619dd7518db3d99462880cfbf87a747b0bde4d139b30efb92a2897d26416.NewRank_AvgRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rank_Eq()(*i3baa4f1f70089ce805f6a5096e8b3a29c96246baf3605f030ef2951df8e61cdd.Rank_EqRequestBuilder) {
    return i3baa4f1f70089ce805f6a5096e8b3a29c96246baf3605f030ef2951df8e61cdd.NewRank_EqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rate()(*ibbfe64cfb60ca74a478178ca79042cf60ba03437dc03a754280148b9b3176587.RateRequestBuilder) {
    return ibbfe64cfb60ca74a478178ca79042cf60ba03437dc03a754280148b9b3176587.NewRateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Received()(*i544cc180c4727e4718f970d997bd828304c5c469cb353a04eb4136d9002fc7d0.ReceivedRequestBuilder) {
    return i544cc180c4727e4718f970d997bd828304c5c469cb353a04eb4136d9002fc7d0.NewReceivedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Replace()(*iaf6aad8ff5f9f5a94542995fec8311088ffaf94e57ba86fefad541089ed2a5e5.ReplaceRequestBuilder) {
    return iaf6aad8ff5f9f5a94542995fec8311088ffaf94e57ba86fefad541089ed2a5e5.NewReplaceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) ReplaceB()(*i68aeb1985b4e03322a4c461c36962d4d5d0c61538152110f09a942acaac1394a.ReplaceBRequestBuilder) {
    return i68aeb1985b4e03322a4c461c36962d4d5d0c61538152110f09a942acaac1394a.NewReplaceBRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rept()(*i225dd267832dae6eab68b4f07ab1629f61b9ae136b37950bc41a4eb2121048ff.ReptRequestBuilder) {
    return i225dd267832dae6eab68b4f07ab1629f61b9ae136b37950bc41a4eb2121048ff.NewReptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Right()(*ib147fed6b5db44580c7acf996a2806f0bb8515459362219dd8cade718c06172f.RightRequestBuilder) {
    return ib147fed6b5db44580c7acf996a2806f0bb8515459362219dd8cade718c06172f.NewRightRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rightb()(*i3f458892db0406f566cc4ae0bfc1ec462432e8d9275bb6e6cd964aaf774a1e54.RightbRequestBuilder) {
    return i3f458892db0406f566cc4ae0bfc1ec462432e8d9275bb6e6cd964aaf774a1e54.NewRightbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Roman()(*i37b034b37e59ca49f27a237e569deee9dd6f9020957c697c75489bbc106b3ccd.RomanRequestBuilder) {
    return i37b034b37e59ca49f27a237e569deee9dd6f9020957c697c75489bbc106b3ccd.NewRomanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Round()(*i7c36657c2cef60c6b31bd5fc09db6d2a130306d77e26c1d597d8361db50f0e59.RoundRequestBuilder) {
    return i7c36657c2cef60c6b31bd5fc09db6d2a130306d77e26c1d597d8361db50f0e59.NewRoundRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RoundDown()(*i1d9c84764ea0a52054b0a1ff763c3e328e31f451bbb49c623958cb42d31d56df.RoundDownRequestBuilder) {
    return i1d9c84764ea0a52054b0a1ff763c3e328e31f451bbb49c623958cb42d31d56df.NewRoundDownRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) RoundUp()(*ieac58a41e841549fbd8af3e54f87582f50250dca151693068af1909ecec35019.RoundUpRequestBuilder) {
    return ieac58a41e841549fbd8af3e54f87582f50250dca151693068af1909ecec35019.NewRoundUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rows()(*i00bd7d5caf4f643e4d966eb8a89a829df9dd5b0dbee54b766d2097ce7036aa2d.RowsRequestBuilder) {
    return i00bd7d5caf4f643e4d966eb8a89a829df9dd5b0dbee54b766d2097ce7036aa2d.NewRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Rri()(*i2193ffa9da564442c3818c80c3a3934f5f1ab209ed23eaa6af49b0fca52698ff.RriRequestBuilder) {
    return i2193ffa9da564442c3818c80c3a3934f5f1ab209ed23eaa6af49b0fca52698ff.NewRriRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sec()(*i01d278483174ac301bcbeac39516b56e97a8a5f1352e96fe34181f0ff94b1ee0.SecRequestBuilder) {
    return i01d278483174ac301bcbeac39516b56e97a8a5f1352e96fe34181f0ff94b1ee0.NewSecRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sech()(*i73debe587dffbb4f4db9d9bed9a0d4d4fecc4b1e819504a326cb68d7a674fd8c.SechRequestBuilder) {
    return i73debe587dffbb4f4db9d9bed9a0d4d4fecc4b1e819504a326cb68d7a674fd8c.NewSechRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Second()(*i09c5c179d03bde65cb63adabd7bd4fbba135050439becd052c9e02d98ccc445d.SecondRequestBuilder) {
    return i09c5c179d03bde65cb63adabd7bd4fbba135050439becd052c9e02d98ccc445d.NewSecondRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SeriesSum()(*ic3583f8f5afc77ad8b847c5324d2d64e7fe873bea81f387539aca6056a3857d1.SeriesSumRequestBuilder) {
    return ic3583f8f5afc77ad8b847c5324d2d64e7fe873bea81f387539aca6056a3857d1.NewSeriesSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sheet()(*i48c062699c214b911f9435316ac4c100152635b2cee12fc5cb6ea644b079214f.SheetRequestBuilder) {
    return i48c062699c214b911f9435316ac4c100152635b2cee12fc5cb6ea644b079214f.NewSheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sheets()(*iff5d1b1deb3b8f2c2d3de6f2d3cbbc8961c3a069683ac3268635178bc77155cc.SheetsRequestBuilder) {
    return iff5d1b1deb3b8f2c2d3de6f2d3cbbc8961c3a069683ac3268635178bc77155cc.NewSheetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sign()(*ibc4fcaf29a5c8243ec98cb8505e6ba9d34dfee284e3358f0f05b1f686a62307e.SignRequestBuilder) {
    return ibc4fcaf29a5c8243ec98cb8505e6ba9d34dfee284e3358f0f05b1f686a62307e.NewSignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sin()(*ifa3194c7763f634c1efa4e8ae657eaa8c118864c0f01ba521f2c5ae992f44bca.SinRequestBuilder) {
    return ifa3194c7763f634c1efa4e8ae657eaa8c118864c0f01ba521f2c5ae992f44bca.NewSinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sinh()(*i05bb4ed647eed6938ebc562b6011f32153484483264b9f46dd53577d8bb071f0.SinhRequestBuilder) {
    return i05bb4ed647eed6938ebc562b6011f32153484483264b9f46dd53577d8bb071f0.NewSinhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Skew()(*i4f07a98ba1b3a8f551935097937b9a4cd8b34d18c75ac58ac182fd2835ed51b2.SkewRequestBuilder) {
    return i4f07a98ba1b3a8f551935097937b9a4cd8b34d18c75ac58ac182fd2835ed51b2.NewSkewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Skew_p()(*ic787c76c3a4d7eafff0f967dfc5e037834a6e91c0a559608a912d2aa9d7de413.Skew_pRequestBuilder) {
    return ic787c76c3a4d7eafff0f967dfc5e037834a6e91c0a559608a912d2aa9d7de413.NewSkew_pRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sln()(*i3c257e281eaba9df97dc69b692bc1e7a702a67ef1f4d7673544a162df195a9c8.SlnRequestBuilder) {
    return i3c257e281eaba9df97dc69b692bc1e7a702a67ef1f4d7673544a162df195a9c8.NewSlnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Small()(*i146daa43f721553e08c4e14d4fe76761df9364bd7bd6b5239d7f796913654307.SmallRequestBuilder) {
    return i146daa43f721553e08c4e14d4fe76761df9364bd7bd6b5239d7f796913654307.NewSmallRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sqrt()(*ie5577128b8e24362fb1c5a485d4911522b5ced09cde805cc1a05969da2097603.SqrtRequestBuilder) {
    return ie5577128b8e24362fb1c5a485d4911522b5ced09cde805cc1a05969da2097603.NewSqrtRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SqrtPi()(*i5e90c8df3a37e03a7f9daa30813e1a9a6651dbf7a729a7284aafa183e84e909d.SqrtPiRequestBuilder) {
    return i5e90c8df3a37e03a7f9daa30813e1a9a6651dbf7a729a7284aafa183e84e909d.NewSqrtPiRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Standardize()(*id5ec5e3a45e7ee1f01952c4a4d9bc19dfe91a032fac85e7b0779e7813ded9f5b.StandardizeRequestBuilder) {
    return id5ec5e3a45e7ee1f01952c4a4d9bc19dfe91a032fac85e7b0779e7813ded9f5b.NewStandardizeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDev_P()(*if1770654587a61f420eb20ffc8a5ea1f8b4c1742716cd5e46b4bff76c3a5264a.StDev_PRequestBuilder) {
    return if1770654587a61f420eb20ffc8a5ea1f8b4c1742716cd5e46b4bff76c3a5264a.NewStDev_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDev_S()(*id108db16dde59511c17ec892beae748163d0a234c67bf6260e480f03b728ca1d.StDev_SRequestBuilder) {
    return id108db16dde59511c17ec892beae748163d0a234c67bf6260e480f03b728ca1d.NewStDev_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDevA()(*if8cd7bff2592738d8202fe4b64adb52b0349bd683eb7b2883788632cf859eb8f.StDevARequestBuilder) {
    return if8cd7bff2592738d8202fe4b64adb52b0349bd683eb7b2883788632cf859eb8f.NewStDevARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) StDevPA()(*i0494ccdd128ce2f5c8821e6a849105650c4b4c0bee4f2530ceb6fbdeb4e8c7c4.StDevPARequestBuilder) {
    return i0494ccdd128ce2f5c8821e6a849105650c4b4c0bee4f2530ceb6fbdeb4e8c7c4.NewStDevPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Substitute()(*iff2f96c56ce6b791c099573e926303c159821091509cb5a3cdffd2974d8cfb13.SubstituteRequestBuilder) {
    return iff2f96c56ce6b791c099573e926303c159821091509cb5a3cdffd2974d8cfb13.NewSubstituteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Subtotal()(*i6d8676c0302c91f80ee45dee8c2947ffc0870c9b7bfaa2dead2be59fa069aa05.SubtotalRequestBuilder) {
    return i6d8676c0302c91f80ee45dee8c2947ffc0870c9b7bfaa2dead2be59fa069aa05.NewSubtotalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Sum()(*ied40e56292d4efe1ae5f8b24055751ee82932520d99ae7cbbd34919617fc342b.SumRequestBuilder) {
    return ied40e56292d4efe1ae5f8b24055751ee82932520d99ae7cbbd34919617fc342b.NewSumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumIf()(*i9591d406740b7642ef993df1ac69b6e39bf409d7daee05f3abc58402bda7e28f.SumIfRequestBuilder) {
    return i9591d406740b7642ef993df1ac69b6e39bf409d7daee05f3abc58402bda7e28f.NewSumIfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumIfs()(*ie95ea4c78471cf9e8da86904a28af4b814e78189e787241116f733848361a7b8.SumIfsRequestBuilder) {
    return ie95ea4c78471cf9e8da86904a28af4b814e78189e787241116f733848361a7b8.NewSumIfsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) SumSq()(*id7a16762b42d506809d333c334ee18086ca130a944d2d04f8c3902f273dc8a55.SumSqRequestBuilder) {
    return id7a16762b42d506809d333c334ee18086ca130a944d2d04f8c3902f273dc8a55.NewSumSqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Syd()(*ief2b188c34ec07e72592e39a3d9a153a92765342857c1a9e782c05d12b30f62e.SydRequestBuilder) {
    return ief2b188c34ec07e72592e39a3d9a153a92765342857c1a9e782c05d12b30f62e.NewSydRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T()(*i41d373dba62b6f70cc8ef9047fdafa598a1a6fd341eb627992cd1d45bbfce880.TRequestBuilder) {
    return i41d373dba62b6f70cc8ef9047fdafa598a1a6fd341eb627992cd1d45bbfce880.NewTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist()(*id85fe98d73d528eb11688fc89fdadcb47512cb9d2f22de02e2af1f3125e360f0.T_DistRequestBuilder) {
    return id85fe98d73d528eb11688fc89fdadcb47512cb9d2f22de02e2af1f3125e360f0.NewT_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist_2T()(*i19b74ab0a3dd49a1976bab9450d4ef387991450a9dd1f439a0f11e82b1287321.T_Dist_2TRequestBuilder) {
    return i19b74ab0a3dd49a1976bab9450d4ef387991450a9dd1f439a0f11e82b1287321.NewT_Dist_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Dist_RT()(*iac41516568b082ded7ba5277f71216c2d4b48c3497906e5bed96d890ad7c4994.T_Dist_RTRequestBuilder) {
    return iac41516568b082ded7ba5277f71216c2d4b48c3497906e5bed96d890ad7c4994.NewT_Dist_RTRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Inv()(*iea0f3b887bb91d47447172b7d3ea7fcec40b7c9dccc043a99fefa454a07cbec7.T_InvRequestBuilder) {
    return iea0f3b887bb91d47447172b7d3ea7fcec40b7c9dccc043a99fefa454a07cbec7.NewT_InvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) T_Inv_2T()(*i8dbc7a210e86514658b6b7b58e46f4977dd33fe20198e1a8cf092493e3ef7122.T_Inv_2TRequestBuilder) {
    return i8dbc7a210e86514658b6b7b58e46f4977dd33fe20198e1a8cf092493e3ef7122.NewT_Inv_2TRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Tan()(*i1325a7cc878f353726715d2b70625d31aa9f68c17022b85fd116b3ec2e2815ca.TanRequestBuilder) {
    return i1325a7cc878f353726715d2b70625d31aa9f68c17022b85fd116b3ec2e2815ca.NewTanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Tanh()(*ic7d1ebf32cd5371fa01a5068d8262f1e6cd7ba173fab384d0eb5c5335a3998e5.TanhRequestBuilder) {
    return ic7d1ebf32cd5371fa01a5068d8262f1e6cd7ba173fab384d0eb5c5335a3998e5.NewTanhRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillEq()(*i0050d32d11128323eba61e1e75eb83c5547c444e23b016815365cda7028c0797.TbillEqRequestBuilder) {
    return i0050d32d11128323eba61e1e75eb83c5547c444e23b016815365cda7028c0797.NewTbillEqRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillPrice()(*id3e78ad2851361e3dea5140cea20ac6e2a699248ed96eae657254dbfd285c917.TbillPriceRequestBuilder) {
    return id3e78ad2851361e3dea5140cea20ac6e2a699248ed96eae657254dbfd285c917.NewTbillPriceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TbillYield()(*ie598c23cd821465563c1a517655127af97b775b456b274545e564c741d25d5cf.TbillYieldRequestBuilder) {
    return ie598c23cd821465563c1a517655127af97b775b456b274545e564c741d25d5cf.NewTbillYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Text()(*ifdfada0f682e13dd6856925da661b74617c1e052d16c650a5b1520225acf60dc.TextRequestBuilder) {
    return ifdfada0f682e13dd6856925da661b74617c1e052d16c650a5b1520225acf60dc.NewTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Time()(*i2b92e5915b7bdf7e95d7baea79b7488f604113d133d1d5a4afebacf7f0a4b470.TimeRequestBuilder) {
    return i2b92e5915b7bdf7e95d7baea79b7488f604113d133d1d5a4afebacf7f0a4b470.NewTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Timevalue()(*i6f940147c420e91e85cb05f8f948d0d2e450e4a2783950aa886229ff61a8189c.TimevalueRequestBuilder) {
    return i6f940147c420e91e85cb05f8f948d0d2e450e4a2783950aa886229ff61a8189c.NewTimevalueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Today()(*ic437e437d0e71b1d3cc7b04d1ab93a5f450917005d23d28cde22e2a557125385.TodayRequestBuilder) {
    return ic437e437d0e71b1d3cc7b04d1ab93a5f450917005d23d28cde22e2a557125385.NewTodayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Trim()(*i7c89957eea7f5d1a9ead132e515642f958a7d519b693bc9aa4a5168ff3d6a7ee.TrimRequestBuilder) {
    return i7c89957eea7f5d1a9ead132e515642f958a7d519b693bc9aa4a5168ff3d6a7ee.NewTrimRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) TrimMean()(*iabaec7f58e32d0ba15b2f608134b387ffc3dbf171ae18d2d78f0fdb2603db827.TrimMeanRequestBuilder) {
    return iabaec7f58e32d0ba15b2f608134b387ffc3dbf171ae18d2d78f0fdb2603db827.NewTrimMeanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) True()(*i7063fea931c4bd333bf58cae4ae95006d880ae51dd69347fcd0112ce942d6e56.TrueRequestBuilder) {
    return i7063fea931c4bd333bf58cae4ae95006d880ae51dd69347fcd0112ce942d6e56.NewTrueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Trunc()(*ib5af4e9a26a97744aea932645181a72915add2b5f542a4712f8a88e0fb711a6a.TruncRequestBuilder) {
    return ib5af4e9a26a97744aea932645181a72915add2b5f542a4712f8a88e0fb711a6a.NewTruncRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Type_escaped()(*i822d92a8a540ae6b23416a1ba6edd6eff35f4abed7953e4928fbd2b4ee76d341.TypeRequestBuilder) {
    return i822d92a8a540ae6b23416a1ba6edd6eff35f4abed7953e4928fbd2b4ee76d341.NewTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Unichar()(*i3f56cdb4a69ef320a75d1586206b7dfb437b5a604fce57a5019673319644e355.UnicharRequestBuilder) {
    return i3f56cdb4a69ef320a75d1586206b7dfb437b5a604fce57a5019673319644e355.NewUnicharRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Unicode()(*ic4eac566db71e14385e8ec72cd1979f90b074926778f8d00063effaaee1753b5.UnicodeRequestBuilder) {
    return ic4eac566db71e14385e8ec72cd1979f90b074926778f8d00063effaaee1753b5.NewUnicodeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Upper()(*i58fe0cf7e33ef2cc42c1aa9ddecea577e48b03f4cdf37d375591026954f46c7c.UpperRequestBuilder) {
    return i58fe0cf7e33ef2cc42c1aa9ddecea577e48b03f4cdf37d375591026954f46c7c.NewUpperRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Usdollar()(*ic1c776e8f5bbf1aad14e8f62953dd29ae80f14bf4e5b2c7caa6d1ada9a7ae721.UsdollarRequestBuilder) {
    return ic1c776e8f5bbf1aad14e8f62953dd29ae80f14bf4e5b2c7caa6d1ada9a7ae721.NewUsdollarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Value()(*i37cb723b274316a2f034aa8b2a909f4d75dfe2fdcbcd35a4fdb5a661e02d602a.ValueRequestBuilder) {
    return i37cb723b274316a2f034aa8b2a909f4d75dfe2fdcbcd35a4fdb5a661e02d602a.NewValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Var_P()(*i80cae408a9fb23c0b3d5973b3b36036f8f86da13604a2df9e4f5d82335af9b89.Var_PRequestBuilder) {
    return i80cae408a9fb23c0b3d5973b3b36036f8f86da13604a2df9e4f5d82335af9b89.NewVar_PRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Var_S()(*if394010a3655e451e9502ee93545f5fe9f2d41f16ca348052e9a061dfb3bae6c.Var_SRequestBuilder) {
    return if394010a3655e451e9502ee93545f5fe9f2d41f16ca348052e9a061dfb3bae6c.NewVar_SRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) VarA()(*i04bc21ac86d9392376c6b121e63d25df0567e9a3fc8da2e91db63efdd803c37b.VarARequestBuilder) {
    return i04bc21ac86d9392376c6b121e63d25df0567e9a3fc8da2e91db63efdd803c37b.NewVarARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) VarPA()(*ie7ed1bf0ad464b9c429a113f7c45516d7c13166cd0c581a3e34aa78ba6423b7b.VarPARequestBuilder) {
    return ie7ed1bf0ad464b9c429a113f7c45516d7c13166cd0c581a3e34aa78ba6423b7b.NewVarPARequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Vdb()(*i114f33a34a95c8aeeb6f01a515b9f584c007b96d3e8469f93ccaf4f74d1473ac.VdbRequestBuilder) {
    return i114f33a34a95c8aeeb6f01a515b9f584c007b96d3e8469f93ccaf4f74d1473ac.NewVdbRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Vlookup()(*id7014948bb05b38a010de7d8ee6c8910b1084bb4974c710ce61cbf7e75c3fe65.VlookupRequestBuilder) {
    return id7014948bb05b38a010de7d8ee6c8910b1084bb4974c710ce61cbf7e75c3fe65.NewVlookupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Weekday()(*ibb5a055522b92a57289f795953bc91a9047dd934fe9577923e13cd420be37ca9.WeekdayRequestBuilder) {
    return ibb5a055522b92a57289f795953bc91a9047dd934fe9577923e13cd420be37ca9.NewWeekdayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WeekNum()(*i5c04cac2d711234a9025c8e14e1c9fddeb0af9d850a46fc99376ee1f28840835.WeekNumRequestBuilder) {
    return i5c04cac2d711234a9025c8e14e1c9fddeb0af9d850a46fc99376ee1f28840835.NewWeekNumRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Weibull_Dist()(*i4ac7e344d13c36c2e3c23d87816150a649beaf2012302f6610653444ecb9b234.Weibull_DistRequestBuilder) {
    return i4ac7e344d13c36c2e3c23d87816150a649beaf2012302f6610653444ecb9b234.NewWeibull_DistRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WorkDay()(*ia59debb70c35806b26ac204748774dcb3d74fce9e80e89b0347efe6d56c2bbc7.WorkDayRequestBuilder) {
    return ia59debb70c35806b26ac204748774dcb3d74fce9e80e89b0347efe6d56c2bbc7.NewWorkDayRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) WorkDay_Intl()(*i486d88d3ffdf365f62b4928cda0b21c0c76426936a8b8ac8ccf19250c5debac9.WorkDay_IntlRequestBuilder) {
    return i486d88d3ffdf365f62b4928cda0b21c0c76426936a8b8ac8ccf19250c5debac9.NewWorkDay_IntlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xirr()(*ia07dbb8476d456306a1137e9feae6f70a68fc248ed129f143d15460b2129f4e2.XirrRequestBuilder) {
    return ia07dbb8476d456306a1137e9feae6f70a68fc248ed129f143d15460b2129f4e2.NewXirrRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xnpv()(*i2b962df7680ae6985d6808308b7cb3aeecdbe628568137708567ef5ed78c52be.XnpvRequestBuilder) {
    return i2b962df7680ae6985d6808308b7cb3aeecdbe628568137708567ef5ed78c52be.NewXnpvRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Xor()(*iacd13e9f55fa4756e8bb88cac5e77f31e49dca8e18f11d86e62e585de0a4b76a.XorRequestBuilder) {
    return iacd13e9f55fa4756e8bb88cac5e77f31e49dca8e18f11d86e62e585de0a4b76a.NewXorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Year()(*i084289164bd03459fb7e41d139af885ebcd76e0cd01fb224a5df9263ede20802.YearRequestBuilder) {
    return i084289164bd03459fb7e41d139af885ebcd76e0cd01fb224a5df9263ede20802.NewYearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YearFrac()(*icd9c390016ac3cd52b33ff2f290463ccc95fb31b19d45dc5aae9850d0dd11c4f.YearFracRequestBuilder) {
    return icd9c390016ac3cd52b33ff2f290463ccc95fb31b19d45dc5aae9850d0dd11c4f.NewYearFracRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Yield()(*i69c2649aba19ccdd2dd0f5bf255ddf8d0aa623a6c022229dbd77634b2c88bed8.YieldRequestBuilder) {
    return i69c2649aba19ccdd2dd0f5bf255ddf8d0aa623a6c022229dbd77634b2c88bed8.NewYieldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YieldDisc()(*i2544fd0935f5ac2f12f79b4a02a43b030e17165c43e65638f812efe9c7d29892.YieldDiscRequestBuilder) {
    return i2544fd0935f5ac2f12f79b4a02a43b030e17165c43e65638f812efe9c7d29892.NewYieldDiscRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) YieldMat()(*i3728955a08a79c9a0b07d504a99948f05e706d80b3740b69a04dbf94261e8dc0.YieldMatRequestBuilder) {
    return i3728955a08a79c9a0b07d504a99948f05e706d80b3740b69a04dbf94261e8dc0.NewYieldMatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *FunctionsRequestBuilder) Z_Test()(*i737c5d48df084950792e74f03a89bbffee5cbd3676f6378d4a47838452db791c.Z_TestRequestBuilder) {
    return i737c5d48df084950792e74f03a89bbffee5cbd3676f6378d4a47838452db791c.NewZ_TestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
