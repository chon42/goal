package algo_reg


type RegEntry struct {
	Name   string `json:"Name"`
	Schema string `json:"Schema"`
	//Interface func(interface{}) interface{}
	Interface interface{}
}
type RegEntryType = RegEntry
type RegMap map[string]RegEntryType

var Empty = RegEntry{}

var CentralRegistry = RegEntry{
	Name:      "Central",
	Schema:    "Registry",
	Interface: make(RegMap),
}

func (reg RegEntry) GetRegEntry(name string, realm string) (RegEntryType, bool) {
	if reg.Schema == "Registry" {
		entry, ok := reg.Interface.(RegMap)
		if ok {
			regEntry, exists := entry[name]
			if realm == reg.Name {
				return regEntry, exists
			} else {
				regEntry.GetRegEntry(name, realm)
			}
		}
	}
	return Empty, false
}
func (reg RegEntry) AddRegEntry(name string,
	realm string, newRegEntry RegEntry) (RegEntryType, bool) {
	if reg.Schema == "Registry" {
		entry, ok := reg.Interface.(RegMap)
		if ok {
			regEntry, exists := entry[name]
			if exists {
				return Empty, false
			} else {
				if realm == reg.Name {
					entry[name] = newRegEntry
					return entry[name], true
				} else {
					regEntry.AddRegEntry(name, realm, newRegEntry)
				}
			}
		}
	}
	return Empty, false
}

func Register(name string, realm string, schema string,
	ifc ...interface{}) (interface{}, bool) {
	//method func(interface{}) interface{},schema string) (interface{},bool) {

	if _, exists := CentralRegistry.GetRegEntry(name, realm); exists {
		return nil, false
	}
	regEntry, ok := CentralRegistry.AddRegEntry(name, realm, RegEntry{
		Name:      name,
		Schema:    schema,
		Interface: ifc,
	})
	if ok {
		return regEntry.Interface, ok
	}

	return nil, false
}
