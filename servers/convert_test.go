package servers_test

import (
	"fmt"
	"testing"

	"github.com/HyeockJinKim/jira-transform/servers"
)

func TestConvert(t *testing.T) {
	jiraText := `
h1. Main Title

h2. Section Title

h3. Subsection

This is a paragraph with _italic text_, *bold text*, and {{inline code}}.

{code:javascript}
function greet() {
    console.log("Hello, JIRA to Markdown!");
}
{code}

{quote}
This is a blockquote in JIRA format.
{quote}

* Unordered list item 1
* Unordered list item 2

# Ordered list item 1
# Ordered list item 2

||*c 1*||*c2*||*c3*||
|1|2|3|
|4|5|6|

[OpenAI|https://www.openai.com]

!https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTBYnUypPWliqJgXA4GqfrO6ptOT5pK2GLJ-A&s!
`

	res := servers.JiraToMD(jiraText)
	fmt.Printf("%s\n", res)
}
