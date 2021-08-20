package main

func duplicate(arr []string) map[string]bool {
	ans := make(map[string]bool)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				ans[arr[i]] = true
			}
		}
	}
	return ans
}