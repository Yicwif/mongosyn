package main

import (
	"github.com/asiainfoLDP/datahub_repository/log"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	C_REPOSITORY            = "repository"
	C_REPOSITORY_DEL        = "repository_del"
	C_DATAITEM              = "dataitem"
	C_DATAITEM_DEL          = "dataitem_del"
	C_REPOSITORY_PERMISSION = "permission_rep"
	C_DATAITEM_PERMISSION   = "permission_item"
	C_SELECT                = "select"
	C_TAG                   = "tag"
	C_STATIS_DAY            = "statis_day"

	COL_REPNAME         = "repository_name"
	COL_REP_ACC         = "repaccesstype"
	COL_REP_ITEMS       = "items"
	COL_ITEM_NAME       = "dataitem_name"
	COL_CH_ITEM_NAME    = "ch_itemname"
	COL_ITEM_ACC        = "itemaccesstype"
	COL_ITEM_TAGS       = "tags"
	COL_COMMENT         = "comment"
	COL_RANK            = "rank"
	COL_PRICE           = "price"
	COL_CREATE_USER     = "create_user"
	COL_REP_COOPERATOR  = "cooperators"
	COL_ITEM_COOPERATOR = "ifcooperate"
	COL_LABEL           = "label"
	COL_OPTIME          = "optime"
	COL_ITEM_META       = "meta"
	COL_ITEM_SAMPLE     = "sample"
	COL_TAG_NAME        = "tag"
	COL_PERMIT_USER     = "user_name"
	COL_PERMIT_REPNAME  = "repository_name"
	COL_PERMIT_ITEMNAME = "dataitem_name"
	COL_PERMIT_WRITE    = "write"
	CMD_INC             = "$inc"
	CMD_ADDTOSET        = "$addToSet"
	CMD_SET             = "$set"
	CMD_UNSET           = "$unset"
	CMD_IN              = "$in"
	CMD_OR              = "$or"
	CMD_REGEX           = "$regex"
	CMD_OPTION          = "$options"
	CMD_AND             = "$and"
	CMD_CASE_ALL        = "$i"
	CMD_PULL            = "$pull"
	CMD_NOTEQUAL        = "$ne"
)

var (
	DB_NAMESPACE_MONGO = "datahub"
	DB_NAME            = "datahub"
	//SERVICE_PORT       = Env("goservice_port", false)

	//Service_Name_Mongo = Env("mongo_service_name", true)

	//DISCOVERY_CONSUL_SERVER_ADDR = Env("CONSUL_SERVER", true)
	//DISCOVERY_CONSUL_SERVER_PORT = Env("CONSUL_DNS_PORT", true)

	db  DB = DB{*connect()}
	Log    = log.NewLogger("http handler")

	//WH_User = []string{"948689655@qq.com", "shendf@primeton.com", "gchf_1@sina.com", "guocf5@cjbigdata.com", "service@cjbigdata.com", "qiuzhihua@gbase.cn", "service-cj@cjbigdata.com", "chengkj@yeah.net", "18810155806@163.com", "lixw@maitewang.com"}

	//GZ_User = []string{"service@gzbdex.com", "wangdc@gzbdex.com", "wang_dachao@163.com", "13902292121@139.com", "wdc-shuiping@163.com", "zhenchao.zhao@kcomber.com", "sjtan@sheendu.com", "info@yiyantang.cn", "hujq@asiainfo.com", "feng.he@heqinuc.com", "liuhf@ibbd.net", "yinhuacheng@linkip.cn", "xiexf@189.cn", "skyrock@139.com", "mailiantao@qq.com", "13822147976@139.com", "shihw11@sina.com", "walrkson@hotmail.com", "838959@qq.com", "56463383@qq.com", "13311288822@189.cn", "7904865@qq.com", "112793863@qq.com", "47181770@qq.com", "yizf@szlanyou.com", "cyllwt@126.com", "woluty@163.com", "a.lun.x@qq.com", "1377388181@qq.com", "cannykin@126.com", "gz20080912@163.com", "358158708@qq.com", "zbc2003@163.com", "hfm500@126.com", "zhengbc@asiainfo.com", "salonzj@qq.com", "liufeng09023@163.com", "tanqi.g@gmail.com", "505467694@qq.com", "zhxzhou@foxmail.com", "693169604@qq.com", "holy4666@163.com", "lws8712@126.com", "jguy@163.com", "hgming@cyberway.net.cn", "389715286@qq.com", "1347555209@qq.com", "wuqiang.123@163.com", "807804546@qq.com", "13837975299@139.com", "15807644873@139.com", "1006322727@qq.com", "250678832@qq.com", "workformydream@yahoo.com", "827931687@qq.com", "wbliu@scnu.edu.cn", "12377536@qq.com", "wqs1603@sina.com", "rulai@iflyun.com", "lenka2010@163.com", "1162350543@qq.com", "guoxiaobin@tass.com.cn", "everforever@163.com", "36319954@qq.com", "962860826@qq.com", "shaoliangfang@126.com", "13763394583@139.com", "richard01lee@126.com", "xuehy@asiainfo.com", "13275201885@163.com", "463526549@qq.com", "153813274@qq.com", "roadlit@qq.com", "1406872781@qq.com", "87855450@qq.com", "254176824@qq.com", "124839545@qq.com"}

	//HEB_User = []string{"lijy@hrbdataex.com"}

	//datahub_User = []string{"hehl@asiainfo.com", "wangjun15@asiainfo.com", "gongjing5@asiainfo.com", "caory@asiainfo.com", "chenggk@asiainfo.com", "songliang@asiainfo.com", "hext3@asiainfo.com", "luozs@asiainfo.com", "shendf@asiainfo.com", "wangzheng3@asiainfo.com", "yepeng@asiainfo.com", "mengjing@asiainfo.com", "liujie15@asiainfo.com", "wangdi6@asiainfo.com", "liuxy10@asiainfo.com", "liuxu@asiainfo.com", "panxy3@asiainfo.com", "chaizs@asiainfo.com", "yuanwm@asiainfo.com", "datahub@asiainfo.com", "shuya@asiainfo.com", "yaxin@asiainfo.com", "huaysj@asiainfo.com", "fangliang@asiainfo.com", "zhongyi@asiainfo.com", "yinlian@asiainfo.com", "zhongke@asiainfo.com", "pengbai@asiainfo.com", "bjjtgl@asiainfo.com", "bjhb@asiainfo.com", "bjedu@asiainfo.com", "bjta@asiainfo.com", "bjrf@asiainfo.com", "bjtsb@asiainfo.com", "datashanghai@asiainfo.com", "dangsha@asiainfo.com", "wangxn5@asiainfo.com", "liull7@asiainfo.com", "cnemc@asiainfo.com", "metdata@asiainfo.com", "icttic@asiainfo.com", "chenyg@asiainfo.com", "inter3i@asiainfo.com", "chelaile@asiainfo.com", "wis@asiainfo.com", "hisign@asiainfo.com", "longcredit@asiainfo.com", "syntun@asiainfo.com", "bjzjkj@asiainfo.com", "getui@asiainfo.com", "transilink@asiainfo.com", "datatang@asiainfo.com", "questmobile@asiainfo.com", "weidata@asiainfo.com", "dratio@asiainfo.com", "cucrz@asiainfo.com", "pcd@asiainfo.com", "wuping3@asiainfo.com", "zhengmq@asiainfo.com", "zhangyx5@asiainfo.com", "liming3@asiainfo.com", "huacj@asiainfo.com", "zhangman3@asiainfo.com", "huaxiang@asiainfo.com", "sundy@asiainfo.com", "huangping@asiainfo.com", "luowei5@asiainfo.com", "wangli20@asiainfo.com", "lizheng3@asiainfo.com", "youyong@asiainfo.com", "weizy@asiainfo.com", "liurj@asiainfo.com", "tangwei@asiainfo.com", "niulin@asiainfo.com", "wangjh11@asiainfo.com", "liuquan2@asiainfo.com", "datawuxi@asiainfo.com", "datazj@asiainfo.com", "zhucce@asiainfo.com", "jinka@asiainfo.com", "zuanshi@asiainfo.com", "aaa@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "wuyw@asiainfo.com", "zhangju@asiainfo.com", "wuqiong@asiainfo.com", "zhangnai@asiainfo.com", "lijing9@asiainfo.com", "wangna7@asiainfo.com", "jinjin@asiainfo.com", "jingxy3@asiainfo.com", "caoxh@asiainfo.com", "xuying@asiainfo.com", "yangyn@asiainfo.com", "ruanys@asiainfo.com", "gaowei@asiainfo.com", "zhangkb@asiainfo.com", "zhouxy@asiainfo.com", "chenff@asiainfo.com", "daimh@asiainfo.com", "wangeh@asiainfo.com", "yangws@asiainfo.com", "shayf@asiainfo.com", "taoxj@asiainfo.com", "zhangyuwen@asiainfo.com", "lij@asiainfo.com", "zoumd@asiainfo.com", "wangdi@asiainfo.com", "penghx2@asiainfo.com", "lizc5@asiainfo.com", "bankofchina@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "financial@asiainfo.com", "guocf5@asiainfo.com", "lijy15@asiainfo.com", "ligh3@asiainfo.com", "chenlei6@asiainfo.com", "qianjf@asiainfo.com", "zhangyu@inter3i.com", "wanghl7@asiainfo.com", "zengzy3@asiainfo.com", "lijl@asiainfo.com", "bbb@asiainfo.com", "liucui@asiainfo.com", "zhangl@mapuni.com", "melvin@syntun.com", "zhangxu@mininglamp.com", "yunpeng.na@bluefocus.com", "zhushaosong@4paradigm.com", "liuqiuwen@hiynn.com", "chenxin@acmr.com.cn", "hong.zheng@kuaiyou.com", "wei.sun@baifendian.com", "zhouquan3@asiainfo.com", "linjiajie@geotmt.com", "dingzhenghong@weidata.cn", "bbb@asiainfo.com", "bbb@asiainfo.com", "repository_9bacb3c5dt@asiainfo.com", "subscription_39d41a795c@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "tangliang6@asiainfo.com", "liull@asiainfo.com", "lijj3@asiainfo.com", "nanhaizf@asiainfo.com", "lijj3@asiainfo.com", "panwj3@asiainfo.com", "cnta@asiainfo.com", "hanzc@asiainfo.com", "fengjun5@asiainfo.com", "penghui.li@huaat.net", "bigdata@chinawiserv.com", "wangyu13@asiainfo.com", "taolei@asiainfo.com", "lipeng7@asiainfo.com", "admin_chenyg@asiainfo.com", "fengq@neunn.com", "liujuan@mapuni.com", "linhuajing@longshine.com", "ivan173@163.com", "zuoli@asiainfo.com", "yuancheng@asiainfo.com", "shilin@asiainfo.com", "jiangjie5@cjbigdata.com", "yaozy@asiainfo.com", "duyl@asiainfo.com", "liuhao3@asiainfo.com", "yuwei@asiainfo.com", "liulf3@asiainfo.com", "liuxk2@asiainfo.com", "yincy@asiainfo.com", "zjuwwq@163.com", "liubq@asiainfo.com", "liyx5@asiainfo.com", "yujin3@asiainfo.com", "guoyy6@asiainfo.com", "wangwei17@asiainfo.com", "hzwenwen@corp.netease.com", "xiaofei9670@foxmail.com", "852028265@qq.com", "wjwxyh_007@126.com", "liulihong@cstor.cn", "aistxianyang@sohu.com", "375921332@qq.com", "nshen_nbip@163.com", "wangchao@daasbank.com", "zhangyz@daasbank.com", "sunhl3@asiainfo.com", "jesse526@163.com", "raoshucheng87@163.com", "zzrjianli@163.com", "chen5575658@163.com", "zhaoyan8@asiainfo.com", "573385501@qq.com", "fangyl5@asiainfo.com", "14682253@qq.com", "lingnanpass@163.com", "273227864@qq.com", "2629314906@qq.com", "jiangnan@finefurther.com", "zh@efala.com", "13560366758@139.com", "wuyivs@126.com", "chen@murren.cn", "huajunsir@zju.edu.cn", "mixdata_dh@asiainfo-mixdata.com", "chenguo3@asiainfo.com", "kkkev@sina.com", "jim_zw@163.com", "gchf7209@sohu.com", "zhaoxy6@asiainfo.com", "2822422913@qq.com", "37574708@qq.com", "zhizaokeji@163.com", "181820265@qq.com", "jingqin@asiainfo.com", "790199623@qq.com", "53901913@qq.com", "15126128@bjtu.edu.cn", "1@1.com", "871810908@qq.com", "cuiwenchang123@163.com", "immortal_phoenix@126.com", "1069478446@qq.com", "chengkj@asiainfo.com", "whulawzy1991@163.com", "info@hrbdataex.com", "mxp@8531.cn", "chengkj99@foxmail.com", "lily001@qq.com", "347566078@qq.com", "lily00052@qq.com", "34566@qq.com", "277081510@qq.com", "bdx-psd@asiainfo.com", "771435128@qq.com", "guolq@asiainfo.com", "wangyd3@asiainfo.com"}

	WH_User = []string{"948689655@qq.com", "shendf@primeton.com", "gchf_1@sina.com", "guocf5@cjbigdata.com", "service@cjbigdata.com", "qiuzhihua@gbase.cn", "service-cj@cjbigdata.com", "chengkj@yeah.net", "18810155806@163.com", "lixw@maitewang.com", "member01@cjbigdata.com", "fengmn@asiainfo.com", "15126128@bjtu.edu.cn", "1312414619@qq.com", "yehovah@sohu.com", "qijing@cjbigdata.com", "whulawzy1991@163.com", "readystar@163.com", "121266485@qq.com", "houxiaosa@51daas.com", "13581884086@126.com", "zxw@tenly.com"}

	GZ_User = []string{"service@gzbdex.com", "wangdc@gzbdex.com", "wang_dachao@163.com", "13902292121@139.com", "wdc-shuiping@163.com", "zhenchao.zhao@kcomber.com", "sjtan@sheendu.com", "info@yiyantang.cn", "hujq@asiainfo.com", "feng.he@heqinuc.com", "liuhf@ibbd.net", "yinhuacheng@linkip.cn", "xiexf@189.cn", "skyrock@139.com", "mailiantao@qq.com", "13822147976@139.com", "shihw11@sina.com", "walrkson@hotmail.com", "838959@qq.com", "56463383@qq.com", "13311288822@189.cn", "7904865@qq.com", "112793863@qq.com", "47181770@qq.com", "yizf@szlanyou.com", "cyllwt@126.com", "woluty@163.com", "a.lun.x@qq.com", "1377388181@qq.com", "cannykin@126.com", "gz20080912@163.com", "358158708@qq.com", "zbc2003@163.com", "hfm500@126.com", "zhengbc@asiainfo.com", "salonzj@qq.com", "liufeng09023@163.com", "tanqi.g@gmail.com", "505467694@qq.com", "zhxzhou@foxmail.com", "693169604@qq.com", "holy4666@163.com", "lws8712@126.com", "jguy@163.com", "hgming@cyberway.net.cn", "389715286@qq.com", "1347555209@qq.com", "wuqiang.123@163.com", "807804546@qq.com", "13837975299@139.com", "15807644873@139.com", "1006322727@qq.com", "250678832@qq.com", "workformydream@yahoo.com", "827931687@qq.com", "wbliu@scnu.edu.cn", "12377536@qq.com", "wqs1603@sina.com", "rulai@iflyun.com", "lenka2010@163.com", "1162350543@qq.com", "guoxiaobin@tass.com.cn", "everforever@163.com", "36319954@qq.com", "962860826@qq.com", "shaoliangfang@126.com", "13763394583@139.com", "richard01lee@126.com", "xuehy@asiainfo.com", "13275201885@163.com", "463526549@qq.com", "153813274@qq.com", "roadlit@qq.com", "1406872781@qq.com", "87855450@qq.com", "254176824@qq.com", "124839545@qq.com", "424112052@qq.com", "caoc@gzbdex.com", "790199623@qq.com", "Zhoukr@gzbdex.com", "hr@gzbdex.com", "chengkj999@foxmail.com", "chengkj@yeah.com", "kangjian281630@hotmail.com", "chengkj99@gmail.com", "15907882262@163.com", "tttick@163.com", "ganry@gzbdex.com", "136343045@qq.com", "13922248989@139.com", "yaochitc@163.com", "821671094@qq.com"}

	HEB_User = []string{"info@hrbdataex.com", "520520520@qq.com", "lijy@hrbdataex.com", "130333999@qq.com", "923423511@qq.com", "835652965@qq.com"}

	datahub_User = []string{"hehl@asiainfo.com", "gongjing5@asiainfo.com", "caory@asiainfo.com", "chenggk@asiainfo.com", "songliang@asiainfo.com", "luozs@asiainfo.com", "yepeng@asiainfo.com", "mengjing@asiainfo.com", "liujie15@asiainfo.com", "wangdi6@asiainfo.com", "liuxu@asiainfo.com", "panxy3@asiainfo.com", "chaizs@asiainfo.com", "yuanwm@asiainfo.com", "datahub@asiainfo.com", "shuya@asiainfo.com", "yaxin@asiainfo.com", "huaysj@asiainfo.com", "fangliang@asiainfo.com", "zhongyi@asiainfo.com", "yinlian@asiainfo.com", "zhongke@asiainfo.com", "pengbai@asiainfo.com", "bjjtgl@asiainfo.com", "bjhb@asiainfo.com", "bjedu@asiainfo.com", "bjta@asiainfo.com", "bjrf@asiainfo.com", "bjtsb@asiainfo.com", "datashanghai@asiainfo.com", "dangsha@asiainfo.com", "wangxn5@asiainfo.com", "liull7@asiainfo.com", "cnemc@asiainfo.com", "metdata@asiainfo.com", "icttic@asiainfo.com", "chenyg@asiainfo.com", "inter3i@asiainfo.com", "chelaile@asiainfo.com", "wis@asiainfo.com", "hisign@asiainfo.com", "longcredit@asiainfo.com", "syntun@asiainfo.com", "bjzjkj@asiainfo.com", "getui@asiainfo.com", "transilink@asiainfo.com", "datatang@asiainfo.com", "questmobile@asiainfo.com", "weidata@asiainfo.com", "dratio@asiainfo.com", "cucrz@asiainfo.com", "pcd@asiainfo.com", "wuping3@asiainfo.com", "zhengmq@asiainfo.com", "zhangyx5@asiainfo.com", "liming3@asiainfo.com", "huacj@asiainfo.com", "zhangman3@asiainfo.com", "huaxiang@asiainfo.com", "sundy@asiainfo.com", "huangping@asiainfo.com", "luowei5@asiainfo.com", "wangli20@asiainfo.com", "lizheng3@asiainfo.com", "youyong@asiainfo.com", "weizy@asiainfo.com", "liurj@asiainfo.com", "tangwei@asiainfo.com", "niulin@asiainfo.com", "wangjh11@asiainfo.com", "liuquan2@asiainfo.com", "datawuxi@asiainfo.com", "datazj@asiainfo.com", "zhucce@asiainfo.com", "jinka@asiainfo.com", "zuanshi@asiainfo.com", "aaa@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "wuyw@asiainfo.com", "zhangju@asiainfo.com", "wuqiong@asiainfo.com", "zhangnai@asiainfo.com", "lijing9@asiainfo.com", "jinjin@asiainfo.com", "jingxy3@asiainfo.com", "caoxh@asiainfo.com", "xuying@asiainfo.com", "yangyn@asiainfo.com", "ruanys@asiainfo.com", "gaowei@asiainfo.com", "zhangkb@asiainfo.com", "zhouxy@asiainfo.com", "chenff@asiainfo.com", "daimh@asiainfo.com", "wangeh@asiainfo.com", "yangws@asiainfo.com", "shayf@asiainfo.com", "taoxj@asiainfo.com", "zhangyuwen@asiainfo.com", "lij@asiainfo.com", "zoumd@asiainfo.com", "wangdi@asiainfo.com", "penghx2@asiainfo.com", "lizc5@asiainfo.com", "bankofchina@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "financial@asiainfo.com", "guocf5@asiainfo.com", "lijy15@asiainfo.com", "ligh3@asiainfo.com", "chenlei6@asiainfo.com", "qianjf@asiainfo.com", "zhangyu@inter3i.com", "wanghl7@asiainfo.com", "zengzy3@asiainfo.com", "lijl@asiainfo.com", "bbb@asiainfo.com", "liucui@asiainfo.com", "zhangl@mapuni.com", "melvin@syntun.com", "zhangxu@mininglamp.com", "yunpeng.na@bluefocus.com", "zhushaosong@4paradigm.com", "liuqiuwen@hiynn.com", "chenxin@acmr.com.cn", "hong.zheng@kuaiyou.com", "wei.sun@baifendian.com", "zhouquan3@asiainfo.com", "linjiajie@geotmt.com", "dingzhenghong@weidata.cn", "bbb@asiainfo.com", "bbb@asiainfo.com", "repository_9bacb3c5dt@asiainfo.com", "subscription_39d41a795c@asiainfo.com", "bbb@asiainfo.com", "bbb@asiainfo.com", "tangliang6@asiainfo.com", "liull@asiainfo.com", "lijj3@asiainfo.com", "nanhaizf@asiainfo.com", "lijj3@asiainfo.com", "panwj3@asiainfo.com", "cnta@asiainfo.com", "hanzc@asiainfo.com", "fengjun5@asiainfo.com", "penghui.li@huaat.net", "bigdata@chinawiserv.com", "wangyu13@asiainfo.com", "taolei@asiainfo.com", "lipeng7@asiainfo.com", "admin_chenyg@asiainfo.com", "fengq@neunn.com", "liujuan@mapuni.com", "linhuajing@longshine.com", "ivan173@163.com", "zuoli@asiainfo.com", "yuancheng@asiainfo.com", "shilin@asiainfo.com", "jiangjie5@cjbigdata.com", "yaozy@asiainfo.com", "duyl@asiainfo.com", "liuhao3@asiainfo.com", "yuwei@asiainfo.com", "liulf3@asiainfo.com", "liuxk2@asiainfo.com", "yincy@asiainfo.com", "zjuwwq@163.com", "liubq@asiainfo.com", "liyx5@asiainfo.com", "yujin3@asiainfo.com", "guoyy6@asiainfo.com", "wangwei17@asiainfo.com", "hzwenwen@corp.netease.com", "xiaofei9670@foxmail.com", "852028265@qq.com", "wjwxyh_007@126.com", "liulihong@cstor.cn", "aistxianyang@sohu.com", "375921332@qq.com", "nshen_nbip@163.com", "wangchao@daasbank.com", "zhangyz@daasbank.com", "sunhl3@asiainfo.com", "jesse526@163.com", "raoshucheng87@163.com", "zzrjianli@163.com", "chen5575658@163.com", "zhaoyan8@asiainfo.com", "573385501@qq.com", "fangyl5@asiainfo.com", "14682253@qq.com", "lingnanpass@163.com", "273227864@qq.com", "2629314906@qq.com", "jiangnan@finefurther.com", "zh@efala.com", "13560366758@139.com", "wuyivs@126.com", "chen@murren.cn", "huajunsir@zju.edu.cn", "mixdata_dh@asiainfo-mixdata.com", "chenguo3@asiainfo.com", "kkkev@sina.com", "jim_zw@163.com", "gchf7209@sohu.com", "zhaoxy6@asiainfo.com", "2822422913@qq.com", "37574708@qq.com", "zhizaokeji@163.com", "181820265@qq.com", "771435128@qq.com", "jingqin@asiainfo.com", "53901913@qq.com", "1@1.com", "871810908@qq.com", "cuiwenchang123@163.com", "immortal_phoenix@126.com", "1069478446@qq.com", "chengkj@asiainfo.com", "mxp@8531.cn", "chengkj99@foxmail.com", "lily001@qq.com", "347566078@qq.com", "lily00052@qq.com", "34566@qq.com", "bdx-psd@asiainfo.com", "Hubei@126.com", "numberone23@sina.cn", "676306924@qq.com", "chengkangjian@sxwdsoft.com", "780364570@qq.com", "1261695934@qq.com", "huyouzhusl@163.com", "781759037@qq.com", "112344@111.com", "18507739693@163.com", "277081510@qq.com", "446034911@qq.com", "sunny.hl@163.com", "282760835@qq.com", "guaiguai0928@126.com", "626060807@qq.com", "wangjz@mapuni.com"}
)

func main() {
	if initDB() == false {
		Log.Fatal("failed to connect mongodb.")
		return
	}

	go refreshDB()

	//getRepoByUser(db.copy(), "yuanwm@asiainfo.com")
	updateAll()

}

func getRepoByUser(db *DB, user string) {
	defer db.Close()
	Q := bson.M{COL_CREATE_USER: "datahub+" + user}
	n, _ := db.DB(DB_NAME).C(C_REPOSITORY).Find(Q).Count()
	Log.Info(n)

	res := []repository{}
	if err := db.DB(DB_NAME).C(C_REPOSITORY).Find(Q).All(&res); err != nil {
		Log.Info(err)
	}
	Log.Info("res:", res)

	db.DB(DB_NAME).C(C_REPOSITORY).UpdateAll(Q, bson.M{CMD_SET: bson.M{COL_CREATE_USER: user}})
}

func updateAll() {
	for _, v := range WH_User {
		updateAllCollection(v, "WH+")
	}

	for _, v := range GZ_User {
		updateAllCollection(v, "GZ+")
	}

	for _, v := range HEB_User {
		updateAllCollection(v, "HEB+")
	}

	for _, v := range datahub_User {
		updateAllCollection(v, "datahub+")
	}
}

func updateAllCollection(user, prefix string) {
	Q := bson.M{COL_CREATE_USER: user}
	U := bson.M{CMD_SET: bson.M{COL_CREATE_USER: prefix + user}}
	if _, e := db.DB(DB_NAME).C(C_REPOSITORY).UpdateAll(Q, U); e != nil {
		Log.Error(e)
	}
	if _, e := db.DB(DB_NAME).C(C_DATAITEM).UpdateAll(Q, U); e != nil {
		Log.Error(e)
	}

	Q_p := bson.M{COL_PERMIT_USER: user}
	U_p := bson.M{CMD_SET: bson.M{COL_PERMIT_USER: prefix + user}}
	if _, e := db.DB(DB_NAME).C(C_REPOSITORY_PERMISSION).UpdateAll(Q_p, U_p); e != nil {
		Log.Error(e)
	}
	if _, e := db.DB(DB_NAME).C(C_DATAITEM_PERMISSION).UpdateAll(Q_p, U_p); e != nil {
		Log.Error(e)
	}
}
