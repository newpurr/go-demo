package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 并发测试
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}

// 基础单元测试
func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"测试Hello World",
			"Hello, world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree(t *testing.T) {
	jsonStr := `[
  {
    "id": 1,
    "channel_name": "[1]公共库",
    "parent_id": 460
  },
  {
    "id": 2,
    "channel_name": "[2]弹药库",
    "parent_id": 460
  },
  {
    "id": 3,
    "channel_name": "[3]E宠商城",
    "parent_id": 461
  },
  {
    "id": 4,
    "channel_name": "[4]E宠三方店",
    "parent_id": 1
  },
  {
    "id": 5,
    "channel_name": "[5]孵化店群(即将废弃)",
    "parent_id": 1
  },
  {
    "id": 7,
    "channel_name": "[7]微店",
    "parent_id": 475
  },
  {
    "id": 8,
    "channel_name": "[8]酒鬼一家",
    "parent_id": 4
  },
  {
    "id": 9,
    "channel_name": "[9]柿子菌",
    "parent_id": 5
  },
  {
    "id": 10,
    "channel_name": "[10]十爷家",
    "parent_id": 4
  },
  {
    "id": 11,
    "channel_name": "[11]京东E宠旗舰店",
    "parent_id": 4
  },
  {
    "id": 12,
    "channel_name": "[12]微店E宠专营店",
    "parent_id": 7
  },
  {
    "id": 13,
    "channel_name": "[13]品牌特卖",
    "parent_id": 3
  },
  {
    "id": 14,
    "channel_name": "[14]E宠团",
    "parent_id": 3
  },
  {
    "id": 15,
    "channel_name": "[15]每日疯抢",
    "parent_id": 3
  },
  {
    "id": 16,
    "channel_name": "[16]0元转盘",
    "parent_id": 3
  },
  {
    "id": 17,
    "channel_name": "[17]礼盒模式",
    "parent_id": 3
  },
  {
    "id": 18,
    "channel_name": "[18]折扣专题",
    "parent_id": 3
  },
  {
    "id": 19,
    "channel_name": "[19]签到",
    "parent_id": 3
  },
  {
    "id": 20,
    "channel_name": "[20]天猫E宠旗舰店",
    "parent_id": 4
  },
  {
    "id": 21,
    "channel_name": "[21]任务量补货",
    "parent_id": 0
  },
  {
    "id": 22,
    "channel_name": "[22]任务量补货渠道",
    "parent_id": 21
  },
  {
    "id": 24,
    "channel_name": "[24]体验馆",
    "parent_id": 3
  },
  {
    "id": 25,
    "channel_name": "[25]其他渠道",
    "parent_id": 475
  },
  {
    "id": 26,
    "channel_name": "[26]其他栏目",
    "parent_id": 25
  },
  {
    "id": 27,
    "channel_name": "[27]上新日",
    "parent_id": 3
  },
  {
    "id": 28,
    "channel_name": "[28]品牌日",
    "parent_id": 3
  },
  {
    "id": 29,
    "channel_name": "[29]会员日",
    "parent_id": 3
  },
  {
    "id": 30,
    "channel_name": "[30]吾哩宠物",
    "parent_id": 4
  },
  {
    "id": 33,
    "channel_name": "[33]它品云店",
    "parent_id": 3
  },
  {
    "id": 34,
    "channel_name": "[34]萌萌圈定制团",
    "parent_id": 3
  },
  {
    "id": 35,
    "channel_name": "[35]预售",
    "parent_id": 3
  },
  {
    "id": 36,
    "channel_name": "[36]CRM促活栏目",
    "parent_id": 3
  },
  {
    "id": 37,
    "channel_name": "[37]CRM0元体验",
    "parent_id": 3
  },
  {
    "id": 38,
    "channel_name": "[38]CRM赠品",
    "parent_id": 3
  },
  {
    "id": 39,
    "channel_name": "[39]缺货预订",
    "parent_id": 3
  },
  {
    "id": 40,
    "channel_name": "[40]市场栏目",
    "parent_id": 3
  },
  {
    "id": 41,
    "channel_name": "[41]双11限时抢购",
    "parent_id": 3
  },
  {
    "id": 43,
    "channel_name": "[43]小红书E宠旗舰店",
    "parent_id": 4
  },
  {
    "id": 44,
    "channel_name": "[44]新客CRM",
    "parent_id": 3
  },
  {
    "id": 45,
    "channel_name": "[45]限时抢购",
    "parent_id": 3
  },
  {
    "id": 46,
    "channel_name": "[46]历史预售-1",
    "parent_id": 3
  },
  {
    "id": 47,
    "channel_name": "[47]历史预售-2",
    "parent_id": 3
  },
  {
    "id": 48,
    "channel_name": "[48]历史预售-3",
    "parent_id": 3
  },
  {
    "id": 50,
    "channel_name": "[50]考拉E宠专营店",
    "parent_id": 4
  },
  {
    "id": 51,
    "channel_name": "[51]清仓",
    "parent_id": 460
  },
  {
    "id": 124,
    "channel_name": "[124]小爪管家批发",
    "parent_id": 461
  },
  {
    "id": 308,
    "channel_name": "[308]调度渠道",
    "parent_id": 460
  },
  {
    "id": 417,
    "channel_name": "[417]自然销售",
    "parent_id": 3
  },
  {
    "id": 433,
    "channel_name": "[433]分销渠道",
    "parent_id": 3
  },
  {
    "id": 434,
    "channel_name": "[434]宠盟国际海外专营店",
    "parent_id": 1809
  },
  {
    "id": 435,
    "channel_name": "[435]天猫petcurean",
    "parent_id": 1809
  },
  {
    "id": 436,
    "channel_name": "[436]哚汪喵",
    "parent_id": 461
  },
  {
    "id": 438,
    "channel_name": "[438]小春有好货",
    "parent_id": 461
  },
  {
    "id": 439,
    "channel_name": "[439]乐仓分销",
    "parent_id": 433
  },
  {
    "id": 440,
    "channel_name": "[440]老和山兄弟批发",
    "parent_id": 433
  },
  {
    "id": 441,
    "channel_name": "[441]海际批发",
    "parent_id": 433
  },
  {
    "id": 442,
    "channel_name": "[442]锦湖吉阳批发",
    "parent_id": 433
  },
  {
    "id": 443,
    "channel_name": "[443]5U批发",
    "parent_id": 433
  },
  {
    "id": 444,
    "channel_name": "[444]汪星达人分销",
    "parent_id": 433
  },
  {
    "id": 445,
    "channel_name": "[445]TG.SILVERLIGHT LTD",
    "parent_id": 433
  },
  {
    "id": 446,
    "channel_name": "[446]GLOBALPETFOODS批发",
    "parent_id": 433
  },
  {
    "id": 447,
    "channel_name": "[447]acana批发",
    "parent_id": 433
  },
  {
    "id": 448,
    "channel_name": "[448]西喵馆",
    "parent_id": 433
  },
  {
    "id": 449,
    "channel_name": "[449]海瑞特批发",
    "parent_id": 433
  },
  {
    "id": 450,
    "channel_name": "[450]杨志伟批发",
    "parent_id": 433
  },
  {
    "id": 451,
    "channel_name": "[451]香港宠物之家批发",
    "parent_id": 433
  },
  {
    "id": 452,
    "channel_name": "[452]欧陆批发",
    "parent_id": 433
  },
  {
    "id": 453,
    "channel_name": "[453]唯品会批发",
    "parent_id": 433
  },
  {
    "id": 454,
    "channel_name": "[454]天猫直营大茂批发",
    "parent_id": 433
  },
  {
    "id": 455,
    "channel_name": "[455]零零鼠分销",
    "parent_id": 433
  },
  {
    "id": 456,
    "channel_name": "[456]天猫跨境直营批发",
    "parent_id": 433
  },
  {
    "id": 457,
    "channel_name": "[457]广州B2B",
    "parent_id": 433
  },
  {
    "id": 458,
    "channel_name": "[458]法蔓分销",
    "parent_id": 433
  },
  {
    "id": 459,
    "channel_name": "[459]宠萌代采",
    "parent_id": 433
  },
  {
    "id": 460,
    "channel_name": "[460]易宠集团",
    "parent_id": 0
  },
  {
    "id": 461,
    "channel_name": "[461]易宠平台",
    "parent_id": 1
  },
  {
    "id": 462,
    "channel_name": "[462]满省",
    "parent_id": 3
  },
  {
    "id": 463,
    "channel_name": "[463]三免一",
    "parent_id": 3
  },
  {
    "id": 464,
    "channel_name": "[464]多件优惠",
    "parent_id": 3
  },
  {
    "id": 465,
    "channel_name": "[465]买赠",
    "parent_id": 3
  },
  {
    "id": 466,
    "channel_name": "[466]礼包",
    "parent_id": 3
  },
  {
    "id": 467,
    "channel_name": "[467]换购",
    "parent_id": 3
  },
  {
    "id": 468,
    "channel_name": "[468]组合",
    "parent_id": 3
  },
  {
    "id": 469,
    "channel_name": "[469]搭配",
    "parent_id": 3
  },
  {
    "id": 470,
    "channel_name": "[470]武汉网红店",
    "parent_id": 4
  },
  {
    "id": 471,
    "channel_name": "[471]贝适安旗舰店",
    "parent_id": 475
  },
  {
    "id": 472,
    "channel_name": "[472]预占库存扣减渠道",
    "parent_id": 460
  },
  {
    "id": 474,
    "channel_name": "[474]猫生狗愿",
    "parent_id": 6721
  },
  {
    "id": 477,
    "channel_name": "[477]生日礼",
    "parent_id": 3
  },
  {
    "id": 478,
    "channel_name": "[478]大促活动",
    "parent_id": 3
  },
  {
    "id": 1066,
    "channel_name": "[1066]宠好货平台",
    "parent_id": 475
  },
  {
    "id": 1067,
    "channel_name": "[1067]宠好货",
    "parent_id": 1066
  },
  {
    "id": 1809,
    "channel_name": "[1809]宠盟国际店",
    "parent_id": 1
  },
  {
    "id": 1917,
    "channel_name": "[1917]小狗在家",
    "parent_id": 1
  },
  {
    "id": 2152,
    "channel_name": "[2152]潮品预售",
    "parent_id": 3
  },
  {
    "id": 2200,
    "channel_name": "[2200]新客复购",
    "parent_id": 3
  },
  {
    "id": 2232,
    "channel_name": "[2232]远集宠物专营店",
    "parent_id": 4
  },
  {
    "id": 2862,
    "channel_name": "[2862]爱库存",
    "parent_id": 433
  },
  {
    "id": 2917,
    "channel_name": "[2917]抖音E宠商城旗舰店",
    "parent_id": 461
  },
  {
    "id": 3069,
    "channel_name": "[3069]有赞E宠旗舰店",
    "parent_id": 461
  },
  {
    "id": 3070,
    "channel_name": "[3070]猫裂变(吸猫菌)",
    "parent_id": 6721
  },
  {
    "id": 3071,
    "channel_name": "[3071]拼多多E宠旗舰店",
    "parent_id": 4
  },
  {
    "id": 3224,
    "channel_name": "[3224]官方小程序活动-折扣",
    "parent_id": 3
  },
  {
    "id": 3969,
    "channel_name": "[3969]远集批发",
    "parent_id": 5
  },
  {
    "id": 4747,
    "channel_name": "[4747]莲花海外旗舰店",
    "parent_id": 1809
  },
  {
    "id": 6243,
    "channel_name": "[6243]宠盟备货库",
    "parent_id": 460
  },
  {
    "id": 6503,
    "channel_name": "[6503]官方小程序活动-礼包",
    "parent_id": 3
  },
  {
    "id": 6716,
    "channel_name": "[6716]会员兑换",
    "parent_id": 3070
  },
  {
    "id": 6717,
    "channel_name": "[6717]E宠订阅",
    "parent_id": 3
  },
  {
    "id": 6718,
    "channel_name": "[6718]分销2部",
    "parent_id": 433
  },
  {
    "id": 6719,
    "channel_name": "[6719]海际（重庆）",
    "parent_id": 1
  },
  {
    "id": 6720,
    "channel_name": "[6720]它品商城",
    "parent_id": 6719
  },
  {
    "id": 6721,
    "channel_name": "[6721]吸猫菌",
    "parent_id": 1
  },
  {
    "id": 6722,
    "channel_name": "[6722]周三购",
    "parent_id": 3
  },
  {
    "id": 6723,
    "channel_name": "[6723]代发分销",
    "parent_id": 1
  },
  {
    "id": 6724,
    "channel_name": "[6724]卜树批发",
    "parent_id": 6806
  },
  {
    "id": 6725,
    "channel_name": "[6725]洋码头E宠商城旗舰店",
    "parent_id": 461
  },
  {
    "id": 6726,
    "channel_name": "[6726]缺货预售",
    "parent_id": 3
  },
  {
    "id": 6727,
    "channel_name": "[6727]宠物研究院",
    "parent_id": 3
  },
  {
    "id": 6728,
    "channel_name": "[6728]海际（广州）",
    "parent_id": 1
  },
  {
    "id": 6729,
    "channel_name": "[6729]活动缺货预售",
    "parent_id": 3
  },
  {
    "id": 6730,
    "channel_name": "[6730]海际代发业务",
    "parent_id": 6723
  },
  {
    "id": 6731,
    "channel_name": "[6731]亚宠云店",
    "parent_id": 1
  },
  {
    "id": 6732,
    "channel_name": "[6732]派小店",
    "parent_id": 6731
  },
  {
    "id": 6733,
    "channel_name": "[6733]它品备货库",
    "parent_id": 6719
  },
  {
    "id": 6734,
    "channel_name": "[6734]非独享活动",
    "parent_id": 3
  },
  {
    "id": 6735,
    "channel_name": "[6735]tikipets旗舰店",
    "parent_id": 4
  },
  {
    "id": 6736,
    "channel_name": "[6736]新它品云店",
    "parent_id": 6719
  },
  {
    "id": 6737,
    "channel_name": "[6737]积分商城",
    "parent_id": 6719
  },
  {
    "id": 6738,
    "channel_name": "[6738]E宠专用渠道【已停用】",
    "parent_id": 6719
  },
  {
    "id": 6739,
    "channel_name": "[6739]芬宠Plus",
    "parent_id": 6731
  },
  {
    "id": 6742,
    "channel_name": "[6742]重庆海际1688代发",
    "parent_id": 6719
  },
  {
    "id": 6743,
    "channel_name": "[6743]奇葩熊",
    "parent_id": 1
  },
  {
    "id": 6744,
    "channel_name": "[6744]杭州魔方",
    "parent_id": 1
  },
  {
    "id": 6745,
    "channel_name": "[6745]生命宝藏旗舰店",
    "parent_id": 6744
  },
  {
    "id": 6746,
    "channel_name": "[6746]海际备货",
    "parent_id": 6744
  },
  {
    "id": 6747,
    "channel_name": "[6747]原始猎食渴望旗舰店",
    "parent_id": 6743
  },
  {
    "id": 6748,
    "channel_name": "[6748]爱肯拿旗舰店",
    "parent_id": 6743
  },
  {
    "id": 6749,
    "channel_name": "[6749]全是爆品HOTSALE",
    "parent_id": 6723
  },
  {
    "id": 6750,
    "channel_name": "[6750]枫趣旗舰店",
    "parent_id": 4
  },
  {
    "id": 6751,
    "channel_name": "[6751]批发渠道",
    "parent_id": 6744
  },
  {
    "id": 6752,
    "channel_name": "[6752]骨头宠物",
    "parent_id": 461
  },
  {
    "id": 6753,
    "channel_name": "[6753]骨头宠物自然销售",
    "parent_id": 417
  },
  {
    "id": 6754,
    "channel_name": "[6754]骨头宠物活动销售",
    "parent_id": 3
  },
  {
    "id": 6756,
    "channel_name": "[6756]骨头宠物-新客频道",
    "parent_id": 3
  },
  {
    "id": 6758,
    "channel_name": "[6758]骨头宠物-免费礼包",
    "parent_id": 3
  },
  {
    "id": 6759,
    "channel_name": "[6759]降龙库存",
    "parent_id": 6719
  },
  {
    "id": 6760,
    "channel_name": "[6760]东部大区",
    "parent_id": 6791
  },
  {
    "id": 6761,
    "channel_name": "[6761]西部大区",
    "parent_id": 6790
  },
  {
    "id": 6762,
    "channel_name": "[6762]南部大区",
    "parent_id": 6791
  },
  {
    "id": 6763,
    "channel_name": "[6763]北部大区",
    "parent_id": 6790
  },
  {
    "id": 6764,
    "channel_name": "[6764]江左盟",
    "parent_id": 6760
  },
  {
    "id": 6765,
    "channel_name": "[6765]桃花岛",
    "parent_id": 6760
  },
  {
    "id": 6766,
    "channel_name": "[6766]六扇门",
    "parent_id": 6762
  },
  {
    "id": 6767,
    "channel_name": "[6767]华山派",
    "parent_id": 6763
  },
  {
    "id": 6768,
    "channel_name": "[6768]乐天派",
    "parent_id": 6763
  },
  {
    "id": 6769,
    "channel_name": "[6769]通吃岛",
    "parent_id": 6763
  },
  {
    "id": 6770,
    "channel_name": "[6770]逍遥派",
    "parent_id": 6763
  },
  {
    "id": 6771,
    "channel_name": "[6771]独孤派",
    "parent_id": 6760
  },
  {
    "id": 6772,
    "channel_name": "[6772]少林派",
    "parent_id": 6760
  },
  {
    "id": 6773,
    "channel_name": "[6773]天地会",
    "parent_id": 6760
  },
  {
    "id": 6774,
    "channel_name": "[6774]衡山派",
    "parent_id": 6762
  },
  {
    "id": 6775,
    "channel_name": "[6775]聚贤庄",
    "parent_id": 6762
  },
  {
    "id": 6776,
    "channel_name": "[6776]情义阁",
    "parent_id": 6762
  },
  {
    "id": 6777,
    "channel_name": "[6777]铁血盟",
    "parent_id": 6762
  },
  {
    "id": 6778,
    "channel_name": "[6778]云霄宫",
    "parent_id": 6762
  },
  {
    "id": 6779,
    "channel_name": "[6779]朝天门",
    "parent_id": 6761
  },
  {
    "id": 6780,
    "channel_name": "[6780]蜀山派",
    "parent_id": 6761
  },
  {
    "id": 6781,
    "channel_name": "[6781]武当派",
    "parent_id": 6761
  },
  {
    "id": 6782,
    "channel_name": "[6782]星宿派",
    "parent_id": 6761
  },
  {
    "id": 6783,
    "channel_name": "[6783]兄弟会",
    "parent_id": 6761
  },
  {
    "id": 6784,
    "channel_name": "[6784]线上销售组",
    "parent_id": 6791
  },
  {
    "id": 6785,
    "channel_name": "[6785]东部大区备货",
    "parent_id": 6760
  },
  {
    "id": 6786,
    "channel_name": "[6786]西部大区备货",
    "parent_id": 6761
  },
  {
    "id": 6787,
    "channel_name": "[6787]南部大区备货",
    "parent_id": 6762
  },
  {
    "id": 6788,
    "channel_name": "[6788]北部大区备货",
    "parent_id": 6763
  },
  {
    "id": 6789,
    "channel_name": "[6789]降龙备货库",
    "parent_id": 6759
  },
  {
    "id": 6790,
    "channel_name": "[6790]销售二部",
    "parent_id": 6759
  },
  {
    "id": 6791,
    "channel_name": "[6791]销售一部",
    "parent_id": 6759
  },
  {
    "id": 6792,
    "channel_name": "[6792]十爷家分销",
    "parent_id": 6723
  },
  {
    "id": 6793,
    "channel_name": "[6793]货权转移专用",
    "parent_id": 1
  },
  {
    "id": 6794,
    "channel_name": "[6794]它食袋",
    "parent_id": 6723
  },
  {
    "id": 6795,
    "channel_name": "[6795]E宠商城抖音旗舰店",
    "parent_id": 6723
  },
  {
    "id": 6796,
    "channel_name": "[6796]萌宠拼拼",
    "parent_id": 6723
  },
  {
    "id": 6797,
    "channel_name": "[6797]海际易宠备货渠道",
    "parent_id": 6720
  },
  {
    "id": 6798,
    "channel_name": "[6798]吸猫菌组合",
    "parent_id": 6721
  },
  {
    "id": 6799,
    "channel_name": "[6799]广州产品测试【大区】",
    "parent_id": 6759
  },
  {
    "id": 6800,
    "channel_name": "[6800]广州测试组【帮派】",
    "parent_id": 6799
  },
  {
    "id": 6801,
    "channel_name": "[6801]广州产品组【帮派】",
    "parent_id": 6799
  },
  {
    "id": 6802,
    "channel_name": "[6802]重庆宠盟",
    "parent_id": 1
  },
  {
    "id": 6803,
    "channel_name": "[6803]喵叔阿达严选",
    "parent_id": 6723
  },
  {
    "id": 6804,
    "channel_name": "[6804]猫奴十三",
    "parent_id": 6723
  },
  {
    "id": 6805,
    "channel_name": "[6805]重庆项目【帮派】",
    "parent_id": 6799
  },
  {
    "id": 6806,
    "channel_name": "[6806]易宠批发",
    "parent_id": 1
  },
  {
    "id": 6807,
    "channel_name": "[6807]于璐批发",
    "parent_id": 6806
  },
  {
    "id": 6808,
    "channel_name": "[6808]芬宠PLUS备用",
    "parent_id": 6731
  },
  {
    "id": 6809,
    "channel_name": "[6809]香港海际",
    "parent_id": 6793
  },
  {
    "id": 6810,
    "channel_name": "[6810]芬宠PLUS奖品",
    "parent_id": 6731
  },
  {
    "id": 6811,
    "channel_name": "[6811] TOUCHDOG宠物社",
    "parent_id": 6723
  },
  {
    "id": 6812,
    "channel_name": "[6812]呆萌虎宠物店8056",
    "parent_id": 6731
  },
  {
    "id": 6813,
    "channel_name": "[6813]毛毛的小店7975",
    "parent_id": 6731
  },
  {
    "id": 6814,
    "channel_name": "[6814]迷7834",
    "parent_id": 6731
  },
  {
    "id": 6815,
    "channel_name": "[6815]萌宠馆7762",
    "parent_id": 6731
  },
  {
    "id": 6816,
    "channel_name": "[6816]MeiMei生活馆cat7741",
    "parent_id": 6731
  },
  {
    "id": 6817,
    "channel_name": "[6817]猫语者7481",
    "parent_id": 6731
  },
  {
    "id": 6818,
    "channel_name": "[6818]Freak7384",
    "parent_id": 6731
  },
  {
    "id": 6819,
    "channel_name": "[6819]一只五花猪8058",
    "parent_id": 6731
  },
  {
    "id": 6820,
    "channel_name": "[6820]渣渣辉的店铺7743",
    "parent_id": 6731
  },
  {
    "id": 6821,
    "channel_name": "[6821]一缸盐汽水7801",
    "parent_id": 6731
  },
  {
    "id": 6822,
    "channel_name": "[6822]糕米滴宠物小铺8021",
    "parent_id": 6731
  },
  {
    "id": 6823,
    "channel_name": "[6823]金铂铂7742",
    "parent_id": 6731
  },
  {
    "id": 6824,
    "channel_name": "[6824]袁雨潇7515",
    "parent_id": 6731
  },
  {
    "id": 6825,
    "channel_name": "[6825]大圣宠物No1店7768",
    "parent_id": 6731
  },
  {
    "id": 6826,
    "channel_name": "[6826]大圣宠物猫咪用品7715",
    "parent_id": 6731
  },
  {
    "id": 6827,
    "channel_name": "[6827]猫科男猫居生活馆",
    "parent_id": 6723
  },
  {
    "id": 0,
    "channel_name": "[0]根节点",
    "parent_id": -1
  }
]`

	var slc []Item
	err := json.Unmarshal([]byte(jsonStr), &slc)
	if err != nil {
		panic(err)
	}

	tree := slcToTree(slc, -1)
	bytes, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type Item struct {
	Id          int    `json:"id"`
	ChannelName string `json:"channel_name"`
	ParentId    int    `json:"parent_id"`
	Children    []Item `json:"children"`
}

func slcToTree(slcAll []Item, pid int) []Item {
	if pid < -1 {
		return nil
	}
	var slc []Item
	for _, item := range slcAll {
		if item.ParentId == pid {
			item.Children = slcToTree(slcAll, item.Id)
			slc = append(slc, item)
		}
	}

	return slc
}
