package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"github.com/miekg/dns"
	"os"
	"time"
)

type dnsEntry struct {
	ip   string
	port string
}

type DB struct {
	mgo.Session
}

func (db *DB) copy() *DB {
	return &DB{*db.Copy()}
}

func getMgoAddr() (string, string) {
	//entryList := dnsExchange(Service_Name_Mongo, DISCOVERY_CONSUL_SERVER_ADDR, DISCOVERY_CONSUL_SERVER_PORT)

	//if len(entryList) > 0 {
	//return entryList[0].ip, entryList[0].port
	//}
	//return "", ""
	return "172.17.0.10", "27017"
}

func initDB() bool {
	var err error
	for i := 0; i < 3; i++ {
		ip, port := getMgoAddr()
		url := fmt.Sprintf(`%s:%s/datahub?maxPoolSize=500`, ip, port)
		if _, err = mgo.Dial(url); err != nil {
			time.Sleep(time.Second * 10)
			continue
		} else {
			break
		}
	}
	if err != nil {
		return false
	}

	return true
}

func refreshDB() {

	for {
		select {
		case <-time.After(time.Second * 5):
			if err := db.Ping(); err != nil {
				Log.Infof("%s db connect err %s", time.Now().Format("2006-01-02 15:04:05"), err)
				db = DB{*connect()}
				db.Refresh()
			}
		}
	}
}

func dnsExchange(srvName, agentIp, agentPort string) []dnsEntry {
	Name := fmt.Sprintf("%s.service.consul", srvName)
	agentAddr := fmt.Sprintf("%s:%s", agentIp, agentPort)

	c := new(dns.Client)
	c.Net = "tcp"

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(Name), dns.TypeSRV)
	m.RecursionDesired = true

	result := []dnsEntry{}
	r, _, err := c.Exchange(m, agentAddr)

	if r == nil {
		Log.Fatalf("dns query %s error: %s\n", srvName, err.Error())
		return result
	}

	if r.Rcode != dns.RcodeSuccess {
		Log.Fatalf("dns query %s error: %v\n", srvName, r.Rcode)
		return result
	}

	for _, ex := range r.Extra {
		if tmp, ok := ex.(*dns.A); ok {
			result = append(result, dnsEntry{ip: tmp.A.String()})
		}
	}

	for i, an := range r.Answer {
		if tmp, ok := an.(*dns.SRV); ok {
			port := fmt.Sprintf("%d", tmp.Port)
			result[i].port = port
		}
	}

	return result
}

func Env(name string, required bool, showLog ...bool) string {
	s := os.Getenv(name)
	if required && s == "" {
		panic("env variable required, " + name)
	}
	if len(showLog) == 0 || showLog[0] {
		Log.Infof("[env][%s] %s\n", name, s)
	}
	return s
}

func connect() *mgo.Session {
	ip, port := getMgoAddr()
	if ip == "" {
		Log.Error("can not init mongo ip")
	}

	if port == "" {
		Log.Error("can not init mongo port")
	}

	url := fmt.Sprintf(`%s:%s/datahub?maxPoolSize=500`, ip, port)
	Log.Infof("[Mongo Addr] %s", url)

	var session *mgo.Session
	var err error
	try := 0
	for {
		ip, port = getMgoAddr()
		url = fmt.Sprintf(`%s:%s/datahub?maxPoolSize=500`, ip, port)
		session, err = mgo.Dial(url)
		if err != nil {
			try++
			Log.Errorf("dial mgo(%s) err %s, already try %d times", url, err.Error(), try)
			time.Sleep(time.Second * 10)
		} else {
			break
		}
	}

	return session
}
