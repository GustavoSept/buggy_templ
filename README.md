## How to run

### Build

Have docker compose installed

```bash
docker compose up -d --build
```

### See

Open it in a browser on `localhost:4444`

## Bug Description

### Initial context

I'm building an app that involves rendering the periodic table. I'm new to templ (this is my first project on it).

### Bug Description

#### The code

Each chemical element is a component. I'm rendering them using this code (notice how, for each sub-div, we're creating a new class):

```templ
package components

import (
    "webapp/components/ui"
    "webapp/utils/types"
    "webapp/utils/constants"
    "strconv"
)

css elementWrapperStyle(element types.Element){
    grid-column: { strconv.Itoa(element.ColNum) };
    grid-row: { strconv.Itoa(element.RowNum) };
}

templ Home() {
    @BaseBare("home") {
        <div class="periodicTableGrid">
            for _, element := range constants.PeriodicTable {
                <div class={ elementWrapperStyle(element) }>
                    @ui.ElementBlock(element)
                </div>
            }
        </div>
    }
}
```

I have a hard-coded Array of Literals containing the periodic table. It looks like this:

```go
var PeriodicTable = []types.Element{
	{Name: "Hydrogen", Symbol: "H", AtomicNumber: 1, Group: 1, Period: 1, Category: "Nonmetal", Color: "#30c93d", ColNum: 1, RowNum: 1},
	{Name: "Helium", Symbol: "He", AtomicNumber: 2, Group: 18, Period: 1, Category: "Noble Gas", Color: "#06bbe8", ColNum: 18, RowNum: 1},
	{Name: "Lithium", Symbol: "Li", AtomicNumber: 3, Group: 1, Period: 2, Category: "Alkali Metal", Color: "#eda30e", ColNum: 1, RowNum: 2},

	// Skipping elements for conciseness

	// This one will be important later
	{Name: "Thallium", Symbol: "Tl", AtomicNumber: 81, Group: 13, Period: 6, Category: "Post-transition Metal", Color: "#2181bb", ColNum: 13, RowNum: 6},
}
```

#### Actual behavior

The issue is that this code consistently generates a single naming clash when templ generates the class for Hydrogen and Thallium specifically.

Here's a comparison of outputs:

| Expected Table                             | Actual Table                             |
| ------------------------------------------ | ---------------------------------------- |
| ![Expected](https://imgur.com/nH7I20M.png) | ![Actual](https://imgur.com/bN8GmO6.png) |

And the generated HTML for the buggy table:
<img src="https://imgur.com/aOfgX2M.png" width="500px">

Notice how Hydrogen and Thallium have the same class name `class="elementWrapperStyle_f781"`. It is correctly generated for Hydrogen, at position (1,1).

#### Things I tried

- **Removing Hydrogen**: Hydrogen is excluded as expected, and Thallium moves back to its expected position!
- **Altering Hydrogen coordinates**: Altering either Hydrogen's ColNum/RolNum (or both), to any coordinate (including 0,0), will move Thallium to its expected position.
  - Moving Hydrogen to 0,0 actually renders the Expected Table properly. But since it's unexpected behavior, I'm not using this solution.
- **Adding a new element to (1,1)**: Thallium keeps being rendered on (1,1).
- **Removing a random element other than Hydrogen**: Thallium keeps being rendered on (1,1).
- **Moving Thallium to another coordinate**: Thallium gets moved as expected to the location.
  - If it's a coordinate that already contains an element, they render one on top of the other (which is expected).
- 🚨 **_Swapping a random element's coordinate with Thallium_** 🚨: Thallium gets placed where expected. The other element is rendered on top of Hydrogen.
  - This leads me to believe the problem internally has something to do with the coordinate(13,6) in specific.
  - The generated class name that is clashing both elements in this case is still the same `elementWrapperStyle_f781` as the Hydrogen/Thallium case.

#### Hypothesis

I don't know exactly how templ works internally, as I'm new to it. I suppose this is a Hashing issue when generating the name, hence why it's so consistent on outputting the same result, given the same input.
Am i just unlucky that this particular combination of bytes is outputting the same Hash for two different rows, as long as they have the same coordinates?

To test this hypothesis, I've added a third random element to the generated classes:

```templ
css elementWrapperStyle(element types.Element) {
    grid-column: { strconv.Itoa(element.ColNum) };
    grid-row: { strconv.Itoa(element.RowNum) };
    background-color: { element.Color };
}
```

And this also solves the issue! All elements are rendered where they should (and have a terrible look to them. But this was just to test the hypothesis).
![Uggly Table](https://imgur.com/wh7AvK5.png)

### More info

**`templ info` output**
Run `templ info` and include the output:

```bash
(✓) os [ goos=linux goarch=amd64 ]
(✓) go [ location=/home/gustsept/.asdf/shims/go version=go version go1.23.2 linux/amd64 ]
(✓) gopls [ location=/home/gustsept/.asdf/installs/golang/1.23.2/packages/bin/gopls version=golang.org/x/tools/gopls v0.16.2 ]
(✓) templ [ location=/home/gustsept/.asdf/installs/golang/1.23.2/packages/bin/templ version=v0.2.793 ]
```

**Desktop (please complete the following information):**

- OS: Linux (Bookworm Debian 12.7)
- templ CLI version (`templ version`): v0.2.793
- Go version (`go version`): go version go1.23.2 linux/amd64
- `gopls` version (`gopls version`): golang.org/x/tools/gopls v0.16.2

###s# Observation
I've fixed my issue, by implementing the coordinates directly in the div's style as an attribute, like this:

```templ
func elementWrapperStyle(element types.Element) templ.Attributes {
    return templ.Attributes{
        "style": "grid-column: " + strconv.Itoa(element.ColNum) + "; grid-row: " + strconv.Itoa(element.RowNum) + ";",
    }
}

templ Home() {
    @BaseBare("home") {
        <div class="periodicTableGrid">
            for _, element := range constants.PeriodicTable {
                <div { elementWrapperStyle(element)... }>
                    @ui.ElementBlock(element)
                </div>
            }
        </div>
    }
}
```

This solves my problem for now. But what if I actually needed dozens or hundreds of different CSS class names?
