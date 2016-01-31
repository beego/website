package models

import (
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego/config"
)

var (
	// HomeNav is navigation for all pages
	HomeNav NavGroup
	// DocNav is navigation for documentation pages
	DocNav NavGroup
)

// ReadNav read navigation data
func ReadNav(iniFile config.Configer) {
	homeNav, _ := iniFile.GetSection("nav")
	docNav, _ := iniFile.GetSection("doc-nav")
	HomeNav = loadNav(homeNav, "nav")
	DocNav = loadNav(docNav, "doc-nav")
}

type (
	// Nav is item of navigation
	Nav struct {
		Order      int
		Key        string
		URL        string
		Children   []*Nav
		i18nPrefix string
	}
	// NavGroup is group as navigation
	NavGroup []*Nav
)

func newNav(key, url, i18n string) *Nav {
	keyData := strings.Split(key, "-")
	order, _ := strconv.Atoi(keyData[len(keyData)-1])
	return &Nav{
		Order:      order,
		Key:        key,
		URL:        url,
		i18nPrefix: i18n,
	}
}

// I18n return the i18n key of this navigation item
func (n *Nav) I18n() string {
	return n.i18nPrefix + "::" + n.Key // in beego ini config, it support getting key in section by "section::key"
}

// IsURL return whether the link is in url, or just separator
func (n *Nav) IsURL() bool {
	return n.URL != "#"
}

// Find find navigation in group by key string
func (ng NavGroup) Find(key string) *Nav {
	for _, n := range ng {
		if n.Key == key {
			return n
		}
	}
	return nil
}

func loadNav(m map[string]string, i18nPrefix string) NavGroup {
	var navData NavGroup
	for k, v := range m {
		navData = append(navData, newNav(k, v, i18nPrefix))
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
