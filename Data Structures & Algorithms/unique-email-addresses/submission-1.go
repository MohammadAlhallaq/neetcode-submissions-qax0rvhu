func numUniqueEmails(emails []string) int {

	result := make(map[string]struct{})

	for _, v := range emails {
		parts := strings.Split(v, "@")
		local := parts[0]
		domain := parts[1]

		if idx := strings.Index(local, "+"); idx != -1 {
			local = local[:idx]
		}


		local = strings.ReplaceAll(local, ".", "")

		normalized := local + "@" + domain
		result[normalized] = struct{}{}
	}
	return len(result)
}