package ns

import "sort"

type Catalog struct {
	berths []string
}

func NewCatalog(berths ...string) *Catalog {
	dedup := map[string]bool{}
	unique := []string{}
	for _, code := range berths {
		if code == "" || dedup[code] {
			continue
		}
		dedup[code] = true
		unique = append(unique, code)
	}
	sort.Strings(unique)
	return &Catalog{berths: unique}
}

func (c *Catalog) Has(code string) bool {
	idx := sort.SearchStrings(c.berths, code)
	return idx < len(c.berths) && c.berths[idx] == code
}

func (c *Catalog) Codes() []string {
	return append([]string(nil), c.berths...)
}

func (c *Catalog) Count() int {
	return len(c.berths)
}
