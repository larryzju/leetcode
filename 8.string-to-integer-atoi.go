func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	
    	i := 0
	for i < len(str) && str[i] == ' ' {
		i+=1
	}
	
	if i >= len(str) {
		return 0
	}
	
	var n int64 = 0
	
	pos := true
	if str[i] == '-' {
		pos = false
		i += 1
	} else if str[i] == '+' {
		i += 1
	}
	
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		n = n*10 + int64(str[i] - '0')
		i += 1
		if pos && n > 0x7FFFFFFF {
			return 0x7FFFFFFF 
		}
		
		if !pos && -n < -0x80000000 {
			return -0x80000000
		}
	}
	

	if !pos {
		n *= -1
	}
	return int(n)
}
