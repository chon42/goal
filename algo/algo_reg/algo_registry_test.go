package algo_reg

import "fmt"

var reg = CentralRegistry

func algo_reg_main() {

	entry, ok := reg.AddRegEntry("test", "Central",
		RegEntry{"test", "int", 1})
	if ok {
		obj, okget := reg.GetRegEntry("test", "Central")
		if okget {
			fmt.Printf("Name: %s,Schema:%s,Value:%v",
				obj.Name, obj.Schema, obj.Interface)
		}
	}
	fmt.Println(entry.Name, entry.Schema, entry.Interface)

}
