package roll

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Tabler interface is used to define all roll.Tables that can be rolled or printed.
type Tabler interface {
	Label() string
	Roll() string
	String() string
}

// TableRegistry allows multiple tables to be registered and called as
// item actions from other tables. It's particularly useful when
// an item.Action needs to refer to its parent table.
/* For Example:

    var r = roll.NewRegistry()

    var t = roll.Table{
	    Name: "Test",
	    ID:   "Parent",
	    Dice: "1d6",
	    Reroll: roll.Reroll{
		    Match: []int{5, 6},
		    Dice:  "1d4",
	    },
	    Items: []roll.TableItem{
		    {Match: []int{1}, Text: "TableItem 1", Action: func() string {
			    tbl, _ := r.Get("Parent")
			    return tbl.Roll()
		    }},
		    {Match: []int{2}, Text: "TableItem 2"},
		    {Match: []int{3}, Text: "TableItem 3"},
		    {Match: []int{4}, Text: "TableItem 4"},
		    {Match: []int{5}, Text: "TableItem 5"},
		    {Match: []int{6}, Text: "TableItem 6"},
	    },
    }

    var t1 = roll.Table{
	    Name: "Subtable 1",
	    ID:   "Child 1",
	    Dice: "1d6",
	    Items: []roll.TableItem{
		    {Match: []int{1}, Text: "TableItem 1.1"},
		    {Match: []int{2}, Text: "TableItem 2.1"},
		    {Match: []int{3}, Text: "TableItem 3.1"},
		    {Match: []int{4}, Text: "TableItem 4.1"},
		    {Match: []int{5}, Text: "TableItem 5.1"},
		    {Match: []int{6}, Text: "TableItem 6.1"},
	    },
    }

    func main() {
	    r.Add(t)
	    r.Add(t1)

	    //	fmt.Printf("%+v\n", t.Items)
	    for i := 0; i < 10; i++ {
		    fmt.Println(t.Roll())
    	}
		}

*/
type TableRegistry map[string]Table

// NewTableRegistry returns a new registry
func NewTableRegistry() TableRegistry {
	return make(TableRegistry)
}

// Add a table to the TableRegistry
func (r TableRegistry) Add(t Table) error {
	if _, ok := r[t.ID]; !ok {
		r[t.ID] = t
		return nil
	}

	return fmt.Errorf("table %s already registered", t.ID)
}

// Remove a table from the TableRegistry
func (r TableRegistry) Remove(id string) error {
	if _, ok := r[id]; ok {
		delete(r, id)
		return nil
	}

	return fmt.Errorf("table %s is not registered", id)
}

// Get a table from the registry
func (r TableRegistry) Get(id string) (Table, error) {
	t, ok := r[id]
	if !ok {
		return Table{}, fmt.Errorf("no table registered with id [%s]", id)
	}

	return t, nil
}

// Table represents a table of text options that can be rolled on. Name is
// optional. Tables are preferable to Lists when using multiple dice to achieve
// a result (i.e 2d6) because their results fall on a bell curve whereas single-die
// rolls have an even probability.
type Table struct {
	ID     string // Shorthand ID for finding subtables
	Name   string
	Dice   Dice
	Mod    int
	Reroll TableReroll
	Items  []TableItem
}

// TableReroll describes conditions under which the table should be rolled on again, using a different dice value
type TableReroll struct {
	Match TableMatchSet
	Dice  Dice
}

// TableItem represents the text and matching numbers from the table
type TableItem struct {
	Match  TableMatchSet
	Text   string
	Action func() string
}

type TableMatchSet []int

func (m TableMatchSet) Contains(n int) bool {
	for _, i := range m {
		if i == n {
			return true
		}
	}

	return false
}

func (m TableMatchSet) String() string {
	var s []string

	for _, n := range m {
		s = append(s, strconv.Itoa(n))
	}

	return strings.Join(s, ", ")
}

// Roll on the table and return the option drawn.
func (t Table) Roll() string {
	out := ""

	r := t.Dice.Roll()
	n := r.Sum() + t.Mod
	if n < t.Dice.Min() {
		n = t.Dice.Min()
	}
	if n > t.Dice.Max() {
		n = t.Dice.Max()
	}

	// Record initial roll result
	for _, i := range t.Items {
		if i.Match.Contains(n) {
			out = i.Text
		}
	}

	// Check for a reroll
	if t.Reroll.Match.Contains(n) {
		r = t.Reroll.Dice.Roll()
		n = r.Sum()
		for _, i := range t.Items {
			if i.Match.Contains(n) {
				if out != "" {
					out += "; "
				}
				out += i.Text
			}
		}
	}

	// Append text for final roll result
	for _, i := range t.Items {
		if i.Match.Contains(n) {
			if i.Action != nil {
				if out != "" {
					out += "; "
				}

				out += i.Action()
			}
			return out
		}
	}

	return ""
}

func (t Table) String() string {
	var (
		buf = new(bytes.Buffer)
		tw  = tabwriter.NewWriter(buf, 2, 2, 1, ' ', 0)
	)

	fmt.Fprintln(tw, "Dice\t|\tText")
	for _, i := range t.Items {
		fmt.Fprintf(tw, "%s\t|\t%s\n", i.Match, i.Text)
	}
	tw.Flush()

	return buf.String()
}

// Label returns the table Name
func (t Table) Label() string {
	return t.Name
}

// List represents a List of strings from which something can be selected at random
type List struct {
	Name  string
	Items []string
}

// Roll returns a random string from List
func (l List) Roll() string {
	if len(l.Items) > 0 {
		return l.Items[rand.Intn(len(l.Items))]
	}

	return ""
}

func (l List) String() string {
	return strings.Join(l.Items, ", ")
}

// Label returns the list Name
func (l List) Label() string {
	return l.Name
}

// MatchRange produces a TableMatchSet from start to end inclusive in order to make it easier
// to match large sets of numbers without having to type them all in
func MatchRange(start, end int) TableMatchSet {
	var out TableMatchSet

	for i := start; i <= end; i++ {
		out = append(out, i)
	}

	return out
}
