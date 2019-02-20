package main

func markdownToHTML(input string) string {
	return composeAll(
		replaceByRules(strongAndEm),
		replaceByRules(strong),
	)(input)
}

func replaceByRules(rule *rule) replaceFunc {
	return func(input string) string {
		return rule.reg.ReplaceAllStringFunc(input, rule.replaceFunc)
	}
}
