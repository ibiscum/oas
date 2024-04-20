package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-faster/jx"
)

type Validation struct {
	Required bool `json:"required"`
}

type Value struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Shortcut string `json:"shortcut"`
}

type Data struct {
	Components        []Data      `json:"components,omitempty"`
	Collapsable       bool        `json:"collapsable"`
	Input             bool        `json:"input"`
	Key               string      `json:"key"`
	TableView         bool        `json:"tableView"`
	Type              string      `json:"type"`
	Value             string      `json:"value"`
	UseLocaleSettings bool        `json:"useLocaleSettings"`
	Values            []Value     `json:"values,omitempty"`
	Validate          *Validation `json:"validate,omitempty"`
}

var input = []byte(`{
    "components": [
        {
            "key": "d1",
            "components": [
                {
                    "key": "custname",
                    "value": "Abraham",
                    "input": true,
                    "tableView": true
                },
                {
                    "key": "type",
                    "type": "radio",
                    "label": "Fisrt",
                    "values": [
                        {
                            "label": "Sole",
                            "value": "sole",
                            "shortcut": ""
                        },
                        {
                            "label": "Bata",
                            "value": "Bata",
                            "shortcut": ""
                        }
                    ],
                    "validate": {
                        "required": true
                    },
                    "tableView": false
                },
                {
                    "key": "registeredField",
                    "value": "reg 111",
                    "input": true
                },
                {
                    "key": "dirc",
                    "value": "abraham"
                },
                {
                    "key": "gst",
                    "value": "textfield",                   
                    "useLocaleSettings": false
                },
                {
                    "key": "pan",
                    "value": "AAAAA0000",                    
                    "useLocaleSettings": false
                }
            ],
            "collapsable": false
        }
    ]
}`)

func main() {
	data := Data{}
	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Printing with simple recursive function")
	// print all components, these could be nested, so let's use a recursive function
	PrintComponents(data.Components)
	// fmt.Println("\n\nPrinting with indented recursion:")
	// PrintComponentsIndent(data.Components, "") // start with indent of 0

	d := jx.DecodeBytes(input)

	var values []jx.Raw
	if err := d.Obj(func(d *jx.Decoder, key string) error {
		switch key {
		case "components":
			// Iterate over each array element.
			return d.Arr(func(d *jx.Decoder) error {
				v, err := d.Raw()
				if err != nil {
					return err
				}
				values = append(values, v)
				return nil
			})
		default:
			// Skip unknown fields if any.
			return d.Skip()
		}
	}); err != nil {
		panic(err)
	}

	for _, v := range values {
		fmt.Println(v)
	}
}

func PrintComponents(data []Data) {
	for _, c := range data {
		if len(c.Components) > 0 {
			PrintComponents(c.Components) // recursive
			continue                      // skip value of this component, remove this line if needed
		}
		val := c.Value // assign string value
		if len(c.Values) > 0 {
			// this component has a slice of values, not a single value
			vals, err := json.MarshalIndent(c.Values, "", "    ") // marshal with indent of 4 spaces, no prefix
			if err != nil {
				fmt.Printf("Oops, looks like we couldn't format something: %+v\n", err)
				return // handle this
			}
			val = string(vals) // marshalled values as string
		}
		fmt.Printf("Key: %s Value: %s\n", c.Key, val) // print output
	}

}

func PrintComponentsIndent(data []Data, indent string) {
	for _, c := range data {
		if len(c.Components) > 0 {
			fmt.Printf("%sComponent block: %s\n", indent, c.Key)
			PrintComponentsIndent(c.Components, indent+"    ")
			continue
		}
		val := c.Value
		if len(c.Values) > 0 {
			// this component has a slice of values, not a single value
			vals, _ := json.MarshalIndent(c.Values, indent, "    ")
			val = string(vals) // marshalled values as string
		}
		fmt.Printf("%sKey: %s Value: %s\n", indent, c.Key, val) // print output
	}

}
