func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	
	t1 := [128]byte{}
	t2 := [128]byte{}

	for i := 0; i < len(s1); i++ {
		t1[s1[i]] += 1
		t2[s2[i]] += 1
	}

	if t1 == t2 {
		return true
	}

	length := len(s1)
	for i := 0; i+length < len(s2); i++ {
		t2[s2[i]] -= 1
		t2[s2[i+length]] += 1
		if t1 == t2 {
			return true
		}
	}
	return false
}
