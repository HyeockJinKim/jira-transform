package convert_test

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

func TestConvert2(t *testing.T) {
	jiraText := `
h2. Motivation

[Why is this Task important? Explain the context and reasons for this task.]

h2. Objective

[Clearly state the purpose or goal of this task.]

h2. Details

[Outline the specific steps or changes needed to complete this task.]

h2. Impact

[Describe how this task may indirectly influence other services, if at all.]



h1. Heading1

h2. Heading2

h3. Heading3

h4. Heading4

h5. Heading5

h6. Heading6

* bullet1
** sub bullet
* bullet2
** sub bullet
*** sub-sub bullet
* bullet3
* bullet4

# index1
## sub index
# index2
## sub index
### sub-sub index
# index3
# index4

{quote}this is quote{quote}



{quote}multi line quote

multi line quote{quote}

----

bold: *bold*, italic: _italic_, strikethrough: -strikethrough-, underline: +underline+,

h2. URL Link

* [https://github.com/lablup/backend.ai|https://github.com/lablup/backend.ai|smart-link]
* [HyeokJin Kim: Github -> JIRA issue Migration 하려합니다.|https://teams.microsoft.com/l/message/19:dbd90d36b60545339faaeed05d6e246b@thread.skype/1733911576205?tenantId=13c6a44d-9b52-4b9e-aa34-0513ee7131f2&groupId=74ae2c4d-ec4d-4fdf-b2c2-f5041d1e8631&parentMessageId=1733911576205&teamName=devops&channelName=General&createdTime=1733911576205]

h2. table

||*abc*||*def*||*ghi*||
|1|2|3|
|4|5|6|

h2. Image (need to review)

!스크린샷 2024-12-18 오후 12.06.04.png|width=809,height=677,alt="스크린샷 2024-12-18 오후 12.06.04.png"!

h2. Check List

* 1 check
* 2 check
* 3 check
* 4 check



{code:python}def test_fn():
	print("test fn"){code}	
`

	res := servers.JiraToMD(jiraText)
	fmt.Printf("%s\n", res)
}
