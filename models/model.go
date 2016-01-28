package models

import (
	"github.com/fuxiaohei/beego/config"
	"sort"
	"strconv"
	"strings"
)

var (
	HomeNav NavGroup
	DocNav  NavGroup
)

func Init() {
	iniFile, err := config.NewConfig("ini", "conf/nav.ini")
	if err != nil {
		panic(err)
	}
	homeNav, _ := iniFile.GetSection("nav")
	docNav, _ := iniFile.GetSection("doc-nav")
	HomeNav = loadNav(homeNav)
	DocNav = loadNav(docNav)
}

type (
	Nav struct {
		Order    int
		Key      string
		URL      string
		Children []*Nav
	}
	NavGroup []*Nav
)

func newNav(key, url string) *Nav {
	keyData := strings.Split(key, "-")
	order, _ := strconv.Atoi(keyData[len(keyData)-1])
	return &Nav{
		Order: order,
		Key:   key,
		URL:   url,
	}
}

func (ng NavGroup) Find(key string) *Nav {
	for _, n := range ng {
		if n.Key == key {
			return n
		}
	}
	return nil
}

func loadNav(m map[string]string) NavGroup {
	var navData NavGroup
	for k, v := range m {
		navData = append(navData, newNav(k, v))
	}

	sort.Sort(keyLengthNavSlice(navData))

	var (
		parentKey string
		topIdx    int
	)
	for i, n := range navData {
		idx := strings.LastIndex(n.Key, "-")
		if idx <= 0 {
			topIdx = i
			break
		}
		parentKey = n.Key[:idx]
		nav := navData.Find(parentKey)
		if nav == nil {
			continue
		}
		nav.Children = append(nav.Children, n)
	}
	navData = navData[topIdx:]
	sortNav(orderNavSlice(navData))
	return NavGroup(navData)
}

type keyLengthNavSlice []*Nav

func (k keyLengthNavSlice) Len() int           { return len(k) }
func (k keyLengthNavSlice) Less(i, j int) bool { return len(k[i].Key) > len(k[j].Key) }
func (k keyLengthNavSlice) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

type orderNavSlice []*Nav

func (o orderNavSlice) Len() int           { return len(o) }
func (o orderNavSlice) Less(i, j int) bool { return o[i].Order < o[j].Order }
func (o orderNavSlice) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

func sortNav(s orderNavSlice) {
	sort.Sort(s)
	for _, n := range s {
		if len(n.Children) > 0 {
			sort.Sort(orderNavSlice(n.Children))
		}
	}
}
